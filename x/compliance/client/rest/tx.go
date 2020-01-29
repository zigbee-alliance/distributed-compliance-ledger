package rest

import (
	"net/http"
	"time"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

type addModelInfoReq struct {
	BaseReq                  rest.BaseReq `json:"base_req"`
	ID                       string       `json:"id"`
	Name                     string       `json:"name"`
	Owner                    string       `json:"owner"`
	Description              string       `json:"description"`
	SKU                      string       `json:"sku"`
	FirmwareVersion          string       `json:"firmware_version"`
	HardwareVersion          string       `json:"hardware_version"`
	CertificateID            string       `json:"certificate_id"`
	CertifiedDate            time.Time    `json:"certified_date"`
	TisOrTrpTestingCompleted bool         `json:"tis_or_trp_testing_completed"`
}

func addModelInfoHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req addModelInfoReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		owner, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		signer, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgAddModelInfo(req.ID, req.Name, owner, req.Description, req.SKU, req.FirmwareVersion,
			req.HardwareVersion, req.CertificateID, req.CertifiedDate, req.TisOrTrpTestingCompleted, signer)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
