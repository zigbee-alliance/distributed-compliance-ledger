package tests

import (
	"math"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/tests/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Main

func TestHandler_ProposeAddDaRootCert(t *testing.T) {
	setup := utils.Setup(t)

	rootCertificate := utils.CreateTestRootCert()

	// propose DA root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(
		setup.Trustee1.String(),
		testconstants.RootCertPem,
		testconstants.Info,
		testconstants.Vid,
		testconstants.CertSchemaVersion,
	)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// Check indexes
	indexes := []utils.TestIndex{
		{Key: types.ProposedCertificateKeyPrefix, Exist: true},
		{Key: types.UniqueCertificateKeyPrefix, Exist: true},
		{Key: types.RejectedCertificateKeyPrefix, Exist: false},
		{Key: types.AllCertificatesKeyPrefix, Exist: false},
		{Key: types.AllCertificatesBySubjectKeyPrefix, Exist: false},
		{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Exist: false},
		{Key: types.ApprovedCertificatesKeyPrefix, Exist: false},
		{Key: types.ApprovedCertificatesBySubjectKeyPrefix, Exist: false},
		{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix, Exist: false},
		{Key: types.ApprovedRootCertificatesKeyPrefix, Exist: false},
	}
	resolvedCertificates := utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)

	// additional checks
	require.Equal(t, proposeAddX509RootCert.Cert, resolvedCertificates.ProposedCertificate.PemCert)
	require.True(t, resolvedCertificates.ProposedCertificate.HasApprovalFrom(proposeAddX509RootCert.Signer))
}

func TestHandler_AddDaRootCert(t *testing.T) {
	setup := utils.Setup(t)

	rootCertificate := utils.CreateTestRootCert()

	// propose add x509 root certificate by trustee
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(
		setup.Trustee1.String(),
		testconstants.RootCertPem,
		testconstants.Info,
		testconstants.Vid,
		testconstants.CertSchemaVersion,
	)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// approve by second trustee
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(),
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.Info,
	)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// Check indexes
	indexes := []utils.TestIndex{
		{Key: types.ProposedCertificateKeyPrefix, Exist: false},
		{Key: types.UniqueCertificateKeyPrefix, Exist: true},
		{Key: types.RejectedCertificateKeyPrefix, Exist: false},
		{Key: types.AllCertificatesKeyPrefix, Exist: true},
		{Key: types.AllCertificatesBySubjectKeyPrefix, Exist: true},
		{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Exist: true},
		{Key: types.ApprovedCertificatesKeyPrefix, Exist: true},
		{Key: types.ApprovedCertificatesBySubjectKeyPrefix, Exist: true},
		{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix, Exist: true},
		{Key: types.ApprovedRootCertificatesKeyPrefix, Exist: true},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)
}

func TestHandler_AddDaRootCert_TwoThirdApprovalsNeeded(t *testing.T) {
	setup := utils.Setup(t)

	rootCertificate := utils.CreateTestRootCert()

	// propose x509 root certificate by account without trustee role
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(
		setup.Trustee1.String(),
		testconstants.RootCertPem,
		testconstants.Info,
		testconstants.Vid,
		testconstants.CertSchemaVersion,
	)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// Create an array of trustee account from 1 to 50
	trusteeAccounts, totalAdditionalTrustees := setup.CreateNTrusteeAccounts()

	// We have 3 Trustees in test setup.
	twoThirds := int(math.Ceil(types.RootCertificateApprovalsPercent * float64(3+totalAdditionalTrustees)))

	// Until we hit 2/3 of the total number of Trustees, we should not be able to approve the certificate
	for i := 1; i < twoThirds-1; i++ {
		approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
			trusteeAccounts[i].String(),
			testconstants.RootSubject,
			testconstants.RootSubjectKeyID,
			testconstants.Info,
		)
		_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
		require.NoError(t, err)

		_, err = utils.QueryApprovedCertificates(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
		require.Error(t, err)
		require.Equal(t, codes.NotFound, status.Code(err))
	}

	// One more approval will move this to approved state from pending
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// Check indexes
	indexes := []utils.TestIndex{
		{Key: types.ProposedCertificateKeyPrefix, Exist: false},
		{Key: types.UniqueCertificateKeyPrefix, Exist: true},
		{Key: types.RejectedCertificateKeyPrefix, Exist: false},
		{Key: types.AllCertificatesKeyPrefix, Exist: true},
		{Key: types.AllCertificatesBySubjectKeyPrefix, Exist: true},
		{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Exist: true},
		{Key: types.ApprovedCertificatesKeyPrefix, Exist: true},
		{Key: types.ApprovedCertificatesBySubjectKeyPrefix, Exist: true},
		{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix, Exist: true},
		{Key: types.ApprovedRootCertificatesKeyPrefix, Exist: true},
	}
	resolvedCertificates := utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)

	// Additional check: Check all approvals are present
	for i := 1; i < twoThirds-1; i++ {
		require.Equal(t, resolvedCertificates.ApprovedCertificates.Certs[0].HasApprovalFrom(trusteeAccounts[i].String()), true)
	}
	require.Equal(t, resolvedCertificates.ApprovedCertificates.Certs[0].HasApprovalFrom(setup.Trustee1.String()), true)
	require.Equal(t, resolvedCertificates.ApprovedCertificates.Certs[0].HasApprovalFrom(setup.Trustee2.String()), true)
}

