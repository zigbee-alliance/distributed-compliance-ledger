package compliance

import (
	"fmt"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(keeper keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgAddModelInfo:
			return handleMsgAddModelInfo(ctx, keeper, msg)
		case types.MsgUpdateModelInfo:
			return handleMsgUpdateModelInfo(ctx, keeper, msg)
		case types.MsgDeleteModelInfo:
			return handleMsgDeleteModelInfo(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized nameservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgAddModelInfo(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgAddModelInfo) sdk.Result {
	if keeper.IsModelInfoPresent(ctx, msg.ID) {
		return types.ErrModelInfoAlreadyExists(types.DefaultCodespace).Result()
	}

	modelInfo := types.NewModelInfo(msg.ID, msg.Family, msg.Cert, msg.Owner)

	keeper.SetModelInfo(ctx, msg.ID, modelInfo)

	return sdk.Result{}
}

func handleMsgUpdateModelInfo(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgUpdateModelInfo) sdk.Result {
	if !keeper.IsModelInfoPresent(ctx, msg.ID) {
		return types.ErrModelInfoDoesNotExist(types.DefaultCodespace).Result()
	}

	modelInfo := keeper.GetModelInfo(ctx, msg.ID)

	if !msg.Owner.Equals(modelInfo.Owner) {
		return sdk.ErrUnauthorized("Incorrect Owner").Result()
	}

	modelInfo.Family = msg.NewFamily
	modelInfo.Cert = msg.NewCert

	keeper.SetModelInfo(ctx, msg.ID, modelInfo)

	return sdk.Result{}
}

func handleMsgDeleteModelInfo(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgDeleteModelInfo) sdk.Result {
	if !keeper.IsModelInfoPresent(ctx, msg.ID) {
		return types.ErrModelInfoDoesNotExist(types.DefaultCodespace).Result()
	}

	modelInfo := keeper.GetModelInfo(ctx, msg.ID)

	if !msg.Owner.Equals(modelInfo.Owner) {
		return sdk.ErrUnauthorized("Incorrect Owner").Result()
	}

	keeper.DeleteModelInfo(ctx, msg.ID)

	return sdk.Result{}
}
