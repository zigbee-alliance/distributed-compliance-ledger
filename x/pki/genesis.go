// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pki

import (
	"fmt"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type GenesisState struct {
	ProposedCertificates           []types.ProposedCertificate           `json:"proposed_certificates"`
	ApprovedCertificatesRecords    []types.Certificates                  `json:"approved_certificates_records"`
	ProposedCertificateRevocations []types.ProposedCertificateRevocation `json:"proposed_certificate_revocations"`
	RevokedCertificatesRecords     []types.Certificates                  `json:"revoked_certificates_records"`
	ChildCertificatesRecords       []types.ChildCertificates             `json:"child_certificates_records"`
}

func NewGenesisState() GenesisState {
	return GenesisState{
		ProposedCertificates:           []types.ProposedCertificate{},
		ApprovedCertificatesRecords:    []types.Certificates{},
		ProposedCertificateRevocations: []types.ProposedCertificateRevocation{},
		RevokedCertificatesRecords:     []types.Certificates{},
		ChildCertificatesRecords:       []types.ChildCertificates{},
	}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.ProposedCertificates {
		if err := validateProposedCertificate(record); err != nil {
			return err
		}
	}

	for _, record := range data.ApprovedCertificatesRecords {
		if err := validateCertificates(record); err != nil {
			return err
		}
	}

	for _, record := range data.ProposedCertificateRevocations {
		if err := validateProposedCertificateRevocation(record); err != nil {
			return err
		}
	}

	for _, record := range data.RevokedCertificatesRecords {
		if err := validateCertificates(record); err != nil {
			return err
		}
	}

	for _, record := range data.ChildCertificatesRecords {
		if err := validateChildCertificates(record); err != nil {
			return err
		}
	}

	return nil
}

func validateProposedCertificate(record types.ProposedCertificate) error {
	if len(record.PemCert) == 0 {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf("Invalid ProposedCertificate: Empty X509 Certificate. Value: %v", record))
	}

	if len(record.Subject) == 0 {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf("Invalid ProposedCertificate: Empty Subject. Value: %v", record))
	}

	if len(record.SubjectKeyID) == 0 {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf("Invalid ProposedCertificate: Empty SubjectKeyID. Value: %v", record))
	}

	if len(record.SerialNumber) == 0 {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf("Invalid ProposedCertificate: Empty SerialNumber. Value: %v", record))
	}

	if record.Owner.Empty() {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf("Invalid ProposedCertificate: Empty Owner. Value: %v", record))
	}

	if record.Approvals == nil {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf("Invalid ProposedCertificate: Approvals is nil. Value: %v", record))
	}

	for _, approval := range record.Approvals {
		if approval.Empty() {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("Invalid ProposedCertificate: Empty Approval. Value: %v", record))
		}
	}

	return nil
}

func validateCertificates(record types.Certificates) error {
	if len(record.Items) == 0 {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf("Invalid Certificates: Empty Items. Value: %v", record))
	}

	for _, certificate := range record.Items {
		if len(certificate.PemCert) == 0 {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("Invalid Certificate: Empty PemCert. Value: %v", certificate))
		}

		if len(certificate.Subject) == 0 {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("Invalid Certificate: Empty Subject. Value: %v", certificate))
		}

		if len(certificate.SubjectKeyID) == 0 {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("Invalid Certificate: Empty SubjectKeyID. Value: %v", certificate))
		}

		if len(certificate.SerialNumber) == 0 {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("Invalid Certificate: Empty SerialNumber. Value: %v", certificate))
		}

		if !certificate.IsRoot {
			if len(certificate.Issuer) == 0 {
				return sdk.ErrUnknownRequest(
					fmt.Sprintf("Invalid Certificate: Empty Issuer. Value: %v", certificate))
			}

			if len(certificate.AuthorityKeyID) == 0 {
				return sdk.ErrUnknownRequest(
					fmt.Sprintf("Invalid Certificate: Empty AuthorityKeyID. Value: %v", certificate))
			}

			if len(certificate.RootSubject) == 0 {
				return sdk.ErrUnknownRequest(
					fmt.Sprintf("Invalid Certificate: Empty RootSubject. Value: %v", certificate))
			}

			if len(certificate.RootSubjectKeyID) == 0 {
				return sdk.ErrUnknownRequest(
					fmt.Sprintf("Invalid Certificate: Empty RootSubjectKeyID. Value: %v", certificate))
			}
		}

		if certificate.Owner.Empty() {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("Invalid Certificate: Empty Owner. Value: %v", certificate))
		}
	}

	return nil
}

