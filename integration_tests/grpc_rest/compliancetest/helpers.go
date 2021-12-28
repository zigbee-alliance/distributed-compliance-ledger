package compliancetest

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	test_dclauth "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/grpc_rest/dclauth"
	test_model "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/grpc_rest/model"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
	compliancetesttypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/types"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func GetAllTestResults(suite *utils.TestSuite) (res []compliancetesttypes.TestingResults, err error) {
	if suite.Rest {
		var resp compliancetesttypes.QueryAllTestingResultsResponse
		err := suite.QueryREST("/dcl/compliancetest/testing-results", &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetTestingResults()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		client := compliancetesttypes.NewQueryClient(grpcConn)
		resp, err := client.TestingResultsAll(
			context.Background(),
			&compliancetesttypes.QueryAllTestingResultsRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetTestingResults()
	}

	return res, nil
}

func GetTestResult(suite *utils.TestSuite, vid int32, pid int32, sv uint32) (res *compliancetesttypes.TestingResults, err error) {
	if suite.Rest {
		var resp compliancetesttypes.QueryGetTestingResultsResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/compliancetest/testing-results/%d/%d/%d",
				vid,
				pid,
				sv,
			),
			&resp,
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetTestingResults()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		client := compliancetesttypes.NewQueryClient(grpcConn)
		resp, err := client.TestingResults(
			context.Background(),
			&compliancetesttypes.QueryGetTestingResultsRequest{
				Vid:             vid,
				Pid:             pid,
				SoftwareVersion: sv,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetTestingResults()
	}

	return res, nil
}

//nolint:funlen
func ComplianceTestDemo(suite *utils.TestSuite) {
	// Query for unknown test results
	_, err := GetTestResult(suite, testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion)
	suite.AssertNotFound(err)

	// Query for empty test results
	testResults, _ := GetAllTestResults(suite)
	require.Equal(suite.T, 0, len(testResults))

	// Alice and Jack are predefined Trustees
	aliceName := testconstants.AliceAccount
	aliceKeyInfo, err := suite.Kr.Key(aliceName)
	require.NoError(suite.T, err)
	aliceAccount, err := test_dclauth.GetAccount(suite, aliceKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	jackName := testconstants.JackAccount
	jackKeyInfo, err := suite.Kr.Key(jackName)
	require.NoError(suite.T, err)
	jackAccount, err := test_dclauth.GetAccount(suite, jackKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	// Register new Vendor account
	vid := int32(tmrand.Uint16())
	vendorName := utils.RandString()
	vendorAccount := test_dclauth.CreateAccount(
		suite,
		vendorName,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		uint64(vid),
		aliceName,
		aliceAccount,
		jackName,
		jackAccount,
	)
	require.NotNil(suite.T, vendorAccount)

	// Register new TestHouse account
	testHouse := utils.RandString()
	testHouseAccount := test_dclauth.CreateAccount(
		suite,
		testHouse,
		dclauthtypes.AccountRoles{dclauthtypes.TestHouse},
		uint64(1),
		aliceName,
		aliceAccount,
		jackName,
		jackAccount,
	)
	require.NotNil(suite.T, vendorAccount)

	// Register second TestHouse account
	secondTestHouse := utils.RandString()
	secondTestHouseAccount := test_dclauth.CreateAccount(
		suite,
		secondTestHouse,
		dclauthtypes.AccountRoles{dclauthtypes.TestHouse},
		uint64(1),
		aliceName,
		aliceAccount,
		jackName,
		jackAccount,
	)
	require.NotNil(suite.T, vendorAccount)

	// Publish model info
	pid := int32(tmrand.Uint16())
	firstModel := test_model.NewMsgCreateModel(vid, pid)
	firstModel.Creator = suite.GetAddress(vendorName).String()
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{firstModel}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Publish modelVersion
	sv := uint32(tmrand.Uint32())
	svs := utils.RandString()
	firstModelVersion := test_model.NewMsgCreateModelVersion(vid, pid, sv, svs)
	firstModelVersion.Creator = suite.GetAddress(vendorName).String()
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{firstModelVersion}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Query for unknown model test results
	_, err = GetTestResult(suite, vid, pid, sv)
	suite.AssertNotFound(err)

	// Query for empty test results
	testResults, _ = GetAllTestResults(suite)
	require.Equal(suite.T, 0, len(testResults))

	// Publish first testing result
	firstTestingResult := compliancetesttypes.MsgAddTestingResult{
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       sv,
		SoftwareVersionString: svs,
		TestResult:            "some test results 1",
		TestDate:              "2020-01-01T00:00:00Z",
		Signer:                testHouseAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&firstTestingResult}, testHouse, testHouseAccount)
	require.NoError(suite.T, err)

	// Check testing result is created
	receivedTestingResult1, _ := GetTestResult(suite, vid, pid, sv)
	require.Equal(suite.T, receivedTestingResult1.Vid, firstTestingResult.Vid)
	require.Equal(suite.T, receivedTestingResult1.Pid, firstTestingResult.Pid)
	require.Equal(suite.T, receivedTestingResult1.SoftwareVersion, firstTestingResult.SoftwareVersion)
	require.Equal(suite.T, receivedTestingResult1.SoftwareVersionString, firstTestingResult.SoftwareVersionString)
	require.Equal(suite.T, 1, len(receivedTestingResult1.Results))
	require.Equal(suite.T, receivedTestingResult1.Results[0].Vid, firstTestingResult.Vid)
	require.Equal(suite.T, receivedTestingResult1.Results[0].Pid, firstTestingResult.Pid)
	require.Equal(suite.T, receivedTestingResult1.Results[0].SoftwareVersion, firstTestingResult.SoftwareVersion)
	require.Equal(suite.T, receivedTestingResult1.Results[0].SoftwareVersionString, firstTestingResult.SoftwareVersionString)
	require.Equal(suite.T, receivedTestingResult1.Results[0].TestResult, firstTestingResult.TestResult)
	require.Equal(suite.T, receivedTestingResult1.Results[0].TestDate, firstTestingResult.TestDate)
	require.Equal(suite.T, receivedTestingResult1.Results[0].Owner, firstTestingResult.Signer)

	// Publish second testing result
	secondTestingResult := compliancetesttypes.MsgAddTestingResult{
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       sv,
		SoftwareVersionString: svs,
		TestResult:            "some test results 2",
		TestDate:              "2021-01-01T00:00:00Z",
		Signer:                secondTestHouseAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&secondTestingResult}, secondTestHouse, secondTestHouseAccount)
	require.NoError(suite.T, err)

	// Check testing result is created
	receivedTestingResult2, _ := GetTestResult(suite, vid, pid, sv)
	require.Equal(suite.T, receivedTestingResult2.Vid, firstTestingResult.Vid)
	require.Equal(suite.T, receivedTestingResult2.Pid, firstTestingResult.Pid)
	require.Equal(suite.T, receivedTestingResult2.SoftwareVersion, firstTestingResult.SoftwareVersion)
	require.Equal(suite.T, receivedTestingResult2.SoftwareVersionString, firstTestingResult.SoftwareVersionString)
	require.Equal(suite.T, 2, len(receivedTestingResult2.Results))

	require.Equal(suite.T, receivedTestingResult2.Results[0].Vid, firstTestingResult.Vid)
	require.Equal(suite.T, receivedTestingResult2.Results[0].Pid, firstTestingResult.Pid)
	require.Equal(suite.T, receivedTestingResult2.Results[0].SoftwareVersion, firstTestingResult.SoftwareVersion)
	require.Equal(suite.T, receivedTestingResult2.Results[0].SoftwareVersionString, firstTestingResult.SoftwareVersionString)
	require.Equal(suite.T, receivedTestingResult2.Results[0].TestResult, firstTestingResult.TestResult)
	require.Equal(suite.T, receivedTestingResult2.Results[0].TestDate, firstTestingResult.TestDate)
	require.Equal(suite.T, receivedTestingResult2.Results[0].Owner, firstTestingResult.Signer)

	require.Equal(suite.T, receivedTestingResult2.Results[1].Vid, secondTestingResult.Vid)
	require.Equal(suite.T, receivedTestingResult2.Results[1].Pid, secondTestingResult.Pid)
	require.Equal(suite.T, receivedTestingResult2.Results[1].SoftwareVersion, secondTestingResult.SoftwareVersion)
	require.Equal(suite.T, receivedTestingResult2.Results[1].SoftwareVersionString, secondTestingResult.SoftwareVersionString)
	require.Equal(suite.T, receivedTestingResult2.Results[1].TestResult, secondTestingResult.TestResult)
	require.Equal(suite.T, receivedTestingResult2.Results[1].TestDate, secondTestingResult.TestDate)
	require.Equal(suite.T, receivedTestingResult2.Results[1].Owner, secondTestingResult.Signer)

	// Publish second model info
	new_pid := int32(tmrand.Uint16())
	secondModel := test_model.NewMsgCreateModel(vid, new_pid)
	secondModel.Creator = vendorAccount.Address
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{secondModel}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Publish second modelVersion
	new_sv := uint32(tmrand.Uint32())
	new_svs := utils.RandString()
	secondModelVersion := test_model.NewMsgCreateModelVersion(vid, new_pid, new_sv, new_svs)
	secondModelVersion.Creator = vendorAccount.Address
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{secondModelVersion}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Publish new testing result
	secondModelTestingResult := compliancetesttypes.MsgAddTestingResult{
		Vid:                   vid,
		Pid:                   new_pid,
		SoftwareVersion:       new_sv,
		SoftwareVersionString: new_svs,
		TestResult:            "some test results 3",
		TestDate:              "2020-06-06T00:00:00Z",
		Signer:                testHouseAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&secondModelTestingResult}, testHouse, testHouseAccount)
	require.NoError(suite.T, err)

	// Check testing result is created
	receivedTestingResult3, _ := GetTestResult(suite, vid, new_pid, new_sv)
	require.Equal(suite.T, receivedTestingResult3.Vid, secondModelTestingResult.Vid)
	require.Equal(suite.T, receivedTestingResult3.Pid, secondModelTestingResult.Pid)
	require.Equal(suite.T, receivedTestingResult3.SoftwareVersion, secondModelTestingResult.SoftwareVersion)
	require.Equal(suite.T, receivedTestingResult3.SoftwareVersionString, secondModelTestingResult.SoftwareVersionString)
	require.Equal(suite.T, 1, len(receivedTestingResult3.Results))
	require.Equal(suite.T, receivedTestingResult3.Results[0].Vid, secondModelTestingResult.Vid)
	require.Equal(suite.T, receivedTestingResult3.Results[0].Pid, secondModelTestingResult.Pid)
	require.Equal(suite.T, receivedTestingResult3.Results[0].SoftwareVersion, secondModelTestingResult.SoftwareVersion)
	require.Equal(suite.T, receivedTestingResult3.Results[0].SoftwareVersionString, secondModelTestingResult.SoftwareVersionString)
	require.Equal(suite.T, receivedTestingResult3.Results[0].TestResult, secondModelTestingResult.TestResult)
	require.Equal(suite.T, receivedTestingResult3.Results[0].TestDate, secondModelTestingResult.TestDate)
	require.Equal(suite.T, receivedTestingResult3.Results[0].Owner, secondModelTestingResult.Signer)

	// Query all test results
	testResults, _ = GetAllTestResults(suite)
	require.Equal(suite.T, 2, len(testResults))

	require.Equal(suite.T, testResults[0].Vid, firstTestingResult.Vid)
	require.Equal(suite.T, testResults[0].Pid, firstTestingResult.Pid)
	require.Equal(suite.T, testResults[0].SoftwareVersion, firstTestingResult.SoftwareVersion)
	require.Equal(suite.T, testResults[0].SoftwareVersionString, firstTestingResult.SoftwareVersionString)
	require.Equal(suite.T, 2, len(testResults[0].Results))

	require.Equal(suite.T, testResults[0].Results[0].Vid, firstTestingResult.Vid)
	require.Equal(suite.T, testResults[0].Results[0].Pid, firstTestingResult.Pid)
	require.Equal(suite.T, testResults[0].Results[0].SoftwareVersion, firstTestingResult.SoftwareVersion)
	require.Equal(suite.T, testResults[0].Results[0].SoftwareVersionString, firstTestingResult.SoftwareVersionString)
	require.Equal(suite.T, testResults[0].Results[0].TestResult, firstTestingResult.TestResult)
	require.Equal(suite.T, testResults[0].Results[0].TestDate, firstTestingResult.TestDate)
	require.Equal(suite.T, testResults[0].Results[0].Owner, firstTestingResult.Signer)

	require.Equal(suite.T, testResults[0].Results[1].Vid, secondTestingResult.Vid)
	require.Equal(suite.T, testResults[0].Results[1].Pid, secondTestingResult.Pid)
	require.Equal(suite.T, testResults[0].Results[1].SoftwareVersion, secondTestingResult.SoftwareVersion)
	require.Equal(suite.T, testResults[0].Results[1].SoftwareVersionString, secondTestingResult.SoftwareVersionString)
	require.Equal(suite.T, testResults[0].Results[1].TestResult, secondTestingResult.TestResult)
	require.Equal(suite.T, testResults[0].Results[1].TestDate, secondTestingResult.TestDate)
	require.Equal(suite.T, testResults[0].Results[1].Owner, secondTestingResult.Signer)

	require.Equal(suite.T, testResults[1].Vid, secondModelTestingResult.Vid)
	require.Equal(suite.T, testResults[1].Pid, secondModelTestingResult.Pid)
	require.Equal(suite.T, testResults[1].SoftwareVersion, secondModelTestingResult.SoftwareVersion)
	require.Equal(suite.T, testResults[1].SoftwareVersionString, secondModelTestingResult.SoftwareVersionString)
	require.Equal(suite.T, 1, len(testResults[1].Results))
	require.Equal(suite.T, testResults[1].Results[0].Vid, secondModelTestingResult.Vid)
	require.Equal(suite.T, testResults[1].Results[0].Pid, secondModelTestingResult.Pid)
	require.Equal(suite.T, testResults[1].Results[0].SoftwareVersion, secondModelTestingResult.SoftwareVersion)
	require.Equal(suite.T, testResults[1].Results[0].SoftwareVersionString, secondModelTestingResult.SoftwareVersionString)
	require.Equal(suite.T, testResults[1].Results[0].TestResult, secondModelTestingResult.TestResult)
	require.Equal(suite.T, testResults[1].Results[0].TestDate, secondModelTestingResult.TestDate)
	require.Equal(suite.T, testResults[1].Results[0].Owner, secondModelTestingResult.Signer)

}
