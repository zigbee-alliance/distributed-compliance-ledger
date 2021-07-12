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

//nolint:testpackage
package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/internal/types"
)

func TestKeeper_Validator_SetGet(t *testing.T) {
	setup := Setup()

	// check if validator present
	require.False(t, setup.ValidatorKeeper.IsValidatorPresent(setup.Ctx, testconstants.ValidatorAddress1))

	// no validator before its created
	require.Panics(t, func() {
		setup.ValidatorKeeper.GetValidator(setup.Ctx, testconstants.ValidatorAddress1)
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
	receivedValidatorPower := setup.ValidatorKeeper.GetLastValidatorPower(setup.Ctx, testconstants.ValidatorAddress1)
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
