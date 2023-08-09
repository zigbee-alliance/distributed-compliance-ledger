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

func TestHandler_UpdatePkiRevocationDistributionPoint_PAA_VID(t *testing.T) {
	var err error
	vendorAcc := GenerateAccAddress()
	addedRevocation := &types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  testconstants.PAACertWithNumericVidVid,
		IsPAA:                true,
		Pid:                  0,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	cases := []struct {
		valid             bool
		name              string
		updatedRevocation types.MsgUpdatePkiRevocationDistributionPoint
		err               error
	}{
		{
			valid: true,
			name:  "Valid: PAAWithVid",
			updatedRevocation: types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:               addedRevocation.Signer,
				Vid:                  addedRevocation.Vid,
				CrlSignerCertificate: addedRevocation.CrlSignerCertificate,
				Label:                addedRevocation.Label,
				DataURL:              addedRevocation.DataURL,
				IssuerSubjectKeyID:   addedRevocation.IssuerSubjectKeyID,
			},
		},
		{
			valid: true,
			name:  "Valid: MinimalParams",
			updatedRevocation: types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:             addedRevocation.Signer,
				Vid:                addedRevocation.Vid,
				Label:              addedRevocation.Label,
				IssuerSubjectKeyID: addedRevocation.IssuerSubjectKeyID,
			},
		},
		{
			valid: true,
			name:  "Valid: AllParams",
			updatedRevocation: types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:             addedRevocation.Signer,
				Vid:                addedRevocation.Vid,
				Label:              addedRevocation.Label,
				DataURL:            addedRevocation.DataURL + "/new",
				IssuerSubjectKeyID: addedRevocation.IssuerSubjectKeyID,
			},
		},
		{
			valid: false,
			name:  "Invalid: DataFieldsProvidedWhenRevocationType1",
			updatedRevocation: types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:             addedRevocation.Signer,
				Vid:                addedRevocation.Vid,
				Label:              addedRevocation.Label,
				DataURL:            addedRevocation.DataURL + "/new",
				DataFileSize:       uint64(123),
				DataDigest:         addedRevocation.DataDigest + "new",
				DataDigestType:     1,
				IssuerSubjectKeyID: addedRevocation.IssuerSubjectKeyID,
			},
			err: pkitypes.ErrDataFieldPresented,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := Setup(t)
			setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.PAACertWithNumericVidVid)

			// propose x509 root certificate by account Trustee1
			proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.PAACertWithNumericVid, testconstants.Info, testconstants.PAACertWithNumericVidVid)
			_, err = setup.Handler(setup.Ctx, proposeAddX509RootCert)
			require.NoError(t, err)

			// approve
			approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
				setup.Trustee2.String(), testconstants.PAACertWithNumericVidSubject, testconstants.PAACertWithNumericVidSubjectKeyID, testconstants.Info)
			_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
			require.NoError(t, err)

			// add revocation
			if addedRevocation != nil {
				_, err = setup.Handler(setup.Ctx, addedRevocation)
				require.NoError(t, err)
			}
			_, err = setup.Handler(setup.Ctx, &tc.updatedRevocation)
			updatedPoint, isFound := setup.Keeper.GetPkiRevocationDistributionPoint(setup.Ctx, addedRevocation.Vid, addedRevocation.Label, addedRevocation.IssuerSubjectKeyID)

			if tc.valid {
				require.NoError(t, err)
				require.True(t, isFound)
				require.Equal(t, updatedPoint.Vid, addedRevocation.Vid)
				require.Equal(t, updatedPoint.Pid, addedRevocation.Pid)
				require.Equal(t, updatedPoint.IsPAA, addedRevocation.IsPAA)
				require.Equal(t, updatedPoint.Label, addedRevocation.Label)
				require.Equal(t, updatedPoint.IssuerSubjectKeyID, addedRevocation.IssuerSubjectKeyID)
				require.Equal(t, updatedPoint.RevocationType, addedRevocation.RevocationType)

				compareUpdatedStringFields(t, addedRevocation.DataURL, tc.updatedRevocation.DataURL, updatedPoint.DataURL)
				compareUpdatedStringFields(t, addedRevocation.DataDigest, tc.updatedRevocation.DataDigest, updatedPoint.DataDigest)
				compareUpdatedStringFields(t, addedRevocation.CrlSignerCertificate, tc.updatedRevocation.CrlSignerCertificate, updatedPoint.CrlSignerCertificate)
				compareUpdatedIntFields(t, int(addedRevocation.DataDigestType), int(tc.updatedRevocation.DataDigestType), int(updatedPoint.DataDigestType))
				compareUpdatedIntFields(t, int(addedRevocation.DataFileSize), int(tc.updatedRevocation.DataFileSize), int(updatedPoint.DataFileSize))
			} else {
				require.ErrorIs(t, err, tc.err)
			}
		})
	}
}

