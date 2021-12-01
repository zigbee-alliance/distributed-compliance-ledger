package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgProposeRevokeAccount = "propose_revoke_account"

var _ sdk.Msg = &MsgProposeRevokeAccount{}

func NewMsgProposeRevokeAccount(signer string, address string) *MsgProposeRevokeAccount {
	return &MsgProposeRevokeAccount{
		Signer:  signer,
		Address: address,
	}
}

func (msg *MsgProposeRevokeAccount) Route() string {
	return RouterKey
}

func (msg *MsgProposeRevokeAccount) Type() string {
	return TypeMsgProposeRevokeAccount
}

func (msg *MsgProposeRevokeAccount) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

func (msg *MsgProposeRevokeAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgProposeRevokeAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}
	return nil
}
