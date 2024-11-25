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

func TestHandler_AddDaIntermediateCert(t *testing.T) {
	setup := utils.Setup(t)

	accAddress := setup.CreateVendorAccount(testconstants.Vid)

	// add DA root certificate
	rootCertOptions := utils.CreateTestRootCertOptions()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// add DA PAI certificate
	addX509Cert := types.NewMsgAddX509Cert(
		accAddress.String(),
		testconstants.IntermediateCertPem,
		testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// Check: DA + All + UniqueCertificate
	utils.EnsureDaIntermediateCertificateExist(
		t,
		setup,
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		testconstants.IntermediateIssuer,
		testconstants.IntermediateSerialNumber,
		false)

	// ChildCertificates: check that child certificates of issuer contains certificate identifier
	utils.EnsureChildCertificateExist(
		t,
		setup,
		testconstants.IntermediateIssuer,
		testconstants.IntermediateAuthorityKeyID,
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
	)

	// Check: ProposedCertificate - empty
	require.False(t, setup.Keeper.IsProposedCertificatePresent(
		setup.Ctx, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID))
}

// Extra cases

func TestHandler_AddDaIntermediateCert_VidScoped(t *testing.T) {
	setup := utils.Setup(t)

	accAddress := setup.CreateVendorAccount(testconstants.PAACertWithNumericVidVid)

	// store root certificate
	rootCertOptions := utils.CreatePAACertWithNumericVidOptions()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// add x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(
		accAddress.String(),
		testconstants.PAICertWithNumericPidVid,
		testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// Check: DA + All + UniqueCertificate
	intermediateCert := utils.EnsureDaIntermediateCertificateExist(
		t,
		setup,
		testconstants.PAICertWithNumericPidVidSubject,
		testconstants.PAICertWithNumericPidVidSubjectKeyID,
		testconstants.PAACertWithNumericVidSubject,
		testconstants.PAICertWithNumericPidVidSerialNumber,
		false)
	require.Equal(t, int32(testconstants.PAICertWithNumericPidVidVid), intermediateCert.Certs[0].Vid)

	// ChildCertificates: check that child certificates of issuer contains certificate identifier
	utils.EnsureChildCertificateExist(
		t,
		setup,
		testconstants.PAACertWithNumericVidSubject,
		testconstants.PAACertWithNumericVidSubjectKeyID,
		testconstants.PAICertWithNumericPidVidSubject,
		testconstants.PAICertWithNumericPidVidSubjectKeyID,
	)

	// Check: ProposedCertificate - empty
	require.False(t, setup.Keeper.IsProposedCertificatePresent(
		setup.Ctx, testconstants.PAICertWithNumericPidVidSubject, testconstants.PAICertWithNumericPidVidSubjectKeyID))
}

func TestHandler_AddDaIntermediateCert_SameSubjectAndSkid_DifferentSerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	vendorAccAddress := setup.CreateVendorAccount(testconstants.Vid)

	// store root certificate
	rootCertOptions := utils.CreateTestRootCertOptions()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// store intermediate certificate with different serial number
	intermediateCertificate := utils.IntermediateCertificateNoVid(vendorAccAddress)
	intermediateCertificate.SerialNumber = utils.SerialNumber
	utils.AddMokedDaCertificate(setup, intermediateCertificate, false)

	// store intermediate certificate second time
	addX509Cert := types.NewMsgAddX509Cert(
		vendorAccAddress.String(),
		testconstants.IntermediateCertPem,
		testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// query All approved certificate
	allApprovedCertificates, _ := utils.QueryAllApprovedCertificates(setup)
	require.Equal(t, 2, len(allApprovedCertificates)) // root + intermediate

	// query All certificate
	allCertificates, _ := utils.QueryAllCertificatesAll(setup)
	require.Equal(t, 2, len(allCertificates)) // root + intermediate

	// check approved certificate
	certificate, _ := utils.QueryApprovedCertificates(setup,
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID)
	require.Equal(t, 2, len(certificate.Certs)) // two intermediates
	require.NotEqual(t, certificate.Certs[0].SerialNumber, certificate.Certs[1].SerialNumber)

	// check global certificate
	globalCertificate, _ := utils.QueryAllCertificates(setup,
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID)
	require.Equal(t, 2, len(globalCertificate.Certs)) // two intermediates
	require.NotEqual(t, globalCertificate.Certs[0].SerialNumber, globalCertificate.Certs[1].SerialNumber)

	// Check indexes by subject key id
	approvedCertificatesBySubjectKeyId, _ := utils.QueryApprovedCertificatesBySubjectKeyID(setup, testconstants.IntermediateSubjectKeyID)
	require.Equal(t, 1, len(approvedCertificatesBySubjectKeyId))
	require.Equal(t, 2, len(approvedCertificatesBySubjectKeyId[0].Certs))

	allCertificatesBySubjectKeyId, _ := utils.QueryAllCertificatesBySubjectKeyID(setup, testconstants.IntermediateSubjectKeyID)
	require.Equal(t, 1, len(allCertificatesBySubjectKeyId))
	require.Equal(t, 2, len(allCertificatesBySubjectKeyId[0].Certs))
}

func TestHandler_AddDaCert_ForTree(t *testing.T) {
	setup := utils.Setup(t)

	vendorAccAddress := setup.CreateVendorAccount(testconstants.Vid)

	// add root x509 certificate
	rootCertOptions := utils.CreateTestRootCertOptions()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// add intermediate x509 certificate
	addIntermediateX509Cert := types.NewMsgAddX509Cert(
		vendorAccAddress.String(),
		testconstants.IntermediateCertPem,
		testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addIntermediateX509Cert)
	require.NoError(t, err)

	// add leaf x509 certificate
	addLeafX509Cert := types.NewMsgAddX509Cert(
		vendorAccAddress.String(),
		testconstants.LeafCertPem,
		testconstants.CertSchemaVersion)
	_, err = setup.Handler(setup.Ctx, addLeafX509Cert)
	require.NoError(t, err)

	// ensure root certificate exist
	utils.EnsureDaRootCertificateExist(
		t,
		setup,
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.RootIssuer,
		testconstants.RootSerialNumber)

	// ensure intermediate certificate exist
	utils.EnsureDaIntermediateCertificateExist(
		t,
		setup,
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		testconstants.IntermediateIssuer,
		testconstants.IntermediateSerialNumber,
		false)

	// ensure leaf certificate exist
	utils.EnsureDaIntermediateCertificateExist(
		t,
		setup,
		testconstants.LeafSubject,
		testconstants.LeafSubjectKeyID,
		testconstants.LeafIssuer,
		testconstants.LeafSerialNumber,
		false)

	// check ChildCertificate - root
	rootCertChildren, _ := utils.QueryChildCertificates(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, 1, len(rootCertChildren.CertIds))
	require.Equal(t,
		utils.CertificateIdentifier(testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID),
		*rootCertChildren.CertIds[0])

	// check ChildCertificate - intermediate
	intermediateCertChildren, _ := utils.QueryChildCertificates(
		setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Equal(t, 1, len(intermediateCertChildren.CertIds))
	require.Equal(t,
		utils.CertificateIdentifier(testconstants.LeafSubject, testconstants.LeafSubjectKeyID),
		*intermediateCertChildren.CertIds[0])

	// check child certificate identifiers of leaf certificate
	leafCertChildren, err := utils.QueryChildCertificates(setup, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
	require.Nil(t, leafCertChildren)
}

//nolint:funlen
func TestHandler_AddX509Cert_EachChildCertRefersToTwoParentCerts(t *testing.T) {
	setup := utils.Setup(t)

	vendorAccAddress := setup.CreateVendorAccount(testconstants.Vid)

	// store root certificate
	rootCert := utils.RootCertificate(setup.Trustee1)
	utils.AddMokedDaCertificate(setup, rootCert, true)

	// store second root certificate
	rootCert = utils.RootCertificate(setup.Trustee1)
	rootCert.SerialNumber = utils.SerialNumber
	utils.AddMokedDaCertificate(setup, rootCert, true)

	// store intermediate certificate (it refers to two parent certificates)
	intermediateCertificate := utils.IntermediateCertificateNoVid(vendorAccAddress)
	intermediateCertificate.SerialNumber = utils.SerialNumber
	utils.AddMokedDaCertificate(setup, intermediateCertificate, true)

	// store second intermediate certificate (it refers to two parent certificates)
	addX509Cert := types.NewMsgAddX509Cert(vendorAccAddress.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// store leaf certificate (it refers to two parent certificates)
	addX509Cert = types.NewMsgAddX509Cert(vendorAccAddress.String(), testconstants.LeafCertPem, testconstants.CertSchemaVersion)
	_, err = setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// query root certificate
	rootCertificates, _ := utils.QueryApprovedCertificates(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, 2, len(rootCertificates.Certs))

	// check child certificate identifiers of root certificate
	rootCertChildren, _ := utils.QueryChildCertificates(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)

	require.Equal(t, 1, len(rootCertChildren.CertIds))
	require.Equal(t,
		utils.CertificateIdentifier(testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID),
		*rootCertChildren.CertIds[0])

	// query intermediate certificate
	intermediateCertificates, _ := utils.QueryApprovedCertificates(
		setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Equal(t, 2, len(intermediateCertificates.Certs))

	// check child certificate identifiers of intermediate certificate
	intermediateCertChildren, _ := utils.QueryChildCertificates(
		setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)

	require.Equal(t, 1, len(intermediateCertChildren.CertIds))
	require.Equal(t,
		utils.CertificateIdentifier(testconstants.LeafSubject, testconstants.LeafSubjectKeyID),
		*intermediateCertChildren.CertIds[0])

	// query leaf certificate
	leafCertificates, _ := utils.QueryApprovedCertificates(setup, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Equal(t, 1, len(leafCertificates.Certs))

	// check child certificate identifiers of intermediate certificate
	leafCertChildren, err := utils.QueryChildCertificates(setup, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
	require.Nil(t, leafCertChildren)
}

func TestHandler_AddDaIntermediateCert_ByNotOwnerButSameVendor(t *testing.T) {
	setup := utils.Setup(t)

	// add two vendors with the same VID
	vendorAccAddress1 := setup.CreateVendorAccount(testconstants.Vid)
	vendorAccAddress2 := setup.CreateVendorAccount(testconstants.Vid)

	// store root certificate
	rootCertOptions := utils.CreateTestRootCertOptions()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// Store an intermediate certificate with the first vendor account as the owner
	intermediateCertificate := utils.IntermediateCertificateNoVid(vendorAccAddress1)
	intermediateCertificate.SerialNumber = utils.SerialNumber
	utils.AddMokedDaCertificate(setup, intermediateCertificate, false)

	// add an intermediate certificate with the same subject and SKID by second vendor account
	addX509Cert := types.NewMsgAddX509Cert(
		vendorAccAddress2.String(),
		testconstants.IntermediateCertPem,
		testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// ensure intermediate certificate exist
	// check list of certificates
	allApprovedCertificates, _ := utils.QueryAllApprovedCertificates(setup)
	require.Equal(t, 2, len(allApprovedCertificates)) // root + intermediate

	// check approved certificate
	certificate, _ := utils.QueryApprovedCertificates(setup,
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID)
	require.Equal(t, 2, len(certificate.Certs)) // two intermediates
	require.NotEqual(t, certificate.Certs[0].SerialNumber, certificate.Certs[1].SerialNumber)

	// Check indexes by subject key id
	approvedCertificatesBySubjectKeyId, _ := utils.QueryApprovedCertificatesBySubjectKeyID(setup, testconstants.IntermediateSubjectKeyID)
	require.Equal(t, 1, len(approvedCertificatesBySubjectKeyId))
	require.Equal(t, 2, len(approvedCertificatesBySubjectKeyId[0].Certs))
}

func TestHandler_AddX509Cert_VIDScopedRoot(t *testing.T) {
	setup := utils.Setup(t)

	// store root certificate
	rootCertOptions := utils.CreatePAACertWithNumericVidOptions()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	accAddress := utils.GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.PAACertWithNumericVidVid)

	// add x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(accAddress.String(), testconstants.PAICertWithNumericPidVid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// query certificate
	certs, _ := utils.QueryAllApprovedCertificates(setup)
	require.Equal(t, 2, len(certs))
	intermediateCerts, _ := utils.QueryApprovedCertificates(setup, testconstants.PAICertWithNumericPidVidSubject, testconstants.PAICertWithNumericPidVidSubjectKeyID)
	require.Equal(t, 1, len(intermediateCerts.Certs))
	require.Equal(t, testconstants.PAICertWithNumericPidVidSubject, intermediateCerts.Certs[0].Subject)
	require.Equal(t, testconstants.PAICertWithNumericPidVidSubjectKeyID, intermediateCerts.Certs[0].SubjectKeyId)
}

func TestHandler_AddX509Cert_NonVIDScopedRoot(t *testing.T) {
	accAddress := utils.GenerateAccAddress()

	cases := []struct {
		name                  string
		rootCertOptions       *utils.RootCertOptions
		childCert             string
		childCertSubject      string
		childCertSubjectKeyID string
		accountVid            int32
	}{
		{
			name:                  "VidScopedChild",
			rootCertOptions:       utils.CreatePAACertNoVidOptions(testconstants.PAICertWithVidVid),
			childCert:             testconstants.PAICertWithNumericVid,
			childCertSubject:      testconstants.PAICertWithNumericVidSubject,
			childCertSubjectKeyID: testconstants.PAICertWithNumericVidSubjectKeyID,
			accountVid:            testconstants.PAICertWithVidVid,
		},
		{
			name:                  "NonVidScopedChild",
			rootCertOptions:       utils.CreateTestRootCertOptions(),
			childCert:             testconstants.IntermediateCertPem,
			childCertSubject:      testconstants.IntermediateSubject,
			childCertSubjectKeyID: testconstants.IntermediateSubjectKeyID,
			accountVid:            testconstants.Vid,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)
			// store root certificate
			utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, tc.rootCertOptions)

			// add vendor account
			setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, tc.accountVid)

			// add x509 certificate
			addX509Cert := types.NewMsgAddX509Cert(accAddress.String(), tc.childCert, testconstants.CertSchemaVersion)
			_, err := setup.Handler(setup.Ctx, addX509Cert)
			require.NoError(t, err)

			// query certificate
			certs, _ := utils.QueryAllApprovedCertificates(setup)
			require.Equal(t, 2, len(certs))
			intermediateCerts, _ := utils.QueryApprovedCertificates(setup, tc.childCertSubject, tc.childCertSubjectKeyID)
			require.Equal(t, 1, len(intermediateCerts.Certs))
			require.Equal(t, tc.childCertSubject, intermediateCerts.Certs[0].Subject)
			require.Equal(t, tc.childCertSubjectKeyID, intermediateCerts.Certs[0].SubjectKeyId)
		})
	}
}

// Error cases

func TestHandler_AddX509Cert_ForInvalidCertificate(t *testing.T) {
	setup := utils.Setup(t)

	accAddress := utils.GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 1)

	// add x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(accAddress.String(), testconstants.StubCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrInvalidCertificate)
}

func TestHandler_AddX509Cert_ForRootCertificate(t *testing.T) {
	setup := utils.Setup(t)

	accAddress := utils.GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 1)

	// add root certificate as leaf x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(accAddress.String(), testconstants.RootCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrNonRootCertificateSelfSigned)
}

func TestHandler_AddX509Cert_ForDuplicate(t *testing.T) {
	setup := utils.Setup(t)

	// store root certificate
	rootCertificate := utils.RootCertificate(setup.Trustee1)
	setup.Keeper.AddAllCertificate(setup.Ctx, rootCertificate)

	accAddress := utils.GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 1)

	// store intermediate certificate
	addX509Cert := types.NewMsgAddX509Cert(accAddress.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// store intermediate certificate second time
	_, err = setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrCertificateAlreadyExists)
}

