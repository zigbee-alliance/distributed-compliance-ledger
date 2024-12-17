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

func TestHandler_ApproveRevokeDaRootCert_NotEnoughApprovals(t *testing.T) {
	setup := utils.Setup(t)

	// propose and approve x509 root certificate
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate)

	// Add 1 more trustee (this will bring the total trustee's to 4)
	setup.CreateTrusteeAccount(1)

	// propose revocation of x509 root certificate
	utils.ProposeRevokeDaRootCertificate(
		setup,
		setup.Trustee1,
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		rootCertificate.SerialNumber,
		false)

	// approve
	utils.ApproveRevokeDaRootCertificate(
		setup,
		setup.Trustee2,
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		rootCertificate.SerialNumber)

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
			{Key: types.ProposedCertificateRevocationKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)
}

func TestHandler_RevokeDaRootCert_BySubjectAndSKID(t *testing.T) {
	setup := utils.Setup(t)

	rootCertificate1 := utils.RootDaCertificateWithSameSubjectAndSKID1(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate1)

	rootCertificate2 := utils.RootDaCertificateWithSameSubjectAndSKID2(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate2)

	// revoke Certificate1 certificate
	utils.ProposeAndApproveCertificateRevocation(
		setup,
		rootCertificate1.Subject,
		rootCertificate1.SubjectKeyId,
		"",
	)

	// Check: Certificate1 is revoked
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.RevokedCertificatesKeyPrefix, Count: 2},
			{Key: types.UniqueCertificateKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateRevocationKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate1, indexes)
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate2, indexes)
}

func TestHandler_RevokeDaRootCert_BySerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	rootCertificate1 := utils.RootDaCertificateWithSameSubjectAndSKID1(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate1)

	rootCertificate2 := utils.RootDaCertificateWithSameSubjectAndSKID2(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate2)

	// revoke Certificate1 certificate
	utils.ProposeAndApproveCertificateRevocation(
		setup,
		rootCertificate1.Subject,
		rootCertificate1.SubjectKeyId,
		rootCertificate1.SerialNumber,
	)

	// Check: Certificate1 - RevokedCertificates - present
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.RevokedCertificatesKeyPrefix},
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix, Count: 1},
			{Key: types.AllCertificatesBySubjectKeyPrefix, Count: 1},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Count: 1},
			{Key: types.ApprovedCertificatesKeyPrefix, Count: 1},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix, Count: 1},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix, Count: 1},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateRevocationKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate1, indexes)
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate2, indexes)

	// revoke Certificate2 certificate
	utils.ProposeAndApproveCertificateRevocation(
		setup,
		rootCertificate2.Subject,
		rootCertificate2.SubjectKeyId,
		rootCertificate2.SerialNumber,
	)

	// Check: Certificate1 is revoked
	indexes = utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.RevokedCertificatesKeyPrefix, Count: 2},
			{Key: types.UniqueCertificateKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateRevocationKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate2, indexes)
}

func TestHandler_RevokeDaRootCert_RevokeChild(t *testing.T) {
	setup := utils.Setup(t)

	// add root x509 certificate
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate)

	// add intermediate x509 certificate
	intermediateCertificate := utils.IntermediateDaCertificate(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, intermediateCertificate)

	// propose revocation of x509 root certificate
	utils.ProposeRevokeDaRootCertificate(
		setup,
		setup.Trustee1,
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		rootCertificate.SerialNumber,
		true)

	// approve
	utils.ApproveRevokeDaRootCertificate(
		setup,
		setup.Trustee2,
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		rootCertificate.SerialNumber)

	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.RevokedCertificatesKeyPrefix},
			{Key: types.UniqueCertificateKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateRevocationKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)
	utils.CheckCertificateStateIndexes(t, setup, intermediateCertificate, indexes)
}

func TestHandler_RevokeDaRootCert_KeepChild(t *testing.T) {
	setup := utils.Setup(t)

	// add root x509 certificate
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate)

	// add intermediate x509 certificate
	intermediateCertificate := utils.IntermediateDaCertificate(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, intermediateCertificate)

	// propose revocation of x509 root certificate
	utils.ProposeRevokeDaRootCertificate(
		setup,
		setup.Trustee1,
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		rootCertificate.SerialNumber,
		false)

	// approve
	utils.ApproveRevokeDaRootCertificate(
		setup,
		setup.Trustee2,
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		rootCertificate.SerialNumber)

	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.RevokedCertificatesKeyPrefix},
			{Key: types.UniqueCertificateKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateRevocationKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)

	indexes = utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ApprovedRootCertificatesKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.RejectedCertificateKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, intermediateCertificate, indexes)
}

func TestHandler_RevokeDaRootCert_BySubjectAndSkid_TwoCertificatesWithSameSkid(t *testing.T) {
	setup := utils.Setup(t)

	// add root certificates
	rootCertificate1 := utils.RootDaCertWithSameSubjectKeyID1(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate1)

	rootCertificate2 := utils.RootDaCertificateWithSameSubjectKeyID2(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate2)

	// revoke Certificate1 certificate
	utils.ProposeAndApproveCertificateRevocation(
		setup,
		rootCertificate1.Subject,
		rootCertificate1.SubjectKeyId,
		"",
	)

	// Check state indexes
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.RevokedCertificatesKeyPrefix},
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},      // another cert with same SKID exists
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix}, // another cert with same SKID exist
		},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateRevocationKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate1, indexes)

	// second still exists
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
			{Key: types.RevokedCertificatesKeyPrefix},
			{Key: types.ProposedCertificateRevocationKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate2, indexes)
}

