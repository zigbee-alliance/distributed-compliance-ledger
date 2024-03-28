package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgRemoveX509Cert = "remove_x_509_cert"

var _ sdk.Msg = &MsgRemoveX509Cert{}

func NewMsgRemoveX509Cert(signer string, subject string, subjectKeyID string, serialNumber string) *MsgRemoveX509Cert {
	return &MsgRemoveX509Cert{
		Signer:       signer,
		Subject:      subject,
		SubjectKeyId: subjectKeyID,
		SerialNumber: serialNumber,
	}
}

func (msg *MsgRemoveX509Cert) Route() string {
	return pkitypes.RouterKey
}

func (msg *MsgRemoveX509Cert) Type() string {
	return TypeMsgRemoveX509Cert
}

func (msg *MsgRemoveX509Cert) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signer}
}

func (msg *MsgRemoveX509Cert) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveX509Cert) ValidateBasic() error {
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
