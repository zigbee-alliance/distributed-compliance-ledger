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

func TestHandler_AddX509Cert(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee1)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	for i, role := range []dclauthtypes.AccountRole{
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

		// query certificate
		certificate, _ := querySingleApprovedCertificate(
			setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
		require.Equal(t, intermediateCertificate(accAddress), *certificate)

		certificateBySubjectKeyID, _ := queryAllApprovedCertificatesBySubjectKeyID(setup, testconstants.IntermediateSubjectKeyID)
		require.Equal(t, 1, len(certificateBySubjectKeyID))
		require.Equal(t, i+1, len(certificateBySubjectKeyID[0].Certs))

		certs := make([]*types.Certificate, 0)
		certs = append(certs, certificate, certificateBySubjectKeyID[0].Certs[i])
		for _, cert := range certs {
			// check
			require.Equal(t, addX509Cert.Cert, cert.PemCert)
			require.Equal(t, addX509Cert.Signer, cert.Owner)
			require.Equal(t, testconstants.IntermediateSubject, cert.Subject)
			require.Equal(t, testconstants.IntermediateSubjectKeyID, cert.SubjectKeyId)
			require.Equal(t, testconstants.IntermediateSerialNumber, cert.SerialNumber)
			require.False(t, cert.IsRoot)
			require.Equal(t, testconstants.IntermediateIssuer, cert.Issuer)
			require.Equal(t, testconstants.IntermediateAuthorityKeyID, cert.AuthorityKeyId)
			require.Equal(t, testconstants.RootSubject, cert.RootSubject)
			require.Equal(t, testconstants.RootSubjectKeyID, cert.RootSubjectKeyId)
		}

		// check that unique certificate key is registered
		require.True(t, setup.Keeper.IsUniqueCertificatePresent(
			setup.Ctx, testconstants.IntermediateIssuer, testconstants.IntermediateSerialNumber))

		// check that child certificates of issuer contains certificate identifier
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

		// cleanup for next iteration
		setup.Keeper.RemoveApprovedCertificates(setup.Ctx,
			testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
		setup.Keeper.RemoveUniqueCertificate(setup.Ctx,
			testconstants.IntermediateIssuer, testconstants.IntermediateSerialNumber)
		setup.Keeper.RemoveChildCertificates(setup.Ctx,
			testconstants.IntermediateIssuer, testconstants.IntermediateAuthorityKeyID)
	}
}

func TestHandler_AddX509Cert_ForInvalidCertificate(t *testing.T) {
	setup := Setup(t)

	// add x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.StubCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrInvalidCertificate.Is(err))
}

func TestHandler_AddX509Cert_ForRootCertificate(t *testing.T) {
	setup := Setup(t)

	// add root certificate as leaf x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.RootCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrInappropriateCertificateType.Is(err))
}

func TestHandler_AddX509Cert_ForDuplicate(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee1)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// store intermediate certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// store intermediate certificate second time
	_, err = setup.Handler(setup.Ctx, addX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateAlreadyExists.Is(err))
}

func TestHandler_AddX509Cert_ForNocCertificate(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee1)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// Store the NOC certificate
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	nocCertificate := intermediateCertificate(vendorAccAddress)
	nocCertificate.SerialNumber = testconstants.TestSerialNumber
	nocCertificate.IsNoc = true

	setup.Keeper.AddApprovedCertificate(setup.Ctx, nocCertificate)
	// TODO: add the certificate to the ICA store after the store is implemented
	uniqueCertificate := types.UniqueCertificate{
		Issuer:       nocCertificate.Issuer,
		SerialNumber: nocCertificate.SerialNumber,
		Present:      true,
	}
	setup.Keeper.SetUniqueCertificate(setup.Ctx, uniqueCertificate)

	// store intermediate certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.True(t, pkitypes.ErrInappropriateCertificateType.Is(err))
}

func TestHandler_AddX509Cert_ForDifferentSerialNumber(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee1)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// store intermediate certificate with different serial number
	intermediateCertificate := intermediateCertificate(setup.Trustee1)
	intermediateCertificate.SerialNumber = SerialNumber
	setup.Keeper.SetUniqueCertificate(
		setup.Ctx,
		uniqueCertificate(intermediateCertificate.Issuer, intermediateCertificate.SerialNumber),
	)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, intermediateCertificate)

	// store intermediate certificate second time
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateCertPem)
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

func TestHandler_AddX509Cert_ForDifferentSerialNumberDifferentSigner(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCertificate := rootCertificate(testconstants.Address1)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// store intermediate certificate with different serial number
	intermediateCertificate := intermediateCertificate(testconstants.Address1)
	intermediateCertificate.SerialNumber = SerialNumber
	setup.Keeper.SetUniqueCertificate(
		setup.Ctx,
		uniqueCertificate(intermediateCertificate.Issuer, intermediateCertificate.SerialNumber),
	)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, intermediateCertificate)

	// store intermediate certificate second time
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_AddX509Cert_ForAbsentDirectParentCert(t *testing.T) {
	setup := Setup(t)

	// add intermediate x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrInvalidCertificate.Is(err))
}

