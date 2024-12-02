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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Main

func TestHandler_RemoveNocRootCert(t *testing.T) {
	setup := utils.Setup(t)

	// add NOC root certificates
	rootCertificate := utils.CreateTestNocRoot1Cert()
	utils.AddNocRootCertificate(setup, setup.Vendor1, rootCertificate.PEM)

	// remove noc root certificate
	utils.RemoveNocRootCertificate(setup, setup.Vendor1, rootCertificate.Subject, rootCertificate.SubjectKeyID, "")

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
	rootCertificate1 := utils.CreateTestNocRoot1Cert()
	utils.AddNocRootCertificate(setup, setup.Vendor1, rootCertificate1.PEM)

	rootCertificate2 := utils.CreateTestNocRoot2Cert()
	utils.AddNocRootCertificate(setup, setup.Vendor1, rootCertificate2.PEM)

	// Add intermediate certificate
	icaCertificate := utils.CreateTestNocIca1Cert()
	utils.AddNocIntermediateCertificate(setup, setup.Vendor1, icaCertificate.PEM)

	// get certificates for further comparison
	nocCerts := setup.Keeper.GetAllNocCertificates(setup.Ctx)
	require.NotNil(t, nocCerts)
	require.Equal(t, 2, len(nocCerts))
	require.Equal(t, 3, len(nocCerts[0].Certs)+len(nocCerts[1].Certs))

	// remove all root nOC certificates but IAC certificate
	utils.RemoveNocRootCertificate(setup, setup.Vendor1, rootCertificate1.Subject, rootCertificate1.SubjectKeyID, "")

	// check that only IAC certificate exists
	nocCerts, _ = utils.QueryAllNocCertificates(setup)
	require.Equal(t, 1, len(nocCerts))
	require.Equal(t, 1, len(nocCerts[0].Certs))
	require.Equal(t, testconstants.NocCert1SerialNumber, nocCerts[0].Certs[0].SerialNumber)

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
	rootCertificate1 := utils.CreateTestNocRoot1Cert()
	utils.AddNocRootCertificate(setup, setup.Vendor1, rootCertificate1.PEM)

	rootCertificate2 := utils.CreateTestNocRoot2Cert()
	utils.AddNocRootCertificate(setup, setup.Vendor1, rootCertificate2.PEM)

	// Add ICA certificates
	icaCertificate := utils.CreateTestNocIca1Cert()
	utils.AddNocIntermediateCertificate(setup, setup.Vendor1, icaCertificate.PEM)

	// remove NOC root certificate by serial number
	utils.RemoveNocRootCertificate(setup, setup.Vendor1, rootCertificate1.Subject, rootCertificate1.SubjectKeyID, rootCertificate1.SerialNumber)

	nocCerts, _ := utils.QueryAllNocCertificates(setup)
	require.Equal(t, 2, len(nocCerts))

	// NocCertificates: Subject and SKID
	nocCertificates, err := utils.QueryNocCertificates(
		setup,
		rootCertificate2.Subject,
		rootCertificate2.SubjectKeyID,
	)
	require.NoError(t, err)
	require.Equal(t, 1, len(nocCertificates.Certs))

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
	utils.RemoveNocRootCertificate(setup, setup.Vendor1, rootCertificate2.Subject, rootCertificate2.SubjectKeyID, rootCertificate2.SerialNumber)

	nocCerts, _ = utils.QueryAllNocCertificates(setup)
	require.Equal(t, 1, len(nocCerts))
	require.Equal(t, 1, len(nocCerts[0].Certs))
	require.Equal(t, testconstants.NocCert1SerialNumber, nocCerts[0].Certs[0].SerialNumber)

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
	rootCertificate1 := utils.CreateTestNocRoot1Cert()
	utils.AddNocRootCertificate(setup, setup.Vendor1, testconstants.NocRootCert1)

	rootCertificate2 := utils.CreateTestNocRoot2Cert()
	utils.AddNocRootCertificate(setup, setup.Vendor1, testconstants.NocRootCert1Copy)

	// Add an intermediate certificate
	icaCertificate := utils.CreateTestNocIca1Cert()
	utils.AddNocIntermediateCertificate(setup, setup.Vendor1, testconstants.NocCert1)

	// revoke NOC root certificates
	utils.RevokeNocRootCertificate(
		setup,
		setup.Vendor1,
		rootCertificate2.Subject,
		rootCertificate2.SubjectKeyID,
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

	revokedCerts, _ := utils.QueryNocRevokedRootCertificates(setup, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(t, 2, len(revokedCerts.Certs))

	// Check that intermediate certificates does not exist
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
		rootCertificate2.SubjectKeyID,
		"",
	)

	allCerts, _ := utils.QueryAllNocCertificates(setup)
	require.Equal(t, 1, len(allCerts))
	require.Equal(t, testconstants.NocCert1SerialNumber, allCerts[0].Certs[0].SerialNumber)

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

// Extra cases

func TestHandler_RemoveNocX509RootCert_RevokedAndActiveCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// add NOC root certificate
	utils.AddNocRootCertificate(setup, setup.Vendor1, testconstants.NocRootCert1)

	// Add an intermediate certificate
	addIcaCert := types.NewMsgAddNocX509IcaCert(setup.Vendor1.String(), testconstants.NocCert1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addIcaCert)
	require.NoError(t, err)

	// get certificates for further comparison
	nocCerts := setup.Keeper.GetAllNocCertificates(setup.Ctx)
	require.NotNil(t, nocCerts)
	require.Equal(t, 2, len(nocCerts))

	// revoke an intermediate certificate
	utils.RevokeNocRootCertificate(
		setup,
		setup.Vendor1,
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		testconstants.NocRootCert1SerialNumber,
		false,
	)

	// Add NOC root certificate with new serial number
	utils.AddNocRootCertificate(setup, setup.Vendor1, testconstants.NocRootCert1Copy)

	certs, _ := utils.QueryNocCertificates(setup, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(t, 1, len(certs.Certs))
	require.Equal(t, testconstants.NocRootCert1CopySerialNumber, certs.Certs[0].SerialNumber)

	// remove NOC root certificate by serial number
	removeIcaCert := types.NewMsgRemoveNocX509RootCert(
		setup.Vendor1.String(),
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		testconstants.NocRootCert1SerialNumber,
	)
	_, err = setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	// check that only one root and IAC certificates exists
	nocCerts, _ = utils.QueryAllNocCertificates(setup)
	require.Equal(t, 2, len(nocCerts))

	certs, _ = utils.QueryNocCertificates(setup, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(t, testconstants.NocRootCert1CopySerialNumber, certs.Certs[0].SerialNumber)
	certs, _ = utils.QueryNocCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, 1, len(certs.Certs))

	_, err = utils.QueryNocRevokedRootCertificates(setup, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificates does not exists
	found := setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SerialNumber)
	require.Equal(t, false, found)
	found = setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocRootCert1Subject, testconstants.NocRootCert1CopySerialNumber)
	require.Equal(t, true, found)

	// query noc certificate by VID
	nocCertificates, err := utils.QueryNocIcaCertificatesByVid(setup, testconstants.Vid)
	require.NoError(t, err)
	require.Equal(t, len(nocCertificates.Certs), 1)
	require.Equal(t, testconstants.NocCert1SerialNumber, nocCertificates.Certs[0].SerialNumber)

	// Add NOC root certificate with new serial number
	utils.AddNocRootCertificate(setup, setup.Vendor1, testconstants.NocRootCert1)

	certs, _ = utils.QueryNocCertificates(setup, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(t, 2, len(certs.Certs))

	// remove NOC root certificates
	removeIcaCert = types.NewMsgRemoveNocX509RootCert(
		setup.Vendor1.String(),
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		"",
	)
	_, err = setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	nocCerts, _ = utils.QueryAllNocCertificates(setup)
	require.Equal(t, 1, len(nocCerts))
	require.Equal(t, 1, len(nocCerts[0].Certs))
	require.Equal(t, testconstants.NocCert1SerialNumber, nocCerts[0].Certs[0].SerialNumber)

	nocCertificates, err = utils.QueryNocIcaCertificatesByVid(setup, testconstants.Vid)
	require.NoError(t, err)
	require.Equal(t, len(nocCertificates.Certs), 1)
	require.Equal(t, testconstants.NocCert1SerialNumber, nocCertificates.Certs[0].SerialNumber)

	// check that IAC certificates can be queried by vid+skid
	certsByVidSkid, _ := utils.QueryNocCertificatesByVidAndSkid(setup, testconstants.Vid, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, 1, len(certsByVidSkid.Certs))
	require.Equal(t, testconstants.NocCert1SerialNumber, certsByVidSkid.Certs[0].SerialNumber)

	// check that root certs removed
	_, err = utils.QueryNocCertificates(setup, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))
	_, err = utils.QueryNocCertificatesBySubject(setup, testconstants.NocRootCert1Subject)
	require.Equal(t, codes.NotFound, status.Code(err))
	certsBySKID, _ := utils.QueryNocCertificatesBySubjectKeyID(setup, testconstants.NocRootCert1SubjectKeyID)
	require.Empty(t, certsBySKID)
	_, err = utils.QueryNocRootCertificatesByVid(setup, testconstants.Vid)
	require.Equal(t, codes.NotFound, status.Code(err))
	_, err = utils.QueryNocRootCertificatesByVid(setup, testconstants.Vid)
	require.Equal(t, codes.NotFound, status.Code(err))
	_, err = utils.QueryNocCertificatesByVidAndSkid(setup, testconstants.Vid, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificates does not exists
	found = setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SerialNumber)
	require.Equal(t, false, found)
	found = setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocRootCert1Subject, testconstants.NocRootCert1CopySerialNumber)
	require.Equal(t, false, found)
}

