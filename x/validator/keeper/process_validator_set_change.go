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

package keeper

import (
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

// Apply and return accumulated updates to the bonded validator set.
// It gets called once after genesis and at every EndBlock.
//
// Only validators that were added or were removed from the validator set are returned to Tendermint.
func (k Keeper) ApplyAndReturnValidatorSetUpdates(ctx sdk.Context) (updates []abci.ValidatorUpdate, err error) {
	// Iterate over validators.
	k.IterateValidators(ctx, func(validator types.Validator) (stop bool) {
		owner := validator.GetOwner()

		// power on the last height.
		lastValidatorPower, _ := k.GetLastValidatorPower(ctx, owner)

		// if last power was more then 0 and potential power 0 it
		// means that validator was jailed or removed within the block.
		if lastValidatorPower.Power > 0 && validator.GetPower() == 0 {
			updates = append(updates, validator.ABCIValidatorUpdateZero())

			// set validator power on lookup index.
			k.RemoveLastValidatorPower(ctx, owner)

			return false
		}

		// if last power was 0 and potential power more then 0 it means that validator was added in the current block.
		if lastValidatorPower.Power == 0 && validator.GetPower() > 0 {
			updates = append(updates, validator.ABCIValidatorUpdate())

			// set validator power on lookup index.
			k.SetLastValidatorPower(ctx, types.NewLastValidatorPower(owner))

			return false
		}

		return false
	})

	return updates, err
}
