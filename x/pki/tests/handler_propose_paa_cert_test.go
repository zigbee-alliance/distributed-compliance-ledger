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

func TestHandler_ProposeAddDaRootCert(t *testing.T) {
	setup := utils.Setup(t)

	// propose DA root certificate
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	proposeAddX509RootCert := utils.ProposeDaRootCertificate(setup, rootCertificate)

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

	require.Equal(t, proposeAddX509RootCert.Cert, resolvedCertificates.ProposedCertificate.PemCert)
	require.True(t, resolvedCertificates.ProposedCertificate.HasApprovalFrom(proposeAddX509RootCert.Signer))
}

func TestHandler_ProposeAddDaRootCert_SameSkidButDifferentSubject(t *testing.T) {
	setup := utils.Setup(t)

	// add Certificate1
	testRootCertificate := utils.RootDaCertWithSameSubjectKeyID1(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, testRootCertificate)

	// add Certificate2
	testRootCertificate2 := utils.RootDaCertificateWithSameSubjectKeyID2(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, testRootCertificate2)

	// Check indexes by subject + subject key id
	allApprovedCertificates, _ := utils.QueryAllProposedCertificates(setup)
	require.Equal(t, 2, len(allApprovedCertificates))

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
	// check for first
	utils.CheckCertificateStateIndexes(t, setup, testRootCertificate, indexes)
	utils.CheckCertificateStateIndexes(t, setup, testRootCertificate2, indexes)
}

func TestHandler_ProposeAddDaRootCert_DifferentSerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	// store root certificate with different serial number
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	rootCertificate.SerialNumber = utils.SerialNumber
	utils.AddMokedDaCertificate(setup, rootCertificate)

	// propose second root certificate
	testRootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, testRootCertificate)

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

func TestHandler_ProposeAddDaRootCert_PreviouslyRejected(t *testing.T) {
	setup := utils.Setup(t)

	// propose x509 root certificate by account Trustee1
	testRootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, testRootCertificate)

	// reject x509 root certificate by account Trustee2
	rejectAddX509RootCert1 := utils.RejectDaRootCertificate(setup, setup.Trustee2, testRootCertificate.Subject, testRootCertificate.SubjectKeyId)

	// reject x509 root certificate by account Trustee3
	rejectAddX509RootCert2 := utils.RejectDaRootCertificate(setup, setup.Trustee3, testRootCertificate.Subject, testRootCertificate.SubjectKeyId)

	// Check state indexes - rejected
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.RejectedCertificateKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, testRootCertificate, indexes)

	// propose again
	proposeAddX509RootCert := utils.ProposeDaRootCertificate(setup, testRootCertificate)

	// Check state indexes - proposed
	indexes = utils.TestIndexes{
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

	require.Equal(t, proposeAddX509RootCert.Cert, resolvedCertificates.ProposedCertificate.PemCert)
	require.True(t, resolvedCertificates.ProposedCertificate.HasApprovalFrom(proposeAddX509RootCert.Signer))
	require.False(t, resolvedCertificates.ProposedCertificate.HasRejectFrom(rejectAddX509RootCert1.Signer))
	require.False(t, resolvedCertificates.ProposedCertificate.HasRejectFrom(rejectAddX509RootCert2.Signer))
}

// Error cases

func TestHandler_ProposeAddDaRootCert_ByNotTrustee(t *testing.T) {
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

func TestHandler_ProposeAddDaRootCert_ForInvalidCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// propose x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.StubCertPem, testconstants.Info, testconstants.Vid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrInvalidCertificate.Is(err))
}

func TestHandler_ProposeAddDaRootCert_ForNonRootCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// propose x509 leaf certificate as root
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.LeafCertPem, testconstants.Info, testconstants.Vid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrInappropriateCertificateType.Is(err))
}

func TestHandler_ProposeAddDaRootCert_ProposedCertificateAlreadyExists(t *testing.T) {
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

func TestHandler_ProposeAddDaRootCert_CertificateAlreadyExists(t *testing.T) {
	setup := utils.Setup(t)

	// store x509 root certificate
	rootCertificate := utils.RootDaCertificate(testconstants.Address1)
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

func TestHandler_ProposeAddDaRootCert_ForNocCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// Store the NOC root certificate
	nocRootCertificate := utils.RootDaCertificate(setup.Vendor1)
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

func TestHandler_ProposeAddDaRootCert_ForDifferentSigner(t *testing.T) {
	setup := utils.Setup(t)

	// store root certificate with different serial number
	rootCertificate := utils.RootDaCertificate(testconstants.Address1)
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