func TestHandler_RemoveNocX509RootCert_ByNotOwnerButSameVendor(t *testing.T) {
	setup := utils.Setup(t)

	// add NOC root certificate
	utils.AddNocRootCertificate(setup, setup.Vendor1, testconstants.NocRootCert1)

	// add first vendor account with VID = 1
	vendorAccAddress1 := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add second vendor account with VID = 1
	vendorAccAddress2 := utils.GenerateAccAddress()
	setup.AddAccount(vendorAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// remove x509 certificate by second vendor account
	removeIcaCert := types.NewMsgRemoveNocX509RootCert(
		vendorAccAddress2.String(),
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		testconstants.NocRootCert1SerialNumber,
	)
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.NoError(t, err)

	// check that certificate removed from 'noc certificates' list
	_, err = utils.QueryNocCertificates(setup, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that certificate removed from 'noc certificates by subject' list
	_, err = utils.QueryNocCertificatesBySubject(setup, testconstants.NocRootCert1Subject)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that certificate removed from 'noc certificates by SKID' list
	nocCerts, err := utils.QueryNocCertificatesBySubjectKeyID(setup, testconstants.NocRootCert1SubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 0, len(nocCerts))

	// query noc certificate by VID
	_, err = utils.QueryNocRootCertificatesByVid(setup, testconstants.Vid)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificate key is not registered
	require.False(t, setup.Keeper.IsUniqueCertificatePresent(setup.Ctx,
		testconstants.NocRootCert1Subject, testconstants.NocRootCert1SerialNumber))
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
	utils.AddNocRootCertificate(setup, setup.Vendor1, testconstants.NocRootCert1)

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
	utils.AddNocRootCertificate(setup, setup.Vendor1, testconstants.NocRootCert1)

	removeIcaCert := types.NewMsgRemoveNocX509RootCert(
		setup.Trustee1.String(), testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID, "")
	_, err := setup.Handler(setup.Ctx, removeIcaCert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_RemoveNocX509RootCert_InvalidSerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	// add NOC root certificate
	utils.AddNocRootCertificate(setup, setup.Vendor1, testconstants.NocRootCert1)

	removeX509Cert := types.NewMsgRemoveNocX509RootCert(
		setup.Vendor1.String(),
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		"invalid")
	_, err := setup.Handler(setup.Ctx, removeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}
