package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Calculate the ValidatorUpdates for the current block
// Called in each EndBlock
func (k Keeper) BlockValidatorUpdates(ctx sdk.Context) []abci.ValidatorUpdate {
	return k.ApplyAndReturnValidatorSetUpdates(ctx)
}

// Apply and return accumulated updates to the bonded validator set. Also,
// * Updates the active valset as keyed by LastValidatorPowerKey.
// * Updates the total power as keyed by LastTotalPowerKey.
// * Updates validator status' according to updated powers.
// * Updates the fee pool bonded vs not-bonded tokens.
// * Updates relevant indices.
// It gets called once after genesis, another time maybe after genesis transactions,
// then once at every EndBlock.
//
// CONTRACT: Only validators with non-zero power or zero-power that were bonded
// at the previous block height or were removed from the validator set entirely
// are returned to Tendermint.
func (k Keeper) ApplyAndReturnValidatorSetUpdates(ctx sdk.Context) (updates []abci.ValidatorUpdate) {
	// Iterate over validators
	k.IterateValidators(ctx, func(validator types.Validator) (stop bool) {
		// power on the last height
		lastPower := k.GetLastValidatorPower(ctx, validator.OperatorAddress)

		if lastPower == 0 {
			updates = append(updates, validator.ABCIValidatorUpdate())

			// set validator power on lookup index
			k.SetLastValidatorPower(ctx, validator.OperatorAddress, validator.GetWeight())
		}

		return false
	})

	return updates
}
