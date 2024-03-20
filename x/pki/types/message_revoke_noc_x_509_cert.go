package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgRevokeNocX509Cert = "revoke_noc_x_509_cert"

var _ sdk.Msg = &MsgRevokeNocX509Cert{}

func NewMsgRevokeNocX509Cert(signer, subject, subjectKeyID, serialNumber, info string, revokeChild bool) *MsgRevokeNocX509Cert {
	return &MsgRevokeNocX509Cert{
		Signer:       signer,
		Subject:      subject,
		SubjectKeyId: subjectKeyID,
		SerialNumber: serialNumber,
		Info:         info,
		Time:         time.Now().Unix(),
		RevokeChild:  revokeChild,
	}
}

func (msg *MsgRevokeNocX509Cert) Route() string {
	return pkitypes.RouterKey
}

func (msg *MsgRevokeNocX509Cert) Type() string {
	return TypeMsgRevokeNocX509Cert
}

func (msg *MsgRevokeNocX509Cert) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signer}
}

func (msg *MsgRevokeNocX509Cert) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgRevokeNocX509Cert) ValidateBasic() error {
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
