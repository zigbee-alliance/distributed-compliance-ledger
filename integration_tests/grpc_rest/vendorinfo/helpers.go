package vendorinfo

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	test_dclauth "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/grpc_rest/dclauth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	vendorinfotypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
)

func NewMsgCreateVendorInfo(vid int32, signer string) *vendorinfotypes.MsgCreateVendorInfo {
	return &vendorinfotypes.MsgCreateVendorInfo{
		Creator:              signer,
		Vid:                  vid,
		VendorName:           testconstants.VendorName,
		CompanyLegalName:     testconstants.CompanyLegalName,
		CompanyPrefferedName: testconstants.CompanyPreferredName,
		VendorLandingPageUrl: testconstants.VendorLandingPageUrl,
	}
}

func NewMsgUpdateVendorInfo(vid int32, signer string) *vendorinfotypes.MsgUpdateVendorInfo {
	return &vendorinfotypes.MsgUpdateVendorInfo{
		Creator:              signer,
		Vid:                  vid,
		VendorName:           testconstants.VendorName + "/new",
		CompanyLegalName:     testconstants.CompanyLegalName + "/new",
		CompanyPrefferedName: testconstants.CompanyPreferredName + "/new",
		VendorLandingPageUrl: testconstants.VendorLandingPageUrl + "/new",
	}
}

func AddVendorInfo(
	suite *utils.TestSuite,
	msg *vendorinfotypes.MsgCreateVendorInfo,
	signerName string,
	signerAccount *dclauthtypes.Account,
) (*sdk.TxResponse, error) {
	msg.Creator = suite.GetAddress(signerName).String()
	return suite.BuildAndBroadcastTx([]sdk.Msg{msg}, signerName, signerAccount)
}

