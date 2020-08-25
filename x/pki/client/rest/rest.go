package rest

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/gorilla/mux"
)

const (
	subject          = "subject"
	subjectKeyID     = "subject_key_id"
	rootSubject      = "root_subject"
	rootSubjectKeyID = "root_subject_key_id"
)

// RegisterRoutes - Central function to define routes that get registered by the main application.
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc(
		fmt.Sprintf("/%s/certs/proposed/root", storeName),
		proposeAddX509RootCertHandler(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		fmt.Sprintf("/%s/certs/proposed/root/{%s}/{%s}", storeName, subject, subjectKeyID),
		approveAddX509RootCertHandler(cliCtx),
	).Methods("PATCH")
	r.HandleFunc(
		fmt.Sprintf("/%s/certs", storeName),
		addX509CertHandler(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		fmt.Sprintf("/%s/certs/revoked", storeName),
		revokeX509CertHandler(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		fmt.Sprintf("/%s/certs/root", storeName),
		getAllX509RootCertsHandler(cliCtx, storeName),
	).Methods("GET")
	r.HandleFunc(
		fmt.Sprintf("/%s/certs/proposed/root/{%s}/{%s}", storeName, subject, subjectKeyID),
		getProposedX509RootCertHandler(cliCtx, storeName),
	).Methods("GET")
	r.HandleFunc(
		fmt.Sprintf("/%s/certs/proposed/root", storeName),
		getAllProposedX509RootCertsHandler(cliCtx, storeName),
	).Methods("GET")
	r.HandleFunc(
		fmt.Sprintf("/%s/certs/{%s}/{%s}", storeName, subject, subjectKeyID),
		getX509CertHandler(cliCtx, storeName),
	).Methods("GET")
	r.HandleFunc(
		fmt.Sprintf("/%s/certs/chain/{%s}/{%s}", storeName, subject, subjectKeyID),
		getX509CertChainHandler(cliCtx, storeName),
	).Methods("GET")
	r.HandleFunc(
		fmt.Sprintf("/%s/certs/{%s}", storeName, subject),
		getAllSubjectX509CertsHandler(cliCtx, storeName),
	).Methods("GET")
	r.HandleFunc(
		fmt.Sprintf("/%s/certs", storeName),
		getAllX509CertsHandler(cliCtx, storeName),
	).Methods("GET")
	r.HandleFunc(
		fmt.Sprintf("/%s/certs/proposed/revoked/root/{%s}/{%s}", storeName, subject, subjectKeyID),
		getProposedX509RootCertToRevokeHandler(cliCtx, storeName),
	).Methods("GET")
	r.HandleFunc(
		fmt.Sprintf("/%s/certs/proposed/revoked/root", storeName),
		getAllProposedX509RootCertsToRevokeHandler(cliCtx, storeName),
	).Methods("GET")
	r.HandleFunc(
		fmt.Sprintf("/%s/certs/revoked/{%s}/{%s}", storeName, subject, subjectKeyID),
		getRevokedX509CertHandler(cliCtx, storeName),
	).Methods("GET")
	r.HandleFunc(
		fmt.Sprintf("/%s/certs/revoked/root", storeName),
		getAllRevokedX509RootCertsHandler(cliCtx, storeName),
	).Methods("GET")
	r.HandleFunc(
		fmt.Sprintf("/%s/certs/revoked", storeName),
		getAllRevokedX509CertsHandler(cliCtx, storeName),
	).Methods("GET")
}
