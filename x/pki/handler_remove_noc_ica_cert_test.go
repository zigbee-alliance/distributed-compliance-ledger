package pki

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

func TestHandler_RemoveNocX509IcaCert_BySubjectAndSKID(t *testing.T) {
	setup := Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	addNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1, vid)

	// Add two intermediate certificates
	addIcaCert := types.NewMsgAddNocX509IcaCert(vendorAccAddress.String(), testconstants.NocCert1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addIcaCert)
	require.NoError(t, err)
	addIcaCert = types.NewMsgAddNocX509IcaCert(vendorAccAddress.String(), testconstants.NocCert1Copy, testconstants.CertSchemaVersion)
	_, err = setup.Handler(setup.Ctx, addIcaCert)
	require.NoError(t, err)

	// Add a leaf certificate
	addIcaLeafCert := types.NewMsgAddNocX509IcaCert(vendorAccAddress.String(), testconstants.NocLeafCert1, testconstants.CertSchemaVersion)
	_, err = setup.Handler(setup.Ctx, addIcaLeafCert)
	require.NoError(t, err)

	// get certificates for further comparison
	nocCerts := setup.Keeper.GetAllNocCertificates(setup.Ctx)
	require.NotNil(t, nocCerts)
	require.Equal(t, 3, len(nocCerts))
	require.Equal(t, 4, len(nocCerts[0].Certs)+len(nocCerts[1].Certs)+len(nocCerts[2].Certs))

	// remove all intermediate certificates but leave leaf certificate
	removeIcaCert := types.NewMsgRemoveNocX509IcaCert(
		vendorAccAddress.String(),
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		"",
	)
	_, err = setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	// check that only root and leaf certificates exists
	nocCerts, _ = queryAllNocCertificates(setup)
	require.Equal(t, 2, len(nocCerts))
	require.Equal(t, 2, len(nocCerts[0].Certs)+len(nocCerts[1].Certs))
	_, err = queryNocCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))
	// check that unique certificates does not exists
	found := setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocCert1Issuer, testconstants.NocCert1SerialNumber)
	require.Equal(t, false, found)
	found = setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocCert1Issuer, testconstants.NocCert1CopySerialNumber)
	require.Equal(t, false, found)

	// check that intermediate certificate can not be queried by vid+skid
	_, err = queryNocCertificatesByVidAndSkid(setup, vid, testconstants.NocCert1SubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	leafCerts, _ := queryNocCertificates(setup, testconstants.NocLeafCert1Subject, testconstants.NocLeafCert1SubjectKeyID)
	require.Equal(t, 1, len(leafCerts.Certs))
	require.Equal(t, testconstants.NocLeafCert1SerialNumber, leafCerts.Certs[0].SerialNumber)

	// query noc certificate by VID
	nocCertificates, err := queryNocCertificatesByVid(setup, vid)
	require.NoError(t, err)
	require.Equal(t, len(nocCertificates.Certs), 1)
	require.Equal(t, testconstants.NocLeafCert1Subject, nocCertificates.Certs[0].Subject)
	require.Equal(t, testconstants.NocLeafCert1SubjectKeyID, nocCertificates.Certs[0].SubjectKeyId)
}

func TestHandler_RemoveNocX509IcaCert_BySerialNumber(t *testing.T) {
	setup := Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	addNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1, vid)

	// Add ICA certificates
	addIcaCert := types.NewMsgAddNocX509IcaCert(vendorAccAddress.String(), testconstants.NocCert1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addIcaCert)
	require.NoError(t, err)

	// Add ICA certificates with sam subject and SKID but different serial number
	addIcaCert = types.NewMsgAddNocX509IcaCert(vendorAccAddress.String(), testconstants.NocCert1Copy, testconstants.CertSchemaVersion)
	_, err = setup.Handler(setup.Ctx, addIcaCert)
	require.NoError(t, err)

	// Add a leaf certificate
	addIcaLeafCert := types.NewMsgAddNocX509IcaCert(vendorAccAddress.String(), testconstants.NocLeafCert1, testconstants.CertSchemaVersion)
	_, err = setup.Handler(setup.Ctx, addIcaLeafCert)
	require.NoError(t, err)

	intermediateCerts, _ := queryNocCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
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
	_, err = setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	// check that only root, intermediate(with serial number 3) and leaf certificates exists
	allCerts, _ := queryAllNocCertificates(setup)
	require.Equal(t, 3, len(allCerts))
	require.Equal(t, 3, len(allCerts[0].Certs)+len(allCerts[1].Certs)+len(allCerts[2].Certs))
	leafCerts, _ := queryNocCertificates(setup, testconstants.NocLeafCert1Subject, testconstants.NocLeafCert1SubjectKeyID)
	require.Equal(t, 1, len(leafCerts.Certs))

	intermediateCerts, _ = queryNocCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, 1, len(intermediateCerts.Certs))
	require.Equal(t, testconstants.NocCert1CopySerialNumber, intermediateCerts.Certs[0].SerialNumber)

	// remove  intermediate certificate by serial number and check that leaf cert is not removed
	removeIcaCert = types.NewMsgRemoveNocX509IcaCert(
		vendorAccAddress.String(),
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1CopySerialNumber,
	)
	_, err = setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	allCerts, _ = queryAllNocCertificates(setup)
	require.Equal(t, 2, len(allCerts))
	require.Equal(t, 2, len(allCerts[0].Certs)+len(allCerts[1].Certs))

	_, err = queryNocCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificates does not exists
	found := setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocCert1Issuer, testconstants.NocCert1SerialNumber)
	require.Equal(t, false, found)
	found = setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocCert1Issuer, testconstants.NocCert1CopySerialNumber)
	require.Equal(t, false, found)

	// check that intermediate certificate can not be queried by vid+skid
	_, err = queryNocCertificatesByVidAndSkid(setup, vid, testconstants.NocCert1SubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	leafCerts, _ = queryNocCertificates(setup, testconstants.NocLeafCert1Subject, testconstants.NocLeafCert1SubjectKeyID)
	require.Equal(t, 1, len(leafCerts.Certs))

	// query noc certificate by VID
	nocCertificates, err := queryNocCertificatesByVid(setup, vid)
	require.NoError(t, err)
	require.Equal(t, len(nocCertificates.Certs), 1)
	require.Equal(t, testconstants.NocLeafCert1Subject, nocCertificates.Certs[0].Subject)
	require.Equal(t, testconstants.NocLeafCert1SubjectKeyID, nocCertificates.Certs[0].SubjectKeyId)
}

