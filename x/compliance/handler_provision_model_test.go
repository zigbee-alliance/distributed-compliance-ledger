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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (setup *TestSetup) RevokeModel(t *testing.T, vid int32, pid int32, softwareVersion uint32, softwareVersionString string, certificationType string, signer sdk.AccAddress) *types.MsgRevokeModel {
	revokeModelMsg := NewMsgRevokeModel(
		vid, pid, softwareVersion, softwareVersionString, certificationType, signer)
	_, err := setup.Handler(setup.Ctx, revokeModelMsg)
	require.NoError(t, err)

	return revokeModelMsg
}

func (setup *TestSetup) ProvisionModel(vid int32, pid int32, softwareVersion uint32, softwareVersionString string, certificationType string, signer sdk.AccAddress) (*types.MsgProvisionModel, error) {
	provisionModelMsg := NewMsgProvisionModel(
		vid, pid, softwareVersion, softwareVersionString, certificationType, signer)
	_, err := setup.Handler(setup.Ctx, provisionModelMsg)

	return provisionModelMsg, err
}

func (setup *TestSetup) ProvisionModelWithAllOptionalFlags(vid int32, pid int32, softwareVersion uint32, softwareVersionString string, certificationType string, signer sdk.AccAddress) (*types.MsgProvisionModel, error) {
	provisionModelMsg := NewMsgProvisionModelWithAllOptionalFlags(
		vid, pid, softwareVersion, softwareVersionString, certificationType, signer)
	_, err := setup.Handler(setup.Ctx, provisionModelMsg)

	return provisionModelMsg, err
}

func (setup *TestSetup) CheckComplianceInfoEqualsProvisionModelMsgData(t *testing.T, provisionModelMsg *types.MsgProvisionModel) {
	vid := provisionModelMsg.Vid
	pid := provisionModelMsg.Pid
	softwareVersion := provisionModelMsg.SoftwareVersion
	certificationType := provisionModelMsg.CertificationType

	// query provisional model
	receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, certificationType)

	// check
	checkProvisionalModelInfo(t, provisionModelMsg, receivedComplianceInfo)
}

func (setup *TestSetup) CheckProvisionalModelAddedToCorrectStore(t *testing.T, provisionModelMsg *types.MsgProvisionModel) {
	vid := provisionModelMsg.Vid
	pid := provisionModelMsg.Pid
	softwareVersion := provisionModelMsg.SoftwareVersion
	certificationType := provisionModelMsg.CertificationType

	provisionalModel, _ := queryProvisionalModel(setup, vid, pid, softwareVersion, certificationType)
	require.True(t, provisionalModel.Value)

	_, err := queryCertifiedModel(setup, vid, pid, softwareVersion, certificationType)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	_, err = queryRevokedModel(setup, vid, pid, softwareVersion, certificationType)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
}

func ProvisionModelSetup(t *testing.T) (*TestSetup, int32, int32, uint32, string, string) {
	setup := Setup(t)

	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)
	certificationType := dclcompltypes.ZigbeeCertificationType

	return setup, vid, pid, softwareVersion, softwareVersionString, certificationType
}

func TestHandler_ProvisionModel_AllCertificationTypes(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, _ := ProvisionModelSetup(t)

	for _, certificationType := range setup.CertificationTypes {
		provisionModelMsg, provisionModelErr := setup.ProvisionModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)

		require.NoError(t, provisionModelErr)
		setup.CheckComplianceInfoEqualsProvisionModelMsgData(t, provisionModelMsg)
		setup.CheckProvisionalModelAddedToCorrectStore(t, provisionModelMsg)
	}
}

func TestHandler_ProvisionModel_WithAllOptionalFlagsForAllCertificationTypes(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, _ := ProvisionModelSetup(t)

	for _, certificationType := range setup.CertificationTypes {
		provisionModelMsg, provisionModelErr := setup.ProvisionModelWithAllOptionalFlags(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)

		require.NoError(t, provisionModelErr)
		setup.CheckComplianceInfoEqualsProvisionModelMsgData(t, provisionModelMsg)
		setup.CheckProvisionalModelAddedToCorrectStore(t, provisionModelMsg)
	}
}

func TestHandler_ProvisionModel_ByWrongRoles(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := ProvisionModelSetup(t)

	accountRoles := []dclauthtypes.AccountRole{
		dclauthtypes.Vendor,
		dclauthtypes.Trustee,
		dclauthtypes.NodeAdmin,
	}

	setup.SetNoModelVersionForKey(vid, pid, softwareVersion)

	for _, role := range accountRoles {
		accAddress := GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role})

		_, provisionModelErr := setup.ProvisionModel(vid, pid, softwareVersion, softwareVersionString, certificationType, accAddress)

		require.Error(t, provisionModelErr)
		require.True(t, sdkerrors.ErrUnauthorized.Is(provisionModelErr))
	}
}

func TestHandler_ProvisionModel_Twice(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := ProvisionModelSetup(t)

	_, provisionModelFisrtTimeErr := setup.ProvisionModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	_, provisionModelSecondTimeErr := setup.ProvisionModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)

	require.NoError(t, provisionModelFisrtTimeErr)
	require.Error(t, provisionModelSecondTimeErr)
}

func TestHandler_ProvisionModel_AlreadyCertified(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := ProvisionModelSetup(t)

	setup.CertifyModel(t, vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	_, provisionModelErr := setup.ProvisionModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)

	require.Error(t, provisionModelErr)
	require.True(t, types.ErrAlreadyCertified.Is(provisionModelErr))
}

func TestHandler_ProvisionModel_AlreadyRevoked(t *testing.T) {
	setup, vid, pid, softwareVersion, softwareVersionString, certificationType := ProvisionModelSetup(t)

	setup.RevokeModel(t, vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)
	_, provisionModelErr := setup.ProvisionModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)

	require.Error(t, provisionModelErr)
	require.True(t, types.ErrAlreadyRevoked.Is(provisionModelErr))
}

func TestHandler_ProvisionModel_MoreThanOneModel(t *testing.T) {
	setup, _, _, _, _, certificationType := ProvisionModelSetup(t)
	modelsQuantity := 5

	for i := 1; i < modelsQuantity; i++ {
		vid := int32(i)
		pid := int32(i)
		softwareVersion := uint32(i)
		softwareVersionString := fmt.Sprint(i)

		setup.AddModelVersion(vid, pid, softwareVersion, softwareVersionString)

		provisionModelMsg, provisionModelErr := setup.ProvisionModel(vid, pid, softwareVersion, softwareVersionString, certificationType, setup.CertificationCenter)

		require.NoError(t, provisionModelErr)
		setup.CheckComplianceInfoEqualsProvisionModelMsgData(t, provisionModelMsg)
		setup.CheckProvisionalModelAddedToCorrectStore(t, provisionModelMsg)
	}
}
