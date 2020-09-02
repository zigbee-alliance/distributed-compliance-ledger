//nolint:testpackage
package keeper

import (
	"testing"

	testconstants "git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/pagination"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
)

func TestQuerier_QueryProposedX509RootCert(t *testing.T) {
	setup := Setup()

	// store proposed certificate
	proposedCertificate := DefaultProposedRootCertificate()
	setup.PkiKeeper.SetProposedCertificate(setup.Ctx, proposedCertificate)

	// query proposed certificate
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryProposedX509RootCert, testconstants.RootSubject, testconstants.RootSubjectKeyID},
		abci.RequestQuery{},
	)

	var receivedProposedCertificate types.ProposedCertificate
	_ = setup.Cdc.UnmarshalJSON(result, &receivedProposedCertificate)

	// check
	require.Equal(t, receivedProposedCertificate.PemCert, proposedCertificate.PemCert)
	require.Equal(t, receivedProposedCertificate.Subject, proposedCertificate.Subject)
	require.Equal(t, receivedProposedCertificate.SubjectKeyID, proposedCertificate.SubjectKeyID)
}

func TestQuerier_QueryProposedX509RootCertForNotFound(t *testing.T) {
	setup := Setup()

	// query proposed certificate
	_, err := setup.Querier(
		setup.Ctx,
		[]string{QueryProposedX509RootCert, testconstants.RootSubject, testconstants.RootSubjectKeyID},
		abci.RequestQuery{},
	)

	// check
	require.NotNil(t, err)
	require.Equal(t, types.CodeProposedCertificateDoesNotExist, err.Code())
}

func TestQuerier_QueryX509Cert(t *testing.T) {
	setup := Setup()

	// store certificate
	certificate := DefaultRootCertificate()
	setup.PkiKeeper.AddApprovedCertificate(setup.Ctx, certificate)

	// query certificate
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryX509Cert, testconstants.RootSubject, testconstants.RootSubjectKeyID},
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

	// query certificate
	_, err := setup.Querier(
		setup.Ctx,
		[]string{QueryX509Cert, testconstants.RootSubject, testconstants.RootSubjectKeyID},
		abci.RequestQuery{},
	)

	// check
	require.NotNil(t, err)
	require.Equal(t, types.CodeCertificateDoesNotExist, err.Code())
}

// nolint:dupl
func TestQuerier_QueryAllProposedX509RootCerts(t *testing.T) {
	setup := Setup()

	// populate store with different certificates
	genCerts := setup.PopulateStoreWithMixedCertificates()

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllProposedX509RootCerts},
		abci.RequestQuery{Data: emptyParams(setup)},
	)

	var listProposedCertificates types.ListProposedCertificates
	_ = setup.Cdc.UnmarshalJSON(result, &listProposedCertificates)

	// check
	require.Equal(t, len(genCerts.ProposedRoots), len(listProposedCertificates.Items))

	for i := 0; i < len(genCerts.ProposedRoots); i++ {
		require.Equal(t, genCerts.ProposedRoots[i].PemCert, listProposedCertificates.Items[i].PemCert)
		require.Equal(t, genCerts.ProposedRoots[i].Subject, listProposedCertificates.Items[i].Subject)
		require.Equal(t, genCerts.ProposedRoots[i].SubjectKeyID, listProposedCertificates.Items[i].SubjectKeyID)
	}
}

func TestQuerier_QueryAllProposedX509RootCertsWithPagination(t *testing.T) {
	setup := Setup()
	skip := 1
	take := 2

	// populate store with different certificates
	genCerts := setup.PopulateStoreWithMixedCertificates()

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllProposedX509RootCerts},
		abci.RequestQuery{Data: paging(setup, skip, take)},
	)

	var listProposedCertificates types.ListProposedCertificates
	_ = setup.Cdc.UnmarshalJSON(result, &listProposedCertificates)

	// check
	require.Equal(t, take, len(listProposedCertificates.Items))

	for i := 0; i < take; i++ {
		require.Equal(t, genCerts.ProposedRoots[skip+i].PemCert, listProposedCertificates.Items[i].PemCert)
		require.Equal(t, genCerts.ProposedRoots[skip+i].Subject, listProposedCertificates.Items[i].Subject)
		require.Equal(t, genCerts.ProposedRoots[skip+i].SubjectKeyID, listProposedCertificates.Items[i].SubjectKeyID)
	}
}

