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

//nolint:maligned
type Model struct {
	VID                                        uint16 `json:"vid"`
	PID                                        uint16 `json:"pid"`
	DeviceTypeID                               uint16 `json:"deviceTypeID,omitempty"`
	ProductName                                string `json:"productName,omitempty"`
	ProductLabel                               string `json:"productLabel,omitempty"`
	PartNumber                                 string `json:"partNumber,omitempty"`
	CommissioningCustomFlow                    uint8  `json:"commissioningCustomFlow,omitempty"`
	CommissioningCustomFlowURL                 string `json:"commissioningCustomFlowURL,omitempty"`
	CommissioningModeInitialStepsHint          uint32 `json:"commissioningModeInitialStepsHint,omitempty"`
	CommissioningModeInitialStepsInstruction   string `json:"commissioningModeInitialStepsInstruction,omitempty"`
	CommissioningModeSecondaryStepsHint        uint32 `json:"commissioningModeSecondaryStepsHint,omitempty"`
	CommissioningModeSecondaryStepsInstruction string `json:"commissioningModeSecondaryStepsInstruction,omitempty"`
	UserManualURL                              string `json:"userManualURL,omitempty"`
	SupportURL                                 string `json:"supportURL,omitempty"`
	ProductURL                                 string `json:"productURL,omitempty"`
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
