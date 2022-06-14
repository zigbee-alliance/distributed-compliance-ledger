package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgCertifyModel = "certify_model"

var _ sdk.Msg = &MsgCertifyModel{}

func NewMsgCertifyModel(
	signer string, vid int32, pid int32, softwareVersion uint32, softwareVersionString string, cdVersionNumber uint32,
	certificationDate string, certificationType string, reason string, programTypeVersion string, cDCertificationID string,
	familyID string, supportedClusters string, compliancePlatformUsed string, compliancePlatformVersion string, osVersion string,
	certificationRoute string,
) *MsgCertifyModel {
	return &MsgCertifyModel{
		Signer:                    signer,
		Vid:                       vid,
		Pid:                       pid,
		SoftwareVersion:           softwareVersion,
		SoftwareVersionString:     softwareVersionString,
		CDVersionNumber:           cdVersionNumber,
		CertificationDate:         certificationDate,
		CertificationType:         certificationType,
		Reason:                    reason,
		ProgramTypeVersion:        programTypeVersion,
		CDCertificationID:         cDCertificationID,
		FamilyID:                  familyID,
		SupportedClusters:         supportedClusters,
		CompliancePlatformUsed:    compliancePlatformUsed,
		CompliancePlatformVersion: compliancePlatformVersion,
		OSVersion:                 osVersion,
		CertificationRoute:        certificationRoute,
	}
}

func (msg *MsgCertifyModel) Route() string {
	return RouterKey
}

func (msg *MsgCertifyModel) Type() string {
	return TypeMsgCertifyModel
}

func (msg *MsgCertifyModel) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signer}
}

func (msg *MsgCertifyModel) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgCertifyModel) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	_, err = time.Parse(time.RFC3339, msg.CertificationDate)
	if err != nil {
		return NewErrInvalidTestDateFormat(msg.CertificationDate)
	}

	if !IsValidCertificationType(msg.CertificationType) {
		return NewErrInvalidCertificationType(msg.CertificationType, CertificationTypesList)
	}

	return nil
}
