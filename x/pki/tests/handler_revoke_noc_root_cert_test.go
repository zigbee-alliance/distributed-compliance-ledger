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

func TestHandler_RevokeNoRootCert(t *testing.T) {
	setup := utils.Setup(t)

	// add the first NOC root certificate
	rootCertificate := utils.CreateTestNocRoot1Cert()
	addNocX509RootCert := types.NewMsgAddNocX509RootCert(
		setup.Vendor1.String(),
		testconstants.NocRootCert1,
		testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
	require.NoError(t, err)

	// Revoke NOC root with subject and subject key id only
	revokeCert := types.NewMsgRevokeNocX509RootCert(
		setup.Vendor1.String(),
		rootCertificate.Subject,
		rootCertificate.SubjectKeyID,
		"",
		testconstants.Info,
		false,
	)
	_, err = setup.Handler(setup.Ctx, revokeCert)
	require.NoError(t, err)

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

func TestHandler_RevokeNocX509RootCert_RevokeDefault(t *testing.T) {
	setup := utils.Setup(t)

	// add the first NOC root certificate
	addNocX509RootCert := types.NewMsgAddNocX509RootCert(setup.Vendor1.String(), testconstants.NocRootCert1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
	require.NoError(t, err)

	// add the second NOC root certificate
	addNocX509RootCert = types.NewMsgAddNocX509RootCert(setup.Vendor1.String(), testconstants.NocRootCert1Copy, testconstants.CertSchemaVersion)
	_, err = setup.Handler(setup.Ctx, addNocX509RootCert)
	require.NoError(t, err)

	// add the third NOC root certificate
	addNocX509RootCert = types.NewMsgAddNocX509RootCert(setup.Vendor1.String(), testconstants.NocRootCert2, testconstants.CertSchemaVersion)
	_, err = setup.Handler(setup.Ctx, addNocX509RootCert)
	require.NoError(t, err)

	// add the first NOC non-root certificate
	addNocX509Cert := types.NewMsgAddNocX509IcaCert(setup.Vendor1.String(), testconstants.NocCert1, testconstants.CertSchemaVersion)
	_, err = setup.Handler(setup.Ctx, addNocX509Cert)
	require.NoError(t, err)

	// add the second NOC non-root certificate
	addNocX509Cert = types.NewMsgAddNocX509IcaCert(setup.Vendor1.String(), testconstants.NocCert2, testconstants.CertSchemaVersion)
	_, err = setup.Handler(setup.Ctx, addNocX509Cert)
	require.NoError(t, err)

	// Revoke NOC root with subject and subject key id only
	revokeCert := types.NewMsgRevokeNocX509RootCert(
		setup.Vendor1.String(),
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		"",
		testconstants.Info,
		false,
	)
	_, err = setup.Handler(setup.Ctx, revokeCert)
	require.NoError(t, err)

	// query all certs
	certs, err := utils.QueryAllNocCertificates(setup)
	require.NoError(t, err)
	require.Equal(t, 3, len(certs))
	require.NotEqual(t, testconstants.NocRootCert1SubjectKeyID, certs[0].SubjectKeyId)
	require.NotEqual(t, testconstants.NocRootCert1SubjectKeyID, certs[1].SubjectKeyId)
	require.NotEqual(t, testconstants.NocRootCert1SubjectKeyID, certs[2].SubjectKeyId)

	revokedNocCerts, err := utils.QueryNocRevokedRootCertificates(setup, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 2, len(revokedNocCerts.Certs))
	require.Equal(t, testconstants.NocRootCert1Subject, revokedNocCerts.Subject)
	require.Equal(t, testconstants.NocRootCert1SubjectKeyID, revokedNocCerts.SubjectKeyId)

	// query that noc root certificate is not added to x509 revoked root certs
	revokedRootCerts, _ := utils.QueryRevokedRootCertificates(setup)
	require.Equal(t, 0, len(revokedRootCerts.Certs))

	// query noc root certificate by Subject
	_, err = utils.QueryNocCertificatesBySubject(setup, testconstants.NocRootCert1Subject)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// query noc root certificate by Subject Key ID
	aprCertsBySubjectKeyID, _ := utils.QueryNocCertificatesBySubjectKeyID(setup, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(t, 0, len(aprCertsBySubjectKeyID))

	// query noc root certificate by VID
	nocRootCerts, err := utils.QueryNocRootCertificatesByVid(setup, testconstants.Vid)
	require.NoError(t, err)
	require.Equal(t, 1, len(nocRootCerts.Certs))
	require.Equal(t, testconstants.NocRootCert2SubjectKeyID, nocRootCerts.Certs[0].SubjectKeyId)

	// query noc certificate by VID and SKID
	_, err = utils.QueryNocCertificatesByVidAndSkid(setup, testconstants.Vid, testconstants.NocRootCert1SubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	nocCertificatesByVidAndSkid, err := utils.QueryNocCertificatesByVidAndSkid(setup, testconstants.Vid, testconstants.NocRootCert2SubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, testconstants.NocRootCert2SubjectKeyID, nocCertificatesByVidAndSkid.SubjectKeyId)
	require.Equal(t, 1, len(nocRootCerts.Certs))
	require.Equal(t, float32(1), nocCertificatesByVidAndSkid.Tq)

	// Child certificate should not be revoked
	_, err = utils.QueryNocRevokedIcaCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))

	// query child of revoked certificate, they should not be revoked
	childCerts, _ := utils.QueryNocCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, 1, len(childCerts.Certs))
	require.Equal(t, testconstants.NocCert1SubjectKeyID, childCerts.SubjectKeyId)

	// check that child cert is not removed
	nocCerts, err := utils.QueryNocIcaCertificatesByVid(setup, testconstants.Vid)
	require.NoError(t, err)
	require.Equal(t, 2, len(nocCerts.Certs))
	require.Equal(t, testconstants.NocCert1SubjectKeyID, nocCerts.Certs[0].SubjectKeyId)

	// check that unique certificate key is removed
	require.False(t,
		setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocRootCert1, testconstants.NocRootCert1SerialNumber))
	require.False(t,
		setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocRootCert1, testconstants.NocRootCert1CopySerialNumber))
}

