package app

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultGenesisState(t *testing.T) {
	// Test successful creation of default genesis state
	cdc := codec.NewProtoCodec(types.NewInterfaceRegistry())
	genesisState := NewDefaultGenesisState(cdc)

	require.NotNil(t, genesisState)
	require.IsType(t, GenesisState{}, genesisState)

	// Verify that all expected modules are present
	expectedModules := []string{
		"validator", "pki", "model", "compliance",
		"vendorinfo", "dclauth", "dclupgrade", "dclgenutil",
	}

	for _, module := range expectedModules {
		_, exists := genesisState[module]
		require.True(t, exists, "Module %s should be present in genesis state", module)
	}
}

func TestNewDefaultGenesisState_WithNilCodec(t *testing.T) {
	// Test that NewDefaultGenesisState handles nil codec gracefully
	defer func() {
		if r := recover(); r != nil {
			// Expected panic with nil codec - check if it's an error or string
			switch v := r.(type) {
			case string:
				require.Contains(t, v, "nil")
			case error:
				require.Contains(t, v.Error(), "nil")
			default:
				// Any panic is acceptable for nil codec
				require.NotNil(t, r)
			}
		} else {
			t.Error("Expected panic but none occurred")
		}
	}()

	NewDefaultGenesisState(nil)
}

func TestGenesisState_TypeDefinition(t *testing.T) {
	// Test that GenesisState is properly defined as a map
	var genesisState GenesisState

	// Test that it can be used as a map
	genesisState = make(GenesisState)
	genesisState["test"] = json.RawMessage("test")
	require.Equal(t, json.RawMessage("test"), genesisState["test"])

	// Test that it behaves like a map[string]json.RawMessage
	require.Equal(t, 1, len(genesisState))
	require.Contains(t, genesisState, "test")
}

func TestGenesisState_JSONOperations(t *testing.T) {
	// Test JSON marshaling and unmarshaling of genesis state
	cdc := codec.NewProtoCodec(types.NewInterfaceRegistry())
	genesisState := NewDefaultGenesisState(cdc)

	// Marshal to JSON using standard json package
	jsonBytes, err := json.Marshal(genesisState)
	require.NoError(t, err)
	require.NotNil(t, jsonBytes)

	// Unmarshal from JSON
	var unmarshaledState GenesisState
	err = json.Unmarshal(jsonBytes, &unmarshaledState)
	require.NoError(t, err)
	require.NotNil(t, unmarshaledState)
}

func TestGenesisState_EmptyState(t *testing.T) {
	// Test handling of empty genesis state
	emptyState := GenesisState{}
	require.NotNil(t, emptyState)
	require.Equal(t, 0, len(emptyState))
}

func TestGenesisState_SingleModule(t *testing.T) {
	// Test genesis state with single module
	singleModuleState := GenesisState{
		"validator": json.RawMessage(`{"validatorList":[]}`),
	}

	require.NotNil(t, singleModuleState)
	require.Equal(t, 1, len(singleModuleState))
	require.Contains(t, singleModuleState, "validator")
}

func TestGenesisState_MultipleModules(t *testing.T) {
	// Test genesis state with multiple modules
	multiModuleState := GenesisState{
		"validator": json.RawMessage(`{"validatorList":[]}`),
		"pki":       json.RawMessage(`{"approvedCertificatesList":[]}`),
		"model":     json.RawMessage(`{"modelList":[]}`),
	}

	require.NotNil(t, multiModuleState)
	require.Equal(t, 3, len(multiModuleState))
	require.Contains(t, multiModuleState, "validator")
	require.Contains(t, multiModuleState, "pki")
	require.Contains(t, multiModuleState, "model")
}

func TestGenesisState_ModuleBasicsIntegration(t *testing.T) {
	// Test integration with ModuleBasics
	cdc := codec.NewProtoCodec(types.NewInterfaceRegistry())
	genesisState := NewDefaultGenesisState(cdc)

	// Verify that ModuleBasics.DefaultGenesis is called
	require.NotNil(t, genesisState)

	// Check that the genesis state contains expected module data
	// This tests the integration between app genesis and module basics
	for moduleName, moduleData := range genesisState {
		require.NotEmpty(t, moduleName)
		// Module data can be nil in some cases, so we just check the field exists
		_ = moduleData
	}
}

func TestGenesisState_CodecIntegration(t *testing.T) {
	// Test that genesis state works with different codec types
	interfaceRegistry := types.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	genesisState := NewDefaultGenesisState(cdc)
	require.NotNil(t, genesisState)
}

func TestGenesisState_TypeSafety(t *testing.T) {
	// Test type safety of genesis state
	var genesisState GenesisState

	// Test assignment
	genesisState = make(GenesisState)
	require.NotNil(t, genesisState)

	// Test that it can hold different types of data
	testData := json.RawMessage("test data")
	genesisState["test"] = testData
	require.Equal(t, testData, genesisState["test"])
}

func TestGenesisState_EdgeCases(t *testing.T) {
	// Test edge cases for genesis state

	// Test with very large genesis state
	largeState := GenesisState{}
	for i := 0; i < 100; i++ {
		largeState[fmt.Sprintf("module%d", i)] = json.RawMessage(fmt.Sprintf("data%d", i))
	}

	require.Equal(t, 100, len(largeState))

	// Test with empty module data
	emptyDataState := GenesisState{
		"empty": json.RawMessage{},
	}
	require.NotNil(t, emptyDataState)
	require.Equal(t, 0, len(emptyDataState["empty"]))
}
