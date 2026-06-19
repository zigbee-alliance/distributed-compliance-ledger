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
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
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

// QueryVendor queries a vendor info record by VID.
func QueryVendor(vid string) ([]byte, error) {
	return utils.ExecuteCLI("query", "vendorinfo", "vendor",
		"--vid", vid,
		"-o", "json",
	)
}

// QueryAllVendors queries all vendor info records.
func QueryAllVendors() ([]byte, error) {
	return utils.ExecuteCLI("query", "vendorinfo", "all-vendors", "-o", "json")
}
