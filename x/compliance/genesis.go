package compliance

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	ModelInfoRecords []ModelInfo `json:"model_info_records"`
}

func NewGenesisState() GenesisState {
	return GenesisState{ModelInfoRecords: []ModelInfo{}}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.ModelInfoRecords {
		if record.ID == "" {
			return fmt.Errorf("invalid ModelInfoRecord: Value: %s. Error: Missing ID", record.ID)
		}

		if record.Family == "" {
			return fmt.Errorf("invalid ModelInfoRecord: Value: %s. Error: Missing Family", record.Family)
		}

		if record.Cert == "" {
			return fmt.Errorf("invalid ModelInfoRecord: Value: %s. Error: Missing Cert", record.Cert)
		}

		if record.Owner == nil {
			return fmt.Errorf("invalid ModelInfoRecord: Value: %s. Error: Missing Owner", record.Owner)
		}
	}

	return nil
}

func DefaultGenesisState() GenesisState {
	return NewGenesisState()
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.ModelInfoRecords {
		keeper.SetModelInfo(ctx, record.ID, record)
	}

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []ModelInfo

	iterator := k.GetModelInfoIDIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		id := string(iterator.Key())
		modelInfo := k.GetModelInfo(ctx, id)
		records = append(records, modelInfo)
	}

	return GenesisState{ModelInfoRecords: records}
}
