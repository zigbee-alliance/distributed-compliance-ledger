package tests

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/tests/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// Main

func TestHandler_RevokeDaIntermediateCert_BySubjectAndSKID(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.RootCertWithVidVid)

	// propose and approve x509 root certificate
	rootCert := utils.RootDaCertificateWithSameSubjectAndSKID1(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// Add intermediate certificates
	testIntermediateCertificate1 := utils.IntermediateDaCertificateWithSameSubjectAndSKID1(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, testIntermediateCertificate1)

	testIntermediateCertificate2 := utils.IntermediateDaCertificateWithSameSubjectAndSKID2(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, testIntermediateCertificate2)

	// revoke only an intermediate certificate
	utils.RevokeDaIntermediateCertificate(
		setup,
		vendorAccAddress,
		testIntermediateCertificate1.Subject,
		testIntermediateCertificate1.SubjectKeyId,
		"",
		false)

	// intermediate and leaf are revoked
	indexes := utils.TestIndexes{
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
}

func TestHandler_RevokeDaIntermediateCert_BySerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.RootCertWithVidVid)

	// propose and approve x509 root certificate
	rootCert := utils.RootDaCertificateWithSameSubjectAndSKID1(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// Add intermediate certificates
	testIntermediateCertificate1 := utils.IntermediateDaCertificateWithSameSubjectAndSKID1(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, testIntermediateCertificate1)

	testIntermediateCertificate2 := utils.IntermediateDaCertificateWithSameSubjectAndSKID2(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, testIntermediateCertificate2)

	// revoke only an intermediate certificate
	utils.RevokeDaIntermediateCertificate(
		setup,
		vendorAccAddress,
		testIntermediateCertificate1.Subject,
		testIntermediateCertificate1.SubjectKeyId,
		testIntermediateCertificate1.SerialNumber,
		false)

	// check indexes for intermediate certificates
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
		},
		Missing: []utils.TestIndex{},
	}
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate1, indexes)
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate2, indexes)

	// revoke intermediate and leaf certificates
	utils.RevokeDaIntermediateCertificate(
		setup,
		vendorAccAddress,
		testIntermediateCertificate2.Subject,
		testIntermediateCertificate2.SubjectKeyId,
		testIntermediateCertificate2.SerialNumber,
		false)

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
}

func TestHandler_RevokeDaIntermediateCert_BySubjectAndSKID_KeepChild(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.RootCertWithVidVid)

	// propose and approve x509 root certificate
	rootCert := utils.RootDaCertificateWithSameSubjectAndSKID1(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// Add intermediate certificates
	intermediateCertificate1 := utils.IntermediateDaCertificateWithSameSubjectAndSKID1(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, intermediateCertificate1)

	intermediateCertificate2 := utils.IntermediateDaCertificateWithSameSubjectAndSKID2(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, intermediateCertificate2)

	// add leaf x509 certificate
	leafCertificate := utils.LeafDaCertificateWithSameSubjectAndSKID(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, leafCertificate)

	// revoke x509 certificate
	utils.RevokeDaIntermediateCertificate(
		setup,
		vendorAccAddress,
		intermediateCertificate1.Subject,
		intermediateCertificate1.SubjectKeyId,
		"",
		false)

	// leaf stays approved
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
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, leafCertificate, indexes)
}

func TestHandler_RevokeDaIntermediateCert_BySerialNumber_KeepChild(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.RootCertWithVidVid)

	// propose and approve x509 root certificate
	rootCert := utils.RootDaCertificateWithSameSubjectAndSKID1(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// Add intermediate certificates
	intermediateCertificate1 := utils.IntermediateDaCertificateWithSameSubjectAndSKID1(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, intermediateCertificate1)

	intermediateCertificate2 := utils.IntermediateDaCertificateWithSameSubjectAndSKID2(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, intermediateCertificate2)

	// add leaf x509 certificate
	leafCertificate := utils.LeafDaCertificateWithSameSubjectAndSKID(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, leafCertificate)

	// revoke x509 certificate
	utils.RevokeDaIntermediateCertificate(
		setup,
		vendorAccAddress,
		intermediateCertificate1.Subject,
		intermediateCertificate1.SubjectKeyId,
		intermediateCertificate1.SerialNumber,
		false)

	// leaf stays approved
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			// {Key: types.AllCertificatesBySubjectKeyPrefix, Count: 2},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			// {Key: types.ApprovedCertificatesBySubjectKeyPrefix, Count: 2},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, leafCertificate, indexes)
}

func TestHandler_RevokeDaIntermediateCert_BySubjectAndSKID_RevokeChild(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.RootCertWithVidVid)

	// propose and approve x509 root certificate
	rootCert := utils.RootDaCertificateWithSameSubjectAndSKID1(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// Add intermediate certificates
	intermediateCertificate1 := utils.IntermediateDaCertificateWithSameSubjectAndSKID1(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, intermediateCertificate1)

	intermediateCertificate2 := utils.IntermediateDaCertificateWithSameSubjectAndSKID2(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, intermediateCertificate2)

	// add leaf x509 certificate
	leafCertificate := utils.LeafDaCertificateWithSameSubjectAndSKID(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, leafCertificate)

	// revoke x509 certificate
	utils.RevokeDaIntermediateCertificate(
		setup,
		vendorAccAddress,
		intermediateCertificate1.Subject,
		intermediateCertificate1.SubjectKeyId,
		"",
		true)

	// leaf stays approved
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, leafCertificate, indexes)
}

