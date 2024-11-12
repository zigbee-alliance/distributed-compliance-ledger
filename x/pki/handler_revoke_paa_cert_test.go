package pki

import (
	"math"
	"math/rand"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Main

func TestHandler_ProposeRevokeX509RootCert_ByTrusteeOwner(t *testing.T) {
	setup := Setup(t)

	// propose x509 root certificate by `setup.Trustee` and approve by another trustee
	rootCertOptions := createTestRootCertOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// propose revocation of x509 root certificate by `setup.Trustee`
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.RootSerialNumber, false, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	// query and check proposed certificate revocation
	proposedRevocation, _ := queryProposedCertificateRevocation(setup, testconstants.RootSerialNumber)
	require.Equal(t, testconstants.RootSubject, proposedRevocation.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, proposedRevocation.SubjectKeyId)
	require.True(t, proposedRevocation.HasRevocationFrom(setup.Trustee1.String()))

	// check that approved certificate still exists
	certificate, _ := querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NotNil(t, certificate)

	// check that revoked certificate does not exist
	_, err = queryRevokedCertificates(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificate key stays registered
	require.True(t,
		setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.RootIssuer, testconstants.RootSerialNumber))
}

func TestHandler_TwoThirdApprovalsNeededForRevokingRootCertification(t *testing.T) {
	setup := Setup(t)

	// propose x509 root certificate by account without trustee role
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// Approve the certificate from Trustee2
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	approvedCertificate, _ := querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, testconstants.RootIssuer, approvedCertificate.Subject)
	require.Equal(t, testconstants.RootSerialNumber, approvedCertificate.SerialNumber)
	require.True(t, approvedCertificate.IsRoot)
	require.True(t, approvedCertificate.HasApprovalFrom(setup.Trustee1.String()))

	// Create an array of trustee account from 1 to 50
	trusteeAccounts := make([]sdk.AccAddress, 50)
	for i := 0; i < 50; i++ {
		trusteeAccounts[i] = GenerateAccAddress()
	}

	totalAdditionalTrustees := rand.Intn(50)
	for i := 0; i < totalAdditionalTrustees; i++ {
		setup.AddAccount(trusteeAccounts[i], []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)
	}

	// We have 3 Trustees in test setup.
	twoThirds := int(math.Ceil(types.RootCertificateApprovalsPercent * float64(3+totalAdditionalTrustees)))

	// Trustee1 proposes to revoke the certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.RootSerialNumber, false, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	// Until we hit 2/3 of the total number of Trustees, we should not be able to revoke the certificate
	// We start the counter from 2 as the proposer is a trustee as well
	for i := 1; i < twoThirds-1; i++ {
		// approve the revocation
		approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
			trusteeAccounts[i].String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.RootSerialNumber, testconstants.Info)
		_, err = setup.Handler(setup.Ctx, approveRevokeX509RootCert)
		require.NoError(t, err)

		// check that the certificate is still not revoked
		approvedCertificate, _ := querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
		require.Equal(t, testconstants.RootIssuer, approvedCertificate.Subject)
		require.Equal(t, testconstants.RootSerialNumber, approvedCertificate.SerialNumber)
		require.True(t, approvedCertificate.IsRoot)
	}

	// One more revoke will revoke the certificate
	approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
		setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.RootSerialNumber, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveRevokeX509RootCert)
	require.NoError(t, err)

	// Check that the certificate is revoked
	ensureDaPaaCertificateDoesNotExist(
		t,
		setup,
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.RootIssuer,
		testconstants.RootSerialNumber,
		true)

	// Check that the certificate is revoked
	revokedCertificate, err := querySingleRevokedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, testconstants.RootIssuer, revokedCertificate.Subject)
	require.Equal(t, testconstants.RootSerialNumber, revokedCertificate.SerialNumber)
	require.True(t, revokedCertificate.IsRoot)
	// Make sure all the approvals are present
	for i := 1; i < twoThirds-1; i++ {
		require.Equal(t, revokedCertificate.HasApprovalFrom(trusteeAccounts[i].String()), true)
	}
	require.Equal(t, revokedCertificate.HasApprovalFrom(setup.Trustee1.String()), true)
	require.Equal(t, revokedCertificate.HasApprovalFrom(setup.Trustee2.String()), true)
}

