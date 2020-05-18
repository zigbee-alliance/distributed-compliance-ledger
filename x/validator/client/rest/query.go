package rest

//nolint:goimports
import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/pagination"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/rest"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"net/http"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"
)

// HTTP request handler to query list of validators.
func getValidatorsHandlerFn(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		state := r.FormValue(state)
		paginationParams := pagination.ParsePaginationParamsFromFlags()
		params := types.NewListValidatorsParams(paginationParams, types.ValidatorState(state))

		restCtx.QueryList(fmt.Sprintf("custom/%s/validators", storeName), params)
	}
}

// HTTP request handler to query the validator information from a given validator address.
func getValidatorHandlerFn(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		vars := restCtx.Variables()
		bech32validatorAddr := vars[validatorAddr]

		validatorAddr, err := sdk.ConsAddressFromBech32(bech32validatorAddr)
		if err != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest, err.Error())
			return
		}

		res, height, err := restCtx.QueryStore(types.GetValidatorKey(validatorAddr), storeName)
		if err != nil || res == nil {
			restCtx.WriteErrorResponse(http.StatusNotFound, types.ErrValidatorDoesNotExist(validatorAddr).Error())
			return
		}

		validator := types.MustUnmarshalBinaryBareValidator(restCtx.Codec(), res)

		restCtx.EncodeAndRespondWithHeight(validator, height)
	}
}
