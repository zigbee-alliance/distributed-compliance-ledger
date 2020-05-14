package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

const (
	node   = "node"
	height = "height"
)

func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/blocks", BlocksHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/status", NodeStatusHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/validator-set", ValidatorSetRequestHandlerFn(cliCtx)).Methods("GET")
}
