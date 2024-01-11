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

	// Revoke certificates with given subject/subjectKeyID
	k.AddRevokedCertificates(ctx, certificates)
	k.RemoveApprovedCertificates(ctx, msg.Subject, msg.SubjectKeyId)

	// Remove certificate identifier from issuer's ChildCertificates record
	certIdentifier := types.CertificateIdentifier{
		Subject:      msg.Subject,
		SubjectKeyId: msg.SubjectKeyId,
	}
	k.RemoveChildCertificate(ctx, certificates.Certs[0].Issuer, certificates.Certs[0].AuthorityKeyId, certIdentifier)

	// revoke all child certificates
	k.RevokeChildCertificates(ctx, msg.Subject, msg.SubjectKeyId)

	// remove from subject -> subject key ID map
	k.RemoveApprovedCertificateBySubject(ctx, msg.Subject, msg.SubjectKeyId)

	// remove from subject key ID -> certificates map
	k.RemoveApprovedCertificatesBySubjectKeyID(ctx, msg.Subject, msg.SubjectKeyId)

	return &types.MsgRevokeX509CertResponse{}, nil
}
