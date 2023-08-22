package compliance

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	dclcompltypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/compliance"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (setup *TestSetup) CertifyModel(vid int32, pid int32, softwareVersion uint32, softwareVersionString string, certificationType string, signer sdk.AccAddress) (*types.MsgCertifyModel, error) {
	certifyModelMsg := NewMsgCertifyModel(
		vid, pid, softwareVersion, softwareVersionString, certificationType, signer)
	_, err := setup.Handler(setup.Ctx, certifyModelMsg)

	return certifyModelMsg, err
}

func (setup *TestSetup) CertifyModelCustomCDCertificateID(vid int32, pid int32, softwareVersion uint32, softwareVersionString string, certificationType string, cDCertificateID string, signer sdk.AccAddress) (*types.MsgCertifyModel, error) {
	certifyModelMsg := NewMsgCertifyModel(
		vid, pid, softwareVersion, softwareVersionString, certificationType, signer)
	certifyModelMsg.CDCertificateId = cDCertificateID
	_, err := setup.Handler(setup.Ctx, certifyModelMsg)

	return certifyModelMsg, err
}

func (setup *TestSetup) CertifyModelByDate(vid int32, pid int32, softwareVersion uint32, softwareVersionString string, certificationType string, certificationDate string, signer sdk.AccAddress) (*types.MsgCertifyModel, error) {
	certifyModelMsg := NewMsgCertifyModel(
		vid, pid, softwareVersion, softwareVersionString, certificationType, signer)
	certifyModelMsg.CertificationDate = certificationDate
	_, err := setup.Handler(setup.Ctx, certifyModelMsg)

	return certifyModelMsg, err
}

func (setup *TestSetup) RevokeModelByDate(vid int32, pid int32, softwareVersion uint32, softwareVersionString string, certificationType string, revocationDate string, signer sdk.AccAddress) (*types.MsgRevokeModel, error) {
	revokeModelMsg := NewMsgRevokeModel(
		vid, pid, softwareVersion, softwareVersionString, certificationType, signer)
	revokeModelMsg.RevocationDate = revocationDate
	_, err := setup.Handler(setup.Ctx, revokeModelMsg)

	return revokeModelMsg, err
}

func (setup *TestSetup) CertifyModelAllOptionalFlags(vid int32, pid int32, softwareVersion uint32, softwareVersionString string, certificationType string, signer sdk.AccAddress) (*types.MsgCertifyModel, error) {
	certifyModelMsg := NewMsgCertifyModelWithAllOptionalFlags(
		vid, pid, softwareVersion, softwareVersionString, certificationType, signer)
	_, err := setup.Handler(setup.Ctx, certifyModelMsg)

	return certifyModelMsg, err
}

func (setup *TestSetup) UpdateComplianceInfoOneOptionalField(vid int32, pid int32, softwareVersion uint32, softwareVersionString string, certificationType string, parentChild string, signer sdk.AccAddress) (*types.MsgUpdateComplianceInfo, error) {
	updateComplianceInfoMsg := NewMsgUpdateComplianceInfo(signer, vid, pid, softwareVersion, softwareVersionString, certificationType)
	updateComplianceInfoMsg.ParentChild = parentChild
	_, err := setup.Handler(setup.Ctx, updateComplianceInfoMsg)

	return updateComplianceInfoMsg, err
}

func (setup *TestSetup) UpdateComplianceInfoNoOptionalFields(vid int32, pid int32, softwareVersion uint32, softwareVersionString string, certificationType string, signer sdk.AccAddress) (*types.MsgUpdateComplianceInfo, error) {
	updateComplianceInfoMsg := NewMsgUpdateComplianceInfo(signer, vid, pid, softwareVersion, softwareVersionString, certificationType)
	_, err := setup.Handler(setup.Ctx, updateComplianceInfoMsg)

	return updateComplianceInfoMsg, err
}

func (setup *TestSetup) UpdateComplianceInfosCDCertificateID(vid int32, pid int32, softwareVersion uint32, softwareVersionString string, certificationType string, CDCertificateID string, signer sdk.AccAddress) (*types.MsgUpdateComplianceInfo, error) {
	updateComplianceInfoMsg := NewMsgUpdateComplianceInfo(signer, vid, pid, softwareVersion, softwareVersionString, certificationType)
	updateComplianceInfoMsg.CDCertificateId = CDCertificateID
	_, err := setup.Handler(setup.Ctx, updateComplianceInfoMsg)

	return updateComplianceInfoMsg, err
}

func (setup *TestSetup) UpdateComplianceInfoWithAllOptionalFlags(vid int32, pid int32, softwareVersion uint32, softwareVersionString string, certificationType string, signer sdk.AccAddress) (*types.MsgUpdateComplianceInfo, error) {
	updateComplianceInfoMsg := NewMsgUpdateComplianceInfoWithAllOptionalFlags(signer, vid, pid, softwareVersion, softwareVersionString, certificationType)
	_, err := setup.Handler(setup.Ctx, updateComplianceInfoMsg)

	return updateComplianceInfoMsg, err
}

