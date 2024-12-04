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

	// add DA root certificate
	testRootCertificate := utils.CreateTestRootCert()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, &testRootCertificate)

	// add DA PAI certificate
	testIntermediateCertificate := utils.CreateTestIntermediateCert()
	utils.AddDaIntermediateCertificate(setup, setup.Vendor1, testIntermediateCertificate.PEM)

	// Check state indexes
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ApprovedRootCertificatesKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.RejectedCertificateKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate, indexes)
}

// Extra cases

func TestHandler_AddDaIntermediateCert_VidScoped(t *testing.T) {
	setup := utils.Setup(t)

	accAddress := setup.CreateVendorAccount(testconstants.PAACertWithNumericVidVid)

	// store root certificate
	testRootCertificate := utils.CreateTestPAACertWithNumericVid()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, &testRootCertificate)

	// add intermediate certificate
	testIntermediateCertificate := utils.CreateTestIntermediateVidScopedCert()
	utils.AddDaIntermediateCertificate(setup, accAddress, testIntermediateCertificate.PEM)

	// Check state indexes
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ApprovedRootCertificatesKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.RejectedCertificateKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate, indexes)
}

func TestHandler_AddDaIntermediateCert_SameSubjectAndSkid_DifferentSerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	// store root certificate
	rootCertOptions := utils.CreateTestRootCert()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, &rootCertOptions)

	// store intermediate certificate with different serial number
	intermediateCertificate := utils.IntermediateCertificateNoVid(setup.Vendor1)
	intermediateCertificate.SerialNumber = utils.SerialNumber
	testIntermediateCertificate2 := utils.CreateTestIntermediateCert()
	testIntermediateCertificate2.SerialNumber = utils.SerialNumber
	utils.AddMokedDaCertificate(setup, intermediateCertificate, false)

	// store intermediate certificate second time
	testIntermediateCertificate1 := utils.CreateTestIntermediateCert()
	utils.AddDaIntermediateCertificate(setup, setup.Vendor1, testIntermediateCertificate1.PEM)

	// query All approved certificate
	allApprovedCertificates, _ := utils.QueryAllApprovedCertificates(setup)
	require.Equal(t, 2, len(allApprovedCertificates)) // root + intermediate

	// query All certificate
	allCertificates, _ := utils.QueryAllCertificatesAll(setup)
	require.Equal(t, 2, len(allCertificates)) // root + intermediate

	// Check indexes for certificate1
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix, Count: 2},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.ApprovedCertificatesKeyPrefix, Count: 2},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.ChildCertificatesKeyPrefix, Count: 1},
		},
		Missing: []utils.TestIndex{
			{Key: types.ApprovedRootCertificatesKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.RejectedCertificateKeyPrefix},
		},
	}
	resolvedCertificates := utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate1, indexes)

	// additional checks
	require.Equal(t, resolvedCertificates.ApprovedCertificates.Certs[0].SerialNumber, testIntermediateCertificate2.SerialNumber)
	require.Equal(t, resolvedCertificates.ApprovedCertificates.Certs[1].SerialNumber, testIntermediateCertificate1.SerialNumber)
	require.NotEqual(
		t,
		resolvedCertificates.ApprovedCertificates.Certs[0].SerialNumber,
		resolvedCertificates.ApprovedCertificates.Certs[1].SerialNumber,
	)

	// Check indexes for certificate2
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate2, indexes)
}

func TestHandler_AddDaCert_ForTree(t *testing.T) {
	setup := utils.Setup(t)

	// add root x509 certificate
	testRootCertificate := utils.CreateTestRootCert()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, &testRootCertificate)

	// add intermediate x509 certificate
	testIntermediateCertificate := utils.CreateTestIntermediateCert()
	utils.AddDaIntermediateCertificate(setup, setup.Vendor1, testIntermediateCertificate.PEM)

	// add leaf x509 certificate
	testLeafCertificate := utils.CreateTestLeafCert()
	utils.AddDaIntermediateCertificate(setup, setup.Vendor1, testLeafCertificate.PEM)

	// Check indexes for root
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
		Missing: []utils.TestIndex{},
	}
	utils.CheckCertificateStateIndexes(t, setup, testRootCertificate, indexes)

	// Check indexes for intermediate
	indexes = utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix, Count: 1},
		},
		Missing: []utils.TestIndex{
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate, indexes)

	// Check indexes for leaf
	utils.CheckCertificateStateIndexes(t, setup, testLeafCertificate, indexes)
}

