//nolint:testpackage
package keeper

//nolint:goimports
import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	"github.com/stretchr/testify/require"
	"testing"
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
