package validator_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		ValidatorList: []types.Validator{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		LastValidatorPowerList: []types.LastValidatorPower{
			{
				ConsensusAddress: "0",
			},
			{
				ConsensusAddress: "1",
			},
		},
		ValidatorSigningInfoList: []types.ValidatorSigningInfo{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		ValidatorMissedBlockBitArrayList: []types.ValidatorMissedBlockBitArray{
			{
				Address: "0",
				Index:   0,
			},
			{
				Address: "1",
				Index:   1,
			},
		},
		ValidatorOwnerList: []types.ValidatorOwner{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ValidatorKeeper(t)
	validator.InitGenesis(ctx, *k, genesisState)
	got := validator.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Len(t, got.ValidatorList, len(genesisState.ValidatorList))
	require.Subset(t, genesisState.ValidatorList, got.ValidatorList)
	require.Len(t, got.LastValidatorPowerList, len(genesisState.LastValidatorPowerList))
	require.Subset(t, genesisState.LastValidatorPowerList, got.LastValidatorPowerList)
	require.Len(t, got.ValidatorSigningInfoList, len(genesisState.ValidatorSigningInfoList))
	require.Subset(t, genesisState.ValidatorSigningInfoList, got.ValidatorSigningInfoList)
	require.Len(t, got.ValidatorMissedBlockBitArrayList, len(genesisState.ValidatorMissedBlockBitArrayList))
	require.Subset(t, genesisState.ValidatorMissedBlockBitArrayList, got.ValidatorMissedBlockBitArrayList)
	require.Len(t, got.ValidatorOwnerList, len(genesisState.ValidatorOwnerList))
	require.Subset(t, genesisState.ValidatorOwnerList, got.ValidatorOwnerList)
	// this line is used by starport scaffolding # genesis/test/assert
}
