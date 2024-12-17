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

func TestHandler_AddNocIntermediateCert(t *testing.T) {
	setup := utils.Setup(t)

	// add NOC root certificate
	rootCertificate := utils.RootNocCertificate1(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate)

	// add NOC ICA certificate
	icaCertificate := utils.IntermediateNocCertificate1(setup.Vendor1)
	utils.AddNocIntermediateCertificate(setup, icaCertificate)

	// Check state indexes
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
			{Key: types.NocRootCertificatesKeyPrefix, Count: 1}, // we created root certificate with same vid
			{Key: types.NocIcaCertificatesKeyPrefix},
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate, indexes)
}

func TestHandler_AddNocIntermediateCert_SameSubjectAndSkid_DifferentSerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	// add NOC root certificate
	rootCertificate := utils.RootNocCertificate1(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate)

	// Store the NOC certificate with different serial number
	intermediateCertificate := utils.IntermediateNocCertificate1(setup.Vendor1)
	intermediateCertificate.SerialNumber = testconstants.TestSerialNumber
	utils.AddMokedNocCertificate(setup, intermediateCertificate)

	// add the new NOC certificate
	icaCertificate := utils.IntermediateNocCertificate1(setup.Vendor1)
	utils.AddNocIntermediateCertificate(setup, icaCertificate)

	// Check state indexes
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix, Count: 2},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.NocCertificatesKeyPrefix, Count: 2},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix, Count: 2},
			{Key: types.NocRootCertificatesKeyPrefix, Count: 1}, // we create root certificate as well but ica should not be there
			{Key: types.NocIcaCertificatesKeyPrefix, Count: 2},
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate, indexes)
}

func TestHandler_AddNocIntermediateCert_ByNotOwnerButSameVendor(t *testing.T) {
	setup := utils.Setup(t)

	// add two vendors with the same VID
	vendorAccAddress1 := setup.CreateVendorAccount(testconstants.Vid)
	vendorAccAddress2 := setup.CreateVendorAccount(testconstants.Vid)

	// add NOC root certificate
	rootCertificate := utils.RootNocCertificate1(vendorAccAddress1)
	utils.AddNocRootCertificate(setup, rootCertificate)

	// add the new NOC certificate by first vendor
	icaCertificate := utils.IntermediateNocCertificate1(vendorAccAddress1)
	utils.AddNocIntermediateCertificate(setup, icaCertificate)

	// add the new NOC certificate by second vendor
	icaCertificate2 := utils.IntermediateNocCertificate1Copy(vendorAccAddress2)
	utils.AddNocIntermediateCertificate(setup, icaCertificate2)

	// Check state indexes
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix, Count: 2},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.NocCertificatesKeyPrefix, Count: 2},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix, Count: 2},
			{Key: types.NocRootCertificatesKeyPrefix, Count: 1}, // we create root certificate as well but ica should not be there
			{Key: types.NocIcaCertificatesKeyPrefix, Count: 2},
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate, indexes)
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate2, indexes)
}

// Error cases

func TestHandler_AddNocIntermediateCert_SenderNotVendor(t *testing.T) {
	setup := utils.Setup(t)

	// add NOC root certificate
	rootCertificate := utils.RootNocCertificate1(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate)

	addNocX509Cert := types.NewMsgAddNocX509IcaCert(setup.Trustee1.String(), testconstants.NocCert1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addNocX509Cert)

	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_AddNocIntermediateCert_Root_VID_Does_Not_Equal_To_AccountVID(t *testing.T) {
	setup := utils.Setup(t)

	// add NOC root certificate
	rootCertificate := utils.RootNocCertificate1(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate)

	newAccAddress := setup.CreateVendorAccount(1111)

	// try to add NOC certificate
	nocX509Cert := types.NewMsgAddNocX509IcaCert(newAccAddress.String(), testconstants.NocCert1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, nocX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrCertVidNotEqualAccountVid)
}

func TestHandler_AddNocIntermediateCert_ForInvalidCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// add x509 certificate
	addX509Cert := types.NewMsgAddNocX509IcaCert(setup.Vendor1.String(), testconstants.StubCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrInvalidCertificate)
}

func TestHandler_AddNocIntermediateCert_ForNocRootCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// try to add root certificate x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Vendor1.String(), testconstants.NocRootCert1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrNonRootCertificateSelfSigned)
}

func TestHandler_AddNocIntermediateCert_ForRootNonNocCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// store root certificate
	rootCert := utils.RootDaCertificateWithVid(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// try to add root certificate x509 certificate
	addX509Cert := types.NewMsgAddNocX509IcaCert(setup.Vendor1.String(), testconstants.IntermediateCertWithVid1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrInappropriateCertificateType)
}

func TestHandler_AddNocIntermediateCert_WhenNocRootCertIsAbsent(t *testing.T) {
	setup := utils.Setup(t)

	// add the new NOC certificate
	addNocX509Cert := types.NewMsgAddNocX509IcaCert(setup.Vendor1.String(), testconstants.NocCert1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addNocX509Cert)

	require.ErrorIs(t, err, pkitypes.ErrCertificateDoesNotExist)
}

