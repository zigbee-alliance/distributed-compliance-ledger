package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateValidator{}

func NewMsgCreateValidator(signer string, pubKey string, description *Description) *MsgCreateValidator {
	return &MsgCreateValidator{
		Signer:      signer,
		PubKey:      pubKey,
		Description: description,
	}
}

func (msg *MsgCreateValidator) Route() string {
	return RouterKey
}

func (msg *MsgCreateValidator) Type() string {
	return "CreateValidator"
}

func (msg *MsgCreateValidator) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

func (msg *MsgCreateValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateValidator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}
	return nil
}
