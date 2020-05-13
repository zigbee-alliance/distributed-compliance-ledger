package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/rpc"

	"github.com/gorilla/mux"
)

const (
	node = "node"
)


func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/blocks", BlocksHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/node-status", NodeStatusHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/validator-set", rpc.LatestValidatorSetRequestHandlerFn(cliCtx)).Methods("GET")
}
