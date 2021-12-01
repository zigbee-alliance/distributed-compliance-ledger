// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	r.HandleFunc(
		"/auth/accounts/proposed/revoked",
		proposedAccountsToRevokeHandler(cliCtx, storeName),
	).Methods("GET")
}
