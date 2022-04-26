package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRejectAddX509RootCert = "reject_add_x_509_root_cert"

var _ sdk.Msg = &MsgRejectAddX509RootCert{}

func NewMsgRejectAddX509RootCert(signer string, cert string) *MsgRejectAddX509RootCert {
	return &MsgRejectAddX509RootCert{
		Signer: signer,
		Cert:   cert,
	}
}

func (msg *MsgRejectAddX509RootCert) Route() string {
	return RouterKey
}

func (msg *MsgRejectAddX509RootCert) Type() string {
	return TypeMsgRejectAddX509RootCert
}

func (msg *MsgRejectAddX509RootCert) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

func (msg *MsgRejectAddX509RootCert) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRejectAddX509RootCert) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}
	return nil
}
