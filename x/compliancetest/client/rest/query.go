package rest

import (
	"encoding/json"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"net/http"
)

func getTestingResultHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx := context.NewCLIContext().WithCodec(cliCtx.Codec)

		vars := mux.Vars(r)
		vid := vars[vid]
		pid := vars[pid]

		res, height, err := cliCtx.QueryStore([]byte(keeper.TestingResultId(vid, pid)), storeName)
		if err != nil || res == nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, types.ErrTestingResultDoesNotExist(vid, pid).Error())
			return
		}

		var testingResult types.TestingResults
		cliCtx.Codec.MustUnmarshalBinaryBare(res, &testingResult)

		out, err := json.Marshal(testingResult)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, out)
	}
}
