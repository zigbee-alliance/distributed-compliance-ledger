package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgRevokeNocX509CRootert = "revoke_noc_x_509_root_cert"

var _ sdk.Msg = &MsgRevokeNocX509RootCert{}

func NewMsgRevokeNocX509RootCert(signer, subject, subjectKeyID, serialNumber, info string, revokeChild bool, schemaVersion uint32) *MsgRevokeNocX509RootCert {
	return &MsgRevokeNocX509RootCert{
		Signer:        signer,
		Subject:       subject,
		SubjectKeyId:  subjectKeyID,
		SerialNumber:  serialNumber,
		Info:          info,
		Time:          time.Now().Unix(),
		RevokeChild:   revokeChild,
		SchemaVersion: schemaVersion,
	}
}

func (msg *MsgRevokeNocX509RootCert) Route() string {
	return pkitypes.RouterKey
}

func (msg *MsgRevokeNocX509RootCert) Type() string {
	return TypeMsgRevokeNocX509CRootert
}

func (msg *MsgRevokeNocX509RootCert) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signer}
}

func (msg *MsgRevokeNocX509RootCert) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgRevokeNocX509RootCert) ValidateBasic() error {
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
