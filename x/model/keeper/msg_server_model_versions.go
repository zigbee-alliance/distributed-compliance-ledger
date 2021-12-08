package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

func (k msgServer) CreateModelVersions(goCtx context.Context, msg *types.MsgCreateModelVersions) (*types.MsgCreateModelVersionsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetModelVersions(
		ctx,
		msg.Vid,
		msg.Pid,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var modelVersions = types.ModelVersions{
		Creator:          msg.Creator,
		Vid:              msg.Vid,
		Pid:              msg.Pid,
		SoftwareVersions: msg.SoftwareVersions,
	}

	k.SetModelVersions(
		ctx,
		modelVersions,
	)
	return &types.MsgCreateModelVersionsResponse{}, nil
}

func (k msgServer) UpdateModelVersions(goCtx context.Context, msg *types.MsgUpdateModelVersions) (*types.MsgUpdateModelVersionsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetModelVersions(
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

	var modelVersions = types.ModelVersions{
		Creator:          msg.Creator,
		Vid:              msg.Vid,
		Pid:              msg.Pid,
		SoftwareVersions: msg.SoftwareVersions,
	}

	k.SetModelVersions(ctx, modelVersions)

	return &types.MsgUpdateModelVersionsResponse{}, nil
}

func (k msgServer) DeleteModelVersions(goCtx context.Context, msg *types.MsgDeleteModelVersions) (*types.MsgDeleteModelVersionsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetModelVersions(
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

	k.RemoveModelVersions(
		ctx,
		msg.Vid,
		msg.Pid,
	)

	return &types.MsgDeleteModelVersionsResponse{}, nil
}
