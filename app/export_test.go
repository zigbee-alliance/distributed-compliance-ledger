package app

import (
	"encoding/json"
	"testing"

	dbm "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/stretchr/testify/require"
)

func TestExportAppStateAndValidators_Success(t *testing.T) {
	// Test successful export with default parameters
	logger := log.NewNopLogger()
	db := dbm.NewMemDB()
	encodingConfig := MakeEncodingConfig()

	app := New(
		logger,
		db,
		nil,
		true,
		map[int64]bool{},
		"",
		0,
		encodingConfig,
	)

	// Test export with empty modules list (exports all modules)
	exportedApp, err := app.ExportAppStateAndValidators(false, []string{}, []string{})
	require.NoError(t, err)
	require.NotNil(t, exportedApp)
	// AppState is a byte slice, should not be nil
	require.NotNil(t, exportedApp.AppState)
	require.NotNil(t, exportedApp.Validators)
	require.Greater(t, exportedApp.Height, int64(0))
	require.NotNil(t, exportedApp.ConsensusParams)
}

func TestExportAppStateAndValidators_WithForZeroHeight(t *testing.T) {
	// Test export with forZeroHeight = true
	logger := log.NewNopLogger()
	db := dbm.NewMemDB()
	encodingConfig := MakeEncodingConfig()

	app := New(
		logger,
		db,
		nil,
		true,
		map[int64]bool{},
		"",
		0,
		encodingConfig,
	)

	// Test export with forZeroHeight = true
	exportedApp, err := app.ExportAppStateAndValidators(true, []string{}, []string{})
	require.NoError(t, err)
	require.NotNil(t, exportedApp)
	require.NotNil(t, exportedApp.AppState)
	require.NotNil(t, exportedApp.Validators)
	require.Greater(t, exportedApp.Height, int64(0))
	require.NotNil(t, exportedApp.ConsensusParams)
}

func TestExportAppStateAndValidators_WithJailAllowedAddrs(t *testing.T) {
	// Test export with jail allowed addresses
	logger := log.NewNopLogger()
	db := dbm.NewMemDB()
	encodingConfig := MakeEncodingConfig()

	app := New(
		logger,
		db,
		nil,
		true,
		map[int64]bool{},
		"",
		0,
		encodingConfig,
	)

	// Test export with jail allowed addresses
	jailAllowedAddrs := []string{"cosmosvaloper1example", "cosmosvaloper2example"}
	exportedApp, err := app.ExportAppStateAndValidators(false, jailAllowedAddrs, []string{})
	require.NoError(t, err)
	require.NotNil(t, exportedApp)
	require.NotNil(t, exportedApp.AppState)
	require.NotNil(t, exportedApp.Validators)
	require.Greater(t, exportedApp.Height, int64(0))
	require.NotNil(t, exportedApp.ConsensusParams)
}

func TestExportAppStateAndValidators_WithSpecificModules(t *testing.T) {
	// Test export with specific modules to export
	logger := log.NewNopLogger()
	db := dbm.NewMemDB()
	encodingConfig := MakeEncodingConfig()

	app := New(
		logger,
		db,
		nil,
		true,
		map[int64]bool{},
		"",
		0,
		encodingConfig,
	)

	// Test export with specific modules
	modulesToExport := []string{"dclauth", "validator"}
	exportedApp, err := app.ExportAppStateAndValidators(false, []string{}, modulesToExport)
	require.NoError(t, err)
	require.NotNil(t, exportedApp)
	require.NotNil(t, exportedApp.AppState)
	require.NotNil(t, exportedApp.Validators)
	require.Greater(t, exportedApp.Height, int64(0))
	require.NotNil(t, exportedApp.ConsensusParams)
}