// nolint:dupl
func TestQuerier_QueryAllX509RootCerts(t *testing.T) {
	setup := Setup()

	// populate store with different certificates
	genCerts := setup.PopulateStoreWithMixedCertificates()

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllX509RootCerts},
		abci.RequestQuery{Data: emptyParams(setup)},
	)

	var listCertificates types.ListCertificates
	_ = setup.Cdc.UnmarshalJSON(result, &listCertificates)

	// check
	require.Equal(t, len(genCerts.ApprovedRoots), len(listCertificates.Items))

	for i := 0; i < len(genCerts.ApprovedRoots); i++ {
		require.Equal(t, genCerts.ApprovedRoots[i].PemCert, listCertificates.Items[i].PemCert)
		require.Equal(t, genCerts.ApprovedRoots[i].Subject, listCertificates.Items[i].Subject)
		require.Equal(t, genCerts.ApprovedRoots[i].SubjectKeyID, listCertificates.Items[i].SubjectKeyID)
	}
}

func TestQuerier_QueryAllX509RootCertsWithPagination(t *testing.T) {
	setup := Setup()
	skip := 1
	take := 2

	// populate store with different certificates
	genCerts := setup.PopulateStoreWithMixedCertificates()

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllX509RootCerts},
		abci.RequestQuery{Data: paging(setup, skip, take)},
	)

	var listCertificates types.ListCertificates
	_ = setup.Cdc.UnmarshalJSON(result, &listCertificates)

	// check
	require.Equal(t, take, len(listCertificates.Items))

	for i := 0; i < take; i++ {
		require.Equal(t, genCerts.ApprovedRoots[skip+i].PemCert, listCertificates.Items[i].PemCert)
		require.Equal(t, genCerts.ApprovedRoots[skip+i].Subject, listCertificates.Items[i].Subject)
		require.Equal(t, genCerts.ApprovedRoots[skip+i].SubjectKeyID, listCertificates.Items[i].SubjectKeyID)
	}
}

func TestQuerier_QueryAllX509Certs(t *testing.T) {
	setup := Setup()

	// populate store with different certificates
	genCerts := setup.PopulateStoreWithMixedCertificates()

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllX509Certs},
		abci.RequestQuery{Data: emptyParams(setup)},
	)

	var listCertificates types.ListCertificates
	_ = setup.Cdc.UnmarshalJSON(result, &listCertificates)

	// check
	allApproved := CombineCertLists(genCerts.ApprovedRoots, genCerts.ApprovedNonRoots)

	require.Equal(t, len(allApproved), len(listCertificates.Items))

	for i := 0; i < len(allApproved); i++ {
		require.Equal(t, allApproved[i].PemCert, listCertificates.Items[i].PemCert)
		require.Equal(t, allApproved[i].Subject, listCertificates.Items[i].Subject)
		require.Equal(t, allApproved[i].SubjectKeyID, listCertificates.Items[i].SubjectKeyID)
	}
}

func TestQuerier_QueryAllX509CertsWithPagination(t *testing.T) {
	setup := Setup()
	skip := 1
	take := 2

	// populate store with different certificates
	genCerts := setup.PopulateStoreWithMixedCertificates()

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllX509Certs},
		abci.RequestQuery{Data: paging(setup, skip, take)},
	)

	var listCertificates types.ListCertificates
	_ = setup.Cdc.UnmarshalJSON(result, &listCertificates)

	// check
	allApproved := CombineCertLists(genCerts.ApprovedRoots, genCerts.ApprovedNonRoots)

	require.Equal(t, take, len(listCertificates.Items))

	for i := 0; i < take; i++ {
		require.Equal(t, allApproved[skip+i].PemCert, listCertificates.Items[i].PemCert)
		require.Equal(t, allApproved[skip+i].Subject, listCertificates.Items[i].Subject)
		require.Equal(t, allApproved[skip+i].SubjectKeyID, listCertificates.Items[i].SubjectKeyID)
	}
}

