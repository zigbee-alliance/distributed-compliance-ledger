package rest

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/rest"
	"net/http"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	restTypes "github.com/cosmos/cosmos-sdk/types/rest"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type CreateValidatorRequest struct {
	BaseReq          restTypes.BaseReq        `json:"base_req"`
	ValidatorAddress sdk.ValAddress           `json:"validator_address"`
	Pubkey           string                   `json:"pubkey"`
	Description      stakingtypes.Description `json:"description"`
}

func createValidatorHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		var req CreateValidatorRequest
		if !restCtx.ReadRESTReq(&req) {
			return
		}

		restCtx, err := restCtx.WithBaseRequest(req.BaseReq)
		if err != nil {
			return
		}

		_, err = sdk.GetConsPubKeyBech32(req.Pubkey)
		if err != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgCreateValidator(req.ValidatorAddress, req.Pubkey, req.Description)

		restCtx.HandleWriteRequest(msg)
	}
}
