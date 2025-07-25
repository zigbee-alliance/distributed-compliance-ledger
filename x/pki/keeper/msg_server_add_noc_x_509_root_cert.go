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

	// decode pem certificate
	x509Certificate, err := x509.DecodeX509Certificate(msg.Cert)
	if err != nil {
		return nil, pkitypes.NewErrInvalidCertificate(err)
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

	msgCertType := types.CertificateType_OperationalPKI
	if msg.IsVidVerificationSigner {
		msgCertType = types.CertificateType_VIDSignerPKI
	}

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
