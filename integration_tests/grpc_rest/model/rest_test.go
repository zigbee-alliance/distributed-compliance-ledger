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

package model_test

import (
	"testing"

	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/grpc_rest/model"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

//nolint:godox
/*
	To Run test you need:
		* Run LocalNet with: `make install && make localnet_init && make localnet_start`
*/

func TestModelDemoREST(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, true)
	model.Demo(&suite)
}

func TestAddModelByNonVendorREST(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, true)
	model.AddModelByNonVendor(&suite)
}

func TestAddModelByDifferentVendorREST(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, true)
	model.AddModelByDifferentVendor(&suite)
}

func TestAddModelTwiceREST(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, true)
	model.AddModelTwice(&suite)
}

func TestGetModelForUnknownREST(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, true)
	model.GetModelForUnknown(&suite)
}

func TestGetModelForInvalidVidPidREST(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, true)
	model.GetModelForInvalidVidPid(&suite)
}

func TestAddModelInHexFormat(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, true)
	model.DemoWithHexVidAndPid(&suite)
}
