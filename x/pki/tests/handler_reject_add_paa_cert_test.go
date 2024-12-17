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

	// propose x509 root certificate by account Trustee1
	testRootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, testRootCertificate)

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

func TestHandler_RejectAddDaRootCert_TwoRejectApprovalsAreNeeded_FiveTrustees(t *testing.T) {
	setup := utils.Setup(t)

	// we have 5 trustees: 1 approval comes from propose => we need 2 rejects to make certificate rejected

	// store 4th trustee
	setup.CreateTrusteeAccount(testconstants.Vid)

	// store 5th trustee
	setup.CreateTrusteeAccount(testconstants.Vid)

	// propose x509 root certificate by account Trustee1
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, rootCertificate)

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

func TestHandler_RejectAddDaRootCert_PreviouslyApprovedByCurrentTrustee_CertificateHasOtherApproval(t *testing.T) {
	setup := utils.Setup(t)

	// Add one more Trustee
	setup.CreateTrusteeAccount(testconstants.Vid)

	// propose add x509 root certificate
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, rootCertificate)

	// approve x509 root certificate by account Trustee2
	utils.ApproveDaRootCertificate(setup, setup.Trustee2, rootCertificate.Subject, rootCertificate.SubjectKeyId)

	// reject x509 root certificate by account Trustee2
	utils.RejectDaRootCertificate(setup, setup.Trustee2, rootCertificate.Subject, rootCertificate.SubjectKeyId)

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

	require.Len(t, resolvedCertificates.ProposedCertificate.Approvals, 1)
	require.Len(t, resolvedCertificates.ProposedCertificate.Rejects, 1)
	require.Equal(t, setup.Trustee1.String(), resolvedCertificates.ProposedCertificate.Approvals[0].Address)
	require.Equal(t, setup.Trustee2.String(), resolvedCertificates.ProposedCertificate.Rejects[0].Address)
}

func TestHandler_RejectAddDaRootCert_PreviouslyApprovedByCurrentTrustee_CertificateHasOtherReject(t *testing.T) {
	setup := utils.Setup(t)

	// Add two more Trustee
	setup.CreateTrusteeAccount(testconstants.Vid)
	setup.CreateTrusteeAccount(testconstants.Vid)
	setup.CreateTrusteeAccount(testconstants.Vid)

	// propose add x509 root certificate
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, rootCertificate)

	// approve x509 root certificate by account Trustee2
	utils.ApproveDaRootCertificate(setup, setup.Trustee2, rootCertificate.Subject, rootCertificate.SubjectKeyId)

	// reject x509 root certificate by account Trustee1
	utils.RejectDaRootCertificate(setup, setup.Trustee1, rootCertificate.Subject, rootCertificate.SubjectKeyId)

	// reject x509 root certificate by account Trustee2
	utils.RejectDaRootCertificate(setup, setup.Trustee2, rootCertificate.Subject, rootCertificate.SubjectKeyId)

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

	require.Len(t, resolvedCertificates.ProposedCertificate.Approvals, 0)
	require.Len(t, resolvedCertificates.ProposedCertificate.Rejects, 2)
	require.Equal(t, setup.Trustee1.String(), resolvedCertificates.ProposedCertificate.Rejects[0].Address)
	require.Equal(t, setup.Trustee2.String(), resolvedCertificates.ProposedCertificate.Rejects[1].Address)
}

func TestHandler_RejectAddDaRootCert_PreviouslyApprovedByCurrentTrustee_CertificateNotHasOtherApproval(t *testing.T) {
	setup := utils.Setup(t)

	// propose add x509 root certificate
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, rootCertificate)

	// reject x509 root certificate by account Trustee1 (who proposed)
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

func TestHandler_RejectAddDaRootCert_ForUnknownProposedCertificate(t *testing.T) {
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

func TestHandler_RejectX509RootCert_TwiceFromTheSameTrustee(t *testing.T) {
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
