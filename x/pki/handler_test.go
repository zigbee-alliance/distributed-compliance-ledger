package pki

import (
	"context"
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	testkeeper "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
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
	Trustee sdk.AccAddress
}

func (setup *TestSetup) AddAccount(
	accAddress sdk.AccAddress,
	roles []dclauthtypes.AccountRole,
) {
	dclauthKeeper := setup.DclauthKeeper

	for _, role := range roles {
		dclauthKeeper.On("HasRole", mock.Anything, accAddress, role).Return(true)
	}
	dclauthKeeper.On("HasRole", mock.Anything, accAddress, mock.Anything).Return(false)
}

func GenerateAccAddress() sdk.AccAddress {
	_, _, accAddress := testdata.KeyTestPubAddr()
	return accAddress
}

func Setup(t *testing.T) *TestSetup {
	dclauthKeeper := &DclauthKeeperMock{}
	keeper, ctx := testkeeper.PkiKeeper(t, dclauthKeeper)

	trustee := GenerateAccAddress()

	setup := &TestSetup{
		T:             t,
		Ctx:           ctx,
		Wctx:          sdk.WrapSDKContext(ctx),
		Keeper:        keeper,
		DclauthKeeper: dclauthKeeper,
		Handler:       NewHandler(*keeper),
		Trustee:       trustee,
	}

	setup.AddAccount(trustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee})

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
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role})

		// propose x509 root certificate
		proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(accAddress.String(), testconstants.RootCertPem)
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
		require.Nil(t, proposedCertificate.Approvals)

		// check that unique certificate key is registered
		require.True(t, setup.Keeper.IsUniqueCertificatePresent(
			setup.Ctx, testconstants.RootIssuer, testconstants.RootSerialNumber))

		// try to query approved certificate
		_, err = querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
		require.Error(t, err)
		require.Equal(t, codes.NotFound, status.Code(err))

		// cleanup for next iteration
		setup.Keeper.RemoveProposedCertificate(setup.Ctx, testconstants.RootSubject, testconstants.RootSubjectKeyID)
		setup.Keeper.RemoveUniqueCertificate(setup.Ctx, testconstants.RootIssuer, testconstants.RootSerialNumber)
	}
}

func TestHandler_ProposeAddX509RootCert_ByTrustee(t *testing.T) {
	setup := Setup(t)

	// propose x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee.String(), testconstants.RootCertPem)
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
	require.Equal(t, []string{proposeAddX509RootCert.Signer}, proposedCertificate.Approvals)

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
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee.String(), testconstants.StubCertPem)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Error(t, err)
	require.True(t, types.ErrInvalidCertificate.Is(err))
}

func TestHandler_ProposeAddX509RootCert_ForNonRootCertificate(t *testing.T) {
	setup := Setup(t)

	// propose x509 leaf certificate as root
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee.String(), testconstants.LeafCertPem)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Error(t, err)
	require.True(t, types.ErrInappropriateCertificateType.Is(err))
}

func TestHandler_ProposeAddX509RootCert_ProposedCertificateAlreadyExists(t *testing.T) {
	setup := Setup(t)

	// propose adding of x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee.String(), testconstants.RootCertPem)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// store another account
	anotherAccount := GenerateAccAddress()
	setup.AddAccount(anotherAccount, []dclauthtypes.AccountRole{dclauthtypes.Vendor})

	// propose adding of the same x509 root certificate again
	proposeAddX509RootCert = types.NewMsgProposeAddX509RootCert(anotherAccount.String(), testconstants.RootCertPem)
	_, err = setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Error(t, err)
	require.True(t, types.ErrProposedCertificateAlreadyExists.Is(err))
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
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee.String(), testconstants.RootCertPem)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Error(t, err)
	require.True(t, types.ErrCertificateAlreadyExists.Is(err))
}

