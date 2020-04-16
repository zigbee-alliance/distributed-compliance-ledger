package rest

import (
	"encoding/json"
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/pagination"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"net/http"
)

func getCertifiedModelsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx := context.NewCLIContext().WithCodec(cliCtx.Codec)

		params, err := pagination.ParsePaginationParamsFromRequest(cliCtx.Codec, r)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/all_certified_models", storeName), params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		respond(w, cliCtx, res, height)
	}
}

func getCertifiedModelHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx := context.NewCLIContext().WithCodec(cliCtx.Codec)

		vars := mux.Vars(r)
		vid := vars[vid]
		pid := vars[pid]

		res, height, err := cliCtx.QueryStore([]byte(keeper.CertifiedModelId(vid, pid)), storeName)
		if err != nil || res == nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, types.ErrDeviceComplianceoDoesNotExist(vid, pid).Error())
			return
		}

		var certifiedModel types.CertifiedModel
		cliCtx.Codec.MustUnmarshalBinaryBare(res, &certifiedModel)

		out, err := json.Marshal(certifiedModel)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		respond(w, cliCtx, out, height)
	}
}

func respond(w http.ResponseWriter, cliCtx context.CLIContext, data []byte, height int64) {
	cliCtx.Height = height
	rest.PostProcessResponse(w, cliCtx, data)

}
