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

//nolint:maligned
type Model struct {
	VID                                        uint16 `json:"vid"`
	PID                                        uint16 `json:"pid"`
	CID                                        uint16 `json:"cid,omitempty"`
	ProductName                                string `json:"productName"`
	Description                                string `json:"description"`
	SKU                                        string `json:"sku"`
	SoftwareVersion                            uint32 `json:"softwareVersion"`
	SoftwareVersionString                      string `json:"softwareVersionString"`
	HardwareVersion                            uint32 `json:"hardwareVersion"`
	HardwareVersionString                      string `json:"hardwareVersionString"`
	CDVersionNumber                            uint16 `json:"CDVersionNumber"`
	FirmwareDigests                            string `json:"firmwareDigests,omitempty"`
	Revoked                                    bool   `json:"revoked"`
	OtaURL                                     string `json:"otaURL,omitempty"`
	OtaChecksum                                string `json:"otaChecksum,omitempty"`
	OtaChecksumType                            string `json:"otaChecksumType,omitempty"`
	OtaBlob                                    string `json:"otaBlob,omitempty"`
	CommissioningCustomFlow                    uint8  `json:"commissioningCustomFlow,omitempty"`
	CommissioningCustomFlowURL                 string `json:"commissioningCustomFlowURL,omitempty"`
	CommissioningModeInitialStepsHint          uint32 `json:"commissioningModeInitialStepsHint,omitempty"`
	CommissioningModeInitialStepsInstruction   string `json:"commissioningModeInitialStepsInstruction,omitempty"`
	CommissioningModeSecondaryStepsHint        uint32 `json:"commissioningModeSecondaryStepsHint,omitempty"`
	CommissioningModeSecondaryStepsInstruction string `json:"commissioningModeSecondaryStepsInstruction,omitempty"`
	ReleaseNotesURL                            string `json:"releaseNotesURL,omitempty"`
	UserManualURL                              string `json:"userManualURL,omitempty"`
	SupportURL                                 string `json:"supportURL,omitempty"`
	ProductURL                                 string `json:"productURL,omitempty"`
	ChipBlob                                   string `json:"chipBlob,omitempty"`
	VendorBlob                                 string `json:"vendorBlob,omitempty"`
}
