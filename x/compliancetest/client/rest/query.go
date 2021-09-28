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
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/conversions"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/rest"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/internal/types"
)

func getTestingResultHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
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

		softwareVersion, err_ := conversions.ParseUInt32FromString("softwareVersion", vars[softwareVersion])
		if err_ != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest, err_.Error())
			return
		}

		res, height, err := restCtx.QueryStore(types.GetTestingResultsKey(vid, pid, softwareVersion), storeName)
		if err != nil || res == nil {
			restCtx.WriteErrorResponse(http.StatusNotFound, types.ErrTestingResultDoesNotExist(vid, pid, softwareVersion).Error())

			return
		}

		var testingResult types.TestingResults

		cliCtx.Codec.MustUnmarshalBinaryBare(res, &testingResult)

		restCtx.EncodeAndRespondWithHeight(testingResult, height)
	}
}
