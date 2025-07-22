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

func TestHandler_RemoveNocRootCert_BySubjectAndSKID(t *testing.T) {
	cases := []CertificateTestCase{
		{
			name:    "OperationalPKI_RemoveNocRootCert_BySubjectAndSKID",
			crtType: types.CertificateType_OperationalPKI,
		},
		{
			name:    "VIDSignerPKI_RemoveNocRootCert_BySubjectAndSKID",
			crtType: types.CertificateType_VIDSignerPKI,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			// add NOC root certificates
			rootCertificate1 := utils.RootNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate1)

			rootCertificate2 := utils.RootNocCertificate1Copy(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate2)

			// get certificates for further comparison
			nocCerts := setup.Keeper.GetAllNocCertificates(setup.Ctx)
			require.NotNil(t, nocCerts)
			require.Equal(t, 1, len(nocCerts))

			// remove all root noc root certificates
			utils.RemoveNocRootCertificate(
				setup,
				setup.Vendor1,
				rootCertificate1.Subject,
				rootCertificate1.SubjectKeyId,
				"",
			)

			// check that only IAC certificate exists
			nocCerts, _ = utils.QueryAllNocCertificates(setup)
			require.Equal(t, 0, len(nocCerts))

			// Check indexes for root certificates - all removed
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
					{Key: types.NocRootCertificatesKeyPrefix},
					{Key: types.UniqueCertificateKeyPrefix},
					{Key: types.RevokedNocIcaCertificatesKeyPrefix},
					{Key: types.RevokedNocRootCertificatesKeyPrefix},
					{Key: types.RevokedCertificatesKeyPrefix},
				},
			}
			utils.CheckCertificateStateIndexes(t, setup, rootCertificate1, indexes)
			utils.CheckCertificateStateIndexes(t, setup, rootCertificate2, indexes)
		})
	}
}

