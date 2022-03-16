package dclupgrade

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/stretchr/testify/require"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	test_dclauth "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/grpc_rest/dclauth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	dclupgradetypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
)

func NewMsgProposeUpgrade(signer string, name string, height int64, info string) *dclupgradetypes.MsgProposeUpgrade {
	return &dclupgradetypes.MsgProposeUpgrade{
		Creator: signer,
		Plan: upgradetypes.Plan{
			Name:   name,
			Height: height,
			Info:   info,
		},
	}
}

func NewMsgApproveUpgrade(signer string, name string) *dclupgradetypes.MsgApproveUpgrade {
	return &dclupgradetypes.MsgApproveUpgrade{
		Creator: signer,
		Name:    name,
	}
}

func ProposeUpgrade(
	suite *utils.TestSuite,
	msg *dclupgradetypes.MsgProposeUpgrade,
	signerName string,
	signerAccount *dclauthtypes.Account,
) (*sdk.TxResponse, error) {
	msg.Creator = suite.GetAddress(signerName).String()

	return suite.BuildAndBroadcastTx([]sdk.Msg{msg}, signerName, signerAccount)
}

func ApproveUpgrade(
	suite *utils.TestSuite,
	msg *dclupgradetypes.MsgApproveUpgrade,
	signerName string,
	signerAccount *dclauthtypes.Account,
) (*sdk.TxResponse, error) {
	msg.Creator = suite.GetAddress(signerName).String()

	return suite.BuildAndBroadcastTx([]sdk.Msg{msg}, signerName, signerAccount)
}

