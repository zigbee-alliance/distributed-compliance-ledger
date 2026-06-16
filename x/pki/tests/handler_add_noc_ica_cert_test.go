package tests

import (
	"errors"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/tests/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

type CertificateTestCase struct {
	name                    string
	crtType                 types.CertificateType
	isVidVerificationSigner bool
}

// Main

func TestHandler_AddNocIntermediateCert(t *testing.T) {
	setup := utils.Setup(t)

	rootCertificate := utils.RootNocCertificate1(setup.Vendor1, types.CertificateType_OperationalPKI)
	utils.AddNocRootCertificate(setup, rootCertificate)

	icaCertificate := utils.IntermediateNocCertificate1(setup.Vendor1, types.CertificateType_OperationalPKI)
	utils.AddNocIntermediateCertificate(setup, icaCertificate)

	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
			{Key: types.NocRootCertificatesKeyPrefix, Count: 1},
			{Key: types.NocIcaCertificatesKeyPrefix},
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate, indexes)
}

// TestHandler_AddNocIntermediateCert_VVSC exercises the
// IsVidVerificationSigner=true branch. The cert profile, fixture PEM, and
// chain shape differ from OperationalPKI; see verifyVVSCCertificate for the
// Matter R1.6 §6.4.10 step 12.a.iii chain semantics.
func TestHandler_AddNocIntermediateCert_VVSC(t *testing.T) {
	setup := utils.Setup(t)

	rootCertificate := utils.RootVvscCertificate1(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate)

	icaCertificate := utils.IntermediateVvscCertificate1(setup.Vendor1)
	utils.AddNocIntermediateCertificate(setup, icaCertificate)

	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix},
			{Key: types.NocRootCertificatesKeyPrefix, Count: 1},
			{Key: types.NocIcaCertificatesKeyPrefix},
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate, indexes)
}

func TestHandler_AddNocIntermediateCert_SameSubjectAndSkid_DifferentSerialNumber(t *testing.T) {
	setup := utils.Setup(t)

	rootCertificate := utils.RootNocCertificate1(setup.Vendor1, types.CertificateType_OperationalPKI)
	utils.AddNocRootCertificate(setup, rootCertificate)

	intermediateCertificate := utils.IntermediateNocCertificate1(setup.Vendor1, types.CertificateType_OperationalPKI)
	utils.AddMokedNocCertificate(setup, intermediateCertificate)

	intermediateCertificate2 := utils.IntermediateNocCertificate1Copy(setup.Vendor1, types.CertificateType_OperationalPKI)
	utils.AddNocIntermediateCertificate(setup, intermediateCertificate2)

	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix, Count: 2},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.NocCertificatesKeyPrefix, Count: 2},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix, Count: 2},
			{Key: types.NocRootCertificatesKeyPrefix, Count: 1},
			{Key: types.NocIcaCertificatesKeyPrefix, Count: 2},
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, intermediateCertificate, indexes)
	utils.CheckCertificateStateIndexes(t, setup, intermediateCertificate2, indexes)
}

func TestHandler_AddNocIntermediateCert_SameSubjectAndSkid_DifferentSerialNumber_VVSC(t *testing.T) {
	setup := utils.Setup(t)

	rootCertificate := utils.RootVvscCertificate1(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate)

	intermediateCertificate := utils.IntermediateVvscCertificate1(setup.Vendor1)
	utils.AddMokedNocCertificate(setup, intermediateCertificate)

	intermediateCertificate2 := utils.IntermediateVvscCertificate1Copy(setup.Vendor1)
	utils.AddNocIntermediateCertificate(setup, intermediateCertificate2)

	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix, Count: 2},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.NocCertificatesKeyPrefix, Count: 2},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix, Count: 2},
			{Key: types.NocRootCertificatesKeyPrefix, Count: 1},
			{Key: types.NocIcaCertificatesKeyPrefix, Count: 2},
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, intermediateCertificate, indexes)
	utils.CheckCertificateStateIndexes(t, setup, intermediateCertificate2, indexes)
}

