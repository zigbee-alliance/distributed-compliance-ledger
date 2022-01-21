package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

func (k msgServer) CreateModelVersion(goCtx context.Context, msg *types.MsgCreateModelVersion) (*types.MsgCreateModelVersionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check signer has enough rights to create model version
	signerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}
	if err := checkModelRights(ctx, k.Keeper, signerAddr, msg.Vid, "MsgCreateModelVersion"); err != nil {
		return nil, err
	}

	// check if model exists
	_, isModelFound := k.GetModel(
		ctx,
		msg.Vid,
		msg.Pid,
	)
	if !isModelFound {
		return nil, types.NewErrModelDoesNotExist(msg.Vid, msg.Pid)
	}

	// check if model version exists
	_, isModelVersionFound := k.GetModelVersion(
		ctx,
		msg.Vid,
		msg.Pid,
		msg.SoftwareVersion,
	)
	if isModelVersionFound {
		return nil, types.NewErrModelVersionAlreadyExists(msg.Vid, msg.Pid, msg.SoftwareVersion)
	}

	// check if maxApllicableSoftwareVersion is less then minApplicableSoftwareVersion
	if msg.MaxApplicableSoftwareVersion < msg.MinApplicableSoftwareVersion {
		return nil, types.NewErrMaxSVLessThanMinSV(msg.MinApplicableSoftwareVersion, msg.MaxApplicableSoftwareVersion)
	}

	modelVersion := types.ModelVersion{
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

	// store new model version
	k.SetModelVersion(
		ctx,
		modelVersion,
	)

	// add model version to a list of all model versions for this vid/pid
	k.AddModelVersion(ctx, msg.Vid, msg.Pid, msg.SoftwareVersion)

	return &types.MsgCreateModelVersionResponse{}, nil
}

func (k msgServer) UpdateModelVersion(goCtx context.Context, msg *types.MsgUpdateModelVersion) (*types.MsgUpdateModelVersionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check signer has enough rights to update model version
	signerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}
	if err := checkModelRights(ctx, k.Keeper, signerAddr, msg.Vid, "MsgCreateModelVersion"); err != nil {
		return nil, err
	}

	// check if model version exists
	modelVersion, isFound := k.GetModelVersion(
		ctx,
		msg.Vid,
		msg.Pid,
		msg.SoftwareVersion,
	)
	if !isFound {
		return nil, types.NewErrModelVersionDoesNotExist(msg.Vid, msg.Pid, msg.SoftwareVersion)
	}

	// Only OtaURL is modifiable field per specs. This can only be modified if this was set initially
	// as otaFileSize, otaChecksum and otaChecksumType are non mutable fields
	if msg.OtaUrl != "" && modelVersion.OtaUrl == "" {
		return nil, types.NewErrOtaURLCannotBeSet(msg.Vid, msg.Pid, msg.SoftwareVersion)
	}

	// update existing model version value only if corresponding value in MsgUpdate is not empty

	// SoftwareVersionValid flag is updated in any case. So pass the existing value to keep it unchanged.
	modelVersion.SoftwareVersionValid = msg.SoftwareVersionValid

	if msg.OtaUrl != "" {
		modelVersion.OtaUrl = msg.OtaUrl
	}

	if msg.MinApplicableSoftwareVersion != 0 {
		modelVersion.MinApplicableSoftwareVersion = msg.MinApplicableSoftwareVersion
	}

	if msg.MaxApplicableSoftwareVersion != 0 {
		modelVersion.MaxApplicableSoftwareVersion = msg.MaxApplicableSoftwareVersion
	}

	if msg.ReleaseNotesUrl != "" {
		modelVersion.ReleaseNotesUrl = msg.ReleaseNotesUrl
	}

	// store updated model version
	k.SetModelVersion(ctx, modelVersion)

	return &types.MsgUpdateModelVersionResponse{}, nil
}
