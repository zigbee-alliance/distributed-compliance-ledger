package compliance

import (
	"context"
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	testkeeper "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	dclcompltypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/compliance"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	modeltypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DclauthKeeperMock struct {
	mock.Mock
}

func (m *DclauthKeeperMock) HasRole(
	ctx sdk.Context,
	addr sdk.AccAddress,
	roleToCheck dclauthtypes.AccountRole,
) bool {
	args := m.Called(ctx, addr, roleToCheck)

	return args.Bool(0)
}

var _ types.DclauthKeeper = &DclauthKeeperMock{}

type ModelKeeperMock struct {
	mock.Mock
}

func (m *ModelKeeperMock) GetModelVersion(
	ctx sdk.Context,
	vid int32,
	pid int32,
	softwareVersion uint32,
) (val modeltypes.ModelVersion, found bool) {
	args := m.Called(ctx, vid, pid, softwareVersion)

	return args.Get(0).(modeltypes.ModelVersion), args.Bool(1)
}

var _ types.ModelKeeper = &ModelKeeperMock{}

type TestSetup struct {
	T *testing.T
	// Cdc         *amino.Codec
	Ctx           sdk.Context
	Wctx          context.Context
	Keeper        *keeper.Keeper
	DclauthKeeper *DclauthKeeperMock
	ModelKeeper   *ModelKeeperMock
	Handler       sdk.Handler
	// Querier     sdk.Querier
	CertificationCenter sdk.AccAddress
	CertificationTypes  dclcompltypes.CertificationTypes
}

func (setup *TestSetup) addAccount(
	accAddress sdk.AccAddress,
	roles []dclauthtypes.AccountRole,
) {
	dclauthKeeper := setup.DclauthKeeper

	for _, role := range roles {
		dclauthKeeper.On("HasRole", mock.Anything, accAddress, role).Return(true)
	}
	dclauthKeeper.On("HasRole", mock.Anything, accAddress, mock.Anything).Return(false)
}

func (setup *TestSetup) addModelVersion(
	vid int32, pid int32, softwareVersion uint32, softwareVersionString string,
) (int32, int32, uint32, string) {
	modelVersion := newModelVersion(vid, pid, softwareVersion, softwareVersionString)

	setup.ModelKeeper.On(
		"GetModelVersion",
		mock.Anything, vid, pid, softwareVersion,
	).Return(*modelVersion, true)

	// return just for convenient re-assignment
	return vid, pid, softwareVersion, softwareVersionString
}

func (setup *TestSetup) setNoModelVersionForKey(
	vid int32,
	pid int32,
	softwareVersion uint32,
) {
	setup.ModelKeeper.On(
		"GetModelVersion",
		mock.Anything, vid, pid, softwareVersion,
	).Return(modeltypes.ModelVersion{}, false)
}

func setup(t *testing.T) *TestSetup {
	t.Helper()
	dclauthKeeper := &DclauthKeeperMock{}
	modelKeeper := &ModelKeeperMock{}
	keeper, ctx := testkeeper.ComplianceKeeper(t, dclauthKeeper, modelKeeper)

	certificationCenter := generateAccAddress()

	certificationTypes := dclcompltypes.CertificationTypes{dclcompltypes.ZigbeeCertificationType, dclcompltypes.MatterCertificationType}

	setup := &TestSetup{
		T:                   t,
		Ctx:                 ctx,
		Wctx:                sdk.WrapSDKContext(ctx),
		Keeper:              keeper,
		DclauthKeeper:       dclauthKeeper,
		ModelKeeper:         modelKeeper,
		Handler:             NewHandler(*keeper),
		CertificationCenter: certificationCenter,
		CertificationTypes:  certificationTypes,
	}

	setup.addAccount(certificationCenter, []dclauthtypes.AccountRole{dclauthtypes.CertificationCenter})

	return setup
}

func queryComplianceInfo(
	setup *TestSetup,
	vid int32,
	pid int32,
	softwareVersion uint32,
	certificationType string,
) (*dclcompltypes.ComplianceInfo, error) {
	req := &types.QueryGetComplianceInfoRequest{
		Vid:               vid,
		Pid:               pid,
		SoftwareVersion:   softwareVersion,
		CertificationType: certificationType,
	}

	resp, err := setup.Keeper.ComplianceInfo(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.ComplianceInfo, nil
}

func queryExistingComplianceInfo(
	setup *TestSetup,
	vid int32,
	pid int32,
	softwareVersion uint32,
	certificationType string,
) *dclcompltypes.ComplianceInfo {
	complianceInfo, err := queryComplianceInfo(setup, vid, pid, softwareVersion, certificationType)
	require.NoError(setup.T, err)

	return complianceInfo
}

func queryDeviceSoftwareCompliance(
	setup *TestSetup,
	cDCertificateID string,
) (*types.DeviceSoftwareCompliance, error) {
	req := &types.QueryGetDeviceSoftwareComplianceRequest{
		CDCertificateId: cDCertificateID,
	}

	resp, err := setup.Keeper.DeviceSoftwareCompliance(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.DeviceSoftwareCompliance, nil
}

func queryExistingDeviceSoftwareCompliance(
	setup *TestSetup,
	cDCertificateID string,
) *types.DeviceSoftwareCompliance {
	deviceSoftwareCompliance, err := queryDeviceSoftwareCompliance(setup, cDCertificateID)
	require.NoError(setup.T, err)

	return deviceSoftwareCompliance
}

func queryProvisionalModel(
	setup *TestSetup,
	vid int32,
	pid int32,
	softwareVersion uint32,
	certificationType string,
) (*types.ProvisionalModel, error) {
	req := &types.QueryGetProvisionalModelRequest{
		Vid:               vid,
		Pid:               pid,
		SoftwareVersion:   softwareVersion,
		CertificationType: certificationType,
	}

	resp, err := setup.Keeper.ProvisionalModel(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.ProvisionalModel, nil
}

func queryCertifiedModel(
	setup *TestSetup,
	vid int32,
	pid int32,
	softwareVersion uint32,
	certificationType string,
) (*types.CertifiedModel, error) {
	req := &types.QueryGetCertifiedModelRequest{
		Vid:               vid,
		Pid:               pid,
		SoftwareVersion:   softwareVersion,
		CertificationType: certificationType,
	}

	resp, err := setup.Keeper.CertifiedModel(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.CertifiedModel, nil
}

func queryRevokedModel(
	setup *TestSetup,
	vid int32,
	pid int32,
	softwareVersion uint32,
	certificationType string,
) (*types.RevokedModel, error) {
	req := &types.QueryGetRevokedModelRequest{
		Vid:               vid,
		Pid:               pid,
		SoftwareVersion:   softwareVersion,
		CertificationType: certificationType,
	}

	resp, err := setup.Keeper.RevokedModel(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.RevokedModel, nil
}

func checkProvisionalModelInfo(
	t *testing.T,
	provisionalModelMsg *types.MsgProvisionModel,
	receivedComplianceInfo *dclcompltypes.ComplianceInfo,
) {
	t.Helper()
	require.Equal(t, provisionalModelMsg.Vid, receivedComplianceInfo.Vid)
	require.Equal(t, provisionalModelMsg.Pid, receivedComplianceInfo.Pid)
	require.Equal(t, dclcompltypes.CodeProvisional, receivedComplianceInfo.SoftwareVersionCertificationStatus)
	require.Equal(t, provisionalModelMsg.ProvisionalDate, receivedComplianceInfo.Date)
	require.Equal(t, provisionalModelMsg.Reason, receivedComplianceInfo.Reason)
	require.Equal(t, provisionalModelMsg.CertificationType, receivedComplianceInfo.CertificationType)
	require.Equal(t, provisionalModelMsg.CDCertificateId, receivedComplianceInfo.CDCertificateId)
	require.Equal(t, provisionalModelMsg.ProgramTypeVersion, receivedComplianceInfo.ProgramTypeVersion)
	require.Equal(t, provisionalModelMsg.FamilyId, receivedComplianceInfo.FamilyId)
	require.Equal(t, provisionalModelMsg.SupportedClusters, receivedComplianceInfo.SupportedClusters)
	require.Equal(t, provisionalModelMsg.CompliantPlatformUsed, receivedComplianceInfo.CompliantPlatformUsed)
	require.Equal(t, provisionalModelMsg.CompliantPlatformVersion, receivedComplianceInfo.CompliantPlatformVersion)
	require.Equal(t, provisionalModelMsg.OSVersion, receivedComplianceInfo.OSVersion)
	require.Equal(t, provisionalModelMsg.CertificationRoute, receivedComplianceInfo.CertificationRoute)
	require.Equal(t, provisionalModelMsg.ProgramType, receivedComplianceInfo.ProgramType)
	require.Equal(t, provisionalModelMsg.Transport, receivedComplianceInfo.Transport)
	require.Equal(t, provisionalModelMsg.ParentChild, receivedComplianceInfo.ParentChild)
	require.Equal(t, provisionalModelMsg.CertificationIdOfSoftwareComponent, receivedComplianceInfo.CertificationIdOfSoftwareComponent)
}

func checkCertifiedModelInfo(
	t *testing.T,
	certifyModelMsg *types.MsgCertifyModel,
	receivedComplianceInfo *dclcompltypes.ComplianceInfo,
) {
	t.Helper()
	require.Equal(t, certifyModelMsg.Vid, receivedComplianceInfo.Vid)
	require.Equal(t, certifyModelMsg.Pid, receivedComplianceInfo.Pid)
	require.Equal(t, dclcompltypes.CodeCertified, receivedComplianceInfo.SoftwareVersionCertificationStatus)
	require.Equal(t, certifyModelMsg.CertificationDate, receivedComplianceInfo.Date)
	require.Equal(t, certifyModelMsg.Reason, receivedComplianceInfo.Reason)
	require.Equal(t, certifyModelMsg.CertificationType, receivedComplianceInfo.CertificationType)
	require.Equal(t, certifyModelMsg.CDCertificateId, receivedComplianceInfo.CDCertificateId)
	require.Equal(t, certifyModelMsg.ProgramTypeVersion, receivedComplianceInfo.ProgramTypeVersion)
	require.Equal(t, certifyModelMsg.FamilyId, receivedComplianceInfo.FamilyId)
	require.Equal(t, certifyModelMsg.SupportedClusters, receivedComplianceInfo.SupportedClusters)
	require.Equal(t, certifyModelMsg.CompliantPlatformUsed, receivedComplianceInfo.CompliantPlatformUsed)
	require.Equal(t, certifyModelMsg.CompliantPlatformVersion, receivedComplianceInfo.CompliantPlatformVersion)
	require.Equal(t, certifyModelMsg.OSVersion, receivedComplianceInfo.OSVersion)
	require.Equal(t, certifyModelMsg.CertificationRoute, receivedComplianceInfo.CertificationRoute)
	require.Equal(t, certifyModelMsg.ProgramType, receivedComplianceInfo.ProgramType)
	require.Equal(t, certifyModelMsg.Transport, receivedComplianceInfo.Transport)
	require.Equal(t, certifyModelMsg.ParentChild, receivedComplianceInfo.ParentChild)
	require.Equal(t, certifyModelMsg.CertificationIdOfSoftwareComponent, receivedComplianceInfo.CertificationIdOfSoftwareComponent)
}

func checkDeviceSoftwareCompliance(
	t *testing.T,
	info *dclcompltypes.ComplianceInfo,
	receivedComplianceInfo *dclcompltypes.ComplianceInfo,
) {
	t.Helper()
	require.Equal(t, info.Vid, receivedComplianceInfo.Vid)
	require.Equal(t, info.Pid, receivedComplianceInfo.Pid)
	require.Equal(t, info.SoftwareVersionCertificationStatus, receivedComplianceInfo.SoftwareVersionCertificationStatus)
	require.Equal(t, info.Date, receivedComplianceInfo.Date)
	require.Equal(t, info.Reason, receivedComplianceInfo.Reason)
	require.Equal(t, info.CertificationType, receivedComplianceInfo.CertificationType)
	require.Equal(t, info.CDCertificateId, receivedComplianceInfo.CDCertificateId)
	require.Equal(t, info.ProgramTypeVersion, receivedComplianceInfo.ProgramTypeVersion)
	require.Equal(t, info.FamilyId, receivedComplianceInfo.FamilyId)
	require.Equal(t, info.SupportedClusters, receivedComplianceInfo.SupportedClusters)
	require.Equal(t, info.CompliantPlatformUsed, receivedComplianceInfo.CompliantPlatformUsed)
	require.Equal(t, info.CompliantPlatformVersion, receivedComplianceInfo.CompliantPlatformVersion)
	require.Equal(t, info.OSVersion, receivedComplianceInfo.OSVersion)
	require.Equal(t, info.CertificationRoute, receivedComplianceInfo.CertificationRoute)
	require.Equal(t, info.ProgramType, receivedComplianceInfo.ProgramType)
	require.Equal(t, info.Transport, receivedComplianceInfo.Transport)
	require.Equal(t, info.ParentChild, receivedComplianceInfo.ParentChild)
	require.Equal(t, info.CertificationIdOfSoftwareComponent, receivedComplianceInfo.CertificationIdOfSoftwareComponent)
}

func checkRevokedModelInfo(
	t *testing.T,
	revokeModelMsg *types.MsgRevokeModel,
	receivedComplianceInfo *dclcompltypes.ComplianceInfo,
) {
	t.Helper()
	require.Equal(t, revokeModelMsg.Vid, receivedComplianceInfo.Vid)
	require.Equal(t, revokeModelMsg.Pid, receivedComplianceInfo.Pid)
	require.Equal(t, dclcompltypes.CodeRevoked, receivedComplianceInfo.SoftwareVersionCertificationStatus)
	require.Equal(t, revokeModelMsg.RevocationDate, receivedComplianceInfo.Date)
	require.Equal(t, revokeModelMsg.Reason, receivedComplianceInfo.Reason)
	require.Equal(t, revokeModelMsg.CertificationType, receivedComplianceInfo.CertificationType)
}

func assertNotFound(t *testing.T, err error) {
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
}

func newMsgProvisionModel(
	vid int32,
	pid int32,
	softwareVersion uint32,
	softwareVersionString string,
	certificationType string,
	signer sdk.AccAddress,
) *types.MsgProvisionModel {
	return &types.MsgProvisionModel{
		Signer:                signer.String(),
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       softwareVersion,
		SoftwareVersionString: softwareVersionString,
		CDVersionNumber:       uint32(testconstants.CdVersionNumber),
		ProvisionalDate:       testconstants.ProvisionalDate,
		CertificationType:     certificationType,
		Reason:                testconstants.Reason,
		CDCertificateId:       testconstants.CDCertificateID,
	}
}

func newMsgUpdateComplianceInfo(
	creator sdk.AccAddress,
	vid int32,
	pid int32,
	softwareVersion uint32,
	softwareVersionString string,
	certificationType string,
) *types.MsgUpdateComplianceInfo {
	return &types.MsgUpdateComplianceInfo{
		Creator:                            creator.String(),
		Vid:                                vid,
		Pid:                                pid,
		SoftwareVersion:                    softwareVersion,
		CertificationType:                  certificationType,
		CDVersionNumber:                    "",
		Date:                               "",
		Reason:                             "",
		Owner:                              "",
		CDCertificateId:                    "",
		CertificationRoute:                 "",
		ProgramType:                        "",
		ProgramTypeVersion:                 "",
		CompliantPlatformUsed:              "",
		CompliantPlatformVersion:           "",
		Transport:                          "",
		FamilyId:                           "",
		SupportedClusters:                  "",
		OSVersion:                          "",
		ParentChild:                        "",
		CertificationIdOfSoftwareComponent: "",
	}
}

func newMsgUpdateComplianceInfoWithAllOptionalFlags(
	creator sdk.AccAddress,
	vid int32,
	pid int32,
	softwareVersion uint32,
	softwareVersionString string,
	certificationType string,
) *types.MsgUpdateComplianceInfo {
	return &types.MsgUpdateComplianceInfo{
		Creator:                            creator.String(),
		Vid:                                vid,
		Pid:                                pid,
		SoftwareVersion:                    softwareVersion,
		CertificationType:                  certificationType,
		CDVersionNumber:                    fmt.Sprint(testconstants.CdVersionNumber),
		Date:                               testconstants.ProvisionalDate,
		CDCertificateId:                    testconstants.CDCertificateID,
		Reason:                             "new Reason",
		CertificationRoute:                 "123",
		Owner:                              "new Owner",
		ProgramType:                        "new programType",
		ProgramTypeVersion:                 "new ProgramTypeVersion",
		CompliantPlatformUsed:              "new CompliantPlatformUsed",
		CompliantPlatformVersion:           "new CompliantPlatformVersion",
		Transport:                          "new Transport",
		FamilyId:                           "new FamilyId",
		SupportedClusters:                  "new SupportedClusters",
		OSVersion:                          "new OSVersion",
		ParentChild:                        "new ParentChild",
		CertificationIdOfSoftwareComponent: "new CertificationIdOfSoftwareComponent",
	}
}

func newMsgProvisionModelWithAllOptionalFlags(
	vid int32,
	pid int32,
	softwareVersion uint32,
	softwareVersionString string,
	certificationType string,
	signer sdk.AccAddress,
) *types.MsgProvisionModel {
	return &types.MsgProvisionModel{
		Signer:                             signer.String(),
		Vid:                                vid,
		Pid:                                pid,
		SoftwareVersion:                    softwareVersion,
		SoftwareVersionString:              softwareVersionString,
		CDVersionNumber:                    uint32(testconstants.CdVersionNumber),
		ProvisionalDate:                    testconstants.ProvisionalDate,
		CertificationType:                  certificationType,
		Reason:                             testconstants.Reason,
		CDCertificateId:                    testconstants.CDCertificateID,
		ProgramTypeVersion:                 testconstants.ProgramTypeVersion,
		FamilyId:                           testconstants.FamilyID,
		SupportedClusters:                  testconstants.SupportedClusters,
		CompliantPlatformUsed:              testconstants.CompliantPlatformUsed,
		CompliantPlatformVersion:           testconstants.CompliantPlatformVersion,
		OSVersion:                          testconstants.OSVersion,
		CertificationRoute:                 testconstants.CertificationRoute,
		ProgramType:                        testconstants.ProgramType,
		Transport:                          testconstants.Transport,
		ParentChild:                        testconstants.ParentChild1,
		CertificationIdOfSoftwareComponent: testconstants.CertificationIDOfSoftwareComponent,
	}
}

func newMsgCertifyModel(
	vid int32,
	pid int32,
	softwareVersion uint32,
	softwareVersionString string,
	certificationType string,
	signer sdk.AccAddress,
) *types.MsgCertifyModel {
	return &types.MsgCertifyModel{
		Signer:                signer.String(),
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       softwareVersion,
		SoftwareVersionString: softwareVersionString,
		CDVersionNumber:       uint32(testconstants.CdVersionNumber),
		CertificationDate:     testconstants.CertificationDate,
		CertificationType:     certificationType,
		Reason:                testconstants.Reason,
		CDCertificateId:       testconstants.CDCertificateID,
	}
}

func newMsgDeleteComplianceInfo(
	vid int32,
	pid int32,
	softwareVersion uint32,
	certificationType string,
	signer sdk.AccAddress,
) *types.MsgDeleteComplianceInfo {
	return &types.MsgDeleteComplianceInfo{
		Creator:           signer.String(),
		Vid:               vid,
		Pid:               pid,
		SoftwareVersion:   softwareVersion,
		CertificationType: certificationType,
	}
}

func newMsgCertifyModelWithAllOptionalFlags(
	vid int32,
	pid int32,
	softwareVersion uint32,
	softwareVersionString string,
	certificationType string,
	signer sdk.AccAddress,
) *types.MsgCertifyModel {
	return &types.MsgCertifyModel{
		Signer:                             signer.String(),
		Vid:                                vid,
		Pid:                                pid,
		SoftwareVersion:                    softwareVersion,
		SoftwareVersionString:              softwareVersionString,
		CDVersionNumber:                    uint32(testconstants.CdVersionNumber),
		CertificationDate:                  testconstants.CertificationDate,
		CertificationType:                  certificationType,
		Reason:                             testconstants.Reason,
		CDCertificateId:                    testconstants.CDCertificateID,
		ProgramTypeVersion:                 testconstants.ProgramTypeVersion,
		FamilyId:                           testconstants.FamilyID,
		SupportedClusters:                  testconstants.SupportedClusters,
		CompliantPlatformUsed:              testconstants.CompliantPlatformUsed,
		CompliantPlatformVersion:           testconstants.CompliantPlatformVersion,
		OSVersion:                          testconstants.OSVersion,
		CertificationRoute:                 testconstants.CertificationRoute,
		ProgramType:                        testconstants.ProgramType,
		Transport:                          testconstants.Transport,
		ParentChild:                        testconstants.ParentChild1,
		CertificationIdOfSoftwareComponent: testconstants.CertificationIDOfSoftwareComponent,
	}
}

func newMsgRevokeModel(
	vid int32,
	pid int32,
	softwareVersion uint32,
	softwareVersionString string,
	certificationType string,
	signer sdk.AccAddress,
) *types.MsgRevokeModel {
	return &types.MsgRevokeModel{
		Signer:                signer.String(),
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       softwareVersion,
		SoftwareVersionString: softwareVersionString,
		CDVersionNumber:       uint32(testconstants.CdVersionNumber),
		RevocationDate:        testconstants.RevocationDate,
		CertificationType:     certificationType,
		Reason:                testconstants.RevocationReason,
	}
}

func newModelVersion(
	vid int32,
	pid int32,
	softwareVersion uint32,
	softwareVersionString string,
) *modeltypes.ModelVersion {
	return &modeltypes.ModelVersion{
		Vid:                          vid,
		Pid:                          pid,
		SoftwareVersion:              softwareVersion,
		SoftwareVersionString:        softwareVersionString,
		CdVersionNumber:              testconstants.CdVersionNumber,
		FirmwareInformation:          testconstants.FirmwareInformation,
		SoftwareVersionValid:         testconstants.SoftwareVersionValid,
		OtaUrl:                       testconstants.OtaURL,
		OtaFileSize:                  testconstants.OtaFileSize,
		OtaChecksum:                  testconstants.OtaChecksum,
		OtaChecksumType:              testconstants.OtaChecksumType,
		MinApplicableSoftwareVersion: testconstants.MinApplicableSoftwareVersion,
		MaxApplicableSoftwareVersion: testconstants.MaxApplicableSoftwareVersion,
		ReleaseNotesUrl:              testconstants.ReleaseNotesURL,
		Creator:                      generateAccAddress().String(),
	}
}

func generateAccAddress() sdk.AccAddress {
	_, _, accAddress := testdata.KeyTestPubAddr()

	return accAddress
}