func TestHandler_AddDaRootCert_FourApprovalsAreNeeded_FiveTrustees(t *testing.T) {
	setup := utils.Setup(t)

	rootCertificate := utils.CreateTestRootCert()

	// we have 5 trustees: 1 approval comes from propose => we need 3 more approvals

	// store 4th trustee
	fourthTrustee := utils.GenerateAccAddress()
	setup.AddAccount(fourthTrustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)

	// store 5th trustee
	fifthTrustee := utils.GenerateAccAddress()
	setup.AddAccount(fifthTrustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)

	// propose x509 root certificate by account Trustee1
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(
		setup.Trustee1.String(),
		testconstants.RootCertPem,
		testconstants.Info,
		testconstants.Vid,
		testconstants.CertSchemaVersion,
	)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// approve x509 root certificate by account Trustee2
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(),
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.Info,
	)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// approve x509 root certificate by account Trustee3
	approveAddX509RootCert = types.NewMsgApproveAddX509RootCert(
		setup.Trustee3.String(),
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.Info,
	)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// reject x509 root certificate by account Trustee4
	rejectAddX509RootCert := types.NewMsgRejectAddX509RootCert(
		fourthTrustee.String(),
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.Info,
	)
	_, err = setup.Handler(setup.Ctx, rejectAddX509RootCert)
	require.NoError(t, err)

	// Check: ProposedCertificate - present because we haven't enough approvals
	indexes := []utils.TestIndex{
		{Key: types.ProposedCertificateKeyPrefix, Exist: true},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)

	// approve x509 root certificate by account Trustee5
	approveAddX509RootCert = types.NewMsgApproveAddX509RootCert(
		fifthTrustee.String(),
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.Info,
	)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// Check indexes
	indexes = []utils.TestIndex{
		{Key: types.ProposedCertificateKeyPrefix, Exist: false},
		{Key: types.UniqueCertificateKeyPrefix, Exist: true},
		{Key: types.RejectedCertificateKeyPrefix, Exist: false},
		{Key: types.AllCertificatesKeyPrefix, Exist: true},
		{Key: types.AllCertificatesBySubjectKeyPrefix, Exist: true},
		{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Exist: true},
		{Key: types.ApprovedCertificatesKeyPrefix, Exist: true},
		{Key: types.ApprovedCertificatesBySubjectKeyPrefix, Exist: true},
		{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix, Exist: true},
		{Key: types.ApprovedRootCertificatesKeyPrefix, Exist: true},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)
}

// Extra cases

func TestHandler_ProposeAddX509RootCert_ForDifferentSerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	testRootCertificate := utils.CreateTestRootCert()
	testRootCertificate.SerialNumber = utils.SerialNumber

	// store root certificate with different serial number
	rootCertificate := utils.RootCertificate(setup.Trustee1)
	rootCertificate.SerialNumber = utils.SerialNumber
	utils.AddMokedDaCertificate(setup, rootCertificate, true)

	// propose second root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(
		setup.Trustee1.String(),
		testconstants.RootCertPem,
		testconstants.Info,
		testconstants.Vid,
		testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// Check indexes
	indexes := []utils.TestIndex{
		{Key: types.ProposedCertificateKeyPrefix, Exist: true}, // we have both: Proposed and Approved
		{Key: types.UniqueCertificateKeyPrefix, Exist: true},
		{Key: types.RejectedCertificateKeyPrefix, Exist: false},
		{Key: types.AllCertificatesKeyPrefix, Exist: true},
		{Key: types.AllCertificatesBySubjectKeyPrefix, Exist: true},
		{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Exist: true},
		{Key: types.ApprovedCertificatesKeyPrefix, Exist: true, Count: 1}, // single approved
		{Key: types.ApprovedCertificatesBySubjectKeyPrefix, Exist: true},
		{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix, Exist: true},
		{Key: types.ApprovedRootCertificatesKeyPrefix, Exist: true},
	}
	resolvedCertificates := utils.CheckCertificateStateIndexes(t, setup, testRootCertificate, indexes)

	// additional check
	require.Equal(t, testconstants.RootSerialNumber, resolvedCertificates.ProposedCertificate.SerialNumber)
}

