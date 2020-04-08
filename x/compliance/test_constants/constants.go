package test_constants

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

var (
	Id                       = "Device Id"
	Name                     = "Device Name"
	Owner                    = sdk.AccAddress([]byte("owner"))
	Description              = "Device Description"
	Sku                      = "RCU2205A"
	FirmwareVersion          = "1.0"
	HardwareVersion          = "2.0"
	CertificateID            = "ZIG12345678"
	CertifiedDate            = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	TisOrTrpTestingCompleted = false
	Signer                   = sdk.AccAddress([]byte("signer"))
)
