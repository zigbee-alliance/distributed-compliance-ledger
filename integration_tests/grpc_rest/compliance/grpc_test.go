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

package compliance_test

import (
	"testing"

	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/grpc_rest/compliance"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

//nolint:godox
/*
	To Run test you need:
		* Run LocalNet with: `make install && make localnet_init && make localnet_start`
	TODO: provide tests for error cases
*/

func TestComplianceDemoTrackComplianceGRPC(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, false)
	compliance.ComplianceDemoTrackCompliance(&suite)
}

func TestComplianceDemoTrackRevocationGRPC(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, false)
	compliance.ComplianceDemoTrackRevocation(&suite)
}

func TestComplianceDemoTrackProvisionGRPC(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, false)
	compliance.ComplianceDemoTrackProvision(&suite)
}
