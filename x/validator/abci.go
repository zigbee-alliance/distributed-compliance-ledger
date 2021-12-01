package validator

import (
	"time"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	// nothing special for now
}

// Called every block, update validator set
func EndBlocker(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	return ApplyAndReturnValidatorSetUpdates(ctx, k)
}

// Apply and return accumulated updates to the bonded validator set.
// It gets called once after genesis and at every EndBlock.
//
// Only validators that were added or were removed from the validator set are returned to Tendermint.
func ApplyAndReturnValidatorSetUpdates(ctx sdk.Context, k keeper.Keeper) (updates []abci.ValidatorUpdate) {
	// Iterate over validators.
	k.IterateValidators(ctx, func(validator types.Validator) (stop bool) {
		owner := validator.GetOwner()

		// power on the last height.
		lastValidatorPower, found := k.GetLastValidatorPower(ctx, owner)

		if !found {
			return false
		}

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

	return updates
}
