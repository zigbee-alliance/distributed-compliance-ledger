package pki

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func TestHandler_RevokeNocX509Cert_SenderNotVendor(t *testing.T) {
	setup := Setup(t)

	accAddress := GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add the new NOC root certificate
	addNocX509RootCert := types.NewMsgAddNocX509RootCert(accAddress.String(), testconstants.NocRootCert1, testconstants.SchemaVersion)
	_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
	require.NoError(t, err)

	revokeCert := types.NewMsgRevokeNocRootX509Cert(
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
	setup := Setup(t)

	accAddress := GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	revokeCert := types.NewMsgRevokeNocX509Cert(
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
	accAddress := GenerateAccAddress()

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
				IsNoc:         true,
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
				IsNoc:         false,
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
				IsNoc:         true,
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
				IsNoc:         true,
				Vid:           testconstants.Vid,
			},
			nocRoorCert: testconstants.NocCert1,
			err:         pkitypes.ErrCertificateDoesNotExist,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := Setup(t)
			setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

			// add the existing certificate
			setup.Keeper.AddApprovedCertificate(setup.Ctx, *tc.existingCert)
			uniqueCertificate := types.UniqueCertificate{
				Issuer:       tc.existingCert.Issuer,
				SerialNumber: tc.existingCert.SerialNumber,
				Present:      true,
			}
			setup.Keeper.SetUniqueCertificate(setup.Ctx, uniqueCertificate)

			revokeCert := types.NewMsgRevokeNocX509Cert(
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

func TestHandler_RevokeNocX509Cert_RevokeDefault(t *testing.T) {
	setup := Setup(t)

	accAddress := GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add the first NOC root certificate
	addNocX509RootCert := types.NewMsgAddNocX509RootCert(accAddress.String(), testconstants.NocRootCert1, testconstants.SchemaVersion)
	_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
	require.NoError(t, err)

	// add the first NOC non-root certificate
	addNocX509Cert := types.NewMsgAddNocX509Cert(accAddress.String(), testconstants.NocCert1, testconstants.SchemaVersion)
	_, err = setup.Handler(setup.Ctx, addNocX509Cert)
	require.NoError(t, err)

	// add the second NOC non-root certificate
	addNocX509Cert = types.NewMsgAddNocX509Cert(accAddress.String(), testconstants.NocCert1Copy, testconstants.SchemaVersion)
	_, err = setup.Handler(setup.Ctx, addNocX509Cert)
	require.NoError(t, err)

	// add the NOC leaf certificate
	addNocX509Cert = types.NewMsgAddNocX509Cert(accAddress.String(), testconstants.NocLeafCert1, testconstants.SchemaVersion)
	_, err = setup.Handler(setup.Ctx, addNocX509Cert)
	require.NoError(t, err)

	// Revoke NOC with subject and subject key id only
	revokeCert := types.NewMsgRevokeNocX509Cert(
		accAddress.String(),
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		"",
		testconstants.Info,
		false,
	)
	_, err = setup.Handler(setup.Ctx, revokeCert)
	require.NoError(t, err)

	revokedNocCerts, err := queryRevokedCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 2, len(revokedNocCerts.Certs))
	require.Equal(t, testconstants.NocCert1Subject, revokedNocCerts.Subject)
	require.Equal(t, testconstants.NocCert1SubjectKeyID, revokedNocCerts.SubjectKeyId)

	// query noc certificate by Subject
	_, err = queryApprovedCertificatesBySubject(setup, testconstants.NocCert1Subject)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// query noc certificate by Subject Key ID
	aprCertsBySubjectKeyID, _ := queryAllApprovedCertificatesBySubjectKeyID(setup, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, 0, len(aprCertsBySubjectKeyID))

	// query noc certificate by VID
	nocCerts, err := queryNocCertificates(setup, testconstants.Vid)
	require.NoError(t, err)
	require.Equal(t, 1, len(nocCerts.Certs))
	require.Equal(t, testconstants.NocLeafCert1SubjectKeyID, nocCerts.Certs[0].SubjectKeyId)

	// Child certificate should not be revoked
	_, err = queryRevokedCertificates(setup, testconstants.NocLeafCert1Subject, testconstants.NocLeafCert1SubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))

	// query all certs
	certs, err := queryAllApprovedCertificates(setup)
	require.NoError(t, err)
	require.Equal(t, 2, len(certs))
	require.NotEqual(t, testconstants.NocCert1SubjectKeyID, certs[0].SubjectKeyId)
	require.NotEqual(t, testconstants.NocCert1SubjectKeyID, certs[1].SubjectKeyId)

	// query child of revoked certificate, they should not be revoked
	childCerts, _ := queryApprovedCertificates(setup, testconstants.NocLeafCert1Subject, testconstants.NocLeafCert1SubjectKeyID)
	require.Equal(t, 1, len(childCerts.Certs))
	require.Equal(t, testconstants.NocLeafCert1SubjectKeyID, childCerts.SubjectKeyId)

	// check that child cert is not removed
	nocCerts, err = queryNocCertificates(setup, testconstants.Vid)
	require.NoError(t, err)
	require.Equal(t, 1, len(nocCerts.Certs))
	require.Equal(t, testconstants.NocLeafCert1SubjectKeyID, nocCerts.Certs[0].SubjectKeyId)

	// check that unique certificate key is removed
	require.False(t,
		setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocCert1, testconstants.NocCert1SerialNumber))
	require.False(t,
		setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocCert1, testconstants.NocCert1CopySerialNumber))
}

