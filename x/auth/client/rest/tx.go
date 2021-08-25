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
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	restTypes "github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/rest"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth/internal/types"
)

type ProposeAddAccountRequest struct {
	BaseReq  restTypes.BaseReq  `json:"base_req"`
	Address  sdk.AccAddress     `json:"address"`
	Pubkey   string             `json:"pubkey"`
	Roles    types.AccountRoles `json:"roles"`
	VendorId uint16             `json:"vendorId"`
}

func proposeAddAccountHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		var req ProposeAddAccountRequest
		if !restCtx.ReadRESTReq(&req) {
			return
		}

		restCtx, err := restCtx.WithBaseRequest(req.BaseReq)
		if err != nil {
			return
		}

		restCtx, err = restCtx.WithSigner()
		if err != nil {
			return
		}

		msg := types.NewMsgProposeAddAccount(req.Address, req.Pubkey, req.Roles, req.VendorId, restCtx.Signer())

		restCtx.HandleWriteRequest(msg)
	}
}

func approveAddAccountHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		vars := restCtx.Variables()

		var req rest.BasicReq
		if !restCtx.ReadRESTReq(&req) {
			return
		}

		restCtx, err := restCtx.WithBaseRequest(req.BaseReq)
		if err != nil {
			return
		}

		restCtx, err = restCtx.WithSigner()
		if err != nil {
			return
		}

		address, err := sdk.AccAddressFromBech32(vars[address])
		if err != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest,
				fmt.Sprintf("Request Parsing Error: %v. valid address must be cpecified", err))

			return
		}

		msg := types.NewMsgApproveAddAccount(address, restCtx.Signer())

		restCtx.HandleWriteRequest(msg)
	}
}

type ProposeRevokeAccountRequest struct {
	BaseReq restTypes.BaseReq `json:"base_req"`
	Address sdk.AccAddress    `json:"address"`
}

func proposeRevokeAccountHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		var req ProposeRevokeAccountRequest
		if !restCtx.ReadRESTReq(&req) {
			return
		}

		restCtx, err := restCtx.WithBaseRequest(req.BaseReq)
		if err != nil {
			return
		}

		restCtx, err = restCtx.WithSigner()
		if err != nil {
			return
		}

		msg := types.NewMsgProposeRevokeAccount(req.Address, restCtx.Signer())

		restCtx.HandleWriteRequest(msg)
	}
}

func approveRevokeAccountHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		vars := restCtx.Variables()

		var req rest.BasicReq
		if !restCtx.ReadRESTReq(&req) {
			return
		}

		restCtx, err := restCtx.WithBaseRequest(req.BaseReq)
		if err != nil {
			return
		}

		restCtx, err = restCtx.WithSigner()
		if err != nil {
			return
		}

		address, err := sdk.AccAddressFromBech32(vars[address])
		if err != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest,
				fmt.Sprintf("Request Parsing Error: %v. valid address must be cpecified", err))

			return
		}

		msg := types.NewMsgApproveRevokeAccount(address, restCtx.Signer())

		restCtx.HandleWriteRequest(msg)
	}
}
