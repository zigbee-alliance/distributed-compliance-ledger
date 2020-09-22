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
	VID                      uint16         `json:"vid"`
	PID                      uint16         `json:"pid"`
	CID                      uint16         `json:"cid,omitempty"`
	Version                  string         `json:"version,omitempty"`
	Name                     string         `json:"name"`
	Description              string         `json:"description"`
	SKU                      string         `json:"sku"`
	HardwareVersion          string         `json:"hardware_version"`
	FirmwareVersion          string         `json:"firmware_version"`
	OtaURL                   string         `json:"ota_url,omitempty"`
	OtaChecksum              string         `json:"ota_checksum,omitempty"`
	OtaChecksumType          string         `json:"ota_checksum_type,omitempty"`
	Custom                   string         `json:"custom,omitempty"`
	TisOrTrpTestingCompleted bool           `json:"tis_or_trp_testing_completed"`
	Owner                    sdk.AccAddress `json:"owner"`
}

func NewModelInfo(
	vid uint16,
	pid uint16,
	cid uint16,
	version string,
	name string,
	description string,
	sku string,
	hardwareVersion string,
	firmwareVersion string,
	otaURL string,
	otaChecksum string,
	otaChecksumType string,
	custom string,
	tisOrTrpTestingCompleted bool,
	owner sdk.AccAddress,
) ModelInfo {
	return ModelInfo{
		VID:                      vid,
		PID:                      pid,
		CID:                      cid,
		Version:                  version,
		Name:                     name,
		Description:              description,
		SKU:                      sku,
		HardwareVersion:          hardwareVersion,
		FirmwareVersion:          firmwareVersion,
		OtaURL:                   otaURL,
		OtaChecksum:              otaChecksum,
		OtaChecksumType:          otaChecksumType,
		Custom:                   custom,
		TisOrTrpTestingCompleted: tisOrTrpTestingCompleted,
		Owner:                    owner,
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
