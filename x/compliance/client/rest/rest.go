package rest

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"

	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/gorilla/mux"
)

const (
	vid = "vid"
	pid = "pid"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc(fmt.Sprintf("/%s/{%s}/{%s}", storeName, vid, pid), getComplianceInfoHandler(cliCtx, storeName), ).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s", storeName), getComplianceInfosHandler(cliCtx, storeName), ).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/%s", storeName, types.Certified), certifyModelHandler(cliCtx)).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/%s/%s/{%s}/{%s}", storeName, types.Certified, vid, pid), getCertifiedModelHandler(cliCtx, storeName), ).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/%s", storeName, types.Certified), getCertifiedModelsHandler(cliCtx, storeName), ).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/%s", storeName, types.Revoked), revokeModelHandler(cliCtx)).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/%s/{%s}/{%s}/{%s}", storeName, types.Revoked, vid, pid), getRevokedModelHandler(cliCtx, storeName), ).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/%s", storeName, types.Revoked), getRevokedModelsHandler(cliCtx, storeName), ).Methods("GET")
}
