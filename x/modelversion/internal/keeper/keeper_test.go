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

//nolint:testpackage
package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelversion/internal/types"
)

func TestKeeper_ModelGetSet(t *testing.T) {
	setup := Setup()

	// check if model present
	require.False(t, setup.ModelVersionKeeper.IsModelVersionPresent(setup.Ctx, testconstants.VID, testconstants.PID, testconstants.SoftwareVersion))

	// no model before its created
	require.Panics(t, func() {
		setup.ModelVersionKeeper.GetModelVersion(setup.Ctx, testconstants.VID, testconstants.PID, testconstants.SoftwareVersion)
	})

	// create model and version
	defaultModelVersion := DefaultModelVersion()
	AddModel(setup, defaultModelVersion.VID, defaultModelVersion.PID)
	setup.ModelVersionKeeper.SetModelVersion(setup.Ctx, DefaultModelVersion())

	// check if model present
	require.True(t, setup.ModelVersionKeeper.IsModelVersionPresent(setup.Ctx, defaultModelVersion.VID, defaultModelVersion.PID, defaultModelVersion.SoftwareVersion))

	// get receivedModelVersion info
	receivedModelVersion := setup.ModelVersionKeeper.GetModelVersion(setup.Ctx, defaultModelVersion.VID, defaultModelVersion.PID, defaultModelVersion.SoftwareVersion)
	require.NotNil(t, receivedModelVersion)
	require.Equal(t, defaultModelVersion.VID, receivedModelVersion.VID)
	require.Equal(t, defaultModelVersion.PID, receivedModelVersion.PID)
	require.Equal(t, defaultModelVersion.SoftwareVersion, receivedModelVersion.SoftwareVersion)
	require.Equal(t, defaultModelVersion.SoftwareVersionString, receivedModelVersion.SoftwareVersionString)
	require.Equal(t, defaultModelVersion.CDVersionNumber, receivedModelVersion.CDVersionNumber)
	require.Equal(t, defaultModelVersion.FirmwareDigests, receivedModelVersion.FirmwareDigests)
	require.Equal(t, defaultModelVersion.SoftwareVersionValid, receivedModelVersion.SoftwareVersionValid)
	require.Equal(t, defaultModelVersion.OtaURL, receivedModelVersion.OtaURL)
	require.Equal(t, defaultModelVersion.OtaFileSize, receivedModelVersion.OtaFileSize)
	require.Equal(t, defaultModelVersion.OtaChecksum, receivedModelVersion.OtaChecksum)
	require.Equal(t, defaultModelVersion.OtaChecksumType, receivedModelVersion.OtaChecksumType)
	require.Equal(t, defaultModelVersion.MinApplicableSoftwareVersion, receivedModelVersion.MinApplicableSoftwareVersion)
	require.Equal(t, defaultModelVersion.MaxApplicableSoftwareVersion, receivedModelVersion.MaxApplicableSoftwareVersion)
	require.Equal(t, defaultModelVersion.ReleaseNotesURL, receivedModelVersion.ReleaseNotesURL)

}

func TestKeeper_ModelVersionIterator(t *testing.T) {
	setup := Setup()

	count := 10

	// add 10 models infos with same VID/PID and check associated products
	PopulateStoreWithModelVersions(setup, count)

	// get total count
	totalModelVersions := setup.ModelVersionKeeper.CountTotalModelVersions(setup.Ctx, testconstants.VID, testconstants.PID)
	require.Equal(t, count, totalModelVersions)

	// get iterator
	var expectedRecords []types.ModelVersion

	setup.ModelVersionKeeper.IterateModelVersions(setup.Ctx, func(model types.ModelVersion) (stop bool) {
		expectedRecords = append(expectedRecords, model)

		return false
	})
	require.Equal(t, count, len(expectedRecords))
}
