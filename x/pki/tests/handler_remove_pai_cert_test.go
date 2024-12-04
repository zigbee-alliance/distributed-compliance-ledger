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

	// propose and approve x509 root certificate
	rootCert := utils.RootCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// Add intermediate certificates
	testIntermediateCertificate := utils.IntermediateCertPem(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, setup.Vendor1, testIntermediateCertificate.PemCert)

	// Remove intermediate certificate
	utils.RemoveDaIntermediateCertificate(
		setup,
		setup.Vendor1,
		testIntermediateCertificate.Subject,
		testIntermediateCertificate.SubjectKeyId,
		"")

	// Check: only one certificate exists - root
	allCerts, _ := utils.QueryAllApprovedCertificates(setup)
	require.Equal(t, 1, len(allCerts))

	// Check indexes for intermediate certificate
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
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate, indexes)
}

func TestHandler_RemoveX509Cert_BySubjectAndSKID_TwoCerts(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.RootCertWithVidVid)

	// propose and approve x509 root certificate
	rootCert := utils.RootCertWithSameSubjectAndSKID1(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// Add two intermediate certificates
	testIntermediateCertificate1 := utils.IntermediateWithSameSubjectAndSKID1(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, vendorAccAddress, testIntermediateCertificate1.PemCert)

	testIntermediateCertificate2 := utils.IntermediateWithSameSubjectAndSKID2(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, vendorAccAddress, testIntermediateCertificate2.PemCert)

	// Add a leaf certificate
	testLeafCertificate := utils.LeafCertWithSameSubjectAndSKID(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, vendorAccAddress, testLeafCertificate.PemCert)

	// get certificates for further comparison
	allCerts := setup.Keeper.GetAllApprovedCertificates(setup.Ctx)
	require.NotNil(t, allCerts)
	require.Equal(t, 3, len(allCerts))
	require.Equal(t, 4, len(allCerts[0].Certs)+len(allCerts[1].Certs)+len(allCerts[2].Certs))

	// remove all intermediate certificates but leave leaf certificate
	utils.RemoveDaIntermediateCertificate(
		setup,
		vendorAccAddress,
		testIntermediateCertificate1.Subject,
		testIntermediateCertificate1.SubjectKeyId,
		"")

	// check that only two certificates exists
	allCerts, _ = utils.QueryAllApprovedCertificates(setup)
	require.Equal(t, 2, len(allCerts))
	require.Equal(t, 2, len(allCerts[0].Certs)+len(allCerts[1].Certs))

	// Check indexes for intermediate certificate
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			// {Key: types.AllCertificatesBySubjectKeyPrefix}, // leaf cert has same subject
			// {Key: types.ApprovedCertificatesBySubjectKeyPrefix}, // leaf cert has same subject
		},
		Missing: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate1, indexes)
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate2, indexes)

	// check that leaf certificate exists
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
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, testLeafCertificate, indexes)
}

