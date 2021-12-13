package keeper

import (
	"fmt"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey

		dclauthKeeper types.DclauthKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,

	dclauthKeeper types.DclauthKeeper,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,

		dclauthKeeper: dclauthKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// TODO issue 99: makes sense to move to a separate file

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
