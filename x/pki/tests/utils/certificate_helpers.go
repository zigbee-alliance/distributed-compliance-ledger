package utils

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func ProposeAndApproveRootCertificateByOptions(
	setup *TestSetup,
	ownerTrustee sdk.AccAddress,
	certificate *RootCertOptions,
) {
	// ensure that `ownerTrustee` is trustee to eventually have enough approvals
	require.True(setup.T, setup.DclauthKeeper.HasRole(setup.Ctx, ownerTrustee, types.RootCertificateApprovalRole))

	// propose x509 root certificate by `ownerTrustee`
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(ownerTrustee.String(), certificate.PemCert, testconstants.Info, certificate.Vid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(setup.T, err)

	// approve x509 root certificate by another trustee
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), certificate.Subject, certificate.SubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(setup.T, err)
}

func ProposeAndApproveRootCertificate(
	setup *TestSetup,
	ownerTrustee sdk.AccAddress,
	certificate types.Certificate,
) {
	// ensure that `ownerTrustee` is trustee to eventually have enough approvals
	require.True(setup.T, setup.DclauthKeeper.HasRole(setup.Ctx, ownerTrustee, types.RootCertificateApprovalRole))

	// propose x509 root certificate by `ownerTrustee`
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(ownerTrustee.String(), certificate.PemCert, testconstants.Info, certificate.Vid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(setup.T, err)

	// approve x509 root certificate by another trustee
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), certificate.Subject, certificate.SubjectKeyId, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(setup.T, err)
}

func ProposeAndApproveCertificateRevocation(
	setup *TestSetup,
	subject string,
	subjectKeyID string,
	serialNumber string,
) {
	// revoke certificate
	revokeX509Cert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(),
		subject,
		subjectKeyID,
		serialNumber,
		false,
		testconstants.Info)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(setup.T, err)

	aprRevokeX509Cert := types.NewMsgApproveRevokeX509RootCert(
		setup.Trustee2.String(),
		subject,
		subjectKeyID,
		serialNumber,
		testconstants.Info)
	_, err = setup.Handler(setup.Ctx, aprRevokeX509Cert)
	require.NoError(setup.T, err)
}

func AddMokedDaCertificate(
	setup *TestSetup,
	certificate types.Certificate,
) {
	setup.Keeper.SetUniqueCertificate(setup.Ctx, UniqueCertificate(certificate.Issuer, certificate.SerialNumber))
	setup.Keeper.StoreDaCertificate(setup.Ctx, certificate, certificate.IsRoot)
}

func AddMokedNocCertificate(
	setup *TestSetup,
	certificate types.Certificate,
) {
	setup.Keeper.SetUniqueCertificate(setup.Ctx, UniqueCertificate(certificate.Issuer, certificate.SerialNumber))
	setup.Keeper.StoreNocCertificate(setup.Ctx, certificate, certificate.IsRoot)
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

func ProposeDaRootCertificate(
	setup *TestSetup,
	certificate types.Certificate,
) *types.MsgProposeAddX509RootCert {
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(
		certificate.Owner,
		certificate.PemCert,
		testconstants.Info,
		testconstants.Vid,
		testconstants.CertSchemaVersion,
	)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(setup.T, err)

	return proposeAddX509RootCert
}

func ApproveDaRootCertificate(
	setup *TestSetup,
	address sdk.AccAddress,
	subject string,
	subjectKeyID string,
) *types.MsgApproveAddX509RootCert {
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		address.String(),
		subject,
		subjectKeyID,
		testconstants.Info,
	)
	_, err := setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(setup.T, err)

	return approveAddX509RootCert
}

func RejectDaRootCertificate(
	setup *TestSetup,
	address sdk.AccAddress,
	subject string,
	subjectKeyID string,
) *types.MsgRejectAddX509RootCert {
	rejectAddX509RootCert := types.NewMsgRejectAddX509RootCert(
		address.String(),
		subject,
		subjectKeyID,
		testconstants.Info,
	)
	_, err := setup.Handler(setup.Ctx, rejectAddX509RootCert)
	require.NoError(setup.T, err)

	return rejectAddX509RootCert
}

func AddDaIntermediateCertificate(
	setup *TestSetup,
	certificate types.Certificate,
) *types.MsgAddX509Cert {
	addX509Cert := types.NewMsgAddX509Cert(certificate.Owner, certificate.PemCert, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(setup.T, err)

	return addX509Cert
}

func ProposeRevokeDaRootCertificate(
	setup *TestSetup,
	address sdk.AccAddress,
	subject string,
	subjectKeyID string,
	serialNumber string,
	revokeChild bool,
) *types.MsgProposeRevokeX509RootCert {
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		address.String(),
		subject,
		subjectKeyID,
		serialNumber,
		revokeChild,
		testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(setup.T, err)

	return proposeRevokeX509RootCert
}

func ApproveRevokeDaRootCertificate(
	setup *TestSetup,
	address sdk.AccAddress,
	subject string,
	subjectKeyID string,
	serialNumber string,
) *types.MsgApproveRevokeX509RootCert {
	approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
		address.String(),
		subject,
		subjectKeyID,
		serialNumber,
		testconstants.Info)
	_, err := setup.Handler(setup.Ctx, approveRevokeX509RootCert)
	require.NoError(setup.T, err)

	return approveRevokeX509RootCert
}

