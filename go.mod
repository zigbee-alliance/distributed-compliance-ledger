module github.com/zigbee-alliance/distributed-compliance-ledger

go 1.13

require (
	github.com/cosmos/cosmos-sdk v0.37.4
	github.com/cosmos/go-bip39 v0.0.0-20180618194314-52158e4697b8
	github.com/gorilla/mux v1.7.3
	github.com/pkg/errors v0.8.1
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.6.1
	github.com/stretchr/testify v1.4.0
	github.com/tendermint/go-amino v0.15.1
	github.com/tendermint/tendermint v0.32.8
	github.com/tendermint/tm-db v0.2.0
	golang.org/x/net v0.0.0-20200625001655-4c5254603344 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

replace github.com/cosmos/cosmos-sdk => github.com/zigbee-alliance/cosmos-sdk v0.37.5-0.20200828165740-d1da07e38b94
