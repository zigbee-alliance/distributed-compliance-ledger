package keeper

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator/internal/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidatorStateChange_ApplyAndReturnValidatorSetUpdates_ForEmpty(t *testing.T) {
	setup := Setup()
	updates := setup.ValidatorKeeper.ApplyAndReturnValidatorSetUpdates(setup.Ctx)
	require.Equal(t, 0, len(updates))
}

func TestValidatorStateChange_ApplyAndReturnValidatorSetUpdates_ForAddedNewValidator(t *testing.T) {
	setup := Setup()

	// create validator
	validator := DefaultValidator()
	setup.ValidatorKeeper.SetValidator(setup.Ctx, validator)

	// ensure last validator power is zero
	require.Equal(t, types.ZeroPower, setup.ValidatorKeeper.GetLastValidatorPower(setup.Ctx, validator.OperatorAddress).Power)

	// check for updates
	updates := setup.ValidatorKeeper.ApplyAndReturnValidatorSetUpdates(setup.Ctx)
	require.Equal(t, 1, len(updates))
	require.Equal(t, validator.ABCIValidatorUpdate(), updates[0])

	// ensure last validator record is set
	require.Equal(t, types.Power, setup.ValidatorKeeper.GetLastValidatorPower(setup.Ctx, validator.OperatorAddress).Power)

	// check for updates second time
	updates = setup.ValidatorKeeper.ApplyAndReturnValidatorSetUpdates(setup.Ctx)
	require.Equal(t, 0, len(updates))
}

func TestValidatorStateChange_ApplyAndReturnValidatorSetUpdates_TwoValidators(t *testing.T) {
	setup := Setup()

	// add 2 validators
	validator1, validator2 := StoreTwoValidators(setup)

	// check for updates
	updates := setup.ValidatorKeeper.ApplyAndReturnValidatorSetUpdates(setup.Ctx)
	require.Equal(t, 2, len(updates))
	require.Equal(t, validator1.ABCIValidatorUpdate(), updates[0])
	require.Equal(t, validator2.ABCIValidatorUpdate(), updates[1])

	// ensure last validator record is set
	require.Equal(t, types.Power, setup.ValidatorKeeper.GetLastValidatorPower(setup.Ctx, validator1.OperatorAddress).Power)
	require.Equal(t, types.Power, setup.ValidatorKeeper.GetLastValidatorPower(setup.Ctx, validator2.OperatorAddress).Power)

	// check for updates second time
	updates = setup.ValidatorKeeper.ApplyAndReturnValidatorSetUpdates(setup.Ctx)
	require.Equal(t, 0, len(updates))
}
