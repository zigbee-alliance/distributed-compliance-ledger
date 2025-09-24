package helpers

import (
	"github.com/cosmos/cosmos-sdk/codec"
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govtypesv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/zigbee-alliance/distributed-compliance-ledger/app"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

var (
	Codec codec.Codec
)

func init() {
	encodingConfig := app.MakeEncodingConfig()
	govtypesv1.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	govtypesv1beta1.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	dclauthtypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	pkitypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	Codec = encodingConfig.Codec
}
