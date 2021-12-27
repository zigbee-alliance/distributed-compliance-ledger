package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/types"
)

func (k msgServer) AddTestingResult(goCtx context.Context, msg *types.MsgAddTestingResult) (*types.MsgAddTestingResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAddTestingResultResponse{}, nil
}
