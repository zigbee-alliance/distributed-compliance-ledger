package compliance

import (
	"fmt"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	CertifiedModelRecords []CertifiedModel `json:"certified_model_records"`
}

func NewGenesisState() GenesisState {
	return GenesisState{CertifiedModelRecords: []CertifiedModel{}}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.CertifiedModelRecords {
		if record.VID == 0 {
			return fmt.Errorf("invalid CertifiedModelRecord: value: %d. Error: Invalid VID: it cannot be 0", record.VID)
		}

		if record.PID == 0 {
			return fmt.Errorf("invalid CertifiedModelRecord: value: %d. Error: Invalid PID: it cannot be 0", record.PID)
		}

		if record.CertificationDate.IsZero() {
			return fmt.Errorf("invalid CertifiedModelRecord: value: %v. Error: Invalid CertificationDate: it cannot be empty", record.CertificationDate)
		}

		if record.CertificationType != "" && record.CertificationType != types.ZbCertificationType {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("invalid CertifiedModelRecord: value: %v. Error: Invalid CertificationType: unknown type; supported types: [%s]", record.CertificationType, types.ZbCertificationType))
		}
	}

	return nil
}

func DefaultGenesisState() GenesisState {
	return NewGenesisState()
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.CertifiedModelRecords {
		keeper.SetCertifiedModel(ctx, record)
	}

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []CertifiedModel

	k.IterateCertifiedModels(ctx, func(deviceCompliance types.CertifiedModel) (stop bool) {
		records = append(records, deviceCompliance)
		return false
	})

	return GenesisState{CertifiedModelRecords: records}
}
