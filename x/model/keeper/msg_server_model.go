package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

func (k msgServer) CreateModel(goCtx context.Context, msg *types.MsgCreateModel) (*types.MsgCreateModelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetModel(
		ctx,
		msg.Vid,
		msg.Pid,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
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

	k.SetModel(
		ctx,
		model,
	)
	return &types.MsgCreateModelResponse{}, nil
}

func (k msgServer) UpdateModel(goCtx context.Context, msg *types.MsgUpdateModel) (*types.MsgUpdateModelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetModel(
		ctx,
		msg.Vid,
		msg.Pid,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
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

	k.SetModel(ctx, model)

	return &types.MsgUpdateModelResponse{}, nil
}

func (k msgServer) DeleteModel(goCtx context.Context, msg *types.MsgDeleteModel) (*types.MsgDeleteModelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetModel(
		ctx,
		msg.Vid,
		msg.Pid,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveModel(
		ctx,
		msg.Vid,
		msg.Pid,
	)

	return &types.MsgDeleteModelResponse{}, nil
}
