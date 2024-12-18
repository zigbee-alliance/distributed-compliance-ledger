package tests

import (
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

func TestHandler_RejectAddDaRootCert(t *testing.T) {
	setup := utils.Setup(t)

	// propose root certificate by account Trustee1
	testRootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, testRootCertificate)

	// reject root certificate by account Trustee2
	utils.RejectDaRootCertificate(setup, setup.Trustee2, testRootCertificate.Subject, testRootCertificate.SubjectKeyId)

	// check state indexes - certificate is proposed
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
	resolvedCertificates := utils.CheckCertificateStateIndexes(t, setup, testRootCertificate, indexes)

	// additional checks - approvals and rejects
	require.Equal(t, setup.Trustee1.String(), resolvedCertificates.ProposedCertificate.Approvals[0].Address)
	require.Equal(t, testconstants.Info, resolvedCertificates.ProposedCertificate.Approvals[0].Info)
	require.Equal(t, setup.Trustee2.String(), resolvedCertificates.ProposedCertificate.Rejects[0].Address)
	require.Equal(t, testconstants.Info, resolvedCertificates.ProposedCertificate.Rejects[0].Info)

	// reject x509 root certificate by account Trustee3
	utils.RejectDaRootCertificate(setup, setup.Trustee3, testRootCertificate.Subject, testRootCertificate.SubjectKeyId)

	// check state indexes - certificate is rejected
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
	resolvedCertificates = utils.CheckCertificateStateIndexes(t, setup, testRootCertificate, indexes)

	// additional checks - approvals and rejects
	require.Equal(t, setup.Trustee1.String(), resolvedCertificates.RejectedCertificate.Certs[0].Approvals[0].Address)
	require.Equal(t, testconstants.Info, resolvedCertificates.RejectedCertificate.Certs[0].Approvals[0].Info)
	require.Equal(t, setup.Trustee2.String(), resolvedCertificates.RejectedCertificate.Certs[0].Rejects[0].Address)
	require.Equal(t, testconstants.Info, resolvedCertificates.RejectedCertificate.Certs[0].Rejects[0].Info)
	require.Equal(t, setup.Trustee3.String(), resolvedCertificates.RejectedCertificate.Certs[0].Rejects[1].Address)
	require.Equal(t, testconstants.Info, resolvedCertificates.RejectedCertificate.Certs[0].Rejects[1].Info)
}

func TestHandler_RejectAddDaRootCert_TwoRejectApprovalsAreNeeded_FiveTrustees(t *testing.T) {
	setup := utils.Setup(t)

	// we have 5 trustees: 1 approval comes from propose => we need 2 rejects to make certificate rejected

	// store 4th trustee
	setup.CreateTrusteeAccount(testconstants.Vid)

	// store 5th trustee
	setup.CreateTrusteeAccount(testconstants.Vid)

	// propose root certificate by account Trustee1
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, rootCertificate)

	// reject root certificate by account Trustee2
	utils.RejectDaRootCertificate(setup, setup.Trustee2, rootCertificate.Subject, rootCertificate.SubjectKeyId)

	// Check state indexes - certificate is proposed
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

	// reject root certificate by account Trustee3
	utils.RejectDaRootCertificate(setup, setup.Trustee3, rootCertificate.Subject, rootCertificate.SubjectKeyId)

	// Check state indexes - certificate is rejected
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
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)
}

func TestHandler_RejectAddDaRootCert_CertificateHasOtherApproval(t *testing.T) {
	setup := utils.Setup(t)

	// add one more Trustee
	setup.CreateTrusteeAccount(testconstants.Vid)

	// propose add root certificate
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, rootCertificate)

	// approve root certificate by account Trustee2
	utils.ApproveDaRootCertificate(setup, setup.Trustee2, rootCertificate.Subject, rootCertificate.SubjectKeyId)

	// reject root certificate by account Trustee2
	utils.RejectDaRootCertificate(setup, setup.Trustee2, rootCertificate.Subject, rootCertificate.SubjectKeyId)

	// check state indexes - certificate is proposed
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

	// additional check - approvals and rejects
	require.Len(t, resolvedCertificates.ProposedCertificate.Approvals, 1)
	require.Len(t, resolvedCertificates.ProposedCertificate.Rejects, 1)
	require.Equal(t, setup.Trustee1.String(), resolvedCertificates.ProposedCertificate.Approvals[0].Address)
	require.Equal(t, setup.Trustee2.String(), resolvedCertificates.ProposedCertificate.Rejects[0].Address)
}

