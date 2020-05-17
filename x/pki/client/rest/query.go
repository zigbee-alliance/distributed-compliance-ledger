package rest

// nolint:goimports
import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/rest"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"net/http"
)

func getAllX509RootCertsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)
		getListCertificates(restCtx,
			fmt.Sprintf("custom/%s/all_x509_root_certs", storeName), "", "")
	}
}

func getAllProposedX509RootCertsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)
		getListCertificates(restCtx,
			fmt.Sprintf("custom/%s/all_proposed_x509_root_certs", storeName), "", "")
	}
}

func getAllX509CertsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)
		rootSubject := r.FormValue(rootSubject)
		rootSubjectKeyID := r.FormValue(rootSubjectKeyID)
		getListCertificates(restCtx, fmt.Sprintf("custom/%s/all_x509_certs", storeName), rootSubject, rootSubjectKeyID)
	}
}

func getAllSubjectX509CertsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)
		vars := restCtx.Variables()
		subject := vars[subject]
		rootSubject := r.FormValue(rootSubject)
		rootSubjectKeyID := r.FormValue(rootSubjectKeyID)
		getListCertificates(restCtx, fmt.Sprintf("custom/%s/all_subject_x509_certs/%s",
			storeName, subject), rootSubject, rootSubjectKeyID)
	}
}

func getProposedX509RootCertHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		vars := restCtx.Variables()
		subject := vars[subject]
		subjectKeyID := vars[subjectKeyID]

		res, height, err := restCtx.QueryStore(types.GetProposedCertificateKey(subject, subjectKeyID), storeName)
		if err != nil || res == nil {
			restCtx.WriteErrorResponse(http.StatusNotFound,
				types.ErrProposedCertificateDoesNotExist(subject, subjectKeyID).Error())
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
		subjectKeyID := vars[subjectKeyID]

		res, height, err := restCtx.QueryStore(types.GetApprovedCertificateKey(subject, subjectKeyID), storeName)
		if err != nil || res == nil {
			restCtx.WriteErrorResponse(http.StatusNotFound,
				types.ErrCertificateDoesNotExist(subject, subjectKeyID).Error())
			return
		}

		var certificate types.Certificates

		cliCtx.Codec.MustUnmarshalBinaryBare(res, &certificate)

		restCtx.EncodeAndRespondWithHeight(certificate, height)
	}
}

func getListCertificates(restCtx rest.RestContext, path string, rootSubject string, rootSubjectKeyID string) {
	paginationParams, err := restCtx.ParsePaginationParams()
	if err != nil {
		return
	}

	params := types.NewListCertificatesQueryParams(paginationParams, rootSubject, rootSubjectKeyID)
	restCtx.QueryList(path, params)
}
