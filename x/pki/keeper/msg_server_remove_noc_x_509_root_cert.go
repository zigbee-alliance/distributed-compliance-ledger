package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func (k msgServer) RemoveNocX509RootCert(goCtx context.Context, msg *types.MsgRemoveNocX509RootCert) (*types.MsgRemoveNocX509RootCertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, pkitypes.NewErrInvalidAddress(err)
	}

	// check if signer has vendor role
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.Vendor) {
		return nil, pkitypes.NewErrUnauthorizedRole("MsgRemoveNocX509RootCert", dclauthtypes.Vendor)
	}

	signerAccount, _ := k.dclauthKeeper.GetAccountO(ctx, signerAddr)
	accountVid := signerAccount.VendorID

	nocCerts, foundActive := k.GetNocCertificates(ctx, msg.Subject, msg.SubjectKeyId)
	revCerts, foundRevoked := k.GetRevokedNocRootCertificates(ctx, msg.Subject, msg.SubjectKeyId)
	certificates := nocCerts.Certs
	certificates = append(certificates, revCerts.Certs...)
	if len(certificates) == 0 {
		return nil, pkitypes.NewErrNocRootCertificateDoesNotExist(msg.Subject, msg.SubjectKeyId)
	}

	cert := certificates[0]
	// Existing certificate must be Root certificate
	if !cert.IsRoot {
		return nil, pkitypes.NewErrMessageExistingCertIsNotRoot(cert.Subject, cert.SubjectKeyId)
	}

	// Existing certificate must be NOC certificate
	if cert.CertificateType != types.CertificateType_OperationalPKI {
		return nil, pkitypes.NewErrProvidedNocCertButExistingNotNoc(msg.Subject, msg.SubjectKeyId)
	}

	// account VID must be same as VID of existing certificates
	if accountVid != cert.Vid {
		return nil, pkitypes.NewErrRevokeCertVidNotEqualToAccountVid(cert.Vid, accountVid)
	}

	certID := types.CertificateIdentifier{
		Subject:      msg.Subject,
		SubjectKeyId: msg.SubjectKeyId,
	}

	if msg.SerialNumber != "" {
		certBySerialNumber, found := FindCertificateInList(msg.SerialNumber, &certificates)
		if !found {
			return nil, pkitypes.NewErrCertificateBySerialNumberDoesNotExist(msg.Subject, msg.SubjectKeyId, msg.SerialNumber)
		}

		// remove from subject with serialNumber map
		k.RemoveUniqueCertificate(ctx, certBySerialNumber.Subject, certBySerialNumber.SerialNumber)

		if foundActive {
			// Remove from lists
			k.RemoveNocCertBySerialNumber(
				ctx,
				certBySerialNumber.Subject,
				certBySerialNumber.SubjectKeyId,
				&nocCerts,
				accountVid,
				msg.SerialNumber,
				cert.Issuer,
				true,
			)
		}

		if foundRevoked {
			RemoveCertFromList(certBySerialNumber.Issuer, certBySerialNumber.SerialNumber, &revCerts.Certs)
			k.removeRevokedNocX509RootCert(ctx, certID, &revCerts)
		}
	} else {
		// remove from revoked noc root certs
		k.RemoveRevokedNocRootCertificates(ctx, certID.Subject, certID.SubjectKeyId)
		// remove from noc certificates map
		k.RemoveNocCertificate(ctx, cert.Subject, cert.SubjectKeyId, accountVid, true)
		// remove from subject with serialNumber map
		for _, cert := range certificates {
			k.RemoveUniqueCertificate(ctx, cert.Subject, cert.SerialNumber)
		}
	}

	return &types.MsgRemoveNocX509RootCertResponse{}, nil
}

func (k msgServer) removeRevokedNocX509RootCert(ctx sdk.Context, certID types.CertificateIdentifier, certificates *types.RevokedNocRootCertificates) {
	if len(certificates.Certs) == 0 {
		k.RemoveRevokedNocRootCertificates(ctx, certID.Subject, certID.SubjectKeyId)
	} else {
		k.SetRevokedNocRootCertificates(ctx, *certificates)
	}
}