func TestHandler_AddX509Cert_ForNoRootCert(t *testing.T) {
	setup := Setup(t)

	// add intermediate certificate
	intermediateCertificate := intermediateCertificate(setup.Trustee1)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, intermediateCertificate)

	// add leaf x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.LeafCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrInvalidCertificate.Is(err))
}

func TestHandler_AddX509Cert_ForFailedCertificateVerification(t *testing.T) {
	setup := Setup(t)

	// add invalid root
	invalidRootCertificate := types.NewRootCertificate(testconstants.StubCertPem,
		testconstants.RootSubject, testconstants.RootSubjectAsText, testconstants.RootSubjectKeyID,
		testconstants.RootSerialNumber, setup.Trustee1.String(), []*types.Grant{}, []*types.Grant{}, testconstants.Vid)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, invalidRootCertificate)

	// add intermediate x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrInvalidCertificate.Is(err))
}

func TestHandler_AddX509Cert_ForTree(t *testing.T) {
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

	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCert)
	setup.Keeper.SetUniqueCertificate(setup.Ctx, uniqueCertificate(rootCert.Subject, rootCert.SerialNumber))

	// store second root certificate
	rootCert = rootCertificate(setup.Trustee1)
	rootCert.SerialNumber = SerialNumber

	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCert)
	setup.Keeper.SetUniqueCertificate(setup.Ctx, uniqueCertificate(rootCert.Subject, rootCert.SerialNumber))

	// store intermediate certificate (it refers to two parent certificates)
	intermediateCertificate := intermediateCertificate(setup.Trustee1)
	intermediateCertificate.SerialNumber = SerialNumber

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
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// store leaf certificate (it refers to two parent certificates)
	addX509Cert = types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.LeafCertPem)
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

func TestHandler_AddX509Cert_ByNotOwner(t *testing.T) {
	setup := Setup(t)

	rootCertificate := rootCertificate(setup.Trustee1)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// add an intermidiate certificate from vendor
	intermediateCertificate := intermediateCertificate(setup.Trustee1)
	intermediateCertificate.SerialNumber = SerialNumber
	setup.Keeper.AddApprovedCertificate(setup.Ctx, intermediateCertificate)
	setup.Keeper.AddApprovedCertificateBySubjectKeyID(setup.Ctx, intermediateCertificate)
	setup.Keeper.SetUniqueCertificate(
		setup.Ctx,
		uniqueCertificate(intermediateCertificate.Issuer, intermediateCertificate.SerialNumber),
	)

	// add an intermediate certificate with the same subject and SKID from the same vendor but under a different account
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee2.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_AddX509Cert_ByNotOwnerButSameVendor(t *testing.T) {
	setup := Setup(t)

	rootCertificate := rootCertificate(setup.Trustee1)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	vendorAccAddress1 := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add an intermidiate certificate from vendor
	intermediateCertificate := intermediateCertificate(vendorAccAddress1)
	intermediateCertificate.SerialNumber = SerialNumber
	setup.Keeper.AddApprovedCertificate(setup.Ctx, intermediateCertificate)
	setup.Keeper.AddApprovedCertificateBySubjectKeyID(setup.Ctx, intermediateCertificate)
	setup.Keeper.SetUniqueCertificate(
		setup.Ctx,
		uniqueCertificate(intermediateCertificate.Issuer, intermediateCertificate.SerialNumber),
	)

	vendorAccAddress2 := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add an intermediate certificate with the same subject and SKID from the same vendor but under a different account
	addX509Cert := types.NewMsgAddX509Cert(vendorAccAddress2.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)
}

func TestHandler_AddX509Cert_ByNotOwnerAndOtherVendor(t *testing.T) {
	setup := Setup(t)

	rootCertificate := rootCertificate(setup.Trustee1)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	vendorAccAddress1 := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add an intermidiate certificate from vendor
	intermediateCertificate := intermediateCertificate(vendorAccAddress1)
	intermediateCertificate.SerialNumber = SerialNumber
	setup.Keeper.AddApprovedCertificate(setup.Ctx, intermediateCertificate)
	setup.Keeper.AddApprovedCertificateBySubjectKeyID(setup.Ctx, intermediateCertificate)
	setup.Keeper.SetUniqueCertificate(
		setup.Ctx,
		uniqueCertificate(intermediateCertificate.Issuer, intermediateCertificate.SerialNumber),
	)

	vendorAccAddress2 := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID1)

	// add an intermediate certificate with the same subject and SKID from the same vendor but under a different account
	addX509Cert := types.NewMsgAddX509Cert(vendorAccAddress2.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}
