package compliance

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/rand"
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

func (setup *TestSetup) AddAccount(
	accAddress sdk.AccAddress,
	roles []dclauthtypes.AccountRole,
) {
	dclauthKeeper := setup.DclauthKeeper

	for _, role := range roles {
		dclauthKeeper.On("HasRole", mock.Anything, accAddress, role).Return(true)
	}
	dclauthKeeper.On("HasRole", mock.Anything, accAddress, mock.Anything).Return(false)
}

func (setup *TestSetup) AddModelVersion(
	vid int32, pid int32, softwareVersion uint32, softwareVersionString string,
) (int32, int32, uint32, string) {
	modelVersion := NewModelVersion(vid, pid, softwareVersion, softwareVersionString)

	setup.ModelKeeper.On(
		"GetModelVersion",
		mock.Anything, vid, pid, softwareVersion,
	).Return(*modelVersion, true)

	// return just for convenient re-assignment
	return vid, pid, softwareVersion, softwareVersionString
}

func (setup *TestSetup) SetNoModelVersionForKey(
	vid int32,
	pid int32,
	softwareVersion uint32,
) {
	setup.ModelKeeper.On(
		"GetModelVersion",
		mock.Anything, vid, pid, softwareVersion,
	).Return(modeltypes.ModelVersion{}, false)
}

func Setup(t *testing.T) *TestSetup {
	t.Helper()
	dclauthKeeper := &DclauthKeeperMock{}
	modelKeeper := &ModelKeeperMock{}
	keeper, ctx := testkeeper.ComplianceKeeper(t, dclauthKeeper, modelKeeper)

	certificationCenter := GenerateAccAddress()

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

	setup.AddAccount(certificationCenter, []dclauthtypes.AccountRole{dclauthtypes.CertificationCenter})

	return setup
}

func TestHandler_CertifyModel_Zigbee(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	// certify model
	certifyModelMsg := NewMsgCertifyModel(
		vid, pid, softwareVersion, softwareVersionString, dclcompltypes.ZigbeeCertificationType, setup.CertificationCenter)
	_, err := setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, err)

	// query certified model
	receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, dclcompltypes.ZigbeeCertificationType)
	checkCertifiedModelInfo(t, certifyModelMsg, receivedComplianceInfo)

	// query device software compliance
	receivedDeviceSoftwareCompliance, _ := queryDeviceSoftwareCompliance(setup, testconstants.CDCertificateID)
	require.Equal(t, receivedDeviceSoftwareCompliance.CDCertificateId, testconstants.CDCertificateID)
	checkDeviceSoftwareCompliance(t, receivedDeviceSoftwareCompliance.ComplianceInfo[0], receivedComplianceInfo)

	// check
	certifiedModel, _ := queryCertifiedModel(setup, vid, pid, softwareVersion, dclcompltypes.ZigbeeCertificationType)
	require.True(t, certifiedModel.Value)

	revokedModel, _ := queryRevokedModel(setup, vid, pid, softwareVersion, dclcompltypes.ZigbeeCertificationType)
	require.False(t, revokedModel.Value)

	provisionalModel, _ := queryProvisionalModel(setup, vid, pid, softwareVersion, dclcompltypes.ZigbeeCertificationType)
	require.False(t, provisionalModel.Value)
}

func TestHandler_CertifyModel_Zigbee_WithAllOptionalFlags(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	// certify model
	certifyModelMsg := NewMsgCertifyModelWithAllOptionalFlags(
		vid, pid, softwareVersion, softwareVersionString, dclcompltypes.ZigbeeCertificationType, setup.CertificationCenter)
	_, err := setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, err)

	// query certified model
	receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, dclcompltypes.ZigbeeCertificationType)
	checkCertifiedModelInfo(t, certifyModelMsg, receivedComplianceInfo)

	// query device software compliance
	receivedDeviceSoftwareCompliance, _ := queryDeviceSoftwareCompliance(setup, testconstants.CDCertificateID)
	require.Equal(t, receivedDeviceSoftwareCompliance.CDCertificateId, testconstants.CDCertificateID)
	checkDeviceSoftwareCompliance(t, receivedDeviceSoftwareCompliance.ComplianceInfo[0], receivedComplianceInfo)

	// check
	certifiedModel, _ := queryCertifiedModel(setup, vid, pid, softwareVersion, dclcompltypes.ZigbeeCertificationType)
	require.True(t, certifiedModel.Value)

	revokedModel, _ := queryRevokedModel(setup, vid, pid, softwareVersion, dclcompltypes.ZigbeeCertificationType)
	require.False(t, revokedModel.Value)

	provisionalModel, _ := queryProvisionalModel(setup, vid, pid, softwareVersion, dclcompltypes.ZigbeeCertificationType)
	require.False(t, provisionalModel.Value)
}

func TestHandler_CertifyModel_Matter(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	// certify model
	certifyModelMsg := NewMsgCertifyModel(
		vid, pid, softwareVersion, softwareVersionString, dclcompltypes.MatterCertificationType, setup.CertificationCenter)
	_, err := setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, err)

	// query certified model
	receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, dclcompltypes.MatterCertificationType)
	checkCertifiedModelInfo(t, certifyModelMsg, receivedComplianceInfo)

	// query device software compliance
	receivedDeviceSoftwareCompliance, _ := queryDeviceSoftwareCompliance(setup, testconstants.CDCertificateID)
	require.Equal(t, receivedDeviceSoftwareCompliance.CDCertificateId, testconstants.CDCertificateID)
	checkDeviceSoftwareCompliance(t, receivedDeviceSoftwareCompliance.ComplianceInfo[0], receivedComplianceInfo)

	// check
	certifiedModel, _ := queryCertifiedModel(setup, vid, pid, softwareVersion, dclcompltypes.MatterCertificationType)
	require.True(t, certifiedModel.Value)

	revokedModel, _ := queryRevokedModel(setup, vid, pid, softwareVersion, dclcompltypes.MatterCertificationType)
	require.False(t, revokedModel.Value)

	provisionalModel, _ := queryProvisionalModel(setup, vid, pid, softwareVersion, dclcompltypes.MatterCertificationType)
	require.False(t, provisionalModel.Value)
}