func TestHandler_AddDaRootCerts_SameSubjectKeyIdButDifferentSubject(t *testing.T) {
	setup := utils.Setup(t)

	testRootCertificate := utils.CreateTestRootCertWithSameSubject()
	testRootCertificate2 := utils.CreateTestRootCertWithSameSubject2()

	// add Certificate1
	rootCertOptions := &utils.RootCertOptions{
		PemCert:      testconstants.PAACertWithSameSubjectID1,
		Subject:      testconstants.PAACertWithSameSubjectID1Subject,
		SubjectKeyID: testconstants.PAACertWithSameSubjectIDSubjectID,
		Info:         testconstants.Info,
		Vid:          testconstants.Vid,
	}
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// add Certificate2
	rootCert2Options := &utils.RootCertOptions{
		PemCert:      testconstants.PAACertWithSameSubjectID2,
		Subject:      testconstants.PAACertWithSameSubjectID2Subject,
		SubjectKeyID: testconstants.PAACertWithSameSubjectIDSubjectID,
		Info:         testconstants.Info,
		Vid:          testconstants.Vid,
	}
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert2Options)

	// Check indexes by subject + subject key id
	allApprovedCertificates, _ := utils.QueryAllApprovedCertificates(setup)
	require.Equal(t, 2, len(allApprovedCertificates))

	allCertificates, _ := utils.QueryAllCertificatesAll(setup)
	require.Equal(t, 2, len(allCertificates))

	// Check indexes
	indexes := []utils.TestIndex{
		{Key: types.ProposedCertificateKeyPrefix, Exist: false},
		{Key: types.UniqueCertificateKeyPrefix, Exist: true},
		{Key: types.RejectedCertificateKeyPrefix, Exist: false},
		{Key: types.AllCertificatesKeyPrefix, Exist: true},
		{Key: types.AllCertificatesBySubjectKeyPrefix, Exist: true},
		{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Exist: true, Count: 2},
		{Key: types.ApprovedCertificatesKeyPrefix, Exist: true},
		{Key: types.ApprovedCertificatesBySubjectKeyPrefix, Exist: true},
		{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix, Exist: true, Count: 2},
		{Key: types.ApprovedRootCertificatesKeyPrefix, Exist: true},
	}
	// check for first
	utils.CheckCertificateStateIndexes(t, setup, testRootCertificate, indexes)
	// check for second
	resolvedCertificates := utils.CheckCertificateStateIndexes(t, setup, testRootCertificate2, indexes)

	// Additional checks
	require.Equal(t, testconstants.PAACertWithSameSubjectIDSubjectID, resolvedCertificates.AllCertificatesBySubjectKeyId[0].SubjectKeyId)
	require.Equal(t, testconstants.PAACertWithSameSubjectID1Subject, resolvedCertificates.AllCertificatesBySubjectKeyId[0].Certs[0].Subject)
	require.Equal(t, testconstants.PAACertWithSameSubjectID2Subject, resolvedCertificates.AllCertificatesBySubjectKeyId[0].Certs[1].Subject)
}

