package compliance

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	testkeeper "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
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
	CertificationTypes  types.CertificationTypes
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
	dclauthKeeper := &DclauthKeeperMock{}
	modelKeeper := &ModelKeeperMock{}
	keeper, ctx := testkeeper.ComplianceKeeper(t, dclauthKeeper, modelKeeper)

	certificationCenter := GenerateAccAddress()

	certificationTypes := types.CertificationTypes{types.ZigbeeCertificationType, types.MatterCertificationType}

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

func TestHandler_ProvisionModel(t *testing.T) {
	setup := Setup(t)

	vid := testconstants.Vid
	pid := testconstants.Pid
	softwareVersion := testconstants.SoftwareVersion
	softwareVersionString := testconstants.SoftwareVersionString

	// set absence of model version
	setup.SetNoModelVersionForKey(vid, pid, softwareVersion)

	for _, certificationType := range setup.CertificationTypes {
		// provision model
		provisionModelMsg := NewMsgProvisionModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		_, err := setup.Handler(setup.Ctx, provisionModelMsg)
		require.NoError(t, err)

		// query provisional model
		receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, certificationType)

		// check
		checkProvisionalModelInfo(t, provisionModelMsg, receivedComplianceInfo)

		provisionalModel, _ := queryProvisionalModel(setup, vid, pid, softwareVersion, certificationType)
		require.True(t, provisionalModel.Value)

		_, err = queryCertifiedModel(setup, vid, pid, softwareVersion, certificationType)
		require.Error(t, err)
		require.Equal(t, codes.NotFound, status.Code(err))

		_, err = queryRevokedModel(setup, vid, pid, softwareVersion, certificationType)
		require.Error(t, err)
		require.Equal(t, codes.NotFound, status.Code(err))
	}
}

func TestHandler_ProvisionModelByDifferentRoles(t *testing.T) {
	setup := Setup(t)

	vid := testconstants.Vid
	pid := testconstants.Pid
	softwareVersion := testconstants.SoftwareVersion
	softwareVersionString := testconstants.SoftwareVersionString

	// set absence of model version
	setup.SetNoModelVersionForKey(vid, pid, softwareVersion)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.Vendor,
		dclauthtypes.Trustee,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role})

		// try to provision model
		for _, certificationType := range setup.CertificationTypes {
			provisionModelMsg := NewMsgProvisionModel(
				vid, pid, softwareVersion, softwareVersionString, certificationType, accAddress)
			_, err := setup.Handler(setup.Ctx, provisionModelMsg)
			require.Error(t, err)
			require.True(t, sdkerrors.ErrUnauthorized.Is(err))
		}
	}
}

func TestHandler_ProvisionModelTwice(t *testing.T) {
	setup := Setup(t)

	vid := testconstants.Vid
	pid := testconstants.Pid
	softwareVersion := testconstants.SoftwareVersion
	softwareVersionString := testconstants.SoftwareVersionString

	// set absence of model version
	setup.SetNoModelVersionForKey(vid, pid, softwareVersion)

	for _, certificationType := range setup.CertificationTypes {
		// provision model
		provisionModelMsg := NewMsgProvisionModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		_, err := setup.Handler(setup.Ctx, provisionModelMsg)
		require.NoError(t, err)

		// provision model second time
		secondProvisionModelMsg := NewMsgProvisionModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		secondProvisionModelMsg.ProvisionalDate = time.Now().UTC().Format(time.RFC3339)
		_, err = setup.Handler(setup.Ctx, secondProvisionModelMsg)
		require.Error(t, err)
		require.True(t, types.ErrAlreadyProvisional.Is(err))
	}
}

func TestHandler_ProvisionCertifiedModel(t *testing.T) {
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

		// try to provision model
		provisionModelMsg := NewMsgProvisionModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		provisionModelMsg.ProvisionalDate = time.Now().UTC().Format(time.RFC3339)
		_, err = setup.Handler(setup.Ctx, provisionModelMsg)
		require.Error(t, err)
		require.True(t, types.ErrAlreadyCertified.Is(err))
	}
}