func TestHandler_RevokeNocX509Cert_RevokeWithChild(t *testing.T) {
	setup := Setup(t)

	accAddress := GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add the first NOC root certificate
	addNocX509RootCert := types.NewMsgAddNocX509RootCert(accAddress.String(), testconstants.NocRootCert1, testconstants.SchemaVersion)
	_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
	require.NoError(t, err)

	// add the first NOC non-root certificate
	addNocX509Cert := types.NewMsgAddNocX509Cert(accAddress.String(), testconstants.NocCert1, testconstants.SchemaVersion)
	_, err = setup.Handler(setup.Ctx, addNocX509Cert)
	require.NoError(t, err)

	// add the second NOC non-root certificate
	addNocX509Cert = types.NewMsgAddNocX509Cert(accAddress.String(), testconstants.NocCert1Copy, testconstants.SchemaVersion)
	_, err = setup.Handler(setup.Ctx, addNocX509Cert)
	require.NoError(t, err)

	// add the NOC leaf certificate
	addNocX509Cert = types.NewMsgAddNocX509Cert(accAddress.String(), testconstants.NocLeafCert1, testconstants.SchemaVersion)
	_, err = setup.Handler(setup.Ctx, addNocX509Cert)
	require.NoError(t, err)

	// Revoke noc with subject and subject key id and its child too
	revokeCert := types.NewMsgRevokeNocX509Cert(
		accAddress.String(),
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		"",
		testconstants.Info,
		true,
	)
	_, err = setup.Handler(setup.Ctx, revokeCert)
	require.NoError(t, err)

	allRevokedCerts, err := queryAllRevokedCertificates(setup)
	require.NoError(t, err)
	require.Equal(t, 2, len(allRevokedCerts))
	require.Equal(t, 3, len(allRevokedCerts[0].Certs)+len(allRevokedCerts[1].Certs))

	revokedNocCerts, err := queryRevokedCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 2, len(revokedNocCerts.Certs))
	require.Equal(t, testconstants.NocCert1Subject, revokedNocCerts.Subject)
	require.Equal(t, testconstants.NocCert1SubjectKeyID, revokedNocCerts.SubjectKeyId)

	// query all certs
	certs, err := queryAllApprovedCertificates(setup)
	require.NoError(t, err)
	require.Equal(t, 1, len(certs))
	require.Equal(t, testconstants.NocRootCert1SubjectKeyID, certs[0].SubjectKeyId)

	// query NOC cert by subject and subject key id
	_, err = queryApprovedCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1CopySubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// query NOC certificate by Subject
	_, err = queryApprovedCertificatesBySubject(setup, testconstants.NocCert1Subject)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	_, err = queryApprovedCertificatesBySubject(setup, testconstants.NocLeafCert1Subject)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// query NOC certificate by Subject Key ID
	aprCertsBySubjectKeyID, _ := queryAllApprovedCertificatesBySubjectKeyID(setup, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, 0, len(aprCertsBySubjectKeyID))

	aprCertsBySubjectKeyID, _ = queryAllApprovedCertificatesBySubjectKeyID(setup, testconstants.NocLeafCert1SubjectKeyID)
	require.Equal(t, 0, len(aprCertsBySubjectKeyID))

	// query noc certificate by VID
	_, err = queryNocCertificates(setup, testconstants.Vid)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// Child certificate should be revoked as well
	revokedLeafCerts, err := queryRevokedCertificates(setup, testconstants.NocLeafCert1Subject, testconstants.NocLeafCert1SubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 1, len(revokedLeafCerts.Certs))
	require.Equal(t, testconstants.NocLeafCert1SubjectKeyID, revokedLeafCerts.SubjectKeyId)

	// query child of revoked certificate, they should be revoked
	_, err = queryApprovedCertificates(setup, testconstants.NocLeafCert1Subject, testconstants.NocLeafCert1SubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificate key is removed
	require.False(t,
		setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocCert1, testconstants.NocCert1SerialNumber))
	require.False(t,
		setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocCert1, testconstants.NocCert1CopySerialNumber))
	require.False(t,
		setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocLeafCert1, testconstants.NocLeafCert1SerialNumber))
}

