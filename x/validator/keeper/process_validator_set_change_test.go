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

package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func TestValidatorSetChange_ApplyAndReturnValidatorSetUpdates_ForEmpty(t *testing.T) {
	setup := testkeeper.Setup(t)
	updates, err := setup.ValidatorKeeper.ApplyAndReturnValidatorSetUpdates(setup.Ctx)
	require.NoError(t, err)
	require.Equal(t, 0, len(updates))
}

func TestValidatorSetChange_ApplyAndReturnValidatorSetUpdates_ForAddedNewValidator(t *testing.T) {
	setup := testkeeper.Setup(t)

	// create validator
	validator := testkeeper.DefaultValidator()
	setup.ValidatorKeeper.SetValidator(setup.Ctx, validator)
	_ = setup.ValidatorKeeper.SetValidatorByConsAddr(setup.Ctx, validator)

	// ensure last validator power is zero
	lastValPower, _ := setup.ValidatorKeeper.GetLastValidatorPower(setup.Ctx, validator.GetOwner())
	require.Equal(t, types.ZeroPower, lastValPower.Power)

	// check for updates
	updates, err := setup.ValidatorKeeper.ApplyAndReturnValidatorSetUpdates(setup.Ctx)
	require.NoError(t, err)
	require.Equal(t, 1, len(updates))
	require.Equal(t, validator.ABCIValidatorUpdate(), updates[0])

	// ensure last validator record is set
	lastValPower, _ = setup.ValidatorKeeper.GetLastValidatorPower(setup.Ctx, validator.GetOwner())
	require.Equal(t, types.Power, lastValPower.Power)

	// check for updates second time
	updates, err = setup.ValidatorKeeper.ApplyAndReturnValidatorSetUpdates(setup.Ctx)
	require.NoError(t, err)
	require.Equal(t, 0, len(updates))
}

func TestValidatorSetChange_ApplyAndReturnValidatorSetUpdates_TwoValidators(t *testing.T) {
	setup := testkeeper.Setup(t)

	// add 2 validators
	validator1, validator2 := testkeeper.StoreTwoValidators(setup)

	// check for updates
	updates, err := setup.ValidatorKeeper.ApplyAndReturnValidatorSetUpdates(setup.Ctx)
	require.NoError(t, err)
	require.Equal(t, 2, len(updates))
	require.Equal(t, validator1.ABCIValidatorUpdate(), updates[1])
	require.Equal(t, validator2.ABCIValidatorUpdate(), updates[0])

	// ensure last validator record is set
	lastValPower1, _ := setup.ValidatorKeeper.GetLastValidatorPower(setup.Ctx, validator1.GetOwner())
	lastValPower2, _ := setup.ValidatorKeeper.GetLastValidatorPower(setup.Ctx, validator2.GetOwner())
	require.Equal(t, types.Power, lastValPower1.Power)
	require.Equal(t, types.Power, lastValPower2.Power)

	// check for updates second time
	updates, err = setup.ValidatorKeeper.ApplyAndReturnValidatorSetUpdates(setup.Ctx)
	require.NoError(t, err)
	require.Equal(t, 0, len(updates))
}

func TestValidatorSetChange_ApplyAndReturnValidatorSetUpdates_ForJailedValidator(t *testing.T) {
	setup := testkeeper.Setup(t)

	// create validator
	validator := testkeeper.DefaultValidator()
	setup.ValidatorKeeper.SetValidator(setup.Ctx, validator)
	_ = setup.ValidatorKeeper.SetValidatorByConsAddr(setup.Ctx, validator)
	setup.ValidatorKeeper.SetLastValidatorPower(setup.Ctx, types.NewLastValidatorPower(validator.GetOwner()))

	// jail validator
	setup.ValidatorKeeper.Jail(setup.Ctx, validator, "some reason")

	// check for updates
	updates, err := setup.ValidatorKeeper.ApplyAndReturnValidatorSetUpdates(setup.Ctx)
	require.NoError(t, err)
	require.Equal(t, 1, len(updates))
	require.Equal(t, validator.ABCIValidatorUpdateZero(), updates[0])

	// ensure last validator record is zeroed
	lastValPower, _ := setup.ValidatorKeeper.GetLastValidatorPower(setup.Ctx, validator.GetOwner())
	require.Equal(t, types.ZeroPower, lastValPower.Power)

	// check for updates second time
	updates, err = setup.ValidatorKeeper.ApplyAndReturnValidatorSetUpdates(setup.Ctx)
	require.NoError(t, err)
	require.Equal(t, 0, len(updates))
}

func TestValidatorSetChange_ApplyAndReturnValidatorSetUpdates_ForJailedAndUnjailedValidator(t *testing.T) {
	setup := testkeeper.Setup(t)

	// create validator
	validator := testkeeper.DefaultValidator()
	setup.ValidatorKeeper.SetValidator(setup.Ctx, validator)
	_ = setup.ValidatorKeeper.SetValidatorByConsAddr(setup.Ctx, validator)
	setup.ValidatorKeeper.SetLastValidatorPower(setup.Ctx, types.NewLastValidatorPower(validator.GetOwner()))

	// jail validator
	setup.ValidatorKeeper.Jail(setup.Ctx, validator, "some reason")

	// check for updates
	updates, err := setup.ValidatorKeeper.ApplyAndReturnValidatorSetUpdates(setup.Ctx)
	require.NoError(t, err)
	require.Equal(t, 1, len(updates))
	require.Equal(t, validator.ABCIValidatorUpdateZero(), updates[0])

	// ensure last validator record is zeroed
	lastValPower, _ := setup.ValidatorKeeper.GetLastValidatorPower(setup.Ctx, validator.GetOwner())
	require.Equal(t, types.ZeroPower, lastValPower.Power)

	// unjail validator
	setup.ValidatorKeeper.Unjail(setup.Ctx, validator)

	// check for updates
	updates, err = setup.ValidatorKeeper.ApplyAndReturnValidatorSetUpdates(setup.Ctx)
	require.NoError(t, err)
	require.Equal(t, 1, len(updates))
	require.Equal(t, validator.ABCIValidatorUpdate(), updates[0])

	// ensure last validator record is not zero
	lastValPower, _ = setup.ValidatorKeeper.GetLastValidatorPower(setup.Ctx, validator.GetOwner())
	require.Equal(t, types.Power, lastValPower.Power)
}
