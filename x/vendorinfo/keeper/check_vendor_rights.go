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
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func checkAddVendorRights(ctx sdk.Context, k Keeper, signer sdk.AccAddress, vid int32) error {
	// sender must have Vendor role and VendorID to add new model
	if !k.dclauthKeeper.HasRole(ctx, signer, dclauthtypes.Vendor) {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("MsgAddVendorInfo transaction should be "+
			"signed by an account with the %s role", dclauthtypes.Vendor))
	}

	// Checks if the the msg creator is the same as the current owner
	if !k.dclauthKeeper.HasVendorID(ctx, signer, vid) {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("MsgAddVendorInfo transaction should be "+
			"signed by a vendor account associated with the vendorID %v ", vid))
	}

	return nil
}

func checkUpdateVendorRights(ctx sdk.Context, k Keeper, signer sdk.AccAddress, vid int32) error {
	// Checks if the the msg creator is the same as the current owner
	if !k.dclauthKeeper.HasVendorID(ctx, signer, vid) {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("MsgAddVendorInfo transaction should be "+
			"signed by a vendor account associated with the vendorID %v ", vid))
	}

	return nil
}
