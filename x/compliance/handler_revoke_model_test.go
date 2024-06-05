package compliance

import (
	"fmt"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	modeltypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

func setupRevokeModel(t *testing.T) (*TestSetup, int32, int32, uint32, string, string) {
	t.Helper()
	setup := setup(t)

	vid, pid, softwareVersion, softwareVersionString := setup.addModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)
	certificationType := types.ZigbeeCertificationType

	return setup, vid, pid, softwareVersion, softwareVersionString, certificationType
}

func (setup *TestSetup) revokeModel(vid int32, pid int32, softwareVersion uint32, softwareVersionString string, certificationType string, signer sdk.AccAddress) (*types.MsgRevokeModel, error) {
	revokeModelMsg := newMsgRevokeModel(vid, pid, softwareVersion, softwareVersionString, certificationType, signer)
	_, err := setup.Handler(setup.Ctx, revokeModelMsg)

	return revokeModelMsg, err
}

func (setup *TestSetup) checkModelRevoked(t *testing.T, revokeModelMsg *types.MsgRevokeModel) {
	t.Helper()
	vid := revokeModelMsg.Vid
	pid := revokeModelMsg.Pid
	softwareVersion := revokeModelMsg.SoftwareVersion
	certificationType := revokeModelMsg.CertificationType

	revokedModel, _ := queryRevokedModel(setup, vid, pid, softwareVersion, certificationType)
	require.True(t, revokedModel.Value)

	_, err := queryCertifiedModel(setup, vid, pid, softwareVersion, certificationType)
	assertNotFound(t, err)

	_, err = queryProvisionalModel(setup, vid, pid, softwareVersion, certificationType)
	assertNotFound(t, err)
}

func (setup *TestSetup) checkModelStatusChangedToRevoked(t *testing.T, revokeModelMsg *types.MsgRevokeModel) {
	t.Helper()
	vid := revokeModelMsg.Vid
	pid := revokeModelMsg.Pid
	softwareVersion := revokeModelMsg.SoftwareVersion
	certificationType := revokeModelMsg.CertificationType

	revokedModel, _ := queryRevokedModel(setup, vid, pid, softwareVersion, certificationType)
	require.True(t, revokedModel.Value)

	provisionalModel, _ := queryProvisionalModel(setup, vid, pid, softwareVersion, certificationType)
	require.False(t, provisionalModel.Value)

	certifiedModel, _ := queryCertifiedModel(setup, vid, pid, softwareVersion, certificationType)
	require.False(t, certifiedModel.Value)
}

func (setup *TestSetup) checkRevokedModelInfoEqualsMessageData(t *testing.T, revokeModelMsg *types.MsgRevokeModel) *types.ComplianceInfo {
	t.Helper()
	vid := revokeModelMsg.Vid
	pid := revokeModelMsg.Pid
	softwareVersion := revokeModelMsg.SoftwareVersion
	certificationType := revokeModelMsg.CertificationType

	receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, certificationType)
	checkRevokedModelInfo(t, revokeModelMsg, receivedComplianceInfo)

	return receivedComplianceInfo
}

func TestHandler_RevokeModel(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupRevokeModel(t)

	revokeModelMsg, revokeModelErr := setup.revokeModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	require.NoError(t, revokeModelErr)

	setup.checkRevokedModelInfoEqualsMessageData(t, revokeModelMsg)
	setup.checkModelRevoked(t, revokeModelMsg)
}

func TestHandler_RevokeCertifiedModel(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupRevokeModel(t)

	certifyModelMsg := newMsgCertifyModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	_, certifyModelErr := setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, certifyModelErr)

	revokeModelMsg, revokeModelErr := setup.revokeModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	require.NoError(t, revokeModelErr)

	_, err := queryDeviceSoftwareCompliance(setup, certifyModelMsg.CDCertificateId)
	assertNotFound(t, err)

	complianceInfo := setup.checkRevokedModelInfoEqualsMessageData(t, revokeModelMsg)
	require.Equal(t, 1, len(complianceInfo.History))
	require.Equal(t, types.CodeCertified, complianceInfo.History[0].SoftwareVersionCertificationStatus)
	require.Equal(t, certifyModelMsg.CertificationDate, complianceInfo.History[0].Date)

	setup.checkModelStatusChangedToRevoked(t, revokeModelMsg)
}

