package main

import (
	"errors"
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/zigbee-alliance/distributed-compliance-ledger/app"
	"github.com/zigbee-alliance/distributed-compliance-ledger/cmd/dcld/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		var e server.ErrorCode
		switch {
		case errors.As(err, &e):
			os.Exit(e.Code)
		default:
			os.Exit(1)
		}
	}
}
