package pki

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func TestHandler_AddNocX509RootCert_SenderNotVendor(t *testing.T) {
	setup := Setup(t)

	addNocX509RootCert := types.NewMsgAddNocX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addNocX509RootCert)

	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_AddNocX509RootCert_InvalidCertificate(t *testing.T) {
	accAddress := GenerateAccAddress()

	cases := []struct {
		name        string
		accountVid  int32
		accountRole dclauthtypes.AccountRole
		nocRoorCert string
		err         error
	}{
		{
			name:        "NotValidPemCertificate",
			accountVid:  testconstants.Vid,
			accountRole: dclauthtypes.Vendor,
			nocRoorCert: testconstants.StubCertPem,
			err:         pkitypes.ErrInvalidCertificate,
		},
		{
			name:        "NonRootCertificate",
			accountVid:  testconstants.Vid,
			accountRole: dclauthtypes.Vendor,
			nocRoorCert: testconstants.IntermediateCertPem,
			err:         pkitypes.ErrRootCertificateIsNotSelfSigned,
		},
		{
			name:        "ExpiredCertificate",
			accountVid:  testconstants.Vid,
			accountRole: dclauthtypes.Vendor,
			nocRoorCert: testconstants.PAACertExpired,
			err:         pkitypes.ErrInvalidCertificate,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := Setup(t)
			setup.AddAccount(accAddress, []dclauthtypes.AccountRole{tc.accountRole}, tc.accountVid)

			addNocX509RootCert := types.NewMsgAddNocX509RootCert(accAddress.String(), tc.nocRoorCert, testconstants.CertSchemaVersion)
			_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestHandler_AddNocX509RootCert_CertificateExist(t *testing.T) {
	accAddress := GenerateAccAddress()

	cases := []struct {
		name         string
		existingCert *types.Certificate
		nocRoorCert  string
		err          error
	}{
		{
			name: "Duplicate",
			existingCert: &types.Certificate{
				Issuer:          testconstants.RootIssuer,
				Subject:         testconstants.RootSubject,
				SubjectAsText:   testconstants.RootSubjectAsText,
				SubjectKeyId:    testconstants.RootSubjectKeyID,
				SerialNumber:    testconstants.RootSerialNumber,
				IsRoot:          true,
				CertificateType: types.CertificateType_OperationalPKI,
				Vid:             testconstants.Vid,
			},
			nocRoorCert: testconstants.RootCertPem,
			err:         pkitypes.ErrCertificateAlreadyExists,
		},
		{
			name: "ExistingNonRootCert",
			existingCert: &types.Certificate{
				Issuer:          testconstants.TestIssuer,
				Subject:         testconstants.RootSubject,
				SubjectAsText:   testconstants.RootSubjectAsText,
				SubjectKeyId:    testconstants.RootSubjectKeyID,
				SerialNumber:    testconstants.TestSerialNumber,
				IsRoot:          false,
				CertificateType: types.CertificateType_OperationalPKI,
				Vid:             testconstants.Vid,
			},
			nocRoorCert: testconstants.RootCertPem,
			err:         sdkerrors.ErrUnauthorized,
		},
		{
			name: "ExistingNotNocCert",
			existingCert: &types.Certificate{
				Issuer:          testconstants.RootIssuer,
				Subject:         testconstants.RootSubject,
				SubjectAsText:   testconstants.RootSubjectAsText,
				SubjectKeyId:    testconstants.RootSubjectKeyID,
				SerialNumber:    testconstants.TestSerialNumber,
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
				Issuer:          testconstants.RootIssuer,
				Subject:         testconstants.RootSubject,
				SubjectAsText:   testconstants.RootSubjectAsText,
				SubjectKeyId:    testconstants.RootSubjectKeyID,
				SerialNumber:    testconstants.GoogleSerialNumber,
				IsRoot:          true,
				CertificateType: types.CertificateType_OperationalPKI,
				Vid:             testconstants.VendorID1,
			},
			nocRoorCert: testconstants.RootCertPem,
			err:         sdkerrors.ErrUnauthorized,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := Setup(t)
			setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

			// add the existing certificate
			setup.Keeper.AddAllCertificate(setup.Ctx, *tc.existingCert)
			uniqueCertificate := types.UniqueCertificate{
				Issuer:       tc.existingCert.Issuer,
				SerialNumber: tc.existingCert.SerialNumber,
				Present:      true,
			}
			setup.Keeper.SetUniqueCertificate(setup.Ctx, uniqueCertificate)

			addNocX509RootCert := types.NewMsgAddNocX509RootCert(accAddress.String(), tc.nocRoorCert, testconstants.CertSchemaVersion)
			_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestHandler_AddNocX509RootCert_AddNew(t *testing.T) {
	setup := Setup(t)

	accAddress := GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// new NOC root certificate
	newNocCertificate := rootCertificate(accAddress)
	newNocCertificate.CertificateType = types.CertificateType_OperationalPKI
	newNocCertificate.Approvals = nil
	newNocCertificate.Rejects = nil

	// add the new NOC root certificate
	addNocX509RootCert := types.NewMsgAddNocX509RootCert(accAddress.String(), newNocCertificate.PemCert, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
	require.NoError(t, err)

	// query noc root certificate by Subject and SKID
	nocCertificate, err := querySingleNocCertificate(setup, newNocCertificate.Subject, newNocCertificate.SubjectKeyId)
	require.NoError(t, err)
	require.Equal(t, &newNocCertificate, nocCertificate)

	// query noc root certificate by Subject
	nocCertificatesBySubject, err := queryNocCertificatesBySubject(setup, newNocCertificate.Subject)
	require.NoError(t, err)
	require.Equal(t, 1, len(nocCertificatesBySubject.SubjectKeyIds))
	require.Equal(t, newNocCertificate.SubjectKeyId, nocCertificatesBySubject.SubjectKeyIds[0])

	nocCertificatesBySubjectKeyID, err := queryAllNocCertificatesBySubjectKeyID(setup, newNocCertificate.SubjectKeyId)
	require.NoError(t, err)
	require.Equal(t, 1, len(nocCertificatesBySubjectKeyID))
	require.Equal(t, 1, len(nocCertificatesBySubjectKeyID[0].Certs))
	require.Equal(t, &newNocCertificate, nocCertificatesBySubjectKeyID[0].Certs[0])

	// query noc root certificate by VID
	nocRootCertificate, err := querySingleNocRootCertificate(setup, testconstants.Vid)
	require.NoError(t, err)
	require.Equal(t, &newNocCertificate, nocRootCertificate)

	// query noc root certificate by VID and SKID
	nocRootCertificate, tq, err := querySingleNocCertificateByVidAndSkid(setup, newNocCertificate.Vid, newNocCertificate.SubjectKeyId)
	require.NoError(t, err)
	require.Equal(t, &newNocCertificate, nocRootCertificate)
	require.Equal(t, float32(1), tq)

	// check that unique certificate key registered
	require.True(t,
		setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.RootIssuer, testconstants.RootSerialNumber))
}

func TestHandler_AddNocX509RootCert_Renew(t *testing.T) {
	setup := Setup(t)

	accAddress := GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// Store the NOC root certificate
	nocRootCertificate := rootCertificate(accAddress)
	nocRootCertificate.SerialNumber = testconstants.TestSerialNumber
	nocRootCertificate.CertificateType = types.CertificateType_OperationalPKI
	nocRootCertificate.Approvals = nil
	nocRootCertificate.Rejects = nil

	setup.Keeper.AddAllCertificate(setup.Ctx, nocRootCertificate)
	setup.Keeper.AddNocCertificate(setup.Ctx, nocRootCertificate)
	setup.Keeper.AddNocRootCertificate(setup.Ctx, nocRootCertificate)
	setup.Keeper.AddNocCertificateBySubject(setup.Ctx, nocRootCertificate)

	uniqueCertificate := types.UniqueCertificate{
		Issuer:       nocRootCertificate.Issuer,
		SerialNumber: nocRootCertificate.SerialNumber,
		Present:      true,
	}
	setup.Keeper.SetUniqueCertificate(setup.Ctx, uniqueCertificate)

	// new NOC root certificate
	newNocCertificate := rootCertificate(accAddress)
	newNocCertificate.CertificateType = types.CertificateType_OperationalPKI
	newNocCertificate.Approvals = nil
	newNocCertificate.Rejects = nil

	// add the new NOC root certificate
	addNocX509RootCert := types.NewMsgAddNocX509RootCert(accAddress.String(), newNocCertificate.PemCert, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
	require.NoError(t, err)

	// query noc root certificate by Subject and SKID
	nocCertificates, err := queryNocCertificates(setup, newNocCertificate.Subject, newNocCertificate.SubjectKeyId)
	require.NoError(t, err)
	require.Equal(t, len(nocCertificates.Certs), 2)
	require.Equal(t, &newNocCertificate, nocCertificates.Certs[1])

	// query noc root certificate by Subject
	nocCertificatesBySubject, err := queryNocCertificatesBySubject(setup, newNocCertificate.Subject)
	require.NoError(t, err)
	require.Equal(t, 1, len(nocCertificatesBySubject.SubjectKeyIds))
	require.Equal(t, newNocCertificate.SubjectKeyId, nocCertificatesBySubject.SubjectKeyIds[0])

	// query noc root certificate by SKID
	nocCertificatesBySubjectKeyID, err := queryAllNocCertificatesBySubjectKeyID(setup, newNocCertificate.SubjectKeyId)
	require.NoError(t, err)
	require.Equal(t, 1, len(nocCertificatesBySubjectKeyID))
	require.Equal(t, 1, len(nocCertificatesBySubjectKeyID[0].Certs))
	require.Equal(t, &newNocCertificate, nocCertificatesBySubjectKeyID[0].Certs[0])

	// query noc root certificate by VID
	nocRootCertificates, err := queryNocRootCertificates(setup, testconstants.Vid)
	require.NoError(t, err)
	require.Equal(t, len(nocRootCertificates.Certs), 2)
	require.Equal(t, &newNocCertificate, nocRootCertificates.Certs[1])

	// query noc root certificate by VID and SKID
	renewedNocRootCertificate, tq, err := querySingleNocCertificateByVidAndSkid(setup, testconstants.Vid, newNocCertificate.SubjectKeyId)
	require.NoError(t, err)
	require.Equal(t, &newNocCertificate, renewedNocRootCertificate)
	require.Equal(t, float32(1), tq)
}
