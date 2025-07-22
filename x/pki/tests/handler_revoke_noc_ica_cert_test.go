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

func TestHandler_RevokeNocIntermediateCert_BySubjectAndSKID(t *testing.T) {

	cases := []CertificateTestCase{
		{
			name:    "OperationalPKI_RevokeNocIntermediateCert_BySubjectAndSKID",
			crtType: types.CertificateType_OperationalPKI,
		},
		{
			name:    "VIDSignerPKI_RevokeNocIntermediateCert_BySubjectAndSKID",
			crtType: types.CertificateType_VIDSignerPKI,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			// add the first NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			// add the first NOC non-root certificate
			icaCertificate1 := utils.IntermediateNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate1)

			// add the second NOC non-root certificate
			icaCertificate2 := utils.IntermediateNocCertificate1Copy(setup.Vendor1, tc.crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate2)

			// Revoke NOC with subject and subject key id only
			utils.RevokeNocIntermediateCertificate(
				setup,
				setup.Vendor1,
				icaCertificate1.Subject,
				icaCertificate1.SubjectKeyId,
				"",
				false)

			// Check indexes - both intermediate are revoked
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
		})
	}
}

func TestHandler_RevokeNocIntermediateCert_BySerialNumber(t *testing.T) {

	cases := []CertificateTestCase{
		{
			name:    "OperationalPKI_RevokeNocIntermediateCert_BySerialNumber",
			crtType: types.CertificateType_OperationalPKI,
		},
		{
			name:    "VIDSignerPKI_RevokeNocIntermediateCert_BySerialNumber",
			crtType: types.CertificateType_VIDSignerPKI,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			// add the first NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			// add the first NOC non-root certificate
			icaCertificate1 := utils.IntermediateNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate1)

			// add the second NOC non-root certificate
			icaCertificate2 := utils.IntermediateNocCertificate1Copy(setup.Vendor1, tc.crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate2)

			// Revoke NOC by serial number only
			utils.RevokeNocIntermediateCertificate(
				setup,
				setup.Vendor1,
				icaCertificate1.Subject,
				icaCertificate1.SubjectKeyId,
				icaCertificate1.SerialNumber,
				false)

			// Check state indexes for intermediate - revoked and approved exist
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
		})
	}
}

