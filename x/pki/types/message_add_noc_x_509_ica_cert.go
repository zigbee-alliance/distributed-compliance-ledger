package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/x509"
)

const TypeMsgAddNocX509IcaCert = "add_noc_x_509_ica_cert"

var _ sdk.Msg = &MsgAddNocX509IcaCert{}

func NewMsgAddNocX509IcaCert(signer string, cert string, certSchemaVersion uint32, isVidVerificationSigner bool) *MsgAddNocX509IcaCert {
	return &MsgAddNocX509IcaCert{
		Signer:                  signer,
		Cert:                    cert,
		CertSchemaVersion:       certSchemaVersion,
		IsVidVerificationSigner: isVidVerificationSigner,
	}
}

func (msg *MsgAddNocX509IcaCert) Route() string {
	return pkitypes.RouterKey
}

func (msg *MsgAddNocX509IcaCert) Type() string {
	return TypeMsgAddNocX509IcaCert
}

func (msg *MsgAddNocX509IcaCert) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signer}
}

func (msg *MsgAddNocX509IcaCert) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddNocX509IcaCert) ValidateBasic() error {
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