func TestQuerier_QueryAllSubjectX509Certs(t *testing.T) {
	setup := Setup()

	// populate store with different certificates
	setup.PopulateStoreWithMixedCertificates()

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllSubjectX509Certs, "DN105"},
		abci.RequestQuery{Data: emptyParams(setup)},
	)

	var listCertificates types.ListCertificates
	_ = setup.Cdc.UnmarshalJSON(result, &listCertificates)

	// check
	require.Equal(t, 1, len(listCertificates.Items))

	require.Equal(t, testconstants.StubCertPem, listCertificates.Items[0].PemCert)
	require.Equal(t, "DN105", listCertificates.Items[0].Subject)
	require.Equal(t, "KeyID105", listCertificates.Items[0].SubjectKeyID)
}

func TestQuerier_QueryAllX509Certs_Filter(t *testing.T) {
	setup := Setup()

	// populate store with different certificates
	setup.PopulateStoreWithMixedCertificates()

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllX509Certs},
		abci.RequestQuery{Data: pagingAndFilter(setup, 0, 0, "DN104", "KeyID104")},
	)

	var listCertificates types.ListCertificates
	_ = setup.Cdc.UnmarshalJSON(result, &listCertificates)

	// check
	require.Equal(t, 3, len(listCertificates.Items))

	require.Equal(t, testconstants.StubCertPem, listCertificates.Items[0].PemCert)
	require.Equal(t, "DN104", listCertificates.Items[0].Subject)
	require.Equal(t, "KeyID104", listCertificates.Items[0].SubjectKeyID)

	require.Equal(t, testconstants.StubCertPem, listCertificates.Items[1].PemCert)
	require.Equal(t, "DN105", listCertificates.Items[1].Subject)
	require.Equal(t, "KeyID105", listCertificates.Items[1].SubjectKeyID)

	require.Equal(t, testconstants.StubCertPem, listCertificates.Items[2].PemCert)
	require.Equal(t, "DN106", listCertificates.Items[2].Subject)
	require.Equal(t, "KeyID106", listCertificates.Items[2].SubjectKeyID)
}

func TestQuerier_QueryProposedX509RootCertRevocation(t *testing.T) {
	setup := Setup()

	// store proposed certificate revocation
	revocation := DefaultProposedRootCertificateRevocation()
	setup.PkiKeeper.SetProposedCertificateRevocation(setup.Ctx, revocation)

	// query proposed certificate revocation
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryProposedX509RootCertRevocation, testconstants.RootSubject, testconstants.RootSubjectKeyID},
		abci.RequestQuery{},
	)

	var receivedRevocation types.ProposedCertificateRevocation
	_ = setup.Cdc.UnmarshalJSON(result, &receivedRevocation)

	// check
	require.Equal(t, receivedRevocation.Subject, revocation.Subject)
	require.Equal(t, receivedRevocation.SubjectKeyID, revocation.SubjectKeyID)
}

func TestQuerier_QueryProposedX509RootCertRevocationForNotFound(t *testing.T) {
	setup := Setup()

	// query proposed certificate revocation
	_, err := setup.Querier(
		setup.Ctx,
		[]string{QueryProposedX509RootCertRevocation, testconstants.RootSubject, testconstants.RootSubjectKeyID},
		abci.RequestQuery{},
	)

	// check
	require.NotNil(t, err)
	require.Equal(t, types.CodeProposedCertificateRevocationDoesNotExist, err.Code())
}

func TestQuerier_QueryRevokedX509Cert(t *testing.T) {
	setup := Setup()

	// store revoked certificate
	certificate := DefaultRootCertificate()
	setup.PkiKeeper.AddRevokedCertificates(setup.Ctx, certificate.Subject, certificate.SubjectKeyID,
		types.NewCertificates([]types.Certificate{certificate}))

	// query revoked certificate
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryRevokedX509Cert, testconstants.RootSubject, testconstants.RootSubjectKeyID},
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

