package keeper

import (
	"context"

	"cosmossdk.io/errors"
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
		return nil, errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, types.RootCertificateApprovalRole) {
		return nil, errors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgApproveRevokeX509RootCert transaction should be signed by "+
				"an account with the \"%s\" role",
			types.RootCertificateApprovalRole,
		)
	}

	// get proposed certificate revocation
	revocation, found := k.GetProposedCertificateRevocation(ctx, msg.Subject, msg.SubjectKeyId, msg.SerialNumber)
	if !found {
		return nil, pkitypes.NewErrProposedCertificateRevocationDoesNotExist(msg.Subject, msg.SubjectKeyId, msg.SerialNumber)
	}

	// check if proposed certificate revocation already has approval form signer
	if revocation.HasApprovalFrom(signerAddr.String()) {
		return nil, errors.Wrapf(sdkerrors.ErrUnauthorized,
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

		certID := types.CertificateIdentifier{
			Subject:      msg.Subject,
			SubjectKeyId: msg.SubjectKeyId,
		}
		k.AddRevokedRootCertificate(ctx, certID)

		k.RemoveProposedCertificateRevocation(ctx, msg.Subject, msg.SubjectKeyId, msg.SerialNumber)

		if msg.SerialNumber != "" {
			err := k.revokeRootCertificateBySerialNumber(ctx, revocation.Approvals, msg.SerialNumber, certificates)
			if err != nil {
				return nil, err
			}
		} else {
			k.revokeRootCertificate(ctx, revocation.Approvals, certificates, revocation.SchemaVersion)
		}

		if revocation.RevokeChild {
			k.RevokeApprovedChildCertificates(ctx, certID.Subject, certID.SubjectKeyId)
		}
	} else {
		k.SetProposedCertificateRevocation(ctx, revocation)
	}

	return &types.MsgApproveRevokeX509RootCertResponse{}, nil
}

func (k msgServer) revokeRootCertificate(
	ctx sdk.Context,
	approvals []*types.Grant,
	certificates types.ApprovedCertificates,
	schemaVersion uint32,
) {
	// Assign the approvals to the root certificate
	for _, cert := range certificates.Certs {
		if cert.IsRoot {
			cert.Approvals = approvals
		}
	}

	// remove from root certs index, add to revoked root certs
	k.AddRevokedCertificates(ctx, types.RevokedCertificates(certificates))
	// remove certificate from da list
	k.RemoveDaCertificate(ctx, certificates.Subject, certificates.SubjectKeyId, true)
}

func (k msgServer) revokeRootCertificateBySerialNumber(
	ctx sdk.Context,
	approvals []*types.Grant,
	serialNumber string,
	certificates types.ApprovedCertificates,
) error {
	cert, found := FindCertificateInList(serialNumber, &certificates.Certs)
	if !found {
		return pkitypes.NewErrCertificateBySerialNumberDoesNotExist(certificates.Subject, certificates.SubjectKeyId, serialNumber)
	}
	cert.Approvals = approvals
	revCert := types.RevokedCertificates{
		Subject:       cert.Subject,
		SubjectKeyId:  cert.SubjectKeyId,
		Certs:         []*types.Certificate{cert},
		SchemaVersion: cert.SchemaVersion,
	}

	// remove from root certs index, add to revoked root certs
	k.AddRevokedCertificates(ctx, revCert)

	// remove from certificate indexes
	k.RemoveDaCertificateBySerialNumber(
		ctx,
		cert.Subject,
		cert.SubjectKeyId,
		&certificates,
		cert.SerialNumber,
		cert.Issuer,
	)

	return nil
}