func TestHandler_RevokeNocIntermediateCert_BySubjectAndSKID_ParentExist(t *testing.T) {

	cases := []CertificateTestCase{
		{
			name:    "OperationalPKI_RevokeNocIntermediateCert_BySubjectAndSKID_ParentExist",
			crtType: types.CertificateType_OperationalPKI,
		},
		{
			name:    "VIDSignerPKI_RevokeNocIntermediateCert_BySubjectAndSKID_ParentExist",
			crtType: types.CertificateType_VIDSignerPKI,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			// add the first NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			// add the NOC non-root certificate
			icaCertificate := utils.IntermediateNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate)

			// Revoke NOC with subject and subject key id only
			utils.RevokeNocIntermediateCertificate(
				setup,
				setup.Vendor1,
				icaCertificate.Subject,
				icaCertificate.SubjectKeyId,
				"",
				false)

			// Check state indexes for root - approved
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

func TestHandler_RevokeNocIntermediateCert_BySerialNumber_ParentExist(t *testing.T) {

	cases := []CertificateTestCase{
		{
			name:    "OperationalPKI_RevokeNocIntermediateCert_BySerialNumber_ParentExist",
			crtType: types.CertificateType_OperationalPKI,
		},
		{
			name:    "VIDSignerPKI_RevokeNocIntermediateCert_BySerialNumber_ParentExist",
			crtType: types.CertificateType_VIDSignerPKI,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			// add the first NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			// add the NOC non-root certificate
			icaCertificate := utils.IntermediateNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate)

			// Revoke NOC with subject and subject key id only
			utils.RevokeNocIntermediateCertificate(
				setup,
				setup.Vendor1,
				icaCertificate.Subject,
				icaCertificate.SubjectKeyId,
				icaCertificate.SerialNumber,
				false)

			// Check state indexes for root - approved
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

func TestHandler_RevokeNocIntermediateCert_BySubjectAndSKID_KeepChild(t *testing.T) {

	cases := []CertificateTestCase{
		{
			name:    "OperationalPKI_RevokeNocIntermediateCert_BySubjectAndSKID_KeepChild",
			crtType: types.CertificateType_OperationalPKI,
		},
		{
			name:    "VIDSignerPKI_RevokeNocIntermediateCert_BySubjectAndSKID_KeepChild",
			crtType: types.CertificateType_VIDSignerPKI,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			// add the first NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			// add the first NOC non-root certificate
			icaCertificate1 := utils.IntermediateNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate1)

			// add the second NOC non-root certificate
			icaCertificate2 := utils.IntermediateNocCertificate1Copy(setup.Vendor1, tc.crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate2)

			// add the NOC leaf certificate
			leafCertificate := utils.LeafNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocIntermediateCertificate(setup, leafCertificate)

			// Revoke NOC with subject and subject key id only
			utils.RevokeNocIntermediateCertificate(
				setup,
				setup.Vendor1,
				icaCertificate1.Subject,
				icaCertificate1.SubjectKeyId,
				"",
				false)

			// Check state indexes for leaf - approved
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
			utils.CheckCertificateStateIndexes(t, setup, leafCertificate, indexes)
		})
	}
}

func TestHandler_RevokeNocIntermediateCert_BySerialNumber_KeepChild(t *testing.T) {

	cases := []CertificateTestCase{
		{
			name:    "OperationalPKI_RevokeNocIntermediateCert_BySerialNumber_KeepChild",
			crtType: types.CertificateType_OperationalPKI,
		},
		{
			name:    "VIDSignerPKI_RevokeNocIntermediateCert_BySerialNumber_KeepChild",
			crtType: types.CertificateType_VIDSignerPKI,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			// add the first NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			// add the first NOC non-root certificate
			icaCertificate1 := utils.IntermediateNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate1)

			// add the second NOC non-root certificate
			icaCertificate2 := utils.IntermediateNocCertificate1Copy(setup.Vendor1, tc.crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate2)

			// add the NOC leaf certificate
			leafCertificate := utils.LeafNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocIntermediateCertificate(setup, leafCertificate)

			// Revoke NOC by serial number only
			utils.RevokeNocIntermediateCertificate(
				setup,
				setup.Vendor1,
				icaCertificate1.Subject,
				icaCertificate1.SubjectKeyId,
				icaCertificate1.SerialNumber,
				false)

			// Check state indexes for leaf - approved
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
					{Key: types.NocIcaCertificatesKeyPrefix, Count: 2},  // intermediate + leaf
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
		})
	}
}

func TestHandler_RevokeNocIntermediateCert_BySubjectAndSKID_RevokeChild(t *testing.T) {

	cases := []CertificateTestCase{
		{
			name:    "OperationalPKI_RevokeNocIntermediateCert_BySubjectAndSKID_RevokeChild",
			crtType: types.CertificateType_OperationalPKI,
		},
		{
			name:    "VIDSignerPKI_RevokeNocIntermediateCert_BySubjectAndSKID_RevokeChild",
			crtType: types.CertificateType_VIDSignerPKI,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			// add the first NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			// add the first NOC non-root certificate
			icaCertificate1 := utils.IntermediateNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate1)

			// add the second NOC non-root certificate
			icaCertificate2 := utils.IntermediateNocCertificate1Copy(setup.Vendor1, tc.crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate2)

			// add the NOC leaf certificate
			leafCertificate := utils.LeafNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocIntermediateCertificate(setup, leafCertificate)

			// Revoke noc with subject and subject key id and its child too
			utils.RevokeNocIntermediateCertificate(
				setup,
				setup.Vendor1,
				icaCertificate1.Subject,
				icaCertificate1.SubjectKeyId,
				"",
				true)

			// Check indexes for child - revoked
			indexes := utils.TestIndexes{
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
		})
	}
}

