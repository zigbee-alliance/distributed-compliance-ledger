package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
)

func (k msgServer) CreateNewVendorInfo(goCtx context.Context, msg *types.MsgCreateNewVendorInfo) (*types.MsgCreateNewVendorInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetNewVendorInfo(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var newVendorInfo = types.NewVendorInfo{
		Creator:    msg.Creator,
		Index:      msg.Index,
		VendorInfo: msg.VendorInfo,
	}

	k.SetNewVendorInfo(
		ctx,
		newVendorInfo,
	)
	return &types.MsgCreateNewVendorInfoResponse{}, nil
}

func (k msgServer) UpdateNewVendorInfo(goCtx context.Context, msg *types.MsgUpdateNewVendorInfo) (*types.MsgUpdateNewVendorInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetNewVendorInfo(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var newVendorInfo = types.NewVendorInfo{
		Creator:    msg.Creator,
		Index:      msg.Index,
		VendorInfo: msg.VendorInfo,
	}

	k.SetNewVendorInfo(ctx, newVendorInfo)

	return &types.MsgUpdateNewVendorInfoResponse{}, nil
}

func (k msgServer) DeleteNewVendorInfo(goCtx context.Context, msg *types.MsgDeleteNewVendorInfo) (*types.MsgDeleteNewVendorInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetNewVendorInfo(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveNewVendorInfo(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteNewVendorInfoResponse{}, nil
}
