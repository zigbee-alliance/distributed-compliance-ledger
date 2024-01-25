package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	dclauthTypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

// jail a validator.
func (k Keeper) Jail(ctx sdk.Context, validator types.Validator, reason string) {
	validator.Power = types.ZeroPower
	validator.Jailed = true
	validator.JailedReason = reason

	k.SetValidator(ctx, validator)
}

// unjail a validator.
func (k Keeper) Unjail(ctx sdk.Context, validator types.Validator) {
	validator.Power = types.Power
	validator.Jailed = false
	validator.JailedReason = ""

	k.SetValidator(ctx, validator)
}

// handle a validator signing two blocks at the same height
// Zeros validator power and jail it. So validator will be removed from validator
// set at the end of the block.
func (k Keeper) HandleDoubleSign(ctx sdk.Context, evidence *evidencetypes.Equivocation) {
	logger := k.Logger(ctx)

	consAddr := evidence.GetConsensusAddress()
	validator, found := k.GetValidatorByConsAddr(ctx, consAddr)

	if !found {
		logger.Error(fmt.Sprintf("Validator by validator address %s not found", consAddr))

		return
	}

	// calculate the age of the evidence.
	infractionHeight := evidence.GetHeight()
	infractionTime := evidence.GetTime()
	ageDuration := ctx.BlockHeader().Time.Sub(infractionTime)
	ageBlocks := ctx.BlockHeader().Height - infractionHeight

	// Reject evidence if the double-sign is too old. Evidence is considered stale
	// if the difference in time and number of blocks is greater than the allowed
	// parameters defined.
	cp := ctx.ConsensusParams()
	if cp != nil && cp.Evidence != nil {
		if ageDuration > cp.Evidence.MaxAgeDuration && ageBlocks > cp.Evidence.MaxAgeNumBlocks {
			logger.Info(
				"ignored equivocation; evidence too old",
				"validator", consAddr,
				"infraction_height", infractionHeight,
				"max_age_num_blocks", cp.Evidence.MaxAgeNumBlocks,
				"infraction_time", infractionTime,
				"max_age_duration", cp.Evidence.MaxAgeDuration,
			)

			return
		}
	}

	// double sign confirmed.
	reason := fmt.Sprintf("Confirmed double sign from %s at height %d, age of duration %d, age of blocks %d", consAddr, infractionHeight, ageDuration, ageBlocks)
	logger.Info(reason)

	// Slash validator.
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			slashingtypes.EventTypeSlash,
			sdk.NewAttribute(slashingtypes.AttributeKeyAddress, consAddr.String()),
			sdk.NewAttribute(slashingtypes.AttributeKeyPower, fmt.Sprintf("%d", evidence.GetValidatorPower())), //nolint:perfsprint
			sdk.NewAttribute(slashingtypes.AttributeKeyReason, slashingtypes.AttributeValueDoubleSign),
			sdk.NewAttribute(slashingtypes.AttributeKeyJailed, consAddr.String()),
		),
	)

	// Jail the validator.
	if !validator.IsJailed() {
		k.Jail(ctx, validator, reason)
	}
	// Revoked Account
	valAddr, err := sdk.ValAddressFromBech32(validator.Owner)
	if err != nil {
		logger.Info("Error:", sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err))
	}

	accAddr := sdk.AccAddress(valAddr)

	// Move account to entity revoked account
	revokedAccount, err := k.dclauthKeeper.AddAccountToRevokedAccount(
		ctx, accAddr, nil, dclauthTypes.RevokedAccount_MaliciousValidator) //nolint:nosnakecase
	if err != nil {
		logger.Info("Error:", err)
	} else {
		k.dclauthKeeper.SetRevokedAccount(ctx, *revokedAccount)
		// delete account record
		k.dclauthKeeper.RemoveAccount(ctx, accAddr)
	}

	logger.Info("Error:", err)
}