func TestHandler_RejectAddDaRootCert(t *testing.T) {
	setup := utils.Setup(t)

	testRootCertificate := utils.CreateTestRootCert()

	// propose x509 root certificate by account Trustee1
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(
		setup.Trustee1.String(),
		testconstants.RootCertPem,
		testconstants.Info,
		testconstants.Vid,
		testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// reject x509 root certificate by account Trustee2
	rejectAddX509RootCert := types.NewMsgRejectAddX509RootCert(
		setup.Trustee2.String(),
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddX509RootCert)
	require.NoError(t, err)

	// certificate should be in the entity <Proposed X509 Root Certificate>, because we haven't enough reject approvals
	indexes := []utils.TestIndex{
		{Key: types.ProposedCertificateKeyPrefix, Exist: true},
		{Key: types.UniqueCertificateKeyPrefix, Exist: true},
		{Key: types.RejectedCertificateKeyPrefix, Exist: false},
		{Key: types.AllCertificatesKeyPrefix, Exist: false},
		{Key: types.AllCertificatesBySubjectKeyPrefix, Exist: false},
		{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Exist: false},
		{Key: types.ApprovedCertificatesKeyPrefix, Exist: false},
		{Key: types.ApprovedCertificatesBySubjectKeyPrefix, Exist: false},
		{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix, Exist: false},
		{Key: types.ApprovedRootCertificatesKeyPrefix, Exist: false},
	}
	// check certificate state indexes
	resolvedCertificates := utils.CheckCertificateStateIndexes(t, setup, testRootCertificate, indexes)

	// additional checks
	require.Equal(t, setup.Trustee1.String(), resolvedCertificates.ProposedCertificate.Approvals[0].Address)
	require.Equal(t, testconstants.Info, resolvedCertificates.ProposedCertificate.Approvals[0].Info)
	require.Equal(t, setup.Trustee2.String(), resolvedCertificates.ProposedCertificate.Rejects[0].Address)
	require.Equal(t, testconstants.Info, resolvedCertificates.ProposedCertificate.Rejects[0].Info)

	// reject x509 root certificate by account Trustee3
	rejectAddX509RootCert = types.NewMsgRejectAddX509RootCert(
		setup.Trustee3.String(),
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddX509RootCert)
	require.NoError(t, err)

	// certificate should not be in the entity <Proposed X509 Root Certificate>, because we have enough reject approvals
	indexes = []utils.TestIndex{
		{Key: types.RejectedCertificateKeyPrefix, Exist: true},
		{Key: types.UniqueCertificateKeyPrefix, Exist: false},
		{Key: types.ProposedCertificateKeyPrefix, Exist: false},
		{Key: types.AllCertificatesKeyPrefix, Exist: false},
		{Key: types.AllCertificatesBySubjectKeyPrefix, Exist: false},
		{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Exist: false},
		{Key: types.ApprovedCertificatesKeyPrefix, Exist: false},
		{Key: types.ApprovedCertificatesBySubjectKeyPrefix, Exist: false},
		{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix, Exist: false},
		{Key: types.ApprovedRootCertificatesKeyPrefix, Exist: false},
	}
	// check certificate state indexes
	resolvedCertificates = utils.CheckCertificateStateIndexes(t, setup, testRootCertificate, indexes)

	// additional checks
	require.Equal(t, setup.Trustee1.String(), resolvedCertificates.RejectedCertificate.Certs[0].Approvals[0].Address)
	require.Equal(t, testconstants.Info, resolvedCertificates.RejectedCertificate.Certs[0].Approvals[0].Info)
	require.Equal(t, setup.Trustee2.String(), resolvedCertificates.RejectedCertificate.Certs[0].Rejects[0].Address)
	require.Equal(t, testconstants.Info, resolvedCertificates.RejectedCertificate.Certs[0].Rejects[0].Info)
	require.Equal(t, setup.Trustee3.String(), resolvedCertificates.RejectedCertificate.Certs[0].Rejects[1].Address)
	require.Equal(t, testconstants.Info, resolvedCertificates.RejectedCertificate.Certs[0].Rejects[1].Info)
}

func TestHandler_ApproveX509RootCertAndRejectX509RootCert_FromTheSameTrustee(t *testing.T) {
	setup := utils.Setup(t)
	// propose add x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.Trustee,
	} {
		accAddress := utils.GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, 1)

		// approve x509 root certificate by account Trustee2
		approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
		_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
		require.NoError(t, err)

		pendingCert, _ := setup.Keeper.GetProposedCertificate(setup.Ctx, testconstants.RootSubject, testconstants.RootSubjectKeyID)
		prevRejectsLen := len(pendingCert.Rejects)
		prevApprovalsLen := len(pendingCert.Approvals)
		// reject x509 root certificate by account Trustee2
		rejectAddX509RootCert := types.NewMsgRejectAddX509RootCert(setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
		_, err = setup.Handler(setup.Ctx, rejectAddX509RootCert)
		require.NoError(t, err)

		pendingCert, found := setup.Keeper.GetProposedCertificate(setup.Ctx, testconstants.RootSubject, testconstants.RootSubjectKeyID)
		require.True(t, found)
		require.Equal(t, len(pendingCert.Rejects), prevRejectsLen+1)
		require.Equal(t, len(pendingCert.Approvals), prevApprovalsLen-1)
	}
}

