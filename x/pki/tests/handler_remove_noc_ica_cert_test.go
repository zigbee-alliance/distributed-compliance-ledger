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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Main

func TestHandler_RemoveNocIntermediateCert(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.Vid)

	// add NOC root certificate
	utils.AddNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1)

	// add intermediate certificate
	icaCertificate := utils.CreateTestNocIca1Cert()
	utils.AddNocIntermediateCertificate(setup, vendorAccAddress, testconstants.NocCert1)

	// remove intermediate certificate
	removeIcaCert := types.NewMsgRemoveNocX509IcaCert(
		vendorAccAddress.String(),
		icaCertificate.Subject,
		icaCertificate.SubjectKeyID,
		"",
	)
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	// Check indexes
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.NocRootCertificatesKeyPrefix, Count: 1}, // root still exits

		},
		Missing: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
			{Key: types.NocIcaCertificatesKeyPrefix},
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.RevokedNocIcaCertificatesKeyPrefix},
			{Key: types.RevokedNocRootCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate, indexes)
}

func TestHandler_RemoveNocX509IcaCert_BySubjectAndSKID(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	rootCertificate := utils.CreateTestNocRoot1Cert()
	utils.AddNocRootCertificate(setup, vendorAccAddress, rootCertificate.PEM)

	// add two intermediate certificates
	icaCertificate1 := utils.CreateTestNocIca1Cert()
	utils.AddNocIntermediateCertificate(setup, vendorAccAddress, icaCertificate1.PEM)

	icaCertificate2 := utils.CreateTestNocIca1CertCopy()
	utils.AddNocIntermediateCertificate(setup, vendorAccAddress, icaCertificate2.PEM)

	// add leaf certificate
	leafCertificate := utils.CreateTestNocLeafCert()
	utils.AddNocIntermediateCertificate(setup, vendorAccAddress, leafCertificate.PEM)

	// get certificates for further comparison
	nocCerts := setup.Keeper.GetAllNocCertificates(setup.Ctx)
	require.NotNil(t, nocCerts)
	require.Equal(t, 3, len(nocCerts))
	require.Equal(t, 4, len(nocCerts[0].Certs)+len(nocCerts[1].Certs)+len(nocCerts[2].Certs))

	// Check indexes for intermediate certificates before removing
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix, Count: 2},
			{Key: types.AllCertificatesBySubjectKeyPrefix, Count: 1},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.NocCertificatesKeyPrefix, Count: 2},
			{Key: types.NocCertificatesBySubjectKeyPrefix, Count: 1},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.NocRootCertificatesKeyPrefix, Count: 1}, // root still exits
			{Key: types.NocIcaCertificatesKeyPrefix, Count: 3},  // 2 inter + leaf certs exist

		},
		Missing: []utils.TestIndex{},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate1, indexes)
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate2, indexes)

	// remove all intermediate certificates but leave leaf certificate (NocCert1 and NocCert1Copy)
	removeIcaCert := types.NewMsgRemoveNocX509IcaCert(
		vendorAccAddress.String(),
		icaCertificate1.Subject,
		icaCertificate1.SubjectKeyID,
		"",
	)
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	// Check indexes for intermediate certificates
	indexes = utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.NocRootCertificatesKeyPrefix, Count: 1}, // root still exits
			{Key: types.NocIcaCertificatesKeyPrefix, Count: 1},  // leaf cert with same vid exist

		},
		Missing: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.RevokedNocIcaCertificatesKeyPrefix},
			{Key: types.RevokedNocRootCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate1, indexes)
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate2, indexes)

	// Check indexes
	indexes = utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
			{Key: types.NocRootCertificatesKeyPrefix, Count: 1}, // root still exits
			{Key: types.NocIcaCertificatesKeyPrefix, Count: 1},  // only leaf exits
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, leafCertificate, indexes)

	// Check that only 2 certificates exists
	nocCerts, _ = utils.QueryAllNocCertificates(setup)
	require.Equal(t, 2, len(nocCerts))
	require.Equal(t, 2, len(nocCerts[0].Certs)+len(nocCerts[1].Certs)) // root + leaf
}

