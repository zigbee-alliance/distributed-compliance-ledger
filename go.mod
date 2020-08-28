module git.dsr-corporation.com/zb-ledger/zb-ledger

go 1.13

require (
	github.com/cosmos/cosmos-sdk v0.37.4
	github.com/cosmos/go-bip39 v0.0.0-20180618194314-52158e4697b8
	github.com/daixiang0/gci v0.2.3 // indirect
	github.com/gorilla/mux v1.7.3
	github.com/pkg/errors v0.8.1
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.6.1
	github.com/stretchr/testify v1.4.0
	github.com/tendermint/go-amino v0.15.1
	github.com/tendermint/tendermint v0.32.8
	github.com/tendermint/tm-db v0.2.0
	mvdan.cc/gofumpt v0.0.0-20200802201014-ab5a8192947d // indirect
)

//github.com/cosmos/cosmos-sdk => github.com/zigbee-alliance/cosmos-sdk multiproofs
replace github.com/cosmos/cosmos-sdk => ../cosmos-sdk
