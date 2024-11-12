package pki

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	testkeeper "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
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

func createRootWithVidOptions() *rootCertOptions {
	return &rootCertOptions{
		pemCert:      testconstants.RootCertWithVid,
		info:         testconstants.Info,
		subject:      testconstants.RootCertWithVidSubject,
		subjectKeyID: testconstants.RootCertWithVidSubjectKeyID,
		vid:          testconstants.RootCertWithVidVid,
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

func proposeAndApproveRootCertificate(setup *TestSetup, ownerTrustee sdk.AccAddress, options *rootCertOptions) {
	// ensure that `ownerTrustee` is trustee to eventually have enough approvals
	require.True(setup.T, setup.DclauthKeeper.HasRole(setup.Ctx, ownerTrustee, types.RootCertificateApprovalRole))

	// propose x509 root certificate by `ownerTrustee`
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(ownerTrustee.String(), options.pemCert, options.info, options.vid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(setup.T, err)

	// approve x509 root certificate by another trustee
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), options.subject, options.subjectKeyID, options.info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(setup.T, err)

	// check that root certificate has been approved
	approvedCertificate, err := queryApprovedCertificates(
		setup, options.subject, options.subjectKeyID)
	require.NoError(setup.T, err)
	require.NotNil(setup.T, approvedCertificate)
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
	return _queryAllApprovedCertificates(setup, "")
}

func queryAllNocCertificates(setup *TestSetup) ([]types.NocCertificates, error) {
	// query all certificates
	return _queryAllNocCertificates(setup, "")
}

func queryAllApprovedCertificatesBySubjectKeyID(setup *TestSetup, subjectKeyID string) ([]types.ApprovedCertificates, error) {
	// query all certificates
	return _queryAllApprovedCertificates(setup, subjectKeyID)
}

func _queryAllApprovedCertificates(setup *TestSetup, subjectKeyID string) ([]types.ApprovedCertificates, error) {
	// query all certificates
	req := &types.QueryAllApprovedCertificatesRequest{
		SubjectKeyId: subjectKeyID,
	}

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

func queryApprovedCertificatesBySubject(
	setup *TestSetup,
	subject string,
) (*types.ApprovedCertificatesBySubject, error) {
	// query certificate
	req := &types.QueryGetApprovedCertificatesBySubjectRequest{
		Subject: subject,
	}

	resp, err := setup.Keeper.ApprovedCertificatesBySubject(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.ApprovedCertificatesBySubject, nil
}

func queryApprovedRootCertificates(
	setup *TestSetup,
	subject string,
	subjectKeyID string,
) ([]*types.Certificate, error) {
	resp, err := queryApprovedCertificates(setup, subject, subjectKeyID)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}
	var list []*types.Certificate
	for _, cert := range resp.Certs {
		if cert.IsRoot {
			list = append(list, cert)
		}
	}

	return list, nil
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
	serialNumber string,
) (*types.ProposedCertificateRevocation, error) {
	// query proposed certificate revocation
	req := &types.QueryGetProposedCertificateRevocationRequest{
		Subject:      testconstants.RootSubject,
		SubjectKeyId: testconstants.RootSubjectKeyID,
		SerialNumber: serialNumber,
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

func queryRevokedRootCertificates(setup *TestSetup) (*types.RevokedRootCertificates, error) {
	// query revoked root certificate
	req := &types.QueryGetRevokedRootCertificatesRequest{}

	resp, err := setup.Keeper.RevokedRootCertificates(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.RevokedRootCertificates, nil
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

func queryAllNocCertificatesBySubjectKeyID(setup *TestSetup, subjectKeyID string) ([]types.NocCertificates, error) {
	// query all noc certificates
	return _queryAllNocCertificates(setup, subjectKeyID)
}

func _queryAllNocCertificates(setup *TestSetup, subjectKeyID string) ([]types.NocCertificates, error) {
	// query all certificates
	req := &types.QueryNocCertificatesRequest{
		SubjectKeyId: subjectKeyID,
	}

	resp, err := setup.Keeper.NocCertificatesAll(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return resp.NocCertificates, nil
}

func querySingleNocCertificate(
	setup *TestSetup,
	subject string,
	subjectKeyID string,
) (*types.Certificate, error) {
	certificates, err := queryNocCertificates(setup, subject, subjectKeyID)
	if err != nil {
		return nil, err
	}

	if len(certificates.Certs) > 1 {
		require.Fail(setup.T, "More than 1 certificate returned")
	}

	return certificates.Certs[0], nil
}

func querySingleNocRootCertificateByVid(
	setup *TestSetup,
	vid int32,
) (*types.Certificate, error) {
	certificates, err := queryNocRootCertificatesByVid(setup, vid)
	if err != nil {
		return nil, err
	}

	if len(certificates.Certs) > 1 {
		require.Fail(setup.T, "More than 1 certificate returned")
	}

	return certificates.Certs[0], nil
}

func queryNocRootCertificatesByVid(
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

func querySingleNocIcaCertificateByVid(
	setup *TestSetup,
	vid int32,
) (*types.Certificate, error) {
	certificates, err := queryNocIcaCertificatesByVid(setup, vid)
	if err != nil {
		return nil, err
	}

	if len(certificates.Certs) > 1 {
		require.Fail(setup.T, "More than 1 certificate returned")
	}

	return certificates.Certs[0], nil
}

func queryNocIcaCertificatesByVid(
	setup *TestSetup,
	vid int32,
) (*types.NocIcaCertificates, error) {
	// query certificate
	req := &types.QueryGetNocIcaCertificatesRequest{Vid: vid}

	resp, err := setup.Keeper.NocIcaCertificates(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.NocIcaCertificates, nil
}

func queryNocCertificates(
	setup *TestSetup,
	subject string,
	subjectKeyID string,
) (*types.NocCertificates, error) {
	// query certificate
	req := &types.QueryGetNocCertificatesRequest{
		Subject:      subject,
		SubjectKeyId: subjectKeyID,
	}

	resp, err := setup.Keeper.NocCertificates(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.NocCertificates, nil
}

func queryNocCertificatesBySubject(
	setup *TestSetup,
	subject string,
) (*types.NocCertificatesBySubject, error) {
	// query certificate
	req := &types.QueryGetNocCertificatesBySubjectRequest{
		Subject: subject,
	}

	resp, err := setup.Keeper.NocCertificatesBySubject(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.NocCertificatesBySubject, nil
}

func querySingleNocCertificateByVidAndSkid(
	setup *TestSetup,
	vid int32,
	subjectKeyID string,
) (*types.Certificate, float32, error) {
	certificates, err := queryNocCertificatesByVidAndSkid(setup, vid, subjectKeyID)
	if err != nil {
		return nil, 0, err
	}

	if len(certificates.Certs) > 1 {
		require.Fail(setup.T, "More than 1 certificate returned")
	}

	return certificates.Certs[0], certificates.Tq, nil
}

func queryNocCertificatesByVidAndSkid(
	setup *TestSetup,
	vid int32,
	subjectKeyID string,
) (*types.NocCertificatesByVidAndSkid, error) {
	// query certificate
	req := &types.QueryGetNocCertificatesByVidAndSkidRequest{Vid: vid, SubjectKeyId: subjectKeyID}

	resp, err := setup.Keeper.NocCertificatesByVidAndSkid(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.NocCertificatesByVidAndSkid, nil
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

func queryRevokedNocRootCertificates(setup *TestSetup, subject, subjectKeyID string) (*types.RevokedNocRootCertificates, error) { //nolint:unparam
	// query certificate
	req := &types.QueryGetRevokedNocRootCertificatesRequest{Subject: subject, SubjectKeyId: subjectKeyID}

	resp, err := setup.Keeper.RevokedNocRootCertificates(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.RevokedNocRootCertificates, nil
}

func queryAllRevokedNocIcaCertificates(setup *TestSetup) ([]types.RevokedNocIcaCertificates, error) { //nolint:unparam
	// query certificate
	req := &types.QueryAllRevokedNocIcaCertificatesRequest{}

	resp, err := setup.Keeper.RevokedNocIcaCertificatesAll(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return resp.RevokedNocIcaCertificates, nil
}

func queryRevokedNocIcaCertificates(setup *TestSetup, subject, subjectKeyID string) (*types.RevokedNocIcaCertificates, error) { //nolint:unparam
	// query certificate
	req := &types.QueryGetRevokedNocIcaCertificatesRequest{Subject: subject, SubjectKeyId: subjectKeyID}

	resp, err := setup.Keeper.RevokedNocIcaCertificates(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.RevokedNocIcaCertificates, nil
}

func queryCertificatesFromAllCertificatesIndex(
	setup *TestSetup,
	subject string,
	subjectKeyID string,
) (*types.AllCertificates, error) {
	// query certificate
	req := &types.QueryGetCertificatesRequest{
		Subject:      subject,
		SubjectKeyId: subjectKeyID,
	}

	resp, err := setup.Keeper.Certificates(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.Certificates, nil
}

func querySingleCertificateFromAllCertificatesIndex(
	setup *TestSetup,
	subject string,
	subjectKeyID string,
) (*types.Certificate, error) {
	certificates, err := queryCertificatesFromAllCertificatesIndex(setup, subject, subjectKeyID)
	if err != nil {
		return nil, err
	}

	if len(certificates.Certs) > 1 {
		require.Fail(setup.T, "More than 1 certificate returned")
	}

	return certificates.Certs[0], nil
}

func queryCertificatesBySubjectFromAllCertificatesIndex(
	setup *TestSetup,
	subject string,
) (*types.AllCertificatesBySubject, error) {
	// query certificate
	req := &types.QueryGetAllCertificatesBySubjectRequest{
		Subject: subject,
	}

	resp, err := setup.Keeper.AllCertificatesBySubject(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.AllCertificatesBySubject, nil
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
		testconstants.SchemaVersion,
	)
}

func intermediateCertificateNoVid(address sdk.AccAddress) types.Certificate {
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
		0,
		testconstants.SchemaVersion,
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

func ensureCertificatePresentInGlobalCertificateIndexes(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyId string,
	serialNumber string,
	skipCheckForSubject bool, // TODO: FIX constants and eliminate this condition
) {
	// AllCertificate: Subject and SKID
	allCertificate, err := querySingleCertificateFromAllCertificatesIndex(setup, subject, subjectKeyId)
	require.NoError(t, err)
	require.Equal(t, subject, allCertificate.Subject)
	require.Equal(t, subjectKeyId, allCertificate.SubjectKeyId)
	require.Equal(t, serialNumber, allCertificate.SerialNumber)

	if !skipCheckForSubject {
		// AllCertificate: Subject
		allCertificatesBySubject, err := queryCertificatesBySubjectFromAllCertificatesIndex(setup, subject)
		require.NoError(t, err)
		require.Equal(t, 1, len(allCertificatesBySubject.SubjectKeyIds))
		require.Equal(t, subjectKeyId, allCertificatesBySubject.SubjectKeyIds[0])
	}
}

func ensureCertificateNotPresentInGlobalCertificateIndexes(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyId string,
	skipCheckForSubject bool, // TODO: FIX constants and eliminate this condition
) {
	// All certificates indexes checks

	// AllCertificate: Subject and SKID
	_, err := querySingleCertificateFromAllCertificatesIndex(setup, subject, subjectKeyId)
	require.Equal(t, codes.NotFound, status.Code(err))

	if !skipCheckForSubject {
		// AllCertificate: Subject
		_, err = queryCertificatesBySubjectFromAllCertificatesIndex(setup, subject)
		require.Equal(t, codes.NotFound, status.Code(err))
	}
}

func ensureCertificatePresentInDaCertificateIndexes(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyId string,
	serialNumber string,
	isRoot bool,
	skipCheckForSubject bool, // TODO: FIX constants and eliminate this condition
) {
	// DaCertificates: Subject and SKID
	approvedCertificate, _ := querySingleApprovedCertificate(setup, subject, subjectKeyId)
	require.Equal(t, subject, approvedCertificate.Subject)
	require.Equal(t, subjectKeyId, approvedCertificate.SubjectKeyId)
	require.Equal(t, serialNumber, approvedCertificate.SerialNumber)
	require.Equal(t, isRoot, approvedCertificate.IsRoot)

	// DaCertificates: SKID
	certificateBySubjectKeyID, _ := queryAllApprovedCertificatesBySubjectKeyID(setup, subjectKeyId)
	require.Equal(t, 1, len(certificateBySubjectKeyID))
	require.Equal(t, 1, len(certificateBySubjectKeyID[0].Certs))

	if !skipCheckForSubject {
		// DACertificates: Subject
		certificatesBySubject, err := queryApprovedCertificatesBySubject(setup, subject)
		require.NoError(t, err)
		require.Equal(t, 1, len(certificatesBySubject.SubjectKeyIds))
		require.Equal(t, subjectKeyId, certificatesBySubject.SubjectKeyIds[0])
	}
}

func ensureCertificatePresentInNocCertificateIndexes(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyId string,
	serialNumber string,
	vid int32,
	isRoot bool,
	skipCheckByVid bool,
) {
	// Noc certificates indexes checks

	// NocCertificates: Subject and SKID
	nocCertificate, err := querySingleNocCertificate(setup, subject, subjectKeyId)
	require.NoError(t, err)
	require.Equal(t, subject, nocCertificate.Subject)
	require.Equal(t, subjectKeyId, nocCertificate.SubjectKeyId)
	require.Equal(t, serialNumber, nocCertificate.SerialNumber)
	require.Equal(t, testconstants.SchemaVersion, nocCertificate.SchemaVersion)

	// NocCertificates: SubjectKeyID
	nocCertificatesBySubjectKeyID, err := queryAllNocCertificatesBySubjectKeyID(setup, subjectKeyId)
	require.NoError(t, err)
	require.Equal(t, 1, len(nocCertificatesBySubjectKeyID))
	require.Equal(t, 1, len(nocCertificatesBySubjectKeyID[0].Certs))
	require.Equal(t, serialNumber, nocCertificatesBySubjectKeyID[0].Certs[0].SerialNumber)

	// NocCertificates: Subject
	nocCertificatesBySubject, err := queryNocCertificatesBySubject(setup, subject)
	require.NoError(t, err)
	require.Equal(t, 1, len(nocCertificatesBySubject.SubjectKeyIds))
	require.Equal(t, subjectKeyId, nocCertificatesBySubject.SubjectKeyIds[0])

	// NocCertificates: VID and SKID
	nocCertificateByVidAndSkid, _, err := querySingleNocCertificateByVidAndSkid(setup, vid, subjectKeyId)
	require.NoError(t, err)
	require.Equal(t, subject, nocCertificateByVidAndSkid.Subject)
	require.Equal(t, subjectKeyId, nocCertificateByVidAndSkid.SubjectKeyId)
	require.Equal(t, serialNumber, nocCertificateByVidAndSkid.SerialNumber)

	if skipCheckByVid {
		return
	}

	// NocCertificates: VID
	if isRoot {
		nocRootCertificate, err := querySingleNocRootCertificateByVid(setup, vid)
		require.NoError(t, err)
		require.Equal(t, serialNumber, nocRootCertificate.SerialNumber)
	} else {
		nocRootCertificate, err := querySingleNocIcaCertificateByVid(setup, vid)
		require.NoError(t, err)
		require.Equal(t, serialNumber, nocRootCertificate.SerialNumber)
	}
}

func ensureCertificateNotPresentInDaCertificateIndexes(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyId string,
	skipCheckForSubject bool, // TODO: FIX constants and eliminate this condition
) {
	// DA certificates indexes checks

	// DaCertificates: Subject and SKID
	_, err := querySingleApprovedCertificate(setup, subject, subjectKeyId)
	require.Equal(t, codes.NotFound, status.Code(err))

	// DaCertificates: SubjectKeyID
	certificatesBySubjectKeyID, _ := queryAllApprovedCertificatesBySubjectKeyID(setup, subjectKeyId)
	require.Equal(t, 0, len(certificatesBySubjectKeyID))

	if !skipCheckForSubject {
		// NocCertificates: Subject
		_, err = queryApprovedCertificatesBySubject(setup, subject)
		require.Equal(t, codes.NotFound, status.Code(err))
	}
}

func ensureCertificateNotPresentInNocCertificateIndexes(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyId string,
	vid int32,
	isRoot bool,
	skipCheckByVid bool,
) {
	// Noc certificates indexes checks

	// NocCertificates: Subject and SKID
	_, err := querySingleNocCertificate(setup, subject, subjectKeyId)
	require.Equal(t, codes.NotFound, status.Code(err))

	// NocCertificates: SubjectKeyID
	certificatesBySubjectKeyID, err := queryAllNocCertificatesBySubjectKeyID(setup, subjectKeyId)
	require.Equal(t, 0, len(certificatesBySubjectKeyID))

	// NocCertificates: Subject
	_, err = queryNocCertificatesBySubject(setup, subject)
	require.Equal(t, codes.NotFound, status.Code(err))

	// NocCertificates: VID and SKID
	_, err = queryNocCertificatesByVidAndSkid(setup, vid, subjectKeyId)
	require.Equal(t, codes.NotFound, status.Code(err))

	// NocCertificates: VID
	if skipCheckByVid {
		return
	}

	if isRoot {
		_, err = querySingleNocRootCertificateByVid(setup, vid)
		require.Equal(t, codes.NotFound, status.Code(err))
	} else {
		_, err = querySingleNocIcaCertificateByVid(setup, vid)
		require.Equal(t, codes.NotFound, status.Code(err))
	}
}

func ensureCertificatePresentInUniqueCertificateIndexes(
	t *testing.T,
	setup *TestSetup,
	issuer string,
	serialNumber string,
) {
	// UniqueCertificate: check that unique certificate key registered
	require.True(t, setup.Keeper.IsUniqueCertificatePresent(
		setup.Ctx, issuer, serialNumber))
}

func ensureCertificateNotPresentInUniqueCertificateIndexes(
	t *testing.T,
	setup *TestSetup,
	issuer string,
	serialNumber string,
	skipCheck bool,
) {
	if !skipCheck {
		// UniqueCertificate: check that unique certificate key registered
		found := setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, issuer, serialNumber)
		require.Equal(t, false, found)
	}
}

func ensureDaPaaCertificateExist(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyId string,
	issuer string,
	serialNumber string,
) {
	// DA certificates indexes checks
	ensureCertificatePresentInDaCertificateIndexes(t, setup, subject, subjectKeyId, serialNumber, true, false)

	// All certificates indexes checks
	ensureCertificatePresentInGlobalCertificateIndexes(t, setup, subject, subjectKeyId, serialNumber, false)

	// UniqueCertificate: check that unique certificate key registered
	ensureCertificatePresentInUniqueCertificateIndexes(t, setup, issuer, serialNumber)
}

func ensureDaPaiCertificateExist(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyId string,
	issuer string,
	serialNumber string,
	skipCheckForSubject bool,
) {
	// DA certificates indexes checks
	ensureCertificatePresentInDaCertificateIndexes(t, setup, subject, subjectKeyId, serialNumber, false, skipCheckForSubject)

	// All certificates indexes checks
	ensureCertificatePresentInGlobalCertificateIndexes(t, setup, subject, subjectKeyId, serialNumber, skipCheckForSubject)

	// UniqueCertificate: check that unique certificate key registered
	ensureCertificatePresentInUniqueCertificateIndexes(t, setup, issuer, serialNumber)
}

func ensureDaPaaCertificateDoesNotExist(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyId string,
	issuer string,
	serialNumber string,
	isRevoked bool,
) {
	// DA certificates indexes checks
	ensureCertificateNotPresentInDaCertificateIndexes(t, setup, subject, subjectKeyId, false)

	// All certificates indexes checks
	ensureCertificateNotPresentInGlobalCertificateIndexes(t, setup, subject, subjectKeyId, false)

	// UniqueCertificate: check that unique certificate key registered
	ensureCertificateNotPresentInUniqueCertificateIndexes(t, setup, issuer, serialNumber, isRevoked)
}

func ensureDaPaiCertificateDoesNotExist(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyId string,
	issuer string,
	serialNumber string,
	skipCheckForUniqueness bool,
	skipCheckForSubject bool,
) {
	// DA certificates indexes checks
	ensureCertificateNotPresentInDaCertificateIndexes(t, setup, subject, subjectKeyId, skipCheckForSubject)

	// All certificates indexes checks
	ensureCertificateNotPresentInGlobalCertificateIndexes(t, setup, subject, subjectKeyId, skipCheckForSubject)

	// UniqueCertificate: check that unique certificate key registered
	ensureCertificateNotPresentInUniqueCertificateIndexes(t, setup, issuer, serialNumber, skipCheckForUniqueness)
}

func ensureNocRootCertificateExist(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyId string,
	issuer string,
	serialNumber string,
	vid int32,
) {
	// Noc certificates indexes checks
	ensureCertificatePresentInNocCertificateIndexes(t, setup, subject, subjectKeyId, serialNumber, vid, true, false)

	// All certificates indexes checks
	ensureCertificatePresentInGlobalCertificateIndexes(t, setup, subject, subjectKeyId, serialNumber, false)

	// UniqueCertificate: check that unique certificate key registered
	ensureCertificatePresentInUniqueCertificateIndexes(t, setup, issuer, serialNumber)
}

func ensureNocIcaCertificateExist(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyId string,
	issuer string,
	serialNumber string,
	vid int32,
	skipCheckByVid bool,
) {
	// Noc certificates indexes checks
	ensureCertificatePresentInNocCertificateIndexes(t, setup, subject, subjectKeyId, serialNumber, vid, false, skipCheckByVid)

	// All certificates indexes checks
	ensureCertificatePresentInGlobalCertificateIndexes(t, setup, subject, subjectKeyId, serialNumber, false)

	// UniqueCertificate: check that unique certificate key registered
	ensureCertificatePresentInUniqueCertificateIndexes(t, setup, issuer, serialNumber)
}

func ensureNocIcaCertificateDoesNotExist(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyId string,
	issuer string,
	serialNumber string,
	vid int32,
	skipCheckByVid bool,
	skipCheckForUniqueness bool,
) {
	// Noc certificates indexes checks
	ensureCertificateNotPresentInNocCertificateIndexes(t, setup, subject, subjectKeyId, vid, false, skipCheckByVid)

	// All certificates indexes checks
	ensureCertificateNotPresentInGlobalCertificateIndexes(t, setup, subject, subjectKeyId, false)

	// UniqueCertificate: check that unique certificate key registered
	ensureCertificateNotPresentInUniqueCertificateIndexes(t, setup, issuer, serialNumber, skipCheckForUniqueness)
}

func ensureNocRootCertificateDoesNotExist(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyId string,
	issuer string,
	serialNumber string,
	vid int32,
	skipCheckByVid bool,
	skipCheckForUniqueness bool,
) {
	// Noc certificates indexes checks
	ensureCertificateNotPresentInNocCertificateIndexes(t, setup, subject, subjectKeyId, vid, true, skipCheckByVid)

	// All certificates indexes checks
	ensureCertificateNotPresentInGlobalCertificateIndexes(t, setup, subject, subjectKeyId, false)

	// UniqueCertificate: check that unique certificate key registered
	ensureCertificateNotPresentInUniqueCertificateIndexes(t, setup, issuer, serialNumber, skipCheckForUniqueness)
}

func addDaPaiCertificate(setup *TestSetup, address sdk.AccAddress, pemCert string) {
	addX509Cert := types.NewMsgAddX509Cert(address.String(), pemCert, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(setup.T, err)
}

func addNocRootCertificate(setup *TestSetup, address sdk.AccAddress, pemCert string) {
	// add the new NOC root certificate
	addNocX509RootCert := types.NewMsgAddNocX509RootCert(address.String(), pemCert, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addNocX509RootCert)
	require.NoError(setup.T, err)
}

func addNocIcaCertificate(setup *TestSetup, address sdk.AccAddress, pemCert string) {
	// add the new NOC root certificate
	nocX509Cert := types.NewMsgAddNocX509IcaCert(address.String(), pemCert, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, nocX509Cert)
	require.NoError(setup.T, err)
}