func TestHandler_CertifyModel_Matter_WithAllOptionalFlags(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	// certify model
	certifyModelMsg := NewMsgCertifyModelWithAllOptionalFlags(
		vid, pid, softwareVersion, softwareVersionString, dclcompltypes.MatterCertificationType, setup.CertificationCenter)
	_, err := setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, err)

	// query certified model
	receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, dclcompltypes.MatterCertificationType)
	checkCertifiedModelInfo(t, certifyModelMsg, receivedComplianceInfo)

	// query device software compliance
	receivedDeviceSoftwareCompliance, _ := queryDeviceSoftwareCompliance(setup, testconstants.CDCertificateID)
	require.Equal(t, receivedDeviceSoftwareCompliance.CDCertificateId, testconstants.CDCertificateID)
	checkDeviceSoftwareCompliance(t, receivedDeviceSoftwareCompliance.ComplianceInfo[0], receivedComplianceInfo)

	// check
	certifiedModel, _ := queryCertifiedModel(setup, vid, pid, softwareVersion, dclcompltypes.MatterCertificationType)
	require.True(t, certifiedModel.Value)

	revokedModel, _ := queryRevokedModel(setup, vid, pid, softwareVersion, dclcompltypes.MatterCertificationType)
	require.False(t, revokedModel.Value)

	provisionalModel, _ := queryProvisionalModel(setup, vid, pid, softwareVersion, dclcompltypes.MatterCertificationType)
	require.False(t, provisionalModel.Value)
}

func TestHandler_CertifyProvisionedModel(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	for index, certificationType := range setup.CertificationTypes {
		// provision model
		provisionModelMsg := NewMsgProvisionModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		_, err := setup.Handler(setup.Ctx, provisionModelMsg)
		require.NoError(t, err)

		// certify model
		certifyModelMsg := NewMsgCertifyModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		certifyModelMsg.CertificationDate = time.Now().UTC().Format(time.RFC3339)
		_, err = setup.Handler(setup.Ctx, certifyModelMsg)
		require.NoError(t, err)

		// query certified model info
		receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, certificationType)
		checkCertifiedModelInfo(t, certifyModelMsg, receivedComplianceInfo)
		require.Equal(t, 1, len(receivedComplianceInfo.History))
		require.Equal(t, dclcompltypes.CodeProvisional, receivedComplianceInfo.History[0].SoftwareVersionCertificationStatus)
		require.Equal(t, provisionModelMsg.ProvisionalDate, receivedComplianceInfo.History[0].Date)

		// query device software compliance
		receivedDeviceSoftwareCompliance, _ := queryDeviceSoftwareCompliance(setup, testconstants.CDCertificateID)
		require.Equal(t, receivedDeviceSoftwareCompliance.CDCertificateId, testconstants.CDCertificateID)
		checkDeviceSoftwareCompliance(t, receivedDeviceSoftwareCompliance.ComplianceInfo[index], receivedComplianceInfo)

		// query certified model
		certifiedModel, _ := queryCertifiedModel(setup, vid, pid, softwareVersion, certificationType)
		require.True(t, certifiedModel.Value)

		// query provisional model
		provisionalModel, _ := queryProvisionalModel(setup, vid, pid, softwareVersion, certificationType)
		require.False(t, provisionalModel.Value)

		// query revoked model
		revokedModel, _ := queryRevokedModel(setup, vid, pid, softwareVersion, certificationType)
		require.False(t, revokedModel.Value)
	}
}

func TestHandler_CertifyProvisionedModel_WithAllOptionalFields(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	for index, certificationType := range setup.CertificationTypes {
		// provision model
		provisionModelMsg := NewMsgProvisionModelWithAllOptionalFlags(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		_, err := setup.Handler(setup.Ctx, provisionModelMsg)
		require.NoError(t, err)

		maxLegthOfField := 64

		// certify model
		certifyModelMsg := NewMsgCertifyModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)

		certifyModelMsg.CertificationDate = time.Now().UTC().Format(time.RFC3339)
		certifyModelMsg.FamilyId = rand.Str(maxLegthOfField)
		certifyModelMsg.SupportedClusters = rand.Str(maxLegthOfField)
		certifyModelMsg.CompliantPlatformUsed = rand.Str(maxLegthOfField)
		certifyModelMsg.CompliantPlatformVersion = rand.Str(maxLegthOfField)
		certifyModelMsg.CertificationRoute = rand.Str(maxLegthOfField)
		certifyModelMsg.OSVersion = rand.Str(maxLegthOfField)
		certifyModelMsg.CertificationRoute = rand.Str(maxLegthOfField)
		certifyModelMsg.ProgramType = rand.Str(maxLegthOfField)
		certifyModelMsg.ProgramTypeVersion = rand.Str(maxLegthOfField)
		certifyModelMsg.Transport = rand.Str(maxLegthOfField)
		certifyModelMsg.ParentChild = testconstants.ParentChild2
		certifyModelMsg.CertificationIdOfSoftwareComponent = rand.Str(maxLegthOfField)

		_, err = setup.Handler(setup.Ctx, certifyModelMsg)
		require.NoError(t, err)

		// query certified model info
		receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, certificationType)
		checkCertifiedModelInfo(t, certifyModelMsg, receivedComplianceInfo)
		require.Equal(t, 1, len(receivedComplianceInfo.History))
		require.Equal(t, dclcompltypes.CodeProvisional, receivedComplianceInfo.History[0].SoftwareVersionCertificationStatus)
		require.Equal(t, provisionModelMsg.ProvisionalDate, receivedComplianceInfo.History[0].Date)

		// query device software compliance
		receivedDeviceSoftwareCompliance, _ := queryDeviceSoftwareCompliance(setup, testconstants.CDCertificateID)
		require.Equal(t, receivedDeviceSoftwareCompliance.CDCertificateId, testconstants.CDCertificateID)
		checkDeviceSoftwareCompliance(t, receivedDeviceSoftwareCompliance.ComplianceInfo[index], receivedComplianceInfo)

		// query certified model
		certifiedModel, _ := queryCertifiedModel(setup, vid, pid, softwareVersion, certificationType)
		require.True(t, certifiedModel.Value)

		// query provisional model
		provisionalModel, _ := queryProvisionalModel(setup, vid, pid, softwareVersion, certificationType)
		require.False(t, provisionalModel.Value)

		// query revoked model
		revokedModel, _ := queryRevokedModel(setup, vid, pid, softwareVersion, certificationType)
		require.False(t, revokedModel.Value)
	}
}

