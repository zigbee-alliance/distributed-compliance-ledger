package pki

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func createAddRevocationPointMessageWithPAACertWithNumericVid(signer string) *types.MsgAddPkiRevocationDistributionPoint {
	return &types.MsgAddPkiRevocationDistributionPoint{
		Signer:               signer,
		Vid:                  testconstants.PAACertWithNumericVidVid,
		IsPAA:                true,
		Pid:                  0,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       types.CRLRevocationType,
	}
}

func createAddRevocationPointMessageWithPAICertWithNumericVidPid(signer string) *types.MsgAddPkiRevocationDistributionPoint {
	return &types.MsgAddPkiRevocationDistributionPoint{
		Signer:               signer,
		Vid:                  testconstants.PAICertWithNumericPidVid_Vid,
		IsPAA:                false,
		Pid:                  testconstants.PAICertWithNumericPidVid_Pid,
		CrlSignerCertificate: testconstants.PAICertWithNumericPidVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       types.CRLRevocationType,
	}
}

func createAddRevocationPointMessageWithPAICertWithVidPid(signer string) *types.MsgAddPkiRevocationDistributionPoint {
	return &types.MsgAddPkiRevocationDistributionPoint{
		Signer:               signer,
		Vid:                  testconstants.PAICertWithPidVid_Vid,
		IsPAA:                false,
		Pid:                  testconstants.PAICertWithPidVid_Pid,
		CrlSignerCertificate: testconstants.PAICertWithPidVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       types.CRLRevocationType,
	}
}

func createAddRevocationPointMessageWithPAACertNoVid(signer string, vid int32) *types.MsgAddPkiRevocationDistributionPoint {
	return &types.MsgAddPkiRevocationDistributionPoint{
		Signer:               signer,
		Vid:                  vid,
		IsPAA:                false,
		Pid:                  0,
		CrlSignerCertificate: testconstants.PAACertNoVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       types.CRLRevocationType,
	}
}

func assertRevocationPointEqual(t *testing.T, expected *types.MsgAddPkiRevocationDistributionPoint, actual *types.PkiRevocationDistributionPoint) {
	require.Equal(t, expected.CrlSignerCertificate, actual.CrlSignerCertificate)
	require.Equal(t, expected.CrlSignerCertificate, actual.CrlSignerCertificate)
	require.Equal(t, expected.DataDigest, actual.DataDigest)
	require.Equal(t, expected.DataDigestType, actual.DataDigestType)
	require.Equal(t, expected.DataFileSize, actual.DataFileSize)
	require.Equal(t, expected.DataURL, actual.DataURL)
	require.Equal(t, expected.IsPAA, actual.IsPAA)
	require.Equal(t, expected.IssuerSubjectKeyID, actual.IssuerSubjectKeyID)
	require.Equal(t, expected.Label, actual.Label)
	require.Equal(t, expected.Pid, actual.Pid)
	require.Equal(t, expected.RevocationType, actual.RevocationType)
	require.Equal(t, expected.Vid, actual.Vid)
}

func TestHandler_AddPkiRevocationDistributionPoint_PAASenderNotVendor(t *testing.T) {
	setup := Setup(t)

	addPkiRevocationDistributionPoint := createAddRevocationPointMessageWithPAACertWithNumericVid(setup.Trustee1.String())
	_, err := setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_AddPkiRevocationDistributionPoint_PAACertEncodesVidSenderVidNotEqualVidField(t *testing.T) {
	setup := Setup(t)

	vendorAcc := setup.AddVendorAccount(1)

	addPkiRevocationDistributionPoint := createAddRevocationPointMessageWithPAACertWithNumericVid(vendorAcc.String())
	_, err := setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrCRLSignerCertificateVidNotEqualAccountVid)
}

func TestHandler_AddPkiRevocationDistributionPoint_PAACertNotFound(t *testing.T) {
	setup := Setup(t)

	vendorAcc := setup.AddVendorAccount(testconstants.PAACertWithNumericVidVid)

	addPkiRevocationDistributionPoint := createAddRevocationPointMessageWithPAACertWithNumericVid(vendorAcc.String())
	_, err := setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrCertificateDoesNotExist)
}