func GetProposedUpgrade(
	suite *utils.TestSuite,
	name string,
) (*dclupgradetypes.ProposedUpgrade, error) {
	var res dclupgradetypes.ProposedUpgrade

	if suite.Rest {
		var resp dclupgradetypes.QueryGetProposedUpgradeResponse
		err := suite.QueryREST(fmt.Sprintf("/dcl/dclupgrade/proposed-upgrades/%s", name), &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetProposedUpgrade()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/dclauth service.
		dclupgradeClient := dclupgradetypes.NewQueryClient(grpcConn)
		resp, err := dclupgradeClient.ProposedUpgrade(
			context.Background(),
			&dclupgradetypes.QueryGetProposedUpgradeRequest{Name: name},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetProposedUpgrade()
	}

	return &res, nil
}

func GetApprovedUpgrade(
	suite *utils.TestSuite,
	name string,
) (*dclupgradetypes.ApprovedUpgrade, error) {
	var res dclupgradetypes.ApprovedUpgrade

	if suite.Rest {
		var resp dclupgradetypes.QueryGetApprovedUpgradeResponse
		err := suite.QueryREST(fmt.Sprintf("/dcl/dclupgrade/approved-upgrades/%s", name), &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetApprovedUpgrade()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/dclauth service.
		dclupgradeClient := dclupgradetypes.NewQueryClient(grpcConn)
		resp, err := dclupgradeClient.ApprovedUpgrade(
			context.Background(),
			&dclupgradetypes.QueryGetApprovedUpgradeRequest{Name: name},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetApprovedUpgrade()
	}

	return &res, nil
}

func GetProposedUpgrades(suite *utils.TestSuite) (res []dclupgradetypes.ProposedUpgrade, err error) {
	if suite.Rest {
		var resp dclupgradetypes.QueryAllProposedUpgradeResponse
		err := suite.QueryREST("/dcl/dclupgrade/proposed-upgrades", &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetProposedUpgrade()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/dclauth service.
		dclupgradeClient := dclupgradetypes.NewQueryClient(grpcConn)
		resp, err := dclupgradeClient.ProposedUpgradeAll(
			context.Background(),
			&dclupgradetypes.QueryAllProposedUpgradeRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetProposedUpgrade()
	}

	return res, nil
}

func GetApprovedUpgrades(suite *utils.TestSuite) (res []dclupgradetypes.ApprovedUpgrade, err error) {
	if suite.Rest {
		var resp dclupgradetypes.QueryAllApprovedUpgradeResponse
		err := suite.QueryREST("/dcl/dclupgrade/approved-upgrades", &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetApprovedUpgrade()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/dclauth service.
		dclupgradeClient := dclupgradetypes.NewQueryClient(grpcConn)
		resp, err := dclupgradeClient.ApprovedUpgradeAll(
			context.Background(),
			&dclupgradetypes.QueryAllApprovedUpgradeRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetApprovedUpgrade()
	}

	return res, nil
}

func DCLUpgradeDemo(suite *utils.TestSuite) {
	// Alice and Bob are predefined Trustees
	aliceName := testconstants.AliceAccount
	aliceKeyInfo, err := suite.Kr.Key(aliceName)
	require.NoError(suite.T, err)
	aliceAccount, err := test_dclauth.GetAccount(suite, aliceKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	bobName := testconstants.BobAccount
	bobKeyInfo, err := suite.Kr.Key(bobName)
	require.NoError(suite.T, err)
	bobAccount, err := test_dclauth.GetAccount(suite, bobKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	// trustee proposes upgrade
	proposeUpgradeMsg := NewMsgProposeUpgrade(aliceAccount.Address, utils.RandString(), 100000, utils.RandString())
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{proposeUpgradeMsg}, aliceName, aliceAccount)
	require.NoError(suite.T, err)

	// Check upgrade is proposed
	proposedUpgrade, err := GetProposedUpgrade(suite, proposeUpgradeMsg.Plan.Name)
	require.NoError(suite.T, err)

	require.Equal(suite.T, proposeUpgradeMsg.Creator, proposedUpgrade.Creator)
	require.Equal(suite.T, proposeUpgradeMsg.Plan, proposedUpgrade.Plan)

	require.Equal(suite.T, 1, len(proposedUpgrade.Approvals))

	require.Equal(suite.T, proposeUpgradeMsg.Creator, proposedUpgrade.Approvals[0].Address)
	require.Equal(suite.T, proposeUpgradeMsg.Time, proposedUpgrade.Approvals[0].Time)
	require.Equal(suite.T, proposeUpgradeMsg.Info, proposedUpgrade.Approvals[0].Info)

	// Get all proposed upgrades
	proposedUpgrades, err := GetProposedUpgrades(suite)
	require.NoError(suite.T, err)
	require.Contains(suite.T, proposedUpgrades, *proposedUpgrade)

	// Get approved upgrade
	_, err = GetApprovedUpgrade(suite, proposeUpgradeMsg.Plan.Name)
	suite.AssertNotFound(err)

	// another trustee approves upgrade
	approveUpgradeMsg := NewMsgApproveUpgrade(bobAccount.Address, proposeUpgradeMsg.Plan.Name)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{approveUpgradeMsg}, bobName, bobAccount)
	require.NoError(suite.T, err)

	// Check upgrade is approved
	approvedUpgrade, err := GetApprovedUpgrade(suite, proposeUpgradeMsg.Plan.Name)
	require.NoError(suite.T, err)

	require.Equal(suite.T, proposeUpgradeMsg.Creator, approvedUpgrade.Creator)
	require.Equal(suite.T, proposeUpgradeMsg.Plan, approvedUpgrade.Plan)

	require.Equal(suite.T, 2, len(approvedUpgrade.Approvals))

	require.Equal(suite.T, proposeUpgradeMsg.Creator, approvedUpgrade.Approvals[0].Address)
	require.Equal(suite.T, proposeUpgradeMsg.Time, approvedUpgrade.Approvals[0].Time)
	require.Equal(suite.T, proposeUpgradeMsg.Info, approvedUpgrade.Approvals[0].Info)

	require.Equal(suite.T, approveUpgradeMsg.Creator, approvedUpgrade.Approvals[1].Address)
	require.Equal(suite.T, approveUpgradeMsg.Time, approvedUpgrade.Approvals[1].Time)
	require.Equal(suite.T, approveUpgradeMsg.Info, approvedUpgrade.Approvals[1].Info)

	// Get all proposed upgrades
	proposedUpgrades, err = GetProposedUpgrades(suite)
	require.NoError(suite.T, err)
	require.NotContains(suite.T, proposedUpgrades, *proposedUpgrade)

	// Get all approved upgrades
	approvedUpgrades, err := GetApprovedUpgrades(suite)
	require.NoError(suite.T, err)
	require.Contains(suite.T, approvedUpgrades, *approvedUpgrade)
}

/* Error cases */

func ProposeUpgradeByNonTrustee(suite *utils.TestSuite) {
	// Alice and Bob are predefined Trustees
	aliceName := testconstants.AliceAccount
	aliceKeyInfo, err := suite.Kr.Key(aliceName)
	require.NoError(suite.T, err)
	aliceAccount, err := test_dclauth.GetAccount(suite, aliceKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	bobName := testconstants.BobAccount
	bobKeyInfo, err := suite.Kr.Key(bobName)
	require.NoError(suite.T, err)
	bobAccount, err := test_dclauth.GetAccount(suite, bobKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	// register new account without Trustee role
	nonTrusteeAccountName := utils.RandString()
	nonTrusteeAccount := test_dclauth.CreateAccount(
		suite,
		nonTrusteeAccountName,
		dclauthtypes.AccountRoles{
			dclauthtypes.CertificationCenter,
			dclauthtypes.Vendor,
			dclauthtypes.NodeAdmin,
		},
		int32(tmrand.Uint16()+1),
		aliceName,
		aliceAccount,
		bobName,
		bobAccount,
		testconstants.Info,
	)

	// try to add proposeUpgradeMsg
	proposeUpgradeMsg := NewMsgProposeUpgrade(nonTrusteeAccount.Address, utils.RandString(), 100000, utils.RandString())
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{proposeUpgradeMsg}, nonTrusteeAccountName, nonTrusteeAccount)
	require.Error(suite.T, err)
	require.ErrorIs(suite.T, err, sdkerrors.ErrUnauthorized)
}

func ApproveUpgradeByNonTrustee(suite *utils.TestSuite) {
	// Alice and Bob are predefined Trustees
	aliceName := testconstants.AliceAccount
	aliceKeyInfo, err := suite.Kr.Key(aliceName)
	require.NoError(suite.T, err)
	aliceAccount, err := test_dclauth.GetAccount(suite, aliceKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	bobName := testconstants.BobAccount
	bobKeyInfo, err := suite.Kr.Key(bobName)
	require.NoError(suite.T, err)
	bobAccount, err := test_dclauth.GetAccount(suite, bobKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	// propose upgrade
	proposeUpgradeMsg := NewMsgProposeUpgrade(aliceAccount.Address, utils.RandString(), 100000, utils.RandString())
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{proposeUpgradeMsg}, aliceName, aliceAccount)
	require.NoError(suite.T, err)

	// register new account without Trustee role
	nonTrusteeAccountName := utils.RandString()
	nonTrusteeAccount := test_dclauth.CreateAccount(
		suite,
		nonTrusteeAccountName,
		dclauthtypes.AccountRoles{
			dclauthtypes.CertificationCenter,
			dclauthtypes.Vendor,
			dclauthtypes.NodeAdmin,
		},
		int32(tmrand.Uint16()+1),
		aliceName,
		aliceAccount,
		bobName,
		bobAccount,
		testconstants.Info,
	)

	// try to approve upgrade
	approveUpgradeMsg := NewMsgApproveUpgrade(nonTrusteeAccount.Address, proposeUpgradeMsg.Plan.Name)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{approveUpgradeMsg}, nonTrusteeAccountName, nonTrusteeAccount)
	require.Error(suite.T, err)
	require.ErrorIs(suite.T, err, sdkerrors.ErrUnauthorized)
}

func ProposeUpgradeTwice(suite *utils.TestSuite) {
	// Alice and Bob are predefined Trustees
	aliceName := testconstants.AliceAccount
	aliceKeyInfo, err := suite.Kr.Key(aliceName)
	require.NoError(suite.T, err)
	aliceAccount, err := test_dclauth.GetAccount(suite, aliceKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	bobName := testconstants.BobAccount
	bobKyInfo, err := suite.Kr.Key(bobName)
	require.NoError(suite.T, err)
	bobAccount, err := test_dclauth.GetAccount(suite, bobKyInfo.GetAddress())
	require.NoError(suite.T, err)

	planName, height, info := utils.RandString(), int64(100000), utils.RandString()

	// trustee proposes upgrade
	proposeUpgradeMsg := NewMsgProposeUpgrade(aliceAccount.Address, planName, height, info)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{proposeUpgradeMsg}, aliceName, aliceAccount)
	require.NoError(suite.T, err)

	// another trustee proposes the same upgrade
	proposeUpgradeMsg.Creator = bobAccount.Address
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{proposeUpgradeMsg}, bobName, bobAccount)
	require.Error(suite.T, err)
	require.ErrorIs(suite.T, err, dclupgradetypes.ErrProposedUpgradeAlreadyExists)
}
