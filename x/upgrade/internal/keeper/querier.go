package keeper

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/upgrade/internal/types"
)

// NewQuerier creates a querier for upgrade cli and REST endpoints
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, sdk.Error) {
		switch path[0] {

		case types.QueryCurrent:
			return queryCurrent(ctx, req, k)

		case types.QueryApplied:
			return queryApplied(ctx, req, k)

		default:
			return nil, types.NewError(types.ErrUnknownRequest, "unknown %s query endpoint: %s", types.ModuleName, path[0])
		}
	}
}

func queryCurrent(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, sdk.Error) {
	plan, has := k.GetUpgradePlan(ctx)
	if !has {
		return nil, nil
	}

	res, err := k.cdc.MarshalJSON(&plan)
	if err != nil {
		return nil, types.NewError(types.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryApplied(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, sdk.Error) {
	var params types.QueryAppliedParams

	err := k.cdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, types.NewError(types.ErrJSONUnmarshal, err.Error())
	}

	applied := k.GetDoneHeight(ctx, params.Name)
	if applied == 0 {
		return nil, nil
	}

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, uint64(applied))

	return bz, nil
}
