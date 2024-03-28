package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/x509"
)

const TypeMsgAddNocX509Cert = "add_noc_x_509_cert"

var _ sdk.Msg = &MsgAddNocX509Cert{}

func NewMsgAddNocX509Cert(signer string, cert string, certSchemaVersion, schemaVersion uint32) *MsgAddNocX509Cert {
	return &MsgAddNocX509Cert{
		Signer:            signer,
		Cert:              cert,
		CertSchemaVersion: certSchemaVersion,
		SchemaVersion:     schemaVersion,
	}
}

func (msg *MsgAddNocX509Cert) Route() string {
	return pkitypes.RouterKey
}

func (msg *MsgAddNocX509Cert) Type() string {
	return TypeMsgAddNocX509Cert
}

func (msg *MsgAddNocX509Cert) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signer}
}

func (msg *MsgAddNocX509Cert) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddNocX509Cert) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	_, err = x509.DecodeX509Certificate(msg.Cert)
	if err != nil {
		return pkitypes.NewErrInvalidCertificate(err)
	}

	return nil
}