func TestHandler_RemoveX509Cert_BySerialNumber_TwoCerts(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.RootCertWithVidVid)

	// propose and approve x509 root certificate
	rootCert := utils.RootCertWithSameSubjectAndSKID1(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// Add intermediate certificates
	testIntermediateCertificate1 := utils.IntermediateWithSameSubjectAndSKID1(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, vendorAccAddress, testIntermediateCertificate1.PemCert)

	testIntermediateCertificate2 := utils.IntermediateWithSameSubjectAndSKID2(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, vendorAccAddress, testIntermediateCertificate2.PemCert)

	// Add a leaf certificate
	testLeafCertificate := utils.LeafCertWithSameSubjectAndSKID(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, vendorAccAddress, testLeafCertificate.PemCert)

	// remove  intermediate certificate by serial number
	utils.RemoveDaIntermediateCertificate(
		setup,
		vendorAccAddress,
		testIntermediateCertificate1.Subject,
		testIntermediateCertificate1.SubjectKeyId,
		testIntermediateCertificate1.SerialNumber)

	// check that only root, intermediate(with serial number 3) and leaf certificates exists
	allCerts, _ := utils.QueryAllApprovedCertificates(setup)
	require.Equal(t, 3, len(allCerts))
	require.Equal(t, 3, len(allCerts[0].Certs)+len(allCerts[1].Certs)+len(allCerts[2].Certs))

	// Check indexes for intermediate certificate 1
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix, Count: 2}, // inter + leaf
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix, Count: 2}, // inter + leaf
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate1, indexes)

	// Check indexes for intermediate certificate 2 (all the same but also UniqueCertificate exists)
	indexes = utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix, Count: 2}, // inter + leaf
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix, Count: 2}, // inter + leaf
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{},
	}
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate2, indexes)

	// check that leaf certificate exists (same as for intermediate 2, skip check by subject)
	indexes = utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{},
	}
	utils.CheckCertificateStateIndexes(t, setup, testLeafCertificate, indexes)

	// remove  intermediate certificate by serial number and check that leaf cert is not removed
	utils.RemoveDaIntermediateCertificate(
		setup,
		vendorAccAddress,
		testIntermediateCertificate2.Subject,
		testIntermediateCertificate2.SubjectKeyId,
		testIntermediateCertificate2.SerialNumber)

	allCerts, _ = utils.QueryAllApprovedCertificates(setup)
	require.Equal(t, 2, len(allCerts))
	require.Equal(t, 2, len(allCerts[0].Certs)+len(allCerts[1].Certs))

	// Check indexes for intermediate certificates
	indexes = utils.TestIndexes{
		Present: []utils.TestIndex{},
		Missing: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate1, indexes)
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate2, indexes)

	// check that leaf certificate exists
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
			{Key: types.ProposedCertificateKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, testLeafCertificate, indexes)
}

func TestHandler_RemoveX509Cert_RevokedCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// propose and approve x509 root certificate
	rootCert := utils.RootCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// Add two intermediate certificates again
	testIntermediateCertificate := utils.IntermediateCertPem(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, setup.Vendor1, testIntermediateCertificate.PemCert)

	// revoke intermediate certificate by serial number
	utils.RevokeDaIntermediateCertificate(
		setup,
		setup.Vendor1,
		testIntermediateCertificate.Subject,
		testIntermediateCertificate.SubjectKeyId,
		testIntermediateCertificate.SerialNumber,
		false)

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
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate, indexes)

	// remove  intermediate certificate by serial number
	utils.RemoveDaIntermediateCertificate(
		setup,
		setup.Vendor1,
		testIntermediateCertificate.Subject,
		testIntermediateCertificate.SubjectKeyId,
		testIntermediateCertificate.SerialNumber)

	indexes = utils.TestIndexes{
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

// Extra cases

func TestHandler_RemoveX509Cert_RevokedAndApprovedCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.RootCertWithVidVid)

	// propose and approve x509 root certificate
	rootCert := utils.RootCertWithSameSubjectAndSKID1(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// Add an intermediate certificate
	testIntermediateCertificate := utils.IntermediateWithSameSubjectAndSKID1(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, vendorAccAddress, testIntermediateCertificate.PemCert)

	// get certificates for further comparison
	allCerts := setup.Keeper.GetAllApprovedCertificates(setup.Ctx)
	require.Equal(t, 2, len(allCerts))

	// revoke an intermediate certificate
	utils.RevokeDaIntermediateCertificate(
		setup,
		vendorAccAddress,
		testIntermediateCertificate.Subject,
		testIntermediateCertificate.SubjectKeyId,
		testIntermediateCertificate.SerialNumber,
		false)

	// Add an intermediate certificate with new serial number
	testIntermediateCertificate2 := utils.IntermediateWithSameSubjectAndSKID2(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, vendorAccAddress, testIntermediateCertificate2.PemCert)

	// check that intermediate certificate 2 exists
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
			{Key: types.ProposedCertificateKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate2, indexes)

	// check that intermediate certificate 1 exists but revoked
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
			{Key: types.RevokedCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate2, indexes)

	// remove an intermediate certificate
	utils.RemoveDaIntermediateCertificate(
		setup,
		vendorAccAddress,
		testIntermediateCertificate2.Subject,
		testIntermediateCertificate2.SubjectKeyId,
		"")

	// check indexes after removing re-activated certificate
	indexes = utils.TestIndexes{
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
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate2, indexes)
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate, indexes)
}

