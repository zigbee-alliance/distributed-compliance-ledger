package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authnext/internal/types"

	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/cosmos/cosmos-sdk/types/rest"
)

func accountHeadersHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		skip, _ := strconv.Atoi(r.FormValue("skip"))
		take, _ := strconv.Atoi(r.FormValue("take"))

		params := types.NewQueryAccountHeadersParams(skip, take)

		data := cliCtx.Codec.MustMarshalJSON(params)

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/account_headers", storeName), data)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func accountHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		accAddr := vars[addrKey]

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/account/%s", storeName, accAddr), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}
