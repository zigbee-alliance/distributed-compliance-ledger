package compliance

import (
	"fmt"
	"testing"
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	modeltypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

func setupCertifyModel(t *testing.T) (*TestSetup, int32, int32, uint32, string, string) {
	t.Helper()
	setup := setup(t)

	vid, pid, softwareVersion, softwareVersionString := setup.addModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)
	certificationType := types.ZigbeeCertificationType

	return setup, vid, pid, softwareVersion, softwareVersionString, certificationType
}

func (setup *TestSetup) checkModelCertified(t *testing.T, certifyModelMsg *types.MsgCertifyModel) {
	t.Helper()

	vid := certifyModelMsg.Vid
	pid := certifyModelMsg.Pid
	softwareVersion := certifyModelMsg.SoftwareVersion
	certificationType := certifyModelMsg.CertificationType

	certifiedModel, _ := queryCertifiedModel(setup, vid, pid, softwareVersion, certificationType)
	require.True(t, certifiedModel.Value)

	revokedModel, _ := queryRevokedModel(setup, vid, pid, softwareVersion, certificationType)
	require.False(t, revokedModel.Value)

	provisionalModel, _ := queryProvisionalModel(setup, vid, pid, softwareVersion, certificationType)
	require.False(t, provisionalModel.Value)
}

func (setup *TestSetup) checkCertifiedModelMatchesMsg(t *testing.T, certifyModelMsg *types.MsgCertifyModel) *types.ComplianceInfo {
	t.Helper()

	vid := certifyModelMsg.Vid
	pid := certifyModelMsg.Pid
	softwareVersion := certifyModelMsg.SoftwareVersion
	certificationType := certifyModelMsg.CertificationType

	receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, certificationType)
	checkCertifiedModelInfo(t, certifyModelMsg, receivedComplianceInfo)

	return receivedComplianceInfo
}

func (setup *TestSetup) checkDeviceSoftwareComplianceMatchesComplianceInfo(t *testing.T, complianceInfo *types.ComplianceInfo) {
	t.Helper()
	receivedDeviceSoftwareCompliance, _ := queryDeviceSoftwareCompliance(setup, complianceInfo.CDCertificateId)
	require.Equal(t, receivedDeviceSoftwareCompliance.CDCertificateId, complianceInfo.CDCertificateId)
	checkDeviceSoftwareCompliance(t, receivedDeviceSoftwareCompliance.ComplianceInfo[0], complianceInfo)
}

func TestHandler_CertifyModel_Zigbee(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupCertifyModel(t)

	certifyModelMsg := newMsgCertifyModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	_, certifyModelErr := setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, certifyModelErr)

	complianceInfo := setup.checkCertifiedModelMatchesMsg(t, certifyModelMsg)
	setup.checkDeviceSoftwareComplianceMatchesComplianceInfo(t, complianceInfo)
	setup.checkModelCertified(t, certifyModelMsg)
}

func TestHandler_CertifyModel_Zigbee_WithAllOptionalFlags(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupCertifyModel(t)

	certifyModelMsg := newMsgCertifyModelWithAllOptionalFlags(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	_, err := setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, err)

	complianceInfo := setup.checkCertifiedModelMatchesMsg(t, certifyModelMsg)
	setup.checkDeviceSoftwareComplianceMatchesComplianceInfo(t, complianceInfo)
	setup.checkModelCertified(t, certifyModelMsg)
}

func TestHandler_CertifyModel_Matter(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, _ := setupCertifyModel(t)
	certificationType := types.MatterCertificationType

	certifyModelMsg := newMsgCertifyModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	_, certifyModelErr := setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, certifyModelErr)

	complianceInfo := setup.checkCertifiedModelMatchesMsg(t, certifyModelMsg)
	setup.checkDeviceSoftwareComplianceMatchesComplianceInfo(t, complianceInfo)
	setup.checkModelCertified(t, certifyModelMsg)
}