func TestHandler_RevokeNocIntermediateCert_BySerialNumber_RevokeChild(t *testing.T) {

	cases := []CertificateTestCase{
		{
			name:    "OperationalPKI_RevokeNocIntermediateCert_BySerialNumber_RevokeChild",
			crtType: types.CertificateType_OperationalPKI,
		},
		{
			name:    "VIDSignerPKI_RevokeNocIntermediateCert_BySerialNumber_RevokeChild",
			crtType: types.CertificateType_VIDSignerPKI,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			// add the first NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			// add the first NOC non-root certificate
			icaCertificate1 := utils.IntermediateNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate1)

			// add the second NOC non-root certificate
			icaCertificate2 := utils.IntermediateNocCertificate1Copy(setup.Vendor1, tc.crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate2)

			// add the NOC leaf certificate
			leafCertificate := utils.LeafNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocIntermediateCertificate(setup, leafCertificate)

			// Revoke noc with subject and subject key id and its child too
			utils.RevokeNocIntermediateCertificate(
				setup,
				setup.Vendor1,
				icaCertificate1.Subject,
				icaCertificate1.SubjectKeyId,
				icaCertificate1.SerialNumber,
				true)

			allRevokedCerts, _ := utils.QueryAllNocRevokedIcaCertificates(setup)
			require.Equal(t, 2, len(allRevokedCerts))

			// Check indexes for child - revoked
			indexes := utils.TestIndexes{
				Present: []utils.TestIndex{
					{Key: types.UniqueCertificateKeyPrefix},
					{Key: types.RevokedNocIcaCertificatesKeyPrefix, Count: 1},
					{Key: types.NocRootCertificatesKeyPrefix, Count: 1}, // root still exits
					{Key: types.NocIcaCertificatesKeyPrefix},            // inter with same vid exists
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
		})
	}
}

func TestHandler_RevokeNocIntermediateCert_ByOtherVendor(t *testing.T) {

	cases := []CertificateTestCase{
		{
			name:    "OperationalPKI_RevokeNocIntermediateCert_ByOtherVendor",
			crtType: types.CertificateType_OperationalPKI,
		},
		{
			name:    "VIDSignerPKI_RevokeNocIntermediateCert_ByOtherVendor",
			crtType: types.CertificateType_VIDSignerPKI,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			// add vendor with same vid
			otherVendor := setup.CreateVendorAccount(testconstants.Vid)

			// add the first NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			// add the NOC non-root certificate
			icaCertificate := utils.IntermediateNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate)

			// Revoke NOC with subject and subject key id only
			utils.RevokeNocIntermediateCertificate(
				setup,
				otherVendor,
				icaCertificate.Subject,
				icaCertificate.SubjectKeyId,
				"",
				false)

			// Check indexes for intermediate - revoked
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
		})
	}
}

// Error cases

func TestHandler_RevokeNocIntermediateCert_SenderNotVendor(t *testing.T) {

	cases := []CertificateTestCase{
		{
			name:    "OperationalPKI_RevokeNocIntermediateCert_SenderNotVendor",
			crtType: types.CertificateType_OperationalPKI,
		},
		{
			name:    "VIDSignerPKI_RevokeNocIntermediateCert_SenderNotVendor",
			crtType: types.CertificateType_VIDSignerPKI,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			// add the first NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			// add the first NOC non-root certificate
			icaCertificate1 := utils.IntermediateNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate1)

			revokeCert := types.NewMsgRevokeNocX509RootCert(
				setup.Trustee1.String(),
				icaCertificate1.Subject,
				icaCertificate1.SubjectKeyId,
				icaCertificate1.SerialNumber,
				"",
				false,
			)
			_, err := setup.Handler(setup.Ctx, revokeCert)
			require.Error(t, err)
			require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
		})
	}
}

