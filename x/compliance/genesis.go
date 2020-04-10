package compliance

import (
	"fmt"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"

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
		if record.VID == 0 {
			return fmt.Errorf("invalid ModelInfoRecord: Value: %d. Error: Invalid VID", record.VID)
		}

		if record.PID == 0 {
			return fmt.Errorf("invalid ModelInfoRecord: Value: %d. Error: Invalid PID", record.PID)
		}

		if record.Name == "" {
			return fmt.Errorf("invalid ModelInfoRecord: Value: %s. Error: Missing Name", record.Name)
		}

		if record.Owner == nil {
			return fmt.Errorf("invalid ModelInfoRecord: Value: %s. Error: Missing Owner", record.Owner)
		}

		if record.Description == "" {
			return fmt.Errorf("invalid ModelInfoRecord: Value: %s. Error: Missing Description", record.Description)
		}

		if record.SKU == "" {
			return fmt.Errorf("invalid ModelInfoRecord: Value: %s. Error: Missing SKU", record.SKU)
		}

		if record.FirmwareVersion == "" {
			return fmt.Errorf("invalid ModelInfoRecord: Value: %s. Error: Missing FirmwareVersion", record.FirmwareVersion)
		}

		if record.HardwareVersion == "" {
			return fmt.Errorf("invalid ModelInfoRecord: Value: %s. Error: Missing HardwareVersion", record.HardwareVersion)
		}

		if record.Custom == "" {
			return fmt.Errorf("invalid ModelInfoRecord: Value: %s. Error: Missing Custom", record.HardwareVersion)
		}
	}

	return nil
}

func DefaultGenesisState() GenesisState {
	return NewGenesisState()
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.ModelInfoRecords {
		keeper.SetModelInfo(ctx, record)
	}

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []ModelInfo

	k.IterateModelInfos(ctx, func(modelInfo types.ModelInfo) (stop bool) {
		records = append(records, modelInfo)
		return false
	})

	return GenesisState{ModelInfoRecords: records}
}
