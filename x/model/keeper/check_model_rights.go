// Copyright 2020 DSR Corporation
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
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func checkModelRights(ctx sdk.Context, k Keeper, signer sdk.AccAddress, vid int32, pid int32, message string) error {
	// sender with VendorAdmin role can add or modify any model
	if k.dclauthKeeper.HasRole(ctx, signer, dclauthtypes.VendorAdmin) {
		return nil
	}
	// sender must have Vendor role to add or modify a model
	if !k.dclauthKeeper.HasRole(ctx, signer, dclauthtypes.Vendor) {
		return errors.Wrapf(sdkerrors.ErrUnauthorized, "%s transaction should be "+
			"signed by an account with the %s role", message, dclauthtypes.Vendor)
	}

	if !k.dclauthKeeper.HasVendorID(ctx, signer, vid) {
		return errors.Wrapf(sdkerrors.ErrUnauthorized, "%s transaction should be "+
			"signed by a vendor account containing the vendorID %v ", message, vid)
	}

	if !k.dclauthKeeper.HasRightsToChange(ctx, signer, pid) {
		return errors.Wrapf(sdkerrors.ErrUnauthorized, "%s transaction should be "+
			"signed by a vendor account who has rights to modify the productID %v ", message, pid)
	}

	return nil
}
