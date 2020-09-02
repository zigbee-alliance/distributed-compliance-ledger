package compliance

import (
	"fmt"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	ComplianceInfoRecords []ComplianceInfo `json:"compliance_model_records"`
}

func NewGenesisState() GenesisState {
	return GenesisState{ComplianceInfoRecords: []ComplianceInfo{}}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.ComplianceInfoRecords {
		if record.VID == 0 {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("Invalid CertifiedModelRecord: value: %d. "+
					"Error: Invalid VID: it cannot be 0", record.VID))
		}

		if record.PID == 0 {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("Invalid CertifiedModelRecord: value: %d. "+
					"Error: Invalid PID: it cannot be 0", record.PID))
		}

		if len(record.State) == 0 {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("Invalid CertifiedModelRecord: value: %d."+
					" Error: Invalid State: it cannot be empty", record.PID))
		}

		if record.Date.IsZero() {
			return sdk.ErrUnknownRequest("Invalid Date: it cannot be empty")
		}

		if record.CertificationType != "" && record.CertificationType != types.ZbCertificationType {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("Invalid CertifiedModelRecord: value: %v."+
					" Error: Invalid CertificationType: "+
					"unknown type; supported types: [%s]", record.CertificationType, types.ZbCertificationType))
		}
	}

	return nil
}

func DefaultGenesisState() GenesisState {
	return NewGenesisState()
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.ComplianceInfoRecords {
		keeper.SetComplianceInfo(ctx, record)
	}

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []ComplianceInfo

	k.IterateComplianceInfos(ctx, "", func(deviceCompliance types.ComplianceInfo) (stop bool) {
		records = append(records, deviceCompliance)
		return false
	})

	return GenesisState{ComplianceInfoRecords: records}
}
