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

	rootCertificate := utils.RootCertificate(setup.Trustee1)

	// propose DA root certificate
	proposeAddX509RootCert := utils.ProposeDaRootCertificate(setup, setup.Trustee1, rootCertificate.PemCert)

	// Check state indexes
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.UniqueCertificateKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.RejectedCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
	}
	resolvedCertificates := utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)

	// additional checks
	require.Equal(t, proposeAddX509RootCert.Cert, resolvedCertificates.ProposedCertificate.PemCert)
	require.True(t, resolvedCertificates.ProposedCertificate.HasApprovalFrom(proposeAddX509RootCert.Signer))
}

func TestHandler_AddDaRootCert(t *testing.T) {
	setup := utils.Setup(t)

	// propose add x509 root certificate by trustee
	rootCertificate := utils.RootCertificate(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, setup.Trustee1, rootCertificate.PemCert)

	// approve by second trustee
	utils.ApproveDaRootCertificate(setup, setup.Trustee2, rootCertificate.Subject, rootCertificate.SubjectKeyId)

	// Check state indexes
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.RejectedCertificateKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)
}

func TestHandler_AddDaRootCert_TwoThirdApprovalsNeeded(t *testing.T) {
	setup := utils.Setup(t)

	// propose x509 root certificate by account without trustee role
	rootCertificate := utils.RootCertificate(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, setup.Trustee1, rootCertificate.PemCert)

	// Create an array of trustee account from 1 to 50
	trusteeAccounts, totalAdditionalTrustees := setup.CreateNTrusteeAccounts()

	// We have 3 Trustees in test setup.
	twoThirds := int(math.Ceil(types.RootCertificateApprovalsPercent * float64(3+totalAdditionalTrustees)))

	// Until we hit 2/3 of the total number of Trustees, we should not be able to approve the certificate
	for i := 1; i < twoThirds-1; i++ {
		utils.ApproveDaRootCertificate(setup, trusteeAccounts[i], rootCertificate.Subject, rootCertificate.SubjectKeyId)

		// Check state indexes
		indexes := utils.TestIndexes{
			Present: []utils.TestIndex{
				{Key: types.UniqueCertificateKeyPrefix},
				{Key: types.ProposedCertificateKeyPrefix},
			},
			Missing: []utils.TestIndex{
				{Key: types.RejectedCertificateKeyPrefix},
				{Key: types.AllCertificatesKeyPrefix},
				{Key: types.AllCertificatesBySubjectKeyPrefix},
				{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
				{Key: types.ApprovedCertificatesKeyPrefix},
				{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
				{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
				{Key: types.ApprovedRootCertificatesKeyPrefix},
			},
		}
		utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)
	}

	// One more approval will move this to approved state from pending
	utils.ApproveDaRootCertificate(setup, setup.Trustee2, rootCertificate.Subject, rootCertificate.SubjectKeyId)

	// Check state indexes
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.RejectedCertificateKeyPrefix},
		},
	}
	resolvedCertificates := utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)

	// Additional checks
	for i := 1; i < twoThirds-1; i++ {
		require.Equal(t, resolvedCertificates.ApprovedCertificates.Certs[0].HasApprovalFrom(trusteeAccounts[i].String()), true)
	}
	require.Equal(t, resolvedCertificates.ApprovedCertificates.Certs[0].HasApprovalFrom(setup.Trustee1.String()), true)
	require.Equal(t, resolvedCertificates.ApprovedCertificates.Certs[0].HasApprovalFrom(setup.Trustee2.String()), true)
}

func TestHandler_AddDaRootCert_FourApprovalsAreNeeded_FiveTrustees(t *testing.T) {
	setup := utils.Setup(t)

	// we have 5 trustees: 1 approval comes from propose => we need 3 more approvals

	// store 4th trustee
	fourthTrustee := utils.GenerateAccAddress()
	setup.AddAccount(fourthTrustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)

	// store 5th trustee
	fifthTrustee := utils.GenerateAccAddress()
	setup.AddAccount(fifthTrustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)

	// propose x509 root certificate by account Trustee1
	rootCertificate := utils.RootCertificate(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, setup.Trustee1, rootCertificate.PemCert)

	// approve x509 root certificate by account Trustee2
	utils.ApproveDaRootCertificate(setup, setup.Trustee2, rootCertificate.Subject, rootCertificate.SubjectKeyId)

	// approve x509 root certificate by account Trustee3
	utils.ApproveDaRootCertificate(setup, setup.Trustee3, rootCertificate.Subject, rootCertificate.SubjectKeyId)

	// reject x509 root certificate by account Trustee4
	utils.RejectDaRootCertificate(setup, fourthTrustee, rootCertificate.Subject, rootCertificate.SubjectKeyId)

	// Check state indexes - certificate is in proposed state
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.UniqueCertificateKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.RejectedCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)

	// approve x509 root certificate by account Trustee5
	utils.ApproveDaRootCertificate(setup, fifthTrustee, rootCertificate.Subject, rootCertificate.SubjectKeyId)

	// Check state indexes
	indexes = utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.RejectedCertificateKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)
}

