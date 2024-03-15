package compliance

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	dclcompltypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/compliance"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func (setup *TestSetup) checkAllComplianceInfoFieldsUpdated(t *testing.T, originalComplianceInfo *dclcompltypes.ComplianceInfo, updatedComplianceInfo *dclcompltypes.ComplianceInfo) {
	require.Equal(t, originalComplianceInfo.Vid, updatedComplianceInfo.Vid)
	require.Equal(t, originalComplianceInfo.Pid, updatedComplianceInfo.Pid)
	require.Equal(t, originalComplianceInfo.SoftwareVersion, updatedComplianceInfo.SoftwareVersion)
	require.Equal(t, originalComplianceInfo.SoftwareVersionString, updatedComplianceInfo.SoftwareVersionString)
	require.Equal(t, originalComplianceInfo.CertificationType, updatedComplianceInfo.CertificationType)
	require.NotEqual(t, originalComplianceInfo.CertificationIdOfSoftwareComponent, updatedComplianceInfo.CertificationIdOfSoftwareComponent)
	require.NotEqual(t, originalComplianceInfo.CertificationRoute, updatedComplianceInfo.CertificationRoute)
	require.NotEqual(t, originalComplianceInfo.CompliantPlatformUsed, updatedComplianceInfo.CompliantPlatformUsed)
	require.NotEqual(t, originalComplianceInfo.CompliantPlatformVersion, updatedComplianceInfo.CompliantPlatformVersion)
	require.NotEqual(t, originalComplianceInfo.Date, updatedComplianceInfo.Date)
	require.NotEqual(t, originalComplianceInfo.FamilyId, updatedComplianceInfo.FamilyId)
	require.NotEqual(t, originalComplianceInfo.OSVersion, updatedComplianceInfo.OSVersion)
	require.NotEqual(t, originalComplianceInfo.ParentChild, updatedComplianceInfo.ParentChild)
	require.NotEqual(t, originalComplianceInfo.ProgramType, updatedComplianceInfo.ProgramType)
	require.NotEqual(t, originalComplianceInfo.ProgramTypeVersion, updatedComplianceInfo.ProgramTypeVersion)
	require.NotEqual(t, originalComplianceInfo.Reason, updatedComplianceInfo.Reason)
	require.NotEqual(t, originalComplianceInfo.SupportedClusters, updatedComplianceInfo.SupportedClusters)
	require.NotEqual(t, originalComplianceInfo.Transport, updatedComplianceInfo.Transport)
	require.Equal(t, originalComplianceInfo.SchemaVersion, updatedComplianceInfo.SchemaVersion)
}

func (setup *TestSetup) checkDeviceSoftwareComplianceUpdated(t *testing.T, originalComplianceInfo *dclcompltypes.ComplianceInfo, updatedDeviceSoftwareCompliance *types.DeviceSoftwareCompliance, isUpdatedMinimally bool) {
	isFound := false

	for _, complianceInfo := range updatedDeviceSoftwareCompliance.ComplianceInfo {
		if complianceInfo.Vid != originalComplianceInfo.Vid ||
			complianceInfo.Pid != originalComplianceInfo.Pid ||
			complianceInfo.SoftwareVersion != originalComplianceInfo.SoftwareVersion ||
			complianceInfo.CertificationType != originalComplianceInfo.CertificationType {
			continue
		}

		isFound = true

		if isUpdatedMinimally {
			require.NotEqual(t, complianceInfo.ParentChild, originalComplianceInfo.ParentChild)
		} else {
			setup.checkAllComplianceInfoFieldsUpdated(t, originalComplianceInfo, complianceInfo)
		}
	}

	// make sure the compliance info in device software compliance did not disappear
	require.True(t, isFound)
}

func setupUpdateComplianceInfo(t *testing.T) (*TestSetup, int32, int32, uint32, string, string, *dclcompltypes.ComplianceInfo, *types.DeviceSoftwareCompliance) {
	setup := setup(t)

	vid, pid, softwareVersion, softwareVersionString := setup.addModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)
	certificationType := dclcompltypes.ZigbeeCertificationType

	certifyModelMsg := newMsgCertifyModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	_, certifyModelErr := setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, certifyModelErr)

	originalComplianceInfo := queryExistingComplianceInfo(setup, vid, pid, softwareVersion, certificationType)
	originalDeviceSoftwareCompliance := queryExistingDeviceSoftwareCompliance(setup, certifyModelMsg.CDCertificateId)

	return setup, vid, pid, softwareVersion, certificationType, certifyModelMsg.CDCertificateId, originalComplianceInfo, originalDeviceSoftwareCompliance
}

