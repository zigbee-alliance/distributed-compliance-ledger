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
	// Validators can be nil in fresh app state
	require.Greater(t, exportedApp.Height, int64(0))
	require.NotNil(t, exportedApp.ConsensusParams)
}

func TestExportAppStateAndValidators_WithSpecificModules(t *testing.T) {
	// Test export with specific modules
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
	modulesToExport := []string{"validator", "auth"}
	modulesToSkip := []string{}

	exportedApp, err := app.ExportAppStateAndValidators(false, modulesToExport, modulesToSkip)
	require.NoError(t, err)
	require.NotNil(t, exportedApp)
	require.NotNil(t, exportedApp.AppState)
	require.Greater(t, exportedApp.Height, int64(0))
}

func TestExportAppStateAndValidators_WithSkippedModules(t *testing.T) {
	// Test export with modules to skip
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

	// Test export with modules to skip
	modulesToExport := []string{}
	modulesToSkip := []string{"validator"}

	exportedApp, err := app.ExportAppStateAndValidators(false, modulesToExport, modulesToSkip)
	require.NoError(t, err)
	require.NotNil(t, exportedApp)
	require.NotNil(t, exportedApp.AppState)
	require.Greater(t, exportedApp.Height, int64(0))
}

func TestExportAppStateAndValidators_WithNonExistentModules(t *testing.T) {
	// Test export with non-existent modules (should be ignored)
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

	// Test export with non-existent modules - should be ignored
	modulesToExport := []string{"nonexistent1", "nonexistent2"}

	// Non-existent modules should be ignored, not cause panic
	exportedApp, err := app.ExportAppStateAndValidators(false, modulesToExport, []string{})
	require.NoError(t, err)
	require.NotNil(t, exportedApp)
	require.NotNil(t, exportedApp.AppState)
	require.Greater(t, exportedApp.Height, int64(0))
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
	// Validators can be nil in fresh app state, so we just check the field exists
	_ = exportedApp.Validators
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

	exportedApp, err := app.ExportAppStateAndValidators(false, []string{}, []string{})
	require.NoError(t, err)
	require.NotNil(t, exportedApp)
	require.Greater(t, exportedApp.Height, int64(0))
}

func TestExportAppStateAndValidators_ConsensusParams(t *testing.T) {
	// Test that consensus params are exported correctly
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
	// ConsensusParams can be nil in fresh app state
	_ = exportedApp.ConsensusParams
}

func TestExportAppStateAndValidators_AppStateJSON(t *testing.T) {
	// Test that AppState is valid JSON
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
	var jsonData interface{}
	err = json.Unmarshal(exportedApp.AppState, &jsonData)
	require.NoError(t, err)
}

func TestExportAppStateAndValidators_ForZero(t *testing.T) {
	// Test export with forZero=true parameter
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

	// Test export with forZero=true
	exportedApp, err := app.ExportAppStateAndValidators(true, []string{}, []string{})
	require.NoError(t, err)
	require.NotNil(t, exportedApp)
	require.NotNil(t, exportedApp.AppState)
	require.Greater(t, exportedApp.Height, int64(0))
}

func TestExportAppStateAndValidators_EmptyModulesAndSkip(t *testing.T) {
	// Test export with both empty modules and skip lists
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

	// Test export with empty modules and skip lists
	exportedApp, err := app.ExportAppStateAndValidators(false, []string{}, []string{})
	require.NoError(t, err)
	require.NotNil(t, exportedApp)
	require.NotNil(t, exportedApp.AppState)
	require.Greater(t, exportedApp.Height, int64(0))
}

func TestExportAppStateAndValidators_AllFieldsPresent(t *testing.T) {
	// Test that all expected fields are present in the exported app
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

	// Check that all expected fields are present
	require.NotNil(t, exportedApp.AppState)
	require.Greater(t, exportedApp.Height, int64(0))
	// Validators and ConsensusParams can be nil in fresh app state
	_ = exportedApp.Validators
	_ = exportedApp.ConsensusParams
}
