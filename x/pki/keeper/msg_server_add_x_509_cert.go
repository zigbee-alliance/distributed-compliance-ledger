package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/x509"
)

func (k msgServer) AddX509Cert(goCtx context.Context, msg *types.MsgAddX509Cert) (*types.MsgAddX509CertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, pkitypes.NewErrInvalidAddress(err)
	}

	// check if signer has vendor role
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.Vendor) {
		return nil, pkitypes.NewErrUnauthorizedRole("MsgAddX509Cert", dclauthtypes.Vendor)
	}

	// Decode pem certificate. This handler accepts both Matter PAIs (cA=TRUE) and Matter
	// DACs (cA=FALSE); VerifyDAChainNonRoot dispatches by the BasicConstraints cA flag
	// and enforces the Matter R1.6 §6.2.2.4 PAI profile for ICAs and the §6.2.2.3 DAC
	// profile for end-entities. VerifyECDSAP256SHA256 enforces the §6.2.2.3/4 ecdsa-
	// with-SHA256 + prime256v1 algorithm requirement; VerifyVersionV3 enforces v3.
	x509Certificate, err := x509.ParseAndValidateCertificate(
		msg.Cert,
		x509.VerifyVersionV3,
		x509.VerifyECDSAP256SHA256,
		x509.VerifyDAChainNonRoot,
		x509.VerifyAtMostOneVIDAndPID,
	)
	if err != nil {
		return nil, err
	}

	// fail if certificate is self-signed
	if x509Certificate.IsSelfSigned() {
		return nil, pkitypes.NewErrNonRootCertificateSelfSigned()
	}

	// check if certificate with Issuer/Serial Number combination already exists
	if k.IsUniqueCertificatePresent(ctx, x509Certificate.Issuer, x509Certificate.SerialNumber) {
		return nil, pkitypes.NewErrCertificateAlreadyExists(x509Certificate.Issuer, x509Certificate.SerialNumber)
	}

	// Get list of certificates for Subject / Subject Key Id combination
	certificates, found := k.GetAllCertificates(ctx, x509Certificate.Subject, x509Certificate.SubjectKeyID)
	if found && len(certificates.Certs) > 0 {
		existingCertificate := certificates.Certs[0]

		// Issuer and authorityKeyID must be the same as ones of existing certificates with the same subject and
		// subjectKeyID. Since the new certificate is not self-signed, we have to ensure that the existing certificates
		// are not self-signed too, consequently are non-root certificates, before to match issuer and authorityKeyID.
		if existingCertificate.IsRoot || x509Certificate.Issuer != existingCertificate.Issuer ||
			x509Certificate.AuthorityKeyID != certificates.Certs[0].AuthorityKeyId {
			return nil, pkitypes.NewErrUnauthorizedCertIssuer(x509Certificate.Subject, x509Certificate.SubjectKeyID)
		}

		// Existing certificate must not be NOC certificate
		if existingCertificate.CertificateType == types.CertificateType_OperationalPKI {
			return nil, pkitypes.NewErrProvidedNotNocCertButExistingNoc(x509Certificate.Subject, x509Certificate.SubjectKeyID)
		}

		if err = k.EnsureVidMatches(ctx, existingCertificate.Owner, msg.Signer); err != nil {
			return nil, err
		}
	}

	// Valid certificate chain must be built for new certificate
	decodedRootCert, err := k.verifyCertificate(ctx, x509Certificate)
	if err != nil {
		return nil, err
	}

	// get the full structure of the root certificate which contains the necessary fields for further validation
	rootCerts, _ := k.GetAllCertificates(ctx, decodedRootCert.Subject, decodedRootCert.SubjectKeyID)
	if len(rootCerts.Certs) == 0 {
		return nil, pkitypes.NewErrRootCertificateDoesNotExist(decodedRootCert.Subject, decodedRootCert.SubjectKeyID)
	}
	rootCert := rootCerts.Certs[0]

	// Root certificate must not be NOC
	if rootCert.CertificateType == types.CertificateType_OperationalPKI {
		return nil, pkitypes.NewErrProvidedNotNocCertButRootIsNoc()
	}

	// VID of account must match to VID of root and provided child certificates
	if err = k.ensureVidMatches(ctx, rootCert, x509Certificate, signerAddr); err != nil {
		return nil, err
	}

	// Matter R1.6 §6.2.2.3 (DAC) 8a + 9a and §6.2.2.4 (PAI) 7a require the new
	// cert's subject VID/PID to match the immediate issuer's. We look up the
	// parent by (Issuer, AuthorityKeyID) — the chain was just verified, so the
	// parent is guaranteed to exist in the store. This runs after
	// ensureVidMatches so the existing root-relative authorization errors
	// surface first for back-compat with negative tests.
	if err = k.verifyImmediateIssuerVidPid(ctx, x509Certificate); err != nil {
		return nil, err
	}

	subjectVid, err := x509.GetVidFromSubject(x509Certificate.SubjectAsText)
	if err != nil {
		return nil, pkitypes.NewErrInvalidCertificate(err)
	}

	// create new certificate
	certificate := types.NewNonRootCertificate(
		msg.Cert,
		x509Certificate.Subject,
		x509Certificate.SubjectAsText,
		x509Certificate.SubjectKeyID,
		x509Certificate.SerialNumber,
		x509Certificate.Issuer,
		x509Certificate.AuthorityKeyID,
		rootCert.Subject,
		rootCert.SubjectKeyId,
		msg.Signer,
		subjectVid,
		msg.CertSchemaVersion,
	)

	// register the unique certificate key
	k.SetUniqueX509Certificate(ctx, x509Certificate)

	// store DA certificate in indexes
	k.StoreDaCertificate(ctx, certificate, false)

	return &types.MsgAddX509CertResponse{}, nil
}

