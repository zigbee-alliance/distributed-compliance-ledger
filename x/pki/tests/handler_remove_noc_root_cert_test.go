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

func TestHandler_RemoveNocRootCert(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.Vid)

	// add NOC root certificates
	rootCertificate := utils.CreateTestNocRoot1Cert()
	utils.AddNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1)

	// remove noc root  certificate
	removeIcaCert := types.NewMsgRemoveNocX509RootCert(
		vendorAccAddress.String(),
		rootCertificate.Subject,
		rootCertificate.SubjectKeyID,
		"",
	)
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	// Check indexes
	indexes := []utils.TestIndex{
		{Key: types.AllCertificatesKeyPrefix, Exist: false},
		{Key: types.AllCertificatesBySubjectKeyPrefix, Exist: false},
		{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Exist: false},
		{Key: types.NocCertificatesKeyPrefix, Exist: false},
		{Key: types.NocCertificatesBySubjectKeyPrefix, Exist: false},
		{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix, Exist: false},
		{Key: types.NocCertificatesByVidAndSkidKeyPrefix, Exist: false},
		{Key: types.NocRootCertificatesKeyPrefix, Exist: false},
		{Key: types.NocIcaCertificatesKeyPrefix, Exist: false},
		{Key: types.UniqueCertificateKeyPrefix, Exist: false},
		{Key: types.ChildCertificatesKeyPrefix, Exist: false},
		{Key: types.RevokedNocIcaCertificatesKeyPrefix, Exist: false},
		{Key: types.RevokedNocRootCertificatesKeyPrefix, Exist: false},
		{Key: types.RevokedCertificatesKeyPrefix, Exist: false},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)
}

func TestHandler_RemoveNocX509RootCert_BySubjectAndSKID(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := setup.CreateVendorAccount(vid)

	// add NOC root certificates
	rootCertificate1 := utils.CreateTestNocRoot1Cert()
	utils.AddNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1)

	rootCertificate2 := utils.CreateTestNocRoot2Cert()
	utils.AddNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1Copy)

	// Add intermediate certificate
	icaCertificate := utils.CreateTestNocIca1Cert()
	utils.AddNocIntermediateCertificate(setup, vendorAccAddress, testconstants.NocCert1)

	// get certificates for further comparison
	nocCerts := setup.Keeper.GetAllNocCertificates(setup.Ctx)
	require.NotNil(t, nocCerts)
	require.Equal(t, 2, len(nocCerts))
	require.Equal(t, 3, len(nocCerts[0].Certs)+len(nocCerts[1].Certs))

	// remove all root nOC certificates but IAC certificate
	removeIcaCert := types.NewMsgRemoveNocX509RootCert(
		vendorAccAddress.String(),
		rootCertificate1.Subject,
		rootCertificate1.SubjectKeyID,
		"",
	)
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	// check that only IAC certificate exists
	nocCerts, _ = utils.QueryAllNocCertificates(setup)
	require.Equal(t, 1, len(nocCerts))
	require.Equal(t, 1, len(nocCerts[0].Certs))
	require.Equal(t, testconstants.NocCert1SerialNumber, nocCerts[0].Certs[0].SerialNumber)

	// Check indexes for root certificates
	indexes := []utils.TestIndex{
		{Key: types.AllCertificatesKeyPrefix, Exist: false},
		{Key: types.AllCertificatesBySubjectKeyPrefix, Exist: false},
		{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Exist: false},
		{Key: types.NocCertificatesKeyPrefix, Exist: false},
		{Key: types.NocCertificatesBySubjectKeyPrefix, Exist: false},
		{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix, Exist: false},
		{Key: types.NocCertificatesByVidAndSkidKeyPrefix, Exist: false},
		{Key: types.NocRootCertificatesKeyPrefix, Exist: false},
		{Key: types.UniqueCertificateKeyPrefix, Exist: false},
		{Key: types.RevokedNocIcaCertificatesKeyPrefix, Exist: false},
		{Key: types.RevokedNocRootCertificatesKeyPrefix, Exist: false},
		{Key: types.RevokedCertificatesKeyPrefix, Exist: false},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate1, indexes)
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate2, indexes)

	// Check indexes for intermediate certificates
	indexes = []utils.TestIndex{
		{Key: types.AllCertificatesKeyPrefix, Exist: true},
		{Key: types.AllCertificatesBySubjectKeyPrefix, Exist: true},
		{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Exist: true},
		{Key: types.NocCertificatesKeyPrefix, Exist: true},
		{Key: types.NocCertificatesBySubjectKeyPrefix, Exist: true},
		{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix, Exist: true},
		{Key: types.NocCertificatesByVidAndSkidKeyPrefix, Exist: true},
		{Key: types.NocIcaCertificatesKeyPrefix, Exist: true},
		{Key: types.UniqueCertificateKeyPrefix, Exist: true},
		{Key: types.ChildCertificatesKeyPrefix, Exist: true},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate, indexes)
}

