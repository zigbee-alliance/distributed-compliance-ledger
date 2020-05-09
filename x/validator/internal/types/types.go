package types

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	tmtypes "github.com/tendermint/tendermint/types"
)

/*
	Validator
*/
type Validator struct {
	Description     Description    `json:"description"`             // description of the validator
	OperatorAddress sdk.ValAddress `json:"operator_address"`        // address of the validator's operator
	ConsensusPubKey string         `json:"consensus_pubkey"`        // the consensus public key of the validator
	Power           int64          `json:"power"`                   // validator consensus power
	Jailed          bool           `json:"jailed"`                  // has the validator been removed from validator set
	JailedReason    string         `json:"jailed_reason,omitempty"` // the reason of validator jailing
}

func NewValidator(address sdk.ValAddress, pubKey string, description Description) Validator {
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

func (v Validator) GetName() string { return v.Description.Name }

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
	Description of Validator
*/
type Description struct {
	Name     string `json:"name"`     // name
	Identity string `json:"identity"` // optional identity signature
	Website  string `json:"website"`  // optional website link
	Details  string `json:"details"`  // optional details
}

// NewDescription returns a new Description with the provided values.
func NewDescription(name, identity, website, details string) Description {
	return Description{
		Name:     name,
		Identity: identity,
		Website:  website,
		Details:  details,
	}
}

const (
	MaxNameLength     = 70
	MaxIdentityLength = 3000
	MaxWebsiteLength  = 140
	MaxDetailsLength  = 280
)

// Ensure the length of a validator's description.
func (d Description) Validate() sdk.Error {
	if len(d.Name) == 0 {
		return sdk.ErrUnknownRequest("Invalid Description Name: it cannot be empty")
	}
	if len(d.Name) > MaxNameLength {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Description Name: received string of lenght %v, max is %v", len(d.Name), MaxNameLength))
	}
	if len(d.Identity) > MaxIdentityLength {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Description Identity: received string of lenght %v, max is %v", len(d.Identity), MaxIdentityLength))
	}
	if len(d.Website) > MaxWebsiteLength {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Description Identity: received string of lenght %v, max is %v", len(d.Website), MaxWebsiteLength))
	}
	if len(d.Details) > MaxDetailsLength {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Description Details: received string of lenght %v, max is %v", len(d.Details), MaxDetailsLength))
	}
	return nil
}

/*
	Validator Signing info
*/
type ValidatorSigningInfo struct {
	Address             sdk.ConsAddress `json:"address"`                          // validator consensus address
	StartHeight         int64           `json:"start_height" yaml:"start_height"` // height at which validator was added
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