func TestHandler_RemoveNocX509IcaCert_BySerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.Vid)

	// add NOC root certificate
	rootCertificate := utils.CreateTestNocRoot1Cert()
	utils.AddNocRootCertificate(setup, vendorAccAddress, rootCertificate.PEM)

	// Add ICA certificates
	icaCertificate1 := utils.CreateTestNocIca1Cert()
	utils.AddNocIntermediateCertificate(setup, vendorAccAddress, icaCertificate1.PEM)

	// Add ICA certificates with sam subject and SKID but different serial number
	icaCertificate2 := utils.CreateTestNocIca1CertCopy()
	utils.AddNocIntermediateCertificate(setup, vendorAccAddress, icaCertificate2.PEM)

	// Add a leaf certificate
	leafCertificate := utils.CreateTestNocLeafCert()
	utils.AddNocIntermediateCertificate(setup, vendorAccAddress, leafCertificate.PEM)

	// Check indexes for intermediate certificates before removing
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix, Count: 2},
			{Key: types.AllCertificatesBySubjectKeyPrefix, Count: 1},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.NocCertificatesKeyPrefix, Count: 2},
			{Key: types.NocCertificatesBySubjectKeyPrefix, Count: 1},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.NocRootCertificatesKeyPrefix, Count: 1}, // root still exits
			{Key: types.NocIcaCertificatesKeyPrefix, Count: 3},  // 2 inter + leaf certs exist

		},
		Missing: []utils.TestIndex{},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate1, indexes)
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate2, indexes)

	// remove ICA certificate by serial number
	removeIcaCert := types.NewMsgRemoveNocX509IcaCert(
		vendorAccAddress.String(),
		icaCertificate1.Subject,
		icaCertificate1.SubjectKeyID,
		icaCertificate1.SerialNumber,
	)
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	// Check indexes for first certificate (second ica exist)
	indexes = utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix, Count: 1},
			{Key: types.AllCertificatesBySubjectKeyPrefix, Count: 1},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Count: 1},
			{Key: types.NocCertificatesKeyPrefix, Count: 1},
			{Key: types.NocCertificatesBySubjectKeyPrefix, Count: 1},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix, Count: 1},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix, Count: 1},
			{Key: types.NocRootCertificatesKeyPrefix, Count: 1}, // root still exits
			{Key: types.NocIcaCertificatesKeyPrefix, Count: 2},  // ica and leaf cert with same vid exist
			{Key: types.ChildCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix}, // removed
			{Key: types.RevokedNocIcaCertificatesKeyPrefix},
			{Key: types.RevokedNocRootCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate1, indexes)

	// Check indexes for second certificate (all same as for ica1 but also UniqueCertificate exists)
	indexes = utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix, Count: 1},
			{Key: types.AllCertificatesBySubjectKeyPrefix, Count: 1},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Count: 1},
			{Key: types.NocCertificatesKeyPrefix, Count: 1},
			{Key: types.NocCertificatesBySubjectKeyPrefix, Count: 1},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix, Count: 1},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix, Count: 1},
			{Key: types.NocRootCertificatesKeyPrefix, Count: 1}, // root still exits
			{Key: types.NocIcaCertificatesKeyPrefix, Count: 2},  // ica and leaf cert with same vid exist
			{Key: types.UniqueCertificateKeyPrefix},             // all same as for ica1 but also UniqueCertificate exists
			{Key: types.ChildCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.RevokedNocIcaCertificatesKeyPrefix},
			{Key: types.RevokedNocRootCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate2, indexes)

	// Check indexes for leaf certificate (all same as for ica2)
	utils.CheckCertificateStateIndexes(t, setup, leafCertificate, indexes)
}