func TestHandler_CertifyProvisionedModelWithAllOptionalFlags(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	for index, certificationType := range setup.CertificationTypes {
		// provision model
		provisionModelMsg := NewMsgProvisionModelWithAllOptionalFlags(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		_, err := setup.Handler(setup.Ctx, provisionModelMsg)
		require.NoError(t, err)

		// certify model
		certifyModelMsg := NewMsgCertifyModelWithAllOptionalFlags(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		certifyModelMsg.CertificationDate = time.Now().UTC().Format(time.RFC3339)
		_, err = setup.Handler(setup.Ctx, certifyModelMsg)
		require.NoError(t, err)

		// query certified model info
		receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, certificationType)
		checkCertifiedModelInfo(t, certifyModelMsg, receivedComplianceInfo)
		require.Equal(t, 1, len(receivedComplianceInfo.History))
		require.Equal(t, dclcompltypes.CodeProvisional, receivedComplianceInfo.History[0].SoftwareVersionCertificationStatus)
		require.Equal(t, provisionModelMsg.ProvisionalDate, receivedComplianceInfo.History[0].Date)

		// query device software compliance
		receivedDeviceSoftwareCompliance, _ := queryDeviceSoftwareCompliance(setup, testconstants.CDCertificateID)
		require.Equal(t, receivedDeviceSoftwareCompliance.CDCertificateId, testconstants.CDCertificateID)
		checkDeviceSoftwareCompliance(t, receivedDeviceSoftwareCompliance.ComplianceInfo[index], receivedComplianceInfo)

		// query certified model
		certifiedModel, _ := queryCertifiedModel(setup, vid, pid, softwareVersion, certificationType)
		require.True(t, certifiedModel.Value)

		// query provisional model
		provisionalModel, _ := queryProvisionalModel(setup, vid, pid, softwareVersion, certificationType)
		require.False(t, provisionalModel.Value)

		// query revoked model
		revokedModel, _ := queryRevokedModel(setup, vid, pid, softwareVersion, certificationType)
		require.False(t, revokedModel.Value)
	}
}

func TestHandler_CertifyModelByDifferentRoles(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.Vendor,
		dclauthtypes.Trustee,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role})

		// try to certify model
		for _, certificationType := range setup.CertificationTypes {
			certifyModelMsg := NewMsgCertifyModel(
				vid, pid, softwareVersion, softwareVersionString, certificationType, accAddress)
			_, err := setup.Handler(setup.Ctx, certifyModelMsg)
			require.Error(t, err)
			require.True(t, sdkerrors.ErrUnauthorized.Is(err))
		}
	}
}

func TestHandler_CertifyModelForUnknownModel(t *testing.T) {
	setup := Setup(t)

	vid := testconstants.Vid
	pid := testconstants.Pid
	softwareVersion := testconstants.SoftwareVersion
	softwareVersionString := testconstants.SoftwareVersionString

	// set presence of model version
	setup.SetNoModelVersionForKey(vid, pid, softwareVersion)

	// try to certify model
	for _, certificationType := range setup.CertificationTypes {
		certifyModelMsg := NewMsgCertifyModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		_, err := setup.Handler(setup.Ctx, certifyModelMsg)
		require.Error(t, err)
		require.True(t, modeltypes.ErrModelVersionDoesNotExist.Is(err))
	}
}

func TestHandler_CertifyModelWithWrongSoftwareVersionString(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	// try to certify model
	for _, certificationType := range setup.CertificationTypes {
		certifyModelMsg := NewMsgCertifyModel(
			vid, pid, softwareVersion, softwareVersionString+"-modified", certificationType, setup.CertificationCenter)
		_, err := setup.Handler(setup.Ctx, certifyModelMsg)
		require.Error(t, err)
		require.True(t, types.ErrModelVersionStringDoesNotMatch.Is(err))
	}
}

func TestHandler_CertifyModelTwice(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	for _, certificationType := range setup.CertificationTypes {
		// certify model
		certifyModelMsg := NewMsgCertifyModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		_, err := setup.Handler(setup.Ctx, certifyModelMsg)
		require.NoError(t, err)

		// certify model second time
		secondCertifyModelMsg := NewMsgCertifyModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		secondCertifyModelMsg.CertificationDate = time.Now().UTC().Format(time.RFC3339)
		_, err = setup.Handler(setup.Ctx, secondCertifyModelMsg)
		require.Error(t, err)
		require.True(t, types.ErrAlreadyCertified.Is(err))
	}
}

func TestHandler_CertifyModelTwiceByDifferentAccounts(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	for _, certificationType := range setup.CertificationTypes {
		// certify model
		certifyModelMsg := NewMsgCertifyModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		_, err := setup.Handler(setup.Ctx, certifyModelMsg)
		require.NoError(t, err)

		// create another certification center account
		accAddress := GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.CertificationCenter})

		// try to certify model again from new account
		secondCertifyModelMsg := NewMsgCertifyModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, accAddress)
		_, err = setup.Handler(setup.Ctx, secondCertifyModelMsg)
		require.Error(t, err)
		require.True(t, types.ErrAlreadyCertified.Is(err))
	}
}

func TestHandler_CertifyDifferentModels(t *testing.T) {
	setup := Setup(t)

	for i := 1; i < 5; i++ {
		// add model version
		vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
			int32(i), int32(i), uint32(i), fmt.Sprint(i))

		for _, certificationType := range setup.CertificationTypes {
			// certify model
			certifyModelMsg := NewMsgCertifyModel(
				vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
			_, err := setup.Handler(setup.Ctx, certifyModelMsg)
			require.NoError(t, err)

			// query certified model
			receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, certificationType)

			// check
			checkCertifiedModelInfo(t, certifyModelMsg, receivedComplianceInfo)
		}
	}
}

