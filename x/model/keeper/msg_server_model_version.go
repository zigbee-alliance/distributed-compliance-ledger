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

	modelVersion := types.ModelVersion{
		Creator:                      msg.Creator,
		Vid:                          msg.Vid,
		Pid:                          msg.Pid,
		SoftwareVersion:              msg.SoftwareVersion,
		SoftwareVersionString:        msg.SoftwareVersionString,
		CdVersionNumber:              msg.CdVersionNumber,
		FirmwareInformation:          msg.FirmwareInformation,
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
	if err := checkModelRights(ctx, k.Keeper, signerAddr, msg.Vid, "MsgUpdateModelVersion"); err != nil {
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

	if msg.MinApplicableSoftwareVersion != 0 && msg.MaxApplicableSoftwareVersion == 0 && msg.MinApplicableSoftwareVersion > modelVersion.MaxApplicableSoftwareVersion {
		return nil, types.NewErrMaxSVLessThanMinSV(msg.MinApplicableSoftwareVersion, modelVersion.MaxApplicableSoftwareVersion)
	}

	if msg.MinApplicableSoftwareVersion == 0 && msg.MaxApplicableSoftwareVersion != 0 &&
		msg.MaxApplicableSoftwareVersion < modelVersion.MinApplicableSoftwareVersion {
		return nil, types.NewErrMaxSVLessThanMinSV(modelVersion.MinApplicableSoftwareVersion, msg.MaxApplicableSoftwareVersion)
	}

	if msg.OtaUrl != "" && modelVersion.OtaUrl != "" && msg.OtaUrl != modelVersion.OtaUrl {
		return nil, types.NewErrOtaURLCannotBeSet(modelVersion.Vid, modelVersion.Pid, modelVersion.SoftwareVersion)
	}

	if msg.OtaFileSize != 0 && modelVersion.OtaFileSize != 0 && msg.OtaFileSize != modelVersion.OtaFileSize {
		return nil, types.NewErrOtaFileSizeCannotBeSet(modelVersion.Vid, modelVersion.Pid, modelVersion.SoftwareVersion)
	}

	if msg.OtaChecksum != "" && modelVersion.OtaChecksum != "" && msg.OtaChecksum != modelVersion.OtaChecksum {
		return nil, types.NewErrOtaChecksumCannotBeSet(modelVersion.Vid, modelVersion.Pid, modelVersion.SoftwareVersion)
	}

	// update existing model version value only if corresponding value in MsgUpdate is not empty

	if msg.OtaUrl != "" && modelVersion.OtaUrl == "" {
		modelVersion.OtaUrl = msg.OtaUrl
	}

	if msg.OtaFileSize != 0 && modelVersion.OtaFileSize == 0 {
		modelVersion.OtaFileSize = msg.OtaFileSize
	}

	if msg.OtaChecksum != "" && modelVersion.OtaChecksum == "" {
		modelVersion.OtaChecksum = msg.OtaChecksum
	}

	// SoftwareVersionValid flag is updated in any case. So pass the existing value to keep it unchanged.
	modelVersion.SoftwareVersionValid = msg.SoftwareVersionValid

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

func (k msgServer) DeleteModelVersion(goCtx context.Context, msg *types.MsgDeleteModelVersion) (*types.MsgDeleteModelVersionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check signer has enough rights to delete model version
	signerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}
	if err := checkModelRights(ctx, k.Keeper, signerAddr, msg.Vid, "MsgDeleteModelVersion"); err != nil {
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

	if msg.Creator != modelVersion.Creator {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "MsgDeleteModelVersion transaction should be "+
			"signed by a modelVersion %d %d %d creator", msg.Vid, msg.Pid, msg.SoftwareVersion)
	}

	isCertified := k.IsModelCertified(ctx, msg.Vid, msg.Pid, msg.SoftwareVersion)
	if isCertified {
		return nil, types.NewErrModelVersionCertified(msg.Vid, msg.Pid, modelVersion.SoftwareVersion)
	}

	// store updated model version
	k.RemoveModelVersion(ctx, msg.Vid, msg.Pid, msg.SoftwareVersion)

	return &types.MsgDeleteModelVersionResponse{}, nil
}

func (k msgServer) IsModelCertified(ctx sdk.Context, vid int32, pid int32, softwareVersion uint32) bool {
	certificationTypes := []string{"zigbee", "matter"}
	for _, certType := range certificationTypes {
		_, isFound := k.complianceKeeper.GetComplianceInfo(ctx, vid, pid, softwareVersion, certType)
		if isFound {
			return true
		}
	}

	return false
}
