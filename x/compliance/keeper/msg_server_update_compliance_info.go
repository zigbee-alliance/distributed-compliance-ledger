package keeper

import (
	"context"
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func (k msgServer) UpdateComplianceInfo(goCtx context.Context, msg *types.MsgUpdateComplianceInfo) (*types.MsgUpdateComplianceInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if signer has enough rights to update model
	signerAddr, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.CertificationCenter) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s transaction should be "+
			"signed by an account with the %s role", "MsgUpdateComplianceInfo", dclauthtypes.CertificationCenter)
	}

	complianceInfo, isFound := k.GetComplianceInfo(ctx, msg.Vid, msg.Pid, msg.SoftwareVersion, msg.CertificationType)

	if !isFound {
		return nil, types.NewErrComplianceInfoDoesNotExist(msg.Vid, msg.Pid, msg.SoftwareVersion, msg.CertificationType)
	}

	if msg.CDCertificateId != "" {
		complianceInfo.CDCertificateId = msg.CDCertificateId
	}

	if msg.CDVersionNumber != math.MaxUint32 {
		complianceInfo.CDVersionNumber = msg.CDVersionNumber
	}

	if msg.CertificationIdOfSoftwareComponent != "" {
		complianceInfo.CertificationIdOfSoftwareComponent = msg.CertificationIdOfSoftwareComponent
	}

	if msg.CertificationRoute != "" {
		complianceInfo.CertificationRoute = msg.CertificationRoute
	}

	if msg.CompliantPlatformUsed != "" {
		complianceInfo.CompliantPlatformUsed = msg.CompliantPlatformUsed
	}

	if msg.CompliantPlatformVersion != "" {
		complianceInfo.CompliantPlatformVersion = msg.CompliantPlatformVersion
	}

	if msg.Date != "" {
		complianceInfo.Date = msg.Date
	}

	if msg.FamilyId != "" {
		complianceInfo.FamilyId = msg.FamilyId
	}

	if msg.OSVersion != "" {
		complianceInfo.OSVersion = msg.OSVersion
	}

	if msg.ParentChild != "" {
		complianceInfo.ParentChild = msg.ParentChild
	}

	if msg.ProgramType != "" {
		complianceInfo.ProgramType = msg.ProgramType
	}

	if msg.ProgramTypeVersion != "" {
		complianceInfo.ProgramTypeVersion = msg.ProgramTypeVersion
	}

	if msg.Reason != "" {
		complianceInfo.Reason = msg.Reason
	}

	if msg.SupportedClusters != "" {
		complianceInfo.SupportedClusters = msg.SupportedClusters
	}

	if msg.Transport != "" {
		complianceInfo.Transport = msg.Transport
	}

	k.SetComplianceInfo(ctx, complianceInfo)

	return &types.MsgUpdateComplianceInfoResponse{}, nil
}