func TestHandler_RemoveNocX509IcaCert_RevokedAndActiveCertificate(t *testing.T) {
	setup := Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	addNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1, vid)

	// Add an intermediate certificate
	addIcaCert := types.NewMsgAddNocX509IcaCert(vendorAccAddress.String(), testconstants.NocCert1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addIcaCert)
	require.NoError(t, err)

	// get certificates for further comparison
	allCerts := setup.Keeper.GetAllNocCertificates(setup.Ctx)
	require.NotNil(t, allCerts)
	require.Equal(t, 2, len(allCerts))
	require.Equal(t, 2, len(allCerts[0].Certs)+len(allCerts[1].Certs))

	// revoke an intermediate certificate
	revokeX509Cert := types.NewMsgRevokeNocX509IcaCert(
		vendorAccAddress.String(),
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1SerialNumber,
		testconstants.Info,
		false,
	)
	_, err = setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(t, err)

	// Add an intermediate certificate with new serial number
	addIcaCert = types.NewMsgAddNocX509IcaCert(vendorAccAddress.String(), testconstants.NocCert1Copy, testconstants.CertSchemaVersion)
	_, err = setup.Handler(setup.Ctx, addIcaCert)
	require.NoError(t, err)

	intermediateCerts, _ := queryNocCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, 1, len(intermediateCerts.Certs))
	require.Equal(t, testconstants.NocCert1Subject, intermediateCerts.Certs[0].Subject)
	require.Equal(t, testconstants.NocCert1SubjectKeyID, intermediateCerts.Certs[0].SubjectKeyId)
	require.Equal(t, testconstants.NocCert1CopySerialNumber, intermediateCerts.Certs[0].SerialNumber)

	// remove an intermediate certificate
	removeIcaCert := types.NewMsgRemoveNocX509IcaCert(
		vendorAccAddress.String(),
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		"",
	)
	_, err = setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	// check that only root and leaf certificates exists
	allCerts, _ = queryAllNocCertificates(setup)
	require.Equal(t, 1, len(allCerts))
	require.Equal(t, true, allCerts[0].Certs[0].IsRoot)
	_, err = queryNocCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))
	_, err = queryRevokedNocIcaCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificates does not exists
	found := setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocCert1Issuer, testconstants.NocCert1SerialNumber)
	require.Equal(t, false, found)
	found = setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocCert1Issuer, testconstants.NocCert1CopySerialNumber)
	require.Equal(t, false, found)

	// check that intermediate certificate can not be queried by vid+skid
	_, err = queryNocCertificatesByVidAndSkid(setup, vid, testconstants.NocCert1SubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// query noc certificate by VID
	_, err = queryNocCertificatesByVid(setup, vid)
	require.Equal(t, codes.NotFound, status.Code(err))
}

