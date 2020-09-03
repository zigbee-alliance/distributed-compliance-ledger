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