func TestHandler_RemoveNocX509IcaCert_RevokedCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.Vid)

	// add NOC root certificate
	rootCertificate := utils.CreateTestNocRoot1Cert()
	utils.AddNocRootCertificate(setup, vendorAccAddress, rootCertificate.PEM)

	// Add an intermediate certificate
	icaCertificate := utils.CreateTestNocIca1Cert()
	utils.AddNocIntermediateCertificate(setup, vendorAccAddress, icaCertificate.PEM)

	// revoke intermediate certificate by serial number
	revokeX509Cert := types.NewMsgRevokeNocX509IcaCert(
		vendorAccAddress.String(),
		icaCertificate.Subject,
		icaCertificate.SubjectKeyID,
		icaCertificate.SerialNumber,
		testconstants.Info,
		false,
	)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(t, err)

	// Check indexes after revocation
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.RevokedNocIcaCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
			{Key: types.NocIcaCertificatesKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.RevokedNocRootCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate, indexes)

	// remove intermediate certificate by serial number
	removeIcaCert := types.NewMsgRemoveNocX509IcaCert(
		vendorAccAddress.String(),
		icaCertificate.Subject,
		icaCertificate.SubjectKeyID,
		icaCertificate.SerialNumber,
	)
	_, err = setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	indexes = utils.TestIndexes{
		Present: []utils.TestIndex{},
		Missing: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
			{Key: types.NocIcaCertificatesKeyPrefix},
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.RevokedNocIcaCertificatesKeyPrefix},
			{Key: types.RevokedNocRootCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate, indexes)
}

// Extra cases

func TestHandler_RemoveNocX509IcaCert_RevokedAndActiveCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.Vid)

	// add NOC root certificate
	rootCertificate := utils.CreateTestNocRoot1Cert()
	utils.AddNocRootCertificate(setup, vendorAccAddress, rootCertificate.PEM)

	// Add an intermediate certificate
	icaCertificate := utils.CreateTestNocIca1Cert()
	utils.AddNocIntermediateCertificate(setup, vendorAccAddress, icaCertificate.PEM)

	// revoke an intermediate certificate
	revokeX509Cert := types.NewMsgRevokeNocX509IcaCert(
		vendorAccAddress.String(),
		icaCertificate.Subject,
		icaCertificate.SubjectKeyID,
		icaCertificate.SerialNumber,
		testconstants.Info,
		false,
	)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(t, err)

	// Check indexes after revocation
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.RevokedNocIcaCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
			{Key: types.NocIcaCertificatesKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.RevokedNocRootCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate, indexes)

	// Add an intermediate certificate with new serial number
	icaCertificate2 := utils.CreateTestNocIca1CertCopy()
	utils.AddNocIntermediateCertificate(setup, vendorAccAddress, icaCertificate2.PEM)

	// Check indexes
	indexes = utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
			{Key: types.NocRootCertificatesKeyPrefix, Count: 1}, // we create root certificate as well but ica should not get there
			{Key: types.NocIcaCertificatesKeyPrefix},
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.RevokedNocIcaCertificatesKeyPrefix}, // we have evoked cert with same id
		},
		Missing: []utils.TestIndex{},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate2, indexes)

	// remove an intermediate certificate
	removeIcaCert := types.NewMsgRemoveNocX509IcaCert(
		vendorAccAddress.String(),
		icaCertificate.Subject,
		icaCertificate.SubjectKeyID,
		"",
	)
	_, err = setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	// check that only root certificates exists
	allCerts, _ := utils.QueryAllNocCertificates(setup)
	require.Equal(t, 1, len(allCerts))
	require.Equal(t, true, allCerts[0].Certs[0].IsRoot)

	indexes = utils.TestIndexes{
		Present: []utils.TestIndex{},
		Missing: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
			{Key: types.NocIcaCertificatesKeyPrefix},
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.RevokedNocIcaCertificatesKeyPrefix},
			{Key: types.RevokedNocRootCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate, indexes)
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate2, indexes)
}

