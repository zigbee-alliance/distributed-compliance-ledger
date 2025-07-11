package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/x509"
)

const TypeMsgAddNocX509RootCert = "add_noc_x_509_root_cert"

var _ sdk.Msg = &MsgAddNocX509RootCert{}

func NewMsgAddNocX509RootCert(signer string, cert string, certSchemaVersion uint32, isVVSC bool) *MsgAddNocX509RootCert {
	return &MsgAddNocX509RootCert{
		Signer:            signer,
		Cert:              cert,
		CertSchemaVersion: certSchemaVersion,
		IsVVSC:            isVVSC,
	}
}

func (msg *MsgAddNocX509RootCert) Route() string {
	return pkitypes.RouterKey
}

func (msg *MsgAddNocX509RootCert) Type() string {
	return TypeMsgAddNocX509RootCert
}

func (msg *MsgAddNocX509RootCert) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signer}
}

func (msg *MsgAddNocX509RootCert) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddNocX509RootCert) ValidateBasic() error {
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
