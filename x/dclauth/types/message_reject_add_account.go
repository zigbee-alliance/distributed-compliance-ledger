package types

import (
	"time"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRejectAddAccount = "reject_add_account"

var _ sdk.Msg = &MsgRejectAddAccount{}

func NewMsgRejectAddAccount(signer sdk.AccAddress, address sdk.AccAddress, info string) *MsgRejectAddAccount {
	return &MsgRejectAddAccount{
		Signer:  signer.String(),
		Address: address.String(),
		Info:    info,
		Time:    time.Now().Unix(),
	}
}

func (msg *MsgRejectAddAccount) Route() string {
	return RouterKey
}

func (msg *MsgRejectAddAccount) Type() string {
	return TypeMsgRejectAddAccount
}

func (msg *MsgRejectAddAccount) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signer}
}

func (msg *MsgRejectAddAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgRejectAddAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	return nil
}
