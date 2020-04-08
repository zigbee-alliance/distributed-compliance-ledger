package compliance

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/test_constants"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"testing"
)

func TestModule_AddGetModelInfo(t *testing.T) {
	setup := Setup()
	owner := setup.Manufacturer()

	// add new model
	modelInfo := TestMsgAddModelInfo(owner)
	result := setup.Handler(setup.Ctx, modelInfo)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query model
	result2, _ := setup.Querier(
		setup.Ctx,
		[]string{keeper.QueryModelInfo, test_constants.Id},
		abci.RequestQuery{},
	)

	var receivedModelInfo types.ModelInfo
	_ = setup.Cdc.UnmarshalJSON(result2, &receivedModelInfo)

	require.Equal(t, receivedModelInfo.ID, modelInfo.ID)
	require.Equal(t, receivedModelInfo.Name, modelInfo.Name)
	require.Equal(t, receivedModelInfo.Owner, modelInfo.Owner)
	require.Equal(t, receivedModelInfo.Description, modelInfo.Description)
	require.Equal(t, receivedModelInfo.SKU, modelInfo.SKU)
	require.Equal(t, receivedModelInfo.FirmwareVersion, modelInfo.FirmwareVersion)
	require.Equal(t, receivedModelInfo.HardwareVersion, modelInfo.HardwareVersion)
	require.Equal(t, receivedModelInfo.CertificateID, modelInfo.CertificateID)
	require.Equal(t, receivedModelInfo.CertifiedDate, modelInfo.CertifiedDate)
	require.Equal(t, receivedModelInfo.TisOrTrpTestingCompleted, modelInfo.TisOrTrpTestingCompleted)
}
