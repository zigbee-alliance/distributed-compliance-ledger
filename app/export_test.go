package app

import (
	"encoding/json"
	"testing"

	dbm "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/stretchr/testify/require"
)

func TestExport_AppStateAndValidators(t *testing.T) {
	positiveTests := []struct {
		name            string
		modulesToExport []string
	}{
		{
			name:            "export_all_modules",
			modulesToExport: []string{"dclupgrade", "vendorinfo", "compliance", "params", "upgrade", "dclgenutil", "model", "dclauth", "validator", "pki"},
		},
		{
			name:            "export_specific_modules",
			modulesToExport: []string{"dclauth", "validator", "pki"},
		},
		{
			name:            "export_without_modules",
			modulesToExport: []string{},
		},
	}

	negativeTests := []struct {
		name            string
		modulesToExport []string
		expectPanic     bool
		err             error
	}{
		{
			name:            "export_non-existent_modules",
			modulesToExport: []string{"non_existent1", "non_existent2"},
			expectPanic:     true,
			err:             nil,
		},
		{
			name:            "export_with_non-existent_modules",
			modulesToExport: []string{"vendorinfo", "non_existent1"},
			expectPanic:     true,
			err:             nil,
		},
	}

	forZeroHeight := false
	jailAllowedAddrs := []string{}
	var jsonData any

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

	for _, tc := range positiveTests {
		t.Run(tc.name, func(t *testing.T) {
			exportedApp, err := app.ExportAppStateAndValidators(forZeroHeight, jailAllowedAddrs, tc.modulesToExport)
			require.NoError(t, err)
			require.NotNil(t, exportedApp)

			// AppState should be non-nil and valid JSON
			require.NotNil(t, exportedApp.AppState)

			require.NoError(t, json.Unmarshal(exportedApp.AppState, &jsonData))

			// Height should be positive (last height + 1)
			require.Greater(t, exportedApp.Height, int64(0))

			// Validators should be nil
			require.Nil(t, exportedApp.Validators)

			// ConsensusParams should be non-nil
			require.NotNil(t, exportedApp.ConsensusParams)
		})
	}

	for _, tc := range negativeTests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expectPanic {
				defer func() {
					r := recover()
					if (r != nil) != tc.expectPanic {
						t.Error("Expected panic but none occurred")
					}
				}()
			}

			exportedApp, err := app.ExportAppStateAndValidators(forZeroHeight, jailAllowedAddrs, tc.modulesToExport)
			require.Nil(t, exportedApp)
			require.Equal(t, tc.err, err)
		})
	}
}
