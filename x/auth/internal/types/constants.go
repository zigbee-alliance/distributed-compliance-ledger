package types

// Default parameter values.
const (
	DefaultMaxMemoCharacters           uint64  = 256
	DefaultTxSizeCostPerByte           uint64  = 10
	DefaultSigVerifyCostED25519        uint64  = 590
	DefaultSigVerifyCostSecp256k1      uint64  = 1000
	DefaultApproveAddAccountPercent    float64 = 0.66
	DefaultApproveRevokeAccountPercent float64 = 0.66
)
