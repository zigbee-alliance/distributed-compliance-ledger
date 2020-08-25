package pki

import (
	"fmt"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	ApprovedCertificateRecords []types.Certificates        `json:"approved_certificate_records"`
	ProposedCertificateRecords []types.ProposedCertificate `json:"proposed_certificate_records"`
	ChildCertificatesRecords   []types.ChildCertificates   `json:"child_certificates_records"`
}

func NewGenesisState() GenesisState {
	return GenesisState{
		ApprovedCertificateRecords: []types.Certificates{},
		ProposedCertificateRecords: []types.ProposedCertificate{},
		ChildCertificatesRecords:   []types.ChildCertificates{},
	}
}

func ValidateGenesis(data GenesisState) error {
	if err := validateApprovedCertificates(data.ApprovedCertificateRecords); err != nil {
		return err
	}

	if err := validateProposedCertificates(data.ProposedCertificateRecords); err != nil {
		return err
	}

	if err := validateChildCertificatesRecords(data.ChildCertificatesRecords); err != nil {
		return err
	}

	return nil
}

func validateApprovedCertificates(approvedCertificateRecords []types.Certificates) error {
	for _, record := range approvedCertificateRecords {
		for _, certificate := range record.Items {
			if len(certificate.PemCert) == 0 {
				return sdk.ErrUnknownRequest(
					fmt.Sprintf("Invalid ApprovedCertificateRecords: value: %s."+
						" Error: Empty X509 Certificate", certificate.PemCert))
			}

			if len(certificate.Subject) == 0 {
				return sdk.ErrUnknownRequest(
					fmt.Sprintf("Invalid ApprovedCertificateRecords: value: %s. "+
						"Error: Empty Subject", certificate.Subject))
			}

			if len(certificate.SubjectKeyID) == 0 {
				return sdk.ErrUnknownRequest(
					fmt.Sprintf("Invalid ApprovedCertificateRecords: value: %s. Error: "+
						"Empty SubjectKeyID", certificate.SubjectKeyID))
			}

			if len(certificate.SerialNumber) == 0 {
				return sdk.ErrUnknownRequest(
					fmt.Sprintf("Invalid ApprovedCertificateRecords: value: %s. "+
						"Error: Empty SerialNumber", certificate.SerialNumber))
			}

			if len(certificate.RootSubjectKeyID) == 0 {
				return sdk.ErrUnknownRequest(
					fmt.Sprintf("Invalid ApprovedCertificateRecords: value: %s. "+
						"Error: Empty RootSubjectId", certificate.RootSubjectKeyID))
			}
		}
	}

	return nil
}

func validateProposedCertificates(proposedCertificateRecords []types.ProposedCertificate) error {
	for _, record := range proposedCertificateRecords {
		if len(record.PemCert) == 0 {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("Invalid ProposedCertificateRecords: value: %s. Error: Empty X509 Certificate", record.PemCert))
		}

		if len(record.Subject) == 0 {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("Invalid ProposedCertificateRecords: value: %s. Error: Empty Subject", record.Subject))
		}

		if len(record.SubjectKeyID) == 0 {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("Invalid ProposedCertificateRecords: value: %s. Error: Empty SubjectKeyID", record.SubjectKeyID))
		}
	}

	return nil
}

func validateChildCertificatesRecords(childCertificatesRecords []types.ChildCertificates) error {
	for _, record := range childCertificatesRecords {
		if len(record.Issuer) == 0 {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("Invalid ChildCertificatesRecords: value: %s. Error: Empty Issuer", record.Issuer))
		}

		if len(record.AuthorityKeyID) == 0 {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("Invalid ChildCertificatesRecords: value: %s. Error: Empty AuthorityKeyId", record.AuthorityKeyID))
		}
	}

	return nil
}

func DefaultGenesisState() GenesisState {
	return NewGenesisState()
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.ProposedCertificateRecords {
		keeper.SetProposedCertificate(ctx, record)
	}

	for _, record := range data.ApprovedCertificateRecords {
		if len(record.Items) > 0 {
			keeper.SetApprovedCertificates(ctx, record.Items[0].Subject, record.Items[0].SubjectKeyID, record)
		}
	}

	for _, record := range data.ChildCertificatesRecords {
		keeper.SetChildCertificates(ctx, record)
	}

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var approvedCertificates []types.Certificates

	var proposedCertificates []types.ProposedCertificate

	var childCertificatesList []types.ChildCertificates

	k.IterateApprovedCertificatesRecords(ctx, "", func(certificates types.Certificates) (stop bool) {
		approvedCertificates = append(approvedCertificates, certificates)
		return false
	})

	k.IterateProposedCertificates(ctx, func(certificate types.ProposedCertificate) (stop bool) {
		proposedCertificates = append(proposedCertificates, certificate)
		return false
	})

	k.IterateChildCertificatesRecords(ctx, func(certificatesList types.ChildCertificates) (stop bool) {
		childCertificatesList = append(childCertificatesList, certificatesList)
		return false
	})

	return GenesisState{
		ApprovedCertificateRecords: approvedCertificates,
		ProposedCertificateRecords: proposedCertificates,
		ChildCertificatesRecords:   childCertificatesList,
	}
}
