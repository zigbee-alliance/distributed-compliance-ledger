module github.com/zigbee-alliance/distributed-compliance-ledger

go 1.16

require (
	github.com/cosmos/cosmos-proto v0.0.0-20210914142853-23ed61ac79ce
	github.com/cosmos/cosmos-sdk v0.44.0
	github.com/cosmos/go-bip39 v1.0.0
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.8.0
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/spm v0.1.5
	github.com/tendermint/tendermint v0.34.13
	github.com/tendermint/tm-db v0.6.4
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa
	google.golang.org/grpc v1.43.0
	sigs.k8s.io/yaml v1.3.0
)

replace (
	github.com/99designs/keyring => github.com/cosmos/keyring v1.1.7-0.20210622111912-ef00f8ac3d76
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)
