package dclauth

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/go-bip39"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

// ProposeAccountOpts holds optional flags for propose-add-account.
// VID > 0 emits --vid (or VIDHex for a hex-formatted value such as "0xA13");
// PidRanges non-empty emits --pid_ranges; Info emits --info.
type ProposeAccountOpts struct {
	Info      string
	VID       int
	VIDHex    string
	PidRanges string
}

func (o ProposeAccountOpts) args() []string {
	var args []string
	if o.Info != "" {
		args = append(args, "--info", o.Info)
	}
	if o.VIDHex != "" {
		args = append(args, "--vid", o.VIDHex)
	} else if o.VID != 0 {
		args = append(args, "--vid", strconv.Itoa(o.VID))
	}
	if o.PidRanges != "" {
		args = append(args, "--pid_ranges", o.PidRanges)
	}

	return args
}

// AccountActionOpts holds optional flags for approve-add-account /
// reject-add-account / propose-revoke-account / approve-revoke-account /
// reject-revoke-account. Info emits --info <value>.
type AccountActionOpts struct {
	Info string
}

func (o AccountActionOpts) args() []string {
	var args []string
	if o.Info != "" {
		args = append(args, "--info", o.Info)
	}

	return args
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

// GetAllAccounts retrieves all accounts through the CLI. The high pagination
// limit avoids the default 100-entry cap.
func GetAllAccounts() ([]dclauthtypes.Account, error) {
	var res dclauthtypes.QueryAllAccountResponse
	if err := cliputils.GetList(&res, "query", "auth", "all-accounts", "-o", "json", "--limit", "10000"); err != nil {
		return nil, err
	}

	return res.Account, nil
}

// GetAccount queries a specific account by address. Returns nil when the
// account does not exist.
func GetAccount(address string) (*dclauthtypes.Account, error) {
	out, err := utils.ExecuteCLI("query", "auth", "account", "--address", address, "-o", "json")
	if err != nil {
		return nil, err
	}
	if utils.IsNotFound(out) {
		return nil, nil //nolint:nilnil // (nil, nil) marks "no record" — established Get* pattern
	}
	out = utils.NormalizeProtoJSON(out)

	// The CLI either wraps the account inside an "account" key or emits the
	// account directly. Try the wrapped shape first; fall back to direct.
	// Account.Address is promoted from the embedded *BaseAccount, so reading
	// it panics when BaseAccount is nil — guard against that here.
	var res struct {
		Account dclauthtypes.Account `json:"account"`
	}
	if err := json.Unmarshal(out, &res); err == nil &&
		res.Account.BaseAccount != nil && res.Account.Address != "" {
		return &res.Account, nil
	}
	var acc dclauthtypes.Account
	if err := json.Unmarshal(out, &acc); err != nil {
		return nil, fmt.Errorf("parse Account: %w, output: %s", err, string(out))
	}
	if acc.BaseAccount == nil {
		return nil, fmt.Errorf("parse Account: missing base_account, output: %s", string(out))
	}

	return &acc, nil
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

// GetProposedAccount queries a proposed (pending) account by address. Returns
// nil when the record does not exist.
func GetProposedAccount(address string) (*dclauthtypes.PendingAccount, error) {
	var res dclauthtypes.PendingAccount
	found, err := cliputils.GetSingle(&res,
		"query", "auth", "proposed-account",
		"--address", address,
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetProposedAccountToRevoke queries a proposed-to-revoke account by address.
// Returns nil when the record does not exist.
func GetProposedAccountToRevoke(address string) (*dclauthtypes.PendingAccountRevocation, error) {
	var res dclauthtypes.PendingAccountRevocation
	found, err := cliputils.GetSingle(&res,
		"query", "auth", "proposed-account-to-revoke",
		"--address", address,
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetRevokedAccount queries a revoked account by address. Returns nil when the
// record does not exist.
func GetRevokedAccount(address string) (*dclauthtypes.RevokedAccount, error) {
	var res dclauthtypes.RevokedAccount
	found, err := cliputils.GetSingle(&res,
		"query", "auth", "revoked-account",
		"--address", address,
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetRejectedAccount queries a rejected account by address. Returns nil when
// the record does not exist.
func GetRejectedAccount(address string) (*dclauthtypes.RejectedAccount, error) {
	var res dclauthtypes.RejectedAccount
	found, err := cliputils.GetSingle(&res,
		"query", "auth", "rejected-account",
		"--address", address,
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetAllProposedAccounts queries all proposed (pending) accounts.
func GetAllProposedAccounts() ([]dclauthtypes.PendingAccount, error) {
	var res dclauthtypes.QueryAllPendingAccountResponse
	if err := cliputils.GetList(&res, "query", "auth", "all-proposed-accounts", "-o", "json"); err != nil {
		return nil, err
	}

	return res.PendingAccount, nil
}

// GetAllProposedAccountsToRevoke queries all accounts proposed to be revoked.
func GetAllProposedAccountsToRevoke() ([]dclauthtypes.PendingAccountRevocation, error) {
	var res dclauthtypes.QueryAllPendingAccountRevocationResponse
	if err := cliputils.GetList(&res, "query", "auth", "all-proposed-accounts-to-revoke", "-o", "json"); err != nil {
		return nil, err
	}

	return res.PendingAccountRevocation, nil
}

// GetAllRevokedAccounts queries all revoked accounts.
func GetAllRevokedAccounts() ([]dclauthtypes.RevokedAccount, error) {
	var res dclauthtypes.QueryAllRevokedAccountResponse
	if err := cliputils.GetList(&res, "query", "auth", "all-revoked-accounts", "-o", "json"); err != nil {
		return nil, err
	}

	return res.RevokedAccount, nil
}

// GetAllRejectedAccounts queries all rejected accounts.
func GetAllRejectedAccounts() ([]dclauthtypes.RejectedAccount, error) {
	var res dclauthtypes.QueryAllRejectedAccountResponse
	if err := cliputils.GetList(&res, "query", "auth", "all-rejected-accounts", "-o", "json"); err != nil {
		return nil, err
	}

	return res.RejectedAccount, nil
}

// containsAccountAddress reports whether list contains an account with the
// given address (either at the embedded BaseAccount level or — for accounts
// returned by the CLI without a Cosmos BaseAccount — at the top-level Address).
func containsAccountAddress(list []dclauthtypes.Account, address string) bool {
	for i := range list {
		if list[i].Address == address {
			return true
		}
		if list[i].BaseAccount != nil && list[i].BaseAccount.Address == address {
			return true
		}
	}

	return false
}

// containsPendingAccountAddress reports whether list has a PendingAccount with
// the given address.
func containsPendingAccountAddress(list []dclauthtypes.PendingAccount, address string) bool {
	for i := range list {
		if list[i].Account != nil && list[i].Account.Address == address {
			return true
		}
	}

	return false
}

// containsPendingAccountRevocationAddress reports the same for
// PendingAccountRevocation.
func containsPendingAccountRevocationAddress(list []dclauthtypes.PendingAccountRevocation, address string) bool {
	for i := range list {
		if list[i].Address == address {
			return true
		}
	}

	return false
}

// containsRevokedAccountAddress reports the same for RevokedAccount.
func containsRevokedAccountAddress(list []dclauthtypes.RevokedAccount, address string) bool {
	for i := range list {
		if list[i].Account != nil && list[i].Account.Address == address {
			return true
		}
	}

	return false
}

// containsRejectedAccountAddress reports the same for RejectedAccount.
func containsRejectedAccountAddress(list []dclauthtypes.RejectedAccount, address string) bool {
	for i := range list {
		if list[i].Account != nil && list[i].Account.Address == address {
			return true
		}
	}

	return false
}