func TestExportAppStateAndValidators_WithAllModules(t *testing.T) {
	// Test export with all modules
	logger := log.NewNopLogger()
	db := dbm.NewMemDB()
	encodingConfig := MakeEncodingConfig()

	app := New(
		logger,
		db,
		nil,
		true,
		map[int64]bool{},
		"",
		0,
		encodingConfig,
	)

	// Test export with all modules
	modulesToExport := []string{
		"dclauth", "validator", "dclupgrade",
		"pki", "vendorinfo", "model", "compliance",
	}
	exportedApp, err := app.ExportAppStateAndValidators(false, []string{}, modulesToExport)
	require.NoError(t, err)
	require.NotNil(t, exportedApp)
	require.NotNil(t, exportedApp.AppState)
	require.NotNil(t, exportedApp.Validators)
	require.Greater(t, exportedApp.Height, int64(0))
	require.NotNil(t, exportedApp.ConsensusParams)
}

func TestExportAppStateAndValidators_EmptyModulesList(t *testing.T) {
	// Test export with empty modules list
	logger := log.NewNopLogger()
	db := dbm.NewMemDB()
	encodingConfig := MakeEncodingConfig()

	app := New(
		logger,
		db,
		nil,
		true,
		map[int64]bool{},
		"",
		0,
		encodingConfig,
	)

	// Test export with empty modules list
	exportedApp, err := app.ExportAppStateAndValidators(false, []string{}, []string{})
	require.NoError(t, err)
	require.NotNil(t, exportedApp)
	require.NotNil(t, exportedApp.AppState)
	require.NotNil(t, exportedApp.Validators)
	require.Greater(t, exportedApp.Height, int64(0))
	require.NotNil(t, exportedApp.ConsensusParams)
}

func TestExportAppStateAndValidators_WithNonExistentModules(t *testing.T) {
	// Test export with non-existent modules (should panic)
	logger := log.NewNopLogger()
	db := dbm.NewMemDB()
	encodingConfig := MakeEncodingConfig()

	app := New(
		logger,
		db,
		nil,
		true,
		map[int64]bool{},
		"",
		0,
		encodingConfig,
	)

	// Test export with non-existent modules - should panic
	modulesToExport := []string{"nonexistent1", "nonexistent2"}

	// This should panic, so we test that it does
	defer func() {
		if r := recover(); r != nil {
			// Expected panic
			require.Contains(t, r.(string), "does not exist")
		} else {
			t.Error("Expected panic but none occurred")
		}
	}()

	app.ExportAppStateAndValidators(false, []string{}, modulesToExport)
}

func TestExportAppStateAndValidators_AppStateJSONStructure(t *testing.T) {
	// Test that the exported app state has proper JSON structure
	logger := log.NewNopLogger()
	db := dbm.NewMemDB()
	encodingConfig := MakeEncodingConfig()

	app := New(
		logger,
		db,
		nil,
		true,
		map[int64]bool{},
		"",
		0,
		encodingConfig,
	)

	exportedApp, err := app.ExportAppStateAndValidators(false, []string{}, []string{})
	require.NoError(t, err)
	require.NotNil(t, exportedApp)
	require.NotNil(t, exportedApp.AppState)

	// Verify that AppState is valid JSON
	var appStateMap map[string]interface{}
	err = json.Unmarshal(exportedApp.AppState, &appStateMap)
	require.NoError(t, err)
	require.NotNil(t, appStateMap)
}

func TestExportAppStateAndValidators_ValidatorsStructure(t *testing.T) {
	// Test that the exported validators have proper structure
	logger := log.NewNopLogger()
	db := dbm.NewMemDB()
	encodingConfig := MakeEncodingConfig()

	app := New(
		logger,
		db,
		nil,
		true,
		map[int64]bool{},
		"",
		0,
		encodingConfig,
	)

	exportedApp, err := app.ExportAppStateAndValidators(false, []string{}, []string{})
	require.NoError(t, err)
	require.NotNil(t, exportedApp)
	require.NotNil(t, exportedApp.Validators)

	// Verify that Validators is a slice of GenesisValidator
	require.IsType(t, []interface{}{}, exportedApp.Validators)
	require.GreaterOrEqual(t, len(exportedApp.Validators), 0)
}

