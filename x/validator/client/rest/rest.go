package rest

import (
	"fmt"
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

const (
	validatorAddr = "validator_addr"
	state         = "state"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc(fmt.Sprintf("/validators"), createValidatorHandlerFn(cliCtx), ).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/validators"), getValidatorsHandlerFn(cliCtx, storeName), ).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/validators/{%s}", validatorAddr), getValidatorHandlerFn(cliCtx, storeName), ).Methods("GET")
}