func TestHandler_AddPkiRevocationDistributionPoint_InvalidCertificate(t *testing.T) {
	setup := Setup(t)

	vendorAcc := setup.AddVendorAccount(testconstants.PAACertWithNumericVidVid)

	addPkiRevocationDistributionPoint := createAddRevocationPointMessageWithPAACertWithNumericVid(vendorAcc.String())
	addPkiRevocationDistributionPoint.CrlSignerCertificate = "invalidpem"
	_, err := setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrInvalidCertificate)
}

func TestHandler_AddPkiRevocationDistributionPoint_PAANotOnLedger(t *testing.T) {
	setup := Setup(t)

	vendorAcc := setup.AddVendorAccount(testconstants.PAACertWithNumericVidVid)

	// propose and approve x509 root certificate
	rootCertOptions := createTestRootCertOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	addPkiRevocationDistributionPoint := createAddRevocationPointMessageWithPAACertWithNumericVid(vendorAcc.String())
	_, err := setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrCertificateDoesNotExist)
}

func TestHandler_AddPkiRevocationDistributionPoint_PAANDifferentPem(t *testing.T) {
	setup := Setup(t)

	vendorAcc := setup.AddVendorAccount(testconstants.PAACertWithNumericVidVid)

	// propose and approve x509 root certificate
	rootCertOptions := createPAACertWithNumericVidOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	addPkiRevocationDistributionPoint := createAddRevocationPointMessageWithPAACertWithNumericVid(vendorAcc.String())
	addPkiRevocationDistributionPoint.CrlSignerCertificate = testconstants.PAACertWithNumericVid1
	_, err := setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrCertificateDoesNotExist)
}

func TestHandler_AddPkiRevocationDistributionPoint_PAISenderNotVendor(t *testing.T) {
	setup := Setup(t)

	addPkiRevocationDistributionPoint := createAddRevocationPointMessageWithPAICertWithNumericVidPid(setup.Trustee1.String())
	_, err := setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint)

	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_AddPkiRevocationDistributionPoint_PAINotChainedBackToDCLCerts(t *testing.T) {
	setup := Setup(t)

	vendorAcc := setup.AddVendorAccount(testconstants.PAACertWithNumericVidVid)

	addPkiRevocationDistributionPoint := createAddRevocationPointMessageWithPAICertWithNumericVidPid(vendorAcc.String())
	_, err := setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrCertNotChainedBack)
}

func TestHandler_AddPkiRevocationDistributionPoint_PAAAlreadyExists(t *testing.T) {
	setup := Setup(t)

	vendorAcc := setup.AddVendorAccount(testconstants.PAACertWithNumericVidVid)

	// propose and approve x509 root certificate
	rootCertOptions := createPAACertWithNumericVidOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	addPkiRevocationDistributionPoint := createAddRevocationPointMessageWithPAACertWithNumericVid(vendorAcc.String())
	_, err := setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	_, err = setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrPkiRevocationDistributionPointAlreadyExists)
}

