package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func (k msgServer) ProposeAddX509RootCert(goCtx context.Context, msg *types.MsgProposeAddX509RootCert) (*types.MsgProposeAddX509RootCertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgProposeAddX509RootCertResponse{}, nil
}
