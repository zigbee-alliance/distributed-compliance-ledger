module github.com/zigbee-alliance/distributed-compliance-ledger

go 1.20

require (
	cosmossdk.io/api v0.3.1
	github.com/cometbft/cometbft v0.37.1
	github.com/cometbft/cometbft-db v0.7.0
	github.com/cosmos/cosmos-sdk v0.47.3
	github.com/cosmos/gogoproto v1.4.10
	github.com/go-playground/locales v0.14.0
	github.com/go-playground/universal-translator v0.18.0
	github.com/go-playground/validator/v10 v10.9.0
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cast v1.5.0
	github.com/spf13/cobra v1.6.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.14.0
	github.com/stretchr/testify v1.8.2
	google.golang.org/grpc v1.56.2
	sigs.k8s.io/yaml v1.3.0
)

require (
	github.com/kr/pretty v0.3.1 // indirect
	golang.org/x/net v0.12.0 // indirect
)

replace (
	cosmossdk.io/simapp => cosmossdk.io/simapp v0.0.0-20240119091742-c5826868d4e5
	github.com/99designs/keyring => github.com/cosmos/keyring v1.1.7-0.20210622111912-ef00f8ac3d76
	github.com/syndtr/goleveldb => github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7
)
