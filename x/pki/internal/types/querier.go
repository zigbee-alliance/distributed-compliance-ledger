package types

import (
	"encoding/json"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/pagination"
)

/*
	Request Payload
*/

// Request Payload for PKI queries (pagination and filters).
type PkiQueryParams struct {
	Skip             int
	Take             int
	RootSubject      string
	RootSubjectKeyID string
}

func NewPkiQueryParams(pagination pagination.PaginationParams,
	rootSubject string, rootSubjectKeyID string) PkiQueryParams {
	return PkiQueryParams{
		Skip:             pagination.Skip,
		Take:             pagination.Take,
		RootSubject:      rootSubject,
		RootSubjectKeyID: rootSubjectKeyID,
	}
}

/*
	Result Payload
*/

// Result Payload for QueryAllX509Certs / QueryAllX509RootCerts / QueryAllSubjectX509Certs /
// QueryAllRevokedX509Certs / QueryAllRevokedX509RootCerts queries.
type ListCertificates struct {
	Total int           `json:"total"`
	Items []Certificate `json:"items"`
}

func NewListCertificates() ListCertificates {
	return ListCertificates{
		Total: 0,
		Items: []Certificate{},
	}
}

// Implement fmt.Stringer.
func (n ListCertificates) String() string {
	res, err := json.Marshal(n)
	if err != nil {
		panic(err)
	}

	return string(res)
}

// Result Payload for QueryAllProposedX509RootCerts query.
type ListProposedCertificates struct {
	Total int                   `json:"total"`
	Items []ProposedCertificate `json:"items"`
}

func NewListProposedCertificates() ListProposedCertificates {
	return ListProposedCertificates{
		Total: 0,
		Items: []ProposedCertificate{},
	}
}

// Implement fmt.Stringer.
func (n ListProposedCertificates) String() string {
	res, err := json.Marshal(n)
	if err != nil {
		panic(err)
	}

	return string(res)
}

// Result Payload for QueryAllProposedX509RootCertRevocations query.
type ListProposedCertificateRevocations struct {
	Total int                             `json:"total"`
	Items []ProposedCertificateRevocation `json:"items"`
}

func NewListProposedCertificateRevocations() ListProposedCertificateRevocations {
	return ListProposedCertificateRevocations{
		Total: 0,
		Items: []ProposedCertificateRevocation{},
	}
}

// Implement fmt.Stringer.
func (n ListProposedCertificateRevocations) String() string {
	res, err := json.Marshal(n)
	if err != nil {
		panic(err)
	}

	return string(res)
}
