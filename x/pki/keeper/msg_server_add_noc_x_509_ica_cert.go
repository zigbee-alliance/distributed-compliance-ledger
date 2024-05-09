package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/x509"
)

func (k msgServer) AddNocX509IcaCert(goCtx context.Context, msg *types.MsgAddNocX509IcaCert) (*types.MsgAddNocX509IcaCertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, pkitypes.NewErrInvalidAddress(err)
	}

	// check if signer has vendor role
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.Vendor) {
		return nil, pkitypes.NewErrUnauthorizedRole("MsgAddNocX509IcaCert", dclauthtypes.Vendor)
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
	signerAccount, _ := k.dclauthKeeper.GetAccountO(ctx, signerAddr)
	accountVid := signerAccount.VendorID

	// Get list of certificates for Subject / Subject Key Id combination
	certificates, _ := k.GetApprovedCertificates(ctx, x509Certificate.Subject, x509Certificate.SubjectKeyID)
	if len(certificates.Certs) > 0 {
		existingCertificate := certificates.Certs[0]

		// Issuer and authorityKeyID must be the same as ones of existing certificates with the same subject and
		// subjectKeyID. Since new certificate is not self-signed, we have to ensure that the existing certificates
		// are not self-signed too, consequently are non-root certificates, before to match issuer and authorityKeyID.
		if existingCertificate.IsRoot || x509Certificate.Issuer != existingCertificate.Issuer ||
			x509Certificate.AuthorityKeyID != certificates.Certs[0].AuthorityKeyId {
			return nil, pkitypes.NewErrUnauthorizedCertIssuer(x509Certificate.Subject, x509Certificate.SubjectKeyID)
		}

		// Existing certificate must be NOC certificate
		if !existingCertificate.IsNoc {
			return nil, pkitypes.NewErrProvidedNocCertButExistingNotNoc(x509Certificate.Subject, x509Certificate.SubjectKeyID)
		}

		// signer VID must be same as VID of existing certificates
		if accountVid != existingCertificate.Vid {
			return nil, pkitypes.NewErrUnauthorizedCertVendor(existingCertificate.Vid)
		}
	}
	// Valid certificate chain must be built for new certificate
	rootCert, err := k.verifyCertificate(ctx, x509Certificate)
	if err != nil {
		return nil, err
	}
	// Check Root and Intermediate certs for VID scoping
	rootCerts, _ := k.GetApprovedCertificates(ctx, rootCert.Subject, rootCert.SubjectKeyID)
	if len(rootCerts.Certs) == 0 {
		return nil, pkitypes.NewErrRootCertificateDoesNotExist(rootCert.Subject, rootCert.SubjectKeyID)
	}
	nocRootCert := rootCerts.Certs[0]
	// Root certificate must be NOC certificate
	if !nocRootCert.IsNoc {
		return nil, pkitypes.NewErrRootOfNocCertIsNotNoc(rootCert.Subject, rootCert.SubjectKeyID)
	}
	// Check VID scoping
	if nocRootCert.Vid != accountVid {
		return nil, pkitypes.NewErrRootCertVidNotEqualToAccountVid(nocRootCert.Vid, accountVid)
	}

	// create new certificate
	certificate := types.NewNocCertificate(
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
		accountVid,
		msg.CertSchemaVersion,
	)

	// Add a NOC certificate to the list of NOC certificates with the same VID
	k.AddNocIcaCertificate(ctx, certificate)

	// append new certificate to list of certificates with the same Subject/SubjectKeyId combination and store updated list
	k.AddApprovedCertificate(ctx, certificate, msg.SchemaVersion)

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

	return &types.MsgAddNocX509IcaCertResponse{}, nil
}
