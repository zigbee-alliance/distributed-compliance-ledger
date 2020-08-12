package rest

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/gorilla/mux"
)

const (
	address = "address"
)

// RegisterRoutes - Central function to define routes that get registered by the main application.
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc(
		"/auth/accounts/proposed",
		proposeAddAccountHandler(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		fmt.Sprintf("/auth/accounts/proposed/{%s}", address),
		approveAddAccountHandler(cliCtx),
	).Methods("PATCH")
	r.HandleFunc(
		"/auth/accounts/proposed",
		proposedAccountsHandler(cliCtx, storeName),
	).Methods("GET")
	r.HandleFunc(
		"/auth/accounts",
		accountsHandler(cliCtx, storeName),
	).Methods("GET")
	r.HandleFunc(
		fmt.Sprintf("/auth/accounts/{%s}", address),
		accountHandler(cliCtx, storeName),
	).Methods("GET")
	r.HandleFunc(
		"/auth/accounts/proposed/revoked",
		proposeRevokeAccountHandler(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		fmt.Sprintf("/auth/accounts/proposed/revoked/{%s}", address),
		approveRevokeAccountHandler(cliCtx),
	).Methods("PATCH")
}
