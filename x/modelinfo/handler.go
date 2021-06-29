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

	model := types.Model{
		VID:                                      msg.VID,
		PID:                                      msg.PID,
		CID:                                      msg.CID,
		Name:                                     msg.Name,
		Description:                              msg.Description,
		SKU:                                      msg.SKU,
		SoftwareVersion:                          msg.SoftwareVersion,
		SoftwareVersionString:                    msg.SoftwareVersionString,
		HardwareVersion:                          msg.HardwareVersion,
		HardwareVersionString:                    msg.HardwareVersionString,
		CDVersionNumber:                          msg.CDVersionNumber,
		FirmwareDigests:                          msg.FirmwareDigests,
		Revoked:                                  msg.Revoked,
		OtaURL:                                   msg.OtaURL,
		OtaChecksum:                              msg.OtaChecksum,
		OtaChecksumType:                          msg.OtaChecksumType,
		OtaBlob:                                  msg.OtaBlob,
		CommissioningCustomFlow:                  msg.CommissioningCustomFlow,
		CommissioningCustomFlowURL:               msg.CommissioningCustomFlowURL,
		CommissioningModeInitialStepsHint:        msg.CommissioningModeInitialStepsHint,
		CommissioningModeInitialStepsInstruction: msg.CommissioningModeInitialStepsInstruction,
		CommissioningModeSecondaryStepsHint:      msg.CommissioningModeSecondaryStepsHint,
		CommissioningModeSecondaryStepsInstruction: msg.CommissioningModeSecondaryStepsInstruction,
		ReleaseNotesURL: msg.ReleaseNotesURL,
		UserManualURL:   msg.UserManualURL,
		SupportURL:      msg.SupportURL,
		ProductURL:      msg.ProductURL,
		ChipBlob:        msg.ChipBlob,
		VendorBlob:      msg.VendorBlob,
	}
	modelInfo := types.NewModelInfo(
		model,
		msg.Signer,
	)

	// store new model
	keeper.SetModelInfo(ctx, modelInfo)

	return sdk.Result{}
}

//nolint:funlen
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

	if msg.OtaURL != "" && modelInfo.Model.OtaURL == "" {
		return types.ErrOtaURLCannotBeSet(msg.VID, msg.PID).Result()
	}

	// updates existing model value only if corresponding value in MsgUpdate is not empty

	if msg.CID != 0 {
		modelInfo.Model.CID = msg.CID
	}

	if msg.Description != "" {
		modelInfo.Model.Description = msg.Description
	}

	if msg.Revoked != modelInfo.Model.Revoked {
		modelInfo.Model.Revoked = msg.Revoked
	}

	if msg.CDVersionNumber != 0 {
		modelInfo.Model.CDVersionNumber = msg.CDVersionNumber
	}

	if msg.OtaURL != "" {
		modelInfo.Model.OtaURL = msg.OtaURL
	}

	if msg.OtaChecksum != "" {
		modelInfo.Model.OtaChecksum = msg.OtaChecksum
	}

	if msg.OtaChecksumType != "" {
		modelInfo.Model.OtaChecksumType = msg.OtaChecksumType
	}

	if msg.CommissioningCustomFlowURL != "" {
		modelInfo.Model.CommissioningCustomFlowURL = msg.CommissioningCustomFlowURL
	}

	if msg.ReleaseNotesURL != "" {
		modelInfo.Model.ReleaseNotesURL = msg.ReleaseNotesURL
	}

	if msg.UserManualURL != "" {
		modelInfo.Model.UserManualURL = msg.UserManualURL
	}

	if msg.SupportURL != "" {
		modelInfo.Model.SupportURL = msg.SupportURL
	}

	if msg.ProductURL != "" {
		modelInfo.Model.ProductURL = msg.ProductURL
	}

	if msg.ChipBlob != "" {
		modelInfo.Model.ChipBlob = msg.ChipBlob
	}

	if msg.VendorBlob != "" {
		modelInfo.Model.VendorBlob = msg.VendorBlob
	}

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
