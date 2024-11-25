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
	utils.AddNocIntermediateCertificate(setup, vendorAccAddress, testconstants.NocCert1)

	// remove intermediate certificate
	removeIcaCert := types.NewMsgRemoveNocX509IcaCert(
		vendorAccAddress.String(),
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		"",
	)
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	// Check: Noc - missing
	utils.EnsureCertificateNotPresentInNocCertificateIndexes(
		t,
		setup,
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.Vid,
		false,
		false,
	)

	// Check: All - missing
	utils.EnsureGlobalCertificateNotExist(
		t,
		setup,
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		false,
		false,
	)

	// Check: UniqueCertificate - missing
	found := setup.Keeper.IsUniqueCertificatePresent(
		setup.Ctx,
		testconstants.NocCert1Issuer,
		testconstants.NocCert1SerialNumber)
	require.False(t, found)

	// Check: RevokedCertificates (ica) - missing
	found = setup.Keeper.IsRevokedNocIcaCertificatePresent(
		setup.Ctx,
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID)
	require.False(t, found)

	// Check: child certificate  - missing
	found = setup.Keeper.IsChildCertificatePresent(
		setup.Ctx,
		testconstants.NocCert1Issuer,
		testconstants.NocCert1AuthorityKeyID)
	require.False(t, found)
}

func TestHandler_RemoveNocX509IcaCert_BySubjectAndSKID(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	utils.AddNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1)

	// add two intermediate certificates
	utils.AddNocIntermediateCertificate(setup, vendorAccAddress, testconstants.NocCert1)
	utils.AddNocIntermediateCertificate(setup, vendorAccAddress, testconstants.NocCert1Copy)
	utils.AddNocIntermediateCertificate(setup, vendorAccAddress, testconstants.NocLeafCert1)

	// get certificates for further comparison
	nocCerts := setup.Keeper.GetAllNocCertificates(setup.Ctx)
	require.NotNil(t, nocCerts)
	require.Equal(t, 3, len(nocCerts))
	require.Equal(t, 4, len(nocCerts[0].Certs)+len(nocCerts[1].Certs)+len(nocCerts[2].Certs))

	// remove all intermediate certificates but leave leaf certificate (NocCert1 and NocCert1Copy)
	removeIcaCert := types.NewMsgRemoveNocX509IcaCert(
		vendorAccAddress.String(),
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		"",
	)
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	// Check that intermediate certificates does not exist
	utils.EnsureNocIntermediateCertificateNotExist(
		t,
		setup,
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1Issuer,
		testconstants.NocCert1SerialNumber,
		vid,
		true, // leaf certificate with the same vid exists
		false)

	utils.EnsureNocIntermediateCertificateNotExist(
		t,
		setup,
		testconstants.NocCert1CopySubject,
		testconstants.NocCert1CopySubjectKeyID,
		testconstants.NocCert1CopyIssuer,
		testconstants.NocCert1CopySerialNumber,
		vid,
		true, // leaf certificate with the same vid exists
		false)

	// Check that leaf certificate exists
	utils.EnsureNocIntermediateCertificateExist(
		t,
		setup,
		testconstants.NocLeafCert1Subject,
		testconstants.NocLeafCert1SubjectKeyID,
		testconstants.NocLeafCert1Issuer,
		testconstants.NocLeafCert1SerialNumber,
		vid,
		false)

	// Check that root certificate exists
	utils.EnsureNocRootCertificateExist(
		t,
		setup,
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SerialNumber,
		vid)

	// Check that only 2 certificates exists
	nocCerts, _ = utils.QueryAllNocCertificates(setup)
	require.Equal(t, 2, len(nocCerts))
	require.Equal(t, 2, len(nocCerts[0].Certs)+len(nocCerts[1].Certs))

	// query noc certificate by VID
	nocCertificates, err := utils.QueryNocIcaCertificatesByVid(setup, vid)
	require.NoError(t, err)
	require.Equal(t, len(nocCertificates.Certs), 1)
	require.Equal(t, testconstants.NocLeafCert1Subject, nocCertificates.Certs[0].Subject)
	require.Equal(t, testconstants.NocLeafCert1SubjectKeyID, nocCertificates.Certs[0].SubjectKeyId)
}

