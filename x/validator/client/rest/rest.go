package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers validator-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc("/validator/validators/{validatorAddr}", createValidatorHandlerFn(cliCtx), ).Methods("POST")
	r.HandleFunc("/validator/validators", validatorsHandlerFn(cliCtx, storeName)).Methods("GET")
	r.HandleFunc("/validator/validators/{validatorAddr}", validatorHandlerFn(cliCtx, storeName)).Methods("GET")
}
