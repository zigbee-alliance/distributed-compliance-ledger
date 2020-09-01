package rest

//nolint:goimports
import (
	"net/http"
	"time"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/conversions"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/rest"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"

	restTypes "github.com/cosmos/cosmos-sdk/types/rest"
)

type CertifyModelRequest struct {
	BaseReq           restTypes.BaseReq `json:"base_req"`
	CertificationDate time.Time         `json:"certification_date"` // rfc3339 encoded date
	Reason            string            `json:"reason,omitempty"`
}

func certifyModelHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		vars := restCtx.Variables()

		vid, err_ := conversions.ParseVID(vars[vid])
		if err_ != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest, err_.Error())
			return
		}

		pid, err_ := conversions.ParsePID(vars[pid])
		if err_ != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest, err_.Error())
			return
		}

		certificationType := types.CertificationType(vars[certificationType])

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

		msg := types.NewMsgCertifyModel(vid, pid, req.CertificationDate,
			certificationType, req.Reason, restCtx.Signer())

		restCtx.HandleWriteRequest(msg)
	}
}

type RevokeModelRequest struct {
	BaseReq        restTypes.BaseReq `json:"base_req"`
	RevocationDate time.Time         `json:"revocation_date"` // rfc3339 encoded date
	Reason         string            `json:"reason,omitempty"`
}

func revokeModelHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		vars := restCtx.Variables()

		vid, err_ := conversions.ParseVID(vars[vid])
		if err_ != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest, err_.Error())
			return
		}

		pid, err_ := conversions.ParsePID(vars[pid])
		if err_ != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest, err_.Error())
			return
		}

		certificationType := types.CertificationType(vars[certificationType])

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

		msg := types.NewMsgRevokeModel(vid, pid, req.RevocationDate,
			certificationType, req.Reason, restCtx.Signer())

		restCtx.HandleWriteRequest(msg)
	}
}