func TestHandler_RemoveNocX509IcaCert_ByNotOwnerButSameVendor(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.Vid)

	// add NOC root certificate
	rootCertificate := utils.CreateTestNocRoot1Cert()
	utils.AddNocRootCertificate(setup, vendorAccAddress, rootCertificate.PEM)

	// add first vendor account with VID = 1
	vendorAccAddress1 := setup.CreateVendorAccount(testconstants.Vid)

	// add ICA certificate by fist vendor account
	addIcaCert := types.NewMsgAddNocX509IcaCert(vendorAccAddress1.String(), testconstants.NocCert1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addIcaCert)
	require.NoError(t, err)

	// add second vendor account with VID = 1
	vendorAccAddress2 := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// remove x509 certificate by second vendor account
	removeIcaCert := types.NewMsgRemoveNocX509IcaCert(
		vendorAccAddress2.String(),
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1SerialNumber,
	)
	_, err = setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	// check that certificate removed from 'noc certificates' list
	_, err = utils.QueryNocCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that certificate removed from 'noc certificates by subject' list
	_, err = utils.QueryNocCertificatesBySubject(setup, testconstants.NocCert1Subject)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that certificate removed from 'noc certificates by SKID' list
	nocCerts, err := utils.QueryNocCertificatesBySubjectKeyID(setup, testconstants.NocCert1SubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 0, len(nocCerts))

	// query noc certificate by VID
	_, err = utils.QueryNocIcaCertificatesByVid(setup, testconstants.Vid)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificate key is not registered
	require.False(t, setup.Keeper.IsUniqueCertificatePresent(setup.Ctx,
		testconstants.NocCert1Issuer, testconstants.NocCert1SerialNumber))

	// check that intermediate certificate can not be queried by vid+skid
	_, err = utils.QueryNocCertificatesByVidAndSkid(setup, testconstants.Vid, testconstants.NocCert1SubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
}

// Error cases

func TestHandler_RemoveNocX509IcaCert_CertificateDoesNotExist(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	removeIcaCert := types.NewMsgRemoveNocX509IcaCert(
		vendorAccAddress.String(), testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID, testconstants.NocCert1SerialNumber)
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_RemoveNocX509IcaCert_EmptyCertificatesList(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	utils.AddNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1)

	setup.Keeper.SetNocIcaCertificates(
		setup.Ctx,
		types.NocIcaCertificates{
			Vid: vid,
		},
	)

	removeIcaCert := types.NewMsgRemoveNocX509IcaCert(
		vendorAccAddress.String(), testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID, "")
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_RemoveNocX509IcaCert_ByOtherVendor(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	utils.AddNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1)

	// add fist vendor account with VID = 1
	vendorAccAddress1 := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add x509 certificate by `setup.Trustee`
	addX509Cert := types.NewMsgAddNocX509IcaCert(vendorAccAddress1.String(), testconstants.NocCert1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// add second vendor account with VID = 1000
	vendorAccAddress2 := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID1)

	// remove ICA certificate by second vendor account
	removeIcaCert := types.NewMsgRemoveNocX509IcaCert(
		vendorAccAddress2.String(), testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID, testconstants.NocCert1SerialNumber)
	_, err = setup.Handler(setup.Ctx, removeIcaCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertVidNotEqualAccountVid.Is(err))
}

func TestHandler_RemoveNocX509IcaCert_SenderNotVendor(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	utils.AddNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1)

	// add x509 certificate
	addX509Cert := types.NewMsgAddNocX509IcaCert(vendorAccAddress.String(), testconstants.NocCert1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	removeIcaCert := types.NewMsgRemoveNocX509IcaCert(
		setup.Trustee1.String(), testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID, "")
	_, err = setup.Handler(setup.Ctx, removeIcaCert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_RemoveNocX509IcaCert_ForNonIcaCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vendorAccAddress := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	setup.Keeper.SetRevokedCertificates(
		setup.Ctx,
		types.RevokedCertificates{
			Subject:      testconstants.IntermediateSubject,
			SubjectKeyId: testconstants.IntermediateSubjectKeyID,
			Certs: []*types.Certificate{{
				CertificateType: types.CertificateType_DeviceAttestationPKI,
			}},
		},
	)

	removeIcaCert := types.NewMsgRemoveNocX509IcaCert(
		vendorAccAddress.String(), testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID, "")
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_RemoveNocX509IcaCert_InvalidSerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	utils.AddNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1)

	addX509Cert := types.NewMsgAddNocX509IcaCert(vendorAccAddress.String(), testconstants.NocCert1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	removeX509Cert := types.NewMsgRemoveNocX509IcaCert(
		vendorAccAddress.String(), testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID, "invalid")
	_, err = setup.Handler(setup.Ctx, removeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}
