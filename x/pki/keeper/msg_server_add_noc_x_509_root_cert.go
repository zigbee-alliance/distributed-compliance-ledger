package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/x509"
)

func (k msgServer) AddNocX509RootCert(goCtx context.Context, msg *types.MsgAddNocX509RootCert) (*types.MsgAddNocX509RootCertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, pkitypes.NewErrInvalidAddress(err)
	}

	// check if signer has vendor role
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.Vendor) {
		return nil, pkitypes.NewErrUnauthorizedRole("MsgAddNocX509RootCert", dclauthtypes.Vendor)
	}

	// Decode the PEM. AddNocX509RootCert handles two distinct Matter R1.6 §6.5.12
	// certificate profiles, selected by msg.IsVidVerificationSigner:
	//   - false → Matter RCAC profile (cA=TRUE, KU keyCertSign+cRLSign[+digitalSignature],
	//     no EKU, SKI+AKI present), enforced by VerifyCAExtensions + VerifyNoEKU.
	//   - true  → Matter VID Verification Signer Certificate (VVSC) profile
	//     (cA=FALSE, KU exactly digitalSignature, SKI+AKI present), enforced by
	//     VerifyVVSCExtensions. The IsSelfSigned() check below restricts this path
	//     to self-issued VVSCs registered as Operational Trust Anchors (§6.4.5.4);
	//     non-self-issued VVSCs go through AddNocX509IcaCert.
	// VerifyECDSAP256SHA256 enforces the §6.5.5/§6.5.8/§6.5.9 ecdsa-with-SHA256 +
	// prime256v1 requirement; VerifyVersionV3 enforces v3.
	options := []x509.ParseAndValidateCertificateOptions{
		x509.VerifyVersionV3,
		x509.VerifyECDSAP256SHA256,
		x509.VerifyAtMostOneVIDAndPID,
	}

	msgCertType := types.CertificateType_OperationalPKI
	if msg.IsVidVerificationSigner {
		msgCertType = types.CertificateType_VIDSignerPKI
		options = append(options, x509.VerifyVVSCExtensions)
	} else {
		options = append(options, x509.VerifyCAExtensions, x509.VerifyNoEKU)
	}

	x509Certificate, err := x509.ParseAndValidateCertificate(msg.Cert, options...)
	if err != nil {
		return nil, err
	}

	// fail if certificate is not self-signed
	if !x509Certificate.IsSelfSigned() {
		return nil, pkitypes.NewErrRootCertificateIsNotSelfSigned()
	}

	// check if certificate with Issuer/Serial Number combination already exists
	if k.IsUniqueCertificatePresent(ctx, x509Certificate.Issuer, x509Certificate.SerialNumber) {
		return nil, pkitypes.NewErrCertificateAlreadyExists(x509Certificate.Issuer, x509Certificate.SerialNumber)
	}

	// verify certificate
	_, err = k.verifyCertificate(ctx, x509Certificate)
	if err != nil {
		return nil, err
	}

	// get signer VID
	signerAccount, _ := k.dclauthKeeper.GetAccountO(ctx, signerAddr)
	signerVid := signerAccount.VendorID

	// Get list of certificates for Subject / Subject Key Id combination
	existingCertificates, found := k.GetAllCertificates(ctx, x509Certificate.Subject, x509Certificate.SubjectKeyID)
	if found && len(existingCertificates.Certs) > 0 {
		existingCertificate := existingCertificates.Certs[0]

		// Issuer and authorityKeyID must be the same as ones of exisiting certificates with the same subject and
		// subjectKeyID. Since new certificate is self-signed, we have to ensure that the exisiting certificates are
		// self-signed too, consequently are root certificates.
		if !existingCertificate.IsRoot {
			return nil, pkitypes.NewErrUnauthorizedCertIssuer(x509Certificate.Subject, x509Certificate.SubjectKeyID)
		}

		// Existing certificate must be NOC certificate
		if existingCertificate.CertificateType != msgCertType {
			return nil, pkitypes.NewErrProvidedNocCertButExistingNotNoc(x509Certificate.Subject, x509Certificate.SubjectKeyID)
		}

		// signer VID must be same as VID of existing certificates
		if signerVid != existingCertificate.Vid {
			return nil, pkitypes.NewErrUnauthorizedCertVendor(existingCertificate.Vid)
		}
	}

	// create new noc root certificate
	certificate := types.NewNocRootCertificate(
		msg.Cert,
		x509Certificate.Subject,
		x509Certificate.SubjectAsText,
		x509Certificate.SubjectKeyID,
		x509Certificate.SerialNumber,
		msg.Signer,
		signerVid,
		msg.CertSchemaVersion,
		msgCertType,
	)

	// register the unique certificate key
	k.SetUniqueX509Certificate(ctx, x509Certificate)

	// store Noc certificate in indexes
	k.StoreNocCertificate(ctx, certificate, true)

	return &types.MsgAddNocX509RootCertResponse{}, nil
}
