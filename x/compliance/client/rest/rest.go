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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/internal/types"
)

const (
	VID = "vid"
	PID = "pid"
	SV  = "softwareVersion"
	HV  = "hardwareVersion"
	CT  = "certification_type"
)

// RegisterRoutes - Central function to define routes that get registered by the main application.
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc(
		fmt.Sprintf("/%s/{%s}/{%s}/{%s}/{%s}/{%s}", storeName,
			VID, PID, SV, HV, CT),
		getComplianceInfoHandler(cliCtx, storeName),
	).Methods("GET")
	r.HandleFunc(
		fmt.Sprintf("/%s", storeName),
		getComplianceInfosHandler(cliCtx, storeName),
	).Methods("GET")
	r.HandleFunc(
		fmt.Sprintf("/%s/%s/{%s}/{%s}/{%s}/{%s}/{%s}", storeName, types.Certified,
			VID, PID, SV, HV, CT),
		certifyModelHandler(cliCtx),
	).Methods("PUT")
	r.HandleFunc(
		fmt.Sprintf("/%s/%s/{%s}/{%s}/{%s}/{%s}/{%s}", storeName, types.Certified,
			VID, PID, SV, HV, CT),
		getCertifiedModelHandler(cliCtx, storeName),
	).Methods("GET")
	r.HandleFunc(
		fmt.Sprintf("/%s/%s", storeName, types.Certified),
		getCertifiedModelsHandler(cliCtx, storeName),
	).Methods("GET")
	r.HandleFunc(
		fmt.Sprintf("/%s/%s/{%s}/{%s}/{%s}/{%s}/{%s}", storeName, types.Revoked,
			VID, PID, SV, HV, CT),
		revokeModelHandler(cliCtx),
	).Methods("PUT")
	r.HandleFunc(
		fmt.Sprintf("/%s/%s/{%s}/{%s}/{%s}/{%s}/{%s}", storeName, types.Revoked,
			VID, PID, SV, HV, CT),
		getRevokedModelHandler(cliCtx, storeName),
	).Methods("GET")
	r.HandleFunc(
		fmt.Sprintf("/%s/%s", storeName, types.Revoked),
		getRevokedModelsHandler(cliCtx, storeName),
	).Methods("GET")
}
