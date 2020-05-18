package rest

//nolint:goimports
import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/rest"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"net/http"
	"time"

	restTypes "github.com/cosmos/cosmos-sdk/types/rest"
)

type CertifyModelRequest struct {
	BaseReq           restTypes.BaseReq       `json:"base_req"`
	VID               uint16                  `json:"vid"`
	PID               uint16                  `json:"pid"`
	CertificationDate time.Time               `json:"certification_date"` // rfc3339 encoded date
	CertificationType types.CertificationType `json:"certification_type"`
	Reason            string                  `json:"reason,omitempty"`
}

func certifyModelHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		var req CertifyModelRequest
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

		msg := types.NewMsgCertifyModel(req.VID, req.PID, req.CertificationDate,
			req.CertificationType, req.Reason, restCtx.Signer())

		restCtx.HandleWriteRequest(msg)
	}
}

type RevokeModelRequest struct {
	BaseReq           restTypes.BaseReq       `json:"base_req"`
	VID               uint16                  `json:"vid"`
	PID               uint16                  `json:"pid"`
	RevocationDate    time.Time               `json:"revocation_date"` // rfc3339 encoded date
	CertificationType types.CertificationType `json:"certification_type"`
	Reason            string                  `json:"reason,omitempty"`
}

func revokeModelHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		var req RevokeModelRequest
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

		msg := types.NewMsgRevokeModel(req.VID, req.PID, req.RevocationDate,
			req.CertificationType, req.Reason, restCtx.Signer())

		restCtx.HandleWriteRequest(msg)
	}
}
