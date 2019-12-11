package compliance

import (
	"fmt"

	"github.com/askolesov/zb-ledger/x/compliance/internal/keeper"
	"github.com/askolesov/zb-ledger/x/compliance/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(keeper keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgAddModelInfo:
			return handleMsgAddModelInfo(ctx, keeper, msg)
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
