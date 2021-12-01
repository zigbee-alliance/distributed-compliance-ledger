package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgProposeAddAccount = "propose_add_account"

var _ sdk.Msg = &MsgProposeAddAccount{}

func NewMsgProposeAddAccount(signer string, address string, pubKey string, roles []string, vendorID uint64) *MsgProposeAddAccount {
	return &MsgProposeAddAccount{
		Signer:   signer,
		Address:  address,
		PubKey:   pubKey,
		Roles:    roles,
		VendorID: vendorID,
	}
}

func (msg *MsgProposeAddAccount) Route() string {
	return RouterKey
}

func (msg *MsgProposeAddAccount) Type() string {
	return TypeMsgProposeAddAccount
}

func (msg *MsgProposeAddAccount) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

func (msg *MsgProposeAddAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgProposeAddAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}
	return nil
}
