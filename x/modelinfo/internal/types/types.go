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
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

/*
	Model Info stored in KVStore
*/
//nolint:maligned
type ModelInfo struct {
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
	Owner                                      sdk.AccAddress `json:"owner"`
}

func NewModelInfo(
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
	owner sdk.AccAddress,

) ModelInfo {
	return ModelInfo{
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
		Owner:           owner,
	}
}

func (d ModelInfo) String() string {
	bytes, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

/*
	List of products for specific Vendor
*/
type VendorProducts struct {
	VID      uint16    `json:"vid"`
	Products []Product `json:"products"`
}

func NewVendorProducts(vid uint16) VendorProducts {
	return VendorProducts{
		VID:      vid,
		Products: []Product{},
	}
}

func (d *VendorProducts) AddVendorProduct(pid Product) {
	d.Products = append(d.Products, pid)
}

func (d *VendorProducts) RemoveVendorProduct(pid uint16) {
	for i, value := range d.Products {
		if pid == value.PID {
			d.Products = append(d.Products[:i], d.Products[i+1:]...)

			return
		}
	}
}

func (d *VendorProducts) IsEmpty() bool {
	return len(d.Products) == 0
}

func (d VendorProducts) String() string {
	bytes, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

// Single Vendor Product.
type Product struct {
	PID   uint16         `json:"pid"`
	Name  string         `json:"name"`
	SKU   string         `json:"sku"`
	Owner sdk.AccAddress `json:"owner"`
}
