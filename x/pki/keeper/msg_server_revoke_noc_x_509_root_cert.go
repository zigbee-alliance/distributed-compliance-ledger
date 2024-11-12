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
		err = k._revokeNocRootCertificateBySerialNumber(ctx, msg.SerialNumber, certificates, cert.Vid)
		if err != nil {
			return nil, err
		}
	} else {
		k._revokeNocRootCertificates(ctx, certificates, cert.Vid)
	}

	if msg.RevokeChild {
		certID := types.CertificateIdentifier{
			Subject:      msg.Subject,
			SubjectKeyId: msg.SubjectKeyId,
		}
		// Remove certificate identifier from issuer's ChildCertificates record
		k.RevokeNocChildCertificates(ctx, certID.Subject, certID.SubjectKeyId)
	}

	return &types.MsgRevokeNocX509RootCertResponse{}, nil
}

func (k msgServer) _revokeNocRootCertificateBySerialNumber(
	ctx sdk.Context,
	serialNumber string,
	certificates types.NocCertificates,
	vid int32,
) error {
	cert, found := findCertificate(serialNumber, &certificates.Certs)
	if !found {
		return pkitypes.NewErrCertificateBySerialNumberDoesNotExist(
			certificates.Subject, certificates.SubjectKeyId, serialNumber,
		)
	}

	revNocCerts := types.RevokedNocRootCertificates{
		Subject:      certificates.Subject,
		SubjectKeyId: certificates.SubjectKeyId,
		Certs:        []*types.Certificate{cert},
	}
	k.AddRevokedNocRootCertificates(ctx, revNocCerts)

	removeCertFromList(cert.Issuer, cert.SerialNumber, &certificates.Certs)

	certID := types.CertificateIdentifier{
		Subject:      cert.Subject,
		SubjectKeyId: cert.SubjectKeyId,
	}

	if len(certificates.Certs) == 0 {
		// Remove certificate from global list
		k.RemoveCertificateFromAllCertificateIndexes(ctx, certID)
		// Remove certificate from noc list
		k.RemoveCertificateFromNocCertificateIndexes(ctx, certID, vid, true)
	} else {
		k.RemoveAllCertificatesBySerialNumber(ctx, cert.Subject, cert.SubjectKeyId, serialNumber)
		k.RemoveNocCertificatesBySerialNumber(ctx, cert.Subject, cert.SubjectKeyId, serialNumber)
		k.RemoveNocRootCertificateBySerialNumber(ctx, cert.Subject, cert.SubjectKeyId, vid, serialNumber)
		k.RemoveNocCertificatesByVidAndSkidBySerialNumber(ctx, vid, cert.Subject, cert.SubjectKeyId, serialNumber)
		k.RemoveNocCertificatesBySubjectKeyIDBySerialNumber(ctx, cert.Subject, cert.SubjectKeyId, serialNumber)
	}

	return nil
}

func (k msgServer) _revokeNocRootCertificates(ctx sdk.Context, certificates types.NocCertificates, vid int32) {
	// Add certs into revoked lists
	k.AddRevokedNocRootCertificates(ctx, types.RevokedNocRootCertificates{
		Subject:      certificates.Subject,
		SubjectKeyId: certificates.SubjectKeyId,
		Certs:        certificates.Certs,
	})

	certID := types.CertificateIdentifier{
		Subject:      certificates.Subject,
		SubjectKeyId: certificates.SubjectKeyId,
	}
	// Remove certificate from global list
	k.RemoveCertificateFromAllCertificateIndexes(ctx, certID)
	// Remove certificate from noc list
	k.RemoveCertificateFromNocCertificateIndexes(ctx, certID, vid, true)
}
