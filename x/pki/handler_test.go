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

// Creates an account with the specified role and permissions as part of the test setup and returns it.
func AddAccountWithRoleAndPermissions(setup *TestSetup, role dclauthtypes.AccountRole, vid int32) sdk.AccAddress {
	accAddress := GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, vid)

	return accAddress
}

// Creates a proposal to add the X509 root certificate.
func proposeX509RootCert(setup *TestSetup, trusteeAddr sdk.AccAddress, root, info string, vid int32) (*types.MsgProposeAddX509RootCert, error) {
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(trusteeAddr.String(), root, info, vid)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)

	return proposeAddX509RootCert, err
}

// Rejects the request to add the X509 root certificate.
func rejectX509RootCert(setup *TestSetup, signerAddr sdk.AccAddress, info string) error {
	rejectAddX509RootCert := types.NewMsgRejectAddX509RootCert(signerAddr.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, info)
	_, err := setup.Handler(setup.Ctx, rejectAddX509RootCert)

	return err
}

// Сreates and sends a message to approve the addition of the X509 root certificate.
func approveX509RootCert(setup *TestSetup, signerAddr sdk.AccAddress, subject, subjectKeyID, info string) (*types.MsgApproveAddX509RootCert, error) {
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(signerAddr.String(), subject, subjectKeyID, info)
	_, err := setup.Handler(setup.Ctx, approveAddX509RootCert)

	return approveAddX509RootCert, err
}

// Сreates a new message to add the X509 certificate.
func addX509Cert(setup *TestSetup, accAddress sdk.AccAddress, cert string) (*types.MsgAddX509Cert, error) {
	addX509Cert := types.NewMsgAddX509Cert(accAddress.String(), cert)
	_, err := setup.Handler(setup.Ctx, addX509Cert)

	return addX509Cert, err
}

// Creates a proposal to revoke the X509 root certificate.
func proposeRevokeX509RootCert(setup *TestSetup, signerAddr sdk.AccAddress, subject, subjectKeyID, info string) error {
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(signerAddr.String(), subject, subjectKeyID, info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)

	return err
}

// Creates a new message to approve revocation of the X509 root certificate.
func approveRevokeX509RootCert(setup *TestSetup, trusteeAddr sdk.AccAddress, info string) error {
	approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(trusteeAddr.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID, info)
	_, err := setup.Handler(setup.Ctx, approveRevokeX509RootCert)

	return err
}

// Creates a new message to revoke the X509 certificate.
func revokeX509Cert(setup *TestSetup, accAddress sdk.AccAddress, subject, keyID, info string) error {
	revokeX509Cert := types.NewMsgRevokeX509Cert(accAddress.String(), subject, keyID, info)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)

	return err
}

// Creates an intermediate certificate, sets it as a unique certificate and adds it to the list of approved certificates.
func storeIntermediateCertificate(setup *TestSetup, accAddress sdk.AccAddress) types.Certificate {
	intermediateCertificate := intermediateCertificate(accAddress)
	intermediateCertificate.SerialNumber = SerialNumber
	setup.Keeper.SetUniqueCertificate(
		setup.Ctx,
		uniqueCertificate(intermediateCertificate.Issuer, intermediateCertificate.SerialNumber),
	)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, intermediateCertificate)

	return intermediateCertificate
}

// Proposes a new x509 root certificate, approves it with another trustee, and ensures the certificate has been approved.
func proposeAndApproveRootCertificate(setup *TestSetup, ownerTrustee sdk.AccAddress, options *rootCertOptions) {
	// ensure that `ownerTrustee` is trustee to eventually have enough approvals
	require.True(setup.T, setup.DclauthKeeper.HasRole(setup.Ctx, ownerTrustee, types.RootCertificateApprovalRole))

	// propose x509 root certificate by `ownerTrustee`
	_, err := proposeX509RootCert(setup, ownerTrustee, options.pemCert, options.info, options.vid)
	require.NoError(setup.T, err)

	// approve x509 root certificate by another trustee
	_, err = approveX509RootCert(setup, setup.Trustee2, options.subject, options.subjectKeyID, options.info)
	require.NoError(setup.T, err)

	// check that root certificate has been approved
	approvedCertificate, err := querySingleApprovedCertificate(
		setup, options.subject, options.subjectKeyID)
	require.NoError(setup.T, err)
	require.NotNil(setup.T, approvedCertificate)
}

func checkProposedCertificate(t *testing.T, proposeAddX509RootCert *types.MsgProposeAddX509RootCert, proposedCertificate *types.ProposedCertificate, trustee1, trustee2 string) {
	t.Helper()

	require.Equal(t, proposeAddX509RootCert.Cert, proposedCertificate.PemCert)
	require.Equal(t, proposeAddX509RootCert.Signer, proposedCertificate.Owner)
	require.Equal(t, testconstants.RootSubject, proposedCertificate.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, proposedCertificate.SubjectKeyId)
	require.Equal(t, testconstants.RootSerialNumber, proposedCertificate.SerialNumber)
	require.Equal(t, trustee1, proposedCertificate.Approvals[0].Address)
	require.Equal(t, testconstants.Info, proposedCertificate.Approvals[0].Info)
	require.Equal(t, trustee2, proposedCertificate.Rejects[0].Address)
	require.Equal(t, testconstants.Info, proposedCertificate.Rejects[0].Info)
}

func checkRejectedCertificate(t *testing.T, proposeAddX509RootCert *types.MsgProposeAddX509RootCert, rejectedCertificate *types.Certificate, trustee1, trustee2, trustee3 string) {
	t.Helper()

	proposedCertificate := &types.ProposedCertificate{
		PemCert:      rejectedCertificate.PemCert,
		Owner:        rejectedCertificate.Owner,
		Subject:      rejectedCertificate.Subject,
		SubjectKeyId: rejectedCertificate.SubjectKeyId,
		SerialNumber: rejectedCertificate.SerialNumber,
		Approvals:    rejectedCertificate.Approvals,
		Rejects:      rejectedCertificate.Rejects,
	}

	checkProposedCertificate(t, proposeAddX509RootCert, proposedCertificate, trustee1, trustee2)

	require.Equal(t, trustee3, rejectedCertificate.Rejects[1].Address)
	require.Equal(t, testconstants.Info, rejectedCertificate.Rejects[1].Info)
}

