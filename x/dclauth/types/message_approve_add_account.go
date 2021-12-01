package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgApproveAddAccount = "approve_add_account"

var _ sdk.Msg = &MsgApproveAddAccount{}

func NewMsgApproveAddAccount(signer string, address string) *MsgApproveAddAccount {
	return &MsgApproveAddAccount{
		Signer:  signer,
		Address: address,
	}
}

func (msg *MsgApproveAddAccount) Route() string {
	return RouterKey
}

func (msg *MsgApproveAddAccount) Type() string {
	return TypeMsgApproveAddAccount
}

func (msg *MsgApproveAddAccount) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

func (msg *MsgApproveAddAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgApproveAddAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}
	return nil
}