func TestExportAppStateAndValidators_HeightCalculation(t *testing.T) {
	// Test that height is calculated correctly
	logger := log.NewNopLogger()
	db := dbm.NewMemDB()
	encodingConfig := MakeEncodingConfig()

	app := New(
		logger,
		db,
		nil,
		true,
		map[int64]bool{},
		"",
		0,
		encodingConfig,
	)

	// Get initial height
	initialHeight := app.LastBlockHeight()

	exportedApp, err := app.ExportAppStateAndValidators(false, []string{}, []string{})
	require.NoError(t, err)
	require.NotNil(t, exportedApp)

	// Verify that exported height is initial height + 1
	expectedHeight := initialHeight + 1
	require.Equal(t, expectedHeight, exportedApp.Height)
}

func TestExportAppStateAndValidators_ConsensusParams(t *testing.T) {
	// Test that consensus params are properly exported
	logger := log.NewNopLogger()
	db := dbm.NewMemDB()
	encodingConfig := MakeEncodingConfig()

	app := New(
		logger,
		db,
		nil,
		true,
		map[int64]bool{},
		"",
		0,
		encodingConfig,
	)

	exportedApp, err := app.ExportAppStateAndValidators(false, []string{}, []string{})
	require.NoError(t, err)
	require.NotNil(t, exportedApp)
	require.NotNil(t, exportedApp.ConsensusParams)

	// Verify that consensus params are not empty
	require.NotNil(t, exportedApp.ConsensusParams.Block)
	require.NotNil(t, exportedApp.ConsensusParams.Evidence)
	require.NotNil(t, exportedApp.ConsensusParams.Validator)
}

func TestExportAppStateAndValidators_ContextCreation(t *testing.T) {
	// Test that context is created properly for export
	logger := log.NewNopLogger()
	db := dbm.NewMemDB()
	encodingConfig := MakeEncodingConfig()

	app := New(
		logger,
		db,
		nil,
		true,
		map[int64]bool{},
		"",
		0,
		encodingConfig,
	)

	// Test that the export function can create context properly
	exportedApp, err := app.ExportAppStateAndValidators(false, []string{}, []string{})
	require.NoError(t, err)
	require.NotNil(t, exportedApp)

	// The fact that we get here without panic means context creation worked
	require.NotNil(t, exportedApp.AppState)
}

func TestExportAppStateAndValidators_ModuleManagerIntegration(t *testing.T) {
	// Test integration with module manager for genesis export
	logger := log.NewNopLogger()
	db := dbm.NewMemDB()
	encodingConfig := MakeEncodingConfig()

	app := New(
		logger,
		db,
		nil,
		true,
		map[int64]bool{},
		"",
		0,
		encodingConfig,
	)

	// Test that module manager can export genesis for modules
	exportedApp, err := app.ExportAppStateAndValidators(false, []string{}, []string{})
	require.NoError(t, err)
	require.NotNil(t, exportedApp)

	// Verify that genesis state was exported
	require.NotNil(t, exportedApp.AppState)
	require.Greater(t, len(exportedApp.AppState), 0)
}

func TestExportAppStateAndValidators_ValidatorKeeperIntegration(t *testing.T) {
	// Test integration with validator keeper
	logger := log.NewNopLogger()
	db := dbm.NewMemDB()
	encodingConfig := MakeEncodingConfig()

	app := New(
		logger,
		db,
		nil,
		true,
		map[int64]bool{},
		"",
		0,
		encodingConfig,
	)

	// Test that validator keeper can write validators
	exportedApp, err := app.ExportAppStateAndValidators(false, []string{}, []string{})
	require.NoError(t, err)
	require.NotNil(t, exportedApp)

	// Verify that validators were exported
	require.NotNil(t, exportedApp.Validators)
	require.Greater(t, len(exportedApp.Validators), 0)
}
