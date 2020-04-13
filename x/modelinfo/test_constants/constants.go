package test_constants

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

var (
	ChainId                        = "zbl-test-chain-id"
	VID                      int16 = 1
	PID                      int16 = 22
	CID                      int16 = 12345
	Name                           = "Device Name"
	Owner                          = Address1
	Description                    = "Device Description"
	Sku                            = "RCU2205A"
	FirmwareVersion                = "1.0"
	HardwareVersion                = "2.0"
	Custom                         = "Custom data"
	CertificateID                  = "ZIG12345678"
	CertifiedDate                  = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	TisOrTrpTestingCompleted       = false
	Address1, _                    = sdk.AccAddressFromBech32("cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz")
	Address2, _                    = sdk.AccAddressFromBech32("cosmos1j8x9urmqs7p44va5p4cu29z6fc3g0cx2c2vxx2")
	Address3, _                    = sdk.AccAddressFromBech32("cosmos1j7tc5f4f54fd8hns42nsavzhadr0gchddz6vfl")
	Signer                         = Address1
)
