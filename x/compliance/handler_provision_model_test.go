package compliance

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	dclcompltypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/compliance"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func provisionModelSetup(t *testing.T) (*TestSetup, int32, int32, uint32, string, string) {
	setup := setup(t)

	vid, pid, softwareVersion, softwareVersionString := setup.addModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)
	certificationType := dclcompltypes.ZigbeeCertificationType

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
	vid := provisionModelMsg.Vid
	pid := provisionModelMsg.Pid
	softwareVersion := provisionModelMsg.SoftwareVersion
	certificationType := provisionModelMsg.CertificationType

	receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, certificationType)

	checkProvisionalModelInfo(t, provisionModelMsg, receivedComplianceInfo)
}

func (setup *TestSetup) checkModelProvisioned(t *testing.T, provisionModelMsg *types.MsgProvisionModel) {
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
	setup, vid, pid, softwareVersion, softwareVersionString, _ := provisionModelSetup(t)

	for _, certificationType := range setup.CertificationTypes {
		provisionModelMsg, provisionModelErr := setup.provisionModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)

		require.NoError(t, provisionModelErr)
		setup.checkComplianceInfoEqualsProvisionModelMsgData(t, provisionModelMsg)
		setup.checkModelProvisioned(t, provisionModelMsg)
	}
}

func TestHandler_ProvisionModel_WithAllOptionalFlagsForAllCertificationTypes(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, _ := provisionModelSetup(t)

	for _, certificationType := range setup.CertificationTypes {
		provisionModelMsg, provisionModelErr := setup.provisionModelWithAllOptionalFlags(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)

		require.NoError(t, provisionModelErr)
		setup.checkComplianceInfoEqualsProvisionModelMsgData(t, provisionModelMsg)
		setup.checkModelProvisioned(t, provisionModelMsg)
	}
}

func TestHandler_ProvisionModel_ByWrongRoles(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := provisionModelSetup(t)

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
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := provisionModelSetup(t)

	_, provisionModelFisrtTimeErr := setup.provisionModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	_, provisionModelSecondTimeErr := setup.provisionModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)

	require.NoError(t, provisionModelFisrtTimeErr)
	require.Error(t, provisionModelSecondTimeErr)
}

func TestHandler_ProvisionModel_AlreadyCertified(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := provisionModelSetup(t)

	_, certifyModelErr := setup.CertifyModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	require.NoError(t, certifyModelErr)
	_, provisionModelErr := setup.provisionModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)

	require.ErrorIs(t, provisionModelErr, types.ErrAlreadyCertified)
}

func TestHandler_ProvisionModel_AlreadyRevoked(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := provisionModelSetup(t)

	_, revokeModelErr := setup.revokeModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	require.NoError(t, revokeModelErr)
	_, provisionModelErr := setup.provisionModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)

	require.ErrorIs(t, provisionModelErr, types.ErrAlreadyRevoked)
}

func TestHandler_ProvisionModel_MoreThanOneModel(t *testing.T) {
	setup := setup(t)
	certificationType := dclcompltypes.ZigbeeCertificationType
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
