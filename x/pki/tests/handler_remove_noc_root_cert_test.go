package tests

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Main

func TestHandler_RemoveNocX509RootCert_BySubjectAndSKID(t *testing.T) {
	setup := Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificates
	addNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1)
	addNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1Copy)

	// Add intermediate certificate
	addNocIcaCertificate(setup, vendorAccAddress, testconstants.NocCert1)

	// get certificates for further comparison
	nocCerts := setup.Keeper.GetAllNocCertificates(setup.Ctx)
	require.NotNil(t, nocCerts)
	require.Equal(t, 2, len(nocCerts))
	require.Equal(t, 3, len(nocCerts[0].Certs)+len(nocCerts[1].Certs))

	// remove all root nOC certificates but IAC certificate
	removeIcaCert := types.NewMsgRemoveNocX509RootCert(
		vendorAccAddress.String(),
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		"",
	)
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	// check that only IAC certificate exists
	nocCerts, _ = queryAllNocCertificates(setup)
	require.Equal(t, 1, len(nocCerts))
	require.Equal(t, 1, len(nocCerts[0].Certs))
	require.Equal(t, testconstants.NocCert1SerialNumber, nocCerts[0].Certs[0].SerialNumber)

	// Check that root certificates does not exist
	ensureNocRootCertificateDoesNotExist(
		t,
		setup,
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		testconstants.NocCert1Issuer,
		testconstants.NocRootCert1SerialNumber,
		vid,
		true, // intermediate certificate with the same vid exists
		false)

	// Check that root copy certificates does not exist
	ensureNocRootCertificateDoesNotExist(
		t,
		setup,
		testconstants.NocRootCert1CopySubject,
		testconstants.NocRootCert1CopySubjectKeyID,
		testconstants.NocCert1Issuer,
		testconstants.NocRootCert1CopySerialNumber,
		vid,
		true, // intermediate certificate with the same vid exists
		false)

	// Check that intermediate certificates does not exist
	ensureNocIcaCertificateExist(
		t,
		setup,
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1Issuer,
		testconstants.NocCert1SerialNumber,
		vid,
		false)
}

func TestHandler_RemoveNocX509RootCert_BySerialNumber(t *testing.T) {
	setup := Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificates
	addNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1)
	addNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1Copy)

	// Add ICA certificates
	addNocIcaCertificate(setup, vendorAccAddress, testconstants.NocCert1)

	// remove NOC root certificate by serial number
	removeIcaCert := types.NewMsgRemoveNocX509RootCert(
		vendorAccAddress.String(),
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		testconstants.NocRootCert1SerialNumber,
	)
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	nocCerts, _ := queryAllNocCertificates(setup)
	require.Equal(t, 2, len(nocCerts))

	// NocCertificates: Subject and SKID
	nocCertificates, err := queryNocCertificates(setup, testconstants.NocRootCert1CopySubject, testconstants.NocRootCert1CopySubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 1, len(nocCertificates.Certs))

	// Check that root copy certificates does not exist
	ensureNocRootCertificateExist(
		t,
		setup,
		testconstants.NocRootCert1CopySubject,
		testconstants.NocRootCert1CopySubjectKeyID,
		testconstants.NocCert1Issuer,
		testconstants.NocRootCert1CopySerialNumber,
		vid)

	// Check that intermediate certificates does not exist
	ensureNocIcaCertificateExist(
		t,
		setup,
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1Issuer,
		testconstants.NocCert1SerialNumber,
		vid,
		false)

	// remove NOC root certificate by serial number and check that IAC cert is not removed
	removeIcaCert = types.NewMsgRemoveNocX509RootCert(
		vendorAccAddress.String(),
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		testconstants.NocRootCert1CopySerialNumber,
	)
	_, err = setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	nocCerts, _ = queryAllNocCertificates(setup)
	require.Equal(t, 1, len(nocCerts))
	require.Equal(t, 1, len(nocCerts[0].Certs))
	require.Equal(t, testconstants.NocCert1SerialNumber, nocCerts[0].Certs[0].SerialNumber)

	// Check that root certificates does not exist
	ensureNocRootCertificateDoesNotExist(
		t,
		setup,
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		testconstants.NocCert1Issuer,
		testconstants.NocRootCert1SerialNumber,
		vid,
		true, // intermediate certificate with the same vid exists
		false)

	// Check that root copy certificates does not exist
	ensureNocRootCertificateDoesNotExist(
		t,
		setup,
		testconstants.NocRootCert1CopySubject,
		testconstants.NocRootCert1CopySubjectKeyID,
		testconstants.NocCert1Issuer,
		testconstants.NocRootCert1CopySerialNumber,
		vid,
		true, // intermediate certificate with the same vid exists
		false)

	// Check that intermediate certificates does not exist
	ensureNocIcaCertificateExist(
		t,
		setup,
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1Issuer,
		testconstants.NocCert1SerialNumber,
		vid,
		false)
}

