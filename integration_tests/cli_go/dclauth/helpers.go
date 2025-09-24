package dclauth_test_cli

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli_go/helpers"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func ProposeAccount(
	address string,
	key string,
	role dclauthtypes.AccountRole,
	from string) (sdk.TxResponse, error) {
	return helpers.Tx(
		"auth",
		"propose-add-account",
		from,
		fmt.Sprintf("--address=%s", address),
		fmt.Sprintf("--pubkey=%s", key),
		fmt.Sprintf("--roles=%s", string(role)))
}

func ApproveAccount(
	address string,
	from string) (sdk.TxResponse, error) {
	return helpers.Tx(
		"auth",
		"approve-add-account",
		from,
		fmt.Sprintf("--address=%s", address))
}

func QueryAccounts() (dclauthtypes.QueryAllAccountResponse, error) {
	res, err := helpers.Query("auth", "all-accounts")
	if err != nil {
		return dclauthtypes.QueryAllAccountResponse{}, err
	}

	var resp dclauthtypes.QueryAllAccountResponse
	err = helpers.Codec.UnmarshalJSON([]byte(res), &resp)
	if err != nil {
		return dclauthtypes.QueryAllAccountResponse{}, err
	}

	return resp, nil
}

func QueryAccount(address string) (dclauthtypes.Account, error) {
	res, err := helpers.Query(
		"auth",
		"account",
		fmt.Sprintf("--address=%s", address))
	if err != nil {
		return dclauthtypes.Account{}, err
	}

	var resp dclauthtypes.Account
	err = helpers.Codec.UnmarshalJSON([]byte(res), &resp)
	if err != nil {
		return dclauthtypes.Account{}, err
	}

	return resp, nil
}

func QueryPendingAccounts() (dclauthtypes.QueryAllPendingAccountResponse, error) {
	res, err := helpers.Query("auth", "all-proposed-accounts")
	if err != nil {
		return dclauthtypes.QueryAllPendingAccountResponse{}, err
	}

	var resp dclauthtypes.QueryAllPendingAccountResponse
	err = helpers.Codec.UnmarshalJSON([]byte(res), &resp)
	if err != nil {
		return dclauthtypes.QueryAllPendingAccountResponse{}, err
	}

	return resp, nil
}

func AccountIsInList(expected string, accounts []dclauthtypes.Account) bool {
	for _, account := range accounts {
		if expected == account.Address {
			return true
		}
	}

	return false
}
