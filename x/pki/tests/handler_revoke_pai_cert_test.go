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

func TestHandler_RevokeDaIntermediateCert(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.RootCertWithVidVid)

	// propose and approve x509 root certificate
	rootCertificate := utils.CreateTestRootCert()
	rootCertOptions := &utils.RootCertOptions{
		PemCert:      testconstants.RootCertPem,
		Subject:      testconstants.RootSubject,
		SubjectKeyID: testconstants.RootSubjectKeyID,
		Info:         testconstants.Info,
		Vid:          testconstants.RootCertWithVidVid,
	}
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// Add intermediate certificate
	intermediateCertificate := utils.CreateTestIntermediateCert()
	utils.AddDaIntermediateCertificate(setup, vendorAccAddress, testconstants.IntermediateCertPem)

	// revoke intermediate certificate
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		vendorAccAddress.String(),
		intermediateCertificate.Subject,
		intermediateCertificate.SubjectKeyID,
		"",
		false,
		testconstants.Info,
	)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(t, err)

	// Check: Certificate is revoked
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateRevocationKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, intermediateCertificate, indexes)

	// Check: Root stays approved
	indexes = utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)
}

func TestHandler_RevokeX509Cert_ForTree(t *testing.T) {
	setup := utils.Setup(t)

	// add root x509 certificate
	rootCertificate := utils.CreateTestRootCert()
	rootCertOptions := utils.CreateTestRootCertOptions()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// add intermediate x509 certificate
	intermediateCertificate := utils.CreateTestIntermediateCert()
	utils.AddDaIntermediateCertificate(setup, setup.Vendor1, testconstants.IntermediateCertPem)

	// add leaf x509 certificate
	leafCertificate := utils.CreateTestLeafCert()
	utils.AddDaIntermediateCertificate(setup, setup.Vendor1, testconstants.LeafCertPem)

	// revoke x509 certificate
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		setup.Vendor1.String(),
		intermediateCertificate.Subject,
		intermediateCertificate.SubjectKeyID,
		"",
		true,
		testconstants.Info,
	)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(t, err)

	// root stays approved
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)

	// intermediate and leaf are revoked
	indexes = utils.TestIndexes{
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
			{Key: types.ApprovedRootCertificatesKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, intermediateCertificate, indexes)
	utils.CheckCertificateStateIndexes(t, setup, leafCertificate, indexes)
}

