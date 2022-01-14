package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
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
