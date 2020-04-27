package keeper

import (
	test_constants "git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/pagination"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"testing"
)

func TestQuerier_QueryProposedX509RootCert(t *testing.T) {
	setup := Setup()

	// store pending certificate
	certificate := DefaultPendingRootCertificate()
	setup.PkiKeeper.SetProposedCertificate(setup.Ctx, certificate)

	// query proposed certificate
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryProposedX509RootCert, test_constants.RootSubject, test_constants.RootSubjectKeyId},
		abci.RequestQuery{},
	)

	var pendingCertificate types.ProposedCertificate
	_ = setup.Cdc.UnmarshalJSON(result, &pendingCertificate)

	// check
	require.Equal(t, pendingCertificate.PemCert, certificate.PemCert)
	require.Equal(t, pendingCertificate.Subject, certificate.Subject)
	require.Equal(t, pendingCertificate.SubjectKeyId, certificate.SubjectKeyId)
}

func TestQuerier_QueryProposedX509RootCertForNotFound(t *testing.T) {
	setup := Setup()

	// query proposed certificate
	_, err := setup.Querier(
		setup.Ctx,
		[]string{QueryProposedX509RootCert, test_constants.RootSubject, test_constants.RootSubjectKeyId},
		abci.RequestQuery{},
	)

	// check
	require.NotNil(t, err)
	require.Equal(t, types.CodePendingCertificateDoesNotExist, err.Code())
}

func TestQuerier_QueryX509Cert(t *testing.T) {
	setup := Setup()

	// store certificate
	certificate := DefaultRootCertificate()
	setup.PkiKeeper.SetCertificate(setup.Ctx, certificate)

	// query certificate
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryX509Cert, test_constants.RootSubject, test_constants.RootSubjectKeyId},
		abci.RequestQuery{},
	)

	var receivedCertificates types.Certificates
	_ = setup.Cdc.UnmarshalJSON(result, &receivedCertificates)

	// check
	require.Equal(t, 1, len(receivedCertificates.Items))
	receivedCertificate := receivedCertificates.Items[0]

	require.Equal(t, certificate.PemCert, receivedCertificate.PemCert)
	require.Equal(t, certificate.Subject, receivedCertificate.Subject)
	require.Equal(t, certificate.SubjectKeyId, receivedCertificate.SubjectKeyId)
}

func TestQuerier_QueryX509CertForNotFound(t *testing.T) {
	setup := Setup()

	// query proposed certificate
	_, err := setup.Querier(
		setup.Ctx,
		[]string{QueryX509Cert, test_constants.RootSubject, test_constants.RootSubjectKeyId},
		abci.RequestQuery{},
	)

	// check
	require.NotNil(t, err)
	require.Equal(t, types.CodeCertificateDoesNotExist, err.Code())
}

func TestQuerier_QueryAllProposedX509RootCerts(t *testing.T) {
	setup := Setup()
	count := 9

	// populate store with different certificates
	_, _, firstId := PopulateStoreWithMixedCertificates(setup, count)

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllProposedX509RootCerts},
		abci.RequestQuery{Data: emptyQueryParams(setup)},
	)

	var listProposedCertificates types.ListProposedCertificates
	_ = setup.Cdc.UnmarshalJSON(result, &listProposedCertificates)

	// check
	expectedCertificate := DefaultPendingRootCertificate()
	require.Equal(t, count/3, len(listProposedCertificates.Items))
	require.Equal(t, expectedCertificate.PemCert, listProposedCertificates.Items[0].PemCert)
	require.Equal(t, string(firstId), listProposedCertificates.Items[0].Subject)
	require.Equal(t, string(firstId), listProposedCertificates.Items[0].SubjectKeyId)
}

func TestQuerier_QueryAllProposedX509RootCertsWithPagination(t *testing.T) {
	setup := Setup()
	count := 9
	skip := 1
	take := 2

	// populate store with different certificates
	_, _, firstId := PopulateStoreWithMixedCertificates(setup, count)

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllProposedX509RootCerts},
		abci.RequestQuery{Data: queryParams(setup, skip, take)},
	)

	var listProposedCertificates types.ListProposedCertificates
	_ = setup.Cdc.UnmarshalJSON(result, &listProposedCertificates)

	// check
	expectedCertificate := DefaultPendingRootCertificate()
	require.Equal(t, take, len(listProposedCertificates.Items))
	require.Equal(t, expectedCertificate.PemCert, listProposedCertificates.Items[0].PemCert)
	require.Equal(t, string(firstId+skip), listProposedCertificates.Items[0].Subject)
	require.Equal(t, string(firstId+skip), listProposedCertificates.Items[0].SubjectKeyId)
}

func TestQuerier_QueryAllX509RootCerts(t *testing.T) {
	setup := Setup()
	count := 9

	// populate store with different certificates
	firstId, _, _ := PopulateStoreWithMixedCertificates(setup, count)

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllX509RootCerts},
		abci.RequestQuery{Data: emptyQueryParams(setup)},
	)

	var listCertificates types.ListCertificates
	_ = setup.Cdc.UnmarshalJSON(result, &listCertificates)

	// check
	expectedCertificate := DefaultRootCertificate()
	require.Equal(t, count/3, len(listCertificates.Items))
	require.Equal(t, expectedCertificate.PemCert, listCertificates.Items[0].PemCert)
	require.Equal(t, string(firstId), listCertificates.Items[0].Subject)
	require.Equal(t, string(firstId), listCertificates.Items[0].SubjectKeyId)
}

