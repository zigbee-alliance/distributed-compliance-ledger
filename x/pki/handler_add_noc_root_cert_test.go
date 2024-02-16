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

func TestHandler_AddNocX509RootCert_SendorNotVendor(t *testing.T) {
	setup := Setup(t)

	addNocX509RootCert := types.NewMsgAddNocX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem)
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

			addNocX509RootCert := types.NewMsgAddNocX509RootCert(accAddress.String(), tc.nocRoorCert)
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
				Issuer:        testconstants.RootIssuer,
				Subject:       testconstants.RootSubject,
				SubjectAsText: testconstants.RootSubjectAsText,
				SubjectKeyId:  testconstants.RootSubjectKeyID,
				SerialNumber:  testconstants.RootSerialNumber,
				IsRoot:        true,
				IsNoc:         true,
				Vid:           testconstants.Vid,
			},
			nocRoorCert: testconstants.RootCertPem,
			err:         pkitypes.ErrCertificateAlreadyExists,
		},
		{
			name: "ExistingNonRootCert",
			existingCert: &types.Certificate{
				Issuer:        testconstants.TestIssuer,
				Subject:       testconstants.RootSubject,
				SubjectAsText: testconstants.RootSubjectAsText,
				SubjectKeyId:  testconstants.RootSubjectKeyID,
				SerialNumber:  testconstants.TestSerialNumber,
				IsRoot:        false,
				IsNoc:         true,
				Vid:           testconstants.Vid,
			},
			nocRoorCert: testconstants.RootCertPem,
			err:         sdkerrors.ErrUnauthorized,
		},
		{
			name: "ExistingNotNocCert",
			existingCert: &types.Certificate{
				Issuer:        testconstants.RootIssuer,
				Subject:       testconstants.RootSubject,
				SubjectAsText: testconstants.RootSubjectAsText,
				SubjectKeyId:  testconstants.RootSubjectKeyID,
				SerialNumber:  testconstants.TestSerialNumber,
				IsRoot:        true,
				IsNoc:         false,
				Vid:           testconstants.Vid,
			},
			nocRoorCert: testconstants.RootCertPem,
			err:         pkitypes.ErrInappropriateCertificateType,
		},
		{
			name: "ExistingCertWithDifferentVid",
			existingCert: &types.Certificate{
				Issuer:        testconstants.RootIssuer,
				Subject:       testconstants.RootSubject,
				SubjectAsText: testconstants.RootSubjectAsText,
				SubjectKeyId:  testconstants.RootSubjectKeyID,
				SerialNumber:  testconstants.GoogleSerialNumber,
				IsRoot:        true,
				IsNoc:         true,
				Vid:           testconstants.VendorID1,
			},
			nocRoorCert: testconstants.RootCertPem,
			err:         pkitypes.ErrCertVidNotEqualAccountVid,
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

			addNocX509RootCert := types.NewMsgAddNocX509RootCert(accAddress.String(), tc.nocRoorCert)
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
	newNocCertificate.IsNoc = true
	newNocCertificate.Approvals = nil
	newNocCertificate.Rejects = nil

	// add the new NOC root certificate
	addNocX509RootCert := types.NewMsgAddNocX509RootCert(accAddress.String(), newNocCertificate.PemCert)
	_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
	require.NoError(t, err)

	// query approved certificate
	approvedCertificate, err := querySingleApprovedCertificate(setup, newNocCertificate.Subject, newNocCertificate.SubjectKeyId)
	require.NoError(t, err)
	require.Equal(t, &newNocCertificate, approvedCertificate)

	// query noc root certificate by VID
	nocRootCertificate, err := querySingleNocRootCertificate(setup, testconstants.Vid)
	require.NoError(t, err)
	require.Equal(t, &newNocCertificate, nocRootCertificate)

	// check that unique certificate key stays registered
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
	nocRootCertificate.IsNoc = true
	nocRootCertificate.Approvals = nil
	nocRootCertificate.Rejects = nil

	setup.Keeper.AddApprovedCertificate(setup.Ctx, nocRootCertificate)
	setup.Keeper.AddNocRootCertificate(setup.Ctx, nocRootCertificate)
	uniqueCertificate := types.UniqueCertificate{
		Issuer:       nocRootCertificate.Issuer,
		SerialNumber: nocRootCertificate.SerialNumber,
		Present:      true,
	}
	setup.Keeper.SetUniqueCertificate(setup.Ctx, uniqueCertificate)

	// new NOC root certificate
	newNocCertificate := rootCertificate(accAddress)
	newNocCertificate.IsNoc = true
	newNocCertificate.Approvals = nil
	newNocCertificate.Rejects = nil

	// add the new NOC root certificate
	addNocX509RootCert := types.NewMsgAddNocX509RootCert(accAddress.String(), newNocCertificate.PemCert)
	_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
	require.NoError(t, err)

	// query approved certificate
	approvedCertificates, err := queryApprovedCertificates(setup, newNocCertificate.Subject, newNocCertificate.SubjectKeyId)
	require.NoError(t, err)
	require.Equal(t, len(approvedCertificates.Certs), 2)
	require.Equal(t, &newNocCertificate, approvedCertificates.Certs[1])

	// query noc root certificate by VID
	nocRootCertificates, err := queryNocRootCertificates(setup, testconstants.Vid)
	require.NoError(t, err)
	require.Equal(t, len(nocRootCertificates.Certs), 2)
	require.Equal(t, &newNocCertificate, nocRootCertificates.Certs[1])
}

func querySingleNocRootCertificate(
	setup *TestSetup,
	vid int32,
) (*types.Certificate, error) {
	certificates, err := queryNocRootCertificates(setup, vid)
	if err != nil {
		return nil, err
	}

	if len(certificates.Certs) > 1 {
		require.Fail(setup.T, "More than 1 certificate returned")
	}

	return certificates.Certs[0], nil
}

func queryNocRootCertificates(
	setup *TestSetup,
	vid int32,
) (*types.NocRootCertificates, error) {
	// query certificate
	req := &types.QueryGetNocRootCertificatesRequest{Vid: vid}

	resp, err := setup.Keeper.NocRootCertificates(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.NocRootCertificates, nil
}
