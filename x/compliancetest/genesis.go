package compliancetest

import (
	"fmt"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	TestingResultRecords []TestingResults `json:"testing_result_records"`
}

func NewGenesisState() GenesisState {
	return GenesisState{TestingResultRecords: []TestingResults{}}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.TestingResultRecords {
		if record.VID == 0 {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid TestingResultRecord: value: %v. "+
				"Error: Invalid VID", record.VID))
		}

		if record.PID == 0 {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid TestingResultRecord: value: %v. "+
				"Error: Invalid PID", record.PID))
		}

		if len(record.Results) == 0 {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid TestingResultRecord: value: %s. "+
				"Error: Missing TestResult", record.Results))
		}
	}

	return nil
}

func DefaultGenesisState() GenesisState {
	return NewGenesisState()
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.TestingResultRecords {
		keeper.SetTestingResults(ctx, record)
	}

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []TestingResults

	k.IterateTestingResults(ctx, func(testingResult types.TestingResults) (stop bool) {
		records = append(records, testingResult)
		return false
	})

	return GenesisState{TestingResultRecords: records}
}
