package pki

import (
	"context"
	"math"
	"math/rand"
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	testkeeper "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const SerialNumber = "12345678"

type DclauthKeeperMock struct {
	mock.Mock
}

func (m *DclauthKeeperMock) HasRole(
	ctx sdk.Context,
	addr sdk.AccAddress,
	roleToCheck dclauthtypes.AccountRole,
) bool {
	args := m.Called(ctx, addr, roleToCheck)

	return args.Bool(0)
}

func (m *DclauthKeeperMock) CountAccountsWithRole(ctx sdk.Context, roleToCount dclauthtypes.AccountRole) int {
	args := m.Called(ctx, roleToCount)

	return args.Int(0)
}

func (m *DclauthKeeperMock) GetAccountO(
	ctx sdk.Context,
	address sdk.AccAddress,
) (val dclauthtypes.Account, found bool) {
	args := m.Called(ctx, address)

	return args.Get(0).(dclauthtypes.Account), args.Bool(1)
}

var _ types.DclauthKeeper = &DclauthKeeperMock{}

type TestSetup struct {
	T *testing.T
	// Cdc         *amino.Codec
	Ctx           sdk.Context
	Wctx          context.Context
	Keeper        *keeper.Keeper
	DclauthKeeper *DclauthKeeperMock
	Handler       sdk.Handler
	// Querier     sdk.Querier
	Trustee1 sdk.AccAddress
	Trustee2 sdk.AccAddress
	Trustee3 sdk.AccAddress
}

// Remove a item from ExpectedCalls Array and return it.
func removeItemFromExpectedCalls(expectedCalls []*mock.Call, methodName string) {
	for i, call := range expectedCalls {
		if call.Method == methodName {
			expectedCalls = append(expectedCalls[:i], expectedCalls[i+1:]...)
		}
	}
}

func (setup *TestSetup) AddAccount(
	accAddress sdk.AccAddress,
	roles []dclauthtypes.AccountRole,
	vid int32,
) {
	dclauthKeeper := setup.DclauthKeeper
	currentTrusteeCount := 0
	// if the CountAccountsWithRole is present get the value from the mock call
	for _, expectedCall := range dclauthKeeper.ExpectedCalls {
		if expectedCall.Method == "CountAccountsWithRole" {
			currentTrusteeCount = dclauthKeeper.CountAccountsWithRole(setup.Ctx, dclauthtypes.Trustee)
		}
	}

	for _, role := range roles {
		dclauthKeeper.On("HasRole", mock.Anything, accAddress, role).Return(true)
		if role == dclauthtypes.Trustee {
			currentTrusteeCount++
			// We remove the call to CountAccountsWithRole from the expected calls and add it back with the new value
			removeItemFromExpectedCalls(dclauthKeeper.ExpectedCalls, "CountAccountsWithRole")
			dclauthKeeper.On("CountAccountsWithRole", setup.Ctx, dclauthtypes.Trustee).Return(currentTrusteeCount)
		}
	}

	dclauthKeeper.On("GetAccountO", setup.Ctx, accAddress).Return(dclauthtypes.Account{VendorID: vid}, true)
	dclauthKeeper.On("HasRole", mock.Anything, accAddress, mock.Anything).Return(false)
}

func GenerateAccAddress() sdk.AccAddress {
	_, _, accAddress := testdata.KeyTestPubAddr()

	return accAddress
}

func Setup(t *testing.T) *TestSetup {
	t.Helper()
	dclauthKeeper := &DclauthKeeperMock{}
	keeper, ctx := testkeeper.PkiKeeper(t, dclauthKeeper)

	setup := &TestSetup{
		T:             t,
		Ctx:           ctx,
		Wctx:          sdk.WrapSDKContext(ctx),
		Keeper:        keeper,
		DclauthKeeper: dclauthKeeper,
		Handler:       NewHandler(*keeper),
		Trustee1:      GenerateAccAddress(),
		Trustee2:      GenerateAccAddress(),
		Trustee3:      GenerateAccAddress(),
	}

	setup.AddAccount(setup.Trustee1, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 65521)
	setup.AddAccount(setup.Trustee2, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)
	setup.AddAccount(setup.Trustee3, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 2)

	return setup
}

func TestHandler_ProposeAddX509RootCert_ByNotTrustee(t *testing.T) {
	setup := Setup(t)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.Vendor,
		dclauthtypes.CertificationCenter,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, 1)

		// propose x509 root certificate
		proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(accAddress.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
		_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
		require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
	}
}

func TestHandler_ProposeAddAndRejectX509RootCert_ByTrustee(t *testing.T) {
	setup := Setup(t)

	// propose x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// reject x509 root certificate
	rejectX509RootCert := types.NewMsgRejectAddX509RootCert(setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectX509RootCert)
	require.NoError(t, err)

	require.False(t, setup.Keeper.IsProposedCertificatePresent(setup.Ctx, testconstants.RootIssuer, testconstants.RootSerialNumber))

	// check that unique certificate key is registered
	require.False(t, setup.Keeper.IsUniqueCertificatePresent(
		setup.Ctx, testconstants.RootIssuer, testconstants.RootSerialNumber))
}

func TestHandler_ProposeAddAndRejectX509RootCert_ByAnotherTrustee(t *testing.T) {
	setup := Setup(t)

	// propose x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// reject x509 root certificate
	rejectX509RootCert := types.NewMsgRejectAddX509RootCert(setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectX509RootCert)
	require.NoError(t, err)

	// query proposed certificate
	proposedCertificate, _ := queryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)

	// check proposed certificate
	require.Equal(t, proposeAddX509RootCert.Cert, proposedCertificate.PemCert)
	require.Equal(t, proposeAddX509RootCert.Signer, proposedCertificate.Owner)
	require.Equal(t, testconstants.RootSubject, proposedCertificate.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, proposedCertificate.SubjectKeyId)
	require.Equal(t, testconstants.RootSerialNumber, proposedCertificate.SerialNumber)
	require.True(t, proposedCertificate.HasApprovalFrom(setup.Trustee1.String()))

	// check that unique certificate key is registered
	require.True(t, setup.Keeper.IsUniqueCertificatePresent(
		setup.Ctx, testconstants.RootIssuer, testconstants.RootSerialNumber))
}

func TestHandler_ProposeAddAndRejectX509RootCertWithApproval_ByTrustee(t *testing.T) {
	setup := Setup(t)

	accAddress := GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)

	// propose x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// reject x509 root certificate
	rejectX509RootCert := types.NewMsgRejectAddX509RootCert(setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectX509RootCert)
	require.NoError(t, err)

	// query proposed certificate
	proposedCertificate, _ := queryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)

	// check proposed certificate
	require.Equal(t, proposeAddX509RootCert.Cert, proposedCertificate.PemCert)
	require.Equal(t, proposeAddX509RootCert.Signer, proposedCertificate.Owner)
	require.Equal(t, testconstants.RootSubject, proposedCertificate.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, proposedCertificate.SubjectKeyId)
	require.Equal(t, testconstants.RootSerialNumber, proposedCertificate.SerialNumber)
	require.True(t, proposedCertificate.HasRejectFrom(setup.Trustee1.String()))
	require.True(t, proposedCertificate.HasApprovalFrom(setup.Trustee2.String()))

	// check that unique certificate key is registered
	require.True(t, setup.Keeper.IsUniqueCertificatePresent(
		setup.Ctx, testconstants.RootIssuer, testconstants.RootSerialNumber))
}

func TestHandler_ProposeAddX509RootCert_ByTrustee(t *testing.T) {
	setup := Setup(t)

	// propose x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// query proposed certificate
	proposedCertificate, _ := queryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)

	// check proposed certificate
	require.Equal(t, proposeAddX509RootCert.Cert, proposedCertificate.PemCert)
	require.Equal(t, proposeAddX509RootCert.Signer, proposedCertificate.Owner)
	require.Equal(t, testconstants.RootSubject, proposedCertificate.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, proposedCertificate.SubjectKeyId)
	require.Equal(t, testconstants.RootSerialNumber, proposedCertificate.SerialNumber)
	require.True(t, proposedCertificate.HasApprovalFrom(proposeAddX509RootCert.Signer))

	// check that unique certificate key is registered
	require.True(t, setup.Keeper.IsUniqueCertificatePresent(
		setup.Ctx, testconstants.RootIssuer, testconstants.RootSerialNumber))

	// query approved certificate
	_, err = querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
}

func TestHandler_ProposeAddX509RootCert_ForInvalidCertificate(t *testing.T) {
	setup := Setup(t)

	// propose x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.StubCertPem, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrInvalidCertificate.Is(err))
}

func TestHandler_ProposeAddX509RootCert_ForNonRootCertificate(t *testing.T) {
	setup := Setup(t)

	// propose x509 leaf certificate as root
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.LeafCertPem, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrInappropriateCertificateType.Is(err))
}

func TestHandler_ProposeAddX509RootCert_ProposedCertificateAlreadyExists(t *testing.T) {
	setup := Setup(t)

	// propose adding of x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// store another account
	anotherAccount := GenerateAccAddress()
	setup.AddAccount(anotherAccount, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)

	// propose adding of the same x509 root certificate again
	proposeAddX509RootCert = types.NewMsgProposeAddX509RootCert(anotherAccount.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
	_, err = setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrProposedCertificateAlreadyExists.Is(err))
}

func TestHandler_ProposeAddX509RootCert_CertificateAlreadyExists(t *testing.T) {
	setup := Setup(t)

	// store x509 root certificate
	rootCertificate := rootCertificate(testconstants.Address1)
	setup.Keeper.SetUniqueCertificate(
		setup.Ctx,
		uniqueCertificate(rootCertificate.Subject, rootCertificate.SerialNumber),
	)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// propose adding of the same x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateAlreadyExists.Is(err))
}

func TestHandler_ProposeAddX509RootCert_ForDifferentSerialNumber(t *testing.T) {
	setup := Setup(t)

	// store root certificate with different serial number
	rootCertificate := rootCertificate(setup.Trustee1)
	rootCertificate.SerialNumber = SerialNumber
	setup.Keeper.SetUniqueCertificate(
		setup.Ctx,
		uniqueCertificate(rootCertificate.Subject, rootCertificate.SerialNumber),
	)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// propose second root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// check
	certificate, _ := querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.True(t, certificate.IsRoot)
	require.Equal(t, testconstants.RootIssuer, certificate.Subject)
	require.Equal(t, SerialNumber, certificate.SerialNumber)

	proposedCertificate, _ := queryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, testconstants.RootIssuer, proposedCertificate.Subject)
	require.Equal(t, testconstants.RootSerialNumber, proposedCertificate.SerialNumber)

	require.NotEqual(t, certificate.SerialNumber, proposedCertificate.SerialNumber)
}

