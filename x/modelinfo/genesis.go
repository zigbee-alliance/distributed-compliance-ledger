package modelinfo

//nolint:goimports
import (
	"fmt"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/types"
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
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid ModelInfoRecord: value: %d. "+
				"Error: Invalid VID", record.VID))
		}

		if record.PID == 0 {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid ModelInfoRecord: value: %d. "+
				"Error: Invalid PID", record.PID))
		}

		if record.Name == "" {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid ModelInfoRecord: value: %s. "+
				"Error: Missing Name", record.Name))
		}

		if record.Owner == nil {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid ModelInfoRecord: value: %s. "+
				"Error: Missing Owner", record.Owner))
		}

		if record.Description == "" {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid ModelInfoRecord: value: %s. "+
				"Error: Missing Description", record.Description))
		}

		if record.SKU == "" {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid ModelInfoRecord: value: %s. "+
				"Error: Missing SKU", record.SKU))
		}

		if record.FirmwareVersion == "" {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid ModelInfoRecord: value: %s."+
				" Error: Missing FirmwareVersion", record.FirmwareVersion))
		}

		if record.HardwareVersion == "" {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid ModelInfoRecord: value: %s. "+
				"Error: Missing HardwareVersion", record.HardwareVersion))
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
