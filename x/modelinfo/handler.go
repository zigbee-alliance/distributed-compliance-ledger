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
	if keeper.IsModelInfoPresent(ctx, msg.VID, msg.PID, msg.SoftwareVersion, msg.HardwareVersion) {
		return types.ErrModelInfoAlreadyExists(msg.VID, msg.PID, msg.SoftwareVersion, msg.HardwareVersion).Result()
	}

	// check sender has enough rights to add model
	if err := checkAddModelRights(ctx, authKeeper, msg.Signer); err != nil {
		return err.Result()
	}

	model := types.Model{
		VID:                                      msg.VID,
		PID:                                      msg.PID,
		CID:                                      msg.CID,
		ProductName:                              msg.ProductName,
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
	if !keeper.IsModelInfoPresent(ctx, msg.Model.VID, msg.Model.PID,
		msg.Model.SoftwareVersion, msg.Model.HardwareVersion) {
		return types.ErrModelInfoDoesNotExist(msg.Model.VID, msg.Model.PID,
			msg.Model.SoftwareVersion, msg.Model.HardwareVersion).Result()
	}

	modelInfo := keeper.GetModelInfo(ctx, msg.Model.VID, msg.Model.PID,
		msg.Model.SoftwareVersion, msg.Model.HardwareVersion)

	// check if sender has enough rights to update model
	if err := checkUpdateModelRights(modelInfo.Owner, msg.Signer); err != nil {
		return err.Result()
	}

	if msg.OtaURL != "" && modelInfo.Model.OtaURL == "" {
		return types.ErrOtaURLCannotBeSet(msg.Model.VID, msg.Model.PID,
		msg.Model.SoftwareVersion, msg.Model.HardwareVersion).Result()
	}

	// updates existing model value only if corresponding value in MsgUpdate is not empty

	if msg.Model.CID != 0 {
		modelInfo.Model.CID = msg.Model.CID
	}

	if msg.Model.Description != "" {
		modelInfo.Model.Description = msg.Model.Description
	}

	if msg.Model.Revoked != modelInfo.Model.Revoked {
		modelInfo.Model.Revoked = msg.Model.Revoked
	}

	if msg.Model.CDVersionNumber != 0 {
		modelInfo.Model.CDVersionNumber = msg.Model.CDVersionNumber
	}

	if msg.Model.OtaURL != "" {
		modelInfo.Model.OtaURL = msg.Model.OtaURL
	}

	if msg.Model.OtaChecksum != "" {
		modelInfo.Model.OtaChecksum = msg.Model.OtaChecksum
	}

	if msg.Model.OtaChecksumType != "" {
		modelInfo.Model.OtaChecksumType = msg.Model.OtaChecksumType
	}

	if msg.Model.CommissioningCustomFlowURL != "" {
		modelInfo.Model.CommissioningCustomFlowURL = msg.Model.CommissioningCustomFlowURL
	}

	if msg.Model.ReleaseNotesURL != "" {
		modelInfo.Model.ReleaseNotesURL = msg.Model.ReleaseNotesURL
	}

	if msg.Model.UserManualURL != "" {
		modelInfo.Model.UserManualURL = msg.Model.UserManualURL
	}

	if msg.Model.SupportURL != "" {
		modelInfo.Model.SupportURL = msg.Model.SupportURL
	}

	if msg.Model.ProductURL != "" {
		modelInfo.Model.ProductURL = msg.Model.ProductURL
	}

	if msg.Model.ChipBlob != "" {
		modelInfo.Model.ChipBlob = msg.Model.ChipBlob
	}

	if msg.Model.VendorBlob != "" {
		modelInfo.Model.VendorBlob = msg.Model.VendorBlob
	}

	// store updated model
	keeper.SetModelInfo(ctx, modelInfo)

	return sdk.Result{}
}

//nolint:unused,deadcode
func handleMsgDeleteModelInfo(ctx sdk.Context, keeper keeper.Keeper, authKeeper auth.Keeper,
	msg types.MsgDeleteModelInfo) sdk.Result {
	// check if model exists
	if !keeper.IsModelInfoPresent(ctx, msg.VID, msg.PID, msg.SoftwareVersion, msg.HardwareVersion) {
		return types.ErrModelInfoDoesNotExist(msg.VID, msg.PID, msg.SoftwareVersion, msg.HardwareVersion).Result()
	}

	modelInfo := keeper.GetModelInfo(ctx, msg.VID, msg.PID, msg.SoftwareVersion, msg.HardwareVersion)

	// check if sender has enough rights to delete model
	if err := checkUpdateModelRights(modelInfo.Owner, msg.Signer); err != nil {
		return err.Result()
	}

	// remove model from the store
	keeper.DeleteModelInfo(ctx, msg.VID, msg.PID, msg.SoftwareVersion, msg.HardwareVersion)

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
