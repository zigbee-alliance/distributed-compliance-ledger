package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgApproveAddAccount = "approve_add_account"

var _ sdk.Msg = &MsgApproveAddAccount{}

func NewMsgApproveAddAccount(signer sdk.AccAddress, address sdk.AccAddress) *MsgApproveAddAccount {
	return &MsgApproveAddAccount{
		Signer:  signer.String(),
		Address: address.String(),
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
	if m.Address.Empty() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Account Address: it cannot be empty")
	}

	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid Account Address: (%s)", err)
	}

	if m.Signer.Empty() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Signer: it cannot be empty")
	}

	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid Signer: (%s)", err)
	}

	return nil
}
