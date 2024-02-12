package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func (k msgServer) AddNocX509RootCert(goCtx context.Context, msg *types.MsgAddNocX509RootCert) (*types.MsgAddNocX509RootCertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAddNocX509RootCertResponse{}, nil
}
