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

func TestHandler_RevokeNoRootCert(t *testing.T) {
	setup := utils.Setup(t)

	// add the first NOC root certificate
	rootCertificate := utils.NocRootCert1(setup.Vendor1)
	utils.AddNocRootCertificate(setup, setup.Vendor1, rootCertificate.PemCert)

	// Revoke NOC root with subject and subject key id only
	utils.RevokeNocRootCertificate(
		setup,
		setup.Vendor1,
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		"",
		false,
	)

	// Check indexes
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.RevokedNocRootCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
			{Key: types.NocRootCertificatesKeyPrefix},
			{Key: types.NocIcaCertificatesKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.RevokedNocIcaCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)
}

func TestHandler_RevokeNocX509RootCert_TwoCerts(t *testing.T) {
	setup := utils.Setup(t)

	// add the first NOC root certificate
	rootCertificate := utils.NocRootCert1(setup.Vendor1)
	utils.AddNocRootCertificate(setup, setup.Vendor1, rootCertificate.PemCert)

	// add the second NOC root certificate
	rootCertificate2 := utils.NocRootCert1Copy(setup.Vendor1)
	utils.AddNocRootCertificate(setup, setup.Vendor1, rootCertificate2.PemCert)

	// add the first NOC non-root certificate
	icaCertificate := utils.NocCertIca1(setup.Vendor1)
	utils.AddNocIntermediateCertificate(setup, setup.Vendor1, icaCertificate.PemCert)

	// Revoke NOC root with subject and subject key id only
	utils.RevokeNocRootCertificate(
		setup,
		setup.Vendor1,
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		"",
		false,
	)

	// Check indexes
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.RevokedNocRootCertificatesKeyPrefix, Count: 2},
			{Key: types.NocIcaCertificatesKeyPrefix}, // inter exists
		},
		Missing: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
			{Key: types.NocRootCertificatesKeyPrefix},
			{Key: types.RevokedNocIcaCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)

	// Check indexes
	indexes = utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
			{Key: types.NocIcaCertificatesKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.NocRootCertificatesKeyPrefix},
			{Key: types.RevokedNocRootCertificatesKeyPrefix},
			{Key: types.RevokedNocIcaCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate, indexes)
}

func TestHandler_RevokeNocX509RootCert_RevokeWithChild(t *testing.T) {
	setup := utils.Setup(t)

	// add the first NOC root certificate
	rootCertificate := utils.NocRootCert1(setup.Vendor1)
	utils.AddNocRootCertificate(setup, setup.Vendor1, rootCertificate.PemCert)

	// add the second NOC root certificate
	rootCertificate2 := utils.NocRootCert1Copy(setup.Vendor1)
	utils.AddNocRootCertificate(setup, setup.Vendor1, rootCertificate2.PemCert)

	// add the first NOC non-root certificate
	icaCertificate := utils.NocCertIca1(setup.Vendor1)
	utils.AddNocIntermediateCertificate(setup, setup.Vendor1, icaCertificate.PemCert)

	// Revoke NOC root with subject and subject key id only
	utils.RevokeNocRootCertificate(
		setup,
		setup.Vendor1,
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		"",
		true,
	)

	// Check indexes
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.RevokedNocRootCertificatesKeyPrefix, Count: 2},
		},
		Missing: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
			{Key: types.NocRootCertificatesKeyPrefix},
			{Key: types.RevokedNocIcaCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
			{Key: types.NocIcaCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)

	// Check indexes
	indexes = utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.RevokedNocIcaCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
			{Key: types.NocRootCertificatesKeyPrefix},
			{Key: types.RevokedNocRootCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
			{Key: types.NocIcaCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate, indexes)
}

func TestHandler_RevokeNocX509RootCert_RevokeWithSerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	// add the first NOC root certificate
	rootCertificate := utils.NocRootCert1(setup.Vendor1)
	utils.AddNocRootCertificate(setup, setup.Vendor1, rootCertificate.PemCert)

	// add the second NOC root certificate
	rootCertificate2 := utils.NocRootCert1Copy(setup.Vendor1)
	utils.AddNocRootCertificate(setup, setup.Vendor1, rootCertificate2.PemCert)

	// add the first NOC non-root certificate
	icaCertificate := utils.NocCertIca1(setup.Vendor1)
	utils.AddNocIntermediateCertificate(setup, setup.Vendor1, icaCertificate.PemCert)

	// Revoke NOC root with subject and subject key id by serial number
	utils.RevokeNocRootCertificate(
		setup,
		setup.Vendor1,
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		rootCertificate.SerialNumber,
		false,
	)

	// Check indexes
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
			{Key: types.RevokedNocRootCertificatesKeyPrefix, Count: 1},
			{Key: types.NocIcaCertificatesKeyPrefix}, // inter exists
		},
		Missing: []utils.TestIndex{
			{Key: types.RevokedNocIcaCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate2, indexes)

	// Check indexes
	indexes = utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
			{Key: types.NocIcaCertificatesKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.NocRootCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.RevokedNocRootCertificatesKeyPrefix},
			{Key: types.RevokedNocIcaCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate, indexes)
}

