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
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelversion/internal/types"
)

func Test_QueryModelVersion(t *testing.T) {
	setup := Setup()

	// add model and version
	modelVersion := AddModelVersion(setup)

	// query model version
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryModelVersion, fmt.Sprintf("%v", modelVersion.VID), fmt.Sprintf("%v", modelVersion.PID), fmt.Sprintf("%v", modelVersion.SoftwareVersion)},
		abci.RequestQuery{},
	)

	var receivedModelVersion types.ModelVersion
	_ = setup.Cdc.UnmarshalJSON(result, &receivedModelVersion)

	// check
	require.Equal(t, receivedModelVersion, modelVersion)

	// Query non existent model version
	result, err := setup.Querier(
		setup.Ctx,
		[]string{QueryModelVersion, fmt.Sprintf("%v", modelVersion.VID), fmt.Sprintf("%v", modelVersion.PID), fmt.Sprintf("%v", 123)},
		abci.RequestQuery{},
	)

	// check
	require.Nil(t, result)
	require.NotNil(t, err)
	require.Equal(t, types.CodeModelVersionDoesNotExist, err.Code())
}

func Test_QueryModelVersions(t *testing.T) {
	setup := Setup()
	count := 5

	// add model and version
	vid, pid := PopulateStoreWithModelVersions(setup, count)

	// query model versions
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllModelVersions, fmt.Sprintf("%v", vid), fmt.Sprintf("%v", pid)},
		abci.RequestQuery{},
	)

	var receivedModelVersions types.ModelVersions
	_ = setup.Cdc.UnmarshalJSON(result, &receivedModelVersions)

	// check
	require.Equal(t, vid, receivedModelVersions.VID)
	require.Equal(t, pid, receivedModelVersions.PID)
	require.Equal(t, count, len(receivedModelVersions.SoftwareVersions))
}
