package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgApproveUpgrade = "approve_upgrade"

var _ sdk.Msg = &MsgApproveUpgrade{}

func NewMsgApproveUpgrade(creator string, name string) *MsgApproveUpgrade {
	return &MsgApproveUpgrade{
		Creator: creator,
		Name:    name,
	}
}

func (msg *MsgApproveUpgrade) Route() string {
	return RouterKey
}

func (msg *MsgApproveUpgrade) Type() string {
	return TypeMsgApproveUpgrade
}

func (msg *MsgApproveUpgrade) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgApproveUpgrade) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgApproveUpgrade) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	return nil
}
