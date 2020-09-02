package rest

//nolint:goimports
import (
	"net/http"
	"time"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/rest"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	restTypes "github.com/cosmos/cosmos-sdk/types/rest"
)

type TestingResultRequest struct {
	BaseReq    restTypes.BaseReq `json:"base_req"`
	VID        uint16            `json:"vid"`
	PID        uint16            `json:"pid"`
	TestResult string            `json:"test_result"`
	TestDate   time.Time         `json:"test_date"` // rfc3339 encoded date
}

func addTestingResultHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		var req TestingResultRequest
		if !restCtx.ReadRESTReq(&req) {
			return
		}

		restCtx, err := restCtx.WithBaseRequest(req.BaseReq)
		if err != nil {
			return
		}

		restCtx, err = restCtx.WithSigner()
		if err != nil {
			return
		}

		msg := types.NewMsgAddTestingResult(req.VID, req.PID, req.TestResult, req.TestDate, restCtx.Signer())

		restCtx.HandleWriteRequest(msg)
	}
}
