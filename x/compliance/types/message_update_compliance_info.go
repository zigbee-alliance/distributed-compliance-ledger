package types

import (
	"strconv"
	"time"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

var _ sdk.Msg = &MsgUpdateComplianceInfo{}

const TypeMsgUpdateComplianceInfo = "update_compliance_info"

func NewMsgUpdateComplianceInfo(creator string, vid int32, pid int32, softwareVersion uint32, certificationType string,
	cDVersionNumber string, date string, reason string, owner string, cDCertificateID string, certificationRoute string,
	programType string, programTypeVersion string, compliantPlatformUsed string, compliantPlatformVersion string,
	transport string, familyID string, supportedClusters string, oSVersion string, parentChild string,
	certificationIDOfSoftwareComponent string, schemaVersion uint32) *MsgUpdateComplianceInfo {
	return &MsgUpdateComplianceInfo{
		Creator:                            creator,
		Vid:                                vid,
		Pid:                                pid,
		SoftwareVersion:                    softwareVersion,
		CertificationType:                  certificationType,
		CDVersionNumber:                    cDVersionNumber,
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
		FamilyId:                           familyID,
		SupportedClusters:                  supportedClusters,
		OSVersion:                          oSVersion,
		ParentChild:                        parentChild,
		CertificationIdOfSoftwareComponent: certificationIDOfSoftwareComponent,
		SchemaVersion:                      schemaVersion,
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
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	err = validator.Validate(msg)

	if err != nil {
		return err
	}

	if msg.Date != "" {
		_, err = time.Parse(time.RFC3339, msg.Date)

		if err != nil {
			return NewErrInvalidTestDateFormat(msg.Date)
		}
	}

	if msg.CDVersionNumber != "" {
		cdVersionNumber, err := strconv.ParseUint(msg.CDVersionNumber, 10, 32)

		if err != nil {
			return NewErrInvalidUint32ForCdVersionNumber(msg.Vid, msg.Pid, msg.SoftwareVersion, msg.CertificationType, msg.CDVersionNumber)
		}

		if cdVersionNumber > 65535 {
			return errors.Wrap(validator.ErrFieldUpperBoundViolated, "CDVersionNumber must not be greater than 65535: field upper bound violatedError")
		}
	}

	if !IsValidCertificationType(msg.CertificationType) {
		return NewErrInvalidCertificationType(msg.CertificationType, CertificationTypesList)
	}

	if !IsValidPFCCertificationRoute(msg.ParentChild) {
		return NewErrInvalidPFCCertificationRoute(msg.ParentChild, PFCCertificationRouteList)
	}

	return nil
}
