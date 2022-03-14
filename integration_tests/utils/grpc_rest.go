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

package utils

import (
	"bufio"
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/simapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/tx"

	//nolint:staticcheck
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"google.golang.org/grpc"
)

// NOTE
// cosmos's 'github.com/cosmos/cosmos-sdk/types/rest' provides Get and Post API
// but they are not convenient enough: http errors (like BadRequest when an entry is missed)
// are hidden there and body is returned in any case

type TestSuite struct {
	T              *testing.T
	EncodingConfig simappparams.EncodingConfig
	ChainID        string
	Kr             keyring.Keyring
	Txf            clienttx.Factory
	Rest           bool
}

func (suite *TestSuite) GetGRPCConn() *grpc.ClientConn {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		"127.0.0.1:26630",   // Or your gRPC server address.
		grpc.WithInsecure(), // The SDK doesn't support any transport security mechanism.
	)
	require.NoError(suite.T, err)
	return grpcConn
}

func SetupTest(t *testing.T, chainID string, rest bool) (suite TestSuite) {
	t.Helper()
	inBuf := bufio.NewReader(os.Stdin)

	// TODO issue 99: pass as an arg
	userHomeDir, err := os.UserHomeDir()
	require.NoError(t, err)

	homeDir := filepath.Join(userHomeDir, ".dcl")

	kr, _ := keyring.New(sdk.KeyringServiceName(), keyring.BackendTest, homeDir, inBuf)

	encConfig := simapp.MakeTestEncodingConfig()
	dclauthtypes.RegisterInterfaces(encConfig.InterfaceRegistry)

	txCfg := encConfig.TxConfig
	txf := clienttx.Factory{}.
		WithChainID(chainID).
		WithTxConfig(txCfg).
		WithSignMode(txCfg.SignModeHandler().DefaultMode()).
		WithKeybase(kr)

	return TestSuite{
		EncodingConfig: encConfig,
		T:              t,
		ChainID:        chainID,
		Kr:             kr,
		Txf:            txf,
		Rest:           rest,
	}
}

func (suite *TestSuite) GetAddress(uid string) sdk.AccAddress {
	signerInfo, err := suite.Kr.Key(uid)
	require.NoError(suite.T, err)
	return signerInfo.GetAddress()
}

// Generates Protobuf-encoded bytes.
func (suite *TestSuite) BuildTx(
	msgs []sdk.Msg, signer string, account *dclauthtypes.Account,
) []byte {
	txfc := suite.Txf

	require.NotEqual(suite.T, 0, account.GetAccountNumber())
	require.NotEqual(suite.T, 0, account.GetSequence())

	txfc = txfc.WithAccountNumber(account.GetAccountNumber()).WithSequence(account.GetSequence())

	txSigned, err := GenTx(
		txfc,
		suite.EncodingConfig.TxConfig,
		msgs,
		signer,
	)
	require.NoError(suite.T, err)
	err = account.SetSequence(account.GetSequence() + 1)
	require.NoError(suite.T, err)

	// Generated Protobuf-encoded bytes.
	txBytes, err := suite.EncodingConfig.TxConfig.TxEncoder()(txSigned)
	require.NoError(suite.T, err)

	return txBytes
}

func (suite *TestSuite) BroadcastTx(txBytes []byte) (*sdk.TxResponse, error) {
	var broadcastResp *tx.BroadcastTxResponse
	var err error

	body := tx.BroadcastTxRequest{
		Mode:    tx.BroadcastMode_BROADCAST_MODE_BLOCK,
		TxBytes: txBytes,
	}

	if suite.Rest {
		var _resp tx.BroadcastTxResponse

		bodyBytes, err := suite.EncodingConfig.Marshaler.MarshalJSON(&body)
		require.NoError(suite.T, err)

		respBytes, err := SendPostRequest("/cosmos/tx/v1beta1/txs", bodyBytes, "", "")
		if err != nil {
			return nil, err
		}
		require.NoError(suite.T, suite.EncodingConfig.Marshaler.UnmarshalJSON(respBytes, &_resp))
		broadcastResp = &_resp
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// Broadcast the tx via gRPC. We create a new client for the Protobuf Tx
		// service.
		txClient := tx.NewServiceClient(grpcConn)
		broadcastResp, err = txClient.BroadcastTx(context.Background(), &body)
		if err != nil {
			return nil, err
		}
	}

	resp := broadcastResp.TxResponse
	if resp.Code != 0 {
		err = sdkerrors.ABCIError(resp.Codespace, resp.Code, resp.RawLog)
		return nil, err
	}

	return resp, nil
}

func (suite *TestSuite) BuildAndBroadcastTx(
	msgs []sdk.Msg, signer string, account *dclauthtypes.Account,
) (*sdk.TxResponse, error) {
	// build Tx
	txBytes := suite.BuildTx(msgs, signer, account)
	// broadcast Tx
	return suite.BroadcastTx(txBytes)
}

func (suite *TestSuite) QueryREST(uri string, resp proto.Message) error {
	respBytes, err := SendGetRequest(uri)
	if err != nil {
		return err
	}
	require.NoError(suite.T, suite.EncodingConfig.Marshaler.UnmarshalJSON(respBytes, resp))
	return nil
}

func (suite *TestSuite) AssertNotFound(err error) {
	require.Error(suite.T, err)
	require.Contains(suite.T, err.Error(), "rpc error: code = NotFound desc = not found")

	if suite.Rest {
		var resterr *RESTError
		if !errors.As(err, &resterr) {
			panic("REST error is not RESTError type")
		}
		require.Equal(suite.T, resterr.resp.StatusCode, 404)
	}
}
