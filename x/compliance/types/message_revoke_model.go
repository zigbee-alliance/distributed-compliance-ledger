package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgRevokeModel = "revoke_model"

var _ sdk.Msg = &MsgRevokeModel{}

func NewMsgRevokeModel(
	signer string, vid int32, pid int32, softwareVersion uint32, softwareVersionString string, cdVersionNumber uint32,
	revocationDate string, certificationType string, reason string, programTypeVersion string, cDCertificationID string,
	familyID string, supportedClusters string, compliancePlatformUsed string, compliancePlatformVersion string, osVersion string,
	certificationRoute string, programType string, transport string, parentChild string,
) *MsgRevokeModel {
	return &MsgRevokeModel{
		Signer:                    signer,
		Vid:                       vid,
		Pid:                       pid,
		SoftwareVersion:           softwareVersion,
		SoftwareVersionString:     softwareVersionString,
		CDVersionNumber:           cdVersionNumber,
		RevocationDate:            revocationDate,
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
		ProgramType:               programType,
		Transport:                 transport,
		ParentChild:               parentChild,
	}
}

func (msg *MsgRevokeModel) Route() string {
	return RouterKey
}

func (msg *MsgRevokeModel) Type() string {
	return TypeMsgRevokeModel
}

func (msg *MsgRevokeModel) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signer}
}

func (msg *MsgRevokeModel) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgRevokeModel) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	_, err = time.Parse(time.RFC3339, msg.RevocationDate)
	if err != nil {
		return NewErrInvalidTestDateFormat(msg.RevocationDate)
	}

	if !IsValidCertificationType(msg.CertificationType) {
		return NewErrInvalidCertificationType(msg.CertificationType, CertificationTypesList)
	}

	return nil
}
