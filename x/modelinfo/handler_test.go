package modelinfo

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/types"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/test_constants"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"testing"
)

func TestHandler_AddModel(t *testing.T) {
	setup := Setup()
	owner := setup.Vendor(test_constants.Address1)

	// add new model
	modelInfo := TestMsgAddModelInfo(owner)
	result := setup.Handler(setup.Ctx, modelInfo)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query model
	receivedModelInfo := queryModelInfo(setup, modelInfo.VID, modelInfo.PID)

	// check
	require.Equal(t, receivedModelInfo.VID, modelInfo.VID)
	require.Equal(t, receivedModelInfo.PID, modelInfo.PID)
	require.Equal(t, receivedModelInfo.Name, modelInfo.Name)
	require.Equal(t, receivedModelInfo.Description, modelInfo.Description)
}

func TestHandler_UpdateModel(t *testing.T) {
	setup := Setup()
	owner := setup.Vendor(test_constants.Address1)

	// try update not present model
	msgUpdatedModelInfo := TestMsgUpdatedModelInfo(owner)
	result := setup.Handler(setup.Ctx, msgUpdatedModelInfo)
	require.Equal(t, types.CodeModelInfoDoesNotExist, result.Code)

	// add new model
	msgAddModelInfo := TestMsgAddModelInfo(owner)
	result = setup.Handler(setup.Ctx, msgAddModelInfo)
	require.Equal(t, sdk.CodeOK, result.Code)

	// update existing model
	result = setup.Handler(setup.Ctx, msgUpdatedModelInfo)
	require.Equal(t, sdk.CodeOK, result.Code)
}

func TestHandler_OnlyOwnerCanUpdateModel(t *testing.T) {
	setup := Setup()
	owner := setup.Vendor(test_constants.Address1)
	administrator := setup.Administrator(test_constants.Address2)
	vendor := setup.Vendor(test_constants.Address3)

	// add new model
	msgAddModelInfo := TestMsgAddModelInfo(owner)
	result := setup.Handler(setup.Ctx, msgAddModelInfo)
	require.Equal(t, sdk.CodeOK, result.Code)

	// update existing model by non-owner administrator
	msgUpdatedModelInfo := TestMsgUpdatedModelInfo(administrator)
	result = setup.Handler(setup.Ctx, msgUpdatedModelInfo)
	require.Equal(t, sdk.CodeUnauthorized, result.Code)

	// update existing model by non-owner vendor
	msgUpdatedModelInfo = TestMsgUpdatedModelInfo(vendor)
	result = setup.Handler(setup.Ctx, msgUpdatedModelInfo)
	require.Equal(t, sdk.CodeUnauthorized, result.Code)

	// owner update existing model
	msgUpdatedModelInfo = TestMsgUpdatedModelInfo(owner)
	result = setup.Handler(setup.Ctx, msgUpdatedModelInfo)
	require.Equal(t, sdk.CodeOK, result.Code)
}

func TestHandler_AddModelWithEmptyOptionalFields(t *testing.T) {
	setup := Setup()
	owner := setup.Vendor(test_constants.Address1)

	// add new model
	modelInfo := TestMsgAddModelInfo(owner)
	modelInfo.CID = 0     // Set empty CID
	modelInfo.Custom = "" // Set empty Custom

	result := setup.Handler(setup.Ctx, modelInfo)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query model
	receivedModelInfo := queryModelInfo(setup, test_constants.VID, test_constants.PID)

	// check
	require.Equal(t, receivedModelInfo.CID, int16(0))
	require.Equal(t, receivedModelInfo.Custom, "")
}

func TestHandler_AddModelByNonVendor(t *testing.T) {
	setup := Setup()

	// add new model
	modelInfo := TestMsgAddModelInfo(test_constants.Address1)
	result := setup.Handler(setup.Ctx, modelInfo)
	require.Equal(t, sdk.CodeUnauthorized, result.Code)
}

func TestHandler_PartiallyUpdateModel(t *testing.T) {
	setup := Setup()
	owner := setup.Vendor(test_constants.Address1)

	// add new model
	msgAddModelInfo := TestMsgAddModelInfo(owner)
	result := setup.Handler(setup.Ctx, msgAddModelInfo)

	// owner update Description of existing model
	msgUpdatedModelInfo := TestMsgUpdatedModelInfo(owner)
	msgUpdatedModelInfo.Description = "New Description"
	msgUpdatedModelInfo.Custom = ""
	msgUpdatedModelInfo.CID = 0
	result = setup.Handler(setup.Ctx, msgUpdatedModelInfo)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query model
	receivedModelInfo := queryModelInfo(setup, msgUpdatedModelInfo.VID, msgUpdatedModelInfo.PID)

	// check
	require.Equal(t, receivedModelInfo.Description, msgUpdatedModelInfo.Description)
	require.Equal(t, receivedModelInfo.Custom, msgAddModelInfo.Custom)
	require.Equal(t, receivedModelInfo.CID, msgAddModelInfo.CID)
}

func queryModelInfo(setup TestSetup, vid int16, pid int16) types.ModelInfo {
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{keeper.QueryModel, fmt.Sprintf("%v", vid), fmt.Sprintf("%v", pid)},
		abci.RequestQuery{},
	)

	var receivedModelInfo types.ModelInfo
	_ = setup.Cdc.UnmarshalJSON(result, &receivedModelInfo)
	return receivedModelInfo
}
