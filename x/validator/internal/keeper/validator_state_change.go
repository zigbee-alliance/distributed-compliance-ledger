package keeper

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	tmtypes "github.com/tendermint/tendermint/types"
	"time"
)

// Calculate the Validators signatures
// Called in each BeginBlock
func (k Keeper) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) {
	// Iterate over all the validators which *should* have signed this block
	// store whether or not they have actually signed it and jail any
	// which have missed too many blocks in a window
	for _, voteInfo := range req.LastCommitInfo.GetVotes() {
		k.HandleValidatorSignature(ctx, voteInfo.Validator.Address, voteInfo.Validator.Power, voteInfo.SignedLastBlock)
	}

	// Iterate through any newly discovered evidence of infraction
	// Slash all validators who contributed to infractions
	for _, evidence := range req.ByzantineValidators {
		switch evidence.Type {
		case tmtypes.ABCIEvidenceTypeDuplicateVote:
			k.HandleDoubleSign(ctx, evidence.Validator.Address, evidence.Height, evidence.Time, evidence.Validator.Power)
		default:
			k.Logger(ctx).Error(fmt.Sprintf("ignored unknown evidence type: %s", evidence.Type))
		}
	}
}

// Calculate the ValidatorUpdates for the current block
// Called in each EndBlock
func (k Keeper) BlockValidatorUpdates(ctx sdk.Context) []abci.ValidatorUpdate {
	return k.ApplyAndReturnValidatorSetUpdates(ctx)
}

// handle a validator signature, must be called once per validator per block
func (k Keeper) HandleValidatorSignature(ctx sdk.Context, addr crypto.Address, power int64, signed bool) {
	logger := k.Logger(ctx)

	height := ctx.BlockHeight()
	consAddr := sdk.ConsAddress(addr)

	if !k.IsValidatorPresent(ctx, consAddr) {
		logger.Error(fmt.Sprintf("Validator by consensus address %s not found", consAddr))
		return
	}

	// fetch signing info
	signInfo := k.GetValidatorSigningInfo(ctx, consAddr)

	// this is a relative index, so it counts blocks the validator *should* have signed
	// will use the 0-value default signing info if not present, except for start height
	index := signInfo.IndexOffset % types.SignedBlocksWindow
	signInfo.IndexOffset++

	// Update signed block bit array & counter
	// This counter just tracks the sum of the bit array
	// That way we avoid needing to read/write the whole array each time
	lastValue := k.GetValidatorMissedBlockBitArray(ctx, consAddr, index)
	missed := !signed
	switch {
	case !lastValue && missed:
		// Array value has changed from not missed to missed, increment counter
		k.SetValidatorMissedBlockBitArray(ctx, consAddr, index, missed)
		signInfo.MissedBlocksCounter++
	case lastValue && !missed:
		// Array value has changed from missed to not missed, decrement counter
		k.SetValidatorMissedBlockBitArray(ctx, consAddr, index, missed)
		signInfo.MissedBlocksCounter--
	default:
		// Array value at this index has not changed, no need to update counter
	}

	if missed {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				slashing.EventTypeLiveness,
				sdk.NewAttribute(slashing.AttributeKeyAddress, consAddr.String()),
				sdk.NewAttribute(slashing.AttributeKeyMissedBlocks, fmt.Sprintf("%d", signInfo.MissedBlocksCounter)),
				sdk.NewAttribute(slashing.AttributeKeyHeight, fmt.Sprintf("%d", height)),
			),
		)

		logger.Info(
			fmt.Sprintf("Absent validator %s at height %d, %d missed, threshold %d", consAddr, height, signInfo.MissedBlocksCounter, types.MinSignedPerWindow))
	}

	minHeight := signInfo.StartHeight + types.SignedBlocksWindow
	maxMissed := types.SignedBlocksWindow - types.MinSignedPerWindow

	// if we are past the minimum height and the validator has missed too many blocks, jail it
	if height > minHeight && signInfo.MissedBlocksCounter > maxMissed {
		validator := k.GetValidator(ctx, consAddr)

		if !validator.IsJailed() {

			// jail the validator
			reason := fmt.Sprintf("Validator \"%v\" passed minimum height \"%v\" and exceeded the maximum number of unsigned blocks \"%v\" within the window in %v",
				consAddr, minHeight, maxMissed, types.SignedBlocksWindow)

			logger.Info(reason)

			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					slashing.EventTypeSlash,
					sdk.NewAttribute(slashing.AttributeKeyAddress, consAddr.String()),
					sdk.NewAttribute(slashing.AttributeKeyPower, fmt.Sprintf("%d", power)),
					sdk.NewAttribute(slashing.AttributeKeyReason, slashing.AttributeValueMissingSignature),
					sdk.NewAttribute(slashing.AttributeKeyJailed, consAddr.String()),
				),
			)
			k.Slash(ctx, consAddr)
			k.Jail(ctx, consAddr, reason)

			// We need to reset the counter & array so that the validator won't be immediately slashed upon rebonding.
			signInfo = signInfo.Reset()
			k.ClearValidatorMissedBlockBitArray(ctx, consAddr)
		} else {
			// Validator already jailed, don't jail again
			logger.Info(
				fmt.Sprintf("Validator %s would have been slashed for downtime, but was either not found in store or already jailed", consAddr),
			)
		}
	}

	// Set the updated signing info
	k.SetValidatorSigningInfo(ctx, signInfo)
}

