package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgEnableValidator = "enable_validator"

var _ sdk.Msg = &MsgEnableValidator{}

func NewMsgEnableValidator(creator sdk.ValAddress) *MsgEnableValidator {
	return &MsgEnableValidator{
		Creator: creator.String(),
	}
}

func (msg *MsgEnableValidator) Route() string {
	return RouterKey
}

func (msg *MsgEnableValidator) Type() string {
	return TypeMsgEnableValidator
}

func (msg *MsgEnableValidator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.ValAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sdk.AccAddress(creator)}
}

func (msg *MsgEnableValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEnableValidator) ValidateBasic() error {
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