func TestHandler_AddNocIntermediateCert_ByNotOwnerButSameVendor(t *testing.T) {
	setup := utils.Setup(t)

	vendorAccAddress1 := setup.CreateVendorAccount(testconstants.Vid)
	vendorAccAddress2 := setup.CreateVendorAccount(testconstants.Vid)

	rootCertificate := utils.RootNocCertificate1(vendorAccAddress1, types.CertificateType_OperationalPKI)
	utils.AddNocRootCertificate(setup, rootCertificate)

	icaCertificate := utils.IntermediateNocCertificate1(vendorAccAddress1, types.CertificateType_OperationalPKI)
	utils.AddNocIntermediateCertificate(setup, icaCertificate)

	icaCertificate2 := utils.IntermediateNocCertificate1Copy(vendorAccAddress2, types.CertificateType_OperationalPKI)
	utils.AddNocIntermediateCertificate(setup, icaCertificate2)

	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix, Count: 2},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.NocCertificatesKeyPrefix, Count: 2},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix, Count: 2},
			{Key: types.NocRootCertificatesKeyPrefix, Count: 1},
			{Key: types.NocIcaCertificatesKeyPrefix, Count: 2},
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate, indexes)
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate2, indexes)
}

func TestHandler_AddNocIntermediateCert_ByNotOwnerButSameVendor_VVSC(t *testing.T) {
	setup := utils.Setup(t)

	vendorAccAddress1 := setup.CreateVendorAccount(testconstants.Vid)
	vendorAccAddress2 := setup.CreateVendorAccount(testconstants.Vid)

	rootCertificate := utils.RootVvscCertificate1(vendorAccAddress1)
	utils.AddNocRootCertificate(setup, rootCertificate)

	icaCertificate := utils.IntermediateVvscCertificate1(vendorAccAddress1)
	utils.AddNocIntermediateCertificate(setup, icaCertificate)

	icaCertificate2 := utils.IntermediateVvscCertificate1Copy(vendorAccAddress2)
	utils.AddNocIntermediateCertificate(setup, icaCertificate2)

	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix, Count: 2},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.NocCertificatesKeyPrefix, Count: 2},
			{Key: types.NocCertificatesBySubjectKeyPrefix},
			{Key: types.NocCertificatesBySubjectKeyIDKeyPrefix, Count: 2},
			{Key: types.NocCertificatesByVidAndSkidKeyPrefix, Count: 2},
			{Key: types.NocRootCertificatesKeyPrefix, Count: 1},
			{Key: types.NocIcaCertificatesKeyPrefix, Count: 2},
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.ChildCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
	}
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate, indexes)
	utils.CheckCertificateStateIndexes(t, setup, icaCertificate2, indexes)
}

// Error cases

func TestHandler_AddNocIntermediateCert_SenderNotVendor(t *testing.T) {
	// Trustee1 has no Vendor role — auth check fires before cert parsing, so
	// the cert PEM and IsVidVerificationSigner flag don't affect the outcome
	// and both code paths share this single coverage.
	setup := utils.Setup(t)

	rootCertificate := utils.RootNocCertificate1(setup.Vendor1, types.CertificateType_OperationalPKI)
	utils.AddNocRootCertificate(setup, rootCertificate)

	addNocX509Cert := types.NewMsgAddNocX509IcaCert(setup.Trustee1.String(), testconstants.NocCert1, testconstants.CertSchemaVersion, false)
	_, err := setup.Handler(setup.Ctx, addNocX509Cert)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_AddNocIntermediateCert_Root_VID_Does_Not_Equal_To_AccountVID(t *testing.T) {
	setup := utils.Setup(t)

	rootCertificate := utils.RootNocCertificate1(setup.Vendor1, types.CertificateType_OperationalPKI)
	utils.AddNocRootCertificate(setup, rootCertificate)

	newAccAddress := setup.CreateVendorAccount(1111)

	nocX509Cert := types.NewMsgAddNocX509IcaCert(newAccAddress.String(), testconstants.NocCert1, testconstants.CertSchemaVersion, false)
	_, err := setup.Handler(setup.Ctx, nocX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrCertVidNotEqualAccountVid)
}

func TestHandler_AddNocIntermediateCert_Root_VID_Does_Not_Equal_To_AccountVID_VVSC(t *testing.T) {
	setup := utils.Setup(t)

	rootCertificate := utils.RootVvscCertificate1(setup.Vendor1)
	utils.AddNocRootCertificate(setup, rootCertificate)

	newAccAddress := setup.CreateVendorAccount(1111)

	nocX509Cert := types.NewMsgAddNocX509IcaCert(newAccAddress.String(), testconstants.VvscIcaCert1, testconstants.CertSchemaVersion, true)
	_, err := setup.Handler(setup.Ctx, nocX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrCertVidNotEqualAccountVid)
}

func TestHandler_AddNocIntermediateCert_ForInvalidCertificate(t *testing.T) {
	// StubCertPem fails at PEM decoding before any profile dispatch, so the
	// IsVidVerificationSigner flag is irrelevant — both paths hit the same
	// ErrInvalidCertificate.
	cases := []CertificateTestCase{
		{
			name:                    "OperationalPKI_AddNocIntermediateCert_ForInvalidCertificate",
			isVidVerificationSigner: false,
		},
		{
			name:                    "VIDSignerPKI_AddNocIntermediateCert_ForInvalidCertificate",
			isVidVerificationSigner: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			addX509Cert := types.NewMsgAddNocX509IcaCert(setup.Vendor1.String(), testconstants.StubCertPem, testconstants.CertSchemaVersion, tc.isVidVerificationSigner)
			_, err := setup.Handler(setup.Ctx, addX509Cert)
			require.ErrorIs(t, err, pkitypes.ErrInvalidCertificate)
		})
	}
}

