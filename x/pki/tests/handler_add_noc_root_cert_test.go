package tests

import (
	"errors"
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

func TestHandler_AddNocRootCert(t *testing.T) {
	setup := utils.Setup(t)

	// add NOC root certificate
	rootCertificate := utils.RootNocCertificate1(setup.Vendor1, types.CertificateType_OperationalPKI)
	utils.AddNocRootCertificate(setup, rootCertificate)

	// check state indexes
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
			{Key: types.NocRootCertificatesKeyPrefix},
			{Key: types.UniqueCertificateKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.NocIcaCertificatesKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)
}

func TestHandler_AddNocRootCert_VVSC(t *testing.T) {
	setup := utils.Setup(t)

	rootCertificate := utils.RootVvscCertificate1(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate)

	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
			{Key: types.NocRootCertificatesKeyPrefix},
			{Key: types.UniqueCertificateKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.NocIcaCertificatesKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)
}

func TestHandler_AddNocRootCert_SameSubjectAndSkid_DifferentSerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	rootCertificate1 := utils.RootNocCertificate1(setup.Vendor1, types.CertificateType_OperationalPKI)
	utils.AddNocRootCertificate(setup, rootCertificate1)

	rootCertificate2 := utils.RootNocCertificate1Copy(setup.Vendor1, types.CertificateType_OperationalPKI)
	utils.AddNocRootCertificate(setup, rootCertificate2)

	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix, Count: 2},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.NocCertificatesKeyPrefix, Count: 2},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.NocRootCertificatesKeyPrefix, Count: 2},
			{Key: types.UniqueCertificateKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.NocIcaCertificatesKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate1, indexes)
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate2, indexes)
}

func TestHandler_AddNocRootCert_SameSubjectAndSkid_DifferentSerialNumber_VVSC(t *testing.T) {
	setup := utils.Setup(t)

	rootCertificate1 := utils.RootVvscCertificate1(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate1)

	rootCertificate2 := utils.RootVvscCertificate1Copy(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate2)

	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix, Count: 2},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.NocCertificatesKeyPrefix, Count: 2},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.NocRootCertificatesKeyPrefix, Count: 2},
			{Key: types.UniqueCertificateKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.NocIcaCertificatesKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate1, indexes)
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate2, indexes)
}

func TestHandler_AddNocRootCert_ByNotOwnerButSameVendor(t *testing.T) {
	setup := utils.Setup(t)

	vendorAccAddress1 := setup.CreateVendorAccount(testconstants.Vid)
	vendorAccAddress2 := setup.CreateVendorAccount(testconstants.Vid)

	rootCertificate1 := utils.RootNocCertificate1(vendorAccAddress1, types.CertificateType_OperationalPKI)
	utils.AddNocRootCertificate(setup, rootCertificate1)

	rootCertificate2 := utils.RootNocCertificate1Copy(vendorAccAddress2, types.CertificateType_OperationalPKI)
	utils.AddNocRootCertificate(setup, rootCertificate2)

	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix, Count: 2},
			{Key: types.AllCertificatesBySubjectKeyPrefix, Count: 1},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.NocCertificatesKeyPrefix, Count: 2},
			{Key: types.NocCertificatesBySubjectKeyPrefix, Count: 1},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.NocRootCertificatesKeyPrefix, Count: 2},
			{Key: types.UniqueCertificateKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.NocIcaCertificatesKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate1, indexes)
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate2, indexes)
}

func TestHandler_AddNocRootCert_ByNotOwnerButSameVendor_VVSC(t *testing.T) {
	setup := utils.Setup(t)

	vendorAccAddress1 := setup.CreateVendorAccount(testconstants.Vid)
	vendorAccAddress2 := setup.CreateVendorAccount(testconstants.Vid)

	rootCertificate1 := utils.RootVvscCertificate1(vendorAccAddress1)
	utils.AddNocRootCertificate(setup, rootCertificate1)

	rootCertificate2 := utils.RootVvscCertificate1Copy(vendorAccAddress2)
	utils.AddNocRootCertificate(setup, rootCertificate2)

	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix, Count: 2},
			{Key: types.AllCertificatesBySubjectKeyPrefix, Count: 1},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.NocCertificatesKeyPrefix, Count: 2},
			{Key: types.NocCertificatesBySubjectKeyPrefix, Count: 1},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.NocRootCertificatesKeyPrefix, Count: 2},
			{Key: types.UniqueCertificateKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.NocIcaCertificatesKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate1, indexes)
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate2, indexes)
}

