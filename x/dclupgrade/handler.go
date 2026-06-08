package dclupgrade

import (
	"fmt"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgProposeUpgrade:
			res, err := msgServer.ProposeUpgrade(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgApproveUpgrade:
			res, err := msgServer.ApproveUpgrade(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRejectUpgrade:
			res, err := msgServer.RejectUpgrade(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
			// this line is used by starport scaffolding # 1
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)

			return nil, errors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
