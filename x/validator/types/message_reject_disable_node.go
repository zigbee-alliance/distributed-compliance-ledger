package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRejectDisableValidator = "reject_disable_validator"

var _ sdk.Msg = &MsgRejectDisableValidator{}

func NewMsgRejectDisableValidator(creator sdk.AccAddress, address sdk.ValAddress, info string) *MsgRejectDisableValidator {
	return &MsgRejectDisableValidator{
		Creator: creator.String(),
		Address: address.String(),
		Info:    info,
		Time:    time.Now().Unix(),
	}
}

func (msg *MsgRejectDisableValidator) Route() string {
	return RouterKey
}

func (msg *MsgRejectDisableValidator) Type() string {
	return TypeMsgRejectDisableValidator
}

func (msg *MsgRejectDisableValidator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{creator}
}

func (msg *MsgRejectDisableValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgRejectDisableValidator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
