package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgProposeRevokeAccount = "propose_revoke_account"

var _ sdk.Msg = &MsgProposeRevokeAccount{}

func NewMsgProposeRevokeAccount(signer sdk.AccAddress, address sdk.AccAddress) *MsgProposeRevokeAccount {
	return &MsgProposeRevokeAccount{
		Signer:  signer.String(),
		Address: address.String(),
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
	if msg.Address == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Account Address: it cannot be empty")
	}

	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid Account Address: (%s)", err)
	}

	if msg.Signer == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Signer: it cannot be empty")
	}

	_, err = sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid Signer: (%s)", err)
	}

	return nil
}
