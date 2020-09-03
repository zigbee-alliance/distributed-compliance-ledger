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
	validatorAddr = "validator_addr"
	state         = "state"
)

// RegisterRoutes - Central function to define routes that get registered by the main application.
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc(
		"/validators",
		createValidatorHandlerFn(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		"/validators",
		getValidatorsHandlerFn(cliCtx, storeName),
	).Methods("GET")
	r.HandleFunc(
		fmt.Sprintf("/validators/{%s}", validatorAddr),
		getValidatorHandlerFn(cliCtx, storeName),
	).Methods("GET")
}
