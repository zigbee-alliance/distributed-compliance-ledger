package compliance

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	dclcompltypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/compliance"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func DeleteComplianceInfoSetup(t *testing.T) (*TestSetup, int32, int32, uint32, string, string) {
	setup := Setup(t)

	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)
	certificationType := dclcompltypes.ZigbeeCertificationType

	return setup, vid, pid, softwareVersion, softwareVersionString, certificationType
}

func (setup *TestSetup) DeleteComplianceInfo(vid int32, pid int32, softwareVersion uint32, certificationType string, signer sdk.AccAddress) (*types.MsgDeleteComplianceInfo, error) {
	deleteComplInfoMsg := NewMsgDeleteComplianceInfo(
		vid, pid, softwareVersion, certificationType, signer)
	_, err := setup.Handler(setup.Ctx, deleteComplInfoMsg)

	return deleteComplInfoMsg, err
}

func (setup *TestSetup) CheckComplianceInfoDeleted(t *testing.T, deleteComplInfoMsg *types.MsgDeleteComplianceInfo) {
	vid := deleteComplInfoMsg.Vid
	pid := deleteComplInfoMsg.Pid
	softwareVersion := deleteComplInfoMsg.SoftwareVersion
	certificationType := deleteComplInfoMsg.CertificationType

	test, err := queryProvisionalModel(setup, vid, pid, softwareVersion, certificationType)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
	_ = test

	_, err = queryCertifiedModel(setup, vid, pid, softwareVersion, certificationType)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	_, err = queryRevokedModel(setup, vid, pid, softwareVersion, certificationType)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
}

func TestHandler_DeleteComplianceInfoForRevokedModel(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := DeleteComplianceInfoSetup(t)

	_, revokeModelErr := setup.RevokeModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	require.NoError(t, revokeModelErr)

	deleteComplInfoMsg, deleteComplInfoError := setup.DeleteComplianceInfo(vid, pid, softwareVersion, certificationType, setup.CertificationCenter)
	require.NoError(t, deleteComplInfoError)

	setup.CheckComplianceInfoDeleted(t, deleteComplInfoMsg)
}

func TestHandler_DeleteComplianceInfoForCertifiedModel(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := DeleteComplianceInfoSetup(t)

	_, certifyModelErr := setup.CertifyModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	require.NoError(t, certifyModelErr)

	deleteComplInfoMsg, deleteComplInfoError := setup.DeleteComplianceInfo(vid, pid, softwareVersion, certificationType, setup.CertificationCenter)
	require.NoError(t, deleteComplInfoError)

	setup.CheckComplianceInfoDeleted(t, deleteComplInfoMsg)
}

func TestHandler_DeleteComplianceInfoForProvisionalModel(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := DeleteComplianceInfoSetup(t)

	_, provisionModelErr := setup.ProvisionModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	require.NoError(t, provisionModelErr)

	deleteComplInfoMsg, deleteComplInfoError := setup.DeleteComplianceInfo(vid, pid, softwareVersion, certificationType, setup.CertificationCenter)
	require.NoError(t, deleteComplInfoError)

	setup.CheckComplianceInfoDeleted(t, deleteComplInfoMsg)
}

func TestHandler_DeleteComplianceInfoDoesNotExist(t *testing.T) {
	setup, vid, pid, softwareVersion, _, certificationType := DeleteComplianceInfoSetup(t)

	_, deleteComplInfoError := setup.DeleteComplianceInfo(vid, pid, softwareVersion, certificationType, setup.CertificationCenter)
	require.ErrorIs(t, deleteComplInfoError, types.ErrComplianceInfoDoesNotExist)
}
