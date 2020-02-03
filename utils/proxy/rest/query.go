package rest

import (
	"net/http"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/cosmos/cosmos-sdk/types/rest"
)

func BlocksHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest,
				sdk.AppendMsgToErr("could not parse query parameters", err.Error()))
			return
		}

		minHeight, _ := strconv.ParseInt(r.FormValue("minHeight"), 10, 64)
		if minHeight < 0 {
			minHeight = 0
		}

		maxHeight, _ := strconv.ParseInt(r.FormValue("maxHeight"), 10, 64)
		if maxHeight < minHeight {
			maxHeight = minHeight
		}

		res, err := cliCtx.Client.BlockchainInfo(minHeight, maxHeight)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}
