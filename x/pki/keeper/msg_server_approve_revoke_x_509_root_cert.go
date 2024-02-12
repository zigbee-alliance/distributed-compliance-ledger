package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func (k msgServer) ApproveRevokeX509RootCert(goCtx context.Context, msg *types.MsgApproveRevokeX509RootCert) (*types.MsgApproveRevokeX509RootCertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if signer has root certificate approval role
	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, types.RootCertificateApprovalRole) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgApproveRevokeX509RootCert transaction should be signed by "+
				"an account with the \"%s\" role",
			types.RootCertificateApprovalRole,
		)
	}

	// get proposed certificate revocation
	revocation, found := k.GetProposedCertificateRevocation(ctx, msg.Subject, msg.SubjectKeyId, msg.SerialNumber)
	if !found {
		return nil, pkitypes.NewErrProposedCertificateRevocationDoesNotExist(msg.Subject, msg.SubjectKeyId)
	}

	// check if proposed certificate revocation already has approval form signer
	if revocation.HasApprovalFrom(signerAddr.String()) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"Certificate revocation associated with subject=%v and subjectKeyID=%v combination "+
				"already has approval from=%v",
			msg.Subject, msg.SubjectKeyId, msg.Signer,
		)
	}

	// append approval
	grant := types.Grant{
		Address: signerAddr.String(),
		Time:    msg.Time,
		Info:    msg.Info,
	}
	revocation.Approvals = append(revocation.Approvals, &grant)

	// check if proposed certificate revocation has enough approvals
	if len(revocation.Approvals) >= k.CertificateApprovalsCount(ctx, k.dclauthKeeper) { //nolint:nestif
		certificates, found := k.GetApprovedCertificates(ctx, msg.Subject, msg.SubjectKeyId)
		if !found {
			return nil, pkitypes.NewErrCertificateDoesNotExist(msg.Subject, msg.SubjectKeyId)
		}

		var certBySerialNumber *types.Certificate
		// Assign the approvals to the root certificate
		for _, cert := range certificates.Certs {
			if cert.IsRoot {
				cert.Approvals = revocation.Approvals
			}
			if msg.SerialNumber != "" && cert.SerialNumber == msg.SerialNumber {
				certBySerialNumber = cert

				break
			}
		}
		certID := types.CertificateIdentifier{
			Subject:      msg.Subject,
			SubjectKeyId: msg.SubjectKeyId,
		}
		k.RemoveProposedCertificateRevocation(ctx, msg.Subject, msg.SubjectKeyId, msg.SerialNumber)
		k.AddRevokedRootCertificate(ctx, certID)

		if certBySerialNumber != nil {
			k._removeAndRevokeBySerialNumber(ctx, certBySerialNumber, certID, certificates)
		} else {
			k._removeAndRevoke(ctx, certID, certificates)
		}
	} else {
		k.SetProposedCertificateRevocation(ctx, revocation)
	}

	return &types.MsgApproveRevokeX509RootCertResponse{}, nil
}

func (k msgServer) _removeAndRevoke(ctx sdk.Context, certID types.CertificateIdentifier, certificates types.ApprovedCertificates) {
	k.AddRevokedCertificates(ctx, certificates)
	k.RemoveApprovedCertificates(ctx, certID.Subject, certID.SubjectKeyId)
	k.RevokeChildCertificates(ctx, certID.Subject, certID.SubjectKeyId)

	// remove from root certs index, add to revoked root certs
	k.RemoveApprovedRootCertificate(ctx, certID)
	// remove from subject -> subject key ID map
	k.RemoveApprovedCertificateBySubject(ctx, certID.Subject, certID.SubjectKeyId)
	// remove from subject key ID -> certificates map
	k.RemoveApprovedCertificatesBySubjectKeyID(ctx, certID.Subject, certID.SubjectKeyId)
}
func (k msgServer) _removeAndRevokeBySerialNumber(ctx sdk.Context, cert *types.Certificate, certID types.CertificateIdentifier, certificates types.ApprovedCertificates) {
	k.AddRevokedCertificates(ctx,
		types.ApprovedCertificates{
			Subject:      cert.Subject,
			SubjectKeyId: cert.SubjectKeyId,
			Certs:        []*types.Certificate{cert},
		})
	k.removeCertFromList(cert.SerialNumber, &certificates)
	if len(certificates.Certs) == 0 {
		k.RemoveApprovedCertificates(ctx, cert.Subject, cert.SubjectKeyId)
		k.RevokeChildCertificates(ctx, cert.Subject, cert.SubjectKeyId)
		k.RemoveApprovedRootCertificate(ctx, certID)
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
