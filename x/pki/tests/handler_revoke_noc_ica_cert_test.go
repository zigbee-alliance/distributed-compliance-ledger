package tests

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/tests/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// Main

func TestHandler_RevokeNocIntermediateCert(t *testing.T) {
	setup := utils.Setup(t)

	accAddress := setup.CreateVendorAccount(testconstants.Vid)

	// add the first NOC root certificate
	utils.AddNocRootCertificate(setup, accAddress, testconstants.NocRootCert1)

	// add the NOC non-root certificate
	icaCertificate := utils.CreateTestNocIca1Cert()
	utils.AddNocIntermediateCertificate(setup, accAddress, testconstants.NocCert1)

	// Revoke NOC with subject and subject key id only
	revokeCert := types.NewMsgRevokeNocX509IcaCert(
		accAddress.String(),
		icaCertificate.Subject,
		icaCertificate.SubjectKeyID,
		"",
		testconstants.Info,
		false,
	)
	_, err := setup.Handler(setup.Ctx, revokeCert)
	require.NoError(t, err)

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

	accAddress := utils.GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add the first NOC root certificate
	utils.AddNocRootCertificate(setup, accAddress, testconstants.NocRootCert1)

	// add the first NOC non-root certificate
	icaCertificate1 := utils.CreateTestNocIca1Cert()
	utils.AddNocIntermediateCertificate(setup, accAddress, testconstants.NocCert1)

	// add the second NOC non-root certificate
	icaCertificate2 := utils.CreateTestNocIca1CertCopy()
	utils.AddNocIntermediateCertificate(setup, accAddress, testconstants.NocCert1Copy)

	// add the NOC leaf certificate
	leafCertificate := utils.CreateTestNocLeafCert()
	utils.AddNocIntermediateCertificate(setup, accAddress, testconstants.NocLeafCert1)

	// Revoke NOC with subject and subject key id only
	revokeCert := types.NewMsgRevokeNocX509IcaCert(
		accAddress.String(),
		icaCertificate1.Subject,
		icaCertificate1.SubjectKeyID,
		"",
		testconstants.Info,
		false,
	)
	_, err := setup.Handler(setup.Ctx, revokeCert)
	require.NoError(t, err)

	// Check that revoked certificates exist
	revokedNocCerts, err := utils.QueryNocRevokedIcaCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 2, len(revokedNocCerts.Certs))
	require.Equal(t, testconstants.NocCert1Subject, revokedNocCerts.Subject)
	require.Equal(t, testconstants.NocCert1SubjectKeyID, revokedNocCerts.SubjectKeyId)

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

	accAddress := utils.GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add the first NOC root certificate
	utils.AddNocRootCertificate(setup, accAddress, testconstants.NocRootCert1)

	// add the first NOC non-root certificate
	icaCertificate1 := utils.CreateTestNocIca1Cert()
	utils.AddNocIntermediateCertificate(setup, accAddress, testconstants.NocCert1)

	// add the second NOC non-root certificate
	icaCertificate2 := utils.CreateTestNocIca1CertCopy()
	utils.AddNocIntermediateCertificate(setup, accAddress, testconstants.NocCert1Copy)

	// add the NOC leaf certificate
	leafCertificate := utils.CreateTestNocLeafCert()
	utils.AddNocIntermediateCertificate(setup, accAddress, testconstants.NocLeafCert1)

	// Revoke noc with subject and subject key id and its child too
	revokeCert := types.NewMsgRevokeNocX509IcaCert(
		accAddress.String(),
		icaCertificate1.Subject,
		icaCertificate1.SubjectKeyID,
		"",
		testconstants.Info,
		true,
	)
	_, err := setup.Handler(setup.Ctx, revokeCert)
	require.NoError(t, err)

	allRevokedCerts, err := utils.QueryAllNocRevokedIcaCertificates(setup)
	require.NoError(t, err)
	require.Equal(t, 2, len(allRevokedCerts))
	require.Equal(t, 3, len(allRevokedCerts[0].Certs)+len(allRevokedCerts[1].Certs))

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

	accAddress := utils.GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add the first NOC root certificate
	utils.AddNocRootCertificate(setup, accAddress, testconstants.NocRootCert1)

	// add the first NOC non-root certificate
	utils.AddNocIntermediateCertificate(setup, accAddress, testconstants.NocCert1)

	// add the second NOC non-root certificate
	utils.AddNocIntermediateCertificate(setup, accAddress, testconstants.NocCert1Copy)

	// add the NOC leaf certificate
	utils.AddNocIntermediateCertificate(setup, accAddress, testconstants.NocLeafCert1)

	// Revoke NOC by serial number only
	revokeCert := types.NewMsgRevokeNocX509IcaCert(
		accAddress.String(),
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1SerialNumber,
		testconstants.Info,
		false,
	)
	_, err := setup.Handler(setup.Ctx, revokeCert)
	require.NoError(t, err)

	revokedNocCerts, err := utils.QueryNocRevokedIcaCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 1, len(revokedNocCerts.Certs))
	require.Equal(t, testconstants.NocCert1SerialNumber, revokedNocCerts.Certs[0].SerialNumber)

	// Child certificate should not be revoked
	_, err = utils.QueryNocRevokedIcaCertificates(setup, testconstants.NocLeafCert1Subject, testconstants.NocLeafCert1SubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))

	// query NOC certificate by Subject
	certsBySubject, err := utils.QueryNocCertificatesBySubject(setup, testconstants.NocCert1Subject)
	require.NoError(t, err)
	require.Equal(t, 1, len(certsBySubject.SubjectKeyIds))

	// query NOC certificate by Subject Key ID
	aprCertsBySubjectKeyID, _ := utils.QueryNocCertificatesBySubjectKeyID(setup, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, 1, len(aprCertsBySubjectKeyID))
	require.Equal(t, 1, len(aprCertsBySubjectKeyID[0].Certs))
	require.Equal(t, testconstants.NocCert1CopySerialNumber, aprCertsBySubjectKeyID[0].Certs[0].SerialNumber)

	// query noc certificate by VID
	nocCerts, err := utils.QueryNocIcaCertificatesByVid(setup, testconstants.Vid)
	require.NoError(t, err)
	require.Equal(t, 2, len(nocCerts.Certs))
	require.NotEqual(t, testconstants.NocCert1SerialNumber, nocCerts.Certs[0].SerialNumber)
	require.NotEqual(t, testconstants.NocCert1SerialNumber, nocCerts.Certs[1].SerialNumber)

	// query all certs
	certs, err := utils.QueryAllNocCertificates(setup)
	require.NoError(t, err)
	require.Equal(t, 3, len(certs))
	require.NotEqual(t, testconstants.NocCert1SerialNumber, certs[0].Certs[0].SerialNumber)
	require.NotEqual(t, testconstants.NocCert1SerialNumber, certs[1].Certs[0].SerialNumber)
	require.NotEqual(t, testconstants.NocCert1SerialNumber, certs[2].Certs[0].SerialNumber)

	// query noc certificate, cert with different serial number should not be removed
	noccCerts, _ := utils.QueryNocCertificates(setup, testconstants.NocCert1CopySubject, testconstants.NocCert1CopySubjectKeyID)
	require.Equal(t, 1, len(noccCerts.Certs))
	require.Equal(t, testconstants.NocCert1CopySerialNumber, noccCerts.Certs[0].SerialNumber)

	// query child certificate, they should not be removed
	childCerts, _ := utils.QueryNocCertificates(setup, testconstants.NocLeafCert1Subject, testconstants.NocLeafCert1SubjectKeyID)
	require.Equal(t, 1, len(childCerts.Certs))
	require.Equal(t, testconstants.NocLeafCert1SubjectKeyID, childCerts.SubjectKeyId)

	// check that leaf certificate can be queried by vid+skid
	certsByVidSkid, _ := utils.QueryNocCertificatesByVidAndSkid(setup, testconstants.Vid, testconstants.NocLeafCert1SubjectKeyID)
	require.Equal(t, 1, len(certsByVidSkid.Certs))
	require.Equal(t, testconstants.NocLeafCert1SerialNumber, certsByVidSkid.Certs[0].SerialNumber)

	// check that unique certificate key is removed
	require.False(t,
		setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocCert1, testconstants.NocCert1SerialNumber))
}

