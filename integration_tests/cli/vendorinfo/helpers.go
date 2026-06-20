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
	"encoding/json"
	"fmt"

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
	Extra                []string
}

func (o VendorOpts) args() []string {
	var args []string
	if o.VID != 0 || o.VIDHex != "" {
		args = append(args, "--vid", flagOrHex(o.VID, o.VIDHex))
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

	return append(args, o.Extra...)
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

// flagOrHex returns hex if non-empty, otherwise the decimal-formatted n.
func flagOrHex(n int, hex string) string {
	if hex != "" {
		return hex
	}

	return itoa(n)
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	neg := false
	if n < 0 {
		neg = true
		n = -n
	}
	var buf [20]byte
	pos := len(buf)
	for n > 0 {
		pos--
		buf[pos] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		pos--
		buf[pos] = '-'
	}

	return string(buf[pos:])
}

// getSingle runs a single-item dcld query and unmarshals into v. Returns
// (false, nil) when the CLI emitted "Not Found".
func getSingle(v interface{}, args ...string) (found bool, err error) {
	out, err := utils.ExecuteCLI(args...)
	if err != nil {
		return false, err
	}
	if utils.IsNotFound(out) {
		return false, nil
	}
	out = utils.NormalizeProtoJSON(out)
	if err := json.Unmarshal(out, v); err != nil {
		return false, fmt.Errorf("parse %T: %w, output: %s", v, err, string(out))
	}

	return true, nil
}

// getList runs an all-* dcld query and unmarshals the wrapper response.
func getList(v interface{}, args ...string) error {
	out, err := utils.ExecuteCLI(args...)
	if err != nil {
		return err
	}
	out = utils.NormalizeProtoJSON(utils.StripPagination(out))
	if err := json.Unmarshal(out, v); err != nil {
		return fmt.Errorf("parse %T: %w, output: %s", v, err, string(out))
	}

	return nil
}

// GetVendor queries a vendor info record by VID. Returns nil when no record
// exists.
func GetVendor(vid int) (*vendorinfotypes.VendorInfo, error) {
	return GetVendorHex(itoa(vid))
}

// GetVendorHex queries a vendor info record using a hex VID string.
func GetVendorHex(vid string) (*vendorinfotypes.VendorInfo, error) {
	var res vendorinfotypes.VendorInfo
	found, err := getSingle(&res,
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
	if err := getList(&res, "query", "vendorinfo", "all-vendors", "-o", "json"); err != nil {
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
