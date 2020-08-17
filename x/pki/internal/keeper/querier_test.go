//nolint:testpackage
package keeper

import (
	"testing"

	test_constants "git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/pagination"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
)

func TestQuerier_QueryProposedX509RootCert(t *testing.T) {
	setup := Setup()

	// store pending certificate
	certificate := DefaultPendingRootCertificate()
	setup.PkiKeeper.SetProposedCertificate(setup.Ctx, certificate)

	// query proposed certificate
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryProposedX509RootCert, test_constants.RootSubject, test_constants.RootSubjectKeyID},
		abci.RequestQuery{},
	)

	var pendingCertificate types.ProposedCertificate
	_ = setup.Cdc.UnmarshalJSON(result, &pendingCertificate)

	// check
	require.Equal(t, pendingCertificate.PemCert, certificate.PemCert)
	require.Equal(t, pendingCertificate.Subject, certificate.Subject)
	require.Equal(t, pendingCertificate.SubjectKeyID, certificate.SubjectKeyID)
}

func TestQuerier_QueryProposedX509RootCertForNotFound(t *testing.T) {
	setup := Setup()

	// query proposed certificate
	_, err := setup.Querier(
		setup.Ctx,
		[]string{QueryProposedX509RootCert, test_constants.RootSubject, test_constants.RootSubjectKeyID},
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
		[]string{QueryX509Cert, test_constants.RootSubject, test_constants.RootSubjectKeyID},
		abci.RequestQuery{},
	)

	var receivedCertificates types.Certificates
	_ = setup.Cdc.UnmarshalJSON(result, &receivedCertificates)

	// check
	require.Equal(t, 1, len(receivedCertificates.Items))
	receivedCertificate := receivedCertificates.Items[0]

	require.Equal(t, certificate.PemCert, receivedCertificate.PemCert)
	require.Equal(t, certificate.Subject, receivedCertificate.Subject)
	require.Equal(t, certificate.SubjectKeyID, receivedCertificate.SubjectKeyID)
}

func TestQuerier_QueryX509CertForNotFound(t *testing.T) {
	setup := Setup()

	// query proposed certificate
	_, err := setup.Querier(
		setup.Ctx,
		[]string{QueryX509Cert, test_constants.RootSubject, test_constants.RootSubjectKeyID},
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
	_, _, firstID := PopulateStoreWithMixedCertificates(setup, count)

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllProposedX509RootCerts},
		abci.RequestQuery{Data: emptyQueryParams(setup)},
	)

	var listProposedCertificates types.ListProposedCertificates
	_ = setup.Cdc.UnmarshalJSON(result, &listProposedCertificates)

	// check
	require.Equal(t, count/3, len(listProposedCertificates.Items))
	require.Equal(t, DefaultPendingRootCertificate().PemCert, listProposedCertificates.Items[0].PemCert)
	require.Equal(t, string(firstID), listProposedCertificates.Items[0].Subject)
	require.Equal(t, string(firstID), listProposedCertificates.Items[0].SubjectKeyID)
}

func TestQuerier_QueryAllProposedX509RootCertsWithPagination(t *testing.T) {
	setup := Setup()
	count := 9
	skip := 1
	take := 2

	// populate store with different certificates
	_, _, firstID := PopulateStoreWithMixedCertificates(setup, count)

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllProposedX509RootCerts},
		abci.RequestQuery{Data: queryParams(setup, skip, take)},
	)

	var listProposedCertificates types.ListProposedCertificates
	_ = setup.Cdc.UnmarshalJSON(result, &listProposedCertificates)

	// check
	require.Equal(t, take, len(listProposedCertificates.Items))
	require.Equal(t, DefaultPendingRootCertificate().PemCert, listProposedCertificates.Items[0].PemCert)
	require.Equal(t, string(firstID+skip), listProposedCertificates.Items[0].Subject)
	require.Equal(t, string(firstID+skip), listProposedCertificates.Items[0].SubjectKeyID)
}

func TestQuerier_QueryAllX509RootCerts(t *testing.T) {
	setup := Setup()
	count := 9

	// populate store with different certificates
	firstID, _, _ := PopulateStoreWithMixedCertificates(setup, count)

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllX509RootCerts},
		abci.RequestQuery{Data: emptyQueryParams(setup)},
	)

	var listCertificates types.ListCertificates
	_ = setup.Cdc.UnmarshalJSON(result, &listCertificates)

	// check
	require.Equal(t, count/3, len(listCertificates.Items))
	require.Equal(t, DefaultRootCertificate().PemCert, listCertificates.Items[0].PemCert)
	require.Equal(t, string(firstID), listCertificates.Items[0].Subject)
	require.Equal(t, string(firstID), listCertificates.Items[0].SubjectKeyID)
}

