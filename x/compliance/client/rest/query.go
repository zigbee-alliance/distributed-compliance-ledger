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

	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/conversions"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/rest"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"
)

func getComplianceInfoHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		getComplianceInfo(cliCtx, w, r, storeName)
	}
}

func getComplianceInfosHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		getAllComplianceInfo(cliCtx, w, r, fmt.Sprintf("custom/%s/all_compliance_info_records", storeName))
	}
}

func getCertifiedModelsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		getAllComplianceInfo(cliCtx, w, r, fmt.Sprintf("custom/%s/all_certified_models", storeName))
	}
}

func getCertifiedModelHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		getComplianceInfoInState(cliCtx, w, r, storeName, types.Certified)
	}
}

func getRevokedModelsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		getAllComplianceInfo(cliCtx, w, r, fmt.Sprintf("custom/%s/all_revoked_models", storeName))
	}
}

func getRevokedModelHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		getComplianceInfoInState(cliCtx, w, r, storeName, types.Revoked)
	}
}

func getComplianceInfoInState(cliCtx context.CLIContext, w http.ResponseWriter, r *http.Request,
	storeName string, state types.ComplianceState) {
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

	certificationType := types.CertificationType(vars[certificationType])

	isInState := types.ComplianceInfoInState{Value: false}

	res, height, err := restCtx.QueryStore(types.GetComplianceInfoKey(certificationType, vid, pid), storeName)
	if res != nil {
		var complianceInfo types.ComplianceInfo

		restCtx.Codec().MustUnmarshalBinaryBare(res, &complianceInfo)

		isInState.Value = complianceInfo.State == state
	}

	if err != nil {
		restCtx.WriteErrorResponse(http.StatusNotFound,
			types.ErrComplianceInfoDoesNotExist(vid, pid, certificationType).Error())

		return
	}

	restCtx.EncodeAndRespondWithHeight(isInState, height)
}

func getComplianceInfo(cliCtx context.CLIContext, w http.ResponseWriter, r *http.Request, storeName string) {
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

	certificationType := types.CertificationType(vars[certificationType])

	res, height, err := restCtx.QueryStore(types.GetComplianceInfoKey(certificationType, vid, pid), storeName)
	if err != nil || res == nil {
		restCtx.WriteErrorResponse(http.StatusNotFound,
			types.ErrComplianceInfoDoesNotExist(vid, pid, certificationType).Error())

		return
	}

	var complianceInfo types.ComplianceInfo

	restCtx.Codec().MustUnmarshalBinaryBare(res, &complianceInfo)

	restCtx.EncodeAndRespondWithHeight(complianceInfo, height)
}

func getAllComplianceInfo(cliCtx context.CLIContext, w http.ResponseWriter, r *http.Request, path string) {
	restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

	paginationParams, err := restCtx.ParsePaginationParams()
	if err != nil {
		return
	}

	certificationType := types.CertificationType(restCtx.Request().FormValue(certificationType))
	params := types.NewListQueryParams(certificationType, paginationParams.Skip, paginationParams.Take)

	restCtx.QueryList(path, params)
}
