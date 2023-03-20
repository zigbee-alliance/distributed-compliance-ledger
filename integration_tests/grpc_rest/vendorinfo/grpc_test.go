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

package vendorinfo_test

import (
	"testing"

	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/grpc_rest/vendorinfo"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

func TestVendorInfoDemoGRPC(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, false)
	vendorinfo.Demo(&suite)
}

func TestAddVendorInfoByNonVendorGRPC(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, false)
	vendorinfo.AddVendorInfoByNonVendor(&suite)
}

func TestAddVendorInfoByDifferentVendorGRPC(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, false)
	vendorinfo.AddVendorInfoByDifferentVendor(&suite)
}

func TestAddVendorInfoByNonVendorAdminGRPC(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, false)
	vendorinfo.AddVendorInfoByNonVendorAdmin(&suite)
}

func TestGetVendorInfoForUnknownGRPC(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, false)
	vendorinfo.GetVendorInfoForUnknown(&suite)
}

func TestGetVendorInfoForInvalidVidGRPC(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, false)
	vendorinfo.GetVendorInfoForInvalidVid(&suite)
}

func TestAddVendorInfoTwiceGRPC(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, false)
	vendorinfo.AddVendorInfoTwice(&suite)
}

func TestAddVendorInfoByVendorAdminGRPC(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, false)
	vendorinfo.AddVendorInfoByVendorAdmin(&suite)
}

func TestUpdateVendorInfoByVendorAdminGRPC(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, false)
	vendorinfo.UpdateVendorInfoByVendorAdmin(&suite)
}
