package types

import (
	"bytes"
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	tmtypes "github.com/tendermint/tendermint/types"
	"strings"
)

var Weight = sdk.NewInt(10)

/*
	Validator
*/
type Validator struct {
	Description     stakingtypes.Description `json:"description"`
	OperatorAddress sdk.ValAddress           `json:"operator_address"`
	ConsensusPubKey string                   `json:"consensus_pubkey"`
	Weight          sdk.Int                  `json:"weight"`
}

func NewValidator(address sdk.ValAddress, pubKey string, description stakingtypes.Description) Validator {
	return Validator{
		Description:     description,
		OperatorAddress: address,
		ConsensusPubKey: pubKey,
		Weight:          Weight,
	}
}

func (v Validator) GetConsAddress() sdk.ConsAddress {
	return sdk.ConsAddress(v.OperatorAddress)
}

func (v Validator) GetConsPubKey() crypto.PubKey {
	return sdk.MustGetConsPubKeyBech32(v.ConsensusPubKey)
}

func (v Validator) GetWeight() int64 { return v.Weight.Int64() }

func (v Validator) GetMoniker() string { return v.Description.Moniker }

// ABCIValidatorUpdate returns an abci.ValidatorUpdate from a validator validator type
// with the full validator power
func (v Validator) ABCIValidatorUpdate() abci.ValidatorUpdate {
	return abci.ValidatorUpdate{
		PubKey: tmtypes.TM2PB.PubKey(v.GetConsPubKey()),
		Power:  v.Weight.Int64(),
	}
}

func (v Validator) String() string {
	bytes, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

/*
	Validators collection
*/
type Validators []Validator

func (v Validators) String() (out string) {
	for _, val := range v {
		out += val.String() + "\n"
	}
	return strings.TrimSpace(out)
}

// Implements sort interface
func (v Validators) Len() int {
	return len(v)
}

// Implements sort interface
func (v Validators) Less(i, j int) bool {
	return bytes.Compare(v[i].OperatorAddress, v[j].OperatorAddress) == -1
}

// Implements sort interface
func (v Validators) Swap(i, j int) {
	it := v[i]
	v[i] = v[j]
	v[j] = it
}

func MustMarshalValidator(cdc *codec.Codec, validator Validator) []byte {
	return cdc.MustMarshalBinaryBare(&validator)
}

func MustUnmarshalValidator(cdc *codec.Codec, value []byte) Validator {
	validator, err := UnmarshalValidator(cdc, value)
	if err != nil {
		panic(err)
	}
	return validator
}

func UnmarshalValidator(cdc *codec.Codec, value []byte) (v Validator, err error) {
	err = cdc.UnmarshalBinaryBare(value, &v)
	return v, err
}
