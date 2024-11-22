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

func UniqueCertificate(issuer string, serialNumber string) types.UniqueCertificate {
	return types.UniqueCertificate{
		Issuer:       issuer,
		SerialNumber: serialNumber,
		Present:      true,
	}
}

func CertificateIdentifier(subject string, subjectKeyID string) types.CertificateIdentifier {
	return types.CertificateIdentifier{
		Subject:      subject,
		SubjectKeyId: subjectKeyID,
	}
}

func AddDaIntermediateCertificate(setup *TestSetup, address sdk.AccAddress, pemCert string) {
	addX509Cert := types.NewMsgAddX509Cert(address.String(), pemCert, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(setup.T, err)
}

func AddNocRootCertificate(setup *TestSetup, address sdk.AccAddress, pemCert string) {
	// add the new NOC root certificate
	addNocX509RootCert := types.NewMsgAddNocX509RootCert(address.String(), pemCert, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
	require.NoError(setup.T, err)
}

func AddNocIntermediateCertificate(setup *TestSetup, address sdk.AccAddress, pemCert string) {
	// add the new NOC root certificate
	nocX509Cert := types.NewMsgAddNocX509IcaCert(address.String(), pemCert, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, nocX509Cert)
	require.NoError(setup.T, err)
}
