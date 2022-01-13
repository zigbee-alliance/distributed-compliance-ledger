// Copyright 2022 DSR Corporation
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

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgCertifyModel = "certify_model"

var _ sdk.Msg = &MsgCertifyModel{}

func NewMsgCertifyModel(signer string, vid int32, pid int32, softwareVersion uint32, softwareVersionString string, cdVersionNumber uint32, certificationDate string, certificationType string, reason string) *MsgCertifyModel {
	return &MsgCertifyModel{
		Signer:                signer,
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       softwareVersion,
		SoftwareVersionString: softwareVersionString,
		CDVersionNumber:       cdVersionNumber,
		CertificationDate:     certificationDate,
		CertificationType:     certificationType,
		Reason:                reason,
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