func TestHandler_RemoveNocX509IcaCert_BySerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	utils.AddNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1)

	// Add ICA certificates
	utils.AddNocIntermediateCertificate(setup, vendorAccAddress, testconstants.NocCert1)

	// Add ICA certificates with sam subject and SKID but different serial number
	utils.AddNocIntermediateCertificate(setup, vendorAccAddress, testconstants.NocCert1Copy)

	// Add a leaf certificate
	utils.AddNocIntermediateCertificate(setup, vendorAccAddress, testconstants.NocLeafCert1)

	// get certificates for further comparison
	intermediateCerts, _ := utils.QueryNocCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, 2, len(intermediateCerts.Certs))
	require.Equal(t, testconstants.NocCert1Subject, intermediateCerts.Certs[0].Subject)
	require.Equal(t, testconstants.NocCert1SubjectKeyID, intermediateCerts.Certs[0].SubjectKeyId)

	// remove ICA certificate by serial number
	removeIcaCert := types.NewMsgRemoveNocX509IcaCert(
		vendorAccAddress.String(),
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1SerialNumber,
	)
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	// Check that only one intermediate certificate exists
	intermediateCerts, _ = utils.QueryNocCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, 1, len(intermediateCerts.Certs))

	globalIntermediateCerts, _ := utils.QueryAllCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, 1, len(globalIntermediateCerts.Certs))

	// check that 3 certificates exists
	allCerts, _ := utils.QueryAllNocCertificates(setup)
	require.Equal(t, 3, len(allCerts))
	require.Equal(t, 3, len(allCerts[0].Certs)+len(allCerts[1].Certs)+len(allCerts[2].Certs))

	// Check that intermediate certificates with NocCert1CopySerialNumber exist
	utils.EnsureNocIntermediateCertificateExist(
		t,
		setup,
		testconstants.NocCert1CopySubject,
		testconstants.NocCert1CopySubjectKeyID,
		testconstants.NocCert1CopyIssuer,
		testconstants.NocCert1CopySerialNumber,
		vid,
		true)

	// Check that leaf certificate exists
	utils.EnsureNocIntermediateCertificateExist(
		t,
		setup,
		testconstants.NocLeafCert1Subject,
		testconstants.NocLeafCert1SubjectKeyID,
		testconstants.NocLeafCert1Issuer,
		testconstants.NocLeafCert1SerialNumber,
		vid,
		true)

	// Check that root certificate exists
	utils.EnsureNocRootCertificateExist(
		t,
		setup,
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SerialNumber,
		vid)

	// remove  intermediate certificate by serial number and check that leaf cert is not removed
	removeIcaCert = types.NewMsgRemoveNocX509IcaCert(
		vendorAccAddress.String(),
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1CopySerialNumber,
	)
	_, err = setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	// check that 2 certificates exists
	allCerts, _ = utils.QueryAllNocCertificates(setup)
	require.Equal(t, 2, len(allCerts))
	require.Equal(t, 2, len(allCerts[0].Certs)+len(allCerts[1].Certs))

	// Check that intermediate certificates with NocCert1SerialNumber does not exist
	utils.EnsureNocIntermediateCertificateNotExist(
		t,
		setup,
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1Issuer,
		testconstants.NocCert1SerialNumber,
		vid,
		true, // leaf certificate with the same vid exists
		false)

	// Check that intermediate certificates with NocCert1CopySerialNumber does not exist
	utils.EnsureNocIntermediateCertificateNotExist(
		t,
		setup,
		testconstants.NocCert1CopySubject,
		testconstants.NocCert1CopySubjectKeyID,
		testconstants.NocCert1CopyIssuer,
		testconstants.NocCert1CopySerialNumber,
		vid,
		true, // leaf certificate with the same vid exists
		false)

	// Check that leaf certificate exists
	utils.EnsureNocIntermediateCertificateExist(
		t,
		setup,
		testconstants.NocLeafCert1Subject,
		testconstants.NocLeafCert1SubjectKeyID,
		testconstants.NocLeafCert1Issuer,
		testconstants.NocLeafCert1SerialNumber,
		vid,
		false)

	// Check that root certificate exists
	utils.EnsureNocRootCertificateExist(
		t,
		setup,
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SerialNumber,
		vid)

	// query noc certificate by VID
	nocCertificates, err := utils.QueryNocIcaCertificatesByVid(setup, vid)
	require.NoError(t, err)
	require.Equal(t, len(nocCertificates.Certs), 1)
	require.Equal(t, testconstants.NocLeafCert1Subject, nocCertificates.Certs[0].Subject)
	require.Equal(t, testconstants.NocLeafCert1SubjectKeyID, nocCertificates.Certs[0].SubjectKeyId)
}

