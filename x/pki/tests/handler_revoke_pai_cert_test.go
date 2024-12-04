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

func TestHandler_RevokeDaIntermediateCert(t *testing.T) {
	setup := utils.Setup(t)

	// propose and approve x509 root certificate
	rootCertificate := utils.CreateTestRootCert()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, &rootCertificate)

	// Add intermediate certificate
	intermediateCertificate := utils.CreateTestIntermediateCert()
	utils.AddDaIntermediateCertificate(setup, setup.Vendor1, intermediateCertificate.PEM)

	// revoke intermediate certificate
	utils.RevokeDaIntermediateCertificate(
		setup,
		setup.Vendor1,
		intermediateCertificate.Subject,
		intermediateCertificate.SubjectKeyID,
		"",
		false)

	// Check: Certificate is revoked
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
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
	utils.CheckCertificateStateIndexes(t, setup, intermediateCertificate, indexes)

	// Check: Root stays approved
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
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)
}

func TestHandler_RevokeX509Cert_ForTree(t *testing.T) {
	setup := utils.Setup(t)

	// add root x509 certificate
	rootCertificate := utils.CreateTestRootCert()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, &rootCertificate)

	// add intermediate x509 certificate
	intermediateCertificate := utils.CreateTestIntermediateCert()
	utils.AddDaIntermediateCertificate(setup, setup.Vendor1, intermediateCertificate.PEM)

	// add leaf x509 certificate
	leafCertificate := utils.CreateTestLeafCert()
	utils.AddDaIntermediateCertificate(setup, setup.Vendor1, leafCertificate.PEM)

	// revoke x509 certificate
	utils.RevokeDaIntermediateCertificate(
		setup,
		setup.Vendor1,
		intermediateCertificate.Subject,
		intermediateCertificate.SubjectKeyID,
		"",
		true)

	// root stays approved
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
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)

	// intermediate and leaf are revoked
	indexes = utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
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
	utils.CheckCertificateStateIndexes(t, setup, intermediateCertificate, indexes)
	utils.CheckCertificateStateIndexes(t, setup, leafCertificate, indexes)
}

func TestHandler_RevokeX509Cert_BySerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.RootCertWithVidVid)

	// propose and approve x509 root certificate
	rootCert := utils.CreateTestRootCertWithSameSubjectAndSkid1()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, &rootCert)

	// Add intermediate certificates
	testIntermediateCertificate1 := utils.CreateTestIntermediateCertWithSameSubjectAndSKID1()
	utils.AddDaIntermediateCertificate(setup, vendorAccAddress, testIntermediateCertificate1.PEM)

	testIntermediateCertificate2 := utils.CreateTestIntermediateCertWithSameSubjectAndSKID2()
	utils.AddDaIntermediateCertificate(setup, vendorAccAddress, testIntermediateCertificate2.PEM)

	// Add a leaf certificate
	testLeafCertificate := utils.CreateTestLeafCertWithSameSubjectAndSKID()
	utils.AddDaIntermediateCertificate(setup, vendorAccAddress, testLeafCertificate.PEM)

	// get certificates for further comparison
	allCerts := setup.Keeper.GetAllApprovedCertificates(setup.Ctx)
	require.Equal(t, 3, len(allCerts))
	require.Equal(t, 4, len(allCerts[0].Certs)+len(allCerts[1].Certs)+len(allCerts[2].Certs))

	// revoke only an intermediate certificate
	utils.RevokeDaIntermediateCertificate(
		setup,
		vendorAccAddress,
		testIntermediateCertificate1.Subject,
		testIntermediateCertificate1.SubjectKeyID,
		testIntermediateCertificate1.SerialNumber,
		false)

	// check indexes for intermediate certificates
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix, Count: 2}, // inter + leaf
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix, Count: 2}, // inter + leaf
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
		},
		Missing: []utils.TestIndex{},
	}
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate1, indexes)
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate2, indexes)

	// check indexes for leaf
	indexes = utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, testLeafCertificate, indexes)

	// revoke intermediate and leaf certificates
	utils.RevokeDaIntermediateCertificate(
		setup,
		vendorAccAddress,
		testIntermediateCertificate2.Subject,
		testIntermediateCertificate2.SubjectKeyID,
		testIntermediateCertificate2.SerialNumber,
		true)

	// intermediate and leaf are revoked
	indexes = utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix, Count: 2},
		},
		Missing: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate1, indexes)
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate2, indexes)

	// intermediate and leaf are revoked
	indexes = utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix, Count: 1},
		},
		Missing: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, testLeafCertificate, indexes)
}

