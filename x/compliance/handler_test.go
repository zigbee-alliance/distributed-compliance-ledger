package compliance

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHandler_HandleAddDeleteModelInfo(t *testing.T) {
	setup := Setup()
	owner := setup.Manufacturer()

	// add new model
	modelInfo := TestMsgAddModelInfo(owner)
	result := setup.Handler(setup.Ctx, modelInfo)
	require.Equal(t, sdk.CodeOK, result.Code)
}

func TestHandler_HandleUpdateModel(t *testing.T) {
	setup := Setup()
	owner := setup.Manufacturer()

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

func TestHandler_OnlyOwnerCanUpdateMessage(t *testing.T) {
	setup := Setup()
	owner := setup.Manufacturer()
	administrator := setup.Administrator()

	// add new model
	msgAddModelInfo := TestMsgAddModelInfo(owner)
	result := setup.Handler(setup.Ctx, msgAddModelInfo)
	require.Equal(t, sdk.CodeOK, result.Code)

	// update existing model
	msgUpdatedModelInfo := TestMsgUpdatedModelInfo(administrator)
	result = setup.Handler(setup.Ctx, msgUpdatedModelInfo)
	require.Equal(t, sdk.CodeUnauthorized, result.Code)

	// owner update existing model
	msgUpdatedModelInfo = TestMsgUpdatedModelInfo(owner)
	result = setup.Handler(setup.Ctx, msgUpdatedModelInfo)
	require.Equal(t, sdk.CodeOK, result.Code)
}
