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

func TestHandler_AddPkiRevocationDistributionPoint_negativeCases(t *testing.T) {
	accAddress := GenerateAccAddress()

	cases := []struct {
		name            string
		accountVid      int32
		accountRole     dclauthtypes.AccountRole
		rootCertOptions *rootCertOptions
		addRevocation   *types.MsgAddPkiRevocationDistributionPoint
		err             error
	}{
		{
			name:          "PAASenderNotVendor",
			accountVid:    testconstants.PAACertWithNumericVidVid,
			accountRole:   dclauthtypes.CertificationCenter,
			addRevocation: createAddRevocationMessageWithPAACertWithNumericVid(accAddress.String()),
			err:           sdkerrors.ErrUnauthorized,
		},
		{
			name:          "PAISenderNotVendor",
			accountVid:    testconstants.PAICertWithNumericPidVid_Vid,
			accountRole:   dclauthtypes.CertificationCenter,
			addRevocation: createAddRevocationMessageWithPAICertWithNumericVidPid(accAddress.String()),
			err:           sdkerrors.ErrUnauthorized,
		},
		{
			name:          "PAACertEncodesVidSenderVidNotEqualVidField",
			accountVid:    testconstants.Vid,
			accountRole:   dclauthtypes.Vendor,
			addRevocation: createAddRevocationMessageWithPAACertWithNumericVid(accAddress.String()),
			err:           pkitypes.ErrCRLSignerCertificateVidNotEqualAccountVid,
		},
		{
			name:          "PAACertNotFound",
			accountVid:    testconstants.PAACertWithNumericVidVid,
			accountRole:   dclauthtypes.Vendor,
			addRevocation: createAddRevocationMessageWithPAACertWithNumericVid(accAddress.String()),
			err:           pkitypes.ErrCertificateDoesNotExist,
		},
		{
			name:          "PAINotChainedBackToDCLCerts",
			accountVid:    testconstants.PAACertWithNumericVidVid,
			accountRole:   dclauthtypes.Vendor,
			addRevocation: createAddRevocationMessageWithPAICertWithNumericVidPid(accAddress.String()),
			err:           pkitypes.ErrCertNotChainedBack,
		},
		{
			name:        "InvalidCertificate",
			accountVid:  testconstants.PAACertWithNumericVidVid,
			accountRole: dclauthtypes.Vendor,
			addRevocation: &types.MsgAddPkiRevocationDistributionPoint{
				Signer:               accAddress.String(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				IsPAA:                true,
				Pid:                  0,
				CrlSignerCertificate: "invalidpem",
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       types.CRLRevocationType,
			},
			err: pkitypes.ErrInvalidCertificate,
		},
		{
			name:            "PAANotOnLedger",
			accountVid:      testconstants.PAACertWithNumericVidVid,
			accountRole:     dclauthtypes.Vendor,
			rootCertOptions: createTestRootCertOptions(),
			addRevocation:   createAddRevocationMessageWithPAACertWithNumericVid(accAddress.String()),
			err:             pkitypes.ErrCertificateDoesNotExist,
		},
		{
			name:            "PAANoVid_LedgerPAANoVid",
			accountVid:      testconstants.Vid,
			accountRole:     dclauthtypes.Vendor,
			rootCertOptions: createPAACertNoVidOptions(),
			addRevocation:   createAddRevocationMessageWithPAACertNoVid(accAddress.String(), testconstants.Vid),
			err:             pkitypes.ErrMessageVidNotEqualRootCertVid,
		},
		{
			name:        "PAANoVid_WrongVID",
			accountVid:  testconstants.Vid,
			accountRole: dclauthtypes.Vendor,
			rootCertOptions: &rootCertOptions{
				pemCert:      testconstants.PAACertNoVid,
				info:         testconstants.Info,
				subject:      testconstants.PAACertNoVidSubject,
				subjectKeyID: testconstants.PAACertNoVidSubjectKeyID,
				vid:          1001,
			},
			addRevocation: createAddRevocationMessageWithPAACertNoVid(accAddress.String(), testconstants.Vid),
			err:           pkitypes.ErrMessageVidNotEqualRootCertVid,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := Setup(t)

			setup.AddAccount(accAddress, []dclauthtypes.AccountRole{tc.accountRole}, tc.accountVid)

			if tc.rootCertOptions != nil {
				proposeAndApproveRootCertificate(setup, setup.Trustee1, tc.rootCertOptions)
			}

			_, err := setup.Handler(setup.Ctx, tc.addRevocation)
			require.ErrorIs(t, err, tc.err)

		})
	}
}

