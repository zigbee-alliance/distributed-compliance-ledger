package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgRejectAddX509RootCert = "reject_add_x_509_root_cert"

var _ sdk.Msg = &MsgRejectAddX509RootCert{}

func NewMsgRejectAddX509RootCert(signer string, subject string, subjectKeyID string, info string, schemaVersion uint32) *MsgRejectAddX509RootCert {
	return &MsgRejectAddX509RootCert{
		Signer:        signer,
		Subject:       subject,
		SubjectKeyId:  subjectKeyID,
		Info:          info,
		Time:          time.Now().Unix(),
		SchemaVersion: schemaVersion,
	}
}

func (msg *MsgRejectAddX509RootCert) Route() string {
	return pkitypes.RouterKey
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

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	return nil
}