func TestHandler_ProposeRevokeX509RootCert_ByTrusteeNotOwner(t *testing.T) {
	setup := Setup(t)

	// propose x509 root certificate by `setup.Trustee` and approve by another trustee
	rootCertOptions := createTestRootCertOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// add another trustee
	anotherTrustee := GenerateAccAddress()
	setup.AddAccount(anotherTrustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)

	// propose revocation of x509 root certificate by new trustee
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		anotherTrustee.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.RootSerialNumber, false, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	// query and check proposed certificate revocation
	proposedRevocation, _ := queryProposedCertificateRevocation(setup, testconstants.RootSerialNumber)
	require.Equal(t, testconstants.RootSubject, proposedRevocation.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, proposedRevocation.SubjectKeyId)
	require.True(t, proposedRevocation.HasRevocationFrom(anotherTrustee.String()))

	// check that approved certificate still exists
	certificate, _ := querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NotNil(t, certificate)

	// check that revoked certificate does not exist
	_, err = queryRevokedCertificates(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificate key stays registered
	require.True(t,
		setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.RootIssuer, testconstants.RootSerialNumber))
}

//nolint:funlen
func TestHandler_ApproveRevokeX509RootCert_ForTree(t *testing.T) {
	setup := Setup(t)

	// add root x509 certificate
	rootCertOptions := createTestRootCertOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// Add vendor account
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add intermediate x509 certificate
	addDaPaiCertificate(setup, vendorAccAddress, testconstants.IntermediateCertPem)

	// add leaf x509 certificate
	addDaPaiCertificate(setup, vendorAccAddress, testconstants.LeafCertPem)

	// propose revocation of x509 root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, "", true, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	// approve
	approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
		setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, "", testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveRevokeX509RootCert)
	require.NoError(t, err)

	// check that root, intermediate and leaf certificates have been revoked
	allRevokedCertificates, _ := queryAllRevokedCertificates(setup)
	require.Equal(t, 3, len(allRevokedCertificates))
	require.Equal(t, testconstants.LeafSubject, allRevokedCertificates[0].Subject)
	require.Equal(t, testconstants.LeafSubjectKeyID, allRevokedCertificates[0].SubjectKeyId)
	require.Equal(t, 1, len(allRevokedCertificates[0].Certs))
	require.Equal(t, testconstants.LeafCertPem, allRevokedCertificates[0].Certs[0].PemCert)
	require.Equal(t, testconstants.RootSubject, allRevokedCertificates[1].Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, allRevokedCertificates[1].SubjectKeyId)
	require.Equal(t, 1, len(allRevokedCertificates[1].Certs))
	require.Equal(t, testconstants.RootCertPem, allRevokedCertificates[1].Certs[0].PemCert)
	require.Equal(t, testconstants.IntermediateSubject, allRevokedCertificates[2].Subject)
	require.Equal(t, testconstants.IntermediateSubjectKeyID, allRevokedCertificates[2].SubjectKeyId)
	require.Equal(t, 1, len(allRevokedCertificates[2].Certs))
	require.Equal(t, testconstants.IntermediateCertPem, allRevokedCertificates[2].Certs[0].PemCert)

	// check that approved certs list is empty
	allApprovedCertificates, err := queryAllApprovedCertificates(setup)
	require.NoError(t, err)
	require.Equal(t, 0, len(allApprovedCertificates))

	// check that no proposed certificate revocations exist
	allProposedCertificateRevocations, err := queryAllProposedCertificateRevocations(setup)
	require.NoError(t, err)
	require.Equal(t, 0, len(allProposedCertificateRevocations))

	// check that no child certificate identifiers are registered for revoked root certificate
	rootCertChildren, err := queryChildCertificates(
		setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
	require.Nil(t, rootCertChildren)

	// check that no child certificate identifiers are registered for revoked intermediate certificate
	intermediateCertChildren, err := queryChildCertificates(
		setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
	require.Nil(t, intermediateCertChildren)

	// check that no child certificate identifiers are registered for revoked leaf certificate
	leafCertChildren, err := queryChildCertificates(
		setup, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
	require.Nil(t, leafCertChildren)

	// check that root certificate does not exist
	ensureDaPaaCertificateDoesNotExist(
		t,
		setup,
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.RootSubject,
		testconstants.RootSerialNumber,
		true)

	// check that intermediate certificate does not exist
	ensureDaPaiCertificateDoesNotExist(
		t,
		setup,
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		testconstants.IntermediateIssuer,
		testconstants.IntermediateSerialNumber,
		true,
		false)

	// check that intermediate certificate does not exist
	ensureDaPaiCertificateDoesNotExist(
		t,
		setup,
		testconstants.LeafSubject,
		testconstants.LeafSubjectKeyID,
		testconstants.LeafIssuer,
		testconstants.LeafSerialNumber,
		true,
		false)
}

func TestHandler_RevokeX509RootCertsBySubjectKeyId(t *testing.T) {
	setup := Setup(t)

	// add root certificates
	rootCertOptions := &rootCertOptions{
		pemCert:      testconstants.PAACertWithSameSubjectID1,
		subject:      testconstants.PAACertWithSameSubjectID1Subject,
		subjectKeyID: testconstants.PAACertWithSameSubjectIDSubjectID,
		info:         testconstants.Info,
		vid:          testconstants.Vid,
	}
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)
	rootCertOptions.pemCert = testconstants.PAACertWithSameSubjectID2
	rootCertOptions.subject = testconstants.PAACertWithSameSubjectID2Subject
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// revoke certificate
	revokeX509Cert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(), testconstants.PAACertWithSameSubjectID1Subject, testconstants.PAACertWithSameSubjectIDSubjectID, "", false, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(t, err)

	aprRevokeX509Cert := types.NewMsgApproveRevokeX509RootCert(
		setup.Trustee2.String(), testconstants.PAACertWithSameSubjectID1Subject, testconstants.PAACertWithSameSubjectIDSubjectID, "", testconstants.Info)
	_, err = setup.Handler(setup.Ctx, aprRevokeX509Cert)
	require.NoError(t, err)

	// check that root certificate has been revoked
	approvedCertificates, _ := queryApprovedCertificates(setup, testconstants.PAACertWithSameSubjectID2Subject, testconstants.PAACertWithSameSubjectIDSubjectID)
	require.Equal(t, 1, len(approvedCertificates.Certs))
	require.Equal(t, testconstants.PAACertWithSameSubjectID2Subject, approvedCertificates.Certs[0].Subject)
	require.Equal(t, testconstants.PAACertWithSameSubjectIDSubjectID, approvedCertificates.SubjectKeyId)

	certsBySubjectKeyID, _ := queryAllApprovedCertificatesBySubjectKeyID(setup, testconstants.PAACertWithSameSubjectIDSubjectID)
	require.Equal(t, 1, len(certsBySubjectKeyID))
	require.Equal(t, 1, len(certsBySubjectKeyID[0].Certs))
	require.Equal(t, testconstants.PAACertWithSameSubjectIDSubjectID, certsBySubjectKeyID[0].SubjectKeyId)
	require.Equal(t, testconstants.PAACertWithSameSubjectID2Subject, certsBySubjectKeyID[0].Certs[0].Subject)

	// check that no proposed certificate revocations have been created
	allProposedCertificateRevocations, _ := queryAllProposedCertificateRevocations(setup)
	require.NoError(t, err)
	require.Equal(t, 0, len(allProposedCertificateRevocations))
}

// Extra cases

func TestHandler_ApproveRevokeX509RootCert_ForNotEnoughApprovals(t *testing.T) {
	setup := Setup(t)

	// propose and approve x509 root certificate
	rootCertOptions := createTestRootCertOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// Add 1 more trustee (this will bring the total trustee's to 4)
	anotherTrustee := GenerateAccAddress()
	setup.AddAccount(anotherTrustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)

	// propose revocation of x509 root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.RootSerialNumber, false, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	// approve
	approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
		setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.RootSerialNumber, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveRevokeX509RootCert)
	require.NoError(t, err)

	// query and check proposed certificate revocation
	proposedRevocation, _ := queryProposedCertificateRevocation(setup, testconstants.RootSerialNumber)
	require.Equal(t, testconstants.RootSubject, proposedRevocation.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, proposedRevocation.SubjectKeyId)
	require.True(t, proposedRevocation.HasRevocationFrom(setup.Trustee1.String()))
	require.True(t, proposedRevocation.HasRevocationFrom(setup.Trustee2.String()))

	// check that approved certificate still exists
	certificate, _ := querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NotNil(t, certificate)

	// check that revoked certificate does not exist
	_, err = queryRevokedCertificates(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificate key stays registered
	require.True(t,
		setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.RootIssuer, testconstants.RootSerialNumber))
}