func (setup *TestSetup) CheckAllComplianceInfoFieldsUpdated(t *testing.T, originalComplianceInfo *dclcompltypes.ComplianceInfo, updatedComplianceInfo *dclcompltypes.ComplianceInfo) {
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
	require.NotEqual(t, originalComplianceInfo.Owner, updatedComplianceInfo.Owner)
	require.NotEqual(t, originalComplianceInfo.ParentChild, updatedComplianceInfo.ParentChild)
	require.NotEqual(t, originalComplianceInfo.ProgramType, updatedComplianceInfo.ProgramType)
	require.NotEqual(t, originalComplianceInfo.ProgramTypeVersion, updatedComplianceInfo.ProgramTypeVersion)
	require.NotEqual(t, originalComplianceInfo.Reason, updatedComplianceInfo.Reason)
	require.Equal(t, originalComplianceInfo.SoftwareVersionCertificationStatus, updatedComplianceInfo.SoftwareVersionCertificationStatus) // TODO: maybe its needed to be updated too?
	require.NotEqual(t, originalComplianceInfo.SupportedClusters, updatedComplianceInfo.SupportedClusters)
	require.NotEqual(t, originalComplianceInfo.Transport, updatedComplianceInfo.Transport)
}

func (setup *TestSetup) CheckDeviceSoftwareComplianceUpdated(t *testing.T, originalComplianceInfo *dclcompltypes.ComplianceInfo, updatedDeviceSoftwareCompliance *types.DeviceSoftwareCompliance, isUpdatedMinimally bool) {
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
			setup.CheckAllComplianceInfoFieldsUpdated(t, originalComplianceInfo, complianceInfo)
		}
	}

	// make sure the compliance info in device software compliance did not disappear
	require.True(t, isFound)
}

func UpdateComplianceInfoSetup(t *testing.T) (*TestSetup, int32, int32, uint32, string, string, string, *dclcompltypes.ComplianceInfo, *types.DeviceSoftwareCompliance) {
	setup := Setup(t)

	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)
	certificationType := dclcompltypes.ZigbeeCertificationType

	certifyModelMsg, certifyModelErr := setup.CertifyModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	require.NoError(t, certifyModelErr)
	originalComplianceInfo := queryExistingComplianceInfo(setup, vid, pid, softwareVersion, certificationType)
	originalDeviceSoftwareCompliance := queryExistingDeviceSoftwareCompliance(setup, certifyModelMsg.CDCertificateId)

	return setup, vid, pid, softwareVersion, softwareVersionString, certificationType, certifyModelMsg.CDCertificateId, originalComplianceInfo, originalDeviceSoftwareCompliance
}

func TestHandler_UpdateComplianceInfo_Minimal(t *testing.T) {
	// only ParentChild field would be updated in this test
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType, cDCertificateID, originalComplianceInfo, _ := UpdateComplianceInfoSetup(t)
	parentChild := testconstants.ParentChild1
	isUpdatedMinimally := true

	_, updateComplianceInfoErr := setup.UpdateComplianceInfoOneOptionalField(vid, pid, softwareVersion, softwareVersionString, certificationType, parentChild, setup.CertificationCenter)
	require.NoError(t, updateComplianceInfoErr)

	updatedComplianceInfo := queryExistingComplianceInfo(setup, vid, pid, softwareVersion, certificationType)
	updatedDeviceSoftwareCompliance := queryExistingDeviceSoftwareCompliance(setup, cDCertificateID)

	require.NotEqual(t, originalComplianceInfo.ParentChild, updatedComplianceInfo.ParentChild)
	setup.CheckDeviceSoftwareComplianceUpdated(t, originalComplianceInfo, updatedDeviceSoftwareCompliance, isUpdatedMinimally)
}

