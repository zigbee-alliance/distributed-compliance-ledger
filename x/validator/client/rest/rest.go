package rest

import (
	"fmt"
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

const (
	validatorAddr = "validator_addr"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc(fmt.Sprintf("/%s/validators/{%s}", storeName, validatorAddr), createValidatorHandlerFn(cliCtx), ).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/validators", storeName), getValidatorsHandlerFn(cliCtx, storeName), ).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/validators/{%s}", storeName, validatorAddr), getValidatorHandlerFn(cliCtx, storeName), ).Methods("GET")
}