func TestHandler(t *testing.T) {
	// setup := Setup(t)

	t.Run("ProposeAddX509RootCert", func(t *testing.T) {
		t.Run("ByNotTrustee", func(t *testing.T) {
			setup := Setup(t)

			for _, role := range []dclauthtypes.AccountRole{
				dclauthtypes.Vendor,
				dclauthtypes.CertificationCenter,
				dclauthtypes.NodeAdmin,
			} {
				accAddress := GenerateAccAddress()
				setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, 1)

				// propose x509 root certificate
				_, err := proposeX509RootCert(setup, accAddress, testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
				require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
			}
		})

		t.Run("ByTrustee", func(t *testing.T) {
			setup := Setup(t)

			// propose x509 root certificate
			_, err := proposeX509RootCert(setup, setup.Trustee1, testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
			require.NoError(t, err)

			// reject x509 root certificate
			err = rejectX509RootCert(setup, setup.Trustee1, testconstants.Info)
			require.NoError(t, err)

			require.False(t, setup.Keeper.IsProposedCertificatePresent(setup.Ctx, testconstants.RootIssuer, testconstants.RootSerialNumber))

			// check that unique certificate key is registered
			require.False(t, setup.Keeper.IsUniqueCertificatePresent(
				setup.Ctx, testconstants.RootIssuer, testconstants.RootSerialNumber))
		})

		t.Run("ForInvalidCertificate", func(t *testing.T) {
			setup := Setup(t)

			// propose x509 root certificate
			_, err := proposeX509RootCert(setup, setup.Trustee1, testconstants.StubCertPem, testconstants.Info, testconstants.Vid)
			require.Error(t, err)
			require.True(t, pkitypes.ErrInvalidCertificate.Is(err))
		})

		t.Run("ForNonRootCertificate", func(t *testing.T) {
			setup := Setup(t)

			// propose x509 root certificate
			_, err := proposeX509RootCert(setup, setup.Trustee1, testconstants.LeafCertPem, testconstants.Info, testconstants.Vid)
			require.Error(t, err)
			require.True(t, pkitypes.ErrInappropriateCertificateType.Is(err))
		})

		t.Run("ProposedCertificateAlreadyExists", func(t *testing.T) {
			setup := Setup(t)

			// propose adding of x509 root certificate
			_, err := proposeX509RootCert(setup, setup.Trustee1, testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
			require.NoError(t, err)

			// store another account
			anotherAccount := AddAccountWithRoleAndPermissions(setup, dclauthtypes.Trustee, 1)

			// propose adding of the same x509 root certificate again
			_, err = proposeX509RootCert(setup, anotherAccount, testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
			require.Error(t, err)
			require.True(t, pkitypes.ErrProposedCertificateAlreadyExists.Is(err))
		})

		t.Run("CertificateAlreadyExists", func(t *testing.T) {
			setup := Setup(t)

			// store x509 root certificate
			rootCertificate := rootCertificate(testconstants.Address1)
			setup.Keeper.SetUniqueCertificate(
				setup.Ctx,
				uniqueCertificate(rootCertificate.Subject, rootCertificate.SerialNumber),
			)
			setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

			// propose adding of the same x509 root certificate
			_, err := proposeX509RootCert(setup, setup.Trustee1, testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
			require.Error(t, err)
			require.True(t, pkitypes.ErrCertificateAlreadyExists.Is(err))
		})

		t.Run("ForDifferentSerialNumber", func(t *testing.T) {
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
			_, err := proposeX509RootCert(setup, setup.Trustee1, testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
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
		})

		t.Run("ForDifferentSerialNumberDifferentSigner", func(t *testing.T) {
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
			_, err := proposeX509RootCert(setup, setup.Trustee1, testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
			require.Error(t, err)
			require.True(t, sdkerrors.ErrUnauthorized.Is(err))
		})
	})

	t.Run("ProposeAddAndRejectX509RootCert", func(t *testing.T) {
		t.Run("ByTrustee", func(t *testing.T) {
			setup := Setup(t)

			// propose x509 root certificate
			_, err := proposeX509RootCert(setup, setup.Trustee1, testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
			require.NoError(t, err)

			// reject x509 root certificate
			err = rejectX509RootCert(setup, setup.Trustee1, testconstants.Info)
			require.NoError(t, err)

			require.False(t, setup.Keeper.IsProposedCertificatePresent(setup.Ctx, testconstants.RootIssuer, testconstants.RootSerialNumber))

			// check that unique certificate key is registered
			require.False(t, setup.Keeper.IsUniqueCertificatePresent(
				setup.Ctx, testconstants.RootIssuer, testconstants.RootSerialNumber))
		})

		t.Run("ByAnotherTrustee", func(t *testing.T) {
			setup := Setup(t)

			// propose x509 root certificate
			proposeAddX509RootCert, err := proposeX509RootCert(setup, setup.Trustee1, testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
			require.NoError(t, err)

			// reject x509 root certificate
			err = rejectX509RootCert(setup, setup.Trustee2, testconstants.Info)
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
		})

		t.Run("WithApproval_ByTrustee", func(t *testing.T) {
			setup := Setup(t)

			_ = AddAccountWithRoleAndPermissions(setup, dclauthtypes.Trustee, 1)

			// propose x509 root certificate
			proposeAddX509RootCert, err := proposeX509RootCert(setup, setup.Trustee1, testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
			require.NoError(t, err)

			// approve
			_, err = approveX509RootCert(setup, setup.Trustee2, testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
			require.NoError(t, err)

			// reject x509 root certificate
			err = rejectX509RootCert(setup, setup.Trustee1, testconstants.Info)
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
		})
	})

	t.Run("ApproveAddX509RootCert", func(t *testing.T) {
		t.Run("ForNotEnoughApprovals", func(t *testing.T) {
			setup := Setup(t)

			// store account without trustee role
			nonTrustee := AddAccountWithRoleAndPermissions(setup, dclauthtypes.Trustee, 1)

			// propose x509 root certificate by account without trustee role
			proposeAddX509RootCert, err := proposeX509RootCert(setup, nonTrustee, testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
			require.NoError(t, err)

			// approve
			_, err = approveX509RootCert(setup, setup.Trustee1, testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
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
			_, err = approveX509RootCert(setup, setup.Trustee2, testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
			require.NoError(t, err)

			// query approved certificate and we should get one back
			approvedCertificate, _ := querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
			require.Equal(t, testconstants.RootIssuer, approvedCertificate.Subject)
			require.Equal(t, testconstants.RootSerialNumber, approvedCertificate.SerialNumber)
			require.True(t, approvedCertificate.IsRoot)
			require.True(t, approvedCertificate.HasApprovalFrom(setup.Trustee1.String()))
			require.True(t, approvedCertificate.HasApprovalFrom(setup.Trustee2.String()))
		})

		t.Run("ForEnoughApprovals", func(t *testing.T) {
			setup := Setup(t)

			// propose add x509 root certificate by trustee
			proposeAddX509RootCert, err := proposeX509RootCert(setup, setup.Trustee1, testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
			require.NoError(t, err)

			// approve by second trustee
			_, err = approveX509RootCert(setup, setup.Trustee2, testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
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
		})

		t.Run("ForUnknownProposedCertificate", func(t *testing.T) {
			setup := Setup(t)

			// approve
			_, err := approveX509RootCert(setup, setup.Trustee1, testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
			require.Error(t, err)
			require.True(t, pkitypes.ErrProposedCertificateDoesNotExist.Is(err))
		})

		t.Run("ByNotTrustee", func(t *testing.T) {
			setup := Setup(t)

			// propose add x509 root certificate
			_, err := proposeX509RootCert(setup, setup.Trustee1, testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
			require.NoError(t, err)

			for _, role := range []dclauthtypes.AccountRole{
				dclauthtypes.Vendor,
				dclauthtypes.CertificationCenter,
				dclauthtypes.NodeAdmin,
			} {
				accAddress := AddAccountWithRoleAndPermissions(setup, role, 1)

				// approve
				_, err = approveX509RootCert(setup, accAddress, testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
				require.Error(t, err)
				require.True(t, sdkerrors.ErrUnauthorized.Is(err))
			}
		})

		t.Run("Twice", func(t *testing.T) {
			setup := Setup(t)

			// store account without Trustee role
			accAddress := AddAccountWithRoleAndPermissions(setup, dclauthtypes.Trustee, 1)

			// propose add x509 root certificate
			_, err := proposeX509RootCert(setup, accAddress, testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
			require.NoError(t, err)

			// approve
			approveAddX509RootCert, err := approveX509RootCert(setup, setup.Trustee1, testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
			require.NoError(t, err)

			// approve second time
			_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
			require.Error(t, err)
			require.True(t, sdkerrors.ErrUnauthorized.Is(err))
		})

		t.Run("FourApprovalsAreNeeded_FiveTrustees", func(t *testing.T) { // trasport
			setup := Setup(t)

			// we have 5 trustees: 1 approval comes from propose => we need 3 more approvals

			// store 4th trustee
			fourthTrustee := AddAccountWithRoleAndPermissions(setup, dclauthtypes.Trustee, 1)

			// store 5th trustee
			fifthTrustee := AddAccountWithRoleAndPermissions(setup, dclauthtypes.Trustee, 1)

			// propose x509 root certificate by account Trustee1
			proposeAddX509RootCert, err := proposeX509RootCert(setup, setup.Trustee1, testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
			require.NoError(t, err)

			// approve x509 root certificate by account Trustee2
			_, err = approveX509RootCert(setup, setup.Trustee2, testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
			require.NoError(t, err)

			// approve x509 root certificate by account Trustee3
			_, err = approveX509RootCert(setup, setup.Trustee3, testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
			require.NoError(t, err)

			// reject x509 root certificate by account Trustee4
			err = rejectX509RootCert(setup, fourthTrustee, testconstants.Info)
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
			_, err = approveX509RootCert(setup, fifthTrustee, testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
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
		})
	})

	t.Run("AddX509Cert", func(t *testing.T) {
		t.Run("AddX509Cert", func(t *testing.T) {
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
				accAddress := AddAccountWithRoleAndPermissions(setup, role, 1)

				// add x509 certificate
				addX509Cert, err := addX509Cert(setup, accAddress, testconstants.IntermediateCertPem)
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
		})

		t.Run("ForInvalidCertificate", func(t *testing.T) {
			setup := Setup(t)

			// add x509 certificate
			_, err := addX509Cert(setup, setup.Trustee1, testconstants.StubCertPem)
			require.Error(t, err)
			require.True(t, pkitypes.ErrInvalidCertificate.Is(err))
		})

		t.Run("ForRootCertificate", func(t *testing.T) {
			setup := Setup(t)

			// add root certificate as leaf x509 certificate
			_, err := addX509Cert(setup, setup.Trustee1, testconstants.RootCertPem)
			require.Error(t, err)
			require.True(t, pkitypes.ErrInappropriateCertificateType.Is(err))
		})

		t.Run("ForDuplicate", func(t *testing.T) {
			setup := Setup(t)

			// store root certificate
			rootCertificate := rootCertificate(setup.Trustee1)
			setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

			// store intermediate certificate
			addX509Cert, err := addX509Cert(setup, setup.Trustee1, testconstants.IntermediateCertPem)
			require.NoError(t, err)

			// store intermediate certificate second time
			_, err = setup.Handler(setup.Ctx, addX509Cert)
			require.Error(t, err)
			require.True(t, pkitypes.ErrCertificateAlreadyExists.Is(err))
		})

		t.Run("ForDifferentSerialNumber", func(t *testing.T) {
			setup := Setup(t)

			// store root certificate
			rootCertificate := rootCertificate(setup.Trustee1)
			setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

			// store intermediate certificate with different serial number
			_ = storeIntermediateCertificate(setup, setup.Trustee1)

			// store intermediate certificate second time
			addX509Cert, err := addX509Cert(setup, setup.Trustee1, testconstants.IntermediateCertPem)
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
		})

		t.Run("ForDifferentSerialNumberDifferentSigner", func(t *testing.T) {
			setup := Setup(t)

			// store root certificate
			rootCertificate := rootCertificate(testconstants.Address1)
			setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

			// store intermediate certificate with different serial number
			_ = storeIntermediateCertificate(setup, testconstants.Address1)

			// store intermediate certificate second time
			_, err := addX509Cert(setup, setup.Trustee1, testconstants.IntermediateCertPem)
			require.Error(t, err)
			require.True(t, sdkerrors.ErrUnauthorized.Is(err))
		})

		t.Run("ForAbsentDirectParentCert", func(t *testing.T) {
			setup := Setup(t)

			// add intermediate x509 certificate
			_, err := addX509Cert(setup, setup.Trustee1, testconstants.IntermediateCertPem)
			require.Error(t, err)
			require.True(t, pkitypes.ErrInvalidCertificate.Is(err))
		})

		t.Run("ForNoRootCert", func(t *testing.T) {
			setup := Setup(t)

			// add intermediate certificate
			intermediateCertificate := intermediateCertificate(setup.Trustee1)
			setup.Keeper.AddApprovedCertificate(setup.Ctx, intermediateCertificate)

			// add leaf x509 certificate
			_, err := addX509Cert(setup, setup.Trustee1, testconstants.LeafCertPem)
			require.Error(t, err)
			require.True(t, pkitypes.ErrInvalidCertificate.Is(err))
		})

		t.Run("ForFailedCertificateVerification", func(t *testing.T) {
			setup := Setup(t)

			// add invalid root
			invalidRootCertificate := types.NewRootCertificate(testconstants.StubCertPem,
				testconstants.RootSubject, testconstants.RootSubjectAsText, testconstants.RootSubjectKeyID,
				testconstants.RootSerialNumber, setup.Trustee1.String(), []*types.Grant{}, []*types.Grant{}, testconstants.Vid)
			setup.Keeper.AddApprovedCertificate(setup.Ctx, invalidRootCertificate)

			// add intermediate x509 certificate
			_, err := addX509Cert(setup, setup.Trustee1, testconstants.IntermediateCertPem)
			require.Error(t, err)
			require.True(t, pkitypes.ErrInvalidCertificate.Is(err))
		})

		t.Run("ForTree", func(t *testing.T) {
			setup := Setup(t)

			// add root x509 certificate
			rootCertOptions := createTestRootCertOptions()
			proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

			// add intermediate x509 certificate
			_, err := addX509Cert(setup, setup.Trustee1, testconstants.IntermediateCertPem)
			require.NoError(t, err)

			// add leaf x509 certificate
			_, err = addX509Cert(setup, setup.Trustee1, testconstants.LeafCertPem)
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
		})

		//nolint:funlen
		t.Run("EachChildCertRefersToTwoParentCerts", func(t *testing.T) {
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
			intermediateCertificate := storeIntermediateCertificate(setup, setup.Trustee1)

			childCertID := certificateIdentifier(intermediateCertificate.Subject, intermediateCertificate.SubjectKeyId)
			rootChildCertificates := types.ChildCertificates{
				Issuer:         intermediateCertificate.Issuer,
				AuthorityKeyId: intermediateCertificate.AuthorityKeyId,
				CertIds:        []*types.CertificateIdentifier{&childCertID},
			}
			setup.Keeper.SetChildCertificates(setup.Ctx, rootChildCertificates)

			// store second intermediate certificate (it refers to two parent certificates)
			_, err := addX509Cert(setup, setup.Trustee1, testconstants.IntermediateCertPem)
			require.NoError(t, err)

			// store leaf certificate (it refers to two parent certificates)
			_, err = addX509Cert(setup, setup.Trustee1, testconstants.LeafCertPem)
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
		})
	})

	t.Run("TwoThirdApprovalsNeeded", func(t *testing.T) {
		t.Run("ForAddingRootCertification", func(t *testing.T) {
			setup := Setup(t)

			// propose x509 root certificate by account without trustee role
			_, err := proposeX509RootCert(setup, setup.Trustee1, testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
			require.NoError(t, err)

			// Create an array of trustee account from 1 to 50
			trusteeAccounts := make([]sdk.AccAddress, 50)

			totalAdditionalTrustees := rand.Intn(50)
			for i := 0; i < totalAdditionalTrustees; i++ {
				trusteeAccounts[i] = AddAccountWithRoleAndPermissions(setup, dclauthtypes.Trustee, 1)
			}

			// We have 3 Trustees in test setup.
			twoThirds := int(math.Ceil(types.RootCertificateApprovalsPercent * float64(3+totalAdditionalTrustees)))

			// Until we hit 2/3 of the total number of Trustees, we should not be able to approve the certificate
			for i := 1; i < twoThirds-1; i++ {
				_, err = approveX509RootCert(setup, trusteeAccounts[i], testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
				require.NoError(t, err)

				_, err = querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
				require.Error(t, err)
				require.Equal(t, codes.NotFound, status.Code(err))
			}

			// One more approval will move this to approved state from pending
			_, err = approveX509RootCert(setup, setup.Trustee2, testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
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
		})

		t.Run("ForRevokingRootCertification", func(t *testing.T) {
			setup := Setup(t)

			// propose x509 root certificate by account without trustee role
			_, err := proposeX509RootCert(setup, setup.Trustee1, testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
			require.NoError(t, err)

			// Approve the certificate from Trustee2
			_, err = approveX509RootCert(setup, setup.Trustee2, testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
			require.NoError(t, err)

			// Check that the certificate is approved
			approvedCertificate, _ := querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
			require.Equal(t, testconstants.RootIssuer, approvedCertificate.Subject)
			require.Equal(t, testconstants.RootSerialNumber, approvedCertificate.SerialNumber)
			require.True(t, approvedCertificate.IsRoot)
			require.True(t, approvedCertificate.HasApprovalFrom(setup.Trustee1.String()))

			// Create an array of trustee account from 1 to 50
			trusteeAccounts := make([]sdk.AccAddress, 50)

			totalAdditionalTrustees := rand.Intn(50)
			for i := 0; i < totalAdditionalTrustees; i++ {
				trusteeAccounts[i] = AddAccountWithRoleAndPermissions(setup, dclauthtypes.Trustee, 1)
			}

			// We have 3 Trustees in test setup.
			twoThirds := int(math.Ceil(types.RootCertificateApprovalsPercent * float64(3+totalAdditionalTrustees)))

			// Trustee1 proposes to revoke the certificate
			err = proposeRevokeX509RootCert(setup, setup.Trustee1, testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
			require.NoError(t, err)

			// Until we hit 2/3 of the total number of Trustees, we should not be able to revoke the certificate
			// We start the counter from 2 as the proposer is a trustee as well
			for i := 1; i < twoThirds-1; i++ {
				// approve the revocation
				err = approveRevokeX509RootCert(setup, trusteeAccounts[i], testconstants.Info)
				require.NoError(t, err)

				// check that the certificate is still not revoked
				approvedCertificate, _ := querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
				require.Equal(t, testconstants.RootIssuer, approvedCertificate.Subject)
				require.Equal(t, testconstants.RootSerialNumber, approvedCertificate.SerialNumber)
				require.True(t, approvedCertificate.IsRoot)
			}

			// One more revoke will revoke the certificate
			err = approveRevokeX509RootCert(setup, setup.Trustee2, testconstants.Info)
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
		})
	})

	t.Run("ProposeRevokeX509RootCert", func(t *testing.T) {
		t.Run("ByTrusteeOwner", func(t *testing.T) {
			setup := Setup(t)

			// propose x509 root certificate by `setup.Trustee` and approve by another trustee
			rootCertOptions := createTestRootCertOptions()
			proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

			// propose revocation of x509 root certificate by `setup.Trustee`
			err := proposeRevokeX509RootCert(setup, setup.Trustee1, testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
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
		})

		t.Run("ByTrusteeNotOwner", func(t *testing.T) {
			setup := Setup(t)

			// propose x509 root certificate by `setup.Trustee` and approve by another trustee
			rootCertOptions := createTestRootCertOptions()
			proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

			// store another trustee
			anotherTrustee := AddAccountWithRoleAndPermissions(setup, dclauthtypes.Trustee, 1)

			// propose revocation of x509 root certificate by new trustee
			err := proposeRevokeX509RootCert(setup, anotherTrustee, testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
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
		})

		t.Run("ByNotTrustee", func(t *testing.T) {
			setup := Setup(t)

			// propose and approve x509 root certificate
			rootCertOptions := createTestRootCertOptions()
			proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

			for _, role := range []dclauthtypes.AccountRole{
				dclauthtypes.Vendor,
				dclauthtypes.CertificationCenter,
				dclauthtypes.NodeAdmin,
			} {
				accAddress := AddAccountWithRoleAndPermissions(setup, role, 1)

				// propose revocation of x509 root certificate
				err := proposeRevokeX509RootCert(setup, accAddress, testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
				require.Error(t, err)
				require.True(t, sdkerrors.ErrUnauthorized.Is(err))
			}
		})

		t.Run("CertificateDoesNotExist", func(t *testing.T) {
			setup := Setup(t)

			// propose revocation of not existing certificate
			err := proposeRevokeX509RootCert(setup, setup.Trustee1, testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
			require.Error(t, err)
			require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
		})

		t.Run("ForProposedCertificate", func(t *testing.T) {
			setup := Setup(t)

			// propose x509 root certificate
			_, err := proposeX509RootCert(setup, setup.Trustee1, testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
			require.NoError(t, err)

			// check that proposed certificate is present
			proposedCertificate, _ := queryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
			require.NotNil(t, proposedCertificate)

			// propose revocation of proposed root certificate
			err = proposeRevokeX509RootCert(setup, setup.Trustee1, testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
			require.Error(t, err)
			require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
		})

		t.Run("ProposedRevocationAlreadyExists", func(t *testing.T) {
			setup := Setup(t)

			// propose and approve x509 root certificate
			rootCertOptions := createTestRootCertOptions()
			proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

			// propose revocation of x509 root certificate
			err := proposeRevokeX509RootCert(setup, setup.Trustee1, testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
			require.NoError(t, err)

			// store another trustee
			anotherTrustee := AddAccountWithRoleAndPermissions(setup, dclauthtypes.Trustee, 1)

			// propose revocation of the same x509 root certificate again
			err = proposeRevokeX509RootCert(setup, anotherTrustee, testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
			require.Error(t, err)
			require.True(t, pkitypes.ErrProposedCertificateRevocationAlreadyExists.Is(err))
		})

		t.Run("ForNonRootCertificate", func(t *testing.T) {
			setup := Setup(t)

			// store x509 root certificate
			rootCertificate := rootCertificate(setup.Trustee1)
			setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

			// store x509 intermediate certificate
			_, err := addX509Cert(setup, setup.Trustee1, testconstants.IntermediateCertPem)
			require.NoError(t, err)

			// propose revocation of x509 intermediate certificate
			err = proposeRevokeX509RootCert(setup, setup.Trustee1, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID, testconstants.Info)
			require.Error(t, err)
			require.True(t, pkitypes.ErrInappropriateCertificateType.Is(err))
		})
	})

	t.Run("ApproveRevokeX509RootCert", func(t *testing.T) {
		t.Run("ForNotEnoughApprovals", func(t *testing.T) {
			setup := Setup(t)

			// propose and approve x509 root certificate
			rootCertOptions := createTestRootCertOptions()
			proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

			// Add 1 more trustee (this will bring the total trustee's to 4)
			_ = AddAccountWithRoleAndPermissions(setup, dclauthtypes.Trustee, 1)

			// propose revocation of x509 root certificate
			err := proposeRevokeX509RootCert(setup, setup.Trustee1, testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
			require.NoError(t, err)

			// approve
			err = approveRevokeX509RootCert(setup, setup.Trustee2, testconstants.Info)
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
		})

		t.Run("ForEnoughApprovals", func(t *testing.T) {
			setup := Setup(t)

			// propose and approve x509 root certificate
			rootCertOptions := createTestRootCertOptions()
			proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

			// propose revocation of x509 root certificate
			err := proposeRevokeX509RootCert(setup, setup.Trustee1, testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
			require.NoError(t, err)

			// get certificate for further comparison
			certificateBeforeRevocation, _ := querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
			require.NotNil(t, certificateBeforeRevocation)

			// approve
			err = approveRevokeX509RootCert(setup, setup.Trustee2, testconstants.Info)
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
		})

		t.Run("ByNotTrustee", func(t *testing.T) {
			setup := Setup(t)

			// propose and approve x509 root certificate
			rootCertOptions := createTestRootCertOptions()
			proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

			// propose revocation of x509 root certificate
			err := proposeRevokeX509RootCert(setup, setup.Trustee1, testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
			require.NoError(t, err)

			for _, role := range []dclauthtypes.AccountRole{
				dclauthtypes.Vendor,
				dclauthtypes.CertificationCenter,
				dclauthtypes.NodeAdmin,
			} {
				accAddress := AddAccountWithRoleAndPermissions(setup, role, 1)

				// approve
				err = approveRevokeX509RootCert(setup, accAddress, testconstants.Info)
				require.Error(t, err)
				require.True(t, sdkerrors.ErrUnauthorized.Is(err))
			}
		})

		t.Run("ProposedRevocationDoesNotExist", func(t *testing.T) {
			setup := Setup(t)

			// propose and approve x509 root certificate
			rootCertOptions := createTestRootCertOptions()
			proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

			// approve revocation of x509 root certificate
			err := approveRevokeX509RootCert(setup, setup.Trustee1, testconstants.Info)
			require.Error(t, err)
			require.True(t, pkitypes.ErrProposedCertificateRevocationDoesNotExist.Is(err))
		})

		t.Run("Twice", func(t *testing.T) {
			setup := Setup(t)

			// propose and approve x509 root certificate
			rootCertOptions := createTestRootCertOptions()
			proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

			// propose revocation of x509 root certificate
			err := proposeRevokeX509RootCert(setup, setup.Trustee1, testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
			require.NoError(t, err)

			// approve revocation by the same trustee
			err = approveRevokeX509RootCert(setup, setup.Trustee1, testconstants.Info)
			require.Error(t, err)
			require.True(t, sdkerrors.ErrUnauthorized.Is(err))
		})

		//nolint:funlen
		t.Run("ForTree", func(t *testing.T) {
			setup := Setup(t)

			// add root x509 certificate
			rootCertOptions := createTestRootCertOptions()
			proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)
			// add intermediate x509 certificate
			_, err := addX509Cert(setup, setup.Trustee1, testconstants.IntermediateCertPem)
			require.NoError(t, err)

			// add leaf x509 certificate
			_, err = addX509Cert(setup, setup.Trustee1, testconstants.LeafCertPem)
			require.NoError(t, err)

			// propose revocation of x509 root certificate
			err = proposeRevokeX509RootCert(setup, setup.Trustee1, testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
			require.NoError(t, err)

			// approve
			err = approveRevokeX509RootCert(setup, setup.Trustee2, testconstants.Info)
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
		})
	})

	t.Run("RevokeX509Cert", func(t *testing.T) {
		t.Run("RevokeX509Cert", func(t *testing.T) {
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
				accAddress := AddAccountWithRoleAndPermissions(setup, role, 1)

				// add x509 certificate
				_, err := addX509Cert(setup, accAddress, testconstants.IntermediateCertPem)
				require.NoError(t, err)

				// get certificate for further comparison
				certificateBeforeRevocation, _ := querySingleApprovedCertificate(
					setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
				require.NotNil(t, certificateBeforeRevocation)

				// revoke x509 certificate
				err = revokeX509Cert(setup, accAddress, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID, testconstants.Info)
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
		})

		t.Run("CertificateDoesNotExist", func(t *testing.T) {
			setup := Setup(t)

			// revoke x509 certificate
			err := revokeX509Cert(setup, setup.Trustee1, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID, testconstants.Info)
			require.Error(t, err)
			require.True(t, pkitypes.ErrCertificateDoesNotExist.Is(err))
		})

		t.Run("ForRootCertificate", func(t *testing.T) {
			setup := Setup(t)

			// propose and approve x509 root certificate
			rootCertOptions := createTestRootCertOptions()
			proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

			// revoke x509 root certificate
			err := revokeX509Cert(setup, setup.Trustee1, testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
			require.Error(t, err)
			require.True(t, pkitypes.ErrInappropriateCertificateType.Is(err))
		})

		t.Run("ByNotOwner", func(t *testing.T) {
			setup := Setup(t)

			// store root certificate
			rootCertificate := rootCertificate(setup.Trustee1)
			setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

			// add x509 certificate by `setup.Trustee`
			_, err := addX509Cert(setup, setup.Trustee1, testconstants.IntermediateCertPem)
			require.NoError(t, err)

			// store another account
			anotherTrustee := AddAccountWithRoleAndPermissions(setup, dclauthtypes.Trustee, 1)

			// revoke x509 certificate by another account
			err = revokeX509Cert(setup, anotherTrustee, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID, testconstants.Info)
			require.Error(t, err)
			require.True(t, sdkerrors.ErrUnauthorized.Is(err))
		})

		t.Run("ForTree", func(t *testing.T) {
			setup := Setup(t)

			// add root x509 certificate
			rootCertOptions := createTestRootCertOptions()
			proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

			// add intermediate x509 certificate
			_, err := addX509Cert(setup, setup.Trustee1, testconstants.IntermediateCertPem)
			require.NoError(t, err)

			// add leaf x509 certificate
			_, err = addX509Cert(setup, setup.Trustee1, testconstants.LeafCertPem)
			require.NoError(t, err)

			// revoke x509 certificate
			err = revokeX509Cert(setup, setup.Trustee1, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID, testconstants.Info)
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
		})
	})

	t.Run("RejectX509RootCert", func(t *testing.T) {
		t.Run("TwoRejectApprovalsAreNeeded", func(t *testing.T) {
			setup := Setup(t)

			// propose x509 root certificate by account Trustee1
			proposeAddX509RootCert, err := proposeX509RootCert(setup, setup.Trustee1, testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
			require.NoError(t, err)

			// reject x509 root certificate by account Trustee2
			err = rejectX509RootCert(setup, setup.Trustee2, testconstants.Info)
			require.NoError(t, err)

			// certificate should be in the entity <Proposed X509 Root Certificate>, because we haven't enough reject approvals
			proposedCertificate, err := queryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
			require.NoError(t, err)

			// check proposed certificate
			checkProposedCertificate(t, proposeAddX509RootCert, proposedCertificate, setup.Trustee1.String(), setup.Trustee2.String())

			// reject x509 root certificate by account Trustee3
			err = rejectX509RootCert(setup, setup.Trustee3, testconstants.Info)
			require.NoError(t, err)

			// certificate should not be in the entity <Proposed X509 Root Certificate>, because we have enough reject approvals
			_, err = queryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
			require.Error(t, err)

			// certificate should be in the entity <Rejected X509 Root Certificate>, because we have enough rejected approvals
			rejectedCertificate, err := queryRejectedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
			require.NoError(t, err)

			// check rejected certificate
			checkRejectedCertificate(t, proposeAddX509RootCert, rejectedCertificate, setup.Trustee1.String(), setup.Trustee2.String(), setup.Trustee3.String())
		})

		t.Run("ByNotTrustee", func(t *testing.T) {
			setup := Setup(t)

			// propose add x509 root certificate
			_, err := proposeX509RootCert(setup, setup.Trustee1, testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
			require.NoError(t, err)

			for _, role := range []dclauthtypes.AccountRole{
				dclauthtypes.Vendor,
				dclauthtypes.CertificationCenter,
				dclauthtypes.NodeAdmin,
			} {
				accAddress := AddAccountWithRoleAndPermissions(setup, role, 1)

				// reject x509 root certificate
				err = rejectX509RootCert(setup, accAddress, testconstants.Info)
				require.Error(t, err)
				require.True(t, sdkerrors.ErrUnauthorized.Is(err))
			}
		})

		t.Run("FromTheSameTrustee", func(t *testing.T) {
			setup := Setup(t)

			// propose add x509 root certificate
			_, err := proposeX509RootCert(setup, setup.Trustee1, testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
			require.NoError(t, err)

			// reject x509 root certificate by account Trustee2
			err = rejectX509RootCert(setup, setup.Trustee2, testconstants.Info)
			require.NoError(t, err)

			// second time reject x509 root certificate by account Trustee2
			err = rejectX509RootCert(setup, setup.Trustee2, testconstants.Info)
			require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
		})

		t.Run("DoubleTime", func(t *testing.T) {
			setup := Setup(t)

			// propose x509 root certificate by account Trustee1
			proposeAddX509RootCert, err := proposeX509RootCert(setup, setup.Trustee1, testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
			require.NoError(t, err)

			// reject x509 root certificate by account Trustee2
			err = rejectX509RootCert(setup, setup.Trustee2, testconstants.Info)
			require.NoError(t, err)

			// certificate should be in the entity <Proposed X509 Root Certificate>, because we haven't enough reject approvals
			proposedCertificate, err := queryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
			require.NoError(t, err)

			// check proposed certificate
			checkProposedCertificate(t, proposeAddX509RootCert, proposedCertificate, setup.Trustee1.String(), setup.Trustee2.String())

			// reject x509 root certificate by account Trustee3
			err = rejectX509RootCert(setup, setup.Trustee3, testconstants.Info)
			require.NoError(t, err)

			// certificate should not be in the entity <Proposed X509 Root Certificate>, because we have enough reject approvals
			_, err = queryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
			require.Error(t, err)

			// certificate should be in the entity <Rejected X509 Root Certificate>, because we have enough rejected approvals
			rejectedCertificate, err := queryRejectedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
			require.NoError(t, err)

			// check rejected certificate
			checkRejectedCertificate(t, proposeAddX509RootCert, rejectedCertificate, setup.Trustee1.String(), setup.Trustee2.String(), setup.Trustee3.String())

			// propose x509 root certificate by account Trustee1
			proposeAddX509RootCert, err = proposeX509RootCert(setup, setup.Trustee1, testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
			require.NoError(t, err)

			// certificate should be in the entity <Proposed X509 Root Certificate>, because we haven't enough reject approvals
			_, err = queryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
			require.NoError(t, err)

			// certificate should not be in the entity <Rejected X509 Root Certificate>, because we have propose that certificate
			_, err = queryRejectedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
			require.Error(t, err)

			// reject x509 root certificate by account Trustee3
			err = rejectX509RootCert(setup, setup.Trustee3, testconstants.Info)
			require.NoError(t, err)

			// reject x509 root certificate by account Trustee2
			err = rejectX509RootCert(setup, setup.Trustee2, testconstants.Info)
			require.NoError(t, err)

			// certificate should be in the entity <Rejected X509 Root Certificate>, because we have enough rejected approvals
			rejectedCertificate, err = queryRejectedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
			require.NoError(t, err)

			// check rejected certificate
			checkRejectedCertificate(t, proposeAddX509RootCert, rejectedCertificate, setup.Trustee1.String(), setup.Trustee3.String(), setup.Trustee2.String())
		})

		t.Run("TwoRejectApprovalsAreNeeded_FiveTrustees", func(t *testing.T) {
			setup := Setup(t)

			// we have 5 trustees: 1 approval comes from propose => we need 2 rejects to make certificate rejected

			// store 4th trustee
			_ = AddAccountWithRoleAndPermissions(setup, dclauthtypes.Trustee, 1)

			// store 5th trustee
			_ = AddAccountWithRoleAndPermissions(setup, dclauthtypes.Trustee, 1)

			// propose x509 root certificate by account Trustee1
			proposeAddX509RootCert, err := proposeX509RootCert(setup, setup.Trustee1, testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
			require.NoError(t, err)

			// reject x509 root certificate by account Trustee2
			err = rejectX509RootCert(setup, setup.Trustee2, testconstants.Info)
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
			err = rejectX509RootCert(setup, setup.Trustee3, testconstants.Info)
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
		})
	})

	t.Run("ApproveX509RootCertAndRejectX509RootCert", func(t *testing.T) {
		t.Run("FromTheSameTrustee", func(t *testing.T) {
			setup := Setup(t)
			// propose add x509 root certificate
			_, err := proposeX509RootCert(setup, setup.Trustee1, testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
			require.NoError(t, err)

			for _, role := range []dclauthtypes.AccountRole{
				dclauthtypes.Trustee,
			} {
				_ = AddAccountWithRoleAndPermissions(setup, role, 1)

				// approve x509 root certificate by account Trustee2
				_, err = approveX509RootCert(setup, setup.Trustee2, testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
				require.NoError(t, err)

				pendingCert, _ := setup.Keeper.GetProposedCertificate(setup.Ctx, testconstants.RootSubject, testconstants.RootSubjectKeyID)
				prevRejectsLen := len(pendingCert.Rejects)
				prevApprovalsLen := len(pendingCert.Approvals)
				// reject x509 root certificate by account Trustee2
				err = rejectX509RootCert(setup, setup.Trustee2, testconstants.Info)
				require.NoError(t, err)

				pendingCert, found := setup.Keeper.GetProposedCertificate(setup.Ctx, testconstants.RootSubject, testconstants.RootSubjectKeyID)
				require.True(t, found)
				require.Equal(t, len(pendingCert.Rejects), prevRejectsLen+1)
				require.Equal(t, len(pendingCert.Approvals), prevApprovalsLen-1)
			}
		})
	})

	t.Run("RejectX509RootCertAndApproveX509RootCert", func(t *testing.T) {
		t.Run("FromTheSameTrustee", func(t *testing.T) {
			setup := Setup(t)
			// propose add x509 root certificate
			_, err := proposeX509RootCert(setup, setup.Trustee1, testconstants.RootCertPem, testconstants.Info, testconstants.Vid)
			require.NoError(t, err)

			for _, role := range []dclauthtypes.AccountRole{
				dclauthtypes.Trustee,
			} {
				_ = AddAccountWithRoleAndPermissions(setup, role, 1)

				// reject x509 root certificate by account Trustee2
				err = rejectX509RootCert(setup, setup.Trustee2, testconstants.Info)
				require.NoError(t, err)

				pendingCert, _ := setup.Keeper.GetProposedCertificate(setup.Ctx, testconstants.RootSubject, testconstants.RootSubjectKeyID)
				prevRejectsLen := len(pendingCert.Rejects)
				prevApprovalsLen := len(pendingCert.Approvals)
				// approve x509 root certificate by account Trustee2
				_, err = approveX509RootCert(setup, setup.Trustee2, testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Info)
				require.NoError(t, err)

				pendingCert, found := setup.Keeper.GetProposedCertificate(setup.Ctx, testconstants.RootSubject, testconstants.RootSubjectKeyID)
				require.True(t, found)
				require.Equal(t, len(pendingCert.Rejects), prevRejectsLen-1)
				require.Equal(t, len(pendingCert.Approvals), prevApprovalsLen+1)
			}
		})
	})

	t.Run("AssignVid", func(t *testing.T) {
		t.Run("SenderNotVendorAdmin", func(t *testing.T) {
			setup := Setup(t)

			assignVid := types.MsgAssignVid{
				Signer:       setup.Trustee1.String(),
				Subject:      testconstants.TestSubject,
				SubjectKeyId: testconstants.TestSubjectKeyID,
				Vid:          testconstants.TestCertPemVid,
			}

			_, err := setup.Handler(setup.Ctx, &assignVid)
			require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
		})

		t.Run("CertificateDoesNotExist", func(t *testing.T) {
			setup := Setup(t)

			vendorAcc := AddAccountWithRoleAndPermissions(setup, dclauthtypes.VendorAdmin, 0)

			assignVid := types.MsgAssignVid{
				Signer:       vendorAcc.String(),
				Subject:      testconstants.TestSubject,
				SubjectKeyId: testconstants.TestSubjectKeyID,
				Vid:          testconstants.TestCertPemVid,
			}

			_, err := setup.Handler(setup.Ctx, &assignVid)
			require.ErrorIs(t, err, pkitypes.ErrCertificateDoesNotExist)
		})

		t.Run("ForNonRootCertificate", func(t *testing.T) {
			setup := Setup(t)

			vendorAcc := AddAccountWithRoleAndPermissions(setup, dclauthtypes.VendorAdmin, 0)

			// propose and approve x509 root certificate
			rootCertOptions := createTestRootCertOptions()
			proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

			// add x509 intermediate certificate
			_, err := addX509Cert(setup, setup.Trustee1, testconstants.IntermediateCertPem)
			require.NoError(t, err)

			assignVid := types.MsgAssignVid{
				Signer:       vendorAcc.String(),
				Subject:      testconstants.IntermediateSubject,
				SubjectKeyId: testconstants.IntermediateSubjectKeyID,
				Vid:          testconstants.PAACertWithNumericVidVid,
			}

			_, err = setup.Handler(setup.Ctx, &assignVid)
			require.ErrorIs(t, err, pkitypes.ErrInappropriateCertificateType)
		})

		t.Run("CertificateAlreadyHasVid", func(t *testing.T) {
			setup := Setup(t)

			vendorAcc := AddAccountWithRoleAndPermissions(setup, dclauthtypes.VendorAdmin, 0)

			// propose and approve x509 root certificate
			rootCertOptions := createPAACertWithNumericVidOptions()
			proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

			assignVid := types.MsgAssignVid{
				Signer:       vendorAcc.String(),
				Subject:      rootCertOptions.subject,
				SubjectKeyId: rootCertOptions.subjectKeyID,
				Vid:          testconstants.PAACertWithNumericVidVid,
			}

			_, err := setup.Handler(setup.Ctx, &assignVid)
			require.ErrorIs(t, err, pkitypes.ErrNotEmptyVid)
		})

		t.Run("MessageVidAndCertificateVidNotEqual", func(t *testing.T) {
			setup := Setup(t)

			vendorAcc := AddAccountWithRoleAndPermissions(setup, dclauthtypes.VendorAdmin, 0)

			// propose and approve x509 root certificate
			rootCertOptions := createPAACertWithNumericVidOptions()
			rootCertOptions.vid = 0
			proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

			assignVid := types.MsgAssignVid{
				Signer:       vendorAcc.String(),
				Subject:      rootCertOptions.subject,
				SubjectKeyId: rootCertOptions.subjectKeyID,
				Vid:          1,
			}

			_, err := setup.Handler(setup.Ctx, &assignVid)
			require.ErrorIs(t, err, pkitypes.ErrCertificateVidNotEqualMsgVid)
		})

		t.Run("certificateWithoutSubjectVid", func(t *testing.T) {
			setup := Setup(t)

			vendorAcc := AddAccountWithRoleAndPermissions(setup, dclauthtypes.VendorAdmin, 0)

			// propose and approve x509 root certificate
			rootCertOptions := createTestRootCertOptions()
			rootCertOptions.vid = 0
			proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

			assignVid := types.MsgAssignVid{
				Signer:       vendorAcc.String(),
				Subject:      rootCertOptions.subject,
				SubjectKeyId: rootCertOptions.subjectKeyID,
				Vid:          testconstants.Vid,
			}

			_, err := setup.Handler(setup.Ctx, &assignVid)
			require.NoError(t, err)

			// query certificate
			certificates, _ := queryApprovedCertificates(setup, rootCertOptions.subject, rootCertOptions.subjectKeyID)

			// check
			require.Equal(t, len(certificates.Certs), 1)
			require.EqualValues(t, certificates.Certs[0].Vid, testconstants.Vid)
		})

		t.Run("certificateWithSubjectVid", func(t *testing.T) {
			setup := Setup(t)

			vendorAcc := AddAccountWithRoleAndPermissions(setup, dclauthtypes.VendorAdmin, 0)

			// propose and approve x509 root certificate
			rootCertOptions := createPAACertWithNumericVidOptions()
			rootCertOptions.vid = 0
			proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

			assignVid := types.MsgAssignVid{
				Signer:       vendorAcc.String(),
				Subject:      rootCertOptions.subject,
				SubjectKeyId: rootCertOptions.subjectKeyID,
				Vid:          testconstants.PAACertWithNumericVidVid,
			}

			_, err := setup.Handler(setup.Ctx, &assignVid)
			require.NoError(t, err)

			// query certificate
			certificates, _ := queryApprovedCertificates(setup, rootCertOptions.subject, rootCertOptions.subjectKeyID)

			// check
			require.Equal(t, len(certificates.Certs), 1)
			require.EqualValues(t, certificates.Certs[0].Vid, testconstants.PAACertWithNumericVidVid)
		})
	})

	t.Run("RevocationPointsByIssuerSubjectKeyID", func(t *testing.T) {
		setup := Setup(t)

		vendorAcc := AddAccountWithRoleAndPermissions(setup, dclauthtypes.Vendor, 65521)

		// propose x509 root certificate by account Trustee1
		_, err := proposeX509RootCert(setup, setup.Trustee1, testconstants.PAACertWithNumericVid, testconstants.Info, testconstants.Vid)
		require.NoError(t, err)

		// approve
		_, err = approveX509RootCert(setup, setup.Trustee2, testconstants.PAACertWithNumericVidSubject, testconstants.PAACertWithNumericVidSubjectKeyID, testconstants.Info)
		require.NoError(t, err)

		revocationPointBySubjectKeyID, isFound := setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
		require.False(t, isFound)
		require.Equal(t, len(revocationPointBySubjectKeyID.Points), 0)

		addPkiRevocationDistributionPoint := types.MsgAddPkiRevocationDistributionPoint{
			Signer:               vendorAcc.String(),
			Vid:                  testconstants.PAACertWithNumericVidVid,
			IsPAA:                true,
			Pid:                  8,
			CrlSignerCertificate: testconstants.PAACertWithNumericVid,
			Label:                "label",
			DataURL:              testconstants.DataURL + "/1",
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
			Vid:                  testconstants.PAACertWithNumericVidVid,
			IsPAA:                true,
			Pid:                  8,
			CrlSignerCertificate: testconstants.PAACertWithNumericVid,
			Label:                "label1",
			DataURL:              testconstants.DataURL + "/2",
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
			Vid:                  testconstants.PAACertWithNumericVidVid,
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
	})
}

type rootCertOptions struct {
	pemCert      string
	info         string
	subject      string
	subjectKeyID string
	vid          int32
}

func createTestRootCertOptions() *rootCertOptions {
	return &rootCertOptions{
		pemCert:      testconstants.RootCertPem,
		info:         testconstants.Info,
		subject:      testconstants.RootSubject,
		subjectKeyID: testconstants.RootSubjectKeyID,
		vid:          testconstants.Vid,
	}
}

func createPAACertWithNumericVidOptions() *rootCertOptions {
	return &rootCertOptions{
		pemCert:      testconstants.PAACertWithNumericVid,
		info:         testconstants.Info,
		subject:      testconstants.PAACertWithNumericVidSubject,
		subjectKeyID: testconstants.PAACertWithNumericVidSubjectKeyID,
		vid:          testconstants.PAACertWithNumericVidVid,
	}
}

func createPAACertNoVidOptions(vid int32) *rootCertOptions {
	return &rootCertOptions{
		pemCert:      testconstants.PAACertNoVid,
		info:         testconstants.Info,
		subject:      testconstants.PAACertNoVidSubject,
		subjectKeyID: testconstants.PAACertNoVidSubjectKeyID,
		vid:          vid,
	}
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
