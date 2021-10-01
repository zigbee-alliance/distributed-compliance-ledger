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

/*
	Vendor
*/

type VendorInfo struct {
	VendorId             uint16 `json:"vendorId" validate:"required"`
	VendorName           string `json:"vendorName" validate:"requiredForAdd,min=2,max=32"`
	CompanyLegalName     string `json:"companyLegalName" validate:"requiredForAdd,min=2,max=64"`
	CompanyPreferredName string `json:"companyPreferredName" validate:"max=64"`
	VendorLandingPageUrl string `json:"vendorLandingPageUrl" validate:"omitempty,max=256,url"`
}

// NewVendor creates a new Vendor object.
func NewVendorInfo(vendorId uint16, vendorName string, companyLegalName string,
	companyPreferredName string, vendorLandingPageUrl string) VendorInfo {
	return VendorInfo{
		VendorId:             vendorId,
		VendorName:           vendorName,
		CompanyLegalName:     companyLegalName,
		CompanyPreferredName: companyPreferredName,
		VendorLandingPageUrl: vendorLandingPageUrl,
	}
}