func TestHandler_RevokeNocX509Cert_RevokeBySerialNumberAndWithChild(t *testing.T) {
	setup := utils.Setup(t)

	accAddress := utils.GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add the first NOC root certificate
	addNocX509RootCert := types.NewMsgAddNocX509RootCert(accAddress.String(), testconstants.NocRootCert1, testconstants.SchemaVersion)
	_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
	require.NoError(t, err)

	// add the first NOC non-root certificate
	addNocX509Cert := types.NewMsgAddNocX509IcaCert(accAddress.String(), testconstants.NocCert1, testconstants.SchemaVersion)
	_, err = setup.Handler(setup.Ctx, addNocX509Cert)
	require.NoError(t, err)

	// add the second NOC non-root certificate
	addNocX509Cert = types.NewMsgAddNocX509IcaCert(accAddress.String(), testconstants.NocCert1Copy, testconstants.SchemaVersion)
	_, err = setup.Handler(setup.Ctx, addNocX509Cert)
	require.NoError(t, err)

	// add the NOC leaf certificate
	addNocX509Cert = types.NewMsgAddNocX509IcaCert(accAddress.String(), testconstants.NocLeafCert1, testconstants.SchemaVersion)
	_, err = setup.Handler(setup.Ctx, addNocX509Cert)
	require.NoError(t, err)

	// Revoke NOC with subject and subject key id and its child too
	revokeCert := types.NewMsgRevokeNocX509IcaCert(
		accAddress.String(),
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1SerialNumber,
		testconstants.Info,
		true,
	)
	_, err = setup.Handler(setup.Ctx, revokeCert)
	require.NoError(t, err)

	allRevokedCerts, err := utils.QueryAllNocRevokedIcaCertificates(setup)
	require.NoError(t, err)
	require.Equal(t, 2, len(allRevokedCerts))
	require.Equal(t, 2, len(allRevokedCerts[0].Certs)+len(allRevokedCerts[1].Certs))

	revokedNocCerts, err := utils.QueryNocRevokedIcaCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 1, len(revokedNocCerts.Certs))
	require.Equal(t, testconstants.NocCert1SerialNumber, revokedNocCerts.Certs[0].SerialNumber)

	// Child certificate should be revoked
	revokedNocCerts, err = utils.QueryNocRevokedIcaCertificates(setup, testconstants.NocLeafCert1Subject, testconstants.NocLeafCert1SubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 1, len(revokedNocCerts.Certs))
	require.Equal(t, testconstants.NocLeafCert1SerialNumber, revokedNocCerts.Certs[0].SerialNumber)

	// query child of revoked certificate, they should be revoked
	_, err = utils.QueryNocCertificates(setup, testconstants.NocLeafCert1Subject, testconstants.NocLeafCert1SubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// query all certs
	certs, err := utils.QueryAllNocCertificates(setup)
	require.NoError(t, err)
	require.Equal(t, 2, len(certs))
	require.NotEqual(t, testconstants.NocCert1SerialNumber, certs[0].Certs[0].SerialNumber)
	require.NotEqual(t, testconstants.NocCert1SerialNumber, certs[1].Certs[0].SerialNumber)

	// query noc certificates
	aprCerts, err := utils.QueryNocCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1CopySubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 1, len(aprCerts.Certs))
	require.Equal(t, testconstants.NocCert1CopySerialNumber, aprCerts.Certs[0].SerialNumber)

	// query noc certificate by Subject
	certsBySubject, err := utils.QueryNocCertificatesBySubject(setup, testconstants.NocCert1Subject)
	require.NoError(t, err)
	require.Equal(t, 1, len(certsBySubject.SubjectKeyIds))

	_, err = utils.QueryNocCertificatesBySubject(setup, testconstants.NocLeafCert1Subject)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// query noc certificate by Subject Key ID
	aprCertsBySubjectKeyID, _ := utils.QueryNocCertificatesBySubjectKeyID(setup, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, 1, len(aprCertsBySubjectKeyID))
	require.Equal(t, testconstants.NocCert1CopySerialNumber, aprCertsBySubjectKeyID[0].Certs[0].SerialNumber)

	aprCertsBySubjectKeyID, _ = utils.QueryNocCertificatesBySubjectKeyID(setup, testconstants.NocLeafCert1SubjectKeyID)
	require.Equal(t, 0, len(aprCertsBySubjectKeyID))

	// query noc certificate by VID
	nocCerts, err := utils.QueryNocIcaCertificatesByVid(setup, testconstants.Vid)
	require.NoError(t, err)
	require.Equal(t, 1, len(nocCerts.Certs))
	require.Equal(t, testconstants.NocCert1CopySerialNumber, nocCerts.Certs[0].SerialNumber)

	// check that leaf certificate can be queried by vid+skid
	certsByVidSkid, _ := utils.QueryNocCertificatesByVidAndSkid(setup, testconstants.Vid, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, 1, len(certsByVidSkid.Certs))
	require.Equal(t, testconstants.NocCert1CopySerialNumber, certsByVidSkid.Certs[0].SerialNumber)

	// check that unique certificate key is removed
	require.False(t,
		setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocCert1, testconstants.NocCert1SerialNumber))
}

// Extra cases

// Error cases

func TestHandler_RevokeNocX509Cert_SenderNotVendor(t *testing.T) {
	setup := utils.Setup(t)

	accAddress := utils.GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add the new NOC root certificate
	addNocX509RootCert := types.NewMsgAddNocX509RootCert(accAddress.String(), testconstants.NocRootCert1, testconstants.CertSchemaVersion)
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

	accAddress := utils.GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	revokeCert := types.NewMsgRevokeNocX509IcaCert(
		accAddress.String(),
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
