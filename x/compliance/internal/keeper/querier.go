package keeper

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/conversions"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
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
			return queryComplianceInfo(ctx, path[1:], req, keeper, "")
		case QueryAllComplianceInfoRecords:
			return queryAllComplianceInfoRecords(ctx, req, keeper, "")
		case QueryCertifiedModel:
			return queryComplianceInfo(ctx, path[1:], req, keeper, types.Certified)
		case QueryAllCertifiedModels:
			return queryAllComplianceInfoRecords(ctx, req, keeper, types.Certified)
		case QueryRevokedModel:
			return queryComplianceInfo(ctx, path[1:], req, keeper, types.Revoked)
		case QueryAllRevokedModels:
			return queryAllComplianceInfoRecords(ctx, req, keeper, types.Revoked)
		default:
			return nil, sdk.ErrUnknownRequest("unknown compliance query endpoint")
		}
	}
}

func queryComplianceInfo(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper, requestedState types.ComplianceState) ([]byte, sdk.Error) {
	var params types.SingleQueryParams

	if err := keeper.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to parse params: %s", err))
	}

	vid, err := conversions.ParseVID(path[0])
	if err != nil {
		return nil, err
	}

	pid, err := conversions.ParsePID(path[1])
	if err != nil {
		return nil, err
	}

	if !keeper.IsComplianceInfoPresent(ctx, params.CertificationType, vid, pid) {
		return nil, types.ErrComplianceInfoDoesNotExist(vid, pid)
	}

	complianceInfo := keeper.GetComplianceInfo(ctx, params.CertificationType, vid, pid)

	if len(requestedState) != 0 && complianceInfo.State != requestedState {
		return nil, types.ErrComplianceInfoDoesNotExist(vid, pid)
	}

	res, err_ := codec.MarshalJSONIndent(keeper.cdc, complianceInfo)
	if err_ != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

func queryAllComplianceInfoRecords(ctx sdk.Context, req abci.RequestQuery, keeper Keeper, requestedState types.ComplianceState) ([]byte, sdk.Error) {
	var params types.ListQueryParams
	if err := keeper.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to parse params: %s", err))
	}

	result := types.ListComplianceInfoItems{
		Total: 0,
		Items: []types.ComplianceInfo{},
	}
	skipped := 0

	keeper.IterateComplianceInfos(ctx, params.CertificationType, func(complianceInfo types.ComplianceInfo) (stop bool) {
		if len(requestedState) != 0 && complianceInfo.State != requestedState {
			return false
		}

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

	res, err := codec.MarshalJSONIndent(keeper.cdc, result)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