func TestHandler_RejectX509RootCertAndApproveX509RootCert_FromTheSameTrustee(t *testing.T) {
	setup := utils.Setup(t)
	// propose add x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.Trustee,
	} {
		accAddress := utils.GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, 1)

		// reject x509 root certificate by account Trustee2
		rejectAddX509RootCert := types.NewMsgRejectAddX509RootCert(setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
		_, err = setup.Handler(setup.Ctx, rejectAddX509RootCert)
		require.NoError(t, err)

		pendingCert, _ := setup.Keeper.GetProposedCertificate(setup.Ctx, testconstants.RootSubject, testconstants.RootSubjectKeyID)
		prevRejectsLen := len(pendingCert.Rejects)
		prevApprovalsLen := len(pendingCert.Approvals)
		// approve x509 root certificate by account Trustee2
		approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
		_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
		require.NoError(t, err)

		pendingCert, found := setup.Keeper.GetProposedCertificate(setup.Ctx, testconstants.RootSubject, testconstants.RootSubjectKeyID)
		require.True(t, found)
		require.Equal(t, len(pendingCert.Rejects), prevRejectsLen-1)
		require.Equal(t, len(pendingCert.Approvals), prevApprovalsLen+1)
	}
}

func TestHandler_RejectX509RootCert_TwoRejectApprovalsAreNeeded_FiveTrustees(t *testing.T) {
	setup := utils.Setup(t)

	// we have 5 trustees: 1 approval comes from propose => we need 2 rejects to make certificate rejected

	// store 4th trustee
	fourthTrustee := utils.GenerateAccAddress()
	setup.AddAccount(fourthTrustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)

	// store 5th trustee
	fifthTrustee := utils.GenerateAccAddress()
	setup.AddAccount(fifthTrustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)

	// propose x509 root certificate by account Trustee1
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// reject x509 root certificate by account Trustee2
	rejectAddX509RootCert := types.NewMsgRejectAddX509RootCert(setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddX509RootCert)
	require.NoError(t, err)

	// certificate should be in the entity <Proposed X509 Root Certificate>, because we haven't enough reject approvals
	proposedCertificate, err := utils.QueryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NoError(t, err)

	// check proposed certificate
	require.Equal(t, proposeAddX509RootCert.Cert, proposedCertificate.PemCert)
	require.Equal(t, proposeAddX509RootCert.Signer, proposedCertificate.Owner)
	require.Equal(t, testconstants.RootSubject, proposedCertificate.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, proposedCertificate.SubjectKeyId)
	require.Equal(t, testconstants.RootSerialNumber, proposedCertificate.SerialNumber)

	// reject x509 root certificate by account Trustee3
	rejectAddX509RootCert = types.NewMsgRejectAddX509RootCert(setup.Trustee3.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddX509RootCert)
	require.NoError(t, err)

	// certificate should be in the entity <Rejected X509 Root Certificate>, because we have enough rejected approvals
	rejectedCertificates, err := utils.QueryRejectedCertificates(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NoError(t, err)

	// check rejected certificate
	rejectedCertificate := rejectedCertificates.Certs[0]
	require.Equal(t, proposeAddX509RootCert.Cert, rejectedCertificate.PemCert)
	require.Equal(t, proposeAddX509RootCert.Signer, rejectedCertificate.Owner)
	require.Equal(t, testconstants.RootSubject, rejectedCertificate.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, rejectedCertificate.SubjectKeyId)
	require.Equal(t, testconstants.RootSerialNumber, rejectedCertificate.SerialNumber)
}

func TestHandler_ProposeAddAndRejectX509RootCert_ByTrustee(t *testing.T) {
	setup := utils.Setup(t)

	// propose x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// reject x509 root certificate
	rejectX509RootCert := types.NewMsgRejectAddX509RootCert(setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectX509RootCert)
	require.NoError(t, err)

	require.False(t, setup.Keeper.IsProposedCertificatePresent(setup.Ctx, testconstants.RootIssuer, testconstants.RootSerialNumber))

	// check that unique certificate key is registered
	require.False(t, setup.Keeper.IsUniqueCertificatePresent(
		setup.Ctx, testconstants.RootIssuer, testconstants.RootSerialNumber))
}

func TestHandler_ProposeAddAndRejectX509RootCert_ByAnotherTrustee(t *testing.T) {
	setup := utils.Setup(t)

	// propose x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// reject x509 root certificate
	rejectX509RootCert := types.NewMsgRejectAddX509RootCert(setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectX509RootCert)
	require.NoError(t, err)

	// query proposed certificate
	proposedCertificate, _ := utils.QueryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)

	// check proposed certificate
	require.Equal(t, proposeAddX509RootCert.Cert, proposedCertificate.PemCert)
	require.Equal(t, proposeAddX509RootCert.Signer, proposedCertificate.Owner)
	require.Equal(t, testconstants.RootSubject, proposedCertificate.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, proposedCertificate.SubjectKeyId)
	require.Equal(t, testconstants.RootSerialNumber, proposedCertificate.SerialNumber)
	require.True(t, proposedCertificate.HasApprovalFrom(setup.Trustee1.String()))

	// check that unique certificate key is registered
	require.True(t, setup.Keeper.IsUniqueCertificatePresent(
		setup.Ctx, testconstants.RootIssuer, testconstants.RootSerialNumber))
}