func TestHandler_CertifyProvisionedModelForCertificationDateBeforeProvisionalDate(t *testing.T) {
	setup := Setup(t)

	certificationTime := time.Now().UTC()
	certificationDate := certificationTime.Format(time.RFC3339)
	provisionalDate := certificationTime.AddDate(0, 0, 1).Format(time.RFC3339)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	for _, certificationType := range setup.CertificationTypes {
		// provision model
		provisionModelMsg := NewMsgProvisionModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		provisionModelMsg.ProvisionalDate = provisionalDate
		_, err := setup.Handler(setup.Ctx, provisionModelMsg)
		require.NoError(t, err)

		// try to certify model
		certifyModelMsg := NewMsgCertifyModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		certifyModelMsg.CertificationDate = certificationDate
		_, err = setup.Handler(setup.Ctx, certifyModelMsg)
		require.Error(t, err)
		require.True(t, types.ErrInconsistentDates.Is(err))
	}
}

func TestHandler_RevokeModel(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	for _, certificationType := range setup.CertificationTypes {
		// revoke model
		revokeModelMsg := NewMsgRevokeModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		_, err := setup.Handler(setup.Ctx, revokeModelMsg)
		require.NoError(t, err)

		// query revoked model
		receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, certificationType)

		// check
		checkRevokedModelInfo(t, revokeModelMsg, receivedComplianceInfo)

		revokedModel, _ := queryRevokedModel(setup, vid, pid, softwareVersion, certificationType)
		require.True(t, revokedModel.Value)

		_, err = queryCertifiedModel(setup, vid, pid, softwareVersion, certificationType)
		require.Error(t, err)
		require.Equal(t, codes.NotFound, status.Code(err))

		_, err = queryProvisionalModel(setup, vid, pid, softwareVersion, certificationType)
		require.Error(t, err)
		require.Equal(t, codes.NotFound, status.Code(err))
	}
}

func TestHandler_RevokeCertifiedModel(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	for _, certificationType := range setup.CertificationTypes {
		// certify model
		certifyModelMsg := NewMsgCertifyModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		_, err := setup.Handler(setup.Ctx, certifyModelMsg)
		require.NoError(t, err)

		// query certified model info
		receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, certificationType)
		checkCertifiedModelInfo(t, certifyModelMsg, receivedComplianceInfo)

		// query device software compliance
		receivedDeviceSoftwareCompliance, _ := queryDeviceSoftwareCompliance(setup, testconstants.CDCertificateID)
		require.Equal(t, receivedDeviceSoftwareCompliance.CDCertificateId, testconstants.CDCertificateID)
		checkDeviceSoftwareCompliance(t, receivedDeviceSoftwareCompliance.ComplianceInfo[0], receivedComplianceInfo)

		// revoke model
		revokeModelMsg := NewMsgRevokeModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		revokeModelMsg.RevocationDate = time.Now().UTC().Format(time.RFC3339)
		_, err = setup.Handler(setup.Ctx, revokeModelMsg)
		require.NoError(t, err)

		// query revoked model info
		receivedComplianceInfo, _ = queryComplianceInfo(setup, vid, pid, softwareVersion, certificationType)
		checkRevokedModelInfo(t, revokeModelMsg, receivedComplianceInfo)
		require.Equal(t, 1, len(receivedComplianceInfo.History))
		require.Equal(t, dclcompltypes.CodeCertified, receivedComplianceInfo.History[0].SoftwareVersionCertificationStatus)
		require.Equal(t, certifyModelMsg.CertificationDate, receivedComplianceInfo.History[0].Date)

		// query device software compliance
		_, err = queryDeviceSoftwareCompliance(setup, testconstants.CDCertificateID)
		require.Error(t, err)
		require.Equal(t, codes.NotFound, status.Code(err))

		// query revoked model
		revokedModel, _ := queryRevokedModel(setup, vid, pid, softwareVersion, certificationType)
		require.True(t, revokedModel.Value)

		// query certified model
		certifiedModel, _ := queryCertifiedModel(setup, vid, pid, softwareVersion, certificationType)
		require.False(t, certifiedModel.Value)

		// query provisional model
		provisionalModel, _ := queryProvisionalModel(setup, vid, pid, softwareVersion, certificationType)
		require.False(t, provisionalModel.Value)
	}
}

func TestHandler_RevokeProvisionedModel(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	for _, certificationType := range setup.CertificationTypes {
		// provision model
		provisionModelMsg := NewMsgProvisionModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		_, err := setup.Handler(setup.Ctx, provisionModelMsg)
		require.NoError(t, err)

		// revoke model
		revokeModelMsg := NewMsgRevokeModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		revokeModelMsg.RevocationDate = time.Now().UTC().Format(time.RFC3339)
		_, err = setup.Handler(setup.Ctx, revokeModelMsg)
		require.NoError(t, err)

		// query revoked model info
		receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, certificationType)

		// check
		checkRevokedModelInfo(t, revokeModelMsg, receivedComplianceInfo)
		require.Equal(t, 1, len(receivedComplianceInfo.History))

		require.Equal(t, dclcompltypes.CodeProvisional, receivedComplianceInfo.History[0].SoftwareVersionCertificationStatus)
		require.Equal(t, provisionModelMsg.ProvisionalDate, receivedComplianceInfo.History[0].Date)

		// query revoked model
		revokedModel, _ := queryRevokedModel(setup, vid, pid, softwareVersion, certificationType)
		require.True(t, revokedModel.Value)

		// query provisional model
		provisionalModel, _ := queryProvisionalModel(setup, vid, pid, softwareVersion, certificationType)
		require.False(t, provisionalModel.Value)

		// query certified model
		certifiedModel, _ := queryCertifiedModel(setup, vid, pid, softwareVersion, certificationType)
		require.False(t, certifiedModel.Value)
	}
}

