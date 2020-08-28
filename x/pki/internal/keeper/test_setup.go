package keeper

import (
	"sort"
	"strconv"

	testconstants "git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

const (
	DN    = "DN"
	KeyID = "KeyID"
)

type GeneratedCertificates struct {
	ProposedRoots           []types.ProposedCertificate
	ApprovedRoots           []types.Certificate
	ApprovedNonRoots        []types.Certificate
	ProposedRootRevocations []types.ProposedCertificateRevocation
	RevokedRoots            []types.Certificate
	RevokedNonRoots         []types.Certificate
}

type Index struct {
	Subject int
}

type Indexes struct {
	Subject int
	Issuer  int
	Root    int
}

type TestSetup struct {
	Cdc       *codec.Codec
	Ctx       sdk.Context
	PkiKeeper Keeper
	Querier   sdk.Querier
}

func (setup TestSetup) PopulateStoreWithMixedCertificates() GeneratedCertificates {
	var genCerts GeneratedCertificates

	// We use the ascending order of indexes and padded numbers for them to ensure the same order of entities
	// in KVStore prefix iterators and in slices of genCerts.

	setup.storeProposedRootCertificate(&genCerts, Index{Subject: 101})
	setup.storeProposedRootCertificate(&genCerts, Index{Subject: 102})
	setup.storeProposedRootCertificate(&genCerts, Index{Subject: 103})

	setup.storeApprovedRootCertificate(&genCerts, Index{Subject: 104})
	{
		setup.storeApprovedNonRootCertificate(&genCerts, Indexes{Subject: 105, Issuer: 104, Root: 104})
		{
			setup.storeApprovedNonRootCertificate(&genCerts, Indexes{Subject: 106, Issuer: 105, Root: 104})
		}
		setup.storeRevokedNonRootCertificate(&genCerts, Indexes{Subject: 107, Issuer: 104, Root: 104})
		{
			setup.storeRevokedNonRootCertificate(&genCerts, Indexes{Subject: 108, Issuer: 107, Root: 104})
		}
	}
	setup.storeApprovedRootCertificate(&genCerts, Index{Subject: 109})
	{
		setup.storeApprovedNonRootCertificate(&genCerts, Indexes{Subject: 110, Issuer: 109, Root: 109})
	}
	setup.storeApprovedRootCertificate(&genCerts, Index{Subject: 111})
	setup.storeRevokedRootCertificate(&genCerts, Index{Subject: 112})
	{
		setup.storeRevokedNonRootCertificate(&genCerts, Indexes{Subject: 113, Issuer: 112, Root: 112})
	}
	setup.storeRevokedRootCertificate(&genCerts, Index{Subject: 114})

	setup.storeProposedRootCertificateRevocation(&genCerts, Index{Subject: 109})
	setup.storeProposedRootCertificateRevocation(&genCerts, Index{Subject: 111})

	return genCerts
}

func (setup TestSetup) storeProposedRootCertificate(genCerts *GeneratedCertificates, index Index) {
	proposedCertificate := createProposedRootCertificate(index)
	setup.PkiKeeper.SetProposedCertificate(setup.Ctx, proposedCertificate)
	genCerts.ProposedRoots = append(genCerts.ProposedRoots, proposedCertificate)
}

func (setup TestSetup) storeApprovedRootCertificate(genCerts *GeneratedCertificates, index Index) {
	certificate := createRootCertificate(index)
	setup.PkiKeeper.AddApprovedCertificate(setup.Ctx, certificate)
	genCerts.ApprovedRoots = append(genCerts.ApprovedRoots, certificate)
}

func (setup TestSetup) storeApprovedNonRootCertificate(genCerts *GeneratedCertificates, indexes Indexes) {
	certificate := createNonRootCertificate(indexes)
	setup.PkiKeeper.AddApprovedCertificate(setup.Ctx, certificate)
	genCerts.ApprovedNonRoots = append(genCerts.ApprovedNonRoots, certificate)
}

func (setup TestSetup) storeProposedRootCertificateRevocation(genCerts *GeneratedCertificates, index Index) {
	revocation := createProposedRootCertificateRevocation(index)
	setup.PkiKeeper.SetProposedCertificateRevocation(setup.Ctx, revocation)
	genCerts.ProposedRootRevocations = append(genCerts.ProposedRootRevocations, revocation)
}

func (setup TestSetup) storeRevokedRootCertificate(genCerts *GeneratedCertificates, index Index) {
	certificate := createRootCertificate(index)

	setup.PkiKeeper.AddRevokedCertificates(setup.Ctx, certificate.Subject, certificate.SubjectKeyID,
		types.NewCertificates([]types.Certificate{certificate}))

	genCerts.RevokedRoots = append(genCerts.RevokedRoots, certificate)
}

func (setup TestSetup) storeRevokedNonRootCertificate(genCerts *GeneratedCertificates, indexes Indexes) {
	certificate := createNonRootCertificate(indexes)

	setup.PkiKeeper.AddRevokedCertificates(setup.Ctx, certificate.Subject, certificate.SubjectKeyID,
		types.NewCertificates([]types.Certificate{certificate}))

	genCerts.RevokedNonRoots = append(genCerts.RevokedNonRoots, certificate)
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

func CombineCertLists(rootCerts []types.Certificate, nonRootCerts []types.Certificate) []types.Certificate {
	concatenation := append(rootCerts, nonRootCerts...)

	sort.Slice(
		concatenation,
		func(i, j int) bool {
			return concatenation[i].SerialNumber < concatenation[j].SerialNumber
		},
	)

	return concatenation
}

func createProposedRootCertificate(index Index) types.ProposedCertificate {
	subjectSuffix := strconv.Itoa(index.Subject)

	return types.ProposedCertificate{
		PemCert:      testconstants.StubCertPem,
		Subject:      DN + subjectSuffix,
		SubjectKeyID: KeyID + subjectSuffix,
		SerialNumber: subjectSuffix,
		Owner:        testconstants.Address1,
		Approvals:    []sdk.AccAddress{},
	}
}

func createRootCertificate(index Index) types.Certificate {
	subjectSuffix := strconv.Itoa(index.Subject)

	return types.Certificate{
		PemCert:      testconstants.StubCertPem,
		Subject:      DN + subjectSuffix,
		SubjectKeyID: KeyID + subjectSuffix,
		SerialNumber: subjectSuffix,
		IsRoot:       true,
		Owner:        testconstants.Address1,
	}
}

func createNonRootCertificate(indexes Indexes) types.Certificate {
	subjectSuffix := strconv.Itoa(indexes.Subject)
	issuerSuffix := strconv.Itoa(indexes.Issuer)
	rootSuffix := strconv.Itoa(indexes.Root)

	return types.Certificate{
		PemCert:          testconstants.StubCertPem,
		Subject:          DN + subjectSuffix,
		SubjectKeyID:     KeyID + subjectSuffix,
		SerialNumber:     subjectSuffix,
		Issuer:           DN + issuerSuffix,
		AuthorityKeyID:   KeyID + issuerSuffix,
		RootSubject:      DN + rootSuffix,
		RootSubjectKeyID: KeyID + rootSuffix,
		IsRoot:           false,
		Owner:            testconstants.Address1,
	}
}

func createProposedRootCertificateRevocation(index Index) types.ProposedCertificateRevocation {
	subjectSuffix := strconv.Itoa(index.Subject)

	return types.ProposedCertificateRevocation{
		Subject:      DN + subjectSuffix,
		SubjectKeyID: KeyID + subjectSuffix,
		Approvals:    []sdk.AccAddress{},
	}
}
