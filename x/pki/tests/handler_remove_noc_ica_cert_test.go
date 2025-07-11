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

func TestHandler_RemoveNocIntermediateCert_BySubjectAndSKID(t *testing.T) {
	for _, crtType := range certificatesTypes {
		t.Run(crtType.String(), func(t *testing.T) {
			setup := utils.Setup(t)

			// add NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			// add two intermediate certificates
			icaCertificate1 := utils.IntermediateNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate1)

			icaCertificate2 := utils.IntermediateNocCertificate1Copy(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate2)

			// check total number of certificates
			nocCerts := setup.Keeper.GetAllNocCertificates(setup.Ctx)
			require.Equal(t, 2, len(nocCerts))

			// remove all intermediate certificates
			utils.RemoveNocIntermediateCertificate(
				setup,
				setup.Vendor1,
				icaCertificate1.Subject,
				icaCertificate1.SubjectKeyId,
				"",
			)

			// Check indexes for intermediate certificates - removed
			indexes := utils.TestIndexes{
				Present: []utils.TestIndex{
					{Key: types.NocRootCertificatesKeyPrefix, Count: 1}, // root still exits
				},
				Missing: []utils.TestIndex{
					{Key: types.AllCertificatesKeyPrefix},
					{Key: types.AllCertificatesBySubjectKeyPrefix},
					{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
					{Key: types.NocCertificatesKeyPrefix},
					{Key: types.NocCertificatesBySubjectKeyPrefix},
					{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
					{Key: types.NocIcaCertificatesKeyPrefix},
					{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
					{Key: types.UniqueCertificateKeyPrefix},
					{Key: types.ChildCertificatesKeyPrefix},
					{Key: types.RevokedNocIcaCertificatesKeyPrefix},
					{Key: types.RevokedNocRootCertificatesKeyPrefix},
					{Key: types.RevokedCertificatesKeyPrefix},
				},
			}
			utils.CheckCertificateStateIndexes(t, setup, icaCertificate1, indexes)
			utils.CheckCertificateStateIndexes(t, setup, icaCertificate2, indexes)

			// Check that only 1 certificate exists (root)
			nocCerts, _ = utils.QueryAllNocCertificates(setup)
			require.Equal(t, 1, len(nocCerts))
		})
	}
}

func TestHandler_RemoveNocIntermediateCert_BySerialNumber(t *testing.T) {
	for _, crtType := range certificatesTypes {
		t.Run(crtType.String(), func(t *testing.T) {
			setup := utils.Setup(t)

			// add NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			// Add ICA certificates
			icaCertificate1 := utils.IntermediateNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate1)

			// Add ICA certificates with sam subject and SKID but different serial number
			icaCertificate2 := utils.IntermediateNocCertificate1Copy(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate2)

			// remove ICA certificate by serial number
			utils.RemoveNocIntermediateCertificate(
				setup,
				setup.Vendor1,
				icaCertificate1.Subject,
				icaCertificate1.SubjectKeyId,
				icaCertificate1.SerialNumber)

			// Check indexes for first certificate - removed (no exist in unique index, but second approved ica exist)
			indexes := utils.TestIndexes{
				Present: []utils.TestIndex{
					{Key: types.AllCertificatesKeyPrefix, Count: 1},
					{Key: types.AllCertificatesBySubjectKeyPrefix, Count: 1},
					{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Count: 1},
					{Key: types.NocCertificatesKeyPrefix, Count: 1},
					{Key: types.NocCertificatesBySubjectKeyPrefix, Count: 1},
					{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix, Count: 1},
					{Key: types.NocCertificatesByVidAndSkidKeyPrefix, Count: 1},
					{Key: types.NocRootCertificatesKeyPrefix, Count: 1}, // root still exits
					{Key: types.NocIcaCertificatesKeyPrefix, Count: 1},
					{Key: types.ChildCertificatesKeyPrefix},
				},
				Missing: []utils.TestIndex{
					{Key: types.UniqueCertificateKeyPrefix}, // removed
					{Key: types.RevokedNocIcaCertificatesKeyPrefix},
					{Key: types.RevokedNocRootCertificatesKeyPrefix},
					{Key: types.RevokedCertificatesKeyPrefix},
				},
			}
			utils.CheckCertificateStateIndexes(t, setup, icaCertificate1, indexes)

			// Check indexes for second certificate (all same as for ica1 but also UniqueCertificate exists)
			indexes = utils.TestIndexes{
				Present: []utils.TestIndex{
					{Key: types.UniqueCertificateKeyPrefix}, // all same as for ica1 but also UniqueCertificate exists
					{Key: types.AllCertificatesKeyPrefix, Count: 1},
					{Key: types.AllCertificatesBySubjectKeyPrefix, Count: 1},
					{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Count: 1},
					{Key: types.NocCertificatesKeyPrefix, Count: 1},
					{Key: types.NocCertificatesBySubjectKeyPrefix, Count: 1},
					{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix, Count: 1},
					{Key: types.NocCertificatesByVidAndSkidKeyPrefix, Count: 1},
					{Key: types.NocRootCertificatesKeyPrefix, Count: 1}, // root still exits
					{Key: types.NocIcaCertificatesKeyPrefix, Count: 1},
					{Key: types.ChildCertificatesKeyPrefix},
				},
				Missing: []utils.TestIndex{
					{Key: types.RevokedNocIcaCertificatesKeyPrefix},
					{Key: types.RevokedNocRootCertificatesKeyPrefix},
					{Key: types.RevokedCertificatesKeyPrefix},
				},
			}
			utils.CheckCertificateStateIndexes(t, setup, icaCertificate2, indexes)
		})
	}
}

func TestHandler_RemoveNocIntermediateCert_BySubjectAndSKID_ParentExist(t *testing.T) {
	for _, crtType := range certificatesTypes {
		t.Run(crtType.String(), func(t *testing.T) {
			setup := utils.Setup(t)

			// add NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			// add two intermediate certificates
			icaCertificate1 := utils.IntermediateNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate1)

			icaCertificate2 := utils.IntermediateNocCertificate1Copy(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate2)

			// add leaf certificate
			leafCertificate := utils.LeafNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, leafCertificate)

			// check total number of certificates
			nocCerts := setup.Keeper.GetAllNocCertificates(setup.Ctx)
			require.Equal(t, 3, len(nocCerts))

			// remove all intermediate certificates but leave leaf certificate (NocCert1 and IntermediateNocCertificate1Copy)
			utils.RemoveNocIntermediateCertificate(
				setup,
				setup.Vendor1,
				icaCertificate1.Subject,
				icaCertificate1.SubjectKeyId,
				"",
			)

			// Check indexes for root certificate - approved
			indexes := utils.TestIndexes{
				Present: []utils.TestIndex{
					{Key: types.AllCertificatesKeyPrefix},
					{Key: types.AllCertificatesBySubjectKeyPrefix},
					{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
					{Key: types.NocCertificatesKeyPrefix},
					{Key: types.NocCertificatesBySubjectKeyPrefix},
					{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
					{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
					{Key: types.NocRootCertificatesKeyPrefix, Count: 1},
					{Key: types.NocIcaCertificatesKeyPrefix, Count: 1},
					{Key: types.UniqueCertificateKeyPrefix},
				},
				Missing: []utils.TestIndex{
					{Key: types.ProposedCertificateKeyPrefix},
					{Key: types.ApprovedCertificatesKeyPrefix},
					{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
					{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
					{Key: types.ApprovedRootCertificatesKeyPrefix},
				},
			}
			utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)

			// Check that only 2 certificates exists
			nocCerts, _ = utils.QueryAllNocCertificates(setup)
			require.Equal(t, 2, len(nocCerts))
		})
	}
}

func TestHandler_RemoveNocIntermediateCert_BySerialNumber_ParentExist(t *testing.T) {
	for _, crtType := range certificatesTypes {
		t.Run(crtType.String(), func(t *testing.T) {
			setup := utils.Setup(t)

			// add NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			// Add ICA certificates
			icaCertificate1 := utils.IntermediateNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate1)

			// Add ICA certificates with sam subject and SKID but different serial number
			icaCertificate2 := utils.IntermediateNocCertificate1Copy(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate2)

			// Add a leaf certificate
			leafCertificate := utils.LeafNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, leafCertificate)

			// remove ICA certificate by serial number
			utils.RemoveNocIntermediateCertificate(
				setup,
				setup.Vendor1,
				icaCertificate1.Subject,
				icaCertificate1.SubjectKeyId,
				icaCertificate1.SerialNumber)

			// Check indexes for root certificate - approved
			indexes := utils.TestIndexes{
				Present: []utils.TestIndex{
					{Key: types.AllCertificatesKeyPrefix, Count: 1},
					{Key: types.AllCertificatesBySubjectKeyPrefix, Count: 1},
					{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Count: 1},
					{Key: types.NocCertificatesKeyPrefix, Count: 1},
					{Key: types.NocCertificatesBySubjectKeyPrefix, Count: 1},
					{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix, Count: 1},
					{Key: types.NocCertificatesByVidAndSkidKeyPrefix, Count: 1},
					{Key: types.NocRootCertificatesKeyPrefix, Count: 1},
					{Key: types.NocIcaCertificatesKeyPrefix, Count: 2}, // root and leaf cert with same vid exist
					{Key: types.UniqueCertificateKeyPrefix},
				},
				Missing: []utils.TestIndex{
					{Key: types.RevokedNocIcaCertificatesKeyPrefix},
					{Key: types.RevokedNocRootCertificatesKeyPrefix},
					{Key: types.RevokedCertificatesKeyPrefix},
				},
			}
			utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)
		})
	}
}