func TestHandler_UpdateComplianceInfo_MinimalOnlyOneParentChildField(t *testing.T) {
	setup, vid, pid, softwareVersion, certificationType, cDCertificateID, originalComplianceInfo, _ := setupUpdateComplianceInfo(t)
	parentChild := testconstants.ParentChild1
	isUpdatedMinimally := true

	updateComplianceInfoMsg := newMsgUpdateComplianceInfo(setup.CertificationCenter, vid, pid, softwareVersion, certificationType)
	updateComplianceInfoMsg.ParentChild = parentChild
	_, updateComplianceInfoErr := setup.Handler(setup.Ctx, updateComplianceInfoMsg)
	require.NoError(t, updateComplianceInfoErr)

	updatedComplianceInfo := queryExistingComplianceInfo(setup, vid, pid, softwareVersion, certificationType)
	updatedDeviceSoftwareCompliance := queryExistingDeviceSoftwareCompliance(setup, cDCertificateID)

	require.NotEqual(t, originalComplianceInfo.ParentChild, updatedComplianceInfo.ParentChild)
	setup.checkDeviceSoftwareComplianceUpdated(t, originalComplianceInfo, updatedDeviceSoftwareCompliance, isUpdatedMinimally)
}

func TestHandler_UpdateComplianceInfo_AllFields(t *testing.T) {
	setup, vid, pid, softwareVersion, certificationType, cDCertificateID, originalComplianceInfo, _ := setupUpdateComplianceInfo(t)
	isUpdatedMinimally := false

	updateComplianceInfoMsg := newMsgUpdateComplianceInfoWithAllOptionalFlags(setup.CertificationCenter, vid, pid, softwareVersion, certificationType)
	_, updateComplianceInfoErr := setup.Handler(setup.Ctx, updateComplianceInfoMsg)
	require.NoError(t, updateComplianceInfoErr)

	updatedComplianceInfo := queryExistingComplianceInfo(setup, vid, pid, softwareVersion, certificationType)
	updatedDeviceSoftwareCompliance := queryExistingDeviceSoftwareCompliance(setup, cDCertificateID)

	setup.checkAllComplianceInfoFieldsUpdated(t, originalComplianceInfo, updatedComplianceInfo)
	setup.checkDeviceSoftwareComplianceUpdated(t, originalComplianceInfo, updatedDeviceSoftwareCompliance, isUpdatedMinimally)
}

func TestHandler_UpdateComplianceInfo_CDCertificateIdChanged(t *testing.T) {
	setup, vid, pid, softwareVersion, certificationType, cDCertificateID, _, originalDeviceSoftwareCompliance := setupUpdateComplianceInfo(t)
	newCDCertificateID := "15DEXD"

	updateComplianceInfoMsg := newMsgUpdateComplianceInfo(setup.CertificationCenter, vid, pid, softwareVersion, certificationType)
	updateComplianceInfoMsg.CDCertificateId = newCDCertificateID
	_, updateComplianceInfoErr := setup.Handler(setup.Ctx, updateComplianceInfoMsg)
	require.NoError(t, updateComplianceInfoErr)

	updatedComplianceInfo := queryExistingComplianceInfo(setup, vid, pid, softwareVersion, certificationType)
	updatedDeviceSoftwareCompliance := queryExistingDeviceSoftwareCompliance(setup, updateComplianceInfoMsg.CDCertificateId)

	_, err := queryDeviceSoftwareCompliance(setup, cDCertificateID)
	assertNotFound(t, err)

	require.Equal(t, len(updatedDeviceSoftwareCompliance.ComplianceInfo), len(originalDeviceSoftwareCompliance.ComplianceInfo))
	require.Equal(t, updatedDeviceSoftwareCompliance.CDCertificateId, updateComplianceInfoMsg.CDCertificateId)
	require.Equal(t, newCDCertificateID, updatedComplianceInfo.CDCertificateId)
}

func TestHandler_UpdateComplianceInfo_ByCertificationCenterNoOptionalFields(t *testing.T) {
	setup, vid, pid, softwareVersion, certificationType, _, originalComplianceInfo, _ := setupUpdateComplianceInfo(t)

	updateComplianceInfoMsg := newMsgUpdateComplianceInfo(setup.CertificationCenter, vid, pid, softwareVersion, certificationType)
	_, updateComplianceInfoErr := setup.Handler(setup.Ctx, updateComplianceInfoMsg)
	require.NoError(t, updateComplianceInfoErr)

	updatedComplianceInfo := queryExistingComplianceInfo(setup, vid, pid, softwareVersion, certificationType)
	require.Equal(t, originalComplianceInfo, updatedComplianceInfo)
}

func TestHandler_UpdateComplianceInfo_NotByCertificationCenter(t *testing.T) {
	setup, vid, pid, softwareVersion, certificationType, _, originalComplianceInfo, _ := setupUpdateComplianceInfo(t)

	nonCertCenterAddress := generateAccAddress()
	setup.DclauthKeeper.On("HasRole", mock.Anything, nonCertCenterAddress, dclauthtypes.CertificationCenter).Return(false)

	updateComplianceInfoMsg := newMsgUpdateComplianceInfoWithAllOptionalFlags(nonCertCenterAddress, vid, pid, softwareVersion, certificationType)
	_, updateComplianceInfoErr := setup.Handler(setup.Ctx, updateComplianceInfoMsg)
	require.ErrorIs(t, updateComplianceInfoErr, sdkerrors.ErrUnauthorized)

	updatedComplianceInfo := queryExistingComplianceInfo(setup, vid, pid, softwareVersion, dclcompltypes.ZigbeeCertificationType)

	require.Equal(t, originalComplianceInfo, updatedComplianceInfo)
}

