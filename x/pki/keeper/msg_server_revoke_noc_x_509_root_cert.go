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

	certificates, _ := k.GetApprovedCertificates(ctx, msg.Subject, msg.SubjectKeyId)
	if len(certificates.Certs) == 0 {
		return nil, pkitypes.NewErrCertificateDoesNotExist(msg.Subject, msg.SubjectKeyId)
	}

	cert := certificates.Certs[0]
	if !cert.IsRoot {
		return nil, pkitypes.NewErrMessageExistingCertIsNotRoot(cert.Subject, cert.SubjectKeyId)
	}
	// Existing certificate must be NOC certificate
	if !cert.IsNoc {
		return nil, pkitypes.NewErrProvidedNocCertButExistingNotNoc(cert.Subject, cert.SubjectKeyId)
	}

	signerAccount, _ := k.dclauthKeeper.GetAccountO(ctx, signerAddr)
	signerVid := signerAccount.VendorID
	// signer VID must be same as VID of existing certificates
	if signerVid != cert.Vid {
		return nil, pkitypes.NewErrRootCertVidNotEqualToAccountVid(cert.Vid, signerVid)
	}

	if msg.SerialNumber != "" {
		err = k._revokeNocRootCertificate(ctx, msg.SerialNumber, certificates, cert.Vid, msg.SchemaVersion)
		if err != nil {
			return nil, err
		}
	} else {
		k._revokeNocRootCertificates(ctx, certificates, cert.Vid, msg.SchemaVersion)
	}

	if msg.RevokeChild {
		certID := types.CertificateIdentifier{
			Subject:      msg.Subject,
			SubjectKeyId: msg.SubjectKeyId,
		}
		// Remove certificate identifier from issuer's ChildCertificates record
		k.RevokeChildCertificates(ctx, certID.Subject, certID.SubjectKeyId, msg.SchemaVersion)
	}

	return &types.MsgRevokeNocX509RootCertResponse{}, nil
}

func (k msgServer) _revokeNocRootCertificates(ctx sdk.Context, certificates types.ApprovedCertificates, vid int32, schemaVersion uint32) {
	// Add certs into revoked lists
	k.AddRevokedCertificates(ctx, certificates, schemaVersion)
	k.AddRevokedNocRootCertificates(ctx, types.RevokedNocRootCertificates{
		Subject:      certificates.Subject,
		SubjectKeyId: certificates.SubjectKeyId,
		Certs:        certificates.Certs,
	})

	// Remove certs from NOC and approved lists
	k.RemoveNocRootCertificate(ctx, vid, certificates.Subject, certificates.SubjectKeyId)
	k.RemoveApprovedCertificates(ctx, certificates.Subject, certificates.SubjectKeyId)
	// remove from subject -> subject key ID map
	k.RemoveApprovedCertificateBySubject(ctx, certificates.Subject, certificates.SubjectKeyId)
	// remove from subject key ID -> certificates map
	k.RemoveApprovedCertificatesBySubjectKeyID(ctx, certificates.Subject, certificates.SubjectKeyId)
	// remove from vid, subject key ID -> certificates map
	k.RemoveNocRootCertificateByVidSubjectAndSkid(ctx, vid, certificates.Subject, certificates.SubjectKeyId)
}

func (k msgServer) _revokeNocRootCertificate(
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
	revNocCerts := types.RevokedNocRootCertificates{
		Subject:      certificates.Subject,
		SubjectKeyId: certificates.SubjectKeyId,
		Certs:        []*types.Certificate{cert},
	}
	k.AddRevokedNocRootCertificates(ctx, revNocCerts)

	removeCertFromList(cert.Issuer, cert.SerialNumber, &certificates.Certs)

	if len(certificates.Certs) == 0 {
		k.RemoveNocRootCertificate(ctx, vid, certificates.Subject, certificates.SubjectKeyId)
		k.RemoveApprovedCertificates(ctx, cert.Subject, cert.SubjectKeyId)
		k.RemoveApprovedCertificateBySubject(ctx, cert.Subject, cert.SubjectKeyId)
		k.RemoveApprovedCertificatesBySubjectKeyID(ctx, cert.Subject, cert.SubjectKeyId)
	} else {
		k.RemoveNocRootCertificateBySerialNumber(ctx, vid, cert.Subject, cert.SubjectKeyId, serialNumber)
		k.SetApprovedCertificates(ctx, certificates)
		k.RemoveApprovedCertificatesBySubjectKeyIDAndSerialNumber(ctx, cert.Subject, cert.SubjectKeyId, serialNumber)
	}

	// remove from vid, subject key ID -> certificates map
	k.RemoveNocRootCertificateByVidSubjectSkidAndSerialNumber(ctx, vid, cert.Subject, cert.SubjectKeyId, serialNumber)

	return nil
}