func TestHandler_ProposeAddX509RootCert_ForDifferentSerialNumber(t *testing.T) {
	setup := Setup(t)

	// store root certificate with different serial number
	rootCertificate := rootCertificate(setup.Trustee)
	rootCertificate.SerialNumber = SerialNumber
	setup.Keeper.SetUniqueCertificate(
		setup.Ctx,
		uniqueCertificate(rootCertificate.Subject, rootCertificate.SerialNumber),
	)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// propose second root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee.String(), testconstants.RootCertPem)
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
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee.String(), testconstants.RootCertPem)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_ApproveAddX509RootCert_ForNotEnoughApprovals(t *testing.T) {
	setup := Setup(t)

	// store account without trustee role
	nonTrustee := GenerateAccAddress()
	setup.AddAccount(nonTrustee, []dclauthtypes.AccountRole{dclauthtypes.Vendor})

	// propose x509 root certificate by account without trustee role
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(nonTrustee.String(), testconstants.RootCertPem)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	// query certificate
	proposedCertificate, _ := queryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, proposeAddX509RootCert.Cert, proposedCertificate.PemCert)
	require.Equal(t, []string{setup.Trustee.String()}, proposedCertificate.Approvals)

	// query approved certificate
	_, err = querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
}

func TestHandler_ApproveAddX509RootCert_ForEnoughApprovals(t *testing.T) {
	setup := Setup(t)

	// propose add x509 root certificate by trustee
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee.String(), testconstants.RootCertPem)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// store second trustee
	secondTrustee := GenerateAccAddress()
	setup.AddAccount(secondTrustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee})

	// approve by second trustee
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		secondTrustee.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID)
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
		setup.Trustee.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID)
	_, err := setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.Error(t, err)
	require.True(t, types.ErrProposedCertificateDoesNotExist.Is(err))
}

func TestHandler_ApproveAddX509RootCert_ByNotTrustee(t *testing.T) {
	setup := Setup(t)

	// propose add x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee.String(), testconstants.RootCertPem)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.Vendor,
		dclauthtypes.CertificationCenter,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role})

		// approve
		approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
			accAddress.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID)
		_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
	}
}

func TestHandler_ApproveAddX509RootCert_Twice(t *testing.T) {
	setup := Setup(t)

	// store account without Trustee role
	accAddress := GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor})

	// propose add x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(accAddress.String(), testconstants.RootCertPem)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID)
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
	rootCertificate := rootCertificate(setup.Trustee)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.Vendor,
		dclauthtypes.CertificationCenter,
		dclauthtypes.Trustee,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role})

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
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee.String(), testconstants.StubCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.Error(t, err)
	require.True(t, types.ErrInvalidCertificate.Is(err))
}

func TestHandler_AddX509Cert_ForRootCertificate(t *testing.T) {
	setup := Setup(t)

	// add root certificate as leaf x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee.String(), testconstants.RootCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.Error(t, err)
	require.True(t, types.ErrInappropriateCertificateType.Is(err))
}

func TestHandler_AddX509Cert_ForDuplicate(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// store intermediate certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// store intermediate certificate second time
	_, err = setup.Handler(setup.Ctx, addX509Cert)
	require.Error(t, err)
	require.True(t, types.ErrCertificateAlreadyExists.Is(err))
}

func TestHandler_AddX509Cert_ForDifferentSerialNumber(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// store intermediate certificate with different serial number
	intermediateCertificate := intermediateCertificate(setup.Trustee)
	intermediateCertificate.SerialNumber = SerialNumber
	setup.Keeper.SetUniqueCertificate(
		setup.Ctx,
		uniqueCertificate(intermediateCertificate.Issuer, intermediateCertificate.SerialNumber),
	)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, intermediateCertificate)

	// store intermediate certificate second time
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee.String(), testconstants.IntermediateCertPem)
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
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_AddX509Cert_ForAbsentDirectParentCert(t *testing.T) {
	setup := Setup(t)

	// add intermediate x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.Error(t, err)
	require.True(t, types.ErrInvalidCertificate.Is(err))
}

func TestHandler_AddX509Cert_ForNoRootCert(t *testing.T) {
	setup := Setup(t)

	// add intermediate certificate
	intermediateCertificate := intermediateCertificate(setup.Trustee)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, intermediateCertificate)

	// add leaf x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee.String(), testconstants.LeafCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.Error(t, err)
	require.True(t, types.ErrInvalidCertificate.Is(err))
}

func TestHandler_AddX509Cert_ForFailedCertificateVerification(t *testing.T) {
	setup := Setup(t)

	// add invalid root
	invalidRootCertificate := types.NewRootCertificate(testconstants.StubCertPem,
		testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.RootSerialNumber, setup.Trustee.String())
	setup.Keeper.AddApprovedCertificate(setup.Ctx, invalidRootCertificate)

	// add intermediate x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.Error(t, err)
	require.True(t, types.ErrInvalidCertificate.Is(err))
}

