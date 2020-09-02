package rest

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

const (
	vid = "vid"
	pid = "pid"
)

// RegisterRoutes - Central function to define routes that get registered by the main application.
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc(
		fmt.Sprintf("/%s/models", storeName),
		addModelHandler(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		fmt.Sprintf("/%s/models", storeName),
		updateModelHandler(cliCtx),
	).Methods("PUT")
	r.HandleFunc(
		fmt.Sprintf("/%s/models", storeName),
		getModelsHandler(cliCtx, storeName),
	).Methods("GET")
	r.HandleFunc(
		fmt.Sprintf("/%s/models/{%s}", storeName, vid),
		getVendorModelsHandler(cliCtx, storeName),
	).Methods("GET")
	r.HandleFunc(
		fmt.Sprintf("/%s/models/{%s}/{%s}", storeName, vid, pid),
		getModelHandler(cliCtx, storeName),
	).Methods("GET")
	r.HandleFunc(
		fmt.Sprintf("/%s/vendors", storeName),
		getVendorsHandler(cliCtx, storeName),
	).Methods("GET")
}