func TestHandler_RemoveNocIntermediateCert_BySubjectAndSKID_ApprovedChildExist(t *testing.T) {
	for _, crtType := range certificatesTypes {
		t.Run(crtType.String(), func(t *testing.T) {
			setup := utils.Setup(t)

			// add NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			// add two intermediate certificates
			icaCertificate1 := utils.IntermediateNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate1)

			icaCertificate2 := utils.IntermediateNocCertificate1Copy(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate2)

			// add leaf certificate
			leafCertificate := utils.LeafNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, leafCertificate)

			// check total number of certificates
			nocCerts := setup.Keeper.GetAllNocCertificates(setup.Ctx)
			require.Equal(t, 3, len(nocCerts))

			// remove all intermediate certificates but leave leaf certificate (NocCert1 and IntermediateNocCertificate1Copy)
			utils.RemoveNocIntermediateCertificate(
				setup,
				setup.Vendor1,
				icaCertificate1.Subject,
				icaCertificate1.SubjectKeyId,
				"",
			)

			// Check indexes for leaf certificate - approved
			indexes := utils.TestIndexes{
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

			// Check that only 2 certificates exists
			nocCerts, _ = utils.QueryAllNocCertificates(setup)
			require.Equal(t, 2, len(nocCerts))
		})
	}
}