func TestHandler_ApproveRevokeX509RootCert_ForEnoughApprovals(t *testing.T) {
	setup := Setup(t)

	// propose and approve x509 root certificate
	rootCertOptions := createTestRootCertOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// propose revocation of x509 root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.RootSerialNumber, false, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	// get certificate for further comparison
	certificateBeforeRevocation, _ := querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NotNil(t, certificateBeforeRevocation)

	// approve
	approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
		setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.RootSerialNumber, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveRevokeX509RootCert)
	require.NoError(t, err)

	// check that proposed certificate revocation does not exist anymore
	_, err = queryProposedCertificateRevocation(setup, testconstants.RootSerialNumber)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that approved certificate does not exist anymore
	_, err = queryApprovedCertificates(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// query and check revoked certificate
	revokedCertificate, _ := querySingleRevokedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, certificateBeforeRevocation, revokedCertificate)

	// check that unique certificate key stays registered
	require.True(t,
		setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.RootIssuer, testconstants.RootSerialNumber))
}

func TestHandler_ApproveRevokeX509RootCert_BySerialNumber(t *testing.T) {
	setup := Setup(t)

	rootCertOpt := &rootCertOptions{
		pemCert:      testconstants.RootCertWithSameSubjectAndSKID1,
		subject:      testconstants.RootCertWithSameSubjectAndSKIDSubject,
		subjectKeyID: testconstants.RootCertWithSameSubjectAndSKIDSubjectKeyID,
		info:         testconstants.Info,
		vid:          testconstants.Vid,
	}
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOpt)
	rootCertOpt.pemCert = testconstants.RootCertWithSameSubjectAndSKID2
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOpt)
	rootSubject := rootCertOpt.subject
	rootSubjectKeyID := rootCertOpt.subjectKeyID

	// Add vendor account
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// Add an intermediate certificate
	addIntermediateX509Cert := types.NewMsgAddX509Cert(vendorAccAddress.String(), testconstants.IntermediateWithSameSubjectAndSKID1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addIntermediateX509Cert)
	require.NoError(t, err)

	intermediateSubject := testconstants.IntermediateCertWithSameSubjectAndSKIDSubject
	intermediateSubjectKeyID := testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID

	// get certificates for further comparison
	certsBeforeRevocation := setup.Keeper.GetAllApprovedCertificates(setup.Ctx)
	require.NotNil(t, certsBeforeRevocation)
	require.Equal(t, 2, len(certsBeforeRevocation))
	require.Equal(t, 3, len(certsBeforeRevocation[0].Certs)+len(certsBeforeRevocation[1].Certs))

	// propose revocation of root certificate with serial number "1"
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(), rootSubject, rootSubjectKeyID, "1", false, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	// approve
	approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
		setup.Trustee2.String(), rootSubject, rootSubjectKeyID, "1", testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveRevokeX509RootCert)
	require.NoError(t, err)

	// check that proposed certificate revocation does not exist anymore
	_, err = queryProposedCertificateRevocation(setup, "1")
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that only two approved certificates exists(root and child certificates)
	rootCerts, _ := queryApprovedRootCertificates(setup, rootSubject, rootSubjectKeyID)
	require.Equal(t, 1, len(rootCerts))
	require.Equal(t, "2", rootCerts[0].SerialNumber)
	certificates, err := queryApprovedCertificates(setup, intermediateSubject, intermediateSubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 1, len(certificates.Certs))

	// query and check revoked certificate
	revokedCertificate, _ := querySingleRevokedCertificate(setup, rootSubject, rootSubjectKeyID)
	require.NotNil(t, revokedCertificate)
	require.Equal(t, "1", revokedCertificate.SerialNumber)

	// propose revocation of root certificate with serial number "2"
	proposeRevokeX509RootCert = types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(), rootSubject, rootSubjectKeyID, "2", true, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	// approve
	approveRevokeX509RootCert = types.NewMsgApproveRevokeX509RootCert(
		setup.Trustee2.String(), rootSubject, rootSubjectKeyID, "2", testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveRevokeX509RootCert)
	require.NoError(t, err)

	// check that proposed certificate revocation does not exist anymore
	_, err = queryProposedCertificateRevocation(setup, "2")
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that approved certificates does not exist anymore
	certsAfterRevocation := setup.Keeper.GetAllApprovedCertificates(setup.Ctx)
	require.Equal(t, 0, len(certsAfterRevocation))
	certsAfterRevocationBySubjectID := setup.Keeper.GetAllApprovedCertificatesBySubjectKeyID(setup.Ctx)
	require.Equal(t, 0, len(certsAfterRevocationBySubjectID))

	// query all revoked certificates
	allRevokedCerts, _ := queryAllRevokedCertificates(setup)
	require.Equal(t, 2, len(allRevokedCerts))

	// query and check revoked root certificates
	revokedCerts, _ := queryRevokedCertificates(setup, rootSubject, rootSubjectKeyID)
	require.Equal(t, 2, len(revokedCerts.Certs))
	require.Equal(t, rootSubject, revokedCerts.Subject)
	require.Equal(t, rootSubjectKeyID, revokedCerts.SubjectKeyId)
	// query and check revoked intermediate certificate
	revokedCerts, _ = queryRevokedCertificates(setup, intermediateSubject, intermediateSubjectKeyID)
	require.Equal(t, 1, len(revokedCerts.Certs))
	require.Equal(t, intermediateSubject, revokedCerts.Subject)
	require.Equal(t, intermediateSubjectKeyID, revokedCerts.SubjectKeyId)
}