func TestHandler_RevokeModelByDifferentRoles(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.Vendor,
		dclauthtypes.Trustee,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role})

		// try to revoke model
		for _, certificationType := range setup.CertificationTypes {
			revokeModelMsg := NewMsgRevokeModel(
				vid, pid, softwareVersion, softwareVersionString, certificationType, accAddress)
			_, err := setup.Handler(setup.Ctx, revokeModelMsg)
			require.Error(t, err)
			require.True(t, sdkerrors.ErrUnauthorized.Is(err))
		}
	}
}

func TestHandler_RevokeModelForUnknownModel(t *testing.T) {
	setup := Setup(t)

	vid := testconstants.Vid
	pid := testconstants.Pid
	softwareVersion := testconstants.SoftwareVersion
	softwareVersionString := testconstants.SoftwareVersionString

	// set presence of model version
	setup.SetNoModelVersionForKey(vid, pid, softwareVersion)

	// try to revoke model
	for _, certificationType := range setup.CertificationTypes {
		revokeModelMsg := NewMsgRevokeModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		_, err := setup.Handler(setup.Ctx, revokeModelMsg)
		require.Error(t, err)
		require.True(t, modeltypes.ErrModelVersionDoesNotExist.Is(err))
	}
}

func TestHandler_RevokeModelWithWrongSoftwareVersionString(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	// try to revoke model
	for _, certificationType := range setup.CertificationTypes {
		revokeModelMsg := NewMsgRevokeModel(
			vid, pid, softwareVersion, softwareVersionString+"-modified", certificationType, setup.CertificationCenter)
		_, err := setup.Handler(setup.Ctx, revokeModelMsg)
		require.Error(t, err)
		require.True(t, types.ErrModelVersionStringDoesNotMatch.Is(err))
	}
}

func TestHandler_RevokeModelTwice(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	for _, certificationType := range setup.CertificationTypes {
		// revoke model
		revokeModelMsg := NewMsgRevokeModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		_, err := setup.Handler(setup.Ctx, revokeModelMsg)
		require.NoError(t, err)

		// revoke model second time
		secondRevokeModelMsg := NewMsgRevokeModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		secondRevokeModelMsg.RevocationDate = time.Now().UTC().Format(time.RFC3339)
		_, err = setup.Handler(setup.Ctx, secondRevokeModelMsg)
		require.Error(t, err)
		require.True(t, types.ErrAlreadyRevoked.Is(err))
	}
}

func TestHandler_RevokeDifferentModels(t *testing.T) {
	setup := Setup(t)

	for i := 1; i < 5; i++ {
		// add model version
		vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
			int32(i), int32(i), uint32(i), fmt.Sprint(i))

		for _, certificationType := range setup.CertificationTypes {
			// revoke model
			revokeModelMsg := NewMsgRevokeModel(
				vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
			_, err := setup.Handler(setup.Ctx, revokeModelMsg)
			require.NoError(t, err)

			// query revoked model
			receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, certificationType)

			// check
			checkRevokedModelInfo(t, revokeModelMsg, receivedComplianceInfo)
		}
	}
}

func TestHandler_RevokeCertifiedModelForRevocationDateBeforeCertificationDate(t *testing.T) {
	setup := Setup(t)

	revocationTime := time.Now().UTC()
	revocationDate := revocationTime.Format(time.RFC3339)
	certificationDate := revocationTime.AddDate(0, 0, 1).Format(time.RFC3339)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	for _, certificationType := range setup.CertificationTypes {
		// certify model
		certifyModelMsg := NewMsgCertifyModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		certifyModelMsg.CertificationDate = certificationDate
		_, err := setup.Handler(setup.Ctx, certifyModelMsg)
		require.NoError(t, err)

		// try to revoke model
		revokeModelMsg := NewMsgRevokeModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		revokeModelMsg.RevocationDate = revocationDate
		_, err = setup.Handler(setup.Ctx, revokeModelMsg)
		require.Error(t, err)
		require.True(t, types.ErrInconsistentDates.Is(err))
	}
}

func TestHandler_RevokeProvisionedModelForRevocationDateBeforeProvisionalDate(t *testing.T) {
	setup := Setup(t)

	revocationTime := time.Now().UTC()
	revocationDate := revocationTime.Format(time.RFC3339)
	provisionalDate := revocationTime.AddDate(0, 0, 1).Format(time.RFC3339)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	for _, certificationType := range setup.CertificationTypes {
		// provision model
		provisionModelMsg := NewMsgProvisionModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		provisionModelMsg.ProvisionalDate = provisionalDate
		_, err := setup.Handler(setup.Ctx, provisionModelMsg)
		require.NoError(t, err)

		// try to revoke model
		revokeModelMsg := NewMsgRevokeModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		revokeModelMsg.RevocationDate = revocationDate
		_, err = setup.Handler(setup.Ctx, revokeModelMsg)
		require.Error(t, err)
		require.True(t, types.ErrInconsistentDates.Is(err))
	}
}

func TestHandler_CertifyRevokedModelForCertificationDateBeforeRevocationDate(t *testing.T) {
	setup := Setup(t)

	certificationTime := time.Now().UTC()
	certificationDate := certificationTime.Format(time.RFC3339)
	revocationDate := certificationTime.AddDate(0, 0, 1).Format(time.RFC3339)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	for _, certificationType := range setup.CertificationTypes {
		// revoke model
		revokeModelMsg := NewMsgRevokeModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		revokeModelMsg.RevocationDate = revocationDate
		_, err := setup.Handler(setup.Ctx, revokeModelMsg)
		require.NoError(t, err)

		// try to cancel model revocation
		certifyModelMsg := NewMsgCertifyModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		certifyModelMsg.CertificationDate = certificationDate
		_, err = setup.Handler(setup.Ctx, certifyModelMsg)
		require.Error(t, err)
		require.True(t, types.ErrInconsistentDates.Is(err))
	}
}

