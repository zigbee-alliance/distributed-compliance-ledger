package model_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		VendorProductsList: []types.VendorProducts{
			{
				Vid: 0,
			},
			{
				Vid: 1,
			},
		},
		ModelList: []types.Model{
			{
				Vid: 0,
				Pid: 0,
			},
			{
				Vid: 1,
				Pid: 1,
			},
		},
		ModelVersionList: []types.ModelVersion{
			{
				Vid:             0,
				Pid:             0,
				SoftwareVersion: 0,
			},
			{
				Vid:             1,
				Pid:             1,
				SoftwareVersion: 1,
			},
		},
		ModelVersionsList: []types.ModelVersions{
			{
				Vid: 0,
				Pid: 0,
			},
			{
				Vid: 1,
				Pid: 1,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ModelKeeper(t, nil)
	model.InitGenesis(ctx, *k, genesisState)
	got := model.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.ElementsMatch(t, genesisState.VendorProductsList, got.VendorProductsList)
	require.ElementsMatch(t, genesisState.ModelList, got.ModelList)
	require.ElementsMatch(t, genesisState.ModelVersionList, got.ModelVersionList)
	require.ElementsMatch(t, genesisState.ModelVersionsList, got.ModelVersionsList)
	// this line is used by starport scaffolding # genesis/test/assert
}
