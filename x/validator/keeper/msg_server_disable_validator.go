package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdkstakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	revokedAccountConstants "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/keeper"
	dclauthTypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func (k msgServer) DisableValidator(goCtx context.Context, msg *types.MsgDisableValidator) (*types.MsgDisableValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, err := sdk.ValAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	// check if message creator has enough rights to propose disable validator
	if !k.dclauthKeeper.HasRole(ctx, sdk.AccAddress(creatorAddr), types.EnableDisableValidatorRole) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"Disable validator transaction should be signed by an account with the %s role",
			types.EnableDisableValidatorRole,
		)
	}

	// check if validator exists
	isFound := k.Keeper.IsValidatorPresent(ctx, creatorAddr)
	if !isFound {
		return nil, sdkstakingtypes.ErrNoValidatorFound
	}

	// check if disabled validator exists
	_, isFound = k.GetDisabledValidator(ctx, msg.Creator)
	if isFound {
		return nil, types.NewErrDisabledValidatorAlreadyExists(msg.Creator)
	}

	creator, err := sdk.ValAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress,
			"Don't convert validator address to ValAddress")
	}

	disabledValidator := types.DisabledValidator{
		Address:             msg.Creator,
		Creator:             sdk.AccAddress(creator).String(),
		DisabledByNodeAdmin: true,
	}

	// Disable validator
	validator, _ := k.GetValidator(ctx, sdk.ValAddress(msg.Creator))
	k.Jail(ctx, validator, "disabled by a node admin")

	// store disabled validator
	k.SetDisabledValidator(ctx, disabledValidator)

	accAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	// get account
	account, ok := k.dclauthKeeper.GetAccountO(ctx, accAddr)

	// check we can get that account or can not
	if !ok {
		return nil, dclauthTypes.ErrAccountDoesNotExist(accAddr)
	}

	// create revoked account record
	revokedAccount := dclauthTypes.NewRevokedAccount(&account, account.Approvals)
	revokedAccount.RevokedReason = revokedAccountConstants.RevokedReasonIsValidatorDisabled
	k.dclauthKeeper.SetRevokedAccount(ctx, *revokedAccount)

	// delete account record
	k.dclauthKeeper.RemoveAccount(ctx, accAddr)

	return &types.MsgDisableValidatorResponse{}, nil
}
