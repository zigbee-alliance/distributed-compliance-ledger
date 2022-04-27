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

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/go-bip39"
	"github.com/stretchr/testify/require"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	modeltypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

const (
	DCLAuthAccountsEndpoint                   = "/dcl/auth/accounts/"
	DCLAuthProposedAccountsEndpoint           = "/dcl/auth/proposed-accounts/"
	DCLAuthProposedRevocationAccountsEndpoint = "/dcl/auth/proposed-revocation-accounts/"
	DCLAuthRevokedAccountsEndpoint            = "/dcl/auth/revoked-accounts/"
	DCLAuthRejectedAccountsEndpoint           = "/dcl/auth/rejected-accounts/"
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

func GetProposedAccount(suite *utils.TestSuite, address sdk.AccAddress) (*dclauthtypes.PendingAccount, error) {
	var res dclauthtypes.PendingAccount

	if suite.Rest {
		var resp dclauthtypes.QueryGetPendingAccountResponse
		err := suite.QueryREST(DCLAuthProposedAccountsEndpoint+address.String(), &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetPendingAccount()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/dclauth service.
		accClient := dclauthtypes.NewQueryClient(grpcConn)
		resp, err := accClient.PendingAccount(
			context.Background(),
			&dclauthtypes.QueryGetPendingAccountRequest{Address: address.String()},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetPendingAccount()
	}

	return &res, nil
}

func GetProposedAccountsToRevoke(suite *utils.TestSuite) (
	res []dclauthtypes.PendingAccountRevocation, err error,
) {
	if suite.Rest {
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

func GetProposedAccountToRevoke(suite *utils.TestSuite, address sdk.AccAddress) (*dclauthtypes.PendingAccountRevocation, error) {
	var res dclauthtypes.PendingAccountRevocation

	if suite.Rest {
		var resp dclauthtypes.QueryGetPendingAccountRevocationResponse
		err := suite.QueryREST(DCLAuthProposedRevocationAccountsEndpoint+address.String(), &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetPendingAccountRevocation()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/dclauth service.
		accClient := dclauthtypes.NewQueryClient(grpcConn)
		resp, err := accClient.PendingAccountRevocation(
			context.Background(),
			&dclauthtypes.QueryGetPendingAccountRevocationRequest{Address: address.String()},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetPendingAccountRevocation()
	}

	return &res, nil
}

func GetRevokedAccount(suite *utils.TestSuite, address sdk.AccAddress) (*dclauthtypes.RevokedAccount, error) {
	var res dclauthtypes.RevokedAccount

	if suite.Rest {
		var resp dclauthtypes.QueryGetRevokedAccountResponse
		err := suite.QueryREST(DCLAuthRevokedAccountsEndpoint+address.String(), &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetRevokedAccount()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC cliennt to query the x/dclauth service.
		accClient := dclauthtypes.NewQueryClient(grpcConn)
		resp, err := accClient.RevokedAccount(
			context.Background(),
			&dclauthtypes.QueryGetRevokedAccountRequest{Address: address.String()},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetRevokedAccount()
	}

	return &res, nil
}

func GetRevokedAccounts(suite *utils.TestSuite) (res []dclauthtypes.RevokedAccount, err error) {
	if suite.Rest {
		var resp dclauthtypes.QueryAllRevokedAccountResponse
		err := suite.QueryREST(DCLAuthRevokedAccountsEndpoint, &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetRevokedAccount()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/dclauth service.
		accClient := dclauthtypes.NewQueryClient(grpcConn)
		resp, err := accClient.RevokedAccountAll(
			context.Background(),
			&dclauthtypes.QueryAllRevokedAccountRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetRevokedAccount()
	}

	return res, nil
}

func GetRejectedAccount(suite *utils.TestSuite, address sdk.AccAddress) (*dclauthtypes.RejectedAccount, error) {
	var res dclauthtypes.RejectedAccount

	if suite.Rest {
		var resp dclauthtypes.QueryGetRejectedAccountResponse
		err := suite.QueryREST(DCLAuthRejectedAccountsEndpoint+address.String(), &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetRejectedAccount()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC cliennt to query the x/dclauth service.
		accClient := dclauthtypes.NewQueryClient(grpcConn)
		resp, err := accClient.RejectedAccount(
			context.Background(),
			&dclauthtypes.QueryGetRejectedAccountRequest{Address: address.String()},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetRejectedAccount()
	}

	return &res, nil
}

func GetRejectedAccounts(suite *utils.TestSuite) (res []dclauthtypes.RejectedAccount, err error) {
	if suite.Rest {
		var resp dclauthtypes.QueryAllRejectedAccountResponse
		err := suite.QueryREST(DCLAuthRejectedAccountsEndpoint, &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetRejectedAccount()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/dclauth service.
		accClient := dclauthtypes.NewQueryClient(grpcConn)
		resp, err := accClient.RejectedAccountAll(
			context.Background(),
			&dclauthtypes.QueryAllRejectedAccountRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetRejectedAccount()
	}

	return res, nil
}

func ProposeAddAccount(
	suite *utils.TestSuite,
	accAddr sdk.AccAddress,
	accKey cryptotypes.PubKey,
	roles dclauthtypes.AccountRoles,
	vendorID int32,
	signerName string,
	signerAccount *dclauthtypes.Account,
	info string,
) (*sdk.TxResponse, error) {
	msg, err := dclauthtypes.NewMsgProposeAddAccount(
		suite.GetAddress(signerName), accAddr, accKey, roles, vendorID, info)
	require.NoError(suite.T, err)

	return suite.BuildAndBroadcastTx([]sdk.Msg{msg}, signerName, signerAccount)
}

func ApproveAddAccount(
	suite *utils.TestSuite,
	accAddr sdk.AccAddress,
	signerName string,
	signerAccount *dclauthtypes.Account,
	info string,
) (*sdk.TxResponse, error) {
	msg := dclauthtypes.NewMsgApproveAddAccount(suite.GetAddress(signerName), accAddr, info)

	return suite.BuildAndBroadcastTx([]sdk.Msg{msg}, signerName, signerAccount)
}

func ProposeRevokeAccount(
	suite *utils.TestSuite,
	accAddr sdk.AccAddress,
	signerName string,
	signerAccount *dclauthtypes.Account,
	info string,
) (*sdk.TxResponse, error) {
	msg := dclauthtypes.NewMsgProposeRevokeAccount(suite.GetAddress(signerName), accAddr, info)

	return suite.BuildAndBroadcastTx([]sdk.Msg{msg}, signerName, signerAccount)
}

func ApproveRevokeAccount(
	suite *utils.TestSuite,
	accAddr sdk.AccAddress,
	signerName string,
	signerAccount *dclauthtypes.Account,
	info string,
) (*sdk.TxResponse, error) {
	msg := dclauthtypes.NewMsgApproveRevokeAccount(suite.GetAddress(signerName), accAddr, info)

	return suite.BuildAndBroadcastTx([]sdk.Msg{msg}, signerName, signerAccount)
}

func RejectAddAccount(
	suite *utils.TestSuite,
	accAddr sdk.AccAddress,
	signerName string,
	signerAccount *dclauthtypes.Account,
	info string,
) (*sdk.TxResponse, error) {
	msg := dclauthtypes.NewMsgRejectAddAccount(suite.GetAddress(signerName), accAddr, info)

	return suite.BuildAndBroadcastTx([]sdk.Msg{msg}, signerName, signerAccount)
}

func CreateAccountInfo(suite *utils.TestSuite, accountName string) keyring.Info {
	entropySeed, err := bip39.NewEntropy(256)
	require.NoError(suite.T, err)

	mnemonic, err := bip39.NewMnemonic(entropySeed)
	require.NoError(suite.T, err)

	accountInfo, err := suite.Kr.NewAccount(accountName, mnemonic, testconstants.Passphrase, sdk.FullFundraiserPath, hd.Secp256k1)
	require.NoError(suite.T, err)

	return accountInfo
}

func CreateAccount(
	suite *utils.TestSuite,
	accountName string,
	roles dclauthtypes.AccountRoles,
	vendorID int32,
	proposerName string,
	proposerAccount *dclauthtypes.Account,
	approverName string,
	approverAccount *dclauthtypes.Account,
	info string,
) *dclauthtypes.Account {
	accountInfo := CreateAccountInfo(suite, accountName)

	_, err := ProposeAddAccount(
		suite,
		accountInfo.GetAddress(),
		accountInfo.GetPubKey(),
		roles,
		vendorID,
		proposerName,
		proposerAccount,
		info,
	)
	require.NoError(suite.T, err)

	_, err = ApproveAddAccount(
		suite,
		accountInfo.GetAddress(),
		approverName,
		approverAccount,
		info,
	)
	require.NoError(suite.T, err)

	account, err := GetAccount(suite, accountInfo.GetAddress())
	require.NoError(suite.T, err)

	return account
}

func NewMsgCreateModel(vid int32, pid int32, signer string) *modeltypes.MsgCreateModel {
	return &modeltypes.MsgCreateModel{
		Creator:                                  signer,
		Vid:                                      vid,
		Pid:                                      pid,
		DeviceTypeId:                             testconstants.DeviceTypeID,
		ProductName:                              utils.RandString(),
		ProductLabel:                             utils.RandString(),
		PartNumber:                               utils.RandString(),
		CommissioningCustomFlow:                  testconstants.CommissioningCustomFlow,
		CommissioningCustomFlowUrl:               testconstants.CommissioningCustomFlowURL,
		CommissioningModeInitialStepsHint:        testconstants.CommissioningModeInitialStepsHint,
		CommissioningModeInitialStepsInstruction: testconstants.CommissioningModeInitialStepsInstruction,
		CommissioningModeSecondaryStepsHint:      testconstants.CommissioningModeSecondaryStepsHint,
		CommissioningModeSecondaryStepsInstruction: testconstants.CommissioningModeSecondaryStepsInstruction,
		UserManualUrl: testconstants.UserManualURL,
		SupportUrl:    testconstants.SupportURL,
		ProductUrl:    testconstants.ProductURL,
		LsfUrl:        testconstants.LsfURL,
	}
}

// Common Test Logic

//nolint:funlen
func AuthDemo(suite *utils.TestSuite) {
	// Jack, Alice and Bob are predefined Trustees
	jackName := testconstants.JackAccount
	jackKeyInfo, err := suite.Kr.Key(jackName)
	require.NoError(suite.T, err)
	jackAccount, err := GetAccount(suite, jackKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	aliceName := testconstants.AliceAccount
	aliceKeyInfo, err := suite.Kr.Key(aliceName)
	require.NoError(suite.T, err)
	aliceAccount, err := GetAccount(suite, aliceKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	bobName := testconstants.BobAccount
	bobKeyInfo, err := suite.Kr.Key(bobName)
	require.NoError(suite.T, err)
	bobAccount, err := GetAccount(suite, bobKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	// Query all active accounts
	inputAccounts, err := GetAccounts(suite)
	require.NoError(suite.T, err)

	// Query all proposed accounts
	inputProposedAccounts, err := GetProposedAccounts(suite)
	require.NoError(suite.T, err)
	require.Equal(suite.T, 0, len(inputProposedAccounts))

	// Query all proposed accounts to revoke
	inputProposedAccountsToRevoke, err := GetProposedAccountsToRevoke(suite)
	require.NoError(suite.T, err)
	require.Equal(suite.T, 0, len(inputProposedAccountsToRevoke))

	// Query all revoked accounts
	inputRevokedAccounts, err := GetRevokedAccounts(suite)
	require.NoError(suite.T, err)
	require.Equal(suite.T, 0, len(inputRevokedAccounts))

	accountName := utils.RandString()
	accountInfo := CreateAccountInfo(suite, accountName)
	testAccPubKey := accountInfo.GetPubKey()
	testAccAddr := accountInfo.GetAddress()

	// Query unknown account
	_, err = GetAccount(suite, testAccAddr)
	suite.AssertNotFound(err)

	// Query unknown proposed account
	_, err = GetProposedAccount(suite, testAccAddr)
	suite.AssertNotFound(err)

	// Query unknown proposed account to revoke
	_, err = GetProposedAccountToRevoke(suite, testAccAddr)
	suite.AssertNotFound(err)

	// Query unknown revoked account
	_, err = GetRevokedAccount(suite, testAccAddr)
	suite.AssertNotFound(err)

	// Jack proposes new account
	vid := int32(tmrand.Uint16())
	_, err = ProposeAddAccount(
		suite,
		testAccAddr, testAccPubKey,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor}, vid,
		jackName, jackAccount,
		testconstants.Info,
	)
	require.NoError(suite.T, err)

	// Query all active accounts
	receivedAccounts, _ := GetAccounts(suite)
	require.Equal(suite.T, len(inputAccounts), len(receivedAccounts))

	// Query all proposed accounts
	receivedProposedAccounts, _ := GetProposedAccounts(suite)
	require.Equal(suite.T, len(inputProposedAccounts)+1, len(receivedProposedAccounts))

	// Query all accounts proposed to be revoked
	receivedProposedAccountsToRevoke, _ := GetProposedAccountsToRevoke(suite)
	require.Equal(suite.T, len(inputProposedAccountsToRevoke), len(receivedProposedAccountsToRevoke))

	// Query all revoked accounts
	receivedRevokedAccounts, _ := GetRevokedAccounts(suite)
	require.Equal(suite.T, len(inputRevokedAccounts), len(receivedRevokedAccounts))

	// Query unknown account
	_, err = GetAccount(suite, testAccAddr)
	suite.AssertNotFound(err)

	// Query unknown proposed account to revoke
	_, err = GetProposedAccountToRevoke(suite, testAccAddr)
	suite.AssertNotFound(err)

	// Query proposed account
	proposedAccount, err := GetProposedAccount(suite, testAccAddr)
	require.NoError(suite.T, err)
	require.Equal(suite.T, testAccAddr, proposedAccount.GetAddress())
	require.Equal(suite.T, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, proposedAccount.GetRoles())

	// Query unknown revoked account
	_, err = GetRevokedAccount(suite, testAccAddr)
	suite.AssertNotFound(err)

	// Alice approves new account
	_, err = ApproveAddAccount(suite, testAccAddr, aliceName, aliceAccount, testconstants.Info)
	require.NoError(suite.T, err)

	// Alice doesn't reject accont, because Alice already approved account
	_, err = ApproveAddAccount(suite, testAccAddr, aliceName, aliceAccount, testconstants.Info)
	require.Error(suite.T, err)

	// Query all active accounts
	receivedAccounts, _ = GetAccounts(suite)
	require.Equal(suite.T, len(inputAccounts)+1, len(receivedAccounts))

	// Query all proposed accounts
	receivedProposedAccounts, _ = GetProposedAccounts(suite)
	require.Equal(suite.T, len(inputProposedAccounts), len(receivedProposedAccounts))

	// Query all accounts proposed to be revoked
	receivedProposedAccountsToRevoke, _ = GetProposedAccountsToRevoke(suite)
	require.Equal(suite.T, len(inputProposedAccountsToRevoke), len(receivedProposedAccountsToRevoke))

	// Query all revoked accounts
	receivedRevokedAccounts, _ = GetRevokedAccounts(suite)
	require.Equal(suite.T, len(inputRevokedAccounts), len(receivedRevokedAccounts))

	// Get new account
	testAccount, err := GetAccount(suite, testAccAddr)
	require.NoError(suite.T, err)
	require.Equal(suite.T, testAccAddr, testAccount.GetAddress())
	require.Equal(suite.T, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testAccount.GetRoles())

	// Query unknown proposed account
	_, err = GetProposedAccount(suite, testAccAddr)
	suite.AssertNotFound(err)

	// Query unknown proposed account to revoke
	_, err = GetProposedAccountToRevoke(suite, testAccAddr)
	suite.AssertNotFound(err)

	// Query unknown revoked account
	_, err = GetRevokedAccount(suite, testAccAddr)
	suite.AssertNotFound(err)

	// Alice proposes to revoke new account
	_, err = ProposeRevokeAccount(suite, testAccAddr, aliceName, aliceAccount, testconstants.Info)
	require.NoError(suite.T, err)

	// Query all active accounts
	receivedAccounts, err = GetAccounts(suite)
	require.NoError(suite.T, err)
	require.Equal(suite.T, len(inputAccounts)+1, len(receivedAccounts))

	// Query all proposed accounts
	receivedProposedAccounts, _ = GetProposedAccounts(suite)
	require.Equal(suite.T, len(inputProposedAccounts), len(receivedProposedAccounts))

	// Query all accounts proposed to be revoked
	receivedProposedAccountsToRevoke, _ = GetProposedAccountsToRevoke(suite)
	require.Equal(suite.T, len(inputProposedAccountsToRevoke)+1, len(receivedProposedAccountsToRevoke))

	// Query all revoked accounts
	receivedRevokedAccounts, _ = GetRevokedAccounts(suite)
	require.Equal(suite.T, len(inputRevokedAccounts), len(receivedRevokedAccounts))

	// Query proposed account to revoke
	proposedToRevokeAccount, err := GetProposedAccountToRevoke(suite, testAccAddr)
	require.NoError(suite.T, err)
	require.Equal(suite.T, testAccAddr.String(), proposedToRevokeAccount.GetAddress())

	// Bob approves to revoke new account
	_, err = ApproveRevokeAccount(suite, testAccAddr, bobName, bobAccount, testconstants.Info)
	require.NoError(suite.T, err)

	// Query all active accounts
	receivedAccounts, err = GetAccounts(suite)
	require.NoError(suite.T, err)
	require.Equal(suite.T, len(inputAccounts), len(receivedAccounts))

	// Query all proposed accounts
	receivedProposedAccounts, _ = GetProposedAccounts(suite)
	require.Equal(suite.T, len(inputProposedAccounts), len(receivedProposedAccounts))

	// Query all accounts proposed to be revoked
	receivedProposedAccountsToRevoke, _ = GetProposedAccountsToRevoke(suite)
	require.Equal(suite.T, len(inputProposedAccountsToRevoke), len(receivedProposedAccountsToRevoke))

	// Query all revoked accounts
	receivedRevokedAccounts, _ = GetRevokedAccounts(suite)
	require.Equal(suite.T, len(inputRevokedAccounts)+1, len(receivedRevokedAccounts))

	// Ensure that new account is not available anymore
	_, err = GetAccount(suite, testAccAddr)
	require.Error(suite.T, err)
	suite.AssertNotFound(err)

	_, err = GetProposedAccount(suite, testAccAddr)
	require.Error(suite.T, err)
	suite.AssertNotFound(err)

	_, err = GetProposedAccountToRevoke(suite, testAccAddr)
	require.Error(suite.T, err)
	suite.AssertNotFound(err)

	// Try to publish another model info by test account.
	// Ensure that the request is responded with not OK status code.
	pid := int32(tmrand.Uint16())
	firstModel := NewMsgCreateModel(vid, pid, testAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{firstModel}, accountName, testAccount)
	require.Error(suite.T, err)
	require.True(suite.T, sdkerrors.ErrUnknownAddress.Is(err))

	// Ensure that new account is exist in revoked account
	receivedRevokedAccount, err := GetRevokedAccount(suite, testAccAddr)
	require.NoError(suite.T, err)
	require.Equal(suite.T, aliceAccount.Address, receivedRevokedAccount.RevokeApprovals[0].Address)
	require.Equal(suite.T, bobAccount.Address, receivedRevokedAccount.RevokeApprovals[1].Address)
	require.Equal(suite.T, dclauthtypes.RevokedAccount_TrusteeVoting, receivedRevokedAccount.Reason)

	// Jack proposes for re-adding the revoked account
	_, err = ProposeAddAccount(
		suite,
		testAccAddr, testAccPubKey,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor}, vid,
		jackName, jackAccount,
		testconstants.Info,
	)
	require.NoError(suite.T, err)

	// Query all active accounts
	receivedAccounts, _ = GetAccounts(suite)
	require.Equal(suite.T, len(inputAccounts), len(receivedAccounts))

	// Query all proposed accounts
	receivedProposedAccounts, _ = GetProposedAccounts(suite)
	require.Equal(suite.T, len(inputProposedAccounts)+1, len(receivedProposedAccounts))

	// Query all accounts proposed to be revoked
	receivedProposedAccountsToRevoke, _ = GetProposedAccountsToRevoke(suite)
	require.Equal(suite.T, len(inputProposedAccountsToRevoke), len(receivedProposedAccountsToRevoke))

	// Query all revoked accounts
	receivedRevokedAccounts, _ = GetRevokedAccounts(suite)
	require.Equal(suite.T, len(inputRevokedAccounts)+1, len(receivedRevokedAccounts))

	// Query unknown account
	_, err = GetAccount(suite, testAccAddr)
	suite.AssertNotFound(err)

	// Query unknown proposed account to revoke
	_, err = GetProposedAccountToRevoke(suite, testAccAddr)
	suite.AssertNotFound(err)

	// Query proposed account
	proposedAccount, err = GetProposedAccount(suite, testAccAddr)
	require.NoError(suite.T, err)
	require.Equal(suite.T, testAccAddr, proposedAccount.GetAddress())
	require.Equal(suite.T, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, proposedAccount.GetRoles())

	// Ensure that new account is exist in revoked account
	receivedRevokedAccount, err = GetRevokedAccount(suite, testAccAddr)
	require.NoError(suite.T, err)
	require.Equal(suite.T, testAccAddr, receivedRevokedAccount.GetAddress())
	require.Equal(suite.T, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, receivedRevokedAccount.GetRoles())

	// Alice approves the revoked account
	_, err = ApproveAddAccount(suite, testAccAddr, aliceName, aliceAccount, testconstants.Info)
	require.NoError(suite.T, err)

	// Alice doesn't reject accont, because Alice already approved account
	_, err = ApproveAddAccount(suite, testAccAddr, aliceName, aliceAccount, testconstants.Info)
	require.Error(suite.T, err)

	// Query all active accounts
	receivedAccounts, _ = GetAccounts(suite)
	require.Equal(suite.T, len(inputAccounts)+1, len(receivedAccounts))

	// Query all proposed accounts
	receivedProposedAccounts, _ = GetProposedAccounts(suite)
	require.Equal(suite.T, len(inputProposedAccounts), len(receivedProposedAccounts))

	// Query all accounts proposed to be revoked
	receivedProposedAccountsToRevoke, _ = GetProposedAccountsToRevoke(suite)
	require.Equal(suite.T, len(inputProposedAccountsToRevoke), len(receivedProposedAccountsToRevoke))

	// Query all revoked accounts
	receivedRevokedAccounts, _ = GetRevokedAccounts(suite)
	require.Equal(suite.T, len(inputRevokedAccounts), len(receivedRevokedAccounts))

	// Get new account
	testAccount, err = GetAccount(suite, testAccAddr)
	require.NoError(suite.T, err)
	require.Equal(suite.T, testAccAddr, testAccount.GetAddress())
	require.Equal(suite.T, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testAccount.GetRoles())

	// Query unknown proposed account
	_, err = GetProposedAccount(suite, testAccAddr)
	suite.AssertNotFound(err)

	// Query unknown proposed account to revoke
	_, err = GetProposedAccountToRevoke(suite, testAccAddr)
	suite.AssertNotFound(err)

	// Query unknown revoked account
	_, err = GetRevokedAccount(suite, testAccAddr)
	suite.AssertNotFound(err)

	// Alice proposes to revoke new account
	_, err = ProposeRevokeAccount(suite, testAccAddr, aliceName, aliceAccount, testconstants.Info)
	require.NoError(suite.T, err)

	// Bob approves to revoke new account
	_, err = ApproveRevokeAccount(suite, testAccAddr, bobName, bobAccount, testconstants.Info)
	require.NoError(suite.T, err)

	// Query all rejected accounts
	inputRejectedAccounts, err := GetRejectedAccounts(suite)
	require.NoError(suite.T, err)
	require.Equal(suite.T, 0, len(inputRejectedAccounts))

	// Query unknow rejected account
	_, err = GetRejectedAccount(suite, testAccAddr)
	suite.AssertNotFound(err)

	// Jack proposes account
	_, err = ProposeAddAccount(
		suite,
		testAccAddr, testAccPubKey,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor}, vid,
		jackName, jackAccount,
		testconstants.Info,
	)
	require.NoError(suite.T, err)

	// Alice rejects account
	_, err = RejectAddAccount(suite, testAccAddr, aliceName, aliceAccount, testconstants.Info)
	require.NoError(suite.T, err)

	// Alice doesn't approve accont, because Alice already rejected account
	_, err = ApproveAddAccount(suite, testAccAddr, aliceName, aliceAccount, testconstants.Info)
	require.Error(suite.T, err)

	// Query all active accounts
	receivedAccounts, _ = GetAccounts(suite)
	require.Equal(suite.T, len(inputAccounts), len(receivedAccounts))

	// Query all proposed accounts
	receivedProposedAccounts, _ = GetProposedAccounts(suite)
	require.Equal(suite.T, len(inputProposedAccounts)+1, len(receivedProposedAccounts))

	// Query all rejected accounts
	receivedRejectedAccounts, err := GetRejectedAccounts(suite)
	require.NoError(suite.T, err)
	require.Equal(suite.T, len(inputRejectedAccounts), len(receivedRejectedAccounts))

	// Query all accounts proposed to be revoked
	receivedProposedAccountsToRevoke, _ = GetProposedAccountsToRevoke(suite)
	require.Equal(suite.T, len(inputProposedAccountsToRevoke), len(receivedProposedAccountsToRevoke))

	// Query all revoked accounts
	receivedRevokedAccounts, _ = GetRevokedAccounts(suite)
	require.Equal(suite.T, len(inputRevokedAccounts)+1, len(receivedRevokedAccounts))

	// Query unknown account
	_, err = GetAccount(suite, testAccAddr)
	suite.AssertNotFound(err)

	// Query unknown rejected account
	_, err = GetRejectedAccount(suite, testAccAddr)
	suite.AssertNotFound(err)

	// Query proposed account
	proposedAccount, err = GetProposedAccount(suite, testAccAddr)
	require.NoError(suite.T, err)
	require.Equal(suite.T, testAccAddr, proposedAccount.GetAddress())
	require.Equal(suite.T, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, proposedAccount.GetRoles())
	require.Equal(suite.T, jackAccount.Address, proposedAccount.Approvals[0].Address)
	require.Equal(suite.T, testconstants.Info, proposedAccount.Approvals[0].Info)
	require.Equal(suite.T, aliceAccount.Address, proposedAccount.Rejects[0].Address)
	require.Equal(suite.T, testconstants.Info, proposedAccount.Rejects[0].Info)

	// Query unknown proposed account to revoke
	_, err = GetProposedAccountToRevoke(suite, testAccAddr)
	suite.AssertNotFound(err)

	// Query unknown revoked account
	_, err = GetRevokedAccount(suite, testAccAddr)
	require.NoError(suite.T, err)

	// Bob rejects account
	_, err = RejectAddAccount(suite, testAccAddr, bobName, bobAccount, testconstants.Info)
	require.NoError(suite.T, err)

	// Bob can not the second time reject the same account
	_, err = ApproveAddAccount(suite, testAccAddr, aliceName, aliceAccount, testconstants.Info)
	require.Error(suite.T, err)

	// Query all active accounts
	receivedAccounts, _ = GetAccounts(suite)
	require.Equal(suite.T, len(inputAccounts), len(receivedAccounts))

	// Query all proposed accounts
	receivedProposedAccounts, _ = GetProposedAccounts(suite)
	require.Equal(suite.T, len(inputProposedAccounts), len(receivedProposedAccounts))

	// Query all rejected accounts
	receivedRejectedAccounts, err = GetRejectedAccounts(suite)
	require.NoError(suite.T, err)
	require.Equal(suite.T, len(inputRejectedAccounts)+1, len(receivedRejectedAccounts))

	// Query all accounts proposed to be revoked
	receivedProposedAccountsToRevoke, _ = GetProposedAccountsToRevoke(suite)
	require.Equal(suite.T, len(inputProposedAccountsToRevoke), len(receivedProposedAccountsToRevoke))

	// Query all revoked accounts
	receivedRevokedAccounts, _ = GetRevokedAccounts(suite)
	require.Equal(suite.T, len(inputRevokedAccounts)+1, len(receivedRevokedAccounts))

	// Query unknown account
	_, err = GetAccount(suite, testAccAddr)
	suite.AssertNotFound(err)

	// Query rejected account
	rejectedAccount, err := GetRejectedAccount(suite, testAccAddr)
	require.NoError(suite.T, err)
	require.Equal(suite.T, testAccAddr, rejectedAccount.GetAddress())
	require.Equal(suite.T, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, rejectedAccount.GetRoles())
	require.Equal(suite.T, jackAccount.Address, rejectedAccount.Approvals[0].Address)
	require.Equal(suite.T, testconstants.Info, rejectedAccount.Approvals[0].Info)
	require.Equal(suite.T, aliceAccount.Address, rejectedAccount.Rejects[0].Address)
	require.Equal(suite.T, testconstants.Info, rejectedAccount.Rejects[0].Info)

	// Query proposed account
	_, err = GetProposedAccount(suite, testAccAddr)
	require.Error(suite.T, err)

	// Query unknown proposed account to revoke
	_, err = GetProposedAccountToRevoke(suite, testAccAddr)
	suite.AssertNotFound(err)

	// Query unknown revoked account
	_, err = GetRevokedAccount(suite, testAccAddr)
	require.NoError(suite.T, err)
}
