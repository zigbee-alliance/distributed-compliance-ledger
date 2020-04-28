package rest

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/rest"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"net/http"
)

func getAllX509RootCertsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)
		getListCertificates(restCtx, fmt.Sprintf("custom/%s/all_x509_root_certs", storeName), "", "")
	}
}

func getAllProposedX509RootCertsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)
		getListCertificates(restCtx, fmt.Sprintf("custom/%s/all_proposed_x509_root_certs", storeName), "", "")
	}
}

func getAllX509CertsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)
		rootSubject := r.FormValue(rootSubject)
		rootSubjectKeyId := r.FormValue(rootSubjectKeyId)
		getListCertificates(restCtx, fmt.Sprintf("custom/%s/all_x509_certs", storeName), rootSubject, rootSubjectKeyId)
	}
}

func getAllSubjectX509CertsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)
		vars := restCtx.Variables()
		subject := vars[subject]
		rootSubject := r.FormValue(rootSubject)
		rootSubjectKeyId := r.FormValue(rootSubjectKeyId)
		getListCertificates(restCtx, fmt.Sprintf("custom/%s/all_subject_x509_certs/%s", storeName, subject), rootSubject, rootSubjectKeyId)
	}
}

func getProposedX509RootCertHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		vars := restCtx.Variables()
		subject := vars[subject]
		subjectKeyId := vars[subjectKeyId]

		res, height, err := restCtx.QueryStore(keeper.ProposedCertificateId(subject, subjectKeyId), storeName)
		if err != nil || res == nil {
			restCtx.WriteErrorResponse(http.StatusNotFound, types.ErrProposedCertificateDoesNotExist(subject, subjectKeyId).Error())
			return
		}

		var proposedCertificate types.ProposedCertificate
		restCtx.Codec().MustUnmarshalBinaryBare(res, &proposedCertificate)

		restCtx.EncodeAndRespondWithHeight(proposedCertificate, height)
	}
}

func getX509CertHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		vars := restCtx.Variables()
		subject := vars[subject]
		subjectKeyId := vars[subjectKeyId]

		res, height, err := restCtx.QueryStore(keeper.CertificateId(subject, subjectKeyId), storeName)
		if err != nil || res == nil {
			restCtx.WriteErrorResponse(http.StatusNotFound, types.ErrCertificateDoesNotExist(subject, subjectKeyId).Error())
			return
		}

		var certificate types.Certificates
		cliCtx.Codec.MustUnmarshalBinaryBare(res, &certificate)

		restCtx.EncodeAndRespondWithHeight(certificate, height)
	}
}

func getListCertificates(restCtx rest.RestContext, path string, rootSubject string, rootSubjectKeyId string) {
	paginationParams, err := restCtx.ParsePaginationParams()
	if err != nil {
		return
	}

	params := types.NewListCertificatesQueryParams(paginationParams, rootSubject, rootSubjectKeyId)
	restCtx.QueryList(path, params)
}
