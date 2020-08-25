package types

// Default parameter values.
const (
	MaxMemoCharacters             uint64  = 256
	TxSizeCostPerByte             uint64  = 10
	DefaultSigVerifyCostED25519   uint64  = 590
	DefaultSigVerifyCostSecp256k1 uint64  = 1000
	AccountApprovalPercent        float64 = 0.66
)