func TestHandler_ProvisionRevokedModel(t *testing.T) {
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

		// try to provision model
		provisionModelMsg := NewMsgProvisionModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		provisionModelMsg.ProvisionalDate = time.Now().UTC().Format(time.RFC3339)
		_, err = setup.Handler(setup.Ctx, provisionModelMsg)
		require.Error(t, err)
		require.True(t, types.ErrAlreadyRevoked.Is(err))
	}
}

func TestHandler_ProvisionDifferentModels(t *testing.T) {
	setup := Setup(t)

	for i := 1; i < 5; i++ {
		vid := int32(i)
		pid := int32(i)
		softwareVersion := uint32(i)
		softwareVersionString := fmt.Sprint(i)

		// set absence of model version
		setup.SetNoModelVersionForKey(vid, pid, softwareVersion)

		for _, certificationType := range setup.CertificationTypes {
			// provision model
			provisionModelMsg := NewMsgProvisionModel(
				vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
			_, err := setup.Handler(setup.Ctx, provisionModelMsg)
			require.NoError(t, err)

			// query provisional model
			receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, certificationType)

			// check
			checkProvisionalModelInfo(t, provisionModelMsg, receivedComplianceInfo)
		}
	}
}

func TestHandler_CertifyModel_Zigbee(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	// certify model
	certifyModelMsg := NewMsgCertifyModel(
		vid, pid, softwareVersion, softwareVersionString, types.ZigbeeCertificationType, setup.CertificationCenter)
	_, err := setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, err)

	// query certified model
	receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, types.ZigbeeCertificationType)

	// check
	checkCertifiedModelInfo(t, certifyModelMsg, receivedComplianceInfo)

	certifiedModel, _ := queryCertifiedModel(setup, vid, pid, softwareVersion, types.ZigbeeCertificationType)
	require.True(t, certifiedModel.Value)

	revokedModel, _ := queryRevokedModel(setup, vid, pid, softwareVersion, types.ZigbeeCertificationType)
	require.False(t, revokedModel.Value)

	provisionalModel, _ := queryProvisionalModel(setup, vid, pid, softwareVersion, types.ZigbeeCertificationType)
	require.False(t, provisionalModel.Value)
}

func TestHandler_CertifyModel_Matter(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	// certify model
	certifyModelMsg := NewMsgCertifyModel(
		vid, pid, softwareVersion, softwareVersionString, types.MatterCertificationType, setup.CertificationCenter)
	_, err := setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, err)

	// query certified model
	receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, types.MatterCertificationType)

	// check
	checkCertifiedModelInfo(t, certifyModelMsg, receivedComplianceInfo)

	certifiedModel, _ := queryCertifiedModel(setup, vid, pid, softwareVersion, types.MatterCertificationType)
	require.True(t, certifiedModel.Value)

	revokedModel, _ := queryRevokedModel(setup, vid, pid, softwareVersion, types.MatterCertificationType)
	require.False(t, revokedModel.Value)

	provisionalModel, _ := queryProvisionalModel(setup, vid, pid, softwareVersion, types.MatterCertificationType)
	require.False(t, provisionalModel.Value)
}

func TestHandler_CertifyProvisionedModel(t *testing.T) {
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

		// certify model
		certifyModelMsg := NewMsgCertifyModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		certifyModelMsg.CertificationDate = time.Now().UTC().Format(time.RFC3339)
		_, err = setup.Handler(setup.Ctx, certifyModelMsg)
		require.NoError(t, err)

		// query certified model info
		receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, certificationType)

		// check
		checkCertifiedModelInfo(t, certifyModelMsg, receivedComplianceInfo)
		require.Equal(t, 1, len(receivedComplianceInfo.History))

		require.Equal(t, types.CodeProvisional, receivedComplianceInfo.History[0].SoftwareVersionCertificationStatus)
		require.Equal(t, provisionModelMsg.ProvisionalDate, receivedComplianceInfo.History[0].Date)

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

	// set absence of model version
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
		require.Equal(t, types.CodeCertified, receivedComplianceInfo.History[0].SoftwareVersionCertificationStatus)
		require.Equal(t, certifyModelMsg.CertificationDate, receivedComplianceInfo.History[0].Date)

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

		require.Equal(t, types.CodeProvisional, receivedComplianceInfo.History[0].SoftwareVersionCertificationStatus)
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

	// set absence of model version
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

		require.Equal(t, types.CodeCertified, receivedComplianceInfo.History[0].SoftwareVersionCertificationStatus)
		require.Equal(t, certifyModelMsg.CertificationDate, receivedComplianceInfo.History[0].Date)

		require.Equal(t, types.CodeRevoked, receivedComplianceInfo.History[1].SoftwareVersionCertificationStatus)
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

