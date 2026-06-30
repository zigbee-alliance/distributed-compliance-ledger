package compliance

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func setupProvisionModel(t *testing.T) (*TestSetup, int32, int32, uint32, string, string) {
	t.Helper()
	setup := setup(t)

	vid, pid, softwareVersion, softwareVersionString := setup.addModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)
	certificationType := types.ZigbeeCertificationType

	return setup, vid, pid, softwareVersion, softwareVersionString, certificationType
}

func (setup *TestSetup) provisionModel(vid int32, pid int32, softwareVersion uint32, softwareVersionString string, certificationType string, signer sdk.AccAddress) (*types.MsgProvisionModel, error) {
	provisionModelMsg := newMsgProvisionModel(
		vid, pid, softwareVersion, softwareVersionString, certificationType, signer)
	_, err := setup.Handler(setup.Ctx, provisionModelMsg)

	return provisionModelMsg, err
}

func (setup *TestSetup) provisionModelByDate(vid int32, pid int32, softwareVersion uint32, softwareVersionString string, certificationType string, provisionalDate string, signer sdk.AccAddress) error {
	provisionModelMsg := newMsgProvisionModel(
		vid, pid, softwareVersion, softwareVersionString, certificationType, signer)
	provisionModelMsg.ProvisionalDate = provisionalDate
	_, err := setup.Handler(setup.Ctx, provisionModelMsg)

	return err
}

func (setup *TestSetup) provisionModelWithAllOptionalFlags(vid int32, pid int32, softwareVersion uint32, softwareVersionString string, certificationType string, signer sdk.AccAddress) (*types.MsgProvisionModel, error) {
	provisionModelMsg := newMsgProvisionModelWithAllOptionalFlags(
		vid, pid, softwareVersion, softwareVersionString, certificationType, signer)
	_, err := setup.Handler(setup.Ctx, provisionModelMsg)

	return provisionModelMsg, err
}

func (setup *TestSetup) checkComplianceInfoEqualsProvisionModelMsgData(t *testing.T, provisionModelMsg *types.MsgProvisionModel) {
	t.Helper()
	vid := provisionModelMsg.Vid
	pid := provisionModelMsg.Pid
	softwareVersion := provisionModelMsg.SoftwareVersion
	certificationType := provisionModelMsg.CertificationType

	receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, certificationType)

	checkProvisionalModelInfo(t, provisionModelMsg, receivedComplianceInfo)
}

func (setup *TestSetup) checkModelProvisioned(t *testing.T, provisionModelMsg *types.MsgProvisionModel) {
	t.Helper()
	vid := provisionModelMsg.Vid
	pid := provisionModelMsg.Pid
	softwareVersion := provisionModelMsg.SoftwareVersion
	certificationType := provisionModelMsg.CertificationType

	provisionalModel, _ := queryProvisionalModel(setup, vid, pid, softwareVersion, certificationType)
	require.True(t, provisionalModel.Value)

	_, err := queryCertifiedModel(setup, vid, pid, softwareVersion, certificationType)
	assertNotFound(t, err)

	_, err = queryRevokedModel(setup, vid, pid, softwareVersion, certificationType)
	assertNotFound(t, err)
}

func TestHandler_ProvisionModel_AllCertificationTypes(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, _ := setupProvisionModel(t)

	for _, certificationType := range setup.CertificationTypes {
		provisionModelMsg, provisionModelErr := setup.provisionModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)

		require.NoError(t, provisionModelErr)
		setup.checkComplianceInfoEqualsProvisionModelMsgData(t, provisionModelMsg)
		setup.checkModelProvisioned(t, provisionModelMsg)
	}
}

func TestHandler_ProvisionModel_WithAllOptionalFlagsForAllCertificationTypes(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, _ := setupProvisionModel(t)

	for _, certificationType := range setup.CertificationTypes {
		provisionModelMsg, provisionModelErr := setup.provisionModelWithAllOptionalFlags(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)

		require.NoError(t, provisionModelErr)
		setup.checkComplianceInfoEqualsProvisionModelMsgData(t, provisionModelMsg)
		setup.checkModelProvisioned(t, provisionModelMsg)
	}
}

func TestHandler_ProvisionModelWhenModelVersionDoesNotExist(t *testing.T) {
	setup := setup(t)
	vid, pid, softwareVersion := testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion
	setup.setNoModelVersionForKey(vid, pid, softwareVersion)

	provisionModelMsg, provisionModelErr := setup.provisionModel(vid, pid, softwareVersion, testconstants.SoftwareVersionString, types.ZigbeeCertificationType, setup.CertificationCenter)
	require.NoError(t, provisionModelErr)

	setup.checkComplianceInfoEqualsProvisionModelMsgData(t, provisionModelMsg)
	setup.checkModelProvisioned(t, provisionModelMsg)
}