func TestHandler_RemoveNocIntermediateCert_BySerialNumber_ApprovedChildExist(t *testing.T) {
	for _, crtType := range certificatesTypes {
		t.Run(crtType.String(), func(t *testing.T) {
			setup := utils.Setup(t)

			// add NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			// Add ICA certificates
			icaCertificate1 := utils.IntermediateNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate1)

			// Add ICA certificates with sam subject and SKID but different serial number
			icaCertificate2 := utils.IntermediateNocCertificate1Copy(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate2)

			// Add a leaf certificate
			leafCertificate := utils.LeafNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, leafCertificate)

			// remove ICA certificate by serial number
			utils.RemoveNocIntermediateCertificate(
				setup,
				setup.Vendor1,
				icaCertificate1.Subject,
				icaCertificate1.SubjectKeyId,
				icaCertificate1.SerialNumber)

			// Check indexes for leaf certificate - approved
			indexes := utils.TestIndexes{
				Present: []utils.TestIndex{
					{Key: types.AllCertificatesKeyPrefix, Count: 1},
					{Key: types.AllCertificatesBySubjectKeyPrefix, Count: 1},
					{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Count: 1},
					{Key: types.NocCertificatesKeyPrefix, Count: 1},
					{Key: types.NocCertificatesBySubjectKeyPrefix, Count: 1},
					{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix, Count: 1},
					{Key: types.NocCertificatesByVidAndSkidKeyPrefix, Count: 1},
					{Key: types.NocRootCertificatesKeyPrefix, Count: 1}, // root still exits
					{Key: types.NocIcaCertificatesKeyPrefix, Count: 2},  // ica and leaf cert with same vid exist
					{Key: types.UniqueCertificateKeyPrefix},
					{Key: types.ChildCertificatesKeyPrefix},
				},
				Missing: []utils.TestIndex{
					{Key: types.RevokedNocIcaCertificatesKeyPrefix},
					{Key: types.RevokedNocRootCertificatesKeyPrefix},
					{Key: types.RevokedCertificatesKeyPrefix},
				},
			}
			utils.CheckCertificateStateIndexes(t, setup, leafCertificate, indexes)
		})
	}
}