func TestHandler_ProposeAddX509RootCert_ForDifferentSerialNumberDifferentSigner(t *testing.T) {
	setup := Setup(t)

	// store root certificate with different serial number
	rootCertificate := rootCertificate(testconstants.Address1)
	rootCertificate.SerialNumber = SerialNumber
	setup.Keeper.SetUniqueCertificate(
		setup.Ctx,
		uniqueCertificate(rootCertificate.Subject, rootCertificate.SerialNumber),
	)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// propose second root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_ApproveAddX509RootCert_ForNotEnoughApprovals(t *testing.T) {
	setup := Setup(t)

	// store account without trustee role
	nonTrustee := GenerateAccAddress()
	setup.AddAccount(nonTrustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)

	// propose x509 root certificate by account without trustee role
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(nonTrustee.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// query certificate
	proposedCertificate, _ := queryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, proposeAddX509RootCert.Cert, proposedCertificate.PemCert)
	require.True(t, proposedCertificate.HasApprovalFrom(setup.Trustee1.String()))

	// query approved certificate
	_, err = querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// approve again from secondTrustee (That makes is 2 trustee's from a total of 3)
	approveAddX509RootCert = types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// query approved certificate and we should get one back
	approvedCertificate, _ := querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, testconstants.RootIssuer, approvedCertificate.Subject)
	require.Equal(t, testconstants.RootSerialNumber, approvedCertificate.SerialNumber)
	require.True(t, approvedCertificate.IsRoot)
	require.True(t, approvedCertificate.HasApprovalFrom(setup.Trustee1.String()))
	require.True(t, approvedCertificate.HasApprovalFrom(setup.Trustee2.String()))
}

func TestHandler_TwoThirdApprovalsNeededForAddingRootCertification(t *testing.T) {
	setup := Setup(t)

	// propose x509 root certificate by account without trustee role
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// Create an array of trustee account from 1 to 50
	trusteeAccounts := make([]sdk.AccAddress, 50)
	for i := 0; i < 50; i++ {
		trusteeAccounts[i] = GenerateAccAddress()
	}

	totalAdditionalTrustees := rand.Intn(50)
	for i := 0; i < totalAdditionalTrustees; i++ {
		setup.AddAccount(trusteeAccounts[i], []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)
	}

	// We have 3 Trustees in test setup.
	twoThirds := int(math.Ceil(types.RootCertificateApprovalsPercent * float64(3+totalAdditionalTrustees)))

	// Until we hit 2/3 of the total number of Trustees, we should not be able to approve the certificate
	for i := 1; i < twoThirds-1; i++ {
		approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
			trusteeAccounts[i].String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
		_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
		require.NoError(t, err)

		_, err = querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
		require.Error(t, err)
		require.Equal(t, codes.NotFound, status.Code(err))
	}

	// One more approval will move this to approved state from pending
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// query approved certificate and we should get one back
	approvedCertificate, _ := querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, testconstants.RootIssuer, approvedCertificate.Subject)
	require.Equal(t, testconstants.RootSerialNumber, approvedCertificate.SerialNumber)
	require.True(t, approvedCertificate.IsRoot)
	// Check all approvals are present
	for i := 1; i < twoThirds-1; i++ {
		require.Equal(t, approvedCertificate.HasApprovalFrom(trusteeAccounts[i].String()), true)
	}
	require.Equal(t, approvedCertificate.HasApprovalFrom(setup.Trustee1.String()), true)
	require.Equal(t, approvedCertificate.HasApprovalFrom(setup.Trustee2.String()), true)
}

func TestHandler_TwoThirdApprovalsNeededForRevokingRootCertification(t *testing.T) {
	setup := Setup(t)

	// propose x509 root certificate by account without trustee role
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// Approve the certificate from Trustee2
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// Check that the certificate is approved
	approvedCertificate, _ := querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, testconstants.RootIssuer, approvedCertificate.Subject)
	require.Equal(t, testconstants.RootSerialNumber, approvedCertificate.SerialNumber)
	require.True(t, approvedCertificate.IsRoot)
	require.True(t, approvedCertificate.HasApprovalFrom(setup.Trustee1.String()))

	// Create an array of trustee account from 1 to 50
	trusteeAccounts := make([]sdk.AccAddress, 50)
	for i := 0; i < 50; i++ {
		trusteeAccounts[i] = GenerateAccAddress()
	}

	totalAdditionalTrustees := rand.Intn(50)
	for i := 0; i < totalAdditionalTrustees; i++ {
		setup.AddAccount(trusteeAccounts[i], []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)
	}

	// We have 3 Trustees in test setup.
	twoThirds := int(math.Ceil(types.RootCertificateApprovalsPercent * float64(3+totalAdditionalTrustees)))

	// Trustee1 proposes to revoke the certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	// Until we hit 2/3 of the total number of Trustees, we should not be able to revoke the certificate
	// We start the counter from 2 as the proposer is a trustee as well
	for i := 1; i < twoThirds-1; i++ {
		// approve the revocation
		approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
			trusteeAccounts[i].String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
		_, err = setup.Handler(setup.Ctx, approveRevokeX509RootCert)
		require.NoError(t, err)

		// check that the certificate is still not revoked
		approvedCertificate, _ := querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
		require.Equal(t, testconstants.RootIssuer, approvedCertificate.Subject)
		require.Equal(t, testconstants.RootSerialNumber, approvedCertificate.SerialNumber)
		require.True(t, approvedCertificate.IsRoot)
	}

	// One more revoke will revoke the certificate
	approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
		setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveRevokeX509RootCert)
	require.NoError(t, err)

	// Check that the certificate is revoked
	_, err = querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// Check that the certificate is revoked
	revokedCertificate, err := querySingleRevokedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, testconstants.RootIssuer, revokedCertificate.Subject)
	require.Equal(t, testconstants.RootSerialNumber, revokedCertificate.SerialNumber)
	require.True(t, revokedCertificate.IsRoot)
	// Make sure all the approvals are present
	for i := 1; i < twoThirds-1; i++ {
		require.Equal(t, revokedCertificate.HasApprovalFrom(trusteeAccounts[i].String()), true)
	}
	require.Equal(t, revokedCertificate.HasApprovalFrom(setup.Trustee1.String()), true)
	require.Equal(t, revokedCertificate.HasApprovalFrom(setup.Trustee2.String()), true)
}

func TestHandler_ApproveAddX509RootCert_ForEnoughApprovals(t *testing.T) {
	setup := Setup(t)

	// propose add x509 root certificate by trustee
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// approve by second trustee
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// query proposed certificate
	_, err = queryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// query approved certificate
	approvedCertificate, _ := querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, proposeAddX509RootCert.Cert, approvedCertificate.PemCert)
	require.Equal(t, proposeAddX509RootCert.Signer, approvedCertificate.Owner)
	require.Equal(t, testconstants.RootSubject, approvedCertificate.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, approvedCertificate.SubjectKeyId)
	require.Equal(t, testconstants.RootSerialNumber, approvedCertificate.SerialNumber)
	require.True(t, approvedCertificate.IsRoot)
	require.Empty(t, approvedCertificate.RootSubject)
	require.Empty(t, approvedCertificate.RootSubjectKeyId)
	require.Empty(t, approvedCertificate.Issuer)
	require.Empty(t, approvedCertificate.AuthorityKeyId)

	// check that unique certificate key is registered
	require.True(t, setup.Keeper.IsUniqueCertificatePresent(
		setup.Ctx, testconstants.RootIssuer, testconstants.RootSerialNumber))
}

func TestHandler_ApproveAddX509RootCert_ForUnknownProposedCertificate(t *testing.T) {
	setup := Setup(t)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrProposedCertificateDoesNotExist.Is(err))
}

func TestHandler_ApproveAddX509RootCert_ByNotTrustee(t *testing.T) {
	setup := Setup(t)

	// propose add x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.Vendor,
		dclauthtypes.CertificationCenter,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, 1)

		// approve
		approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
			accAddress.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
		_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
	}
}

func TestHandler_ApproveAddX509RootCert_Twice(t *testing.T) {
	setup := Setup(t)

	// store account without Trustee role
	accAddress := GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)

	// propose add x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(accAddress.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// approve second time
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_AddX509Cert(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee1)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.Vendor,
		dclauthtypes.CertificationCenter,
		dclauthtypes.Trustee,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, 1)

		// add x509 certificate
		addX509Cert := types.NewMsgAddX509Cert(accAddress.String(), testconstants.IntermediateCertPem)
		_, err := setup.Handler(setup.Ctx, addX509Cert)
		require.NoError(t, err)

		// query certificate
		certificate, _ := querySingleApprovedCertificate(
			setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)

		// check
		require.Equal(t, addX509Cert.Cert, certificate.PemCert)
		require.Equal(t, addX509Cert.Signer, certificate.Owner)
		require.Equal(t, testconstants.IntermediateSubject, certificate.Subject)
		require.Equal(t, testconstants.IntermediateSubjectKeyID, certificate.SubjectKeyId)
		require.Equal(t, testconstants.IntermediateSerialNumber, certificate.SerialNumber)
		require.False(t, certificate.IsRoot)
		require.Equal(t, testconstants.IntermediateIssuer, certificate.Issuer)
		require.Equal(t, testconstants.IntermediateAuthorityKeyID, certificate.AuthorityKeyId)
		require.Equal(t, testconstants.RootSubject, certificate.RootSubject)
		require.Equal(t, testconstants.RootSubjectKeyID, certificate.RootSubjectKeyId)

		// check that unique certificate key is registered
		require.True(t, setup.Keeper.IsUniqueCertificatePresent(
			setup.Ctx, testconstants.IntermediateIssuer, testconstants.IntermediateSerialNumber))

		// check that child certificates of issuer contains certificate identifier
		issuerChildren, _ := queryChildCertificates(
			setup, testconstants.IntermediateIssuer, testconstants.IntermediateAuthorityKeyID)
		require.Equal(t, 1, len(issuerChildren.CertIds))
		require.Equal(t,
			&types.CertificateIdentifier{
				Subject:      testconstants.IntermediateSubject,
				SubjectKeyId: testconstants.IntermediateSubjectKeyID,
			},
			issuerChildren.CertIds[0])

		// check that no proposed certificate has been created
		_, err = queryProposedCertificate(setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
		require.Error(t, err)
		require.Equal(t, codes.NotFound, status.Code(err))

		// cleanup for next iteration
		setup.Keeper.RemoveApprovedCertificates(setup.Ctx,
			testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
		setup.Keeper.RemoveUniqueCertificate(setup.Ctx,
			testconstants.IntermediateIssuer, testconstants.IntermediateSerialNumber)
		setup.Keeper.RemoveChildCertificates(setup.Ctx,
			testconstants.IntermediateIssuer, testconstants.IntermediateAuthorityKeyID)
	}
}

