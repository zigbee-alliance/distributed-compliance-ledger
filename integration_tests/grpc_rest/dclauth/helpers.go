// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dclauth

import (
	"context"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	"github.com/stretchr/testify/require"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"

	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

const (
	DCLAuthAccountsEndpoint                   = "/dcl/auth/accounts/"
	DCLAuthProposedAccountsEndpoint           = "/dcl/auth/proposed-accounts"
	DCLAuthProposedRevocationAccountsEndpoint = "/dcl/auth/proposed-revocation-accounts"
)

//nolint:godox
/*
	To Run test you need:
		* Run LocalNet with: `make install && make localnet_init && make localnet_start`

	TODO: provide tests for error cases
*/

func GetAccount(suite *utils.TestSuite, address sdk.AccAddress) (*dclauthtypes.Account, error) {
	var res dclauthtypes.Account

	if suite.Rest {
		// TODO issue 99: explore the way how to get the endpoint from proto-
		//      instead of the hard coded value (the same for all rest queries)
		var resp dclauthtypes.QueryGetAccountResponse
		err := suite.QueryREST(DCLAuthAccountsEndpoint+address.String(), &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetAccount()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/dclauth service.
		accClient := dclauthtypes.NewQueryClient(grpcConn)
		resp, err := accClient.Account(
			context.Background(),
			&dclauthtypes.QueryGetAccountRequest{Address: address.String()},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetAccount()
	}

	return &res, nil

}

func GetAccounts(suite *utils.TestSuite) (res []dclauthtypes.Account, err error) {
	if suite.Rest {
		// TODO issue 99: explore the way how to get the endpoint from proto-
		//      instead of the hard coded value (the same for all rest queries)
		var resp dclauthtypes.QueryAllAccountResponse
		err := suite.QueryREST(DCLAuthAccountsEndpoint, &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetAccount()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/dclauth service.
		accClient := dclauthtypes.NewQueryClient(grpcConn)
		resp, err := accClient.AccountAll(
			context.Background(),
			&dclauthtypes.QueryAllAccountRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetAccount()
	}

	return res, nil
}

func GetProposedAccounts(suite *utils.TestSuite) (res []dclauthtypes.PendingAccount, err error) {
	if suite.Rest {
		// TODO issue 99: explore the way how to get the endpoint from proto-
		//      instead of the hard coded value (the same for all rest queries)
		var resp dclauthtypes.QueryAllPendingAccountResponse
		err := suite.QueryREST(DCLAuthProposedAccountsEndpoint, &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetPendingAccount()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/dclauth service.
		accClient := dclauthtypes.NewQueryClient(grpcConn)
		resp, err := accClient.PendingAccountAll(
			context.Background(),
			&dclauthtypes.QueryAllPendingAccountRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetPendingAccount()
	}

	return res, nil
}

func GetProposedAccountsToRevoke(suite *utils.TestSuite) (
	res []dclauthtypes.PendingAccountRevocation, err error,
) {
	if suite.Rest {
		// TODO issue 99: explore the way how to get the endpoint from proto-
		//      instead of the hard coded value (the same for all rest queries)
		var resp dclauthtypes.QueryAllPendingAccountRevocationResponse
		err := suite.QueryREST(DCLAuthProposedRevocationAccountsEndpoint, &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetPendingAccountRevocation()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/dclauth service.
		accClient := dclauthtypes.NewQueryClient(grpcConn)
		resp, err := accClient.PendingAccountRevocationAll(
			context.Background(),
			&dclauthtypes.QueryAllPendingAccountRevocationRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetPendingAccountRevocation()
	}

	return res, nil
}

// TODO issue 99: add support for query accounts stat

func ProposeAddAccount(
	suite *utils.TestSuite,
	signer string,
	accAddr sdk.AccAddress, accKey cryptotypes.PubKey,
	roles dclauthtypes.AccountRoles, vendorID uint64,
	accnum uint64, sequence uint64,
) (*sdk.TxResponse, error) {
	msg, err := dclauthtypes.NewMsgProposeAddAccount(
		suite.GetAddress(signer), accAddr, accKey, roles, vendorID)
	require.NoError(suite.T, err)
	return suite.BuildAndBroadcastTx(signer, []sdk.Msg{msg}, accnum, sequence)
}

func ApproveAddAccount(
	suite *utils.TestSuite,
	signer string, accAddr sdk.AccAddress, accnum uint64, sequence uint64,
) (*sdk.TxResponse, error) {
	msg := dclauthtypes.NewMsgApproveAddAccount(suite.GetAddress(signer), accAddr)
	return suite.BuildAndBroadcastTx(signer, []sdk.Msg{msg}, accnum, sequence)
}

func ProposeRevokeAccount(
	suite *utils.TestSuite,
	signer string, accAddr sdk.AccAddress, accnum uint64, sequence uint64,
) (*sdk.TxResponse, error) {
	msg := dclauthtypes.NewMsgProposeRevokeAccount(suite.GetAddress(signer), accAddr)
	return suite.BuildAndBroadcastTx(signer, []sdk.Msg{msg}, accnum, sequence)
}

func ApproveRevokeAccount(
	suite *utils.TestSuite,
	signer string, accAddr sdk.AccAddress, accnum uint64, sequence uint64,
) (*sdk.TxResponse, error) {
	msg := dclauthtypes.NewMsgApproveRevokeAccount(suite.GetAddress(signer), accAddr)
	return suite.BuildAndBroadcastTx(signer, []sdk.Msg{msg}, accnum, sequence)
}

// Common Test Logic

//nolint:funlen
func AuthDemo(suite *utils.TestSuite) {
	jackName := testconstants.JackAccount
	aliceName := testconstants.AliceAccount
	bobName := testconstants.BobAccount

	// Query all active accounts
	inputAccounts, err := GetAccounts(suite)
	require.NoError(suite.T, err)
	require.Equal(suite.T, 4, len(inputAccounts))

	// build map with an acc address as a key
	accDataInitial := make(map[string]dclauthtypes.Account)
	for _, acc := range inputAccounts {
		accDataInitial[acc.GetAddress().String()] = acc
	}

	// Jack, Alice and Bob are predefined Trustees
	jackKeyInfo, err := suite.Kr.Key(jackName)
	require.NoError(suite.T, err)
	require.Contains(suite.T, accDataInitial, jackKeyInfo.GetAddress().String())
	jackSequence := accDataInitial[jackKeyInfo.GetAddress().String()].GetSequence()
	jackAccNum := accDataInitial[jackKeyInfo.GetAddress().String()].GetAccountNumber()

	aliceKeyInfo, err := suite.Kr.Key(aliceName)
	require.NoError(suite.T, err)
	require.Contains(suite.T, accDataInitial, aliceKeyInfo.GetAddress().String())
	aliceSequence := accDataInitial[aliceKeyInfo.GetAddress().String()].GetSequence()
	aliceAccNum := accDataInitial[aliceKeyInfo.GetAddress().String()].GetAccountNumber()

	bobKeyInfo, err := suite.Kr.Key(bobName)
	require.NoError(suite.T, err)
	require.Contains(suite.T, accDataInitial, bobKeyInfo.GetAddress().String())
	bobSequence := accDataInitial[bobKeyInfo.GetAddress().String()].GetSequence()
	bobAccNum := accDataInitial[bobKeyInfo.GetAddress().String()].GetAccountNumber()

	// Query all proposed accounts
	inputProposedAccounts, err := GetProposedAccounts(suite)
	require.NoError(suite.T, err)
	require.Equal(suite.T, 0, len(inputProposedAccounts))

	// Query all proposed accounts to revoke
	inputProposedAccountsToRevoke, err := GetProposedAccountsToRevoke(suite)
	require.NoError(suite.T, err)
	require.Equal(suite.T, 0, len(inputProposedAccountsToRevoke))

	_, testAccPubKey, testAccAddr := testdata.KeyTestPubAddr()

	// Jack proposes new account
	_, err = ProposeAddAccount(
		suite,
		jackName, testAccAddr, testAccPubKey,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor}, testconstants.VID,
		jackAccNum, jackSequence,
	)
	require.NoError(suite.T, err)
	jackSequence += 1

	// Query all active accounts
	receivedAccounts, _ := GetAccounts(suite)
	require.Equal(suite.T, len(inputAccounts), len(receivedAccounts))

	// Query all proposed accounts
	receivedProposedAccounts, _ := GetProposedAccounts(suite)
	require.Equal(suite.T, len(inputProposedAccounts)+1, len(receivedProposedAccounts))

	// Query all accounts proposed to be revoked
	receivedProposedAccountsToRevoke, _ := GetProposedAccountsToRevoke(suite)
	require.Equal(suite.T, len(inputProposedAccountsToRevoke), len(receivedProposedAccountsToRevoke))

	// Alice approves new account
	_, err = ApproveAddAccount(suite, aliceName, testAccAddr, aliceAccNum, aliceSequence)
	require.NoError(suite.T, err)
	aliceSequence += 1

	// Query all active accounts
	receivedAccounts, _ = GetAccounts(suite)
	require.Equal(suite.T, len(inputAccounts)+1, len(receivedAccounts))

	// Query all proposed accounts
	receivedProposedAccounts, _ = GetProposedAccounts(suite)
	require.Equal(suite.T, len(inputProposedAccounts), len(receivedProposedAccounts))

	// Query all accounts proposed to be revoked
	receivedProposedAccountsToRevoke, _ = GetProposedAccountsToRevoke(suite)
	require.Equal(suite.T, len(inputProposedAccountsToRevoke), len(receivedProposedAccountsToRevoke))

	// Get info for new account
	testAccount, err := GetAccount(suite, testAccAddr)
	require.NoError(suite.T, err)
	require.Equal(suite.T, testAccAddr, testAccount.GetAddress())
	require.Equal(suite.T, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testAccount.GetRoles())

	// FIXME issue 99: enable once implemented
	/*
		// Publish model info by test account
		model := NewMsgAddModel(suite, testAccountKeyInfo.Address, testconstants.VID)
		_, _ = AddModel(suite, model, testAccountKeyInfo)

		// Check model is created
		receivedModel, _ := GetModel(suite, model.VID, model.PID)
		require.Equal(suite.T, receivedModel.VID, model.VID)
		require.Equal(suite.T, receivedModel.PID, model.PID)
		require.Equal(suite.T, receivedModel.ProductName, model.ProductName)
	*/

	// Alice proposes to revoke new account
	_, err = ProposeRevokeAccount(suite, aliceName, testAccAddr, aliceAccNum, aliceSequence)
	require.NoError(suite.T, err)
	aliceSequence += 1

	// Query all active accounts
	receivedAccounts, err = GetAccounts(suite)
	require.Equal(suite.T, len(inputAccounts)+1, len(receivedAccounts))

	// Query all proposed accounts
	receivedProposedAccounts, _ = GetProposedAccounts(suite)
	require.Equal(suite.T, len(inputProposedAccounts), len(receivedProposedAccounts))

	// Query all accounts proposed to be revoked
	receivedProposedAccountsToRevoke, _ = GetProposedAccountsToRevoke(suite)
	require.Equal(suite.T, len(inputProposedAccountsToRevoke)+1, len(receivedProposedAccountsToRevoke))

	// Bob approves to revoke new account
	_, err = ApproveRevokeAccount(suite, bobName, testAccAddr, bobAccNum, bobSequence)
	require.NoError(suite.T, err)
	bobSequence += 1

	// Query all active accounts
	receivedAccounts, err = GetAccounts(suite)
	require.Equal(suite.T, len(inputAccounts), len(receivedAccounts))

	// Query all proposed accounts
	receivedProposedAccounts, _ = GetProposedAccounts(suite)
	require.Equal(suite.T, len(inputProposedAccounts), len(receivedProposedAccounts))

	// Query all accounts proposed to be revoked
	receivedProposedAccountsToRevoke, _ = GetProposedAccountsToRevoke(suite)
	require.Equal(suite.T, len(inputProposedAccountsToRevoke), len(receivedProposedAccountsToRevoke))

	// Ensure that new account is not available anymore
	_, err = GetAccount(suite, testAccAddr)
	require.Error(suite.T, err)
	require.Contains(suite.T, err.Error(), "rpc error: code = InvalidArgument desc = not found: invalid request")

	// FIXME issue 99: enable once implemented
	/*
		// Try to publish another model info by test account.
		// Ensure that the request is responded with not OK status code.
		model = NewMsgAddModel(suite, testAccountKeyInfo.Address, testconstants.VID)
		_, code = AddModel(suite, model, testAccountKeyInfo)
		require.NotEqual(suite.T, http.StatusOK, code)
	*/
}