func TestHandler_UpdatePkiRevocationDistributionPoint_PAA_NOVID(t *testing.T) {
	var err error
	vendorAcc := GenerateAccAddress()
	addedRevocation := &types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  1001,
		IsPAA:                true,
		Pid:                  0,
		CrlSignerCertificate: testconstants.PAACertNoVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	cases := []struct {
		valid             bool
		name              string
		updatedRevocation types.MsgUpdatePkiRevocationDistributionPoint
		err               error
	}{
		{
			valid: true,
			name:  "Valid: Same PAA",
			updatedRevocation: types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:               addedRevocation.Signer,
				Vid:                  addedRevocation.Vid,
				CrlSignerCertificate: addedRevocation.CrlSignerCertificate,
				Label:                addedRevocation.Label,
				DataURL:              addedRevocation.DataURL,
				IssuerSubjectKeyID:   addedRevocation.IssuerSubjectKeyID,
			},
		},
		{
			valid: true,
			name:  "Valid: MinimalParams",
			updatedRevocation: types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:             addedRevocation.Signer,
				Vid:                addedRevocation.Vid,
				Label:              addedRevocation.Label,
				IssuerSubjectKeyID: addedRevocation.IssuerSubjectKeyID,
			},
		},
		{
			valid: true,
			name:  "Valid: AllParams",
			updatedRevocation: types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:             addedRevocation.Signer,
				Vid:                addedRevocation.Vid,
				Label:              addedRevocation.Label,
				DataURL:            addedRevocation.DataURL + "/new",
				IssuerSubjectKeyID: addedRevocation.IssuerSubjectKeyID,
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := Setup(t)
			setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 1001)

			// propose x509 root certificate by account Trustee1
			proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.PAACertNoVid, testconstants.Info, 1001)
			_, err = setup.Handler(setup.Ctx, proposeAddX509RootCert)
			require.NoError(t, err)

			// approve
			approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
				setup.Trustee2.String(), testconstants.PAACertNoVidSubject, testconstants.PAACertNoVidSubjectKeyID, testconstants.Info)
			_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
			require.NoError(t, err)

			// add revocation
			if addedRevocation != nil {
				_, err = setup.Handler(setup.Ctx, addedRevocation)
				require.NoError(t, err)
			}
			_, err = setup.Handler(setup.Ctx, &tc.updatedRevocation)
			updatedPoint, isFound := setup.Keeper.GetPkiRevocationDistributionPoint(setup.Ctx, addedRevocation.Vid, addedRevocation.Label, addedRevocation.IssuerSubjectKeyID)

			if tc.valid {
				require.NoError(t, err)
				require.True(t, isFound)
				require.Equal(t, updatedPoint.Vid, addedRevocation.Vid)
				require.Equal(t, updatedPoint.Pid, addedRevocation.Pid)
				require.Equal(t, updatedPoint.IsPAA, addedRevocation.IsPAA)
				require.Equal(t, updatedPoint.Label, addedRevocation.Label)
				require.Equal(t, updatedPoint.IssuerSubjectKeyID, addedRevocation.IssuerSubjectKeyID)
				require.Equal(t, updatedPoint.RevocationType, addedRevocation.RevocationType)

				compareUpdatedStringFields(t, addedRevocation.DataURL, tc.updatedRevocation.DataURL, updatedPoint.DataURL)
				compareUpdatedStringFields(t, addedRevocation.DataDigest, tc.updatedRevocation.DataDigest, updatedPoint.DataDigest)
				compareUpdatedStringFields(t, addedRevocation.CrlSignerCertificate, tc.updatedRevocation.CrlSignerCertificate, updatedPoint.CrlSignerCertificate)
				compareUpdatedIntFields(t, int(addedRevocation.DataDigestType), int(tc.updatedRevocation.DataDigestType), int(updatedPoint.DataDigestType))
				compareUpdatedIntFields(t, int(addedRevocation.DataFileSize), int(tc.updatedRevocation.DataFileSize), int(updatedPoint.DataFileSize))
			} else {
				require.ErrorIs(t, err, tc.err)
			}
		})
	}
}

