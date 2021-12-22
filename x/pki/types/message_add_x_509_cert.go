package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddX509Cert = "add_x_509_cert"

var _ sdk.Msg = &MsgAddX509Cert{}

func NewMsgAddX509Cert(signer string, cert string) *MsgAddX509Cert {
	return &MsgAddX509Cert{
		Signer: signer,
		Cert:   cert,
	}
}

func (msg *MsgAddX509Cert) Route() string {
	return RouterKey
}

func (msg *MsgAddX509Cert) Type() string {
	return TypeMsgAddX509Cert
}

func (msg *MsgAddX509Cert) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

func (msg *MsgAddX509Cert) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddX509Cert) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}
	return nil
}