func TestHandler_RevokeProvisionedModel(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupRevokeModel(t)

	provisionModelMsg, provisionModelErr := setup.provisionModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	require.NoError(t, provisionModelErr)

	revokeModelMsg, revokeModelErr := setup.revokeModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	require.NoError(t, revokeModelErr)

	complianceInfo := setup.checkRevokedModelInfoEqualsMessageData(t, revokeModelMsg)
	require.Equal(t, 1, len(complianceInfo.History))
	require.Equal(t, types.CodeProvisional, complianceInfo.History[0].SoftwareVersionCertificationStatus)
	require.Equal(t, provisionModelMsg.ProvisionalDate, complianceInfo.History[0].Date)

	setup.checkModelStatusChangedToRevoked(t, revokeModelMsg)
}

func TestHandler_RevokeModelByDifferentRoles(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupRevokeModel(t)
	accountRoles := []dclauthtypes.AccountRole{
		dclauthtypes.Vendor,
		dclauthtypes.Trustee,
		dclauthtypes.NodeAdmin,
	}

	for _, role := range accountRoles {
		accAddress := generateAccAddress()
		setup.addAccount(accAddress, []dclauthtypes.AccountRole{role})

		_, revokeModelErr := setup.revokeModel(vid, pid, softwareVersion, softwareVersionString, certificationType, accAddress)
		require.Error(t, revokeModelErr)
		require.True(t, sdkerrors.ErrUnauthorized.Is(revokeModelErr))
	}
}

func TestHandler_RevokeModelForUnknownModel(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupRevokeModel(t)
	nonExistentPid := pid + 1

	setup.setNoModelVersionForKey(vid, nonExistentPid, softwareVersion)

	_, revokeModelErr := setup.revokeModel(vid, nonExistentPid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	require.Error(t, revokeModelErr)
	require.True(t, modeltypes.ErrModelVersionDoesNotExist.Is(revokeModelErr))
}

func TestHandler_RevokeModelWithWrongSoftwareVersionString(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupRevokeModel(t)
	nonExistentSVS := softwareVersionString + "-modified"

	_, revokeModelErr := setup.revokeModel(vid, pid, softwareVersion, nonExistentSVS, certificationType, setup.CertificationCenter)
	require.Error(t, revokeModelErr)
	require.True(t, types.ErrModelVersionStringDoesNotMatch.Is(revokeModelErr))
}

func TestHandler_RevokeModelTwice(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupRevokeModel(t)

	_, revokeModelErr := setup.revokeModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	require.NoError(t, revokeModelErr)

	_, revokeModelErr = setup.revokeModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	require.ErrorIs(t, revokeModelErr, types.ErrAlreadyRevoked)
}

func TestHandler_RevokeDifferentModels(t *testing.T) {
	setup := setup(t)
	certificationType := types.ZigbeeCertificationType
	modelVersionsQuantity := 5

	for i := 1; i < modelVersionsQuantity; i++ {
		vid := int32(i)
		pid := int32(i)
		softwareVersion := uint32(i)
		softwareVersionString := fmt.Sprint(i)

		setup.addModelVersion(vid, pid, softwareVersion, softwareVersionString)

		revokeModelMsg, revokeModelErr := setup.revokeModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)

		require.NoError(t, revokeModelErr)
		setup.checkModelRevoked(t, revokeModelMsg)
		setup.checkRevokedModelInfoEqualsMessageData(t, revokeModelMsg)
	}
}