func TestHandler_AddX509Cert_ForInvalidCertificate(t *testing.T) {
	setup := Setup(t)

	// add x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.StubCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrInvalidCertificate.Is(err))
}

func TestHandler_AddX509Cert_ForRootCertificate(t *testing.T) {
	setup := Setup(t)

	// add root certificate as leaf x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.RootCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrInappropriateCertificateType.Is(err))
}

func TestHandler_AddX509Cert_ForDuplicate(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee1)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// store intermediate certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// store intermediate certificate second time
	_, err = setup.Handler(setup.Ctx, addX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateAlreadyExists.Is(err))
}

func TestHandler_AddX509Cert_ForDifferentSerialNumber(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee1)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// store intermediate certificate with different serial number
	intermediateCertificate := intermediateCertificate(setup.Trustee1)
	intermediateCertificate.SerialNumber = SerialNumber
	setup.Keeper.SetUniqueCertificate(
		setup.Ctx,
		uniqueCertificate(intermediateCertificate.Issuer, intermediateCertificate.SerialNumber),
	)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, intermediateCertificate)

	// store intermediate certificate second time
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// query certificate
	certificates, _ := queryApprovedCertificates(setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)

	// check
	require.Equal(t, 2, len(certificates.Certs))
	require.NotEqual(t, certificates.Certs[0].SerialNumber, certificates.Certs[1].SerialNumber)

	for _, certificate := range certificates.Certs {
		require.Equal(t, addX509Cert.Cert, certificate.PemCert)
		require.Equal(t, addX509Cert.Signer, certificate.Owner)
		require.Equal(t, testconstants.IntermediateSubject, certificate.Subject)
		require.Equal(t, testconstants.IntermediateSubjectKeyID, certificate.SubjectKeyId)
		require.False(t, certificate.IsRoot)
		require.Equal(t, testconstants.RootSubject, certificate.RootSubject)
		require.Equal(t, testconstants.RootSubjectKeyID, certificate.RootSubjectKeyId)
		require.Equal(t, testconstants.IntermediateIssuer, certificate.Issuer)
		require.Equal(t, testconstants.IntermediateAuthorityKeyID, certificate.AuthorityKeyId)
	}
}

func TestHandler_AddX509Cert_ForDifferentSerialNumberDifferentSigner(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCertificate := rootCertificate(testconstants.Address1)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// store intermediate certificate with different serial number
	intermediateCertificate := intermediateCertificate(testconstants.Address1)
	intermediateCertificate.SerialNumber = SerialNumber
	setup.Keeper.SetUniqueCertificate(
		setup.Ctx,
		uniqueCertificate(intermediateCertificate.Issuer, intermediateCertificate.SerialNumber),
	)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, intermediateCertificate)

	// store intermediate certificate second time
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_AddX509Cert_ForAbsentDirectParentCert(t *testing.T) {
	setup := Setup(t)

	// add intermediate x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrInvalidCertificate.Is(err))
}

func TestHandler_AddX509Cert_ForNoRootCert(t *testing.T) {
	setup := Setup(t)

	// add intermediate certificate
	intermediateCertificate := intermediateCertificate(setup.Trustee1)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, intermediateCertificate)

	// add leaf x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.LeafCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrInvalidCertificate.Is(err))
}

func TestHandler_AddX509Cert_ForFailedCertificateVerification(t *testing.T) {
	setup := Setup(t)

	// add invalid root
	invalidRootCertificate := types.NewRootCertificate(testconstants.StubCertPem,
		testconstants.RootSubject, testconstants.RootSubjectAsText, testconstants.RootSubjectKeyID,
		testconstants.RootSerialNumber, setup.Trustee1.String(), []*types.Grant{}, []*types.Grant{}, testconstants.Vid)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, invalidRootCertificate)

	// add intermediate x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrInvalidCertificate.Is(err))
}

func TestHandler_AddX509Cert_ForTree(t *testing.T) {
	setup := Setup(t)

	// add root x509 certificate
	proposeAndApproveRootCertificate(setup, setup.Trustee1)

	// add intermediate x509 certificate
	addIntermediateX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addIntermediateX509Cert)
	require.NoError(t, err)

	// add leaf x509 certificate
	addLeafX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.LeafCertPem)
	_, err = setup.Handler(setup.Ctx, addLeafX509Cert)
	require.NoError(t, err)

	// query root certificate
	rootCertificate, _ := querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, testconstants.RootCertPem, rootCertificate.PemCert)

	// check child certificate identifiers of root certificate
	rootCertChildren, _ := queryChildCertificates(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)

	require.Equal(t, 1, len(rootCertChildren.CertIds))
	require.Equal(t,
		certificateIdentifier(testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID),
		*rootCertChildren.CertIds[0])

	// query intermediate certificate
	intermediateCertificate, _ := querySingleApprovedCertificate(setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Equal(t, testconstants.IntermediateCertPem, intermediateCertificate.PemCert)

	// check child certificate identifiers of intermediate certificate
	intermediateCertChildren, _ := queryChildCertificates(
		setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)

	require.Equal(t, 1, len(intermediateCertChildren.CertIds))
	require.Equal(t,
		certificateIdentifier(testconstants.LeafSubject, testconstants.LeafSubjectKeyID),
		*intermediateCertChildren.CertIds[0])

	// query leaf certificate
	leafCertificate, _ := querySingleApprovedCertificate(setup, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Equal(t, testconstants.LeafCertPem, leafCertificate.PemCert)

	// check child certificate identifiers of leaf certificate
	leafCertChildren, err := queryChildCertificates(setup, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
	require.Nil(t, leafCertChildren)
}

//nolint:funlen
func TestHandler_AddX509Cert_EachChildCertRefersToTwoParentCerts(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCert := rootCertificate(setup.Trustee1)

	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCert)
	setup.Keeper.SetUniqueCertificate(setup.Ctx, uniqueCertificate(rootCert.Subject, rootCert.SerialNumber))

	// store second root certificate
	rootCert = rootCertificate(setup.Trustee1)
	rootCert.SerialNumber = SerialNumber

	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCert)
	setup.Keeper.SetUniqueCertificate(setup.Ctx, uniqueCertificate(rootCert.Subject, rootCert.SerialNumber))

	// store intermediate certificate (it refers to two parent certificates)
	intermediateCertificate := intermediateCertificate(setup.Trustee1)
	intermediateCertificate.SerialNumber = SerialNumber

	setup.Keeper.AddApprovedCertificate(setup.Ctx, intermediateCertificate)
	setup.Keeper.SetUniqueCertificate(
		setup.Ctx,
		uniqueCertificate(intermediateCertificate.Issuer, intermediateCertificate.SerialNumber),
	)

	childCertID := certificateIdentifier(intermediateCertificate.Subject, intermediateCertificate.SubjectKeyId)
	rootChildCertificates := types.ChildCertificates{
		Issuer:         intermediateCertificate.Issuer,
		AuthorityKeyId: intermediateCertificate.AuthorityKeyId,
		CertIds:        []*types.CertificateIdentifier{&childCertID},
	}
	setup.Keeper.SetChildCertificates(setup.Ctx, rootChildCertificates)

	// store second intermediate certificate (it refers to two parent certificates)
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// store leaf certificate (it refers to two parent certificates)
	addX509Cert = types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.LeafCertPem)
	_, err = setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// query root certificate
	rootCertificates, _ := queryApprovedCertificates(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, 2, len(rootCertificates.Certs))

	// check child certificate identifiers of root certificate
	rootCertChildren, _ := queryChildCertificates(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)

	require.Equal(t, 1, len(rootCertChildren.CertIds))
	require.Equal(t,
		certificateIdentifier(testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID),
		*rootCertChildren.CertIds[0])

	// query intermediate certificate
	intermediateCertificates, _ := queryApprovedCertificates(
		setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Equal(t, 2, len(intermediateCertificates.Certs))

	// check child certificate identifiers of intermediate certificate
	intermediateCertChildren, _ := queryChildCertificates(
		setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)

	require.Equal(t, 1, len(intermediateCertChildren.CertIds))
	require.Equal(t,
		certificateIdentifier(testconstants.LeafSubject, testconstants.LeafSubjectKeyID),
		*intermediateCertChildren.CertIds[0])

	// query leaf certificate
	leafCertificates, _ := queryApprovedCertificates(setup, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Equal(t, 1, len(leafCertificates.Certs))

	// check child certificate identifiers of intermediate certificate
	leafCertChildren, err := queryChildCertificates(setup, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
	require.Nil(t, leafCertChildren)
}

func TestHandler_ProposeRevokeX509RootCert_ByTrusteeOwner(t *testing.T) {
	setup := Setup(t)

	// propose x509 root certificate by `setup.Trustee` and approve by another trustee
	proposeAndApproveRootCertificate(setup, setup.Trustee1)

	// propose revocation of x509 root certificate by `setup.Trustee`
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	// query and check proposed certificate revocation
	proposedRevocation, _ := queryProposedCertificateRevocation(setup)
	require.Equal(t, testconstants.RootSubject, proposedRevocation.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, proposedRevocation.SubjectKeyId)
	require.True(t, proposedRevocation.HasRevocationFrom(setup.Trustee1.String()))

	// check that approved certificate still exists
	certificate, _ := querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NotNil(t, certificate)

	// check that revoked certificate does not exist
	_, err = queryRevokedCertificates(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificate key stays registered
	require.True(t,
		setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.RootIssuer, testconstants.RootSerialNumber))
}

func TestHandler_ProposeRevokeX509RootCert_ByTrusteeNotOwner(t *testing.T) {
	setup := Setup(t)

	// propose x509 root certificate by `setup.Trustee` and approve by another trustee
	proposeAndApproveRootCertificate(setup, setup.Trustee1)

	// store another trustee
	anotherTrustee := GenerateAccAddress()
	setup.AddAccount(anotherTrustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)

	// propose revocation of x509 root certificate by new trustee
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		anotherTrustee.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	// query and check proposed certificate revocation
	proposedRevocation, _ := queryProposedCertificateRevocation(setup)
	require.Equal(t, testconstants.RootSubject, proposedRevocation.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, proposedRevocation.SubjectKeyId)
	require.True(t, proposedRevocation.HasRevocationFrom(anotherTrustee.String()))

	// check that approved certificate still exists
	certificate, _ := querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NotNil(t, certificate)

	// check that revoked certificate does not exist
	_, err = queryRevokedCertificates(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificate key stays registered
	require.True(t,
		setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.RootIssuer, testconstants.RootSerialNumber))
}

func TestHandler_ProposeRevokeX509RootCert_ByNotTrustee(t *testing.T) {
	setup := Setup(t)

	// propose and approve x509 root certificate
	proposeAndApproveRootCertificate(setup, setup.Trustee1)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.Vendor,
		dclauthtypes.CertificationCenter,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, 1)

		// propose revocation of x509 root certificate
		proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
			accAddress.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
		_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
	}
}

