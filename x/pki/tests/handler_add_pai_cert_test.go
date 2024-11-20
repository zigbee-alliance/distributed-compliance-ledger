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

func TestHandler_AddX509Cert(t *testing.T) {
	setup := Setup(t)

	accAddress := GenerateAccAddress()
	vid := testconstants.Vid
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add DA PAA certificate
	rootCertOptions := createTestRootCertOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// add DA PAI certificate
	addX509Cert := types.NewMsgAddX509Cert(accAddress.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// Check that root certificate exists
	ensureDaPaiCertificateExist(
		t,
		setup,
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectKeyID,
		testconstants.IntermediateIssuer,
		testconstants.IntermediateSerialNumber,
		false)

	// ChildCertificates: check that child certificates of issuer contains certificate identifier
	issuerChildren, _ := queryChildCertificates(
		setup, testconstants.IntermediateIssuer, testconstants.IntermediateAuthorityKeyID)
	require.Equal(t, 1, len(issuerChildren.CertIds))
	require.Equal(t,
		&types.CertificateIdentifier{
			Subject:      testconstants.IntermediateSubject,
			SubjectKeyId: testconstants.IntermediateSubjectKeyID,
		},
		issuerChildren.CertIds[0])

	// check that no proposed certificate has been created
	_, err = queryProposedCertificate(setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
}

// Extra cases

func TestHandler_AddX509Cert_VIDScoped(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCertOptions := createPAACertWithNumericVidOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	accAddress := GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.PAACertWithNumericVidVid)

	// add x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(accAddress.String(), testconstants.PAICertWithNumericPidVid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// query certificate
	intermediateCerts, _ := queryApprovedCertificates(setup, testconstants.PAICertWithNumericPidVidSubject, testconstants.PAICertWithNumericPidVidSubjectKeyID)
	require.Equal(t, 1, len(intermediateCerts.Certs))
	require.Equal(t, testconstants.PAICertWithNumericPidVidSubject, intermediateCerts.Certs[0].Subject)
	require.Equal(t, testconstants.PAICertWithNumericPidVidSubjectKeyID, intermediateCerts.Certs[0].SubjectKeyId)
	require.Equal(t, int32(testconstants.PAICertWithNumericPidVidVid), intermediateCerts.Certs[0].Vid)
}

func TestHandler_AddX509Cert_ForDifferentSerialNumber(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee1)
	setup.Keeper.AddAllCertificate(setup.Ctx, rootCertificate)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// store intermediate certificate with different serial number
	intermediateCertificate := intermediateCertificateNoVid(vendorAccAddress)
	intermediateCertificate.SerialNumber = SerialNumber
	setup.Keeper.SetUniqueCertificate(
		setup.Ctx,
		uniqueCertificate(intermediateCertificate.Issuer, intermediateCertificate.SerialNumber),
	)
	setup.Keeper.AddAllCertificate(setup.Ctx, intermediateCertificate)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, intermediateCertificate)

	// store intermediate certificate second time
	addX509Cert := types.NewMsgAddX509Cert(vendorAccAddress.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// query certificate
	certificates, _ := queryApprovedCertificates(setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)

	// check
	require.Equal(t, 2, len(certificates.Certs))
	require.NotEqual(t, certificates.Certs[0].SerialNumber, certificates.Certs[1].SerialNumber)

	for _, certificate := range certificates.Certs {
		require.Equal(t, addX509Cert.Cert, certificate.PemCert)
		require.Equal(t, addX509Cert.Signer, certificate.Owner)
		require.Equal(t, testconstants.IntermediateSubject, certificate.Subject)
		require.Equal(t, testconstants.IntermediateSubjectKeyID, certificate.SubjectKeyId)
		require.False(t, certificate.IsRoot)
		require.Equal(t, testconstants.RootSubject, certificate.RootSubject)
		require.Equal(t, testconstants.RootSubjectKeyID, certificate.RootSubjectKeyId)
		require.Equal(t, testconstants.IntermediateIssuer, certificate.Issuer)
		require.Equal(t, testconstants.IntermediateAuthorityKeyID, certificate.AuthorityKeyId)
	}
}

