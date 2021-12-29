package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
)

func (k msgServer) CreateVendorInfo(goCtx context.Context, msg *types.MsgCreateVendorInfo) (*types.MsgCreateVendorInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetVendorInfo(
		ctx,
		msg.VendorID,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	// check if creator has enough rights to create vendorinfo
	if err := checkAddVendorRights(ctx, k.Keeper, msg.GetSigners()[0], msg.VendorID); err != nil {
		return nil, err
	}

	vendorInfo := types.VendorInfo{
		Creator:              msg.Creator,
		VendorID:             msg.VendorID,
		VendorName:           msg.VendorName,
		CompanyLegalName:     msg.CompanyLegalName,
		CompanyPrefferedName: msg.CompanyPrefferedName,
		VendorLandingPageURL: msg.VendorLandingPageURL,
	}

	k.SetVendorInfo(
		ctx,
		vendorInfo,
	)
	return &types.MsgCreateVendorInfoResponse{}, nil
}

func (k msgServer) UpdateVendorInfo(goCtx context.Context, msg *types.MsgUpdateVendorInfo) (*types.MsgUpdateVendorInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	_, isFound := k.GetVendorInfo(
		ctx,
		msg.VendorID,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// check if creator has enough rights to update vendorinfo
	if err := checkUpdateVendorRights(ctx, k.Keeper, msg.GetSigners()[0], msg.VendorID); err != nil {
		return nil, err
	}

	vendorInfo := types.VendorInfo{
		Creator:              msg.Creator,
		VendorID:             msg.VendorID,
		VendorName:           msg.VendorName,
		CompanyLegalName:     msg.CompanyLegalName,
		CompanyPrefferedName: msg.CompanyPrefferedName,
		VendorLandingPageURL: msg.VendorLandingPageURL,
	}

	k.SetVendorInfo(ctx, vendorInfo)

	return &types.MsgUpdateVendorInfoResponse{}, nil
}

func (k msgServer) DeleteVendorInfo(goCtx context.Context, msg *types.MsgDeleteVendorInfo) (*types.MsgDeleteVendorInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetVendorInfo(
		ctx,
		msg.VendorID,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveVendorInfo(
		ctx,
		msg.VendorID,
	)

	return &types.MsgDeleteVendorInfoResponse{}, nil
}
