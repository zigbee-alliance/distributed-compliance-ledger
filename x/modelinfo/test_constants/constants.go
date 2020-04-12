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
	Owner, _                       = sdk.AccAddressFromBech32("cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz")
	Description                    = "Device Description"
	Sku                            = "RCU2205A"
	FirmwareVersion                = "1.0"
	HardwareVersion                = "2.0"
	Custom                         = "Custom data"
	CertificateID                  = "ZIG12345678"
	CertifiedDate                  = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	TisOrTrpTestingCompleted       = false
	Signer, _                      = sdk.AccAddressFromBech32("cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz")
)