func TestHandler_CertifyModel_Matter_WithAllOptionalFlags(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, _ := setupCertifyModel(t)
	certificationType := types.MatterCertificationType

	certifyModelMsg := newMsgCertifyModelWithAllOptionalFlags(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	_, certifyModelErr := setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, certifyModelErr)

	complianceInfo := setup.checkCertifiedModelMatchesMsg(t, certifyModelMsg)
	setup.checkDeviceSoftwareComplianceMatchesComplianceInfo(t, complianceInfo)
	setup.checkModelCertified(t, certifyModelMsg)
}

func TestHandler_CertifyProvisionedModel(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupCertifyModel(t)

	provisionModelMsg, provisionModelErr := setup.provisionModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	require.NoError(t, provisionModelErr)

	certifyModelMsg := newMsgCertifyModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	_, certifyModelErr := setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, certifyModelErr)

	complianceInfo := setup.checkCertifiedModelMatchesMsg(t, certifyModelMsg)
	setup.checkModelCertified(t, certifyModelMsg)
	setup.checkDeviceSoftwareComplianceMatchesComplianceInfo(t, complianceInfo)
	require.Equal(t, 1, len(complianceInfo.History))
	require.Equal(t, types.CodeProvisional, complianceInfo.History[0].SoftwareVersionCertificationStatus)
	require.Equal(t, provisionModelMsg.ProvisionalDate, complianceInfo.History[0].Date)
}

func TestHandler_CertifyProvisionedModelWithAllOptionalFlags(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupCertifyModel(t)

	provisionModelMsg, provisionModelErr := setup.provisionModelWithAllOptionalFlags(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	require.NoError(t, provisionModelErr)

	certifyModelMsg := newMsgCertifyModelWithAllOptionalFlags(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	_, certifyModelErr := setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, certifyModelErr)

	complianceInfo := setup.checkCertifiedModelMatchesMsg(t, certifyModelMsg)
	setup.checkModelCertified(t, certifyModelMsg)
	setup.checkDeviceSoftwareComplianceMatchesComplianceInfo(t, complianceInfo)
	require.Equal(t, 1, len(complianceInfo.History))
	require.Equal(t, types.CodeProvisional, complianceInfo.History[0].SoftwareVersionCertificationStatus)
	require.Equal(t, provisionModelMsg.ProvisionalDate, complianceInfo.History[0].Date)
}

func TestHandler_CertifyModelByDifferentRoles(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupCertifyModel(t)
	accountRoles := []dclauthtypes.AccountRole{
		dclauthtypes.Vendor,
		dclauthtypes.Trustee,
		dclauthtypes.NodeAdmin,
	}

	for _, role := range accountRoles {
		accAddress := generateAccAddress()
		setup.addAccount(accAddress, []dclauthtypes.AccountRole{role})

		certifyModelMsg := newMsgCertifyModel(vid, pid, softwareVersion, softwareVersionString, certificationType, accAddress)
		_, certifyModelErr := setup.Handler(setup.Ctx, certifyModelMsg)
		require.ErrorIs(t, certifyModelErr, sdkerrors.ErrUnauthorized)
	}
}

func TestHandler_CertifyModelForUnknownModel(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupCertifyModel(t)
	nonExistentPid := pid + 1

	setup.setNoModelVersionForKey(vid, nonExistentPid, softwareVersion)

	certifyModelMsg := newMsgCertifyModel(vid, nonExistentPid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	_, certifyModelErr := setup.Handler(setup.Ctx, certifyModelMsg)
	require.ErrorIs(t, certifyModelErr, modeltypes.ErrModelVersionDoesNotExist)
}

func TestHandler_CertifyModelWithWrongSoftwareVersionString(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupCertifyModel(t)
	nonExistentSoftwareVersionStr := softwareVersionString + "modified"

	certifyModelMsg := newMsgCertifyModel(vid, pid, softwareVersion, nonExistentSoftwareVersionStr, certificationType, setup.CertificationCenter)
	_, certifyModelErr := setup.Handler(setup.Ctx, certifyModelMsg)
	require.ErrorIs(t, certifyModelErr, types.ErrModelVersionStringDoesNotMatch)
}

func TestHandler_CertifyModelTwice(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupCertifyModel(t)

	certifyModelMsg := newMsgCertifyModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	_, certifyModelErr := setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, certifyModelErr)

	_ = newMsgCertifyModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	_, certifyModelErr = setup.Handler(setup.Ctx, certifyModelMsg)
	require.ErrorIs(t, certifyModelErr, types.ErrAlreadyCertified)
}

func TestHandler_CertifyModelTwiceByDifferentAccounts(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupCertifyModel(t)
	accAddress := generateAccAddress()
	setup.addAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.CertificationCenter})

	certifyModelMsg := newMsgCertifyModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	_, certifyModelErr := setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, certifyModelErr)

	_ = newMsgCertifyModel(vid, pid, softwareVersion, softwareVersionString, certificationType, accAddress)
	_, certifyModelErr = setup.Handler(setup.Ctx, certifyModelMsg)
	require.ErrorIs(t, certifyModelErr, types.ErrAlreadyCertified)
}

