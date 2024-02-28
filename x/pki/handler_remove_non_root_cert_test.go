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

func TestHandler_RemoveX509Cert_BySerialNumber(t *testing.T) {
	setup := Setup(t)
	// propose and approve x509 root certificate
	rootCertOptions := &rootCertOptions{
		pemCert:      testconstants.RootCertWithSameSubjectAndSKID1,
		subject:      testconstants.RootCertWithSameSubjectAndSKIDSubject,
		subjectKeyID: testconstants.RootCertWithSameSubjectAndSKIDSubjectKeyID,
		info:         testconstants.Info,
		vid:          65521,
	}
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// Add two intermediate certificates again
	addIntermediateX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateWithSameSubjectAndSKID1)
	_, err := setup.Handler(setup.Ctx, addIntermediateX509Cert)
	require.NoError(t, err)
	addIntermediateX509Cert = types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateWithSameSubjectAndSKID2)
	_, err = setup.Handler(setup.Ctx, addIntermediateX509Cert)
	require.NoError(t, err)

	// Add a leaf certificate
	addLeafX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.LeafCertWithSameSubjectAndSKID)
	_, err = setup.Handler(setup.Ctx, addLeafX509Cert)
	require.NoError(t, err)

	intermediateCerts, _ := queryApprovedCertificates(setup, testconstants.IntermediateCertWithSameSubjectAndSKIDSubject, testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(t, 2, len(intermediateCerts.Certs))
	require.Equal(t, testconstants.IntermediateCertWithSameSubjectAndSKIDSubject, intermediateCerts.Certs[0].Subject)
	require.Equal(t, testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID, intermediateCerts.Certs[0].SubjectKeyId)

	// remove  intermediate certificate by serial number
	removeX509Cert := types.NewMsgRemoveX509Cert(
		setup.Trustee1.String(),
		testconstants.IntermediateCertWithSameSubjectAndSKIDSubject,
		testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID,
		testconstants.IntermediateCertWithSameSubjectAndSKID1SerialNumber,
	)
	_, err = setup.Handler(setup.Ctx, removeX509Cert)
	require.NoError(t, err)

	// check that only root, intermediate(with serial number 3) and leaf certificates exists
	allCerts, _ := queryAllApprovedCertificates(setup)
	require.Equal(t, 3, len(allCerts))
	require.Equal(t, 3, len(allCerts[0].Certs)+len(allCerts[1].Certs)+len(allCerts[2].Certs))
	leafCerts, _ := queryApprovedCertificates(setup, testconstants.LeafCertWithSameSubjectAndSKIDSubject, testconstants.LeafCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(t, 1, len(leafCerts.Certs))

	intermediateCerts, _ = queryApprovedCertificates(setup, testconstants.IntermediateCertWithSameSubjectAndSKIDSubject, testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(t, 1, len(intermediateCerts.Certs))
	require.Equal(t, testconstants.IntermediateCertWithSameSubjectAndSKID2SerialNumber, intermediateCerts.Certs[0].SerialNumber)

	// remove  intermediate certificate by serial number and check that leaf cert is not removed
	removeX509Cert = types.NewMsgRemoveX509Cert(
		setup.Trustee1.String(),
		testconstants.IntermediateCertWithSameSubjectAndSKIDSubject,
		testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID,
		testconstants.IntermediateCertWithSameSubjectAndSKID2SerialNumber,
	)
	_, err = setup.Handler(setup.Ctx, removeX509Cert)
	require.NoError(t, err)

	allCerts, _ = queryAllApprovedCertificates(setup)
	require.Equal(t, 2, len(allCerts))
	require.Equal(t, 2, len(allCerts[0].Certs)+len(allCerts[1].Certs))

	_, err = queryApprovedCertificates(setup, testconstants.IntermediateCertWithSameSubjectAndSKIDSubject, testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificates does not exists
	found := setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.RootCertWithSameSubjectAndSKIDSubject, testconstants.IntermediateCertWithSameSubjectAndSKID1SerialNumber)
	require.Equal(t, false, found)
	found = setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.RootCertWithSameSubjectAndSKIDSubject, testconstants.IntermediateCertWithSameSubjectAndSKID2SerialNumber)
	require.Equal(t, false, found)

	leafCerts, _ = queryApprovedCertificates(setup, testconstants.LeafCertWithSameSubjectAndSKIDSubject, testconstants.LeafCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(t, 1, len(leafCerts.Certs))
}

func TestHandler_RemoveX509Cert_RevokedAndApprovedCertificate(t *testing.T) {
	setup := Setup(t)
	// propose and approve x509 root certificate
	rootCertOptions := &rootCertOptions{
		pemCert:      testconstants.RootCertWithSameSubjectAndSKID1,
		subject:      testconstants.RootCertWithSameSubjectAndSKIDSubject,
		subjectKeyID: testconstants.RootCertWithSameSubjectAndSKIDSubjectKeyID,
		info:         testconstants.Info,
		vid:          65521,
	}
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// Add an intermediate certificate
	addIntermediateX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateWithSameSubjectAndSKID1)
	_, err := setup.Handler(setup.Ctx, addIntermediateX509Cert)
	require.NoError(t, err)

	// get certificates for further comparison
	allCerts := setup.Keeper.GetAllApprovedCertificates(setup.Ctx)
	require.NotNil(t, allCerts)
	require.Equal(t, 2, len(allCerts))
	require.Equal(t, 2, len(allCerts[0].Certs)+len(allCerts[1].Certs))

	// revoke an intermediate certificate
	revokeX509Cert := types.NewMsgRemoveX509Cert(
		setup.Trustee1.String(),
		testconstants.IntermediateCertWithSameSubjectAndSKIDSubject,
		testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID,
		testconstants.IntermediateCertWithSameSubjectAndSKID1SerialNumber,
	)
	_, err = setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(t, err)

	// Add an intermediate certificate with new serial number
	addIntermediateX509Cert = types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateWithSameSubjectAndSKID2)
	_, err = setup.Handler(setup.Ctx, addIntermediateX509Cert)
	require.NoError(t, err)

	intermediateCerts, _ := queryApprovedCertificates(setup, testconstants.IntermediateCertWithSameSubjectAndSKIDSubject, testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(t, 1, len(intermediateCerts.Certs))
	require.Equal(t, testconstants.IntermediateCertWithSameSubjectAndSKIDSubject, intermediateCerts.Certs[0].Subject)
	require.Equal(t, testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID, intermediateCerts.Certs[0].SubjectKeyId)
	require.Equal(t, testconstants.IntermediateCertWithSameSubjectAndSKID2SerialNumber, intermediateCerts.Certs[0].SerialNumber)

	// remove an intermediate certificate
	removeX509Cert := types.NewMsgRemoveX509Cert(
		setup.Trustee1.String(),
		testconstants.IntermediateCertWithSameSubjectAndSKIDSubject,
		testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID,
		testconstants.IntermediateCertWithSameSubjectAndSKID2SerialNumber,
	)
	_, err = setup.Handler(setup.Ctx, removeX509Cert)
	require.NoError(t, err)

	// check that only root and leaf certificates exists
	allCerts, _ = queryAllApprovedCertificates(setup)
	require.Equal(t, 1, len(allCerts))
	require.Equal(t, true, allCerts[0].Certs[0].IsRoot)
	_, err = queryApprovedCertificates(setup, testconstants.IntermediateCertWithSameSubjectAndSKIDSubject, testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))
	_, err = queryRevokedCertificates(setup, testconstants.IntermediateCertWithSameSubjectAndSKIDSubject, testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificates does not exists
	found := setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.RootCertWithSameSubjectAndSKIDSubject, testconstants.IntermediateCertWithSameSubjectAndSKID1SerialNumber)
	require.Equal(t, false, found)
	found = setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.RootCertWithSameSubjectAndSKIDSubject, testconstants.IntermediateCertWithSameSubjectAndSKID2SerialNumber)
	require.Equal(t, false, found)
}

