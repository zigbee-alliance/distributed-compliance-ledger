package rest

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/auth"
	restutils "git.dsr-corporation.com/zb-ledger/zb-ledger/utils/tx/rest"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
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

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		processMessage(cliCtx, w, r, baseReq, msg, from)
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

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		processMessage(cliCtx, w, r, baseReq, msg, from)
	}
}

func processMessage(cliCtx context.CLIContext, w http.ResponseWriter, r *http.Request, baseReq rest.BaseReq, msg sdk.Msg, signer sdk.AccAddress) {
	account, passphrase, err := auth.GetCredentialsFromRequest(r)
	if err != nil { // No credentials - just generate request message
		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
		return
	}

	// Credentials are found - sign and broadcast message
	res, err_ := restutils.SignAndBroadcastMessage(cliCtx, baseReq.ChainID, signer, account, passphrase, []sdk.Msg{msg})
	if err_ != nil {
		rest.WriteErrorResponse(w, http.StatusInternalServerError, err_.Error())
		return
	}

	rest.PostProcessResponse(w, cliCtx, res)
}
