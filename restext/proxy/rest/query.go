package rest

import (
	"net/http"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/client/context"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/rest"
)

func BlocksHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		err := restCtx.Request().ParseForm()
		if err != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest,
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

		res, err := restCtx.BlockchainInfo(minHeight, maxHeight)
		if err != nil {
			restCtx.WriteErrorResponse(http.StatusNotFound, err.Error())
			return
		}

		restCtx.PostProcessResponse(res)
	}
}

func NodeStatusHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		err := restCtx.Request().ParseForm()
		if err != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest,
				sdk.AppendMsgToErr("could not parse query parameters", err.Error()))
			return
		}

		nodeIp := "tcp://localhost:26657" // default node
		if ip := r.FormValue(node); len(ip) > 0 {
			nodeIp = ip
		}

		restCtx = restCtx.WithNodeURI(nodeIp)

		status, err := restCtx.NodeStatus()
		if err != nil {
			return
		}

		restCtx.RespondWithHeight(status, status.SyncInfo.LatestBlockHeight)
	}
}