// Error cases

func TestHandler_AddNocRootCert_SenderNotVendor(t *testing.T) {
	cases := []CertificateTestCase{
		{
			name:                    "OperationalPKI_AddNocRootCert_SenderNotVendor",
			isVidVerificationSigner: false,
		},
		{
			name:                    "VIDSignerPKI_AddNocRootCert_SenderNotVendor",
			isVidVerificationSigner: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			addNocX509RootCert := types.NewMsgAddNocX509RootCert(
				setup.Trustee1.String(),
				testconstants.RootCertPem,
				testconstants.CertSchemaVersion, tc.isVidVerificationSigner)
			_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
			require.Error(t, err)
			require.True(t, sdkerrors.ErrUnauthorized.Is(err))
		})
	}
}

func TestHandler_AddNocRootCert_InvalidCertificate(t *testing.T) {
	accAddress := utils.GenerateAccAddress()

	cases := []struct {
		name        string
		accountVid  int32
		accountRole dclauthtypes.AccountRole
		nocRoorCert string
		err         error
	}{
		{
			name:        "NotValidPemCertificate",
			accountVid:  testconstants.Vid,
			accountRole: dclauthtypes.Vendor,
			nocRoorCert: testconstants.StubCertPem,
			err:         pkitypes.ErrInvalidCertificate,
		},
		{
			name:        "NonRootCertificate",
			accountVid:  testconstants.Vid,
			accountRole: dclauthtypes.Vendor,
			nocRoorCert: testconstants.NocCert2,
			err:         pkitypes.ErrRootCertificateIsNotSelfSigned,
		},
		{
			name:        "NonCACertificate",
			accountVid:  testconstants.Vid,
			accountRole: dclauthtypes.Vendor,
			nocRoorCert: testconstants.LeafCertPem,
			err:         pkitypes.ErrInappropriateCertificateType,
		},
		{
			name:        "ExpiredCertificate",
			accountVid:  testconstants.Vid,
			accountRole: dclauthtypes.Vendor,
			nocRoorCert: testconstants.PAACertExpired,
			err:         pkitypes.ErrInvalidCertificate,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)
			setup.AddAccount(accAddress, []dclauthtypes.AccountRole{tc.accountRole}, tc.accountVid)

			addNocX509RootCert := types.NewMsgAddNocX509RootCert(accAddress.String(), tc.nocRoorCert, testconstants.CertSchemaVersion, false)
			_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestHandler_AddNocRootCert_InvalidCertificate_VVSC(t *testing.T) {
	accAddress := utils.GenerateAccAddress()

	cases := []struct {
		name        string
		accountVid  int32
		accountRole dclauthtypes.AccountRole
		nocRoorCert string
		err         error
	}{
		{
			name:        "NotValidPemCertificate",
			accountVid:  testconstants.Vid,
			accountRole: dclauthtypes.Vendor,
			nocRoorCert: testconstants.StubCertPem,
			err:         pkitypes.ErrInvalidCertificate,
		},
		{
			name:        "NonRootCertificate",
			accountVid:  testconstants.Vid,
			accountRole: dclauthtypes.Vendor,
			nocRoorCert: testconstants.VvscIcaCert1,
			err:         pkitypes.ErrRootCertificateIsNotSelfSigned,
		},
		{
			name:        "NonVVSCProfile_CA_True",
			accountVid:  testconstants.Vid,
			accountRole: dclauthtypes.Vendor,
			nocRoorCert: testconstants.NocCert2,
			err:         pkitypes.ErrInappropriateCertificateType,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)
			setup.AddAccount(accAddress, []dclauthtypes.AccountRole{tc.accountRole}, tc.accountVid)

			addNocX509RootCert := types.NewMsgAddNocX509RootCert(accAddress.String(), tc.nocRoorCert, testconstants.CertSchemaVersion, true)
			_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestHandler_AddNocRootCert_CertificateExist(t *testing.T) {
	runCertificateExistCases(t,
		certificateExistCase{
			subject:       testconstants.NocRootCert3Subject,
			subjectAsText: testconstants.NocRootCert3SubjectAsText,
			subjectKeyID:  testconstants.NocRootCert3SubjectKeyID,
			serialNumber:  testconstants.NocRootCert3SerialNumber,
			submitPem:     testconstants.NocRootCert3,
			crtType:       types.CertificateType_OperationalPKI,
			isVVSC:        false,
		})
}

// TestHandler_AddNocRootCert_CertificateExist_VVSC re-runs the existing-row
// collision cases for the IsVidVerificationSigner=true branch. The pre-seeded
// existing cert needs the VVSC-shape SubjectKeyID and a VIDSignerPKI cert
// type; the submitted PEM is a self-signed VVSC instead of NocRootCert3 so
// the request reaches the existence checks instead of failing at
// VerifyVVSCExtensions.
func TestHandler_AddNocRootCert_CertificateExist_VVSC(t *testing.T) {
	runCertificateExistCases(t,
		certificateExistCase{
			subject:       testconstants.VvscRootCert2Subject,
			subjectAsText: testconstants.VvscRootCert2SubjectAsText,
			subjectKeyID:  testconstants.VvscRootCert2SubjectKeyID,
			serialNumber:  testconstants.VvscRootCert2SerialNumber,
			submitPem:     testconstants.VvscRootCert2,
			crtType:       types.CertificateType_VIDSignerPKI,
			isVVSC:        true,
		})
}

type certificateExistCase struct {
	subject       string
	subjectAsText string
	subjectKeyID  string
	serialNumber  string
	submitPem     string
	crtType       types.CertificateType
	isVVSC        bool
}

func runCertificateExistCases(t *testing.T, base certificateExistCase) {
	t.Helper()
	accAddress := utils.GenerateAccAddress()

	cases := []struct {
		name         string
		existingCert *types.Certificate
		err          error
	}{
		{
			name: "Duplicate",
			existingCert: &types.Certificate{
				Issuer:        base.subject,
				Subject:       base.subject,
				SubjectAsText: base.subjectAsText,
				SubjectKeyId:  base.subjectKeyID,
				SerialNumber:  base.serialNumber,
				IsRoot:        true,
				Vid:           testconstants.Vid,
			},
			err: pkitypes.ErrCertificateAlreadyExists,
		},
		{
			name: "ExistingNonRootCert",
			existingCert: &types.Certificate{
				Issuer:        testconstants.TestIssuer,
				Subject:       base.subject,
				SubjectAsText: base.subjectAsText,
				SubjectKeyId:  base.subjectKeyID,
				SerialNumber:  testconstants.TestSerialNumber,
				IsRoot:        false,
				Vid:           testconstants.Vid,
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			name: "ExistingNotNocCert",
			existingCert: &types.Certificate{
				Issuer:        base.subject,
				Subject:       base.subject,
				SubjectAsText: base.subjectAsText,
				SubjectKeyId:  base.subjectKeyID,
				SerialNumber:  testconstants.TestSerialNumber,
				IsRoot:        true,
				Vid:           testconstants.Vid,
			},
			err: pkitypes.ErrInappropriateCertificateType,
		},
		{
			name: "ExistingCertWithDifferentVid",
			existingCert: &types.Certificate{
				Issuer:        base.subject,
				Subject:       base.subject,
				SubjectAsText: base.subjectAsText,
				SubjectKeyId:  base.subjectKeyID,
				SerialNumber:  testconstants.GoogleSerialNumber,
				IsRoot:        true,
				Vid:           testconstants.VendorID1,
			},
			err: sdkerrors.ErrUnauthorized,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)
			setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

			existingCert := *tc.existingCert

			if errors.Is(tc.err, pkitypes.ErrInappropriateCertificateType) {
				existingCert.CertificateType = types.CertificateType_DeviceAttestationPKI
			} else {
				existingCert.CertificateType = base.crtType
			}

			setup.Keeper.AddAllCertificate(setup.Ctx, existingCert)
			uniqueCertificate := types.UniqueCertificate{
				Issuer:       existingCert.Issuer,
				SerialNumber: existingCert.SerialNumber,
				Present:      true,
			}
			setup.Keeper.SetUniqueCertificate(setup.Ctx, uniqueCertificate)

			addNocX509RootCert := types.NewMsgAddNocX509RootCert(accAddress.String(), base.submitPem, testconstants.CertSchemaVersion, base.isVVSC)
			_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