func TestHandler_RemoveNocIntermediateCert_BySubjectAndSKID_RevokedChildExist(t *testing.T) {
	for _, crtType := range certificatesTypes {
		t.Run(crtType.String(), func(t *testing.T) {
			setup := utils.Setup(t)

			// add NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			// add two intermediate certificates
			icaCertificate1 := utils.IntermediateNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate1)

			icaCertificate2 := utils.IntermediateNocCertificate1Copy(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate2)

			// add leaf certificate
			leafCertificate := utils.LeafNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, leafCertificate)

			// check total number of certificates
			nocCerts := setup.Keeper.GetAllNocCertificates(setup.Ctx)
			require.Equal(t, 3, len(nocCerts))

			// revoke leaf certificate
			utils.RevokeNocIntermediateCertificate(
				setup,
				setup.Vendor1,
				leafCertificate.Subject,
				leafCertificate.SubjectKeyId,
				"",
				false,
			)

			// remove all intermediate certificates but leave leaf certificate (NocCert1 and IntermediateNocCertificate1Copy)
			utils.RemoveNocIntermediateCertificate(
				setup,
				setup.Vendor1,
				icaCertificate1.Subject,
				icaCertificate1.SubjectKeyId,
				"",
			)

			// Check indexes for leaf certificate - revoked
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
				},
			}
			utils.CheckCertificateStateIndexes(t, setup, leafCertificate, indexes)
		})
	}
}

func TestHandler_RemoveNocIntermediateCert_BySerialNumber_RevokedChildExist(t *testing.T) {
	for _, crtType := range certificatesTypes {
		t.Run(crtType.String(), func(t *testing.T) {
			setup := utils.Setup(t)

			// add NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			// Add ICA certificates
			icaCertificate1 := utils.IntermediateNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate1)

			// Add ICA certificates with sam subject and SKID but different serial number
			icaCertificate2 := utils.IntermediateNocCertificate1Copy(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate2)

			// Add a leaf certificate
			leafCertificate := utils.LeafNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, leafCertificate)

			// revoke leaf certificate
			utils.RevokeNocIntermediateCertificate(
				setup,
				setup.Vendor1,
				leafCertificate.Subject,
				leafCertificate.SubjectKeyId,
				"",
				false,
			)

			// remove ICA certificate by serial number
			utils.RemoveNocIntermediateCertificate(
				setup,
				setup.Vendor1,
				icaCertificate1.Subject,
				icaCertificate1.SubjectKeyId,
				icaCertificate1.SerialNumber)

			// Check indexes for leaf certificate- revoked
			indexes := utils.TestIndexes{
				Present: []utils.TestIndex{
					{Key: types.UniqueCertificateKeyPrefix},
					{Key: types.RevokedNocIcaCertificatesKeyPrefix},
					{Key: types.NocRootCertificatesKeyPrefix, Count: 1}, // root still exits
					{Key: types.NocIcaCertificatesKeyPrefix},            // single intermediate exists
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
				},
			}
			utils.CheckCertificateStateIndexes(t, setup, leafCertificate, indexes)
		})
	}
}

func TestHandler_RemoveNocIntermediateCert_BySubjectAndSKID_RevokedCertificate(t *testing.T) {
	for _, crtType := range certificatesTypes {
		t.Run(crtType.String(), func(t *testing.T) {
			setup := utils.Setup(t)

			// add NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			// Add an intermediate certificate
			icaCertificate1 := utils.IntermediateNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate1)

			// Add an intermediate certificate
			icaCertificate2 := utils.IntermediateNocCertificate1Copy(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate2)

			// revoke intermediate certificate by serial number
			utils.RevokeNocIntermediateCertificate(
				setup,
				setup.Vendor1,
				icaCertificate1.Subject,
				icaCertificate1.SubjectKeyId,
				"",
				false,
			)

			// remove ICA certificate by serial number
			utils.RemoveNocIntermediateCertificate(
				setup,
				setup.Vendor1,
				icaCertificate1.Subject,
				icaCertificate1.SubjectKeyId,
				"")

			// Check indexes after revocation - removed
			indexes := utils.TestIndexes{
				Present: []utils.TestIndex{},
				Missing: []utils.TestIndex{
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
					{Key: types.RevokedNocRootCertificatesKeyPrefix},
					{Key: types.RevokedCertificatesKeyPrefix},
				},
			}
			utils.CheckCertificateStateIndexes(t, setup, icaCertificate1, indexes)
			utils.CheckCertificateStateIndexes(t, setup, icaCertificate2, indexes)
		})
	}
}