// Error cases

func TestHandler_ProposeRevokeX509RootCert_ByNotTrustee(t *testing.T) {
	setup := Setup(t)

	// propose and approve x509 root certificate
	rootCertOptions := createTestRootCertOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.Vendor,
		dclauthtypes.CertificationCenter,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, 1)

		// propose revocation of x509 root certificate
		proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
			accAddress.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.RootSerialNumber, false, testconstants.Info)
		_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
	}
}

func TestHandler_ProposeRevokeX509RootCert_CertificateDoesNotExist(t *testing.T) {
	setup := Setup(t)

	// propose revocation of not existing certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.RootSerialNumber, false, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_ProposeRevokeX509RootCert_CertificateDoesNotExistBySerialNumber(t *testing.T) {
	setup := Setup(t)
	// propose and approve x509 root certificate
	rootCertOptions := createTestRootCertOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// revoke x509 certificate
	revokeX509Cert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(),
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		"invalid",
		false,
		testconstants.Info,
	)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_ProposeRevokeX509RootCert_ForProposedCertificate(t *testing.T) {
	setup := Setup(t)

	// propose x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// check that proposed certificate is present
	proposedCertificate, _ := queryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NotNil(t, proposedCertificate)

	// propose revocation of proposed root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.RootSerialNumber, false, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_ProposeRevokeX509RootCert_ProposedRevocationAlreadyExists(t *testing.T) {
	setup := Setup(t)

	// propose and approve x509 root certificate
	rootCertOptions := createTestRootCertOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// propose revocation of x509 root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.RootSerialNumber, false, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	// store another trustee
	anotherTrustee := GenerateAccAddress()
	setup.AddAccount(anotherTrustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)

	// propose revocation of the same x509 root certificate again
	proposeRevokeX509RootCert = types.NewMsgProposeRevokeX509RootCert(
		anotherTrustee.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.RootSerialNumber, false, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrProposedCertificateRevocationAlreadyExists.Is(err))
}