func TestHandler_RemoveNocX509RootCert_BySerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificates
	rootCertificate1 := utils.CreateTestNocRoot1Cert()
	utils.AddNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1)

	rootCertificate2 := utils.CreateTestNocRoot2Cert()
	utils.AddNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1Copy)

	// Add ICA certificates
	icaCertificate := utils.CreateTestNocIca1Cert()
	utils.AddNocIntermediateCertificate(setup, vendorAccAddress, testconstants.NocCert1)

	// remove NOC root certificate by serial number
	removeIcaCert := types.NewMsgRemoveNocX509RootCert(
		vendorAccAddress.String(),
		rootCertificate1.Subject,
		rootCertificate1.SubjectKeyID,
		rootCertificate1.SerialNumber,
	)
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	nocCerts, _ := utils.QueryAllNocCertificates(setup)
	require.Equal(t, 2, len(nocCerts))

	// NocCertificates: Subject and SKID
	nocCertificates, err := utils.QueryNocCertificates(
		setup,
		testconstants.NocRootCert1CopySubject,
		testconstants.NocRootCert1CopySubjectKeyID,
	)
	require.NoError(t, err)
	require.Equal(t, 1, len(nocCertificates.Certs))

	// Check indexes for root certificates
	indexes := []utils.TestIndex{
		{Key: types.UniqueCertificateKeyPrefix, Exist: false},
		{Key: types.AllCertificatesKeyPrefix, Exist: true},
		{Key: types.AllCertificatesBySubjectKeyPrefix, Exist: true},
		{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Exist: true},
		{Key: types.NocCertificatesKeyPrefix, Exist: true},
		{Key: types.NocCertificatesBySubjectKeyPrefix, Exist: true},
		{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix, Exist: true},
		{Key: types.NocCertificatesByVidAndSkidKeyPrefix, Exist: true},
		{Key: types.NocIcaCertificatesKeyPrefix, Exist: true},
		{Key: types.RevokedNocRootCertificatesKeyPrefix, Exist: true},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate1, indexes)

	indexes = []utils.TestIndex{
		{Key: types.AllCertificatesKeyPrefix, Exist: true},
		{Key: types.AllCertificatesBySubjectKeyPrefix, Exist: true},
		{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Exist: true},
		{Key: types.NocCertificatesKeyPrefix, Exist: true},
		{Key: types.NocCertificatesBySubjectKeyPrefix, Exist: true},
		{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix, Exist: true},
		{Key: types.NocCertificatesByVidAndSkidKeyPrefix, Exist: true},
		{Key: types.NocIcaCertificatesKeyPrefix, Exist: true},
		{Key: types.UniqueCertificateKeyPrefix, Exist: true},
		{Key: types.ChildCertificatesKeyPrefix, Exist: true},
		{Key: types.RevokedNocRootCertificatesKeyPrefix, Exist: true},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate2, indexes)

	// remove NOC root certificate by serial number and check that IAC cert is not removed
	removeIcaCert = types.NewMsgRemoveNocX509RootCert(
		vendorAccAddress.String(),
		rootCertificate2.Subject,
		rootCertificate2.SubjectKeyID,
		rootCertificate2.SerialNumber,
	)
	_, err = setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	nocCerts, _ = utils.QueryAllNocCertificates(setup)
	require.Equal(t, 1, len(nocCerts))
	require.Equal(t, 1, len(nocCerts[0].Certs))
	require.Equal(t, testconstants.NocCert1SerialNumber, nocCerts[0].Certs[0].SerialNumber)

	// Check indexes for root certificates
	indexes = []utils.TestIndex{
		{Key: types.AllCertificatesKeyPrefix, Exist: false},
		{Key: types.AllCertificatesBySubjectKeyPrefix, Exist: false},
		{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Exist: false},
		{Key: types.NocCertificatesKeyPrefix, Exist: false},
		{Key: types.NocCertificatesBySubjectKeyPrefix, Exist: false},
		{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix, Exist: false},
		{Key: types.NocCertificatesByVidAndSkidKeyPrefix, Exist: false},
		{Key: types.NocRootCertificatesKeyPrefix, Exist: false},
		{Key: types.UniqueCertificateKeyPrefix, Exist: false},
		{Key: types.RevokedNocIcaCertificatesKeyPrefix, Exist: false},
		{Key: types.RevokedNocRootCertificatesKeyPrefix, Exist: false},
		{Key: types.RevokedCertificatesKeyPrefix, Exist: false},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate1, indexes)
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate2, indexes)

	// Check indexes for intermediate certificates
	indexes = []utils.TestIndex{
		{Key: types.AllCertificatesKeyPrefix, Exist: true},
		{Key: types.AllCertificatesBySubjectKeyPrefix, Exist: true},
		{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Exist: true},
		{Key: types.NocCertificatesKeyPrefix, Exist: true},
		{Key: types.NocCertificatesBySubjectKeyPrefix, Exist: true},
		{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix, Exist: true},
		{Key: types.NocCertificatesByVidAndSkidKeyPrefix, Exist: true},
		{Key: types.NocIcaCertificatesKeyPrefix, Exist: true},
		{Key: types.UniqueCertificateKeyPrefix, Exist: true},
		{Key: types.ChildCertificatesKeyPrefix, Exist: true},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate, indexes)
}