func TestHandler_ProposeRevokeX509RootCert_CertificateDoesNotExist(t *testing.T) {
	setup := Setup(t)

	// propose revocation of not existing certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_ProposeRevokeX509RootCert_ForProposedCertificate(t *testing.T) {
	setup := Setup(t)

	// propose x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// check that proposed certificate is present
	proposedCertificate, _ := queryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NotNil(t, proposedCertificate)

	// propose revocation of proposed root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_ProposeRevokeX509RootCert_ProposedRevocationAlreadyExists(t *testing.T) {
	setup := Setup(t)

	// propose and approve x509 root certificate
	proposeAndApproveRootCertificate(setup, setup.Trustee1)

	// propose revocation of x509 root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	// store another trustee
	anotherTrustee := GenerateAccAddress()
	setup.AddAccount(anotherTrustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)

	// propose revocation of the same x509 root certificate again
	proposeRevokeX509RootCert = types.NewMsgProposeRevokeX509RootCert(
		anotherTrustee.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrProposedCertificateRevocationAlreadyExists.Is(err))
}

func TestHandler_ProposeRevokeX509RootCert_ForNonRootCertificate(t *testing.T) {
	setup := Setup(t)

	// store x509 root certificate
	rootCertificate := rootCertificate(setup.Trustee1)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// store x509 intermediate certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// propose revocation of x509 intermediate certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(), testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrInappropriateCertificateType.Is(err))
}

func TestHandler_ApproveRevokeX509RootCert_ForNotEnoughApprovals(t *testing.T) {
	setup := Setup(t)

	// propose and approve x509 root certificate
	proposeAndApproveRootCertificate(setup, setup.Trustee1)

	// Add 1 more trustee (this will bring the total trustee's to 4)
	anotherTrustee := GenerateAccAddress()
	setup.AddAccount(anotherTrustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)

	// propose revocation of x509 root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	// approve
	approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
		setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveRevokeX509RootCert)
	require.NoError(t, err)

	// query and check proposed certificate revocation
	proposedRevocation, _ := queryProposedCertificateRevocation(setup)
	require.Equal(t, testconstants.RootSubject, proposedRevocation.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, proposedRevocation.SubjectKeyId)
	require.True(t, proposedRevocation.HasRevocationFrom(setup.Trustee1.String()))
	require.True(t, proposedRevocation.HasRevocationFrom(setup.Trustee2.String()))

	// check that approved certificate still exists
	certificate, _ := querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NotNil(t, certificate)

	// check that revoked certificate does not exist
	_, err = queryRevokedCertificates(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that unique certificate key stays registered
	require.True(t,
		setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.RootIssuer, testconstants.RootSerialNumber))
}

func TestHandler_ApproveRevokeX509RootCert_ForEnoughApprovals(t *testing.T) {
	setup := Setup(t)

	// propose and approve x509 root certificate
	proposeAndApproveRootCertificate(setup, setup.Trustee1)

	// propose revocation of x509 root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	// get certificate for further comparison
	certificateBeforeRevocation, _ := querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NotNil(t, certificateBeforeRevocation)

	// approve
	approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
		setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveRevokeX509RootCert)
	require.NoError(t, err)

	// check that proposed certificate revocation does not exist anymore
	_, err = queryProposedCertificateRevocation(setup)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that approved certificate does not exist anymore
	_, err = queryApprovedCertificates(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// query and check revoked certificate
	revokedCertificate, _ := querySingleRevokedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, certificateBeforeRevocation, revokedCertificate)

	// check that unique certificate key stays registered
	require.True(t,
		setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, testconstants.RootIssuer, testconstants.RootSerialNumber))
}

func TestHandler_ApproveRevokeX509RootCert_ByNotTrustee(t *testing.T) {
	setup := Setup(t)

	// propose and approve x509 root certificate
	proposeAndApproveRootCertificate(setup, setup.Trustee1)

	// propose revocation of x509 root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.Vendor,
		dclauthtypes.CertificationCenter,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, 1)

		// approve
		approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
			accAddress.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
		_, err = setup.Handler(setup.Ctx, approveRevokeX509RootCert)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
	}
}

func TestHandler_ApproveRevokeX509RootCert_ProposedRevocationDoesNotExist(t *testing.T) {
	setup := Setup(t)

	// propose and approve x509 root certificate
	proposeAndApproveRootCertificate(setup, setup.Trustee1)

	// approve revocation of x509 root certificate
	approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, approveRevokeX509RootCert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrProposedCertificateRevocationDoesNotExist.Is(err))
}

