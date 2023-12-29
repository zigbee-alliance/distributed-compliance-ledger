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

func TestModelDemoGRPC(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, false)
	model.Demo(&suite)
}

func TestAddModelByNonVendorGRPC(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, false)
	model.AddModelByNonVendor(&suite)
}

func TestAddModelByDifferentVendorGRPC(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, false)
	model.AddModelByDifferentVendor(&suite)
}

func TestAddModelTwiceGRPC(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, false)
	model.AddModelTwice(&suite)
}

func TestGetModelForUnknownGRPC(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, false)
	model.GetModelForUnknown(&suite)
}

func TestGetModelForInvalidVidPidGRPC(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, false)
	model.GetModelForInvalidVidPid(&suite)
}

func TestDeleteModelWithAssociatedModelVersionsGRPC(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, false)
	model.DeleteModelWithAssociatedModelVersions(&suite)
}

func TestDeleteModelVersionGRPC(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, false)
	model.DeleteModelVersion(&suite)
}

func TestDeleteModelVersionDifferentVidGRPC(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, false)
	model.DeleteModelVersionDifferentVid(&suite)
}

func TestDeleteModelVersionDoesNotExistGRPC(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, false)
	model.DeleteModelVersionDoesNotExist(&suite)
}

func TestDeleteModelVersionNotByCreatorGRPC(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, false)
	model.DeleteModelVersionNotByCreator(&suite)
}

func TestDeleteModelVersionCertifiedGRPC(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, false)
	model.DeleteModelVersionCertified(&suite)
}
func TestAddModelByVendorProductIdsGRPC(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, false)
	model.AddModelByVendorWithProductIds(&suite)
}

func TestUpdateByVendorWithProductIdsGRPC(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, false)
	model.UpdateByVendorWithProductIds(&suite)
}

func TestAddModelByVendorWithNonAssociatedProductIdsGRPC(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, false)
	model.AddModelByVendorWithNonAssociatedProductIds(&suite)
}

func TestUpdateModelByVendorWithNonAssociatedProductIdsGRPC(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, false)
	model.UpdateModelByVendorWithNonAssociatedProductIds(&suite)
}

func TestDeleteModelByVendorWithNonAssociatedProductIdsGRPC(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, false)
	model.DeleteModelByVendorWithNonAssociatedProductIds(&suite)
}
