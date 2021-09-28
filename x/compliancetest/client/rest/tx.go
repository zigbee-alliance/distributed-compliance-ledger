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
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/rest"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/internal/types"
)

type TestingResultRequest struct {
	BaseReq               restTypes.BaseReq `json:"base_req"`
	VID                   uint16            `json:"vid"`
	PID                   uint16            `json:"pid"`
	SoftwareVersion       uint32            `json:"softwareVersion"`
	SoftwareVersionString string            `json:"softwareVersionString"`
	TestResult            string            `json:"test_result"`
	TestDate              time.Time         `json:"test_date"` // rfc3339 encoded date
}

func addTestingResultHandler(cliCtx context.CLIContext) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		var req TestingResultRequest
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

		msg := types.NewMsgAddTestingResult(req.VID, req.PID,
			req.SoftwareVersion, req.SoftwareVersionString, req.TestResult, req.TestDate, restCtx.Signer())

		restCtx.HandleWriteRequest(msg)
	}
}