func TestHandler_ProposeAddAndRejectX509RootCertWithApproval_ByTrustee(t *testing.T) {
	setup := utils.Setup(t)

	accAddress := utils.GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)
	// propose x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// reject x509 root certificate
	rejectX509RootCert := types.NewMsgRejectAddX509RootCert(setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectX509RootCert)
	require.NoError(t, err)

	// query proposed certificate
	proposedCertificate, _ := utils.QueryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)

	// check proposed certificate
	require.Equal(t, proposeAddX509RootCert.Cert, proposedCertificate.PemCert)
	require.Equal(t, proposeAddX509RootCert.Signer, proposedCertificate.Owner)
	require.Equal(t, testconstants.RootSubject, proposedCertificate.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, proposedCertificate.SubjectKeyId)
	require.Equal(t, testconstants.RootSerialNumber, proposedCertificate.SerialNumber)
	require.True(t, proposedCertificate.HasRejectFrom(setup.Trustee1.String()))
	require.True(t, proposedCertificate.HasApprovalFrom(setup.Trustee2.String()))

	// check that unique certificate key is registered
	require.True(t, setup.Keeper.IsUniqueCertificatePresent(
		setup.Ctx, testconstants.RootIssuer, testconstants.RootSerialNumber))
}

// Error cases

func TestHandler_ProposeAddX509RootCert_ByNotTrustee(t *testing.T) {
	setup := utils.Setup(t)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.Vendor,
		dclauthtypes.CertificationCenter,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := utils.GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, 1)

		// propose x509 root certificate
		proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(accAddress.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid, testconstants.CertSchemaVersion)
		_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
		require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
	}
}

func TestHandler_ProposeAddX509RootCert_ForInvalidCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// propose x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.StubCertPem, testconstants.Info, testconstants.Vid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrInvalidCertificate.Is(err))
}

func TestHandler_ProposeAddX509RootCert_ForNonRootCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// propose x509 leaf certificate as root
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.LeafCertPem, testconstants.Info, testconstants.Vid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrInappropriateCertificateType.Is(err))
}

func TestHandler_ProposeAddX509RootCert_ProposedCertificateAlreadyExists(t *testing.T) {
	setup := utils.Setup(t)

	// propose adding of x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// store another account
	anotherAccount := utils.GenerateAccAddress()
	setup.AddAccount(anotherAccount, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)

	// propose adding of the same x509 root certificate again
	proposeAddX509RootCert = types.NewMsgProposeAddX509RootCert(anotherAccount.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid, testconstants.CertSchemaVersion)
	_, err = setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrProposedCertificateAlreadyExists.Is(err))
}

func TestHandler_ProposeAddX509RootCert_CertificateAlreadyExists(t *testing.T) {
	setup := utils.Setup(t)

	// store x509 root certificate
	rootCertificate := utils.RootCertificate(testconstants.Address1)
	setup.Keeper.SetUniqueCertificate(
		setup.Ctx,
		utils.UniqueCertificate(rootCertificate.Subject, rootCertificate.SerialNumber),
	)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// propose adding of the same x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateAlreadyExists.Is(err))
}

func TestHandler_ProposeAddX509RootCert_ForNocCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// Store the NOC root certificate
	vendorAccAddress := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	nocRootCertificate := utils.RootCertificate(vendorAccAddress)
	nocRootCertificate.SerialNumber = testconstants.TestSerialNumber
	nocRootCertificate.CertificateType = types.CertificateType_OperationalPKI
	nocRootCertificate.Approvals = nil
	nocRootCertificate.Rejects = nil

	setup.Keeper.AddAllCertificate(setup.Ctx, nocRootCertificate)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, nocRootCertificate)
	setup.Keeper.AddNocRootCertificate(setup.Ctx, nocRootCertificate)
	uniqueCertificate := types.UniqueCertificate{
		Issuer:       nocRootCertificate.Issuer,
		SerialNumber: nocRootCertificate.SerialNumber,
		Present:      true,
	}
	setup.Keeper.SetUniqueCertificate(setup.Ctx, uniqueCertificate)

	// propose a new root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.True(t, pkitypes.ErrInappropriateCertificateType.Is(err))
}

