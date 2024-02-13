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
				Vid:           testconstants.Vid,
			},
			nocRoorCert: testconstants.RootCertPem,
			err:         pkitypes.ErrCertificateAlreadyExists,
		},
		{
			name: "ExistingNonRootCert",
			existingCert: &types.Certificate{
				Issuer:        testconstants.GoogleIssuer,
				Subject:       testconstants.RootSubject,
				SubjectAsText: testconstants.RootSubjectAsText,
				SubjectKeyId:  testconstants.RootSubjectKeyID,
				SerialNumber:  testconstants.GoogleSerialNumber,
				IsRoot:        false,
				Vid:           testconstants.Vid,
			},
			nocRoorCert: testconstants.RootCertPem,
			err:         sdkerrors.ErrUnauthorized,
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

			// add the existing certificate key
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

func TestHandler_AddNocX509RootCert(t *testing.T) {
	setup := Setup(t)

	accAddress := GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	addNocX509RootCert := types.NewMsgAddNocX509RootCert(accAddress.String(), testconstants.RootCertPem)
	_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
	require.NoError(t, err)

	// query approved certificate
	approvedCertificate, err := querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, testconstants.RootCertPem, approvedCertificate.PemCert)
	require.Equal(t, testconstants.RootSubject, approvedCertificate.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, approvedCertificate.SubjectKeyId)
	require.Equal(t, testconstants.RootSerialNumber, approvedCertificate.SerialNumber)
	require.Equal(t, testconstants.RootIssuer, approvedCertificate.Issuer)

	// query noc root certificate by VID
	rootNocCertificate, err := querySingleNocRootCertificate(setup, testconstants.Vid)
	require.NoError(t, err)
	require.Equal(t, testconstants.RootCertPem, rootNocCertificate.PemCert)
	require.Equal(t, testconstants.RootSubject, rootNocCertificate.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, rootNocCertificate.SubjectKeyId)
	require.Equal(t, testconstants.RootSerialNumber, rootNocCertificate.SerialNumber)
	require.Equal(t, testconstants.RootIssuer, rootNocCertificate.Issuer)
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
