package compliance

import (
	"fmt"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(keeper keeper.Keeper, authzKeeper authz.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgAddModelInfo:
			return handleMsgAddModelInfo(ctx, keeper, authzKeeper, msg)
		case types.MsgUpdateModelInfo:
			return handleMsgUpdateModelInfo(ctx, keeper, authzKeeper, msg)
		case types.MsgDeleteModelInfo:
			return handleMsgDeleteModelInfo(ctx, keeper, authzKeeper, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized nameservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgAddModelInfo(ctx sdk.Context, keeper keeper.Keeper, authzKeeper authz.Keeper,
	msg types.MsgAddModelInfo) sdk.Result {
	if keeper.IsModelInfoPresent(ctx, msg.ID) {
		return types.ErrModelInfoAlreadyExists(types.DefaultCodespace).Result()
	}

	if err := checkAuthorized(ctx, authzKeeper, msg.Owner, msg.Signer); err != nil {
		return err.Result()
	}

	modelInfo := types.NewModelInfo(msg.ID, msg.Family, msg.Cert, msg.Owner)

	keeper.SetModelInfo(ctx, modelInfo)

	return sdk.Result{}
}

func handleMsgUpdateModelInfo(ctx sdk.Context, keeper keeper.Keeper, authzKeeper authz.Keeper,
	msg types.MsgUpdateModelInfo) sdk.Result {
	if !keeper.IsModelInfoPresent(ctx, msg.ID) {
		return types.ErrModelInfoDoesNotExist(types.DefaultCodespace).Result()
	}

	modelInfo := keeper.GetModelInfo(ctx, msg.ID)

	if err := checkAuthorized(ctx, authzKeeper, modelInfo.Owner, msg.Signer); err != nil {
		return err.Result()
	}

	modelInfo.Owner = msg.NewOwner
	modelInfo.Family = msg.NewFamily
	modelInfo.Cert = msg.NewCert

	keeper.SetModelInfo(ctx, modelInfo)

	return sdk.Result{}
}

func handleMsgDeleteModelInfo(ctx sdk.Context, keeper keeper.Keeper, authzKeeper authz.Keeper,
	msg types.MsgDeleteModelInfo) sdk.Result {
	if !keeper.IsModelInfoPresent(ctx, msg.ID) {
		return types.ErrModelInfoDoesNotExist(types.DefaultCodespace).Result()
	}

	modelInfo := keeper.GetModelInfo(ctx, msg.ID)

	if err := checkAuthorized(ctx, authzKeeper, modelInfo.Owner, msg.Signer); err != nil {
		return err.Result()
	}

	keeper.DeleteModelInfo(ctx, msg.ID)

	return sdk.Result{}
}

func checkAuthorized(ctx sdk.Context, authzKeeper authz.Keeper, owner sdk.AccAddress, signer sdk.AccAddress) sdk.Error {
	isAuthorized := false

	if signer.Equals(owner) && authzKeeper.HasRole(ctx, signer, authz.Manufacturer) {
		isAuthorized = true
	}

	if authzKeeper.HasRole(ctx, signer, authz.Administrator) {
		isAuthorized = true
	}

	if !isAuthorized {
		return sdk.ErrUnauthorized("tx should be signed either by owner or by administrator")
	}

	return nil
}
