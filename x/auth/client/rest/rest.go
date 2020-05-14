package rest

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/gorilla/mux"
)

const (
	addrKey = "address"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc("/auth/accounts",
		accountsHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/auth/account/{%s}", addrKey),
		accountHandler(cliCtx, storeName)).Methods("GET")
}
