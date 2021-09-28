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

package model

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/internal/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/internal/types"
)

func NewHandler(keeper keeper.Keeper, authKeeper auth.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgAddModel:
			return handleMsgAddModel(ctx, keeper, authKeeper, msg)
		case types.MsgUpdateModel:
			return handleMsgUpdateModel(ctx, keeper, authKeeper, msg)
			/*		case type.MsgDeleteModel:
					return handleMsgDeleteModel(ctx, keeper, authKeeper, msg)*/
		default:
			errMsg := fmt.Sprintf("unrecognized nameservice Msg type: %v", msg.Type())

			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgAddModel(ctx sdk.Context, keeper keeper.Keeper, authKeeper auth.Keeper,
	msg types.MsgAddModel) sdk.Result {
	// check if model already exists
	if keeper.IsModelPresent(ctx, msg.VID, msg.PID) {
		return types.ErrModelAlreadyExists(msg.VID, msg.PID).Result()
	}

	// check sender has enough rights to add model
	if err := checkModelRights(ctx, authKeeper, msg.Signer, msg.VID, "msgAddModel"); err != nil {
		return err.Result()
	}

	model := types.Model{
		VID:                                      msg.VID,
		PID:                                      msg.PID,
		DeviceTypeID:                             msg.DeviceTypeID,
		ProductName:                              msg.ProductName,
		ProductLabel:                             msg.ProductLabel,
		PartNumber:                               msg.PartNumber,
		CommissioningCustomFlow:                  msg.CommissioningCustomFlow,
		CommissioningCustomFlowURL:               msg.CommissioningCustomFlowURL,
		CommissioningModeInitialStepsHint:        msg.CommissioningModeInitialStepsHint,
		CommissioningModeInitialStepsInstruction: msg.CommissioningModeInitialStepsInstruction,
		CommissioningModeSecondaryStepsHint:      msg.CommissioningModeSecondaryStepsHint,
		CommissioningModeSecondaryStepsInstruction: msg.CommissioningModeSecondaryStepsInstruction,
		UserManualURL: msg.UserManualURL,
		SupportURL:    msg.SupportURL,
		ProductURL:    msg.ProductURL,
	}

	// store new model
	keeper.Logger(ctx).Info("Creating a new model",
		"Model :", model.String())

	keeper.SetModel(ctx, model)

	return sdk.Result{}
}

//nolint:funlen
func handleMsgUpdateModel(ctx sdk.Context, keeper keeper.Keeper, authKeeper auth.Keeper,
	msg types.MsgUpdateModel) sdk.Result {
	// check if model exists
	if !keeper.IsModelPresent(ctx, msg.VID, msg.PID) {
		return types.ErrModelDoesNotExist(msg.VID, msg.PID).Result()
	}

	model := keeper.GetModel(ctx, msg.VID, msg.PID)

	// check if sender has enough rights to update model
	if err := checkModelRights(ctx, authKeeper, msg.Signer, msg.VID, "msgUpdateModel"); err != nil {
		return err.Result()
	}

	// updates existing model value only if corresponding value in MsgUpdate is not empty

	if msg.DeviceTypeID != 0 {
		model.DeviceTypeID = msg.DeviceTypeID
	}

	if msg.ProductLabel != "" {
		model.ProductLabel = msg.ProductLabel
	}

	if msg.CommissioningCustomFlowURL != "" {
		model.CommissioningCustomFlowURL = msg.CommissioningCustomFlowURL
	}

	if msg.UserManualURL != "" {
		model.UserManualURL = msg.UserManualURL
	}

	if msg.SupportURL != "" {
		model.SupportURL = msg.SupportURL
	}

	if msg.ProductURL != "" {
		model.ProductURL = msg.ProductURL
	}

	// store updated model
	keeper.SetModel(ctx, model)

	return sdk.Result{}
}

//nolint:unused,deadcode
func handleMsgDeleteModel(ctx sdk.Context, keeper keeper.Keeper, authKeeper auth.Keeper,
	msg types.MsgDeleteModel) sdk.Result {
	// check if model exists
	if !keeper.IsModelPresent(ctx, msg.VID, msg.PID) {
		return types.ErrModelDoesNotExist(msg.VID, msg.PID).Result()
	}

	// check if sender has enough rights to delete model
	if err := checkModelRights(ctx, authKeeper, msg.Signer, msg.VID, "msgDeleteModel"); err != nil {
		return err.Result()
	}

	// remove model from the store
	keeper.DeleteModel(ctx, msg.VID, msg.PID)

	return sdk.Result{}
}

func checkModelRights(ctx sdk.Context, authKeeper auth.Keeper, signer sdk.AccAddress, vid uint16, message string) sdk.Error {
	// sender must have Vendor role to add new model
	if !authKeeper.HasRole(ctx, signer, auth.Vendor) {
		return sdk.ErrUnauthorized(fmt.Sprintf("%s transaction should be "+
			"signed by an account with the %s role", message, auth.Vendor))
	}
	if !authKeeper.HasVendorId(ctx, signer, vid) {
		return sdk.ErrUnauthorized(fmt.Sprintf("%s transaction should be "+
			"signed by an vendor account containing the vendorId %v ", message, vid))
	}

	return nil
}