func TestHandler_AddNocIntermediateCert_CertificateExist(t *testing.T) {
	accAddress := utils.GenerateAccAddress()

	cases := []struct {
		name         string
		existingCert *types.Certificate
		nocCert      string
		err          error
	}{
		{
			name: "Duplicate",
			existingCert: &types.Certificate{
				Issuer:          testconstants.NocRootCert1Subject,
				AuthorityKeyId:  testconstants.NocRootCert1SubjectKeyID,
				Subject:         testconstants.NocCert1Subject,
				SubjectAsText:   testconstants.NocCert1SubjectAsText,
				SubjectKeyId:    testconstants.NocCert1SubjectKeyID,
				SerialNumber:    testconstants.NocCert1SerialNumber,
				IsRoot:          false,
				CertificateType: types.CertificateType_OperationalPKI,
				Vid:             testconstants.Vid,
			},
			nocCert: testconstants.NocCert1,
			err:     pkitypes.ErrCertificateAlreadyExists,
		},
		{
			name: "ExistingIsRootCert",
			existingCert: &types.Certificate{
				Issuer:          testconstants.NocRootCert1Subject,
				AuthorityKeyId:  testconstants.NocRootCert1SubjectKeyID,
				Subject:         testconstants.NocCert1Subject,
				SubjectAsText:   testconstants.NocCert1SubjectAsText,
				SubjectKeyId:    testconstants.NocCert1SubjectKeyID,
				SerialNumber:    testconstants.NocRootCert1SerialNumber,
				IsRoot:          true,
				CertificateType: types.CertificateType_OperationalPKI,
				Vid:             testconstants.Vid,
			},
			nocCert: testconstants.NocCert1,
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			name: "ExistingWithDifferentIssuer",
			existingCert: &types.Certificate{
				Issuer:          testconstants.RootIssuer,
				AuthorityKeyId:  testconstants.NocRootCert1SubjectKeyID,
				Subject:         testconstants.NocCert1Subject,
				SubjectAsText:   testconstants.NocCert1SubjectAsText,
				SubjectKeyId:    testconstants.NocCert1SubjectKeyID,
				SerialNumber:    "1234",
				IsRoot:          false,
				CertificateType: types.CertificateType_OperationalPKI,
				Vid:             testconstants.Vid,
			},
			nocCert: testconstants.NocCert1,
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			name: "ExistingWithDifferentAuthorityKeyId",
			existingCert: &types.Certificate{
				Issuer:          testconstants.NocRootCert1Subject,
				AuthorityKeyId:  testconstants.RootSubjectKeyID,
				Subject:         testconstants.NocCert1Subject,
				SubjectAsText:   testconstants.NocCert1SubjectAsText,
				SubjectKeyId:    testconstants.NocCert1SubjectKeyID,
				SerialNumber:    "1234",
				IsRoot:          false,
				CertificateType: types.CertificateType_OperationalPKI,
				Vid:             testconstants.Vid,
			},
			nocCert: testconstants.NocCert1,
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			name: "ExistingNotNocCert",
			existingCert: &types.Certificate{
				Issuer:          testconstants.NocRootCert1Subject,
				AuthorityKeyId:  testconstants.NocRootCert1SubjectKeyID,
				Subject:         testconstants.NocCert1Subject,
				SubjectAsText:   testconstants.NocCert1SubjectAsText,
				SubjectKeyId:    testconstants.NocCert1SubjectKeyID,
				SerialNumber:    "1234",
				IsRoot:          false,
				CertificateType: types.CertificateType_DeviceAttestationPKI,
				Vid:             testconstants.Vid,
			},
			nocCert: testconstants.NocCert1,
			err:     pkitypes.ErrInappropriateCertificateType,
		},
		{
			name: "ExistingCertWithDifferentVid",
			existingCert: &types.Certificate{
				Issuer:          testconstants.NocRootCert1Subject,
				AuthorityKeyId:  testconstants.NocRootCert1SubjectKeyID,
				Subject:         testconstants.NocCert1Subject,
				SubjectAsText:   testconstants.NocCert1SubjectAsText,
				SubjectKeyId:    testconstants.NocCert1SubjectKeyID,
				SerialNumber:    "1234",
				IsRoot:          false,
				CertificateType: types.CertificateType_OperationalPKI,
				Vid:             testconstants.VendorID1,
			},
			nocCert: testconstants.NocCert1,
			err:     sdkerrors.ErrUnauthorized,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)
			vid := testconstants.Vid
			setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

			// add NOC root certificate
			rootCertificate := utils.RootNocCertificate1(accAddress)
			utils.AddNocRootCertificate(setup, rootCertificate)

			// add the existing certificate
			setup.Keeper.AddAllCertificate(setup.Ctx, *tc.existingCert)
			uniqueCertificate := types.UniqueCertificate{
				Issuer:       tc.existingCert.Issuer,
				SerialNumber: tc.existingCert.SerialNumber,
				Present:      true,
			}
			setup.Keeper.SetUniqueCertificate(setup.Ctx, uniqueCertificate)

			addNocX509Cert := types.NewMsgAddNocX509IcaCert(accAddress.String(), tc.nocCert, testconstants.CertSchemaVersion)
			_, err := setup.Handler(setup.Ctx, addNocX509Cert)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
