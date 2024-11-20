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

func TestHandler_RevokeX509Cert(t *testing.T) {
	setup := Setup(t)

	// Add vendor account
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.RootCertWithVidVid)

	// propose and approve x509 root certificate
	rootCertOptions := &rootCertOptions{
		pemCert:      testconstants.RootCertPem,
		subject:      testconstants.RootSubject,
		subjectKeyID: testconstants.RootSubjectKeyID,
		info:         testconstants.Info,
		vid:          testconstants.RootCertWithVidVid,
	}
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// Add two intermediate certificates again
	addDaPaiCertificate(setup, vendorAccAddress, testconstants.IntermediateCertPem)

	// revoke x509 certificate
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		vendorAccAddress.String(),
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		"",
		false,
		testconstants.Info,
	)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(t, err)

	// check that intermediate certificate has been revoked
	allRevokedCertificates, _ := queryAllRevokedCertificates(setup)
	require.Equal(t, 1, len(allRevokedCertificates))
	require.Equal(t, testconstants.IntermediateSubject, allRevokedCertificates[0].Subject)
	require.Equal(t, testconstants.IntermediateSubjectKeyID, allRevokedCertificates[0].SubjectKeyId)
	require.Equal(t, 1, len(allRevokedCertificates[0].Certs))

	ensureDaPaiCertificateDoesNotExist(
		t,
		setup,
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		testconstants.IntermediateIssuer,
		testconstants.IntermediateSerialNumber,
		true,
		false)

	// check that root certificate stays approved
	ensureDaPaaCertificateExist(
		t,
		setup,
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.RootSubject,
		testconstants.RootSerialNumber)

	// check that no proposed certificate revocations have been created
	allProposedCertificateRevocations, _ := queryAllProposedCertificateRevocations(setup)
	require.NoError(t, err)
	require.Equal(t, 0, len(allProposedCertificateRevocations))

	// check that child certificate identifiers list of issuer do not exist anymore
	_, err = queryChildCertificates(setup, testconstants.IntermediateIssuer, testconstants.IntermediateAuthorityKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificate key stays registered
	require.True(t, setup.Keeper.IsUniqueCertificatePresent(setup.Ctx,
		testconstants.IntermediateIssuer, testconstants.IntermediateSerialNumber))
}

