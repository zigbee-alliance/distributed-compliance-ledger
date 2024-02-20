package pki

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	test_dclauth "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/grpc_rest/dclauth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

//nolint:godox
/*
	To Run test you need:
		* Run LocalNet with: `make install && make localnet_init && make localnet_start`

	TODO: provide tests for error cases
*/

func GetAllNocX509RootCerts(suite *utils.TestSuite) (res []pkitypes.NocRootCertificates, err error) {
	if suite.Rest {
		var resp pkitypes.QueryAllNocRootCertificatesResponse
		err := suite.QueryREST("/dcl/pki/noc-root-certificates/", &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetNocRootCertificates()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.NocRootCertificatesAll(
			context.Background(),
			&pkitypes.QueryAllNocRootCertificatesRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetNocRootCertificates()
	}

	return res, nil
}

func GetNocX509RootCerts(suite *utils.TestSuite, vendorID int32) (*pkitypes.NocRootCertificates, error) {
	var res pkitypes.NocRootCertificates
	if suite.Rest {
		var resp pkitypes.QueryGetNocRootCertificatesResponse
		err := suite.QueryREST(fmt.Sprintf("/dcl/pki/noc-root-certificates/%v", vendorID), &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetNocRootCertificates()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.NocRootCertificates(
			context.Background(),
			&pkitypes.QueryGetNocRootCertificatesRequest{Vid: vendorID},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetNocRootCertificates()
	}

	return &res, nil
}

//nolint:funlen
func NocCertDemo(suite *utils.TestSuite) {
	// All requests return empty or 404 value
	allNocCertificates, _ := GetAllNocX509RootCerts(suite)
	require.Equal(suite.T, 0, len(allNocCertificates))

	_, err := GetNocX509RootCerts(suite, testconstants.Vid)
	suite.AssertNotFound(err)

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

	// Register first Vendor account
	vid1 := int32(tmrand.Uint16())
	vendor1Name := utils.RandString()
	vendor1Account := test_dclauth.CreateVendorAccount(
		suite,
		vendor1Name,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		vid1,
		testconstants.ProductIDsEmpty,
		aliceName,
		aliceAccount,
		jackName,
		jackAccount,
		testconstants.Info,
	)
	require.NotNil(suite.T, vendor1Account)

	// Register second Vendor account
	vid2 := int32(tmrand.Uint16())
	vendor2Name := utils.RandString()
	vendor2Account := test_dclauth.CreateVendorAccount(
		suite,
		vendor2Name,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		vid2,
		testconstants.ProductIDsEmpty,
		aliceName,
		aliceAccount,
		jackName,
		jackAccount,
		testconstants.Info,
	)
	require.NotNil(suite.T, vendor2Account)

	// Try to add inermidiate cert
	msgAddNocRootCertificate := pkitypes.MsgAddNocX509RootCert{
		Signer: vendor1Account.Address,
		Cert:   testconstants.IntermediateCertPem,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddNocRootCertificate}, vendor1Name, vendor1Account)
	require.Error(suite.T, err)

	// Add first NOC certificate by first vendor
	msgAddNocRootCertificate = pkitypes.MsgAddNocX509RootCert{
		Signer: vendor1Account.Address,
		Cert:   testconstants.NocRootCert1,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddNocRootCertificate}, vendor1Name, vendor1Account)
	require.NoError(suite.T, err)

	// Add second NOC certificate by first vendor
	msgAddNocRootCertificate = pkitypes.MsgAddNocX509RootCert{
		Signer: vendor1Account.Address,
		Cert:   testconstants.NocRootCert2,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddNocRootCertificate}, vendor1Name, vendor1Account)
	require.NoError(suite.T, err)

	// Add third NOC certificate by second vendor
	msgAddNocRootCertificate = pkitypes.MsgAddNocX509RootCert{
		Signer: vendor2Account.Address,
		Cert:   testconstants.NocRootCert3,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddNocRootCertificate}, vendor2Name, vendor2Account)
	require.NoError(suite.T, err)

	// Request NOC root certificate by VID1
	nocCertificates, _ := GetNocX509RootCerts(suite, vid1)
	require.Equal(suite.T, 2, len(nocCertificates.Certs))
	require.Equal(suite.T, testconstants.NocRootCert1Subject, nocCertificates.Certs[0].Subject)
	require.Equal(suite.T, testconstants.NocRootCert1SubjectKeyID, nocCertificates.Certs[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.NocRootCert2Subject, nocCertificates.Certs[1].Subject)
	require.Equal(suite.T, testconstants.NocRootCert2SubjectKeyID, nocCertificates.Certs[1].SubjectKeyId)

	// Request All NOC root certificate
	allNocCertificates, _ = GetAllNocX509RootCerts(suite)
	require.Equal(suite.T, 2, len(allNocCertificates))

	var (
		certsWithVid1 []*pkitypes.Certificate
		certsWithVid2 []*pkitypes.Certificate
	)

	if allNocCertificates[0].Vid == vid1 {
		certsWithVid1 = allNocCertificates[0].Certs
		certsWithVid2 = allNocCertificates[1].Certs
	} else {
		certsWithVid1 = allNocCertificates[1].Certs
		certsWithVid2 = allNocCertificates[0].Certs
	}

	require.Equal(suite.T, 2, len(certsWithVid1))
	require.Equal(suite.T, 1, len(certsWithVid2))
	require.Equal(suite.T, testconstants.NocRootCert1Subject, certsWithVid1[0].Subject)
	require.Equal(suite.T, testconstants.NocRootCert1SubjectKeyID, certsWithVid1[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.NocRootCert2Subject, certsWithVid1[1].Subject)
	require.Equal(suite.T, testconstants.NocRootCert2SubjectKeyID, certsWithVid1[1].SubjectKeyId)
	require.Equal(suite.T, testconstants.NocRootCert3Subject, certsWithVid2[0].Subject)
	require.Equal(suite.T, testconstants.NocRootCert3SubjectKeyID, certsWithVid2[0].SubjectKeyId)

	// Request NOC root certificate by Subject and SubjectKeyID
	certificate, _ := GetX509Cert(suite, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(suite.T, testconstants.NocRootCert1Subject, certificate.Subject)
	require.Equal(suite.T, testconstants.NocRootCert1SubjectKeyID, certificate.SubjectKeyId)
	require.Equal(suite.T, testconstants.NocRootCert1SubjectAsText, certificate.Certs[0].SubjectAsText)
	require.Equal(suite.T, 1, len(certificate.Certs))
	require.Equal(suite.T, testconstants.NocRootCert1, certificate.Certs[0].PemCert)
	require.Equal(suite.T, vendor1Account.Address, certificate.Certs[0].Owner)
	require.True(suite.T, certificate.Certs[0].IsRoot)

	// Request NOC root certificate by Subject
	subjectCertificates, _ := GetAllX509CertsBySubject(suite, testconstants.NocRootCert1Subject)
	require.Equal(suite.T, testconstants.NocRootCert1Subject, subjectCertificates.Subject)
	require.Equal(suite.T, 1, len(subjectCertificates.SubjectKeyIds))
	require.Equal(suite.T, testconstants.NocRootCert1SubjectKeyID, subjectCertificates.SubjectKeyIds[0])

	// Request NOC root certificate by SubjectKeyID
	certsBySubjectKeyID, _ := GetAllX509certsBySubjectKeyID(suite, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(suite.T, 1, len(certsBySubjectKeyID))
	require.Equal(suite.T, 1, len(certsBySubjectKeyID[0].Certs))
	require.Equal(suite.T, testconstants.NocRootCert1SubjectKeyID, certsBySubjectKeyID[0].Certs[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.NocRootCert1Subject, certsBySubjectKeyID[0].Certs[0].Subject)
	require.Equal(suite.T, testconstants.NocRootCert1SubjectAsText, certsBySubjectKeyID[0].Certs[0].SubjectAsText)
}
