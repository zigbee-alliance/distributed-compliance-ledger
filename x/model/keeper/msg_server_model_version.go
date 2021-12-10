package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

func (k msgServer) CreateModelVersion(goCtx context.Context, msg *types.MsgCreateModelVersion) (*types.MsgCreateModelVersionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetModelVersion(
		ctx,
		msg.Vid,
		msg.Pid,
		msg.SoftwareVersion,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var modelVersion = types.ModelVersion{
		Creator:                      msg.Creator,
		Vid:                          msg.Vid,
		Pid:                          msg.Pid,
		SoftwareVersion:              msg.SoftwareVersion,
		SoftwareVersionString:        msg.SoftwareVersionString,
		CdVersionNumber:              msg.CdVersionNumber,
		FirmwareDigests:              msg.FirmwareDigests,
		SoftwareVersionValid:         msg.SoftwareVersionValid,
		OtaUrl:                       msg.OtaUrl,
		OtaFileSize:                  msg.OtaFileSize,
		OtaChecksum:                  msg.OtaChecksum,
		OtaChecksumType:              msg.OtaChecksumType,
		MinApplicableSoftwareVersion: msg.MinApplicableSoftwareVersion,
		MaxApplicableSoftwareVersion: msg.MaxApplicableSoftwareVersion,
		ReleaseNotesUrl:              msg.ReleaseNotesUrl,
	}

	k.SetModelVersion(
		ctx,
		modelVersion,
	)
	return &types.MsgCreateModelVersionResponse{}, nil
}

func (k msgServer) UpdateModelVersion(goCtx context.Context, msg *types.MsgUpdateModelVersion) (*types.MsgUpdateModelVersionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	modelVersion, isFound := k.GetModelVersion(
		ctx,
		msg.Vid,
		msg.Pid,
		msg.SoftwareVersion,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != modelVersion.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	modelVersion.SoftwareVersionValid = msg.SoftwareVersionValid
	modelVersion.OtaUrl = msg.OtaUrl
	modelVersion.MinApplicableSoftwareVersion = msg.MinApplicableSoftwareVersion
	modelVersion.MaxApplicableSoftwareVersion = msg.MaxApplicableSoftwareVersion
	modelVersion.ReleaseNotesUrl = msg.ReleaseNotesUrl

	k.SetModelVersion(ctx, modelVersion)

	return &types.MsgUpdateModelVersionResponse{}, nil
}
