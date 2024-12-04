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

func TestHandler_RemoveNocRootCert(t *testing.T) {
	setup := utils.Setup(t)

	// add NOC root certificates
	rootCertificate := utils.RootNocCertificate1(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate)

	// remove noc root certificate
	utils.RemoveNocRootCertificate(setup, setup.Vendor1, rootCertificate.Subject, rootCertificate.SubjectKeyId, "")

	// Check indexes
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
			{Key: types.NocIcaCertificatesKeyPrefix},
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
			{Key: types.RevokedNocIcaCertificatesKeyPrefix},
			{Key: types.RevokedNocRootCertificatesKeyPrefix},
			{Key: types.RevokedCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)
}

func TestHandler_RemoveNocX509RootCert_BySubjectAndSKID(t *testing.T) {
	setup := utils.Setup(t)

	// add NOC root certificates
	rootCertificate1 := utils.RootNocCertificate1(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate1)

	rootCertificate2 := utils.RootNocCertificate1Copy(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate2)

	// Add intermediate certificate
	icaCertificate := utils.IntermediateNocCertificate1(setup.Vendor1)
	utils.AddNocIntermediateCertificate(setup, icaCertificate)

	// get certificates for further comparison
	nocCerts := setup.Keeper.GetAllNocCertificates(setup.Ctx)
	require.NotNil(t, nocCerts)
	require.Equal(t, 2, len(nocCerts))
	require.Equal(t, 3, len(nocCerts[0].Certs)+len(nocCerts[1].Certs))

	// remove all root noc root certificates
	utils.RemoveNocRootCertificate(setup, setup.Vendor1, rootCertificate1.Subject, rootCertificate1.SubjectKeyId, "")

	// check that only IAC certificate exists
	nocCerts, _ = utils.QueryAllNocCertificates(setup)
	require.Equal(t, 1, len(nocCerts))
	require.Equal(t, 1, len(nocCerts[0].Certs))

	// Check indexes for root certificates
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

	// Check indexes for intermediate certificates
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
}

func TestHandler_RemoveNocX509RootCert_BySerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	// add NOC root certificates
	rootCertificate1 := utils.RootNocCertificate1(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate1)

	rootCertificate2 := utils.RootNocCertificate1Copy(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate2)

	// Add ICA certificates
	icaCertificate := utils.IntermediateNocCertificate1(setup.Vendor1)
	utils.AddNocIntermediateCertificate(setup, icaCertificate)

	// get certificates for further comparison
	nocCerts := setup.Keeper.GetAllNocCertificates(setup.Ctx)
	require.NotNil(t, nocCerts)
	require.Equal(t, 2, len(nocCerts))
	require.Equal(t, 3, len(nocCerts[0].Certs)+len(nocCerts[1].Certs))

	// remove NOC root certificate by serial number
	utils.RemoveNocRootCertificate(
		setup,
		setup.Vendor1,
		rootCertificate1.Subject,
		rootCertificate1.SubjectKeyId,
		rootCertificate1.SerialNumber)

	// check total
	nocCerts, _ = utils.QueryAllNocCertificates(setup)
	require.Equal(t, 2, len(nocCerts))
	require.Equal(t, 2, len(nocCerts[0].Certs)+len(nocCerts[1].Certs))

	// Check indexes for root certificates
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
		},
		Missing: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.RevokedNocRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate1, indexes)

	// same but unique does not exist
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
		},
		Missing: []utils.TestIndex{
			{Key: types.RevokedNocRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate2, indexes)

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

	// Check indexes for intermediate certificates
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
}

func TestHandler_RemoveNocX509RootCert_RevokedCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// add NOC root certificate
	rootCertificate1 := utils.RootNocCertificate1(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate1)

	rootCertificate2 := utils.RootNocCertificate1Copy(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate2)

	// Add an intermediate certificate
	icaCertificate := utils.IntermediateNocCertificate1(setup.Vendor1)
	utils.AddNocIntermediateCertificate(setup, icaCertificate)

	// revoke NOC root certificates
	utils.RevokeNocRootCertificate(
		setup,
		setup.Vendor1,
		rootCertificate2.Subject,
		rootCertificate2.SubjectKeyId,
		"",
		false,
	)

	// Check indexes for root certificates
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
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate1, indexes)
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate2, indexes)

	// Check indexes for intermediate certificate
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

	// remove NOC root certificates
	utils.RemoveNocRootCertificate(
		setup,
		setup.Vendor1,
		rootCertificate2.Subject,
		rootCertificate2.SubjectKeyId,
		"",
	)

	// Check indexes for root certificates
	indexes = utils.TestIndexes{
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

	// Check that intermediate certificates still exist
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
}