func TestHandler_CertifyRevokedModelThatWasCertifiedEarlier(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	for _, certificationType := range setup.CertificationTypes {
		// certify model
		certifyModelMsg := NewMsgCertifyModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		_, err := setup.Handler(setup.Ctx, certifyModelMsg)
		require.NoError(t, err)

		// revoke model
		revokeModelMsg := NewMsgRevokeModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		revokeModelMsg.RevocationDate = time.Now().UTC().Format(time.RFC3339)
		_, err = setup.Handler(setup.Ctx, revokeModelMsg)
		require.NoError(t, err)

		// certify model again
		secondCertifyModelMsg := NewMsgCertifyModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		secondCertifyModelMsg.CertificationDate = time.Now().UTC().Format(time.RFC3339)
		_, err = setup.Handler(setup.Ctx, secondCertifyModelMsg)
		require.NoError(t, err)

		// query certified model info
		receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, certificationType)

		// check
		checkCertifiedModelInfo(t, secondCertifyModelMsg, receivedComplianceInfo)
		require.Equal(t, 2, len(receivedComplianceInfo.History))

		require.Equal(t, dclcompltypes.CodeCertified, receivedComplianceInfo.History[0].SoftwareVersionCertificationStatus)
		require.Equal(t, certifyModelMsg.CertificationDate, receivedComplianceInfo.History[0].Date)

		require.Equal(t, dclcompltypes.CodeRevoked, receivedComplianceInfo.History[1].SoftwareVersionCertificationStatus)
		require.Equal(t, revokeModelMsg.RevocationDate, receivedComplianceInfo.History[1].Date)

		// query certified model
		certifiedModel, _ := queryCertifiedModel(setup, vid, pid, softwareVersion, certificationType)
		require.True(t, certifiedModel.Value)

		// query revoked model
		revokedModel, _ := queryRevokedModel(setup, vid, pid, softwareVersion, certificationType)
		require.False(t, revokedModel.Value)

		// query provisional model
		provisionalModel, _ := queryProvisionalModel(setup, vid, pid, softwareVersion, certificationType)
		require.False(t, provisionalModel.Value)
	}
}

func TestHandler_CDCertificateIDUpdateChangesOnlyOneComplianceInfo(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid1, pid1, softwareVersion1, softwareVersionString1 := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	// certify first model version
	certifyModelMsg := NewMsgCertifyModel(
		vid1, pid1, softwareVersion1, softwareVersionString1, dclcompltypes.ZigbeeCertificationType, setup.CertificationCenter)
	_, err := setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, err)

	// add second model version
	vid2, pid2, softwareVersion2, softwareVersionString2 := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion+1, testconstants.SoftwareVersionString)

	// certify second model version
	certifyModelMsg = NewMsgCertifyModel(
		vid2, pid2, softwareVersion2, softwareVersionString2, dclcompltypes.ZigbeeCertificationType, setup.CertificationCenter)
	_, err = setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, err)

	// Update compliance info of first model version
	updateComplianceInfoMsg := NewMsgUpdateComplianceInfo(setup.CertificationCenter, vid1, pid1, softwareVersion1, softwareVersionString1, dclcompltypes.ZigbeeCertificationType)
	updateComplianceInfoMsg.CDCertificateId += "new"
	_, err = setup.Handler(setup.Ctx, updateComplianceInfoMsg)
	require.NoError(t, err)

	firstComplianceInfo, err := queryComplianceInfo(setup, vid1, pid1, softwareVersion1, testconstants.CertificationType)
	require.NoError(t, err)

	secondComplianceInfo, err := queryComplianceInfo(setup, vid2, pid2, softwareVersion2, testconstants.CertificationType)
	require.NoError(t, err)

	require.Equal(t, updateComplianceInfoMsg.CDCertificateId, firstComplianceInfo.CDCertificateId)
	require.Equal(t, testconstants.CDCertificateID, secondComplianceInfo.CDCertificateId)
}

func TestHandler_UpdateWithTheSameCDCertificateID(t *testing.T) {
	setup := Setup(t)

	const cdCertID = "cdCertID"

	// add model version
	vid1, pid1, softwareVersion1, softwareVersionString1 := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	// certify first model version
	certifyModelMsg := NewMsgCertifyModel(
		vid1, pid1, softwareVersion1, softwareVersionString1, dclcompltypes.ZigbeeCertificationType, setup.CertificationCenter)
	certifyModelMsg.CDCertificateId = cdCertID

	_, err := setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, err)

	// add second model version
	vid2, pid2, softwareVersion2, softwareVersionString2 := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion+1, testconstants.SoftwareVersionString)

	// certify second model version
	certifyModelMsg = NewMsgCertifyModel(
		vid2, pid2, softwareVersion2, softwareVersionString2, dclcompltypes.ZigbeeCertificationType, setup.CertificationCenter)
	certifyModelMsg.CDCertificateId = cdCertID

	_, err = setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, err)

	// Update compliance info of first model version
	updateComplianceInfoMsg := &types.MsgUpdateComplianceInfo{
		Creator:           setup.CertificationCenter.String(),
		Vid:               testconstants.Vid,
		Pid:               pid1,
		SoftwareVersion:   softwareVersion1,
		CertificationType: dclcompltypes.ZigbeeCertificationType,
		CDCertificateId:   cdCertID,
	}

	originalDeviceSoftwareCompliance, _ := queryDeviceSoftwareCompliance(setup, cdCertID)

	_, err = setup.Handler(setup.Ctx, updateComplianceInfoMsg)
	require.NoError(t, err)

	newDeviceSoftwareCompliance, _ := queryDeviceSoftwareCompliance(setup, updateComplianceInfoMsg.CDCertificateId)

	firstComplianceInfo, err := queryComplianceInfo(setup, vid1, pid1, softwareVersion1, testconstants.CertificationType)
	require.NoError(t, err)

	secondComplianceInfo, err := queryComplianceInfo(setup, vid2, pid2, softwareVersion2, testconstants.CertificationType)
	require.NoError(t, err)

	require.Equal(t, updateComplianceInfoMsg.CDCertificateId, firstComplianceInfo.CDCertificateId)
	require.Equal(t, cdCertID, secondComplianceInfo.CDCertificateId)

	require.Equal(t, true, reflect.DeepEqual(originalDeviceSoftwareCompliance, newDeviceSoftwareCompliance))
}

