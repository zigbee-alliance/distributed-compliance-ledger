package rest

import (
	"fmt"
	restutils "git.dsr-corporation.com/zb-ledger/zb-ledger/utils/tx/rest"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"net/http"
)

type AddCertificateRequest struct {
	BaseReq     rest.BaseReq `json:"base_req"`
	Certificate string       `json:"cert"`
}

type ApproveCertificateRequest struct {
	BaseReq      rest.BaseReq `json:"base_req"`
	Subject      string       `json:"subject"`
	SubjectKeyId string       `json:"subject_key_id"`
}

func proposeAddX509RootCertHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx := context.NewCLIContext().WithCodec(cliCtx.Codec)
		var req AddCertificateRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		from, err := sdk.AccAddressFromBech32(baseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Request Parsing Error: %v. `from` must be a valid address", err))
			return
		}

		msg := types.NewMsgProposeAddX509RootCert(req.Certificate, from)

		restutils.ProcessMessage(cliCtx, w, r, baseReq, msg, from)
	}
}

func approveAddX509RootCertHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx := context.NewCLIContext().WithCodec(cliCtx.Codec)
		var req ApproveCertificateRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		from, err := sdk.AccAddressFromBech32(baseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Request Parsing Error: %v. `from` must be a valid address", err))
			return
		}

		msg := types.NewMsgApproveAddX509RootCert(req.Subject, req.SubjectKeyId, from)

		restutils.ProcessMessage(cliCtx, w, r, baseReq, msg, from)
	}
}

func addX509CertHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx := context.NewCLIContext().WithCodec(cliCtx.Codec)
		var req AddCertificateRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		from, err := sdk.AccAddressFromBech32(baseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Request Parsing Error: %v. `from` must be a valid address", err))
			return
		}

		msg := types.NewMsgAddX509Cert(req.Certificate, from)

		restutils.ProcessMessage(cliCtx, w, r, baseReq, msg, from)
	}
}
