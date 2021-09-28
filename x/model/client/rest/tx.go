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
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	restTypes "github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/rest"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/internal/types"
)

type AddModelRequest struct {
	types.Model
	BaseReq restTypes.BaseReq `json:"base_req"`
}

type UpdateModelRequest struct {
	types.Model
	BaseReq restTypes.BaseReq `json:"base_req"`
}

func addModelHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		var req AddModelRequest
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

		model := types.Model{
			VID:                                      req.VID,
			PID:                                      req.PID,
			DeviceTypeID:                             req.DeviceTypeID,
			ProductName:                              req.ProductName,
			ProductLabel:                             req.ProductLabel,
			PartNumber:                               req.PartNumber,
			CommissioningCustomFlow:                  req.CommissioningCustomFlow,
			CommissioningCustomFlowURL:               req.CommissioningCustomFlowURL,
			CommissioningModeInitialStepsHint:        req.CommissioningModeInitialStepsHint,
			CommissioningModeInitialStepsInstruction: req.CommissioningModeInitialStepsInstruction,
			CommissioningModeSecondaryStepsHint:      req.CommissioningModeSecondaryStepsHint,
			CommissioningModeSecondaryStepsInstruction: req.CommissioningModeSecondaryStepsInstruction,
			UserManualURL: req.UserManualURL,
			SupportURL:    req.SupportURL,
			ProductURL:    req.ProductURL,
		}

		msg := types.NewMsgAddModel(model, restCtx.Signer())

		restCtx.HandleWriteRequest(msg)
	}
}

func updateModelHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		var req UpdateModelRequest
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

		model := types.Model{
			VID:                        req.VID,
			PID:                        req.PID,
			DeviceTypeID:               req.DeviceTypeID,
			ProductName:                req.ProductName,
			ProductLabel:               req.ProductLabel,
			CommissioningCustomFlowURL: req.CommissioningCustomFlowURL,
			UserManualURL:              req.UserManualURL,
			SupportURL:                 req.SupportURL,
			ProductURL:                 req.ProductURL,
		}

		msg := types.NewMsgUpdateModel(model, restCtx.Signer())
		restCtx.HandleWriteRequest(msg)
	}
}
