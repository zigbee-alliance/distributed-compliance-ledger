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

package vendorinfo

import (
	"strconv"

	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
	vendorinfotypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
)

// VendorOpts holds parameters for add-vendor / update-vendor.
// VID is required for both (>0 emits --vid); other fields are only emitted
// when non-empty.
type VendorOpts struct {
	VID                  int
	VIDHex               string
	VendorName           string
	CompanyLegalName     string
	CompanyPreferredName string
	VendorLandingPageURL string
	SchemaVersion        string
}

func (o VendorOpts) args() []string {
	var args []string
	if o.VID != 0 || o.VIDHex != "" {
		args = append(args, "--vid", cliputils.FlagOrHex(o.VID, o.VIDHex))
	}
	if o.VendorName != "" {
		args = append(args, "--vendorName", o.VendorName)
	}
	if o.CompanyLegalName != "" {
		args = append(args, "--companyLegalName", o.CompanyLegalName)
	}
	if o.CompanyPreferredName != "" {
		args = append(args, "--companyPreferredName", o.CompanyPreferredName)
	}
	if o.VendorLandingPageURL != "" {
		args = append(args, "--vendorLandingPageURL", o.VendorLandingPageURL)
	}
	if o.SchemaVersion != "" {
		args = append(args, "--schemaVersion", o.SchemaVersion)
	}

	return args
}

// AddVendor adds a vendor info record.
func AddVendor(from string, opts VendorOpts) (*utils.TxResult, error) {
	args := []string{"tx", "vendorinfo", "add-vendor", "--from", from}
	args = append(args, opts.args()...)

	return utils.ExecuteTx(args...)
}

// UpdateVendor updates a vendor info record.
func UpdateVendor(from string, opts VendorOpts) (*utils.TxResult, error) {
	args := []string{"tx", "vendorinfo", "update-vendor", "--from", from}
	args = append(args, opts.args()...)

	return utils.ExecuteTx(args...)
}

// GetVendor queries a vendor info record by VID. Returns nil when no record
// exists.
func GetVendor(vid int) (*vendorinfotypes.VendorInfo, error) {
	return GetVendorHex(strconv.Itoa(vid))
}

// GetVendorHex queries a vendor info record using a hex VID string.
func GetVendorHex(vid string) (*vendorinfotypes.VendorInfo, error) {
	var res vendorinfotypes.VendorInfo
	found, err := cliputils.GetSingle(&res,
		"query", "vendorinfo", "vendor",
		"--vid", vid,
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetAllVendors queries all vendor info records.
func GetAllVendors() ([]vendorinfotypes.VendorInfo, error) {
	var res vendorinfotypes.QueryAllVendorInfoResponse
	if err := cliputils.GetList(&res, "query", "vendorinfo", "all-vendors", "-o", "json"); err != nil {
		return nil, err
	}

	return res.VendorInfo, nil
}

// containsVendorByID reports whether list has a VendorInfo with the given VID.
func containsVendorByID(list []vendorinfotypes.VendorInfo, vid int32) bool {
	for i := range list {
		if list[i].VendorID == vid {
			return true
		}
	}

	return false
}