func TestHandler_ApproveRevokeX509RootCert_Twice(t *testing.T) {
	setup := Setup(t)

	// propose and approve x509 root certificate
	proposeAndApproveRootCertificate(setup, setup.Trustee1)

	// propose revocation of x509 root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	// approve revocation by the same trustee
	approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveRevokeX509RootCert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

//nolint:funlen
func TestHandler_ApproveRevokeX509RootCert_ForTree(t *testing.T) {
	setup := Setup(t)

	// add root x509 certificate
	proposeAndApproveRootCertificate(setup, setup.Trustee1)

	// add intermediate x509 certificate
	addIntermediateX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addIntermediateX509Cert)
	require.NoError(t, err)

	// add leaf x509 certificate
	addLeafX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.LeafCertPem)
	_, err = setup.Handler(setup.Ctx, addLeafX509Cert)
	require.NoError(t, err)

	// propose revocation of x509 root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	// approve
	approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
		setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveRevokeX509RootCert)
	require.NoError(t, err)

	// check that root, intermediate and leaf certificates have been revoked
	allRevokedCertificates, _ := queryAllRevokedCertificates(setup)
	require.Equal(t, 3, len(allRevokedCertificates))
	require.Equal(t, testconstants.LeafSubject, allRevokedCertificates[0].Subject)
	require.Equal(t, testconstants.LeafSubjectKeyID, allRevokedCertificates[0].SubjectKeyId)
	require.Equal(t, 1, len(allRevokedCertificates[0].Certs))
	require.Equal(t, testconstants.LeafCertPem, allRevokedCertificates[0].Certs[0].PemCert)
	require.Equal(t, testconstants.RootSubject, allRevokedCertificates[1].Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, allRevokedCertificates[1].SubjectKeyId)
	require.Equal(t, 1, len(allRevokedCertificates[1].Certs))
	require.Equal(t, testconstants.RootCertPem, allRevokedCertificates[1].Certs[0].PemCert)
	require.Equal(t, testconstants.IntermediateSubject, allRevokedCertificates[2].Subject)
	require.Equal(t, testconstants.IntermediateSubjectKeyID, allRevokedCertificates[2].SubjectKeyId)
	require.Equal(t, 1, len(allRevokedCertificates[2].Certs))
	require.Equal(t, testconstants.IntermediateCertPem, allRevokedCertificates[2].Certs[0].PemCert)

	// check that no certificates stays approved
	allApprovedCertificates, err := queryAllApprovedCertificates(setup)
	require.NoError(t, err)
	require.Equal(t, 0, len(allApprovedCertificates))

	// check that no proposed certificate revocations exist
	allProposedCertificateRevocations, err := queryAllProposedCertificateRevocations(setup)
	require.NoError(t, err)
	require.Equal(t, 0, len(allProposedCertificateRevocations))

	// check that no child certificate identifiers are registered for revoked root certificate
	rootCertChildren, err := queryChildCertificates(
		setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
	require.Nil(t, rootCertChildren)

	// check that no child certificate identifiers are registered for revoked intermediate certificate
	intermediateCertChildren, err := queryChildCertificates(
		setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
	require.Nil(t, intermediateCertChildren)

	// check that no child certificate identifiers are registered for revoked leaf certificate
	leafCertChildren, err := queryChildCertificates(
		setup, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
	require.Nil(t, leafCertChildren)
}

func TestHandler_RevokeX509Cert(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee1)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.Vendor,
		dclauthtypes.CertificationCenter,
		dclauthtypes.Trustee,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, 1)

		// add x509 certificate
		addX509Cert := types.NewMsgAddX509Cert(accAddress.String(), testconstants.IntermediateCertPem)
		_, err := setup.Handler(setup.Ctx, addX509Cert)
		require.NoError(t, err)

		// get certificate for further comparison
		certificateBeforeRevocation, _ := querySingleApprovedCertificate(
			setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
		require.NotNil(t, certificateBeforeRevocation)

		// revoke x509 certificate
		revokeX509Cert := types.NewMsgRevokeX509Cert(
			accAddress.String(), testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID, testconstants.Info)
		_, err = setup.Handler(setup.Ctx, revokeX509Cert)
		require.NoError(t, err)

		// check that intermediate certificate has been revoked
		allRevokedCertificates, _ := queryAllRevokedCertificates(setup)
		require.Equal(t, 1, len(allRevokedCertificates))
		require.Equal(t, testconstants.IntermediateSubject, allRevokedCertificates[0].Subject)
		require.Equal(t, testconstants.IntermediateSubjectKeyID, allRevokedCertificates[0].SubjectKeyId)
		require.Equal(t, 1, len(allRevokedCertificates[0].Certs))
		require.Equal(t, *certificateBeforeRevocation, *allRevokedCertificates[0].Certs[0])

		// check that root certificate stays approved
		allApprovedCertificates, _ := queryAllApprovedCertificates(setup)
		require.Equal(t, 1, len(allApprovedCertificates))
		require.Equal(t, testconstants.IntermediateSubject, allRevokedCertificates[0].Subject)
		require.Equal(t, testconstants.IntermediateSubjectKeyID, allRevokedCertificates[0].SubjectKeyId)

		// check that no proposed certificate revocations have been created
		allProposedCertificateRevocations, _ := queryAllProposedCertificateRevocations(setup)
		require.NoError(t, err)
		require.Equal(t, 0, len(allProposedCertificateRevocations))

		// check that child certificate identifiers list of issuer do not exist anymore
		_, err = queryChildCertificates(setup, testconstants.IntermediateIssuer, testconstants.IntermediateAuthorityKeyID)
		require.Error(t, err)
		require.Equal(t, codes.NotFound, status.Code(err))

		// check that unique certificate key stays registered
		require.True(t, setup.Keeper.IsUniqueCertificatePresent(setup.Ctx,
			testconstants.IntermediateIssuer, testconstants.IntermediateSerialNumber))

		// cleanup for next iteration
		setup.Keeper.RemoveRevokedCertificates(setup.Ctx,
			testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
		setup.Keeper.RemoveUniqueCertificate(setup.Ctx,
			testconstants.IntermediateIssuer, testconstants.IntermediateSerialNumber)
	}
}

func TestHandler_RevokeX509Cert_CertificateDoesNotExist(t *testing.T) {
	setup := Setup(t)

	// revoke x509 certificate
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		setup.Trustee1.String(), testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_RevokeX509Cert_ForRootCertificate(t *testing.T) {
	setup := Setup(t)

	// propose and approve x509 root certificate
	proposeAndApproveRootCertificate(setup, setup.Trustee1)

	// revoke x509 root certificate
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		setup.Trustee1.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.Error(t, err)
	require.True(t, pkitypes.ErrInappropriateCertificateType.Is(err))
}

func TestHandler_RevokeX509Cert_ByNotOwner(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee1)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// add x509 certificate by `setup.Trustee`
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// store another account
	anotherTrustee := GenerateAccAddress()
	setup.AddAccount(anotherTrustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)

	// revoke x509 certificate by another account
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		anotherTrustee.String(), testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, revokeX509Cert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_RevokeX509Cert_ForTree(t *testing.T) {
	setup := Setup(t)

	// add root x509 certificate
	proposeAndApproveRootCertificate(setup, setup.Trustee1)

	// add intermediate x509 certificate
	addIntermediateX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addIntermediateX509Cert)
	require.NoError(t, err)

	// add leaf x509 certificate
	addLeafX509Cert := types.NewMsgAddX509Cert(setup.Trustee1.String(), testconstants.LeafCertPem)
	_, err = setup.Handler(setup.Ctx, addLeafX509Cert)
	require.NoError(t, err)

	// revoke x509 certificate
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		setup.Trustee1.String(), testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(t, err)

	// check that intermediate and leaf certificates have been revoked
	allRevokedCertificates, _ := queryAllRevokedCertificates(setup)
	require.Equal(t, 2, len(allRevokedCertificates))
	require.Equal(t, testconstants.LeafSubject, allRevokedCertificates[0].Subject)
	require.Equal(t, testconstants.LeafSubjectKeyID, allRevokedCertificates[0].SubjectKeyId)
	require.Equal(t, 1, len(allRevokedCertificates[0].Certs))
	require.Equal(t, testconstants.LeafCertPem, allRevokedCertificates[0].Certs[0].PemCert)
	require.Equal(t, testconstants.IntermediateSubject, allRevokedCertificates[1].Subject)
	require.Equal(t, testconstants.IntermediateSubjectKeyID, allRevokedCertificates[1].SubjectKeyId)
	require.Equal(t, 1, len(allRevokedCertificates[1].Certs))
	require.Equal(t, testconstants.IntermediateCertPem, allRevokedCertificates[1].Certs[0].PemCert)

	// check that root certificate stays approved
	allApprovedCertificates, _ := queryAllApprovedCertificates(setup)
	require.Equal(t, 1, len(allApprovedCertificates))
	require.Equal(t, testconstants.RootSubject, allApprovedCertificates[0].Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, allApprovedCertificates[0].SubjectKeyId)

	// check that no proposed certificate revocations have been created
	allProposedCertificateRevocations, _ := queryAllProposedCertificateRevocations(setup)
	require.NoError(t, err)
	require.Equal(t, 0, len(allProposedCertificateRevocations))

	// check that no child certificate identifiers are now registered for root certificate
	_, err = queryChildCertificates(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that no child certificate identifiers are registered for revoked intermediate certificate
	_, err = queryChildCertificates(setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// check that no child certificate identifiers are registered for revoked leaf certificate
	_, err = queryChildCertificates(setup, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
}

func TestHandler_RejectX509RootCert_TwoRejectApprovalsAreNeeded(t *testing.T) {
	setup := Setup(t)

	// propose x509 root certificate by account Trustee1
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// reject x509 root certificate by account Trustee2
	rejectAddX509RootCert := types.NewMsgRejectAddX509RootCert(setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddX509RootCert)
	require.NoError(t, err)

	// certificate should be in the entity <Proposed X509 Root Certificate>, because we haven't enough reject approvals
	proposedCertificate, err := queryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NoError(t, err)

	// check proposed certificate
	require.Equal(t, proposeAddX509RootCert.Cert, proposedCertificate.PemCert)
	require.Equal(t, proposeAddX509RootCert.Signer, proposedCertificate.Owner)
	require.Equal(t, testconstants.RootSubject, proposedCertificate.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, proposedCertificate.SubjectKeyId)
	require.Equal(t, testconstants.RootSerialNumber, proposedCertificate.SerialNumber)
	require.Equal(t, setup.Trustee1.String(), proposedCertificate.Approvals[0].Address)
	require.Equal(t, testconstants.Info, proposedCertificate.Approvals[0].Info)
	require.Equal(t, setup.Trustee2.String(), proposedCertificate.Rejects[0].Address)
	require.Equal(t, testconstants.Info, proposedCertificate.Rejects[0].Info)

	// reject x509 root certificate by account Trustee3
	rejectAddX509RootCert = types.NewMsgRejectAddX509RootCert(setup.Trustee3.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddX509RootCert)
	require.NoError(t, err)

	// certificate should not be in the entity <Proposed X509 Root Certificate>, because we have enough reject approvals
	_, err = queryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Error(t, err)

	// certificate should be in the entity <Rejected X509 Root Certificate>, because we have enough rejected approvals
	rejectedCertificate, err := queryRejectedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NoError(t, err)

	// check rejected certificate
	require.Equal(t, proposeAddX509RootCert.Cert, rejectedCertificate.PemCert)
	require.Equal(t, proposeAddX509RootCert.Signer, rejectedCertificate.Owner)
	require.Equal(t, testconstants.RootSubject, rejectedCertificate.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, rejectedCertificate.SubjectKeyId)
	require.Equal(t, testconstants.RootSerialNumber, rejectedCertificate.SerialNumber)
	require.Equal(t, setup.Trustee1.String(), rejectedCertificate.Approvals[0].Address)
	require.Equal(t, testconstants.Info, rejectedCertificate.Approvals[0].Info)
	require.Equal(t, setup.Trustee2.String(), rejectedCertificate.Rejects[0].Address)
	require.Equal(t, testconstants.Info, rejectedCertificate.Rejects[0].Info)
	require.Equal(t, setup.Trustee3.String(), rejectedCertificate.Rejects[1].Address)
	require.Equal(t, testconstants.Info, rejectedCertificate.Rejects[1].Info)
}

func TestHandler_RejectX509RootCert_ByNotTrustee(t *testing.T) {
	setup := Setup(t)

	// propose add x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.Vendor,
		dclauthtypes.CertificationCenter,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, 1)

		// reject x509 root certificate
		approveAddX509RootCert := types.NewMsgRejectAddX509RootCert(
			accAddress.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
		_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
	}
}

func TestHandler_Duplicate_RejectX509RootCert_FromTheSameTrustee(t *testing.T) {
	setup := Setup(t)

	// propose add x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// reject x509 root certificate by account Trustee2
	rejectAddX509RootCert := types.NewMsgRejectAddX509RootCert(setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddX509RootCert)
	require.NoError(t, err)

	// second time reject x509 root certificate by account Trustee2
	rejectAddX509RootCert = types.NewMsgRejectAddX509RootCert(setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddX509RootCert)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_ApproveX509RootCertAndRejectX509RootCert_FromTheSameTrustee(t *testing.T) {
	setup := Setup(t)
	// propose add x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.Trustee,
	} {
		accAddress := GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, 1)

		// approve x509 root certificate by account Trustee2
		approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
		_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
		require.NoError(t, err)

		pendingCert, _ := setup.Keeper.GetProposedCertificate(setup.Ctx, testconstants.RootSubject, testconstants.RootSubjectKeyID)
		prevRejectsLen := len(pendingCert.Rejects)
		prevApprovalsLen := len(pendingCert.Approvals)
		// reject x509 root certificate by account Trustee2
		rejectAddX509RootCert := types.NewMsgRejectAddX509RootCert(setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
		_, err = setup.Handler(setup.Ctx, rejectAddX509RootCert)
		require.NoError(t, err)

		pendingCert, found := setup.Keeper.GetProposedCertificate(setup.Ctx, testconstants.RootSubject, testconstants.RootSubjectKeyID)
		require.True(t, found)
		require.Equal(t, len(pendingCert.Rejects), prevRejectsLen+1)
		require.Equal(t, len(pendingCert.Approvals), prevApprovalsLen-1)
	}
}

func TestHandler_RejectX509RootCertAndApproveX509RootCert_FromTheSameTrustee(t *testing.T) {
	setup := Setup(t)
	// propose add x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.Trustee,
	} {
		accAddress := GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, 1)

		// reject x509 root certificate by account Trustee2
		rejectAddX509RootCert := types.NewMsgRejectAddX509RootCert(setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
		_, err = setup.Handler(setup.Ctx, rejectAddX509RootCert)
		require.NoError(t, err)

		pendingCert, _ := setup.Keeper.GetProposedCertificate(setup.Ctx, testconstants.RootSubject, testconstants.RootSubjectKeyID)
		prevRejectsLen := len(pendingCert.Rejects)
		prevApprovalsLen := len(pendingCert.Approvals)
		// approve x509 root certificate by account Trustee2
		approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
		_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
		require.NoError(t, err)

		pendingCert, found := setup.Keeper.GetProposedCertificate(setup.Ctx, testconstants.RootSubject, testconstants.RootSubjectKeyID)
		require.True(t, found)
		require.Equal(t, len(pendingCert.Rejects), prevRejectsLen-1)
		require.Equal(t, len(pendingCert.Approvals), prevApprovalsLen+1)
	}
}

func TestHandler_DoubleTimeRejectX509RootCert(t *testing.T) {
	setup := Setup(t)

	// propose x509 root certificate by account Trustee1
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// reject x509 root certificate by account Trustee2
	rejectAddX509RootCert := types.NewMsgRejectAddX509RootCert(setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddX509RootCert)
	require.NoError(t, err)

	// certificate should be in the entity <Proposed X509 Root Certificate>, because we haven't enough reject approvals
	proposedCertificate, err := queryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NoError(t, err)

	// check proposed certificate
	require.Equal(t, proposeAddX509RootCert.Cert, proposedCertificate.PemCert)
	require.Equal(t, proposeAddX509RootCert.Signer, proposedCertificate.Owner)
	require.Equal(t, testconstants.RootSubject, proposedCertificate.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, proposedCertificate.SubjectKeyId)
	require.Equal(t, testconstants.RootSerialNumber, proposedCertificate.SerialNumber)
	require.Equal(t, setup.Trustee1.String(), proposedCertificate.Approvals[0].Address)
	require.Equal(t, testconstants.Info, proposedCertificate.Approvals[0].Info)
	require.Equal(t, setup.Trustee2.String(), proposedCertificate.Rejects[0].Address)
	require.Equal(t, testconstants.Info, proposedCertificate.Rejects[0].Info)

	// reject x509 root certificate by account Trustee3
	rejectAddX509RootCert = types.NewMsgRejectAddX509RootCert(setup.Trustee3.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddX509RootCert)
	require.NoError(t, err)

	// certificate should not be in the entity <Proposed X509 Root Certificate>, because we have enough reject approvals
	_, err = queryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Error(t, err)

	// certificate should be in the entity <Rejected X509 Root Certificate>, because we have enough rejected approvals
	rejectedCertificate, err := queryRejectedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NoError(t, err)

	// check rejected certificate
	require.Equal(t, proposeAddX509RootCert.Cert, rejectedCertificate.PemCert)
	require.Equal(t, proposeAddX509RootCert.Signer, rejectedCertificate.Owner)
	require.Equal(t, testconstants.RootSubject, rejectedCertificate.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, rejectedCertificate.SubjectKeyId)
	require.Equal(t, testconstants.RootSerialNumber, rejectedCertificate.SerialNumber)
	require.Equal(t, setup.Trustee1.String(), rejectedCertificate.Approvals[0].Address)
	require.Equal(t, testconstants.Info, rejectedCertificate.Approvals[0].Info)
	require.Equal(t, setup.Trustee2.String(), rejectedCertificate.Rejects[0].Address)
	require.Equal(t, testconstants.Info, rejectedCertificate.Rejects[0].Info)
	require.Equal(t, setup.Trustee3.String(), rejectedCertificate.Rejects[1].Address)
	require.Equal(t, testconstants.Info, rejectedCertificate.Rejects[1].Info)

	// propose x509 root certificate by account Trustee1
	proposeAddX509RootCert = types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
	_, err = setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// certificate should be in the entity <Proposed X509 Root Certificate>, because we haven't enough reject approvals
	_, err = queryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NoError(t, err)

	// certificate should not be in the entity <Rejected X509 Root Certificate>, because we have propose that certificate
	_, err = queryRejectedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Error(t, err)

	// reject x509 root certificate by account Trustee3
	rejectAddX509RootCert = types.NewMsgRejectAddX509RootCert(setup.Trustee3.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddX509RootCert)
	require.NoError(t, err)

	// reject x509 root certificate by account Trustee2
	rejectAddX509RootCert = types.NewMsgRejectAddX509RootCert(setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddX509RootCert)
	require.NoError(t, err)

	// certificate should be in the entity <Rejected X509 Root Certificate>, because we have enough rejected approvals
	rejectedCertificate, err = queryRejectedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NoError(t, err)

	// check rejected certificate
	require.Equal(t, proposeAddX509RootCert.Cert, rejectedCertificate.PemCert)
	require.Equal(t, proposeAddX509RootCert.Signer, rejectedCertificate.Owner)
	require.Equal(t, testconstants.RootSubject, rejectedCertificate.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, rejectedCertificate.SubjectKeyId)
	require.Equal(t, testconstants.RootSerialNumber, rejectedCertificate.SerialNumber)
	require.Equal(t, setup.Trustee1.String(), rejectedCertificate.Approvals[0].Address)
	require.Equal(t, testconstants.Info, rejectedCertificate.Approvals[0].Info)
	require.Equal(t, setup.Trustee3.String(), rejectedCertificate.Rejects[0].Address)
	require.Equal(t, testconstants.Info, rejectedCertificate.Rejects[0].Info)
	require.Equal(t, setup.Trustee2.String(), rejectedCertificate.Rejects[1].Address)
	require.Equal(t, testconstants.Info, rejectedCertificate.Rejects[1].Info)
}

func proposeAndApproveRootCertificate(setup *TestSetup, ownerTrustee sdk.AccAddress) {
	// ensure that `ownerTrustee` is trustee to eventually have enough approvals
	require.True(setup.T, setup.DclauthKeeper.HasRole(setup.Ctx, ownerTrustee, types.RootCertificateApprovalRole))

	// propose x509 root certificate by `ownerTrustee`
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(ownerTrustee.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(setup.T, err)

	// approve x509 root certificate by another trustee
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(setup.T, err)

	// check that root certificate has been approved
	approvedCertificate, err := querySingleApprovedCertificate(
		setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NoError(setup.T, err)
	require.NotNil(setup.T, approvedCertificate)
}

func TestHandler_RejectX509RootCert_TwoRejectApprovalsAreNeeded_FiveTrustees(t *testing.T) {
	setup := Setup(t)

	// we have 5 trustees: 1 approval comes from propose => we need 2 rejects to make certificate rejected

	// store 4th trustee
	fourthTrustee := GenerateAccAddress()
	setup.AddAccount(fourthTrustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)

	// store 5th trustee
	fifthTrustee := GenerateAccAddress()
	setup.AddAccount(fifthTrustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)

	// propose x509 root certificate by account Trustee1
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// reject x509 root certificate by account Trustee2
	rejectAddX509RootCert := types.NewMsgRejectAddX509RootCert(setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddX509RootCert)
	require.NoError(t, err)

	// certificate should be in the entity <Proposed X509 Root Certificate>, because we haven't enough reject approvals
	proposedCertificate, err := queryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NoError(t, err)

	// check proposed certificate
	require.Equal(t, proposeAddX509RootCert.Cert, proposedCertificate.PemCert)
	require.Equal(t, proposeAddX509RootCert.Signer, proposedCertificate.Owner)
	require.Equal(t, testconstants.RootSubject, proposedCertificate.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, proposedCertificate.SubjectKeyId)
	require.Equal(t, testconstants.RootSerialNumber, proposedCertificate.SerialNumber)

	// reject x509 root certificate by account Trustee3
	rejectAddX509RootCert = types.NewMsgRejectAddX509RootCert(setup.Trustee3.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddX509RootCert)
	require.NoError(t, err)

	// certificate should be in the entity <Rejected X509 Root Certificate>, because we have enough rejected approvals
	rejectedCertificate, err := queryRejectedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NoError(t, err)

	// check rejected certificate
	require.Equal(t, proposeAddX509RootCert.Cert, rejectedCertificate.PemCert)
	require.Equal(t, proposeAddX509RootCert.Signer, rejectedCertificate.Owner)
	require.Equal(t, testconstants.RootSubject, rejectedCertificate.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, rejectedCertificate.SubjectKeyId)
	require.Equal(t, testconstants.RootSerialNumber, rejectedCertificate.SerialNumber)
}

func TestHandler_ApproveX509RootCert_FourApprovalsAreNeeded_FiveTrustees(t *testing.T) {
	setup := Setup(t)

	// we have 5 trustees: 1 approval comes from propose => we need 3 more approvals

	// store 4th trustee
	fourthTrustee := GenerateAccAddress()
	setup.AddAccount(fourthTrustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)

	// store 5th trustee
	fifthTrustee := GenerateAccAddress()
	setup.AddAccount(fifthTrustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)

	// propose x509 root certificate by account Trustee1
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// approve x509 root certificate by account Trustee2
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(setup.Trustee2.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// approve x509 root certificate by account Trustee3
	approveAddX509RootCert = types.NewMsgApproveAddX509RootCert(setup.Trustee3.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// reject x509 root certificate by account Trustee4
	rejectAddX509RootCert := types.NewMsgRejectAddX509RootCert(fourthTrustee.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddX509RootCert)
	require.NoError(t, err)

	// certificate should be in the entity <Proposed X509 Root Certificate>, because we haven't enough approvals
	proposedCertificate, err := queryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NoError(t, err)

	// check proposed certificate
	require.Equal(t, proposeAddX509RootCert.Cert, proposedCertificate.PemCert)
	require.Equal(t, proposeAddX509RootCert.Signer, proposedCertificate.Owner)
	require.Equal(t, testconstants.RootSubject, proposedCertificate.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, proposedCertificate.SubjectKeyId)
	require.Equal(t, testconstants.RootSerialNumber, proposedCertificate.SerialNumber)

	// approve x509 root certificate by account Trustee5
	approveAddX509RootCert = types.NewMsgApproveAddX509RootCert(fifthTrustee.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// certificate should be in the entity <X509 Root Certificate>, because we have enough approvals
	approvedCertificate, err := querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NoError(t, err)

	// check certificate
	require.Equal(t, proposeAddX509RootCert.Cert, approvedCertificate.PemCert)
	require.Equal(t, proposeAddX509RootCert.Signer, approvedCertificate.Owner)
	require.Equal(t, testconstants.RootSubject, approvedCertificate.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, approvedCertificate.SubjectKeyId)
	require.Equal(t, testconstants.RootSerialNumber, approvedCertificate.SerialNumber)
}

func TestHandler_AddPkiRevocationDistributionPoint_PAASenderNotVendor(t *testing.T) {
	setup := Setup(t)

	addPkiRevocationDistributionPoint := types.MsgAddPkiRevocationDistributionPoint{
		Signer:               setup.Trustee1.String(),
		Vid:                  65521,
		IsPAA:                true,
		Pid:                  8,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err := setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_AddPkiRevocationDistributionPoint_PAACertEncodesVidSenderVidNotEqualVidField(t *testing.T) {
	setup := Setup(t)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 1)

	addPkiRevocationDistributionPoint := types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  65521,
		IsPAA:                true,
		Pid:                  8,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err := setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrCRLSignerCertificateVidNotEqualAccountVid)
}

func TestHandler_AddPkiRevocationDistributionPoint_PAACertNotFound(t *testing.T) {
	setup := Setup(t)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 65521)

	addPkiRevocationDistributionPoint := types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  65521,
		IsPAA:                true,
		Pid:                  8,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err := setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrCertificateDoesNotExist)
}

func TestHandler_AddPkiRevocationDistributionPoint_PAAPemValueOfStoredCertNotEqualToCrlSignerCertifcate(t *testing.T) {
	setup := Setup(t)

	// propose x509 root certificate by account Trustee1
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.PAACertWithNumericVid, testconstants.Info, testconstants.Vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.PAACertWithNumericVidSubject, testconstants.PAACertWithNumericVidSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 65522)

	addPkiRevocationDistributionPoint := types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  65522,
		IsPAA:                true,
		Pid:                  8,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid1,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrCertificateDoesNotExist)
}

func TestHandler_AddPkiRevocationDistributionPoint_PAISenderNotVendor(t *testing.T) {
	setup := Setup(t)

	addPkiRevocationDistributionPoint := types.MsgAddPkiRevocationDistributionPoint{
		Signer:               setup.Trustee1.String(),
		Vid:                  65521,
		IsPAA:                false,
		Pid:                  8,
		CrlSignerCertificate: testconstants.PAICertWithNumericPidVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err := setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_AddPkiRevocationDistributionPoint_PAINotChainedBackToDCLCerts(t *testing.T) {
	setup := Setup(t)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 65521)

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
	_, err := setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrInvalidCertificate)
}

func TestHandler_AddPkiRevocationDistributionPoint_PAAAlreadyExists(t *testing.T) {
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
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	addPkiRevocationDistributionPoint = types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  65521,
		IsPAA:                true,
		Pid:                  8,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrPkiRevocationDistributionPointAlreadyExists)
}

func TestHandler_AddPkiRevocationDistributionPoint_PAAWithVid(t *testing.T) {
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
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	revocationPoint, isFound := setup.Keeper.GetPkiRevocationDistributionPoint(setup.Ctx, 65521, "label", testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	require.Equal(t, revocationPoint.CrlSignerCertificate, addPkiRevocationDistributionPoint.CrlSignerCertificate)
	require.Equal(t, revocationPoint.DataDigest, addPkiRevocationDistributionPoint.DataDigest)
	require.Equal(t, revocationPoint.DataDigestType, addPkiRevocationDistributionPoint.DataDigestType)
	require.Equal(t, revocationPoint.DataFileSize, addPkiRevocationDistributionPoint.DataFileSize)
	require.Equal(t, revocationPoint.DataURL, addPkiRevocationDistributionPoint.DataURL)
	require.Equal(t, revocationPoint.IsPAA, addPkiRevocationDistributionPoint.IsPAA)
	require.Equal(t, revocationPoint.IssuerSubjectKeyID, addPkiRevocationDistributionPoint.IssuerSubjectKeyID)
	require.Equal(t, revocationPoint.Label, addPkiRevocationDistributionPoint.Label)
	require.Equal(t, revocationPoint.Pid, addPkiRevocationDistributionPoint.Pid)
	require.Equal(t, revocationPoint.RevocationType, addPkiRevocationDistributionPoint.RevocationType)
	require.Equal(t, revocationPoint.Vid, addPkiRevocationDistributionPoint.Vid)

	revocationPointBySubjectKeyID, isFound := setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	require.Equal(t, revocationPointBySubjectKeyID.Points[0].CrlSignerCertificate, addPkiRevocationDistributionPoint.CrlSignerCertificate)
	require.Equal(t, revocationPointBySubjectKeyID.Points[0].DataDigest, addPkiRevocationDistributionPoint.DataDigest)
	require.Equal(t, revocationPointBySubjectKeyID.Points[0].DataDigestType, addPkiRevocationDistributionPoint.DataDigestType)
	require.Equal(t, revocationPointBySubjectKeyID.Points[0].DataFileSize, addPkiRevocationDistributionPoint.DataFileSize)
	require.Equal(t, revocationPointBySubjectKeyID.Points[0].DataURL, addPkiRevocationDistributionPoint.DataURL)
	require.Equal(t, revocationPointBySubjectKeyID.Points[0].IsPAA, addPkiRevocationDistributionPoint.IsPAA)
	require.Equal(t, revocationPointBySubjectKeyID.Points[0].IssuerSubjectKeyID, addPkiRevocationDistributionPoint.IssuerSubjectKeyID)
	require.Equal(t, revocationPointBySubjectKeyID.Points[0].Label, addPkiRevocationDistributionPoint.Label)
	require.Equal(t, revocationPointBySubjectKeyID.Points[0].Pid, addPkiRevocationDistributionPoint.Pid)
	require.Equal(t, revocationPointBySubjectKeyID.Points[0].RevocationType, addPkiRevocationDistributionPoint.RevocationType)
	require.Equal(t, revocationPointBySubjectKeyID.Points[0].Vid, addPkiRevocationDistributionPoint.Vid)
}

func TestHandler_AddPkiRevocationDistributionPoint_PAIWithNumericVidPid(t *testing.T) {
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

	addX509Cert := types.NewMsgAddX509Cert(vendorAcc.String(), testconstants.PAICertWithNumericPidVid)
	_, err = setup.Handler(setup.Ctx, addX509Cert)
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
}

func TestHandler_AddPkiRevocationDistributionPoint_PAIWithStringVidPid(t *testing.T) {
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
}

func TestHandler_AddPkiRevocationDistributionPoint_DataURLNotUnique(t *testing.T) {
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

	addPkiRevocationDistributionPoint = types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  65522,
		IsPAA:                false,
		Pid:                  8,
		CrlSignerCertificate: testconstants.PAICertWithPidVid,
		Label:                "label-new",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrPkiRevocationDistributionPointAlreadyExists)
}

func TestHandler_AddPkiRevocationDistributionPoint_PAIWithVid(t *testing.T) {
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
		CrlSignerCertificate: testconstants.PAICertWithVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.NoError(t, err)
}

func TestHandler_RevocationPointsByIssuerSubjectKeyID(t *testing.T) {
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

	revocationPointBySubjectKeyID, isFound := setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
	require.False(t, isFound)
	require.Equal(t, len(revocationPointBySubjectKeyID.Points), 0)

	addPkiRevocationDistributionPoint := types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  65521,
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

	revocationPointBySubjectKeyID, isFound = setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	require.Equal(t, len(revocationPointBySubjectKeyID.Points), 1)

	addPkiRevocationDistributionPoint = types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  65521,
		IsPAA:                true,
		Pid:                  8,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                "label1",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	revocationPointBySubjectKeyID, isFound = setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	require.Equal(t, len(revocationPointBySubjectKeyID.Points), 2)

	dataURLNew := testconstants.DataURL + "/new"
	updatePkiRevocationDistributionPoint := types.MsgUpdatePkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  65521,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                "label",
		DataURL:              dataURLNew,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
	}
	_, err = setup.Handler(setup.Ctx, &updatePkiRevocationDistributionPoint)
	require.NoError(t, err)

	revocationPointBySubjectKeyID, isFound = setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	require.Equal(t, len(revocationPointBySubjectKeyID.Points), 2)
	require.Equal(t, revocationPointBySubjectKeyID.Points[0].CrlSignerCertificate, updatePkiRevocationDistributionPoint.CrlSignerCertificate)
	require.Equal(t, revocationPointBySubjectKeyID.Points[0].DataURL, updatePkiRevocationDistributionPoint.DataURL)

	deletePkiRevocationDistributionPoint := types.MsgDeletePkiRevocationDistributionPoint{
		Signer:             vendorAcc.String(),
		Vid:                65521,
		Label:              "label",
		IssuerSubjectKeyID: testconstants.SubjectKeyIDWithoutColons,
	}
	_, err = setup.Handler(setup.Ctx, &deletePkiRevocationDistributionPoint)
	require.NoError(t, err)

	revocationPointBySubjectKeyID, isFound = setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	require.Equal(t, len(revocationPointBySubjectKeyID.Points), 1)
}

func TestHandler_UpdatePkiRevocationDistributionPoint(t *testing.T) {
	var err error
	vendorAcc := GenerateAccAddress()
	const vid int32 = 65521
	addedRevocationPoints := []types.MsgAddPkiRevocationDistributionPoint{
		{
			Signer:               vendorAcc.String(),
			Vid:                  vid,
			IsPAA:                true,
			Pid:                  8,
			CrlSignerCertificate: testconstants.PAACertWithNumericVid,
			Label:                testconstants.ProductLabel + "-1",
			DataURL:              testconstants.DataURL + "/1",
			IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
			RevocationType:       1,
		},
		{
			Signer:               vendorAcc.String(),
			Vid:                  vid,
			IsPAA:                true,
			Pid:                  8,
			CrlSignerCertificate: testconstants.PAACertWithNumericVid,
			Label:                testconstants.ProductLabel + "-2",
			DataURL:              testconstants.DataURL + "/2",
			IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
			RevocationType:       1,
		},
	}
	addedRevocation := addedRevocationPoints[0]
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
			name:  "Valid: DataFieldsProvidedWhenRevocationType1",
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
		{
			valid: false,
			name:  "Valid: not unique DataURL for Issuer",
			updatedRevocation: types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:             addedRevocation.Signer,
				Vid:                addedRevocation.Vid,
				Label:              addedRevocation.Label,
				DataURL:            addedRevocationPoints[1].DataURL,
				IssuerSubjectKeyID: addedRevocation.IssuerSubjectKeyID,
			},
			err: pkitypes.ErrPkiRevocationDistributionPointAlreadyExists,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := Setup(t)
			setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

			// propose x509 root certificate by account Trustee1
			proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.PAACertWithNumericVid, testconstants.Info, testconstants.Vid)
			_, err = setup.Handler(setup.Ctx, proposeAddX509RootCert)
			require.NoError(t, err)

			// approve
			approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
				setup.Trustee2.String(), testconstants.PAACertWithNumericVidSubject, testconstants.PAACertWithNumericVidSubjectKeyID, testconstants.Info)
			_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
			require.NoError(t, err)

			// add revocation
			for _, revocationPoint := range addedRevocationPoints {
				_, err = setup.Handler(setup.Ctx, &revocationPoint)
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

func TestHandler_UpdatePkiRevocationDistributionPoint_PAIWithVidPid(t *testing.T) {
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
		Signer:               vendorAcc.String(),
		Vid:                  65521,
		CrlSignerCertificate: testconstants.PAICertWithNumericPidVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
	}
	_, err = setup.Handler(setup.Ctx, &updatePkiRevocationDistributionPoint)
	require.NoError(t, err)
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
		Vid:                  65521,
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
		Vid:                  65521,
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
		Vid:                  65521,
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
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.PAACertWithNumericVid, testconstants.Info, testconstants.Vid)
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