func TestHandler_ProvisionModel_ByWrongRoles(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupProvisionModel(t)

	accountRoles := []dclauthtypes.AccountRole{
		dclauthtypes.Vendor,
		dclauthtypes.Trustee,
		dclauthtypes.NodeAdmin,
	}

	setup.setNoModelVersionForKey(vid, pid, softwareVersion)

	for _, role := range accountRoles {
		accAddress := generateAccAddress()
		setup.addAccount(accAddress, []dclauthtypes.AccountRole{role})

		_, provisionModelErr := setup.provisionModel(vid, pid, softwareVersion, softwareVersionString, certificationType, accAddress)

		require.ErrorIs(t, provisionModelErr, sdkerrors.ErrUnauthorized)
	}
}

func TestHandler_ProvisionModel_Twice(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupProvisionModel(t)

	_, provisionModelFisrtTimeErr := setup.provisionModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	_, provisionModelSecondTimeErr := setup.provisionModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)

	require.NoError(t, provisionModelFisrtTimeErr)
	require.Error(t, provisionModelSecondTimeErr)
}

func TestHandler_ProvisionModel_AlreadyCertified(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupProvisionModel(t)

	certifyModelMsg := newMsgCertifyModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	_, certifyModelErr := setup.Handler(setup.Ctx, certifyModelMsg)
	require.NoError(t, certifyModelErr)
	_, provisionModelErr := setup.provisionModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)

	require.ErrorIs(t, provisionModelErr, types.ErrAlreadyCertified)
}

func TestHandler_ProvisionModel_AlreadyRevoked(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupProvisionModel(t)

	_, revokeModelErr := setup.revokeModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	require.NoError(t, revokeModelErr)
	_, provisionModelErr := setup.provisionModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)

	require.ErrorIs(t, provisionModelErr, types.ErrAlreadyRevoked)
}

func TestHandler_ProvisionModel_MoreThanOneModel(t *testing.T) {
	setup := setup(t)
	certificationType := types.ZigbeeCertificationType
	modelVersionsQuantity := 5

	for i := 1; i < modelVersionsQuantity; i++ {
		vid := int32(i)
		pid := int32(i)
		softwareVersion := uint32(i)
		softwareVersionString := fmt.Sprint(i)

		setup.addModelVersion(vid, pid, softwareVersion, softwareVersionString)

		provisionModelMsg, provisionModelErr := setup.provisionModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)

		require.NoError(t, provisionModelErr)
		setup.checkComplianceInfoEqualsProvisionModelMsgData(t, provisionModelMsg)
		setup.checkModelProvisioned(t, provisionModelMsg)
	}
}

func TestHandler_SchemaVersion_ProvisionModel_StampsCurrentOnCreate(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupProvisionModel(t)

	_, err := setup.provisionModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	require.NoError(t, err)

	complianceInfo := queryExistingComplianceInfo(setup, vid, pid, softwareVersion, certificationType)
	require.Equal(t, complianceInfo.CurrentSchemaVersion(), complianceInfo.SchemaVersion)

	provisionalModel, _ := queryProvisionalModel(setup, vid, pid, softwareVersion, certificationType)
	require.Equal(t, provisionalModel.CurrentSchemaVersion(), provisionalModel.SchemaVersion)
}

// Covers the default switch arm: existing compliance info whose status is none of
// provisional/certified/revoked.
func TestHandler_ProvisionModel_ExistingInfoUnknownStatus(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := setupProvisionModel(t)

	setup.Keeper.SetComplianceInfo(setup.Ctx, &types.ComplianceInfo{
		Vid:                                vid,
		Pid:                                pid,
		SoftwareVersion:                    softwareVersion,
		CertificationType:                  certificationType,
		SoftwareVersionCertificationStatus: 99, // not provisional/certified/revoked
	})

	_, err := setup.provisionModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	require.ErrorIs(t, err, types.ErrComplianceInfoAlreadyExist)
}

// Covers the branch where the stored model version's software version string does
// not match the one in the message.
func TestHandler_ProvisionModel_ModelVersionStringMismatch(t *testing.T) {
	setup, vid, pid, softwareVersion, _, certificationType := setupProvisionModel(t)

	_, err := setup.provisionModel(vid, pid, softwareVersion, "different-version-string", certificationType, setup.CertificationCenter)
	require.ErrorIs(t, err, types.ErrModelVersionStringDoesNotMatch)
}

// Covers the branch where the stored model version's CD version number does not
// match the one in the message (software version string still matches).
func TestHandler_ProvisionModel_CdVersionNumberMismatch(t *testing.T) {
	setup := setup(t)
	vid, pid, softwareVersion := testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion

	modelVersion := newModelVersion(vid, pid, softwareVersion, testconstants.SoftwareVersionString)
	modelVersion.CdVersionNumber = testconstants.CdVersionNumber + 1
	setup.ModelKeeper.On("GetModelVersion", mock.Anything, vid, pid, softwareVersion).Return(*modelVersion, true)

	_, err := setup.provisionModel(vid, pid, softwareVersion, testconstants.SoftwareVersionString, types.ZigbeeCertificationType, setup.CertificationCenter)
	require.ErrorIs(t, err, types.ErrModelVersionStringDoesNotMatch)
}
