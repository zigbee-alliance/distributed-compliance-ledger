package pki

import (
	"fmt"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	ApprovedCertificateRecords []types.Certificate         `json:"approved_certificate_records"`
	PendingCertificateRecords  []types.ProposedCertificate `json:"pending_certificate_records"`
	ChildCertificatesRecords   []types.ChildCertificates   `json:"child_certificates_records"`
}

func NewGenesisState() GenesisState {
	return GenesisState{
		ApprovedCertificateRecords: []types.Certificate{},
		PendingCertificateRecords:  []types.ProposedCertificate{},
		ChildCertificatesRecords:   []types.ChildCertificates{},
	}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.ApprovedCertificateRecords {
		if len(record.PemCert) == 0 {
			return fmt.Errorf("invalid ApprovedCertificateRecords: value: %s. Error: Empty X509 Certificate", record.PemCert)
		}
		if record.Type != types.RootCertificate && record.Type != types.IntermediateCertificate {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("invalid ApprovedCertificateRecords: value: %v. Error: Invalid Certificate Type: unknown type; supported types: [%s,%s]", record.Type, types.RootCertificate, types.IntermediateCertificate))
		}
		if len(record.Subject) == 0 {
			return fmt.Errorf("invalid ApprovedCertificateRecords: value: %s. Error: Empty Subject", record.Subject)
		}
		if len(record.SubjectKeyId) == 0 {
			return fmt.Errorf("invalid ApprovedCertificateRecords: value: %s. Error: Empty SubjectKeyId", record.SubjectKeyId)
		}
		if len(record.SerialNumber) == 0 {
			return fmt.Errorf("invalid ApprovedCertificateRecords: value: %s. Error: Empty SerialNumber", record.SerialNumber)
		}
		if len(record.RootSubjectId) == 0 {
			return fmt.Errorf("invalid ApprovedCertificateRecords: value: %s. Error: Empty RootSubjectId", record.RootSubjectId)
		}
	}

	for _, record := range data.PendingCertificateRecords {
		if len(record.PemCert) == 0 {
			return fmt.Errorf("invalid PendingCertificateRecords: value: %s. Error: Empty X509 Certificate", record.PemCert)
		}
		if len(record.Subject) == 0 {
			return fmt.Errorf("invalid PendingCertificateRecords: value: %s. Error: Empty Subject", record.Subject)
		}
		if len(record.SubjectKeyId) == 0 {
			return fmt.Errorf("invalid PendingCertificateRecords: value: %s. Error: Empty SubjectKeyId", record.SubjectKeyId)
		}
	}

	for _, record := range data.ChildCertificatesRecords {
		if len(record.Subject) == 0 {
			return fmt.Errorf("invalid ChildCertificatesRecords: value: %s. Error: Empty Subject", record.Subject)
		}
		if len(record.SubjectKeyId) == 0 {
			return fmt.Errorf("invalid ChildCertificatesRecords: value: %s. Error: Empty SubjectKeyId", record.SubjectKeyId)
		}
	}

	return nil
}

func DefaultGenesisState() GenesisState {
	return NewGenesisState()
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {

	for _, record := range data.PendingCertificateRecords {
		keeper.SetProposedCertificate(ctx, record)
	}

	for _, record := range data.ApprovedCertificateRecords {
		keeper.SetCertificate(ctx, record)
	}

	for _, record := range data.ChildCertificatesRecords {
		keeper.SetChildCertificatesList(ctx, record)
	}

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var approvedCertificates []types.Certificate
	var pendingCertificates []types.ProposedCertificate
	var childCertificatesList []types.ChildCertificates

	k.IterateCertificates(ctx, "", func(certificate types.Certificate) (stop bool) {
		approvedCertificates = append(approvedCertificates, certificate)
		return false
	})

	k.IterateProposedCertificates(ctx, func(certificate types.ProposedCertificate) (stop bool) {
		pendingCertificates = append(pendingCertificates, certificate)
		return false
	})

	k.IterateChildCertificatesRecords(ctx, func(certificatesList types.ChildCertificates) (stop bool) {
		childCertificatesList = append(childCertificatesList, certificatesList)
		return false
	})

	return GenesisState{
		ApprovedCertificateRecords: approvedCertificates,
		PendingCertificateRecords:  pendingCertificates,
		ChildCertificatesRecords:   childCertificatesList,
	}
}
