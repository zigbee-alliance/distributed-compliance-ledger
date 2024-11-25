package utils

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func ProposeAndApproveRootCertificate(setup *TestSetup, ownerTrustee sdk.AccAddress, options *RootCertOptions) {
	// ensure that `ownerTrustee` is trustee to eventually have enough approvals
	require.True(setup.T, setup.DclauthKeeper.HasRole(setup.Ctx, ownerTrustee, types.RootCertificateApprovalRole))

	// propose x509 root certificate by `ownerTrustee`
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(ownerTrustee.String(), options.PemCert, options.Info, options.Vid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(setup.T, err)

	// approve x509 root certificate by another trustee
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), options.Subject, options.SubjectKeyID, options.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(setup.T, err)

	// check that root certificate has been approved
	approvedCertificate, err := QueryApprovedCertificates(
		setup, options.Subject, options.SubjectKeyID)
	require.NoError(setup.T, err)
	require.NotNil(setup.T, approvedCertificate)
}

func AddMokedDaCertificate(
	setup *TestSetup,
	certificate types.Certificate,
	isRoot bool,
) {
	setup.Keeper.SetUniqueCertificate(setup.Ctx, UniqueCertificate(certificate.Subject, certificate.SerialNumber))
	setup.Keeper.StoreDaCertificate(setup.Ctx, certificate, isRoot)
}
