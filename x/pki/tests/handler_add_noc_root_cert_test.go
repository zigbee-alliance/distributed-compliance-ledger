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
	cases := []CertificateTestCase{
		{
			name:    "OperationalPKI_AddNocRootCert",
			crtType: types.CertificateType_OperationalPKI,
		},
		{
			name:    "VIDSignerPKI_AddNocRootCert",
			crtType: types.CertificateType_VIDSignerPKI,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			// add NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, tc.crtType)
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
		})
	}
}

func TestHandler_AddNocRootCert_SameSubjectAndSkid_DifferentSerialNumber(t *testing.T) {
	cases := []CertificateTestCase{
		{
			name:    "OperationalPKI_AddNocRootCert_SameSubjectAndSkid_DifferentSerialNumber",
			crtType: types.CertificateType_OperationalPKI,
		},
		{
			name:    "VIDSignerPKI_AddNocRootCert_SameSubjectAndSkid_DifferentSerialNumber",
			crtType: types.CertificateType_VIDSignerPKI,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			// Store the NOC root certificate
			rootCertificate1 := utils.RootNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate1)

			// add second NOC root certificate
			rootCertificate2 := utils.RootNocCertificate1Copy(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate2)

			// Check state indexes
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
		})
	}
}

func TestHandler_AddNocRootCert_ByNotOwnerButSameVendor(t *testing.T) {
	cases := []CertificateTestCase{
		{
			name:    "OperationalPKI_AddNocRootCert_ByNotOwnerButSameVendor",
			crtType: types.CertificateType_OperationalPKI,
		},
		{
			name:    "VIDSignerPKI_AddNocRootCert_ByNotOwnerButSameVendor",
			crtType: types.CertificateType_VIDSignerPKI,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			// add two vendors with the same VID
			vendorAccAddress1 := setup.CreateVendorAccount(testconstants.Vid)
			vendorAccAddress2 := setup.CreateVendorAccount(testconstants.Vid)

			// add NOC root certificate
			rootCertificate1 := utils.RootNocCertificate1(vendorAccAddress1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate1)

			// add second NOC root certificate by other vendor
			rootCertificate2 := utils.RootNocCertificate1Copy(vendorAccAddress2, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate2)

			// Check state indexes
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
		})
	}
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
			nocRoorCert: testconstants.IntermediateCertPem,
			err:         pkitypes.ErrRootCertificateIsNotSelfSigned,
		},
		{
			name:        "ExpiredCertificate",
			accountVid:  testconstants.Vid,
			accountRole: dclauthtypes.Vendor,
			nocRoorCert: testconstants.PAACertExpired,
			err:         pkitypes.ErrInvalidCertificate,
		},
	}

	crtCases := []CertificateTestCase{
		{
			name:                    "OperationalPKI",
			isVidVerificationSigner: false,
		},
		{
			name:                    "VIDSignerPKI",
			isVidVerificationSigner: true,
		},
	}

	for _, tc := range cases {
		for _, tcc := range crtCases {
			t.Run(tcc.name+"_"+tc.name, func(t *testing.T) {
				setup := utils.Setup(t)
				setup.AddAccount(accAddress, []dclauthtypes.AccountRole{tc.accountRole}, tc.accountVid)

				addNocX509RootCert := types.NewMsgAddNocX509RootCert(accAddress.String(), tc.nocRoorCert, testconstants.CertSchemaVersion, tcc.isVidVerificationSigner)
				_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
				require.ErrorIs(t, err, tc.err)
			})
		}
	}
}

func TestHandler_AddNocRootCert_CertificateExist(t *testing.T) {
	accAddress := utils.GenerateAccAddress()

	cases := []struct {
		name         string
		existingCert *types.Certificate
		nocRoorCert  string
		err          error
	}{
		{
			name: "Duplicate",
			existingCert: &types.Certificate{
				Issuer:        testconstants.RootIssuer,
				Subject:       testconstants.RootSubject,
				SubjectAsText: testconstants.RootSubjectAsText,
				SubjectKeyId:  testconstants.RootSubjectKeyID,
				SerialNumber:  testconstants.RootSerialNumber,
				IsRoot:        true,
				Vid:           testconstants.Vid,
			},
			nocRoorCert: testconstants.RootCertPem,
			err:         pkitypes.ErrCertificateAlreadyExists,
		},
		{
			name: "ExistingNonRootCert",
			existingCert: &types.Certificate{
				Issuer:        testconstants.TestIssuer,
				Subject:       testconstants.RootSubject,
				SubjectAsText: testconstants.RootSubjectAsText,
				SubjectKeyId:  testconstants.RootSubjectKeyID,
				SerialNumber:  testconstants.TestSerialNumber,
				IsRoot:        false,
				Vid:           testconstants.Vid,
			},
			nocRoorCert: testconstants.RootCertPem,
			err:         sdkerrors.ErrUnauthorized,
		},
		{
			name: "ExistingNotNocCert",
			existingCert: &types.Certificate{
				Issuer:        testconstants.RootIssuer,
				Subject:       testconstants.RootSubject,
				SubjectAsText: testconstants.RootSubjectAsText,
				SubjectKeyId:  testconstants.RootSubjectKeyID,
				SerialNumber:  testconstants.TestSerialNumber,
				IsRoot:        true,
				Vid:           testconstants.Vid,
			},
			nocRoorCert: testconstants.RootCertPem,
			err:         pkitypes.ErrInappropriateCertificateType,
		},
		{
			name: "ExistingCertWithDifferentVid",
			existingCert: &types.Certificate{
				Issuer:        testconstants.RootIssuer,
				Subject:       testconstants.RootSubject,
				SubjectAsText: testconstants.RootSubjectAsText,
				SubjectKeyId:  testconstants.RootSubjectKeyID,
				SerialNumber:  testconstants.GoogleSerialNumber,
				IsRoot:        true,
				Vid:           testconstants.VendorID1,
			},
			nocRoorCert: testconstants.RootCertPem,
			err:         sdkerrors.ErrUnauthorized,
		},
	}

	crtCases := []CertificateTestCase{
		{
			name:                    "OperationalPKI",
			crtType:                 types.CertificateType_OperationalPKI,
			isVidVerificationSigner: false,
		},
		{
			name:                    "VIDSignerPKI",
			crtType:                 types.CertificateType_VIDSignerPKI,
			isVidVerificationSigner: true,
		},
	}

	for _, tc := range cases {
		for _, tcc := range crtCases {
			t.Run(tcc.name+"_"+tc.name, func(t *testing.T) {
				setup := utils.Setup(t)
				setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

				existingCert := *tc.existingCert

				// the test for this error requires different types
				if errors.Is(tc.err, pkitypes.ErrInappropriateCertificateType) {
					existingCert.CertificateType = types.CertificateType_DeviceAttestationPKI
				} else {
					existingCert.CertificateType = tcc.crtType
				}

				// add the existing certificate
				setup.Keeper.AddAllCertificate(setup.Ctx, existingCert)
				uniqueCertificate := types.UniqueCertificate{
					Issuer:       existingCert.Issuer,
					SerialNumber: existingCert.SerialNumber,
					Present:      true,
				}
				setup.Keeper.SetUniqueCertificate(setup.Ctx, uniqueCertificate)

				addNocX509RootCert := types.NewMsgAddNocX509RootCert(accAddress.String(), tc.nocRoorCert, testconstants.CertSchemaVersion, tcc.isVidVerificationSigner)
				_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
				require.ErrorIs(t, err, tc.err)
			})
		}
	}
}
