package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
)

func (k msgServer) CertifyModel(goCtx context.Context, msg *types.MsgCertifyModel) (*types.MsgCertifyModelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCertifyModelResponse{}, nil
}
