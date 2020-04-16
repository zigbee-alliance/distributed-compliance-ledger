package keeper

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKeeper_CertifiedModelGetSet(t *testing.T) {
	setup := Setup()

	// check if certified model present
	require.False(t, setup.CompliancetKeeper.IsCertifiedModelPresent(setup.Ctx, test_constants.VID, test_constants.PID))

	// no certified model before its created
	require.Panics(t, func() {
		setup.CompliancetKeeper.GetCertifiedModel(setup.Ctx, test_constants.VID, test_constants.PID)
	})

	// create certified model
	certifiedModel := DefaultCertifiedModel()
	setup.CompliancetKeeper.SetCertifiedModel(setup.Ctx, certifiedModel)

	// check if certified model present
	require.True(t, setup.CompliancetKeeper.IsCertifiedModelPresent(setup.Ctx, test_constants.VID, test_constants.PID))

	// get certified model
	receivedCertifiedModel := setup.CompliancetKeeper.GetCertifiedModel(setup.Ctx, test_constants.VID, test_constants.PID)
	require.Equal(t, test_constants.VID, receivedCertifiedModel.VID)
	require.Equal(t, test_constants.PID, receivedCertifiedModel.PID)
	require.Equal(t, test_constants.CertificationDate, receivedCertifiedModel.CertificationDate)
	require.Equal(t, test_constants.CertificationType, receivedCertifiedModel.CertificationType)
}

func TestKeeper_CertifiedModelIterator(t *testing.T) {
	setup := Setup()

	count := 10

	// add 10 models
	PopulateStoreWithCertifiedModels(setup, count)

	// get total count
	totalModes := setup.CompliancetKeeper.CountTotalCertifiedModel(setup.Ctx)
	require.Equal(t, count, totalModes)

	// get iterator
	var expectedRecords []types.CertifiedModel

	setup.CompliancetKeeper.IterateCertifiedModels(setup.Ctx, func(modelInfo types.CertifiedModel) (stop bool) {
		expectedRecords = append(expectedRecords, modelInfo)
		return false
	})
	require.Equal(t, count, len(expectedRecords))
}
