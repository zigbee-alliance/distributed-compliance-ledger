package rest

import (
	"net/http"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/rest"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/keys"
	keys2 "github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/gorilla/mux"
)

// Lists all keys in the local keychain.
func KeysHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		kb, err := keys.NewKeyBaseFromHomeFlag()
		if err != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest, err.Error())

			return
		}

		infos, err := kb.List()
		if err != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest, err.Error())

			return
		}

		outputs, err := keys2.Bech32KeysOutput(infos)
		if err != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest, err.Error())

			return
		}

		resp := resultKeyInfos{
			Total: len(infos),
			Items: outputs,
		}

		restCtx.PostProcessResponseBare(&resp)
	}
}

// Lists all keys in the local keychain.
func KeyHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		vars := mux.Vars(r)
		keyName := vars[keyNameKey]

		kb, err := keys.NewKeyBaseFromHomeFlag()
		if err != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest, err.Error())

			return
		}

		keyInfo, err := kb.Get(keyName)
		if err != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest, err.Error())

			return
		}

		keyOutput, err := keys2.Bech32KeyOutput(keyInfo)
		if err != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest, err.Error())

			return
		}

		restCtx.PostProcessResponseBare(&keyOutput)
	}
}
