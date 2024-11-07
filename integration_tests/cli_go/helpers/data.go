package helpers

import (
	tmrand "github.com/cometbft/cometbft/libs/rand"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

func RandomString() string {
	return utils.RandString()
}

func RandomVid() int32 {
	return int32(tmrand.Uint16())
}
