package simulation

import (
	"math/rand"
	"strconv"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	appEncoding "github.com/zigbee-alliance/distributed-compliance-ledger/app"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

// FIXME issue 110: fix dependencies on AccountKeeper and BankKeeper

// Prevent strconv unused error.
var _ = strconv.IntSize

func SimulateMsgCreateModel(
	// ak types.AccountKeeper,
	// bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		i := r.Int()
		msg := &types.MsgCreateModel{
			Creator: simAccount.Address.String(),
			Vid:     int32(i),
			Pid:     int32(i),
		}

		_, found := k.GetModel(ctx, msg.Vid, msg.Pid)
		if found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Model already exist"), nil, nil
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           appEncoding.MakeEncodingConfig().TxConfig,
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

func SimulateMsgUpdateModel(
	// ak types.AccountKeeper,
	// bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount = simtypes.Account{}
			model      = types.Model{}
			msg        = &types.MsgUpdateModel{}
			allModel   = k.GetAllModel(ctx)
			found      = false
		)

		for _, obj := range allModel {
			simAccount, found = FindAccount(accs, obj.Creator)
			if found {
				model = obj

				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "model creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()

		msg.Vid = model.Vid
		msg.Pid = model.Pid

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           appEncoding.MakeEncodingConfig().TxConfig,
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

func SimulateMsgDeleteModel(
	// ak types.AccountKeeper,
	// bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount = simtypes.Account{}
			model      = types.Model{}
			msg        = &types.MsgUpdateModel{}
			allModel   = k.GetAllModel(ctx)
			found      = false
		)

		for _, obj := range allModel {
			simAccount, found = FindAccount(accs, obj.Creator)
			if found {
				model = obj

				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "model creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()

		msg.Vid = model.Vid
		msg.Pid = model.Pid

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           appEncoding.MakeEncodingConfig().TxConfig,
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