func TestHandler_DeletePkiRevocationDistributionPoint_NotFound(t *testing.T) {
	setup := Setup(t)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 65521)

	deletePkiRevocationDistributionPoint := types.MsgDeletePkiRevocationDistributionPoint{
		Signer:             vendorAcc.String(),
		Vid:                65521,
		Label:              "label",
		IssuerSubjectKeyID: testconstants.SubjectKeyIDWithoutColons,
	}
	_, err := setup.Handler(setup.Ctx, &deletePkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrPkiRevocationDistributionPointDoesNotExists)
}

func TestHandler_DeletePkiRevocationDistributionPoint_PAA(t *testing.T) {
	setup := Setup(t)
	const vid int32 = 65521
	const label string = "label"

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
		Vid:                  vid,
		IsPAA:                true,
		Pid:                  8,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                label,
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	deletePkiRevocationDistributionPoint := types.MsgDeletePkiRevocationDistributionPoint{
		Signer:             vendorAcc.String(),
		Vid:                vid,
		Label:              label,
		IssuerSubjectKeyID: testconstants.SubjectKeyIDWithoutColons,
	}
	_, err = setup.Handler(setup.Ctx, &deletePkiRevocationDistributionPoint)
	require.NoError(t, err)

	_, isFound := setup.Keeper.GetPkiRevocationDistributionPoint(setup.Ctx, vid, label, testconstants.SubjectKeyIDWithoutColons)
	require.False(t, isFound)

	_, isFound = setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
	require.False(t, isFound)
}

