package test_constants

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

var (
	VID                      int16 = 1
	PID                      int16 = 22
	CID                      int16 = 12345
	Name                           = "Device Name"
	Owner                          = sdk.AccAddress([]byte("me"))
	Description                    = "Device Description"
	Sku                            = "RCU2205A"
	FirmwareVersion                = "1.0"
	HardwareVersion                = "2.0"
	Custom                         = "Custom data"
	CertificateID                  = "ZIG12345678"
	CertifiedDate                  = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	TisOrTrpTestingCompleted       = false
	Signer                         = sdk.AccAddress([]byte("me"))
)