func TestQuerier_QueryAllX509RootCertsWithPagination(t *testing.T) {
	setup := Setup()
	count := 9
	skip := 1
	take := 2

	// populate store with different certificates
	firstID, _, _ := PopulateStoreWithMixedCertificates(setup, count)

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllX509RootCerts},
		abci.RequestQuery{Data: queryParams(setup, skip, take)},
	)

	var listCertificates types.ListCertificates
	_ = setup.Cdc.UnmarshalJSON(result, &listCertificates)

	// check
	require.Equal(t, take, len(listCertificates.Items))
	require.Equal(t, DefaultRootCertificate().PemCert, listCertificates.Items[0].PemCert)
	require.Equal(t, string(firstID+skip), listCertificates.Items[0].Subject)
	require.Equal(t, string(firstID+skip), listCertificates.Items[0].SubjectKeyID)
}

func TestQuerier_QueryAllX509Certs(t *testing.T) {
	setup := Setup()
	count := 9

	// populate store with different certificates
	firstRootID, firstIntermediateID, _ := PopulateStoreWithMixedCertificates(setup, count)

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
	require.Equal(t, string(firstRootID), listCertificates.Items[0].Subject)
	require.Equal(t, string(firstRootID), listCertificates.Items[0].SubjectKeyID)

	// check first intermediate
	index := count / 3
	expectedCertificate = DefaultIntermediateCertificate()
	require.Equal(t, expectedCertificate.PemCert, listCertificates.Items[index].PemCert)
	require.Equal(t, string(firstIntermediateID), listCertificates.Items[index].Subject)
	require.Equal(t, string(firstIntermediateID), listCertificates.Items[index].SubjectKeyID)
}

func TestQuerier_QueryAllX509CertsWithPagination(t *testing.T) {
	setup := Setup()
	count := 9
	skip := 1
	take := 2

	// populate store with different certificates
	firstRootID, _, _ := PopulateStoreWithMixedCertificates(setup, count)

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
	require.Equal(t, DefaultRootCertificate().PemCert, listCertificates.Items[0].PemCert)
	require.Equal(t, string(firstRootID+skip), listCertificates.Items[0].Subject)
	require.Equal(t, string(firstRootID+skip), listCertificates.Items[0].SubjectKeyID)
}

func TestQuerier_QueryAllSubjectX509Certs(t *testing.T) {
	setup := Setup()
	count := 9

	// populate store with different certificates
	firstRootID, _, _ := PopulateStoreWithMixedCertificates(setup, count)

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllSubjectX509Certs, string(firstRootID)},
		abci.RequestQuery{Data: emptyQueryParams(setup)},
	)

	var listCertificates types.ListCertificates
	_ = setup.Cdc.UnmarshalJSON(result, &listCertificates)

	// check
	require.Equal(t, 1, len(listCertificates.Items))

	// check first root
	expectedCertificate := DefaultRootCertificate()
	require.Equal(t, expectedCertificate.PemCert, listCertificates.Items[0].PemCert)
	require.Equal(t, string(firstRootID), listCertificates.Items[0].Subject)
	require.Equal(t, string(firstRootID), listCertificates.Items[0].SubjectKeyID)
}

func TestQuerier_QueryAllSubjectX509Certs_Filer(t *testing.T) {
	setup := Setup()
	count := 9

	// populate store with different certificates
	firstRootID, _, _ := PopulateStoreWithMixedCertificates(setup, count)

	paginationParams := pagination.NewPaginationParams(0, 0)

	params := setup.Cdc.MustMarshalJSON(types.NewListCertificatesQueryParams(paginationParams, string(firstRootID+1), ""))

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllSubjectX509Certs, string(firstRootID)},
		abci.RequestQuery{Data: params},
	)

	var listCertificates types.ListCertificates
	_ = setup.Cdc.UnmarshalJSON(result, &listCertificates)

	// check
	require.Equal(t, 0, len(listCertificates.Items))
}

func emptyQueryParams(setup TestSetup) []byte {
	paginationParams := pagination.NewPaginationParams(0, 0)
	return setup.Cdc.MustMarshalJSON(types.NewListCertificatesQueryParams(paginationParams, "", ""))
}

func queryParams(setup TestSetup, skip int, take int) []byte {
	paginationParams := pagination.NewPaginationParams(skip, take)
	return setup.Cdc.MustMarshalJSON(types.NewListCertificatesQueryParams(paginationParams, "", ""))
}
