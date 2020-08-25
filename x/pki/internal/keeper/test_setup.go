package keeper

//nolint:goimports
import (
	testconstants "git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
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
	dbStore.MountStoreWithDB(pkiKey, sdk.StoreTypeIAVL, nil)
	_ = dbStore.LoadLatestVersion()

	// Init Keepers
	pkiKeeper := NewKeeper(pkiKey, cdc)

	// Init Querier
	querier := NewQuerier(pkiKeeper)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: testconstants.ChainID}, false, log.NewNopLogger())

	setup := TestSetup{
		Cdc:       cdc,
		Ctx:       ctx,
		PkiKeeper: pkiKeeper,
		Querier:   querier,
	}

	return setup
}

func DefaultLeafCertificate() types.Certificate {
	return types.NewNonRootCertificate(
		testconstants.LeafCertPem,
		testconstants.LeafSubject,
		testconstants.LeafSubjectKeyID,
		testconstants.LeafSerialNumber,
		testconstants.LeafIssuer,
		testconstants.LeafAuthorityKeyID,
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.Address1)
}

func DefaultRootCertificate() types.Certificate {
	return types.NewRootCertificate(
		testconstants.RootCertPem,
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.RootSerialNumber,
		testconstants.Address1)
}

func DefaultProposedRootCertificate() types.ProposedCertificate {
	return types.NewProposedCertificate(
		testconstants.RootCertPem,
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.RootSerialNumber,
		testconstants.Address1)
}

// add n Mixed Certificates into store {SubjectKeyID: "1".."n"}.
func PopulateStoreWithMixedCertificates(setup TestSetup, count int) (int, int, int) {
	n := count / 3
	firstID := 1
	firstIDRoot := firstID
	firstIDLeaf := firstID + n
	firstIDProposed := firstID + n*2
	populateStoreWithCertificates(setup, n, DefaultRootCertificate(), firstIDRoot)
	populateStoreWithCertificates(setup, n+n, DefaultLeafCertificate(), firstIDLeaf)
	populateStoreWithProposedCertificates(setup, n+n*2, DefaultProposedRootCertificate(), firstIDProposed)

	return firstIDRoot, firstIDLeaf, firstIDProposed
}

// add n Certificates into store {SubjectKeyID: "1".."n"}.
func populateStoreWithCertificates(setup TestSetup, count int, certificate types.Certificate, firstID int) int {
	for i := firstID; i <= count; i++ {
		certificate.Subject = string(i)
		certificate.SubjectKeyID = string(i)
		certificate.SerialNumber = string(i)
		certificate.RootSubject = string(i)
		certificate.RootSubjectKeyID = string(i)
		setup.PkiKeeper.AddApprovedCertificate(setup.Ctx, certificate)
	}

	return firstID
}

// add n Proposed Root Certificates into store {SubjectKeyID: "1".."n"}.
func populateStoreWithProposedCertificates(setup TestSetup,
	count int, certificate types.ProposedCertificate, firstID int) int {
	for i := firstID; i <= count; i++ {
		certificate.Subject = string(i)
		certificate.SubjectKeyID = string(i)
		certificate.SerialNumber = string(i)
		setup.PkiKeeper.SetProposedCertificate(setup.Ctx, certificate)
	}

	return firstID
}
