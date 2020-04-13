package keeper

import (
	"fmt"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryTestingResult = "testresult"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryTestingResult:
			return queryTestingResult(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown compliancetest query endpoint")
		}
	}
}

func queryTestingResult(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	vid, err := types.ParseVID(path[0])
	if err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to parse vid: %s", err))
	}

	pid, err := types.ParsePID(path[1])
	if err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to parse pid: %s", err))
	}

	if !keeper.IsTestingResultsPresents(ctx, vid, pid) {
		return nil, types.ErrTestingResultDoesNotExist()
	}

	testingResult := keeper.GetTestingResults(ctx, vid, pid)

	res, err := codec.MarshalJSONIndent(keeper.cdc, testingResult)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
