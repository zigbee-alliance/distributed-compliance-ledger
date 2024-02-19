package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func (k msgServer) RevokeX509Cert(goCtx context.Context, msg *types.MsgRevokeX509Cert) (*types.MsgRevokeX509CertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	certificates, found := k.GetApprovedCertificates(ctx, msg.Subject, msg.SubjectKeyId)
	if !found {
		return nil, pkitypes.NewErrCertificateDoesNotExist(msg.Subject, msg.SubjectKeyId)
	}

	if certificates.Certs[0].IsRoot {
		return nil, pkitypes.NewErrInappropriateCertificateType(
			fmt.Sprintf("Inappropriate Certificate Type: Certificate with subject=%v and subjectKeyID=%v "+
				"is a root certificate. To propose revocation of a root certificate please use "+
				"`PROPOSE_REVOKE_X509_ROOT_CERT` transaction.", msg.Subject, msg.SubjectKeyId),
		)
	}

	if msg.Signer != certificates.Certs[0].Owner {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"Only owner can revoke certificate using `REVOKE_X509_CERT`",
		)
	}

	certIdentifier := types.CertificateIdentifier{
		Subject:      msg.Subject,
		SubjectKeyId: msg.SubjectKeyId,
	}
	var certBySerialNumber *types.Certificate

	if msg.SerialNumber != "" {
		certBySerialNumber, found = findCertificate(msg.SerialNumber, &certificates.Certs)
		if !found {
			return nil, pkitypes.NewErrCertificateBySerialNumberDoesNotExist(msg.Subject, msg.SubjectKeyId, msg.SerialNumber)
		}
	}

	if certBySerialNumber != nil {
		k._makeX509CertRevoked(ctx, certBySerialNumber, certIdentifier, certificates)
	} else {
		k._makeX509CertsRevoked(ctx, certIdentifier, certificates)
	}

	if msg.RevokeChild {
		// Remove certificate identifier from issuer's ChildCertificates record
		k.RevokeChildCertificates(ctx, certIdentifier.Subject, certIdentifier.SubjectKeyId)
	}

	return &types.MsgRevokeX509CertResponse{}, nil
}

func (k msgServer) _makeX509CertsRevoked(ctx sdk.Context, certID types.CertificateIdentifier, certificates types.ApprovedCertificates) {
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
func (k msgServer) _makeX509CertRevoked(ctx sdk.Context, cert *types.Certificate, certID types.CertificateIdentifier, certificates types.ApprovedCertificates) {
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