func TestHandler_AddX509Cert_ForTree(t *testing.T) {
	setup := Setup(t)

	// add root x509 certificate
	proposeAndApproveRootCertificate(setup, setup.Trustee)

	// add intermediate x509 certificate
	addIntermediateX509Cert := types.NewMsgAddX509Cert(setup.Trustee.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addIntermediateX509Cert)
	require.NoError(t, err)

	// add leaf x509 certificate
	addLeafX509Cert := types.NewMsgAddX509Cert(setup.Trustee.String(), testconstants.LeafCertPem)
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
	intermediateCertificate, _ :=
		querySingleApprovedCertificate(setup, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
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
	rootCert := rootCertificate(setup.Trustee)

	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCert)
	setup.Keeper.SetUniqueCertificate(setup.Ctx, uniqueCertificate(rootCert.Subject, rootCert.SerialNumber))

	// store second root certificate
	rootCert = rootCertificate(setup.Trustee)
	rootCert.SerialNumber = SerialNumber

	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCert)
	setup.Keeper.SetUniqueCertificate(setup.Ctx, uniqueCertificate(rootCert.Subject, rootCert.SerialNumber))

	// store intermediate certificate (it refers to two parent certificates)
	intermediateCertificate := intermediateCertificate(setup.Trustee)
	intermediateCertificate.SerialNumber = SerialNumber

	setup.Keeper.AddApprovedCertificate(setup.Ctx, intermediateCertificate)
	setup.Keeper.SetUniqueCertificate(
		setup.Ctx,
		uniqueCertificate(intermediateCertificate.Issuer, intermediateCertificate.SerialNumber),
	)

	childCertId := certificateIdentifier(intermediateCertificate.Subject, intermediateCertificate.SubjectKeyId)
	rootChildCertificates := types.ChildCertificates{
		Issuer:         intermediateCertificate.Issuer,
		AuthorityKeyId: intermediateCertificate.AuthorityKeyId,
		CertIds:        []*types.CertificateIdentifier{&childCertId},
	}
	setup.Keeper.SetChildCertificates(setup.Ctx, rootChildCertificates)

	// store second intermediate certificate (it refers to two parent certificates)
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// store leaf certificate (it refers to two parent certificates)
	addX509Cert = types.NewMsgAddX509Cert(setup.Trustee.String(), testconstants.LeafCertPem)
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
	proposeAndApproveRootCertificate(setup, setup.Trustee)

	// propose revocation of x509 root certificate by `setup.Trustee`
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	// query and check proposed certificate revocation
	proposedRevocation, _ := queryProposedCertificateRevocation(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, testconstants.RootSubject, proposedRevocation.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, proposedRevocation.SubjectKeyId)
	require.Equal(t, []string{setup.Trustee.String()}, proposedRevocation.Approvals)

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
	proposeAndApproveRootCertificate(setup, setup.Trustee)

	// store another trustee
	anotherTrustee := GenerateAccAddress()
	setup.AddAccount(anotherTrustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee})

	// propose revocation of x509 root certificate by new trustee
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		anotherTrustee.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	// query and check proposed certificate revocation
	proposedRevocation, _ := queryProposedCertificateRevocation(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, testconstants.RootSubject, proposedRevocation.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, proposedRevocation.SubjectKeyId)
	require.Equal(t, []string{anotherTrustee.String()}, proposedRevocation.Approvals)

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
	proposeAndApproveRootCertificate(setup, setup.Trustee)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.Vendor,
		dclauthtypes.CertificationCenter,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role})

		// propose revocation of x509 root certificate
		proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
			accAddress.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID)
		_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
	}
}

