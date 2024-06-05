package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgDeleteComplianceInfo = "delete_compliance_info"

var _ sdk.Msg = &MsgDeleteComplianceInfo{}

func NewMsgDeleteComplianceInfo(creator string, vid int32, pid int32, softwareVersion uint32, certificationType string) *MsgDeleteComplianceInfo {
	return &MsgDeleteComplianceInfo{
		Creator:           creator,
		Vid:               vid,
		Pid:               pid,
		SoftwareVersion:   softwareVersion,
		CertificationType: certificationType,
	}
}

func (msg *MsgDeleteComplianceInfo) Route() string {
	return RouterKey
}

func (msg *MsgDeleteComplianceInfo) Type() string {
	return TypeMsgDeleteComplianceInfo
}

func (msg *MsgDeleteComplianceInfo) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signer}
}

func (msg *MsgDeleteComplianceInfo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteComplianceInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	if !IsValidCertificationType(msg.CertificationType) {
		return NewErrInvalidCertificationType(msg.CertificationType, CertificationTypesList)
	}

	return nil
}
