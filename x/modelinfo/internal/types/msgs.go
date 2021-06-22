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
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName

//nolint:maligned
type MsgAddModelInfo struct {
	VID                                        uint16         `json:"vid"`
	PID                                        uint16         `json:"pid"`
	CID                                        uint16         `json:"cid,omitempty"`
	Name                                       string         `json:"name"`
	Description                                string         `json:"description"`
	SKU                                        string         `json:"sku"`
	SoftwareVersion                            uint32         `json:"software_version"`
	SoftwareVersionString                      string         `json:"software_version_string"`
	HardwareVersion                            uint32         `json:"hardware_version"`
	HardwareVersionString                      string         `json:"hardware_version_string"`
	CDVersionNumber                            uint16         `json:"cd_version_number"`
	FirmwareDigests                            string         `json:"firmware_digests,omitempty"`
	Revoked                                    bool           `json:"revoked"`
	OtaURL                                     string         `json:"ota_url,omitempty"`
	OtaChecksum                                string         `json:"ota_checksum,omitempty"`
	OtaChecksumType                            string         `json:"ota_checksum_type,omitempty"`
	OtaBlob                                    string         `json:"ota_blob,omitempty"`
	CommissioningCustomFlow                    uint8          `json:"commission_custom_flow,omitempty"`
	CommissioningCustomFlowUrl                 string         `json:"commission_custom_flow_url,omitempty"`
	CommissioningModeInitialStepsHint          uint32         `json:"commisioning_mode_initial_steps_hint,omitempty"`
	CommissioningModeInitialStepsInstruction   string         `json:"commisioning_mode_initial_steps_instruction,omitempty"`
	CommissioningModeSecondaryStepsHint        uint32         `json:"commisioning_mode_secondary_steps_hint,omitempty"`
	CommissioningModeSecondaryStepsInstruction string         `json:"commisioning_mode_secondary_steps_instruction,omitempty"`
	ReleaseNotesUrl                            string         `json:"release_notes_url,omitempty"`
	UserManualUrl                              string         `json:"user_manual_url,omitempty"`
	SupportUrl                                 string         `json:"support_url,omitempty"`
	ProductURL                                 string         `json:"product_url,omitempty"`
	ChipBlob                                   string         `json:"chip_blob,omitempty"`
	VendorBlob                                 string         `json:"vendor_blob,omitempty"`
	Signer                                     sdk.AccAddress `json:"signer"`
}

func NewMsgAddModelInfo(
	vid uint16,
	pid uint16,
	cid uint16,
	name string,
	description string,
	sku string,
	softwareVersion uint32,
	softwareVersionString string,
	hardwareVersion uint32,
	hardwareVersionString string,
	cDVersionNumber uint16,
	firmwareDigests string,
	revoked bool,
	otaURL string,
	otaChecksum string,
	otaChecksumType string,
	otaBlob string,
	commissioningCustomFlow uint8,
	commissioningCustomFlowUrl string,
	commissioningModeInitialStepsHint uint32,
	commissioningModeInitialStepsInstruction string,
	commissioningModeSecondaryStepsHint uint32,
	commissioningModeSecondaryStepsInstruction string,
	releaseNotesUrl string,
	userManualUrl string,
	supportUrl string,
	productURL string,
	chipBlob string,
	vendorBlob string,
	signer sdk.AccAddress,
) MsgAddModelInfo {
	return MsgAddModelInfo{
		VID:                                      vid,
		PID:                                      pid,
		CID:                                      cid,
		Name:                                     name,
		Description:                              description,
		SKU:                                      sku,
		SoftwareVersion:                          softwareVersion,
		SoftwareVersionString:                    softwareVersionString,
		HardwareVersion:                          hardwareVersion,
		HardwareVersionString:                    hardwareVersionString,
		CDVersionNumber:                          cDVersionNumber,
		FirmwareDigests:                          firmwareDigests,
		Revoked:                                  revoked,
		OtaURL:                                   otaURL,
		OtaChecksum:                              otaChecksum,
		OtaChecksumType:                          otaChecksumType,
		OtaBlob:                                  otaBlob,
		CommissioningCustomFlow:                  commissioningCustomFlow,
		CommissioningCustomFlowUrl:               commissioningCustomFlowUrl,
		CommissioningModeInitialStepsHint:        commissioningModeInitialStepsHint,
		CommissioningModeInitialStepsInstruction: commissioningModeInitialStepsInstruction,
		CommissioningModeSecondaryStepsHint:      commissioningModeSecondaryStepsHint,
		CommissioningModeSecondaryStepsInstruction: commissioningModeSecondaryStepsInstruction,
		ReleaseNotesUrl: releaseNotesUrl,
		UserManualUrl:   userManualUrl,
		SupportUrl:      supportUrl,
		ProductURL:      productURL,
		ChipBlob:        chipBlob,
		VendorBlob:      vendorBlob,
		Signer:          signer,
	}
}

