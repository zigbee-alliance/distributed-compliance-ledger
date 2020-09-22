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

package modelinfo

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelinfo/internal/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelinfo/internal/types"
)

func NewHandler(keeper keeper.Keeper, authKeeper auth.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgAddModelInfo:
			return handleMsgAddModelInfo(ctx, keeper, authKeeper, msg)
		case types.MsgUpdateModelInfo:
			return handleMsgUpdateModelInfo(ctx, keeper, authKeeper, msg)
			/*		case type.MsgDeleteModelInfo:
					return handleMsgDeleteModelInfo(ctx, keeper, authKeeper, msg)*/
		default:
			errMsg := fmt.Sprintf("unrecognized nameservice Msg type: %v", msg.Type())

			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgAddModelInfo(ctx sdk.Context, keeper keeper.Keeper, authKeeper auth.Keeper,
	msg types.MsgAddModelInfo) sdk.Result {
	// check if model already exists
	if keeper.IsModelInfoPresent(ctx, msg.VID, msg.PID) {
		return types.ErrModelInfoAlreadyExists(msg.VID, msg.PID).Result()
	}

	// check sender has enough rights to add model
	if err := checkAddModelRights(ctx, authKeeper, msg.Signer); err != nil {
		return err.Result()
	}

	modelInfo := types.NewModelInfo(
		msg.VID,
		msg.PID,
		msg.CID,
		msg.Version,
		msg.Name,
		msg.Description,
		msg.SKU,
		msg.HardwareVersion,
		msg.FirmwareVersion,
		msg.OtaURL,
		msg.OtaChecksum,
		msg.OtaChecksumType,
		msg.Custom,
		msg.TisOrTrpTestingCompleted,
		msg.Signer,
	)

	// store new model
	keeper.SetModelInfo(ctx, modelInfo)

	return sdk.Result{}
}

func handleMsgUpdateModelInfo(ctx sdk.Context, keeper keeper.Keeper, authKeeper auth.Keeper,
	msg types.MsgUpdateModelInfo) sdk.Result {
	// check if model exists
	if !keeper.IsModelInfoPresent(ctx, msg.VID, msg.PID) {
		return types.ErrModelInfoDoesNotExist(msg.VID, msg.PID).Result()
	}

	modelInfo := keeper.GetModelInfo(ctx, msg.VID, msg.PID)

	// check if sender has enough rights to update model
	if err := checkUpdateModelRights(modelInfo.Owner, msg.Signer); err != nil {
		return err.Result()
	}

	if msg.OtaURL != "" && modelInfo.OtaURL == "" {
		return types.ErrOtaURLCannotBeSet(msg.VID, msg.PID).Result()
	}

	// updates existing model value only if corresponding value in MsgUpdate is not empty

	if msg.CID != 0 {
		modelInfo.CID = msg.CID
	}

	if msg.Description != "" {
		modelInfo.Description = msg.Description
	}

	if msg.OtaURL != "" {
		modelInfo.OtaURL = msg.OtaURL
	}

	if msg.Custom != "" {
		modelInfo.Custom = msg.Custom
	}

	modelInfo.TisOrTrpTestingCompleted = msg.TisOrTrpTestingCompleted

	// store updated model
	keeper.SetModelInfo(ctx, modelInfo)

	return sdk.Result{}
}

//nolint:unused,deadcode
func handleMsgDeleteModelInfo(ctx sdk.Context, keeper keeper.Keeper, authKeeper auth.Keeper,
	msg types.MsgDeleteModelInfo) sdk.Result {
	// check if model exists
	if !keeper.IsModelInfoPresent(ctx, msg.VID, msg.PID) {
		return types.ErrModelInfoDoesNotExist(msg.VID, msg.PID).Result()
	}

	modelInfo := keeper.GetModelInfo(ctx, msg.VID, msg.PID)

	// check if sender has enough rights to delete model
	if err := checkUpdateModelRights(modelInfo.Owner, msg.Signer); err != nil {
		return err.Result()
	}

	// remove model from the store
	keeper.DeleteModelInfo(ctx, msg.VID, msg.PID)

	return sdk.Result{}
}

func checkAddModelRights(ctx sdk.Context, authKeeper auth.Keeper, signer sdk.AccAddress) sdk.Error {
	// sender must have Vendor role to add new model
	if !authKeeper.HasRole(ctx, signer, auth.Vendor) {
		return sdk.ErrUnauthorized(fmt.Sprintf("MsgAddModelInfo transaction should be "+
			"signed by an account with the %s role", auth.Vendor))
	}

	return nil
}

func checkUpdateModelRights(owner sdk.AccAddress, signer sdk.AccAddress) sdk.Error {
	// sender must be equal to owner to edit model
	if !signer.Equals(owner) {
		return sdk.ErrUnauthorized("MsgUpdateModelInfo tx should be signed by owner")
	}

	return nil
}
