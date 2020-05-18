package types

//nolint:goimports
import (
	"encoding/binary"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName is the name of the validator module.
	ModuleName = "validator"

	// StoreKey is the string store representation.
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the validator module.
	QuerierRoute = ModuleName

	// RouterKey is the msg router key for validator module.
	RouterKey = ModuleName
)

var (
	ValidatorPrefix          = []byte{0x01} // prefix for each key to a validator
	ValidatorLastPowerPrefix = []byte{0x02} // prefix for each key to a validator index, by last power

	ValidatorOwnerPrefix               = []byte{0x05} // prefix for validator owner
	ValidatorSigningInfoPrefix         = []byte{0x06} // prefix for validator signing info
	ValidatorMissedBlockBitArrayPrefix = []byte{0x07} // prefix for validator missed blocks

)

// Key builder for Validator record.
func GetValidatorKey(addr sdk.ConsAddress) []byte {
	return append(ValidatorPrefix, addr.Bytes()...)
}

// Key builder for Last Validator Power record.
func GetValidatorLastPowerKey(addr sdk.ConsAddress) []byte {
	return append(ValidatorLastPowerPrefix, addr.Bytes()...)
}

// Key builder for Validator signing info record.
func GetValidatorSigningInfoKey(addr sdk.ConsAddress) []byte {
	return append(ValidatorSigningInfoPrefix, addr.Bytes()...)
}

// Key builder for Validator owner record.
func GetValidatorOwnerKey(addr sdk.AccAddress) []byte {
	return append(ValidatorOwnerPrefix, addr.Bytes()...)
}

func GetValidatorMissedBlockBitArrayPrefixKey(v sdk.ConsAddress) []byte {
	return append(ValidatorMissedBlockBitArrayPrefix, v.Bytes()...)
}

// Key builder for Validator Missed blocks.
func GetValidatorMissedBlockBitArrayKey(v sdk.ConsAddress, i int64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(i))

	return append(GetValidatorMissedBlockBitArrayPrefixKey(v), b...)
}