func TestHandler_UpdatePkiRevocationDistributionPoint_PAI_VIDPID(t *testing.T) {
	var err error
	vendorAcc := GenerateAccAddress()
	addedRevocation := &types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  65521,
		IsPAA:                false,
		Pid:                  8,
		CrlSignerCertificate: testconstants.PAICertWithNumericPidVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	cases := []struct {
		valid             bool
		name              string
		updatedRevocation types.MsgUpdatePkiRevocationDistributionPoint
		err               error
	}{
		{
			valid: true,
			name:  "Valid: Same PAI",
			updatedRevocation: types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:               addedRevocation.Signer,
				Vid:                  addedRevocation.Vid,
				CrlSignerCertificate: addedRevocation.CrlSignerCertificate,
				Label:                addedRevocation.Label,
				DataURL:              addedRevocation.DataURL,
				IssuerSubjectKeyID:   addedRevocation.IssuerSubjectKeyID,
			},
		},
		{
			valid: true,
			name:  "Valid: MinimalParams",
			updatedRevocation: types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:             addedRevocation.Signer,
				Vid:                addedRevocation.Vid,
				Label:              addedRevocation.Label,
				IssuerSubjectKeyID: addedRevocation.IssuerSubjectKeyID,
			},
		},
		{
			valid: true,
			name:  "Valid: AllParams",
			updatedRevocation: types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:             addedRevocation.Signer,
				Vid:                addedRevocation.Vid,
				Label:              addedRevocation.Label,
				DataURL:            addedRevocation.DataURL + "/new",
				IssuerSubjectKeyID: addedRevocation.IssuerSubjectKeyID,
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := Setup(t)
			setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 65521)

			// propose x509 root certificate by account Trustee1
			proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.PAACertWithNumericVid, testconstants.Info, testconstants.PAACertWithNumericVidVid)
			_, err = setup.Handler(setup.Ctx, proposeAddX509RootCert)
			require.NoError(t, err)

			// approve
			approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
				setup.Trustee2.String(), testconstants.PAACertWithNumericVidSubject, testconstants.PAACertWithNumericVidSubjectKeyID, testconstants.Info)
			_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
			require.NoError(t, err)

			// add revocation
			if addedRevocation != nil {
				_, err = setup.Handler(setup.Ctx, addedRevocation)
				require.NoError(t, err)
			}
			_, err = setup.Handler(setup.Ctx, &tc.updatedRevocation)
			updatedPoint, isFound := setup.Keeper.GetPkiRevocationDistributionPoint(setup.Ctx, addedRevocation.Vid, addedRevocation.Label, addedRevocation.IssuerSubjectKeyID)

			if tc.valid {
				require.NoError(t, err)
				require.True(t, isFound)
				require.Equal(t, updatedPoint.Vid, addedRevocation.Vid)
				require.Equal(t, updatedPoint.Pid, addedRevocation.Pid)
				require.Equal(t, updatedPoint.IsPAA, addedRevocation.IsPAA)
				require.Equal(t, updatedPoint.Label, addedRevocation.Label)
				require.Equal(t, updatedPoint.IssuerSubjectKeyID, addedRevocation.IssuerSubjectKeyID)
				require.Equal(t, updatedPoint.RevocationType, addedRevocation.RevocationType)

				compareUpdatedStringFields(t, addedRevocation.DataURL, tc.updatedRevocation.DataURL, updatedPoint.DataURL)
				compareUpdatedStringFields(t, addedRevocation.DataDigest, tc.updatedRevocation.DataDigest, updatedPoint.DataDigest)
				compareUpdatedStringFields(t, addedRevocation.CrlSignerCertificate, tc.updatedRevocation.CrlSignerCertificate, updatedPoint.CrlSignerCertificate)
				compareUpdatedIntFields(t, int(addedRevocation.DataDigestType), int(tc.updatedRevocation.DataDigestType), int(updatedPoint.DataDigestType))
				compareUpdatedIntFields(t, int(addedRevocation.DataFileSize), int(tc.updatedRevocation.DataFileSize), int(updatedPoint.DataFileSize))
			} else {
				require.ErrorIs(t, err, tc.err)
			}
		})
	}
}

