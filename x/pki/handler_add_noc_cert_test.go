package pki

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func TestHandler_AddNocX509Cert_SenderNotVendor(t *testing.T) {
	setup := Setup(t)

	addNocX509Cert := types.NewMsgAddNocX509Cert(setup.Trustee1.String(), testconstants.NocCert1)
	_, err := setup.Handler(setup.Ctx, addNocX509Cert)

	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_AddNocX509Cert_AddNew(t *testing.T) {
	setup := Setup(t)

	accAddress := GenerateAccAddress()
	vid := testconstants.Vid
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	addNocRootCertificate(setup, accAddress, testconstants.NocRootCert1, vid)

	// add the new NOC certificate
	nocX509Cert := types.NewMsgAddNocX509Cert(accAddress.String(), testconstants.NocCert1)
	_, err := setup.Handler(setup.Ctx, nocX509Cert)
	require.NoError(t, err)

	// query noc root certificate by Subject and SKID
	approvedCertificate, err := querySingleApprovedCertificate(setup, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, testconstants.NocCert1Subject, approvedCertificate.Subject)
	require.Equal(t, testconstants.NocCert1SubjectKeyID, approvedCertificate.SubjectKeyId)
	require.Equal(t, testconstants.NocCert1SerialNumber, approvedCertificate.SerialNumber)

	// query noc root certificate by SubjectKeyID
	approvedCertificatesBySubjectKeyID, err := queryAllApprovedCertificatesBySubjectKeyID(setup, testconstants.NocCert1SubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, 1, len(approvedCertificatesBySubjectKeyID))
	require.Equal(t, 1, len(approvedCertificatesBySubjectKeyID[0].Certs))
	require.Equal(t, testconstants.NocCert1SerialNumber, approvedCertificatesBySubjectKeyID[0].Certs[0].SerialNumber)

	// query noc root certificate by VID
	nocRootCertificate, err := querySingleNocCertificate(setup, vid)
	require.NoError(t, err)
	require.Equal(t, testconstants.NocCert1SerialNumber, nocRootCertificate.SerialNumber)

	// check that child certificates of issuer contains certificate identifier
	issuerChildren, _ := queryChildCertificates(
		setup, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(t, 1, len(issuerChildren.CertIds))
	require.Equal(t,
		&types.CertificateIdentifier{
			Subject:      testconstants.NocCert1Subject,
			SubjectKeyId: testconstants.NocCert1SubjectKeyID,
		},
		issuerChildren.CertIds[0])

	// check that unique certificate key registered
	require.True(t,
		setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.NocCert1Issuer, testconstants.NocCert1SerialNumber))
}

func TestHandler_AddNocX509Cert_Renew(t *testing.T) {
	setup := Setup(t)

	accAddress := GenerateAccAddress()
	vid := testconstants.Vid
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add NOC root certificate
	addNocRootCertificate(setup, accAddress, testconstants.NocRootCert1, vid)

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
	)
	newNocCertificate.SerialNumber = testconstants.TestSerialNumber

	setup.Keeper.AddApprovedCertificate(setup.Ctx, newNocCertificate)
	setup.Keeper.AddApprovedCertificateBySubjectKeyID(setup.Ctx, newNocCertificate)
	setup.Keeper.AddApprovedCertificateBySubject(setup.Ctx, newNocCertificate.Subject, newNocCertificate.SubjectKeyId)
	setup.Keeper.AddNocCertificate(setup.Ctx, newNocCertificate)
	uniqueCertificate := types.UniqueCertificate{
		Issuer:       newNocCertificate.Issuer,
		SerialNumber: newNocCertificate.SerialNumber,
		Present:      true,
	}
	setup.Keeper.SetUniqueCertificate(setup.Ctx, uniqueCertificate)

	// add the new NOC certificate
	addNocX509Cert := types.NewMsgAddNocX509Cert(accAddress.String(), testconstants.NocCert1)
	_, err := setup.Handler(setup.Ctx, addNocX509Cert)
	require.NoError(t, err)

	// query noc certificate by Subject and SKID
	approvedCertificates, err := queryApprovedCertificates(setup, newNocCertificate.Subject, newNocCertificate.SubjectKeyId)
	require.NoError(t, err)
	require.Equal(t, len(approvedCertificates.Certs), 2)
	require.Equal(t, &newNocCertificate, approvedCertificates.Certs[0])

	// query noc certificate by Subject
	approvedCertificatesBySubject, err := queryApprovedCertificatesBySubject(setup, newNocCertificate.Subject)
	require.NoError(t, err)
	require.Equal(t, 1, len(approvedCertificatesBySubject.SubjectKeyIds))

	// query noc certificate by SKID
	approvedCertificatesBySubjectKeyID, err := queryAllApprovedCertificatesBySubjectKeyID(setup, newNocCertificate.SubjectKeyId)
	require.NoError(t, err)
	require.Equal(t, 1, len(approvedCertificatesBySubjectKeyID))
	require.Equal(t, 2, len(approvedCertificatesBySubjectKeyID[0].Certs))
	require.Equal(t, testconstants.NocCert1Subject, approvedCertificatesBySubjectKeyID[0].Certs[0].Subject)
	require.Equal(t, testconstants.NocCert1SubjectKeyID, approvedCertificatesBySubjectKeyID[0].Certs[0].SubjectKeyId)
	require.Equal(t, vid, approvedCertificatesBySubjectKeyID[0].Certs[0].Vid)

	// query noc certificate by VID
	nocCertificates, err := queryNocCertificates(setup, testconstants.Vid)
	require.NoError(t, err)
	require.Equal(t, len(nocCertificates.Certs), 2)
	require.Equal(t, testconstants.NocCert1Subject, nocCertificates.Certs[0].Subject)
	require.Equal(t, testconstants.NocCert1SubjectKeyID, nocCertificates.Certs[0].SubjectKeyId)
	require.Equal(t, vid, nocCertificates.Certs[0].Vid)
}

