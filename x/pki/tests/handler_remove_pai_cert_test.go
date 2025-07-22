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

func TestHandler_RemoveDaIntermediateCert_BySubjectAndSKID(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.RootCertWithVidVid)

	// Add root certificate
	rootCert := utils.RootDaCertificateWithSameSubjectAndSKID1(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// Add two intermediate certificates
	testIntermediateCertificate1 := utils.IntermediateDaCertificateWithSameSubjectAndSKID1(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, testIntermediateCertificate1)

	testIntermediateCertificate2 := utils.IntermediateDaCertificateWithSameSubjectAndSKID2(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, testIntermediateCertificate2)

	// remove all intermediate certificates
	utils.RemoveDaIntermediateCertificate(
		setup,
		vendorAccAddress,
		testIntermediateCertificate1.Subject,
		testIntermediateCertificate1.SubjectKeyId,
		"")

	// Check state indexes - intermediate certificate are removed
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{},
		Missing: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
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
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate1, indexes)
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate2, indexes)
}

func TestHandler_RemoveDaIntermediateCert_BySerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.RootCertWithVidVid)

	// Add root certificate
	rootCert := utils.RootDaCertificateWithSameSubjectAndSKID1(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// Add two intermediate certificates
	testIntermediateCertificate1 := utils.IntermediateDaCertificateWithSameSubjectAndSKID1(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, testIntermediateCertificate1)

	testIntermediateCertificate2 := utils.IntermediateDaCertificateWithSameSubjectAndSKID2(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, testIntermediateCertificate2)

	// remove intermediate certificate by serial number
	utils.RemoveDaIntermediateCertificate(
		setup,
		vendorAccAddress,
		testIntermediateCertificate1.Subject,
		testIntermediateCertificate1.SubjectKeyId,
		testIntermediateCertificate1.SerialNumber)

	// Check state indexes - intermediate certificate1 removed but there is another approved
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate1, indexes)

	// Check state indexes - intermediate certificate2 approved (all the same but also UniqueCertificate exists)
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
		Missing: []utils.TestIndex{},
	}
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate2, indexes)
}

func TestHandler_RemoveDaIntermediateCert_BySubjectAndSKID_ParentExist(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.RootCertWithVidVid)

	// Add root certificate
	rootCert := utils.RootDaCertificateWithSameSubjectAndSKID1(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// Add two intermediate certificates
	testIntermediateCertificate1 := utils.IntermediateDaCertificateWithSameSubjectAndSKID1(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, testIntermediateCertificate1)

	testIntermediateCertificate2 := utils.IntermediateDaCertificateWithSameSubjectAndSKID2(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, testIntermediateCertificate2)

	// remove all intermediate certificates
	utils.RemoveDaIntermediateCertificate(
		setup,
		vendorAccAddress,
		testIntermediateCertificate1.Subject,
		testIntermediateCertificate1.SubjectKeyId,
		"")

	// Check state indexes - parent stays approved
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
	utils.CheckCertificateStateIndexes(t, setup, rootCert, indexes)
}

func TestHandler_RemoveDaIntermediateCert_BySerialNumber_ParentExist(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.RootCertWithVidVid)

	// Add root certificate
	rootCert := utils.RootDaCertificateWithSameSubjectAndSKID1(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// Add two intermediate certificates
	testIntermediateCertificate1 := utils.IntermediateDaCertificateWithSameSubjectAndSKID1(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, testIntermediateCertificate1)

	testIntermediateCertificate2 := utils.IntermediateDaCertificateWithSameSubjectAndSKID2(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, testIntermediateCertificate2)

	// remove intermediate certificate by serial number
	utils.RemoveDaIntermediateCertificate(
		setup,
		vendorAccAddress,
		testIntermediateCertificate1.Subject,
		testIntermediateCertificate1.SubjectKeyId,
		testIntermediateCertificate1.SerialNumber)

	// Check state indexes - parent stays approved
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
	utils.CheckCertificateStateIndexes(t, setup, rootCert, indexes)
}

func TestHandler_RemoveDaIntermediateCert_BySubjectAndSKID_RevokedCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// Add root certificate
	rootCert := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// Add intermediate certificates
	testIntermediateCertificate := utils.IntermediateDaCertificate(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, testIntermediateCertificate)

	// Revoke intermediate certificate
	utils.RevokeDaIntermediateCertificate(
		setup,
		setup.Vendor1,
		testIntermediateCertificate.Subject,
		testIntermediateCertificate.SubjectKeyId,
		"",
		false)

	// Remove intermediate certificate
	utils.RemoveDaIntermediateCertificate(
		setup,
		setup.Vendor1,
		testIntermediateCertificate.Subject,
		testIntermediateCertificate.SubjectKeyId,
		testIntermediateCertificate.SerialNumber)

	// Check state indexes - certificate is removed
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{},
		Missing: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate, indexes)
}

