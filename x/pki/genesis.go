package pki

import (
	"fmt"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	ApprovedCertificateRecords []types.Certificates       `json:"approved_certificate_records"`
	PendingCertificateRecords  []types.ProposedCertificate `json:"pending_certificate_records"`
	ChildCertificatesRecords   []types.ChildCertificates   `json:"child_certificates_records"`
}

func NewGenesisState() GenesisState {
	return GenesisState{
		ApprovedCertificateRecords: []types.Certificates{},
		PendingCertificateRecords:  []types.ProposedCertificate{},
		ChildCertificatesRecords:   []types.ChildCertificates{},
	}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.ApprovedCertificateRecords {
		for _, certificate := range record.Items {
			if len(certificate.PemCert) == 0 {
				return fmt.Errorf("invalid ApprovedCertificateRecords: value: %s. Error: Empty X509 Certificate", certificate.PemCert)
			}
			if certificate.Type != types.RootCertificate && certificate.Type != types.IntermediateCertificate {
				return sdk.ErrUnknownRequest(
					fmt.Sprintf("invalid ApprovedCertificateRecords: value: %v. Error: Invalid Certificate Type: unknown type; supported types: [%s,%s]", certificate.Type, types.RootCertificate, types.IntermediateCertificate))
			}
			if len(certificate.Subject) == 0 {
				return fmt.Errorf("invalid ApprovedCertificateRecords: value: %s. Error: Empty Subject", certificate.Subject)
			}
			if len(certificate.SubjectKeyId) == 0 {
				return fmt.Errorf("invalid ApprovedCertificateRecords: value: %s. Error: Empty SubjectKeyId", certificate.SubjectKeyId)
			}
			if len(certificate.SerialNumber) == 0 {
				return fmt.Errorf("invalid ApprovedCertificateRecords: value: %s. Error: Empty SerialNumber", certificate.SerialNumber)
			}
			if len(certificate.RootSubjectKeyId) == 0 {
				return fmt.Errorf("invalid ApprovedCertificateRecords: value: %s. Error: Empty RootSubjectId", certificate.RootSubjectKeyId)
			}
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
		if len(record.Items) > 0 {
			keeper.SetCertificates(ctx, record.Items[0].Subject, record.Items[0].SubjectKeyId, record)
		}
	}

	for _, record := range data.ChildCertificatesRecords {
		keeper.SetChildCertificatesList(ctx, record)
	}

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var approvedCertificates []types.Certificates
	var pendingCertificates []types.ProposedCertificate
	var childCertificatesList []types.ChildCertificates

	k.IterateCertificates(ctx, "", func(certificates types.Certificates) (stop bool) {
		approvedCertificates = append(approvedCertificates, certificates)
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