func TestHandler_RemoveX509Cert_ByNotOwnerButSameVendor(t *testing.T) {
	setup := utils.Setup(t)

	// store root certificate
	rootCert := utils.RootCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// add x509 certificate by fist vendor account
	testIntermediateCertificate := utils.IntermediateCertPem(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, setup.Vendor1, testIntermediateCertificate.PemCert)

	// add second vendor account with VID = 1
	vendorAccAddress2 := setup.CreateVendorAccount(testconstants.Vid)

	// remove x509 certificate by second vendor account
	utils.RemoveDaIntermediateCertificate(
		setup,
		vendorAccAddress2,
		testIntermediateCertificate.Subject,
		testIntermediateCertificate.SubjectKeyId,
		testIntermediateCertificate.SerialNumber)

	// check state indexes
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

func TestHandler_RemoveX509Cert_CertificateDoesNotExist(t *testing.T) {
	setup := utils.Setup(t)

	removeX509Cert := types.NewMsgRemoveX509Cert(
		setup.Vendor1.String(),
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		testconstants.IntermediateSerialNumber)
	_, err := setup.Handler(setup.Ctx, removeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_RemoveX509Cert_EmptyCertificatesList(t *testing.T) {
	setup := utils.Setup(t)

	rootCertificate := utils.RootCertificate(setup.Trustee1)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	setup.Keeper.SetApprovedCertificates(
		setup.Ctx,
		types.ApprovedCertificates{
			Subject:      testconstants.IntermediateSubject,
			SubjectKeyId: testconstants.IntermediateSubjectKeyID,
		},
	)

	removeX509Cert := types.NewMsgRemoveX509Cert(
		setup.Vendor1.String(),
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		"")
	_, err := setup.Handler(setup.Ctx, removeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_RemoveX509Cert_ByOtherVendor(t *testing.T) {
	setup := utils.Setup(t)

	// store root certificate
	rootCertificate := utils.RootCertificate(setup.Trustee1)
	setup.Keeper.AddAllCertificate(setup.Ctx, rootCertificate)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// add fist vendor account with VID = 1
	vendorAccAddress1 := setup.CreateVendorAccount(testconstants.Vid)

	// add x509 certificate by `setup.Trustee`
	addX509Cert := types.NewMsgAddX509Cert(vendorAccAddress1.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// add scond vendor account with VID = 1000
	vendorAccAddress2 := setup.CreateVendorAccount(testconstants.VendorID1)

	// revoke x509 certificate by second vendor account
	removeX509Cert := types.NewMsgRemoveX509Cert(
		vendorAccAddress2.String(), testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID, testconstants.IntermediateSerialNumber)
	_, err = setup.Handler(setup.Ctx, removeX509Cert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_RemoveX509Cert_SenderNotVendor(t *testing.T) {
	setup := utils.Setup(t)

	// store root certificate
	rootCert := utils.RootCertWithVid(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.RootCertWithVidVid)

	// add x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(vendorAccAddress.String(), testconstants.IntermediateCertWithVid1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	removeX509Cert := types.NewMsgRemoveX509Cert(
		setup.Trustee1.String(), testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID, "invalid")
	_, err = setup.Handler(setup.Ctx, removeX509Cert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_RemoveX509Cert_ForRootCertificate(t *testing.T) {
	setup := utils.Setup(t)

	rootCert := utils.RootCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	removeX509Cert := types.NewMsgRemoveX509Cert(
		setup.Vendor1.String(),
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.RootSerialNumber)
	_, err := setup.Handler(setup.Ctx, removeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrInappropriateCertificateType.Is(err))
}

func TestHandler_RemoveX509Cert_InvalidSerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	rootCert := utils.RootCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	addX509Cert := types.NewMsgAddX509Cert(setup.Vendor1.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	removeX509Cert := types.NewMsgRemoveX509Cert(
		setup.Vendor1.String(),
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		"invalid")
	_, err = setup.Handler(setup.Ctx, removeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_RemoveX509Cert_ForNocIcaCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// add NOC root certificate
	utils.AddNocRootCertificate(setup, setup.Vendor1, testconstants.NocRootCert1)

	// Add ICA certificate
	addX509Cert := types.NewMsgAddNocX509IcaCert(setup.Vendor1.String(), testconstants.NocCert1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	removeX509Cert := types.NewMsgRemoveX509Cert(
		setup.Vendor1.String(),
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1SerialNumber)
	_, err = setup.Handler(setup.Ctx, removeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}
