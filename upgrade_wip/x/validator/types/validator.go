package types

import (
	"fmt"

	abci "github.com/tendermint/tendermint/abci/types"
	"sigs.k8s.io/yaml"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
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

func NewValidator(owner sdk.ValAddress, pubKey cryptotypes.PubKey, description Description) Validator {
	pkAny, err := codectypes.NewAnyWithValue(pubKey)
	if err != nil {
		return Validator{}, err
	}

	return Validator{
		Owner:       owner.String(),
		Description: description,
		Pubkey:      pkAny,
		Power:       Power,
		Jailed:      false,
	}
}

func (v Validator) GetConsAddress() (sdk.ConsAddress, error) {
	pk, ok := v.Pubkey.GetCachedValue().(cryptotypes.PubKey)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "expecting cryptotypes.PubKey, got %T", pk)
	}

	return sdk.ConsAddress(pk.Address()), nil
}

func (v Validator) GetConsPubKey() (cryptotypes.PubKey, error) {
	pk, ok := v.Pubkey.GetCachedValue().(cryptotypes.PubKey)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "expecting cryptotypes.PubKey, got %T", pk)
	}

	return pk, nil

}

func (v Validator) GetPower() int32 { return v.Power }

func (v Validator) GetOwner() sdk.ValAddress {
	if v.Owner == "" {
		return nil
	}
	addr, err := sdk.ValAddressFromBech32(v.Owner)
	if err != nil {
		panic(err)
	}
	return addr
}

func (v Validator) GetName() string { return v.Description.Name }

func (v Validator) IsJailed() bool { return v.Jailed }

// ABCI ValidatorUpdate message to add new validator to validator set.
func (v Validator) ABCIValidatorUpdate() abci.ValidatorUpdate {
	tmProtoPk, err := v.TmConsPublicKey()
	if err != nil {
		panic(err)
	}

	return abci.ValidatorUpdate{
		PubKey: tmProtoPk,
		Power:  v.GetPower(),
	}
}

// ABCI ValidatorUpdate message to remove validator from validator set.
func (v Validator) ABCIValidatorUpdateZero() abci.ValidatorUpdate {
	tmProtoPk, err := v.TmConsPublicKey()
	if err != nil {
		panic(err)
	}

	return abci.ValidatorUpdate{
		PubKey: tmProtoPk,
		Power:  ZeroPower,
	}
}

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

// TmConsPublicKey casts Validator.Pubkey to tmprotocrypto.PubKey.
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
	pk, ok := v.Pubkey.GetCachedValue().(cryptotypes.PubKey)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "expecting cryptotypes.PubKey, got %T", pk)
	}

	return sdk.ConsAddress(pk.Address()), nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (v Validator) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var pk cryptotypes.PubKey
	return unpacker.UnpackAny(v.Pubkey, &pk)
}

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

func (d Description) String() string {
	out, _ := yaml.Marshal(d)
	return string(out)
}

// ============= LastValidatorPower ================
// needed for taking validator set updates

func NewLastValidatorPower(owner sdk.ValAddress) LastValidatorPower {
	return LastValidatorPower{
		Owner: address.String(),
		Power: Power,
	}
}

func (vp LastValidatorPower) GetOwner() sdk.ValAddress {
	if vp.Owner == "" {
		return nil
	}
	addr, err := sdk.ValAddressFromBech32(vp.Owner)
	if err != nil {
		panic(err)
	}
	return addr
}

func (vp LastValidatorPower) GetPower() Power { return vp.Power }

func (vp LastValidatorPower) String() string {
	out, _ := yaml.Marshal(vp)
	return string(out)
}