func compareUpdatedStringFields(t *testing.T, oldValue string, newValue string, updatedValue string) {
	if newValue == "" {
		require.Equal(t, oldValue, updatedValue)
	} else {
		require.Equal(t, newValue, updatedValue)
	}
}

func compareUpdatedIntFields(t *testing.T, oldValue int, newValue int, updatedValue int) {
	if newValue == 0 {
		require.Equal(t, oldValue, updatedValue)
	} else {
		require.Equal(t, newValue, updatedValue)
	}
}

func TestHandler_UpdatePkiRevocationDistributionPoint_PAIWithoutPid(t *testing.T) {
	setup := Setup(t)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 65522)

	// propose x509 root certificate by account Trustee1
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.PAACertNoVid, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.PAACertNoVidSubject, testconstants.PAACertNoVidSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// add intermediate certificate
	addX509Cert := types.NewMsgAddX509Cert(vendorAcc.String(), testconstants.PAICertWithPidVid)
	_, err = setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	addPkiRevocationDistributionPoint := types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  65522,
		IsPAA:                false,
		CrlSignerCertificate: testconstants.PAICertWithPidVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	updatePkiRevocationDistributionPoint := types.MsgUpdatePkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  65522,
		CrlSignerCertificate: testconstants.PAICertWithVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
	}
	_, err = setup.Handler(setup.Ctx, &updatePkiRevocationDistributionPoint)
	require.NoError(t, err)
}

func TestHandler_UpdatePkiRevocationDistributionPoint_NotFound(t *testing.T) {
	setup := Setup(t)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 65521)

	updatePkiRevocationDistributionPoint := types.MsgUpdatePkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  testconstants.PAACertWithNumericVidVid,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
	}
	_, err := setup.Handler(setup.Ctx, &updatePkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrPkiRevocationDistributionPointDoesNotExists)
}

func TestHandler_UpdatePkiRevocationDistributionPoint_PAANewCertificateNotPAA(t *testing.T) {
	setup := Setup(t)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 65521)

	// propose x509 root certificate by account Trustee1
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.PAACertWithNumericVid, testconstants.Info, testconstants.PAACertWithNumericVidVid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.PAACertWithNumericVidSubject, testconstants.PAACertWithNumericVidSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	addPkiRevocationDistributionPoint := types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  testconstants.PAACertWithNumericVidVid,
		IsPAA:                true,
		Pid:                  8,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	updatePkiRevocationDistributionPoint := types.MsgUpdatePkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  65521,
		CrlSignerCertificate: testconstants.PAICertWithNumericPidVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
	}
	_, err = setup.Handler(setup.Ctx, &updatePkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrRootCertificateIsNotSelfSigned)
}

func TestHandler_UpdatePkiRevocationDistributionPoint_PAASenderNotVendor(t *testing.T) {
	setup := Setup(t)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 65521)

	// propose x509 root certificate by account Trustee1
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.PAACertWithNumericVid, testconstants.Info, testconstants.PAACertWithNumericVidVid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.PAACertWithNumericVidSubject, testconstants.PAACertWithNumericVidSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	addPkiRevocationDistributionPoint := types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  testconstants.PAACertWithNumericVidVid,
		IsPAA:                true,
		Pid:                  8,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	updatePkiRevocationDistributionPoint := types.MsgUpdatePkiRevocationDistributionPoint{
		Signer:               setup.Trustee1.String(),
		Vid:                  testconstants.PAACertWithNumericVidVid,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
	}
	_, err = setup.Handler(setup.Ctx, &updatePkiRevocationDistributionPoint)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_UpdatePkiRevocationDistributionPoint_PAASenderVidNotEqualCertVid(t *testing.T) {
	setup := Setup(t)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 65521)

	// propose x509 root certificate by account Trustee1
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.PAACertWithNumericVid, testconstants.Info, testconstants.PAACertWithNumericVidVid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.PAACertWithNumericVidSubject, testconstants.PAACertWithNumericVidSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	addPkiRevocationDistributionPoint := types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  testconstants.PAACertWithNumericVidVid,
		IsPAA:                true,
		Pid:                  8,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	newVendorAcc := GenerateAccAddress()
	setup.AddAccount(newVendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 65522)

	updatePkiRevocationDistributionPoint := types.MsgUpdatePkiRevocationDistributionPoint{
		Signer:               newVendorAcc.String(),
		Vid:                  testconstants.PAACertWithNumericVidVid,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
	}
	_, err = setup.Handler(setup.Ctx, &updatePkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrCRLSignerCertificateVidNotEqualAccountVid)
}