func TestHandler_ProposeAddX509RootCert_ForDifferentSerialNumberDifferentSigner(t *testing.T) {
	setup := utils.Setup(t)

	// store root certificate with different serial number
	rootCertificate := utils.RootCertificate(testconstants.Address1)
	rootCertificate.SerialNumber = utils.SerialNumber
	setup.Keeper.SetUniqueCertificate(
		setup.Ctx,
		utils.UniqueCertificate(rootCertificate.Subject, rootCertificate.SerialNumber),
	)
	setup.Keeper.AddAllCertificate(setup.Ctx, rootCertificate)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// propose second root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_ApproveAddX509RootCert_ForNotEnoughApprovals(t *testing.T) {
	setup := utils.Setup(t)

	// store account without trustee role
	nonTrustee := utils.GenerateAccAddress()
	setup.AddAccount(nonTrustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)

	// propose x509 root certificate by account without trustee role
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(nonTrustee.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// query certificate
	proposedCertificate, _ := utils.QueryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, proposeAddX509RootCert.Cert, proposedCertificate.PemCert)
	require.True(t, proposedCertificate.HasApprovalFrom(setup.Trustee1.String()))

	// query approved certificate
	_, err = utils.QueryApprovedCertificates(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// approve again from secondTrustee (That makes is 2 trustee's from a total of 3)
	approveAddX509RootCert = types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// query approved certificate and we should get one back
	approvedCertificate, _ := utils.QueryApprovedCertificates(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	for _, cert := range approvedCertificate.Certs {
		// check
		require.Equal(t, testconstants.RootIssuer, cert.Subject)
		require.Equal(t, testconstants.RootSerialNumber, cert.SerialNumber)
		require.True(t, cert.IsRoot)
		require.True(t, cert.HasApprovalFrom(setup.Trustee1.String()))
		require.True(t, cert.HasApprovalFrom(setup.Trustee2.String()))
	}
}

func TestHandler_ApproveAddX509RootCert_ForUnknownProposedCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrProposedCertificateDoesNotExist.Is(err))
}

func TestHandler_ApproveAddX509RootCert_ByNotTrustee(t *testing.T) {
	setup := utils.Setup(t)

	// propose add x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.Vendor,
		dclauthtypes.CertificationCenter,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := utils.GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, 1)

		// approve
		approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
			accAddress.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
		_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
	}
}

func TestHandler_ApproveAddX509RootCert_Twice(t *testing.T) {
	setup := utils.Setup(t)

	// store account without Trustee role
	accAddress := utils.GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)

	// propose add x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(accAddress.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// approve second time
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_RejectX509RootCert_ByNotTrustee(t *testing.T) {
	setup := utils.Setup(t)

	// propose add x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.Vendor,
		dclauthtypes.CertificationCenter,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := utils.GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, 1)

		// reject x509 root certificate
		approveAddX509RootCert := types.NewMsgRejectAddX509RootCert(
			accAddress.String(),
			testconstants.RootSubject,
			testconstants.RootSubjectKeyID,
			testconstants.Info,
		)
		_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
	}
}

