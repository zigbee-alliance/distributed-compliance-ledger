package genutil

import (
	"encoding/json"
	"fmt"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/genutil/types"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authexported "github.com/cosmos/cosmos-sdk/x/auth/exported"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// ValidateAccountInGenesis checks that the provided key has sufficient
// coins in the genesis accounts.
func ValidateAccountInGenesis(appGenesisState map[string]json.RawMessage,
	genAccIterator types.GenesisAccountsIterator,
	key sdk.Address, cdc *codec.Codec) error {
	accountIsInGenesis := false

	validatorDataBz := appGenesisState[validator.ModuleName]

	var validatorData validator.GenesisState

	cdc.MustUnmarshalJSON(validatorDataBz, &validatorData)

	genUtilDataBz := appGenesisState[validator.ModuleName]

	var genesisState GenesisState

	cdc.MustUnmarshalJSON(genUtilDataBz, &genesisState)

	genAccIterator.IterateGenesisAccounts(cdc, appGenesisState, func(acc authexported.Account) (stop bool) {
		accAddress := acc.GetAddress()

		if accAddress.Equals(key) {
			accountIsInGenesis = true
			return true
		}
		return false
	})

	if !accountIsInGenesis {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf("Error account %s in not in the app_state.accounts array of genesis.json", key))
	}

	return nil
}

type deliverTxfn func(abci.RequestDeliverTx) abci.ResponseDeliverTx

// DeliverGenTxs - deliver a genesis transaction.
func DeliverGenTxs(ctx sdk.Context, cdc *codec.Codec, genTxs []json.RawMessage,
	validatorKeeper types.ValidatorKeeper, deliverTx deliverTxfn) []abci.ValidatorUpdate {
	for _, genTx := range genTxs {
		var tx authtypes.StdTx

		cdc.MustUnmarshalJSON(genTx, &tx)

		bz := cdc.MustMarshalBinaryLengthPrefixed(tx)
		res := deliverTx(abci.RequestDeliverTx{Tx: bz})

		if !res.IsOK() {
			panic(res.Log)
		}
	}

	return validatorKeeper.ApplyAndReturnValidatorSetUpdates(ctx)
}
