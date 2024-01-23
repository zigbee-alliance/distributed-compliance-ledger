package main

import (
	"github.com/cosmos/cosmos-sdk/server"
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/zigbee-alliance/distributed-compliance-ledger/app"
	"github.com/zigbee-alliance/distributed-compliance-ledger/cmd/dcld/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
