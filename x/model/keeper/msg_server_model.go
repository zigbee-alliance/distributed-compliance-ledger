package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

func (k msgServer) CreateModel(goCtx context.Context, msg *types.MsgCreateModel) (*types.MsgCreateModelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if signer has enough rights to create model
	if err := checkModelRights(ctx, k.Keeper, msg.GetSigners()[0], msg.Vid, "MsgCreateModel"); err != nil {
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

	var model = types.Model{
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
	}

	// store new model
	k.SetModel(
		ctx,
		model,
	)

	return &types.MsgCreateModelResponse{}, nil
}

func (k msgServer) UpdateModel(goCtx context.Context, msg *types.MsgUpdateModel) (*types.MsgUpdateModelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if signer has enough rights to update model
	if err := checkModelRights(ctx, k.Keeper, msg.GetSigners()[0], msg.Vid, "MsgUpdateModel"); err != nil {
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

	// store updated model
	k.SetModel(ctx, model)

	return &types.MsgUpdateModelResponse{}, nil
}

func (k msgServer) DeleteModel(goCtx context.Context, msg *types.MsgDeleteModel) (*types.MsgDeleteModelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if signer has enough rights to delete model
	if err := checkModelRights(ctx, k.Keeper, msg.GetSigners()[0], msg.Vid, "MsgDeleteModel"); err != nil {
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

	// remove model from store
	k.RemoveModel(
		ctx,
		msg.Vid,
		msg.Pid,
	)

	return &types.MsgDeleteModelResponse{}, nil
}
