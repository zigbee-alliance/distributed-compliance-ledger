package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgProposeDisableValidator = "propose_disable_validator"

var _ sdk.Msg = &MsgProposeDisableValidator{}

func NewMsgProposeDisableValidator(creator string, address string) *MsgProposeDisableValidator {
	return &MsgProposeDisableValidator{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgProposeDisableValidator) Route() string {
	return RouterKey
}

func (msg *MsgProposeDisableValidator) Type() string {
	return TypeMsgProposeDisableValidator
}

func (msg *MsgProposeDisableValidator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgProposeDisableValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgProposeDisableValidator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = sdk.ConsAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	return nil
}
