package rest

import (
	"fmt"
	"net/http"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/rest"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
		if err != nil {
			restCtx.WriteErrorResponse(http.StatusNotFound, err.Error())
			return
		}

		var account types.Account

		cliCtx.Codec.MustUnmarshalBinaryBare(res, &account)
		restCtx.RespondWithHeight(types.ZBAccount(account), height)
	}
}