func TestHandler_RemoveNocX509IcaCert_RevokedCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	utils.AddNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1)

	// Add an intermediate certificate
	utils.AddNocIntermediateCertificate(setup, vendorAccAddress, testconstants.NocCert1)

	// Check that certificate exists
	utils.EnsureNocIntermediateCertificateExist(
		t,
		setup,
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1Issuer,
		testconstants.NocCert1SerialNumber,
		vid,
		false)

	// revoke intermediate certificate by serial number
	revokeX509Cert := types.NewMsgRevokeNocX509IcaCert(
		vendorAccAddress.String(),
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1SerialNumber,
		testconstants.Info,
		false,
	)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(t, err)

	// Check that certificate does not exist
	utils.EnsureNocIntermediateCertificateNotExist(
		t,
		setup,
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1Issuer,
		testconstants.NocCert1SerialNumber,
		vid,
		false,
		true)

	// Check that revoked certificate exists
	revokedCerts, _ := utils.QueryNocRevokedIcaCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, 1, len(revokedCerts.Certs))
	require.Equal(t, testconstants.NocCert1Subject, revokedCerts.Certs[0].Subject)
	require.Equal(t, testconstants.NocCert1SubjectKeyID, revokedCerts.Certs[0].SubjectKeyId)

	// remove intermediate certificate by serial number
	removeIcaCert := types.NewMsgRemoveNocX509IcaCert(
		vendorAccAddress.String(),
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1SerialNumber,
	)
	_, err = setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	// only one root certificate exist
	allCerts, _ := utils.QueryAllNocCertificates(setup)
	require.Equal(t, 1, len(allCerts))
	require.Equal(t, true, allCerts[0].Certs[0].IsRoot)

	// Check that certificate does not exist
	utils.EnsureNocIntermediateCertificateNotExist(
		t,
		setup,
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1Issuer,
		testconstants.NocCert1SerialNumber,
		vid,
		false,
		false)

	// Check that revoked certificate does not exist
	_, err = utils.QueryNocRevokedIcaCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificate does not exists
	found := setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocCert1Issuer, testconstants.NocCert1SerialNumber)
	require.Equal(t, false, found)
}

// Extra cases

func TestHandler_RemoveNocX509IcaCert_RevokedAndActiveCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	utils.AddNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1)

	// Add an intermediate certificate
	utils.AddNocIntermediateCertificate(setup, vendorAccAddress, testconstants.NocCert1)

	// Check that certificate exists
	utils.EnsureNocIntermediateCertificateExist(
		t,
		setup,
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1Issuer,
		testconstants.NocCert1SerialNumber,
		vid,
		false)

	// revoke an intermediate certificate
	revokeX509Cert := types.NewMsgRevokeNocX509IcaCert(
		vendorAccAddress.String(),
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1SerialNumber,
		testconstants.Info,
		false,
	)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(t, err)

	// Check that certificate does not exist
	utils.EnsureNocIntermediateCertificateNotExist(
		t,
		setup,
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1Issuer,
		testconstants.NocCert1SerialNumber,
		vid,
		false,
		true) // revocation does not remove uniqueness identifier

	// Check that revoked certificate exists
	revokedNocCerts, err := utils.QueryNocRevokedIcaCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 1, len(revokedNocCerts.Certs))

	// Add an intermediate certificate with new serial number
	utils.AddNocIntermediateCertificate(setup, vendorAccAddress, testconstants.NocCert1Copy)

	// Ensure that only 1 certificate exists
	intermediateCerts, _ := utils.QueryNocCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, 1, len(intermediateCerts.Certs))

	// Check that certificate exists (with new serial number)
	utils.EnsureNocIntermediateCertificateExist(
		t,
		setup,
		testconstants.NocCert1CopySubject,
		testconstants.NocCert1CopySubjectKeyID,
		testconstants.NocCert1CopyIssuer,
		testconstants.NocCert1CopySerialNumber,
		vid,
		false)

	// remove an intermediate certificate
	removeIcaCert := types.NewMsgRemoveNocX509IcaCert(
		vendorAccAddress.String(),
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		"",
	)
	_, err = setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	// check that only root certificates exists
	allCerts, _ := utils.QueryAllNocCertificates(setup)
	require.Equal(t, 1, len(allCerts))
	require.Equal(t, true, allCerts[0].Certs[0].IsRoot)

	// Check that certificate does not exist
	utils.EnsureNocIntermediateCertificateNotExist(
		t,
		setup,
		testconstants.NocCert1CopySubject,
		testconstants.NocCert1CopySubjectKeyID,
		testconstants.NocCert1CopyIssuer,
		testconstants.NocCert1CopySerialNumber,
		vid,
		false,
		false)

	// Check that revoked certificate does not exist
	_, err = utils.QueryNocRevokedIcaCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))
}

func TestHandler_RemoveNocX509IcaCert_ByNotOwnerButSameVendor(t *testing.T) {
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
	_, err = utils.QueryNocIcaCertificatesByVid(setup, vid)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificate key is not registered
	require.False(t, setup.Keeper.IsUniqueCertificatePresent(setup.Ctx,
		testconstants.NocCert1Issuer, testconstants.NocCert1SerialNumber))

	// check that intermediate certificate can not be queried by vid+skid
	_, err = utils.QueryNocCertificatesByVidAndSkid(setup, vid, testconstants.NocCert1SubjectKeyID)
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