func TestHandler_CertifyRevokedModel(t *testing.T) {
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

		// cancel model revocation
		certifyModelMsg := NewMsgCertifyModel(
			vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
		certifyModelMsg.CertificationDate = time.Now().UTC().Format(time.RFC3339)
		_, err = setup.Handler(setup.Ctx, certifyModelMsg)
		require.NoError(t, err)

		// query certified model info
		receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, certificationType)

		// check
		checkCertifiedModelInfo(t, certifyModelMsg, receivedComplianceInfo)
		require.Equal(t, 1, len(receivedComplianceInfo.History))

		require.Equal(t, types.CodeRevoked, receivedComplianceInfo.History[0].SoftwareVersionCertificationStatus)
		require.Equal(t, revokeModelMsg.RevocationDate, receivedComplianceInfo.History[0].Date)

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

func queryComplianceInfo(
	setup *TestSetup,
	vid int32,
	pid int32,
	softwareVersion uint32,
	certificationType string,
) (*types.ComplianceInfo, error) {

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
	receivedComplianceInfo *types.ComplianceInfo,
) {
	require.Equal(t, provisionalModelMsg.Vid, receivedComplianceInfo.Vid)
	require.Equal(t, provisionalModelMsg.Pid, receivedComplianceInfo.Pid)
	require.Equal(t, types.CodeProvisional, receivedComplianceInfo.SoftwareVersionCertificationStatus)
	require.Equal(t, provisionalModelMsg.ProvisionalDate, receivedComplianceInfo.Date)
	require.Equal(t, provisionalModelMsg.Reason, receivedComplianceInfo.Reason)
	require.Equal(t, provisionalModelMsg.CertificationType, receivedComplianceInfo.CertificationType)
}

func checkCertifiedModelInfo(
	t *testing.T,
	certifyModelMsg *types.MsgCertifyModel,
	receivedComplianceInfo *types.ComplianceInfo,
) {
	require.Equal(t, certifyModelMsg.Vid, receivedComplianceInfo.Vid)
	require.Equal(t, certifyModelMsg.Pid, receivedComplianceInfo.Pid)
	require.Equal(t, types.CodeCertified, receivedComplianceInfo.SoftwareVersionCertificationStatus)
	require.Equal(t, certifyModelMsg.CertificationDate, receivedComplianceInfo.Date)
	require.Equal(t, certifyModelMsg.Reason, receivedComplianceInfo.Reason)
	require.Equal(t, certifyModelMsg.CertificationType, receivedComplianceInfo.CertificationType)
}

func checkRevokedModelInfo(
	t *testing.T,
	revokeModelMsg *types.MsgRevokeModel,
	receivedComplianceInfo *types.ComplianceInfo,
) {
	require.Equal(t, revokeModelMsg.Vid, receivedComplianceInfo.Vid)
	require.Equal(t, revokeModelMsg.Pid, receivedComplianceInfo.Pid)
	require.Equal(t, types.CodeRevoked, receivedComplianceInfo.SoftwareVersionCertificationStatus)
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
		FirmwareDigests:              testconstants.FirmwareDigests,
		SoftwareVersionValid:         testconstants.SoftwareVersionValid,
		OtaUrl:                       testconstants.OtaUrl,
		OtaFileSize:                  testconstants.OtaFileSize,
		OtaChecksum:                  testconstants.OtaChecksum,
		OtaChecksumType:              testconstants.OtaChecksumType,
		MinApplicableSoftwareVersion: testconstants.MinApplicableSoftwareVersion,
		MaxApplicableSoftwareVersion: testconstants.MaxApplicableSoftwareVersion,
		ReleaseNotesUrl:              testconstants.ReleaseNotesUrl,
		Creator:                      GenerateAccAddress().String(),
	}
}

func GenerateAccAddress() sdk.AccAddress {
	_, _, accAddress := testdata.KeyTestPubAddr()
	return accAddress
}
