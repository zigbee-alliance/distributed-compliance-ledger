package keeper

import (
	"fmt"

	"github.com/askolesov/zb-ledger/x/authnext/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryAccountHeaders = "account_headers"
)

func NewQuerier(accKeeper types.AccountKeeper, cdc *codec.Codec) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryAccountHeaders:
			return queryAccountHeaders(ctx, req, accKeeper, cdc)
		default:
			return nil, sdk.ErrUnknownRequest("unknown authnext query endpoint")
		}
	}
}

func queryAccountHeaders(ctx sdk.Context, req abci.RequestQuery, accKeeper types.AccountKeeper,
	cdc *codec.Codec) ([]byte, sdk.Error) {
	var params types.QueryAccountHeadersParams
	if err := cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to parse params: %s", err))
	}

	var result types.QueryAccountHeadersResult

	skipped := 0

	accKeeper.IterateAccounts(ctx, func(account exported.Account) (stop bool) {
		if skipped < params.Skip {
			skipped++
			return false
		}

		if len(result) < params.Count || params.Count == 0 {
			header := types.AccountHeader{
				Address: account.GetAddress(),
				PubKey:  account.GetPubKey(),
			}

			result = append(result, header)
			return false
		}

		return true
	})

	res, err := codec.MarshalJSONIndent(cdc, result)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
