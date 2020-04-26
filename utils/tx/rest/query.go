package rest

import (
	"encoding/base64"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/rest"
	"io/ioutil"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/cosmos/cosmos-sdk/x/auth/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/cosmos/cosmos-sdk/client/context"
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

func SignTxHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		err := r.ParseForm()
		if err != nil {
			restCtx.WriteErrorResponse(http.StatusBadRequest,
				sdk.AppendMsgToErr("could not parse query parameters", err.Error()))
			return
		}

		name, passphrase := r.FormValue("name"), r.FormValue("passphrase")

		var signMsg types.StdSignMsg
		if !restCtx.ReadRESTReq(&signMsg) {
			return
		}

		txBldr := auth.NewTxBuilderFromCLI().
			WithTxEncoder(utils.GetTxEncoder(restCtx.Codec())).
			WithAccountNumber(signMsg.AccountNumber).
			WithSequence(signMsg.Sequence).
			WithChainID(signMsg.ChainID)

		stdTx := auth.StdTx{
			Msgs:       signMsg.Msgs,
			Fee:        signMsg.Fee,
			Signatures: nil,
			Memo:       signMsg.Memo,
		}

		signedStdTx, err := txBldr.SignStdTx(name, passphrase, stdTx, false)

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

		var stdTx types.StdTx
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