func TestHandler_RemoveNocX509RootCert_RevokedCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.Vid)

	// add NOC root certificate
	rootCertificate1 := utils.CreateTestNocRoot1Cert()
	utils.AddNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1)

	rootCertificate2 := utils.CreateTestNocRoot2Cert()
	utils.AddNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1Copy)

	// Add an intermediate certificate
	icaCertificate := utils.CreateTestNocIca1Cert()
	utils.AddNocIntermediateCertificate(setup, vendorAccAddress, testconstants.NocCert1)

	// revoke NOC root certificates
	revokeX509Cert := types.NewMsgRevokeNocX509RootCert(
		vendorAccAddress.String(),
		rootCertificate1.Subject,
		rootCertificate1.SubjectKeyID,
		"",
		testconstants.Info,
		false,
	)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(t, err)

	// Check indexes for root certificates
	indexes := []utils.TestIndex{
		{Key: types.UniqueCertificateKeyPrefix, Exist: true},
		{Key: types.RevokedNocRootCertificatesKeyPrefix, Exist: true},
		{Key: types.AllCertificatesKeyPrefix, Exist: false},
		{Key: types.AllCertificatesBySubjectKeyPrefix, Exist: false},
		{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Exist: false},
		{Key: types.NocCertificatesKeyPrefix, Exist: false},
		{Key: types.NocCertificatesBySubjectKeyPrefix, Exist: false},
		{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix, Exist: false},
		{Key: types.NocCertificatesByVidAndSkidKeyPrefix, Exist: false},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate1, indexes)
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate2, indexes)

	revokedCerts, _ := utils.QueryNocRevokedRootCertificates(setup, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(t, 2, len(revokedCerts.Certs))

	// Check that intermediate certificates does not exist
	indexes = []utils.TestIndex{
		{Key: types.AllCertificatesKeyPrefix, Exist: true},
		{Key: types.AllCertificatesBySubjectKeyPrefix, Exist: true},
		{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Exist: true},
		{Key: types.NocCertificatesKeyPrefix, Exist: true},
		{Key: types.NocCertificatesBySubjectKeyPrefix, Exist: true},
		{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix, Exist: true},
		{Key: types.NocCertificatesByVidAndSkidKeyPrefix, Exist: true},
		{Key: types.NocIcaCertificatesKeyPrefix, Exist: true},
		{Key: types.UniqueCertificateKeyPrefix, Exist: true},
		{Key: types.ChildCertificatesKeyPrefix, Exist: true},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate, indexes)

	// remove NOC root certificates
	removeIcaCert := types.NewMsgRemoveNocX509RootCert(
		vendorAccAddress.String(),
		rootCertificate1.Subject,
		rootCertificate1.SubjectKeyID,
		"",
	)
	_, err = setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	allCerts, _ := utils.QueryAllNocCertificates(setup)
	require.Equal(t, 1, len(allCerts))
	require.Equal(t, testconstants.NocCert1SerialNumber, allCerts[0].Certs[0].SerialNumber)

	// Check indexes for root certificates
	indexes = []utils.TestIndex{
		{Key: types.UniqueCertificateKeyPrefix, Exist: false},
		{Key: types.AllCertificatesKeyPrefix, Exist: false},
		{Key: types.AllCertificatesBySubjectKeyPrefix, Exist: false},
		{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Exist: false},
		{Key: types.NocCertificatesKeyPrefix, Exist: false},
		{Key: types.NocCertificatesBySubjectKeyPrefix, Exist: false},
		{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix, Exist: false},
		{Key: types.NocCertificatesByVidAndSkidKeyPrefix, Exist: false},
		{Key: types.RevokedNocRootCertificatesKeyPrefix, Exist: false},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate1, indexes)
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate2, indexes)

	// Check that intermediate certificates still exist
	indexes = []utils.TestIndex{
		{Key: types.AllCertificatesKeyPrefix, Exist: true},
		{Key: types.AllCertificatesBySubjectKeyPrefix, Exist: true},
		{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Exist: true},
		{Key: types.NocCertificatesKeyPrefix, Exist: true},
		{Key: types.NocCertificatesBySubjectKeyPrefix, Exist: true},
		{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix, Exist: true},
		{Key: types.NocCertificatesByVidAndSkidKeyPrefix, Exist: true},
		{Key: types.NocIcaCertificatesKeyPrefix, Exist: true},
		{Key: types.UniqueCertificateKeyPrefix, Exist: true},
		{Key: types.ChildCertificatesKeyPrefix, Exist: true},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate, indexes)
}

