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
	require.False(t, setup.ValidatorKeeper.IsValidatorPresent(setup.Ctx, test_constants.ValAddress1))

	// no validator before its created
	require.Panics(t, func() {
		setup.ValidatorKeeper.GetValidator(setup.Ctx, test_constants.ValAddress1)
	})

	// create validator
	validator := DefaultValidator()
	setup.ValidatorKeeper.SetValidator(setup.Ctx, validator)

	// check if validator present
	require.True(t, setup.ValidatorKeeper.IsValidatorPresent(setup.Ctx, validator.OperatorAddress))

	// get validator
	receivedValidator := setup.ValidatorKeeper.GetValidator(setup.Ctx, validator.OperatorAddress)
	require.Equal(t, validator, receivedValidator)

	// get all validators
	validators, total := setup.ValidatorKeeper.GetAllValidators(setup.Ctx)
	require.Equal(t, 1, total)
	require.Equal(t, 1, len(validators))
	require.Equal(t, validator, validators[0])
}

func TestKeeper_ValidatorByConsAddr_SetGet(t *testing.T) {
	setup := Setup()

	// check if validator present
	require.False(t, setup.ValidatorKeeper.IsValidatorByConsAddrPresent(setup.Ctx, test_constants.ConsensusAddress1))

	// no validator before its created
	require.Panics(t, func() {
		setup.ValidatorKeeper.GetValidatorByConsAddr(setup.Ctx, test_constants.ConsensusAddress1)
	})

	// set validator by consensus address record
	validator := DefaultValidator()
	setup.ValidatorKeeper.SetValidatorByConsAddr(setup.Ctx, validator)

	// check if record present
	require.True(t, setup.ValidatorKeeper.IsValidatorByConsAddrPresent(setup.Ctx, test_constants.ConsensusAddress1))

	// get validator by consensus address for missed validator record
	require.Panics(t, func() {
		setup.ValidatorKeeper.GetValidatorByConsAddr(setup.Ctx, test_constants.ConsensusAddress1)
	})

	// set validator
	setup.ValidatorKeeper.SetValidator(setup.Ctx, validator)

	// check if record present
	require.True(t, setup.ValidatorKeeper.IsValidatorByConsAddrPresent(setup.Ctx, test_constants.ConsensusAddress1))

	// get validator by consensus address
	receivedValidator := setup.ValidatorKeeper.GetValidatorByConsAddr(setup.Ctx, test_constants.ConsensusAddress1)
	require.Equal(t, validator, receivedValidator)
}

func TestKeeper_LastValidatorPower_SetGet(t *testing.T) {
	setup := Setup()

	// empty validator power before it set
	receivedValidatorPower := setup.ValidatorKeeper.GetLastValidatorPower(setup.Ctx, test_constants.ValAddress1)
	require.Equal(t, types.ZeroPower, receivedValidatorPower.Power)

	// set validator and power
	validator := DefaultValidator()
	validatorPower := DefaultValidatorPower()

	setup.ValidatorKeeper.SetValidator(setup.Ctx, validator)
	setup.ValidatorKeeper.SetLastValidatorPower(setup.Ctx, validatorPower)

	// get validator power
	receivedValidatorPower = setup.ValidatorKeeper.GetLastValidatorPower(setup.Ctx, validatorPower.OperatorAddress)
	require.Equal(t, receivedValidatorPower.OperatorAddress, validatorPower.OperatorAddress)
	require.Equal(t, types.Power, validatorPower.Power)

	// get all validator powers
	validatorPowers := setup.ValidatorKeeper.GetLastValidatorPowers(setup.Ctx)
	require.Equal(t, 1, len(validatorPowers))
	require.Equal(t, validatorPower.OperatorAddress, validatorPowers[0].OperatorAddress)

	// get all last validators
	validators := setup.ValidatorKeeper.GetAllLastValidators(setup.Ctx)
	require.Equal(t, 1, len(validatorPowers))
	require.Equal(t, validator, validators[0])
}
