package rest

import (
	"net/http"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/rest"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	restTypes "github.com/cosmos/cosmos-sdk/types/rest"
)

type ProposeAddRootCertificateRequest struct {
	BaseReq restTypes.BaseReq `json:"base_req"`
	Cert    string            `json:"cert"`
}

type AddCertificateRequest struct {
	BaseReq restTypes.BaseReq `json:"base_req"`
	Cert    string            `json:"cert"`
}

type RevokeCertificateRequest struct {
	BaseReq      restTypes.BaseReq `json:"base_req"`
	Subject      string            `json:"subject"`
	SubjectKeyID string            `json:"subject_key_id"`
}

func proposeAddX509RootCertHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		var req ProposeAddRootCertificateRequest
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

		msg := types.NewMsgProposeAddX509RootCert(req.Cert, restCtx.Signer())

		restCtx.HandleWriteRequest(msg)
	}
}

func approveAddX509RootCertHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		vars := restCtx.Variables()

		var req rest.BasicReq
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

		msg := types.NewMsgApproveAddX509RootCert(vars[subject], vars[subjectKeyID], restCtx.Signer())

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

		msg := types.NewMsgAddX509Cert(req.Cert, restCtx.Signer())

		restCtx.HandleWriteRequest(msg)
	}
}

func revokeX509CertHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		var req RevokeCertificateRequest
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

		msg := types.NewMsgRevokeX509Cert(req.Subject, req.SubjectKeyID, restCtx.Signer())

		restCtx.HandleWriteRequest(msg)
	}
}