func TestHandler_ProposeRevokeX509RootCert_CertificateDoesNotExist(t *testing.T) {
	setup := Setup(t)

	// propose revocation of not existing certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Error(t, err)
	require.True(t, types.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_ProposeRevokeX509RootCert_ForProposedCertificate(t *testing.T) {
	setup := Setup(t)

	// propose x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee.String(), testconstants.RootCertPem)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// check that proposed certificate is present
	proposedCertificate, _ := queryProposedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NotNil(t, proposedCertificate)

	// propose revocation of proposed root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID)
	_, err = setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Error(t, err)
	require.True(t, types.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_ProposeRevokeX509RootCert_ProposedRevocationAlreadyExists(t *testing.T) {
	setup := Setup(t)

	// propose and approve x509 root certificate
	proposeAndApproveRootCertificate(setup, setup.Trustee)

	// propose revocation of x509 root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	// store another trustee
	anotherTrustee := GenerateAccAddress()
	setup.AddAccount(anotherTrustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee})

	// propose revocation of the same x509 root certificate again
	proposeRevokeX509RootCert = types.NewMsgProposeRevokeX509RootCert(
		anotherTrustee.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID)
	_, err = setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Error(t, err)
	require.True(t, types.ErrProposedCertificateRevocationAlreadyExists.Is(err))
}

func TestHandler_ProposeRevokeX509RootCert_ForNonRootCertificate(t *testing.T) {
	setup := Setup(t)

	// store x509 root certificate
	rootCertificate := rootCertificate(setup.Trustee)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// store x509 intermediate certificate
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// propose revocation of x509 intermediate certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee.String(), testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	_, err = setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Error(t, err)
	require.True(t, types.ErrInappropriateCertificateType.Is(err))
}

func TestHandler_ApproveRevokeX509RootCert_ForNotEnoughApprovals(t *testing.T) {
	setup := Setup(t)

	// propose and approve x509 root certificate
	proposeAndApproveRootCertificate(setup, setup.Trustee)

	// increase the number of approvals required for root certificates control to three
	oldRootCertificateApprovals := types.RootCertificateApprovals

	defer func() {
		types.RootCertificateApprovals = oldRootCertificateApprovals
	}()

	types.RootCertificateApprovals = 3

	// propose revocation of x509 root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	// store second trustee
	anotherTrustee := GenerateAccAddress()
	setup.AddAccount(anotherTrustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee})

	// approve
	approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
		anotherTrustee.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID)
	_, err = setup.Handler(setup.Ctx, approveRevokeX509RootCert)
	require.NoError(t, err)

	// query and check proposed certificate revocation
	proposedRevocation, _ := queryProposedCertificateRevocation(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, testconstants.RootSubject, proposedRevocation.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, proposedRevocation.SubjectKeyId)
	require.Equal(t, []string{setup.Trustee.String(), anotherTrustee.String()}, proposedRevocation.Approvals)

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
	proposeAndApproveRootCertificate(setup, setup.Trustee)

	// propose revocation of x509 root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	// get certificate for further comparison
	certificateBeforeRevocation, _ :=
		querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NotNil(t, certificateBeforeRevocation)

	// store second trustee
	anotherTrustee := GenerateAccAddress()
	setup.AddAccount(anotherTrustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee})

	// approve
	approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
		anotherTrustee.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID)
	_, err = setup.Handler(setup.Ctx, approveRevokeX509RootCert)
	require.NoError(t, err)

	// check that proposed certificate revocation does not exist anymore
	_, err = queryProposedCertificateRevocation(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
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
	proposeAndApproveRootCertificate(setup, setup.Trustee)

	// propose revocation of x509 root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.Vendor,
		dclauthtypes.CertificationCenter,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role})

		// approve
		approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
			accAddress.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID)
		_, err = setup.Handler(setup.Ctx, approveRevokeX509RootCert)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
	}
}

func TestHandler_ApproveRevokeX509RootCert_ProposedRevocationDoesNotExist(t *testing.T) {
	setup := Setup(t)

	// propose and approve x509 root certificate
	proposeAndApproveRootCertificate(setup, setup.Trustee)

	// approve revocation of x509 root certificate
	approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
		setup.Trustee.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID)
	_, err := setup.Handler(setup.Ctx, approveRevokeX509RootCert)
	require.Error(t, err)
	require.True(t, types.ErrProposedCertificateRevocationDoesNotExist.Is(err))
}

