package utils

import (
	"github.com/cosmos/cosmos-sdk/client"
	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenTx generates a signed mock transaction.
func GenTx(txf clienttx.Factory, _ client.TxConfig, msgs []sdk.Msg, signer string) (sdk.Tx, error) {
	// tx, err := txf.BuildUnsignedTx(msgs...)
	tx, err := txf.BuildUnsignedTx(msgs...)
	if err != nil {
		return nil, err
	}

	err = clienttx.Sign(txf, signer, tx, true)
	if err != nil {
		return nil, err
	}

	return tx.GetTx(), nil
}