func TestHandler_RejectAddDaRootCert_CertificateHasOtherReject(t *testing.T) {
	setup := utils.Setup(t)

	// Add more Trustees
	setup.CreateTrusteeAccount(testconstants.Vid)
	setup.CreateTrusteeAccount(testconstants.Vid)
	setup.CreateTrusteeAccount(testconstants.Vid)

	// propose add root certificate
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, rootCertificate)

	// approve root certificate by account Trustee2
	utils.ApproveDaRootCertificate(setup, setup.Trustee2, rootCertificate.Subject, rootCertificate.SubjectKeyId)

	// reject root certificate by account Trustee1
	utils.RejectDaRootCertificate(setup, setup.Trustee1, rootCertificate.Subject, rootCertificate.SubjectKeyId)

	// reject root certificate by account Trustee2
	utils.RejectDaRootCertificate(setup, setup.Trustee2, rootCertificate.Subject, rootCertificate.SubjectKeyId)

	// check state indexes - certificate is proposed
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

	// additional check - approvals and rejects
	require.Len(t, resolvedCertificates.ProposedCertificate.Approvals, 0)
	require.Len(t, resolvedCertificates.ProposedCertificate.Rejects, 2)
	require.Equal(t, setup.Trustee1.String(), resolvedCertificates.ProposedCertificate.Rejects[0].Address)
	require.Equal(t, setup.Trustee2.String(), resolvedCertificates.ProposedCertificate.Rejects[1].Address)
}

func TestHandler_RejectAddDaRootCert_CertificateNotHasOtherApprovalAndRejects(t *testing.T) {
	setup := utils.Setup(t)

	// propose add root certificate
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, rootCertificate)

	// reject root certificate by account Trustee1 (who proposed)
	utils.RejectDaRootCertificate(setup, setup.Trustee1, rootCertificate.Subject, rootCertificate.SubjectKeyId)

	// check certificate state indexes - certificate removed
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.UniqueCertificateKeyPrefix},
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

// Error cases

func TestHandler_RejectAddDaRootCert_UnknownProposedCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// approve
	rejectAddX509RootCert := types.NewMsgRejectAddX509RootCert(
		setup.Trustee1.String(),
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.Info)
	_, err := setup.Handler(setup.Ctx, rejectAddX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrProposedCertificateDoesNotExist.Is(err))
}

func TestHandler_RejectAddDaRootCert_ByNotTrustee(t *testing.T) {
	setup := utils.Setup(t)

	// propose add x509 root certificate
	testRootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, testRootCertificate)

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
			testRootCertificate.Subject,
			testRootCertificate.SubjectKeyId,
			testconstants.Info,
		)
		_, err := setup.Handler(setup.Ctx, approveAddX509RootCert)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
	}
}

func TestHandler_RejectAddDaRootCert_Twice(t *testing.T) {
	setup := utils.Setup(t)

	// propose add root certificate
	testRootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, testRootCertificate)

	// reject root certificate by account Trustee2
	rejectAddX509RootCert := types.NewMsgRejectAddX509RootCert(
		setup.Trustee2.String(),
		testRootCertificate.Subject,
		testRootCertificate.SubjectKeyId,
		testconstants.Info)
	_, err := setup.Handler(setup.Ctx, rejectAddX509RootCert)
	require.NoError(t, err)

	// second time reject root certificate by account Trustee2
	_, err = setup.Handler(setup.Ctx, rejectAddX509RootCert)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}
