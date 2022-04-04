package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgDisableValidator = "disable_validator"

var _ sdk.Msg = &MsgDisableValidator{}

func NewMsgDisableValidator(creator sdk.ValAddress) *MsgDisableValidator {
	return &MsgDisableValidator{
		Creator: creator.String(),
	}
}

func (msg *MsgDisableValidator) Route() string {
	return RouterKey
}

func (msg *MsgDisableValidator) Type() string {
	return TypeMsgDisableValidator
}

func (msg *MsgDisableValidator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.ValAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sdk.AccAddress(creator)}
}

func (msg *MsgDisableValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDisableValidator) ValidateBasic() error {
	_, err := sdk.ValAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	return nil
}
