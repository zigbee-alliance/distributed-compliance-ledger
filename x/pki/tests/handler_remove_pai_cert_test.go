package tests

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/tests/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Main

func TestHandler_RemoveDaIntermediateCert_BySubjectAndSKID(t *testing.T) {
	setup := utils.Setup(t)

	// propose and approve x509 root certificate
	rootCertOptions := utils.CreateTestRootCertOptions()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// Add intermediate certificates
	testIntermediateCertificate := utils.CreateTestIntermediateCert()
	utils.AddDaIntermediateCertificate(setup, setup.Vendor1, testconstants.IntermediateCertPem)

	// Remove intermediate certificate
	removeX509Cert := types.NewMsgRemoveX509Cert(
		setup.Vendor1.String(),
		testIntermediateCertificate.Subject,
		testIntermediateCertificate.SubjectKeyID,
		"",
	)
	_, err := setup.Handler(setup.Ctx, removeX509Cert)
	require.NoError(t, err)

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
	rootCertOptions := &utils.RootCertOptions{
		PemCert:      testconstants.RootCertWithSameSubjectAndSKID1,
		Subject:      testconstants.RootCertWithSameSubjectAndSKIDSubject,
		SubjectKeyID: testconstants.RootCertWithSameSubjectAndSKIDSubjectKeyID,
		Info:         testconstants.Info,
		Vid:          testconstants.RootCertWithVidVid,
	}
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// Add two intermediate certificates
	testIntermediateCertificate1 := utils.CreateTestIntermediateCertWithSameSubjectAndSKID1()
	utils.AddDaIntermediateCertificate(setup, vendorAccAddress, testconstants.IntermediateWithSameSubjectAndSKID1)

	testIntermediateCertificate2 := utils.CreateTestIntermediateCertWithSameSubjectAndSKID2()
	utils.AddDaIntermediateCertificate(setup, vendorAccAddress, testconstants.IntermediateWithSameSubjectAndSKID2)

	// Add a leaf certificate
	testLeafCertificate := utils.CreateTestLeafCertWithSameSubjectAndSKID()
	utils.AddDaIntermediateCertificate(setup, vendorAccAddress, testconstants.LeafCertWithSameSubjectAndSKID)

	// get certificates for further comparison
	allCerts := setup.Keeper.GetAllApprovedCertificates(setup.Ctx)
	require.NotNil(t, allCerts)
	require.Equal(t, 3, len(allCerts))
	require.Equal(t, 4, len(allCerts[0].Certs)+len(allCerts[1].Certs)+len(allCerts[2].Certs))

	// remove all intermediate certificates but leave leaf certificate
	removeX509Cert := types.NewMsgRemoveX509Cert(
		vendorAccAddress.String(),
		testIntermediateCertificate1.Subject,
		testIntermediateCertificate1.SubjectKeyID,
		"",
	)
	_, err := setup.Handler(setup.Ctx, removeX509Cert)
	require.NoError(t, err)

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
	rootCertOptions := &utils.RootCertOptions{
		PemCert:      testconstants.RootCertWithSameSubjectAndSKID1,
		Subject:      testconstants.RootCertWithSameSubjectAndSKIDSubject,
		SubjectKeyID: testconstants.RootCertWithSameSubjectAndSKIDSubjectKeyID,
		Info:         testconstants.Info,
		Vid:          testconstants.RootCertWithVidVid,
	}
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// Add intermediate certificates
	testIntermediateCertificate1 := utils.CreateTestIntermediateCertWithSameSubjectAndSKID1()
	utils.AddDaIntermediateCertificate(setup, vendorAccAddress, testconstants.IntermediateWithSameSubjectAndSKID1)

	testIntermediateCertificate2 := utils.CreateTestIntermediateCertWithSameSubjectAndSKID2()
	utils.AddDaIntermediateCertificate(setup, vendorAccAddress, testconstants.IntermediateWithSameSubjectAndSKID2)

	// Add a leaf certificate
	testLeafCertificate := utils.CreateTestLeafCertWithSameSubjectAndSKID()
	utils.AddDaIntermediateCertificate(setup, vendorAccAddress, testconstants.LeafCertWithSameSubjectAndSKID)

	// remove  intermediate certificate by serial number
	removeX509Cert := types.NewMsgRemoveX509Cert(
		vendorAccAddress.String(),
		testIntermediateCertificate1.Subject,
		testIntermediateCertificate1.SubjectKeyID,
		testIntermediateCertificate1.SerialNumber,
	)
	_, err := setup.Handler(setup.Ctx, removeX509Cert)
	require.NoError(t, err)

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
	removeX509Cert = types.NewMsgRemoveX509Cert(
		vendorAccAddress.String(),
		testIntermediateCertificate2.Subject,
		testIntermediateCertificate2.SubjectKeyID,
		testIntermediateCertificate2.SerialNumber,
	)
	_, err = setup.Handler(setup.Ctx, removeX509Cert)
	require.NoError(t, err)

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

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.RootCertWithVidVid)

	// propose and approve x509 root certificate
	rootCertOptions := &utils.RootCertOptions{
		PemCert:      testconstants.RootCertPem,
		Subject:      testconstants.RootSubject,
		SubjectKeyID: testconstants.RootSubjectKeyID,
		Info:         testconstants.Info,
		Vid:          testconstants.RootCertWithVidVid,
	}
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// Add two intermediate certificates again
	testIntermediateCertificate := utils.CreateTestIntermediateCert()
	utils.AddDaIntermediateCertificate(setup, vendorAccAddress, testconstants.IntermediateCertPem)

	// revoke intermediate certificate by serial number
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		vendorAccAddress.String(),
		testIntermediateCertificate.Subject,
		testIntermediateCertificate.SubjectKeyID,
		testIntermediateCertificate.SerialNumber,
		false,
		testconstants.Info,
	)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(t, err)

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
	removeX509Cert := types.NewMsgRemoveX509Cert(
		vendorAccAddress.String(),
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		testconstants.IntermediateSerialNumber,
	)
	_, err = setup.Handler(setup.Ctx, removeX509Cert)
	require.NoError(t, err)

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
	// propose and approve x509 root certificate
	rootCertOptions := &utils.RootCertOptions{
		PemCert:      testconstants.RootCertWithSameSubjectAndSKID1,
		Subject:      testconstants.RootCertWithSameSubjectAndSKIDSubject,
		SubjectKeyID: testconstants.RootCertWithSameSubjectAndSKIDSubjectKeyID,
		Info:         testconstants.Info,
		Vid:          testconstants.RootCertWithVidVid,
	}
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.RootCertWithVidVid)

	// Add an intermediate certificate
	addIntermediateX509Cert := types.NewMsgAddX509Cert(vendorAccAddress.String(), testconstants.IntermediateWithSameSubjectAndSKID1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addIntermediateX509Cert)
	require.NoError(t, err)

	// get certificates for further comparison
	allCerts := setup.Keeper.GetAllApprovedCertificates(setup.Ctx)
	require.NotNil(t, allCerts)
	require.Equal(t, 2, len(allCerts))
	require.Equal(t, 2, len(allCerts[0].Certs)+len(allCerts[1].Certs))

	// revoke an intermediate certificate
	revokeX509Cert := types.NewMsgRemoveX509Cert(
		vendorAccAddress.String(),
		testconstants.IntermediateCertWithSameSubjectAndSKIDSubject,
		testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID,
		testconstants.IntermediateCertWithSameSubjectAndSKID1SerialNumber,
	)
	_, err = setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(t, err)

	// Add an intermediate certificate with new serial number
	addIntermediateX509Cert = types.NewMsgAddX509Cert(vendorAccAddress.String(), testconstants.IntermediateWithSameSubjectAndSKID2, testconstants.CertSchemaVersion)
	_, err = setup.Handler(setup.Ctx, addIntermediateX509Cert)
	require.NoError(t, err)

	intermediateCerts, _ := utils.QueryApprovedCertificates(setup, testconstants.IntermediateCertWithSameSubjectAndSKIDSubject, testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(t, 1, len(intermediateCerts.Certs))
	require.Equal(t, testconstants.IntermediateCertWithSameSubjectAndSKIDSubject, intermediateCerts.Certs[0].Subject)
	require.Equal(t, testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID, intermediateCerts.Certs[0].SubjectKeyId)
	require.Equal(t, testconstants.IntermediateCertWithSameSubjectAndSKID2SerialNumber, intermediateCerts.Certs[0].SerialNumber)

	// remove an intermediate certificate
	removeX509Cert := types.NewMsgRemoveX509Cert(
		vendorAccAddress.String(),
		testconstants.IntermediateCertWithSameSubjectAndSKIDSubject,
		testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID,
		testconstants.IntermediateCertWithSameSubjectAndSKID2SerialNumber,
	)
	_, err = setup.Handler(setup.Ctx, removeX509Cert)
	require.NoError(t, err)

	// check that only root and leaf certificates exists
	allCerts, _ = utils.QueryAllApprovedCertificates(setup)
	require.Equal(t, 1, len(allCerts))
	require.Equal(t, true, allCerts[0].Certs[0].IsRoot)
	_, err = utils.QueryApprovedCertificates(setup, testconstants.IntermediateCertWithSameSubjectAndSKIDSubject, testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))
	_, err = utils.QueryRevokedCertificates(setup, testconstants.IntermediateCertWithSameSubjectAndSKIDSubject, testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificates does not exists
	found := setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.RootCertWithSameSubjectAndSKIDSubject, testconstants.IntermediateCertWithSameSubjectAndSKID1SerialNumber)
	require.Equal(t, false, found)
	found = setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.RootCertWithSameSubjectAndSKIDSubject, testconstants.IntermediateCertWithSameSubjectAndSKID2SerialNumber)
	require.Equal(t, false, found)
}

