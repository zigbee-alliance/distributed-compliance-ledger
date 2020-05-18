package rest

//nolint:goimports
import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"

	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/gorilla/mux"
)

const (
	vid               = "vid"
	pid               = "pid"
	certificationType = "certification_type"
)

// RegisterRoutes - Central function to define routes that get registered by the main application.
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc(
		fmt.Sprintf("/%s/{%s}/{%s}/{%s}", storeName, vid, pid, certificationType),
		getComplianceInfoHandler(cliCtx, storeName),
	).Methods("GET")
	r.HandleFunc(
		fmt.Sprintf("/%s", storeName),
		getComplianceInfosHandler(cliCtx, storeName),
	).Methods("GET")
	r.HandleFunc(
		fmt.Sprintf("/%s/%s/{%s}/{%s}/{%s}", storeName, types.Certified, vid, pid, certificationType),
		certifyModelHandler(cliCtx),
	).Methods("PUT")
	r.HandleFunc(
		fmt.Sprintf("/%s/%s/{%s}/{%s}/{%s}", storeName, types.Certified, vid, pid, certificationType),
		getCertifiedModelHandler(cliCtx, storeName),
	).Methods("GET")
	r.HandleFunc(
		fmt.Sprintf("/%s/%s", storeName, types.Certified),
		getCertifiedModelsHandler(cliCtx, storeName),
	).Methods("GET")
	r.HandleFunc(
		fmt.Sprintf("/%s/%s/{%s}/{%s}/{%s}", storeName, types.Revoked, vid, pid, certificationType),
		revokeModelHandler(cliCtx),
	).Methods("PUT")
	r.HandleFunc(
		fmt.Sprintf("/%s/{%s}/{%s}/{%s}/{%s}", storeName, types.Revoked, vid, pid, certificationType),
		getRevokedModelHandler(cliCtx, storeName),
	).Methods("GET")
	r.HandleFunc(
		fmt.Sprintf("/%s/%s", storeName, types.Revoked),
		getRevokedModelsHandler(cliCtx, storeName),
	).Methods("GET")
}
