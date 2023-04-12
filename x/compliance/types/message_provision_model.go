package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	dclcompltypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/compliance"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgProvisionModel = "provision_model"

var _ sdk.Msg = &MsgProvisionModel{}

func NewMsgProvisionModel(
	signer string, vid int32, pid int32, softwareVersion uint32, softwareVersionString string, cDVersionNumber uint32,
	provisionalDate string, certificationType string, reason string, programTypeVersion string, cDCertificateID string,
	familyID string, supportedClusters string, compliantPlatformUsed string, compliantPlatformVersion string, osVersion string,
	certificationRoute string, programType string, transport string, parentChild string, certificationIDOfSoftwareComponent string,
) *MsgProvisionModel {
	return &MsgProvisionModel{
		Signer:                             signer,
		Vid:                                vid,
		Pid:                                pid,
		SoftwareVersion:                    softwareVersion,
		SoftwareVersionString:              softwareVersionString,
		CDVersionNumber:                    cDVersionNumber,
		ProvisionalDate:                    provisionalDate,
		CertificationType:                  certificationType,
		Reason:                             reason,
		ProgramTypeVersion:                 programTypeVersion,
		CDCertificateId:                    cDCertificateID,
		FamilyId:                           familyID,
		SupportedClusters:                  supportedClusters,
		CompliantPlatformUsed:              compliantPlatformUsed,
		CompliantPlatformVersion:           compliantPlatformVersion,
		OSVersion:                          osVersion,
		CertificationRoute:                 certificationRoute,
		ProgramType:                        programType,
		Transport:                          transport,
		ParentChild:                        parentChild,
		CertificationIdOfSoftwareComponent: certificationIDOfSoftwareComponent,
	}
}

func (msg *MsgProvisionModel) Route() string {
	return RouterKey
}

func (msg *MsgProvisionModel) Type() string {
	return TypeMsgProvisionModel
}

func (msg *MsgProvisionModel) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signer}
}

func (msg *MsgProvisionModel) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgProvisionModel) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	_, err = time.Parse(time.RFC3339, msg.ProvisionalDate)
	if err != nil {
		return NewErrInvalidTestDateFormat(msg.ProvisionalDate)
	}

	if !dclcompltypes.IsValidCertificationType(msg.CertificationType) {
		return NewErrInvalidCertificationType(msg.CertificationType, dclcompltypes.CertificationTypesList)
	}

	if !dclcompltypes.IsValidPFCCertificationRoute(msg.ParentChild) {
		return NewErrInvalidPFCCertificationRoute(msg.ParentChild, dclcompltypes.PFCCertificationRouteList)
	}

	return nil
}
