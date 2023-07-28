package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/x509"
)

const TypeMsgAssignVid = "assign_vid"

var _ sdk.Msg = &MsgAssignVid{}

func NewMsgAssignVid(signer string, subject string, subjectKeyId string, vid int32) *MsgAssignVid {
	return &MsgAssignVid{
		Signer:       signer,
		Subject:      subject,
		SubjectKeyId: subjectKeyId,
		Vid:          vid,
	}
}

func (msg *MsgAssignVid) Route() string {
	return pkitypes.RouterKey
}

func (msg *MsgAssignVid) Type() string {
	return TypeMsgAssignVid
}

func (msg *MsgAssignVid) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

func (msg *MsgAssignVid) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAssignVid) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	subjectVid, err := x509.GetVidFromSubject(msg.Subject)
	if err != nil {
		return pkitypes.NewErrInvalidCertificate(err)
	}
	
	if subjectVid != 0 && subjectVid != msg.Vid {
		return pkitypes.NewErrCertificateVidNotEqualMsgVid(fmt.Sprintf("Certificate VID=%d is not equal to the msg VID=%d", subjectVid, msg.Vid))
	}

	return nil
}
