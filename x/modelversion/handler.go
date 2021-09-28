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

package modelversion

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelversion/internal/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelversion/internal/types"
)

func NewHandler(keeper keeper.Keeper, authKeeper auth.Keeper, modelKeeper model.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {

		switch msg := msg.(type) {
		case types.MsgAddModelVersion:
			return handleMsgAddModelVersion(ctx, keeper, authKeeper, modelKeeper, msg)
		case types.MsgUpdateModelVersion:
			return handleMsgUpdateModelVersion(ctx, keeper, authKeeper, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized nameservice Msg type: %v", msg.Type())

			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgAddModelVersion(ctx sdk.Context, keeper keeper.Keeper, authKeeper auth.Keeper, modelKeeper model.Keeper,
	msg types.MsgAddModelVersion) sdk.Result {

	// check sender has enough rights to add model
	if err := checkModelRights(ctx, authKeeper, msg.Signer, msg.VID); err != nil {
		return err.Result()
	}

	// check if model exists
	if !modelKeeper.IsModelPresent(ctx, msg.VID, msg.PID) {
		return types.ErrModelDoesNotExist(msg.VID, msg.PID).Result()
	}

	// check if model version already exists
	if keeper.IsModelVersionPresent(ctx, msg.VID, msg.PID, msg.SoftwareVersion) {
		return types.ErrModelVersionAlreadyExists(msg.VID, msg.PID, msg.SoftwareVersion).Result()
	}

	modelVersion := types.ModelVersion{
		VID:                          msg.VID,
		PID:                          msg.PID,
		SoftwareVersion:              msg.SoftwareVersion,
		SoftwareVersionString:        msg.SoftwareVersionString,
		CDVersionNumber:              msg.CDVersionNumber,
		FirmwareDigests:              msg.FirmwareDigests,
		SoftwareVersionValid:         msg.SoftwareVersionValid,
		OtaURL:                       msg.OtaURL,
		OtaFileSize:                  msg.OtaFileSize,
		OtaChecksum:                  msg.OtaChecksum,
		OtaChecksumType:              msg.OtaChecksumType,
		MinApplicableSoftwareVersion: msg.MinApplicableSoftwareVersion,
		MaxApplicableSoftwareVersion: msg.MaxApplicableSoftwareVersion,
		ReleaseNotesURL:              msg.ReleaseNotesURL,
	}

	// store new model version
	keeper.Logger(ctx).Info("Creating a new model version",
		"ModelVersion :", modelVersion.String())

	keeper.SetModelVersion(ctx, modelVersion)
	return sdk.Result{}
}

//nolint:funlen
func handleMsgUpdateModelVersion(ctx sdk.Context, keeper keeper.Keeper, authKeeper auth.Keeper,
	msg types.MsgUpdateModelVersion) sdk.Result {
	// check if model exists
	if !keeper.IsModelVersionPresent(ctx, msg.VID, msg.PID, msg.SoftwareVersion) {
		return types.ErrModelVersionDoesNotExist(msg.VID, msg.PID, msg.SoftwareVersion).Result()
	}

	modelVersion := keeper.GetModelVersion(ctx, msg.VID, msg.PID, msg.SoftwareVersion)

	// check if sender has enough rights to update model
	if err := checkModelRights(ctx, authKeeper, msg.Signer, msg.VID); err != nil {
		return err.Result()
	}

	if msg.OtaURL != "" && modelVersion.OtaURL == "" {
		return types.ErrOtaURLCannotBeSet(msg.VID, msg.PID, msg.SoftwareVersion).Result()
	}

	// updates existing model version value only if corresponding value in MsgUpdate is not empty
	// p.s. only mutable fields are updated.

	if msg.SoftwareVersionValid != modelVersion.SoftwareVersionValid {
		modelVersion.SoftwareVersionValid = msg.SoftwareVersionValid
	}

	if msg.OtaURL != "" {
		modelVersion.OtaURL = msg.OtaURL
	}

	if msg.MinApplicableSoftwareVersion != 0 {
		modelVersion.MinApplicableSoftwareVersion = msg.MinApplicableSoftwareVersion
	}

	if msg.MaxApplicableSoftwareVersion != 0 {
		modelVersion.MinApplicableSoftwareVersion = msg.MaxApplicableSoftwareVersion
	}

	if msg.ReleaseNotesURL != "" {
		modelVersion.ReleaseNotesURL = msg.ReleaseNotesURL
	}

	// store updated model
	keeper.SetModelVersion(ctx, modelVersion)

	return sdk.Result{}
}

func checkModelRights(ctx sdk.Context, authKeeper auth.Keeper, signer sdk.AccAddress, vid uint16) sdk.Error {
	// sender must have Vendor role to add new model
	if !authKeeper.HasRole(ctx, signer, auth.Vendor) {
		return sdk.ErrUnauthorized(fmt.Sprintf("ModelVersion Add/Update transaction should be "+
			"signed by an account with the %s role", auth.Vendor))
	}
	if !authKeeper.HasVendorId(ctx, signer, vid) {
		return sdk.ErrUnauthorized(fmt.Sprintf("ModelVersion Add/Update transaction should be "+
			"signed by an vendor account containing the vendorId %v ", vid))
	}

	return nil
}
