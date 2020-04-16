package rest

import (
	"fmt"
	restutils "git.dsr-corporation.com/zb-ledger/zb-ledger/utils/tx/rest"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"net/http"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

type TestingResultRequest struct {
	BaseReq    rest.BaseReq `json:"base_req"`
	VID        int16        `json:"vid"`
	PID        int16        `json:"pid"`
	TestResult string       `json:"test_result"`
	TestDate   time.Time    `json:"test_date"` // rfc3339 encoded date
}

func addTestingResultHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx := context.NewCLIContext().WithCodec(cliCtx.Codec)
		var req TestingResultRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		from, err := sdk.AccAddressFromBech32(baseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Request Parsing Error: %v. `from` must be a valid address", err))
			return
		}

		msg := types.NewMsgAddTestingResult(req.VID, req.PID, req.TestResult, req.TestDate, from)

		restutils.ProcessMessage(cliCtx, w, r, baseReq, msg, from)
	}
}
