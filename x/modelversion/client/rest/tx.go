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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelversion/internal/types"
)

type AddModelVersionRequest struct {
	types.ModelVersion
	BaseReq restTypes.BaseReq `json:"base_req"`
}

type UpdateModelVersionRequest struct {
	types.ModelVersion
	BaseReq restTypes.BaseReq `json:"base_req"`
}

func addModelVersionHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		var req AddModelVersionRequest
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

		modelVersion := types.ModelVersion{
			VID:                          req.VID,
			PID:                          req.PID,
			SoftwareVersion:              req.SoftwareVersion,
			SoftwareVersionString:        req.SoftwareVersionString,
			CDVersionNumber:              req.CDVersionNumber,
			FirmwareDigests:              req.FirmwareDigests,
			SoftwareVersionValid:         req.SoftwareVersionValid,
			OtaURL:                       req.OtaURL,
			OtaFileSize:                  req.OtaFileSize,
			OtaChecksum:                  req.OtaChecksum,
			OtaChecksumType:              req.OtaChecksumType,
			MinApplicableSoftwareVersion: req.MinApplicableSoftwareVersion,
			MaxApplicableSoftwareVersion: req.MaxApplicableSoftwareVersion,
			ReleaseNotesURL:              req.ReleaseNotesURL,
		}

		msg := types.NewMsgAddModelVersion(modelVersion, restCtx.Signer())
		restCtx.HandleWriteRequest(msg)
	}
}

func updateModelVersionHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		var req UpdateModelVersionRequest
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

		modelVersion := types.ModelVersion{
			VID:                          req.ModelVersion.VID,
			PID:                          req.ModelVersion.PID,
			SoftwareVersion:              req.ModelVersion.SoftwareVersion,
			SoftwareVersionValid:         req.ModelVersion.SoftwareVersionValid,
			OtaURL:                       req.ModelVersion.OtaURL,
			MinApplicableSoftwareVersion: req.ModelVersion.MinApplicableSoftwareVersion,
			MaxApplicableSoftwareVersion: req.ModelVersion.MaxApplicableSoftwareVersion,
			ReleaseNotesURL:              req.ModelVersion.ReleaseNotesURL,
		}

		msg := types.NewMsgUpdateModelVersion(modelVersion, restCtx.Signer())
		restCtx.HandleWriteRequest(msg)
	}
}
