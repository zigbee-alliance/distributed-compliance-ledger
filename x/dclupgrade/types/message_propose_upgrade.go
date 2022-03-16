package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgProposeUpgrade = "propose_upgrade"

var _ sdk.Msg = &MsgProposeUpgrade{}

func NewMsgProposeUpgrade(creator string, plan Plan, info string) *MsgProposeUpgrade {
	return &MsgProposeUpgrade{
		Creator: creator,
		Plan:    plan,
		Info:    info,
		Time:    time.Now().Unix(),
	}
}

func (msg *MsgProposeUpgrade) Route() string {
	return RouterKey
}

func (msg *MsgProposeUpgrade) Type() string {
	return TypeMsgProposeUpgrade
}

func (msg *MsgProposeUpgrade) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{creator}
}

func (msg *MsgProposeUpgrade) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgProposeUpgrade) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	err = msg.Plan.ValidateBasic()
	if err != nil {
		return err
	}

	return nil
}
