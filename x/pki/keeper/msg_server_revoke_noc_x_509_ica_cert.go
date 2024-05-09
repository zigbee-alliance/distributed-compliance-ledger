package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func (k msgServer) RevokeNocX509IcaCert(goCtx context.Context, msg *types.MsgRevokeNocX509IcaCert) (*types.MsgRevokeNocX509IcaCertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, pkitypes.NewErrInvalidAddress(err)
	}
	// check if signer has vendor role
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.Vendor) {
		return nil, pkitypes.NewErrUnauthorizedRole("MsgRevokeNocX509Cert", dclauthtypes.Vendor)
	}

	certificates, _ := k.GetApprovedCertificates(ctx, msg.Subject, msg.SubjectKeyId)
	if len(certificates.Certs) == 0 {
		return nil, pkitypes.NewErrCertificateDoesNotExist(msg.Subject, msg.SubjectKeyId)
	}

	cert := certificates.Certs[0]
	if cert.IsRoot {
		return nil, pkitypes.NewErrMessageExpectedNonRoot(msg.Subject, msg.SubjectKeyId)
	}
	// Existing certificate must be NOC certificate
	if !cert.IsNoc {
		return nil, pkitypes.NewErrProvidedNocCertButExistingNotNoc(cert.Subject, cert.SubjectKeyId)
	}

	signerAccount, _ := k.dclauthKeeper.GetAccountO(ctx, signerAddr)
	signerVid := signerAccount.VendorID
	// signer VID must be same as VID of existing certificates
	if signerVid != cert.Vid {
		return nil, pkitypes.NewErrRevokeCertVidNotEqualToAccountVid(cert.Vid, signerVid)
	}

	if msg.SerialNumber != "" {
		err = k._revokeNocCertificate(ctx, msg.SerialNumber, certificates, cert.Vid, msg.SchemaVersion)
		if err != nil {
			return nil, err
		}
	} else {
		k._revokeNocCertificates(ctx, certificates, cert.Vid, msg.SchemaVersion)
	}

	if msg.RevokeChild {
		// Remove certificate identifier from issuer's ChildCertificates record
		k.RevokeChildCertificates(ctx, msg.Subject, msg.SubjectKeyId, msg.SchemaVersion)
	}

	return &types.MsgRevokeNocX509IcaCertResponse{}, nil
}

func (k msgServer) _revokeNocCertificate(
	ctx sdk.Context,
	serialNumber string,
	certificates types.ApprovedCertificates,
	vid int32,
	schemaVersion uint32,
) error {
	cert, found := findCertificate(serialNumber, &certificates.Certs)
	if !found {
		return pkitypes.NewErrCertificateBySerialNumberDoesNotExist(
			certificates.Subject, certificates.SubjectKeyId, serialNumber,
		)
	}

	revCerts := types.ApprovedCertificates{
		Subject:      cert.Subject,
		SubjectKeyId: cert.SubjectKeyId,
		Certs:        []*types.Certificate{cert},
	}
	k.AddRevokedCertificates(ctx, revCerts, schemaVersion)

	removeCertFromList(cert.Issuer, cert.SerialNumber, &certificates.Certs)
	if len(certificates.Certs) == 0 {
		k.RemoveNocIcaCertificate(ctx, certificates.Subject, certificates.SubjectKeyId, vid)
		k.RemoveApprovedCertificates(ctx, cert.Subject, cert.SubjectKeyId)
		k.RemoveApprovedCertificateBySubject(ctx, cert.Subject, cert.SubjectKeyId)
		k.RemoveApprovedCertificatesBySubjectKeyID(ctx, cert.Subject, cert.SubjectKeyId)
	} else {
		k.RemoveNocIcaCertificateBySerialNumber(ctx, vid, cert.Subject, cert.SubjectKeyId, serialNumber)
		k.RemoveApprovedCertificatesBySubjectKeyIDAndSerialNumber(ctx, cert.Subject, cert.SubjectKeyId, serialNumber)
		k.SetApprovedCertificates(ctx, certificates)
	}

	return nil
}

func (k msgServer) _revokeNocCertificates(ctx sdk.Context, certificates types.ApprovedCertificates, vid int32, schemaVersion uint32) {
	// Add certs into revoked lists
	k.AddRevokedCertificates(ctx, certificates, schemaVersion)
	// remove cert from NOC certs list
	k.RemoveNocIcaCertificate(ctx, certificates.Subject, certificates.SubjectKeyId, vid)
	// remove cert from approved certs list
	k.RemoveApprovedCertificates(ctx, certificates.Subject, certificates.SubjectKeyId)
	// remove from subject -> subject key ID map
	k.RemoveApprovedCertificateBySubject(ctx, certificates.Subject, certificates.SubjectKeyId)
	// remove from subject key ID -> certificates map
	k.RemoveApprovedCertificatesBySubjectKeyID(ctx, certificates.Subject, certificates.SubjectKeyId)
}
