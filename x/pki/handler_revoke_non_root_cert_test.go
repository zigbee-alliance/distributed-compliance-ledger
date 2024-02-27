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

func TestHandler_RevokeX509CertRoot_CertificateDoesNotExistBySerialNumber(t *testing.T) {
	setup := Setup(t)
	// propose and approve x509 root certificate
	rootCertOptions := createTestRootCertOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// revoke x509 certificate
	revokeX509Cert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, "invalid", testconstants.Info)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}
func TestHandler_RevokeX509Cert_CertificateDoesNotExist(t *testing.T) {
	setup := Setup(t)

	// revoke x509 certificate
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		setup.Trustee1.String(), testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID, testconstants.IntermediateSerialNumber, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_RevokeX509Cert_CertificateDoesNotExistBySerialNumber(t *testing.T) {
	setup := Setup(t)
	// propose and approve x509 root certificate
	rootCertOptions := createTestRootCertOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// Add intermediate certificate
	addIntermediateX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addIntermediateX509Cert)
	require.NoError(t, err)

	// revoke x509 certificate
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		setup.Trustee1.String(), testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID, "invalid", testconstants.Info)
	_, err = setup.Handler(setup.Ctx, revokeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_RevokeX509Cert_ForRootCertificate(t *testing.T) {
	setup := Setup(t)

	// propose and approve x509 root certificate
	rootCertOptions := createTestRootCertOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// revoke x509 root certificate
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.RootSerialNumber, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrInappropriateCertificateType.Is(err))
}

func TestHandler_RevokeX509Cert_ForTree(t *testing.T) {
	setup := Setup(t)

	// add root x509 certificate
	rootCertOptions := createTestRootCertOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// add intermediate x509 certificate
	addIntermediateX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addIntermediateX509Cert)
	require.NoError(t, err)

	// add leaf x509 certificate
	addLeafX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.LeafCertPem)
	_, err = setup.Handler(setup.Ctx, addLeafX509Cert)
	require.NoError(t, err)

	// check that intermediate nd leaf certificates removed from subject-key-id -> certs map
	certs, _ := queryAllApprovedCertificatesBySubjectKeyID(setup, testconstants.IntermediateSubjectKeyID)
	require.Equal(t, 1, len(certs))
	certs, _ = queryAllApprovedCertificatesBySubjectKeyID(setup, testconstants.LeafSubjectKeyID)
	require.Equal(t, 1, len(certs))

	// revoke x509 certificate
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		setup.Trustee1.String(), testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID, "", testconstants.Info)
	_, err = setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(t, err)

	// check that intermediate and leaf certificates have been revoked
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
	allApprovedCertificates, _ := queryAllApprovedCertificates(setup)
	require.Equal(t, 1, len(allApprovedCertificates))
	require.Equal(t, testconstants.RootSubject, allApprovedCertificates[0].Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, allApprovedCertificates[0].SubjectKeyId)
	// check that root certificate(by subject key id) stays approved
	allApprovedCertificates, _ = queryAllApprovedCertificatesBySubjectKeyID(setup, testconstants.RootSubjectKeyID)
	require.Equal(t, 1, len(allApprovedCertificates))
	require.Equal(t, testconstants.RootSubjectKeyID, allApprovedCertificates[0].SubjectKeyId)
	// check that intermediate and leaf certificates removed from subject-key-id -> certs map
	allApprovedCertificates, _ = queryAllApprovedCertificatesBySubjectKeyID(setup, testconstants.IntermediateSubjectKeyID)
	require.Equal(t, 0, len(allApprovedCertificates))
	allApprovedCertificates, _ = queryAllApprovedCertificatesBySubjectKeyID(setup, testconstants.LeafSubjectKeyID)
	require.Equal(t, 0, len(allApprovedCertificates))

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

