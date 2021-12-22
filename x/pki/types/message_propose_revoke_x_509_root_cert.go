package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgProposeRevokeX509RootCert = "propose_revoke_x_509_root_cert"

var _ sdk.Msg = &MsgProposeRevokeX509RootCert{}

func NewMsgProposeRevokeX509RootCert(signer string, subject string, subjectKeyId string) *MsgProposeRevokeX509RootCert {
	return &MsgProposeRevokeX509RootCert{
		Signer:       signer,
		Subject:      subject,
		SubjectKeyId: subjectKeyId,
	}
}

func (msg *MsgProposeRevokeX509RootCert) Route() string {
	return RouterKey
}

func (msg *MsgProposeRevokeX509RootCert) Type() string {
	return TypeMsgProposeRevokeX509RootCert
}

func (msg *MsgProposeRevokeX509RootCert) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

func (msg *MsgProposeRevokeX509RootCert) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgProposeRevokeX509RootCert) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}
	return nil
}
