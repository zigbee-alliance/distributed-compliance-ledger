package rest

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// EncodeTxRequestHandlerFn returns the decode tx REST handler. In particular,
// it takes a base64-encoded bytes, decodes it using the Amino wire protocol,
// and responds with JSON-encoded transaction.
func DecodeTxRequestHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		var req DecodeTxsRequest

		err = cliCtx.Codec.UnmarshalJSON(body, &req)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		resp := DecodeTxsResponse{
			Txs: []auth.StdTx{},
		}

		for _, base64str := range req.Txs {
			tx, err := decodeTx(cliCtx.Codec, base64str)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}

			resp.Txs = append(resp.Txs, tx)
		}

		rest.PostProcessResponseBare(w, cliCtx, &resp)
	}
}

func decodeTx(cdc *codec.Codec, base64str string) (tx auth.StdTx, err error) {
	var res types.StdTx

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
