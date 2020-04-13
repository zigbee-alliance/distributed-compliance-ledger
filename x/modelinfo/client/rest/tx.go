package rest

import (
	restutils "git.dsr-corporation.com/zb-ledger/zb-ledger/utils/tx/rest"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"net/http"

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
	TisOrTrpTestingCompleted bool           `json:"tis_or_trp_testing_completed"`
	Signer                   sdk.AccAddress `json:"signer"`
	Account                  string         `json:"account,omitempty"`
	Passphrase               string         `json:"passphrase,omitempty"`
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
			req.HardwareVersion, req.Custom, req.TisOrTrpTestingCompleted, req.Signer)

		err := msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		processMessage(cliCtx, w, req, msg)
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

		msg := types.NewMsgUpdateModelInfo(req.VID, req.PID, req.Description, req.Custom, req.TisOrTrpTestingCompleted, cliCtx.GetFromAddress())

		err := msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		processMessage(cliCtx, w, req, msg)
	}
}

func processMessage(cliCtx context.CLIContext, w http.ResponseWriter, req ModelInfoRequest, msg sdk.Msg) {
	if len(req.Passphrase) == 0 || len(req.Account) == 0 {
		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
		return
	}

	res, err_ := restutils.SignAndBroadcastMessage(cliCtx, req.Account, req.Passphrase, req.BaseReq.AccountNumber,
		req.BaseReq.Sequence, req.BaseReq.ChainID, []sdk.Msg{msg})
	if err_ != nil {
		rest.WriteErrorResponse(w, http.StatusInternalServerError, err_.Error())
		return
	}

	rest.PostProcessResponse(w, cliCtx, res)
}
