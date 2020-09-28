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
	BaseReq                  restTypes.BaseReq `json:"base_req"`
	VID                      uint16            `json:"vid"`
	PID                      uint16            `json:"pid"`
	CID                      uint16            `json:"cid,omitempty"`
	Version                  string            `json:"version,omitempty"`
	Name                     string            `json:"name"`
	Description              string            `json:"description"`
	SKU                      string            `json:"sku"`
	HardwareVersion          string            `json:"hardware_version"`
	FirmwareVersion          string            `json:"firmware_version"`
	OtaURL                   string            `json:"ota_url,omitempty"`
	OtaChecksum              string            `json:"ota_checksum,omitempty"`
	OtaChecksumType          string            `json:"ota_checksum_type,omitempty"`
	Custom                   string            `json:"custom,omitempty"`
	TisOrTrpTestingCompleted bool              `json:"tis_or_trp_testing_completed"`
}

//nolint:maligned
type UpdateModelInfoRequest struct {
	BaseReq                  restTypes.BaseReq `json:"base_req"`
	VID                      uint16            `json:"vid"`
	PID                      uint16            `json:"pid"`
	CID                      uint16            `json:"cid,omitempty"`
	Description              string            `json:"description,omitempty"`
	OtaURL                   string            `json:"ota_url,omitempty"`
	Custom                   string            `json:"custom,omitempty"`
	TisOrTrpTestingCompleted bool              `json:"tis_or_trp_testing_completed"`
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

		msg := types.NewMsgAddModelInfo(req.VID, req.PID, req.CID, req.Version,
			req.Name, req.Description, req.SKU, req.HardwareVersion,
			req.FirmwareVersion, req.OtaURL, req.OtaChecksum, req.OtaChecksumType,
			req.Custom, req.TisOrTrpTestingCompleted, restCtx.Signer())

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

		msg := types.NewMsgUpdateModelInfo(req.VID, req.PID, req.CID, req.Description,
			req.OtaURL, req.Custom, req.TisOrTrpTestingCompleted, restCtx.Signer())

		restCtx.HandleWriteRequest(msg)
	}
}