func TestHandler_RevokeCertifiedModelForRevocationDateBeforeCertificationDate(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupRevokeModel(t)

	revocationTime := time.Now().UTC()
	revocationDate := revocationTime.Format(time.RFC3339)
	certificationDate := revocationTime.AddDate(0, 0, 1).Format(time.RFC3339)

	certifyModelMsg := newMsgCertifyModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	certifyModelMsg.CertificationDate = certificationDate
	_, certifyModelErr := setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, certifyModelErr)

	revokeModelMsg := newMsgRevokeModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	revokeModelMsg.RevocationDate = revocationDate
	_, revokeModelErr := setup.Handler(setup.Ctx, revokeModelMsg)
	require.Error(t, revokeModelErr)
	require.True(t, types.ErrInconsistentDates.Is(revokeModelErr))
}

func TestHandler_RevokeProvisionedModelForRevocationDateBeforeProvisionalDate(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupRevokeModel(t)

	revocationTime := time.Now().UTC()
	revocationDate := revocationTime.Format(time.RFC3339)
	provisionalDate := revocationTime.AddDate(0, 0, 1).Format(time.RFC3339)

	provisionModelErr := setup.provisionModelByDate(vid, pid, softwareVersion, softwareVersionString, certificationType, provisionalDate, setup.CertificationCenter)
	require.NoError(t, provisionModelErr)

	revokeModelMsg := newMsgRevokeModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	revokeModelMsg.RevocationDate = revocationDate
	_, revokeModelErr := setup.Handler(setup.Ctx, revokeModelMsg)
	require.Error(t, revokeModelErr)
	require.True(t, types.ErrInconsistentDates.Is(revokeModelErr))
}

func TestHandler_CertifyRevokedModelForCertificationDateBeforeRevocationDate(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupRevokeModel(t)

	certificationTime := time.Now().UTC()
	certificationDate := certificationTime.Format(time.RFC3339)
	revocationDate := certificationTime.AddDate(0, 0, 1).Format(time.RFC3339)

	revokeModelMsg := newMsgRevokeModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	revokeModelMsg.RevocationDate = revocationDate
	_, revokeModelErr := setup.Handler(setup.Ctx, revokeModelMsg)
	require.NoError(t, revokeModelErr)

	certifyModelMsg := newMsgCertifyModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	certifyModelMsg.CertificationDate = certificationDate
	_, certifyModelErr := setup.Handler(setup.Ctx, certifyModelMsg)
	require.Error(t, certifyModelErr)
	require.True(t, types.ErrInconsistentDates.Is(certifyModelErr))
}

func TestHandler_CertifyRevokedModelThatWasCertifiedEarlier(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupRevokeModel(t)
	revocationDate := time.Now().UTC().Format(time.RFC3339)
	certificationDate := time.Now().UTC().Format(time.RFC3339)

	certifyModelMsg := newMsgCertifyModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	_, certifyModelErr := setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, certifyModelErr)

	revokeModelMsg := newMsgRevokeModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	revokeModelMsg.RevocationDate = revocationDate
	_, revokeModelErr := setup.Handler(setup.Ctx, revokeModelMsg)
	require.NoError(t, revokeModelErr)

	secondCertifyModelMsg := newMsgCertifyModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	secondCertifyModelMsg.CertificationDate = certificationDate
	_, certifyModelErr = setup.Handler(setup.Ctx, secondCertifyModelMsg)
	require.NoError(t, certifyModelErr)

	complianceInfo := setup.checkCertifiedModelMatchesMsg(t, secondCertifyModelMsg)
	setup.checkModelCertified(t, secondCertifyModelMsg)

	require.Equal(t, 2, len(complianceInfo.History))
	require.Equal(t, types.CodeCertified, complianceInfo.History[0].SoftwareVersionCertificationStatus)
	require.Equal(t, certifyModelMsg.CertificationDate, complianceInfo.History[0].Date)
	require.Equal(t, types.CodeRevoked, complianceInfo.History[1].SoftwareVersionCertificationStatus)
	require.Equal(t, revokeModelMsg.RevocationDate, complianceInfo.History[1].Date)
}
