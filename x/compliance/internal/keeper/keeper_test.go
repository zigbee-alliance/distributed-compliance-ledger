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

//nolint:goimports
import (
	"testing"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	"github.com/stretchr/testify/require"
)

func TestKeeper_ComplianceInfoGetSet(t *testing.T) {
	setup := Setup()

	// check if compliance info present
	require.False(t, setup.CompliancetKeeper.IsComplianceInfoPresent(setup.Ctx,
		types.CertificationType(testconstants.CertificationType), testconstants.VID, testconstants.PID))

	// no compliance info before its created
	require.Panics(t, func() {
		setup.CompliancetKeeper.GetComplianceInfo(setup.Ctx,
			types.CertificationType(testconstants.CertificationType), testconstants.VID, testconstants.PID)
	})

	// create compliance info
	certifiedModel := DefaultCertifiedModel()
	setup.CompliancetKeeper.SetComplianceInfo(setup.Ctx, certifiedModel)

	// check if compliance info present
	require.True(t, setup.CompliancetKeeper.IsComplianceInfoPresent(setup.Ctx,
		types.CertificationType(testconstants.CertificationType), testconstants.VID, testconstants.PID))

	// get compliance info
	receivedComplianceInfo := setup.CompliancetKeeper.GetComplianceInfo(setup.Ctx,
		types.CertificationType(testconstants.CertificationType), testconstants.VID, testconstants.PID)
	CheckComplianceInfo(t, certifiedModel, receivedComplianceInfo)
}

func TestKeeper_ComplianceInfoIterator(t *testing.T) {
	setup := Setup()

	count := 10

	// add 10 models
	PopulateStoreWithMixedModels(setup, count)

	// get total count
	totalModes := setup.CompliancetKeeper.CountTotalComplianceInfo(
		setup.Ctx, types.CertificationType(testconstants.CertificationType))
	require.Equal(t, count, totalModes)

	// get iterator
	var expectedRecords []types.ComplianceInfo

	setup.CompliancetKeeper.IterateComplianceInfos(setup.Ctx, types.CertificationType(testconstants.CertificationType),
		func(modelInfo types.ComplianceInfo) (stop bool) {
			expectedRecords = append(expectedRecords, modelInfo)

			return false
		})
	require.Equal(t, count, len(expectedRecords))
}

func TestKeeper_TwoComplianceInfoWithDifferentType(t *testing.T) {
	setup := Setup()

	// create zb compliance info
	zbCertifiedModel := DefaultCertifiedModel()
	setup.CompliancetKeeper.SetComplianceInfo(setup.Ctx, zbCertifiedModel)

	// create other compliance info
	otherCertifiedModel := DefaultCertifiedModel()
	otherCertifiedModel.CertificationType = "Other"
	setup.CompliancetKeeper.SetComplianceInfo(setup.Ctx, otherCertifiedModel)

	// get zb compliance info
	receivedComplianceInfo := setup.CompliancetKeeper.GetComplianceInfo(setup.Ctx,
		zbCertifiedModel.CertificationType, zbCertifiedModel.VID, zbCertifiedModel.PID)
	CheckComplianceInfo(t, zbCertifiedModel, receivedComplianceInfo)

	// get other compliance info
	receivedComplianceInfo = setup.CompliancetKeeper.GetComplianceInfo(setup.Ctx,
		otherCertifiedModel.CertificationType, otherCertifiedModel.VID, otherCertifiedModel.PID)
	CheckComplianceInfo(t, otherCertifiedModel, receivedComplianceInfo)
}
