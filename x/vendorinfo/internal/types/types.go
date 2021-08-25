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
	VendorId             uint16 `json:"vendorId"`
	VendorName           string `json:"vendorName"`
	CompanyLegalName     string `json:"companyLegaName"`
	CompanyPreferredName string `json:"companyPreferredName"`
	VendorLandingPageUrl string `json:"vendorLandingPageUrl"`
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

// Validate checks for errors on the vesting and module account parameters.
func (vendor VendorInfo) Validate() error {
	// if acc.Address == nil {
	// 	return sdk.ErrUnknownRequest(
	// 		fmt.Sprintf("Invalid Account: Value: %s. Error: Missing Address", acc.Address))
	// }

	// if acc.PubKey == nil {
	// 	return sdk.ErrUnknownRequest(
	// 		fmt.Sprintf("Invalid Account: Value: %s. Error: Missing PubKey", acc.PubKey))
	// }

	// if err := acc.Roles.Validate(); err != nil {
	// 	return err
	// }

	// // If creating an account with Vendor Role, we need to have a associated VendorId
	// if acc.HasRole(Vendor) && acc.VendorId <= 0 {
	// 	return ErrMissingVendorIdForVendorAccount()
	// }

	return nil
}