func TestHandler_AddNocX509Cert_Root_VID_Does_Not_Equal_To_AccountVID(t *testing.T) {
	setup := Setup(t)

	accAddress := GenerateAccAddress()
	vid := testconstants.Vid
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add NOC root certificate
	addNocRootCertificate(setup, accAddress, testconstants.NocRootCert1, vid)

	newAccAddress := GenerateAccAddress()
	setup.AddAccount(newAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 1111)

	// try to add NOC certificate
	nocX509Cert := types.NewMsgAddNocX509Cert(newAccAddress.String(), testconstants.NocCert1)
	_, err := setup.Handler(setup.Ctx, nocX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertVidNotEqualAccountVid.Is(err))
}

func TestHandler_AddNocX509Cert_ForInvalidCertificate(t *testing.T) {
	setup := Setup(t)

	accAddress := GenerateAccAddress()
	vid := testconstants.Vid
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add x509 certificate
	addX509Cert := types.NewMsgAddNocX509Cert(accAddress.String(), testconstants.StubCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrInvalidCertificate.Is(err))
}

func TestHandler_AddXNoc509Cert_ForNocRootCertificate(t *testing.T) {
	setup := Setup(t)

	accAddress := GenerateAccAddress()
	vid := testconstants.Vid
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// try to add root certificate x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(accAddress.String(), testconstants.NocRootCert1)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrInappropriateCertificateType.Is(err))
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
	addX509Cert := types.NewMsgAddNocX509Cert(accAddress.String(), testconstants.IntermediateCertWithVid1)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.Error(t, err)
	require.ErrorIs(t, err, pkitypes.ErrInappropriateCertificateType)
}

