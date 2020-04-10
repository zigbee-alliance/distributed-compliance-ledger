package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"

	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/cosmos/cosmos-sdk/types/rest"
)

func modelInfoHeadersHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		skip, _ := strconv.Atoi(r.FormValue("skip"))
		take, _ := strconv.Atoi(r.FormValue("take"))

		params := types.NewQueryModelInfoHeadersParams(skip, take)

		data := cliCtx.Codec.MustMarshalJSON(params)

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/model_info_headers", storeName), data)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func modelInfoHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		vid := vars[vid]
		pid := vars[pid]

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/model_info/%s/%s", storeName, vid, pid), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}