func TestHandler_RevokeDaIntermediateCert_BySerialNumber_RevokeChild(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.RootCertWithVidVid)

	// propose and approve x509 root certificate
	rootCert := utils.RootDaCertificateWithSameSubjectAndSKID1(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// Add intermediate certificates
	intermediateCertificate1 := utils.IntermediateDaCertificateWithSameSubjectAndSKID1(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, intermediateCertificate1)

	intermediateCertificate2 := utils.IntermediateDaCertificateWithSameSubjectAndSKID2(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, intermediateCertificate2)

	// add leaf x509 certificate
	leafCertificate := utils.LeafDaCertificateWithSameSubjectAndSKID(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, leafCertificate)

	// revoke x509 certificate
	utils.RevokeDaIntermediateCertificate(
		setup,
		vendorAccAddress,
		intermediateCertificate1.Subject,
		intermediateCertificate1.SubjectKeyId,
		intermediateCertificate1.SerialNumber,
		true)

	// leaf stays approved
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			// {Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			// {Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, leafCertificate, indexes)
}

func TestHandler_RevokeDaIntermediateCert_BySubjectAndSKID_ParentExist(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.RootCertWithVidVid)

	// propose and approve x509 root certificate
	rootCert := utils.RootDaCertificateWithSameSubjectAndSKID1(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// Add intermediate certificates
	intermediateCertificate1 := utils.IntermediateDaCertificateWithSameSubjectAndSKID1(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, intermediateCertificate1)

	intermediateCertificate2 := utils.IntermediateDaCertificateWithSameSubjectAndSKID2(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, intermediateCertificate2)

	// revoke x509 certificate
	utils.RevokeDaIntermediateCertificate(
		setup,
		vendorAccAddress,
		intermediateCertificate1.Subject,
		intermediateCertificate1.SubjectKeyId,
		"",
		false)

	// leaf stays approved
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
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCert, indexes)
}

func TestHandler_RevokeDaIntermediateCert_BySerialNumber_ParentExist(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.RootCertWithVidVid)

	// propose and approve x509 root certificate
	rootCert := utils.RootDaCertificateWithSameSubjectAndSKID1(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// Add intermediate certificates
	intermediateCertificate1 := utils.IntermediateDaCertificateWithSameSubjectAndSKID1(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, intermediateCertificate1)

	intermediateCertificate2 := utils.IntermediateDaCertificateWithSameSubjectAndSKID2(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, intermediateCertificate2)

	// revoke x509 certificate
	utils.RevokeDaIntermediateCertificate(
		setup,
		vendorAccAddress,
		intermediateCertificate1.Subject,
		intermediateCertificate1.SubjectKeyId,
		intermediateCertificate1.SerialNumber,
		false)

	// leaf stays approved
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
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCert, indexes)
}

func TestHandler_RevokeDaIntermediateCert_ByNotOwnerButSameVendor(t *testing.T) {
	setup := utils.Setup(t)

	// store root certificate
	rootCert := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// add x509 certificate by first vendor account
	intermediateCertificate := utils.IntermediateDaCertificate(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, intermediateCertificate)

	// add second vendor account with VID = 1
	vendorAccAddress2 := setup.CreateVendorAccount(testconstants.Vid)

	// revoke x509 certificate by second vendor account
	utils.RevokeDaIntermediateCertificate(
		setup,
		vendorAccAddress2,
		intermediateCertificate.Subject,
		intermediateCertificate.SubjectKeyId,
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

func TestHandler_RevokeDaIntermediateCert_CertificateDoesNotExist(t *testing.T) {
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

func TestHandler_RevokeDaIntermediateCert_CertificateDoesNotExistBySerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	// propose and approve x509 root certificate
	rootCert := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// Add intermediate certificate
	intermediateCertificate := utils.IntermediateDaCertificate(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, intermediateCertificate)

	// revoke x509 certificate
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		setup.Vendor1.String(),
		intermediateCertificate.Subject,
		intermediateCertificate.SubjectKeyId,
		"invalid",
		false,
		testconstants.Info,
	)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_RevokeDaIntermediateCert_ForRootCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// propose and approve x509 root certificate
	rootCert := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// revoke x509 root certificate
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		setup.Vendor1.String(),
		rootCert.Subject,
		rootCert.SubjectKeyId,
		rootCert.SerialNumber,
		false,
		testconstants.Info,
	)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrInappropriateCertificateType.Is(err))
}

func TestHandler_RevokeDaIntermediateCert_ByVendorWithOtherVid(t *testing.T) {
	setup := utils.Setup(t)

	// propose and approve x509 root certificate
	rootCert := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// Add intermediate certificate
	intermediateCertificate := utils.IntermediateDaCertificate(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, intermediateCertificate)

	// add second vendor account with VID = 1000
	vendorAccAddress2 := setup.CreateVendorAccount(testconstants.VendorID1)

	// revoke x509 certificate by second vendor account
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		vendorAccAddress2.String(),
		intermediateCertificate.Subject,
		intermediateCertificate.SubjectKeyId,
		intermediateCertificate.SerialNumber,
		false,
		testconstants.Info,
	)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_RevokeDaIntermediateCert_SenderNotVendor(t *testing.T) {
	setup := utils.Setup(t)

	// propose and approve x509 root certificate
	rootCert := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// Add intermediate certificate
	intermediateCertificate := utils.IntermediateDaCertificate(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, intermediateCertificate)

	// Try to revoke By Trustee
	removeX509Cert := types.NewMsgRevokeX509Cert(
		setup.Trustee1.String(),
		intermediateCertificate.Subject,
		intermediateCertificate.SubjectKeyId,
		intermediateCertificate.SerialNumber,
		false,
		testconstants.Info,
	)
	_, err := setup.Handler(setup.Ctx, removeX509Cert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}
