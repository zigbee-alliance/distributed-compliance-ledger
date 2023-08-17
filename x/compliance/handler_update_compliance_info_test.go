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

func (setup *TestSetup) CertifyModel(t *testing.T, vid int32, pid int32, softwareVersion uint32, softwareVersionString string, certificationType string, signer sdk.AccAddress) *types.MsgCertifyModel {
	certifyModelMsg := NewMsgCertifyModel(
		vid, pid, softwareVersion, softwareVersionString, certificationType, signer)
	_, err := setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, err)

	return certifyModelMsg
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
	require.Equal(t, true, isFound)
}

func UpdateComplianceInfoSetup(t *testing.T) (*TestSetup, int32, int32, uint32, string, string, string, *dclcompltypes.ComplianceInfo, *types.DeviceSoftwareCompliance) {
	setup := Setup(t)

	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)
	certificationType := dclcompltypes.ZigbeeCertificationType

	certifyModelMsg := setup.CertifyModel(t, vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
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
