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
)

// Main

func TestHandler_AddDaRootCert(t *testing.T) {
	setup := utils.Setup(t)

	// propose add x509 root certificate by trustee
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, rootCertificate)

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
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, rootCertificate)

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
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, rootCertificate)

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

func TestHandler_AddDaRootCert_SameSkid_DifferentSubject(t *testing.T) {
	setup := utils.Setup(t)

	// add Certificate1
	testRootCertificate := utils.RootDaCertWithSameSubjectKeyID1(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, testRootCertificate)

	// add Certificate2
	testRootCertificate2 := utils.RootDaCertificateWithSameSubjectKeyID2(setup.Trustee1)
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

func TestHandler_AddDaRootCert_SameSubjectAndSkid_DifferentSerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	rootCertificate1 := utils.RootDaCertificateWithSameSubjectAndSKID1(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate1)

	rootCertificate2 := utils.RootDaCertificateWithSameSubjectAndSKID2(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate2)

	// Check:
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix, Count: 2},
			{Key: types.AllCertificatesBySubjectKeyPrefix, Count: 1},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.ApprovedCertificatesKeyPrefix, Count: 2},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix, Count: 1},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateRevocationKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate1, indexes)
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate2, indexes)
}

func TestHandler_ApproveAddDaRootCert_PreviouslyRejectedByCurrentTrustee(t *testing.T) {
	setup := utils.Setup(t)

	// Add one more Trustee
	setup.CreateTrusteeAccount(testconstants.Vid)

	// propose add x509 root certificate
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, rootCertificate)

	// reject x509 root certificate by account Trustee2
	utils.RejectDaRootCertificate(setup, setup.Trustee2, rootCertificate.Subject, rootCertificate.SubjectKeyId)

	// approve x509 root certificate by account Trustee2
	utils.ApproveDaRootCertificate(setup, setup.Trustee2, rootCertificate.Subject, rootCertificate.SubjectKeyId)

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
	resolvedCertificates := utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)

	require.Len(t, resolvedCertificates.ProposedCertificate.Approvals, 2)
	require.Len(t, resolvedCertificates.ProposedCertificate.Rejects, 0)
	require.Equal(t, setup.Trustee1.String(), resolvedCertificates.ProposedCertificate.Approvals[0].Address)
	require.Equal(t, setup.Trustee2.String(), resolvedCertificates.ProposedCertificate.Approvals[1].Address)
}

// Error cases

func TestHandler_ApproveAddDaRootCert_ForUnknownProposedCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrProposedCertificateDoesNotExist.Is(err))
}

func TestHandler_ApproveAddDaRootCert_ByNotTrustee(t *testing.T) {
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

func TestHandler_ApproveAddDaRootCert_Twice(t *testing.T) {
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
