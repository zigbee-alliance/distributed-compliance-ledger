package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgRevokeX509Cert = "revoke_x_509_cert"

var _ sdk.Msg = &MsgRevokeX509Cert{}

func NewMsgRevokeX509Cert(signer string, subject string, subjectKeyId string, info string) *MsgRevokeX509Cert {
	return &MsgRevokeX509Cert{
		Signer:       signer,
		Subject:      subject,
		SubjectKeyId: subjectKeyId,
		Info:         info,
		Time:         time.Now().Unix(),
	}
}

func (msg *MsgRevokeX509Cert) Route() string {
	return RouterKey
}

func (msg *MsgRevokeX509Cert) Type() string {
	return TypeMsgRevokeX509Cert
}

func (msg *MsgRevokeX509Cert) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signer}
}

func (msg *MsgRevokeX509Cert) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgRevokeX509Cert) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	return nil
}
