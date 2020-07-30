//nolint:testpackage
package modelinfo

//nolint:goimports
import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"testing"
)

func TestHandler_AddModel(t *testing.T) {
	setup := Setup()

	// add new model
	modelInfo := TestMsgAddModelInfo(setup.Vendor)
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

	// try update not present model
	msgUpdatedModelInfo := TestMsgUpdatedModelInfo(setup.Vendor)
	result := setup.Handler(setup.Ctx, msgUpdatedModelInfo)
	require.Equal(t, types.CodeModelInfoDoesNotExist, result.Code)

	// add new model
	msgAddModelInfo := TestMsgAddModelInfo(setup.Vendor)
	result = setup.Handler(setup.Ctx, msgAddModelInfo)
	require.Equal(t, sdk.CodeOK, result.Code)

	// update existing model
	result = setup.Handler(setup.Ctx, msgUpdatedModelInfo)
	require.Equal(t, sdk.CodeOK, result.Code)
}

func TestHandler_OnlyOwnerCanUpdateModel(t *testing.T) {
	setup := Setup()

	// add new model
	msgAddModelInfo := TestMsgAddModelInfo(setup.Vendor)
	result := setup.Handler(setup.Ctx, msgAddModelInfo)
	require.Equal(t, sdk.CodeOK, result.Code)

	for _, role := range []auth.AccountRole{auth.Trustee, auth.TestHouse, auth.Vendor} {
		// store account
		account := auth.NewAccount(testconstants.Address3, testconstants.PubKey3, auth.AccountRoles{role})
		setup.authKeeper.SetAccount(setup.Ctx, account)

		// update existing model by not owner
		msgUpdatedModelInfo := TestMsgUpdatedModelInfo(testconstants.Address3)
		result = setup.Handler(setup.Ctx, msgUpdatedModelInfo)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}

	// owner update existing model
	msgUpdatedModelInfo := TestMsgUpdatedModelInfo(setup.Vendor)
	result = setup.Handler(setup.Ctx, msgUpdatedModelInfo)
	require.Equal(t, sdk.CodeOK, result.Code)
}

func TestHandler_AddModelWithEmptyOptionalFields(t *testing.T) {
	setup := Setup()

	// add new model
	modelInfo := TestMsgAddModelInfo(setup.Vendor)
	modelInfo.CID = 0     // Set empty CID
	modelInfo.Custom = "" // Set empty Custom

	result := setup.Handler(setup.Ctx, modelInfo)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query model
	receivedModelInfo := queryModelInfo(setup, testconstants.VID, testconstants.PID)

	// check
	require.Equal(t, receivedModelInfo.CID, uint16(0))
	require.Equal(t, receivedModelInfo.Custom, "")
}

func TestHandler_AddModelByNonVendor(t *testing.T) {
	setup := Setup()

	for _, role := range []auth.AccountRole{auth.Trustee, auth.TestHouse} {
		// store account
		account := auth.NewAccount(testconstants.Address3, testconstants.PubKey3, auth.AccountRoles{role})
		setup.authKeeper.SetAccount(setup.Ctx, account)

		// add new model
		modelInfo := TestMsgAddModelInfo(testconstants.Address3)
		result := setup.Handler(setup.Ctx, modelInfo)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}
}

func TestHandler_PartiallyUpdateModel(t *testing.T) {
	setup := Setup()

	// add new model
	msgAddModelInfo := TestMsgAddModelInfo(setup.Vendor)
	result := setup.Handler(setup.Ctx, msgAddModelInfo)

	// owner update Description of existing model
	msgUpdatedModelInfo := TestMsgUpdatedModelInfo(setup.Vendor)
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

func queryModelInfo(setup TestSetup, vid uint16, pid uint16) types.ModelInfo {
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{keeper.QueryModel, fmt.Sprintf("%v", vid), fmt.Sprintf("%v", pid)},
		abci.RequestQuery{},
	)

	var receivedModelInfo types.ModelInfo
	_ = setup.Cdc.UnmarshalJSON(result, &receivedModelInfo)

	return receivedModelInfo
}
