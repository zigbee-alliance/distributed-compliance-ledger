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
		return nil, pkitypes.NewErrUnauthorizedRole("MsgAddNocX509RootCert", dclauthtypes.Vendor)
	}

	// decode pem certificate
	x509Certificate, err := x509.DecodeX509Certificate(msg.Cert)
	if err != nil {
		return nil, pkitypes.NewErrInvalidCertificate(err)
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
	certificates, found := k.GetApprovedCertificates(ctx, x509Certificate.Subject, x509Certificate.SubjectKeyID)
	if found && len(certificates.Certs) > 0 {
		existingCertificate := certificates.Certs[0]

		// Issuer and authorityKeyID must be the same as ones of exisiting certificates with the same subject and
		// subjectKeyID. Since new certificate is not self-signed, we have to ensure that the exisiting certificates
		// are not self-signed too, consequently are non-root certificates, before to match issuer and authorityKeyID.
		if existingCertificate.IsRoot || x509Certificate.Issuer != existingCertificate.Issuer ||
			x509Certificate.AuthorityKeyID != certificates.Certs[0].AuthorityKeyId {
			return nil, pkitypes.NewErrUnauthorizedCertIssuer(x509Certificate.Subject, x509Certificate.SubjectKeyID)
		}

		// Existing certificate must not be NOC certificate
		if existingCertificate.IsNoc {
			return nil, pkitypes.NewErrProvidedNotNocCertButExistingNoc(x509Certificate.Subject, x509Certificate.SubjectKeyID)
		}

		if err = k.EnsureSenderAndOwnerVidMatch(ctx, existingCertificate, msg.Signer); err != nil {
			return nil, err
		}
	}

	// Valid certificate chain must be built for new certificate
	decodedRootCert, err := k.verifyCertificate(ctx, x509Certificate)
	if err != nil {
		return nil, err
	}

	// get the full structure of the root certificate which contains the necessary fields for further validation
	approvedRootCerts, _ := k.GetApprovedCertificates(ctx, decodedRootCert.Subject, decodedRootCert.SubjectKeyID)
	if len(approvedRootCerts.Certs) == 0 {
		return nil, pkitypes.NewErrRootCertificateDoesNotExist(decodedRootCert.Subject, decodedRootCert.SubjectKeyID)
	}
	rootCert := approvedRootCerts.Certs[0]

	// Root certificate must not be NOC
	if rootCert.IsNoc {
		return nil, pkitypes.NewErrProvidedNotNocCertButExistingNoc(x509Certificate.Subject, x509Certificate.SubjectKeyID)
	}

	k.ensureCertsAndSenderVidMatch(ctx, rootCert, x509Certificate, signerAddr)

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
	)

	// append new certificate to list of certificates with the same Subject/SubjectKeyId combination and store updated list
	k.AddApprovedCertificate(ctx, certificate)

	// add the certificate identifier to the issuer's Child Certificates record
	certificateIdentifier := types.CertificateIdentifier{
		Subject:      certificate.Subject,
		SubjectKeyId: certificate.SubjectKeyId,
	}
	k.AddChildCertificate(ctx, certificate.Issuer, certificate.AuthorityKeyId, certificateIdentifier)

	// register the unique certificate key
	uniqueCertificate := types.UniqueCertificate{
		Issuer:       x509Certificate.Issuer,
		SerialNumber: x509Certificate.SerialNumber,
		Present:      true,
	}
	k.SetUniqueCertificate(ctx, uniqueCertificate)

	// add to subject -> subject key ID map
	k.AddApprovedCertificateBySubject(ctx, certificate.Subject, certificate.SubjectKeyId)

	// add to subject key ID -> certificates map
	k.AddApprovedCertificateBySubjectKeyID(ctx, certificate)

	return &types.MsgAddX509CertResponse{}, nil
}

func (k msgServer) ensureCertsAndSenderVidMatch(
	ctx sdk.Context,
	rootCert *types.Certificate,
	childCert *x509.Certificate,
	signerAddr sdk.AccAddress,
) error {
	// Check Root and Intermediate certs for VID scoping
	rootVid, err := x509.GetVidFromSubject(x509.ToSubjectAsText(rootCert.SubjectAsText))
	if err != nil {
		return pkitypes.NewErrInvalidVidFormat(err)
	}
	childVid, err := x509.GetVidFromSubject(childCert.SubjectAsText)
	if err != nil {
		return pkitypes.NewErrInvalidVidFormat(err)
	}

	signerAccount, _ := k.dclauthKeeper.GetAccountO(ctx, signerAddr)
	accountVID := signerAccount.VendorID

	if rootVid != 0 {
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
		if childVid != 0 && childVid != accountVID {
			return pkitypes.NewErrCertVidNotEqualToAccountVid(accountVID, childVid)
		}

		// Only a Vendor associated with root certificate VID can add an intermediate certificate.
		if rootCert.Vid != accountVID {
			return pkitypes.NewErrRootCertVidNotEqualToAccountVid(rootCert.Vid, accountVID)
		}
	}

	return nil
}
