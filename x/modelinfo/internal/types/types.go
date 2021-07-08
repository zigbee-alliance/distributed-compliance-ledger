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
	Model Model          `json:"model"`
	Owner sdk.AccAddress `json:"owner"`
}

func NewModelInfo(
	model Model,
	owner sdk.AccAddress,

) ModelInfo {
	return ModelInfo{
		Model: model,
		Owner: owner,
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

func (d *VendorProducts) RemoveVendorProduct(pid uint16,
	softwareVersion uint32, hardwareVersion uint32) {
	for i, value := range d.Products {
		if pid == value.PID &&
			softwareVersion == value.SoftwareVersion &&
			hardwareVersion == value.HardwareVersion {
			d.Products = append(d.Products[:i], d.Products[i+1:]...)
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

/*
	List of versions for specific Product
*/
type ProductVersions struct {
	VID      uint16    `json:"vid"`
	PID      uint16    `json:"pid"`
	Versions []Version `json:"products"`
}

func NewProductVersion(vid uint16, pid uint16) ProductVersions {
	return ProductVersions{
		VID:      vid,
		PID:      pid,
		Versions: []Version{},
	}
}

func (d *ProductVersions) AddProductVersion(version Version) {
	d.Versions = append(d.Versions, version)
}

func (d *ProductVersions) RemoveProductVersion(softwareVersion uint32, hardwareVersion uint32) {
	for i, value := range d.Versions {
		if softwareVersion == value.SoftwareVersion && hardwareVersion == value.HardwareVersion {
			d.Versions = append(d.Versions[:i], d.Versions[i+1:]...)

			return
		}
	}
}

func (d *ProductVersions) IsEmpty() bool {
	return len(d.Versions) == 0
}

func (d ProductVersions) String() string {
	bytes, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

// Single Vendor Product.
type Product struct {
	PID             uint16         `json:"pid"`
	Name            string         `json:"name"`
	SKU             string         `json:"sku"`
	SoftwareVersion uint32         `json:"softwareVersion"`
	HardwareVersion uint32         `json:"hardwareVersion"`
	Owner           sdk.AccAddress `json:"owner"`
}

// Single Product Version.
type Version struct {
	SoftwareVersion uint32         `json:"softwareVersion"`
	HardwareVersion uint32         `json:"hardwareVersion"`
	Name            string         `json:"name"`
	SKU             string         `json:"sku"`
	Owner           sdk.AccAddress `json:"owner"`
}