func TestHandler_AddXNoc509Cert_WhenNocRootCertIsAbsent(t *testing.T) {
	setup := Setup(t)

	accAddress := GenerateAccAddress()
	vid := testconstants.Vid
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	// add the new NOC certificate
	addNocX509Cert := types.NewMsgAddNocX509Cert(accAddress.String(), testconstants.NocCert1)
	_, err := setup.Handler(setup.Ctx, addNocX509Cert)

	require.ErrorIs(t, err, pkitypes.ErrInvalidCertificate)
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
				Issuer:         testconstants.NocRootCert1Subject,
				AuthorityKeyId: testconstants.NocRootCert1SubjectKeyID,
				Subject:        testconstants.NocCert1Subject,
				SubjectAsText:  testconstants.NocCert1SubjectAsText,
				SubjectKeyId:   testconstants.NocCert1SubjectKeyID,
				SerialNumber:   testconstants.NocCert1SerialNumber,
				IsRoot:         false,
				IsNoc:          true,
				Vid:            testconstants.Vid,
			},
			nocCert: testconstants.NocCert1,
			err:     pkitypes.ErrCertificateAlreadyExists,
		},
		{
			name: "ExistingIsRootCert",
			existingCert: &types.Certificate{
				Issuer:         testconstants.NocRootCert1Subject,
				AuthorityKeyId: testconstants.NocRootCert1SubjectKeyID,
				Subject:        testconstants.NocCert1Subject,
				SubjectAsText:  testconstants.NocCert1SubjectAsText,
				SubjectKeyId:   testconstants.NocCert1SubjectKeyID,
				SerialNumber:   testconstants.NocRootCert1SerialNumber,
				IsRoot:         true,
				IsNoc:          true,
				Vid:            testconstants.Vid,
			},
			nocCert: testconstants.NocCert1,
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			name: "ExistingWithDifferentIssuer",
			existingCert: &types.Certificate{
				Issuer:         testconstants.RootIssuer,
				AuthorityKeyId: testconstants.NocRootCert1SubjectKeyID,
				Subject:        testconstants.NocCert1Subject,
				SubjectAsText:  testconstants.NocCert1SubjectAsText,
				SubjectKeyId:   testconstants.NocCert1SubjectKeyID,
				SerialNumber:   "1234",
				IsRoot:         false,
				IsNoc:          true,
				Vid:            testconstants.Vid,
			},
			nocCert: testconstants.NocCert1,
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			name: "ExistingWithDifferentAuthorityKeyId",
			existingCert: &types.Certificate{
				Issuer:         testconstants.NocRootCert1Subject,
				AuthorityKeyId: testconstants.RootSubjectKeyID,
				Subject:        testconstants.NocCert1Subject,
				SubjectAsText:  testconstants.NocCert1SubjectAsText,
				SubjectKeyId:   testconstants.NocCert1SubjectKeyID,
				SerialNumber:   "1234",
				IsRoot:         false,
				IsNoc:          true,
				Vid:            testconstants.Vid,
			},
			nocCert: testconstants.NocCert1,
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			name: "ExistingNotNocCert",
			existingCert: &types.Certificate{
				Issuer:         testconstants.NocRootCert1Subject,
				AuthorityKeyId: testconstants.NocRootCert1SubjectKeyID,
				Subject:        testconstants.NocCert1Subject,
				SubjectAsText:  testconstants.NocCert1SubjectAsText,
				SubjectKeyId:   testconstants.NocCert1SubjectKeyID,
				SerialNumber:   "1234",
				IsRoot:         false,
				IsNoc:          false,
				Vid:            testconstants.Vid,
			},
			nocCert: testconstants.NocCert1,
			err:     pkitypes.ErrInappropriateCertificateType,
		},
		{
			name: "ExistingCertWithDifferentVid",
			existingCert: &types.Certificate{
				Issuer:         testconstants.NocRootCert1Subject,
				AuthorityKeyId: testconstants.NocRootCert1SubjectKeyID,
				Subject:        testconstants.NocCert1Subject,
				SubjectAsText:  testconstants.NocCert1SubjectAsText,
				SubjectKeyId:   testconstants.NocCert1SubjectKeyID,
				SerialNumber:   "1234",
				IsRoot:         false,
				IsNoc:          true,
				Vid:            testconstants.VendorID1,
			},
			nocCert: testconstants.NocCert1,
			err:     pkitypes.ErrCertVidNotEqualAccountVid,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := Setup(t)
			vid := testconstants.Vid
			setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

			// add NOC root certificate
			addNocRootCertificate(setup, accAddress, testconstants.NocRootCert1, vid)

			// add the existing certificate
			setup.Keeper.AddApprovedCertificate(setup.Ctx, *tc.existingCert)
			uniqueCertificate := types.UniqueCertificate{
				Issuer:       tc.existingCert.Issuer,
				SerialNumber: tc.existingCert.SerialNumber,
				Present:      true,
			}
			setup.Keeper.SetUniqueCertificate(setup.Ctx, uniqueCertificate)

			addNocX509Cert := types.NewMsgAddNocX509Cert(accAddress.String(), tc.nocCert)
			_, err := setup.Handler(setup.Ctx, addNocX509Cert)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func addNocRootCertificate(setup *TestSetup, address sdk.AccAddress, pemCert string, vid int32) { //nolint:unparam
	// add the new NOC root certificate
	addNocX509RootCert := types.NewMsgAddNocX509RootCert(address.String(), pemCert)
	_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
	require.NoError(setup.T, err)

	// check that noc certificate has been added
	nocCerts, err := queryNocRootCertificates(setup, vid)
	require.NoError(setup.T, err)
	require.NotNil(setup.T, nocCerts)
}

func querySingleNocCertificate(
	setup *TestSetup,
	vid int32,
) (*types.Certificate, error) {
	certificates, err := queryNocCertificates(setup, vid)
	if err != nil {
		return nil, err
	}

	if len(certificates.Certs) > 1 {
		require.Fail(setup.T, "More than 1 certificate returned")
	}

	return certificates.Certs[0], nil
}

func queryNocCertificates(
	setup *TestSetup,
	vid int32,
) (*types.NocCertificates, error) {
	// query certificate
	req := &types.QueryGetNocCertificatesRequest{Vid: vid}

	resp, err := setup.Keeper.NocCertificates(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.NocCertificates, nil
}