func TestHandler_DeletePkiRevocationDistributionPoint_PAI(t *testing.T) {
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

	// add intermediate certificate
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

	deletePkiRevocationDistributionPoint := types.MsgDeletePkiRevocationDistributionPoint{
		Signer:             vendorAcc.String(),
		Vid:                65521,
		Label:              "label",
		IssuerSubjectKeyID: testconstants.SubjectKeyIDWithoutColons,
	}
	_, err = setup.Handler(setup.Ctx, &deletePkiRevocationDistributionPoint)
	require.NoError(t, err)
}

func TestHandler_DeletePkiRevocationDistributionPoint_PAASenderNotVendor(t *testing.T) {
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
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	deletePkiRevocationDistributionPoint := types.MsgDeletePkiRevocationDistributionPoint{
		Signer:             setup.Trustee1.String(),
		Vid:                65521,
		Label:              "label",
		IssuerSubjectKeyID: testconstants.SubjectKeyIDWithoutColons,
	}
	_, err = setup.Handler(setup.Ctx, &deletePkiRevocationDistributionPoint)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_DeletePkiRevocationDistributionPoint_PAASenderVidNotEqualCertVid(t *testing.T) {
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

	deletePkiRevocationDistributionPoint := types.MsgDeletePkiRevocationDistributionPoint{
		Signer:             newVendorAcc.String(),
		Vid:                65521,
		Label:              "label",
		IssuerSubjectKeyID: testconstants.SubjectKeyIDWithoutColons,
	}
	_, err = setup.Handler(setup.Ctx, &deletePkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrCRLSignerCertificateVidNotEqualAccountVid)
}

func TestHandler_DeletePkiRevocationDistributionPoint_PAISenderNotVendor(t *testing.T) {
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

	// add intermediate certificate
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

	deletePkiRevocationDistributionPoint := types.MsgDeletePkiRevocationDistributionPoint{
		Signer:             setup.Trustee1.String(),
		Vid:                65521,
		Label:              "label",
		IssuerSubjectKeyID: testconstants.SubjectKeyIDWithoutColons,
	}
	_, err = setup.Handler(setup.Ctx, &deletePkiRevocationDistributionPoint)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_DeletePkiRevocationDistributionPoint_PAISenderVidNotEqualCertVid(t *testing.T) {
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

	// add intermediate certificate
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

	newVendorAcc := GenerateAccAddress()
	setup.AddAccount(newVendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 65522)

	deletePkiRevocationDistributionPoint := types.MsgDeletePkiRevocationDistributionPoint{
		Signer:             newVendorAcc.String(),
		Vid:                65521,
		Label:              "label",
		IssuerSubjectKeyID: testconstants.SubjectKeyIDWithoutColons,
	}
	_, err = setup.Handler(setup.Ctx, &deletePkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrCRLSignerCertificateVidNotEqualAccountVid)
}

func queryProposedCertificate(
	setup *TestSetup,
	subject string,
	subjectKeyID string,
) (*types.ProposedCertificate, error) {
	// query proposed certificate
	req := &types.QueryGetProposedCertificateRequest{
		Subject:      subject,
		SubjectKeyId: subjectKeyID,
	}

	resp, err := setup.Keeper.ProposedCertificate(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.ProposedCertificate, nil
}

func queryAllApprovedCertificates(setup *TestSetup) ([]types.ApprovedCertificates, error) {
	// query all certificates
	req := &types.QueryAllApprovedCertificatesRequest{}

	resp, err := setup.Keeper.ApprovedCertificatesAll(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return resp.ApprovedCertificates, nil
}

func querySingleApprovedCertificate(
	setup *TestSetup,
	subject string,
	subjectKeyID string,
) (*types.Certificate, error) {
	certificates, err := queryApprovedCertificates(setup, subject, subjectKeyID)
	if err != nil {
		return nil, err
	}

	if len(certificates.Certs) > 1 {
		require.Fail(setup.T, "More than 1 certificate returned")
	}

	return certificates.Certs[0], nil
}

func queryApprovedCertificates(
	setup *TestSetup,
	subject string,
	subjectKeyID string,
) (*types.ApprovedCertificates, error) {
	// query certificate
	req := &types.QueryGetApprovedCertificatesRequest{
		Subject:      subject,
		SubjectKeyId: subjectKeyID,
	}

	resp, err := setup.Keeper.ApprovedCertificates(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.ApprovedCertificates, nil
}

func queryAllProposedCertificateRevocations(setup *TestSetup) ([]types.ProposedCertificateRevocation, error) {
	// query all proposed certificate revocations
	req := &types.QueryAllProposedCertificateRevocationRequest{}

	resp, err := setup.Keeper.ProposedCertificateRevocationAll(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return resp.ProposedCertificateRevocation, nil
}

func queryProposedCertificateRevocation(
	setup *TestSetup,
) (*types.ProposedCertificateRevocation, error) {
	// query proposed certificate revocation
	req := &types.QueryGetProposedCertificateRevocationRequest{
		Subject:      testconstants.RootSubject,
		SubjectKeyId: testconstants.RootSubjectKeyID,
	}

	resp, err := setup.Keeper.ProposedCertificateRevocation(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.ProposedCertificateRevocation, nil
}

func queryAllRevokedCertificates(setup *TestSetup) ([]types.RevokedCertificates, error) {
	// query all revoked certificates
	req := &types.QueryAllRevokedCertificatesRequest{}

	resp, err := setup.Keeper.RevokedCertificatesAll(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return resp.RevokedCertificates, nil
}

func querySingleRevokedCertificate(
	setup *TestSetup,
	subject string,
	subjectKeyID string,
) (*types.Certificate, error) {
	certificates, err := queryRevokedCertificates(setup, subject, subjectKeyID)
	if err != nil {
		return nil, err
	}

	if len(certificates.Certs) > 1 {
		require.Fail(setup.T, "More than 1 certificate returned")
	}

	return certificates.Certs[0], nil
}

func queryRevokedCertificates(
	setup *TestSetup,
	subject string,
	subjectKeyID string,
) (*types.RevokedCertificates, error) {
	// query revoked certificate
	req := &types.QueryGetRevokedCertificatesRequest{
		Subject:      subject,
		SubjectKeyId: subjectKeyID,
	}

	resp, err := setup.Keeper.RevokedCertificates(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.RevokedCertificates, nil
}

func queryChildCertificates(
	setup *TestSetup,
	issuer string,
	authorityKeyID string,
) (*types.ChildCertificates, error) {
	// query certificate
	req := &types.QueryGetChildCertificatesRequest{
		Issuer:         issuer,
		AuthorityKeyId: authorityKeyID,
	}

	resp, err := setup.Keeper.ChildCertificates(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.ChildCertificates, nil
}

//nolint:unparam
func queryRejectedCertificate(
	setup *TestSetup,
	subject string,
	subjectKeyID string,
) (*types.Certificate, error) {
	certificates, err := queryRejectedCertificates(setup, subject, subjectKeyID)
	if err != nil {
		return nil, err
	}

	if len(certificates.Certs) > 1 {
		require.Fail(setup.T, "More than 1 certificate returned")
	}

	return certificates.Certs[0], nil
}

func queryRejectedCertificates(
	setup *TestSetup,
	subject string,
	subjectKeyID string,
) (*types.RejectedCertificate, error) {
	req := &types.QueryGetRejectedCertificatesRequest{
		Subject:      subject,
		SubjectKeyId: subjectKeyID,
	}

	resp, err := setup.Keeper.RejectedCertificate(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.RejectedCertificate, nil
}

func rootCertificate(address sdk.AccAddress) types.Certificate {
	return types.NewRootCertificate(
		testconstants.RootCertPem,
		testconstants.RootSubject,
		testconstants.RootSubjectAsText,
		testconstants.RootSubjectKeyID,
		testconstants.RootSerialNumber,
		address.String(),
		[]*types.Grant{},
		[]*types.Grant{},
		testconstants.Vid,
	)
}

func intermediateCertificate(address sdk.AccAddress) types.Certificate {
	return types.NewNonRootCertificate(
		testconstants.IntermediateCertPem,
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectAsText,
		testconstants.IntermediateSubjectKeyID,
		testconstants.IntermediateSerialNumber,
		testconstants.IntermediateIssuer,
		testconstants.IntermediateAuthorityKeyID,
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		address.String(),
	)
}

func uniqueCertificate(issuer string, serialNumber string) types.UniqueCertificate {
	return types.UniqueCertificate{
		Issuer:       issuer,
		SerialNumber: serialNumber,
		Present:      true,
	}
}

func certificateIdentifier(subject string, subjectKeyID string) types.CertificateIdentifier {
	return types.CertificateIdentifier{
		Subject:      subject,
		SubjectKeyId: subjectKeyID,
	}
}
