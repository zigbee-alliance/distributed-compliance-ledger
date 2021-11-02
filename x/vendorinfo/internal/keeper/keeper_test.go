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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/internal/types"
)

func TestKeeper_VendorInfoGetSet(t *testing.T) {
	setup := Setup()

	// check if vendor is present
	require.False(t, setup.VendorInfoKeeper.IsVendorInfoPresent(setup.Ctx, testconstants.VID))

	// no vendor info before its created
	require.Panics(t, func() {
		setup.VendorInfoKeeper.GetVendorInfo(setup.Ctx, testconstants.VID)
	})

	// create vendor info
	setup.VendorInfoKeeper.SetVendorInfo(setup.Ctx, DefaultVendorInfo())

	// check if vendor info is present
	require.True(t, setup.VendorInfoKeeper.IsVendorInfoPresent(setup.Ctx, testconstants.VID))

	// get vendorInfo info
	vendorInfo := setup.VendorInfoKeeper.GetVendorInfo(setup.Ctx, testconstants.VID)
	require.NotNil(t, vendorInfo)
	require.Equal(t, testconstants.VID, vendorInfo.VendorID)
	require.Equal(t, testconstants.VendorName, vendorInfo.VendorName)
	require.Equal(t, testconstants.CompanyLegalName, vendorInfo.CompanyLegalName)
	require.Equal(t, testconstants.CompanyPreferredName, vendorInfo.CompanyPreferredName)
	require.Equal(t, testconstants.VendorLandingPageURL, vendorInfo.VendorLandingPageURL)

	// Update the vendorInfo record with new values
	vendorInfo.VendorName = testconstants.VendorName + "updated"
	vendorInfo.CompanyLegalName = testconstants.CompanyLegalName + "updated"
	vendorInfo.CompanyPreferredName = testconstants.CompanyPreferredName + "updated"
	vendorInfo.VendorLandingPageURL = testconstants.VendorLandingPageURL + "updated"
	setup.VendorInfoKeeper.SetVendorInfo(setup.Ctx, vendorInfo)

	updateVendorInfo := setup.VendorInfoKeeper.GetVendorInfo(setup.Ctx, testconstants.VID)
	require.Equal(t, testconstants.VID, updateVendorInfo.VendorID)
	require.Equal(t, testconstants.VendorName+"updated", updateVendorInfo.VendorName)
	require.Equal(t, testconstants.CompanyLegalName+"updated", updateVendorInfo.CompanyLegalName)
	require.Equal(t, testconstants.CompanyPreferredName+"updated", updateVendorInfo.CompanyPreferredName)
	require.Equal(t, testconstants.VendorLandingPageURL+"updated", updateVendorInfo.VendorLandingPageURL)
}

func TestKeeper_VendorInfoIterator(t *testing.T) {
	setup := Setup()

	count := 10

	// add 10 models infos with same VID and check associated products
	PopulateStoreWithVendorInfo(setup, count)

	// get total count
	totalVendorInfos := setup.VendorInfoKeeper.CountTotalVendorInfos(setup.Ctx)
	require.Equal(t, count, totalVendorInfos)

	// get iterator
	var expectedRecords []types.VendorInfo

	setup.VendorInfoKeeper.IterateVendorInfos(setup.Ctx, func(model types.VendorInfo) (stop bool) {
		expectedRecords = append(expectedRecords, model)

		return false
	})
	require.Equal(t, count, len(expectedRecords))
}
