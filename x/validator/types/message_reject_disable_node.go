package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRejectDisableNode = "reject_disable_node"

var _ sdk.Msg = &MsgRejectDisableNode{}

func NewMsgRejectDisableNode(creator string, address string, info string, time int64) *MsgRejectDisableNode {
	return &MsgRejectDisableNode{
		Creator: creator,
		Address: address,
		Info:    info,
		Time:    time,
	}
}

func (msg *MsgRejectDisableNode) Route() string {
	return RouterKey
}

func (msg *MsgRejectDisableNode) Type() string {
	return TypeMsgRejectDisableNode
}

func (msg *MsgRejectDisableNode) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRejectDisableNode) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRejectDisableNode) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