func TestHandler_RevokeNocX509RootCert_RevokeWithChild(t *testing.T) {
	setup := utils.Setup(t)

	// add the first NOC root certificate
	addNocX509RootCert := types.NewMsgAddNocX509RootCert(setup.Vendor1.String(), testconstants.NocRootCert1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
	require.NoError(t, err)

	// add the second NOC root certificate
	addNocX509RootCert = types.NewMsgAddNocX509RootCert(setup.Vendor1.String(), testconstants.NocRootCert1Copy, testconstants.CertSchemaVersion)
	_, err = setup.Handler(setup.Ctx, addNocX509RootCert)
	require.NoError(t, err)

	// add the first NOC non-root certificate
	addNocX509Cert := types.NewMsgAddNocX509IcaCert(setup.Vendor1.String(), testconstants.NocCert1, testconstants.CertSchemaVersion)
	_, err = setup.Handler(setup.Ctx, addNocX509Cert)
	require.NoError(t, err)

	// Revoke NOC root with subject and subject key id and its child too
	revokeCert := types.NewMsgRevokeNocX509RootCert(
		setup.Vendor1.String(),
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		"",
		testconstants.Info,
		true,
	)
	_, err = setup.Handler(setup.Ctx, revokeCert)
	require.NoError(t, err)

	// query all certs
	certs, err := utils.QueryAllNocCertificates(setup)
	require.NoError(t, err)
	require.Equal(t, 0, len(certs))

	revokedNocCerts, err := utils.QueryNocRevokedRootCertificates(setup, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 2, len(revokedNocCerts.Certs))
	require.Equal(t, testconstants.NocRootCert1Subject, revokedNocCerts.Subject)
	require.Equal(t, testconstants.NocRootCert1SubjectKeyID, revokedNocCerts.SubjectKeyId)

	// query that noc root certificate is not added to x509 revoked root certs
	revokedRootCerts, _ := utils.QueryRevokedRootCertificates(setup)
	require.Equal(t, 0, len(revokedRootCerts.Certs))

	// query noc root certificate by Subject
	_, err = utils.QueryNocCertificatesBySubject(setup, testconstants.NocRootCert1Subject)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// query child noc certificate by Subject
	_, err = utils.QueryNocCertificatesBySubject(setup, testconstants.NocCert1Subject)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// query noc root certificate by VID
	_, err = utils.QueryNocRootCertificatesByVid(setup, testconstants.Vid)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// query noc certificate by VID and SKID
	_, err = utils.QueryNocCertificatesByVidAndSkid(setup, testconstants.Vid, testconstants.NocRootCert1SubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// query noc root certificate by Subject Key ID
	aprCertsBySubjectKeyID, _ := utils.QueryNocCertificatesBySubjectKeyID(setup, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(t, 0, len(aprCertsBySubjectKeyID))

	// Child certificate should be revoked as well
	revokedChildCerts, err := utils.QueryNocRevokedIcaCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 1, len(revokedChildCerts.Certs))
	require.Equal(t, testconstants.NocCert1SubjectKeyID, revokedChildCerts.SubjectKeyId)

	// query child noc certificate by Subject Key ID
	aprCertsBySubjectKeyID, _ = utils.QueryNocCertificatesBySubjectKeyID(setup, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, 0, len(aprCertsBySubjectKeyID))

	_, err = utils.QueryNocCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that child noc cert also removed
	_, err = utils.QueryNocIcaCertificatesByVid(setup, testconstants.Vid)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificate key is removed
	require.False(t,
		setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocRootCert1, testconstants.NocRootCert1SerialNumber))
	require.False(t,
		setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocRootCert1, testconstants.NocRootCert1CopySerialNumber))

	// check that unique child certificate key is removed
	require.False(t,
		setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocCert1, testconstants.NocCert1SerialNumber))
}

