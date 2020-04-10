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

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc(fmt.Sprintf("/%s/model_info", storeName), modelInfoHeadersHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/model_info/{%s}/{%s}", storeName, vid, pid), modelInfoHandler(cliCtx,
		storeName)).Methods("GET")
}