func TestHandler_RemoveDaIntermediateCert_BySerialNumber_RevokedCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// Add root certificate
	rootCert := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// Add intermediate certificates again
	testIntermediateCertificate := utils.IntermediateDaCertificate(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, testIntermediateCertificate)

	// revoke intermediate certificate by serial number
	utils.RevokeDaIntermediateCertificate(
		setup,
		setup.Vendor1,
		testIntermediateCertificate.Subject,
		testIntermediateCertificate.SubjectKeyId,
		testIntermediateCertificate.SerialNumber,
		false)

	// remove  intermediate certificate by serial number
	utils.RemoveDaIntermediateCertificate(
		setup,
		setup.Vendor1,
		testIntermediateCertificate.Subject,
		testIntermediateCertificate.SubjectKeyId,
		testIntermediateCertificate.SerialNumber)

	// Check state indexes - certificate is removed
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{},
		Missing: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate, indexes)
}

func TestHandler_RemoveDaIntermediateCert_BySubjectAndSKID_ApprovedChildExist(t *testing.T) {
	setup := utils.Setup(t)

	// add root certificate
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate)

	// add intermediate certificate
	intermediateCertificate := utils.IntermediateDaCertificate(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, intermediateCertificate)

	// add leaf certificate
	leafCertificate := utils.LeafCertificate(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, leafCertificate)

	// revoke intermediate certificate
	utils.RemoveDaIntermediateCertificate(
		setup,
		setup.Vendor1,
		intermediateCertificate.Subject,
		intermediateCertificate.SubjectKeyId,
		"")

	// check state indexes - leaf stays approved
	indexes := utils.TestIndexes{
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
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, leafCertificate, indexes)
}

func TestHandler_RemoveDaIntermediateCert_BySerialNumber_ApprovedChildExist(t *testing.T) {
	setup := utils.Setup(t)

	// add root certificate
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate)

	// add intermediate certificate
	intermediateCertificate := utils.IntermediateDaCertificate(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, intermediateCertificate)

	// add leaf certificate
	leafCertificate := utils.LeafCertificate(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, leafCertificate)

	// revoke intermediate certificate
	utils.RemoveDaIntermediateCertificate(
		setup,
		setup.Vendor1,
		intermediateCertificate.Subject,
		intermediateCertificate.SubjectKeyId,
		intermediateCertificate.SerialNumber)

	// check state indexes - leaf stays approved
	indexes := utils.TestIndexes{
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
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, leafCertificate, indexes)
}

func TestHandler_RemoveDaIntermediateCert_BySubjectAndSKID_RevokedChildExist(t *testing.T) {
	setup := utils.Setup(t)

	// add root certificate
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate)

	// add intermediate certificate
	intermediateCertificate := utils.IntermediateDaCertificate(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, intermediateCertificate)

	// add leaf certificate
	leafCertificate := utils.LeafCertificate(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, leafCertificate)

	// revoke leaf certificate
	utils.RevokeDaIntermediateCertificate(
		setup,
		setup.Vendor1,
		leafCertificate.Subject,
		leafCertificate.SubjectKeyId,
		"",
		true)

	// revoke intermediate certificate
	utils.RemoveDaIntermediateCertificate(
		setup,
		setup.Vendor1,
		intermediateCertificate.Subject,
		intermediateCertificate.SubjectKeyId,
		"")

	// check state indexes - leaf certificate stays revoked
	indexes := utils.TestIndexes{
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
	utils.CheckCertificateStateIndexes(t, setup, leafCertificate, indexes)
}

func TestHandler_RemoveDaIntermediateCert_BySerialNumber_RevokedChildExist(t *testing.T) {
	setup := utils.Setup(t)

	// add root certificate
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate)

	// add intermediate certificate
	intermediateCertificate := utils.IntermediateDaCertificate(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, intermediateCertificate)

	// add leaf certificate
	leafCertificate := utils.LeafCertificate(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, leafCertificate)

	// revoke certificate
	utils.RevokeDaIntermediateCertificate(
		setup,
		setup.Vendor1,
		leafCertificate.Subject,
		leafCertificate.SubjectKeyId,
		"",
		true)

	// revoke certificate
	utils.RemoveDaIntermediateCertificate(
		setup,
		setup.Vendor1,
		intermediateCertificate.Subject,
		intermediateCertificate.SubjectKeyId,
		intermediateCertificate.SerialNumber)

	// check state indexes - leaf certificate stays revoked
	indexes := utils.TestIndexes{
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
	utils.CheckCertificateStateIndexes(t, setup, leafCertificate, indexes)
}

func TestHandler_RemoveDaIntermediateCert_BySubjectAndSKID_RevokedAndActiveCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.RootCertWithVidVid)

	// Add root certificate
	rootCert := utils.RootDaCertificateWithSameSubjectAndSKID1(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// Add two intermediate certificate
	testIntermediateCertificate := utils.IntermediateDaCertificateWithSameSubjectAndSKID1(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, testIntermediateCertificate)

	testIntermediateCertificate2 := utils.IntermediateDaCertificateWithSameSubjectAndSKID2(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, testIntermediateCertificate2)

	// revoke an intermediate certificate
	utils.RevokeDaIntermediateCertificate(
		setup,
		vendorAccAddress,
		testIntermediateCertificate.Subject,
		testIntermediateCertificate.SubjectKeyId,
		testIntermediateCertificate.SerialNumber,
		false)

	// revoke certificate
	utils.RemoveDaIntermediateCertificate(
		setup,
		vendorAccAddress,
		testIntermediateCertificate.Subject,
		testIntermediateCertificate.SubjectKeyId,
		"")

	// check state indexes - both certificates removed
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{},
		Missing: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate, indexes)
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate2, indexes)
}