func TestHandler_CertifyDifferentModels(t *testing.T) {
	setup := setup(t)
	modelVersionsQuantity := 5

	for i := 1; i < modelVersionsQuantity; i++ {
		vid, pid, softwareVersion, softwareVersionString := setup.addModelVersion(int32(i), int32(i), uint32(i), fmt.Sprint(i))

		for _, certificationType := range setup.CertificationTypes {
			certifyModelMsg := newMsgCertifyModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
			_, certifyModelErr := setup.Handler(setup.Ctx, certifyModelMsg)
			require.NoError(t, certifyModelErr)
			setup.checkCertifiedModelMatchesMsg(t, certifyModelMsg)
			setup.checkModelCertified(t, certifyModelMsg)
		}
	}
}

func TestHandler_CertifyProvisionedModelForCertificationDateBeforeProvisionalDate(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupCertifyModel(t)

	certificationTime := time.Now().UTC()
	certificationDate := certificationTime.Format(time.RFC3339)
	provisionalDate := certificationTime.AddDate(0, 0, 1).Format(time.RFC3339)

	provisionModelErr := setup.provisionModelByDate(vid, pid, softwareVersion, softwareVersionString, certificationType, provisionalDate, setup.CertificationCenter)
	require.NoError(t, provisionModelErr)

	certifyModelMsg := newMsgCertifyModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	certifyModelMsg.CertificationDate = certificationDate
	_, certifyModelErr := setup.Handler(setup.Ctx, certifyModelMsg)
	require.ErrorIs(t, certifyModelErr, types.ErrInconsistentDates)
}

func TestHandler_CertifyRevokedModel(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupCertifyModel(t)
	certificationTime := time.Now().UTC()
	certificationDate := certificationTime.Format(time.RFC3339)

	revokeModelMsg, revokeModelErr := setup.revokeModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	require.NoError(t, revokeModelErr)

	certifyModelMsg := newMsgCertifyModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	certifyModelMsg.CertificationDate = certificationDate
	_, certifyModelErr := setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, certifyModelErr)

	receivedComplianceInfo := setup.checkCertifiedModelMatchesMsg(t, certifyModelMsg)
	require.Equal(t, 1, len(receivedComplianceInfo.History))
	require.Equal(t, types.CodeRevoked, receivedComplianceInfo.History[0].SoftwareVersionCertificationStatus)
	require.Equal(t, revokeModelMsg.RevocationDate, receivedComplianceInfo.History[0].Date)

	receivedDeviceSoftwareCompliance, _ := queryDeviceSoftwareCompliance(setup, testconstants.CDCertificateID)
	require.Equal(t, receivedDeviceSoftwareCompliance.CDCertificateId, testconstants.CDCertificateID)
	checkDeviceSoftwareCompliance(t, receivedDeviceSoftwareCompliance.ComplianceInfo[0], receivedComplianceInfo)

	setup.checkModelCertified(t, certifyModelMsg)
}
