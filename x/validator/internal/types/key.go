package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleNa,e is the name of the validator module
	ModuleName = "validator"

	// StoreKey is the string store representation
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the validator module
	QuerierRoute = ModuleName

	// RouterKey is the msg router key for validator module
	RouterKey = ModuleName
)

var (
	ValidatorKey           = []byte{0x1} // prefix for each key to a validator
	ValidatorByConsAddrKey = []byte{0x2} // prefix for each key to a validator index, by pubkey
	ValidatorLastPowerKey  = []byte{0x3} // prefix for each key to a validator index, by last power
)

// Key builder for Validator record
func GetValidatorKey(addr sdk.ValAddress) []byte {
	return append(ValidatorKey, addr.Bytes()...)
}

// Key builder for Consensus Address to Validator Address mapping record
func GetValidatorByConsAddrKey(addr sdk.ConsAddress) []byte {
	return append(ValidatorByConsAddrKey, addr.Bytes()...)
}

// Key builder for Last Validator Power record
func GetValidatorLastPowerKey(addr sdk.ValAddress) []byte {
	return append(ValidatorLastPowerKey, addr.Bytes()...)
}
