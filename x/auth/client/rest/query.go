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

package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/rest"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth/internal/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth/internal/types"
)

func accountsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		params, err := restCtx.ParsePaginationParams()
		if err != nil {
			return
		}

		restCtx.QueryList(fmt.Sprintf("custom/%s/%s", storeName, keeper.QueryAllAccounts), params)
	}
}

func accountsRangeHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		valueUnmarshaler := func(bytes []byte) json.RawMessage {
			value := types.Account{}
			restCtx.Codec().MustUnmarshalBinaryBare(bytes, &value)

			// the trick to prevent appending of `type` field by cdc
			return codec.Cdc.MustMarshalJSON(value)
		}

		restCtx.QueryRangeWithTotalAndHandleIO(
			storeName, types.AccountPrefix, types.AccountsTotalKey, valueUnmarshaler)
	}
}

func proposedAccountsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		params, err := restCtx.ParsePaginationParams()
		if err != nil {
			return
		}

		restCtx.QueryList(fmt.Sprintf("custom/%s/%s", storeName, keeper.QueryAllPendingAccounts), params)
	}
}

func proposedAccountsRangeHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		valueUnmarshaler := func(bytes []byte) json.RawMessage {
			value := types.PendingAccount{}
			restCtx.Codec().MustUnmarshalBinaryBare(bytes, &value)

			return restCtx.Codec().MustMarshalJSON(value)
		}

		restCtx.QueryRangeWithTotalAndHandleIO(
			storeName, types.PendingAccountPrefix, types.PendingAccountsTotalKey, valueUnmarshaler)
	}
}

func proposedAccountsToRevokeHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		params, err := restCtx.ParsePaginationParams()
		if err != nil {
			return
		}

		restCtx.QueryList(fmt.Sprintf("custom/%s/%s", storeName, keeper.QueryAllPendingAccountRevocations), params)
	}
}

func proposedAccountsToRevokeRangeHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		valueUnmarshaler := func(bytes []byte) json.RawMessage {
			value := types.PendingAccountRevocation{}
			restCtx.Codec().MustUnmarshalBinaryBare(bytes, &value)

			return restCtx.Codec().MustMarshalJSON(value)
		}

		restCtx.QueryRangeWithTotalAndHandleIO(
			storeName, types.PendingAccountRevocationPrefix, types.PendingAccountRevocationsTotalKey, valueUnmarshaler)
	}
}

func accountHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		vars := restCtx.Variables()
		accAddr := vars[address]

		address, err := sdk.AccAddressFromBech32(accAddr)
		if err != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest, sdk.ErrInvalidAddress(accAddr).Error())

			return
		}

		res, height, err := cliCtx.QueryStore(types.GetAccountKey(address), storeName)
		if err != nil || res == nil {
			restCtx.WriteErrorResponse(http.StatusNotFound, types.ErrAccountDoesNotExist(address).Error())

			return
		}

		var account types.Account

		cliCtx.Codec.MustUnmarshalBinaryBare(res, &account)
		// the trick to prevent appending of `type` field by cdc
		restCtx.RespondWithHeight(types.ZBAccount(account), height)
	}
}