func TestQuerier_QueryRevokedX509CertForNotFound(t *testing.T) {
	setup := Setup()

	// query revoked certificate
	_, err := setup.Querier(
		setup.Ctx,
		[]string{QueryRevokedX509Cert, testconstants.RootSubject, testconstants.RootSubjectKeyID},
		abci.RequestQuery{},
	)

	// check
	require.NotNil(t, err)
	require.Equal(t, types.CodeRevokedCertificateDoesNotExist, err.Code())
}

// nolint:dupl
func TestQuerier_QueryAllProposedX509RootCertRevocations(t *testing.T) {
	setup := Setup()

	// populate store with different certificates
	genCerts := setup.PopulateStoreWithMixedCertificates()

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllProposedX509RootCertRevocations},
		abci.RequestQuery{Data: emptyParams(setup)},
	)

	var listProposedRevocations types.ListProposedCertificateRevocations
	_ = setup.Cdc.UnmarshalJSON(result, &listProposedRevocations)

	// check
	require.Equal(t, len(genCerts.ProposedRootRevocations), len(listProposedRevocations.Items))

	for i := 0; i < len(genCerts.ProposedRootRevocations); i++ {
		require.Equal(t, genCerts.ProposedRootRevocations[i].Subject, listProposedRevocations.Items[i].Subject)
		require.Equal(t, genCerts.ProposedRootRevocations[i].SubjectKeyID, listProposedRevocations.Items[i].SubjectKeyID)
		require.Equal(t, genCerts.ProposedRootRevocations[i].Approvals, listProposedRevocations.Items[i].Approvals)
	}
}

func TestQuerier_QueryAllProposedX509RootCertRevocationsWithPagination(t *testing.T) {
	setup := Setup()
	skip := 1
	take := 1

	// populate store with different certificates
	genCerts := setup.PopulateStoreWithMixedCertificates()

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllProposedX509RootCertRevocations},
		abci.RequestQuery{Data: paging(setup, skip, take)},
	)

	var listProposedRevocations types.ListProposedCertificateRevocations
	_ = setup.Cdc.UnmarshalJSON(result, &listProposedRevocations)

	// check
	require.Equal(t, take, len(listProposedRevocations.Items))

	for i := 0; i < take; i++ {
		require.Equal(t, genCerts.ProposedRootRevocations[skip+i].Subject, listProposedRevocations.Items[i].Subject)
		require.Equal(t, genCerts.ProposedRootRevocations[skip+i].SubjectKeyID, listProposedRevocations.Items[i].SubjectKeyID)
		require.Equal(t, genCerts.ProposedRootRevocations[skip+i].Approvals, listProposedRevocations.Items[i].Approvals)
	}
}

// nolint:dupl
func TestQuerier_QueryAllRevokedX509RootCerts(t *testing.T) {
	setup := Setup()

	// populate store with different certificates
	genCerts := setup.PopulateStoreWithMixedCertificates()

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllRevokedX509RootCerts},
		abci.RequestQuery{Data: emptyParams(setup)},
	)

	var listCertificates types.ListCertificates
	_ = setup.Cdc.UnmarshalJSON(result, &listCertificates)

	// check
	require.Equal(t, len(genCerts.RevokedRoots), len(listCertificates.Items))

	for i := 0; i < len(genCerts.RevokedRoots); i++ {
		require.Equal(t, genCerts.RevokedRoots[i].PemCert, listCertificates.Items[i].PemCert)
		require.Equal(t, genCerts.RevokedRoots[i].Subject, listCertificates.Items[i].Subject)
		require.Equal(t, genCerts.RevokedRoots[i].SubjectKeyID, listCertificates.Items[i].SubjectKeyID)
	}
}

func TestQuerier_QueryAllRevokedX509RootCertsWithPagination(t *testing.T) {
	setup := Setup()
	skip := 1
	take := 1

	// populate store with different certificates
	genCerts := setup.PopulateStoreWithMixedCertificates()

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllRevokedX509RootCerts},
		abci.RequestQuery{Data: paging(setup, skip, take)},
	)

	var listCertificates types.ListCertificates
	_ = setup.Cdc.UnmarshalJSON(result, &listCertificates)

	// check
	require.Equal(t, take, len(listCertificates.Items))

	for i := 0; i < take; i++ {
		require.Equal(t, genCerts.RevokedRoots[skip+i].PemCert, listCertificates.Items[i].PemCert)
		require.Equal(t, genCerts.RevokedRoots[skip+i].Subject, listCertificates.Items[i].Subject)
		require.Equal(t, genCerts.RevokedRoots[skip+i].SubjectKeyID, listCertificates.Items[i].SubjectKeyID)
	}
}

