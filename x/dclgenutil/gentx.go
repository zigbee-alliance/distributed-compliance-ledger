package dclgenutil

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclgenutil/types"
)

// SetGenTxsInAppGenesisState - sets the genesis transactions in the app genesis state.
func SetGenTxsInAppGenesisState(
	cdc codec.JSONCodec, txJSONEncoder sdk.TxEncoder, appGenesisState map[string]json.RawMessage, genTxs []sdk.Tx,
) (map[string]json.RawMessage, error) {
	genesisState := types.GetGenesisStateFromAppState(cdc, appGenesisState)
	genTxsBz := make([]json.RawMessage, 0, len(genTxs))

	for _, genTx := range genTxs {
		txBz, err := txJSONEncoder(genTx)
		if err != nil {
			return appGenesisState, err
		}

		genTxsBz = append(genTxsBz, txBz)
	}

	genesisState.GenTxs = genTxsBz
	return types.SetGenesisStateInAppState(cdc, appGenesisState, genesisState), nil
}

// ValidateAccountInGenesis checks that the provided account is presented in
// the set of genesis accounts.
func ValidateAccountInGenesis(
	appGenesisState map[string]json.RawMessage, genAccIterator dclauthtypes.GenesisAccountsIterator,
	addr sdk.Address, cdc codec.JSONCodec,
) error {
	var err error

	accountIsInGenesis := false

	genAccIterator.IterateGenesisAccounts(cdc, appGenesisState,
		func(acc dclauthtypes.GenesisAccount) (stop bool) {
			accAddress := acc.GetAddress()

			// ensure that account is in genesis
			if accAddress.Equals(addr) {
				accountIsInGenesis = true
				return true
			}

			return false
		},
	)

	if err != nil {
		return err
	}

	if !accountIsInGenesis {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
			"Error account %s in not in the app_state.accounts array of genesis.json", addr,
		)
	}

	return nil
}

type deliverTxfn func(abci.RequestDeliverTx) abci.ResponseDeliverTx

// DeliverGenTxs iterates over all genesis txs, decodes each into a Tx and
// invokes the provided deliverTxfn with the decoded Tx. It returns the result
// of the validator module's ApplyAndReturnValidatorSetUpdates.
func DeliverGenTxs(
	ctx sdk.Context, genTxs []json.RawMessage,
	validatorKeeper types.ValidatorKeeper, deliverTx deliverTxfn,
	txEncodingConfig client.TxEncodingConfig,
) ([]abci.ValidatorUpdate, error) {
	for _, genTx := range genTxs {
		tx, err := txEncodingConfig.TxJSONDecoder()(genTx)
		if err != nil {
			panic(err)
		}

		bz, err := txEncodingConfig.TxEncoder()(tx)
		if err != nil {
			panic(err)
		}

		res := deliverTx(abci.RequestDeliverTx{Tx: bz})
		if !res.IsOK() {
			panic(res.Log)
		}
	}

	return validatorKeeper.ApplyAndReturnValidatorSetUpdates(ctx)
}
