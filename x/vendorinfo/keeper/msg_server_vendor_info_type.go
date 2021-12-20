package keeper

import (
	"context"

    "github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)


func (k msgServer) CreateVendorInfoType(goCtx context.Context,  msg *types.MsgCreateVendorInfoType) (*types.MsgCreateVendorInfoTypeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    // Check if the value already exists
    _, isFound := k.GetVendorInfoType(
        ctx,
        msg.Index,
        )
    if isFound {
        return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
    }

    var vendorInfoType = types.VendorInfoType{
        Creator: msg.Creator,
        Index: msg.Index,
        VendorID: msg.VendorID,
        VendorName: msg.VendorName,
        CompanyLegalName: msg.CompanyLegalName,
        CompanyPrefferedName: msg.CompanyPrefferedName,
        VendorLandingPageURL: msg.VendorLandingPageURL,
        
    }

   k.SetVendorInfoType(
   		ctx,
   		vendorInfoType,
   	)
	return &types.MsgCreateVendorInfoTypeResponse{}, nil
}

func (k msgServer) UpdateVendorInfoType(goCtx context.Context,  msg *types.MsgUpdateVendorInfoType) (*types.MsgUpdateVendorInfoTypeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    // Check if the value exists
    valFound, isFound := k.GetVendorInfoType(
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

    var vendorInfoType = types.VendorInfoType{
		Creator: msg.Creator,
		Index: msg.Index,
        VendorID: msg.VendorID,
		VendorName: msg.VendorName,
		CompanyLegalName: msg.CompanyLegalName,
		CompanyPrefferedName: msg.CompanyPrefferedName,
		VendorLandingPageURL: msg.VendorLandingPageURL,
		
	}

	k.SetVendorInfoType(ctx, vendorInfoType)

	return &types.MsgUpdateVendorInfoTypeResponse{}, nil
}

func (k msgServer) DeleteVendorInfoType(goCtx context.Context,  msg *types.MsgDeleteVendorInfoType) (*types.MsgDeleteVendorInfoTypeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    // Check if the value exists
    valFound, isFound := k.GetVendorInfoType(
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

	k.RemoveVendorInfoType(
	    ctx,
	msg.Index,
    )

	return &types.MsgDeleteVendorInfoTypeResponse{}, nil
}