func TestHandler_RemoveNocX509IcaCert_RevokedCertificate(t *testing.T) {
	setup := Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	addNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1, vid)

	// Add an intermediate certificate
	addIcaCert := types.NewMsgAddNocX509IcaCert(vendorAccAddress.String(), testconstants.NocCert1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addIcaCert)
	require.NoError(t, err)

	intermediateCerts, _ := queryNocCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, 1, len(intermediateCerts.Certs))
	require.Equal(t, testconstants.NocCert1Subject, intermediateCerts.Certs[0].Subject)
	require.Equal(t, testconstants.NocCert1SubjectKeyID, intermediateCerts.Certs[0].SubjectKeyId)

	// revoke intermediate certificate by serial number
	revokeX509Cert := types.NewMsgRevokeNocX509IcaCert(
		vendorAccAddress.String(),
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1SerialNumber,
		testconstants.Info,
		false,
	)
	_, err = setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(t, err)

	_, err = queryNocCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))

	revokedCerts, _ := queryRevokedNocIcaCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, 1, len(revokedCerts.Certs))
	require.Equal(t, testconstants.NocCert1Subject, revokedCerts.Certs[0].Subject)
	require.Equal(t, testconstants.NocCert1SubjectKeyID, revokedCerts.Certs[0].SubjectKeyId)

	// remove  intermediate certificate by serial number
	removeIcaCert := types.NewMsgRemoveNocX509IcaCert(
		vendorAccAddress.String(),
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1SerialNumber,
	)
	_, err = setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	allCerts, _ := queryAllNocCertificates(setup)
	require.Equal(t, 1, len(allCerts))
	require.Equal(t, true, allCerts[0].Certs[0].IsRoot)

	_, err = queryNocCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))
	_, err = queryRevokedNocIcaCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificate does not exists
	found := setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocCert1Issuer, testconstants.NocCert1SerialNumber)
	require.Equal(t, false, found)

	// check that intermediate certificate can not be queried by vid+skid
	_, err = queryNocCertificatesByVidAndSkid(setup, vid, testconstants.NocCert1SubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// query noc certificate by VID
	_, err = queryNocCertificatesByVid(setup, vid)
	require.Equal(t, codes.NotFound, status.Code(err))
}

func TestHandler_RemoveNocX509IcaCert_ByNotOwnerButSameVendor(t *testing.T) {
	setup := Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	addNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1, vid)

	// add first vendor account with VID = 1
	vendorAccAddress1 := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add ICA certificate by fist vendor account
	addIcaCert := types.NewMsgAddNocX509IcaCert(vendorAccAddress1.String(), testconstants.NocCert1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addIcaCert)
	require.NoError(t, err)

	// add second vendor account with VID = 1
	vendorAccAddress2 := GenerateAccAddress()
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
	_, err = queryNocCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that certificate removed from 'noc certificates by subject' list
	_, err = queryNocCertificatesBySubject(setup, testconstants.NocCert1Subject)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that certificate removed from 'noc certificates by SKID' list
	nocCerts, err := queryAllNocCertificatesBySubjectKeyID(setup, testconstants.NocCert1SubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 0, len(nocCerts))

	// query noc certificate by VID
	_, err = queryNocCertificatesByVid(setup, vid)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificate key is not registered
	require.False(t, setup.Keeper.IsUniqueCertificatePresent(setup.Ctx,
		testconstants.NocCert1Issuer, testconstants.NocCert1SerialNumber))

	// check that intermediate certificate can not be queried by vid+skid
	_, err = queryNocCertificatesByVidAndSkid(setup, vid, testconstants.NocCert1SubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
}

func TestHandler_RemoveNocX509IcaCert_CertificateDoesNotExist(t *testing.T) {
	setup := Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	removeIcaCert := types.NewMsgRemoveNocX509IcaCert(
		vendorAccAddress.String(), testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID, testconstants.NocCert1SerialNumber)
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_RemoveNocX509IcaCert_EmptyCertificatesList(t *testing.T) {
	setup := Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	addNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1, vid)

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
	setup := Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	addNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1, vid)

	// add fist vendor account with VID = 1
	vendorAccAddress1 := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add x509 certificate by `setup.Trustee`
	addX509Cert := types.NewMsgAddNocX509IcaCert(vendorAccAddress1.String(), testconstants.NocCert1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// add second vendor account with VID = 1000
	vendorAccAddress2 := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID1)

	// remove ICA certificate by second vendor account
	removeIcaCert := types.NewMsgRemoveNocX509IcaCert(
		vendorAccAddress2.String(), testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID, testconstants.NocCert1SerialNumber)
	_, err = setup.Handler(setup.Ctx, removeIcaCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertVidNotEqualAccountVid.Is(err))
}

func TestHandler_RemoveNocX509IcaCert_SenderNotVendor(t *testing.T) {
	setup := Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	addNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1, vid)

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
	setup := Setup(t)

	// Add vendor account
	vendorAccAddress := GenerateAccAddress()
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
	setup := Setup(t)

	// Add vendor account
	vid := testconstants.Vid
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	addNocRootCertificate(setup, vendorAccAddress, testconstants.NocRootCert1, vid)

	addX509Cert := types.NewMsgAddNocX509IcaCert(vendorAccAddress.String(), testconstants.NocCert1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	removeX509Cert := types.NewMsgRemoveNocX509IcaCert(
		vendorAccAddress.String(), testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID, "invalid")
	_, err = setup.Handler(setup.Ctx, removeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}
