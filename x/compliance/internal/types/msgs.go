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
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName

type MsgCertifyModel struct {
	VID                   uint16            `json:"vid"`
	PID                   uint16            `json:"pid"`
	SoftwareVersion       uint32            `json:"softwareVersion"`
	SoftwareVersionString string            `json:"softwareVersionString"`
	CertificationDate     time.Time         `json:"certification_date"` // rfc3339 encoded date
	CertificationType     CertificationType `json:"certification_type"`
	Reason                string            `json:"reason,omitempty"`
	Signer                sdk.AccAddress    `json:"signer"`
}

func NewMsgCertifyModel(vid uint16, pid uint16,
	softwareVersion uint32, softwareVersionString string,
	certificationDate time.Time, certificationType CertificationType,
	reason string, signer sdk.AccAddress) MsgCertifyModel {
	return MsgCertifyModel{
		VID:                   vid,
		PID:                   pid,
		SoftwareVersion:       softwareVersion,
		SoftwareVersionString: softwareVersionString,
		CertificationDate:     certificationDate,
		CertificationType:     certificationType,
		Reason:                reason,
		Signer:                signer,
	}
}

func (m MsgCertifyModel) Route() string {
	return RouterKey
}

func (m MsgCertifyModel) Type() string {
	return "certify_model"
}

func (m MsgCertifyModel) ValidateBasic() sdk.Error {
	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress("Invalid Signer: it cannot be empty")
	}

	if m.VID == 0 {
		return sdk.ErrUnknownRequest("Invalid VID: it must be non zero 16-bit unsigned integer")
	}

	if m.PID == 0 {
		return sdk.ErrUnknownRequest("Invalid PID: it must be non zero 16-bit unsigned integer")
	}

	if m.SoftwareVersion == 0 {
		return sdk.ErrUnknownRequest("Invalid SoftwareVersion: it must be non zero 32-bit unsigned integer")
	}

	if m.CertificationDate.IsZero() {
		return sdk.ErrUnknownRequest("Invalid CertificationDate: it cannot be empty")
	}

	if m.CertificationType != ZbCertificationType {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf("Invalid CertificationType: \"%s\". Supported types: [%s]", m.CertificationType, ZbCertificationType))
	}

	return nil
}

func (m MsgCertifyModel) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgCertifyModel) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

type MsgRevokeModel struct {
	VID                   uint16 `json:"vid"`
	PID                   uint16 `json:"pid"`
	SoftwareVersion       uint32 `json:"softwareVersion"`
	SoftwareVersionString string `json:"softwareVersionString"`

	RevocationDate    time.Time         `json:"revocation_date"` // rfc3339 encoded date
	CertificationType CertificationType `json:"certification_type"`
	Reason            string            `json:"reason,omitempty"`
	Signer            sdk.AccAddress    `json:"signer"`
}

func NewMsgRevokeModel(vid uint16, pid uint16,
	softwareVersion uint32,
	revocationDate time.Time, certificationType CertificationType,
	revocationReason string, signer sdk.AccAddress) MsgRevokeModel {
	return MsgRevokeModel{
		VID:               vid,
		PID:               pid,
		SoftwareVersion:   softwareVersion,
		RevocationDate:    revocationDate,
		CertificationType: certificationType,
		Reason:            revocationReason,
		Signer:            signer,
	}
}

func (m MsgRevokeModel) Route() string {
	return RouterKey
}

func (m MsgRevokeModel) Type() string {
	return "revoke_model"
}

func (m MsgRevokeModel) ValidateBasic() sdk.Error {
	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress("Invalid Signer: it cannot be empty")
	}

	if m.VID == 0 {
		return sdk.ErrUnknownRequest("Invalid VID: it must be non zero 16-bit unsigned integer")
	}

	if m.PID == 0 {
		return sdk.ErrUnknownRequest("Invalid PID: it must be non zero 16-bit unsigned integer")
	}

	if m.SoftwareVersion == 0 {
		return sdk.ErrUnknownRequest("Invalid SoftwareVersion: it must be non zero 32-bit unsigned integer")
	}

	if m.RevocationDate.IsZero() {
		return sdk.ErrUnknownRequest("Invalid RevocationDate: it cannot be empty")
	}

	if m.CertificationType != ZbCertificationType {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf("Invalid CertificationType: \"%s\". Supported types: [%s]", m.CertificationType, ZbCertificationType))
	}

	return nil
}

func (m MsgRevokeModel) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgRevokeModel) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}
