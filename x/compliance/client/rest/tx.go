package rest

import (
	"fmt"
	restutils "git.dsr-corporation.com/zb-ledger/zb-ledger/utils/tx/rest"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"net/http"
	"time"
)

type CertifyModelRequest struct {
	BaseReq           rest.BaseReq            `json:"base_req"`
	VID               int16                   `json:"vid"`
	PID               int16                   `json:"pid"`
	CertificationDate time.Time               `json:"certification_date"` // rfc3339 encoded date
	CertificationType types.CertificationType `json:"certification_type,omitempty"`
	Reason            string                  `json:"reason,omitempty"`
}

func certifyModelHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CertifyModelRequest
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

		msg := types.NewMsgCertifyModel(req.VID, req.PID, req.CertificationDate, req.CertificationType, req.Reason, from)

		restutils.ProcessMessage(cliCtx, w, r, baseReq, msg, from)
	}
}

type RevokeModelRequest struct {
	BaseReq           rest.BaseReq            `json:"base_req"`
	VID               int16                   `json:"vid"`
	PID               int16                   `json:"pid"`
	RevocationDate    time.Time               `json:"revocation_date"` // rfc3339 encoded date
	CertificationType types.CertificationType `json:"certification_type,omitempty"`
	Reason            string                  `json:"reason,omitempty"`
}

func revokeModelHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RevokeModelRequest
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

		msg := types.NewMsgRevokeModel(req.VID, req.PID, req.RevocationDate, req.CertificationType, req.Reason, from)

		restutils.ProcessMessage(cliCtx, w, r, baseReq, msg, from)
	}
}
