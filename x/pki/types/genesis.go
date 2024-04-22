package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index.
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state.
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ApprovedCertificatesList:                                []ApprovedCertificates{},
		ProposedCertificateList:                                 []ProposedCertificate{},
		ChildCertificatesList:                                   []ChildCertificates{},
		ProposedCertificateRevocationList:                       []ProposedCertificateRevocation{},
		RevokedCertificatesList:                                 []RevokedCertificates{},
		UniqueCertificateList:                                   []UniqueCertificate{},
		ApprovedRootCertificates:                                nil,
		RevokedRootCertificates:                                 nil,
		ApprovedCertificatesBySubjectList:                       []ApprovedCertificatesBySubject{},
		RejectedCertificateList:                                 []RejectedCertificate{},
		PkiRevocationDistributionPointList:                      []PkiRevocationDistributionPoint{},
		PkiRevocationDistributionPointsByIssuerSubjectKeyIDList: []PkiRevocationDistributionPointsByIssuerSubjectKeyID{},
		ApprovedCertificatesBySubjectKeyIdList:                  []ApprovedCertificatesBySubjectKeyId{},
		NocRootCertificatesList:                                 []NocRootCertificates{},
		NocIcaCertificatesList:                                  []NocIcaCertificates{},
		RevokedNocRootCertificatesList:                          []RevokedNocRootCertificates{},
		NocRootCertificatesByVidAndSkidList:                     []NocRootCertificatesByVidAndSkid{},
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error { //nolint:gocyclo,vet
	// Check for duplicated index in approvedCertificates
	approvedCertificatesIndexMap := make(map[string]struct{})

	for _, elem := range gs.ApprovedCertificatesList {
		index := string(ApprovedCertificatesKey(elem.Subject, elem.SubjectKeyId))
		if _, ok := approvedCertificatesIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for approvedCertificates")
		}

		approvedCertificatesIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in proposedCertificate
	proposedCertificateIndexMap := make(map[string]struct{})

	for _, elem := range gs.ProposedCertificateList {
		index := string(ProposedCertificateKey(elem.Subject, elem.SubjectKeyId))
		if _, ok := proposedCertificateIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for proposedCertificate")
		}

		proposedCertificateIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in childCertificates
	childCertificatesIndexMap := make(map[string]struct{})

	for _, elem := range gs.ChildCertificatesList {
		index := string(ChildCertificatesKey(elem.Issuer, elem.AuthorityKeyId))
		if _, ok := childCertificatesIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for childCertificates")
		}
		childCertificatesIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in proposedCertificateRevocation
	proposedCertificateRevocationIndexMap := make(map[string]struct{})

	for _, elem := range gs.ProposedCertificateRevocationList {
		index := string(ProposedCertificateRevocationKey(elem.Subject, elem.SubjectKeyId, elem.SerialNumber))
		if _, ok := proposedCertificateRevocationIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for proposedCertificateRevocation")
		}
		proposedCertificateRevocationIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in revokedCertificates
	revokedCertificatesIndexMap := make(map[string]struct{})

	for _, elem := range gs.RevokedCertificatesList {
		index := string(RevokedCertificatesKey(elem.Subject, elem.SubjectKeyId))
		if _, ok := revokedCertificatesIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for revokedCertificates")
		}
		revokedCertificatesIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in uniqueCertificate
	uniqueCertificateIndexMap := make(map[string]struct{})

	for _, elem := range gs.UniqueCertificateList {
		index := string(UniqueCertificateKey(elem.Issuer, elem.SerialNumber))
		if _, ok := uniqueCertificateIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for uniqueCertificate")
		}
		uniqueCertificateIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in approvedCertificatesBySubject
	approvedCertificatesBySubjectIndexMap := make(map[string]struct{})

	for _, elem := range gs.ApprovedCertificatesBySubjectList {
		index := string(ApprovedCertificatesBySubjectKey(elem.Subject))
		if _, ok := approvedCertificatesBySubjectIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for approvedCertificatesBySubject")
		}
		approvedCertificatesBySubjectIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in rejectedCertificate
	rejectedCertificateIndexMap := make(map[string]struct{})

	for _, elem := range gs.RejectedCertificateList {
		index := string(RejectedCertificateKey(elem.Subject, elem.SubjectKeyId))
		if _, ok := rejectedCertificateIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for rejectedCertificate")
		}
		rejectedCertificateIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in pKIRevocationDistributionPoint
	pKIRevocationDistributionPointIndexMap := make(map[string]struct{})

	for _, elem := range gs.PkiRevocationDistributionPointList {
		index := string(PkiRevocationDistributionPointKey(elem.Vid, elem.Label, elem.IssuerSubjectKeyID))
		if _, ok := pKIRevocationDistributionPointIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for pKIRevocationDistributionPoint")
		}
		pKIRevocationDistributionPointIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in pkiRevocationDistributionPointsByIssuerSubjectKeyID
	pkiRevocationDistributionPointsByIssuerSubjectKeyIDIndexMap := make(map[string]struct{})

	for _, elem := range gs.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList {
		index := string(PkiRevocationDistributionPointsByIssuerSubjectKeyIDKey(elem.IssuerSubjectKeyID))
		if _, ok := pkiRevocationDistributionPointsByIssuerSubjectKeyIDIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for pkiRevocationDistributionPointsByIssuerSubjectKeyID")
		}
		pkiRevocationDistributionPointsByIssuerSubjectKeyIDIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in approvedCertificatesBySubjectKeyId
	approvedCertificatesBySubjectKeyIDIndexMap := make(map[string]struct{})

	for _, elem := range gs.ApprovedCertificatesBySubjectKeyIdList {
		index := string(ApprovedCertificatesBySubjectKeyIDKey(elem.SubjectKeyId))
		if _, ok := approvedCertificatesBySubjectKeyIDIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for approvedCertificatesBySubjectKeyId")
		}
		approvedCertificatesBySubjectKeyIDIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in nocRootCertificates
	nocRootCertificatesIndexMap := make(map[string]struct{})

	for _, elem := range gs.NocRootCertificatesList {
		index := string(NocRootCertificatesKey(elem.Vid))
		if _, ok := nocRootCertificatesIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for nocRootCertificates")
		}
		nocRootCertificatesIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in nocCertificates
	nocCertificatesIndexMap := make(map[string]struct{})

	for _, elem := range gs.NocIcaCertificatesList {
		index := string(NocIcaCertificatesKey(elem.Vid))
		if _, ok := nocCertificatesIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for nocCertificates")
		}
		nocCertificatesIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in revokedNocRootCertificates
	revokedNocRootCertificatesIndexMap := make(map[string]struct{})

	for _, elem := range gs.RevokedNocRootCertificatesList {
		index := string(RevokedNocRootCertificatesKey(elem.Subject, elem.SubjectKeyId))
		if _, ok := revokedNocRootCertificatesIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for revokedNocRootCertificates")
		}
		revokedNocRootCertificatesIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in nocRootCertificatesByVidAndSkid
	nocRootCertificatesByVidAndSkidIndexMap := make(map[string]struct{})

	for _, elem := range gs.NocRootCertificatesByVidAndSkidList {
		index := string(NocRootCertificatesByVidAndSkidKey(elem.Vid, elem.SubjectKeyId))
		if _, ok := nocRootCertificatesByVidAndSkidIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for nocRootCertificatesByVidAndSkid")
		}
		nocRootCertificatesByVidAndSkidIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