func TestHandler_RemoveNocX509RootCert_RevokedCertificate(t *testing.T) {
	setup := Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	addNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1)
	addNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1Copy)

	// Add an intermediate certificate
	addNocIcaCertificate(setup, vendorAccAddress, testconstants.NocCert1)

	// revoke NOC root certificates
	revokeX509Cert := types.NewMsgRevokeNocX509RootCert(
		vendorAccAddress.String(),
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		"",
		testconstants.Info,
		false,
	)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(t, err)

	// Check that root copy certificates does not exist
	ensureNocRootCertificateDoesNotExist(
		t,
		setup,
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		testconstants.NocCert1Issuer,
		testconstants.NocRootCert1SerialNumber,
		vid,
		true, // intermediate certificate with the same vid exists
		true)

	// Check that root copy certificates does not exist
	ensureNocRootCertificateDoesNotExist(
		t,
		setup,
		testconstants.NocRootCert1CopySubject,
		testconstants.NocRootCert1CopySubjectKeyID,
		testconstants.NocCert1Issuer,
		testconstants.NocRootCert1CopySerialNumber,
		vid,
		true, // intermediate certificate with the same vid exists
		true)

	revokedCerts, _ := queryRevokedNocRootCertificates(setup, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(t, 2, len(revokedCerts.Certs))
	require.Equal(t, testconstants.NocRootCert1Subject, revokedCerts.Certs[0].Subject)
	require.Equal(t, testconstants.NocRootCert1SubjectKeyID, revokedCerts.Certs[0].SubjectKeyId)
	require.Equal(t, testconstants.NocRootCert1CopySubject, revokedCerts.Certs[1].Subject)
	require.Equal(t, testconstants.NocRootCert1CopySubjectKeyID, revokedCerts.Certs[1].SubjectKeyId)

	// Check that intermediate certificates does not exist
	ensureNocIcaCertificateExist(
		t,
		setup,
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1Issuer,
		testconstants.NocCert1SerialNumber,
		vid,
		false)

	// remove NOC root certificates
	removeIcaCert := types.NewMsgRemoveNocX509RootCert(
		vendorAccAddress.String(),
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		"",
	)
	_, err = setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	allCerts, _ := queryAllNocCertificates(setup)
	require.Equal(t, 1, len(allCerts))
	require.Equal(t, testconstants.NocCert1SerialNumber, allCerts[0].Certs[0].SerialNumber)

	// Check that intermediate certificates does not exist
	ensureNocIcaCertificateExist(
		t,
		setup,
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1Issuer,
		testconstants.NocCert1SerialNumber,
		vid,
		false)

	// Check that root copy certificates does not exist
	ensureNocRootCertificateDoesNotExist(
		t,
		setup,
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		testconstants.NocCert1Issuer,
		testconstants.NocRootCert1SerialNumber,
		vid,
		true, // intermediate certificate with the same vid exists
		true)

	// Check that root copy certificates does not exist
	ensureNocRootCertificateDoesNotExist(
		t,
		setup,
		testconstants.NocRootCert1CopySubject,
		testconstants.NocRootCert1CopySubjectKeyID,
		testconstants.NocCert1Issuer,
		testconstants.NocRootCert1CopySerialNumber,
		vid,
		true, // intermediate certificate with the same vid exists
		true)

	// Check that revoked certificate does not exist
	_, err = queryRevokedNocRootCertificates(setup, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))
}

// Extra cases

