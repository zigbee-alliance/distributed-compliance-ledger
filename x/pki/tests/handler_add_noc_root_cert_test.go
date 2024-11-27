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

func TestHandler_AddNocRootCert(t *testing.T) {
	setup := utils.Setup(t)

	accAddress := setup.CreateVendorAccount(testconstants.Vid)

	// add NOC root certificate
	rootCertificate := utils.CreateTestNocRoot1Cert()
	utils.AddNocRootCertificate(setup, accAddress, testconstants.NocRootCert1)

	// Check indexes
	indexes := []utils.TestIndex{
		{Key: types.AllCertificatesKeyPrefix, Exist: true},
		{Key: types.AllCertificatesBySubjectKeyPrefix, Exist: true},
		{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Exist: true},
		{Key: types.NocCertificatesKeyPrefix, Exist: true},
		{Key: types.NocCertificatesBySubjectKeyPrefix, Exist: true},
		{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix, Exist: true},
		{Key: types.NocCertificatesByVidAndSkidKeyPrefix, Exist: true},
		{Key: types.NocRootCertificatesKeyPrefix, Exist: true},
		{Key: types.NocIcaCertificatesKeyPrefix, Exist: false},
		{Key: types.UniqueCertificateKeyPrefix, Exist: true},
		{Key: types.ProposedCertificateKeyPrefix, Exist: false},
		{Key: types.ApprovedCertificatesKeyPrefix, Exist: false},
		{Key: types.ApprovedCertificatesBySubjectKeyPrefix, Exist: false},
		{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix, Exist: false},
		{Key: types.ApprovedRootCertificatesKeyPrefix, Exist: false},
	}
	utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)
}

// Extra cases

func TestHandler_AddNocX509RootCert_Renew(t *testing.T) {
	setup := utils.Setup(t)

	accAddress := utils.GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// Store the NOC root certificate
	nocRootCertificate := utils.RootCertificate(accAddress)
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
	newNocCertificate := utils.RootCertificate(accAddress)
	newNocCertificate.CertificateType = types.CertificateType_OperationalPKI
	newNocCertificate.Approvals = nil
	newNocCertificate.Rejects = nil

	// add the new NOC root certificate
	addNocX509RootCert := types.NewMsgAddNocX509RootCert(accAddress.String(), newNocCertificate.PemCert, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
	require.NoError(t, err)

	// query noc root certificate by Subject and SKID
	nocCertificates, err := utils.QueryNocCertificates(setup, newNocCertificate.Subject, newNocCertificate.SubjectKeyId)
	require.NoError(t, err)
	require.Equal(t, len(nocCertificates.Certs), 2)
	require.Equal(t, &newNocCertificate, nocCertificates.Certs[1])

	// query noc root certificate by Subject
	nocCertificatesBySubject, err := utils.QueryNocCertificatesBySubject(setup, newNocCertificate.Subject)
	require.NoError(t, err)
	require.Equal(t, 1, len(nocCertificatesBySubject.SubjectKeyIds))
	require.Equal(t, newNocCertificate.SubjectKeyId, nocCertificatesBySubject.SubjectKeyIds[0])

	// query noc root certificate by SKID
	nocCertificatesBySubjectKeyID, err := utils.QueryNocCertificatesBySubjectKeyID(setup, newNocCertificate.SubjectKeyId)
	require.NoError(t, err)
	require.Equal(t, 1, len(nocCertificatesBySubjectKeyID))
	require.Equal(t, 1, len(nocCertificatesBySubjectKeyID[0].Certs))
	require.Equal(t, &newNocCertificate, nocCertificatesBySubjectKeyID[0].Certs[0])

	// query noc root certificate by VID
	nocRootCertificates, err := utils.QueryNocRootCertificatesByVid(setup, testconstants.Vid)
	require.NoError(t, err)
	require.Equal(t, len(nocRootCertificates.Certs), 2)
	require.Equal(t, &newNocCertificate, nocRootCertificates.Certs[1])

	// query noc root certificate by VID and SKID
	renewedNocRootCertificate, err := utils.QueryNocCertificatesByVidAndSkid(setup, testconstants.Vid, newNocCertificate.SubjectKeyId)
	require.NoError(t, err)
	require.Equal(t, &newNocCertificate, renewedNocRootCertificate.Certs[0])
}

// Error cases

func TestHandler_AddNocX509RootCert_SenderNotVendor(t *testing.T) {
	setup := utils.Setup(t)

	addNocX509RootCert := types.NewMsgAddNocX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addNocX509RootCert)

	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_AddNocX509RootCert_InvalidCertificate(t *testing.T) {
	accAddress := utils.GenerateAccAddress()

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
			setup := utils.Setup(t)
			setup.AddAccount(accAddress, []dclauthtypes.AccountRole{tc.accountRole}, tc.accountVid)

			addNocX509RootCert := types.NewMsgAddNocX509RootCert(accAddress.String(), tc.nocRoorCert, testconstants.CertSchemaVersion)
			_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestHandler_AddNocX509RootCert_CertificateExist(t *testing.T) {
	accAddress := utils.GenerateAccAddress()

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
			setup := utils.Setup(t)
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
