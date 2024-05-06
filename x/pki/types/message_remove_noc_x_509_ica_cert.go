package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgRemoveNocX509IcaCert = "remove_noc_x_509_ica_cert"

var _ sdk.Msg = &MsgRemoveNocX509IcaCert{}

func NewMsgRemoveNocX509IcaCert(signer string, subject string, subjectKeyID string, serialNumber string) *MsgRemoveNocX509IcaCert {
	return &MsgRemoveNocX509IcaCert{
		Signer:       signer,
		Subject:      subject,
		SubjectKeyId: subjectKeyID,
		SerialNumber: serialNumber,
	}
}

func (msg *MsgRemoveNocX509IcaCert) Route() string {
	return pkitypes.RouterKey
}

func (msg *MsgRemoveNocX509IcaCert) Type() string {
	return TypeMsgRemoveNocX509IcaCert
}

func (msg *MsgRemoveNocX509IcaCert) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signer}
}

func (msg *MsgRemoveNocX509IcaCert) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveNocX509IcaCert) ValidateBasic() error {
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