func TestHandler_RevokeNocX509RootCert_RevokeWithSerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	// add the first NOC root certificate
	addNocX509RootCert := types.NewMsgAddNocX509RootCert(setup.Vendor1.String(), testconstants.NocRootCert1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
	require.NoError(t, err)

	// add the second NOC root certificate
	addNocX509RootCert = types.NewMsgAddNocX509RootCert(setup.Vendor1.String(), testconstants.NocRootCert1Copy, testconstants.CertSchemaVersion)
	_, err = setup.Handler(setup.Ctx, addNocX509RootCert)
	require.NoError(t, err)

	// add the first NOC non-root certificate
	addNocX509Cert := types.NewMsgAddNocX509IcaCert(setup.Vendor1.String(), testconstants.NocCert1, testconstants.CertSchemaVersion)
	_, err = setup.Handler(setup.Ctx, addNocX509Cert)
	require.NoError(t, err)

	// Revoke NOC root with subject and subject key id by serial number
	revokeCert := types.NewMsgRevokeNocX509RootCert(
		setup.Vendor1.String(),
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		testconstants.NocRootCert1SerialNumber,
		testconstants.Info,
		false,
	)
	_, err = setup.Handler(setup.Ctx, revokeCert)
	require.NoError(t, err)

	// Check that cert is added to revoked lists
	revokedNocCerts, err := utils.QueryNocRevokedRootCertificates(setup, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 1, len(revokedNocCerts.Certs))
	require.Equal(t, testconstants.NocRootCert1SerialNumber, revokedNocCerts.Certs[0].SerialNumber)

	// query that noc root certificate is not added to x509 revoked root certs
	revokedRootCerts, _ := utils.QueryRevokedRootCertificates(setup)
	require.Equal(t, 0, len(revokedRootCerts.Certs))

	// Check that cert is removed from noc lists
	rootCerts, err := utils.QueryNocCertificates(setup, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 1, len(rootCerts.Certs))
	require.Equal(t, testconstants.NocRootCert1CopySerialNumber, rootCerts.Certs[0].SerialNumber)

	// Check that root with different serial number still exits
	certsBySubject, err := utils.QueryNocCertificatesBySubject(setup, testconstants.NocRootCert1Subject)
	require.NoError(t, err)
	require.Equal(t, 1, len(certsBySubject.SubjectKeyIds))
	require.Equal(t, testconstants.NocRootCert1Subject, certsBySubject.Subject)

	aprCertsBySubjectKeyID, _ := utils.QueryNocCertificatesBySubjectKeyID(setup, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(t, 1, len(aprCertsBySubjectKeyID))
	require.Equal(t, testconstants.NocRootCert1CopySerialNumber, aprCertsBySubjectKeyID[0].Certs[0].SerialNumber)

	// query noc root certificate by VID should return only one root cert
	revNocRoot, err := utils.QueryNocRootCertificatesByVid(setup, testconstants.Vid)
	require.NoError(t, err)
	require.Equal(t, 1, len(revNocRoot.Certs))
	require.Equal(t, testconstants.NocRootCert1CopySerialNumber, revNocRoot.Certs[0].SerialNumber)

	// query noc certificate by VID and SKID
	nocCertificatesByVidAndSkid, err := utils.QueryNocCertificatesByVidAndSkid(setup, testconstants.Vid, testconstants.NocRootCert1SubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, testconstants.NocRootCert1SubjectKeyID, nocCertificatesByVidAndSkid.SubjectKeyId)
	require.Equal(t, 1, len(revNocRoot.Certs))
	require.Equal(t, float32(1), nocCertificatesByVidAndSkid.Tq)
	require.Equal(t, testconstants.NocRootCert1CopySerialNumber, nocCertificatesByVidAndSkid.Certs[0].SerialNumber)

	// Child certificate should not be revoked
	_, err = utils.QueryNocRevokedIcaCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))

	// query child of revoked certificate, they should not be revoked
	childCerts, _ := utils.QueryNocCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, 1, len(childCerts.Certs))
	require.Equal(t, testconstants.NocCert1SubjectKeyID, childCerts.SubjectKeyId)

	// check that child cert is not removed
	nocCerts, err := utils.QueryNocIcaCertificatesByVid(setup, testconstants.Vid)
	require.NoError(t, err)
	require.Equal(t, 1, len(nocCerts.Certs))
	require.Equal(t, testconstants.NocCert1SubjectKeyID, nocCerts.Certs[0].SubjectKeyId)

	// check that unique certificate key is removed
	require.False(t,
		setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocRootCert1, testconstants.NocRootCert1SerialNumber))
}

