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
	"time"

	"github.com/cosmos/cosmos-sdk/client/context"
	restTypes "github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/conversions"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/rest"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/internal/types"
)

type CertifyModelRequest struct {
	BaseReq           restTypes.BaseReq `json:"base_req"`
	CertificationDate time.Time         `json:"certification_date"` // rfc3339 encoded date
	Reason            string            `json:"reason,omitempty"`
}

// nolint:dupl
func certifyModelHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		vars := restCtx.Variables()

		vid, err_ := conversions.ParseVID(vars[vid])
		if err_ != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest, err_.Error())

			return
		}

		pid, err_ := conversions.ParsePID(vars[pid])
		if err_ != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest, err_.Error())

			return
		}

		softwareVersion, err_ := conversions.ParseUInt32FromString(softwareVersion, vars[softwareVersion])
		if err_ != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest, err_.Error())

			return
		}

		softwareVersionString := vars[softwareVersionString]

		certificationType := types.CertificationType(vars[certificationType])

		var req CertifyModelRequest
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

		msg := types.NewMsgCertifyModel(vid, pid, softwareVersion, softwareVersionString, req.CertificationDate,
			certificationType, req.Reason, restCtx.Signer())

		restCtx.HandleWriteRequest(msg)
	}
}

type RevokeModelRequest struct {
	BaseReq        restTypes.BaseReq `json:"base_req"`
	RevocationDate time.Time         `json:"revocation_date"` // rfc3339 encoded date
	Reason         string            `json:"reason,omitempty"`
}

// nolint:dupl
func revokeModelHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		vars := restCtx.Variables()

		vid, err_ := conversions.ParseVID(vars[vid])
		if err_ != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest, err_.Error())

			return
		}

		pid, err_ := conversions.ParsePID(vars[pid])
		if err_ != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest, err_.Error())

			return
		}

		softwareVersion, err_ := conversions.ParseUInt32FromString(softwareVersion, vars[softwareVersion])
		if err_ != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest, err_.Error())

			return
		}

		certificationType := types.CertificationType(vars[certificationType])

		var req RevokeModelRequest
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

		msg := types.NewMsgRevokeModel(vid, pid, softwareVersion, req.RevocationDate,
			certificationType, req.Reason, restCtx.Signer())

		restCtx.HandleWriteRequest(msg)
	}
}