func TestHandler_ApproveRevokeX509RootCert_Twice(t *testing.T) {
	setup := Setup(t)

	// propose and approve x509 root certificate
	proposeAndApproveRootCertificate(setup, setup.Trustee)

	// propose revocation of x509 root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID)
	_, err := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	// approve revocation by the same trustee
	approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
		setup.Trustee.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID)
	_, err = setup.Handler(setup.Ctx, approveRevokeX509RootCert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

//nolint:funlen
func TestHandler_ApproveRevokeX509RootCert_ForTree(t *testing.T) {
	setup := Setup(t)

	// add root x509 certificate
	proposeAndApproveRootCertificate(setup, setup.Trustee)

	// add intermediate x509 certificate
	addIntermediateX509Cert := types.NewMsgAddX509Cert(setup.Trustee.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addIntermediateX509Cert)
	require.NoError(t, err)

	// add leaf x509 certificate
	addLeafX509Cert := types.NewMsgAddX509Cert(setup.Trustee.String(), testconstants.LeafCertPem)
	_, err = setup.Handler(setup.Ctx, addLeafX509Cert)
	require.NoError(t, err)

	// propose revocation of x509 root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		setup.Trustee.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID)
	_, err = setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.NoError(t, err)

	// store second trustee
	anotherTrustee := GenerateAccAddress()
	setup.AddAccount(anotherTrustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee})

	// approve
	approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
		anotherTrustee.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID)
	_, err = setup.Handler(setup.Ctx, approveRevokeX509RootCert)
	require.NoError(t, err)

	// check that root, intermediate and leaf certificates have been revoked
	allRevokedCertificates, _ := queryAllRevokedCertificates(setup)
	require.Equal(t, 3, len(allRevokedCertificates))
	require.Equal(t, testconstants.IntermediateSubject, allRevokedCertificates[0].Subject)
	require.Equal(t, testconstants.IntermediateSubjectKeyID, allRevokedCertificates[0].SubjectKeyId)
	require.Equal(t, 1, len(allRevokedCertificates[0].Certs))
	require.Equal(t, testconstants.IntermediateCertPem, allRevokedCertificates[0].Certs[0].PemCert)
	require.Equal(t, testconstants.LeafSubject, allRevokedCertificates[1].Subject)
	require.Equal(t, testconstants.LeafSubjectKeyID, allRevokedCertificates[1].SubjectKeyId)
	require.Equal(t, 1, len(allRevokedCertificates[1].Certs))
	require.Equal(t, testconstants.LeafCertPem, allRevokedCertificates[1].Certs[0].PemCert)
	require.Equal(t, testconstants.RootSubject, allRevokedCertificates[2].Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, allRevokedCertificates[2].SubjectKeyId)
	require.Equal(t, 1, len(allRevokedCertificates[2].Certs))
	require.Equal(t, testconstants.RootCertPem, allRevokedCertificates[2].Certs[0].PemCert)

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
	rootCertificate := rootCertificate(setup.Trustee)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.Vendor,
		dclauthtypes.CertificationCenter,
		dclauthtypes.Trustee,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role})

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
			accAddress.String(), testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
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
		setup.Trustee.String(), testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.Error(t, err)
	require.True(t, types.ErrCertificateDoesNotExist.Is(err))
}

func TestHandler_RevokeX509Cert_ForRootCertificate(t *testing.T) {
	setup := Setup(t)

	// propose and approve x509 root certificate
	proposeAndApproveRootCertificate(setup, setup.Trustee)

	// revoke x509 root certificate
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		setup.Trustee.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID)
	_, err := setup.Handler(setup.Ctx, revokeX509Cert)
	require.Error(t, err)
	require.True(t, types.ErrInappropriateCertificateType.Is(err))
}