func TestHandler_RevokeNocX509RootCert_RevokeWithSerialNumberAndChild(t *testing.T) {
	setup := utils.Setup(t)

	// add the first NOC root certificate
	addNocX509RootCert := types.NewMsgAddNocX509RootCert(setup.Vendor1.String(), testconstants.NocRootCert1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
	require.NoError(t, err)

	// add the second NOC root certificate
	addNocX509RootCert = types.NewMsgAddNocX509RootCert(setup.Vendor1.String(), testconstants.NocRootCert1Copy, testconstants.CertSchemaVersion)
	_, err = setup.Handler(setup.Ctx, addNocX509RootCert)
	require.NoError(t, err)

	// add the first NOC non-root certificate
	addNocX509Cert := types.NewMsgAddNocX509IcaCert(setup.Vendor1.String(), testconstants.NocCert1, testconstants.CertSchemaVersion)
	_, err = setup.Handler(setup.Ctx, addNocX509Cert)
	require.NoError(t, err)

	// Revoke NOC root with subject and subject key id by serial number
	revokeCert := types.NewMsgRevokeNocX509RootCert(
		setup.Vendor1.String(),
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		testconstants.NocRootCert1SerialNumber,
		testconstants.Info,
		true,
	)
	_, err = setup.Handler(setup.Ctx, revokeCert)
	require.NoError(t, err)

	// Check that cert is added to revoked lists
	revokedNocCerts, err := utils.QueryNocRevokedRootCertificates(setup, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 1, len(revokedNocCerts.Certs))
	require.Equal(t, testconstants.NocRootCert1SerialNumber, revokedNocCerts.Certs[0].SerialNumber)

	// query that noc root certificate is not added to x509 revoked root certs
	revokedRootCerts, _ := utils.QueryRevokedRootCertificates(setup)
	require.Equal(t, 0, len(revokedRootCerts.Certs))

	// Check that root with different serial number still exits
	rootCerts, err := utils.QueryNocCertificates(setup, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 1, len(rootCerts.Certs))
	require.Equal(t, testconstants.NocRootCert1CopySerialNumber, rootCerts.Certs[0].SerialNumber)

	certsBySubject, err := utils.QueryNocCertificatesBySubject(setup, testconstants.NocRootCert1Subject)
	require.NoError(t, err)
	require.Equal(t, 1, len(certsBySubject.SubjectKeyIds))
	require.Equal(t, testconstants.NocRootCert1Subject, certsBySubject.Subject)

	aprCertsBySubjectKeyID, _ := utils.QueryNocCertificatesBySubjectKeyID(setup, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(t, 1, len(aprCertsBySubjectKeyID))
	require.Equal(t, testconstants.NocRootCert1CopySerialNumber, aprCertsBySubjectKeyID[0].Certs[0].SerialNumber)

	// query noc root certificate by VID should return only one root cert
	revNocRoot, err := utils.QueryNocRootCertificatesByVid(setup, testconstants.Vid)
	require.NoError(t, err)
	require.Equal(t, 1, len(revNocRoot.Certs))
	require.Equal(t, testconstants.NocRootCert1CopySerialNumber, revNocRoot.Certs[0].SerialNumber)

	// Child certificate should be revoked as well
	revokedCerts, err := utils.QueryNocRevokedIcaCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 1, len(revokedCerts.Certs))
	require.Equal(t, testconstants.NocCert1SubjectKeyID, revokedCerts.SubjectKeyId)

	// query child of revoked certificate, they should be removed as well
	_, err = utils.QueryNocCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))

	_, err = utils.QueryNocCertificatesBySubject(setup, testconstants.NocCert1Subject)
	require.Equal(t, codes.NotFound, status.Code(err))

	aprCertsBySubjectKeyID, _ = utils.QueryNocCertificatesBySubjectKeyID(setup, testconstants.NocCert1Subject)
	require.Equal(t, 0, len(aprCertsBySubjectKeyID))

	_, err = utils.QueryNocIcaCertificatesByVid(setup, testconstants.Vid)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificate key is removed
	require.False(t,
		setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocRootCert1, testconstants.NocRootCert1SerialNumber))
	require.False(t,
		setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocCert1, testconstants.NocCert1SerialNumber))
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
