package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDisableValidator = "disable_validator"

var _ sdk.Msg = &MsgDisableValidator{}

func NewMsgDisableValidator(creator string, address string) *MsgDisableValidator {
	return &MsgDisableValidator{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgDisableValidator) Route() string {
	return RouterKey
}

func (msg *MsgDisableValidator) Type() string {
	return TypeMsgDisableValidator
}

func (msg *MsgDisableValidator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDisableValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDisableValidator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
