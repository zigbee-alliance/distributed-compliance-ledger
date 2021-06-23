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

//nolint:maligned
type AddModelInfoRequest struct {
	types.Model
	BaseReq restTypes.BaseReq `json:"base_req"`
}

//nolint:maligned
type UpdateModelInfoRequest struct {
	types.Model
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
			VID:                                      req.VID,
			PID:                                      req.PID,
			CID:                                      req.CID,
			Name:                                     req.Name,
			Description:                              req.Description,
			SKU:                                      req.SKU,
			SoftwareVersion:                          req.SoftwareVersion,
			SoftwareVersionString:                    req.SoftwareVersionString,
			HardwareVersion:                          req.HardwareVersion,
			HardwareVersionString:                    req.HardwareVersionString,
			CDVersionNumber:                          req.CDVersionNumber,
			FirmwareDigests:                          req.FirmwareDigests,
			Revoked:                                  req.Revoked,
			OtaURL:                                   req.OtaURL,
			OtaChecksum:                              req.OtaChecksum,
			OtaChecksumType:                          req.OtaChecksumType,
			OtaBlob:                                  req.OtaBlob,
			CommissioningCustomFlow:                  req.CommissioningCustomFlow,
			CommissioningCustomFlowUrl:               req.CommissioningCustomFlowUrl,
			CommissioningModeInitialStepsHint:        req.CommissioningModeInitialStepsHint,
			CommissioningModeInitialStepsInstruction: req.CommissioningModeInitialStepsInstruction,
			CommissioningModeSecondaryStepsHint:      req.CommissioningModeSecondaryStepsHint,
			CommissioningModeSecondaryStepsInstruction: req.CommissioningModeSecondaryStepsInstruction,
			ReleaseNotesUrl: req.ReleaseNotesUrl,
			UserManualUrl:   req.UserManualUrl,
			SupportUrl:      req.SupportUrl,
			ProductURL:      req.ProductURL,
			ChipBlob:        req.ChipBlob,
			VendorBlob:      req.VendorBlob,
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
			VID:                        req.VID,
			PID:                        req.PID,
			CID:                        req.CID,
			Description:                req.Description,
			CDVersionNumber:            req.CDVersionNumber,
			Revoked:                    req.Revoked,
			OtaURL:                     req.OtaURL,
			OtaChecksum:                req.OtaChecksum,
			OtaChecksumType:            req.OtaChecksumType,
			OtaBlob:                    req.OtaBlob,
			CommissioningCustomFlowUrl: req.CommissioningCustomFlowUrl,
			ReleaseNotesUrl:            req.ReleaseNotesUrl,
			UserManualUrl:              req.UserManualUrl,
			SupportUrl:                 req.SupportUrl,
			ProductURL:                 req.ProductURL,
			ChipBlob:                   req.ChipBlob,
			VendorBlob:                 req.VendorBlob,
		}

		msg := types.NewMsgUpdateModelInfo(model, restCtx.Signer())
		restCtx.HandleWriteRequest(msg)
	}
}