func TestHandler_RevokeNocX509RootCert_RevokeWithSerialNumberAndChild(t *testing.T) {
	setup := utils.Setup(t)

	// add the first NOC root certificate
	rootCertificate := utils.NocRootCert1(setup.Vendor1)
	utils.AddNocRootCertificate(setup, setup.Vendor1, rootCertificate.PemCert)

	// add the second NOC root certificate
	rootCertificate2 := utils.NocRootCert1Copy(setup.Vendor1)
	utils.AddNocRootCertificate(setup, setup.Vendor1, rootCertificate2.PemCert)

	// add the first NOC non-root certificate
	icaCertificate := utils.NocCertIca1(setup.Vendor1)
	utils.AddNocIntermediateCertificate(setup, setup.Vendor1, icaCertificate.PemCert)

	// Revoke NOC root with subject and subject key id by serial number
	utils.RevokeNocRootCertificate(
		setup,
		setup.Vendor1,
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		rootCertificate.SerialNumber,
		true,
	)

	// Check indexes
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
			{Key: types.RevokedNocRootCertificatesKeyPrefix, Count: 1},
		},
		Missing: []utils.TestIndex{
			{Key: types.RevokedNocIcaCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
			{Key: types.NocIcaCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate2, indexes)

	// Check indexes
	indexes = utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.RevokedNocIcaCertificatesKeyPrefix},
			{Key: types.NocRootCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
			{Key: types.RevokedNocRootCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
			{Key: types.NocIcaCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate, indexes)
}

// Extra cases

// Error cases

func TestHandler_RevokeNocX509RootCert_SenderNotVendor(t *testing.T) {
	setup := utils.Setup(t)

	// add the new NOC root certificate
	addNocX509RootCert := types.NewMsgAddNocX509RootCert(setup.Vendor1.String(), testconstants.NocRootCert1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
	require.NoError(t, err)

	revokeCert := types.NewMsgRevokeNocX509RootCert(
		setup.Trustee1.String(),
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		testconstants.NocRootCert1SerialNumber,
		"",
		false,
	)
	_, err = setup.Handler(setup.Ctx, revokeCert)

	require.Error(t, err)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_RevokeNocX509RootCert_CertificateDoesNotExist(t *testing.T) {
	setup := utils.Setup(t)

	revokeCert := types.NewMsgRevokeNocX509RootCert(
		setup.Vendor1.String(),
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		testconstants.NocRootCert1SerialNumber,
		"",
		false,
	)
	_, err := setup.Handler(setup.Ctx, revokeCert)

	require.Error(t, err)
	require.ErrorIs(t, err, pkitypes.ErrCertificateDoesNotExist)
}

func TestHandler_RevokeNocX509RootCert_CertificateExists(t *testing.T) {
	accAddress := utils.GenerateAccAddress()

	cases := []struct {
		name         string
		existingCert *types.Certificate
		nocRoorCert  string
		err          error
	}{
		{
			name: "ExistingNonRootCert",
			existingCert: &types.Certificate{
				Issuer:          testconstants.NocRootCert1Subject,
				Subject:         testconstants.NocRootCert1Subject,
				SubjectAsText:   testconstants.NocRootCert1SubjectAsText,
				SubjectKeyId:    testconstants.NocRootCert1SubjectKeyID,
				SerialNumber:    testconstants.NocRootCert1SerialNumber,
				IsRoot:          false,
				CertificateType: types.CertificateType_OperationalPKI,
				Vid:             testconstants.Vid,
			},
			nocRoorCert: testconstants.RootCertPem,
			err:         pkitypes.ErrInappropriateCertificateType,
		},
		{
			name: "ExistingNotNocCert",
			existingCert: &types.Certificate{
				Issuer:          testconstants.NocRootCert1Subject,
				Subject:         testconstants.NocRootCert1Subject,
				SubjectAsText:   testconstants.NocRootCert1SubjectAsText,
				SubjectKeyId:    testconstants.NocRootCert1SubjectKeyID,
				SerialNumber:    testconstants.NocRootCert1SerialNumber,
				IsRoot:          true,
				CertificateType: types.CertificateType_DeviceAttestationPKI,
				Vid:             testconstants.Vid,
			},
			nocRoorCert: testconstants.RootCertPem,
			err:         pkitypes.ErrInappropriateCertificateType,
		},
		{
			name: "ExistingCertWithDifferentVid",
			existingCert: &types.Certificate{
				Issuer:          testconstants.NocRootCert1Subject,
				Subject:         testconstants.NocRootCert1Subject,
				SubjectAsText:   testconstants.NocRootCert1SubjectAsText,
				SubjectKeyId:    testconstants.NocRootCert1SubjectKeyID,
				SerialNumber:    testconstants.NocRootCert1SerialNumber,
				IsRoot:          true,
				CertificateType: types.CertificateType_OperationalPKI,
				Vid:             testconstants.VendorID1,
			},
			nocRoorCert: testconstants.RootCertPem,
			err:         pkitypes.ErrCertVidNotEqualAccountVid,
		},
		{
			name: "ExistingCertWithDifferentSerialNumber",
			existingCert: &types.Certificate{
				Issuer:          testconstants.NocRootCert1Subject,
				Subject:         testconstants.NocRootCert1Subject,
				SubjectAsText:   testconstants.NocRootCert1SubjectAsText,
				SubjectKeyId:    testconstants.NocRootCert1SubjectKeyID,
				SerialNumber:    "1234567",
				IsRoot:          true,
				CertificateType: types.CertificateType_OperationalPKI,
				Vid:             testconstants.Vid,
			},
			nocRoorCert: testconstants.RootCertPem,
			err:         pkitypes.ErrCertificateDoesNotExist,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)
			setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

			// add the existing certificate
			setup.Keeper.AddNocCertificate(setup.Ctx, *tc.existingCert)
			uniqueCertificate := types.UniqueCertificate{
				Issuer:       tc.existingCert.Issuer,
				SerialNumber: tc.existingCert.SerialNumber,
				Present:      true,
			}
			setup.Keeper.SetUniqueCertificate(setup.Ctx, uniqueCertificate)

			revokeCert := types.NewMsgRevokeNocX509RootCert(
				accAddress.String(),
				testconstants.NocRootCert1Subject,
				testconstants.NocRootCert1SubjectKeyID,
				testconstants.NocRootCert1SerialNumber,
				"",
				false,
			)
			_, err := setup.Handler(setup.Ctx, revokeCert)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
