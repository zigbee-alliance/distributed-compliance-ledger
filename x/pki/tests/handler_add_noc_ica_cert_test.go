package tests

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// Main

func TestHandler_AddNocX509Cert_AddNewIca(t *testing.T) {
	setup := Setup(t)

	accAddress := GenerateAccAddress()
	vid := testconstants.Vid
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	addNocRootCertificate(setup, accAddress, testconstants.NocRootCert1)

	// add NOC ICA certificate
	addNocIcaCertificate(setup, accAddress, testconstants.NocCert1)

	ensureNocIcaCertificateExist(
		t,
		setup,
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1Issuer,
		testconstants.NocCert1SerialNumber,
		vid,
		false)

	// ChildCertificates: check that child certificates of issuer contains certificate identifier
	issuerChildren, _ := queryChildCertificates(
		setup, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(t, 1, len(issuerChildren.CertIds))
	require.Equal(t,
		&types.CertificateIdentifier{
			Subject:      testconstants.NocCert1Subject,
			SubjectKeyId: testconstants.NocCert1SubjectKeyID,
		},
		issuerChildren.CertIds[0])
}

// Extra cases

func TestHandler_AddNocX509Cert_Renew(t *testing.T) {
	setup := Setup(t)

	accAddress := GenerateAccAddress()
	vid := testconstants.Vid
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add NOC root certificate
	addNocRootCertificate(setup, accAddress, testconstants.NocRootCert1)

	// Store the NOC certificate
	newNocCertificate := types.NewNocCertificate(
		testconstants.NocCert1,
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectAsText,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1SerialNumber,
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		accAddress.String(),
		vid,
		testconstants.SchemaVersion,
	)
	newNocCertificate.SerialNumber = testconstants.TestSerialNumber

	setup.Keeper.AddAllCertificate(setup.Ctx, newNocCertificate)
	setup.Keeper.AddNocCertificate(setup.Ctx, newNocCertificate)
	setup.Keeper.AddNocCertificateBySubjectKeyID(setup.Ctx, newNocCertificate)
	setup.Keeper.AddNocCertificateBySubject(setup.Ctx, newNocCertificate)
	setup.Keeper.AddNocIcaCertificate(setup.Ctx, newNocCertificate)
	uniqueCertificate := types.UniqueCertificate{
		Issuer:       newNocCertificate.Issuer,
		SerialNumber: newNocCertificate.SerialNumber,
		Present:      true,
	}
	setup.Keeper.SetUniqueCertificate(setup.Ctx, uniqueCertificate)

	// add the new NOC certificate
	addNocX509Cert := types.NewMsgAddNocX509IcaCert(accAddress.String(), testconstants.NocCert1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addNocX509Cert)
	require.NoError(t, err)

	// query noc certificate by Subject and SKID
	nocCertificates, err := queryNocCertificates(setup, newNocCertificate.Subject, newNocCertificate.SubjectKeyId)
	require.NoError(t, err)
	require.Equal(t, len(nocCertificates.Certs), 2)
	require.Equal(t, &newNocCertificate, nocCertificates.Certs[0])

	// query noc certificate by Subject
	nocCertificatesBySubject, err := queryNocCertificatesBySubject(setup, newNocCertificate.Subject)
	require.NoError(t, err)
	require.Equal(t, 1, len(nocCertificatesBySubject.SubjectKeyIds))

	// query noc certificate by SKID
	nocCertificatesBySubjectKeyID, err := queryAllNocCertificatesBySubjectKeyID(setup, newNocCertificate.SubjectKeyId)
	require.NoError(t, err)
	require.Equal(t, 1, len(nocCertificatesBySubjectKeyID))
	require.Equal(t, 2, len(nocCertificatesBySubjectKeyID[0].Certs))
	require.Equal(t, testconstants.NocCert1Subject, nocCertificatesBySubjectKeyID[0].Certs[0].Subject)
	require.Equal(t, testconstants.NocCert1SubjectKeyID, nocCertificatesBySubjectKeyID[0].Certs[0].SubjectKeyId)
	require.Equal(t, vid, nocCertificatesBySubjectKeyID[0].Certs[0].Vid)

	// query noc certificate by VID
	nocCertificatesByVid, err := queryNocIcaCertificatesByVid(setup, testconstants.Vid)
	require.NoError(t, err)
	require.Equal(t, len(nocCertificatesByVid.Certs), 2)
	require.Equal(t, testconstants.NocCert1Subject, nocCertificatesByVid.Certs[0].Subject)
	require.Equal(t, testconstants.NocCert1SubjectKeyID, nocCertificatesByVid.Certs[0].SubjectKeyId)
	require.Equal(t, vid, nocCertificatesByVid.Certs[0].Vid)
}

// Error cases

func TestHandler_AddNocX509Cert_SenderNotVendor(t *testing.T) {
	setup := Setup(t)

	addNocX509Cert := types.NewMsgAddNocX509IcaCert(setup.Trustee1.String(), testconstants.NocCert1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addNocX509Cert)

	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_AddNocX509Cert_Root_VID_Does_Not_Equal_To_AccountVID(t *testing.T) {
	setup := Setup(t)

	accAddress := GenerateAccAddress()
	vid := testconstants.Vid
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	addNocRootCertificate(setup, accAddress, testconstants.NocRootCert1)

	newAccAddress := GenerateAccAddress()
	setup.AddAccount(newAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 1111)

	// try to add NOC certificate
	nocX509Cert := types.NewMsgAddNocX509IcaCert(newAccAddress.String(), testconstants.NocCert1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, nocX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrCertVidNotEqualAccountVid)
}

func TestHandler_AddNocX509Cert_ForInvalidCertificate(t *testing.T) {
	setup := Setup(t)

	accAddress := GenerateAccAddress()
	vid := testconstants.Vid
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add x509 certificate
	addX509Cert := types.NewMsgAddNocX509IcaCert(accAddress.String(), testconstants.StubCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrInvalidCertificate)
}

func TestHandler_AddXNoc509Cert_ForNocRootCertificate(t *testing.T) {
	setup := Setup(t)

	accAddress := GenerateAccAddress()
	vid := testconstants.Vid
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// try to add root certificate x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(accAddress.String(), testconstants.NocRootCert1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrNonRootCertificateSelfSigned)
}

func TestHandler_AddXNoc509Cert_ForRootNonNocCertificate(t *testing.T) {
	setup := Setup(t)

	accAddress := GenerateAccAddress()
	vid := testconstants.Vid
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// store root certificate
	rootCertOptions := &rootCertOptions{
		pemCert:      testconstants.RootCertWithVid,
		info:         testconstants.Info,
		subject:      testconstants.RootCertWithVidSubject,
		subjectKeyID: testconstants.RootCertWithVidSubjectKeyID,
		vid:          testconstants.RootCertWithVidVid,
	}
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// try to add root certificate x509 certificate
	addX509Cert := types.NewMsgAddNocX509IcaCert(accAddress.String(), testconstants.IntermediateCertWithVid1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrInappropriateCertificateType)
}

func TestHandler_AddXNoc509Cert_WhenNocRootCertIsAbsent(t *testing.T) {
	setup := Setup(t)

	accAddress := GenerateAccAddress()
	vid := testconstants.Vid
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add the new NOC certificate
	addNocX509Cert := types.NewMsgAddNocX509IcaCert(accAddress.String(), testconstants.NocCert1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addNocX509Cert)

	require.ErrorIs(t, err, pkitypes.ErrCertificateDoesNotExist)
}

func TestHandler_AddNocX509Cert_CertificateExist(t *testing.T) {
	accAddress := GenerateAccAddress()

	cases := []struct {
		name         string
		existingCert *types.Certificate
		nocCert      string
		err          error
	}{
		{
			name: "Duplicate",
			existingCert: &types.Certificate{
				Issuer:          testconstants.NocRootCert1Subject,
				AuthorityKeyId:  testconstants.NocRootCert1SubjectKeyID,
				Subject:         testconstants.NocCert1Subject,
				SubjectAsText:   testconstants.NocCert1SubjectAsText,
				SubjectKeyId:    testconstants.NocCert1SubjectKeyID,
				SerialNumber:    testconstants.NocCert1SerialNumber,
				IsRoot:          false,
				CertificateType: types.CertificateType_OperationalPKI,
				Vid:             testconstants.Vid,
			},
			nocCert: testconstants.NocCert1,
			err:     pkitypes.ErrCertificateAlreadyExists,
		},
		{
			name: "ExistingIsRootCert",
			existingCert: &types.Certificate{
				Issuer:          testconstants.NocRootCert1Subject,
				AuthorityKeyId:  testconstants.NocRootCert1SubjectKeyID,
				Subject:         testconstants.NocCert1Subject,
				SubjectAsText:   testconstants.NocCert1SubjectAsText,
				SubjectKeyId:    testconstants.NocCert1SubjectKeyID,
				SerialNumber:    testconstants.NocRootCert1SerialNumber,
				IsRoot:          true,
				CertificateType: types.CertificateType_OperationalPKI,
				Vid:             testconstants.Vid,
			},
			nocCert: testconstants.NocCert1,
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			name: "ExistingWithDifferentIssuer",
			existingCert: &types.Certificate{
				Issuer:          testconstants.RootIssuer,
				AuthorityKeyId:  testconstants.NocRootCert1SubjectKeyID,
				Subject:         testconstants.NocCert1Subject,
				SubjectAsText:   testconstants.NocCert1SubjectAsText,
				SubjectKeyId:    testconstants.NocCert1SubjectKeyID,
				SerialNumber:    "1234",
				IsRoot:          false,
				CertificateType: types.CertificateType_OperationalPKI,
				Vid:             testconstants.Vid,
			},
			nocCert: testconstants.NocCert1,
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			name: "ExistingWithDifferentAuthorityKeyId",
			existingCert: &types.Certificate{
				Issuer:          testconstants.NocRootCert1Subject,
				AuthorityKeyId:  testconstants.RootSubjectKeyID,
				Subject:         testconstants.NocCert1Subject,
				SubjectAsText:   testconstants.NocCert1SubjectAsText,
				SubjectKeyId:    testconstants.NocCert1SubjectKeyID,
				SerialNumber:    "1234",
				IsRoot:          false,
				CertificateType: types.CertificateType_OperationalPKI,
				Vid:             testconstants.Vid,
			},
			nocCert: testconstants.NocCert1,
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			name: "ExistingNotNocCert",
			existingCert: &types.Certificate{
				Issuer:          testconstants.NocRootCert1Subject,
				AuthorityKeyId:  testconstants.NocRootCert1SubjectKeyID,
				Subject:         testconstants.NocCert1Subject,
				SubjectAsText:   testconstants.NocCert1SubjectAsText,
				SubjectKeyId:    testconstants.NocCert1SubjectKeyID,
				SerialNumber:    "1234",
				IsRoot:          false,
				CertificateType: types.CertificateType_DeviceAttestationPKI,
				Vid:             testconstants.Vid,
			},
			nocCert: testconstants.NocCert1,
			err:     pkitypes.ErrInappropriateCertificateType,
		},
		{
			name: "ExistingCertWithDifferentVid",
			existingCert: &types.Certificate{
				Issuer:          testconstants.NocRootCert1Subject,
				AuthorityKeyId:  testconstants.NocRootCert1SubjectKeyID,
				Subject:         testconstants.NocCert1Subject,
				SubjectAsText:   testconstants.NocCert1SubjectAsText,
				SubjectKeyId:    testconstants.NocCert1SubjectKeyID,
				SerialNumber:    "1234",
				IsRoot:          false,
				CertificateType: types.CertificateType_OperationalPKI,
				Vid:             testconstants.VendorID1,
			},
			nocCert: testconstants.NocCert1,
			err:     sdkerrors.ErrUnauthorized,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := Setup(t)
			vid := testconstants.Vid
			setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

			// add NOC root certificate
			addNocRootCertificate(setup, accAddress, testconstants.NocRootCert1)

			// add the existing certificate
			setup.Keeper.AddAllCertificate(setup.Ctx, *tc.existingCert)
			uniqueCertificate := types.UniqueCertificate{
				Issuer:       tc.existingCert.Issuer,
				SerialNumber: tc.existingCert.SerialNumber,
				Present:      true,
			}
			setup.Keeper.SetUniqueCertificate(setup.Ctx, uniqueCertificate)

			addNocX509Cert := types.NewMsgAddNocX509IcaCert(accAddress.String(), tc.nocCert, testconstants.CertSchemaVersion)
			_, err := setup.Handler(setup.Ctx, addNocX509Cert)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
