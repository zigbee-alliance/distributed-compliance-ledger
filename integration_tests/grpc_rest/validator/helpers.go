package validator

import (
	"context"
	"fmt"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	validatortypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func GetDisabledValidator(suite *utils.TestSuite, address sdk.ValAddress) (*validatortypes.DisabledValidator, error) {
	var res validatortypes.DisabledValidator

	if suite.Rest {
		var resp validatortypes.QueryGetDisabledValidatorResponse
		err := suite.QueryREST(fmt.Sprintf("/dcl/validator/disabled_validators/%s", address.String()), &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetDisabledValidator()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/validator service.
		validatorClient := validatortypes.NewQueryClient(grpcConn)
		resp, err := validatorClient.DisabledValidator(
			context.Background(),
			&validatortypes.QueryGetDisabledValidatorRequest{Address: address.String()},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetDisabledValidator()
	}

	return &res, nil
}

func GetDisabledValidators(suite *utils.TestSuite) (res []validatortypes.DisabledValidator, err error) {
	if suite.Rest {
		var resp validatortypes.QueryAllDisabledValidatorResponse
		err := suite.QueryREST("/dcl/validator/disabled_validators", &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetDisabledValidator()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/validator service.
		validatorClient := validatortypes.NewQueryClient(grpcConn)
		resp, err := validatorClient.DisabledValidatorAll(
			context.Background(),
			&validatortypes.QueryAllDisabledValidatorRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetDisabledValidator()
	}

	return res, nil
}

func GetProposedValidatorToDisable(suite *utils.TestSuite, address sdk.ValAddress) (*validatortypes.ProposedDisableValidator, error) {
	var res validatortypes.ProposedDisableValidator

	if suite.Rest {
		var resp validatortypes.QueryGetProposedDisableValidatorResponse
		err := suite.QueryREST(fmt.Sprintf("/dcl/validator/proposed_disable_validators/%s", address.String()), &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetProposedDisableValidator()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/validator service.
		validatorClient := validatortypes.NewQueryClient(grpcConn)
		resp, err := validatorClient.ProposedDisableValidator(
			context.Background(),
			&validatortypes.QueryGetProposedDisableValidatorRequest{Address: address.String()},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetProposedDisableValidator()
	}

	return &res, nil
}

func GetProposedValidatorsToDisable(suite *utils.TestSuite) (
	res []validatortypes.ProposedDisableValidator, err error,
) {
	if suite.Rest {
		var resp validatortypes.QueryAllProposedDisableValidatorResponse
		err := suite.QueryREST("/dcl/validator/proposed_disable_validators", &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetProposedDisableValidator()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/validator service.
		validatorClient := validatortypes.NewQueryClient(grpcConn)
		resp, err := validatorClient.ProposedDisableValidatorAll(
			context.Background(),
			&validatortypes.QueryAllProposedDisableValidatorRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetProposedDisableValidator()
	}

	return res, nil
}

func CreateValidator(
	suite *utils.TestSuite,
	valAddr sdk.ValAddress,
	signerName string,
	signerAccount *dclauthtypes.Account,
	pubkey cryptotypes.PubKey,
	moniker string,
) (*sdk.TxResponse, error) {
	msg, err := validatortypes.NewMsgCreateValidator(valAddr, pubkey, &validatortypes.Description{Moniker: moniker})
	require.NoError(suite.T, err)
	return suite.BuildAndBroadcastTx([]sdk.Msg{msg}, signerName, signerAccount)
}

func DisableValidator(
	suite *utils.TestSuite,
	valAddr sdk.ValAddress,
	signerName string,
	signerAccount *dclauthtypes.Account,
) (*sdk.TxResponse, error) {
	msg := validatortypes.NewMsgDisableValidator(valAddr)
	return suite.BuildAndBroadcastTx([]sdk.Msg{msg}, signerName, signerAccount)
}

func EnableValidator(
	suite *utils.TestSuite,
	valAddr sdk.ValAddress,
	signerName string,
	signerAccount *dclauthtypes.Account,
) (*sdk.TxResponse, error) {
	msg := validatortypes.NewMsgEnableValidator(valAddr)
	return suite.BuildAndBroadcastTx([]sdk.Msg{msg}, signerName, signerAccount)
}

func ProposeDisableValidator(
	suite *utils.TestSuite,
	valAddr sdk.ValAddress,
	signerName string,
	signerAccount *dclauthtypes.Account,
	info string,
) (*sdk.TxResponse, error) {
	msg := validatortypes.NewMsgProposeDisableValidator(suite.GetAddress(signerName), valAddr, info)
	return suite.BuildAndBroadcastTx([]sdk.Msg{msg}, signerName, signerAccount)
}

func ApproveDisableValidator(
	suite *utils.TestSuite,
	accAddr sdk.ValAddress,
	signerName string,
	signerAccount *dclauthtypes.Account,
	info string,
) (*sdk.TxResponse, error) {
	msg := validatortypes.NewMsgApproveDisableValidator(suite.GetAddress(signerName), accAddr, info)
	return suite.BuildAndBroadcastTx([]sdk.Msg{msg}, signerName, signerAccount)
}

// Common Test Logic

//nolint:funlen
func ValidatorDemo(suite *utils.TestSuite) {
	// // Jack, Alice and Bob are predefined Trustees
	// jackName := testconstants.JackAccount
	// jackKeyInfo, err := suite.Kr.Key(jackName)
	// require.NoError(suite.T, err)
	// jackAccount, err := test_dclauth.GetAccount(suite, jackKeyInfo.GetAddress())
	// require.NoError(suite.T, err)

	// aliceName := testconstants.AliceAccount
	// aliceKeyInfo, err := suite.Kr.Key(aliceName)
	// require.NoError(suite.T, err)
	// aliceAccount, err := test_dclauth.GetAccount(suite, aliceKeyInfo.GetAddress())
	// require.NoError(suite.T, err)

	// bobName := testconstants.BobAccount
	// _, err = suite.Kr.Key(bobName)
	// require.NoError(suite.T, err)

	// // Register new Vendor account
	// vid := int32(tmrand.Uint16())
	// nodeAdminName := utils.RandString()
	// nodeAdminAcc := test_dclauth.CreateAccount(
	// 	suite,
	// 	nodeAdminName,
	// 	dclauthtypes.AccountRoles{dclauthtypes.NodeAdmin},
	// 	vid,
	// 	aliceName,
	// 	aliceAccount,
	// 	jackName,
	// 	jackAccount,
	// 	testconstants.Info,
	// )
	// nodeAdminAddr, err := sdk.AccAddressFromBech32(nodeAdminAcc.Address)
	// require.NoError(suite.T, err)
	// validatorAddr := sdk.ValAddress(nodeAdminAddr)
	// _, err = CreateValidator(suite, validatorAddr, nodeAdminName, nodeAdminAcc, nil, "test123")
	// require.NoError(suite.T, err)

	// // Query all proposed disable validators
	// proposedDisableValidators, err := GetProposedValidatorsToDisable(suite)
	// require.NoError(suite.T, err)
	// require.Equal(suite.T, 0, len(proposedDisableValidators))

	// // Query unknown disable validator
	// _, err = GetDisabledValidator(suite, sdk.ValAddress(testconstants.ValidatorAddress1))
	// suite.AssertNotFound(err)

	// // Query unknown proposed disable validator
	// _, err = GetProposedValidatorToDisable(suite, sdk.ValAddress(testconstants.ValidatorAddress1))
	// suite.AssertNotFound(err)

	// _, err = DisableValidator(suite, validatorAddr, nodeAdminName, nodeAdminAcc)
	// require.NoError(suite.T, err)

	// // Query disabled validator
	// disabledValidator, err := GetDisabledValidator(suite, validatorAddr)
	// require.NoError(suite.T, err)
	// require.True(suite.T, disabledValidator.DisabledByNodeAdmin)
	// require.Equal(suite.T, validatorAddr, disabledValidator.Address)
	// require.Empty(suite.T, disabledValidator.Approvals)

	// // Query all disabled validators
	// disabledValidators, err := GetDisabledValidators(suite)
	// require.NoError(suite.T, err)
	// require.Equal(suite.T, 1, len(disabledValidators))

	// // Enable validator
	// _, err = EnableValidator(suite, validatorAddr, nodeAdminName, nodeAdminAcc)
	// require.NoError(suite.T, err)

	// // Query all disabled validators
	// disabledValidators, err = GetDisabledValidators(suite)
	// require.NoError(suite.T, err)
	// require.Empty(suite.T, disabledValidators)

	// // Propose disable validator
	// _, err = ProposeDisableValidator(suite, validatorAddr, aliceName, aliceAccount, testconstants.Info)
	// require.NoError(suite.T, err)

	// // Query all validators proposed to be disabled
	// proposedValidatorsToDisable, err := GetProposedValidatorsToDisable(suite)
	// require.NoError(suite.T, err)
	// require.Equal(suite.T, 1, len(proposedValidatorsToDisable))

	// // Query proposed disable validator
	// proposedValidatorToDisable, err := GetProposedValidatorToDisable(suite, validatorAddr)
	// require.NoError(suite.T, err)
	// require.Equal(suite.T, validatorAddr, proposedValidatorToDisable.Address)
	// require.Equal(suite.T, aliceAccount.Address, proposedValidatorToDisable.Creator)
	// require.Equal(suite.T, 1, len(proposedValidatorToDisable.Approvals))

	// // Approve new disable validator
	// _, err = ApproveDisableValidator(suite, validatorAddr, jackName, jackAccount, testconstants.Info)
	// require.NoError(suite.T, err)

	// // Query disabled validator
	// disabledValidator, err = GetDisabledValidator(suite, validatorAddr)
	// require.NoError(suite.T, err)
	// require.False(suite.T, disabledValidator.DisabledByNodeAdmin)
	// require.Equal(suite.T, validatorAddr, disabledValidator.Address)
	// require.Equal(suite.T, 2, len(disabledValidator.Approvals))

	// // Query all disabled validators
	// disabledValidators, err = GetDisabledValidators(suite)
	// require.NoError(suite.T, err)
	// require.Equal(suite.T, 1, len(disabledValidators))

	// // Query all accounts proposed to be revoked
	// proposedValidatorsToDisable, _ = GetProposedValidatorsToDisable(suite)
	// require.Equal(suite.T, 0, len(proposedValidatorsToDisable))
}