func TestHandler_AddX509Cert_ForTree(t *testing.T) {
	setup := Setup(t)

	// add root x509 certificate
	rootCertOptions := createTestRootCertOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add intermediate x509 certificate
	addIntermediateX509Cert := types.NewMsgAddX509Cert(vendorAccAddress.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addIntermediateX509Cert)
	require.NoError(t, err)

	// add leaf x509 certificate
	addLeafX509Cert := types.NewMsgAddX509Cert(vendorAccAddress.String(), testconstants.LeafCertPem, testconstants.CertSchemaVersion)
	_, err = setup.Handler(setup.Ctx, addLeafX509Cert)
	require.NoError(t, err)

	// query root certificate
	rootCertificate, _ := querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, testconstants.RootCertPem, rootCertificate.PemCert)

	// check child certificate identifiers of root certificate
	rootCertChildren, _ := queryChildCertificates(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)

	require.Equal(t, 1, len(rootCertChildren.CertIds))
	require.Equal(t,
		certificateIdentifier(testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID),
		*rootCertChildren.CertIds[0])

	// query intermediate certificate
	intermediateCertificate, _ := querySingleApprovedCertificate(setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Equal(t, testconstants.IntermediateCertPem, intermediateCertificate.PemCert)

	// check child certificate identifiers of intermediate certificate
	intermediateCertChildren, _ := queryChildCertificates(
		setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)

	require.Equal(t, 1, len(intermediateCertChildren.CertIds))
	require.Equal(t,
		certificateIdentifier(testconstants.LeafSubject, testconstants.LeafSubjectKeyID),
		*intermediateCertChildren.CertIds[0])

	// query leaf certificate
	leafCertificate, _ := querySingleApprovedCertificate(setup, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Equal(t, testconstants.LeafCertPem, leafCertificate.PemCert)

	// check child certificate identifiers of leaf certificate
	leafCertChildren, err := queryChildCertificates(setup, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
	require.Nil(t, leafCertChildren)
}

//nolint:funlen
func TestHandler_AddX509Cert_EachChildCertRefersToTwoParentCerts(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCert := rootCertificate(setup.Trustee1)

	setup.Keeper.AddAllCertificate(setup.Ctx, rootCert)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCert)
	setup.Keeper.SetUniqueCertificate(setup.Ctx, uniqueCertificate(rootCert.Subject, rootCert.SerialNumber))

	// store second root certificate
	rootCert = rootCertificate(setup.Trustee1)
	rootCert.SerialNumber = SerialNumber

	setup.Keeper.AddAllCertificate(setup.Ctx, rootCert)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCert)
	setup.Keeper.SetUniqueCertificate(setup.Ctx, uniqueCertificate(rootCert.Subject, rootCert.SerialNumber))

	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// store intermediate certificate (it refers to two parent certificates)
	intermediateCertificate := intermediateCertificateNoVid(vendorAccAddress)
	intermediateCertificate.SerialNumber = SerialNumber

	setup.Keeper.AddAllCertificate(setup.Ctx, intermediateCertificate)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, intermediateCertificate)
	setup.Keeper.SetUniqueCertificate(
		setup.Ctx,
		uniqueCertificate(intermediateCertificate.Issuer, intermediateCertificate.SerialNumber),
	)

	childCertID := certificateIdentifier(intermediateCertificate.Subject, intermediateCertificate.SubjectKeyId)
	rootChildCertificates := types.ChildCertificates{
		Issuer:         intermediateCertificate.Issuer,
		AuthorityKeyId: intermediateCertificate.AuthorityKeyId,
		CertIds:        []*types.CertificateIdentifier{&childCertID},
	}
	setup.Keeper.SetChildCertificates(setup.Ctx, rootChildCertificates)

	// store second intermediate certificate (it refers to two parent certificates)
	addX509Cert := types.NewMsgAddX509Cert(vendorAccAddress.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// store leaf certificate (it refers to two parent certificates)
	addX509Cert = types.NewMsgAddX509Cert(vendorAccAddress.String(), testconstants.LeafCertPem, testconstants.CertSchemaVersion)
	_, err = setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// query root certificate
	rootCertificates, _ := queryApprovedCertificates(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, 2, len(rootCertificates.Certs))

	// check child certificate identifiers of root certificate
	rootCertChildren, _ := queryChildCertificates(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)

	require.Equal(t, 1, len(rootCertChildren.CertIds))
	require.Equal(t,
		certificateIdentifier(testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID),
		*rootCertChildren.CertIds[0])

	// query intermediate certificate
	intermediateCertificates, _ := queryApprovedCertificates(
		setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Equal(t, 2, len(intermediateCertificates.Certs))

	// check child certificate identifiers of intermediate certificate
	intermediateCertChildren, _ := queryChildCertificates(
		setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)

	require.Equal(t, 1, len(intermediateCertChildren.CertIds))
	require.Equal(t,
		certificateIdentifier(testconstants.LeafSubject, testconstants.LeafSubjectKeyID),
		*intermediateCertChildren.CertIds[0])

	// query leaf certificate
	leafCertificates, _ := queryApprovedCertificates(setup, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Equal(t, 1, len(leafCertificates.Certs))

	// check child certificate identifiers of intermediate certificate
	leafCertChildren, err := queryChildCertificates(setup, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
	require.Nil(t, leafCertChildren)
}

func TestHandler_AddX509Cert_ByNotOwnerButSameVendor(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee1)
	setup.Keeper.AddAllCertificate(setup.Ctx, rootCertificate)

	// add first vendor account with VID = 1
	vendorAccAddress1 := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// Store an intermediate certificate with the first vendor account as the owner
	intermediateCertificate := intermediateCertificateNoVid(vendorAccAddress1)
	intermediateCertificate.SerialNumber = SerialNumber
	setup.Keeper.AddAllCertificate(setup.Ctx, intermediateCertificate)
	setup.Keeper.AddApprovedCertificateBySubjectKeyID(setup.Ctx, intermediateCertificate)
	setup.Keeper.SetUniqueCertificate(
		setup.Ctx,
		uniqueCertificate(intermediateCertificate.Issuer, intermediateCertificate.SerialNumber),
	)

	// add second vendor account with VID = 1
	vendorAccAddress2 := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add an intermediate certificate with the same subject and SKID by second vendor account
	addX509Cert := types.NewMsgAddX509Cert(vendorAccAddress2.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)
}

func TestHandler_AddX509Cert_VIDScopedRoot(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCertOptions := createPAACertWithNumericVidOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	accAddress := GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.PAACertWithNumericVidVid)

	// add x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(accAddress.String(), testconstants.PAICertWithNumericPidVid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// query certificate
	certs, _ := queryAllApprovedCertificates(setup)
	require.Equal(t, 2, len(certs))
	intermediateCerts, _ := queryApprovedCertificates(setup, testconstants.PAICertWithNumericPidVidSubject, testconstants.PAICertWithNumericPidVidSubjectKeyID)
	require.Equal(t, 1, len(intermediateCerts.Certs))
	require.Equal(t, testconstants.PAICertWithNumericPidVidSubject, intermediateCerts.Certs[0].Subject)
	require.Equal(t, testconstants.PAICertWithNumericPidVidSubjectKeyID, intermediateCerts.Certs[0].SubjectKeyId)
}

func TestHandler_AddX509Cert_NonVIDScopedRoot(t *testing.T) {
	accAddress := GenerateAccAddress()

	cases := []struct {
		name                  string
		rootCertOptions       *rootCertOptions
		childCert             string
		childCertSubject      string
		childCertSubjectKeyID string
		accountVid            int32
	}{
		{
			name:                  "VidScopedChild",
			rootCertOptions:       createPAACertNoVidOptions(testconstants.PAICertWithVidVid),
			childCert:             testconstants.PAICertWithNumericVid,
			childCertSubject:      testconstants.PAICertWithNumericVidSubject,
			childCertSubjectKeyID: testconstants.PAICertWithNumericVidSubjectKeyID,
			accountVid:            testconstants.PAICertWithVidVid,
		},
		{
			name:                  "NonVidScopedChild",
			rootCertOptions:       createTestRootCertOptions(),
			childCert:             testconstants.IntermediateCertPem,
			childCertSubject:      testconstants.IntermediateSubject,
			childCertSubjectKeyID: testconstants.IntermediateSubjectKeyID,
			accountVid:            testconstants.Vid,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := Setup(t)
			// store root certificate
			proposeAndApproveRootCertificate(setup, setup.Trustee1, tc.rootCertOptions)

			// add vendor account
			setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, tc.accountVid)

			// add x509 certificate
			addX509Cert := types.NewMsgAddX509Cert(accAddress.String(), tc.childCert, testconstants.CertSchemaVersion)
			_, err := setup.Handler(setup.Ctx, addX509Cert)
			require.NoError(t, err)

			// query certificate
			certs, _ := queryAllApprovedCertificates(setup)
			require.Equal(t, 2, len(certs))
			intermediateCerts, _ := queryApprovedCertificates(setup, tc.childCertSubject, tc.childCertSubjectKeyID)
			require.Equal(t, 1, len(intermediateCerts.Certs))
			require.Equal(t, tc.childCertSubject, intermediateCerts.Certs[0].Subject)
			require.Equal(t, tc.childCertSubjectKeyID, intermediateCerts.Certs[0].SubjectKeyId)
		})
	}
}

