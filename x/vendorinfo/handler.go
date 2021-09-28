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

package vendorinfo

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/internal/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/internal/types"
)

func NewHandler(keeper keeper.Keeper, authKeeper auth.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgAddVendorInfo:
			return handleMsgAddVendorInfo(ctx, keeper, authKeeper, msg)
		case types.MsgUpdateVendorInfo:
			return handleMsgUpdateVendorInfo(ctx, keeper, authKeeper, msg)
			/*		case type.MsgDeleteModel:
					return handleMsgDeleteModel(ctx, keeper, authKeeper, msg)*/
		default:
			errMsg := fmt.Sprintf("unrecognized nameservice Msg type: %v", msg.Type())

			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgAddVendorInfo(ctx sdk.Context, keeper keeper.Keeper, authKeeper auth.Keeper,
	msg types.MsgAddVendorInfo) sdk.Result {
	// check if model already exists
	if keeper.IsVendorInfoPresent(ctx, msg.VendorId) {
		return types.ErrVendorInfoAlreadyExists(msg.VendorId).Result()
	}

	// check sender has enough rights to add model
	if err := checkAddVendorRights(ctx, authKeeper, msg.Signer, msg.VendorId); err != nil {
		return err.Result()
	}

	vendorInfo := types.NewVendorInfo(msg.VendorId, msg.VendorName, msg.CompanyLegalName, msg.CompanyPreferredName, msg.VendorLandingPageUrl)

	// store new vendorInfo
	keeper.SetVendorInfo(ctx, vendorInfo)

	return sdk.Result{}
}

//nolint:funlen
func handleMsgUpdateVendorInfo(ctx sdk.Context, keeper keeper.Keeper, authKeeper auth.Keeper,
	msg types.MsgUpdateVendorInfo) sdk.Result {
	// check if Vendor exists
	if !keeper.IsVendorInfoPresent(ctx, msg.VendorId) {
		return types.ErrVendorInfoDoesNotExist(msg.VendorId).Result()
	}

	vendorInfo := keeper.GetVendorInfo(ctx, msg.VendorId)

	// check if sender has enough rights to update model
	if err := checkUpdateVendorRights(ctx, authKeeper, msg.Signer, msg.VendorId); err != nil {
		return err.Result()
	}

	// updates existing vendor value only if corresponding value in MsgUpdate is not empty
	// VendorName           string `json:"vendorName"`
	// CompanyLegalName     string `json:"companyLegalName"`
	// CompanyPreferredName string `json:"companyPreferredName"`
	// VendorLandingPageUrl string `json:"vendorLandingPageUrl"`

	if msg.VendorName != "" {
		vendorInfo.VendorName = msg.VendorName
	}

	if msg.CompanyLegalName != "" {
		vendorInfo.CompanyLegalName = msg.CompanyLegalName
	}

	if msg.CompanyPreferredName != "" {
		vendorInfo.CompanyPreferredName = msg.CompanyPreferredName
	}

	if msg.VendorLandingPageUrl != "" {
		vendorInfo.VendorLandingPageUrl = msg.VendorLandingPageUrl
	}

	// store updated model
	keeper.SetVendorInfo(ctx, vendorInfo)

	return sdk.Result{}
}

func checkAddVendorRights(ctx sdk.Context, authKeeper auth.Keeper, signer sdk.AccAddress, vid uint16) sdk.Error {
	// sender must have Vendor role and VendorI to add new model
	if !authKeeper.HasRole(ctx, signer, auth.Vendor) {
		return sdk.ErrUnauthorized(fmt.Sprintf("MsgAddVendorInfo transaction should be "+
			"signed by an account with the %s role", auth.Vendor))
	}
	if !authKeeper.HasVendorId(ctx, signer, vid) {
		return sdk.ErrUnauthorized(fmt.Sprintf("MsgAddVendorInfo transaction should be "+
			"signed by an vendor account associated with the vendorId %v ", vid))
	}

	return nil
}

func checkUpdateVendorRights(ctx sdk.Context, authKeeper auth.Keeper, signer sdk.AccAddress, vid uint16) sdk.Error {
	// sender must be equal to owner to edit vendor info
	if !authKeeper.HasVendorId(ctx, signer, vid) {
		return sdk.ErrUnauthorized(fmt.Sprintf("MsgAddVendorInfo transaction should be "+
			"signed by an vendor account associated with the vendorId %v ", vid))
	}

	return nil
}
