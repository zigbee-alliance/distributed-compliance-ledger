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
	Description  Description     `json:"description"`             // description of the validator
	Address      sdk.ConsAddress `json:"validator_address"`       // the consensus address of the tendermint validator
	PubKey       string          `json:"validator_pubkey"`        // the consensus public key of the tendermint validator
	Power        int64           `json:"power"`                   // validator consensus power
	Jailed       bool            `json:"jailed"`                  // has the validator been removed from validator set
	JailedReason string          `json:"jailed_reason,omitempty"` // the reason of validator jailing
	Owner        sdk.AccAddress  `json:"owner"`                   // the account address of validator owner
}

func NewValidator(address sdk.ConsAddress, pubKey string, description Description, owner sdk.AccAddress) Validator {
	return Validator{
		Description: description,
		Address:     address,
		PubKey:      pubKey,
		Power:       Power,
		Jailed:      false,
		Owner:       owner,
	}
}

func (v Validator) GetConsAddress() sdk.ConsAddress {
	return sdk.ConsAddress(v.GetConsPubKey().Address())
}

func (v Validator) GetConsPubKey() crypto.PubKey {
	return sdk.MustGetConsPubKeyBech32(v.PubKey)
}

func (v Validator) GetPower() int64 { return v.Power }

func (v Validator) GeOwner() sdk.AccAddress { return v.Owner }

func (v Validator) GetName() string { return v.Description.Name }

func (v Validator) IsJailed() bool { return v.Jailed }

// ABCI ValidatorUpdate message to add new validator to validator set.
func (v Validator) ABCIValidatorUpdate() abci.ValidatorUpdate {
	return abci.ValidatorUpdate{
		PubKey: tmtypes.TM2PB.PubKey(v.GetConsPubKey()),
		Power:  v.GetPower(),
	}
}

// ABCI ValidatorUpdate message to remove validator from validator set.
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
	ConsensusAddress sdk.ConsAddress `json:"address"`
	Power            int64           `json:"power"`
}

func NewLastValidatorPower(address sdk.ConsAddress) LastValidatorPower {
	return LastValidatorPower{
		ConsensusAddress: address,
		Power:            Power,
	}
}

/*
	Description of Validator
*/
type Description struct {
	// name.
	Name string `json:"name"`
	// optional identity signature.
	Identity string `json:"identity,omitempty"`
	// optional website link.
	Website string `json:"website,omitempty"`
	// optional details.
	Details string `json:"details,omitempty"`
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
		return sdk.ErrUnknownRequest(fmt.Sprintf(
			"Invalid Description Name: received string of length %v, max is %v", len(d.Name), MaxNameLength))
	}

	if len(d.Identity) > MaxIdentityLength {
		return sdk.ErrUnknownRequest(fmt.Sprintf(
			"Invalid Description Identity: "+
				"received string of length %v, max is %v", len(d.Identity), MaxIdentityLength))
	}

	if len(d.Website) > MaxWebsiteLength {
		return sdk.ErrUnknownRequest(fmt.Sprintf(
			"Invalid Description Website: "+
				"received string of length %v, max is %v", len(d.Website), MaxWebsiteLength))
	}

	if len(d.Details) > MaxDetailsLength {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Description Details: "+
			"received string of length %v, max is %v", len(d.Details), MaxDetailsLength))
	}

	return nil
}

/*
	Validator Signing info
*/
type ValidatorSigningInfo struct {
	// validator consensus address.
	Address sdk.ConsAddress `json:"address"`
	// height at which validator was added.
	StartHeight int64 `json:"start_height"`
	// index offset into signed block bit array.
	IndexOffset int64 `json:"index_offset"`
	// missed blocks counter (to avoid scanning the array every time).
	MissedBlocksCounter int64 `json:"missed_blocks_counter"`
}

func NewValidatorSigningInfo(address sdk.ConsAddress, startHeight int64) ValidatorSigningInfo {
	return ValidatorSigningInfo{
		Address:             address,
		StartHeight:         startHeight,
		IndexOffset:         0,
		MissedBlocksCounter: 0,
	}
}

func (v ValidatorSigningInfo) Reset() ValidatorSigningInfo {
	v.MissedBlocksCounter = 0
	v.IndexOffset = 0

	return v
}

func (v ValidatorSigningInfo) String() string {
	bytes, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
