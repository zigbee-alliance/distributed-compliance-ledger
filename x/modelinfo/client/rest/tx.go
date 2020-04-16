package rest

import (
	restutils "git.dsr-corporation.com/zb-ledger/zb-ledger/utils/tx/rest"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

type ModelInfoRequest struct {
	BaseReq                  rest.BaseReq `json:"base_req"`
	VID                      int16        `json:"vid"`
	PID                      int16        `json:"pid"`
	CID                      int16        `json:"cid,omitempty"`
	Name                     string       `json:"name"`
	Description              string       `json:"description"`
	SKU                      string       `json:"sku"`
	FirmwareVersion          string       `json:"firmware_version"`
	HardwareVersion          string       `json:"hardware_version"`
	Custom                   string       `json:"custom,omitempty"`
	TisOrTrpTestingCompleted bool         `json:"tis_or_trp_testing_completed"`
}

func addModelHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ModelInfoRequest
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

		msg := types.NewMsgAddModelInfo(req.VID, req.PID, req.CID, req.Name, req.Description, req.SKU, req.FirmwareVersion,
			req.HardwareVersion, req.Custom, req.TisOrTrpTestingCompleted, from)

		restutils.ProcessMessage(cliCtx, w, r, baseReq, msg, from)
	}
}

func updateModelHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ModelInfoRequest

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

		msg := types.NewMsgUpdateModelInfo(req.VID, req.PID, req.CID, req.Description, req.Custom, req.TisOrTrpTestingCompleted, from)

		restutils.ProcessMessage(cliCtx, w, r, baseReq, msg, from)
	}
}