func TestHandler_UpdateToAnotherCDCertificateID(t *testing.T) {
	setup := Setup(t)

	const (
		cdCertID1 = "cdCertID1"
		cdCertID2 = "cdCertID2"
	)

	// add model version
	vid1, pid1, softwareVersion1, softwareVersionString1 := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	// certify first model version
	certifyModelMsg := NewMsgCertifyModel(
		vid1, pid1, softwareVersion1, softwareVersionString1, dclcompltypes.ZigbeeCertificationType, setup.CertificationCenter)
	certifyModelMsg.CDCertificateId = cdCertID1

	_, err := setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, err)

	// add second model version
	vid2, pid2, softwareVersion2, softwareVersionString2 := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion+1, testconstants.SoftwareVersionString)

	// certify second model version
	certifyModelMsg = NewMsgCertifyModel(
		vid2, pid2, softwareVersion2, softwareVersionString2, dclcompltypes.ZigbeeCertificationType, setup.CertificationCenter)
	certifyModelMsg.CDCertificateId = cdCertID2

	_, err = setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, err)

	// Update compliance info of first model version
	updateComplianceInfoMsg := &types.MsgUpdateComplianceInfo{
		Creator:           setup.CertificationCenter.String(),
		Vid:               testconstants.Vid,
		Pid:               pid1,
		SoftwareVersion:   softwareVersion1,
		CertificationType: dclcompltypes.ZigbeeCertificationType,
		CDCertificateId:   cdCertID2,
	}

	originalDeviceSoftwareCompliance1, _ := queryDeviceSoftwareCompliance(setup, cdCertID1)
	originalDeviceSoftwareCompliance2, _ := queryDeviceSoftwareCompliance(setup, cdCertID2)

	_, err = setup.Handler(setup.Ctx, updateComplianceInfoMsg)
	require.NoError(t, err)

	newDeviceSoftwareCompliance1, _ := queryDeviceSoftwareCompliance(setup, cdCertID1)
	newDeviceSoftwareCompliance2, _ := queryDeviceSoftwareCompliance(setup, cdCertID2)

	firstComplianceInfo, err := queryComplianceInfo(setup, vid1, pid1, softwareVersion1, testconstants.CertificationType)
	require.NoError(t, err)

	secondComplianceInfo, err := queryComplianceInfo(setup, vid2, pid2, softwareVersion2, testconstants.CertificationType)
	require.NoError(t, err)

	require.Equal(t, updateComplianceInfoMsg.CDCertificateId, firstComplianceInfo.CDCertificateId)
	require.Equal(t, cdCertID2, secondComplianceInfo.CDCertificateId)

	require.Equal(t, cdCertID2, firstComplianceInfo.CDCertificateId)
	require.Equal(t, cdCertID2, newDeviceSoftwareCompliance2.ComplianceInfo[1].CDCertificateId)

	require.Equal(t, 1, len(originalDeviceSoftwareCompliance1.ComplianceInfo))
	require.Equal(t, 1, len(originalDeviceSoftwareCompliance2.ComplianceInfo))
	require.Nil(t, newDeviceSoftwareCompliance1)
	require.Equal(t, 2, len(newDeviceSoftwareCompliance2.ComplianceInfo))
	require.Equal(t, originalDeviceSoftwareCompliance2.ComplianceInfo[0], newDeviceSoftwareCompliance2.ComplianceInfo[0])

	cdCertificateIDExcluded := newDeviceSoftwareCompliance2.ComplianceInfo[1]
	cdCertificateIDExcluded.CDCertificateId = cdCertID1
	require.Equal(t, originalDeviceSoftwareCompliance1.ComplianceInfo[0], cdCertificateIDExcluded)
}

func TestHandler_CertifyRevokedModel(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	for index, certificationType := range setup.CertificationTypes {
		// revoke model
		revokeModelMsg := NewMsgRevokeModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		_, err := setup.Handler(setup.Ctx, revokeModelMsg)
		require.NoError(t, err)

		// cancel model revocation
		certifyModelMsg := NewMsgCertifyModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		certifyModelMsg.CertificationDate = time.Now().UTC().Format(time.RFC3339)
		_, err = setup.Handler(setup.Ctx, certifyModelMsg)
		require.NoError(t, err)

		// query certified model info
		receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, certificationType)
		require.Equal(t, certifyModelMsg.Vid, receivedComplianceInfo.Vid)
		require.Equal(t, certifyModelMsg.Pid, receivedComplianceInfo.Pid)
		require.Equal(t, dclcompltypes.CodeCertified, receivedComplianceInfo.SoftwareVersionCertificationStatus)
		require.Equal(t, certifyModelMsg.CertificationDate, receivedComplianceInfo.Date)
		require.Equal(t, certifyModelMsg.CDCertificateId, receivedComplianceInfo.CDCertificateId)
		require.Equal(t, certifyModelMsg.Reason, receivedComplianceInfo.Reason)
		require.Equal(t, certifyModelMsg.CertificationType, receivedComplianceInfo.CertificationType)
		require.Equal(t, 1, len(receivedComplianceInfo.History))
		require.Equal(t, dclcompltypes.CodeRevoked, receivedComplianceInfo.History[0].SoftwareVersionCertificationStatus)
		require.Equal(t, revokeModelMsg.RevocationDate, receivedComplianceInfo.History[0].Date)

		// query device software compliance
		receivedDeviceSoftwareCompliance, _ := queryDeviceSoftwareCompliance(setup, testconstants.CDCertificateID)
		require.Equal(t, receivedDeviceSoftwareCompliance.CDCertificateId, testconstants.CDCertificateID)
		checkDeviceSoftwareCompliance(t, receivedDeviceSoftwareCompliance.ComplianceInfo[index], receivedComplianceInfo)

		// query certified model
		certifiedModel, _ := queryCertifiedModel(setup, vid, pid, softwareVersion, certificationType)
		require.True(t, certifiedModel.Value)

		// query revoked model
		revokedModel, _ := queryRevokedModel(setup, vid, pid, softwareVersion, certificationType)
		require.False(t, revokedModel.Value)

		// query provisional model
		provisionalModel, _ := queryProvisionalModel(setup, vid, pid, softwareVersion, certificationType)
		require.False(t, provisionalModel.Value)
	}
}

