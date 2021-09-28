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

package compliancetest

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/internal/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/internal/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelversion"
)

func NewHandler(keeper keeper.Keeper, modelversionKeeper modelversion.Keeper, authKeeper auth.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		keeper.Logger(ctx).Info("Inside the handleMsgAddTestingCenter...")
		switch msg := msg.(type) {
		case types.MsgAddTestingResult:
			return handleMsgAddTestingResult(ctx, keeper, modelversionKeeper, authKeeper, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized compliancetest Msg type: %v", msg.Type())

			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgAddTestingResult(ctx sdk.Context, keeper keeper.Keeper, modelversionKeeper modelversion.Keeper,
	authKeeper auth.Keeper, msg types.MsgAddTestingResult) sdk.Result {
	// check if sender has enough rights to add testing results
	keeper.Logger(ctx).Info("Inside the handleMsgAddTestingCenter")
	if err := checkAddTestingResultRights(ctx, authKeeper, msg.Signer); err != nil {
		return err.Result()
	}

	// check that corresponding model exists on the ledger
	if !modelversionKeeper.IsModelVersionPresent(ctx, msg.VID, msg.PID, msg.SoftwareVersion) {
		return modelversion.ErrModelVersionDoesNotExist(msg.VID, msg.PID, msg.SoftwareVersion).Result()
	}

	testingResult := types.NewTestingResult(
		msg.VID,
		msg.PID,
		msg.SoftwareVersion,
		msg.SoftwareVersionString,
		msg.Signer,
		msg.TestResult,
		msg.TestDate,
	)

	// store testing results. it extends existing value if testing results already exists
	keeper.AddTestingResult(ctx, testingResult)

	return sdk.Result{}
}

func checkAddTestingResultRights(ctx sdk.Context, authKeeper auth.Keeper, signer sdk.AccAddress) sdk.Error {
	// sender must have TestHouse role to add new model
	if !authKeeper.HasRole(ctx, signer, auth.TestHouse) {
		return sdk.ErrUnauthorized(fmt.Sprintf(
			"MsgAddTestingResult transaction should be signed by an account with the %s role", auth.TestHouse))
	}

	return nil
}