// Error cases

func TestHandler_AddX509Cert_ForInvalidCertificate(t *testing.T) {
	setup := Setup(t)

	accAddress := GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 1)

	// add x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(accAddress.String(), testconstants.StubCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrInvalidCertificate)
}

func TestHandler_AddX509Cert_ForRootCertificate(t *testing.T) {
	setup := Setup(t)

	accAddress := GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 1)

	// add root certificate as leaf x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(accAddress.String(), testconstants.RootCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrNonRootCertificateSelfSigned)
}

func TestHandler_AddX509Cert_ForDuplicate(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee1)
	setup.Keeper.AddAllCertificate(setup.Ctx, rootCertificate)

	accAddress := GenerateAccAddress()
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
	setup := Setup(t)

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee1)
	setup.Keeper.AddAllCertificate(setup.Ctx, rootCertificate)

	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// Store the NOC certificate
	nocCertificate := intermediateCertificateNoVid(vendorAccAddress)
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
	setup := Setup(t)

	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add intermediate certificate
	intermediateCertificate := intermediateCertificateNoVid(vendorAccAddress)
	setup.Keeper.AddAllCertificate(setup.Ctx, intermediateCertificate)

	// add leaf x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(vendorAccAddress.String(), testconstants.LeafCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrInvalidCertificate)
}