func TestHandler_RemoveNocX509RootCert_RevokedWithChildCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// add NOC root certificate
	rootCertificate := utils.RootNocCertificate1(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate)

	// Add an intermediate certificate
	icaCertificate := utils.IntermediateNocCertificate1(setup.Vendor1)
	utils.AddNocIntermediateCertificate(setup, icaCertificate)

	// revoke NOC root certificates
	utils.RevokeNocRootCertificate(
		setup,
		setup.Vendor1,
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		"",
		true,
	)

	// Check indexes for root certificates
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
			{Key: types.RevokedNocIcaCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)

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
			{Key: types.RevokedNocRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate, indexes)

	// remove NOC root certificates
	utils.RemoveNocRootCertificate(
		setup,
		setup.Vendor1,
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		"",
	)

	// Check indexes for root certificates
	indexes = utils.TestIndexes{
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
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)

	// Check that intermediate certificates still is revoked
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
			{Key: types.RevokedNocRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate, indexes)
}

// Extra cases

func TestHandler_RemoveNocX509RootCert_RevokedAndActiveCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// add NOC root certificate
	rootCertificate := utils.RootNocCertificate1(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate)

	// Add an intermediate certificate
	icaCertificate := utils.IntermediateNocCertificate1(setup.Vendor1)
	utils.AddNocIntermediateCertificate(setup, icaCertificate)

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
	rootCertificate2 := utils.RootNocCertificate1Copy(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate2)

	// remove NOC root certificate by serial number
	utils.RemoveNocRootCertificate(
		setup,
		setup.Vendor1,
		rootCertificate.Subject,
		rootCertificate.SubjectKeyId,
		rootCertificate.SerialNumber,
	)

	// Check indexes for re-activated root certificates
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
			{Key: types.UniqueCertificateKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.RevokedNocRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate2, indexes)

	// Check indexes for deleted root certificates
	indexes = utils.TestIndexes{
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
			{Key: types.RevokedNocRootCertificatesKeyPrefix},
			{Key: types.UniqueCertificateKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)

	// remove NOC root certificates
	utils.RemoveNocRootCertificate(
		setup,
		setup.Vendor1,
		rootCertificate2.Subject,
		rootCertificate2.SubjectKeyId,
		"",
	)

	// Check indexes for root certificates (after deletion re-activated)
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
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.RevokedNocRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate2, indexes)
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)
}

func TestHandler_RemoveNocX509RootCert_ByNotOwnerButSameVendor(t *testing.T) {
	setup := utils.Setup(t)

	// add NOC root certificate
	rootCertificate := utils.RootNocCertificate1(setup.Vendor1)
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

	// Check indexes for root certificates
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
}

// Error cases
func TestHandler_RemoveNocX509RootCert_CertificateDoesNotExist(t *testing.T) {
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

func TestHandler_RemoveNocX509RootCert_EmptyCertificatesList(t *testing.T) {
	setup := utils.Setup(t)

	setup.Keeper.SetNocRootCertificates(
		setup.Ctx,
		types.NocRootCertificates{
			Vid: testconstants.Vid,
		},
	)

	removeIcaCert := types.NewMsgRemoveNocX509RootCert(
		setup.Vendor1.String(),
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		"")
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_RemoveNocX509RootCert_ByOtherVendor(t *testing.T) {
	setup := utils.Setup(t)

	// add NOC root certificate
	rootCertificate := utils.RootNocCertificate1(setup.Vendor1)
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
}

func TestHandler_RemoveNocX509RootCert_SenderNotVendor(t *testing.T) {
	setup := utils.Setup(t)

	// add NOC root certificate
	rootCertificate := utils.RootNocCertificate1(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate)

	removeIcaCert := types.NewMsgRemoveNocX509RootCert(
		setup.Trustee1.String(), testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID, "")
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_RemoveNocX509RootCert_InvalidSerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	// add NOC root certificate
	rootCertificate := utils.RootNocCertificate1(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate)

	removeX509Cert := types.NewMsgRemoveNocX509RootCert(
		setup.Vendor1.String(),
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		"invalid")
	_, err := setup.Handler(setup.Ctx, removeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}