func TestHandler_AddX509Cert_ForExistingNocCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// store root certificate
	rootCertificate := utils.RootCertificate(setup.Trustee1)
	setup.Keeper.AddAllCertificate(setup.Ctx, rootCertificate)

	vendorAccAddress := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// Store the NOC certificate
	nocCertificate := utils.IntermediateCertificateNoVid(vendorAccAddress)
	nocCertificate.SerialNumber = testconstants.TestSerialNumber
	nocCertificate.CertificateType = types.CertificateType_OperationalPKI

	setup.Keeper.AddAllCertificate(setup.Ctx, nocCertificate)
	setup.Keeper.AddNocIcaCertificate(setup.Ctx, nocCertificate)
	uniqueCertificate := types.UniqueCertificate{
		Issuer:       nocCertificate.Issuer,
		SerialNumber: nocCertificate.SerialNumber,
		Present:      true,
	}
	setup.Keeper.SetUniqueCertificate(setup.Ctx, uniqueCertificate)

	// store intermediate certificate
	addX509Cert := types.NewMsgAddX509Cert(vendorAccAddress.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrInappropriateCertificateType)
}

func TestHandler_AddX509Cert_NoRootCert(t *testing.T) {
	setup := utils.Setup(t)

	vendorAccAddress := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add intermediate certificate
	intermediateCertificate := utils.IntermediateCertificateNoVid(vendorAccAddress)
	setup.Keeper.AddAllCertificate(setup.Ctx, intermediateCertificate)

	// add leaf x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(vendorAccAddress.String(), testconstants.LeafCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrInvalidCertificate)
}