func TestHandler_RemoveNocRootCert_BySerialNumber(t *testing.T) {
	cases := []CertificateTestCase{
		{
			name:    "OperationalPKI_RemoveNocRootCert_BySerialNumber",
			crtType: types.CertificateType_OperationalPKI,
		},
		{
			name:    "VIDSignerPKI_RemoveNocRootCert_BySerialNumber",
			crtType: types.CertificateType_VIDSignerPKI,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			// add NOC root certificates
			rootCertificate1 := utils.RootNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate1)

			rootCertificate2 := utils.RootNocCertificate1Copy(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate2)

			// get certificates for further comparison
			nocCerts := setup.Keeper.GetAllNocCertificates(setup.Ctx)
			require.NotNil(t, nocCerts)
			require.Equal(t, 1, len(nocCerts))

			// remove NOC root certificate by serial number
			utils.RemoveNocRootCertificate(
				setup,
				setup.Vendor1,
				rootCertificate1.Subject,
				rootCertificate1.SubjectKeyId,
				rootCertificate1.SerialNumber)

			// Check indexes for root certificate1 - unique does not exist (another approved exists)
			indexes := utils.TestIndexes{
				Present: []utils.TestIndex{
					{Key: types.AllCertificatesKeyPrefix},
					{Key: types.AllCertificatesBySubjectKeyPrefix},
					{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
					{Key: types.NocCertificatesKeyPrefix},
					{Key: types.NocCertificatesBySubjectKeyPrefix},
					{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
					{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
				},
				Missing: []utils.TestIndex{
					{Key: types.UniqueCertificateKeyPrefix},
					{Key: types.RevokedNocRootCertificatesKeyPrefix},
					{Key: types.NocIcaCertificatesKeyPrefix},
				},
			}
			utils.CheckCertificateStateIndexes(t, setup, rootCertificate1, indexes)

			// Check indexes for root certificate2 - approved
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
				},
				Missing: []utils.TestIndex{
					{Key: types.RevokedNocRootCertificatesKeyPrefix},
					{Key: types.NocIcaCertificatesKeyPrefix},
				},
			}
			utils.CheckCertificateStateIndexes(t, setup, rootCertificate2, indexes)

			// remove second NOC root certificate by serial number and check that IAC cert is not removed
			utils.RemoveNocRootCertificate(
				setup,
				setup.Vendor1,
				rootCertificate2.Subject,
				rootCertificate2.SubjectKeyId,
				rootCertificate2.SerialNumber)

			// check total
			nocCerts, _ = utils.QueryAllNocCertificates(setup)
			require.Equal(t, 0, len(nocCerts))

			// Check indexes for root certificates
			indexes = utils.TestIndexes{
				Present: []utils.TestIndex{},
				Missing: []utils.TestIndex{
					{Key: types.AllCertificatesKeyPrefix},
					{Key: types.AllCertificatesBySubjectKeyPrefix},
					{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
					{Key: types.NocCertificatesKeyPrefix},
					{Key: types.NocCertificatesBySubjectKeyPrefix},
					{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
					{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
					{Key: types.NocRootCertificatesKeyPrefix},
					{Key: types.UniqueCertificateKeyPrefix},
					{Key: types.RevokedNocIcaCertificatesKeyPrefix},
					{Key: types.RevokedNocRootCertificatesKeyPrefix},
					{Key: types.RevokedCertificatesKeyPrefix},
				},
			}
			utils.CheckCertificateStateIndexes(t, setup, rootCertificate1, indexes)
			utils.CheckCertificateStateIndexes(t, setup, rootCertificate2, indexes)
		})
	}
}

func TestHandler_RemoveNocRootCert_BySubjectAndSKID_ChildExist(t *testing.T) {
	cases := []CertificateTestCase{
		{
			name:    "OperationalPKI_RemoveNocRootCert_BySubjectAndSKID_ChildExist",
			crtType: types.CertificateType_OperationalPKI,
		},
		{
			name:    "VIDSignerPKI_RemoveNocRootCert_BySubjectAndSKID_ChildExist",
			crtType: types.CertificateType_VIDSignerPKI,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			// add NOC root certificates
			rootCertificate1 := utils.RootNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate1)

			rootCertificate2 := utils.RootNocCertificate1Copy(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate2)

			// Add intermediate certificate
			icaCertificate := utils.IntermediateNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate)

			// remove all root noc root certificates
			utils.RemoveNocRootCertificate(
				setup,
				setup.Vendor1,
				rootCertificate1.Subject,
				rootCertificate1.SubjectKeyId,
				"",
			)

			// check that only IAC certificate exists
			nocCerts, _ := utils.QueryAllNocCertificates(setup)
			require.Equal(t, 1, len(nocCerts))
			require.Equal(t, 1, len(nocCerts[0].Certs))

			// Check state indexes for intermediate certificates - approved
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
				Missing: []utils.TestIndex{},
			}
			utils.CheckCertificateStateIndexes(t, setup, icaCertificate, indexes)
		})
	}
}

func TestHandler_RemoveNocRootCert_BySerialNumber_ChildExist(t *testing.T) {
	cases := []CertificateTestCase{
		{
			name:    "OperationalPKI_RemoveNocRootCert_BySerialNumber_ChildExist",
			crtType: types.CertificateType_OperationalPKI,
		},
		{
			name:    "VIDSignerPKI_RemoveNocRootCert_BySerialNumber_ChildExist",
			crtType: types.CertificateType_VIDSignerPKI,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			// add NOC root certificates
			rootCertificate1 := utils.RootNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate1)

			rootCertificate2 := utils.RootNocCertificate1Copy(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate2)

			// Add ICA certificates
			icaCertificate := utils.IntermediateNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate)

			// remove NOC root certificate by serial number
			utils.RemoveNocRootCertificate(
				setup,
				setup.Vendor1,
				rootCertificate1.Subject,
				rootCertificate1.SubjectKeyId,
				rootCertificate1.SerialNumber)

			// check total
			nocCerts, _ := utils.QueryAllNocCertificates(setup)
			require.Equal(t, 2, len(nocCerts))

			// Check indexes for intermediate certificates - approved
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
				Missing: []utils.TestIndex{},
			}
			utils.CheckCertificateStateIndexes(t, setup, icaCertificate, indexes)

			// remove NOC root certificate by serial number and check that IAC cert is not removed
			utils.RemoveNocRootCertificate(
				setup,
				setup.Vendor1,
				rootCertificate2.Subject,
				rootCertificate2.SubjectKeyId,
				rootCertificate2.SerialNumber)

			// check total
			nocCerts, _ = utils.QueryAllNocCertificates(setup)
			require.Equal(t, 1, len(nocCerts))
			require.Equal(t, 1, len(nocCerts[0].Certs))

			// Check indexes for intermediate certificates - approved
			indexes = utils.TestIndexes{
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
				Missing: []utils.TestIndex{},
			}
			utils.CheckCertificateStateIndexes(t, setup, icaCertificate, indexes)
		})
	}
}