func GetVendorInfo(
	suite *utils.TestSuite,
	vid int32,
) (*vendorinfotypes.VendorInfo, error) {
	var res vendorinfotypes.VendorInfo

	if suite.Rest {
		var resp vendorinfotypes.QueryGetVendorInfoResponse
		err := suite.QueryREST(fmt.Sprintf("/dcl/vendorinfo/vendors/%v", vid), &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetVendorInfo()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/dclauth service.
		vendorinfoClient := vendorinfotypes.NewQueryClient(grpcConn)
		resp, err := vendorinfoClient.VendorInfo(
			context.Background(),
			&vendorinfotypes.QueryGetVendorInfoRequest{Vid: vid},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetVendorInfo()
	}

	return &res, nil
}

func GetVendorInfos(suite *utils.TestSuite) (res []vendorinfotypes.VendorInfo, err error) {
	if suite.Rest {
		var resp vendorinfotypes.QueryAllVendorInfoResponse
		err := suite.QueryREST("/dcl/vendorinfo/vendors", &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetVendorInfo()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/dclauth service.
		vendorinfoClient := vendorinfotypes.NewQueryClient(grpcConn)
		resp, err := vendorinfoClient.VendorInfoAll(
			context.Background(),
			&vendorinfotypes.QueryAllVendorInfoRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetVendorInfo()
	}

	return res, nil
}

func VendorInfoDemo(suite *utils.TestSuite) {
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

	// Register new Vendor account
	vid := int32(tmrand.Uint16())
	vendorName := utils.RandString()
	vendorAccount := test_dclauth.CreateAccount(
		suite,
		vendorName,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		vid,
		aliceName,
		aliceAccount,
		bobName,
		bobAccount,
		testconstants.Info,
	)
	require.NotNil(suite.T, vendorAccount)

	// New vendor adds first vendorinfo
	createFirstVendorInfoMsg := NewMsgCreateVendorInfo(vid, vendorAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{createFirstVendorInfoMsg}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Check first vendorinfo is added
	receivedVendorInfo, err := GetVendorInfo(suite, createFirstVendorInfoMsg.Vid)
	require.NoError(suite.T, err)
	require.Equal(suite.T, createFirstVendorInfoMsg.Vid, receivedVendorInfo.Vid)
	require.Equal(suite.T, createFirstVendorInfoMsg.VendorName, receivedVendorInfo.VendorName)
	require.Equal(suite.T, createFirstVendorInfoMsg.CompanyLegalName, receivedVendorInfo.CompanyLegalName)
	require.Equal(suite.T, createFirstVendorInfoMsg.CompanyLegalName, receivedVendorInfo.CompanyLegalName)
	require.Equal(suite.T, createFirstVendorInfoMsg.VendorLandingPageUrl, receivedVendorInfo.VendorLandingPageUrl)

	// Get all vendorinfos
	_, err = GetVendorInfos(suite)
	require.NoError(suite.T, err)
}

/* Error cases */

func AddVendorInfoByNonVendor(suite *utils.TestSuite) {
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

	// register new account without Vendor role
	nonVendorAccountNamew := utils.RandString()
	vid := int32(tmrand.Uint16())
	nonVendorAccount := test_dclauth.CreateAccount(
		suite,
		nonVendorAccountNamew,
		dclauthtypes.AccountRoles{dclauthtypes.CertificationCenter},
		vid,
		aliceName,
		aliceAccount,
		bobName,
		bobAccount,
		testconstants.Info,
	)

	require.NotContains(suite.T, nonVendorAccount.Roles, dclauthtypes.Vendor)

	// try to add createVendorInfoMsg
	createVendorInfoMsg := NewMsgCreateVendorInfo(vid, nonVendorAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{createVendorInfoMsg}, nonVendorAccountNamew, nonVendorAccount)
	require.Error(suite.T, err)
	require.True(suite.T, sdkerrors.ErrUnauthorized.Is(err))
}

func AddVendorInfoByDifferentVendor(suite *utils.TestSuite) {
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

	// register new Vendor account
	vendorName := utils.RandString()
	vid := int32(tmrand.Uint16())
	vendorAccount := test_dclauth.CreateAccount(
		suite,
		vendorName,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		vid+1,
		aliceName,
		aliceAccount,
		bobName,
		bobAccount,
		testconstants.Info,
	)

	// try to add createVendorInfoMsg
	createVendorInfoMsg := NewMsgCreateVendorInfo(vid, vendorAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{createVendorInfoMsg}, vendorName, vendorAccount)
	require.Error(suite.T, err)
	require.True(suite.T, sdkerrors.ErrUnauthorized.Is(err))
}

func AddVendorInfoTwice(suite *utils.TestSuite) {
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

	// register new Vendor account
	vendorName := utils.RandString()
	vid := int32(tmrand.Uint16())
	vendorAccount := test_dclauth.CreateAccount(
		suite,
		vendorName,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		vid,
		aliceName,
		aliceAccount,
		bobName,
		bobAccount,
		testconstants.Info,
	)

	// add vendorinfo
	createVendorInfoMsg := NewMsgCreateVendorInfo(vid, vendorAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{createVendorInfoMsg}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// add the same vendorinfo second time
	_, err = AddVendorInfo(suite, createVendorInfoMsg, vendorName, vendorAccount)
	require.Error(suite.T, err)
	require.True(suite.T, sdkerrors.ErrInvalidRequest.Is(err))
}

func GetVendorInfoForUnknown(suite *utils.TestSuite) {
	_, err := GetVendorInfo(suite, int32(tmrand.Uint16()))
	require.Error(suite.T, err)
	suite.AssertNotFound(err)
}

func GetVendorInfoForInvalidVid(suite *utils.TestSuite) {
	// zero vid
	_, err := GetVendorInfo(suite, 0)
	require.Error(suite.T, err)
	// FIXME: Consider adding validation for queries.
	// require.True(suite.T, sdkerrors.ErrInvalidRequest.Is(err))
	suite.AssertNotFound(err)
}
