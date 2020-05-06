package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/tendermint/tendermint/crypto"
)

type MsgCreateValidator struct {
	Description      stakingtypes.Description `json:"description"`
	ValidatorAddress sdk.ValAddress           `json:"validator_address"`
	PubKey           string                   `json:"pubkey"` // MustBech32ifyConsPub
}

func NewMsgCreateValidator(valAddr sdk.ValAddress, pubKey string, description stakingtypes.Description) MsgCreateValidator {
	return MsgCreateValidator{
		Description:      description,
		ValidatorAddress: valAddr,
		PubKey:           pubKey,
	}
}

func (m MsgCreateValidator) Route() string { return RouterKey }

func (m MsgCreateValidator) Type() string { return "create_validator" }

func (m MsgCreateValidator) ValidateBasic() sdk.Error {
	if m.ValidatorAddress.Empty() {
		return sdk.ErrUnknownRequest("Invalid Validator OperatorAddress: it cannot be empty")
	}
	if m.Description == (stakingtypes.Description{}) {
		return sdk.ErrUnknownRequest("Invalid Description: it cannot be empty")
	}
	if _, err := m.Description.EnsureLength(); err != nil {
		return err
	}
	return nil
}

func (m MsgCreateValidator) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(m.ValidatorAddress)}
}

func (m MsgCreateValidator) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgCreateValidator) GetPubKey() crypto.PubKey {
	return sdk.MustGetConsPubKeyBech32(m.PubKey)
}
