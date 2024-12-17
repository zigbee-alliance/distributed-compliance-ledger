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
)

// Main

func TestHandler_AddDaIntermediateCert(t *testing.T) {
	setup := utils.Setup(t)

	// add DA root certificate
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate)

	// add DA PAI certificate
	testIntermediateCertificate := utils.IntermediateDaCertificate(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, testIntermediateCertificate)

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

func TestHandler_AddDaIntermediateCert_VidScoped(t *testing.T) {
	setup := utils.Setup(t)

	accAddress := setup.CreateVendorAccount(testconstants.PAACertWithNumericVidVid)

	// store root certificate
	testRootCertificate := utils.RootDaCertificateWithNumericVid(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, testRootCertificate)

	// add intermediate certificate
	testIntermediateCertificate := utils.IntermediateDaCertificateWithNumericPidVid(accAddress)
	utils.AddDaIntermediateCertificate(setup, testIntermediateCertificate)

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
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate)

	// store intermediate certificate with different serial number
	intermediateCertificate := utils.IntermediateDaCertificate(setup.Vendor1)
	intermediateCertificate.SerialNumber = utils.SerialNumber
	utils.AddMokedDaCertificate(setup, intermediateCertificate)

	// store intermediate certificate second time
	testIntermediateCertificate1 := utils.IntermediateDaCertificate(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, testIntermediateCertificate1)

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
	require.Equal(t, resolvedCertificates.ApprovedCertificates.Certs[0].SerialNumber, intermediateCertificate.SerialNumber)
	require.Equal(t, resolvedCertificates.ApprovedCertificates.Certs[1].SerialNumber, testIntermediateCertificate1.SerialNumber)
	require.NotEqual(
		t,
		resolvedCertificates.ApprovedCertificates.Certs[0].SerialNumber,
		resolvedCertificates.ApprovedCertificates.Certs[1].SerialNumber,
	)

	// Check indexes for certificate2
	utils.CheckCertificateStateIndexes(t, setup, intermediateCertificate, indexes)
}

func TestHandler_AddDaIntermediateCert_ForTree(t *testing.T) {
	setup := utils.Setup(t)

	// add root x509 certificate
	testRootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, testRootCertificate)

	// add intermediate x509 certificate
	testIntermediateCertificate := utils.IntermediateDaCertificate(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, testIntermediateCertificate)

	// add leaf x509 certificate
	testLeafCertificate := utils.LeafCertificate(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, testLeafCertificate)

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

func TestHandler_AddDaIntermediateCert_ByNotOwnerButSameVendor(t *testing.T) {
	setup := utils.Setup(t)

	// add two vendors with the same VID
	vendorAccAddress1 := setup.CreateVendorAccount(testconstants.Vid)
	vendorAccAddress2 := setup.CreateVendorAccount(testconstants.Vid)

	// store root certificate
	testRootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, testRootCertificate)

	// Store an intermediate certificate with the first vendor account as the owner
	intermediateCertificate := utils.IntermediateDaCertificate(vendorAccAddress1)
	intermediateCertificate.SerialNumber = utils.SerialNumber
	utils.AddMokedDaCertificate(setup, intermediateCertificate)

	// add an intermediate certificate with the same subject and SKID by second vendor account
	testIntermediateCertificate := utils.IntermediateDaCertificate(vendorAccAddress2)
	utils.AddDaIntermediateCertificate(setup, testIntermediateCertificate)

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

func TestHandler_AddDaIntermediateCert_VIDScopedRoot(t *testing.T) {
	setup := utils.Setup(t)

	accAddress := setup.CreateVendorAccount(testconstants.PAACertWithNumericVidVid)

	// store root certificate
	rootCert := utils.RootDaCertificateWithNumericVid(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// add x509 certificate
	testIntermediateCertificate := utils.IntermediateDaCertificateWithNumericPidVid(accAddress)
	utils.AddDaIntermediateCertificate(setup, testIntermediateCertificate)

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

func TestHandler_AddDaIntermediateCert_NonVIDScopedRoot(t *testing.T) {
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
			addX509Cert := types.NewMsgAddX509Cert(accAddress.String(), tc.childCert, testconstants.CertSchemaVersion)
			_, err := setup.Handler(setup.Ctx, addX509Cert)
			require.NoError(setup.T, err)

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

func TestHandler_AddDaIntermediateCert_ForInvalidCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// add x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Vendor1.String(), testconstants.StubCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrInvalidCertificate)
}

func TestHandler_AddDaIntermediateCert_ForRootCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// add root certificate as leaf x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Vendor1.String(), testconstants.RootCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrNonRootCertificateSelfSigned)
}

func TestHandler_AddDaIntermediateCert_ForDuplicate(t *testing.T) {
	setup := utils.Setup(t)

	// store root certificate
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	setup.Keeper.AddAllCertificate(setup.Ctx, rootCertificate)

	// store intermediate certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Vendor1.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// store intermediate certificate second time
	_, err = setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrCertificateAlreadyExists)
}

func TestHandler_AddDaIntermediateCert_RootIsNoc(t *testing.T) {
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

func TestHandler_AddDaIntermediateCert_ForAbsentDirectParentCert(t *testing.T) {
	setup := utils.Setup(t)

	// add intermediate x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Vendor1.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrCertificateDoesNotExist)
}

func TestHandler_AddDaIntermediateCert_ByOtherVendor(t *testing.T) {
	setup := utils.Setup(t)

	// Add vendor account
	vendorAccAddress := setup.CreateVendorAccount(testconstants.RootCertWithVidVid)

	// propose and approve x509 root certificate
	rootCert := utils.RootDaCertificateWithSameSubjectAndSKID1(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// Add intermediate certificate
	testIntermediateCertificate1 := utils.IntermediateDaCertificateWithSameSubjectAndSKID1(vendorAccAddress)
	utils.AddDaIntermediateCertificate(setup, testIntermediateCertificate1)

	// add second vendor account with VID = 1000
	vendorAccAddress2 := setup.CreateVendorAdminAccount(testconstants.VendorID1)

	// add second intermediate certificates with same Subject/SKID
	testIntermediateCertificate2 := utils.IntermediateDaCertificateWithSameSubjectAndSKID2(vendorAccAddress2)
	addX509Cert := types.NewMsgAddX509Cert(
		vendorAccAddress2.String(),
		testIntermediateCertificate2.PemCert,
		testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_AddDaIntermediateCert_SenderNotVendor(t *testing.T) {
	setup := utils.Setup(t)

	// store root certificate
	rootCert := utils.RootDaCertificateWithVid(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// add x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateCertWithVid1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_AddDaIntermediateCert_VIDScopedRoot_NegativeCases(t *testing.T) {
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

func TestHandler_AddDaIntermediateCert_NonVIDScopedRoot_NegativeCases(t *testing.T) {
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
