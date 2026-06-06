// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	"encoding/json"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultGenesisState_ProtoCodec(t *testing.T) {
	expectedModules := []string{
		"validator", "pki", "model", "compliance",
		"vendorinfo", "dclauth", "dclupgrade", "dclgenutil",
	}

	cdc := codec.NewProtoCodec(types.NewInterfaceRegistry())

	state := NewDefaultGenesisState(cdc)

	require.NotNil(t, state)

	require.IsType(t, GenesisState{}, state)
	for _, module := range expectedModules {
		_, exists := state[module]
		require.True(t, exists, "Module %s should be present in genesis state", module)
	}

	jsonBytes, err := json.Marshal(state)
	require.NoError(t, err)
	require.NotNil(t, jsonBytes)

	var unmarshaledState GenesisState
	err = json.Unmarshal(jsonBytes, &unmarshaledState)
	require.NoError(t, err)
	require.NotNil(t, unmarshaledState)
}

func TestNewDefaultGenesisState_NilCodec(t *testing.T) {
	cdc := func() codec.JSONCodec { return nil }()

	defer func() {
		r := recover()
		if r == nil {
			t.Error("Expected panic but none occurred")
		}
	}()

	_ = NewDefaultGenesisState(cdc)
}
