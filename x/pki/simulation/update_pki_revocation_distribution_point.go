package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func SimulateMsgUpdatePkiRevocationDistributionPoint(
	_ keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgUpdatePkiRevocationDistributionPoint{
			Signer: simAccount.Address.String(),
		}

		// TODO: Handling the UpdatePkiRevocationDistributionPoint simulation

		return simtypes.NoOpMsg(pkitypes.ModuleName, msg.Type(), "UpdatePkiRevocationDistributionPoint simulation not implemented"), nil, nil
	}
}