func TestHandler_RevokeX509Cert_ByNotOwner(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee1)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// add x509 certificate by `setup.Trustee`
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// store another account
	anotherTrustee := GenerateAccAddress()
	setup.AddAccount(anotherTrustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)

	// revoke x509 certificate by another account
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		anotherTrustee.String(), testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID, testconstants.IntermediateSerialNumber, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, revokeX509Cert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_RevokeX509Cert_ByNotOwnerButSameVendor(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee1)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// add first vendor's account
	vendorAccAddress1 := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add x509 certificate by vendor`
	addX509Cert := types.NewMsgAddX509Cert(vendorAccAddress1.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// add second vendor's account
	vendorAccAddress2 := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// revoke x509 certificate by another account
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		vendorAccAddress2.String(),
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		testconstants.IntermediateSerialNumber,
		testconstants.Info,
	)
	_, err = setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(t, err)

	// check that intermediate cer

	// check that intermediate certificate has been added to revoked list
	revokedCertificates, _ := queryRevokedCertificates(setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Equal(t, testconstants.IntermediateSubject, revokedCertificates.Subject)
	require.Equal(t, testconstants.IntermediateSubjectKeyID, revokedCertificates.SubjectKeyId)
	require.Equal(t, 1, len(revokedCertificates.Certs))
	require.Equal(t, intermediateCertificate(vendorAccAddress1), *revokedCertificates.Certs[0])

	// check that revoked certificate remover from approved certificates list
	_, err = queryApprovedCertificates(setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that revoked certificate removed from approved certificates by subject list
	_, err = queryApprovedCertificatesBySubject(setup, testconstants.IntermediateSubject)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that revoked certificate removed from approved certificates by SKID list
	approvedCerts, err := queryAllApprovedCertificatesBySubjectKeyID(setup, testconstants.IntermediateSubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 0, len(approvedCerts))

	// check that child certificate identifiers list of issuer do not exist anymore
	_, err = queryChildCertificates(setup, testconstants.IntermediateIssuer, testconstants.IntermediateAuthorityKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificate key stays registered
	require.True(t, setup.Keeper.IsUniqueCertificatePresent(setup.Ctx,
		testconstants.IntermediateIssuer, testconstants.IntermediateSerialNumber))
}

func TestHandler_RevokeX509Cert_ByNotOwnerAndOtherVendor(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee1)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	vendorAccAddress1 := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add x509 certificate by `setup.Trustee`
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// store another account
	anotherTrustee := GenerateAccAddress()
	setup.AddAccount(anotherTrustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)

	vendorAccAddress2 := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// revoke x509 certificate by another account
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		anotherTrustee.String(), testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID, testconstants.IntermediateSerialNumber, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, revokeX509Cert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_RevokeX509Cert(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee1)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.Vendor,
		dclauthtypes.CertificationCenter,
		dclauthtypes.Trustee,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, 1)

		// add x509 certificate
		addX509Cert := types.NewMsgAddX509Cert(accAddress.String(), testconstants.IntermediateCertPem)
		_, err := setup.Handler(setup.Ctx, addX509Cert)
		require.NoError(t, err)

		// get intermediate certificate by subject-key-id
		certsBySubjectKeyID, _ := queryAllApprovedCertificatesBySubjectKeyID(setup, testconstants.IntermediateSubjectKeyID)
		require.Equal(t, 1, len(certsBySubjectKeyID))
		// get certificate for further comparison
		certificateBeforeRevocation, _ := querySingleApprovedCertificate(
			setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
		require.NotNil(t, certificateBeforeRevocation)

		// revoke x509 certificate
		revokeX509Cert := types.NewMsgRevokeX509Cert(
			accAddress.String(), testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID, "", testconstants.Info)
		_, err = setup.Handler(setup.Ctx, revokeX509Cert)
		require.NoError(t, err)

		// check that intermediate certificate has been revoked
		allRevokedCertificates, _ := queryAllRevokedCertificates(setup)
		require.Equal(t, 1, len(allRevokedCertificates))
		require.Equal(t, testconstants.IntermediateSubject, allRevokedCertificates[0].Subject)
		require.Equal(t, testconstants.IntermediateSubjectKeyID, allRevokedCertificates[0].SubjectKeyId)
		require.Equal(t, 1, len(allRevokedCertificates[0].Certs))
		require.Equal(t, *certificateBeforeRevocation, *allRevokedCertificates[0].Certs[0])

		// check that root certificate stays approved
		allApprovedCertificates, _ := queryAllApprovedCertificates(setup)
		require.Equal(t, 1, len(allApprovedCertificates))
		require.Equal(t, testconstants.IntermediateSubject, allRevokedCertificates[0].Subject)
		require.Equal(t, testconstants.IntermediateSubjectKeyID, allRevokedCertificates[0].SubjectKeyId)

		// check that intermediate certificate removed from subject-key-id -> certs map
		certsBySubjectKeyID, _ = queryAllApprovedCertificatesBySubjectKeyID(setup, testconstants.IntermediateSubjectKeyID)
		require.Equal(t, 0, len(certsBySubjectKeyID))

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

		// cleanup for next iteration
		setup.Keeper.RemoveRevokedCertificates(setup.Ctx,
			testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
		setup.Keeper.RemoveUniqueCertificate(setup.Ctx,
			testconstants.IntermediateIssuer, testconstants.IntermediateSerialNumber)
	}
}

func TestHandler_RevokeX509Cert_BySerialNumber(t *testing.T) {
	setup := Setup(t)
	// propose and approve x509 root certificate
	rootCertOptions := createTestRootCertOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// Add two intermediate certificates
	addIntermediateX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addIntermediateX509Cert)
	require.NoError(t, err)
	intermediateCertificate := intermediateCertificate(setup.Trustee1)
	intermediateCertificate.SerialNumber = SerialNumber
	setup.Keeper.AddApprovedCertificate(setup.Ctx, intermediateCertificate)
	setup.Keeper.AddApprovedCertificateBySubjectKeyID(setup.Ctx, intermediateCertificate)
	setup.Keeper.SetUniqueCertificate(
		setup.Ctx,
		uniqueCertificate(intermediateCertificate.Issuer, intermediateCertificate.SerialNumber),
	)
	// Add a leaf certificate
	addLeafX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.LeafCertPem)
	_, err = setup.Handler(setup.Ctx, addLeafX509Cert)
	require.NoError(t, err)

	// get certificates for further comparison
	allCerts := setup.Keeper.GetAllApprovedCertificates(setup.Ctx)
	require.NotNil(t, allCerts)
	require.Equal(t, 3, len(allCerts))
	require.Equal(t, 4, len(allCerts[0].Certs)+len(allCerts[1].Certs)+len(allCerts[2].Certs))

	// revoke intermediate certificate
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		setup.Trustee1.String(), testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID, testconstants.IntermediateSerialNumber, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(t, err)

	// check that proposed certificate revocation does not exist anymore
	_, err = queryProposedCertificateRevocation(setup, testconstants.IntermediateSerialNumber)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that only root and two intermediate certificates exists
	allCerts, _ = queryAllApprovedCertificates(setup)
	require.Equal(t, 3, len(allCerts))
	require.Equal(t, 3, len(allCerts[0].Certs)+len(allCerts[1].Certs)+len(allCerts[2].Certs))
	intermediateCerts, _ := queryApprovedCertificates(setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Equal(t, 1, len(intermediateCerts.Certs))
	require.Equal(t, SerialNumber, intermediateCerts.Certs[0].SerialNumber)

	// query and check revoked certificate
	revokedCertificate, _ := querySingleRevokedCertificate(setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.NotNil(t, revokedCertificate)
	require.Equal(t, testconstants.IntermediateSubject, revokedCertificate.Subject)
	require.Equal(t, testconstants.IntermediateSubjectKeyID, revokedCertificate.SubjectKeyId)
	require.Equal(t, testconstants.IntermediateSerialNumber, revokedCertificate.SerialNumber)

	// revoke intermediate certificate
	revokeX509Cert = types.NewMsgRevokeX509Cert(
		setup.Trustee1.String(), testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID, SerialNumber, testconstants.Info)
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
}