func TestHandler_RemoveNocRootCert_BySubjectAndSKID_RevokedCertificate(t *testing.T) {
	cases := []CertificateTestCase{
		{
			name:    "OperationalPKI_RemoveNocRootCert_BySubjectAndSKID_RevokedCertificate",
			crtType: types.CertificateType_OperationalPKI,
		},
		{
			name:    "VIDSignerPKI_RemoveNocRootCert_BySubjectAndSKID_RevokedCertificate",
			crtType: types.CertificateType_VIDSignerPKI,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			// add NOC root certificate
			rootCertificate1 := utils.RootNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate1)

			rootCertificate2 := utils.RootNocCertificate1Copy(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate2)

			// revoke NOC root certificates
			utils.RevokeNocRootCertificate(
				setup,
				setup.Vendor1,
				rootCertificate2.Subject,
				rootCertificate2.SubjectKeyId,
				"",
				false,
			)

			// remove NOC root certificates
			utils.RemoveNocRootCertificate(
				setup,
				setup.Vendor1,
				rootCertificate2.Subject,
				rootCertificate2.SubjectKeyId,
				"",
			)

			// Check indexes for root certificates - removed
			indexes := utils.TestIndexes{
				Present: []utils.TestIndex{},
				Missing: []utils.TestIndex{
					{Key: types.UniqueCertificateKeyPrefix},
					{Key: types.AllCertificatesKeyPrefix},
					{Key: types.AllCertificatesBySubjectKeyPrefix},
					{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
					{Key: types.NocCertificatesKeyPrefix},
					{Key: types.NocCertificatesBySubjectKeyPrefix},
					{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
					{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
					{Key: types.RevokedNocRootCertificatesKeyPrefix},
				},
			}
			utils.CheckCertificateStateIndexes(t, setup, rootCertificate1, indexes)
			utils.CheckCertificateStateIndexes(t, setup, rootCertificate2, indexes)
		})
	}
}

func TestHandler_RemoveNocRootCert_BySerialNumber_RevokedCertificate(t *testing.T) {
	cases := []CertificateTestCase{
		{
			name:    "OperationalPKI_RemoveNocRootCert_BySerialNumber_RevokedCertificate",
			crtType: types.CertificateType_OperationalPKI,
		},
		{
			name:    "VIDSignerPKI_RemoveNocRootCert_BySerialNumber_RevokedCertificate",
			crtType: types.CertificateType_VIDSignerPKI,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			// add NOC root certificate
			rootCertificate1 := utils.RootNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate1)

			rootCertificate2 := utils.RootNocCertificate1Copy(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate2)

			// revoke NOC root certificates
			utils.RevokeNocRootCertificate(
				setup,
				setup.Vendor1,
				rootCertificate2.Subject,
				rootCertificate2.SubjectKeyId,
				"",
				false,
			)

			// remove NOC root certificates
			utils.RemoveNocRootCertificate(
				setup,
				setup.Vendor1,
				rootCertificate2.Subject,
				rootCertificate2.SubjectKeyId,
				rootCertificate2.SerialNumber,
			)

			// Check indexes for root certificate1 - revoked
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
				},
			}
			utils.CheckCertificateStateIndexes(t, setup, rootCertificate1, indexes)

			// Check indexes for root certificate2 - removed
			indexes = utils.TestIndexes{
				Present: []utils.TestIndex{
					// another root with same vid exists
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
					{Key: types.UniqueCertificateKeyPrefix},
				},
			}
			utils.CheckCertificateStateIndexes(t, setup, rootCertificate2, indexes)
		})
	}
}

