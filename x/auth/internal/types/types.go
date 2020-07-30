package types

//nolint:goimports
import (
	"encoding/json"
	"fmt"
	"github.com/tendermint/tendermint/crypto"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

/*
	Roles assigning to Account
*/

type AccountRole string

const (
	Vendor                AccountRole = "Vendor"
	TestHouse             AccountRole = "TestHouse"
	ZBCertificationCenter AccountRole = "ZBCertificationCenter"
	Trustee               AccountRole = "Trustee"
	NodeAdmin             AccountRole = "NodeAdmin"
)

var Roles = AccountRoles{Vendor, TestHouse, ZBCertificationCenter, Trustee, NodeAdmin}

func (lt AccountRole) Validate() sdk.Error {
	switch lt {
	case Vendor, TestHouse, ZBCertificationCenter, Trustee, NodeAdmin:
		return nil
	}

	return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Account Role: \"%v\". Supported roles: [%v]", lt, Roles))
}

/*
	List of Roles
*/

type AccountRoles []AccountRole

// Validate checks for errors on the account roles.
func (acc AccountRoles) Validate() sdk.Error {
	for _, role := range acc {
		if err := role.Validate(); err != nil {
			return err
		}
	}

	return nil
}

/*
	Pending Account
*/
type PendingAccount struct {
	Address   sdk.AccAddress   `json:"address"`
	PubKey    crypto.PubKey    `json:"public_key"`
	Roles     AccountRoles     `json:"roles"`
	Approvals []sdk.AccAddress `json:"approvals"`
}

// NewPendingAccount creates a new PendingAccount object
func NewPendingAccount(address sdk.AccAddress, pubKey crypto.PubKey,
	roles AccountRoles, approval sdk.AccAddress) PendingAccount {
	return PendingAccount{
		Address:   address,
		PubKey:    pubKey,
		Roles:     roles,
		Approvals: []sdk.AccAddress{approval},
	}
}

// String implements fmt.Stringer
func (acc PendingAccount) String() string {
	bytes, err := json.Marshal(acc)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

// Validate checks for errors on the vesting and module account parameters.
func (acc PendingAccount) Validate() error {
	if acc.Address == nil {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf("Invalid Account: Value: %s. Error: Missing Address", acc.Address))
	}

	if acc.PubKey == nil {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf("Invalid Account: Value: %s. Error: Missing PubKey", acc.PubKey))
	}

	if err := acc.Roles.Validate(); err != nil {
		return err
	}

	return nil
}

func (acc PendingAccount) HasApprovalFrom(address sdk.AccAddress) bool {
	for _, approval := range acc.Approvals {
		if approval.Equals(address) {
			return true
		}
	}
	return false
}

/*
	Account
*/
type Account struct {
	Address       sdk.AccAddress `json:"address"`
	PubKey        crypto.PubKey  `json:"public_key"`
	AccountNumber uint64         `json:"account_number"`
	Sequence      uint64         `json:"sequence"`
	Roles         AccountRoles   `json:"roles"`
}

// NewAccount creates a new Account object.
func NewAccount(address sdk.AccAddress, pubKey crypto.PubKey, roles AccountRoles) Account {
	return Account{
		Address: address,
		PubKey:  pubKey,
		Roles:   roles,
	}
}

// String implements fmt.Stringer.
func (acc Account) String() string {
	bytes, err := json.Marshal(acc)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

// Validate checks for errors on the vesting and module account parameters.
func (acc Account) Validate() error {
	if acc.Address == nil {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf("Invalid Accounts: Value: %s. Error: Missing Address", acc.Address))
	}

	if acc.PubKey == nil {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf("Invalid Accounts: Value: %s. Error: Missing PubKey", acc.PubKey))
	}

	if err := acc.Roles.Validate(); err != nil {
		return err
	}

	return nil
}

func (acc Account) HasRole(targetRole AccountRole) bool {
	for _, role := range acc.Roles {
		if role == targetRole {
			return true
		}
	}
	return false
}

func (acc Account) GetAddress() sdk.AccAddress {
	return acc.Address
}

// SetAddress - Implements sdk.Account.
func (acc *Account) SetAddress(addr sdk.AccAddress) error {
	if len(acc.Address) != 0 {
		return sdk.ErrInvalidAddress("Cannot override Account address")
	}

	acc.Address = addr

	return nil
}

// GetPubKey - Implements sdk.Account.
func (acc Account) GetPubKey() crypto.PubKey {
	return acc.PubKey
}

// SetPubKey - Implements sdk.Account.
func (acc *Account) SetPubKey(pubKey crypto.PubKey) error {
	acc.PubKey = pubKey
	return nil
}

// GetCoins - Implements sdk.Account.
func (acc *Account) GetCoins() sdk.Coins {
	return nil
}

// SetCoins - Implements sdk.Account.
func (acc *Account) SetCoins(coins sdk.Coins) error {
	return nil
}

// GetAccountNumber - Implements Account.
func (acc *Account) GetAccountNumber() uint64 {
	return acc.AccountNumber
}

// SetAccountNumber - Implements Account.
func (acc *Account) SetAccountNumber(accNumber uint64) error {
	acc.AccountNumber = accNumber
	return nil
}

// GetSequence - Implements sdk.Account.
func (acc *Account) GetSequence() uint64 {
	return acc.Sequence
}

// SetSequence - Implements sdk.Account.
func (acc *Account) SetSequence(seq uint64) error {
	acc.Sequence = seq
	return nil
}

// SpendableCoins returns the total set of spendable coins. For a base account,
// this is simply the base coins.
func (acc *Account) SpendableCoins(_ time.Time) sdk.Coins {
	return nil
}
