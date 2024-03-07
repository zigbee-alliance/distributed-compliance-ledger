package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func (k msgServer) RevokeX509Cert(goCtx context.Context, msg *types.MsgRevokeX509Cert) (*types.MsgRevokeX509CertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, pkitypes.NewErrInvalidAddress(err)
	}

	// check if signer has vendor role
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.Vendor) {
		return nil, pkitypes.NewErrUnauthorizedRole("MsgRevokeX509Cert", dclauthtypes.Vendor)
	}

	certificates, _ := k.GetApprovedCertificates(ctx, msg.Subject, msg.SubjectKeyId)
	if len(certificates.Certs) == 0 {
		return nil, pkitypes.NewErrCertificateDoesNotExist(msg.Subject, msg.SubjectKeyId)
	}

	if certificates.Certs[0].IsRoot {
		return nil, pkitypes.NewErrMessageRemoveRoot(msg.Subject, msg.SubjectKeyId)
	}

	if err := k.EnsureVidMatches(ctx, certificates.Certs[0].Owner, msg.Signer); err != nil {
		return nil, err
	}

	certIdentifier := types.CertificateIdentifier{
		Subject:      msg.Subject,
		SubjectKeyId: msg.SubjectKeyId,
	}

	if msg.SerialNumber != "" {
		certBySerialNumber, found := findCertificate(msg.SerialNumber, &certificates.Certs)
		if !found {
			return nil, pkitypes.NewErrCertificateBySerialNumberDoesNotExist(msg.Subject, msg.SubjectKeyId, msg.SerialNumber)
		}

		k._revokeX509Certificate(ctx, certBySerialNumber, certIdentifier, certificates)
	} else {
		k._revokeX509Certificates(ctx, certIdentifier, certificates)
	}

	if msg.RevokeChild {
		// Remove certificate identifier from issuer's ChildCertificates record
		k.RevokeChildCertificates(ctx, certIdentifier.Subject, certIdentifier.SubjectKeyId)
	}

	return &types.MsgRevokeX509CertResponse{}, nil
}

func (k msgServer) _revokeX509Certificates(ctx sdk.Context, certID types.CertificateIdentifier, certificates types.ApprovedCertificates) {
	// Revoke certificates with given subject/subjectKeyID
	k.AddRevokedCertificates(ctx, certificates)
	k.RemoveApprovedCertificates(ctx, certID.Subject, certID.SubjectKeyId)
	// Remove certificate identifier from issuer's ChildCertificates record
	k.RemoveChildCertificate(ctx, certificates.Certs[0].Issuer, certificates.Certs[0].AuthorityKeyId, certID)
	// remove from subject -> subject key ID map
	k.RemoveApprovedCertificateBySubject(ctx, certID.Subject, certID.SubjectKeyId)
	// remove from subject key ID -> certificates map
	k.RemoveApprovedCertificatesBySubjectKeyID(ctx, certID.Subject, certID.SubjectKeyId)
}
func (k msgServer) _revokeX509Certificate(ctx sdk.Context, cert *types.Certificate, certID types.CertificateIdentifier, certificates types.ApprovedCertificates) {
	k.AddRevokedCertificates(ctx,
		types.ApprovedCertificates{
			Subject:      cert.Subject,
			SubjectKeyId: cert.SubjectKeyId,
			Certs:        []*types.Certificate{cert},
		})
	k.removeCertFromList(cert.Issuer, cert.SerialNumber, &certificates)
	if len(certificates.Certs) == 0 {
		k.RemoveApprovedCertificates(ctx, cert.Subject, cert.SubjectKeyId)
		// Remove certificate identifier from issuer's ChildCertificates record
		k.RemoveChildCertificate(ctx, cert.Issuer, cert.AuthorityKeyId, certID)

		k.RemoveApprovedCertificateBySubject(ctx, cert.Subject, cert.SubjectKeyId)
		k.RemoveApprovedCertificatesBySubjectKeyID(ctx, cert.Subject, cert.SubjectKeyId)
	} else {
		k.SetApprovedCertificates(ctx, certificates)
		k.SetApprovedCertificatesBySubjectKeyID(
			ctx,
			types.ApprovedCertificatesBySubjectKeyId{SubjectKeyId: cert.SubjectKeyId, Certs: certificates.Certs},
		)
	}
}
