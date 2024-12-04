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

// Propose

func TestHandler_ProposeRevokeDaRootCert(t *testing.T) {
	setup := utils.Setup(t)

	// propose x509 root certificate by `setup.Trustee` and approve by another trustee
	rootCertificate := utils.RootCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate)

	// propose revocation of x509 root certificate by `setup.Trustee`
	utils.ProposeRevokeDaRootCertificate(
		setup,
		setup.Trustee1,
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		rootCertificate.SerialNumber,
		false)

	// Check: Certificate is proposed to revoke
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
	resolvedCertificates := utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)

	// additional check
	require.True(t, resolvedCertificates.ProposedRevocation.HasRevocationFrom(setup.Trustee1.String()))
}

func TestHandler_ProposeRevokeDaRootCert_ByTrusteeNotOwner(t *testing.T) {
	setup := utils.Setup(t)

	// propose x509 root certificate by `setup.Trustee` and approve by another trustee
	rootCertificate := utils.RootCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate)

	// add another trustee
	anotherTrustee := setup.CreateTrusteeAccount(1)

	// propose revocation of x509 root certificate by new trustee
	utils.ProposeRevokeDaRootCertificate(
		setup,
		anotherTrustee,
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		rootCertificate.SerialNumber,
		false)

	// Check: Certificate is proposed to revoke
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
	resolvedCertificates := utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)

	// additional check
	require.True(t, resolvedCertificates.ProposedRevocation.HasRevocationFrom(anotherTrustee.String()))
}

// Propose + Approve

func TestHandler_RevokeDaRootCert(t *testing.T) {
	setup := utils.Setup(t)

	// propose x509 root certificate by `setup.Trustee` and approve by another trustee
	rootCertificate := utils.RootCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate)

	// revoke certificate
	proposeAndApproveCertificateRevocation(
		t,
		setup,
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		"",
	)

	// Check state indexes
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
}

func TestHandler_RevokeDaRootCert_BySubjectAndSkid_WhenTwoCertsWithSameSkidExist(t *testing.T) {
	setup := utils.Setup(t)

	// add root certificates
	rootCertificate1 := utils.PAACertWithSameSubjectID1(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate1)

	rootCertificate2 := utils.PAACertWithSameSubjectID2(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate2)

	// revoke Certificate1 certificate
	proposeAndApproveCertificateRevocation(
		t,
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

func TestHandler_RevokeDaRootCert_BySerialNumber_WhenTwoCertsWithSameSubjectAndSkidExist(t *testing.T) {
	setup := utils.Setup(t)

	rootCertificate1 := utils.RootCertWithSameSubjectAndSKID1(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate1)

	rootCertificate2 := utils.RootCertWithSameSubjectAndSKID2(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate2)

	// revoke Certificate1 certificate
	proposeAndApproveCertificateRevocation(
		t,
		setup,
		testconstants.RootCertWithSameSubjectAndSKIDSubject,
		testconstants.RootCertWithSameSubjectAndSKIDSubjectKeyID,
		testconstants.RootCertWithSameSubjectAndSKID1SerialNumber,
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
	proposeAndApproveCertificateRevocation(
		t,
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

func TestHandler_RevokeDaRootCert_TwoThirdApprovalsNeeded(t *testing.T) {
	setup := utils.Setup(t)

	// add root x509 certificate
	rootCertificate := utils.RootCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate)

	// root exists
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
		utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)
	}

	// One more revoke will revoke the certificate
	utils.ApproveRevokeDaRootCertificate(
		setup,
		setup.Trustee2,
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		rootCertificate.SerialNumber)

	indexes = utils.TestIndexes{
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
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)

	// Check: Revoked - present
	revokedCertificate, err := utils.QueryRevokedCertificates(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, testconstants.RootIssuer, revokedCertificate.Certs[0].Subject)
	require.Equal(t, testconstants.RootSerialNumber, revokedCertificate.Certs[0].SerialNumber)
	require.True(t, revokedCertificate.Certs[0].IsRoot)

	// Make sure all the approvals are present
	for i := 1; i < twoThirds-1; i++ {
		require.Equal(t, revokedCertificate.Certs[0].HasApprovalFrom(trusteeAccounts[i].String()), true)
	}
	require.Equal(t, revokedCertificate.Certs[0].HasApprovalFrom(setup.Trustee1.String()), true)
	require.Equal(t, revokedCertificate.Certs[0].HasApprovalFrom(setup.Trustee2.String()), true)
}

//nolint:funlen
func TestHandler_RevokeDaRootCert_ForTree(t *testing.T) {
	setup := utils.Setup(t)

	// add root x509 certificate
	rootCertificate := utils.RootCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate)

	// add intermediate x509 certificate
	intermediateCertificate := utils.IntermediateCertPem(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, setup.Vendor1, intermediateCertificate.PemCert)

	// add leaf x509 certificate
	leafCertificate := utils.LeafCertPem(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, setup.Vendor1, leafCertificate.PemCert)

	// propose revocation of x509 root certificate
	utils.ProposeRevokeDaRootCertificate(
		setup,
		setup.Trustee1,
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		"",
		true)

	// approve
	utils.ApproveRevokeDaRootCertificate(
		setup,
		setup.Trustee2,
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		"")

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
	utils.CheckCertificateStateIndexes(t, setup, leafCertificate, indexes)
}

// Extra cases

func TestHandler_ApproveRevokeX509RootCert_ForNotEnoughApprovals(t *testing.T) {
	setup := utils.Setup(t)

	// propose and approve x509 root certificate
	rootCertificate := utils.RootCertificate(setup.Trustee1)
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

// Error cases

func TestHandler_ProposeRevokeX509RootCert_ByNotTrustee(t *testing.T) {
	setup := utils.Setup(t)

	// propose and approve x509 root certificate
	rootCert := utils.RootCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.Vendor,
		dclauthtypes.CertificationCenter,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := utils.GenerateAccAddress()
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
	setup := utils.Setup(t)

	// propose revocation of not existing certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(),
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.RootSerialNumber,
		false,
		testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_ProposeRevokeX509RootCert_CertificateDoesNotExistBySerialNumber(t *testing.T) {
	setup := utils.Setup(t)
	// propose and approve x509 root certificate
	rootCert := utils.RootCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

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
	setup := utils.Setup(t)

	// propose x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(
		setup.Trustee1.String(),
		testconstants.RootCertPem,
		testconstants.Info,
		testconstants.Vid,
		testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// check that proposed certificate is present
	proposedCertificate, _ := utils.QueryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NotNil(t, proposedCertificate)

	// propose revocation of proposed root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(),
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.RootSerialNumber,
		false,
		testconstants.Info)
	_, err = setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_ProposeRevokeX509RootCert_ProposedRevocationAlreadyExists(t *testing.T) {
	setup := utils.Setup(t)

	// propose and approve x509 root certificate
	rootCert := utils.RootCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// propose revocation of x509 root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.RootSerialNumber, false, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	// store another trustee
	anotherTrustee := setup.CreateTrusteeAccount(1)

	// propose revocation of the same x509 root certificate again
	proposeRevokeX509RootCert = types.NewMsgProposeRevokeX509RootCert(
		anotherTrustee.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.RootSerialNumber, false, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrProposedCertificateRevocationAlreadyExists.Is(err))
}

func TestHandler_ProposeRevokeX509RootCert_ForNonRootCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// store x509 root certificate
	rootCertificate := utils.RootCertificate(setup.Trustee1)
	setup.Keeper.AddAllCertificate(setup.Ctx, rootCertificate)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.Vid)

	// store x509 intermediate certificate
	addX509Cert := types.NewMsgAddX509Cert(vendorAccAddress.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// propose revocation of x509 intermediate certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(),
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		testconstants.RootSerialNumber,
		false,
		testconstants.Info)
	_, err = setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrInappropriateCertificateType.Is(err))
}

