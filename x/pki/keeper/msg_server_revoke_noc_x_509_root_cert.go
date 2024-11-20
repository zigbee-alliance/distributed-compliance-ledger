package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func (k msgServer) RevokeNocX509RootCert(goCtx context.Context, msg *types.MsgRevokeNocX509RootCert) (*types.MsgRevokeNocX509RootCertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, pkitypes.NewErrInvalidAddress(err)
	}

	// check if signer has vendor role
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.Vendor) {
		return nil, pkitypes.NewErrUnauthorizedRole("MsgRevokeNocX509RootCert", dclauthtypes.Vendor)
	}

	certificates, _ := k.GetNocCertificates(ctx, msg.Subject, msg.SubjectKeyId)
	if len(certificates.Certs) == 0 {
		return nil, pkitypes.NewErrNocRootCertificateDoesNotExist(msg.Subject, msg.SubjectKeyId)
	}

	cert := certificates.Certs[0]
	if !cert.IsRoot {
		return nil, pkitypes.NewErrMessageExistingCertIsNotRoot(cert.Subject, cert.SubjectKeyId)
	}
	// Existing certificate must be NOC certificate
	if cert.CertificateType != types.CertificateType_OperationalPKI {
		return nil, pkitypes.NewErrProvidedNocCertButExistingNotNoc(cert.Subject, cert.SubjectKeyId)
	}

	signerAccount, _ := k.dclauthKeeper.GetAccountO(ctx, signerAddr)
	signerVid := signerAccount.VendorID
	// signer VID must be same as VID of existing certificates
	if signerVid != cert.Vid {
		return nil, pkitypes.NewErrRevokeRootCertVidNotEqualToAccountVid(cert.Vid, signerVid)
	}

	if msg.SerialNumber != "" {
		err = k.revokeNocRootCertificateBySerialNumber(ctx, msg.SerialNumber, certificates, cert.Vid)
		if err != nil {
			return nil, err
		}
	} else {
		k.revokeNocRootCertificate(ctx, certificates, cert.Vid)
	}

	if msg.RevokeChild {
		// Remove certificate identifier from issuer's ChildCertificates record
		k.RevokeNocChildCertificates(ctx, msg.Subject, msg.SubjectKeyId)
	}

	return &types.MsgRevokeNocX509RootCertResponse{}, nil
}

func (k msgServer) revokeNocRootCertificateBySerialNumber(
	ctx sdk.Context,
	serialNumber string,
	certificates types.NocCertificates,
	vid int32,
) error {
	cert, found := FindCertificateInList(serialNumber, &certificates.Certs)
	if !found {
		return pkitypes.NewErrCertificateBySerialNumberDoesNotExist(
			certificates.Subject, certificates.SubjectKeyId, serialNumber,
		)
	}

	k.AddRevokedNocRootCertificates(ctx, types.RevokedNocRootCertificates{
		Subject:      certificates.Subject,
		SubjectKeyId: certificates.SubjectKeyId,
		Certs:        []*types.Certificate{cert},
	})

	k.RemoveNocCertBySerialNumber(
		ctx,
		cert.Subject,
		cert.SubjectKeyId,
		&certificates,
		vid,
		serialNumber,
		cert.Issuer,
		true,
	)

	return nil
}

func (k msgServer) revokeNocRootCertificate(ctx sdk.Context, certificates types.NocCertificates, vid int32) {
	// Add certs into revoked lists
	k.AddRevokedNocRootCertificates(ctx, types.RevokedNocRootCertificates{
		Subject:      certificates.Subject,
		SubjectKeyId: certificates.SubjectKeyId,
		Certs:        certificates.Certs,
	})
	// Remove certificate from noc list
	k.RemoveNocCertificate(ctx, certificates.Subject, certificates.SubjectKeyId, vid, true)
}