func TestHandler_ProposeRevokeX509RootCert_ForNonRootCertificate(t *testing.T) {
	setup := Setup(t)

	// store x509 root certificate
	rootCertificate := rootCertificate(setup.Trustee1)
	setup.Keeper.AddAllCertificate(setup.Ctx, rootCertificate)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// Add vendor account
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// store x509 intermediate certificate
	addX509Cert := types.NewMsgAddX509Cert(vendorAccAddress.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// propose revocation of x509 intermediate certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(), testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID, testconstants.RootSerialNumber, false, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrInappropriateCertificateType.Is(err))
}

func TestHandler_ApproveRevokeX509RootCert_ByNotTrustee(t *testing.T) {
	setup := Setup(t)

	// propose and approve x509 root certificate
	rootCertOptions := createTestRootCertOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// propose revocation of x509 root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.RootSerialNumber, false, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.Vendor,
		dclauthtypes.CertificationCenter,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, 1)

		// approve
		approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
			accAddress.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.RootSerialNumber, testconstants.Info)
		_, err = setup.Handler(setup.Ctx, approveRevokeX509RootCert)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
	}
}

func TestHandler_ApproveRevokeX509RootCert_ProposedRevocationDoesNotExist(t *testing.T) {
	setup := Setup(t)

	// propose and approve x509 root certificate
	rootCertOptions := createTestRootCertOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// approve revocation of x509 root certificate
	approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.RootSerialNumber, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, approveRevokeX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrProposedCertificateRevocationDoesNotExist.Is(err))
}

