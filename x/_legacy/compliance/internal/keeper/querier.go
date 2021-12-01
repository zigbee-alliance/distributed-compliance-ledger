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
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/conversions"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/internal/types"
)

const (
	QueryComplianceInfo           = "compliance_info"
	QueryAllComplianceInfoRecords = "all_compliance_info_records"
	QueryCertifiedModel           = "certified_model"
	QueryAllCertifiedModels       = "all_certified_models"
	QueryRevokedModel             = "revoked_model"
	QueryAllRevokedModels         = "all_revoked_models"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryComplianceInfo:
			return queryComplianceInfo(ctx, path[1:], keeper, 0)
		case QueryAllComplianceInfoRecords:
			return queryAllComplianceInfoRecords(ctx, req, keeper)
		case QueryCertifiedModel:
			return queryComplianceInfo(ctx, path[1:], keeper, types.CodeCertified)
		case QueryAllCertifiedModels:
			return queryAllComplianceInfoInStateRecords(ctx, req, keeper, types.CodeCertified)
		case QueryRevokedModel:
			return queryComplianceInfo(ctx, path[1:], keeper, types.CodeRevoked)
		case QueryAllRevokedModels:
			return queryAllComplianceInfoInStateRecords(ctx, req, keeper, types.CodeRevoked)
		default:
			return nil, sdk.ErrUnknownRequest("unknown compliance query endpoint")
		}
	}
}

func queryComplianceInfo(ctx sdk.Context, path []string, keeper Keeper,
	requestedStatus types.SoftwareVersionCertificationStatus) (res []byte, err sdk.Error) {
	vid, err := conversions.ParseVID(path[0])
	if err != nil {
		return nil, err
	}

	pid, err := conversions.ParsePID(path[1])
	if err != nil {
		return nil, err
	}

	softwareVersion, err := conversions.ParseUInt32FromString("SoftwareVersion", path[2])
	if err != nil {
		return nil, err
	}

	certificationType := types.CertificationType(path[3])

	if !keeper.IsComplianceInfoPresent(ctx, certificationType, vid, pid, softwareVersion) {
		return nil, types.ErrComplianceInfoDoesNotExist(vid, pid, softwareVersion, certificationType)
	}

	complianceInfo := keeper.GetComplianceInfo(ctx, certificationType, vid, pid, softwareVersion)

	if requestedStatus != 0 {
		if complianceInfo.SoftwareVersionCertificationStatus != requestedStatus {
			return nil, types.ErrComplianceInfoDoesNotExist(vid, pid, softwareVersion, certificationType)
		}

		res = codec.MustMarshalJSONIndent(keeper.cdc, types.ComplianceInfoInState{Value: true})
	} else {
		res = codec.MustMarshalJSONIndent(keeper.cdc, complianceInfo)
	}

	return res, nil
}

func queryAllComplianceInfoRecords(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	var params types.ListQueryParams
	if err := keeper.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("failed to parse request params: %s", err))
	}

	result := types.ListComplianceInfoItems{
		Total: 0,
		Items: []types.ComplianceInfo{},
	}
	skipped := 0

	keeper.IterateComplianceInfos(ctx, params.CertificationType, func(complianceInfo types.ComplianceInfo) (stop bool) {
		result.Total++

		if skipped < params.Skip {
			skipped++

			return false
		}

		if len(result.Items) < params.Take || params.Take == 0 {
			result.Items = append(result.Items, complianceInfo)

			return false
		}

		return false
	})

	res = codec.MustMarshalJSONIndent(keeper.cdc, result)

	return res, nil
}

func queryAllComplianceInfoInStateRecords(ctx sdk.Context, req abci.RequestQuery, keeper Keeper,
	requestedStatus types.SoftwareVersionCertificationStatus) (res []byte, err sdk.Error) {
	var params types.ListQueryParams
	if err := keeper.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("failed to parse request params: %s", err))
	}

	result := types.ListComplianceInfoKeyItems{
		Total: 0,
		Items: []types.ComplianceInfoKey{},
	}
	skipped := 0

	keeper.IterateComplianceInfos(ctx, params.CertificationType, func(complianceInfo types.ComplianceInfo) (stop bool) {
		if requestedStatus != 0 && complianceInfo.SoftwareVersionCertificationStatus != requestedStatus {
			return false
		}

		result.Total++

		if skipped < params.Skip {
			skipped++

			return false
		}

		if len(result.Items) < params.Take || params.Take == 0 {
			result.Items = append(result.Items, types.ComplianceInfoKey{
				VID:               complianceInfo.VID,
				PID:               complianceInfo.PID,
				SoftwareVersion:   complianceInfo.SoftwareVersion,
				CertificationType: complianceInfo.CertificationType,
			})

			return false
		}

		return false
	})

	res = codec.MustMarshalJSONIndent(keeper.cdc, result)

	return res, nil
}
