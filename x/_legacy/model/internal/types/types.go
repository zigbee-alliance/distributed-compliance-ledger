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
)

//nolint:maligned,lll
type Model struct {
	VID                                        uint16 `json:"vid" validate:"required"`
	PID                                        uint16 `json:"pid" validate:"required"`
	DeviceTypeID                               uint16 `json:"deviceTypeID,omitempty" validate:"requiredForAdd"`
	ProductName                                string `json:"productName,omitempty" validate:"requiredForAdd,max=128"`
	ProductLabel                               string `json:"productLabel,omitempty" validate:"requiredForAdd,max=256"`
	PartNumber                                 string `json:"partNumber,omitempty" validate:"requiredForAdd,max=32"`
	CommissioningCustomFlow                    uint8  `json:"commissioningCustomFlow" validate:"max=2"`
	CommissioningCustomFlowURL                 string `json:"commissioningCustomFlowURL,omitempty" validate:"omitempty,url,startsnotwith=http:,max=256,required_if=commissioningCustomFlow 2"`
	CommissioningModeInitialStepsHint          uint32 `json:"commissioningModeInitialStepsHint,omitempty"`
	CommissioningModeInitialStepsInstruction   string `json:"commissioningModeInitialStepsInstruction,omitempty" validate:"max=1024"`
	CommissioningModeSecondaryStepsHint        uint32 `json:"commissioningModeSecondaryStepsHint,omitempty"`
	CommissioningModeSecondaryStepsInstruction string `json:"commissioningModeSecondaryStepsInstruction,omitempty" validate:"max=1024"`
	UserManualURL                              string `json:"userManualURL,omitempty" validate:"omitempty,url,startsnotwith=http:,max=256"`
	SupportURL                                 string `json:"supportURL,omitempty" validate:"omitempty,url,startsnotwith=http:,max=256"`
	ProductURL                                 string `json:"productURL,omitempty" validate:"omitempty,url,startsnotwith=http:,max=256"`
}

func (d Model) String() string {
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
	PID        uint16 `json:"pid"`
	Name       string `json:"name"`
	PartNumber string `json:"partNumber"`
}

// Model Versions.
//nolint:maligned,lll
type ModelVersion struct {
	VID                          uint16 `json:"vid" validate:"required"`
	PID                          uint16 `json:"pid" validate:"required"`
	SoftwareVersion              uint32 `json:"softwareVersion" validate:"required"`
	SoftwareVersionString        string `json:"softwareVersionString,omitempty" validate:"requiredForAdd,max=64"`
	CDVersionNumber              uint16 `json:"CDVersionNumber,omitempty" validate:"requiredForAdd"`
	FirmwareDigests              string `json:"firmwareDigests,omitempty" validate:"max=512"`
	SoftwareVersionValid         bool   `json:"softwareVersionValid"`
	OtaURL                       string `json:"otaURL,omitempty" validate:"omitempty,url,startsnotwith=http:,max=256"`
	OtaFileSize                  uint64 `json:"otaFileSize,omitempty" validate:"required_with=OtaURL"`
	OtaChecksum                  string `json:"otaChecksum,omitempty" validate:"required_with=OtaURL,max=64"`
	OtaChecksumType              uint16 `json:"otaChecksumType,omitempty" validate:"required_with=OtaURL"`
	MinApplicableSoftwareVersion uint32 `json:"minApplicableSoftwareVersion,omitempty" validate:"requiredForAdd"`
	MaxApplicableSoftwareVersion uint32 `json:"maxApplicableSoftwareVersion,omitempty" validate:"requiredForAdd,gtecsfield=MinApplicableSoftwareVersion"`
	ReleaseNotesURL              string `json:"releaseNotesURL,omitempty" validate:"omitempty,url,startsnotwith=http:,max=256"`
}

func (d ModelVersion) String() string {
	bytes, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

type ModelVersions struct {
	VID              uint16   `json:"vid"`
	PID              uint16   `json:"pid"`
	SoftwareVersions []uint32 `json:"softwareVersions"`
}
