package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgApproveDisableValidator = "approve_disable_validator"

var _ sdk.Msg = &MsgApproveDisableValidator{}

func NewMsgApproveDisableValidator(creator string, address string) *MsgApproveDisableValidator {
	return &MsgApproveDisableValidator{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgApproveDisableValidator) Route() string {
	return RouterKey
}

func (msg *MsgApproveDisableValidator) Type() string {
	return TypeMsgApproveDisableValidator
}

func (msg *MsgApproveDisableValidator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgApproveDisableValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgApproveDisableValidator) ValidateBasic() error {
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