package keeper

import (
	dclcompltypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/compliance"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) dclcompltypes.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ dclcompltypes.MsgServer = msgServer{}
