module git.dsr-corporation.com/zb-ledger/zb-ledger

go 1.13

require (
	github.com/cosmos/cosmos-sdk v0.34.4-0.20200819073641-f02b0b574501
	github.com/cosmos/go-bip39 v0.0.0-20180819234021-555e2067c45d
	github.com/etcd-io/bbolt v1.3.3 // indirect
	github.com/gorilla/mux v1.7.4
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	github.com/stumble/gorocksdb v0.0.3 // indirect
	github.com/tendermint/go-amino v0.15.1
	github.com/tendermint/tendermint v0.34.0-rc3
	github.com/tendermint/tm-db v0.6.1
)

replace github.com/cosmos/cosmos-sdk => ../cosmos-sdk