// Extra cases

func TestHandler_ProposeAddX509RootCert_ForDifferentSerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	// store root certificate with different serial number
	rootCertificate := utils.RootCertificate(setup.Trustee1)
	rootCertificate.SerialNumber = utils.SerialNumber
	utils.AddMokedDaCertificate(setup, rootCertificate, true)

	// propose second root certificate
	testRootCertificate := utils.RootCertificate(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, setup.Trustee1, testRootCertificate.PemCert)

	// Check state indexes
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.ProposedCertificateKeyPrefix}, // we have both: Proposed and Approved
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix, Count: 1}, // single approved
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.RejectedCertificateKeyPrefix},
		},
	}
	resolvedCertificates := utils.CheckCertificateStateIndexes(t, setup, testRootCertificate, indexes)

	// additional check
	require.Equal(t, testRootCertificate.SerialNumber, resolvedCertificates.ProposedCertificate.SerialNumber)
}

func TestHandler_AddDaRootCerts_SameSubjectKeyIdButDifferentSubject(t *testing.T) {
	setup := utils.Setup(t)

	// add Certificate1
	testRootCertificate := utils.PAACertWithSameSubjectID1(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, testRootCertificate)

	// add Certificate2
	testRootCertificate2 := utils.PAACertWithSameSubjectID2(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, testRootCertificate2)

	// Check indexes by subject + subject key id
	allApprovedCertificates, _ := utils.QueryAllApprovedCertificates(setup)
	require.Equal(t, 2, len(allApprovedCertificates))

	allCertificates, _ := utils.QueryAllCertificatesAll(setup)
	require.Equal(t, 2, len(allCertificates))

	// Check state indexes
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.RejectedCertificateKeyPrefix},
		},
	}
	// check for first
	utils.CheckCertificateStateIndexes(t, setup, testRootCertificate, indexes)
	// check for second
	resolvedCertificates := utils.CheckCertificateStateIndexes(t, setup, testRootCertificate2, indexes)

	// Additional checks
	require.Equal(t, testRootCertificate.SubjectKeyId, resolvedCertificates.AllCertificatesBySubjectKeyID[0].SubjectKeyId)
	require.Equal(t, testRootCertificate.Subject, resolvedCertificates.AllCertificatesBySubjectKeyID[0].Certs[0].Subject)
	require.Equal(t, testRootCertificate2.Subject, resolvedCertificates.AllCertificatesBySubjectKeyID[0].Certs[1].Subject)
}

func TestHandler_RejectAddDaRootCert(t *testing.T) {
	setup := utils.Setup(t)

	// propose x509 root certificate by account Trustee1
	testRootCertificate := utils.RootCertificate(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, setup.Trustee1, testRootCertificate.PemCert)

	// reject x509 root certificate by account Trustee2
	utils.RejectDaRootCertificate(setup, setup.Trustee2, testRootCertificate.Subject, testRootCertificate.SubjectKeyId)

	// certificate should be in the entity <Proposed X509 Root Certificate>, because we haven't enough reject approvals
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.UniqueCertificateKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.RejectedCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
	}
	// check certificate state indexes
	resolvedCertificates := utils.CheckCertificateStateIndexes(t, setup, testRootCertificate, indexes)

	// additional checks
	require.Equal(t, setup.Trustee1.String(), resolvedCertificates.ProposedCertificate.Approvals[0].Address)
	require.Equal(t, testconstants.Info, resolvedCertificates.ProposedCertificate.Approvals[0].Info)
	require.Equal(t, setup.Trustee2.String(), resolvedCertificates.ProposedCertificate.Rejects[0].Address)
	require.Equal(t, testconstants.Info, resolvedCertificates.ProposedCertificate.Rejects[0].Info)

	// reject x509 root certificate by account Trustee3
	utils.RejectDaRootCertificate(setup, setup.Trustee3, testRootCertificate.Subject, testRootCertificate.SubjectKeyId)

	// certificate should not be in the entity <Proposed X509 Root Certificate>, because we have enough reject approvals
	indexes = utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.RejectedCertificateKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
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
	rootCertificate := utils.RootCertificate(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, setup.Trustee1, rootCertificate.PemCert)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.Trustee,
	} {
		accAddress := utils.GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, 1)

		// approve x509 root certificate by account Trustee2
		utils.ApproveDaRootCertificate(setup, setup.Trustee2, rootCertificate.Subject, rootCertificate.SubjectKeyId)

		pendingCert, _ := setup.Keeper.GetProposedCertificate(
			setup.Ctx,
			rootCertificate.Subject,
			rootCertificate.SubjectKeyId)
		prevRejectsLen := len(pendingCert.Rejects)
		prevApprovalsLen := len(pendingCert.Approvals)

		// reject x509 root certificate by account Trustee2
		utils.RejectDaRootCertificate(setup, setup.Trustee2, rootCertificate.Subject, rootCertificate.SubjectKeyId)

		pendingCert, found := setup.Keeper.GetProposedCertificate(setup.Ctx,
			rootCertificate.Subject,
			rootCertificate.SubjectKeyId)
		require.True(t, found)
		require.Equal(t, len(pendingCert.Rejects), prevRejectsLen+1)
		require.Equal(t, len(pendingCert.Approvals), prevApprovalsLen-1)
	}
}