func TestHandler_AddX509Cert_RootIsNoc(t *testing.T) {
	setup := utils.Setup(t)

	accAddress := utils.GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.IntermediateCertWithVid1Vid)

	// Add NOC root certificate
	addNocX509RootCert := types.NewMsgAddNocX509RootCert(accAddress.String(), testconstants.RootCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
	require.NoError(t, err)

	// add x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(accAddress.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err = setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrInappropriateCertificateType)
}

func TestHandler_AddX509Cert_ForAbsentDirectParentCert(t *testing.T) {
	setup := utils.Setup(t)

	vendorAccAddress := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add intermediate x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(vendorAccAddress.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrCertificateDoesNotExist)
}

func TestHandler_AddX509Cert_ForFailedCertificateVerification(t *testing.T) {
	setup := utils.Setup(t)

	// add invalid root
	invalidRootCertificate := types.NewRootCertificate(testconstants.StubCertPem,
		testconstants.RootSubject, testconstants.RootSubjectAsText, testconstants.RootSubjectKeyID,
		testconstants.RootSerialNumber, setup.Trustee1.String(), []*types.Grant{}, []*types.Grant{}, testconstants.Vid, testconstants.SchemaVersion)
	setup.Keeper.AddAllCertificate(setup.Ctx, invalidRootCertificate)

	vendorAccAddress := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add intermediate x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(vendorAccAddress.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrInvalidCertificate)
}

