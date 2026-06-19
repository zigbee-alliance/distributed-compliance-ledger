package dclauth

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/go-bip39"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

// ProposeAccountOpts holds optional flags for propose-add-account.
// VID > 0 emits --vid; PidRanges non-empty emits --pid_ranges; Info emits --info.
type ProposeAccountOpts struct {
	Info      string
	VID       int
	PidRanges string
	Extra     []string
}

func (o ProposeAccountOpts) args() []string {
	var args []string
	if o.Info != "" {
		args = append(args, "--info", o.Info)
	}
	if o.VID != 0 {
		args = append(args, "--vid", itoa(o.VID))
	}
	if o.PidRanges != "" {
		args = append(args, "--pid_ranges", o.PidRanges)
	}

	return append(args, o.Extra...)
}

// AccountActionOpts holds optional flags for approve-add-account /
// reject-add-account / propose-revoke-account / approve-revoke-account /
// reject-revoke-account. Info emits --info <value>.
type AccountActionOpts struct {
	Info  string
	Extra []string
}

func (o AccountActionOpts) args() []string {
	var args []string
	if o.Info != "" {
		args = append(args, "--info", o.Info)
	}

	return append(args, o.Extra...)
}

// ProposeAccount executes the CLI command to propose adding a new account.
func ProposeAccount(address, pubkey, roles, from string, opts ...ProposeAccountOpts) (*utils.TxResult, error) {
	args := []string{
		"tx", "auth", "propose-add-account",
		"--address", address,
		"--pubkey", pubkey,
		"--roles", roles,
		"--from", from,
	}
	for _, o := range opts {
		args = append(args, o.args()...)
	}

	return utils.ExecuteTx(args...)
}

// ApproveAccount executes the CLI command to approve adding an account.
func ApproveAccount(address, from string, opts ...AccountActionOpts) (*utils.TxResult, error) {
	args := []string{
		"tx", "auth", "approve-add-account",
		"--address", address,
		"--from", from,
	}
	for _, o := range opts {
		args = append(args, o.args()...)
	}

	return utils.ExecuteTx(args...)
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	neg := false
	if n < 0 {
		neg = true
		n = -n
	}
	var buf [20]byte
	pos := len(buf)
	for n > 0 {
		pos--
		buf[pos] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		pos--
		buf[pos] = '-'
	}

	return string(buf[pos:])
}

