package types

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"time"

	abci "github.com/tendermint/tendermint/abci/types"
	tmprotocrypto "github.com/tendermint/tendermint/proto/tendermint/crypto"
	"sigs.k8s.io/yaml"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// TODO implement ValidatorI as export interface with:
//	- GetConsAddress
//	- IsJailed
//	- GetPower
//	- ...
//
//	additionally:
//	- MinEqual
//	- Equal

// TODO consider to implement collection of validators:
//	- Validators []Validator

// ============= Validator ================

func NewValidator(address sdk.ConsAddress, pubKey string, description Description, owner sdk.AccAddress) Validator {
	pkAny, err := codectypes.NewAnyWithValue(pubKey)
	if err != nil {
		return Validator{}, err
	}

	return Validator{
		Description: description,
		Address:     address.String(),
		PubKey:      pubKey,
		// Pubkey:   pkAny,
		Power:  Power,
		Jailed: false,
		Owner:  owner.String(),
	}
}

// FIXME issue 99
func (v Validator) GetConsAddress() sdk.ConsAddress {
	return sdk.ConsAddress(v.GetConsPubKey().Address())
}

// FIXME issue 99
func (v Validator) GetConsPubKey() crypto.PubKey {
	return sdk.MustGetConsPubKeyBech32(v.PubKey)
}

func (v Validator) GetPower() int32 { return v.Power }

func (v Validator) GetOwner() sdk.AccAddress {
	if v.Address == "" {
		return nil
	}
	addr, err := sdk.AccAddressFromBech32(v.Address)
	if err != nil {
		panic(err)
	}
	return addr
}

func (v Validator) GetName() string { return v.Description.Name }

func (v Validator) IsJailed() bool { return v.Jailed }

// FIXME issue 99
// ABCI ValidatorUpdate message to add new validator to validator set.
func (v Validator) ABCIValidatorUpdate() abci.ValidatorUpdate {
	/*
		tmProtoPk, err := v.TmConsPublicKey()
		if err != nil {
			panic(err)
		}
	*/

	return abci.ValidatorUpdate{
		PubKey: tmtypes.TM2PB.PubKey(v.GetConsPubKey()),
		Power:  v.GetPower(),
	}
}

// ABCI ValidatorUpdate message to remove validator from validator set.
func (v Validator) ABCIValidatorUpdateZero() abci.ValidatorUpdate {
	/*
		tmProtoPk, err := v.TmConsPublicKey()
		if err != nil {
			panic(err)
		}
	*/

	return abci.ValidatorUpdate{
		PubKey: tmtypes.TM2PB.PubKey(v.GetConsPubKey()),
		Power:  ZeroPower,
	}
}

// FIXME issue 99
func (v Validator) String() string {
	bz, err := codec.ProtoMarshalJSON(&v, nil)
	if err != nil {
		panic(err)
	}

	out, err := yaml.JSONToYAML(bz)
	if err != nil {
		panic(err)
	}

	return string(out)
}

// return the redelegation
func MustMarshalValidator(cdc codec.BinaryCodec, validator *Validator) []byte {
	return cdc.MustMarshal(validator)
}

// unmarshal a redelegation from a store value
func MustUnmarshalValidator(cdc codec.BinaryCodec, value []byte) Validator {
	validator, err := UnmarshalValidator(cdc, value)
	if err != nil {
		panic(err)
	}

	return validator
}

// unmarshal a redelegation from a store value
func UnmarshalValidator(cdc codec.BinaryCodec, value []byte) (v Validator, err error) {
	err = cdc.Unmarshal(value, &v)
	return v, err
}

/* FIXME issue 99

// ConsPubKey returns the validator PubKey as a cryptotypes.PubKey.
func (v Validator) ConsPubKey() (cryptotypes.PubKey, error) {
	pk, ok := v.ConsensusPubkey.GetCachedValue().(cryptotypes.PubKey)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "expecting cryptotypes.PubKey, got %T", pk)
	}

	return pk, nil

}

// TmConsPublicKey casts Validator.ConsensusPubkey to tmprotocrypto.PubKey.
func (v Validator) TmConsPublicKey() (tmprotocrypto.PublicKey, error) {
	pk, err := v.ConsPubKey()
	if err != nil {
		return tmprotocrypto.PublicKey{}, err
	}

	tmPk, err := cryptocodec.ToTmProtoPublicKey(pk)
	if err != nil {
		return tmprotocrypto.PublicKey{}, err
	}

	return tmPk, nil
}

// GetConsAddr extracts Consensus key address
func (v Validator) GetConsAddr() (sdk.ConsAddress, error) {
	pk, ok := v.ConsensusPubkey.GetCachedValue().(cryptotypes.PubKey)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "expecting cryptotypes.PubKey, got %T", pk)
	}

	return sdk.ConsAddress(pk.Address()), nil
}

*/

// ============= Description of Validator ================

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

// TODO issue 99 review
// String implements the Stringer interface for a Description object.
func (d Description) String() string {
	out, _ := yaml.Marshal(d)
	return string(out)
}

// ============= LastValidatorPower ================
// needed for taking validator set updates

func NewLastValidatorPower(address sdk.ConsAddress) LastValidatorPower {
	return LastValidatorPower{
		ConsensusAddress: address.String(),
		Power:            Power,
	}
}

// ============= ValidatorSigninginfo ================

func NewValidatorSigningInfo(address sdk.ConsAddress, startHeight uint64) ValidatorSigningInfo {
	return ValidatorSigningInfo{
		Address:             address.String(),
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

// TODO issue 99 review
func (v ValidatorSigningInfo) String() string {
	out, _ := yaml.Marshal(v)
	return string(out)
}

// ============= ValidatorMissedBlockBitArray ================

func NewValidatorMissedBlockBitArray(address sdk.ConsAddress, index uint64) ValidatorMissedBlockBitArray {
	return ValidatorMissedBlockBitArray{
		Address: address.String(),
		Index:   index,
	}
}

func (v ValidatorMissedBlockBitArray) String() string {
	out, _ := yaml.Marshal(v)
	return string(out)
}

// ============= ValidatorOwner ================

func NewValidatorOwner(address sdk.AccAddress, index uint64) ValidatorOwner {
	return ValidatorOwner{
		Address: address.String(),
	}
}

func (v ValidatorOwner) String() string {
	out, _ := yaml.Marshal(v)
	return string(out)
}