func TestHandler_Duplicate_RejectX509RootCert_FromTheSameTrustee(t *testing.T) {
	setup := utils.Setup(t)

	// propose add x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// reject x509 root certificate by account Trustee2
	rejectAddX509RootCert := types.NewMsgRejectAddX509RootCert(setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddX509RootCert)
	require.NoError(t, err)

	// second time reject x509 root certificate by account Trustee2
	rejectAddX509RootCert = types.NewMsgRejectAddX509RootCert(setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddX509RootCert)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_DoubleTimeRejectX509RootCert(t *testing.T) {
	setup := utils.Setup(t)

	// propose x509 root certificate by account Trustee1
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// reject x509 root certificate by account Trustee2
	rejectAddX509RootCert := types.NewMsgRejectAddX509RootCert(setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddX509RootCert)
	require.NoError(t, err)

	// certificate should be in the entity <Proposed X509 Root Certificate>, because we haven't enough reject approvals
	proposedCertificate, err := utils.QueryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NoError(t, err)

	// check proposed certificate
	require.Equal(t, proposeAddX509RootCert.Cert, proposedCertificate.PemCert)
	require.Equal(t, proposeAddX509RootCert.Signer, proposedCertificate.Owner)
	require.Equal(t, testconstants.RootSubject, proposedCertificate.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, proposedCertificate.SubjectKeyId)
	require.Equal(t, testconstants.RootSerialNumber, proposedCertificate.SerialNumber)
	require.Equal(t, setup.Trustee1.String(), proposedCertificate.Approvals[0].Address)
	require.Equal(t, testconstants.Info, proposedCertificate.Approvals[0].Info)
	require.Equal(t, setup.Trustee2.String(), proposedCertificate.Rejects[0].Address)
	require.Equal(t, testconstants.Info, proposedCertificate.Rejects[0].Info)

	// reject x509 root certificate by account Trustee3
	rejectAddX509RootCert = types.NewMsgRejectAddX509RootCert(setup.Trustee3.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddX509RootCert)
	require.NoError(t, err)

	// certificate should not be in the entity <Proposed X509 Root Certificate>, because we have enough reject approvals
	_, err = utils.QueryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Error(t, err)

	// certificate should be in the entity <Rejected X509 Root Certificate>, because we have enough rejected approvals
	rejectedCertificates, err := utils.QueryRejectedCertificates(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NoError(t, err)

	// check rejected certificate
	rejectedCertificate := rejectedCertificates.Certs[0]
	require.Equal(t, proposeAddX509RootCert.Cert, rejectedCertificate.PemCert)
	require.Equal(t, proposeAddX509RootCert.Signer, rejectedCertificate.Owner)
	require.Equal(t, testconstants.RootSubject, rejectedCertificate.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, rejectedCertificate.SubjectKeyId)
	require.Equal(t, testconstants.RootSerialNumber, rejectedCertificate.SerialNumber)
	require.Equal(t, setup.Trustee1.String(), rejectedCertificate.Approvals[0].Address)
	require.Equal(t, testconstants.Info, rejectedCertificate.Approvals[0].Info)
	require.Equal(t, setup.Trustee2.String(), rejectedCertificate.Rejects[0].Address)
	require.Equal(t, testconstants.Info, rejectedCertificate.Rejects[0].Info)
	require.Equal(t, setup.Trustee3.String(), rejectedCertificate.Rejects[1].Address)
	require.Equal(t, testconstants.Info, rejectedCertificate.Rejects[1].Info)

	// propose x509 root certificate by account Trustee1
	proposeAddX509RootCert = types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid, testconstants.CertSchemaVersion)
	_, err = setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// certificate should be in the entity <Proposed X509 Root Certificate>, because we haven't enough reject approvals
	_, err = utils.QueryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NoError(t, err)

	// certificate should not be in the entity <Rejected X509 Root Certificate>, because we have propose that certificate
	_, err = utils.QueryRejectedCertificates(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Error(t, err)

	// reject x509 root certificate by account Trustee3
	rejectAddX509RootCert = types.NewMsgRejectAddX509RootCert(setup.Trustee3.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddX509RootCert)
	require.NoError(t, err)

	// reject x509 root certificate by account Trustee2
	rejectAddX509RootCert = types.NewMsgRejectAddX509RootCert(setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddX509RootCert)
	require.NoError(t, err)

	// certificate should be in the entity <Rejected X509 Root Certificate>, because we have enough rejected approvals
	rejectedCertificates, err = utils.QueryRejectedCertificates(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NoError(t, err)

	// check rejected certificate
	rejectedCertificate = rejectedCertificates.Certs[0]
	require.Equal(t, proposeAddX509RootCert.Cert, rejectedCertificate.PemCert)
	require.Equal(t, proposeAddX509RootCert.Signer, rejectedCertificate.Owner)
	require.Equal(t, testconstants.RootSubject, rejectedCertificate.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, rejectedCertificate.SubjectKeyId)
	require.Equal(t, testconstants.RootSerialNumber, rejectedCertificate.SerialNumber)
	require.Equal(t, setup.Trustee1.String(), rejectedCertificate.Approvals[0].Address)
	require.Equal(t, testconstants.Info, rejectedCertificate.Approvals[0].Info)
	require.Equal(t, setup.Trustee3.String(), rejectedCertificate.Rejects[0].Address)
	require.Equal(t, testconstants.Info, rejectedCertificate.Rejects[0].Info)
	require.Equal(t, setup.Trustee2.String(), rejectedCertificate.Rejects[1].Address)
	require.Equal(t, testconstants.Info, rejectedCertificate.Rejects[1].Info)
}