func TestHandler_UpdatePkiRevocationDistributionPoint_PAISenderIsNotVendor(t *testing.T) {
	setup := Setup(t)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 65521)

	// propose x509 root certificate by account Trustee1
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.PAACertWithNumericVid, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.PAACertWithNumericVidSubject, testconstants.PAACertWithNumericVidSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	addPkiRevocationDistributionPoint := types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  65521,
		IsPAA:                true,
		Pid:                  8,
		CrlSignerCertificate: testconstants.PAICertWithNumericPidVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	updatePkiRevocationDistributionPoint := types.MsgUpdatePkiRevocationDistributionPoint{
		Signer:               setup.Trustee1.String(),
		Vid:                  65521,
		CrlSignerCertificate: testconstants.PAICertWithNumericPidVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
	}
	_, err = setup.Handler(setup.Ctx, &updatePkiRevocationDistributionPoint)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_UpdatePkiRevocationDistributionPoint_PAISenderVidNotEqualCertVid(t *testing.T) {
	setup := Setup(t)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 65521)

	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.PAACertNoVid, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.PAACertNoVidSubject, testconstants.PAACertNoVidSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	addX509Cert := types.NewMsgAddX509Cert(vendorAcc.String(), testconstants.PAICertWithPidVid)
	_, err = setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	addPkiRevocationDistributionPoint := types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  65521,
		IsPAA:                false,
		Pid:                  8,
		CrlSignerCertificate: testconstants.PAICertWithPidVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	newVendorAcc := GenerateAccAddress()
	setup.AddAccount(newVendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 65522)

	updatePkiRevocationDistributionPoint := types.MsgUpdatePkiRevocationDistributionPoint{
		Signer:               newVendorAcc.String(),
		Vid:                  65521,
		CrlSignerCertificate: testconstants.PAICertWithPidVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
	}
	_, err = setup.Handler(setup.Ctx, &updatePkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrCRLSignerCertificateVidNotEqualAccountVid)
}

func TestHandler_UpdatePkiRevocationDistributionPoint_PAICertVidNotEqualMsgVid(t *testing.T) {
	setup := Setup(t)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 65521)

	// propose x509 root certificate by account Trustee1
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.PAACertWithNumericVid, testconstants.Info, testconstants.PAACertWithNumericVidVid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.PAACertWithNumericVidSubject, testconstants.PAACertWithNumericVidSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	addX509Cert := types.NewMsgAddX509Cert(vendorAcc.String(), testconstants.PAICertWithNumericPidVid)
	_, err = setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	addPkiRevocationDistributionPoint := types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  65521,
		IsPAA:                false,
		Pid:                  8,
		CrlSignerCertificate: testconstants.PAICertWithNumericPidVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	updatePkiRevocationDistributionPoint := types.MsgUpdatePkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  65522,
		CrlSignerCertificate: testconstants.PAICertWithNumericPidVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
	}
	_, err = setup.Handler(setup.Ctx, &updatePkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrPkiRevocationDistributionPointDoesNotExists)
}

func TestHandler_UpdatePkiRevocationDistributionPoint_PAIPidNotFoundInNewCert(t *testing.T) {
	setup := Setup(t)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 65522)

	// propose x509 root certificate by account Trustee1
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.PAACertNoVid, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.PAACertNoVidSubject, testconstants.PAACertNoVidSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// add intermediate certificate
	addX509Cert := types.NewMsgAddX509Cert(vendorAcc.String(), testconstants.PAICertWithPidVid)
	_, err = setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	addPkiRevocationDistributionPoint := types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  65522,
		IsPAA:                false,
		Pid:                  8,
		CrlSignerCertificate: testconstants.PAICertWithPidVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	updatePkiRevocationDistributionPoint := types.MsgUpdatePkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  65522,
		CrlSignerCertificate: testconstants.PAICertWithVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
	}
	_, err = setup.Handler(setup.Ctx, &updatePkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrPidNotFound)
}