func (m MsgAddModelInfo) Route() string {
	return RouterKey
}

func (m MsgAddModelInfo) Type() string {
	return "add_model_info"
}

func (m MsgAddModelInfo) ValidateBasic() sdk.Error {
	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress("Invalid Signer: it cannot be empty")
	}

	if m.VID == 0 {
		return sdk.ErrUnknownRequest("Invalid VID: it must be non-zero 16-bit unsigned integer")
	}

	if m.PID == 0 {
		return sdk.ErrUnknownRequest("Invalid PID: it must be non-zero 16-bit unsigned integer")
	}

	if len(m.Name) == 0 {
		return sdk.ErrUnknownRequest("Invalid Name: it cannot be empty")
	}

	if len(m.Description) == 0 {
		return sdk.ErrUnknownRequest("Invalid Description: it cannot be empty")
	}

	if len(m.SKU) == 0 {
		return sdk.ErrUnknownRequest("Invalid SKU: it cannot be empty")
	}

	if m.SoftwareVersion == 0 {
		return sdk.ErrUnknownRequest("Invalid SoftwareVersion: it must be non-zero 32-bit unsigned integer")
	}

	if len(m.SoftwareVersionString) == 0 {
		return sdk.ErrUnknownRequest("Invalid SoftwareVersionString: it cannot be empty")
	}

	if m.HardwareVersion == 0 {
		return sdk.ErrUnknownRequest("Invalid HardwareVersion: it must be non-zero 32-bit unsigned integer")
	}

	if len(m.HardwareVersionString) == 0 {
		return sdk.ErrUnknownRequest("Invalid HardwareVersionString: it cannot be empty")
	}

	if m.CDVersionNumber == 0 {
		return sdk.ErrUnknownRequest("Invalid CDVersionNumber: it must be non-zero 16-bit unsigned integer")
	}

	if m.OtaURL != "" || m.OtaChecksum != "" || m.OtaChecksumType != "" {
		if m.OtaURL == "" || m.OtaChecksum == "" || m.OtaChecksumType == "" {
			return sdk.ErrUnknownRequest("Invalid MsgAddModelInfo: the fields OtaURL, OtaChecksum and " +
				"OtaChecksumType must be either specified together, or not specified together")
		}
	}

	return nil
}

func (m MsgAddModelInfo) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgAddModelInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