func TestHandler_RevokeX509Cert_ByNotOwner(t *testing.T) {
	setup := Setup(t)

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee)
	setup.Keeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// add x509 certificate by `setup.Trustee`
	addX509Cert := types.NewMsgAddX509Cert(setup.Trustee.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	// store another account
	anotherTrustee := GenerateAccAddress()
	setup.AddAccount(anotherTrustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee})

	// revoke x509 certificate by another account
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		anotherTrustee.String(), testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	_, err = setup.Handler(setup.Ctx, revokeX509Cert)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_RevokeX509Cert_ForTree(t *testing.T) {
	setup := Setup(t)

	// add root x509 certificate
	proposeAndApproveRootCertificate(setup, setup.Trustee)

	// add intermediate x509 certificate
	addIntermediateX509Cert := types.NewMsgAddX509Cert(setup.Trustee.String(), testconstants.IntermediateCertPem)
	_, err := setup.Handler(setup.Ctx, addIntermediateX509Cert)
	require.NoError(t, err)

	// add leaf x509 certificate
	addLeafX509Cert := types.NewMsgAddX509Cert(setup.Trustee.String(), testconstants.LeafCertPem)
	_, err = setup.Handler(setup.Ctx, addLeafX509Cert)
	require.NoError(t, err)

	// revoke x509 certificate
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		setup.Trustee.String(), testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	_, err = setup.Handler(setup.Ctx, revokeX509Cert)
	require.NoError(t, err)

	// check that intermediate and leaf certificates have been revoked
	allRevokedCertificates, _ := queryAllRevokedCertificates(setup)
	require.Equal(t, 2, len(allRevokedCertificates))
	require.Equal(t, testconstants.IntermediateSubject, allRevokedCertificates[0].Subject)
	require.Equal(t, testconstants.IntermediateSubjectKeyID, allRevokedCertificates[0].SubjectKeyId)
	require.Equal(t, 1, len(allRevokedCertificates[0].Certs))
	require.Equal(t, testconstants.IntermediateCertPem, allRevokedCertificates[0].Certs[0].PemCert)
	require.Equal(t, testconstants.LeafSubject, allRevokedCertificates[1].Subject)
	require.Equal(t, testconstants.LeafSubjectKeyID, allRevokedCertificates[1].SubjectKeyId)
	require.Equal(t, 1, len(allRevokedCertificates[1].Certs))
	require.Equal(t, testconstants.LeafCertPem, allRevokedCertificates[1].Certs[0].PemCert)

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

func proposeAndApproveRootCertificate(setup *TestSetup, ownerTrustee sdk.AccAddress) {
	// ensure that `ownerTrustee` is trustee to eventually have enough approvals
	require.True(setup.T, setup.DclauthKeeper.HasRole(setup.Ctx, ownerTrustee, types.RootCertificateApprovalRole))

	// propose x509 root certificate by `ownerTrustee`
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(ownerTrustee.String(), testconstants.RootCertPem)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(setup.T, err)

	// store another trustee account
	anotherTrustee := GenerateAccAddress()
	setup.AddAccount(anotherTrustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee})

	// approve x509 root certificate by another trustee
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		anotherTrustee.String(), testconstants.RootSubject, testconstants.RootSubjectKeyID)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(setup.T, err)

	// check that root certificate has been approved
	approvedCertificate, err := querySingleApprovedCertificate(
		setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NoError(setup.T, err)
	require.NotNil(setup.T, approvedCertificate)
}

func queryProposedCertificate(
	setup *TestSetup,
	subject string,
	subjectKeyId string,
) (*types.ProposedCertificate, error) {
	// query proposed certificate
	req := &types.QueryGetProposedCertificateRequest{
		Subject:      subject,
		SubjectKeyId: subjectKeyId,
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
	subjectKeyId string,
) (*types.Certificate, error) {
	certificates, err := queryApprovedCertificates(setup, subject, subjectKeyId)
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
	subjectKeyId string,
) (*types.ApprovedCertificates, error) {
	// query certificate
	req := &types.QueryGetApprovedCertificatesRequest{
		Subject:      subject,
		SubjectKeyId: subjectKeyId,
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
	subject string,
	subjectKeyId string,
) (*types.ProposedCertificateRevocation, error) {
	// query proposed certificate revocation
	req := &types.QueryGetProposedCertificateRevocationRequest{
		Subject:      subject,
		SubjectKeyId: subjectKeyId,
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
	subjectKeyId string,
) (*types.Certificate, error) {
	certificates, err := queryRevokedCertificates(setup, subject, subjectKeyId)
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
	subjectKeyId string,
) (*types.RevokedCertificates, error) {
	// query revoked certificate
	req := &types.QueryGetRevokedCertificatesRequest{
		Subject:      subject,
		SubjectKeyId: subjectKeyId,
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
	authorityKeyId string,
) (*types.ChildCertificates, error) {
	// query certificate
	req := &types.QueryGetChildCertificatesRequest{
		Issuer:         issuer,
		AuthorityKeyId: authorityKeyId,
	}

	resp, err := setup.Keeper.ChildCertificates(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)
		return nil, err
	}

	require.NotNil(setup.T, resp)
	return &resp.ChildCertificates, nil
}

func rootCertificate(address sdk.AccAddress) types.Certificate {
	return types.NewRootCertificate(
		testconstants.RootCertPem,
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.RootSerialNumber,
		address.String(),
	)
}

func intermediateCertificate(address sdk.AccAddress) types.Certificate {
	return types.NewNonRootCertificate(
		testconstants.IntermediateCertPem,
		testconstants.IntermediateSubject,
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

func certificateIdentifier(subject string, subjectKeyId string) types.CertificateIdentifier {
	return types.CertificateIdentifier{
		Subject:      subject,
		SubjectKeyId: subjectKeyId,
	}
}