func TestHandler_CDCertificateIDUpdateChangesOnlyOneComplianceInfo(t *testing.T) {
	setup, vid1, pid1, softwareVersion1, certificationType, _, originalComplianceInfo, _ := setupUpdateComplianceInfo(t)
	softwareVersionString1 := originalComplianceInfo.SoftwareVersionString
	newCDCertificateID := originalComplianceInfo.CDCertificateId + "new"

	vid2, pid2, softwareVersion2, softwareVersionString2 := setup.addModelVersion(vid1+1, pid1+1, softwareVersion1+1, softwareVersionString1)
	certifyModelMsg := newMsgCertifyModel(vid2, pid2, softwareVersion2, softwareVersionString2, certificationType, setup.CertificationCenter)
	_, certifyModelErr := setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, certifyModelErr)

	updateComplianceInfoMsg := newMsgUpdateComplianceInfo(setup.CertificationCenter, vid1, pid1, softwareVersion1, certificationType)
	updateComplianceInfoMsg.CDCertificateId = newCDCertificateID
	_, updateComplianceInfoErr := setup.Handler(setup.Ctx, updateComplianceInfoMsg)
	require.NoError(t, updateComplianceInfoErr)

	firstComplianceInfo, err := queryComplianceInfo(setup, vid1, pid1, softwareVersion1, testconstants.CertificationType)
	require.NoError(t, err)

	secondComplianceInfo, err := queryComplianceInfo(setup, vid2, pid2, softwareVersion2, testconstants.CertificationType)
	require.NoError(t, err)

	require.Equal(t, updateComplianceInfoMsg.CDCertificateId, firstComplianceInfo.CDCertificateId)
	require.NotEqual(t, firstComplianceInfo.CDCertificateId, secondComplianceInfo.CDCertificateId)
}

func TestHandler_UpdateToAnotherCDCertificateID(t *testing.T) {
	setup, vid1, pid1, softwareVersion1, certificationType, _, originalComplianceInfo, _ := setupUpdateComplianceInfo(t)
	softwareVersionString1 := originalComplianceInfo.SoftwareVersionString
	cDCertificateID1 := originalComplianceInfo.CDCertificateId

	vid2, pid2, softwareVersion2, softwareVersionString2 := setup.addModelVersion(vid1+1, pid1+1, softwareVersion1+1, softwareVersionString1)
	cDCertificateID2 := cDCertificateID1 + "new"
	certifyModelMsg := newMsgCertifyModel(vid2, pid2, softwareVersion2, softwareVersionString2, certificationType, setup.CertificationCenter)
	certifyModelMsg.CDCertificateId = cDCertificateID2
	_, certifyModelErr := setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, certifyModelErr)

	originalDeviceSoftwareCompliance2, _ := queryDeviceSoftwareCompliance(setup, cDCertificateID2)

	updateComplianceInfoMsg := newMsgUpdateComplianceInfo(setup.CertificationCenter, vid2, pid2, softwareVersion2, certificationType)
	updateComplianceInfoMsg.CDCertificateId = cDCertificateID1
	_, updateComplianceInfoErr := setup.Handler(setup.Ctx, updateComplianceInfoMsg)
	require.NoError(t, updateComplianceInfoErr)

	newDeviceSoftwareCompliance1, _ := queryDeviceSoftwareCompliance(setup, cDCertificateID1)
	newDeviceSoftwareCompliance2, _ := queryDeviceSoftwareCompliance(setup, cDCertificateID2)
	complianceInfo1, _ := queryComplianceInfo(setup, vid1, pid1, softwareVersion1, testconstants.CertificationType)
	complianceInfo2, _ := queryComplianceInfo(setup, vid2, pid2, softwareVersion2, testconstants.CertificationType)

	require.Equal(t, complianceInfo1.CDCertificateId, complianceInfo2.CDCertificateId)
	require.Equal(t, cDCertificateID1, newDeviceSoftwareCompliance1.ComplianceInfo[1].CDCertificateId)
	require.Nil(t, newDeviceSoftwareCompliance2)
	require.Equal(t, 2, len(newDeviceSoftwareCompliance1.ComplianceInfo))
	cdCertificateIDExcluded := originalDeviceSoftwareCompliance2.ComplianceInfo[0]
	cdCertificateIDExcluded.CDCertificateId = cDCertificateID1
	require.Equal(t, cdCertificateIDExcluded, newDeviceSoftwareCompliance1.ComplianceInfo[1])
}