func TestHandler_RemoveNocRootCert_BySubjectAndSKID_RevokedAndActiveCertificate(t *testing.T) {
	cases := []CertificateTestCase{
		{
			name:    "OperationalPKI_RemoveNocRootCert_BySubjectAndSKID_RevokedAndActiveCertificate",
			crtType: types.CertificateType_OperationalPKI,
		},
		{
			name:    "VIDSignerPKI_RemoveNocRootCert_BySubjectAndSKID_RevokedAndActiveCertificate",
			crtType: types.CertificateType_VIDSignerPKI,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			// add NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			// revoke an intermediate certificate
			utils.RevokeNocRootCertificate(
				setup,
				setup.Vendor1,
				rootCertificate.Subject,
				rootCertificate.SubjectKeyId,
				rootCertificate.SerialNumber,
				false,
			)

			// Add NOC root certificate with new serial number
			rootCertificate2 := utils.RootNocCertificate1Copy(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate2)

			// remove NOC root certificate by serial number
			utils.RemoveNocRootCertificate(
				setup,
				setup.Vendor1,
				rootCertificate.Subject,
				rootCertificate.SubjectKeyId,
				"",
			)

			// Check indexes for root certificates - removed
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
					{Key: types.UniqueCertificateKeyPrefix},
					{Key: types.RevokedNocRootCertificatesKeyPrefix},
				},
			}
			utils.CheckCertificateStateIndexes(t, setup, rootCertificate2, indexes)
			utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)
		})
	}
}

func TestHandler_RemoveNocRootCert_ByNotOwnerButSameVendor(t *testing.T) {
	cases := []CertificateTestCase{
		{
			name:    "OperationalPKI_RemoveNocRootCert_ByNotOwnerButSameVendor",
			crtType: types.CertificateType_OperationalPKI,
		},
		{
			name:    "VIDSignerPKI_RemoveNocRootCert_ByNotOwnerButSameVendor",
			crtType: types.CertificateType_VIDSignerPKI,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			// add NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			// add second vendor account with VID = 1
			vendorAccAddress2 := setup.CreateVendorAccount(testconstants.Vid)

			// remove x509 certificate by second vendor account
			utils.RemoveNocRootCertificate(
				setup,
				vendorAccAddress2,
				rootCertificate.Subject,
				rootCertificate.SubjectKeyId,
				rootCertificate.SerialNumber,
			)

			// Check indexes for root certificates - removed
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
					{Key: types.UniqueCertificateKeyPrefix},
					{Key: types.RevokedNocRootCertificatesKeyPrefix},
				},
			}
			utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)
		})
	}
}

