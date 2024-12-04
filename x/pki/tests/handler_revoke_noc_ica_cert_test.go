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

func TestHandler_RevokeNocIntermediateCert(t *testing.T) {
	setup := utils.Setup(t)

	// add the first NOC root certificate
	rootCertificate := utils.CreateTestNocRoot1Cert()
	utils.AddNocRootCertificate(setup, setup.Vendor1, rootCertificate.PEM)

	// add the NOC non-root certificate
	icaCertificate := utils.CreateTestNocIca1Cert()
	utils.AddNocIntermediateCertificate(setup, setup.Vendor1, icaCertificate.PEM)

	// Revoke NOC with subject and subject key id only
	utils.RevokeNocIntermediateCertificate(
		setup,
		setup.Vendor1,
		icaCertificate.Subject,
		icaCertificate.SubjectKeyID,
		"",
		false)

	// Check indexes
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.RevokedNocIcaCertificatesKeyPrefix},
			{Key: types.NocRootCertificatesKeyPrefix, Count: 1}, // root still exits
		},
		Missing: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
			{Key: types.NocIcaCertificatesKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.RevokedNocRootCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate, indexes)
}

func TestHandler_RevokeNocX509Cert_RevokeDefault(t *testing.T) {
	setup := utils.Setup(t)

	// add the first NOC root certificate
	utils.AddNocRootCertificate(setup, setup.Vendor1, testconstants.NocRootCert1)

	// add the first NOC non-root certificate
	icaCertificate1 := utils.CreateTestNocIca1Cert()
	utils.AddNocIntermediateCertificate(setup, setup.Vendor1, testconstants.NocCert1)

	// add the second NOC non-root certificate
	icaCertificate2 := utils.CreateTestNocIca1CertCopy()
	utils.AddNocIntermediateCertificate(setup, setup.Vendor1, testconstants.NocCert1Copy)

	// add the NOC leaf certificate
	leafCertificate := utils.CreateTestNocLeafCert()
	utils.AddNocIntermediateCertificate(setup, setup.Vendor1, testconstants.NocLeafCert1)

	// Revoke NOC with subject and subject key id only
	utils.RevokeNocIntermediateCertificate(
		setup,
		setup.Vendor1,
		icaCertificate1.Subject,
		icaCertificate1.SubjectKeyID,
		"",
		false)

	// Check indexes
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.RevokedNocIcaCertificatesKeyPrefix, Count: 2},
			{Key: types.NocRootCertificatesKeyPrefix, Count: 1}, // root still exits
			{Key: types.NocIcaCertificatesKeyPrefix},            // leaf still exists
		},
		Missing: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.RevokedNocRootCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate1, indexes)
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate2, indexes)

	// Check indexes for leaf
	indexes = utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
			{Key: types.NocRootCertificatesKeyPrefix, Count: 1}, // root still exits
			{Key: types.NocIcaCertificatesKeyPrefix, Count: 1},  // only leaf exits
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
	utils.CheckCertificateStateIndexes(t, setup, leafCertificate, indexes)
}

func TestHandler_RevokeNocX509Cert_RevokeWithChild(t *testing.T) {
	setup := utils.Setup(t)

	// add the first NOC root certificate
	utils.AddNocRootCertificate(setup, setup.Vendor1, testconstants.NocRootCert1)

	// add the first NOC non-root certificate
	icaCertificate1 := utils.CreateTestNocIca1Cert()
	utils.AddNocIntermediateCertificate(setup, setup.Vendor1, icaCertificate1.PEM)

	// add the second NOC non-root certificate
	icaCertificate2 := utils.CreateTestNocIca1CertCopy()
	utils.AddNocIntermediateCertificate(setup, setup.Vendor1, icaCertificate2.PEM)

	// add the NOC leaf certificate
	leafCertificate := utils.CreateTestNocLeafCert()
	utils.AddNocIntermediateCertificate(setup, setup.Vendor1, leafCertificate.PEM)

	// Revoke noc with subject and subject key id and its child too
	utils.RevokeNocIntermediateCertificate(
		setup,
		setup.Vendor1,
		icaCertificate1.Subject,
		icaCertificate1.SubjectKeyID,
		"",
		true)

	allRevokedCerts, _ := utils.QueryAllNocRevokedIcaCertificates(setup)
	require.Equal(t, 2, len(allRevokedCerts))

	// Check indexes
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.RevokedNocIcaCertificatesKeyPrefix, Count: 2},
			{Key: types.NocRootCertificatesKeyPrefix, Count: 1}, // root still exits
		},
		Missing: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
			{Key: types.NocIcaCertificatesKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.RevokedNocRootCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate1, indexes)
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate2, indexes)

	indexes = utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.RevokedNocIcaCertificatesKeyPrefix, Count: 1},
			{Key: types.NocRootCertificatesKeyPrefix, Count: 1}, // root still exits
		},
		Missing: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
			{Key: types.NocIcaCertificatesKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.RevokedNocRootCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, leafCertificate, indexes)
}