//nolint:maligned
type MsgUpdateModelInfo struct {
	VID                        uint16         `json:"vid"`
	PID                        uint16         `json:"pid"`
	CID                        uint16         `json:"cid,omitempty"`
	Description                string         `json:"description,omitempty"`
	CDVersionNumber            uint16         `json:"cd_version_number"`
	Revoked                    bool           `json:"revoked"`
	OtaURL                     string         `json:"ota_url,omitempty"`
	OtaChecksum                string         `json:"ota_checksum,omitempty"`
	OtaChecksumType            string         `json:"ota_checksum_type,omitempty"`
	OtaBlob                    string         `json:"ota_blob,omitempty"`
	CommissioningCustomFlowUrl string         `json:"commission_custom_flow_url,omitempty"`
	ReleaseNotesUrl            string         `json:"release_notes_url,omitempty"`
	UserManualUrl              string         `json:"user_manual_url,omitempty"`
	SupportUrl                 string         `json:"support_url,omitempty"`
	ProductURL                 string         `json:"product_url,omitempty"`
	ChipBlob                   string         `json:"chip_blob,omitempty"`
	VendorBlob                 string         `json:"vendor_blob,omitempty"`
	Signer                     sdk.AccAddress `json:"signer"`
}

func NewMsgUpdateModelInfo(
	vid uint16,
	pid uint16,
	cid uint16,
	description string,
	cDVersionNumber uint16,
	revoked bool,
	otaURL string,
	otaChecksum string,
	otaChecksumType string,
	otaBlob string,
	commissioningCustomFlowUrl string,
	releaseNotesUrl string,
	userManualUrl string,
	supportUrl string,
	productURL string,
	chipBlob string,
	vendorBlob string,
	signer sdk.AccAddress,
) MsgUpdateModelInfo {
	return MsgUpdateModelInfo{
		VID:                        vid,
		PID:                        pid,
		CID:                        cid,
		Description:                description,
		CDVersionNumber:            cDVersionNumber,
		Revoked:                    revoked,
		OtaURL:                     otaURL,
		OtaChecksum:                otaChecksum,
		OtaChecksumType:            otaChecksumType,
		OtaBlob:                    otaBlob,
		CommissioningCustomFlowUrl: commissioningCustomFlowUrl,
		ReleaseNotesUrl:            releaseNotesUrl,
		UserManualUrl:              userManualUrl,
		SupportUrl:                 supportUrl,
		ProductURL:                 productURL,
		ChipBlob:                   chipBlob,
		VendorBlob:                 vendorBlob,
		Signer:                     signer,
	}
}

func (m MsgUpdateModelInfo) Route() string {
	return RouterKey
}

func (m MsgUpdateModelInfo) Type() string {
	return "update_model_info"
}

func (m MsgUpdateModelInfo) ValidateBasic() sdk.Error {
	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress("Invalid Signer: it cannot be empty")
	}

	if m.VID == 0 {
		return sdk.ErrUnknownRequest("Invalid VID: it must be non-zero 16-bit unsigned integer")
	}

	if m.PID == 0 {
		return sdk.ErrUnknownRequest("Invalid PID: it must be non-zero 16-bit unsigned integer")
	}

	return nil
}

func (m MsgUpdateModelInfo) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgUpdateModelInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

type MsgDeleteModelInfo struct {
	VID    uint16         `json:"vid"`
	PID    uint16         `json:"pid"`
	Signer sdk.AccAddress `json:"signer"`
}

func NewMsgDeleteModelInfo(vid uint16, pid uint16, signer sdk.AccAddress) MsgDeleteModelInfo {
	return MsgDeleteModelInfo{
		VID:    vid,
		PID:    pid,
		Signer: signer,
	}
}

func (m MsgDeleteModelInfo) Route() string {
	return RouterKey
}

func (m MsgDeleteModelInfo) Type() string {
	return "delete_model_info"
}

func (m MsgDeleteModelInfo) ValidateBasic() sdk.Error {
	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress("Invalid Signer: it cannot be empty")
	}

	if m.VID == 0 {
		return sdk.ErrUnknownRequest("Invalid VID: it must be non-zero 16-bit unsigned integer")
	}

	if m.PID == 0 {
		return sdk.ErrUnknownRequest("Invalid PID: it must be non-zero 16-bit unsigned integer")
	}

	return nil
}

func (m MsgDeleteModelInfo) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgDeleteModelInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}
