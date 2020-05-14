package rest

//nolint:goimports
import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"net/http"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/rest"
	"github.com/cosmos/cosmos-sdk/client/context"
)

func accountsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		params, err := restCtx.ParsePaginationParams()
		if err != nil {
			return
		}

		restCtx.QueryList(fmt.Sprintf("custom/%s/accounts", storeName), params)
	}
}

func accountHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		vars := restCtx.Variables()
		accAddr := vars[addrKey]

		address, err := sdk.AccAddressFromBech32(accAddr)
		if err != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest, sdk.ErrInvalidAddress(accAddr).Error())
			return
		}

		res, height, err := cliCtx.QueryStore(types.GetAccountKey(address), storeName)
		if err != nil {
			restCtx.WriteErrorResponse(http.StatusNotFound, err.Error())
			return
		}

		restCtx.RespondWithHeight(res, height)
	}
}
