package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/gorilla/mux"
)

func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/tx/decode", DecodeTxRequestHandlerFn(cliCtx)).Methods("POST")
	r.HandleFunc("/tx/sign", SignMessageHandlerFn(cliCtx)).Methods("POST")
	r.HandleFunc("/tx/broadcast", BroadcastTxHandlerFn(cliCtx)).Methods("POST")
}
