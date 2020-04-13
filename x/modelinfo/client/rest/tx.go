package rest

import (
	"net/http"
	"time"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

type ModelInfoRequest struct {
	BaseReq                  rest.BaseReq   `json:"base_req"`
	VID                      int16          `json:"vid"`
	PID                      int16          `json:"pid"`
	CID                      int16          `json:"cid,omitempty"`
	Name                     string         `json:"name"`
	Description              string         `json:"description"`
	SKU                      string         `json:"sku"`
	FirmwareVersion          string         `json:"firmware_version"`
	HardwareVersion          string         `json:"hardware_version"`
	Custom                   string         `json:"custom,omitempty"`
	CertificateID            string         `json:"certificate_id,omitempty"`
	CertifiedDate            time.Time      `json:"certified_date,omitempty"`
	TisOrTrpTestingCompleted bool           `json:"tis_or_trp_testing_completed"`
	Signer                   sdk.AccAddress `json:"signer"`
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

		msg := types.NewMsgAddModelInfo(req.VID, req.PID, req.CID, req.Name, req.Description, req.SKU, req.FirmwareVersion,
			req.HardwareVersion, req.Custom, req.CertificateID, req.CertifiedDate, req.TisOrTrpTestingCompleted, cliCtx.GetFromAddress())

		err := msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
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

		msg := types.NewMsgUpdateModelInfo(req.VID, req.PID, req.CID, req.Name, req.Description, req.SKU, req.FirmwareVersion,
			req.HardwareVersion, req.Custom, req.CertificateID, req.CertifiedDate, req.TisOrTrpTestingCompleted, cliCtx.GetFromAddress())

		err := msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