func TestHandler_AddNocIntermediateCert_ForNonCACertificate(t *testing.T) {
	// LeafCertWithoutBasicConstraints has no BC extension. The two profiles
	// report this differently:
	//   - OperationalPKI (VerifyCAExtensions) treats a missing BC as "not a CA"
	//     and returns ErrInappropriateCertificateType (wrong cert TYPE).
	//   - VIDSignerPKI (VerifyVVSCExtensions → verifyEndEntityExtensions)
	//     treats a missing BC as a structural defect and returns
	//     ErrInvalidCertificate.
	cases := []struct {
		name                    string
		isVidVerificationSigner bool
		expectedErr             error
	}{
		{
			name:                    "OperationalPKI_AddNocIntermediateCert_ForNonCACertificate",
			isVidVerificationSigner: false,
			expectedErr:             pkitypes.ErrInappropriateCertificateType,
		},
		{
			name:                    "VIDSignerPKI_AddNocIntermediateCert_ForNonCACertificate",
			isVidVerificationSigner: true,
			expectedErr:             pkitypes.ErrInvalidCertificate,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			addX509Cert := types.NewMsgAddNocX509IcaCert(setup.Vendor1.String(), testconstants.LeafCertWithoutBasicConstraints, testconstants.CertSchemaVersion, tc.isVidVerificationSigner)
			_, err := setup.Handler(setup.Ctx, addX509Cert)
			require.ErrorIs(t, err, tc.expectedErr)
		})
	}
}

func TestHandler_AddNocIntermediateCert_ForNocRootCertificate(t *testing.T) {
	setup := utils.Setup(t)

	// try to add root certificate x509 certificate. NocRootCert1 is a NOC root
	// (cA=TRUE without pathLenConstraint), so the §6.2.2.4 PAI rule 9 check
	// rejects it before the IsSelfSigned check can fire. Either rejection mode
	// is acceptable as the cert is unfit for the DA chain either way.
	addX509Cert := types.NewMsgAddX509Cert(setup.Vendor1.String(), testconstants.NocRootCert1, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrInvalidCertificate)
	require.Contains(t, err.Error(), "PAI: BasicConstraints pathLenConstraint SHALL be present and set to 0")
}

func TestHandler_AddNocIntermediateCert_ForRootNonNocCertificate(t *testing.T) {
	// IntermediateCertWithVid1 is cA=TRUE; the OperationalPKI path fails the
	// "root must be NOC" check, while the VIDSignerPKI path fails the VVSC
	// profile check earlier. Same error code, different rule.
	cases := []CertificateTestCase{
		{
			name:                    "OperationalPKI_AddNocIntermediateCert_ForRootNonNocCertificate",
			isVidVerificationSigner: false,
		},
		{
			name:                    "VIDSignerPKI_AddNocIntermediateCert_ForRootNonNocCertificate",
			isVidVerificationSigner: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			// store root certificate
			rootCert := utils.RootDaCertificateWithVid(setup.Trustee1)
			utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

			// try to add root certificate x509 certificate
			addX509Cert := types.NewMsgAddNocX509IcaCert(setup.Vendor1.String(), testconstants.IntermediateCertWithVid1, testconstants.CertSchemaVersion, tc.isVidVerificationSigner)
			_, err := setup.Handler(setup.Ctx, addX509Cert)
			require.ErrorIs(t, err, pkitypes.ErrInappropriateCertificateType)
		})
	}
}

