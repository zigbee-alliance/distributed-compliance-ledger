package tests

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/tests/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// Main

func TestHandler_RevokeNocRootCert_BySubjectAndSKID(t *testing.T) {
	setup := utils.Setup(t)

	// add the first NOC root certificate
	rootCertificate := utils.RootNocCertificate1(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate)

	// add the second NOC root certificate
	rootCertificate2 := utils.RootNocCertificate1Copy(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate2)

	// Revoke NOC root with subject and subject key id only
	utils.RevokeNocRootCertificate(
		setup,
		setup.Vendor1,
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		"",
		false,
	)

	// Check indexes - both revoked
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
}

func TestHandler_RevokeNocRootCert_BySerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	// add the first NOC root certificate
	rootCertificate := utils.RootNocCertificate1(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate)

	// add the second NOC root certificate
	rootCertificate2 := utils.RootNocCertificate1Copy(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate2)

	// Revoke NOC root with subject and subject key id by serial number
	utils.RevokeNocRootCertificate(
		setup,
		setup.Vendor1,
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		rootCertificate.SerialNumber,
		false,
	)

	// Check indexes - both approved and revoked exist
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
			{Key: types.NocIcaCertificatesKeyPrefix},
			{Key: types.RevokedNocIcaCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate2, indexes)
}

func TestHandler_RevokeNocRootCert_BySubjectAndSKID_KeepChild(t *testing.T) {
	setup := utils.Setup(t)

	// add the first NOC root certificate
	rootCertificate1 := utils.RootNocCertificate1(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate1)

	// add the second NOC root certificate
	rootCertificate2 := utils.RootNocCertificate1Copy(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate2)

	// add the first NOC non-root certificate
	icaCertificate1 := utils.IntermediateNocCertificate1(setup.Vendor1)
	utils.AddNocIntermediateCertificate(setup, icaCertificate1)

	// Revoke NOC with subject and subject key id only
	utils.RevokeNocRootCertificate(
		setup,
		setup.Vendor1,
		rootCertificate1.Subject,
		rootCertificate1.SubjectKeyId,
		"",
		false)

	// Check state indexes for intermediate certificate - stays approved
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
			{Key: types.NocIcaCertificatesKeyPrefix},
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.NocRootCertificatesKeyPrefix},
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate1, indexes)
}

func TestHandler_RevokeNocRootCert_BySerialNumber_KeepChild(t *testing.T) {
	setup := utils.Setup(t)

	// add the first NOC root certificate
	rootCertificate1 := utils.RootNocCertificate1(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate1)

	// add the second NOC root certificate
	rootCertificate2 := utils.RootNocCertificate1Copy(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate2)

	// add the first NOC non-root certificate
	icaCertificate1 := utils.IntermediateNocCertificate1(setup.Vendor1)
	utils.AddNocIntermediateCertificate(setup, icaCertificate1)

	// Revoke NOC with subject and subject key id only
	utils.RevokeNocRootCertificate(
		setup,
		setup.Vendor1,
		rootCertificate1.Subject,
		rootCertificate1.SubjectKeyId,
		rootCertificate1.SerialNumber,
		false)

	// Check state indexes for intermediate certificate - stays approved
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
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate1, indexes)
}

func TestHandler_RevokeNocRootCert_BySubjectAndSKID_RevokeChild(t *testing.T) {
	setup := utils.Setup(t)

	// add the first NOC root certificate
	rootCertificate1 := utils.RootNocCertificate1(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate1)

	// the second NOC root certificate
	rootCertificate2 := utils.RootNocCertificate1Copy(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate2)

	// add the NOC intermediate certificate
	icaCertificate1 := utils.IntermediateNocCertificate1(setup.Vendor1)
	utils.AddNocIntermediateCertificate(setup, icaCertificate1)

	// Revoke noc with subject and subject key id and its child too
	utils.RevokeNocRootCertificate(
		setup,
		setup.Vendor1,
		rootCertificate1.Subject,
		rootCertificate1.SubjectKeyId,
		"",
		true)

	// Check indexes for intermediate certificate - revoked
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.RevokedNocIcaCertificatesKeyPrefix, Count: 1},
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
			{Key: types.NocRootCertificatesKeyPrefix}, // root also revoked
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate1, indexes)
}

func TestHandler_RevokeNocRootCert_BySerialNumber_RevokeChild(t *testing.T) {
	setup := utils.Setup(t)

	// add the first NOC root certificate
	rootCertificate1 := utils.RootNocCertificate1(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate1)

	// the second NOC root certificate
	rootCertificate2 := utils.RootNocCertificate1Copy(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate2)

	// add the NOC intermediate certificate
	icaCertificate1 := utils.IntermediateNocCertificate1(setup.Vendor1)
	utils.AddNocIntermediateCertificate(setup, icaCertificate1)

	// Revoke noc with subject and subject key id and its child too
	utils.RevokeNocRootCertificate(
		setup,
		setup.Vendor1,
		rootCertificate1.Subject,
		rootCertificate1.SubjectKeyId,
		rootCertificate1.SerialNumber,
		true)

	// Check indexes for intermediate certificates - revoked
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.RevokedNocIcaCertificatesKeyPrefix, Count: 1},
			{Key: types.NocRootCertificatesKeyPrefix, Count: 1}, // root with same vid still exits
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
}

func TestHandler_RevokeNocRootCert_OtherVendor(t *testing.T) {
	setup := utils.Setup(t)

	otherVendorAddress := setup.CreateVendorAccount(testconstants.Vid)

	// add the first NOC root certificate
	rootCertificate := utils.RootNocCertificate1(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate)

	// Revoke NOC root with subject and subject key id only
	utils.RevokeNocRootCertificate(
		setup,
		otherVendorAddress,
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		"",
		false,
	)

	// Check state indexes - intermediate certificate revoked
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

// Error cases

func TestHandler_RevokeNocRootCert_SenderNotVendor(t *testing.T) {
	setup := utils.Setup(t)

	// add the new NOC root certificate
	rootCertificate := utils.RootNocCertificate1(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate)

	revokeCert := types.NewMsgRevokeNocX509RootCert(
		setup.Trustee1.String(),
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		rootCertificate.SerialNumber,
		"",
		false,
	)
	_, err := setup.Handler(setup.Ctx, revokeCert)
	require.Error(t, err)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_RevokeNocRootCert_CertificateDoesNotExist(t *testing.T) {
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

func TestHandler_RevokeNocRootCert_CertificateExists(t *testing.T) {
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

			accAddress := setup.CreateVendorAccount(testconstants.Vid)

			// add the existing certificate
			utils.AddMokedNocCertificate(setup, *tc.existingCert)

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
