package rest

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/gorilla/mux"
)

const (
	keyNameKey = "name"
)

func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/key", KeysHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/key/{%s}", keyNameKey), KeyHandlerFn(cliCtx)).Methods("GET")
}