//nolint:funlen
func TestHandler_AddX509Cert_EachChildCertRefersToTwoParentCerts(t *testing.T) {
	setup := utils.Setup(t)

	// store root certificate
	rootCert := utils.RootCertificate(setup.Trustee1)
	utils.AddMokedDaCertificate(setup, rootCert, true)

	// store second root certificate
	rootCert = utils.RootCertificate(setup.Trustee1)
	rootCert.SerialNumber = utils.SerialNumber
	utils.AddMokedDaCertificate(setup, rootCert, true)

	// store intermediate certificate (it refers to two parent certificates)
	intermediateCertificate := utils.IntermediateCertificateNoVid(setup.Vendor1)
	intermediateCertificate.SerialNumber = utils.SerialNumber
	utils.AddMokedDaCertificate(setup, intermediateCertificate, true)

	// store second intermediate certificate (it refers to two parent certificates)
	utils.AddDaIntermediateCertificate(setup, setup.Vendor1, testconstants.IntermediateCertPem)

	// store leaf certificate (it refers to two parent certificates)
	utils.AddDaIntermediateCertificate(setup, setup.Vendor1, testconstants.LeafCertPem)

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
	testRootCertificate := utils.CreateTestRootCert()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, &testRootCertificate)

	// Store an intermediate certificate with the first vendor account as the owner
	intermediateCertificate := utils.IntermediateCertificateNoVid(vendorAccAddress1)
	intermediateCertificate.SerialNumber = utils.SerialNumber
	utils.AddMokedDaCertificate(setup, intermediateCertificate, false)

	// add an intermediate certificate with the same subject and SKID by second vendor account
	testIntermediateCertificate := utils.CreateTestIntermediateCert()
	utils.AddDaIntermediateCertificate(setup, vendorAccAddress2, testIntermediateCertificate.PEM)

	// Check state indexes
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix, Count: 2},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.ApprovedCertificatesKeyPrefix, Count: 2},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.ChildCertificatesKeyPrefix, Count: 1},
		},
		Missing: []utils.TestIndex{
			{Key: types.ApprovedRootCertificatesKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.RejectedCertificateKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate, indexes)
}

func TestHandler_AddX509Cert_VIDScopedRoot(t *testing.T) {
	setup := utils.Setup(t)

	accAddress := setup.CreateVendorAccount(testconstants.PAACertWithNumericVidVid)

	// store root certificate
	rootCert := utils.CreateTestPAACertWithNumericVid()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, &rootCert)

	// add x509 certificate
	testIntermediateCertificate := utils.CreateTestIntermediateVidScopedCert()
	utils.AddDaIntermediateCertificate(setup, accAddress, testIntermediateCertificate.PEM)

	// Check state indexes
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ApprovedRootCertificatesKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.RejectedCertificateKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, testIntermediateCertificate, indexes)
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
			utils.ProposeAndApproveRootCertificateByOptions(setup, setup.Trustee1, tc.rootCertOptions)

			// add vendor account
			setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, tc.accountVid)

			// add x509 certificate
			utils.AddDaIntermediateCertificate(setup, accAddress, tc.childCert)

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

	accAddress := setup.CreateVendorAccount(1)

	// add x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(accAddress.String(), testconstants.StubCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrInvalidCertificate)
}

func TestHandler_AddX509Cert_ForRootCertificate(t *testing.T) {
	setup := utils.Setup(t)

	accAddress := setup.CreateVendorAccount(1)

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

	accAddress := setup.CreateVendorAccount(1)

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

	// Store the NOC certificate
	nocCertificate := utils.IntermediateCertificateNoVid(setup.Vendor1)
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
	addX509Cert := types.NewMsgAddX509Cert(setup.Vendor1.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrInappropriateCertificateType)
}

func TestHandler_AddX509Cert_NoRootCert(t *testing.T) {
	setup := utils.Setup(t)

	// add intermediate certificate
	intermediateCertificate := utils.IntermediateCertificateNoVid(setup.Vendor1)
	setup.Keeper.AddAllCertificate(setup.Ctx, intermediateCertificate)

	// add leaf x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Vendor1.String(), testconstants.LeafCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrInvalidCertificate)
}

func TestHandler_AddX509Cert_RootIsNoc(t *testing.T) {
	setup := utils.Setup(t)

	accAddress := setup.CreateVendorAccount(testconstants.IntermediateCertWithVid1Vid)

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

	// add intermediate x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Vendor1.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
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

	// add intermediate x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Vendor1.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrInvalidCertificate)
}

func TestHandler_AddX509Cert_ByOtherVendor(t *testing.T) {
	setup := utils.Setup(t)

	// store root certificate
	rootCertificate := utils.RootCertificate(setup.Trustee1)
	setup.Keeper.AddAllCertificate(setup.Ctx, rootCertificate)

	// add first vendor account with VID = 1
	vendorAccAddress1 := setup.CreateVendorAccount(testconstants.Vid)

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
	rootCert := utils.CreateTestRootCertWithVid()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, &rootCert)

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
			utils.ProposeAndApproveRootCertificateByOptions(setup, setup.Trustee1, tc.rootCertOptions)

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
			utils.ProposeAndApproveRootCertificateByOptions(setup, setup.Trustee1, tc.rootCertOptions)

			// add vendor account
			setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, tc.accountVid)

			// add x509 certificate
			addX509Cert := types.NewMsgAddX509Cert(accAddress.String(), tc.childCert, testconstants.CertSchemaVersion)
			_, err := setup.Handler(setup.Ctx, addX509Cert)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
