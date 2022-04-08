package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRejectAddAccount = "reject_add_account"

var _ sdk.Msg = &MsgRejectAddAccount{}

func NewMsgRejectAddAccount(signer string, address string) *MsgRejectAddAccount {
	return &MsgRejectAddAccount{
		Signer:  signer,
		Address: address,
	}
}

func (msg *MsgRejectAddAccount) Route() string {
	return RouterKey
}

func (msg *MsgRejectAddAccount) Type() string {
	return TypeMsgRejectAddAccount
}

func (msg *MsgRejectAddAccount) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signer}
}

func (msg *MsgRejectAddAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgRejectAddAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	return nil
}
