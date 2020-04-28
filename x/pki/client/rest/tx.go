package rest

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/rest"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	restTypes "github.com/cosmos/cosmos-sdk/types/rest"
	"net/http"
)

type AddCertificateRequest struct {
	BaseReq     restTypes.BaseReq `json:"base_req"`
	Certificate string            `json:"cert"`
}

type ApproveCertificateRequest struct {
	BaseReq      restTypes.BaseReq `json:"base_req"`
	Subject      string            `json:"subject"`
	SubjectKeyId string            `json:"subject_key_id"`
}

func proposeAddX509RootCertHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		var req AddCertificateRequest
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

		msg := types.NewMsgProposeAddX509RootCert(req.Certificate, restCtx.Signer())

		restCtx.HandleWriteRequest(msg)
	}
}

func approveAddX509RootCertHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		var req ApproveCertificateRequest
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

		msg := types.NewMsgApproveAddX509RootCert(req.Subject, req.SubjectKeyId, restCtx.Signer())

		restCtx.HandleWriteRequest(msg)
	}
}

func addX509CertHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		var req AddCertificateRequest
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

		msg := types.NewMsgAddX509Cert(req.Certificate, restCtx.Signer())

		restCtx.HandleWriteRequest(msg)
	}
}
