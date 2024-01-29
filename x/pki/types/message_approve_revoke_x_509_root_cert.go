package types

import (
	"time"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgApproveRevokeX509RootCert = "approve_revoke_x_509_root_cert"

var _ sdk.Msg = &MsgApproveRevokeX509RootCert{}

func NewMsgApproveRevokeX509RootCert(signer string, subject string, subjectKeyID string, info string) *MsgApproveRevokeX509RootCert {
	return &MsgApproveRevokeX509RootCert{
		Signer:       signer,
		Subject:      subject,
		SubjectKeyId: subjectKeyID,
		Info:         info,
		Time:         time.Now().Unix(),
	}
}

func (msg *MsgApproveRevokeX509RootCert) Route() string {
	return pkitypes.RouterKey
}

func (msg *MsgApproveRevokeX509RootCert) Type() string {
	return TypeMsgApproveRevokeX509RootCert
}

func (msg *MsgApproveRevokeX509RootCert) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signer}
}

func (msg *MsgApproveRevokeX509RootCert) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgApproveRevokeX509RootCert) ValidateBasic() error {
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
