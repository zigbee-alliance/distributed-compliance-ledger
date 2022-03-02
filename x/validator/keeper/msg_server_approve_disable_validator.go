package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func (k msgServer) ApproveDisableValidator(goCtx context.Context, msg *types.MsgApproveDisableValidator) (*types.MsgApproveDisableValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgApproveDisableValidatorResponse{}, nil
}
