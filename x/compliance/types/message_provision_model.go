// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"time"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgProvisionModel = "provision_model"

var _ sdk.Msg = &MsgProvisionModel{}

func NewMsgProvisionModel(
	signer string, vid int32, pid int32, softwareVersion uint32, softwareVersionString string, cDVersionNumber uint32,
	provisionalDate string, certificationType string, reason string, programTypeVersion string, cDCertificateID string,
	familyID string, supportedClusters string, compliantPlatformUsed string, compliantPlatformVersion string, osVersion string,
	certificationRoute string, programType string, transport string, parentChild string, certificationIDOfSoftwareComponent string,
	schemaVersion uint32,
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
		SchemaVersion:                      schemaVersion,
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
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	_, err = time.Parse(time.RFC3339, msg.ProvisionalDate)
	if err != nil {
		return NewErrInvalidTestDateFormat(msg.ProvisionalDate)
	}

	if !IsValidCertificationType(msg.CertificationType) {
		return NewErrInvalidCertificationType(msg.CertificationType, CertificationTypesList)
	}

	if !IsValidPFCCertificationRoute(msg.ParentChild) {
		return NewErrInvalidPFCCertificationRoute(msg.ParentChild, PFCCertificationRouteList)
	}

	return nil
}