func TestHandler_RemoveX509Cert_ByNotOwnerButSameVendor(t *testing.T) {
	setup := utils.Setup(t)

	// store root certificate
	rootCertificate := utils.RootCertificate(setup.Trustee1)
	setup.Keeper.AddAllCertificate(setup.Ctx, rootCertificate)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// add first vendor account with VID = 1
	vendorAccAddress1 := setup.CreateVendorAccount(testconstants.Vid)

	// add x509 certificate by fist vendor account
	addX509Cert := types.NewMsgAddX509Cert(vendorAccAddress1.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// add second vendor account with VID = 1
	vendorAccAddress2 := setup.CreateVendorAccount(testconstants.Vid)

	// remove x509 certificate by second vendor account
	removeX509Cert := types.NewMsgRemoveX509Cert(
		vendorAccAddress2.String(),
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		testconstants.IntermediateSerialNumber,
	)
	_, err = setup.Handler(setup.Ctx, removeX509Cert)
	require.NoError(t, err)

	// check that certificate removed from 'approved certificates' list
	_, err = utils.QueryApprovedCertificates(setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that certificate removed from 'approved certificates by subject' list
	_, err = utils.QueryApprovedCertificatesBySubject(setup, testconstants.IntermediateSubject)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that certificate removed from 'approved certificates by SKID' list
	approvedCerts, err := utils.QueryApprovedCertificatesBySubjectKeyID(setup, testconstants.IntermediateSubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 0, len(approvedCerts))

	// check that unique certificate key is not registered
	require.False(t, setup.Keeper.IsUniqueCertificatePresent(setup.Ctx,
		testconstants.IntermediateIssuer, testconstants.IntermediateSerialNumber))
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
	rootCertOptions := utils.CreateRootWithVidOptions()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

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

	rootCertOptions := utils.CreateTestRootCertOptions()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

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

	rootCertOptions := utils.CreateTestRootCertOptions()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

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