func TestQuerier_QueryAllRevokedX509Certs(t *testing.T) {
	setup := Setup()

	// populate store with different certificates
	genCerts := setup.PopulateStoreWithMixedCertificates()

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllRevokedX509Certs},
		abci.RequestQuery{Data: emptyParams(setup)},
	)

	var listCertificates types.ListCertificates
	_ = setup.Cdc.UnmarshalJSON(result, &listCertificates)

	// check
	allRevoked := CombineCertLists(genCerts.RevokedRoots, genCerts.RevokedNonRoots)

	require.Equal(t, len(allRevoked), len(listCertificates.Items))

	for i := 0; i < len(allRevoked); i++ {
		require.Equal(t, allRevoked[i].PemCert, listCertificates.Items[i].PemCert)
		require.Equal(t, allRevoked[i].Subject, listCertificates.Items[i].Subject)
		require.Equal(t, allRevoked[i].SubjectKeyID, listCertificates.Items[i].SubjectKeyID)
	}
}

func TestQuerier_QueryAllRevokedX509CertsWithPagination(t *testing.T) {
	setup := Setup()
	skip := 1
	take := 2

	// populate store with different certificates
	genCerts := setup.PopulateStoreWithMixedCertificates()

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllRevokedX509Certs},
		abci.RequestQuery{Data: paging(setup, skip, take)},
	)

	var listCertificates types.ListCertificates
	_ = setup.Cdc.UnmarshalJSON(result, &listCertificates)

	// check
	allRevoked := CombineCertLists(genCerts.RevokedRoots, genCerts.RevokedNonRoots)

	require.Equal(t, take, len(listCertificates.Items))

	for i := 0; i < take; i++ {
		require.Equal(t, allRevoked[skip+i].PemCert, listCertificates.Items[i].PemCert)
		require.Equal(t, allRevoked[skip+i].Subject, listCertificates.Items[i].Subject)
		require.Equal(t, allRevoked[skip+i].SubjectKeyID, listCertificates.Items[i].SubjectKeyID)
	}
}

func TestQuerier_QueryAllRevokedX509Certs_Filter(t *testing.T) {
	setup := Setup()

	// populate store with different certificates
	setup.PopulateStoreWithMixedCertificates()

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllRevokedX509Certs},
		abci.RequestQuery{Data: pagingAndFilter(setup, 0, 0, "DN104", "KeyID104")},
	)

	var listCertificates types.ListCertificates
	_ = setup.Cdc.UnmarshalJSON(result, &listCertificates)

	// check
	require.Equal(t, 2, len(listCertificates.Items))

	require.Equal(t, testconstants.StubCertPem, listCertificates.Items[0].PemCert)
	require.Equal(t, "DN107", listCertificates.Items[0].Subject)
	require.Equal(t, "KeyID107", listCertificates.Items[0].SubjectKeyID)

	require.Equal(t, testconstants.StubCertPem, listCertificates.Items[1].PemCert)
	require.Equal(t, "DN108", listCertificates.Items[1].Subject)
	require.Equal(t, "KeyID108", listCertificates.Items[1].SubjectKeyID)
}

func emptyParams(setup TestSetup) []byte {
	return paging(setup, 0, 0)
}

func paging(setup TestSetup, skip int, take int) []byte {
	return pagingAndFilter(setup, skip, take, "", "")
}

func pagingAndFilter(setup TestSetup, skip int, take int, rootSubject string, rootSubjectKeyID string) []byte {
	paginationParams := pagination.NewPaginationParams(skip, take)

	return setup.Cdc.MustMarshalJSON(types.NewPkiQueryParams(paginationParams, rootSubject, rootSubjectKeyID))
}
