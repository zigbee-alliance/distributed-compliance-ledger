package modelinfo

import (
	"fmt"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(keeper keeper.Keeper, authzKeeper authz.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgAddModelInfo:
			return handleMsgAddModelInfo(ctx, keeper, authzKeeper, msg)
		case types.MsgUpdateModelInfo:
			return handleMsgUpdateModelInfo(ctx, keeper, authzKeeper, msg)
			/*		case type.MsgDeleteModelInfo:
					return handleMsgDeleteModelInfo(ctx, keeper, authzKeeper, msg)*/
		default:
			errMsg := fmt.Sprintf("unrecognized nameservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgAddModelInfo(ctx sdk.Context, keeper keeper.Keeper, authzKeeper authz.Keeper,
	msg types.MsgAddModelInfo) sdk.Result {
	if err := msg.ValidateBasic(); err != nil {
		return err.Result()
	} // TODO: investigate whether we need to call it here explicitly or tandermind/sdk already does it before

	if keeper.IsModelInfoPresent(ctx, msg.VID, msg.PID) {
		return types.ErrModelInfoAlreadyExists(msg.VID, msg.PID).Result()
	}

	if err := checkAddModelRights(ctx, authzKeeper, msg.Signer); err != nil {
		return err.Result()
	}

	modelInfo := types.NewModelInfo(
		msg.VID,
		msg.PID,
		msg.CID,
		msg.Name,
		msg.Signer,
		msg.Description,
		msg.SKU,
		msg.FirmwareVersion,
		msg.HardwareVersion,
		msg.Custom,
		msg.TisOrTrpTestingCompleted,
	)

	keeper.SetModelInfo(ctx, modelInfo)

	return sdk.Result{}
}

func handleMsgUpdateModelInfo(ctx sdk.Context, keeper keeper.Keeper, authzKeeper authz.Keeper,
	msg types.MsgUpdateModelInfo) sdk.Result {
	if err := msg.ValidateBasic(); err != nil {
		return err.Result()
	}

	if !keeper.IsModelInfoPresent(ctx, msg.VID, msg.PID) {
		return types.ErrModelInfoDoesNotExist(msg.VID, msg.PID).Result()
	}

	modelInfo := keeper.GetModelInfo(ctx, msg.VID, msg.PID)

	if err := checkUpdateModelRights(modelInfo.Owner, msg.Signer); err != nil {
		return err.Result()
	}

	CID := modelInfo.CID
	if msg.CID != 0 {
		CID = msg.CID
	}

	description := modelInfo.Description
	if len(msg.Description) != 0 {
		description = msg.Description
	}

	custom := modelInfo.Custom
	if len(msg.Custom) != 0 {
		custom = msg.Custom
	}

	modelInfo = types.NewModelInfo(
		msg.VID,
		msg.PID,
		CID,
		modelInfo.Name,
		msg.Signer,
		description,
		modelInfo.SKU,
		modelInfo.FirmwareVersion,
		modelInfo.HardwareVersion,
		custom,
		msg.TisOrTrpTestingCompleted,
	)

	keeper.SetModelInfo(ctx, modelInfo)

	return sdk.Result{}
}

func handleMsgDeleteModelInfo(ctx sdk.Context, keeper keeper.Keeper, authzKeeper authz.Keeper,
	msg types.MsgDeleteModelInfo) sdk.Result {
	if err := msg.ValidateBasic(); err != nil {
		return err.Result()
	}

	if !keeper.IsModelInfoPresent(ctx, msg.VID, msg.PID) {
		return types.ErrModelInfoDoesNotExist(msg.VID, msg.PID).Result()
	}

	modelInfo := keeper.GetModelInfo(ctx, msg.VID, msg.PID)

	if err := checkUpdateModelRights(modelInfo.Owner, msg.Signer); err != nil {
		return err.Result()
	}

	keeper.DeleteModelInfo(ctx, msg.VID, msg.PID)

	return sdk.Result{}
}

func checkAddModelRights(ctx sdk.Context, authzKeeper authz.Keeper, signer sdk.AccAddress) sdk.Error {
	if !authzKeeper.HasRole(ctx, signer, authz.Vendor) {
		return sdk.ErrUnauthorized("MsgAddModelInfo transaction should be signed by an account with the vendor role")
	}

	return nil
}

func checkUpdateModelRights(owner sdk.AccAddress, signer sdk.AccAddress) sdk.Error {
	if !signer.Equals(owner) {
		return sdk.ErrUnauthorized("MsgUpdateModelInfo tx should be signed by owner")
	}

	return nil
}