// Extra cases

func TestHandler_RevokeX509Cert_ByNotOwnerButSameVendor(t *testing.T) {
	setup := utils.Setup(t)

	// store root certificate
	rootCert := utils.CreateTestRootCert()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, &rootCert)

	// add x509 certificate by first vendor account
	intermediateCertificate := utils.CreateTestIntermediateCert()
	utils.AddDaIntermediateCertificate(setup, setup.Vendor1, intermediateCertificate.PEM)

	// add second vendor account with VID = 1
	vendorAccAddress2 := setup.CreateVendorAccount(testconstants.Vid)

	// revoke x509 certificate by second vendor account
	utils.RevokeDaIntermediateCertificate(
		setup,
		vendorAccAddress2,
		intermediateCertificate.Subject,
		intermediateCertificate.SubjectKeyID,
		intermediateCertificate.SerialNumber,
		false)

	// Check: Certificate is revoked
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
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
	utils.CheckCertificateStateIndexes(t, setup, intermediateCertificate, indexes)
}

// Error cases

func TestHandler_RevokeX509Cert_CertificateDoesNotExist(t *testing.T) {
	setup := utils.Setup(t)

	// revoke x509 certificate
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		setup.Vendor1.String(),
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		testconstants.IntermediateSerialNumber,
		false,
		testconstants.Info,
	)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_RevokeX509Cert_CertificateDoesNotExistBySerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	// propose and approve x509 root certificate
	rootCert := utils.CreateTestRootCert()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, &rootCert)

	// Add intermediate certificate
	addIntermediateX509Cert := types.NewMsgAddX509Cert(
		setup.Vendor1.String(),
		testconstants.IntermediateCertPem,
		testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addIntermediateX509Cert)
	require.NoError(t, err)

	// revoke x509 certificate
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		setup.Vendor1.String(),
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		"invalid",
		false,
		testconstants.Info,
	)
	_, err = setup.Handler(setup.Ctx, revokeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_RevokeX509Cert_ForRootCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// propose and approve x509 root certificate
	rootCert := utils.CreateTestRootCert()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, &rootCert)

	// revoke x509 root certificate
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		setup.Vendor1.String(),
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.RootSerialNumber,
		false,
		testconstants.Info,
	)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrInappropriateCertificateType.Is(err))
}

func TestHandler_RevokeX509Cert_ByOtherVendor(t *testing.T) {
	setup := utils.Setup(t)

	// store root certificate
	rootCertificate := utils.RootCertificate(setup.Trustee1)
	setup.Keeper.AddAllCertificate(setup.Ctx, rootCertificate)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// add x509 certificate by first vendor account
	addX509Cert := types.NewMsgAddX509Cert(setup.Vendor1.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// add second vendor account with VID = 1000
	vendorAccAddress2 := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID1)

	// revoke x509 certificate by second vendor account
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		vendorAccAddress2.String(),
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		testconstants.IntermediateSerialNumber,
		false,
		testconstants.Info,
	)
	_, err = setup.Handler(setup.Ctx, revokeX509Cert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_RevokeX509Cert_SenderNotVendor(t *testing.T) {
	setup := utils.Setup(t)

	// store root certificate
	rootCert := utils.CreateTestRootCertWithVid()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, &rootCert)

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.RootCertWithVidVid)

	// add x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(vendorAccAddress.String(), testconstants.IntermediateCertWithVid1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	removeX509Cert := types.NewMsgRevokeX509Cert(
		setup.Trustee1.String(),
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		testconstants.IntermediateSerialNumber,
		false,
		testconstants.Info,
	)
	_, err = setup.Handler(setup.Ctx, removeX509Cert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}
