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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/internal/types"
)

type AddVendorInfoRequest struct {
	VendorInfo types.VendorInfo  `json:"vendor"`
	BaseReq    restTypes.BaseReq `json:"base_req"`
}

type UpdateModelRequest struct {
	VendorInfo types.VendorInfo  `json:"vendor"`
	BaseReq    restTypes.BaseReq `json:"base_req"`
}

func addVendorHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		var req AddVendorInfoRequest
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

		vendorInfo := types.VendorInfo{
			VendorId:             req.VendorInfo.VendorId,
			VendorName:           req.VendorInfo.VendorName,
			CompanyLegalName:     req.VendorInfo.CompanyLegalName,
			CompanyPreferredName: req.VendorInfo.CompanyPreferredName,
			VendorLandingPageUrl: req.VendorInfo.VendorLandingPageUrl,
		}

		msg := types.NewMsgAddVendorInfo(vendorInfo, restCtx.Signer())

		restCtx.HandleWriteRequest(msg)
	}
}

func updateVendorHandler(cliCtx context.CLIContext) http.HandlerFunc {
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

		vendorInfo := types.VendorInfo{
			VendorId:             req.VendorInfo.VendorId,
			VendorName:           req.VendorInfo.VendorName,
			CompanyLegalName:     req.VendorInfo.CompanyLegalName,
			CompanyPreferredName: req.VendorInfo.CompanyPreferredName,
			VendorLandingPageUrl: req.VendorInfo.VendorLandingPageUrl,
		}

		msg := types.NewMsgUpdateVendorInfo(vendorInfo, restCtx.Signer())
		restCtx.HandleWriteRequest(msg)
	}
}