// handle a validator signing two blocks at the same height
// Zeros validator power and jail it. So validator will be removed from validator set at the end of the block
func (k Keeper) HandleDoubleSign(ctx sdk.Context, addr crypto.Address, infractionHeight int64, timestamp time.Time, power int64) {
	logger := k.Logger(ctx)

	consAddr := sdk.ConsAddress(addr)

	if !k.IsValidatorPresent(ctx, consAddr) {
		logger.Error(fmt.Sprintf("Validator by consensus address %s not found", consAddr))
		return
	}

	// calculate the age of the evidence
	age := ctx.BlockHeader().Time.Sub(timestamp)

	// Reject evidence if the double-sign is too old
	if age > types.MaxEvidenceAge {
		logger.Info(fmt.Sprintf("Ignored double sign from %s at height %d, age of %d past max age of %d",
			consAddr, infractionHeight, age, types.MaxEvidenceAge))
		return
	}

	// double sign confirmed
	reason := fmt.Sprintf("Confirmed double sign from %s at height %d, age of %d", consAddr, infractionHeight, age)
	logger.Info(reason)

	// Slash validator
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			slashing.EventTypeSlash,
			sdk.NewAttribute(slashing.AttributeKeyAddress, consAddr.String()),
			sdk.NewAttribute(slashing.AttributeKeyPower, fmt.Sprintf("%d", power)),
			sdk.NewAttribute(slashing.AttributeKeyReason, slashing.AttributeValueDoubleSign),
			sdk.NewAttribute(slashing.AttributeKeyJailed, consAddr.String()),
		),
	)
	k.Slash(ctx, consAddr)
	k.Jail(ctx, consAddr, reason)
}

// Apply and return accumulated updates to the bonded validator set.
// It gets called once after genesis and at every EndBlock.
//
// Only validators that were added or were removed from the validator set are returned to Tendermint.
func (k Keeper) ApplyAndReturnValidatorSetUpdates(ctx sdk.Context) (updates []abci.ValidatorUpdate) {
	// Iterate over validators
	k.IterateValidators(ctx, func(validator types.Validator) (stop bool) {
		// power on the last height
		lastValidatorPower := k.GetLastValidatorPower(ctx, validator.Address)

		// if last power was more then 0 and potential power 0 it means that validator was jailed or removed within the block.
		if lastValidatorPower.Power > 0 && validator.GetPower() == 0 {
			updates = append(updates, validator.ABCIValidatorUpdateZero())

			// set validator power on lookup index
			k.DeleteLastValidatorPower(ctx, validator.Address)

			return false
		}

		// if last power was 0 and potential power more then 0 it means that validator was added in the current block.
		if lastValidatorPower.Power == 0 && validator.GetPower() > 0 {
			updates = append(updates, validator.ABCIValidatorUpdate())

			// set validator power on lookup index
			k.SetLastValidatorPower(ctx, types.NewLastValidatorPower(validator.Address))

			// init signing info for validator
			signingInfo := types.NewValidatorSigningInfo(validator.GetConsAddress(), ctx.BlockHeight())
			k.SetValidatorSigningInfo(ctx, signingInfo)

			return false
		}

		return false
	})

	return updates
}