func TestHandler_RevokeNocX509Cert_RevokeBySerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	// add the first NOC root certificate
	utils.AddNocRootCertificate(setup, setup.Vendor1, testconstants.NocRootCert1)

	// add the first NOC non-root certificate
	icaCertificate1 := utils.CreateTestNocIca1Cert()
	utils.AddNocIntermediateCertificate(setup, setup.Vendor1, icaCertificate1.PEM)

	// add the second NOC non-root certificate
	icaCertificate2 := utils.CreateTestNocIca1CertCopy()
	utils.AddNocIntermediateCertificate(setup, setup.Vendor1, icaCertificate2.PEM)

	// add the NOC leaf certificate
	leafCertificate := utils.CreateTestNocLeafCert()
	utils.AddNocIntermediateCertificate(setup, setup.Vendor1, leafCertificate.PEM)

	// Revoke NOC by serial number only
	utils.RevokeNocIntermediateCertificate(
		setup,
		setup.Vendor1,
		icaCertificate1.Subject,
		icaCertificate1.SubjectKeyID,
		icaCertificate1.SerialNumber,
		false)

	// Check indexes for intermediate after revocation
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.RevokedNocIcaCertificatesKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
			{Key: types.NocIcaCertificatesKeyPrefix, Count: 2}, // intermediate + leaf
			{Key: types.ChildCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.RevokedNocRootCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate1, indexes)
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate2, indexes)

	// Check indexes for leaf
	indexes = utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
			{Key: types.NocIcaCertificatesKeyPrefix, Count: 2}, // inter + leaf
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.UniqueCertificateKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.RevokedNocIcaCertificatesKeyPrefix},
			{Key: types.RevokedNocRootCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, leafCertificate, indexes)
}

func TestHandler_RevokeNocX509Cert_RevokeBySerialNumberAndWithChild(t *testing.T) {
	setup := utils.Setup(t)

	// add the first NOC root certificate
	utils.AddNocRootCertificate(setup, setup.Vendor1, testconstants.NocRootCert1)

	// add the first NOC non-root certificate
	icaCertificate1 := utils.CreateTestNocIca1Cert()
	utils.AddNocIntermediateCertificate(setup, setup.Vendor1, icaCertificate1.PEM)

	// add the second NOC non-root certificate
	icaCertificate2 := utils.CreateTestNocIca1CertCopy()
	utils.AddNocIntermediateCertificate(setup, setup.Vendor1, icaCertificate2.PEM)

	// add the NOC leaf certificate
	leafCertificate := utils.CreateTestNocLeafCert()
	utils.AddNocIntermediateCertificate(setup, setup.Vendor1, leafCertificate.PEM)

	// Revoke NOC with subject and subject key id and its child too
	utils.RevokeNocIntermediateCertificate(
		setup,
		setup.Vendor1,
		icaCertificate1.Subject,
		icaCertificate1.SubjectKeyID,
		icaCertificate1.SerialNumber,
		true)

	// Check indexes certificates
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.RevokedNocIcaCertificatesKeyPrefix},
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
			{Key: types.RevokedNocRootCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate1, indexes)
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate2, indexes)

	indexes = utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.RevokedNocIcaCertificatesKeyPrefix},
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
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.RevokedNocRootCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, leafCertificate, indexes)
}

// Extra cases

// Error cases

