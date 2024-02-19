package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/x509"
)

func (k msgServer) AddX509Cert(goCtx context.Context, msg *types.MsgAddX509Cert) (*types.MsgAddX509CertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// decode pem certificate
	x509Certificate, err := x509.DecodeX509Certificate(msg.Cert)
	if err != nil {
		return nil, pkitypes.NewErrInvalidCertificate(err)
	}

	// fail if certificate is self-signed
	if x509Certificate.IsSelfSigned() {
		return nil, pkitypes.NewErrInappropriateCertificateType(
			"Inappropriate Certificate Type: Passed certificate is self-signed, " +
				"so it cannot be added to the system as a non-root certificate. " +
				"To propose adding a root certificate please use `PROPOSE_ADD_X509_ROOT_CERT` transaction.")
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

		// signer must be same as owner of existing certificates
		if msg.Signer != existingCertificate.Owner {
			return nil, pkitypes.NewErrUnauthorizedCertOwner(x509Certificate.Subject, x509Certificate.SubjectKeyID)
		}
	}

	// Valid certificate chain must be built for new certificate
	rootCert, err := k.verifyCertificate(ctx, x509Certificate)
	if err != nil {
		return nil, err
	}
	// Check Root and Intermediate certs for VID scoping
	rootVid, err := x509.GetVidFromSubject(x509.ToSubjectAsText(rootCert.SubjectAsText))
	if err != nil {
		return nil, pkitypes.NewErrInvalidVidFormat(err)
	}
	childVid, err := x509.GetVidFromSubject(x509Certificate.SubjectAsText)
	if err != nil {
		return nil, pkitypes.NewErrInvalidVidFormat(err)
	}
	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, pkitypes.NewErrInvalidAddress(err)
	}

	signerAccount, _ := k.dclauthKeeper.GetAccountO(ctx, signerAddr)
	accountVID := signerAccount.VendorID

	if rootVid != 0 {
		// If added under a VID scoped root CA: Intermediate cert must be also VID scoped to the same VID as a root one.
		// Only a Vendor associated with this VID can add an intermediate certificate. So `rootVid == childVid == accountVID`
		// condition must hold
		if rootVid != childVid || rootVid != accountVID {
			return nil, pkitypes.NewErrRootCertVidNotEqualToAccountVidOrCertVid(rootVid, accountVID, childVid)
		}
		// If added under a non-VID scoped root CA associated with a VID: Intermediate cert must be either VID scoped to the same VID, or non-VID scoped.
		// Only a Vendor associated with this VID can add an intermediate certificate.
	} else if childVid != 0 && childVid != accountVID {
		return nil, pkitypes.NewErrAccountVidNotEqualToCertVid(accountVID, childVid)
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
		rootCert.SubjectKeyID,
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
