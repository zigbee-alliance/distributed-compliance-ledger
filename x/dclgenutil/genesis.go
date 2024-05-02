package dclgenutil

import (
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclgenutil/types"
)

// InitGenesis - initialize accounts and deliver genesis transactions.
func InitGenesis(
	ctx sdk.Context, validatorKeeper types.ValidatorKeeper,
	deliverTx deliverTxfn, genesisState types.GenesisState,
	txEncodingConfig client.TxEncodingConfig,
) (validators []abci.ValidatorUpdate, err error) {
	if len(genesisState.GenTxs) > 0 {
		validators, err = DeliverGenTxs(ctx, genesisState.GenTxs, validatorKeeper, deliverTx, txEncodingConfig)
	}

	return
}
