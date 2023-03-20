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

func TestVendorInfoDemoREST(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, true)
	vendorinfo.Demo(&suite)
}

func TestAddVendorInfoByNonVendorREST(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, true)
	vendorinfo.AddVendorInfoByNonVendor(&suite)
}

func TestAddVendorInfoByDifferentVendorREST(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, true)
	vendorinfo.AddVendorInfoByDifferentVendor(&suite)
}

func TestAddVendorInfoByNonVendorAdminREST(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, true)
	vendorinfo.AddVendorInfoByNonVendorAdmin(&suite)
}

func TestGetVendorInfoForUnknownREST(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, true)
	vendorinfo.GetVendorInfoForUnknown(&suite)
}

func TestGetVendorInfoForInvalidVidREST(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, true)
	vendorinfo.GetVendorInfoForInvalidVid(&suite)
}

func TestAddVendorInfoTwiceREST(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, true)
	vendorinfo.AddVendorInfoTwice(&suite)
}

func TestDemoWithHexVid(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, true)
	vendorinfo.DemoWithHexVid(&suite)
}

func TestAddVendorInfoByVendorAdminREST(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, true)
	vendorinfo.AddVendorInfoByVendorAdmin(&suite)
}

func TestUpdateVendorInfoByVendorAdminREST(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, true)
	vendorinfo.UpdateVendorInfoByVendorAdmin(&suite)
}
