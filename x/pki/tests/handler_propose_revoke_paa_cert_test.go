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

func TestHandler_ProposeRevokeDaRootCert(t *testing.T) {
	setup := utils.Setup(t)

	// add root certificate
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate)

	// propose revocation of root certificate by the same trustee
	utils.ProposeRevokeDaRootCertificate(
		setup,
		setup.Trustee1,
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		rootCertificate.SerialNumber,
		false)

	// Check state indexes - certificate is proposed to revoke (but stays approved)
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.ProposedCertificateRevocationKeyPrefix},
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
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
			{Key: types.RevokedRootCertificatesKeyPrefix},
		},
	}
	resolvedCertificates := utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)

	// additional check - revocation approval exists
	require.True(t, resolvedCertificates.ProposedRevocation.HasRevocationFrom(setup.Trustee1.String()))
}

func TestHandler_ProposeRevokeDaRootCert_TwoCertificates(t *testing.T) {
	setup := utils.Setup(t)

	// add two root certificates
	rootCertificate1 := utils.RootDaCertificateWithSameSubjectAndSKID1(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate1)

	rootCertificate2 := utils.RootDaCertificateWithSameSubjectAndSKID2(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate2)

	// propose revocation of first certificate by `setup.Trustee`
	utils.ProposeRevokeDaRootCertificate(
		setup,
		setup.Trustee1,
		rootCertificate1.Subject,
		rootCertificate1.SubjectKeyId,
		rootCertificate1.SerialNumber,
		false)

	// Check state indexes - certificate1 is proposed to revoke (but stays approved)
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.ProposedCertificateRevocationKeyPrefix, Count: 1},
			{Key: types.UniqueCertificateKeyPrefix, Count: 1},
			{Key: types.AllCertificatesKeyPrefix, Count: 2},
			{Key: types.AllCertificatesBySubjectKeyPrefix, Count: 1},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.ApprovedCertificatesKeyPrefix, Count: 2},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix, Count: 1},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.ApprovedRootCertificatesKeyPrefix, Count: 1},
		},
		Missing: []utils.TestIndex{
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
			{Key: types.RevokedRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate1, indexes)

	// Check state indexes - certificate2 is not proposed to revoke
	indexes = utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix, Count: 1},
			{Key: types.AllCertificatesKeyPrefix, Count: 2},
			{Key: types.AllCertificatesBySubjectKeyPrefix, Count: 1},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.ApprovedCertificatesKeyPrefix, Count: 2},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix, Count: 1},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.ApprovedRootCertificatesKeyPrefix, Count: 1},
		},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateRevocationKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate2, indexes)
}

func TestHandler_ProposeRevokeDaRootCert_KeepChild(t *testing.T) {
	setup := utils.Setup(t)

	// add root certificate
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate)

	// add intermediate certificate
	intermediateCertificate := utils.IntermediateDaCertificate(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, intermediateCertificate)

	// propose revocation of root certificate
	utils.ProposeRevokeDaRootCertificate(
		setup,
		setup.Trustee1,
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		rootCertificate.SerialNumber,
		false)

	// Check state indexes - intermediate certificates stays approved
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateRevocationKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, intermediateCertificate, indexes)
}

func TestHandler_ProposeRevokeDaRootCert_RevokeChild(t *testing.T) {
	setup := utils.Setup(t)

	// add root certificate
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate)

	// add intermediate certificate
	intermediateCertificate := utils.IntermediateDaCertificate(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, intermediateCertificate)

	// propose revocation of root certificate
	utils.ProposeRevokeDaRootCertificate(
		setup,
		setup.Trustee1,
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		rootCertificate.SerialNumber,
		true)

	// Check state indexes - intermediate stays approved - not affected at propose step
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateRevocationKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, intermediateCertificate, indexes)
}

func TestHandler_ProposeRevokeDaRootCert_ByTrusteeNotOwner(t *testing.T) {
	setup := utils.Setup(t)

	// propose root certificate by `setup.Trustee`
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate)

	// propose revocation of x509 root certificate by new trustee
	utils.ProposeRevokeDaRootCertificate(
		setup,
		setup.Trustee3,
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		rootCertificate.SerialNumber,
		false)

	// Check state indexes - certificate is proposed to revoke
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.ProposedCertificateRevocationKeyPrefix},
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
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	resolvedCertificates := utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)

	// additional check
	require.True(t, resolvedCertificates.ProposedRevocation.HasRevocationFrom(setup.Trustee3.String()))
}

// Error cases

func TestHandler_ProposeRevokeDaRootCert_ByNotTrustee(t *testing.T) {
	setup := utils.Setup(t)

	// propose and approve x509 root certificate
	rootCert := utils.RootDaCertificate(setup.Trustee1)
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
			accAddress.String(),
			rootCert.Subject,
			rootCert.SubjectKeyId,
			rootCert.SerialNumber,
			false,
			testconstants.Info)
		_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
	}
}

func TestHandler_ProposeRevokeDaRootCert_CertificateDoesNotExist(t *testing.T) {
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

func TestHandler_ProposeRevokeDaRootCert_CertificateDoesNotExistBySerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	// propose and approve x509 root certificate
	rootCert := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// revoke x509 certificate
	revokeX509Cert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(),
		rootCert.Subject,
		rootCert.RootSubjectKeyId,
		"invalid",
		false,
		testconstants.Info,
	)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_ProposeRevokeDaRootCert_ForProposedCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// propose x509 root certificate
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeDaRootCertificate(setup, rootCertificate)

	// propose revocation of proposed root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(),
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		rootCertificate.SerialNumber,
		false,
		testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_ProposeRevokeDaRootCert_ProposedRevocationAlreadyExists(t *testing.T) {
	setup := utils.Setup(t)

	// propose and approve x509 root certificate
	rootCert := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// propose revocation of x509 root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(),
		rootCert.Subject,
		rootCert.SubjectKeyId,
		rootCert.SerialNumber,
		false,
		testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	// propose revocation of the same x509 root certificate again
	proposeRevokeX509RootCert = types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee2.String(),
		rootCert.Subject,
		rootCert.SubjectKeyId,
		rootCert.SerialNumber,
		false,
		testconstants.Info)
	_, err = setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrProposedCertificateRevocationAlreadyExists.Is(err))
}

func TestHandler_ProposeRevokeDaRootCert_ForNonRootCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// add DA root certificate
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate)

	// add DA PAI certificate
	testIntermediateCertificate := utils.IntermediateDaCertificate(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, testIntermediateCertificate)

	// propose revocation of x509 intermediate certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(),
		testIntermediateCertificate.Subject,
		testIntermediateCertificate.SubjectKeyId,
		testIntermediateCertificate.SerialNumber,
		false,
		testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrInappropriateCertificateType.Is(err))
}