func TestHandler_RemoveX509Cert_RevokedCertificate(t *testing.T) {
	setup := Setup(t)
	// propose and approve x509 root certificate
	rootCertOptions := &rootCertOptions{
		pemCert:      testconstants.RootCertPem,
		subject:      testconstants.RootSubject,
		subjectKeyID: testconstants.RootSubjectKeyID,
		info:         testconstants.Info,
		vid:          65521,
	}
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// Add two intermediate certificates again
	addIntermediateX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addIntermediateX509Cert)
	require.NoError(t, err)

	intermediateCerts, _ := queryApprovedCertificates(setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Equal(t, 1, len(intermediateCerts.Certs))
	require.Equal(t, testconstants.IntermediateSubject, intermediateCerts.Certs[0].Subject)
	require.Equal(t, testconstants.IntermediateSubjectKeyID, intermediateCerts.Certs[0].SubjectKeyId)

	// revoke intermediate certificate by serial number
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		setup.Trustee1.String(),
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		testconstants.IntermediateSerialNumber,
		testconstants.Info,
	)
	_, err = setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(t, err)

	_, err = queryApprovedCertificates(setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))

	revokedCerts, _ := queryRevokedCertificates(setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Equal(t, 1, len(revokedCerts.Certs))
	require.Equal(t, testconstants.IntermediateSubject, revokedCerts.Certs[0].Subject)
	require.Equal(t, testconstants.IntermediateSubjectKeyID, revokedCerts.Certs[0].SubjectKeyId)

	// remove  intermediate certificate by serial number
	removeX509Cert := types.NewMsgRemoveX509Cert(
		setup.Trustee1.String(),
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		testconstants.IntermediateSerialNumber,
	)
	_, err = setup.Handler(setup.Ctx, removeX509Cert)
	require.NoError(t, err)

	allCerts, _ := queryAllApprovedCertificates(setup)
	require.Equal(t, 1, len(allCerts))
	require.Equal(t, true, allCerts[0].Certs[0].IsRoot)

	_, err = queryApprovedCertificates(setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))
	_, err = queryRevokedCertificates(setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificate does not exists
	found := setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.IntermediateIssuer, testconstants.IntermediateSerialNumber)
	require.Equal(t, false, found)
}

