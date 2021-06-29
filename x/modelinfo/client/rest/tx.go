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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelinfo/internal/types"
)

type AddModelInfoRequest struct {
	Model   types.Model       `json:"model"`
	BaseReq restTypes.BaseReq `json:"base_req"`
}

type UpdateModelInfoRequest struct {
	Model   types.Model       `json:"model"`
	BaseReq restTypes.BaseReq `json:"base_req"`
}

func addModelHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		var req AddModelInfoRequest
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
			VID:                                      req.Model.VID,
			PID:                                      req.Model.PID,
			CID:                                      req.Model.CID,
			Name:                                     req.Model.Name,
			Description:                              req.Model.Description,
			SKU:                                      req.Model.SKU,
			SoftwareVersion:                          req.Model.SoftwareVersion,
			SoftwareVersionString:                    req.Model.SoftwareVersionString,
			HardwareVersion:                          req.Model.HardwareVersion,
			HardwareVersionString:                    req.Model.HardwareVersionString,
			CDVersionNumber:                          req.Model.CDVersionNumber,
			FirmwareDigests:                          req.Model.FirmwareDigests,
			Revoked:                                  req.Model.Revoked,
			OtaURL:                                   req.Model.OtaURL,
			OtaChecksum:                              req.Model.OtaChecksum,
			OtaChecksumType:                          req.Model.OtaChecksumType,
			OtaBlob:                                  req.Model.OtaBlob,
			CommissioningCustomFlow:                  req.Model.CommissioningCustomFlow,
			CommissioningCustomFlowURL:               req.Model.CommissioningCustomFlowURL,
			CommissioningModeInitialStepsHint:        req.Model.CommissioningModeInitialStepsHint,
			CommissioningModeInitialStepsInstruction: req.Model.CommissioningModeInitialStepsInstruction,
			CommissioningModeSecondaryStepsHint:      req.Model.CommissioningModeSecondaryStepsHint,
			CommissioningModeSecondaryStepsInstruction: req.Model.CommissioningModeSecondaryStepsInstruction,
			ReleaseNotesURL: req.Model.ReleaseNotesURL,
			UserManualURL:   req.Model.UserManualURL,
			SupportURL:      req.Model.SupportURL,
			ProductURL:      req.Model.ProductURL,
			ChipBlob:        req.Model.ChipBlob,
			VendorBlob:      req.Model.VendorBlob,
		}

		msg := types.NewMsgAddModelInfo(model, restCtx.Signer())

		restCtx.HandleWriteRequest(msg)
	}
}

func updateModelHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		var req UpdateModelInfoRequest
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
			VID:                        req.Model.VID,
			PID:                        req.Model.PID,
			CID:                        req.Model.CID,
			Description:                req.Model.Description,
			CDVersionNumber:            req.Model.CDVersionNumber,
			Revoked:                    req.Model.Revoked,
			OtaURL:                     req.Model.OtaURL,
			OtaChecksum:                req.Model.OtaChecksum,
			OtaChecksumType:            req.Model.OtaChecksumType,
			OtaBlob:                    req.Model.OtaBlob,
			CommissioningCustomFlowURL: req.Model.CommissioningCustomFlowURL,
			ReleaseNotesURL:            req.Model.ReleaseNotesURL,
			UserManualURL:              req.Model.UserManualURL,
			SupportURL:                 req.Model.SupportURL,
			ProductURL:                 req.Model.ProductURL,
			ChipBlob:                   req.Model.ChipBlob,
			VendorBlob:                 req.Model.VendorBlob,
		}

		msg := types.NewMsgUpdateModelInfo(model, restCtx.Signer())
		restCtx.HandleWriteRequest(msg)
	}
}
