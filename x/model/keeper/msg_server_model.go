package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) CreateModel(goCtx context.Context, msg *types.MsgCreateModel) (*types.MsgCreateModelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if signer has enough rights to create model
	signerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}
	if err := checkModelRights(ctx, k.Keeper, signerAddr, msg.Vid, "MsgCreateModel"); err != nil {
		return nil, err
	}

	// check if model exists
	_, isFound := k.GetModel(
		ctx,
		msg.Vid,
		msg.Pid,
	)
	if isFound {
		return nil, types.NewErrModelAlreadyExists(msg.Vid, msg.Pid)
	}

	model := types.Model{
		Creator:                                  msg.Creator,
		Vid:                                      msg.Vid,
		Pid:                                      msg.Pid,
		DeviceTypeId:                             msg.DeviceTypeId,
		ProductName:                              msg.ProductName,
		ProductLabel:                             msg.ProductLabel,
		PartNumber:                               msg.PartNumber,
		CommissioningCustomFlow:                  msg.CommissioningCustomFlow,
		CommissioningCustomFlowUrl:               msg.CommissioningCustomFlowUrl,
		CommissioningModeInitialStepsHint:        msg.CommissioningModeInitialStepsHint,
		CommissioningModeInitialStepsInstruction: msg.CommissioningModeInitialStepsInstruction,
		CommissioningModeSecondaryStepsHint:      msg.CommissioningModeSecondaryStepsHint,
		CommissioningModeSecondaryStepsInstruction: msg.CommissioningModeSecondaryStepsInstruction,
		UserManualUrl: msg.UserManualUrl,
		SupportUrl:    msg.SupportUrl,
		ProductUrl:    msg.ProductUrl,
		LsfUrl:        msg.LsfUrl,
	}

	// if LsfUrl is not empty, we set lsfRevision to default value of 1
	if model.LsfUrl != "" {
		model.LsfRevision = 1
	}

	// store new model
	k.SetModel(
		ctx,
		model,
	)

	// store new product in VendorProducts
	k.SetVendorProduct(ctx, model.Vid, types.Product{
		Pid:        model.Pid,
		Name:       model.ProductName,
		PartNumber: model.PartNumber,
	})

	return &types.MsgCreateModelResponse{}, nil
}

func (k msgServer) UpdateModel(goCtx context.Context, msg *types.MsgUpdateModel) (*types.MsgUpdateModelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if signer has enough rights to update model
	signerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}
	if err := checkModelRights(ctx, k.Keeper, signerAddr, msg.Vid, "MsgUpdateModel"); err != nil {
		return nil, err
	}

	// check if model exists
	model, isFound := k.GetModel(
		ctx,
		msg.Vid,
		msg.Pid,
	)
	if !isFound {
		return nil, types.NewErrModelDoesNotExist(msg.Vid, msg.Pid)
	}

	// update existing model value only if corresponding value in MsgUpdate is not empty

	if msg.ProductName != "" {
		model.ProductName = msg.ProductName
	}

	if msg.ProductLabel != "" {
		model.ProductLabel = msg.ProductLabel
	}

	if msg.PartNumber != "" {
		model.PartNumber = msg.PartNumber
	}

	if msg.CommissioningCustomFlowUrl != "" {
		model.CommissioningCustomFlowUrl = msg.CommissioningCustomFlowUrl
	}

	if msg.CommissioningModeInitialStepsInstruction != "" {
		model.CommissioningModeInitialStepsInstruction = msg.CommissioningModeInitialStepsInstruction
	}

	if msg.CommissioningModeSecondaryStepsInstruction != "" {
		model.CommissioningModeSecondaryStepsInstruction = msg.CommissioningModeSecondaryStepsInstruction
	}

	if msg.UserManualUrl != "" {
		model.UserManualUrl = msg.UserManualUrl
	}

	if msg.SupportUrl != "" {
		model.SupportUrl = msg.SupportUrl
	}

	if msg.ProductUrl != "" {
		model.ProductUrl = msg.ProductUrl
	}

	if msg.LsfRevision > 0 {
		// If lsfRevision is set but no lsfURL is provided or present in model
		if msg.LsfUrl == "" && model.LsfUrl == "" {
			return nil, types.NewErrLsfRevisionIsNotAllowed()
		}
		if msg.LsfRevision != model.LsfRevision+1 {
			return nil, types.NewErrLsfRevisionIsNotValid(model.LsfRevision, msg.LsfRevision)
		}
		model.LsfRevision = msg.LsfRevision
	}

	if msg.LsfUrl != "" {
		model.LsfUrl = msg.LsfUrl
		// If lsfRevision is not present, we set it to default value of 1
		if model.LsfRevision == 0 {
			model.LsfRevision = 1
		}
	}

	// store updated model
	k.SetModel(ctx, model)

	// store updated product in VendorProducts
	k.SetVendorProduct(ctx, model.Vid, types.Product{
		Pid:        model.Pid,
		Name:       model.ProductName,
		PartNumber: model.PartNumber,
	})

	return &types.MsgUpdateModelResponse{}, nil
}

func (k msgServer) DeleteModel(goCtx context.Context, msg *types.MsgDeleteModel) (*types.MsgDeleteModelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if signer has enough rights to delete model
	signerAddr, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	if err := checkModelRights(ctx, k.Keeper, signerAddr, msg.Vid, "MsgDeleteModel"); err != nil {
		return nil, err
	}

	// check if model exists
	_, isFound := k.GetModel(
		ctx,
		msg.Vid,
		msg.Pid,
	)

	if !isFound {
		return nil, types.NewErrModelDoesNotExist(msg.Vid, msg.Pid)
	}

	modelVersions, err := k.ModelVersions(goCtx, &types.QueryGetModelVersionsRequest{
		Vid: msg.Vid,
		Pid: msg.Pid,
	})

	if err != nil && status.Code(err) == codes.InvalidArgument {
		return nil, err
	}

	if modelVersions != nil {
		if err = removeAssociatedModelVersions(k, goCtx, modelVersions.ModelVersions, *msg); err != nil {
			return nil, err
		}
	}

	// remove model from store
	k.RemoveModel(
		ctx,
		msg.Vid,
		msg.Pid,
	)

	// remove product from VendorProducts
	k.RemoveVendorProduct(ctx, msg.Vid, msg.Pid)

	return &types.MsgDeleteModelResponse{}, nil
}

func removeAssociatedModelVersions(k msgServer, goCtx context.Context, modelVersions types.ModelVersions, msg types.MsgDeleteModel) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if no model version has certification record
	for _, softwareVersion := range modelVersions.SoftwareVersions {
		if k.IsModelCertified(ctx, msg.Vid, msg.Pid, softwareVersion) {
			return types.NewErrModelVersionCertified(msg.Vid, msg.Pid, softwareVersion)
		}
	}

	// remove modelVersion for each softwareVersion
	for _, softwareVersion := range modelVersions.SoftwareVersions {
		msgDeleteModelVersion := types.NewMsgDeleteModelVersion(
			msg.Creator,
			msg.Vid,
			msg.Pid,
			softwareVersion,
		)

		k.DeleteModelVersion(goCtx, msgDeleteModelVersion)
	}

	// remove modelVersions record
	k.RemoveModelVersions(ctx, msg.Vid, msg.Pid)

	return nil
}
