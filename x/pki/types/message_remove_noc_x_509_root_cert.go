package types

import (
	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgRemoveNocX509RootCert = "remove_noc_x_509_root_cert"

var _ sdk.Msg = &MsgRemoveNocX509RootCert{}

func NewMsgRemoveNocX509RootCert(signer string, subject string, subjectKeyID string, serialNumber string) *MsgRemoveNocX509RootCert {
	return &MsgRemoveNocX509RootCert{
		Signer:       signer,
		Subject:      subject,
		SubjectKeyId: subjectKeyID,
		SerialNumber: serialNumber,
	}
}

func (msg *MsgRemoveNocX509RootCert) Route() string {
	return pkitypes.RouterKey
}

func (msg *MsgRemoveNocX509RootCert) Type() string {
	return TypeMsgRemoveNocX509RootCert
}

func (msg *MsgRemoveNocX509RootCert) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signer}
}

func (msg *MsgRemoveNocX509RootCert) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveNocX509RootCert) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	return nil
}
