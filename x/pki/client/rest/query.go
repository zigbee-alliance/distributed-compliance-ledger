package rest

import (
	"encoding/json"
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/pagination"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"net/http"
)

func getAllX509RootCertsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		getListCertificates(cliCtx, w, r, fmt.Sprintf("custom/%s/all_x509_root_certs", storeName))
	}
}

func getAllProposedX509RootCertsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		getListCertificates(cliCtx, w, r, fmt.Sprintf("custom/%s/all_proposed_x509_root_certs", storeName))
	}
}

func getAllX509CertsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		getListCertificates(cliCtx, w, r, fmt.Sprintf("custom/%s/all_x509_certs", storeName))
	}
}

func getAllSubjectX509CertsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		getListCertificates(cliCtx, w, r, fmt.Sprintf("custom/%s/all_subject_x509_certs/%s", storeName, subject))
	}
}

func getProposedX509RootCertHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx := context.NewCLIContext().WithCodec(cliCtx.Codec)

		vars := mux.Vars(r)
		subject := vars[subject]
		subjectKeyId := vars[subjectKeyId]

		res, height, err := cliCtx.QueryStore([]byte(keeper.ProposedCertificateId(subject, subjectKeyId)), storeName)
		if err != nil || res == nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, types.ErrProposedCertificateDoesNotExist(subject, subjectKeyId).Error())
			return
		}

		var proposedCertificate types.ProposedCertificate
		cliCtx.Codec.MustUnmarshalBinaryBare(res, &proposedCertificate)

		out, err := json.Marshal(proposedCertificate)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, out)
	}
}

func getX509CertHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx := context.NewCLIContext().WithCodec(cliCtx.Codec)

		vars := mux.Vars(r)
		subject := vars[subject]
		subjectKeyId := vars[subjectKeyId]

		res, height, err := cliCtx.QueryStore([]byte(keeper.CertificateId(subject, subjectKeyId)), storeName)
		if err != nil || res == nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, types.ErrCertificateDoesNotExist(subject, subjectKeyId).Error())
			return
		}

		var certificate types.Certificate
		cliCtx.Codec.MustUnmarshalBinaryBare(res, &res)

		out, err := json.Marshal(certificate)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, out)
	}
}

func getListCertificates(cliCtx context.CLIContext, w http.ResponseWriter, r *http.Request, path string) {
	cliCtx = context.NewCLIContext().WithCodec(cliCtx.Codec)

	subjectKeyId := r.FormValue(subjectKeyId)
	serialNumber := r.FormValue(serialNumber)
	rootSubjectKeyId := r.FormValue(rootSubjectKeyId)

	paginationParams, err := pagination.ParsePaginationParamsFromRequest(cliCtx.Codec, r)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	params := types.NewListCertificatesQueryParams(paginationParams, subjectKeyId, serialNumber, rootSubjectKeyId)

	res, height, err := cliCtx.QueryWithData(path, cliCtx.Codec.MustMarshalJSON(params))
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	cliCtx.Height = height
	rest.PostProcessResponse(w, cliCtx, res)
}
