package keeper

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

type TestSetup struct {
	Cdc       *codec.Codec
	Ctx       sdk.Context
	PkiKeeper Keeper
	Querier   sdk.Querier
}

func Setup() TestSetup {
	// Init Codec
	cdc := codec.New()
	sdk.RegisterCodec(cdc)

	// Init KVSore
	db := dbm.NewMemDB()
	dbStore := store.NewCommitMultiStore(db)
	pkiKey := sdk.NewKVStoreKey(types.StoreKey)
	dbStore.MountStoreWithDB(pkiKey, sdk.StoreTypeIAVL, db)
	_ = dbStore.LoadLatestVersion()

	// Init Keepers
	pkiKeeper := NewKeeper(pkiKey, cdc)

	// Init Querier
	querier := NewQuerier(pkiKeeper)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: test_constants.ChainId}, false, log.NewNopLogger())

	setup := TestSetup{
		Cdc:       cdc,
		Ctx:       ctx,
		PkiKeeper: pkiKeeper,
		Querier:   querier,
	}
	return setup
}

func DefaultIntermediateCertificate() types.Certificate {
	return types.NewIntermediateCertificate(
		test_constants.LeafCertPem,
		test_constants.LeafSubject,
		test_constants.LeafSubjectKeyId,
		test_constants.LeafSerialNumber,
		test_constants.RootSubjectKeyId,
		test_constants.Address1)
}

func DefaultRootCertificate() types.Certificate {
	return types.NewRootCertificate(
		test_constants.RootCertPem,
		test_constants.RootSubject,
		test_constants.RootSubjectKeyId,
		test_constants.RootSerialNumber,
		test_constants.Address1)
}

func DefaultPendingRootCertificate() types.ProposedCertificate {
	return types.NewProposedCertificate(
		test_constants.RootCertPem,
		test_constants.RootSubject,
		test_constants.RootSubjectKeyId,
		test_constants.RootSerialNumber,
		test_constants.Address1)
}

// add n Mixed Certificates into store {SubjectKeyId: "1".."n"}
func PopulateStoreWithMixedCertificates(setup TestSetup, count int) (int, int, int) {
	n := count / 3
	firstId := 1
	firstIdRoot := firstId
	firstIdLeaf := firstId + n
	firstIdPending := firstId + n*2
	populateStoreWithCertificates(setup, n, DefaultRootCertificate(), firstIdRoot)
	populateStoreWithCertificates(setup, n+n, DefaultIntermediateCertificate(), firstIdLeaf)
	populateStoreWithPendingCertificates(setup, n+n*2, DefaultPendingRootCertificate(), firstIdPending)
	return firstIdRoot, firstIdLeaf, firstIdPending
}

// add n Certificates into store {SubjectKeyId: "1".."n"}
func populateStoreWithCertificates(setup TestSetup, count int, certificate types.Certificate, firstId int) int {
	for i := firstId; i <= count; i++ {
		certificate.Subject = string(i)
		certificate.SubjectKeyId = string(i)
		certificate.SerialNumber = string(i)
		certificate.RootSubjectId = string(i)
		setup.PkiKeeper.SetCertificate(setup.Ctx, certificate)
	}
	return firstId
}

// add n Pending Root Certificates into store {SubjectKeyId: "1".."n"}
func populateStoreWithPendingCertificates(setup TestSetup, count int, certificate types.ProposedCertificate, firstId int) int {
	for i := firstId; i <= count; i++ {
		certificate.Subject = string(i)
		certificate.SubjectKeyId = string(i)
		certificate.SerialNumber = string(i)
		setup.PkiKeeper.SetProposedCertificate(setup.Ctx, certificate)
	}
	return firstId
}