func TestHandler_RemoveNocX509RootCert_RevokedAndActiveCertificate(t *testing.T) {
	setup := Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	addNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1)

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
	addNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1Copy)

	certs, _ := queryNocCertificates(setup, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
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
	nocCerts, _ = queryAllNocCertificates(setup)
	require.Equal(t, 2, len(nocCerts))

	certs, _ = queryNocCertificates(setup, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(t, testconstants.NocRootCert1CopySerialNumber, certs.Certs[0].SerialNumber)
	certs, _ = queryNocCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, 1, len(certs.Certs))

	_, err = queryRevokedNocRootCertificates(setup, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificates does not exists
	found := setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SerialNumber)
	require.Equal(t, false, found)
	found = setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocRootCert1Subject, testconstants.NocRootCert1CopySerialNumber)
	require.Equal(t, true, found)

	// query noc certificate by VID
	nocCertificates, err := queryNocIcaCertificatesByVid(setup, vid)
	require.NoError(t, err)
	require.Equal(t, len(nocCertificates.Certs), 1)
	require.Equal(t, testconstants.NocCert1SerialNumber, nocCertificates.Certs[0].SerialNumber)

	// Add NOC root certificate with new serial number
	addNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1)

	certs, _ = queryNocCertificates(setup, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
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

	nocCerts, _ = queryAllNocCertificates(setup)
	require.Equal(t, 1, len(nocCerts))
	require.Equal(t, 1, len(nocCerts[0].Certs))
	require.Equal(t, testconstants.NocCert1SerialNumber, nocCerts[0].Certs[0].SerialNumber)

	nocCertificates, err = queryNocIcaCertificatesByVid(setup, vid)
	require.NoError(t, err)
	require.Equal(t, len(nocCertificates.Certs), 1)
	require.Equal(t, testconstants.NocCert1SerialNumber, nocCertificates.Certs[0].SerialNumber)

	// check that IAC certificates can be queried by vid+skid
	certsByVidSkid, _ := queryNocCertificatesByVidAndSkid(setup, vid, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, 1, len(certsByVidSkid.Certs))
	require.Equal(t, testconstants.NocCert1SerialNumber, certsByVidSkid.Certs[0].SerialNumber)

	// check that root certs removed
	_, err = queryNocCertificates(setup, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))
	_, err = queryNocCertificatesBySubject(setup, testconstants.NocRootCert1Subject)
	require.Equal(t, codes.NotFound, status.Code(err))
	certsBySKID, _ := queryAllNocCertificatesBySubjectKeyID(setup, testconstants.NocRootCert1SubjectKeyID)
	require.Empty(t, certsBySKID)
	_, err = queryNocRootCertificates(setup, vid)
	require.Equal(t, codes.NotFound, status.Code(err))
	_, err = queryNocCertificatesByVidAndSkid(setup, vid, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificates does not exists
	found = setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SerialNumber)
	require.Equal(t, false, found)
	found = setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocRootCert1Subject, testconstants.NocRootCert1CopySerialNumber)
	require.Equal(t, false, found)
}

func TestHandler_RemoveNocX509RootCert_ByNotOwnerButSameVendor(t *testing.T) {
	setup := Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	addNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1)

	// add first vendor account with VID = 1
	vendorAccAddress1 := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add second vendor account with VID = 1
	vendorAccAddress2 := GenerateAccAddress()
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
	_, err = queryNocCertificates(setup, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that certificate removed from 'noc certificates by subject' list
	_, err = queryNocCertificatesBySubject(setup, testconstants.NocRootCert1Subject)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that certificate removed from 'noc certificates by SKID' list
	nocCerts, err := queryAllNocCertificatesBySubjectKeyID(setup, testconstants.NocRootCert1SubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 0, len(nocCerts))

	// query noc certificate by VID
	_, err = queryNocRootCertificates(setup, vid)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificate key is not registered
	require.False(t, setup.Keeper.IsUniqueCertificatePresent(setup.Ctx,
		testconstants.NocRootCert1Subject, testconstants.NocRootCert1SerialNumber))
}

// Error cases
func TestHandler_RemoveNocX509RootCert_CertificateDoesNotExist(t *testing.T) {
	setup := Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	removeIcaCert := types.NewMsgRemoveNocX509RootCert(
		vendorAccAddress.String(), testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID, testconstants.NocRootCert1SerialNumber)
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_RemoveNocX509RootCert_EmptyCertificatesList(t *testing.T) {
	setup := Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := GenerateAccAddress()
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
	setup := Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	addNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1)

	// add fist vendor account with VID = 1
	vendorAccAddress1 := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add second vendor account with VID = 1000
	vendorAccAddress2 := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID1)

	// remove ICA certificate by second vendor account
	removeIcaCert := types.NewMsgRemoveNocX509RootCert(
		vendorAccAddress2.String(), testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID, testconstants.NocRootCert1SerialNumber)
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertVidNotEqualAccountVid.Is(err))
}

func TestHandler_RemoveNocX509RootCert_SenderNotVendor(t *testing.T) {
	setup := Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	addNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1)

	removeIcaCert := types.NewMsgRemoveNocX509RootCert(
		setup.Trustee1.String(), testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID, "")
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_RemoveNocX509RootCert_InvalidSerialNumber(t *testing.T) {
	setup := Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	addNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1)

	removeX509Cert := types.NewMsgRemoveNocX509RootCert(
		vendorAccAddress.String(), testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID, "invalid")
	_, err := setup.Handler(setup.Ctx, removeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}
