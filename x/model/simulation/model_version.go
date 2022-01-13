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

package simulation

import (
	"math/rand"
	"strconv"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

// FIXME issue 110: fix dependencies on AccountKeeper and BankKeeper

// Prevent strconv unused error.
var _ = strconv.IntSize

func SimulateMsgCreateModelVersion(
	// ak types.AccountKeeper,
	// bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		i := r.Int()
		msg := &types.MsgCreateModelVersion{
			Creator:         simAccount.Address.String(),
			Vid:             int32(i),
			Pid:             int32(i),
			SoftwareVersion: uint32(i),
		}

		_, found := k.GetModelVersion(ctx, msg.Vid, msg.Pid, msg.SoftwareVersion)
		if found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "ModelVersion already exist"), nil, nil
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			// AccountKeeper:   ak,
			// Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgUpdateModelVersion(
	// ak types.AccountKeeper,
	// bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount      = simtypes.Account{}
			modelVersion    = types.ModelVersion{}
			msg             = &types.MsgUpdateModelVersion{}
			allModelVersion = k.GetAllModelVersion(ctx)
			found           = false
		)
		for _, obj := range allModelVersion {
			simAccount, found = FindAccount(accs, obj.Creator)
			if found {
				modelVersion = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "modelVersion creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()

		msg.Vid = modelVersion.Vid
		msg.Pid = modelVersion.Pid
		msg.SoftwareVersion = modelVersion.SoftwareVersion

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			// AccountKeeper:   ak,
			// Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgDeleteModelVersion(
	// ak types.AccountKeeper,
	// bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount      = simtypes.Account{}
			modelVersion    = types.ModelVersion{}
			msg             = &types.MsgUpdateModelVersion{}
			allModelVersion = k.GetAllModelVersion(ctx)
			found           = false
		)
		for _, obj := range allModelVersion {
			simAccount, found = FindAccount(accs, obj.Creator)
			if found {
				modelVersion = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "modelVersion creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()

		msg.Vid = modelVersion.Vid
		msg.Pid = modelVersion.Pid
		msg.SoftwareVersion = modelVersion.SoftwareVersion

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			// AccountKeeper:   ak,
			// Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}
