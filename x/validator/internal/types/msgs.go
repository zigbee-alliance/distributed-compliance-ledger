package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/tendermint/tendermint/crypto"
)

type MsgCreateValidator struct {
	ValidatorAddress sdk.ValAddress           `json:"validator_address"`
	PubKey           string                   `json:"pubkey"`
	Description      stakingtypes.Description `json:"description"`
}

func NewMsgCreateValidator(valAddr sdk.ValAddress, pubKey string, description stakingtypes.Description) MsgCreateValidator {
	return MsgCreateValidator{
		ValidatorAddress: valAddr,
		PubKey:           pubKey,
		Description:      description,
	}
}

func (m MsgCreateValidator) Route() string { return RouterKey }

func (m MsgCreateValidator) Type() string { return EventTypeCreateValidator }

func (m MsgCreateValidator) ValidateBasic() sdk.Error {
	if m.ValidatorAddress.Empty() {
		return sdk.ErrUnknownRequest("Invalid Validator OperatorAddress: it cannot be empty")
	}
	if _, err := sdk.GetConsPubKeyBech32(m.PubKey); err != nil {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Validator Public Key: %v", err))
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