func TestHandler_RevokeNocX509Cert_RevokeBySerialNumber(t *testing.T) {
	setup := Setup(t)

	accAddress := GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add the first NOC root certificate
	addNocX509RootCert := types.NewMsgAddNocX509RootCert(accAddress.String(), testconstants.NocRootCert1, testconstants.SchemaVersion)
	_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
	require.NoError(t, err)

	// add the first NOC non-root certificate
	addNocX509Cert := types.NewMsgAddNocX509Cert(accAddress.String(), testconstants.NocCert1, testconstants.SchemaVersion)
	_, err = setup.Handler(setup.Ctx, addNocX509Cert)
	require.NoError(t, err)

	// add the second NOC non-root certificate
	addNocX509Cert = types.NewMsgAddNocX509Cert(accAddress.String(), testconstants.NocCert1Copy, testconstants.SchemaVersion)
	_, err = setup.Handler(setup.Ctx, addNocX509Cert)
	require.NoError(t, err)

	// add the NOC leaf certificate
	addNocX509Cert = types.NewMsgAddNocX509Cert(accAddress.String(), testconstants.NocLeafCert1, testconstants.SchemaVersion)
	_, err = setup.Handler(setup.Ctx, addNocX509Cert)
	require.NoError(t, err)

	// Revoke NOC by serial number only
	revokeCert := types.NewMsgRevokeNocX509Cert(
		accAddress.String(),
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1SerialNumber,
		testconstants.Info,
		false,
	)
	_, err = setup.Handler(setup.Ctx, revokeCert)
	require.NoError(t, err)

	revokedNocCerts, err := queryRevokedCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 1, len(revokedNocCerts.Certs))
	require.Equal(t, testconstants.NocCert1SerialNumber, revokedNocCerts.Certs[0].SerialNumber)

	// Child certificate should not be revoked
	_, err = queryRevokedCertificates(setup, testconstants.NocLeafCert1Subject, testconstants.NocLeafCert1SubjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))

	// query NOC certificate by Subject
	certsBySubject, err := queryApprovedCertificatesBySubject(setup, testconstants.NocCert1Subject)
	require.NoError(t, err)
	require.Equal(t, 1, len(certsBySubject.SubjectKeyIds))

	// query NOC certificate by Subject Key ID
	aprCertsBySubjectKeyID, _ := queryAllApprovedCertificatesBySubjectKeyID(setup, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, 1, len(aprCertsBySubjectKeyID))
	require.Equal(t, 1, len(aprCertsBySubjectKeyID[0].Certs))
	require.Equal(t, testconstants.NocCert1CopySerialNumber, aprCertsBySubjectKeyID[0].Certs[0].SerialNumber)

	// query noc certificate by VID
	nocCerts, err := queryNocCertificates(setup, testconstants.Vid)
	require.NoError(t, err)
	require.Equal(t, 2, len(nocCerts.Certs))
	require.NotEqual(t, testconstants.NocCert1SerialNumber, nocCerts.Certs[0].SerialNumber)
	require.NotEqual(t, testconstants.NocCert1SerialNumber, nocCerts.Certs[1].SerialNumber)

	// query all certs
	certs, err := queryAllApprovedCertificates(setup)
	require.NoError(t, err)
	require.Equal(t, 3, len(certs))
	require.NotEqual(t, testconstants.NocCert1SerialNumber, certs[0].Certs[0].SerialNumber)
	require.NotEqual(t, testconstants.NocCert1SerialNumber, certs[1].Certs[0].SerialNumber)
	require.NotEqual(t, testconstants.NocCert1SerialNumber, certs[2].Certs[0].SerialNumber)

	// query approved certificate, cert with different serial number should not be removed
	approvedCerts, _ := queryApprovedCertificates(setup, testconstants.NocCert1CopySubject, testconstants.NocCert1CopySubjectKeyID)
	require.Equal(t, 1, len(approvedCerts.Certs))
	require.Equal(t, testconstants.NocCert1CopySerialNumber, approvedCerts.Certs[0].SerialNumber)

	// query child certificate, they should not be removed
	childCerts, _ := queryApprovedCertificates(setup, testconstants.NocLeafCert1Subject, testconstants.NocLeafCert1SubjectKeyID)
	require.Equal(t, 1, len(childCerts.Certs))
	require.Equal(t, testconstants.NocLeafCert1SubjectKeyID, childCerts.SubjectKeyId)

	// check that unique certificate key is removed
	require.False(t,
		setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocCert1, testconstants.NocCert1SerialNumber))
}

