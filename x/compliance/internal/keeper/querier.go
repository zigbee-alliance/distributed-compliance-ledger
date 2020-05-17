package keeper

//nolint:goimports
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
			return queryComplianceInfo(ctx, path[1:], keeper, "")
		case QueryAllComplianceInfoRecords:
			return queryAllComplianceInfoRecords(ctx, req, keeper)
		case QueryCertifiedModel:
			return queryComplianceInfo(ctx, path[1:], keeper, types.Certified)
		case QueryAllCertifiedModels:
			return queryAllComplianceInfoInStateRecords(ctx, req, keeper, types.Certified)
		case QueryRevokedModel:
			return queryComplianceInfo(ctx, path[1:], keeper, types.Revoked)
		case QueryAllRevokedModels:
			return queryAllComplianceInfoInStateRecords(ctx, req, keeper, types.Revoked)
		default:
			return nil, sdk.ErrUnknownRequest("unknown compliance query endpoint")
		}
	}
}

func queryComplianceInfo(ctx sdk.Context, path []string, keeper Keeper,
	requestedState types.ComplianceState) (res []byte, err sdk.Error) {
	vid, err := conversions.ParseVID(path[0])
	if err != nil {
		return nil, err
	}

	pid, err := conversions.ParsePID(path[1])
	if err != nil {
		return nil, err
	}

	certificationType := types.CertificationType(path[2])

	if !keeper.IsComplianceInfoPresent(ctx, certificationType, vid, pid) {
		return nil, types.ErrComplianceInfoDoesNotExist(vid, pid, certificationType)
	}

	complianceInfo := keeper.GetComplianceInfo(ctx, certificationType, vid, pid)

	if len(requestedState) != 0 {
		if complianceInfo.State != requestedState {
			return nil, types.ErrComplianceInfoDoesNotExist(vid, pid, certificationType)
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
	requestedState types.ComplianceState) (res []byte, err sdk.Error) {
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
		if len(requestedState) != 0 && complianceInfo.State != requestedState {
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
				CertificationType: complianceInfo.CertificationType,
			})
			return false
		}

		return false
	})

	res = codec.MustMarshalJSONIndent(keeper.cdc, result)

	return res, nil
}
