// Copyright 2022 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

// func (k msgServer) DeleteVendorInfo(goCtx context.Context, msg *types.MsgDeleteVendorInfo) (*types.MsgDeleteVendorInfoResponse, error) {
// 	ctx := sdk.UnwrapSDKContext(goCtx)

// 	// Check if the value exists
// 	valFound, isFound := k.GetVendorInfo(
// 		ctx,
// 		msg.VendorID,
// 	)
// 	if !isFound {
// 		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
// 	}

// 	// Checks if the the msg creator is the same as the current owner
// 	if msg.Creator != valFound.Creator {
// 		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
// 	}

// 	k.RemoveVendorInfo(
// 		ctx,
// 		msg.VendorID,
// 	)

// 	return &types.MsgDeleteVendorInfoResponse{}, nil
// }