func TestHandler_AddNocIntermediateCert_WhenNocRootCertIsAbsent(t *testing.T) {
	setup := utils.Setup(t)

	addNocX509Cert := types.NewMsgAddNocX509IcaCert(setup.Vendor1.String(), testconstants.NocCert1, testconstants.CertSchemaVersion, false)
	_, err := setup.Handler(setup.Ctx, addNocX509Cert)
	require.ErrorIs(t, err, pkitypes.ErrCertificateDoesNotExist)
}

func TestHandler_AddNocIntermediateCert_WhenNocRootCertIsAbsent_VVSC(t *testing.T) {
	setup := utils.Setup(t)

	addNocX509Cert := types.NewMsgAddNocX509IcaCert(setup.Vendor1.String(), testconstants.VvscIcaCert1, testconstants.CertSchemaVersion, true)
	_, err := setup.Handler(setup.Ctx, addNocX509Cert)
	// verifyVVSCCertificate wraps NewErrRootCertificateDoesNotExist, which
	// itself wraps the base ErrCertificateDoesNotExist registered code.
	require.ErrorIs(t, err, pkitypes.ErrCertificateDoesNotExist)
}

type intermediateCertificateExistFixture struct {
	rootIssuer           string
	rootSubject          string
	rootSubjectKeyID     string
	childIssuer          string
	childSubject         string
	childSubjectAsText   string
	childSubjectKeyID    string
	childSerialNumber    string
	submitPem            string
	crtType              types.CertificateType
	isVVSC               bool
	rootCertificateAdder func(addr sdk.AccAddress) types.Certificate
}

func operationalPKIIntermediateFixture() intermediateCertificateExistFixture {
	return intermediateCertificateExistFixture{
		rootIssuer:         testconstants.NocRootCert1Subject,
		rootSubject:        testconstants.NocRootCert1Subject,
		rootSubjectKeyID:   testconstants.NocRootCert1SubjectKeyID,
		childIssuer:        testconstants.NocRootCert1Subject,
		childSubject:       testconstants.NocCert1Subject,
		childSubjectAsText: testconstants.NocCert1SubjectAsText,
		childSubjectKeyID:  testconstants.NocCert1SubjectKeyID,
		childSerialNumber:  testconstants.NocCert1SerialNumber,
		submitPem:          testconstants.NocCert1,
		crtType:            types.CertificateType_OperationalPKI,
		isVVSC:             false,
		rootCertificateAdder: func(addr sdk.AccAddress) types.Certificate {
			return utils.RootNocCertificate1(addr, types.CertificateType_OperationalPKI)
		},
	}
}

func vvscIntermediateFixture() intermediateCertificateExistFixture {
	return intermediateCertificateExistFixture{
		rootIssuer:           testconstants.VvscRootCert1Subject,
		rootSubject:          testconstants.VvscRootCert1Subject,
		rootSubjectKeyID:     testconstants.VvscRootCert1SubjectKeyID,
		childIssuer:          testconstants.VvscRootCert1Subject,
		childSubject:         testconstants.VvscIcaCert1Subject,
		childSubjectAsText:   testconstants.VvscIcaCert1SubjectAsText,
		childSubjectKeyID:    testconstants.VvscIcaCert1SubjectKeyID,
		childSerialNumber:    testconstants.VvscIcaCert1SerialNumber,
		submitPem:            testconstants.VvscIcaCert1,
		crtType:              types.CertificateType_VIDSignerPKI,
		isVVSC:               true,
		rootCertificateAdder: utils.RootVvscCertificate1,
	}
}

func TestHandler_AddNocIntermediateCert_CertificateExist(t *testing.T) {
	runIntermediateCertificateExistCases(t, operationalPKIIntermediateFixture())
}

// TestHandler_AddNocIntermediateCert_CertificateExist_VVSC reuses the
// existing-row collision matrix from the OperationalPKI variant against the
// VVSC fixtures: VvscRootCert1 + VvscIcaCert1 in place of NocRootCert1 +
// NocCert1, with the existing on-ledger row's CertificateType seeded as
// VIDSignerPKI so the type-match check inside AddNocX509IcaCert is satisfied.
func TestHandler_AddNocIntermediateCert_CertificateExist_VVSC(t *testing.T) {
	runIntermediateCertificateExistCases(t, vvscIntermediateFixture())
}