func TestHandler_ApproveRevokeX509RootCert_Twice(t *testing.T) {
	setup := Setup(t)

	// propose and approve x509 root certificate
	rootCertOptions := createTestRootCertOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// propose revocation of x509 root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.RootSerialNumber, false, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	// approve revocation by the same trustee
	approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.RootSerialNumber, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveRevokeX509RootCert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_RevocationPointsByIssuerSubjectKeyID(t *testing.T) {
	setup := Setup(t)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 65521)

	// propose x509 root certificate by account Trustee1
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.PAACertWithNumericVid, testconstants.Info, testconstants.Vid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.PAACertWithNumericVidSubject, testconstants.PAACertWithNumericVidSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	revocationPointBySubjectKeyID, isFound := setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
	require.False(t, isFound)
	require.Equal(t, len(revocationPointBySubjectKeyID.Points), 0)

	addPkiRevocationDistributionPoint := types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  testconstants.PAACertWithNumericVidVid,
		IsPAA:                true,
		Pid:                  8,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                "label",
		DataURL:              testconstants.DataURL + "/1",
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	revocationPointBySubjectKeyID, isFound = setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	require.Equal(t, len(revocationPointBySubjectKeyID.Points), 1)

	addPkiRevocationDistributionPoint = types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  testconstants.PAACertWithNumericVidVid,
		IsPAA:                true,
		Pid:                  8,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                "label1",
		DataURL:              testconstants.DataURL + "/2",
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	revocationPointBySubjectKeyID, isFound = setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	require.Equal(t, len(revocationPointBySubjectKeyID.Points), 2)

	dataURLNew := testconstants.DataURL + "/new"
	updatePkiRevocationDistributionPoint := types.MsgUpdatePkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  testconstants.PAACertWithNumericVidVid,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                "label",
		DataURL:              dataURLNew,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
	}
	_, err = setup.Handler(setup.Ctx, &updatePkiRevocationDistributionPoint)
	require.NoError(t, err)

	revocationPointBySubjectKeyID, isFound = setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	require.Equal(t, len(revocationPointBySubjectKeyID.Points), 2)
	require.Equal(t, revocationPointBySubjectKeyID.Points[0].CrlSignerCertificate, updatePkiRevocationDistributionPoint.CrlSignerCertificate)
	require.Equal(t, revocationPointBySubjectKeyID.Points[0].DataURL, updatePkiRevocationDistributionPoint.DataURL)

	deletePkiRevocationDistributionPoint := types.MsgDeletePkiRevocationDistributionPoint{
		Signer:             vendorAcc.String(),
		Vid:                65521,
		Label:              "label",
		IssuerSubjectKeyID: testconstants.SubjectKeyIDWithoutColons,
	}
	_, err = setup.Handler(setup.Ctx, &deletePkiRevocationDistributionPoint)
	require.NoError(t, err)

	revocationPointBySubjectKeyID, isFound = setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	require.Equal(t, len(revocationPointBySubjectKeyID.Points), 1)
}

