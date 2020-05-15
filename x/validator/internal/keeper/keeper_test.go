package keeper

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator/internal/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKeeper_Validator_SetGet(t *testing.T) {
	setup := Setup()

	// check if validator present
	require.False(t, setup.ValidatorKeeper.IsValidatorPresent(setup.Ctx, test_constants.ValidatorAddress1))

	// no validator before its created
	require.Panics(t, func() {
		setup.ValidatorKeeper.GetValidator(setup.Ctx, test_constants.ValidatorAddress1)
	})

	// create validator
	validator := DefaultValidator()
	setup.ValidatorKeeper.SetValidator(setup.Ctx, validator)

	// check if validator present
	require.True(t, setup.ValidatorKeeper.IsValidatorPresent(setup.Ctx, validator.Address))

	// get validator
	receivedValidator := setup.ValidatorKeeper.GetValidator(setup.Ctx, validator.Address)
	require.Equal(t, validator, receivedValidator)

	// get all validators
	validators := setup.ValidatorKeeper.GetAllValidators(setup.Ctx)
	require.Equal(t, 1, len(validators))
	require.Equal(t, validator, validators[0])
}

func TestKeeper_LastValidatorPower_SetGet(t *testing.T) {
	setup := Setup()

	// empty validator power before it set
	receivedValidatorPower := setup.ValidatorKeeper.GetLastValidatorPower(setup.Ctx, test_constants.ValidatorAddress1)
	require.Equal(t, types.ZeroPower, receivedValidatorPower.Power)

	// set validator and power
	validator := DefaultValidator()
	validatorPower := DefaultValidatorPower()

	setup.ValidatorKeeper.SetValidator(setup.Ctx, validator)
	setup.ValidatorKeeper.SetLastValidatorPower(setup.Ctx, validatorPower)

	// get validator power
	receivedValidatorPower = setup.ValidatorKeeper.GetLastValidatorPower(setup.Ctx, validatorPower.ConsensusAddress)
	require.Equal(t, receivedValidatorPower.ConsensusAddress, validatorPower.ConsensusAddress)
	require.Equal(t, types.Power, validatorPower.Power)

	// get all validator powers
	validatorPowers := setup.ValidatorKeeper.GetLastValidatorPowers(setup.Ctx)
	require.Equal(t, 1, len(validatorPowers))
	require.Equal(t, validatorPower.ConsensusAddress, validatorPowers[0].ConsensusAddress)

	// get all last validators
	validators := setup.ValidatorKeeper.GetAllLastValidators(setup.Ctx)
	require.Equal(t, 1, len(validatorPowers))
	require.Equal(t, validator, validators[0])
}
