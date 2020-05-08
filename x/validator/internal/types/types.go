package types

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	tmtypes "github.com/tendermint/tendermint/types"
)

/*
	Validator
*/
type Validator struct {
	Description     stakingtypes.Description `json:"description"`
	OperatorAddress sdk.ValAddress           `json:"operator_address"`
	ConsensusPubKey string                   `json:"consensus_pubkey"`
	Power           int64                    `json:"power"`
	Jailed          bool                     `json:"jailed"`
	JailedReason    string                   `json:"jailed_reason,omitempty"`
}

func NewValidator(address sdk.ValAddress, pubKey string, description stakingtypes.Description) Validator {
	return Validator{
		Description:     description,
		OperatorAddress: address,
		ConsensusPubKey: pubKey,
		Power:           Power,
		Jailed:          false,
	}
}

func (v Validator) GetOperatorAddress() sdk.ValAddress {
	return v.OperatorAddress
}

func (v Validator) GetConsAddress() sdk.ConsAddress {
	return sdk.ConsAddress(v.GetConsPubKey().Address())
}

func (v Validator) GetConsPubKey() crypto.PubKey {
	return sdk.MustGetConsPubKeyBech32(v.ConsensusPubKey)
}

func (v Validator) GetPower() int64 { return v.Power }

func (v Validator) GetMoniker() string { return v.Description.Moniker }

func (v Validator) IsJailed() bool { return v.Jailed }

// ABCI ValidatorUpdate message to add new validator to validator set
func (v Validator) ABCIValidatorUpdate() abci.ValidatorUpdate {
	return abci.ValidatorUpdate{
		PubKey: tmtypes.TM2PB.PubKey(v.GetConsPubKey()),
		Power:  Power,
	}
}

// ABCI ValidatorUpdate message to remove validator from validator set
func (v Validator) ABCIValidatorUpdateZero() abci.ValidatorUpdate {
	return abci.ValidatorUpdate{
		PubKey: tmtypes.TM2PB.PubKey(v.GetConsPubKey()),
		Power:  ZeroPower,
	}
}

func (v Validator) String() string {
	bytes, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func MustMarshalValidator(cdc *codec.Codec, validator Validator) []byte {
	return cdc.MustMarshalBinaryBare(&validator)
}

func MustUnmarshalBinaryBareValidator(cdc *codec.Codec, value []byte) Validator {
	validator, err := UnmarshalBinaryBareValidator(cdc, value)
	if err != nil {
		panic(err)
	}
	return validator
}

func UnmarshalBinaryBareValidator(cdc *codec.Codec, value []byte) (v Validator, err error) {
	err = cdc.UnmarshalBinaryBare(value, &v)
	return v, err
}

/*
	Last Validator power. needed for taking validator set updates
*/
type LastValidatorPower struct {
	OperatorAddress sdk.ValAddress `json:"operator_address"`
	Power           int64          `json:"power"`
}

func NewLastValidatorPower(address sdk.ValAddress) LastValidatorPower {
	return LastValidatorPower{
		OperatorAddress: address,
		Power:           Power,
	}
}

/*
	Validator Signing info
*/
// Signing info for a validator
type ValidatorSigningInfo struct {
	Address             sdk.ConsAddress `json:"address"`                          // validator consensus address
	StartHeight         int64           `json:"start_height" yaml:"start_height"` // height at which validator was first a candidate OR was unjailed
	IndexOffset         int64           `json:"index_offset"`                     // index offset into signed block bit array
	MissedBlocksCounter int64           `json:"missed_blocks_counter"`            // missed blocks counter (to avoid scanning the array every time)
}

func NewValidatorSigningInfo(condAddr sdk.ConsAddress, startHeight int64) ValidatorSigningInfo {
	return ValidatorSigningInfo{
		Address:             condAddr,
		StartHeight:         startHeight,
		IndexOffset:         0,
		MissedBlocksCounter: 0,
	}
}

func (v ValidatorSigningInfo) String() string {
	bytes, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