func TestHandler_RemoveNocIntermediateCert_BySerialNumber_RevokedCertificate(t *testing.T) {
	for _, crtType := range certificatesTypes {
		t.Run(crtType.String(), func(t *testing.T) {
			setup := utils.Setup(t)

			// add NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			// Add an intermediate certificate
			icaCertificate1 := utils.IntermediateNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate1)

			// Add an intermediate certificate
			icaCertificate2 := utils.IntermediateNocCertificate1Copy(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate2)

			// revoke intermediate certificate by serial number
			utils.RevokeNocIntermediateCertificate(
				setup,
				setup.Vendor1,
				icaCertificate1.Subject,
				icaCertificate1.SubjectKeyId,
				icaCertificate1.SerialNumber,
				false,
			)

			// remove ICA certificate by serial number
			utils.RemoveNocIntermediateCertificate(
				setup,
				setup.Vendor1,
				icaCertificate1.Subject,
				icaCertificate1.SubjectKeyId,
				icaCertificate1.SerialNumber)

			// Check indexes for certificate 1 - removed (unique does not exist but another approved exists)
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
					{Key: types.ChildCertificatesKeyPrefix},
				},
				Missing: []utils.TestIndex{
					{Key: types.UniqueCertificateKeyPrefix},
					{Key: types.RevokedNocIcaCertificatesKeyPrefix},
					{Key: types.RevokedNocRootCertificatesKeyPrefix},
					{Key: types.RevokedCertificatesKeyPrefix},
				},
			}
			utils.CheckCertificateStateIndexes(t, setup, icaCertificate1, indexes)

			// Check indexes for certificate 1 - approved
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
					{Key: types.RevokedNocRootCertificatesKeyPrefix},
					{Key: types.RevokedCertificatesKeyPrefix},
				},
			}
			utils.CheckCertificateStateIndexes(t, setup, icaCertificate2, indexes)
		})
	}
}

func TestHandler_RemoveNocIntermediateCert_BySubjectAndSKID_RevokedAndActiveCertificate(t *testing.T) {
	for _, crtType := range certificatesTypes {
		t.Run(crtType.String(), func(t *testing.T) {
			setup := utils.Setup(t)

			// add NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			// Add an intermediate certificate
			icaCertificate := utils.IntermediateNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate)

			// revoke an intermediate certificate
			utils.RevokeNocIntermediateCertificate(
				setup,
				setup.Vendor1,
				icaCertificate.Subject,
				icaCertificate.SubjectKeyId,
				icaCertificate.SerialNumber,
				false,
			)

			// Add an intermediate certificate with new serial number
			icaCertificate2 := utils.IntermediateNocCertificate1Copy(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate2)

			// remove an intermediate certificate
			utils.RemoveNocIntermediateCertificate(
				setup,
				setup.Vendor1,
				icaCertificate.Subject,
				icaCertificate.SubjectKeyId,
				"",
			)

			// check that only root certificates exists
			allCerts, _ := utils.QueryAllNocCertificates(setup)
			require.Equal(t, 1, len(allCerts))

			// check state indexes for intermediate certificates
			indexes := utils.TestIndexes{
				Present: []utils.TestIndex{},
				Missing: []utils.TestIndex{
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
					{Key: types.RevokedNocIcaCertificatesKeyPrefix},
					{Key: types.RevokedNocRootCertificatesKeyPrefix},
					{Key: types.RevokedCertificatesKeyPrefix},
				},
			}
			utils.CheckCertificateStateIndexes(t, setup, icaCertificate, indexes)
			utils.CheckCertificateStateIndexes(t, setup, icaCertificate2, indexes)
		})
	}
}

func TestHandler_RemoveNocIntermediateCert_ByNotOwnerButSameVendor(t *testing.T) {
	for _, crtType := range certificatesTypes {
		t.Run(crtType.String(), func(t *testing.T) {
			setup := utils.Setup(t)

			// add NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			// add ICA certificate by fist vendor account
			icaCertificate := utils.IntermediateNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate)

			// add second vendor account with VID = 1
			vendorAccAddress2 := utils.GenerateAccAddress()
			setup.AddAccount(vendorAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

			// remove certificate by second vendor account
			utils.RemoveNocIntermediateCertificate(
				setup,
				vendorAccAddress2,
				icaCertificate.Subject,
				icaCertificate.SubjectKeyId,
				icaCertificate.SerialNumber,
			)

			// check state indexes for intermediate certificates - removed
			indexes := utils.TestIndexes{
				Present: []utils.TestIndex{},
				Missing: []utils.TestIndex{
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
					{Key: types.RevokedNocIcaCertificatesKeyPrefix},
					{Key: types.RevokedNocRootCertificatesKeyPrefix},
					{Key: types.RevokedCertificatesKeyPrefix},
				},
			}
			utils.CheckCertificateStateIndexes(t, setup, icaCertificate, indexes)
		})
	}
}

