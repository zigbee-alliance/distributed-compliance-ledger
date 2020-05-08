package keeper

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// Calculate the ValidatorUpdates for the current block
// Called in each EndBlock
func (k Keeper) BlockValidatorUpdates(ctx sdk.Context) []abci.ValidatorUpdate {
	return k.ApplyAndReturnValidatorSetUpdates(ctx)
}

// Apply and return accumulated updates to the bonded validator set.
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
		lastValidatorPower := k.GetLastValidatorPower(ctx, validator.OperatorAddress)

		// if last power was 0 it means that validator was added in the current block. additionally ensure that potential power more than 0
		if lastValidatorPower.Power == 0 && validator.GetPower() > 0 {
			updates = append(updates, validator.ABCIValidatorUpdate())

			// set validator power on lookup index
			k.SetLastValidatorPower(ctx, types.NewLastValidatorPower(validator.OperatorAddress))
		}

		return false
	})

	return updates
}
