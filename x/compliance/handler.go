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
			return handleMsgAddDevice(ctx, keeper, msg)
		case types.MsgAddCompliance:
			return handleMsgAddCompliance(ctx, keeper, msg)
		case types.MsgApproveCompliance:
			return handleMsgApproveCompliance(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized nameservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgAddDevice(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgAddModelInfo) sdk.Result {

	if !keeper.IsModelInfoPresent(ctx, msg.ID) {
		return types.ErrDeviceAlreaadyExists(types.DefaultCodespace).Result()
	}

	device := types.NewDevice(msg.Family, msg.Cert, msg.Owner)

	keeper.SetModelInfo(ctx, msg.ID, *device)

	return sdk.Result{}
}

func handleMsgAddCompliance(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgAddCompliance) sdk.Result {

}

func handleMsgApproveCompliance(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgApproveCompliance) sdk.Result {

}