func TestQuerier_QueryAllX509RootCertsWithPagination(t *testing.T) {
	setup := Setup()
	count := 9
	skip := 1
	take := 2

	// populate store with different certificates
	firstId, _, _ := PopulateStoreWithMixedCertificates(setup, count)

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllX509RootCerts},
		abci.RequestQuery{Data: queryParams(setup, skip, take)},
	)

	var listCertificates types.ListCertificates
	_ = setup.Cdc.UnmarshalJSON(result, &listCertificates)

	// check
	expectedCertificate := DefaultRootCertificate()
	require.Equal(t, take, len(listCertificates.Items))
	require.Equal(t, expectedCertificate.PemCert, listCertificates.Items[0].PemCert)
	require.Equal(t, string(firstId+skip), listCertificates.Items[0].Subject)
	require.Equal(t, string(firstId+skip), listCertificates.Items[0].SubjectKeyId)
}

func TestQuerier_QueryAllX509Certs(t *testing.T) {
	setup := Setup()
	count := 9

	// populate store with different certificates
	firstRootId, firstIntermediateId, _ := PopulateStoreWithMixedCertificates(setup, count)

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllX509Certs},
		abci.RequestQuery{Data: emptyQueryParams(setup)},
	)

	var listCertificates types.ListCertificates
	_ = setup.Cdc.UnmarshalJSON(result, &listCertificates)

	// check
	require.Equal(t, count/3*2, len(listCertificates.Items))

	// check first root
	expectedCertificate := DefaultRootCertificate()
	require.Equal(t, expectedCertificate.PemCert, listCertificates.Items[0].PemCert)
	require.Equal(t, string(firstRootId), listCertificates.Items[0].Subject)
	require.Equal(t, string(firstRootId), listCertificates.Items[0].SubjectKeyId)

	// check first intermediate
	index := count / 3
	expectedCertificate = DefaultIntermediateCertificate()
	require.Equal(t, expectedCertificate.PemCert, listCertificates.Items[index].PemCert)
	require.Equal(t, string(firstIntermediateId), listCertificates.Items[index].Subject)
	require.Equal(t, string(firstIntermediateId), listCertificates.Items[index].SubjectKeyId)
}

func TestQuerier_QueryAllX509CertsWithPagination(t *testing.T) {
	setup := Setup()
	count := 9
	skip := 1
	take := 2

	// populate store with different certificates
	firstRootId, _, _ := PopulateStoreWithMixedCertificates(setup, count)

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllX509Certs},
		abci.RequestQuery{Data: queryParams(setup, skip, take)},
	)

	var listCertificates types.ListCertificates
	_ = setup.Cdc.UnmarshalJSON(result, &listCertificates)

	// check
	require.Equal(t, take, len(listCertificates.Items))
	expectedCertificate := DefaultRootCertificate()
	require.Equal(t, expectedCertificate.PemCert, listCertificates.Items[0].PemCert)
	require.Equal(t, string(firstRootId+skip), listCertificates.Items[0].Subject)
	require.Equal(t, string(firstRootId+skip), listCertificates.Items[0].SubjectKeyId)
}

func TestQuerier_QueryAllSubjectX509Certs(t *testing.T) {
	setup := Setup()
	count := 9

	// populate store with different certificates
	firstRootId, _, _ := PopulateStoreWithMixedCertificates(setup, count)

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllSubjectX509Certs, string(firstRootId)},
		abci.RequestQuery{Data: emptyQueryParams(setup)},
	)

	var listCertificates types.ListCertificates
	_ = setup.Cdc.UnmarshalJSON(result, &listCertificates)

	// check
	require.Equal(t, 1, len(listCertificates.Items))

	// check first root
	expectedCertificate := DefaultRootCertificate()
	require.Equal(t, expectedCertificate.PemCert, listCertificates.Items[0].PemCert)
	require.Equal(t, string(firstRootId), listCertificates.Items[0].Subject)
	require.Equal(t, string(firstRootId), listCertificates.Items[0].SubjectKeyId)
}

func TestQuerier_QueryAllSubjectX509Certs_Filer(t *testing.T) {
	setup := Setup()
	count := 9

	// populate store with different certificates
	firstRootId, _, _ := PopulateStoreWithMixedCertificates(setup, count)

	paginationParams := pagination.NewPaginationParams(0, 0)

	params, _ := setup.Cdc.MarshalJSON(types.NewListCertificatesQueryParams(paginationParams, string(firstRootId+1), ""))

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllSubjectX509Certs, string(firstRootId)},
		abci.RequestQuery{Data: params},
	)

	var listCertificates types.ListCertificates
	_ = setup.Cdc.UnmarshalJSON(result, &listCertificates)

	// check
	require.Equal(t, 0, len(listCertificates.Items))
}

func emptyQueryParams(setup TestSetup) []byte {
	paginationParams := pagination.NewPaginationParams(0, 0)
	res, _ := setup.Cdc.MarshalJSON(types.NewListCertificatesQueryParams(paginationParams, "", ""))
	return res
}

func queryParams(setup TestSetup, skip int, take int) []byte {
	paginationParams := pagination.NewPaginationParams(skip, take)
	res, _ := setup.Cdc.MarshalJSON(types.NewListCertificatesQueryParams(paginationParams, "", ""))
	return res
}
