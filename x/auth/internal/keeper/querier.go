package keeper

//nolint:goimports
import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/pagination"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryAccount             = "account"
	QueryAllAccounts         = "all_accounts"
	QueryAllProposedAccounts = "all_proposed_accounts"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryAccount:
			return queryAccount(ctx, req, keeper)
		case QueryAllAccounts:
			return queryAllAccounts(ctx, req, keeper)
		case QueryAllProposedAccounts:
			return queryAllProposedAccounts(ctx, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown auth query endpoint")
		}
	}
}

func queryAccount(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var params types.QueryAccountParams
	if err := keeper.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("failed to parse request params: %s", err))
	}

	if !keeper.IsAccountPresent(ctx, params.Address) {
		return nil, types.ErrAccountDoesNotExist(params.Address.String())
	}

	account := keeper.GetAccount(ctx, params.Address)

	res := codec.MustMarshalJSONIndent(keeper.cdc, account)

	return res, nil
}

func queryAllAccounts(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	var params pagination.PaginationParams
	if err := keeper.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("failed to parse request params: %s", err))
	}

	result := types.ListAccountItems{
		Total: 0,
		Items: []types.Account{},
	}
	skipped := 0

	keeper.IterateAccounts(ctx, func(account types.Account) (stop bool) {
		result.Total++

		if skipped < params.Skip {
			skipped++
			return false
		}

		if len(result.Items) < params.Take || params.Take == 0 {
			result.Items = append(result.Items, account)
			return false
		}

		return false
	})

	res = codec.MustMarshalJSONIndent(keeper.cdc, result)

	return res, nil
}

func queryAllProposedAccounts(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	var params pagination.PaginationParams
	if err := keeper.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("failed to parse request params: %s", err))
	}

	result := types.ListProposedAccountItems{
		Total: 0,
		Items: []types.PendingAccount{},
	}
	skipped := 0

	keeper.IterateProposedAccounts(ctx, func(account types.PendingAccount) (stop bool) {
		result.Total++

		if skipped < params.Skip {
			skipped++
			return false
		}

		if len(result.Items) < params.Take || params.Take == 0 {
			result.Items = append(result.Items, account)
			return false
		}

		return false
	})

	res = codec.MustMarshalJSONIndent(keeper.cdc, result)

	return res, nil
}
