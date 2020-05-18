package rest

//nolint:goimports
import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/rpc"
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

		if ip := r.FormValue(node); len(ip) > 0 {
			restCtx = restCtx.WithNodeURI(ip)
		}

		status, err := restCtx.NodeStatus()
		if err != nil {
			return
		}

		restCtx.RespondWithHeight(status, status.SyncInfo.LatestBlockHeight)
	}
}

// Validator Set at a height REST handler.
func ValidatorSetRequestHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		err := restCtx.Request().ParseForm()
		if err != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest,
				sdk.AppendMsgToErr("could not parse query parameters", err.Error()))
			return
		}

		var height_ int64

		if h := r.FormValue(height); len(h) > 0 {
			height_, err = strconv.ParseInt(h, 10, 64)
			if err != nil {
				restCtx.WriteErrorResponse(http.StatusBadRequest, "Invalid height: it must be integer")
				return
			}
		}

		chainHeight, err := restCtx.GetChainHeight()
		if err != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest, err.Error())
			return
		}

		if height_ > chainHeight {
			restCtx.WriteErrorResponse(http.StatusNotFound,
				fmt.Sprintf("Invalid height: It must not be bigger then the chain height: \"%v\"", chainHeight))
			return
		}

		if height_ == 0 {
			height_ = chainHeight
		}

		output, err := rpc.GetValidators(cliCtx, &height_)
		if err != nil {
			restCtx.WriteErrorResponse(http.StatusInternalServerError, err.Error())
			return
		}

		restCtx.RespondWithHeight(output, height_)
	}
}