func TestHandler_ApproveRevokeX509RootCert_ByNotTrustee(t *testing.T) {
	setup := utils.Setup(t)

	// propose and approve x509 root certificate
	rootCert := utils.RootCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// propose revocation of x509 root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(),
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.RootSerialNumber,
		false,
		testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

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
			testconstants.RootSubject,
			testconstants.RootSubjectKeyID,
			testconstants.RootSerialNumber,
			testconstants.Info)
		_, err = setup.Handler(setup.Ctx, approveRevokeX509RootCert)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
	}
}

func TestHandler_ApproveRevokeX509RootCert_ProposedRevocationDoesNotExist(t *testing.T) {
	setup := utils.Setup(t)

	// propose and approve x509 root certificate
	rootCert := utils.RootCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// approve revocation of x509 root certificate
	approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
		setup.Trustee1.String(),
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.RootSerialNumber,
		testconstants.Info)
	_, err := setup.Handler(setup.Ctx, approveRevokeX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrProposedCertificateRevocationDoesNotExist.Is(err))
}

func TestHandler_ApproveRevokeX509RootCert_Twice(t *testing.T) {
	setup := utils.Setup(t)

	// propose and approve x509 root certificate
	rootCert := utils.RootCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// propose revocation of x509 root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(),
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.RootSerialNumber,
		false,
		testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	// approve revocation by the same trustee
	approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
		setup.Trustee1.String(),
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.RootSerialNumber,
		testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveRevokeX509RootCert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_RevocationPointsByIssuerSubjectKeyID(t *testing.T) {
	setup := utils.Setup(t)

	vendorAcc := setup.CreateVendorAccount(65521)

	// propose x509 root certificate by account Trustee1
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(
		setup.Trustee1.String(),
		testconstants.PAACertWithNumericVid,
		testconstants.Info,
		testconstants.Vid,
		testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(),
		testconstants.PAACertWithNumericVidSubject,
		testconstants.PAACertWithNumericVidSubjectKeyID,
		testconstants.Info)
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
	setup := utils.Setup(t)

	vendorAcc := setup.CreateVendorAccount(65521)

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
	setup := utils.Setup(t)

	vendorAcc := setup.CreateVendorAccount(65521)

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

func proposeAndApproveCertificateRevocation(
	t *testing.T,
	setup *utils.TestSetup,
	subject string,
	subjectKeyID string,
	serialNumber string,
) {
	t.Helper()

	// revoke certificate
	revokeX509Cert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(),
		subject,
		subjectKeyID,
		serialNumber,
		false,
		testconstants.Info)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(t, err)

	aprRevokeX509Cert := types.NewMsgApproveRevokeX509RootCert(
		setup.Trustee2.String(),
		subject,
		subjectKeyID,
		serialNumber,
		testconstants.Info)
	_, err = setup.Handler(setup.Ctx, aprRevokeX509Cert)
	require.NoError(t, err)
}