func TestHandler_AddRevocationPointForSameCertificateWithDifferentWhitespaces(t *testing.T) {
	setup := Setup(t)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 65521)

	// propose x509 root certificate by account Trustee1
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.PAACertWithNumericVid, testconstants.Info, testconstants.Vid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.PAACertWithNumericVidSubject, testconstants.PAACertWithNumericVidSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	addPkiRevocationDistributionPoint := types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  testconstants.PAACertWithNumericVidVid,
		IsPAA:                true,
		Pid:                  8,
		CrlSignerCertificate: testconstants.PAACertWithNumericVidDifferentWhitespaces,
		Label:                "label",
		DataURL:              testconstants.DataURL + "/1",
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	revocationPointBySubjectKeyID, isFound := setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	require.Equal(t, len(revocationPointBySubjectKeyID.Points), 1)
	require.Equal(t, revocationPointBySubjectKeyID.Points[0].CrlSignerCertificate, addPkiRevocationDistributionPoint.CrlSignerCertificate)
}

func TestHandler_UpdateRevocationPointForSameCertificateWithDifferentWhitespaces(t *testing.T) {
	setup := Setup(t)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 65521)

	// propose x509 root certificate by account Trustee1
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.PAACertWithNumericVid, testconstants.Info, testconstants.Vid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.PAACertWithNumericVidSubject, testconstants.PAACertWithNumericVidSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	addPkiRevocationDistributionPoint := types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  testconstants.PAACertWithNumericVidVid,
		IsPAA:                true,
		Pid:                  8,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                "label",
		DataURL:              testconstants.DataURL + "/1",
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	revocationPointBySubjectKeyID, isFound := setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	require.Equal(t, len(revocationPointBySubjectKeyID.Points), 1)

	dataURLNew := testconstants.DataURL + "/new"
	updatePkiRevocationDistributionPoint := types.MsgUpdatePkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  testconstants.PAACertWithNumericVidVid,
		CrlSignerCertificate: testconstants.PAACertWithNumericVidDifferentWhitespaces,
		Label:                "label",
		DataURL:              dataURLNew,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
	}
	_, err = setup.Handler(setup.Ctx, &updatePkiRevocationDistributionPoint)
	require.NoError(t, err)

	revocationPointBySubjectKeyID, isFound = setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	require.Equal(t, revocationPointBySubjectKeyID.Points[0].CrlSignerCertificate, updatePkiRevocationDistributionPoint.CrlSignerCertificate)
	require.Equal(t, revocationPointBySubjectKeyID.Points[0].DataURL, updatePkiRevocationDistributionPoint.DataURL)
}
