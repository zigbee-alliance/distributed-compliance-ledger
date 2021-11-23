package validator

import (
	"time"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
)

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	// nothing special for now
}

// Called every block, update validator set
func EndBlocker(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	return k.ApplyAndReturnValidatorSetUpdates(ctx)
}

// Apply and return accumulated updates to the bonded validator set.
// It gets called once after genesis and at every EndBlock.
//
// Only validators that were added or were removed from the validator set are returned to Tendermint.
func (k Keeper) ApplyAndReturnValidatorSetUpdates(ctx sdk.Context) (updates []abci.ValidatorUpdate) {
	// Iterate over validators.
	k.IterateValidators(ctx, func(validator types.Validator) (stop bool) {
		// power on the last height.
		lastValidatorPower := k.GetLastValidatorPower(ctx, validator.Address)

		// if last power was more then 0 and potential power 0 it
		// means that validator was jailed or removed within the block.
		if lastValidatorPower.Power > 0 && validator.GetPower() == 0 {
			updates = append(updates, validator.ABCIValidatorUpdateZero())

			// set validator power on lookup index.
			k.DeleteLastValidatorPower(ctx, validator.Address)

			return false
		}

		// if last power was 0 and potential power more then 0 it means that validator was added in the current block.
		if lastValidatorPower.Power == 0 && validator.GetPower() > 0 {
			updates = append(updates, validator.ABCIValidatorUpdate())

			// set validator power on lookup index.
			k.SetLastValidatorPower(ctx, types.NewLastValidatorPower(validator.Address))

			// init signing info for validator.
			signingInfo := types.NewValidatorSigningInfo(validator.GetConsAddress(), ctx.BlockHeight())
			k.SetValidatorSigningInfo(ctx, signingInfo)

			return false
		}

		return false
	})

	return updates
}