func TestHandler_AddX509Cert_ByOtherVendor(t *testing.T) {
	setup := utils.Setup(t)

	// store root certificate
	rootCertificate := utils.RootCertificate(setup.Trustee1)
	setup.Keeper.AddAllCertificate(setup.Ctx, rootCertificate)

	// add first vendor account with VID = 1
	vendorAccAddress1 := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// Store an intermediate certificate with the first vendor account as the owner
	intermediateCertificate := utils.IntermediateCertificateNoVid(vendorAccAddress1)
	intermediateCertificate.SerialNumber = utils.SerialNumber
	setup.Keeper.AddAllCertificate(setup.Ctx, intermediateCertificate)
	setup.Keeper.AddApprovedCertificateBySubjectKeyID(setup.Ctx, intermediateCertificate)
	setup.Keeper.SetUniqueCertificate(
		setup.Ctx,
		utils.UniqueCertificate(intermediateCertificate.Issuer, intermediateCertificate.SerialNumber),
	)

	// add seconf vendor account with VID = 1000
	vendorAccAddress2 := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID1)

	// add an intermediate certificate with the same subject and SKID by second vendor account
	addX509Cert := types.NewMsgAddX509Cert(vendorAccAddress2.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_AddX509Cert_SenderNotVendor(t *testing.T) {
	setup := utils.Setup(t)

	// store root certificate
	rootCertOptions := utils.CreateRootWithVidOptions()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// add x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateCertWithVid1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_AddX509Cert_VIDScopedRoot_NegativeCases(t *testing.T) {
	accAddress := utils.GenerateAccAddress()

	cases := []struct {
		name            string
		rootCertOptions *utils.RootCertOptions
		childCert       string
		accountVid      int32
		err             error
	}{
		{
			name:            "IncorrectChildVid",
			rootCertOptions: utils.CreateRootWithVidOptions(),
			childCert:       testconstants.IntermediateCertWithVid2,
			accountVid:      testconstants.RootCertWithVidVid,
			err:             pkitypes.ErrCertVidNotEqualToRootVid,
		},
		{
			name:            "IncorrectAccountVid",
			rootCertOptions: utils.CreateRootWithVidOptions(),
			childCert:       testconstants.IntermediateCertWithVid1,
			accountVid:      testconstants.Vid,
			err:             pkitypes.ErrCertVidNotEqualAccountVid,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)
			// store root certificate
			utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, tc.rootCertOptions)

			// add vendor account
			setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, tc.accountVid)

			// add x509 certificate
			addX509Cert := types.NewMsgAddX509Cert(accAddress.String(), tc.childCert, testconstants.CertSchemaVersion)
			_, err := setup.Handler(setup.Ctx, addX509Cert)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestHandler_AddX509Cert_NonVIDScopedRoot_NegativeCases(t *testing.T) {
	accAddress := utils.GenerateAccAddress()

	cases := []struct {
		name            string
		rootCertOptions *utils.RootCertOptions
		childCert       string
		accountVid      int32
		err             error
	}{
		{
			name:            "IncorrectChildVid",
			rootCertOptions: utils.CreatePAACertNoVidOptions(testconstants.Vid),
			childCert:       testconstants.PAICertWithNumericVid,
			accountVid:      testconstants.Vid,
			err:             pkitypes.ErrCertVidNotEqualToRootVid,
		},
		{
			name:            "IncorrectAccountVid",
			rootCertOptions: utils.CreatePAACertNoVidOptions(testconstants.PAICertWithVidVid),
			childCert:       testconstants.PAICertWithNumericVid,
			accountVid:      testconstants.Vid,
			err:             pkitypes.ErrCertVidNotEqualAccountVid,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)
			// store root certificate
			utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, tc.rootCertOptions)

			// add vendor account
			setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, tc.accountVid)

			// add x509 certificate
			addX509Cert := types.NewMsgAddX509Cert(accAddress.String(), tc.childCert, testconstants.CertSchemaVersion)
			_, err := setup.Handler(setup.Ctx, addX509Cert)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
