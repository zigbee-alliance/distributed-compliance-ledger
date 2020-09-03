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