// QueryAccounts retrieves all accounts through the CLI.
func QueryAccounts() (*dclauthtypes.QueryAllAccountResponse, error) {
	out, err := utils.ExecuteCLI("query", "auth", "all-accounts", "-o", "json")
	if err != nil {
		return nil, err
	}
	var res dclauthtypes.QueryAllAccountResponse
	if err := json.Unmarshal(out, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// QueryAccount queries a specific account by address.
func QueryAccount(address string) (*dclauthtypes.Account, error) {
	out, err := utils.ExecuteCLI("query", "auth", "account", "--address", address, "-o", "json")
	if err != nil {
		return nil, err
	}

	// CLI generally wraps the account inside a response map or directly outputs the model.
	var res struct {
		Account dclauthtypes.Account `json:"account"`
	}
	if err := json.Unmarshal(out, &res); err != nil {
		var acc dclauthtypes.Account
		if err2 := json.Unmarshal(out, &acc); err2 == nil {
			return &acc, nil
		}

		return nil, err
	}

	return &res.Account, nil
}

// AccountIsInList is a utility to check if an address string is in the account list.
func AccountIsInList(address string, accounts []dclauthtypes.Account) bool {
	for _, acc := range accounts {
		if acc.Address == address || acc.BaseAccount.Address == address {
			return true
		}
	}

	return false
}

// CreateAccountInfo provisions a new key in the test suite keyring with a random name.
// It proxies to the original helper to avoid duplication but allows the test to call it neatly.
func CreateAccountInfo(suite *utils.TestSuite) keyring.Record {
	accountName := utils.RandString()
	entropySeed, err := bip39.NewEntropy(256)
	if err != nil {
		suite.T.Fatal(err)
	}

	mnemonic, err := bip39.NewMnemonic(entropySeed)
	if err != nil {
		suite.T.Fatal(err)
	}

	accountInfo, err := suite.Kr.NewAccount(accountName, mnemonic, testconstants.Passphrase, sdk.FullFundraiserPath, hd.Secp256k1)
	if err != nil {
		suite.T.Fatal(err)
	}

	return *accountInfo
}

// ProposeRevokeAccount executes the CLI command to propose revoking an account.
func ProposeRevokeAccount(address, from string, opts ...AccountActionOpts) (*utils.TxResult, error) {
	args := []string{
		"tx", "auth", "propose-revoke-account",
		"--address", address,
		"--from", from,
	}
	for _, o := range opts {
		args = append(args, o.args()...)
	}

	return utils.ExecuteTx(args...)
}

// ApproveRevokeAccount executes the CLI command to approve revoking an account.
func ApproveRevokeAccount(address, from string, opts ...AccountActionOpts) (*utils.TxResult, error) {
	args := []string{
		"tx", "auth", "approve-revoke-account",
		"--address", address,
		"--from", from,
	}
	for _, o := range opts {
		args = append(args, o.args()...)
	}

	return utils.ExecuteTx(args...)
}

// RejectAccount executes the CLI command to reject adding an account.
func RejectAccount(address, from string, opts ...AccountActionOpts) (*utils.TxResult, error) {
	args := []string{
		"tx", "auth", "reject-add-account",
		"--address", address,
		"--from", from,
	}
	for _, o := range opts {
		args = append(args, o.args()...)
	}

	return utils.ExecuteTx(args...)
}

// RejectRevokeAccount executes the CLI command to reject a revocation proposal.
func RejectRevokeAccount(address, from string, opts ...AccountActionOpts) (*utils.TxResult, error) {
	args := []string{
		"tx", "auth", "reject-revoke-account",
		"--address", address,
		"--from", from,
	}
	for _, o := range opts {
		args = append(args, o.args()...)
	}

	return utils.ExecuteTx(args...)
}

// QueryProposedAccount queries a proposed (pending) account by address.
func QueryProposedAccount(address string) ([]byte, error) {
	return utils.ExecuteCLI("query", "auth", "proposed-account", "--address", address, "-o", "json")
}

// QueryProposedAccountToRevoke queries a proposed-to-revoke account by address.
func QueryProposedAccountToRevoke(address string) ([]byte, error) {
	return utils.ExecuteCLI("query", "auth", "proposed-account-to-revoke", "--address", address, "-o", "json")
}

// QueryRevokedAccount queries a revoked account by address.
func QueryRevokedAccount(address string) ([]byte, error) {
	return utils.ExecuteCLI("query", "auth", "revoked-account", "--address", address, "-o", "json")
}

// QueryRejectedAccount queries a rejected account by address.
func QueryRejectedAccount(address string) ([]byte, error) {
	return utils.ExecuteCLI("query", "auth", "rejected-account", "--address", address, "-o", "json")
}

// QueryAllProposedAccounts queries all proposed (pending) accounts.
func QueryAllProposedAccounts() ([]byte, error) {
	return utils.ExecuteCLI("query", "auth", "all-proposed-accounts", "-o", "json")
}

// QueryAllProposedAccountsToRevoke queries all accounts proposed to be revoked.
func QueryAllProposedAccountsToRevoke() ([]byte, error) {
	return utils.ExecuteCLI("query", "auth", "all-proposed-accounts-to-revoke", "-o", "json")
}

// QueryAllRevokedAccounts queries all revoked accounts.
func QueryAllRevokedAccounts() ([]byte, error) {
	return utils.ExecuteCLI("query", "auth", "all-revoked-accounts", "-o", "json")
}

// QueryAllRejectedAccounts queries all rejected accounts.
func QueryAllRejectedAccounts() ([]byte, error) {
	return utils.ExecuteCLI("query", "auth", "all-rejected-accounts", "-o", "json")
}

// QueryAccountRaw queries a specific account by address and returns raw bytes.
func QueryAccountRaw(address string) ([]byte, error) {
	return utils.ExecuteCLI("query", "auth", "account", "--address", address, "-o", "json")
}

// QueryAllAccountsRaw retrieves all accounts as raw bytes.
// Uses a high limit to avoid the default 100-entry pagination cap.
func QueryAllAccountsRaw() ([]byte, error) {
	return utils.ExecuteCLI("query", "auth", "all-accounts", "-o", "json", "--limit", "10000")
}

// GetAddress returns the address string for a keyring key name.
func GetAddress(name string) (string, error) {
	out, err := utils.ExecuteCLI("keys", "show", name, "-a", "--keyring-backend", "test")
	if err != nil {
		return "", err
	}
	s := string(out)
	for len(s) > 0 && (s[len(s)-1] == '\n' || s[len(s)-1] == '\r' || s[len(s)-1] == ' ') {
		s = s[:len(s)-1]
	}

	return s, nil
}

// GetPubkey returns the pubkey string for a keyring key name.
func GetPubkey(name string) (string, error) {
	out, err := utils.ExecuteCLI("keys", "show", name, "-p", "--keyring-backend", "test")
	if err != nil {
		return "", err
	}
	s := string(out)
	for len(s) > 0 && (s[len(s)-1] == '\n' || s[len(s)-1] == '\r' || s[len(s)-1] == ' ') {
		s = s[:len(s)-1]
	}

	return s, nil
}

// AddKey generates a new key in the test keyring with the given name.
// Any pre-existing key with the same name is deleted first.
func AddKey(name string) error {
	_, _ = utils.ExecuteCLI("keys", "delete", name, "--keyring-backend", "test", "-y")
	_, err := utils.ExecuteCLI("keys", "add", name, "--keyring-backend", "test", "--no-backup")

	return err
}

// QueryAccountResponse is a lightweight wrapper to unmarshal the account query output.
func QueryAccountResponse(address string) (*dclauthtypes.Account, error) {
	out, err := utils.ExecuteCLI("query", "auth", "account", "--address", address, "-o", "json")
	if err != nil {
		return nil, err
	}
	var res struct {
		Account dclauthtypes.Account `json:"account"`
	}
	if err := json.Unmarshal(out, &res); err != nil {
		var acc dclauthtypes.Account
		if err2 := json.Unmarshal(out, &acc); err2 == nil {
			return &acc, nil
		}

		return nil, err
	}

	return &res.Account, nil
}