func validateProposedCertificateRevocation(record types.ProposedCertificateRevocation) error {
	if len(record.Subject) == 0 {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf("Invalid ProposedCertificateRevocation: Empty Subject. Value: %v", record))
	}

	if len(record.SubjectKeyID) == 0 {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf("Invalid ProposedCertificateRevocation: Empty SubjectKeyID. Value: %v", record))
	}

	if record.Approvals == nil {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf("Invalid ProposedCertificateRevocation: Approvals is nil. Value: %v", record))
	}

	for _, approval := range record.Approvals {
		if approval.Empty() {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("Invalid ProposedCertificateRevocation: Empty Approval. Value: %v", record))
		}
	}

	return nil
}

func validateChildCertificates(record types.ChildCertificates) error {
	if len(record.Issuer) == 0 {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf("Invalid ChildCertificates: Empty Issuer. Value: %v", record))
	}

	if len(record.AuthorityKeyID) == 0 {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf("Invalid ChildCertificates: Empty AuthorityKeyID. Value: %v", record))
	}

	if len(record.CertIdentifiers) == 0 {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf("Invalid ChildCertificates: Empty CertIdentifiers. Value: %v", record))
	}

	for _, certIdentifier := range record.CertIdentifiers {
		if len(certIdentifier.Subject) == 0 {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("Invalid ChildCertificates: Empty CertIdentifier.Subject. Value: %v", record))
		}

		if len(certIdentifier.SubjectKeyID) == 0 {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("Invalid ChildCertificates: Empty CertIdentifier.SubjectKeyID. Value: %v", record))
		}
	}

	return nil
}

func DefaultGenesisState() GenesisState {
	return NewGenesisState()
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	for _, record := range data.ProposedCertificates {
		keeper.SetProposedCertificate(ctx, record)
	}

	for _, record := range data.ApprovedCertificatesRecords {
		if len(record.Items) > 0 {
			keeper.SetApprovedCertificates(ctx, record.Items[0].Subject, record.Items[0].SubjectKeyID, record)
		}
	}

	for _, record := range data.ProposedCertificateRevocations {
		keeper.SetProposedCertificateRevocation(ctx, record)
	}

	for _, record := range data.RevokedCertificatesRecords {
		if len(record.Items) > 0 {
			keeper.SetRevokedCertificates(ctx, record.Items[0].Subject, record.Items[0].SubjectKeyID, record)
		}
	}

	for _, record := range data.ChildCertificatesRecords {
		keeper.SetChildCertificates(ctx, record)
	}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var (
		proposedCertificates           []types.ProposedCertificate
		approvedCertificatesRecords    []types.Certificates
		proposedCertificateRevocations []types.ProposedCertificateRevocation
		revokedCertificatesRecords     []types.Certificates
		childCertificatesRecords       []types.ChildCertificates
	)

	k.IterateProposedCertificates(ctx, func(value types.ProposedCertificate) (stop bool) {
		proposedCertificates = append(proposedCertificates, value)

		return false
	})

	k.IterateApprovedCertificatesRecords(ctx, "", func(value types.Certificates) (stop bool) {
		approvedCertificatesRecords = append(approvedCertificatesRecords, value)

		return false
	})

	k.IterateProposedCertificateRevocations(ctx, func(value types.ProposedCertificateRevocation) (stop bool) {
		proposedCertificateRevocations = append(proposedCertificateRevocations, value)

		return false
	})

	k.IterateRevokedCertificatesRecords(ctx, "", func(value types.Certificates) (stop bool) {
		revokedCertificatesRecords = append(revokedCertificatesRecords, value)

		return false
	})

	k.IterateChildCertificatesRecords(ctx, func(value types.ChildCertificates) (stop bool) {
		childCertificatesRecords = append(childCertificatesRecords, value)

		return false
	})

	return GenesisState{
		ProposedCertificates:           proposedCertificates,
		ApprovedCertificatesRecords:    approvedCertificatesRecords,
		ProposedCertificateRevocations: proposedCertificateRevocations,
		RevokedCertificatesRecords:     revokedCertificatesRecords,
		ChildCertificatesRecords:       childCertificatesRecords,
	}
}
