package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgApproveAddX509RootCert = "approve_add_x_509_root_cert"

var _ sdk.Msg = &MsgApproveAddX509RootCert{}

func NewMsgApproveAddX509RootCert(signer string, subject string, subjectKeyId string) *MsgApproveAddX509RootCert {
	return &MsgApproveAddX509RootCert{
		Signer:       signer,
		Subject:      subject,
		SubjectKeyId: subjectKeyId,
	}
}

func (msg *MsgApproveAddX509RootCert) Route() string {
	return RouterKey
}

func (msg *MsgApproveAddX509RootCert) Type() string {
	return TypeMsgApproveAddX509RootCert
}

func (msg *MsgApproveAddX509RootCert) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

func (msg *MsgApproveAddX509RootCert) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgApproveAddX509RootCert) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}
	return nil
}