func TestHandler_DeleteComplianceInfoForRevokedModel(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion+2, testconstants.SoftwareVersionString)

	// Revoke model
	revokeModelMsg := NewMsgRevokeModel(
		vid, pid, softwareVersion, softwareVersionString, dclcompltypes.MatterCertificationType, setup.CertificationCenter)
	_, err := setup.Handler(setup.Ctx, revokeModelMsg)
	require.NoError(t, err)

	deleteComplInfoMsg := NewMsgDeleteComplianceInfo(vid, pid, softwareVersion, dclcompltypes.MatterCertificationType, setup.CertificationCenter)
	_, err = setup.Handler(setup.Ctx, deleteComplInfoMsg)
	require.NoError(t, err)

	revokedModel, err := queryRevokedModel(setup, vid, pid, softwareVersion, dclcompltypes.MatterCertificationType)
	require.Equal(setup.T, err.Error(), "rpc error: code = NotFound desc = not found")
	require.Nil(setup.T, revokedModel)

	deviceSoftwareCompliance, err := queryDeviceSoftwareCompliance(setup, testconstants.CDCertificateID)
	require.Equal(setup.T, err.Error(), "rpc error: code = NotFound desc = not found")
	require.Nil(setup.T, deviceSoftwareCompliance)

	complianceInfo, err := queryComplianceInfo(setup, vid, pid, softwareVersion, dclcompltypes.MatterCertificationType)
	require.Equal(setup.T, err.Error(), "rpc error: code = NotFound desc = not found")
	require.Nil(setup.T, complianceInfo)
}

func TestHandler_DeleteComplianceInfoForCertifiedModel(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	// certify model
	certifyModelMsg := NewMsgCertifyModel(
		vid, pid, softwareVersion, softwareVersionString, dclcompltypes.MatterCertificationType, setup.CertificationCenter)
	_, err := setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, err)
	deleteComplInfoMsg := NewMsgDeleteComplianceInfo(vid, pid, softwareVersion, dclcompltypes.MatterCertificationType, setup.CertificationCenter)
	_, err = setup.Handler(setup.Ctx, deleteComplInfoMsg)
	require.NoError(t, err)

	// check
	certifiedModel, err := queryCertifiedModel(setup, vid, pid, softwareVersion, dclcompltypes.MatterCertificationType)
	require.Equal(setup.T, err.Error(), "rpc error: code = NotFound desc = not found")
	require.Nil(setup.T, certifiedModel)

	deviceSoftwareCompliance, err := queryDeviceSoftwareCompliance(setup, certifyModelMsg.CDCertificateId)
	require.Equal(setup.T, err.Error(), "rpc error: code = NotFound desc = not found")
	require.Nil(setup.T, deviceSoftwareCompliance)

	complianceInfo, err := queryComplianceInfo(setup, vid, pid, softwareVersion, dclcompltypes.MatterCertificationType)
	require.Equal(setup.T, err.Error(), "rpc error: code = NotFound desc = not found")
	require.Nil(setup.T, complianceInfo)
}

func TestHandler_DeleteComplianceInfoForProvisionalModel(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion+1, testconstants.SoftwareVersionString)

	// Provision model
	provisionModelMsg := NewMsgProvisionModel(
		vid, pid, softwareVersion, softwareVersionString, dclcompltypes.MatterCertificationType, setup.CertificationCenter)
	_, err := setup.Handler(setup.Ctx, provisionModelMsg)
	require.NoError(t, err)

	deleteComplInfoMsg := NewMsgDeleteComplianceInfo(vid, pid, softwareVersion, dclcompltypes.MatterCertificationType, setup.CertificationCenter)
	_, err = setup.Handler(setup.Ctx, deleteComplInfoMsg)
	require.NoError(t, err)

	provisionalModel, err := queryProvisionalModel(setup, vid, pid, softwareVersion, dclcompltypes.MatterCertificationType)
	require.Equal(setup.T, err.Error(), "rpc error: code = NotFound desc = not found")
	require.Nil(setup.T, provisionalModel)

	deviceSoftwareCompliance, err := queryDeviceSoftwareCompliance(setup, provisionModelMsg.CDCertificateId)
	require.Equal(setup.T, err.Error(), "rpc error: code = NotFound desc = not found")
	require.Nil(setup.T, deviceSoftwareCompliance)

	complianceInfo, err := queryComplianceInfo(setup, vid, pid, softwareVersion, dclcompltypes.MatterCertificationType)
	require.Equal(setup.T, err.Error(), "rpc error: code = NotFound desc = not found")
	require.Nil(setup.T, complianceInfo)
}

func TestHandler_DeleteComplianceInfoDoesNotExist(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid, pid, softwareVersion, _ := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	deleteComplInfoMsg := NewMsgDeleteComplianceInfo(vid, pid, softwareVersion, dclcompltypes.MatterCertificationType, setup.CertificationCenter)
	_, err := setup.Handler(setup.Ctx, deleteComplInfoMsg)
	require.ErrorIs(t, err, types.ErrComplianceInfoDoesNotExist)
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

func NewMsgProvisionModel(
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

func NewMsgUpdateComplianceInfo(
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

func NewMsgUpdateComplianceInfoWithAllOptionalFlags(
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

func NewMsgProvisionModelWithAllOptionalFlags(
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

func NewMsgCertifyModel(
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

func NewMsgDeleteComplianceInfo(
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

func NewMsgCertifyModelWithAllOptionalFlags(
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

func NewMsgRevokeModel(
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

func NewModelVersion(
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
		Creator:                      GenerateAccAddress().String(),
	}
}

func GenerateAccAddress() sdk.AccAddress {
	_, _, accAddress := testdata.KeyTestPubAddr()

	return accAddress
}