func TestHandler_UpdateComplianceInfo_AllFields(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType, cDCertificateID, originalComplianceInfo, _ := UpdateComplianceInfoSetup(t)
	isUpdatedMinimally := false

	_, updateComplianceInfoErr := setup.UpdateComplianceInfoWithAllOptionalFlags(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	require.NoError(t, updateComplianceInfoErr)

	updatedComplianceInfo := queryExistingComplianceInfo(setup, vid, pid, softwareVersion, certificationType)
	updatedDeviceSoftwareCompliance := queryExistingDeviceSoftwareCompliance(setup, cDCertificateID)

	setup.CheckAllComplianceInfoFieldsUpdated(t, originalComplianceInfo, updatedComplianceInfo)
	setup.CheckDeviceSoftwareComplianceUpdated(t, originalComplianceInfo, updatedDeviceSoftwareCompliance, isUpdatedMinimally)
}

func TestHandler_UpdateComplianceInfo_CDCertificateIdChanged(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType, cDCertificateID, _, originalDeviceSoftwareCompliance := UpdateComplianceInfoSetup(t)
	newCDCertificateID := "15DEXD"

	updateComplianceInfoMsg, updateComplianceInfoErr := setup.UpdateComplianceInfosCDCertificateID(vid, pid, softwareVersion, softwareVersionString, certificationType, newCDCertificateID, setup.CertificationCenter)
	require.NoError(t, updateComplianceInfoErr)

	updatedComplianceInfo := queryExistingComplianceInfo(setup, vid, pid, softwareVersion, certificationType)
	updatedDeviceSoftwareCompliance := queryExistingDeviceSoftwareCompliance(setup, updateComplianceInfoMsg.CDCertificateId)

	_, err := queryDeviceSoftwareCompliance(setup, cDCertificateID)
	require.Equal(t, codes.NotFound, status.Code(err))

	require.Equal(t, len(updatedDeviceSoftwareCompliance.ComplianceInfo), len(originalDeviceSoftwareCompliance.ComplianceInfo))
	require.Equal(t, updatedDeviceSoftwareCompliance.CDCertificateId, updateComplianceInfoMsg.CDCertificateId)
	require.Equal(t, newCDCertificateID, updatedComplianceInfo.CDCertificateId)
}

func TestHandler_UpdateComplianceInfo_ByCertificationCenterNoOptionalFields(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType, _, originalComplianceInfo, _ := UpdateComplianceInfoSetup(t)

	_, updateComplianceInfoErr := setup.UpdateComplianceInfoNoOptionalFields(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	require.NoError(t, updateComplianceInfoErr)

	updatedComplianceInfo := queryExistingComplianceInfo(setup, vid, pid, softwareVersion, certificationType)
	require.Equal(t, originalComplianceInfo, updatedComplianceInfo)
}

func TestHandler_UpdateComplianceInfo_NotByCertificationCenter(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType, _, originalComplianceInfo, _ := UpdateComplianceInfoSetup(t)

	nonCertCenterAddress := GenerateAccAddress()
	setup.DclauthKeeper.On("HasRole", mock.Anything, nonCertCenterAddress, dclauthtypes.CertificationCenter).Return(false)

	_, updateComplianceInfoErr := setup.UpdateComplianceInfoWithAllOptionalFlags(vid, pid, softwareVersion, softwareVersionString, certificationType, nonCertCenterAddress)
	require.ErrorIs(t, updateComplianceInfoErr, sdkerrors.ErrUnauthorized)

	updatedComplianceInfo := queryExistingComplianceInfo(setup, vid, pid, softwareVersion, dclcompltypes.ZigbeeCertificationType)

	require.Equal(t, originalComplianceInfo, updatedComplianceInfo)
}

func TestHandler_CDCertificateIDUpdateChangesOnlyOneComplianceInfo(t *testing.T) {
	setup, vid1, pid1, softwareVersion1, softwareVersionString1, certificationType, _, originalComplianceInfo, _ := UpdateComplianceInfoSetup(t)
	newCDCertificateID := originalComplianceInfo.CDCertificateId + "new"

	vid2, pid2, softwareVersion2, softwareVersionString2 := setup.AddModelVersion(vid1+1, pid1+1, softwareVersion1+1, softwareVersionString1)
	_, certifyModelErr := setup.CertifyModel(vid2, pid2, softwareVersion2, softwareVersionString2, certificationType, setup.CertificationCenter)
	require.NoError(t, certifyModelErr)

	updateComplianceInfoMsg, updateComplianceInfoErr := setup.UpdateComplianceInfosCDCertificateID(vid1, pid1, softwareVersion1, softwareVersionString1, certificationType, newCDCertificateID, setup.CertificationCenter)
	require.NoError(t, updateComplianceInfoErr)

	firstComplianceInfo, err := queryComplianceInfo(setup, vid1, pid1, softwareVersion1, testconstants.CertificationType)
	require.NoError(t, err)

	secondComplianceInfo, err := queryComplianceInfo(setup, vid2, pid2, softwareVersion2, testconstants.CertificationType)
	require.NoError(t, err)

	require.Equal(t, updateComplianceInfoMsg.CDCertificateId, firstComplianceInfo.CDCertificateId)
	require.NotEqual(t, firstComplianceInfo.CDCertificateId, secondComplianceInfo.CDCertificateId)
}

func TestHandler_UpdateToAnotherCDCertificateID(t *testing.T) {
	setup, vid1, pid1, softwareVersion1, softwareVersionString1, certificationType, _, originalComplianceInfo, _ := UpdateComplianceInfoSetup(t)
	cDCertificateID1 := originalComplianceInfo.CDCertificateId

	vid2, pid2, softwareVersion2, softwareVersionString2 := setup.AddModelVersion(vid1+1, pid1+1, softwareVersion1+1, softwareVersionString1)
	cDCertificateID2 := cDCertificateID1 + "new"
	_, certifyModelErr := setup.CertifyModelCustomCDCertificateID(vid2, pid2, softwareVersion2, softwareVersionString2, certificationType, cDCertificateID2, setup.CertificationCenter)
	require.NoError(t, certifyModelErr)

	originalDeviceSoftwareCompliance2, _ := queryDeviceSoftwareCompliance(setup, cDCertificateID2)

	setup.UpdateComplianceInfosCDCertificateID(vid2, pid2, softwareVersion2, softwareVersionString2, certificationType, cDCertificateID1, setup.CertificationCenter)

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
