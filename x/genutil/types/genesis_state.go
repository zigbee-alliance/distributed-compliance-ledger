package types

//nolint:goimports
import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	cosmosgenutil "github.com/cosmos/cosmos-sdk/x/genutil"
)

// ValidateGenesis validates GenTx transactions.
func ValidateGenesis(genesisState cosmosgenutil.GenesisState) error {
	for i, genTx := range genesisState.GenTxs {
		var tx authtypes.StdTx
		if err := ModuleCdc.UnmarshalJSON(genTx, &tx); err != nil {
			return err
		}

		msgs := tx.GetMsgs()
		if len(msgs) != 1 {
			return sdk.ErrUnknownRequest("must provide genesis StdTx with exactly 1 CreateValidator message")
		}

		if _, ok := msgs[0].(validator.MsgCreateValidator); !ok {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("genesis transaction %v does not contain a MsgCreateValidator", i))
		}
	}

	return nil
}
