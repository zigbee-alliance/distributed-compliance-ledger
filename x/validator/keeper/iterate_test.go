package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func TestValidatorIterationAndRemove(t *testing.T) {
	setup := testkeeper.Setup(t)
	k := setup.ValidatorKeeper
	ctx := setup.Ctx

	v1, v2 := testkeeper.StoreTwoValidators(setup)
	require.NoError(t, k.SetValidatorByConsAddr(ctx, v1))
	require.NoError(t, k.SetValidatorByConsAddr(ctx, v2))
	k.SetLastValidatorPower(ctx, types.NewLastValidatorPower(v1.GetOwner()))
	k.SetLastValidatorPower(ctx, types.NewLastValidatorPower(v2.GetOwner()))

	require.Len(t, k.GetAllValidator(ctx), 2)
	require.Len(t, k.GetAllLastValidators(ctx), 2)

	var iterated int
	k.IterateValidators(ctx, func(types.Validator) bool {
		iterated++

		return false
	})
	require.Equal(t, 2, iterated)

	// IterateValidators with early stop.
	var stopped int
	k.IterateValidators(ctx, func(types.Validator) bool {
		stopped++

		return true
	})
	require.Equal(t, 1, stopped)

	// RemoveValidator requires the cached consensus pub key, which the
	// StoreTwoValidators fixtures carry.
	k.RemoveValidator(ctx, v1.GetOwner())
	_, found := k.GetValidator(ctx, v1.GetOwner())
	require.False(t, found)

	// Removing a non-existent validator is a no-op.
	k.RemoveValidator(ctx, v1.GetOwner())
}

func TestSetLastValidatorPower_PanicsOnInvalidOwner(t *testing.T) {
	k, ctx := testkeeper.ValidatorKeeper(t, nil)
	require.Panics(t, func() {
		k.SetLastValidatorPower(ctx, types.LastValidatorPower{Owner: "invalid-address"})
	})
}
