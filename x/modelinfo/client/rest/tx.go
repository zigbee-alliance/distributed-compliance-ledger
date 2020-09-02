package rest

//nolint:goimports
import (
	"net/http"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/rest"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	restTypes "github.com/cosmos/cosmos-sdk/types/rest"
)

type ModelInfoRequest struct {
	BaseReq                  restTypes.BaseReq `json:"base_req"`
	Name                     string            `json:"name"`
	Description              string            `json:"description"`
	SKU                      string            `json:"sku"`
	FirmwareVersion          string            `json:"firmware_version"`
	HardwareVersion          string            `json:"hardware_version"`
	Custom                   string            `json:"custom,omitempty"`
	TisOrTrpTestingCompleted bool              `json:"tis_or_trp_testing_completed"`
	VID                      uint16            `json:"vid"`
	PID                      uint16            `json:"pid"`
	CID                      uint16            `json:"cid,omitempty"`
}

func addModelHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		var req ModelInfoRequest
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

		msg := types.NewMsgAddModelInfo(req.VID, req.PID, req.CID, req.Name, req.Description, req.SKU,
			req.FirmwareVersion, req.HardwareVersion, req.Custom, req.TisOrTrpTestingCompleted, restCtx.Signer())

		restCtx.HandleWriteRequest(msg)
	}
}

func updateModelHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		var req ModelInfoRequest
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

		msg := types.NewMsgUpdateModelInfo(req.VID, req.PID, req.CID, req.Description,
			req.Custom, req.TisOrTrpTestingCompleted, restCtx.Signer())

		restCtx.HandleWriteRequest(msg)
	}
}