func TestHandler_AddPkiRevocationDistributionPoint_PAAAlreadyExists(t *testing.T) {
	setup := Setup(t)

	accAddress := GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.PAACertWithNumericVidVid)

	// propose and approve x509 root certificate
	rootCertOptions := createPAACertWithNumericVidOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	addPkiRevocationDistributionPoint := createAddRevocationMessageWithPAACertWithNumericVid(accAddress.String())

	_, err := setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	_, err = setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrPkiRevocationDistributionPointAlreadyExists)
}

func TestHandler_AddPkiRevocationDistributionPoint_positiveCases(t *testing.T) {
	vendorAcc := GenerateAccAddress()

	cases := []struct {
		name            string
		rootCertOptions *rootCertOptions
		addRevocation   *types.MsgAddPkiRevocationDistributionPoint
	}{
		{
			name:            "PAAWithVid",
			rootCertOptions: createPAACertWithNumericVidOptions(),
			addRevocation:   createAddRevocationMessageWithPAACertWithNumericVid(vendorAcc.String()),
		},
		{
			name:            "PAIWithNumericVidPid",
			rootCertOptions: createPAACertWithNumericVidOptions(),
			addRevocation:   createAddRevocationMessageWithPAICertWithNumericVidPid(vendorAcc.String()),
		},
		{
			name:            "PAIWithStringVidPid",
			rootCertOptions: createPAACertNoVidOptions(),
			addRevocation:   createAddRevocationMessageWithPAICertWithVidPid(vendorAcc.String()),
		},
		{
			name: "PAANoVid",
			rootCertOptions: &rootCertOptions{
				pemCert:      testconstants.PAACertNoVid,
				info:         testconstants.Info,
				subject:      testconstants.PAACertNoVidSubject,
				subjectKeyID: testconstants.PAACertNoVidSubjectKeyID,
				vid:          1001,
			},
			addRevocation: createAddRevocationMessageWithPAACertNoVid(vendorAcc.String(), 1001),
		},
		{
			name:            "PAIWithVid",
			rootCertOptions: createPAACertNoVidOptions(),
			addRevocation: &types.MsgAddPkiRevocationDistributionPoint{
				Signer:               vendorAcc.String(),
				Vid:                  testconstants.PAICertWithVid_Vid,
				IsPAA:                false,
				Pid:                  0,
				CrlSignerCertificate: testconstants.PAICertWithVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       types.CRLRevocationType,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := Setup(t)
			setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, tc.addRevocation.Vid)

			proposeAndApproveRootCertificate(setup, setup.Trustee1, tc.rootCertOptions)

			_, err := setup.Handler(setup.Ctx, tc.addRevocation)
			require.NoError(t, err)

			revocationPoint, isFound := setup.Keeper.GetPkiRevocationDistributionPoint(setup.Ctx, tc.addRevocation.Vid, "label", testconstants.SubjectKeyIDWithoutColons)
			require.True(t, isFound)
			assertRevocationPointEqual(t, tc.addRevocation, &revocationPoint)

			revocationPointBySubjectKeyID, isFound := setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
			require.True(t, isFound)
			assertRevocationPointEqual(t, tc.addRevocation, revocationPointBySubjectKeyID.Points[0])
		})
	}
}