func TestHandler_AddX509Cert_RootIsNoc(t *testing.T) {
	setup := Setup(t)

	accAddress := GenerateAccAddress()
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
	setup := Setup(t)

	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add intermediate x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(vendorAccAddress.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrCertificateDoesNotExist)
}

func TestHandler_AddX509Cert_ForFailedCertificateVerification(t *testing.T) {
	setup := Setup(t)

	// add invalid root
	invalidRootCertificate := types.NewRootCertificate(testconstants.StubCertPem,
		testconstants.RootSubject, testconstants.RootSubjectAsText, testconstants.RootSubjectKeyID,
		testconstants.RootSerialNumber, setup.Trustee1.String(), []*types.Grant{}, []*types.Grant{}, testconstants.Vid, testconstants.SchemaVersion)
	setup.Keeper.AddAllCertificate(setup.Ctx, invalidRootCertificate)

	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add intermediate x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(vendorAccAddress.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrInvalidCertificate)
}

func TestHandler_AddX509Cert_ByOtherVendor(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee1)
	setup.Keeper.AddAllCertificate(setup.Ctx, rootCertificate)

	// add first vendor account with VID = 1
	vendorAccAddress1 := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// Store an intermediate certificate with the first vendor account as the owner
	intermediateCertificate := intermediateCertificateNoVid(vendorAccAddress1)
	intermediateCertificate.SerialNumber = SerialNumber
	setup.Keeper.AddAllCertificate(setup.Ctx, intermediateCertificate)
	setup.Keeper.AddApprovedCertificateBySubjectKeyID(setup.Ctx, intermediateCertificate)
	setup.Keeper.SetUniqueCertificate(
		setup.Ctx,
		uniqueCertificate(intermediateCertificate.Issuer, intermediateCertificate.SerialNumber),
	)

	// add seconf vendor account with VID = 1000
	vendorAccAddress2 := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID1)

	// add an intermediate certificate with the same subject and SKID by second vendor account
	addX509Cert := types.NewMsgAddX509Cert(vendorAccAddress2.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_AddX509Cert_SenderNotVendor(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCertOptions := createRootWithVidOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// add x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateCertWithVid1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_AddX509Cert_VIDScopedRoot_NegativeCases(t *testing.T) {
	accAddress := GenerateAccAddress()

	cases := []struct {
		name            string
		rootCertOptions *rootCertOptions
		childCert       string
		accountVid      int32
		err             error
	}{
		{
			name:            "IncorrectChildVid",
			rootCertOptions: createRootWithVidOptions(),
			childCert:       testconstants.IntermediateCertWithVid2,
			accountVid:      testconstants.RootCertWithVidVid,
			err:             pkitypes.ErrCertVidNotEqualToRootVid,
		},
		{
			name:            "IncorrectAccountVid",
			rootCertOptions: createRootWithVidOptions(),
			childCert:       testconstants.IntermediateCertWithVid1,
			accountVid:      testconstants.Vid,
			err:             pkitypes.ErrCertVidNotEqualAccountVid,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := Setup(t)
			// store root certificate
			proposeAndApproveRootCertificate(setup, setup.Trustee1, tc.rootCertOptions)

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
	accAddress := GenerateAccAddress()

	cases := []struct {
		name            string
		rootCertOptions *rootCertOptions
		childCert       string
		accountVid      int32
		err             error
	}{
		{
			name:            "IncorrectChildVid",
			rootCertOptions: createPAACertNoVidOptions(testconstants.Vid),
			childCert:       testconstants.PAICertWithNumericVid,
			accountVid:      testconstants.Vid,
			err:             pkitypes.ErrCertVidNotEqualToRootVid,
		},
		{
			name:            "IncorrectAccountVid",
			rootCertOptions: createPAACertNoVidOptions(testconstants.PAICertWithVidVid),
			childCert:       testconstants.PAICertWithNumericVid,
			accountVid:      testconstants.Vid,
			err:             pkitypes.ErrCertVidNotEqualAccountVid,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := Setup(t)
			// store root certificate
			proposeAndApproveRootCertificate(setup, setup.Trustee1, tc.rootCertOptions)

			// add vendor account
			setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, tc.accountVid)

			// add x509 certificate
			addX509Cert := types.NewMsgAddX509Cert(accAddress.String(), tc.childCert, testconstants.CertSchemaVersion)
			_, err := setup.Handler(setup.Ctx, addX509Cert)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