func TestHandler_RevokeNocX509Cert_SenderNotVendor(t *testing.T) {
	setup := utils.Setup(t)

	// add the new NOC root certificate
	addNocX509RootCert := types.NewMsgAddNocX509RootCert(setup.Vendor1.String(), testconstants.NocRootCert1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
	require.NoError(t, err)

	revokeCert := types.NewMsgRevokeNocX509RootCert(
		setup.Trustee1.String(),
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1SerialNumber,
		"",
		false,
	)
	_, err = setup.Handler(setup.Ctx, revokeCert)

	require.Error(t, err)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_RevokeNocX509Cert_CertificateDoesNotExist(t *testing.T) {
	setup := utils.Setup(t)

	revokeCert := types.NewMsgRevokeNocX509IcaCert(
		setup.Vendor1.String(),
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1SerialNumber,
		"",
		false,
	)
	_, err := setup.Handler(setup.Ctx, revokeCert)

	require.Error(t, err)
	require.ErrorIs(t, err, pkitypes.ErrCertificateDoesNotExist)
}

func TestHandler_RevokeNocX509Cert_CertificateExists(t *testing.T) {
	accAddress := utils.GenerateAccAddress()

	cases := []struct {
		name         string
		existingCert *types.Certificate
		nocRoorCert  string
		err          error
	}{
		{
			name: "ExistingRootCert",
			existingCert: &types.Certificate{
				Issuer:          testconstants.NocCert1Subject,
				Subject:         testconstants.NocCert1Subject,
				SubjectAsText:   testconstants.NocCert1SubjectAsText,
				SubjectKeyId:    testconstants.NocCert1SubjectKeyID,
				SerialNumber:    testconstants.NocCert1SerialNumber,
				IsRoot:          true,
				CertificateType: types.CertificateType_OperationalPKI,
				Vid:             testconstants.Vid,
			},
			nocRoorCert: testconstants.NocRootCert1,
			err:         pkitypes.ErrInappropriateCertificateType,
		},
		{
			name: "ExistingNotNocCert",
			existingCert: &types.Certificate{
				Issuer:          testconstants.NocCert1Subject,
				Subject:         testconstants.NocCert1Subject,
				SubjectAsText:   testconstants.NocCert1SubjectAsText,
				SubjectKeyId:    testconstants.NocCert1SubjectKeyID,
				SerialNumber:    testconstants.NocCert1SerialNumber,
				IsRoot:          true,
				CertificateType: types.CertificateType_DeviceAttestationPKI,
				Vid:             testconstants.Vid,
			},
			nocRoorCert: testconstants.NocCert1,
			err:         pkitypes.ErrInappropriateCertificateType,
		},
		{
			name: "ExistingCertWithDifferentVid",
			existingCert: &types.Certificate{
				Issuer:          testconstants.NocCert1Subject,
				Subject:         testconstants.NocCert1Subject,
				SubjectAsText:   testconstants.NocCert1SubjectAsText,
				SubjectKeyId:    testconstants.NocCert1SubjectKeyID,
				SerialNumber:    testconstants.NocCert1SerialNumber,
				IsRoot:          false,
				CertificateType: types.CertificateType_OperationalPKI,
				Vid:             testconstants.VendorID1,
			},
			nocRoorCert: testconstants.NocCert1,
			err:         pkitypes.ErrCertVidNotEqualAccountVid,
		},
		{
			name: "ExistingCertWithDifferentSerialNumber",
			existingCert: &types.Certificate{
				Issuer:          testconstants.NocCert1Subject,
				Subject:         testconstants.NocCert1Subject,
				SubjectAsText:   testconstants.NocCert1SubjectAsText,
				SubjectKeyId:    testconstants.NocCert1SubjectKeyID,
				SerialNumber:    "1234567",
				IsRoot:          false,
				CertificateType: types.CertificateType_OperationalPKI,
				Vid:             testconstants.Vid,
			},
			nocRoorCert: testconstants.NocCert1,
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

			revokeCert := types.NewMsgRevokeNocX509IcaCert(
				accAddress.String(),
				testconstants.NocCert1Subject,
				testconstants.NocCert1SubjectKeyID,
				testconstants.NocCert1SerialNumber,
				"",
				false,
			)
			_, err := setup.Handler(setup.Ctx, revokeCert)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
