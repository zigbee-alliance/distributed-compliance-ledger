package rest

import (
	"encoding/json"
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/pagination"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"net/http"
)

func getComplianceInfoHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		getComplianceInfo(cliCtx, w, r, storeName, "")
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
		getComplianceInfo(cliCtx, w, r, storeName, types.Certified)
	}
}

func getRevokedModelsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		getAllComplianceInfo(cliCtx, w, r, fmt.Sprintf("custom/%s/all_revoked_models", storeName))
	}
}

func getRevokedModelHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		getComplianceInfo(cliCtx, w, r, storeName, types.Revoked)
	}
}

func getComplianceInfo(cliCtx context.CLIContext, w http.ResponseWriter, r *http.Request, storeName string, state types.ComplianceState) {
	cliCtx = context.NewCLIContext().WithCodec(cliCtx.Codec)

	vars := mux.Vars(r)
	vid := vars[vid]
	pid := vars[pid]
	certificationType := types.CertificationType(r.FormValue("certification_type"))

	res, height, err := cliCtx.QueryStore([]byte(keeper.ComplianceInfoId(certificationType, vid, pid)), storeName)
	if err != nil || res == nil {
		rest.WriteErrorResponse(w, http.StatusNotFound, types.ErrComplianceInfoDoesNotExist(vid, pid).Error())
		return
	}

	var complianceInfo types.ComplianceInfo
	cliCtx.Codec.MustUnmarshalBinaryBare(res, &complianceInfo)

	if len(state) > 0 && complianceInfo.State != state {
		rest.WriteErrorResponse(w, http.StatusNotFound, types.ErrComplianceInfoDoesNotExist(vid, pid).Error())
		return
	}

	out, err := json.Marshal(complianceInfo)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	cliCtx.Height = height
	rest.PostProcessResponse(w, cliCtx, out)
}

func getAllComplianceInfo(cliCtx context.CLIContext, w http.ResponseWriter, r *http.Request, path string) {
	cliCtx = context.NewCLIContext().WithCodec(cliCtx.Codec)

	paginationParams, err := pagination.ParsePaginationParamsFromRequest(cliCtx.Codec, r)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	certificationType := types.CertificationType(r.FormValue("certification_type"))

	params := types.NewListQueryParams(certificationType, paginationParams.Skip, paginationParams.Take)

	res, height, err := cliCtx.QueryWithData(path, cliCtx.Codec.MustMarshalJSON(params))
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	cliCtx.Height = height
	rest.PostProcessResponse(w, cliCtx, res)
}