func TestHandler_UpdatePkiRevocationDistributionPoint_PAANotOnLedger(t *testing.T) {
	setup := Setup(t)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 65521)

	// propose x509 root certificate by account Trustee1
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.PAACertNoVid, testconstants.Info, 65521)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.PAACertNoVidSubject, testconstants.PAACertNoVidSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	addPkiRevocationDistributionPoint := types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  65521,
		IsPAA:                true,
		Pid:                  0,
		CrlSignerCertificate: testconstants.PAACertNoVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	updatePkiRevocationDistributionPoint := types.MsgUpdatePkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  testconstants.PAACertWithNumericVidVid,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
	}
	_, err = setup.Handler(setup.Ctx, &updatePkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrCertificateDoesNotExist)
}

func TestHandler_UpdatePkiRevocationDistributionPoint_PAI_NotChainedOnLedger(t *testing.T) {
	setup := Setup(t)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 65521)

	// propose x509 root certificate by account Trustee1
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.PAACertWithNumericVid, testconstants.Info, testconstants.PAACertWithNumericVidVid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.PAACertWithNumericVidSubject, testconstants.PAACertWithNumericVidSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	addPkiRevocationDistributionPoint := &types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  65521,
		IsPAA:                false,
		Pid:                  8,
		CrlSignerCertificate: testconstants.PAICertWithNumericPidVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	proposeRevokeRootCert := types.NewMsgProposeRevokeX509RootCert(setup.Trustee1.String(), testconstants.PAACertWithNumericVidSubject, testconstants.PAACertWithNumericVidSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, proposeRevokeRootCert)
	require.NoError(t, err)

	approveRevokeRootCert := types.NewMsgApproveRevokeX509RootCert(
		setup.Trustee2.String(), testconstants.PAACertWithNumericVidSubject, testconstants.PAACertWithNumericVidSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveRevokeRootCert)
	require.NoError(t, err)

	updatePkiRevocationDistributionPoint := types.MsgUpdatePkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  65521,
		CrlSignerCertificate: testconstants.PAICertWithNumericPidVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
	}
	_, err = setup.Handler(setup.Ctx, &updatePkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrCertNotChainedBack)
}

func TestHandler_UpdatePkiRevocationDistributionPoint_PAA_NOVID_TO_PAA_NOVID(t *testing.T) {
	setup := Setup(t)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 1001)

	// add PAA NOVID 1
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.PAACertNoVid, testconstants.Info, 1001)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.PAACertNoVidSubject, testconstants.PAACertNoVidSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// add PAA NOVID 2
	proposeAddX509RootCert = types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, 1001)
	_, err = setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)
	approveAddX509RootCert = types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// add Revocation Point PAA NOVID 1
	addPkiRevocationDistributionPoint := types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  1001,
		IsPAA:                true,
		Pid:                  0,
		CrlSignerCertificate: testconstants.PAACertNoVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	// update to PAA NOVID 2
	updatePkiRevocationDistributionPoint := types.MsgUpdatePkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  addPkiRevocationDistributionPoint.Vid,
		CrlSignerCertificate: testconstants.RootCertPem,
		Label:                addPkiRevocationDistributionPoint.Label,
		IssuerSubjectKeyID:   addPkiRevocationDistributionPoint.IssuerSubjectKeyID,
	}
	_, err = setup.Handler(setup.Ctx, &updatePkiRevocationDistributionPoint)
	require.NoError(t, err)
}

func TestHandler_UpdatePkiRevocationDistributionPoint_PAA_NOVID_DifferentVID(t *testing.T) {
	setup := Setup(t)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 1001)

	// add PAA NOVID 1 with vid=1001
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.PAACertNoVid, testconstants.Info, 1001)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.PAACertNoVidSubject, testconstants.PAACertNoVidSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// add PAA NOVID 2 with vid=1002
	proposeAddX509RootCert = types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, 1002)
	_, err = setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)
	approveAddX509RootCert = types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// add Revocation Point PAA NOVID 1
	addPkiRevocationDistributionPoint := types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  1001,
		IsPAA:                true,
		Pid:                  0,
		CrlSignerCertificate: testconstants.PAACertNoVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	// update to PAA NOVID 2 (with different vid)
	updatePkiRevocationDistributionPoint := types.MsgUpdatePkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  addPkiRevocationDistributionPoint.Vid,
		CrlSignerCertificate: testconstants.RootCertPem,
		Label:                addPkiRevocationDistributionPoint.Label,
		IssuerSubjectKeyID:   addPkiRevocationDistributionPoint.IssuerSubjectKeyID,
	}
	_, err = setup.Handler(setup.Ctx, &updatePkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrMessageVidNotEqualRootCertVid)
}