func runIntermediateCertificateExistCases(t *testing.T, fx intermediateCertificateExistFixture) {
	t.Helper()
	accAddress := utils.GenerateAccAddress()

	cases := []struct {
		name         string
		existingCert *types.Certificate
		err          error
	}{
		{
			name: "Duplicate",
			existingCert: &types.Certificate{
				Issuer:         fx.childIssuer,
				AuthorityKeyId: fx.rootSubjectKeyID,
				Subject:        fx.childSubject,
				SubjectAsText:  fx.childSubjectAsText,
				SubjectKeyId:   fx.childSubjectKeyID,
				SerialNumber:   fx.childSerialNumber,
				IsRoot:         false,
				Vid:            testconstants.Vid,
			},
			err: pkitypes.ErrCertificateAlreadyExists,
		},
		{
			name: "ExistingIsRootCert",
			existingCert: &types.Certificate{
				Issuer:         fx.childIssuer,
				AuthorityKeyId: fx.rootSubjectKeyID,
				Subject:        fx.childSubject,
				SubjectAsText:  fx.childSubjectAsText,
				SubjectKeyId:   fx.childSubjectKeyID,
				SerialNumber:   testconstants.NocRootCert1SerialNumber,
				IsRoot:         true,
				Vid:            testconstants.Vid,
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			name: "ExistingWithDifferentIssuer",
			existingCert: &types.Certificate{
				Issuer:         testconstants.RootIssuer,
				AuthorityKeyId: fx.rootSubjectKeyID,
				Subject:        fx.childSubject,
				SubjectAsText:  fx.childSubjectAsText,
				SubjectKeyId:   fx.childSubjectKeyID,
				SerialNumber:   "1234",
				IsRoot:         false,
				Vid:            testconstants.Vid,
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			name: "ExistingWithDifferentAuthorityKeyId",
			existingCert: &types.Certificate{
				Issuer:         fx.childIssuer,
				AuthorityKeyId: testconstants.RootSubjectKeyID,
				Subject:        fx.childSubject,
				SubjectAsText:  fx.childSubjectAsText,
				SubjectKeyId:   fx.childSubjectKeyID,
				SerialNumber:   "1234",
				IsRoot:         false,
				Vid:            testconstants.Vid,
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			name: "ExistingNotNocCert",
			existingCert: &types.Certificate{
				Issuer:         fx.childIssuer,
				AuthorityKeyId: fx.rootSubjectKeyID,
				Subject:        fx.childSubject,
				SubjectAsText:  fx.childSubjectAsText,
				SubjectKeyId:   fx.childSubjectKeyID,
				SerialNumber:   "1234",
				IsRoot:         false,
				Vid:            testconstants.Vid,
			},
			err: pkitypes.ErrInappropriateCertificateType,
		},
		{
			name: "ExistingCertWithDifferentVid",
			existingCert: &types.Certificate{
				Issuer:         fx.childIssuer,
				AuthorityKeyId: fx.rootSubjectKeyID,
				Subject:        fx.childSubject,
				SubjectAsText:  fx.childSubjectAsText,
				SubjectKeyId:   fx.childSubjectKeyID,
				SerialNumber:   "1234",
				IsRoot:         false,
				Vid:            testconstants.VendorID1,
			},
			err: sdkerrors.ErrUnauthorized,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)
			vid := testconstants.Vid
			setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

			rootCertificate := fx.rootCertificateAdder(accAddress)
			utils.AddNocRootCertificate(setup, rootCertificate)

			existingCert := *tc.existingCert

			if errors.Is(tc.err, pkitypes.ErrInappropriateCertificateType) {
				existingCert.CertificateType = types.CertificateType_DeviceAttestationPKI
			} else {
				existingCert.CertificateType = fx.crtType
			}

			setup.Keeper.AddAllCertificate(setup.Ctx, existingCert)
			uniqueCertificate := types.UniqueCertificate{
				Issuer:       existingCert.Issuer,
				SerialNumber: existingCert.SerialNumber,
				Present:      true,
			}
			setup.Keeper.SetUniqueCertificate(setup.Ctx, uniqueCertificate)

			addNocX509Cert := types.NewMsgAddNocX509IcaCert(accAddress.String(), fx.submitPem, testconstants.CertSchemaVersion, fx.isVVSC)
			_, err := setup.Handler(setup.Ctx, addNocX509Cert)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