// verifyImmediateIssuerVidPid enforces Matter R1.6 §6.2.2.3 8a/9a and §6.2.2.4
// 7a against the certificate's *immediate* issuer (one hop up in the chain),
// distinct from ensureVidMatches which compares against the root.
//
// The parent is looked up by (cert.Issuer, cert.AuthorityKeyID); if no parent
// is found at that exact key the chain check upstream would already have
// failed, so this returns nil and lets the upstream error speak. When a parent
// is found, the actual VID/PID equality check is delegated to
// x509.VerifyVidPidConsistency, which only enforces rules the spec actually
// states (no symmetric checks).
func (k msgServer) verifyImmediateIssuerVidPid(ctx sdk.Context, childCert *x509.Certificate) error {
	parents, found := k.GetAllCertificates(ctx, childCert.Issuer, childCert.AuthorityKeyID)
	if !found || len(parents.Certs) == 0 {
		return nil
	}
	parent := parents.Certs[0]

	// Stored SubjectAsText is already in the readable `vid=0x..` form, so this
	// re-projection is a no-op in practice — but propagate any error in case
	// the row was written before FormatOID started validating its input.
	parentSubjectAsText, err := x509.ToSubjectAsText(parent.SubjectAsText)
	if err != nil {
		return pkitypes.NewErrInvalidCertificate(err)
	}

	return x509.VerifyVidPidConsistency(childCert.SubjectAsText, parentSubjectAsText)
}

func (k msgServer) ensureVidMatches(
	ctx sdk.Context,
	rootCert *types.Certificate,
	childCert *x509.Certificate,
	signerAddr sdk.AccAddress,
) error {
	// Check Root and Intermediate certs for VID scoping
	rootSubjectAsText, err := x509.ToSubjectAsText(rootCert.SubjectAsText)
	if err != nil {
		return pkitypes.NewErrInvalidCertificate(err)
	}
	rootVid, err := x509.GetVidFromSubject(rootSubjectAsText)
	if err != nil {
		return pkitypes.NewErrInvalidVidFormat(err)
	}
	childVid, err := x509.GetVidFromSubject(childCert.SubjectAsText)
	if err != nil {
		return pkitypes.NewErrInvalidVidFormat(err)
	}

	signerAccount, _ := k.dclauthKeeper.GetAccountO(ctx, signerAddr)
	accountVID := signerAccount.VendorID

	if rootVid != 0 { //nolint:nestif
		// If added under a VID scoped root CA:
		// Child certificate must be also VID scoped to the same VID as a root one
		if rootVid != childVid {
			return pkitypes.NewErrRootCertVidNotEqualToCertVid(rootVid, childVid)
		}

		// Only a Vendor associated with root certificate's VID can add an intermediate certificate
		if rootVid != accountVID {
			return pkitypes.NewErrRootCertVidNotEqualToAccountVid(rootVid, accountVID)
		}
	} else {
		// If added under a non-VID scoped root CA:
		// Child certificate must be either VID scoped to the same VID, or non-VID scoped.
		if childVid != 0 && childVid != rootCert.Vid {
			return pkitypes.NewErrRootCertVidNotEqualToCertVid(accountVID, childVid)
		}

		// Only a Vendor associated with root certificate VID can add an intermediate certificate.
		if rootCert.Vid != accountVID {
			return pkitypes.NewErrRootCertVidNotEqualToAccountVid(rootCert.Vid, accountVID)
		}
	}

	return nil
}
