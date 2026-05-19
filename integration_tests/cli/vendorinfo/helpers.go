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

// AddVendor adds a vendor info record.
func AddVendor(from string, extra ...string) (*utils.TxResult, error) {
	args := []string{"tx", "vendorinfo", "add-vendor", "--from", from}
	args = append(args, extra...)

	return utils.ExecuteTx(args...)
}

// UpdateVendor updates a vendor info record.
func UpdateVendor(from string, extra ...string) (*utils.TxResult, error) {
	args := []string{"tx", "vendorinfo", "update-vendor", "--from", from}
	args = append(args, extra...)

	return utils.ExecuteTx(args...)
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