// Error cases

func TestHandler_RemoveNocIntermediateCert_CertificateDoesNotExist(t *testing.T) {
	setup := utils.Setup(t)

	removeIcaCert := types.NewMsgRemoveNocX509IcaCert(
		setup.Vendor1.String(),
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1SerialNumber)
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_RemoveNocIntermediateCert_ByOtherVendor(t *testing.T) {
	for _, crtType := range certificatesTypes {
		t.Run(crtType.String(), func(t *testing.T) {
			setup := utils.Setup(t)

			// add NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			// add two intermediate certificates
			icaCertificate1 := utils.IntermediateNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate1)

			// add fist vendor account with VID = 1
			vendorAccAddress1 := setup.CreateVendorAccount(testconstants.VendorID1)

			// remove ICA certificate by second vendor account
			removeIcaCert := types.NewMsgRemoveNocX509IcaCert(
				vendorAccAddress1.String(),
				icaCertificate1.Subject,
				icaCertificate1.SubjectKeyId,
				icaCertificate1.SerialNumber)
			_, err := setup.Handler(setup.Ctx, removeIcaCert)
			require.Error(t, err)
			require.True(t, pkitypes.ErrCertVidNotEqualAccountVid.Is(err))
		})
	}
}

func TestHandler_RemoveNocIntermediateCert_SenderNotVendor(t *testing.T) {
	for _, crtType := range certificatesTypes {
		t.Run(crtType.String(), func(t *testing.T) {
			setup := utils.Setup(t)

			// add NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			// add two intermediate certificates
			icaCertificate1 := utils.IntermediateNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate1)

			removeIcaCert := types.NewMsgRemoveNocX509IcaCert(
				setup.Trustee1.String(),
				icaCertificate1.Subject,
				icaCertificate1.SubjectKeyId,
				"")
			_, err := setup.Handler(setup.Ctx, removeIcaCert)
			require.Error(t, err)
			require.True(t, sdkerrors.ErrUnauthorized.Is(err))
		})
	}
}

func TestHandler_RemoveNocIntermediateCert_ForNonIcaCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// propose and approve x509 root certificate
	rootCert := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	// Add intermediate certificates
	testIntermediateCertificate := utils.IntermediateDaCertificate(setup.Vendor1)
	utils.AddDaIntermediateCertificate(setup, testIntermediateCertificate)

	removeIcaCert := types.NewMsgRemoveNocX509IcaCert(
		setup.Vendor1.String(),
		testIntermediateCertificate.Subject,
		testIntermediateCertificate.SubjectKeyId,
		"")
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_RemoveNocIntermediateCert_InvalidSerialNumber(t *testing.T) {
	for _, crtType := range certificatesTypes {
		t.Run(crtType.String(), func(t *testing.T) {
			setup := utils.Setup(t)

			// add NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			// add two intermediate certificates
			icaCertificate1 := utils.IntermediateNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate1)

			removeX509Cert := types.NewMsgRemoveNocX509IcaCert(
				setup.Vendor1.String(),
				icaCertificate1.Subject,
				icaCertificate1.SubjectKeyId,
				"invalid")
			_, err := setup.Handler(setup.Ctx, removeX509Cert)
			require.Error(t, err)
			require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
		})
	}
}

func TestHandler_RemoveNocIntermediateCert_ForRoot(t *testing.T) {
	for _, crtType := range certificatesTypes {
		t.Run(crtType.String(), func(t *testing.T) {
			setup := utils.Setup(t)

			// add NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			removeX509Cert := types.NewMsgRemoveNocX509IcaCert(
				setup.Vendor1.String(),
				rootCertificate.Subject,
				rootCertificate.SubjectKeyId,
				"")
			_, err := setup.Handler(setup.Ctx, removeX509Cert)
			require.Error(t, err)
			require.True(t, pkitypes.ErrInappropriateCertificateType.Is(err))
		})
	}
}