func TestHandler_RejectX509RootCertAndApproveX509RootCert_FromTheSameTrustee(t *testing.T) {
	setup := utils.Setup(t)

	// propose add x509 root certificate
	rootCertificate := utils.RootCertificate(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, setup.Trustee1, rootCertificate.PemCert)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.Trustee,
	} {
		accAddress := utils.GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, 1)

		// reject x509 root certificate by account Trustee2
		utils.RejectDaRootCertificate(setup, setup.Trustee2, rootCertificate.Subject, rootCertificate.SubjectKeyId)

		pendingCert, _ := setup.Keeper.GetProposedCertificate(
			setup.Ctx,
			rootCertificate.Subject,
			rootCertificate.SubjectKeyId)
		prevRejectsLen := len(pendingCert.Rejects)
		prevApprovalsLen := len(pendingCert.Approvals)

		// approve x509 root certificate by account Trustee2
		utils.ApproveDaRootCertificate(setup, setup.Trustee2, rootCertificate.Subject, rootCertificate.SubjectKeyId)

		pendingCert, found := setup.Keeper.GetProposedCertificate(
			setup.Ctx,
			rootCertificate.Subject,
			rootCertificate.SubjectKeyId)
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
	rootCertificate := utils.RootCertificate(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, setup.Trustee1, rootCertificate.PemCert)

	// reject x509 root certificate by account Trustee2
	utils.RejectDaRootCertificate(setup, setup.Trustee2, rootCertificate.Subject, rootCertificate.SubjectKeyId)

	// Check state indexes
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.UniqueCertificateKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.RejectedCertificateKeyPrefix}, // not rejected yet
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)

	// reject x509 root certificate by account Trustee3
	utils.RejectDaRootCertificate(setup, setup.Trustee3, rootCertificate.Subject, rootCertificate.SubjectKeyId)

	// Check state indexes
	indexes = utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.RejectedCertificateKeyPrefix}, // certificate is rejected now
		},
		Missing: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)
}

func TestHandler_ProposeAddAndRejectX509RootCert_ByTrustee(t *testing.T) {
	setup := utils.Setup(t)

	// propose x509 root certificate
	rootCertificate := utils.RootCertificate(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, setup.Trustee1, rootCertificate.PemCert)

	// reject x509 root certificate
	utils.RejectDaRootCertificate(setup, setup.Trustee1, rootCertificate.Subject, rootCertificate.SubjectKeyId)

	// Check state indexes
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{},
		Missing: []utils.TestIndex{
			{Key: types.RejectedCertificateKeyPrefix}, // certificates do not get into rejected collection because there were no approvals before
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)
}

func TestHandler_ProposeAddAndRejectX509RootCert_ByAnotherTrustee(t *testing.T) {
	setup := utils.Setup(t)

	// propose x509 root certificate
	rootCertificate := utils.RootCertificate(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, setup.Trustee1, rootCertificate.PemCert)

	// reject x509 root certificate
	utils.RejectDaRootCertificate(setup, setup.Trustee2, rootCertificate.Subject, rootCertificate.SubjectKeyId)

	// Check state indexes
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.RejectedCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)
}

func TestHandler_ProposeAddAndRejectX509RootCertWithApproval_ByTrustee(t *testing.T) {
	setup := utils.Setup(t)

	// add another trustee
	setup.CreateTrusteeAccount(1)

	// propose x509 root certificate
	rootCertificate := utils.RootCertificate(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, setup.Trustee1, rootCertificate.PemCert)

	// approve
	utils.ApproveDaRootCertificate(setup, setup.Trustee2, rootCertificate.Subject, rootCertificate.SubjectKeyId)

	// reject x509 root certificate
	utils.RejectDaRootCertificate(setup, setup.Trustee1, rootCertificate.Subject, rootCertificate.SubjectKeyId)

	// Check state indexes
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.RejectedCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
	}
	resolvedCertificates := utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)

	// additional checks
	require.True(t, resolvedCertificates.ProposedCertificate.HasRejectFrom(setup.Trustee1.String()))
	require.True(t, resolvedCertificates.ProposedCertificate.HasApprovalFrom(setup.Trustee2.String()))
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
	nocRootCertificate := utils.RootCertificate(setup.Vendor1)
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
	nonTrustee := setup.CreateTrusteeAccount(1)

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
	accAddress := setup.CreateTrusteeAccount(1)

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