func TestHandler_UpdatePkiRevocationDistributionPoint_PAA_NOVID_TO_PAA_VID(t *testing.T) {
	setup := Setup(t)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.PAACertWithNumericVidVid)

	// add PAA NOVID
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.PAACertNoVid, testconstants.Info, testconstants.PAACertWithNumericVidVid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.PAACertNoVidSubject, testconstants.PAACertNoVidSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// add PAA VID
	proposeAddX509RootCert = types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.PAACertWithNumericVid, testconstants.Info, testconstants.PAACertWithNumericVidVid)
	_, err = setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)
	approveAddX509RootCert = types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.PAACertWithNumericVidSubject, testconstants.PAACertWithNumericVidSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// add Revocation Point PAA NOVID
	addPkiRevocationDistributionPoint := types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  testconstants.PAACertWithNumericVidVid,
		IsPAA:                true,
		Pid:                  0,
		CrlSignerCertificate: testconstants.PAACertNoVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	// uodpate to PAA VID with same VID
	updatePkiRevocationDistributionPoint := types.MsgUpdatePkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  addPkiRevocationDistributionPoint.Vid,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                addPkiRevocationDistributionPoint.Label,
		IssuerSubjectKeyID:   addPkiRevocationDistributionPoint.IssuerSubjectKeyID,
	}
	_, err = setup.Handler(setup.Ctx, &updatePkiRevocationDistributionPoint)
	require.NoError(t, err)
}

func TestHandler_UpdatePkiRevocationDistributionPoint_PAA_VID_TO_PAA_NOVID(t *testing.T) {
	setup := Setup(t)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.PAACertWithNumericVidVid)

	// add PAA NOVID
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.PAACertNoVid, testconstants.Info, testconstants.PAACertWithNumericVidVid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.PAACertNoVidSubject, testconstants.PAACertNoVidSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// add PAA VID
	proposeAddX509RootCert = types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.PAACertWithNumericVid, testconstants.Info, testconstants.PAACertWithNumericVidVid)
	_, err = setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)
	approveAddX509RootCert = types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.PAACertWithNumericVidSubject, testconstants.PAACertWithNumericVidSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// add Revocation Point PAA VID
	addPkiRevocationDistributionPoint := types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  testconstants.PAACertWithNumericVidVid,
		IsPAA:                true,
		Pid:                  0,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	// update to PAA NOVID with same VID
	updatePkiRevocationDistributionPoint := types.MsgUpdatePkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  addPkiRevocationDistributionPoint.Vid,
		CrlSignerCertificate: testconstants.PAACertNoVid,
		Label:                addPkiRevocationDistributionPoint.Label,
		IssuerSubjectKeyID:   addPkiRevocationDistributionPoint.IssuerSubjectKeyID,
	}
	_, err = setup.Handler(setup.Ctx, &updatePkiRevocationDistributionPoint)
	require.NoError(t, err)
}

func TestHandler_UpdatePkiRevocationDistributionPoint_PAI_VID_TO_PAI_NOVID(t *testing.T) {
	setup := Setup(t)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 65521)

	// add PAA for PAI_VID
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.PAACertWithNumericVid, testconstants.Info, testconstants.PAACertWithNumericVidVid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.PAACertWithNumericVidSubject, testconstants.PAACertWithNumericVidSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// add PAA for PAI_NOVID
	proposeAddX509RootCert = types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.PAACertWithNumericVidVid)
	_, err = setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)
	approveAddX509RootCert = types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// add Revocation Point PAI_VID
	addPkiRevocationDistributionPoint := types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  65521,
		IsPAA:                true,
		Pid:                  0,
		CrlSignerCertificate: testconstants.PAICertWithNumericPidVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	// update Revocation Point to PAI_NOVID
	updatePkiRevocationDistributionPoint := types.MsgUpdatePkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  65521,
		CrlSignerCertificate: testconstants.IntermediateCertPem,
		Label:                addPkiRevocationDistributionPoint.Label,
		IssuerSubjectKeyID:   addPkiRevocationDistributionPoint.IssuerSubjectKeyID,
	}
	_, err = setup.Handler(setup.Ctx, &updatePkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrCRLSignerCertificateVidNotEqualRevocationPointVid)
}
