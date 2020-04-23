package types

import (
	"encoding/json"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/pagination"
)

/*
	Request Payload
*/

// Request Payload for QueryAllSubjectX509Certs(pagination and filtering) query
type ListCertificatesQueryParams struct {
	Skip             int
	Take             int
	RootSubject      string
	RootSubjectKeyId string
}

func NewListCertificatesQueryParams(pagination pagination.PaginationParams, rootSubject string, rootSubjectKeyId string) ListCertificatesQueryParams {
	return ListCertificatesQueryParams{
		Skip:             pagination.Skip,
		Take:             pagination.Take,
		RootSubject:      rootSubject,
		RootSubjectKeyId: rootSubjectKeyId,
	}
}

/*
	Result Payload
*/

// Result Payload for QueryAllX509RootCerts / QueryAllX509Certs / QueryAllSubjectX509Certs queries
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

// Implement fmt.Stringer
func (n ListCertificates) String() string {
	res, err := json.Marshal(n)

	if err != nil {
		panic(err)
	}

	return string(res)
}

// Result Payload for QueryAllProposedX509RootCerts query
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

// Implement fmt.Stringer
func (n ListProposedCertificates) String() string {
	res, err := json.Marshal(n)

	if err != nil {
		panic(err)
	}

	return string(res)
}