func TestHandler_AddPkiRevocationDistributionPoint_PAAWithVid(t *testing.T) {
	setup := Setup(t)

	vendorAcc := setup.AddVendorAccount(testconstants.PAACertWithNumericVidVid)

	// propose and approve x509 root certificate
	rootCertOptions := createPAACertWithNumericVidOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	addPkiRevocationDistributionPoint := createAddRevocationPointMessageWithPAACertWithNumericVid(vendorAcc.String())
	_, err := setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	revocationPoint, isFound := setup.Keeper.GetPkiRevocationDistributionPoint(setup.Ctx, testconstants.PAACertWithNumericVidVid, "label", testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	assertRevocationPointEqual(t, addPkiRevocationDistributionPoint, &revocationPoint)

	revocationPointBySubjectKeyID, isFound := setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	assertRevocationPointEqual(t, addPkiRevocationDistributionPoint, revocationPointBySubjectKeyID.Points[0])
}

func TestHandler_AddPkiRevocationDistributionPoint_PAIWithNumericVidPid(t *testing.T) {
	setup := Setup(t)

	vendorAcc := setup.AddVendorAccount(testconstants.PAACertWithNumericVidVid)

	// propose and approve x509 root certificate
	rootCertOptions := createPAACertWithNumericVidOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	addPkiRevocationDistributionPoint := createAddRevocationPointMessageWithPAICertWithNumericVidPid(vendorAcc.String())
	_, err := setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	revocationPoint, isFound := setup.Keeper.GetPkiRevocationDistributionPoint(setup.Ctx, testconstants.PAACertWithNumericVidVid, "label", testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	assertRevocationPointEqual(t, addPkiRevocationDistributionPoint, &revocationPoint)

	revocationPointBySubjectKeyID, isFound := setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	assertRevocationPointEqual(t, addPkiRevocationDistributionPoint, revocationPointBySubjectKeyID.Points[0])
}

func TestHandler_AddPkiRevocationDistributionPoint_PAIWithStringVidPid(t *testing.T) {
	setup := Setup(t)

	vendorAcc := setup.AddVendorAccount(testconstants.PAICertWithPidVid_Vid)

	// propose and approve x509 root certificate
	rootCertOptions := createPAACertNoVidOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	addPkiRevocationDistributionPoint := createAddRevocationPointMessageWithPAICertWithVidPid(vendorAcc.String())
	_, err := setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	revocationPoint, isFound := setup.Keeper.GetPkiRevocationDistributionPoint(setup.Ctx, testconstants.PAICertWithPidVid_Vid, "label", testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	assertRevocationPointEqual(t, addPkiRevocationDistributionPoint, &revocationPoint)

	revocationPointBySubjectKeyID, isFound := setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	assertRevocationPointEqual(t, addPkiRevocationDistributionPoint, revocationPointBySubjectKeyID.Points[0])
}

func TestHandler_AddPkiRevocationDistributionPoint_PAIWithVid(t *testing.T) {
	setup := Setup(t)

	vendorAcc := setup.AddVendorAccount(testconstants.PAICertWithPidVid_Vid)

	/// propose and approve x509 root certificate
	rootCertOptions := createPAACertNoVidOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	addPkiRevocationDistributionPoint := createAddRevocationPointMessageWithPAICertWithVidPid(vendorAcc.String())
	_, err := setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	revocationPoint, isFound := setup.Keeper.GetPkiRevocationDistributionPoint(setup.Ctx, testconstants.PAICertWithPidVid_Vid, "label", testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	assertRevocationPointEqual(t, addPkiRevocationDistributionPoint, &revocationPoint)

	revocationPointBySubjectKeyID, isFound := setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	assertRevocationPointEqual(t, addPkiRevocationDistributionPoint, revocationPointBySubjectKeyID.Points[0])
}

func TestHandler_AddPkiRevocationDistributionPoint_PAA_NoVid(t *testing.T) {
	setup := Setup(t)
	vid := int32(1001)

	vendorAcc := setup.AddVendorAccount(vid)

	// propose and approve x509 root certificate
	rootCertOptions := createPAACertNoVidOptions()
	rootCertOptions.vid = vid
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	addPkiRevocationDistributionPoint := createAddRevocationPointMessageWithPAACertNoVid(vendorAcc.String(), vid)
	_, err := setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	revocationPoint, isFound := setup.Keeper.GetPkiRevocationDistributionPoint(setup.Ctx, vid, "label", testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	assertRevocationPointEqual(t, addPkiRevocationDistributionPoint, &revocationPoint)

	revocationPointBySubjectKeyID, isFound := setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	assertRevocationPointEqual(t, addPkiRevocationDistributionPoint, revocationPointBySubjectKeyID.Points[0])
}

func TestHandler_AddPkiRevocationDistributionPoint_PAANoVid_LedgerPAANoVid(t *testing.T) {
	setup := Setup(t)
	vid := int32(1001)

	vendorAcc := setup.AddVendorAccount(vid)

	// propose and approve x509 root certificate
	rootCertOptions := createPAACertNoVidOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	addPkiRevocationDistributionPoint := createAddRevocationPointMessageWithPAACertNoVid(vendorAcc.String(), vid)
	_, err := setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrMessageVidNotEqualRootCertVid)
}

func TestHandler_AddPkiRevocationDistributionPoint_PAANoVid_WrongVID(t *testing.T) {
	setup := Setup(t)
	vid := int32(1001)

	vendorAcc := setup.AddVendorAccount(vid)

	// propose and approve x509 root certificate
	rootCertOptions := createPAACertNoVidOptions()
	rootCertOptions.vid = 1002
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	addPkiRevocationDistributionPoint := createAddRevocationPointMessageWithPAACertNoVid(vendorAcc.String(), vid)
	_, err := setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrMessageVidNotEqualRootCertVid)
}