func TestHandler_RemoveDaIntermediateCert_ByNotOwnerButSameVendor(t *testing.T) {
	setup := utils.Setup(t)

	// store root certificate
	rootCert := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// add certificate by fist vendor account
	testIntermediateCertificate := utils.IntermediateDaCertificate(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, testIntermediateCertificate)

	// add second vendor account with VID = 1
	vendorAccAddress2 := setup.CreateVendorAccount(testconstants.Vid)

	// remove certificate by second vendor account
	utils.RemoveDaIntermediateCertificate(
		setup,
		vendorAccAddress2,
		testIntermediateCertificate.Subject,
		testIntermediateCertificate.SubjectKeyId,
		testIntermediateCertificate.SerialNumber)

	// check state indexes - certificate is removed
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{},
		Missing: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate, indexes)
}

// Error cases

func TestHandler_RemoveDaIntermediateCert_CertificateDoesNotExist(t *testing.T) {
	setup := utils.Setup(t)

	removeX509Cert := types.NewMsgRemoveX509Cert(
		setup.Vendor1.String(),
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		"")
	_, err := setup.Handler(setup.Ctx, removeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_RemoveDaIntermediateCert_InvalidSerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	// propose and approve x509 root certificate
	rootCert := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// Add intermediate certificates
	testIntermediateCertificate := utils.IntermediateDaCertificate(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, testIntermediateCertificate)

	// remove intermediate certificate
	removeX509Cert := types.NewMsgRemoveX509Cert(
		setup.Vendor1.String(),
		testIntermediateCertificate.Subject,
		testIntermediateCertificate.SubjectKeyId,
		"invalid")
	_, err := setup.Handler(setup.Ctx, removeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_RemoveDaIntermediateCert_ByOtherVendor(t *testing.T) {
	setup := utils.Setup(t)

	// add root certificate
	rootCert := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// add intermediate certificates
	testIntermediateCertificate := utils.IntermediateDaCertificate(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, testIntermediateCertificate)

	// add second vendor account with VID = 1000
	vendorAccAddress2 := setup.CreateVendorAccount(testconstants.VendorID1)

	// revoke x509 certificate by second vendor account
	removeX509Cert := types.NewMsgRemoveX509Cert(
		vendorAccAddress2.String(),
		testIntermediateCertificate.Subject,
		testIntermediateCertificate.SubjectKeyId,
		testIntermediateCertificate.SerialNumber,
	)
	_, err := setup.Handler(setup.Ctx, removeX509Cert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_RemoveDaIntermediateCert_SenderNotVendor(t *testing.T) {
	setup := utils.Setup(t)

	// add root certificate
	rootCert := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// add intermediate certificates
	testIntermediateCertificate := utils.IntermediateDaCertificate(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, testIntermediateCertificate)

	removeX509Cert := types.NewMsgRemoveX509Cert(
		setup.Trustee1.String(),
		testIntermediateCertificate.Subject,
		testIntermediateCertificate.SubjectKeyId,
		"invalid")
	_, err := setup.Handler(setup.Ctx, removeX509Cert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_RemoveDaIntermediateCert_ForRootCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// add intermediate certificates
	rootCert := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	removeX509Cert := types.NewMsgRemoveX509Cert(
		setup.Vendor1.String(),
		rootCert.Subject,
		rootCert.SubjectKeyId,
		rootCert.SerialNumber)
	_, err := setup.Handler(setup.Ctx, removeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrInappropriateCertificateType.Is(err))
}

func TestHandler_RemoveDaIntermediateCert_ForNocIcaCertificate(t *testing.T) {
	cases := []CertificateTestCase{
		{
			name:    "OperationalPKI_RemoveDaIntermediateCert_ForNocIcaCertificate",
			crtType: types.CertificateType_OperationalPKI,
		},
		{
			name:    "VIDSignerPKI_RemoveDaIntermediateCert_ForNocIcaCertificate",
			crtType: types.CertificateType_VIDSignerPKI,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			// add NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			// Add ICA certificate
			icaCertificate := utils.IntermediateNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate)

			// Try to remove NOC ICA certificate
			removeX509Cert := types.NewMsgRemoveX509Cert(
				setup.Vendor1.String(),
				icaCertificate.Subject,
				icaCertificate.SubjectKeyId,
				icaCertificate.SerialNumber)
			_, err := setup.Handler(setup.Ctx, removeX509Cert)
			require.Error(t, err)
			require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
		})
	}
}