func TestHandler_RevokeNocX509Cert_RevokeBySerialNumberAndWithChild(t *testing.T) {
	setup := Setup(t)

	accAddress := GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add the first NOC root certificate
	addNocX509RootCert := types.NewMsgAddNocX509RootCert(accAddress.String(), testconstants.NocRootCert1, testconstants.SchemaVersion)
	_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
	require.NoError(t, err)

	// add the first NOC non-root certificate
	addNocX509Cert := types.NewMsgAddNocX509Cert(accAddress.String(), testconstants.NocCert1, testconstants.SchemaVersion)
	_, err = setup.Handler(setup.Ctx, addNocX509Cert)
	require.NoError(t, err)

	// add the second NOC non-root certificate
	addNocX509Cert = types.NewMsgAddNocX509Cert(accAddress.String(), testconstants.NocCert1Copy, testconstants.SchemaVersion)
	_, err = setup.Handler(setup.Ctx, addNocX509Cert)
	require.NoError(t, err)

	// add the NOC leaf certificate
	addNocX509Cert = types.NewMsgAddNocX509Cert(accAddress.String(), testconstants.NocLeafCert1, testconstants.SchemaVersion)
	_, err = setup.Handler(setup.Ctx, addNocX509Cert)
	require.NoError(t, err)

	// Revoke NOC with subject and subject key id and its child too
	revokeCert := types.NewMsgRevokeNocX509Cert(
		accAddress.String(),
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1SerialNumber,
		testconstants.Info,
		true,
	)
	_, err = setup.Handler(setup.Ctx, revokeCert)
	require.NoError(t, err)

	allRevokedCerts, err := queryAllRevokedCertificates(setup)
	require.NoError(t, err)
	require.Equal(t, 2, len(allRevokedCerts))
	require.Equal(t, 2, len(allRevokedCerts[0].Certs)+len(allRevokedCerts[1].Certs))

	revokedNocCerts, err := queryRevokedCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 1, len(revokedNocCerts.Certs))
	require.Equal(t, testconstants.NocCert1SerialNumber, revokedNocCerts.Certs[0].SerialNumber)

	// Child certificate should be revoked
	revokedNocCerts, err = queryRevokedCertificates(setup, testconstants.NocLeafCert1Subject, testconstants.NocLeafCert1SubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 1, len(revokedNocCerts.Certs))
	require.Equal(t, testconstants.NocLeafCert1SerialNumber, revokedNocCerts.Certs[0].SerialNumber)

	// query child of revoked certificate, they should be revoked
	_, err = queryApprovedCertificates(setup, testconstants.NocLeafCert1Subject, testconstants.NocLeafCert1SubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// query all certs
	certs, err := queryAllApprovedCertificates(setup)
	require.NoError(t, err)
	require.Equal(t, 2, len(certs))
	require.NotEqual(t, testconstants.NocCert1SerialNumber, certs[0].Certs[0].SerialNumber)
	require.NotEqual(t, testconstants.NocCert1SerialNumber, certs[1].Certs[0].SerialNumber)

	// query approved certificates
	aprCerts, err := queryApprovedCertificates(setup, testconstants.NocCert1Subject, testconstants.NocCert1CopySubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 1, len(aprCerts.Certs))
	require.Equal(t, testconstants.NocCert1CopySerialNumber, aprCerts.Certs[0].SerialNumber)

	// query noc certificate by Subject
	certsBySubject, err := queryApprovedCertificatesBySubject(setup, testconstants.NocCert1Subject)
	require.NoError(t, err)
	require.Equal(t, 1, len(certsBySubject.SubjectKeyIds))

	_, err = queryApprovedCertificatesBySubject(setup, testconstants.NocLeafCert1Subject)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// query noc certificate by Subject Key ID
	aprCertsBySubjectKeyID, _ := queryAllApprovedCertificatesBySubjectKeyID(setup, testconstants.NocCert1SubjectKeyID)
	require.Equal(t, 1, len(aprCertsBySubjectKeyID))
	require.Equal(t, testconstants.NocCert1CopySerialNumber, aprCertsBySubjectKeyID[0].Certs[0].SerialNumber)

	aprCertsBySubjectKeyID, _ = queryAllApprovedCertificatesBySubjectKeyID(setup, testconstants.NocLeafCert1SubjectKeyID)
	require.Equal(t, 0, len(aprCertsBySubjectKeyID))

	// query noc certificate by VID
	nocCerts, err := queryNocCertificates(setup, testconstants.Vid)
	require.NoError(t, err)
	require.Equal(t, 1, len(nocCerts.Certs))
	require.Equal(t, testconstants.NocCert1CopySerialNumber, nocCerts.Certs[0].SerialNumber)

	// check that unique certificate key is removed
	require.False(t,
		setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocCert1, testconstants.NocCert1SerialNumber))
}
