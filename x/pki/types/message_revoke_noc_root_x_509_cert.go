package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgRevokeNocRootX509Cert = "revoke_noc_root_x_509_cert"

var _ sdk.Msg = &MsgRevokeNocRootX509Cert{}

func NewMsgRevokeNocRootX509Cert(signer, subject, subjectKeyID, serialNumber, info string, revokeChild bool, schemaVersion uint32) *MsgRevokeNocRootX509Cert {
	return &MsgRevokeNocRootX509Cert{
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

func (msg *MsgRevokeNocRootX509Cert) Route() string {
	return pkitypes.RouterKey
}

func (msg *MsgRevokeNocRootX509Cert) Type() string {
	return TypeMsgRevokeNocRootX509Cert
}

func (msg *MsgRevokeNocRootX509Cert) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signer}
}

func (msg *MsgRevokeNocRootX509Cert) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgRevokeNocRootX509Cert) ValidateBasic() error {
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
