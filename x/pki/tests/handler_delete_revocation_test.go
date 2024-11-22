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

func TestHandler_DeletePkiRevocationDistributionPoint_NegativeCases(t *testing.T) {
	accAddress := GenerateAccAddress()
	vendorAcc := GenerateAccAddress()

	cases := []struct {
		name             string
		accountVid       int32
		accountRole      dclauthtypes.AccountRole
		vendorAccVid     int32
		rootCertOptions  *rootCertOptions
		addRevocation    *types.MsgAddPkiRevocationDistributionPoint
		deleteRevocation *types.MsgDeletePkiRevocationDistributionPoint
		err              error
	}{
		{
			name:            "PAASenderNotVendor",
			accountVid:      testconstants.PAACertWithNumericVidVid,
			accountRole:     dclauthtypes.CertificationCenter,
			vendorAccVid:    testconstants.PAACertWithNumericVidVid,
			rootCertOptions: createPAACertWithNumericVidOptions(),
			addRevocation:   createAddRevocationMessageWithPAACertWithNumericVid(vendorAcc.String()),
			deleteRevocation: &types.MsgDeletePkiRevocationDistributionPoint{
				Signer:             accAddress.String(),
				Vid:                testconstants.PAACertWithNumericVidVid,
				Label:              label,
				IssuerSubjectKeyID: testconstants.SubjectKeyIDWithoutColons,
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			name:            "PAISenderNotVendor",
			accountVid:      testconstants.PAACertWithNumericVidVid,
			accountRole:     dclauthtypes.CertificationCenter,
			vendorAccVid:    testconstants.PAACertWithNumericVidVid,
			rootCertOptions: createPAACertWithNumericVidOptions(),
			addRevocation:   createAddRevocationMessageWithPAICertWithNumericVidPid(vendorAcc.String()),
			deleteRevocation: &types.MsgDeletePkiRevocationDistributionPoint{
				Signer:             accAddress.String(),
				Vid:                testconstants.PAICertWithNumericPidVidVid,
				Label:              label,
				IssuerSubjectKeyID: testconstants.SubjectKeyIDWithoutColons,
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			name:         "RevocationPointNotFound",
			vendorAccVid: testconstants.VendorID1,
			deleteRevocation: &types.MsgDeletePkiRevocationDistributionPoint{
				Signer:             vendorAcc.String(),
				Vid:                testconstants.VendorID1,
				Label:              label,
				IssuerSubjectKeyID: testconstants.SubjectKeyIDWithoutColons,
			},
			err: pkitypes.ErrPkiRevocationDistributionPointDoesNotExists,
		},
		{
			name:            "PAASenderVidNotEqualCertVid",
			accountVid:      testconstants.VendorID1,
			accountRole:     dclauthtypes.Vendor,
			vendorAccVid:    testconstants.PAACertWithNumericVidVid,
			rootCertOptions: createPAACertWithNumericVidOptions(),
			addRevocation:   createAddRevocationMessageWithPAICertWithNumericVidPid(vendorAcc.String()),
			deleteRevocation: &types.MsgDeletePkiRevocationDistributionPoint{
				Signer:             accAddress.String(),
				Vid:                testconstants.PAACertWithNumericVidVid,
				Label:              label,
				IssuerSubjectKeyID: testconstants.SubjectKeyIDWithoutColons,
			},
			err: pkitypes.ErrMessageVidNotEqualAccountVid,
		},
		{
			name:            "PAISenderVidNotEqualCertVid",
			accountVid:      testconstants.VendorID1,
			accountRole:     dclauthtypes.Vendor,
			vendorAccVid:    testconstants.PAICertWithNumericPidVidVid,
			rootCertOptions: createPAACertWithNumericVidOptions(),
			addRevocation:   createAddRevocationMessageWithPAICertWithNumericVidPid(vendorAcc.String()),
			deleteRevocation: &types.MsgDeletePkiRevocationDistributionPoint{
				Signer:             accAddress.String(),
				Vid:                testconstants.PAICertWithNumericPidVidVid,
				Label:              label,
				IssuerSubjectKeyID: testconstants.SubjectKeyIDWithoutColons,
			},
			err: pkitypes.ErrMessageVidNotEqualAccountVid,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := Setup(t)

			setup.AddAccount(accAddress, []dclauthtypes.AccountRole{tc.accountRole}, tc.accountVid)
			setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, tc.vendorAccVid)

			if tc.rootCertOptions != nil {
				proposeAndApproveRootCertificate(setup, setup.Trustee1, tc.rootCertOptions)
			}

			if tc.addRevocation != nil {
				_, err := setup.Handler(setup.Ctx, tc.addRevocation)
				require.NoError(t, err)
			}

			_, err := setup.Handler(setup.Ctx, tc.deleteRevocation)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestHandler_DeletePkiRevocationDistributionPoint_PositiveCases(t *testing.T) {
	vendorAcc := GenerateAccAddress()

	cases := []struct {
		name             string
		rootCertOptions  *rootCertOptions
		addRevocation    *types.MsgAddPkiRevocationDistributionPoint
		deleteRevocation *types.MsgDeletePkiRevocationDistributionPoint
	}{
		{
			name:            "PAA",
			rootCertOptions: createPAACertWithNumericVidOptions(),
			addRevocation:   createAddRevocationMessageWithPAACertWithNumericVid(vendorAcc.String()),
			deleteRevocation: &types.MsgDeletePkiRevocationDistributionPoint{
				Signer:             vendorAcc.String(),
				Vid:                testconstants.PAACertWithNumericVidVid,
				Label:              label,
				IssuerSubjectKeyID: testconstants.SubjectKeyIDWithoutColons,
			},
		},
		{
			name:            "PAI",
			rootCertOptions: createPAACertWithNumericVidOptions(),
			addRevocation:   createAddRevocationMessageWithPAICertWithNumericVidPid(vendorAcc.String()),
			deleteRevocation: &types.MsgDeletePkiRevocationDistributionPoint{
				Signer:             vendorAcc.String(),
				Vid:                testconstants.PAICertWithNumericPidVidVid,
				Label:              label,
				IssuerSubjectKeyID: testconstants.SubjectKeyIDWithoutColons,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := Setup(t)

			setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, tc.deleteRevocation.Vid)

			proposeAndApproveRootCertificate(setup, setup.Trustee1, tc.rootCertOptions)

			_, err := setup.Handler(setup.Ctx, tc.addRevocation)
			require.NoError(t, err)

			_, err = setup.Handler(setup.Ctx, tc.deleteRevocation)
			require.NoError(t, err)

			_, isFound := setup.Keeper.GetPkiRevocationDistributionPoint(setup.Ctx, testconstants.PAACertWithNumericVidVid, label, testconstants.SubjectKeyIDWithoutColons)
			require.False(t, isFound)

			_, isFound = setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
			require.False(t, isFound)
		})
	}
}

func TestHandler_DeletePkiRevocationDistributionPoint_Multiple_SameIssuerSubjectKeyId(t *testing.T) {
	setup := Setup(t)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.PAACertWithNumericVidVid)

	// add PAA NOVID
	rootCertOptions := createPAACertNoVidOptions(testconstants.PAACertWithNumericVidVid)
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// add PAA VID
	rootCertOptions = createPAACertWithNumericVidOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// add Revocation Point PAA NOVID
	addRevocationPAANoVid := createAddRevocationMessageWithPAACertNoVid(vendorAcc.String(), testconstants.PAACertWithNumericVidVid)
	_, err := setup.Handler(setup.Ctx, addRevocationPAANoVid)
	require.NoError(t, err)

	// add Revocation Point PAA VID
	addRevocationPAAWithVid := createAddRevocationMessageWithPAACertWithNumericVid(vendorAcc.String())
	addRevocationPAAWithVid.Label = "label2"
	addRevocationPAAWithVid.DataURL = testconstants.DataURL2
	_, err = setup.Handler(setup.Ctx, addRevocationPAAWithVid)
	require.NoError(t, err)

	revocationPointBySubjectKeyID, isFound := setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	require.Equal(t, len(revocationPointBySubjectKeyID.Points), 2)
	assertRevocationPointEqual(t, addRevocationPAANoVid, revocationPointBySubjectKeyID.Points[0])
	assertRevocationPointEqual(t, addRevocationPAAWithVid, revocationPointBySubjectKeyID.Points[1])

	deleteRevocationPAANoVid := types.MsgDeletePkiRevocationDistributionPoint{
		Signer:             vendorAcc.String(),
		Vid:                addRevocationPAANoVid.Vid,
		Label:              addRevocationPAANoVid.Label,
		IssuerSubjectKeyID: testconstants.SubjectKeyIDWithoutColons,
	}
	_, err = setup.Handler(setup.Ctx, &deleteRevocationPAANoVid)
	require.NoError(t, err)

	revocationPointBySubjectKeyID, isFound = setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	require.Equal(t, len(revocationPointBySubjectKeyID.Points), 1)
	assertRevocationPointEqual(t, addRevocationPAAWithVid, revocationPointBySubjectKeyID.Points[0])

	deletePkiRevocationDistributionPoint := types.MsgDeletePkiRevocationDistributionPoint{
		Signer:             vendorAcc.String(),
		Vid:                addRevocationPAAWithVid.Vid,
		Label:              addRevocationPAAWithVid.Label,
		IssuerSubjectKeyID: testconstants.SubjectKeyIDWithoutColons,
	}
	_, err = setup.Handler(setup.Ctx, &deletePkiRevocationDistributionPoint)
	require.NoError(t, err)

	_, isFound = setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
	require.False(t, isFound)
}
