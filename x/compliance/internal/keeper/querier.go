package keeper

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/pagination"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/conversions"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryCertifiedModel     = "certified_model"
	QueryAllCertifiedModels = "all_certified_models"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryCertifiedModel:
			return queryCertifiedModel(ctx, path[1:], req, keeper)
		case QueryAllCertifiedModels:
			return queryAllCertifiedModels(ctx, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown compliance query endpoint")
		}
	}
}

func queryCertifiedModel(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	vid, err := conversions.ParseVID(path[0])
	if err != nil {
		return nil, err
	}

	pid, err := conversions.ParsePID(path[1])
	if err != nil {
		return nil, err
	}

	if !keeper.IsCertifiedModelPresent(ctx, vid, pid) {
		return nil, types.ErrDeviceComplianceoDoesNotExist()
	}

	certifiedModel := keeper.GetCertifiedModel(ctx, vid, pid)

	res, err_ := codec.MarshalJSONIndent(keeper.cdc, certifiedModel)
	if err_ != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

func queryAllCertifiedModels(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var params pagination.PaginationParams
	if err := keeper.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to parse params: %s", err))
	}

	result := types.ListCertifiedModelItems{
		Total: keeper.CountTotalCertifiedModel(ctx),
		Items: []types.CertifiedModel{},
	}

	skipped := 0

	keeper.IterateCertifiedModels(ctx, func(certifiedModel types.CertifiedModel) (stop bool) {
		if skipped < params.Skip {
			skipped++
			return false
		}
		if len(result.Items) < params.Take || params.Take == 0 {
			result.Items = append(result.Items, certifiedModel)
			return false
		}

		return true
	})

	res, err := codec.MarshalJSONIndent(keeper.cdc, result)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