func TestHandler_RevokeDaRootCert_TwoThirdApprovalsNeeded(t *testing.T) {
	setup := utils.Setup(t)

	// add root x509 certificate
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate)

	// Create an array of trustee account from 1 to 50
	trusteeAccounts, totalAdditionalTrustees := setup.CreateNTrusteeAccounts()

	// We have 3 Trustees in test setup.
	twoThirds := int(math.Ceil(types.RootCertificateApprovalsPercent * float64(3+totalAdditionalTrustees)))

	// Trustee1 proposes to revoke the certificate
	utils.ProposeRevokeDaRootCertificate(
		setup,
		setup.Trustee1,
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		rootCertificate.SerialNumber,
		false)

	// Until we hit 2/3 of the total number of Trustees, we should not be able to revoke the certificate
	// We start the counter from 2 as the proposer is a trustee as well
	for i := 1; i < twoThirds-1; i++ {
		// approve the revocation
		utils.ApproveRevokeDaRootCertificate(
			setup,
			trusteeAccounts[i],
			rootCertificate.Subject,
			rootCertificate.SubjectKeyId,
			rootCertificate.SerialNumber)

		// check that the certificate is still not revoked
		indexes := utils.TestIndexes{
			Present: []utils.TestIndex{
				{Key: types.UniqueCertificateKeyPrefix},
				{Key: types.ProposedCertificateRevocationKeyPrefix},
				{Key: types.AllCertificatesKeyPrefix},
				{Key: types.AllCertificatesBySubjectKeyPrefix},
				{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
				{Key: types.ApprovedCertificatesKeyPrefix},
				{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
				{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
				{Key: types.ApprovedRootCertificatesKeyPrefix},
			},
			Missing: []utils.TestIndex{
				{Key: types.ChildCertificatesKeyPrefix},
				{Key: types.ProposedCertificateKeyPrefix},
				{Key: types.RevokedCertificatesKeyPrefix},
			},
		}
		utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)
	}

	// One more revoke will revoke the certificate
	utils.ApproveRevokeDaRootCertificate(
		setup,
		setup.Trustee2,
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		rootCertificate.SerialNumber)

	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.RevokedCertificatesKeyPrefix},
			{Key: types.RevokedRootCertificatesKeyPrefix},
			{Key: types.UniqueCertificateKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateRevocationKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
		},
	}
	resolvedCertificates := utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)

	// Make sure all the approvals are present
	for i := 1; i < twoThirds-1; i++ {
		require.Equal(t, resolvedCertificates.RevokedCertificates.Certs[0].HasApprovalFrom(trusteeAccounts[i].String()), true)
	}
	require.Equal(t, resolvedCertificates.RevokedCertificates.Certs[0].HasApprovalFrom(setup.Trustee1.String()), true)
	require.Equal(t, resolvedCertificates.RevokedCertificates.Certs[0].HasApprovalFrom(setup.Trustee2.String()), true)
}

// Error cases

func TestHandler_ApproveRevokeDaRootCert_ByNotTrustee(t *testing.T) {
	setup := utils.Setup(t)

	// propose and approve x509 root certificate
	rootCert := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// propose revocation of x509 root certificate
	utils.ProposeRevokeDaRootCertificate(
		setup,
		setup.Trustee1,
		rootCert.Subject,
		rootCert.SubjectKeyId,
		"",
		false)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.Vendor,
		dclauthtypes.CertificationCenter,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := utils.GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, 1)

		// approve
		approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
			accAddress.String(),
			rootCert.Subject,
			rootCert.SubjectKeyId,
			rootCert.SerialNumber,
			testconstants.Info)
		_, err := setup.Handler(setup.Ctx, approveRevokeX509RootCert)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
	}
}

func TestHandler_ApproveRevokeDaRootCert_ProposedRevocationDoesNotExist(t *testing.T) {
	setup := utils.Setup(t)

	// propose and approve x509 root certificate
	rootCert := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// approve revocation of x509 root certificate
	approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
		setup.Trustee1.String(),
		rootCert.Subject,
		rootCert.SubjectKeyId,
		rootCert.SerialNumber,
		testconstants.Info)
	_, err := setup.Handler(setup.Ctx, approveRevokeX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrProposedCertificateRevocationDoesNotExist.Is(err))
}

func TestHandler_ApproveRevokeDaRootCert_BySerialNumber_ProposedRevocationDoesNotExist(t *testing.T) {
	setup := utils.Setup(t)

	// propose and approve x509 root certificate
	rootCert := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// propose certificate revocation
	utils.ProposeAndApproveCertificateRevocation(
		setup,
		rootCert.Subject,
		rootCert.SubjectKeyId,
		"",
	)

	// approve revocation of x509 root certificate
	approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
		setup.Trustee1.String(),
		rootCert.Subject,
		rootCert.SubjectKeyId,
		"invalid",
		testconstants.Info)
	_, err := setup.Handler(setup.Ctx, approveRevokeX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrProposedCertificateRevocationDoesNotExist.Is(err))
}

func TestHandler_ApproveRevokeDaRootCert_Twice(t *testing.T) {
	setup := utils.Setup(t)

	// propose and approve x509 root certificate
	rootCert := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// propose revocation of x509 root certificate
	utils.ProposeRevokeDaRootCertificate(
		setup,
		setup.Trustee1,
		rootCert.Subject,
		rootCert.SubjectKeyId,
		rootCert.SerialNumber,
		false,
	)

	// approve revocation by the same trustee
	approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
		setup.Trustee1.String(),
		rootCert.Subject,
		rootCert.SubjectKeyId,
		rootCert.SerialNumber,
		testconstants.Info)
	_, err := setup.Handler(setup.Ctx, approveRevokeX509RootCert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}
