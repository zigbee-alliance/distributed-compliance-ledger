package compliance

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	dclcompltypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/compliance"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
)

func setupDeleteComplianceInfo(t *testing.T) (*TestSetup, int32, int32, uint32, string, string) {
	setup := setup(t)

	vid, pid, softwareVersion, softwareVersionString := setup.addModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)
	certificationType := dclcompltypes.ZigbeeCertificationType

	return setup, vid, pid, softwareVersion, softwareVersionString, certificationType
}

func (setup *TestSetup) deleteComplianceInfo(vid int32, pid int32, softwareVersion uint32, certificationType string, signer sdk.AccAddress) (*types.MsgDeleteComplianceInfo, error) {
	deleteComplInfoMsg := newMsgDeleteComplianceInfo(
		vid, pid, softwareVersion, certificationType, signer)
	_, err := setup.Handler(setup.Ctx, deleteComplInfoMsg)

	return deleteComplInfoMsg, err
}

func (setup *TestSetup) checkComplianceInfoDeleted(t *testing.T, deleteComplInfoMsg *types.MsgDeleteComplianceInfo) {
	vid := deleteComplInfoMsg.Vid
	pid := deleteComplInfoMsg.Pid
	softwareVersion := deleteComplInfoMsg.SoftwareVersion
	certificationType := deleteComplInfoMsg.CertificationType

	_, err := queryProvisionalModel(setup, vid, pid, softwareVersion, certificationType)
	assertNotFound(t, err)

	_, err = queryCertifiedModel(setup, vid, pid, softwareVersion, certificationType)
	assertNotFound(t, err)

	_, err = queryRevokedModel(setup, vid, pid, softwareVersion, certificationType)
	assertNotFound(t, err)
}

func TestHandler_DeleteComplianceInfoForRevokedModel(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupDeleteComplianceInfo(t)

	_, revokeModelErr := setup.revokeModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	require.NoError(t, revokeModelErr)

	deleteComplInfoMsg, deleteComplInfoError := setup.deleteComplianceInfo(vid, pid, softwareVersion, certificationType, setup.CertificationCenter)
	require.NoError(t, deleteComplInfoError)

	setup.checkComplianceInfoDeleted(t, deleteComplInfoMsg)
}

func TestHandler_DeleteComplianceInfoForCertifiedModel(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupDeleteComplianceInfo(t)

	certifyModelMsg := newMsgCertifyModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	_, certifyModelErr := setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, certifyModelErr)

	deleteComplInfoMsg, deleteComplInfoError := setup.deleteComplianceInfo(vid, pid, softwareVersion, certificationType, setup.CertificationCenter)
	require.NoError(t, deleteComplInfoError)

	setup.checkComplianceInfoDeleted(t, deleteComplInfoMsg)
}

func TestHandler_DeleteComplianceInfoForProvisionalModel(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupDeleteComplianceInfo(t)

	_, provisionModelErr := setup.provisionModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	require.NoError(t, provisionModelErr)

	deleteComplInfoMsg, deleteComplInfoError := setup.deleteComplianceInfo(vid, pid, softwareVersion, certificationType, setup.CertificationCenter)
	require.NoError(t, deleteComplInfoError)

	setup.checkComplianceInfoDeleted(t, deleteComplInfoMsg)
}

func TestHandler_DeleteComplianceInfoDoesNotExist(t *testing.T) {
	setup, vid, pid, softwareVersion, _, certificationType := setupDeleteComplianceInfo(t)

	_, deleteComplInfoError := setup.deleteComplianceInfo(vid, pid, softwareVersion, certificationType, setup.CertificationCenter)
	require.ErrorIs(t, deleteComplInfoError, types.ErrComplianceInfoDoesNotExist)
}
