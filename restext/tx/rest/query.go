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
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/rest"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	restTypes "github.com/cosmos/cosmos-sdk/types/rest"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// EncodeTxRequestHandlerFn returns the decode tx REST handler. In particular,
// it takes a base64-encoded bytes, decodes it using the Amino wire protocol,
// and responds with JSON-encoded transaction.
func DecodeTxRequestHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest, err.Error())

			return
		}

		var req DecodeTxsRequest

		err = restCtx.Codec().UnmarshalJSON(body, &req)
		if err != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest, err.Error())

			return
		}

		resp := DecodeTxsResponse{
			Txs: []auth.StdTx{},
		}

		for _, base64str := range req.Txs {
			tx, err := decodeTx(restCtx.Codec(), base64str)
			if err != nil {
				restCtx.WriteErrorResponse(http.StatusBadRequest, err.Error())

				return
			}

			resp.Txs = append(resp.Txs, tx)
		}

		restCtx.PostProcessResponseBare(&resp)
	}
}

func decodeTx(cdc *codec.Codec, base64str string) (tx auth.StdTx, err error) {
	var res auth.StdTx

	bytes, err := base64.StdEncoding.DecodeString(base64str)
	if err != nil {
		return auth.StdTx{}, err
	}

	err = cdc.UnmarshalBinaryLengthPrefixed(bytes, &res)
	if err != nil {
		return auth.StdTx{}, err
	}

	return res, nil
}

type SignMessageRequest struct {
	BaseReq restTypes.BaseReq `json:"base_req"`
	Txn     Txn               `json:"txn"`
}

type Txn struct {
	Type_ string     `json:"type"`
	Value auth.StdTx `json:"value"`
}

func SignMessageHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		var req SignMessageRequest
		if !restCtx.ReadRESTReq(&req) {
			return
		}

		println(fmt.Sprintf("%v", req))

		err := r.ParseForm()
		if err != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest,
				sdk.AppendMsgToErr("could not parse query parameters", err.Error()))

			return
		}

		account, passphrase, ok := restCtx.BasicAuth()
		if !ok {
			restCtx.WriteErrorResponse(http.StatusBadRequest, "Could not find credentials to use")

			return
		}

		restCtx, err = restCtx.WithBaseRequest(req.BaseReq)
		if err != nil {
			return
		}

		restCtx, err = restCtx.WithSigner()
		if err != nil {
			return
		}

		txBldr, err := restCtx.TxnBuilder()
		if err != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest, err.Error())

			return
		}

		signedStdTx, err := txBldr.SignStdTx(account, passphrase, req.Txn.Value, false)
		if err != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest, err.Error())

			return
		}

		restCtx.PostProcessResponse(signedStdTx)
	}
}

func BroadcastTxHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		var stdTx auth.StdTx
		if !restCtx.ReadRESTReq(&stdTx) {
			return
		}

		txBytes, err := restCtx.Codec().MarshalBinaryLengthPrefixed(stdTx)
		if err != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest, err.Error())

			return
		}

		res, err := restCtx.BroadcastMessage(txBytes)
		if err != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest, err.Error())

			return
		}

		restCtx.PostProcessResponse(res)
	}
}