func TestHandler_RevokeX509Cert_BySerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	// propose and approve x509 root certificate
	rootCertOptions := utils.CreateTestRootCertOptions()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// Add two intermediate certificates
	utils.AddDaIntermediateCertificate(setup, setup.Vendor1, testconstants.IntermediateCertPem)

	intermediateCertificate := utils.IntermediateCertificateNoVid(setup.Vendor1)
	intermediateCertificate.SerialNumber = utils.SerialNumber
	setup.Keeper.AddAllCertificate(setup.Ctx, intermediateCertificate)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, intermediateCertificate)
	setup.Keeper.AddApprovedCertificateBySubjectKeyID(setup.Ctx, intermediateCertificate)
	setup.Keeper.SetUniqueCertificate(
		setup.Ctx,
		utils.UniqueCertificate(intermediateCertificate.Issuer, intermediateCertificate.SerialNumber),
	)

	// Add a leaf certificate
	utils.AddDaIntermediateCertificate(setup, setup.Vendor1, testconstants.LeafCertPem)

	// get certificates for further comparison
	allCerts := setup.Keeper.GetAllApprovedCertificates(setup.Ctx)
	require.NotNil(t, allCerts)
	require.Equal(t, 3, len(allCerts))
	require.Equal(t, 4, len(allCerts[0].Certs)+len(allCerts[1].Certs)+len(allCerts[2].Certs))

	// revoke only an intermediate certificate
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		setup.Vendor1.String(),
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		testconstants.IntermediateSerialNumber,
		false,
		testconstants.Info,
	)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(t, err)

	// check that proposed certificate revocation does not exist anymore
	_, err = utils.QueryProposedCertificateRevocation(
		setup,
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		testconstants.IntermediateSerialNumber,
	)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that only root, intermediate and leaf certificates exists
	allCerts, _ = utils.QueryAllApprovedCertificates(setup)
	require.Equal(t, 3, len(allCerts))
	require.Equal(t, 3, len(allCerts[0].Certs)+len(allCerts[1].Certs)+len(allCerts[2].Certs))

	intermediateCerts, _ := utils.QueryApprovedCertificates(setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Equal(t, 1, len(intermediateCerts.Certs))
	require.Equal(t, utils.SerialNumber, intermediateCerts.Certs[0].SerialNumber)

	leafCerts, _ := utils.QueryApprovedCertificates(setup, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Equal(t, 1, len(leafCerts.Certs))
	require.Equal(t, testconstants.LeafSerialNumber, leafCerts.Certs[0].SerialNumber)

	// query and check revoked certificate
	revokedCertificate, _ := utils.QueryRevokedCertificates(setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.NotNil(t, revokedCertificate)
	require.Equal(t, testconstants.IntermediateSubject, revokedCertificate.Certs[0].Subject)
	require.Equal(t, testconstants.IntermediateSubjectKeyID, revokedCertificate.Certs[0].SubjectKeyId)
	require.Equal(t, testconstants.IntermediateSerialNumber, revokedCertificate.Certs[0].SerialNumber)

	// revoke intermediate and leaf certificates
	revokeX509Cert = types.NewMsgRevokeX509Cert(
		setup.Vendor1.String(),
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		utils.SerialNumber,
		true,
		testconstants.Info,
	)
	_, err = setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(t, err)

	_, err = utils.QueryProposedCertificateRevocation(
		setup,
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		testconstants.IntermediateSerialNumber,
	)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that only root certificate exists
	certsAfterRevocation := setup.Keeper.GetAllApprovedCertificates(setup.Ctx)
	require.Equal(t, 1, len(certsAfterRevocation))
	require.Equal(t, 1, len(certsAfterRevocation[0].Certs))
	require.Equal(t, testconstants.RootSerialNumber, certsAfterRevocation[0].Certs[0].SerialNumber)

	// query and check revoked certificate
	revokedCerts, _ := utils.QueryRevokedCertificates(setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Equal(t, 2, len(revokedCerts.Certs))
	require.Equal(t, testconstants.IntermediateSubject, revokedCerts.Subject)
	require.Equal(t, testconstants.IntermediateSubjectKeyID, revokedCerts.SubjectKeyId)

	// query and check revoked certificate
	revokedCerts, _ = utils.QueryRevokedCertificates(setup, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Equal(t, 1, len(revokedCerts.Certs))
	require.Equal(t, testconstants.LeafSubject, revokedCerts.Subject)
	require.Equal(t, testconstants.LeafSubjectKeyID, revokedCerts.SubjectKeyId)
}

// Extra cases

func TestHandler_RevokeX509Cert_ByNotOwnerButSameVendor(t *testing.T) {
	setup := utils.Setup(t)

	// store root certificate
	rootCertificate := utils.RootCertificate(setup.Trustee1)
	setup.Keeper.AddAllCertificate(setup.Ctx, rootCertificate)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// add first vendor account with VID = 1
	vendorAccAddress1 := setup.CreateVendorAccount(testconstants.Vid)

	// add x509 certificate by first vendor account
	addX509Cert := types.NewMsgAddX509Cert(vendorAccAddress1.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// add second vendor account with VID = 1
	vendorAccAddress2 := setup.CreateVendorAccount(testconstants.Vid)

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
	revokedCertificates, _ := utils.QueryRevokedCertificates(setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Equal(t, testconstants.IntermediateSubject, revokedCertificates.Subject)
	require.Equal(t, testconstants.IntermediateSubjectKeyID, revokedCertificates.SubjectKeyId)
	require.Equal(t, 1, len(revokedCertificates.Certs))
	require.Equal(t, utils.IntermediateCertificateNoVid(vendorAccAddress1), *revokedCertificates.Certs[0])

	// check that revoked certificate removed from approved certificates list
	_, err = utils.QueryApprovedCertificates(setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that revoked certificate removed from 'approved certificates' by subject list
	_, err = utils.QueryApprovedCertificatesBySubject(setup, testconstants.IntermediateSubject)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that revoked certificate removed from 'approved certificates' by SKID list
	approvedCerts, err := utils.QueryApprovedCertificatesBySubjectKeyID(setup, testconstants.IntermediateSubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 0, len(approvedCerts))

	// check that unique certificate key stays registered
	require.True(t, setup.Keeper.IsUniqueCertificatePresent(setup.Ctx,
		testconstants.IntermediateIssuer, testconstants.IntermediateSerialNumber))
}

// Error cases

func TestHandler_RevokeX509Cert_CertificateDoesNotExist(t *testing.T) {
	setup := utils.Setup(t)

	// revoke x509 certificate
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		setup.Vendor1.String(),
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
	setup := utils.Setup(t)

	// propose and approve x509 root certificate
	rootCertOptions := utils.CreateTestRootCertOptions()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// Add intermediate certificate
	addIntermediateX509Cert := types.NewMsgAddX509Cert(
		setup.Vendor1.String(),
		testconstants.IntermediateCertPem,
		testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addIntermediateX509Cert)
	require.NoError(t, err)

	// revoke x509 certificate
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		setup.Vendor1.String(),
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
	setup := utils.Setup(t)

	// propose and approve x509 root certificate
	rootCertOptions := utils.CreateTestRootCertOptions()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// revoke x509 root certificate
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		setup.Vendor1.String(),
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
	setup := utils.Setup(t)

	// store root certificate
	rootCertificate := utils.RootCertificate(setup.Trustee1)
	setup.Keeper.AddAllCertificate(setup.Ctx, rootCertificate)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// add x509 certificate by first vendor account
	addX509Cert := types.NewMsgAddX509Cert(setup.Vendor1.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// add second vendor account with VID = 1000
	vendorAccAddress2 := utils.GenerateAccAddress()
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
