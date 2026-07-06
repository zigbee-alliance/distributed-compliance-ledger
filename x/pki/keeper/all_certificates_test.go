package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/tests/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNCertificates(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.AllCertificates {
	items := make([]types.AllCertificates, n)
	for i := range items {
		items[i].Subject = strconv.Itoa(i)
		items[i].SubjectKeyId = strconv.Itoa(i)

		keeper.SetAllCertificates(ctx, items[i])
	}

	return items
}

func TestCertificatesGet(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNCertificates(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAllCertificates(ctx,
			item.Subject,
			item.SubjectKeyId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestCertificatesRemove(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNCertificates(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAllCertificates(ctx,
			item.Subject,
			item.SubjectKeyId,
		)
		_, found := keeper.GetAllCertificates(ctx,
			item.Subject,
			item.SubjectKeyId,
		)
		require.False(t, found)
	}
}

func TestCertificatesGetAll(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNCertificates(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAllCertificates(ctx)),
	)
}

func TestVerifyVVSCCertificate_ParentNotVidSigner(t *testing.T) {
	setup := utils.Setup(t)

	// parent exists at the ICA's issuer/AKI but is not a VID-signer cert, so the
	// VVSC walker skips it and no chain can be built.
	parent := types.Certificate{
		Subject:         testconstants.VvscIcaCert1Issuer,
		SubjectKeyId:    testconstants.VvscIcaCert1AuthorityKeyID,
		CertificateType: types.CertificateType_OperationalPKI,
		PemCert:         testconstants.VvscRootCert1,
	}
	setup.Keeper.AddAllCertificate(setup.Ctx, parent)

	addMsg := types.NewMsgAddNocX509IcaCert(
		setup.Vendor1.String(), testconstants.VvscIcaCert1, testconstants.CertSchemaVersion, true)
	_, err := setup.Handler(setup.Ctx, addMsg)
	require.ErrorIs(t, err, pkitypes.ErrVVSCChainVerificationFailed)
}

func TestVerifyVVSCCertificate_ParentUndecodable(t *testing.T) {
	setup := utils.Setup(t)

	// parent is a VID-signer entry but its stored PEM cannot be decoded.
	parent := types.Certificate{
		Subject:         testconstants.VvscIcaCert1Issuer,
		SubjectKeyId:    testconstants.VvscIcaCert1AuthorityKeyID,
		CertificateType: types.CertificateType_VIDSignerPKI,
		PemCert:         "not a certificate",
	}
	setup.Keeper.AddAllCertificate(setup.Ctx, parent)

	addMsg := types.NewMsgAddNocX509IcaCert(
		setup.Vendor1.String(), testconstants.VvscIcaCert1, testconstants.CertSchemaVersion, true)
	_, err := setup.Handler(setup.Ctx, addMsg)
	require.ErrorIs(t, err, pkitypes.ErrVVSCChainVerificationFailed)
}

func TestVerifyVVSCCertificate_ParentSignatureMismatch(t *testing.T) {
	setup := utils.Setup(t)

	// parent is a valid VID-signer cert but did not sign the ICA, so the VVSC
	// signature check fails and no chain can be built.
	parent := types.Certificate{
		Subject:         testconstants.VvscIcaCert1Issuer,
		SubjectKeyId:    testconstants.VvscIcaCert1AuthorityKeyID,
		CertificateType: types.CertificateType_VIDSignerPKI,
		PemCert:         testconstants.VvscRootCert2,
	}
	setup.Keeper.AddAllCertificate(setup.Ctx, parent)

	addMsg := types.NewMsgAddNocX509IcaCert(
		setup.Vendor1.String(), testconstants.VvscIcaCert1, testconstants.CertSchemaVersion, true)
	_, err := setup.Handler(setup.Ctx, addMsg)
	require.ErrorIs(t, err, pkitypes.ErrVVSCChainVerificationFailed)
}
