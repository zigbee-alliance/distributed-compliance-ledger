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
	// Set all the uniqueCertificate
	for _, elem := range genState.UniqueCertificateList {
		k.SetUniqueCertificate(ctx, elem)
	}
	// Set if defined
	if genState.ApprovedRootCertificates != nil {
		k.SetApprovedRootCertificates(ctx, *genState.ApprovedRootCertificates)
	}
	// Set if defined
	if genState.RevokedRootCertificates != nil {
		k.SetRevokedRootCertificates(ctx, *genState.RevokedRootCertificates)
	}
	// Set all the approvedCertificatesBySubject
	for _, elem := range genState.ApprovedCertificatesBySubjectList {
		k.SetApprovedCertificatesBySubject(ctx, elem)
	}
	// Set all the rejectedCertificate
	for _, elem := range genState.RejectedCertificateList {
		k.SetRejectedCertificate(ctx, elem)
	}
	// Set all the pKIRevocationDistributionPoint
	for _, elem := range genState.PkiRevocationDistributionPointList {
		k.SetPkiRevocationDistributionPoint(ctx, elem)
	}
	// Set all the pkiRevocationDistributionPointsByIssuerSubjectKeyID
	for _, elem := range genState.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList {
		k.SetPkiRevocationDistributionPointsByIssuerSubjectKeyID(ctx, elem)
	}
	// Set all the approvedCertificatesBySubjectKeyId
	for _, elem := range genState.ApprovedCertificatesBySubjectKeyIdList {
		k.SetApprovedCertificatesBySubjectKeyID(ctx, elem)
	}
	// Set all the nocRootCertificates
	for _, elem := range genState.NocRootCertificatesList {
		k.SetNocRootCertificates(ctx, elem)
	}
	// Set all the nocIcaCertificates
	for _, elem := range genState.NocIcaCertificatesList {
		k.SetNocIcaCertificates(ctx, elem)
	}
	// Set all the revokedNocRootCertificates
	for _, elem := range genState.RevokedNocRootCertificatesList {
		k.SetRevokedNocRootCertificates(ctx, elem)
	}
	// Set all the nocCertificatesByVidAndSkid
	for _, elem := range genState.NocCertificatesByVidAndSkidList {
		k.SetNocCertificatesByVidAndSkid(ctx, elem)
	}
	// Set all the nocCertificatesBySubjectKeyId
	for _, elem := range genState.NocCertificatesBySubjectKeyIDList {
		k.SetNocCertificatesBySubjectKeyID(ctx, elem)
	}
	// Set all the nocCertificates
	for _, elem := range genState.NocCertificatesList {
		k.SetNocCertificates(ctx, elem)
	}
	// Set all the nocCertificatesBySubject
	for _, elem := range genState.NocCertificatesBySubjectList {
		k.SetNocCertificatesBySubject(ctx, elem)
	}
	// Set all the certificates
	for _, elem := range genState.CertificatesList {
		k.SetAllCertificates(ctx, elem)
	}
	// Set all the revokedNocIcaCertificates
	for _, elem := range genState.RevokedNocIcaCertificatesList {
		k.SetRevokedNocIcaCertificates(ctx, elem)
	}
	// Set all the allCertificatesBySubject
	for _, elem := range genState.AllCertificatesBySubjectList {
		k.SetAllCertificatesBySubject(ctx, elem)
	}
	for _, elem := range genState.AllCertificatesBySubjectKeyIdList {
		k.SetAllCertificatesBySubjectKeyID(ctx, elem)
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
	genesis.UniqueCertificateList = k.GetAllUniqueCertificate(ctx)
	// Get all approvedRootCertificates
	approvedRootCertificates, found := k.GetApprovedRootCertificates(ctx)
	if found {
		genesis.ApprovedRootCertificates = &approvedRootCertificates
	}
	// Get all revokedRootCertificates
	revokedRootCertificates, found := k.GetRevokedRootCertificates(ctx)
	if found {
		genesis.RevokedRootCertificates = &revokedRootCertificates
	}
	genesis.ApprovedCertificatesBySubjectList = k.GetAllApprovedCertificatesBySubject(ctx)
	genesis.RejectedCertificateList = k.GetAllRejectedCertificate(ctx)
	genesis.PkiRevocationDistributionPointList = k.GetAllPkiRevocationDistributionPoint(ctx)
	genesis.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList = k.GetAllPkiRevocationDistributionPointsByIssuerSubjectKeyID(ctx)
	genesis.ApprovedCertificatesBySubjectKeyIdList = k.GetAllApprovedCertificatesBySubjectKeyID(ctx)
	genesis.NocRootCertificatesList = k.GetAllNocRootCertificates(ctx)
	genesis.NocIcaCertificatesList = k.GetAllNocIcaCertificates(ctx)
	genesis.RevokedNocRootCertificatesList = k.GetAllRevokedNocRootCertificates(ctx)
	genesis.NocCertificatesByVidAndSkidList = k.GetAllNocCertificatesByVidAndSkid(ctx)
	genesis.NocCertificatesList = k.GetAllNocCertificates(ctx)
	genesis.NocCertificatesBySubjectList = k.GetAllNocCertificatesBySubject(ctx)
	genesis.NocCertificatesBySubjectKeyIDList = k.GetAllNocCertificatesBySubjectKeyID(ctx)
	genesis.CertificatesList = k.GetAllAllCertificates(ctx)
	genesis.RevokedNocIcaCertificatesList = k.GetAllRevokedNocIcaCertificates(ctx)
	genesis.AllCertificatesBySubjectList = k.GetAllAllCertificatesBySubject(ctx)
	genesis.AllCertificatesBySubjectKeyIdList = k.GetAllAllCertificatesBySubjectKeyID(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
