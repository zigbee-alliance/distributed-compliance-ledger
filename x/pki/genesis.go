package pki

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the approvedCertificates
	for _, elem := range genState.ApprovedCertificatesList {
		k.SetApprovedCertificates(ctx, elem)
	}
	// Set all the proposedCertificate
	for _, elem := range genState.ProposedCertificateList {
		k.SetProposedCertificate(ctx, elem)
	}
	// Set all the childCertificates
	for _, elem := range genState.ChildCertificatesList {
		k.SetChildCertificates(ctx, elem)
	}
	// Set all the proposedCertificateRevocation
	for _, elem := range genState.ProposedCertificateRevocationList {
		k.SetProposedCertificateRevocation(ctx, elem)
	}
	// Set all the revokedCertificates
	for _, elem := range genState.RevokedCertificatesList {
		k.SetRevokedCertificates(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.ApprovedCertificatesList = k.GetAllApprovedCertificates(ctx)
	genesis.ProposedCertificateList = k.GetAllProposedCertificate(ctx)
	genesis.ChildCertificatesList = k.GetAllChildCertificates(ctx)
	genesis.ProposedCertificateRevocationList = k.GetAllProposedCertificateRevocation(ctx)
	genesis.RevokedCertificatesList = k.GetAllRevokedCertificates(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