func RemoveDaIntermediateCertificate(
	setup *TestSetup,
	address sdk.AccAddress,
	subject string,
	subjectKeyID string,
	serialNumber string,
) *types.MsgRemoveX509Cert {
	removeCert := types.NewMsgRemoveX509Cert(
		address.String(),
		subject,
		subjectKeyID,
		serialNumber,
	)
	_, err := setup.Handler(setup.Ctx, removeCert)
	require.NoError(setup.T, err)

	return removeCert
}

func RevokeDaIntermediateCertificate(
	setup *TestSetup,
	address sdk.AccAddress,
	subject string,
	subjectKeyID string,
	serialNumber string,
	revokedChild bool,
) *types.MsgRevokeX509Cert {
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		address.String(),
		subject,
		subjectKeyID,
		serialNumber,
		revokedChild,
		testconstants.Info,
	)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(setup.T, err)

	return revokeX509Cert
}

func AddNocRootCertificate(
	setup *TestSetup,
	certificate types.Certificate,
) *types.MsgAddNocX509RootCert {
	var isVidVerificationSigner = false
	if certificate.CertificateType == types.CertificateType_VIDSignerPKI {
		isVidVerificationSigner = true
	}

	addNocX509RootCert := types.NewMsgAddNocX509RootCert(certificate.Owner, certificate.PemCert, testconstants.CertSchemaVersion, isVidVerificationSigner)
	_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
	require.NoError(setup.T, err)

	return addNocX509RootCert
}

func AddNocIntermediateCertificate(
	setup *TestSetup,
	certificate types.Certificate,
) *types.MsgAddNocX509IcaCert {
	var isVidVerificationSigner = false
	if certificate.CertificateType == types.CertificateType_VIDSignerPKI {
		isVidVerificationSigner = true
	}

	nocX509Cert := types.NewMsgAddNocX509IcaCert(certificate.Owner, certificate.PemCert, testconstants.CertSchemaVersion, isVidVerificationSigner)
	_, err := setup.Handler(setup.Ctx, nocX509Cert)
	require.NoError(setup.T, err)

	return nocX509Cert
}

func RemoveNocIntermediateCertificate(
	setup *TestSetup,
	address sdk.AccAddress,
	subject string,
	subjectKeyID string,
	serialNumber string,
) *types.MsgRemoveNocX509IcaCert {
	removeIcaCert := types.NewMsgRemoveNocX509IcaCert(
		address.String(),
		subject,
		subjectKeyID,
		serialNumber,
	)
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(setup.T, err)

	return removeIcaCert
}

func RevokeNocIntermediateCertificate(
	setup *TestSetup,
	address sdk.AccAddress,
	subject string,
	subjectKeyID string,
	serialNumber string,
	revokedChild bool,
) *types.MsgRevokeNocX509IcaCert {
	revokeX509Cert := types.NewMsgRevokeNocX509IcaCert(
		address.String(),
		subject,
		subjectKeyID,
		serialNumber,
		testconstants.Info,
		revokedChild,
	)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(setup.T, err)

	return revokeX509Cert
}

func RemoveNocRootCertificate(
	setup *TestSetup,
	address sdk.AccAddress,
	subject string,
	subjectKeyID string,
	serialNumber string,
) *types.MsgRemoveNocX509RootCert {
	removeRootCert := types.NewMsgRemoveNocX509RootCert(
		address.String(),
		subject,
		subjectKeyID,
		serialNumber,
	)
	_, err := setup.Handler(setup.Ctx, removeRootCert)
	require.NoError(setup.T, err)

	return removeRootCert
}

func RevokeNocRootCertificate(
	setup *TestSetup,
	address sdk.AccAddress,
	subject string,
	subjectKeyID string,
	serialNumber string,
	revokedChild bool,
) *types.MsgRevokeNocX509RootCert {
	revokeX509Cert := types.NewMsgRevokeNocX509RootCert(
		address.String(),
		subject,
		subjectKeyID,
		serialNumber,
		testconstants.Info,
		revokedChild,
	)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(setup.T, err)

	return revokeX509Cert
}

func AssignCertificateVid(
	setup *TestSetup,
	address sdk.AccAddress,
	subject string,
	subjectKeyID string,
	vid int32,
) *types.MsgAssignVid {
	assignVid := types.NewMsgAssignVid(
		address.String(),
		subject,
		subjectKeyID,
		vid,
	)
	_, err := setup.Handler(setup.Ctx, assignVid)
	require.NoError(setup.T, err)

	return assignVid
}
