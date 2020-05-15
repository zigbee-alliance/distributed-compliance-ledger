package keeper

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator/internal/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKeeper_ValidatorSigningInfo_SetGet(t *testing.T) {
	setup := Setup()

	height := int64(1)

	// no signing info before its created
	receivedSigningInfo := setup.ValidatorKeeper.GetValidatorSigningInfo(setup.Ctx, test_constants.ValidatorAddress1)
	require.Equal(t, types.ValidatorSigningInfo{}, receivedSigningInfo)

	// set signing info
	signingInfo := types.NewValidatorSigningInfo(test_constants.ValidatorAddress1, height)
	setup.ValidatorKeeper.SetValidatorSigningInfo(setup.Ctx, signingInfo)

	// get signing info
	receivedSigningInfo = setup.ValidatorKeeper.GetValidatorSigningInfo(setup.Ctx, test_constants.ValidatorAddress1)
	require.Equal(t, signingInfo, receivedSigningInfo)
}

func TestKeeper_ValidatorMissedBlock_SetGet(t *testing.T) {
	setup := Setup()

	index := int64(1)
	index2 := int64(2)

	// false for non existing block
	require.False(t, setup.ValidatorKeeper.GetValidatorMissedBlockBitArray(setup.Ctx, test_constants.ValidatorAddress1, index))

	// set two indexes
	setup.ValidatorKeeper.SetValidatorMissedBlockBitArray(setup.Ctx, test_constants.ValidatorAddress1, index, true)
	setup.ValidatorKeeper.SetValidatorMissedBlockBitArray(setup.Ctx, test_constants.ValidatorAddress1, index2, false)

	// check indexes
	require.True(t, setup.ValidatorKeeper.GetValidatorMissedBlockBitArray(setup.Ctx, test_constants.ValidatorAddress1, index))
	require.False(t, setup.ValidatorKeeper.GetValidatorMissedBlockBitArray(setup.Ctx, test_constants.ValidatorAddress1, index2))

	// overwrite index
	setup.ValidatorKeeper.SetValidatorMissedBlockBitArray(setup.Ctx, test_constants.ValidatorAddress1, index2, true)

	// check indexes
	require.True(t, setup.ValidatorKeeper.GetValidatorMissedBlockBitArray(setup.Ctx, test_constants.ValidatorAddress1, index))
	require.True(t, setup.ValidatorKeeper.GetValidatorMissedBlockBitArray(setup.Ctx, test_constants.ValidatorAddress1, index2))

	// iterate over window
	countIndexes := 0
	missedCount := 0
	setup.ValidatorKeeper.IterateValidatorMissedBlockBitArray(setup.Ctx, test_constants.ValidatorAddress1,
		func(index int64, missed bool) (stop bool) {
			countIndexes++
			if missed {
				missedCount++
			}
			return false
		})

	require.Equal(t, 2, countIndexes)
	require.Equal(t, 2, missedCount)

	// clear indexes
	setup.ValidatorKeeper.ClearValidatorMissedBlockBitArray(setup.Ctx, test_constants.ValidatorAddress1)

	// check indexes
	require.False(t, setup.ValidatorKeeper.GetValidatorMissedBlockBitArray(setup.Ctx, test_constants.ValidatorAddress1, index))
	require.False(t, setup.ValidatorKeeper.GetValidatorMissedBlockBitArray(setup.Ctx, test_constants.ValidatorAddress1, index2))
}
