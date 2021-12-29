package utils

import (
	"github.com/cosmos/cosmos-sdk/client"
	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SimAppChainID hardcoded chainID for simulation.
const (
	DefaultGenTxGas = 1000000
	SimAppChainID   = "simulation-app"
)

// GenTx generates a signed mock transaction.
func GenTx(txf clienttx.Factory, gen client.TxConfig, msgs []sdk.Msg, feeAmt sdk.Coins, gas uint64, signer string) (sdk.Tx, error) {
	tx, err := txf.BuildUnsignedTx(msgs...)
	if err != nil {
		return nil, err
	}

	// memo := simulation.RandStringOfLength(r, simulation.RandIntBetween(r, 0, 100))
	// tx.SetMemo(memo)
	tx.SetFeeAmount(feeAmt)
	tx.SetGasLimit(gas)

	err = clienttx.Sign(txf, signer, tx, true)
	if err != nil {
		return nil, err
	}

	return tx.GetTx(), nil
}
