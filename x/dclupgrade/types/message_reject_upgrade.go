package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRejectUpgrade = "reject_upgrade"

var _ sdk.Msg = &MsgRejectUpgrade{}

func NewMsgRejectUpgrade(creator string, name string) *MsgRejectUpgrade {
	return &MsgRejectUpgrade{
		Creator: creator,
		Name:    name,
	}
}

func (msg *MsgRejectUpgrade) Route() string {
	return RouterKey
}

func (msg *MsgRejectUpgrade) Type() string {
	return TypeMsgRejectUpgrade
}

func (msg *MsgRejectUpgrade) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRejectUpgrade) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRejectUpgrade) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
