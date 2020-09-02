package types

import (
	"encoding/json"
)

/*
	Request Payload
*/

// Request Payload for QueryAllComplianceInfoRecords/QueryAllCertifiedModels/QueryAllRevokedModels
//(pagination and filtering) query.
type ListQueryParams struct {
	CertificationType CertificationType
	Skip              int
	Take              int
}

func NewListQueryParams(certificationType CertificationType, skip int, take int) ListQueryParams {
	return ListQueryParams{
		CertificationType: certificationType,
		Skip:              skip,
		Take:              take,
	}
}

/*
	Response Payload
*/

// Response Payload for QueryAllComplianceInfoRecords query.
type ListComplianceInfoItems struct {
	Total int              `json:"total"`
	Items []ComplianceInfo `json:"items"`
}

// Implement fmt.Stringer.
func (n ListComplianceInfoItems) String() string {
	res, err := json.Marshal(n)
	if err != nil {
		panic(err)
	}

	return string(res)
}

// Response Payload for QueryAllCertifiedModels/QueryAllRevokedModels queries.
type ListComplianceInfoKeyItems struct {
	Total int                 `json:"total"`
	Items []ComplianceInfoKey `json:"items"`
}

// Implement fmt.Stringer.
func (n ListComplianceInfoKeyItems) String() string {
	res, err := json.Marshal(n)
	if err != nil {
		panic(err)
	}

	return string(res)
}

// Response Payload for QueryCertifiedModel/QueryRevokedModel queries.
type ComplianceInfoInState struct {
	Value bool `json:"value"`
}

func (n ComplianceInfoInState) String() string {
	res, err := json.Marshal(n)
	if err != nil {
		panic(err)
	}

	return string(res)
}
