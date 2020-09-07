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

package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/internal/types"
)

const (
	QueryAllProposedX509RootCerts           = "all_proposed_x509_root_certs"
	QueryProposedX509RootCert               = "proposed_x509_root_cert"
	QueryX509Cert                           = "x509_cert"
	QueryAllX509RootCerts                   = "all_x509_root_certs"
	QueryAllX509Certs                       = "all_x509_certs"
	QueryAllSubjectX509Certs                = "all_subject_x509_certs"
	QueryAllProposedX509RootCertRevocations = "all_proposed_x509_root_cert_revocations"
	QueryProposedX509RootCertRevocation     = "proposed_x509_root_cert_revocation"
	QueryAllRevokedX509Certs                = "all_revoked_x509_certs"
	QueryAllRevokedX509RootCerts            = "all_revoked_x509_root_certs"
	QueryRevokedX509Cert                    = "revoked_x509_cert"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryAllProposedX509RootCerts:
			return queryAllProposedX509RootCerts(ctx, req, keeper)
		case QueryProposedX509RootCert:
			return queryProposedX509RootCert(ctx, path[1:], keeper)
		case QueryX509Cert:
			return queryX509Cert(ctx, path[1:], keeper)
		case QueryAllX509RootCerts:
			return queryAllX509RootCerts(ctx, req, keeper)
		case QueryAllX509Certs:
			return queryAllX509Certs(ctx, req, keeper)
		case QueryAllSubjectX509Certs:
			return queryAllSubjectX509Certs(ctx, path[1:], req, keeper)
		case QueryAllProposedX509RootCertRevocations:
			return queryAllProposedX509RootCertRevocations(ctx, req, keeper)
		case QueryProposedX509RootCertRevocation:
			return queryProposedX509RootCertRevocation(ctx, path[1:], keeper)
		case QueryAllRevokedX509Certs:
			return queryAllRevokedX509Certs(ctx, req, keeper)
		case QueryAllRevokedX509RootCerts:
			return queryAllRevokedX509RootCerts(ctx, req, keeper)
		case QueryRevokedX509Cert:
			return queryRevokedX509Cert(ctx, path[1:], keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown pki query endpoint")
		}
	}
}

// nolint:dupl
func queryAllProposedX509RootCerts(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	var params types.PkiQueryParams
	if err := keeper.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("Failed to parse request params: %s", err))
	}

	result := types.NewListProposedCertificates()

	skipped := 0

	keeper.IterateProposedCertificates(ctx, func(certificate types.ProposedCertificate) (stop bool) {
		result.Total++

		if skipped < params.Skip {
			skipped++

			return false
		}

		if len(result.Items) < params.Take || params.Take == 0 {
			result.Items = append(result.Items, certificate)

			return false
		}

		return false
	})

	res = codec.MustMarshalJSONIndent(keeper.cdc, result)

	return res, nil
}

func queryProposedX509RootCert(ctx sdk.Context, path []string, keeper Keeper) (res []byte, err sdk.Error) {
	subject := path[0]
	subjectKeyID := path[1]

	if !keeper.IsProposedCertificatePresent(ctx, subject, subjectKeyID) {
		return nil, types.ErrProposedCertificateDoesNotExist(subject, subjectKeyID)
	}

	certificate := keeper.GetProposedCertificate(ctx, subject, subjectKeyID)

	res = codec.MustMarshalJSONIndent(keeper.cdc, certificate)

	return res, nil
}

func queryX509Cert(ctx sdk.Context, path []string, keeper Keeper) (res []byte, err sdk.Error) {
	subject := path[0]
	subjectKeyID := path[1]

	if !keeper.IsApprovedCertificatesPresent(ctx, subject, subjectKeyID) {
		return nil, types.ErrCertificateDoesNotExist(subject, subjectKeyID)
	}

	certificates := keeper.GetApprovedCertificates(ctx, subject, subjectKeyID)

	res = codec.MustMarshalJSONIndent(keeper.cdc, certificates)

	return res, nil
}

func queryAllX509RootCerts(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	return queryX509Certs(ctx, req, keeper, true, false, "")
}

func queryAllX509Certs(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	return queryX509Certs(ctx, req, keeper, false, false, "")
}

