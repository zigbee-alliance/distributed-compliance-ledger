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

type addModelReq struct {
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
		var req addModelReq

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

type updateModelReq struct {
	BaseReq                     rest.BaseReq   `json:"base_req"`
	VID                         int16          `json:"vid"`
	PID                         int16          `json:"pid"`
	NewCID                      int16          `json:"new_cid,omitempty"`
	NewDescription              string         `json:"new_description"`
	NewCustom                   string         `json:"new_custom,omitempty"`
	NewCertificateID            string         `json:"new_certificate_id,omitempty"`
	NewCertifiedDate            time.Time      `json:"new_certified_date,omitempty"`
	NewTisOrTrpTestingCompleted bool           `json:"new_tis_or_trp_testing_completed"`
	Signer                      sdk.AccAddress `json:"signer"`
}

func updateModelHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req updateModelReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgUpdateModelInfo(req.VID, req.PID, req.NewCID, req.NewDescription, req.NewCustom,
			req.NewCertificateID, req.NewCertifiedDate, req.NewTisOrTrpTestingCompleted, cliCtx.GetFromAddress())

		err := msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