func TestHandler_RevokeX509Cert_ForTree(t *testing.T) {
	setup := Setup(t)

	// Add vendor account
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add root x509 certificate
	rootCertOptions := createTestRootCertOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// add intermediate x509 certificate
	addDaPaiCertificate(setup, vendorAccAddress, testconstants.IntermediateCertPem)

	// add leaf x509 certificate
	addDaPaiCertificate(setup, vendorAccAddress, testconstants.LeafCertPem)

	// revoke x509 certificate
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		vendorAccAddress.String(),
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		"",
		true,
		testconstants.Info,
	)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(t, err)

	// check that intermediate certificate has been revoked
	allRevokedCertificates, _ := queryAllRevokedCertificates(setup)
	require.Equal(t, 2, len(allRevokedCertificates))
	require.Equal(t, testconstants.LeafSubject, allRevokedCertificates[0].Subject)
	require.Equal(t, testconstants.LeafSubjectKeyID, allRevokedCertificates[0].SubjectKeyId)
	require.Equal(t, 1, len(allRevokedCertificates[0].Certs))
	require.Equal(t, testconstants.LeafCertPem, allRevokedCertificates[0].Certs[0].PemCert)
	require.Equal(t, testconstants.IntermediateSubject, allRevokedCertificates[1].Subject)
	require.Equal(t, testconstants.IntermediateSubjectKeyID, allRevokedCertificates[1].SubjectKeyId)
	require.Equal(t, 1, len(allRevokedCertificates[1].Certs))
	require.Equal(t, testconstants.IntermediateCertPem, allRevokedCertificates[1].Certs[0].PemCert)

	// check that root certificate stays approved
	ensureDaPaaCertificateExist(
		t,
		setup,
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.RootSubject,
		testconstants.RootSerialNumber)

	// check that no proposed certificate revocations have been created
	allProposedCertificateRevocations, _ := queryAllProposedCertificateRevocations(setup)
	require.NoError(t, err)
	require.Equal(t, 0, len(allProposedCertificateRevocations))

	// check that no child certificate identifiers are now registered for root certificate
	_, err = queryChildCertificates(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that no child certificate identifiers are registered for revoked intermediate certificate
	_, err = queryChildCertificates(setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that no child certificate identifiers are registered for revoked leaf certificate
	_, err = queryChildCertificates(setup, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
}

func TestHandler_RevokeX509Cert_BySerialNumber(t *testing.T) {
	setup := Setup(t)

	// Add vendor account
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// propose and approve x509 root certificate
	rootCertOptions := createTestRootCertOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// Add two intermediate certificates
	addDaPaiCertificate(setup, vendorAccAddress, testconstants.IntermediateCertPem)

	intermediateCertificate := intermediateCertificateNoVid(vendorAccAddress)
	intermediateCertificate.SerialNumber = SerialNumber
	setup.Keeper.AddAllCertificate(setup.Ctx, intermediateCertificate)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, intermediateCertificate)
	setup.Keeper.AddApprovedCertificateBySubjectKeyID(setup.Ctx, intermediateCertificate)
	setup.Keeper.SetUniqueCertificate(
		setup.Ctx,
		uniqueCertificate(intermediateCertificate.Issuer, intermediateCertificate.SerialNumber),
	)

	// Add a leaf certificate
	addDaPaiCertificate(setup, vendorAccAddress, testconstants.LeafCertPem)

	// get certificates for further comparison
	allCerts := setup.Keeper.GetAllApprovedCertificates(setup.Ctx)
	require.NotNil(t, allCerts)
	require.Equal(t, 3, len(allCerts))
	require.Equal(t, 4, len(allCerts[0].Certs)+len(allCerts[1].Certs)+len(allCerts[2].Certs))

	// revoke only an intermediate certificate
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		vendorAccAddress.String(),
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		testconstants.IntermediateSerialNumber,
		false,
		testconstants.Info,
	)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(t, err)

	// check that proposed certificate revocation does not exist anymore
	_, err = queryProposedCertificateRevocation(setup, testconstants.IntermediateSerialNumber)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that only root, intermediate and leaf certificates exists
	allCerts, _ = queryAllApprovedCertificates(setup)
	require.Equal(t, 3, len(allCerts))
	require.Equal(t, 3, len(allCerts[0].Certs)+len(allCerts[1].Certs)+len(allCerts[2].Certs))

	intermediateCerts, _ := queryApprovedCertificates(setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Equal(t, 1, len(intermediateCerts.Certs))
	require.Equal(t, SerialNumber, intermediateCerts.Certs[0].SerialNumber)

	leafCerts, _ := queryApprovedCertificates(setup, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Equal(t, 1, len(leafCerts.Certs))
	require.Equal(t, testconstants.LeafSerialNumber, leafCerts.Certs[0].SerialNumber)

	// query and check revoked certificate
	revokedCertificate, _ := querySingleRevokedCertificate(setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.NotNil(t, revokedCertificate)
	require.Equal(t, testconstants.IntermediateSubject, revokedCertificate.Subject)
	require.Equal(t, testconstants.IntermediateSubjectKeyID, revokedCertificate.SubjectKeyId)
	require.Equal(t, testconstants.IntermediateSerialNumber, revokedCertificate.SerialNumber)

	// revoke intermediate and leaf certificates
	revokeX509Cert = types.NewMsgRevokeX509Cert(
		vendorAccAddress.String(),
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		SerialNumber,
		true,
		testconstants.Info,
	)
	_, err = setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(t, err)

	_, err = queryProposedCertificateRevocation(setup, testconstants.IntermediateSerialNumber)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that only root certificate exists
	certsAfterRevocation := setup.Keeper.GetAllApprovedCertificates(setup.Ctx)
	require.Equal(t, 1, len(certsAfterRevocation))
	require.Equal(t, 1, len(certsAfterRevocation[0].Certs))
	require.Equal(t, testconstants.RootSerialNumber, certsAfterRevocation[0].Certs[0].SerialNumber)

	// query and check revoked certificate
	revokedCerts, _ := queryRevokedCertificates(setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Equal(t, 2, len(revokedCerts.Certs))
	require.Equal(t, testconstants.IntermediateSubject, revokedCerts.Subject)
	require.Equal(t, testconstants.IntermediateSubjectKeyID, revokedCerts.SubjectKeyId)

	// query and check revoked certificate
	revokedCerts, _ = queryRevokedCertificates(setup, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Equal(t, 1, len(revokedCerts.Certs))
	require.Equal(t, testconstants.LeafSubject, revokedCerts.Subject)
	require.Equal(t, testconstants.LeafSubjectKeyID, revokedCerts.SubjectKeyId)
}

// Extra cases

func TestHandler_RevokeX509Cert_ByNotOwnerButSameVendor(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee1)
	setup.Keeper.AddAllCertificate(setup.Ctx, rootCertificate)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// add first vendor account with VID = 1
	vendorAccAddress1 := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add x509 certificate by first vendor account
	addX509Cert := types.NewMsgAddX509Cert(vendorAccAddress1.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// add second vendor account with VID = 1
	vendorAccAddress2 := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// revoke x509 certificate by second vendor account
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		vendorAccAddress2.String(),
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		testconstants.IntermediateSerialNumber,
		false,
		testconstants.Info,
	)
	_, err = setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(t, err)

	// check that intermediate certificate has been added to revoked list
	revokedCertificates, _ := queryRevokedCertificates(setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Equal(t, testconstants.IntermediateSubject, revokedCertificates.Subject)
	require.Equal(t, testconstants.IntermediateSubjectKeyID, revokedCertificates.SubjectKeyId)
	require.Equal(t, 1, len(revokedCertificates.Certs))
	require.Equal(t, intermediateCertificateNoVid(vendorAccAddress1), *revokedCertificates.Certs[0])

	// check that revoked certificate removed from approved certificates list
	_, err = queryApprovedCertificates(setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that revoked certificate removed from 'approved certificates' by subject list
	_, err = queryApprovedCertificatesBySubject(setup, testconstants.IntermediateSubject)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that revoked certificate removed from 'approved certificates' by SKID list
	approvedCerts, err := queryAllApprovedCertificatesBySubjectKeyID(setup, testconstants.IntermediateSubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 0, len(approvedCerts))

	// check that unique certificate key stays registered
	require.True(t, setup.Keeper.IsUniqueCertificatePresent(setup.Ctx,
		testconstants.IntermediateIssuer, testconstants.IntermediateSerialNumber))
}

// Error cases

func TestHandler_RevokeX509Cert_CertificateDoesNotExist(t *testing.T) {
	setup := Setup(t)

	// Add vendor account
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// revoke x509 certificate
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		vendorAccAddress.String(),
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

func TestHandler_RevokeX509Cert_CertificateDoesNotExistBySerialNumber(t *testing.T) {
	setup := Setup(t)
	// propose and approve x509 root certificate
	rootCertOptions := createTestRootCertOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// Add vendor account
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// Add intermediate certificate
	addIntermediateX509Cert := types.NewMsgAddX509Cert(vendorAccAddress.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addIntermediateX509Cert)
	require.NoError(t, err)

	// revoke x509 certificate
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		vendorAccAddress.String(),
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		"invalid",
		false,
		testconstants.Info,
	)
	_, err = setup.Handler(setup.Ctx, revokeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_RevokeX509Cert_ForRootCertificate(t *testing.T) {
	setup := Setup(t)

	// propose and approve x509 root certificate
	rootCertOptions := createTestRootCertOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// Add vendor account
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// revoke x509 root certificate
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		vendorAccAddress.String(),
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.RootSerialNumber,
		false,
		testconstants.Info,
	)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrInappropriateCertificateType.Is(err))
}

func TestHandler_RevokeX509Cert_ByOtherVendor(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee1)
	setup.Keeper.AddAllCertificate(setup.Ctx, rootCertificate)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// add first vendor account with VID = 1
	vendorAccAddress1 := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add x509 certificate by first vendor account
	addX509Cert := types.NewMsgAddX509Cert(vendorAccAddress1.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// add second vendor account with VID = 1000
	vendorAccAddress2 := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID1)

	// revoke x509 certificate by second vendor account
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		vendorAccAddress2.String(),
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		testconstants.IntermediateSerialNumber,
		false,
		testconstants.Info,
	)
	_, err = setup.Handler(setup.Ctx, revokeX509Cert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_RevokeX509Cert_SenderNotVendor(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCertOptions := createRootWithVidOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// Add vendor account
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.RootCertWithVidVid)

	// add x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(vendorAccAddress.String(), testconstants.IntermediateCertWithVid1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	removeX509Cert := types.NewMsgRevokeX509Cert(
		setup.Trustee1.String(),
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		testconstants.IntermediateSerialNumber,
		false,
		testconstants.Info,
	)
	_, err = setup.Handler(setup.Ctx, removeX509Cert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}