func queryAllSubjectX509Certs(ctx sdk.Context, path []string,
	req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	subject := path[0]

	return queryX509Certs(ctx, req, keeper, false, false, subject)
}

// nolint:gocognit
func queryX509Certs(ctx sdk.Context, req abci.RequestQuery, keeper Keeper,
	onlyRoot bool, revoked bool, iteratorPrefix string) (res []byte, err sdk.Error) {
	var params types.PkiQueryParams
	if err := keeper.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("Failed to parse request params: %s", err))
	}

	result := types.NewListCertificates()

	skipped := 0

	process := func(certificates types.Certificates) (stop bool) {
		for _, certificate := range certificates.Items {
			// filter by certificate type (Root/Any)
			if onlyRoot && !certificate.IsRoot {
				return false
			}

			// filter by root subject
			if len(params.RootSubject) > 0 {
				if !certificate.IsRoot && certificate.RootSubject != params.RootSubject ||
					certificate.IsRoot && certificate.Subject != params.RootSubject {
					return false
				}
			}

			// filter by root subject key id
			if len(params.RootSubjectKeyID) > 0 {
				if !certificate.IsRoot && certificate.RootSubjectKeyID != params.RootSubjectKeyID ||
					certificate.IsRoot && certificate.SubjectKeyID != params.RootSubjectKeyID {
					return false
				}
			}

			result.Total++

			if skipped < params.Skip {
				skipped++

				return false
			}

			if len(result.Items) < params.Take || params.Take == 0 {
				result.Items = append(result.Items, certificate)

				return false
			}
		}

		return false
	}

	if revoked {
		keeper.IterateRevokedCertificatesRecords(ctx, iteratorPrefix, process)
	} else {
		keeper.IterateApprovedCertificatesRecords(ctx, iteratorPrefix, process)
	}

	res = codec.MustMarshalJSONIndent(keeper.cdc, result)

	return res, nil
}

// nolint:dupl
func queryAllProposedX509RootCertRevocations(ctx sdk.Context, req abci.RequestQuery,
	keeper Keeper) (res []byte, err sdk.Error) {
	var params types.PkiQueryParams
	if err := keeper.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("Failed to parse request params: %s", err))
	}

	result := types.NewListProposedCertificateRevocations()

	skipped := 0

	keeper.IterateProposedCertificateRevocations(ctx, func(revocation types.ProposedCertificateRevocation) (stop bool) {
		result.Total++

		if skipped < params.Skip {
			skipped++

			return false
		}

		if len(result.Items) < params.Take || params.Take == 0 {
			result.Items = append(result.Items, revocation)

			return false
		}

		return false
	})

	res = codec.MustMarshalJSONIndent(keeper.cdc, result)

	return res, nil
}

func queryProposedX509RootCertRevocation(ctx sdk.Context, path []string, keeper Keeper) (res []byte, err sdk.Error) {
	subject := path[0]
	subjectKeyID := path[1]

	if !keeper.IsProposedCertificateRevocationPresent(ctx, subject, subjectKeyID) {
		return nil, types.ErrProposedCertificateRevocationDoesNotExist(subject, subjectKeyID)
	}

	revocation := keeper.GetProposedCertificateRevocation(ctx, subject, subjectKeyID)

	res = codec.MustMarshalJSONIndent(keeper.cdc, revocation)

	return res, nil
}

func queryAllRevokedX509Certs(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	return queryX509Certs(ctx, req, keeper, false, true, "")
}

func queryAllRevokedX509RootCerts(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	return queryX509Certs(ctx, req, keeper, true, true, "")
}

func queryRevokedX509Cert(ctx sdk.Context, path []string, keeper Keeper) (res []byte, err sdk.Error) {
	subject := path[0]
	subjectKeyID := path[1]

	if !keeper.IsRevokedCertificatesPresent(ctx, subject, subjectKeyID) {
		return nil, types.ErrRevokedCertificateDoesNotExist(subject, subjectKeyID)
	}

	certificate := keeper.GetRevokedCertificates(ctx, subject, subjectKeyID)

	res = codec.MustMarshalJSONIndent(keeper.cdc, certificate)

	return res, nil
}