func TestHandler_RevokeNocIntermediateCert_CertificateDoesNotExist(t *testing.T) {
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

func TestHandler_RevokeNocIntermediateCert_CertificateExists(t *testing.T) {
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
				Issuer:        testconstants.NocCert1Subject,
				Subject:       testconstants.NocCert1Subject,
				SubjectAsText: testconstants.NocCert1SubjectAsText,
				SubjectKeyId:  testconstants.NocCert1SubjectKeyID,
				SerialNumber:  testconstants.NocCert1SerialNumber,
				IsRoot:        true,
				Vid:           testconstants.Vid,
			},
			nocRoorCert: testconstants.NocRootCert1,
			err:         pkitypes.ErrInappropriateCertificateType,
		},
		{
			name: "ExistingNotNocCert",
			existingCert: &types.Certificate{
				Issuer:        testconstants.NocCert1Subject,
				Subject:       testconstants.NocCert1Subject,
				SubjectAsText: testconstants.NocCert1SubjectAsText,
				SubjectKeyId:  testconstants.NocCert1SubjectKeyID,
				SerialNumber:  testconstants.NocCert1SerialNumber,
				IsRoot:        true,
				Vid:           testconstants.Vid,
			},
			nocRoorCert: testconstants.NocCert1,
			err:         pkitypes.ErrInappropriateCertificateType,
		},
		{
			name: "ExistingCertWithDifferentVid",
			existingCert: &types.Certificate{
				Issuer:        testconstants.NocCert1Subject,
				Subject:       testconstants.NocCert1Subject,
				SubjectAsText: testconstants.NocCert1SubjectAsText,
				SubjectKeyId:  testconstants.NocCert1SubjectKeyID,
				SerialNumber:  testconstants.NocCert1SerialNumber,
				IsRoot:        false,
				Vid:           testconstants.VendorID1,
			},
			nocRoorCert: testconstants.NocCert1,
			err:         pkitypes.ErrCertVidNotEqualAccountVid,
		},
		{
			name: "ExistingCertWithDifferentSerialNumber",
			existingCert: &types.Certificate{
				Issuer:        testconstants.NocCert1Subject,
				Subject:       testconstants.NocCert1Subject,
				SubjectAsText: testconstants.NocCert1SubjectAsText,
				SubjectKeyId:  testconstants.NocCert1SubjectKeyID,
				SerialNumber:  "1234567",
				IsRoot:        false,
				Vid:           testconstants.Vid,
			},
			nocRoorCert: testconstants.NocCert1,
			err:         pkitypes.ErrCertificateDoesNotExist,
		},
	}

	crtCases := []CertificateTestCase{
		{
			name:    "OperationalPKI",
			crtType: types.CertificateType_OperationalPKI,
		},
		{
			name:    "VIDSignerPKI",
			crtType: types.CertificateType_VIDSignerPKI,
		},
	}

	for _, tc := range cases {
		for _, tcc := range crtCases {
			t.Run(tcc.name+"_"+tc.name, func(t *testing.T) {
				setup := utils.Setup(t)
				setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

				existingCert := *tc.existingCert

				// the test for this error requires different types
				if errors.Is(tc.err, pkitypes.ErrInappropriateCertificateType) && tc.existingCert.IsRoot {
					existingCert.CertificateType = types.CertificateType_DeviceAttestationPKI
				} else {
					existingCert.CertificateType = tcc.crtType
				}

				// add the existing certificate
				setup.Keeper.AddNocCertificate(setup.Ctx, existingCert)
				uniqueCertificate := types.UniqueCertificate{
					Issuer:       existingCert.Issuer,
					SerialNumber: existingCert.SerialNumber,
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
}
