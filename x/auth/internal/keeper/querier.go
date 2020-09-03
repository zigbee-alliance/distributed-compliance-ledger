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

	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/pagination"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryAccount                      = "account"
	QueryAllAccounts                  = "all_accounts"
	QueryAllPendingAccounts           = "all_pending_accounts"
	QueryAllPendingAccountRevocations = "all_pending_account_revocations"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryAccount:
			return queryAccount(ctx, req, keeper)
		case QueryAllAccounts:
			return queryAllAccounts(ctx, req, keeper)
		case QueryAllPendingAccounts:
			return queryAllPendingAccounts(ctx, req, keeper)
		case QueryAllPendingAccountRevocations:
			return queryAllPendingAccountRevocations(ctx, req, keeper)
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

// nolint:dupl
func queryAllAccounts(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	var params pagination.PaginationParams
	if err := keeper.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("failed to parse request params: %s", err))
	}

	result := types.ListAccounts{
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

// nolint:dupl
func queryAllPendingAccounts(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	var params pagination.PaginationParams
	if err := keeper.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("failed to parse request params: %s", err))
	}

	result := types.ListPendingAccounts{
		Total: 0,
		Items: []types.PendingAccount{},
	}
	skipped := 0

	keeper.IteratePendingAccounts(ctx, func(pendAcc types.PendingAccount) (stop bool) {
		result.Total++

		if skipped < params.Skip {
			skipped++

			return false
		}

		if len(result.Items) < params.Take || params.Take == 0 {
			result.Items = append(result.Items, pendAcc)

			return false
		}

		return false
	})

	res = codec.MustMarshalJSONIndent(keeper.cdc, result)

	return res, nil
}

// nolint:dupl
func queryAllPendingAccountRevocations(ctx sdk.Context,
	req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	var params pagination.PaginationParams
	if err := keeper.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("failed to parse request params: %s", err))
	}

	result := types.ListPendingAccountRevocations{
		Total: 0,
		Items: []types.PendingAccountRevocation{},
	}
	skipped := 0

	keeper.IteratePendingAccountRevocations(ctx, func(revocation types.PendingAccountRevocation) (stop bool) {
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