// Extra cases

func TestHandler_RemoveNocX509RootCert_RevokedAndActiveCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	utils.AddNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1)

	// Add an intermediate certificate
	addIcaCert := types.NewMsgAddNocX509IcaCert(vendorAccAddress.String(), testconstants.NocCert1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addIcaCert)
	require.NoError(t, err)

	// get certificates for further comparison
	nocCerts := setup.Keeper.GetAllNocCertificates(setup.Ctx)
	require.NotNil(t, nocCerts)
	require.Equal(t, 2, len(nocCerts))

	// revoke an intermediate certificate
	revokeX509Cert := types.NewMsgRevokeNocX509RootCert(
		vendorAccAddress.String(),
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		testconstants.NocRootCert1SerialNumber,
		testconstants.Info,
		false,
	)
	_, err = setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(t, err)

	// Add NOC root certificate with new serial number
	utils.AddNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1Copy)

	certs, _ := utils.QueryNocCertificates(setup, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(t, 1, len(certs.Certs))
	require.Equal(t, testconstants.NocRootCert1CopySerialNumber, certs.Certs[0].SerialNumber)

	// remove NOC root certificate by serial number
	removeIcaCert := types.NewMsgRemoveNocX509RootCert(
		vendorAccAddress.String(),
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		testconstants.NocRootCert1SerialNumber,
	)
	_, err = setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	// check that only one root and IAC certificates exists
	nocCerts, _ = utils.QueryAllNocCertificates(setup)
	require.Equal(t, 2, len(nocCerts))

	certs, _ = utils.QueryNocCertificates(setup, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(t, testconstants.NocRootCert1CopySerialNumber, certs.Certs[0].SerialNumber)
	certs, _ = utils.QueryNocCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, 1, len(certs.Certs))

	_, err = utils.QueryNocRevokedRootCertificates(setup, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificates does not exists
	found := setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SerialNumber)
	require.Equal(t, false, found)
	found = setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocRootCert1Subject, testconstants.NocRootCert1CopySerialNumber)
	require.Equal(t, true, found)

	// query noc certificate by VID
	nocCertificates, err := utils.QueryNocIcaCertificatesByVid(setup, vid)
	require.NoError(t, err)
	require.Equal(t, len(nocCertificates.Certs), 1)
	require.Equal(t, testconstants.NocCert1SerialNumber, nocCertificates.Certs[0].SerialNumber)

	// Add NOC root certificate with new serial number
	utils.AddNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1)

	certs, _ = utils.QueryNocCertificates(setup, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(t, 2, len(certs.Certs))

	// remove NOC root certificates
	removeIcaCert = types.NewMsgRemoveNocX509RootCert(
		vendorAccAddress.String(),
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		"",
	)
	_, err = setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	nocCerts, _ = utils.QueryAllNocCertificates(setup)
	require.Equal(t, 1, len(nocCerts))
	require.Equal(t, 1, len(nocCerts[0].Certs))
	require.Equal(t, testconstants.NocCert1SerialNumber, nocCerts[0].Certs[0].SerialNumber)

	nocCertificates, err = utils.QueryNocIcaCertificatesByVid(setup, vid)
	require.NoError(t, err)
	require.Equal(t, len(nocCertificates.Certs), 1)
	require.Equal(t, testconstants.NocCert1SerialNumber, nocCertificates.Certs[0].SerialNumber)

	// check that IAC certificates can be queried by vid+skid
	certsByVidSkid, _ := utils.QueryNocCertificatesByVidAndSkid(setup, vid, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, 1, len(certsByVidSkid.Certs))
	require.Equal(t, testconstants.NocCert1SerialNumber, certsByVidSkid.Certs[0].SerialNumber)

	// check that root certs removed
	_, err = utils.QueryNocCertificates(setup, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))
	_, err = utils.QueryNocCertificatesBySubject(setup, testconstants.NocRootCert1Subject)
	require.Equal(t, codes.NotFound, status.Code(err))
	certsBySKID, _ := utils.QueryNocCertificatesBySubjectKeyID(setup, testconstants.NocRootCert1SubjectKeyID)
	require.Empty(t, certsBySKID)
	_, err = utils.QueryNocRootCertificatesByVid(setup, vid)
	require.Equal(t, codes.NotFound, status.Code(err))
	_, err = utils.QueryNocRootCertificatesByVid(setup, vid)
	require.Equal(t, codes.NotFound, status.Code(err))
	_, err = utils.QueryNocCertificatesByVidAndSkid(setup, vid, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificates does not exists
	found = setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SerialNumber)
	require.Equal(t, false, found)
	found = setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocRootCert1Subject, testconstants.NocRootCert1CopySerialNumber)
	require.Equal(t, false, found)
}