func TestHandler_RemoveX509Cert_CertificateDoesNotExist(t *testing.T) {
	setup := Setup(t)

	removeX509Cert := types.NewMsgRemoveX509Cert(
		setup.Trustee1.String(), testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID, testconstants.IntermediateSerialNumber)
	_, err := setup.Handler(setup.Ctx, removeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_RemoveX509Cert_EmptyCertificatesList(t *testing.T) {
	setup := Setup(t)

	rootCertificate := rootCertificate(setup.Trustee1)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	setup.Keeper.SetApprovedCertificates(
		setup.Ctx,
		types.ApprovedCertificates{
			Subject:      testconstants.IntermediateSubject,
			SubjectKeyId: testconstants.IntermediateSubjectKeyID,
		},
	)

	removeX509Cert := types.NewMsgRemoveX509Cert(
		setup.Trustee1.String(), testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID, "")
	_, err := setup.Handler(setup.Ctx, removeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_RemoveX509Cert_ByNotOwner(t *testing.T) {
	setup := Setup(t)

	rootCertificate := rootCertificate(setup.Trustee1)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	anotherTrustee := GenerateAccAddress()
	setup.AddAccount(anotherTrustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)

	removeX509Cert := types.NewMsgRemoveX509Cert(
		anotherTrustee.String(), testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID, "")
	_, err = setup.Handler(setup.Ctx, removeX509Cert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_RemoveX509Cert_ByNotOwnerButSameVendor(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee1)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// add first vendor account with VID = 1
	vendorAccAddress1 := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add x509 certificate by fist vendor account
	addX509Cert := types.NewMsgAddX509Cert(vendorAccAddress1.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// add second vendor account with VID = 1
	vendorAccAddress2 := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

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
	_, err = queryApprovedCertificates(setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that certificate removed from 'approved certificates by subject' list
	_, err = queryApprovedCertificatesBySubject(setup, testconstants.IntermediateSubject)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that certificate removed from 'approved certificates by SKID' list
	approvedCerts, err := queryAllApprovedCertificatesBySubjectKeyID(setup, testconstants.IntermediateSubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 0, len(approvedCerts))

	// check that unique certificate key is not registered
	require.False(t, setup.Keeper.IsUniqueCertificatePresent(setup.Ctx,
		testconstants.IntermediateIssuer, testconstants.IntermediateSerialNumber))
}

func TestHandler_RemoveX509Cert_ByNotOwnerAndOtherVendor(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee1)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// add fist vendor account with VID = 1
	vendorAccAddress1 := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add x509 certificate by `setup.Trustee`
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// add scond vendor account with VID = 1000
	vendorAccAddress2 := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID1)

	// revoke x509 certificate by second vendor account
	removeX509Cert := types.NewMsgRemoveX509Cert(
		vendorAccAddress2.String(), testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID, testconstants.IntermediateSerialNumber)
	_, err = setup.Handler(setup.Ctx, removeX509Cert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_RemoveX509Cert_ForRootCertificate(t *testing.T) {
	setup := Setup(t)

	rootCertOptions := createTestRootCertOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	removeX509Cert := types.NewMsgRemoveX509Cert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.RootSerialNumber)
	_, err := setup.Handler(setup.Ctx, removeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrInappropriateCertificateType.Is(err))
}

func TestHandler_RemoveX509Cert_InvalidSerialNumber(t *testing.T) {
	setup := Setup(t)

	rootCertOptions := createTestRootCertOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	removeX509Cert := types.NewMsgRemoveX509Cert(
		setup.Trustee1.String(), testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID, "invalid")
	_, err = setup.Handler(setup.Ctx, removeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_RemoveX509Cert_BySubjectAndSKID(t *testing.T) {
	setup := Setup(t)
	// propose and approve x509 root certificate
	rootCertOptions := &rootCertOptions{
		pemCert:      testconstants.RootCertWithSameSubjectAndSKID1,
		subject:      testconstants.RootCertWithSameSubjectAndSKIDSubject,
		subjectKeyID: testconstants.RootCertWithSameSubjectAndSKIDSubjectKeyID,
		info:         testconstants.Info,
		vid:          65521,
	}
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// Add two intermediate certificates
	addIntermediateX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateWithSameSubjectAndSKID1)
	_, err := setup.Handler(setup.Ctx, addIntermediateX509Cert)
	require.NoError(t, err)
	addIntermediateX509Cert = types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateWithSameSubjectAndSKID2)
	_, err = setup.Handler(setup.Ctx, addIntermediateX509Cert)
	require.NoError(t, err)

	// Add a leaf certificate
	addLeafX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.LeafCertWithSameSubjectAndSKID)
	_, err = setup.Handler(setup.Ctx, addLeafX509Cert)
	require.NoError(t, err)

	// get certificates for further comparison
	allCerts := setup.Keeper.GetAllApprovedCertificates(setup.Ctx)
	require.NotNil(t, allCerts)
	require.Equal(t, 3, len(allCerts))
	require.Equal(t, 4, len(allCerts[0].Certs)+len(allCerts[1].Certs)+len(allCerts[2].Certs))

	// remove all intermediate certificates but leave leaf certificate
	removeX509Cert := types.NewMsgRemoveX509Cert(
		setup.Trustee1.String(),
		testconstants.IntermediateCertWithSameSubjectAndSKIDSubject,
		testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID,
		"",
	)
	_, err = setup.Handler(setup.Ctx, removeX509Cert)
	require.NoError(t, err)

	// check that only root and leaf certificates exists
	allCerts, _ = queryAllApprovedCertificates(setup)
	require.Equal(t, 2, len(allCerts))
	require.Equal(t, 2, len(allCerts[0].Certs)+len(allCerts[1].Certs))
	_, err = queryApprovedCertificates(setup, testconstants.IntermediateCertWithSameSubjectAndSKIDSubject, testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))
	// check that unique certificates does not exists
	found := setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.RootCertWithSameSubjectAndSKIDSubject, testconstants.IntermediateCertWithSameSubjectAndSKID1SerialNumber)
	require.Equal(t, false, found)
	found = setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.RootCertWithSameSubjectAndSKIDSubject, testconstants.IntermediateCertWithSameSubjectAndSKID2SerialNumber)
	require.Equal(t, false, found)

	leafCerts, _ := queryApprovedCertificates(setup, testconstants.LeafCertWithSameSubjectAndSKIDSubject, testconstants.LeafCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(t, 1, len(leafCerts.Certs))
	require.Equal(t, testconstants.LeafCertWithSameSubjectAndSKIDSerialNumber, leafCerts.Certs[0].SerialNumber)
}
