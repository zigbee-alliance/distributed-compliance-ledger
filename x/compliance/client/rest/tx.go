package rest

import (
	restutils "git.dsr-corporation.com/zb-ledger/zb-ledger/utils/tx/rest"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"net/http"
	"time"
)

type CertifiedModelRequest struct {
	BaseReq           rest.BaseReq `json:"base_req"`
	VID               int16        `json:"vid"`
	PID               int16        `json:"pid"`
	CertificationDate time.Time    `json:"certification_date"` // rfc3339 encoded date
	CertificationType string       `json:"certification_type,omitempty"`
}

func certifyModelHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CertifiedModelRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		from, err := sdk.AccAddressFromBech32(baseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		msg := types.NewMsgCertifyModel(req.VID, req.PID, req.CertificationDate, req.CertificationType, from)

		restutils.ProcessMessage(cliCtx, w, r, baseReq, msg, from)
	}
}