func TestHandler_RemoveNocX509RootCert_ByNotOwnerButSameVendor(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	utils.AddNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1)

	// add first vendor account with VID = 1
	vendorAccAddress1 := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add second vendor account with VID = 1
	vendorAccAddress2 := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// remove x509 certificate by second vendor account
	removeIcaCert := types.NewMsgRemoveNocX509RootCert(
		vendorAccAddress2.String(),
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		testconstants.NocRootCert1SerialNumber,
	)
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	// check that certificate removed from 'noc certificates' list
	_, err = utils.QueryNocCertificates(setup, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that certificate removed from 'noc certificates by subject' list
	_, err = utils.QueryNocCertificatesBySubject(setup, testconstants.NocRootCert1Subject)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that certificate removed from 'noc certificates by SKID' list
	nocCerts, err := utils.QueryNocCertificatesBySubjectKeyID(setup, testconstants.NocRootCert1SubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 0, len(nocCerts))

	// query noc certificate by VID
	_, err = utils.QueryNocRootCertificatesByVid(setup, vid)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificate key is not registered
	require.False(t, setup.Keeper.IsUniqueCertificatePresent(setup.Ctx,
		testconstants.NocRootCert1Subject, testconstants.NocRootCert1SerialNumber))
}

// Error cases
func TestHandler_RemoveNocX509RootCert_CertificateDoesNotExist(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	removeIcaCert := types.NewMsgRemoveNocX509RootCert(
		vendorAccAddress.String(), testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID, testconstants.NocRootCert1SerialNumber)
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_RemoveNocX509RootCert_EmptyCertificatesList(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	setup.Keeper.SetNocRootCertificates(
		setup.Ctx,
		types.NocRootCertificates{
			Vid: vid,
		},
	)

	removeIcaCert := types.NewMsgRemoveNocX509RootCert(
		vendorAccAddress.String(), testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID, "")
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_RemoveNocX509RootCert_ByOtherVendor(t *testing.T) {
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

	// add second vendor account with VID = 1000
	vendorAccAddress2 := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID1)

	// remove ICA certificate by second vendor account
	removeIcaCert := types.NewMsgRemoveNocX509RootCert(
		vendorAccAddress2.String(), testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID, testconstants.NocRootCert1SerialNumber)
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertVidNotEqualAccountVid.Is(err))
}

func TestHandler_RemoveNocX509RootCert_SenderNotVendor(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	utils.AddNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1)

	removeIcaCert := types.NewMsgRemoveNocX509RootCert(
		setup.Trustee1.String(), testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID, "")
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_RemoveNocX509RootCert_InvalidSerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	utils.AddNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1)

	removeX509Cert := types.NewMsgRemoveNocX509RootCert(
		vendorAccAddress.String(), testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID, "invalid")
	_, err := setup.Handler(setup.Ctx, removeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}