// Error cases
func TestHandler_RemoveNocRootCert_CertificateDoesNotExist(t *testing.T) {
	setup := utils.Setup(t)

	removeIcaCert := types.NewMsgRemoveNocX509RootCert(
		setup.Vendor1.String(),
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		testconstants.NocRootCert1SerialNumber)
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_RemoveNocRootCert_ByOtherVendor(t *testing.T) {
	cases := []CertificateTestCase{
		{
			name:    "OperationalPKI_RemoveNocRootCert_CertificateDoesNotExist",
			crtType: types.CertificateType_OperationalPKI,
		},
		{
			name:    "VIDSignerPKI_RemoveNocRootCert_CertificateDoesNotExist",
			crtType: types.CertificateType_VIDSignerPKI,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			// add NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			// add fist vendor account with VID = 1
			vendorAccAddress1 := utils.GenerateAccAddress()
			setup.AddAccount(vendorAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

			// add second vendor account with VID = 1000
			vendorAccAddress2 := utils.GenerateAccAddress()
			setup.AddAccount(vendorAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID1)

			// remove ICA certificate by second vendor account
			removeIcaCert := types.NewMsgRemoveNocX509RootCert(
				vendorAccAddress2.String(), testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID, testconstants.NocRootCert1SerialNumber)
			_, err := setup.Handler(setup.Ctx, removeIcaCert)
			require.Error(t, err)
			require.True(t, pkitypes.ErrCertVidNotEqualAccountVid.Is(err))
		})
	}
}

func TestHandler_RemoveNocRootCert_SenderNotVendor(t *testing.T) {
	cases := []CertificateTestCase{
		{
			name:    "OperationalPKI_RemoveNocRootCert_SenderNotVendor",
			crtType: types.CertificateType_OperationalPKI,
		},
		{
			name:    "VIDSignerPKI_RemoveNocRootCert_SenderNotVendor",
			crtType: types.CertificateType_VIDSignerPKI,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			// add NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			removeIcaCert := types.NewMsgRemoveNocX509RootCert(
				setup.Trustee1.String(), testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID, "")
			_, err := setup.Handler(setup.Ctx, removeIcaCert)
			require.Error(t, err)
			require.True(t, sdkerrors.ErrUnauthorized.Is(err))
		})
	}
}

func TestHandler_RemoveNocRootCert_InvalidSerialNumber(t *testing.T) {
	cases := []CertificateTestCase{
		{
			name:    "OperationalPKI_RemoveNocRootCert_InvalidSerialNumber",
			crtType: types.CertificateType_OperationalPKI,
		},
		{
			name:    "VIDSignerPKI_RemoveNocRootCert_InvalidSerialNumber",
			crtType: types.CertificateType_VIDSignerPKI,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			// add NOC root certificate
			rootCertificate := utils.RootNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate)

			removeX509Cert := types.NewMsgRemoveNocX509RootCert(
				setup.Vendor1.String(),
				testconstants.NocRootCert1Subject,
				testconstants.NocRootCert1SubjectKeyID,
				"invalid")
			_, err := setup.Handler(setup.Ctx, removeX509Cert)
			require.Error(t, err)
			require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
		})
	}
}

func TestHandler_RemoveNocRootCert_IntermediateCertificate(t *testing.T) {
	cases := []CertificateTestCase{
		{
			name:    "OperationalPKI_RemoveNocRootCert_IntermediateCertificate",
			crtType: types.CertificateType_OperationalPKI,
		},
		{
			name:    "VIDSignerPKI_RemoveNocRootCert_IntermediateCertificate",
			crtType: types.CertificateType_VIDSignerPKI,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			// add NOC root certificates
			rootCertificate1 := utils.RootNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocRootCertificate(setup, rootCertificate1)

			// Add ICA certificates
			icaCertificate := utils.IntermediateNocCertificate1(setup.Vendor1, tc.crtType)
			utils.AddNocIntermediateCertificate(setup, icaCertificate)

			removeX509Cert := types.NewMsgRemoveNocX509RootCert(
				setup.Vendor1.String(),
				icaCertificate.Subject,
				icaCertificate.SubjectKeyId,
				"")
			_, err := setup.Handler(setup.Ctx, removeX509Cert)
			require.Error(t, err)
			require.True(t, pkitypes.ErrInappropriateCertificateType.Is(err))
		})
	}
}

func TestHandler_RemoveNocRootCert_DaCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// add DA root certificate
	rootCertificate := utils.RootDaCertificate(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertificate)

	removeX509Cert := types.NewMsgRemoveNocX509RootCert(
		setup.Vendor1.String(),
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		"")
	_, err := setup.Handler(setup.Ctx, removeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}
