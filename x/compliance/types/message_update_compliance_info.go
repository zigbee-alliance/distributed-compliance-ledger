package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateComplianceInfo = "update_compliance_info"

var _ sdk.Msg = &MsgUpdateComplianceInfo{}

func NewMsgUpdateComplianceInfo(creator string, vid int32, pid int32, softwareVersion uint32, softwareVersionString string, certificationType string, cDVersionNumber uint32, softwareVersionCertificationStatus uint32, date string, reason string, owner string, cDCertificateID string, certificationRoute string, programType string, programTypeVersion string, compliantPlatformUsed string, compliantPlatformVersion string, transport string, familyId string, supportedClusters string, oSVersion string, parentChild string, certificationIdOfSoftwareComponent string) *MsgUpdateComplianceInfo {
	return &MsgUpdateComplianceInfo{
		Creator:                            creator,
		Vid:                                vid,
		Pid:                                pid,
		SoftwareVersion:                    softwareVersion,
		SoftwareVersionString:              softwareVersionString,
		CertificationType:                  certificationType,
		CDVersionNumber:                    cDVersionNumber,
		SoftwareVersionCertificationStatus: softwareVersionCertificationStatus,
		Date:                               date,
		Reason:                             reason,
		Owner:                              owner,
		CDCertificateId:                    cDCertificateID,
		CertificationRoute:                 certificationRoute,
		ProgramType:                        programType,
		ProgramTypeVersion:                 programTypeVersion,
		CompliantPlatformUsed:              compliantPlatformUsed,
		CompliantPlatformVersion:           compliantPlatformVersion,
		Transport:                          transport,
		FamilyId:                           familyId,
		SupportedClusters:                  supportedClusters,
		OSVersion:                          oSVersion,
		ParentChild:                        parentChild,
		CertificationIdOfSoftwareComponent: certificationIdOfSoftwareComponent,
	}
}

func (msg *MsgUpdateComplianceInfo) Route() string {
	return RouterKey
}

func (msg *MsgUpdateComplianceInfo) Type() string {
	return TypeMsgUpdateComplianceInfo
}

func (msg *MsgUpdateComplianceInfo) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateComplianceInfo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateComplianceInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
