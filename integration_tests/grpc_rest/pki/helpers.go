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

package pki

import (
	"context"
	"fmt"
	"net/url"

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

func GetAllProposedX509RootCerts(suite *utils.TestSuite) (res []pkitypes.ProposedCertificate, err error) {
	if suite.Rest {
		var resp pkitypes.QueryAllProposedCertificateResponse
		err := suite.QueryREST("/dcl/pki/proposed-certificates", &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetProposedCertificate()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.ProposedCertificateAll(
			context.Background(),
			&pkitypes.QueryAllProposedCertificateRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetProposedCertificate()
	}

	return res, nil
}

func GetProposedX509RootCert(suite *utils.TestSuite, subject string, subjectKeyID string) (*pkitypes.ProposedCertificate, error) {
	var res pkitypes.ProposedCertificate
	if suite.Rest {
		var resp pkitypes.QueryGetProposedCertificateResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/pki/proposed-certificates/%s/%s",
				url.QueryEscape(subject), url.QueryEscape(subjectKeyID),
			),
			&resp,
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetProposedCertificate()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.ProposedCertificate(
			context.Background(),
			&pkitypes.QueryGetProposedCertificateRequest{
				Subject:      subject,
				SubjectKeyId: subjectKeyID,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetProposedCertificate()
	}

	return &res, nil
}

func GetAllX509Certs(suite *utils.TestSuite) (res []pkitypes.ApprovedCertificates, err error) {
	if suite.Rest {
		var resp pkitypes.QueryAllApprovedCertificatesResponse
		err := suite.QueryREST("/dcl/pki/certificates", &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetApprovedCertificates()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.ApprovedCertificatesAll(
			context.Background(),
			&pkitypes.QueryAllApprovedCertificatesRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetApprovedCertificates()
	}

	return res, nil
}

func GetX509Cert(suite *utils.TestSuite, subject string, subjectKeyID string) (*pkitypes.ApprovedCertificates, error) {
	var res pkitypes.ApprovedCertificates
	if suite.Rest {
		var resp pkitypes.QueryGetApprovedCertificatesResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/pki/certificates/%s/%s",
				url.QueryEscape(subject), url.QueryEscape(subjectKeyID),
			),
			&resp,
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetApprovedCertificates()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.ApprovedCertificates(
			context.Background(),
			&pkitypes.QueryGetApprovedCertificatesRequest{
				Subject:      subject,
				SubjectKeyId: subjectKeyID,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetApprovedCertificates()
	}

	return &res, nil
}

func GetAllRevokedX509Certs(suite *utils.TestSuite) (res []pkitypes.RevokedCertificates, err error) {
	if suite.Rest {
		var resp pkitypes.QueryAllRevokedCertificatesResponse
		err := suite.QueryREST("/dcl/pki/revoked-certificates", &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetRevokedCertificates()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.RevokedCertificatesAll(
			context.Background(),
			&pkitypes.QueryAllRevokedCertificatesRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetRevokedCertificates()
	}

	return res, nil
}

func GetRevokedX509Cert(suite *utils.TestSuite, subject string, subjectKeyID string) (*pkitypes.RevokedCertificates, error) {
	var res pkitypes.RevokedCertificates
	if suite.Rest {
		var resp pkitypes.QueryGetRevokedCertificatesResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/pki/revoked-certificates/%s/%s",
				url.QueryEscape(subject), url.QueryEscape(subjectKeyID),
			),
			&resp,
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetRevokedCertificates()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.RevokedCertificates(
			context.Background(),
			&pkitypes.QueryGetRevokedCertificatesRequest{
				Subject:      subject,
				SubjectKeyId: subjectKeyID,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetRevokedCertificates()
	}

	return &res, nil
}

func GetAllProposedRevocationX509Certs(suite *utils.TestSuite) (res []pkitypes.ProposedCertificateRevocation, err error) {
	if suite.Rest {
		var resp pkitypes.QueryAllProposedCertificateRevocationResponse
		err := suite.QueryREST("/dcl/pki/proposed-revocation-certificates", &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetProposedCertificateRevocation()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.ProposedCertificateRevocationAll(
			context.Background(),
			&pkitypes.QueryAllProposedCertificateRevocationRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetProposedCertificateRevocation()
	}

	return res, nil
}

func GetProposedRevocationX509Cert(suite *utils.TestSuite, subject string, subjectKeyID string) (*pkitypes.ProposedCertificateRevocation, error) {
	var res pkitypes.ProposedCertificateRevocation
	if suite.Rest {
		var resp pkitypes.QueryGetProposedCertificateRevocationResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/pki/proposed-revocation-certificates/%s/%s",
				url.QueryEscape(subject), url.QueryEscape(subjectKeyID),
			),
			&resp,
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetProposedCertificateRevocation()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.ProposedCertificateRevocation(
			context.Background(),
			&pkitypes.QueryGetProposedCertificateRevocationRequest{
				Subject:      subject,
				SubjectKeyId: subjectKeyID,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetProposedCertificateRevocation()
	}

	return &res, nil
}

func GetAllRootX509Certs(suite *utils.TestSuite) (res pkitypes.ApprovedRootCertificates, err error) {
	if suite.Rest {
		var resp pkitypes.QueryGetApprovedRootCertificatesResponse
		err := suite.QueryREST("/dcl/pki/root-certificates", &resp)
		if err != nil {
			return pkitypes.ApprovedRootCertificates{}, err
		}
		res = resp.GetApprovedRootCertificates()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.ApprovedRootCertificates(
			context.Background(),
			&pkitypes.QueryGetApprovedRootCertificatesRequest{},
		)
		if err != nil {
			return pkitypes.ApprovedRootCertificates{}, err
		}
		res = resp.GetApprovedRootCertificates()
	}

	return res, nil
}

func GetAllRevokedRootX509Certs(suite *utils.TestSuite) (res pkitypes.RevokedRootCertificates, err error) {
	if suite.Rest {
		var resp pkitypes.QueryGetRevokedRootCertificatesResponse
		err := suite.QueryREST("/dcl/pki/revoked-root-certificates", &resp)
		if err != nil {
			return pkitypes.RevokedRootCertificates{}, err
		}
		res = resp.GetRevokedRootCertificates()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.RevokedRootCertificates(
			context.Background(),
			&pkitypes.QueryGetRevokedRootCertificatesRequest{},
		)
		if err != nil {
			return pkitypes.RevokedRootCertificates{}, err
		}
		res = resp.GetRevokedRootCertificates()
	}

	return res, nil
}

func GetAllX509CertsBySubject(suite *utils.TestSuite, subject string) (*pkitypes.ApprovedCertificatesBySubject, error) {
	var res pkitypes.ApprovedCertificatesBySubject
	if suite.Rest {
		var resp pkitypes.QueryGetApprovedCertificatesBySubjectResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/pki/certificates/%s",
				url.QueryEscape(subject),
			),
			&resp,
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetApprovedCertificatesBySubject()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.ApprovedCertificatesBySubject(
			context.Background(),
			&pkitypes.QueryGetApprovedCertificatesBySubjectRequest{
				Subject: subject,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetApprovedCertificatesBySubject()
	}

	return &res, nil
}

func GetAllChildX509Certs(suite *utils.TestSuite, subject string, subjectKeyID string) (*pkitypes.ChildCertificates, error) {
	var res pkitypes.ChildCertificates
	if suite.Rest {
		var resp pkitypes.QueryGetChildCertificatesResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/pki/child-certificates/%s/%s",
				url.QueryEscape(subject), url.QueryEscape(subjectKeyID),
			),
			&resp,
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetChildCertificates()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.ChildCertificates(
			context.Background(),
			&pkitypes.QueryGetChildCertificatesRequest{
				Issuer:         subject,
				AuthorityKeyId: subjectKeyID,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetChildCertificates()
	}

	return &res, nil
}

func GetAllRejectedX509RootCerts(suite *utils.TestSuite) (res []pkitypes.RejectedCertificate, err error) {
	if suite.Rest {
		var resp pkitypes.QueryAllRejectedCertificatesResponse
		err := suite.QueryREST("dcl/pki/rejected-certificates", &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetRejectedCertificate()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.RejectedCertificateAll(
			context.Background(),
			&pkitypes.QueryAllRejectedCertificatesRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetRejectedCertificate()
	}

	return res, nil
}

func GetRejectedX509RootCert(suite *utils.TestSuite, subject string, subjectKeyID string) (*pkitypes.RejectedCertificate, error) {
	var res pkitypes.RejectedCertificate
	if suite.Rest {
		var resp pkitypes.QueryGetRejectedCertificatesResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"dcl/pki/rejected-certificates/%s/%s",
				url.QueryEscape(subject), url.QueryEscape(subjectKeyID),
			),
			&resp,
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetRejectedCertificate()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.RejectedCertificate(
			context.Background(),
			&pkitypes.QueryGetRejectedCertificatesRequest{
				Subject:      subject,
				SubjectKeyId: subjectKeyID,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetRejectedCertificate()
	}

	return &res, nil
}

//nolint:funlen
func Demo(suite *utils.TestSuite) {
	// All requests return empty or 404 value
	proposedCertificates, _ := GetAllProposedX509RootCerts(suite)
	require.Equal(suite.T, 0, len(proposedCertificates))

	_, err := GetProposedX509RootCert(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	suite.AssertNotFound(err)

	certificates, _ := GetAllX509Certs(suite)
	require.Equal(suite.T, 0, len(certificates))

	_, err = GetX509Cert(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	suite.AssertNotFound(err)

	revokedCertificates, _ := GetAllRevokedX509Certs(suite)
	require.Equal(suite.T, 0, len(revokedCertificates))

	_, err = GetRevokedX509Cert(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	suite.AssertNotFound(err)

	proposedRevocationCertificates, _ := GetAllProposedRevocationX509Certs(suite)
	require.Equal(suite.T, 0, len(proposedRevocationCertificates))

	_, err = GetProposedRevocationX509Cert(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	suite.AssertNotFound(err)

	rootCertificates, _ := GetAllRootX509Certs(suite)
	require.Equal(suite.T, 0, len(rootCertificates.Certs))

	revokedRootCertificates, _ := GetAllRevokedRootX509Certs(suite)
	require.Equal(suite.T, 0, len(revokedRootCertificates.Certs))

	_, err = GetAllChildX509Certs(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	suite.AssertNotFound(err)

	_, err = GetAllX509CertsBySubject(suite, testconstants.RootSubject)
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

	// Register new Vendor account
	vid := int32(tmrand.Uint16())
	vendorName := utils.RandString()
	vendorAccount := test_dclauth.CreateVendorAccount(
		suite,
		vendorName,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		vid,
		aliceName,
		aliceAccount,
		jackName,
		jackAccount,
		testconstants.Info,
	)
	require.NotNil(suite.T, vendorAccount)

	// Vendor (Not Trustee) propose Root certificate
	msgProposeAddX509RootCert := pkitypes.MsgProposeAddX509RootCert{
		Cert:   testconstants.RootCertPem,
		Signer: vendorAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgProposeAddX509RootCert}, vendorName, vendorAccount)
	require.Error(suite.T, err)

	// Jack (Trustee) propose Root certificate
	msgProposeAddX509RootCert = pkitypes.MsgProposeAddX509RootCert{
		Cert:   testconstants.RootCertPem,
		Signer: jackAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgProposeAddX509RootCert}, jackName, jackAccount)
	require.NoError(suite.T, err)

	// Request all proposed Root certificates
	proposedCertificates, _ = GetAllProposedX509RootCerts(suite)
	require.Equal(suite.T, 1, len(proposedCertificates))
	require.Equal(suite.T, testconstants.RootSubject, proposedCertificates[0].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, proposedCertificates[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.RootSubjectAsText, proposedCertificates[0].SubjectAsText)

	// Request all approved certificates
	certificates, _ = GetAllX509Certs(suite)
	require.Equal(suite.T, 0, len(certificates))

	// Request proposed Root certificate
	proposedCertificate, _ := GetProposedX509RootCert(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(suite.T, testconstants.RootCertPem, proposedCertificate.PemCert)
	require.Equal(suite.T, jackAccount.Address, proposedCertificate.Owner)
	require.Equal(suite.T, 1, len(proposedCertificate.Approvals))

	// Request all proposed Root certificates
	proposedCertificates, _ = GetAllProposedX509RootCerts(suite)
	require.Equal(suite.T, 1, len(proposedCertificates))
	require.Equal(suite.T, testconstants.RootSubject, proposedCertificates[0].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, proposedCertificates[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.RootSubjectAsText, proposedCertificates[0].SubjectAsText)

	// Request all approved certificates
	certificates, _ = GetAllX509Certs(suite)
	require.Equal(suite.T, 0, len(certificates))

	// Request proposed Root certificate
	proposedCertificate, _ = GetProposedX509RootCert(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(suite.T, testconstants.RootCertPem, proposedCertificate.PemCert)
	require.Equal(suite.T, jackAccount.Address, proposedCertificate.Owner)
	require.True(suite.T, proposedCertificate.HasApprovalFrom(jackAccount.Address))

	// Alice (Trustee) approve Root certificate
	secondMsgApproveAddX509RootCert := pkitypes.MsgApproveAddX509RootCert{
		Subject:      proposedCertificate.Subject,
		SubjectKeyId: proposedCertificate.SubjectKeyId,
		Signer:       aliceAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&secondMsgApproveAddX509RootCert}, aliceName, aliceAccount)
	require.NoError(suite.T, err)

	// Request all proposed Root certificates
	proposedCertificates, _ = GetAllProposedX509RootCerts(suite)
	require.Equal(suite.T, 0, len(proposedCertificates))

	// Request all approved certificates
	certificates, _ = GetAllX509Certs(suite)
	require.Equal(suite.T, 1, len(certificates))
	require.Equal(suite.T, testconstants.RootSubject, certificates[0].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, certificates[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.RootSubjectAsText, certificates[0].Certs[0].SubjectAsText)

	// Request all approved Root certificates
	rootCertificates, _ = GetAllRootX509Certs(suite)
	require.Equal(suite.T, 1, len(rootCertificates.Certs))
	require.Equal(suite.T, testconstants.RootSubject, rootCertificates.Certs[0].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, rootCertificates.Certs[0].SubjectKeyId)

	// Request approved Root certificate
	certificate, _ := GetX509Cert(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(suite.T, testconstants.RootSubject, certificate.Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, certificate.SubjectKeyId)
	require.Equal(suite.T, testconstants.RootSubjectAsText, certificate.Certs[0].SubjectAsText)
	require.Equal(suite.T, 1, len(certificate.Certs))
	require.Equal(suite.T, testconstants.RootCertPem, certificate.Certs[0].PemCert)
	require.Equal(suite.T, jackAccount.Address, certificate.Certs[0].Owner)
	require.True(suite.T, certificate.Certs[0].IsRoot)

	// User (Not Trustee) add Intermediate certificate
	msgAddX509Cert := pkitypes.MsgAddX509Cert{
		Cert:   testconstants.IntermediateCertPem,
		Signer: vendorAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddX509Cert}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Request all proposed Root certificates
	proposedCertificates, _ = GetAllProposedX509RootCerts(suite)
	require.Equal(suite.T, 0, len(proposedCertificates))

	// Request all approved certificates
	certificates, _ = GetAllX509Certs(suite)
	require.Equal(suite.T, 2, len(certificates))
	require.Equal(suite.T, testconstants.RootSubject, certificates[0].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, certificates[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.RootSubjectAsText, certificates[0].Certs[0].SubjectAsText)
	require.Equal(suite.T, testconstants.IntermediateSubject, certificates[1].Subject)
	require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, certificates[1].SubjectKeyId)
	require.Equal(suite.T, testconstants.IntermediateSubjectAsText, certificates[1].Certs[0].SubjectAsText)

	// Request all approved Root certificates
	rootCertificates, _ = GetAllRootX509Certs(suite)
	require.Equal(suite.T, 1, len(rootCertificates.Certs))
	require.Equal(suite.T, testconstants.RootSubject, rootCertificates.Certs[0].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, rootCertificates.Certs[0].SubjectKeyId)

	// Request Intermediate certificate
	certificate, _ = GetX509Cert(suite, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Equal(suite.T, testconstants.IntermediateSubject, certificate.Subject)
	require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, certificate.SubjectKeyId)
	require.Equal(suite.T, testconstants.IntermediateSubjectAsText, certificate.Certs[0].SubjectAsText)
	require.Equal(suite.T, 1, len(certificate.Certs))
	require.Equal(suite.T, testconstants.IntermediateCertPem, certificate.Certs[0].PemCert)
	require.Equal(suite.T, vendorAccount.Address, certificate.Certs[0].Owner)
	require.False(suite.T, certificate.Certs[0].IsRoot)

	// Alice (Trustee) add Leaf certificate
	secondMsgAddX509Cert := pkitypes.MsgAddX509Cert{
		Cert:   testconstants.LeafCertPem,
		Signer: aliceAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&secondMsgAddX509Cert}, aliceName, aliceAccount)
	require.NoError(suite.T, err)

	// Request all proposed Root certificates
	proposedCertificates, _ = GetAllProposedX509RootCerts(suite)
	require.Equal(suite.T, 0, len(proposedCertificates))

	// Request all approved certificates
	certificates, _ = GetAllX509Certs(suite)
	require.Equal(suite.T, 3, len(certificates))
	require.Equal(suite.T, testconstants.LeafSubject, certificates[0].Subject)
	require.Equal(suite.T, testconstants.LeafSubjectKeyID, certificates[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.LeafSubjectAsText, certificates[0].Certs[0].SubjectAsText)
	require.Equal(suite.T, testconstants.RootSubject, certificates[1].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, certificates[1].SubjectKeyId)
	require.Equal(suite.T, testconstants.RootSubjectAsText, certificates[1].Certs[0].SubjectAsText)
	require.Equal(suite.T, testconstants.IntermediateSubject, certificates[2].Subject)
	require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, certificates[2].SubjectKeyId)
	require.Equal(suite.T, testconstants.IntermediateSubjectAsText, certificates[2].Certs[0].SubjectAsText)

	// Request all approved Root certificates
	rootCertificates, _ = GetAllRootX509Certs(suite)
	require.Equal(suite.T, 1, len(rootCertificates.Certs))
	require.Equal(suite.T, testconstants.RootSubject, rootCertificates.Certs[0].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, rootCertificates.Certs[0].SubjectKeyId)

	// Request Leaf certificate
	certificate, _ = GetX509Cert(suite, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Equal(suite.T, testconstants.LeafSubject, certificate.Subject)
	require.Equal(suite.T, testconstants.LeafSubjectKeyID, certificate.SubjectKeyId)
	require.Equal(suite.T, 1, len(certificate.Certs))
	require.Equal(suite.T, testconstants.LeafCertPem, certificate.Certs[0].PemCert)
	require.Equal(suite.T, aliceAccount.Address, certificate.Certs[0].Owner)
	require.Equal(suite.T, testconstants.LeafSubject, certificate.Certs[0].Subject)
	require.Equal(suite.T, testconstants.LeafSubjectKeyID, certificate.Certs[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.LeafSubjectAsText, certificate.Certs[0].SubjectAsText)
	require.False(suite.T, certificate.Certs[0].IsRoot)

	// Request all Subject certificates
	subjectCertificates, _ := GetAllX509CertsBySubject(suite, testconstants.LeafSubject)
	require.Equal(suite.T, testconstants.LeafSubject, subjectCertificates.Subject)
	require.Equal(suite.T, 1, len(subjectCertificates.SubjectKeyIds))
	require.Equal(suite.T, testconstants.LeafSubjectKeyID, subjectCertificates.SubjectKeyIds[0])

	subjectCertificates, _ = GetAllX509CertsBySubject(suite, testconstants.IntermediateSubject)
	require.Equal(suite.T, testconstants.IntermediateSubject, subjectCertificates.Subject)
	require.Equal(suite.T, 1, len(subjectCertificates.SubjectKeyIds))
	require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, subjectCertificates.SubjectKeyIds[0])

	subjectCertificates, _ = GetAllX509CertsBySubject(suite, testconstants.RootSubject)
	require.Equal(suite.T, testconstants.RootSubject, subjectCertificates.Subject)
	require.Equal(suite.T, 1, len(subjectCertificates.SubjectKeyIds))
	require.Equal(suite.T, testconstants.RootSubjectKeyID, subjectCertificates.SubjectKeyIds[0])

	// Request all Root certificates proposed to revoke
	proposedRevocationCertificates, _ = GetAllProposedRevocationX509Certs(suite)
	require.Equal(suite.T, 0, len(proposedRevocationCertificates))

	// Request all revoked certificates
	revokedCertificates, _ = GetAllRevokedX509Certs(suite)
	require.Equal(suite.T, 0, len(revokedCertificates))

	// Request all revoked Root certificates
	revokedRootCertificates, _ = GetAllRevokedRootX509Certs(suite)
	require.Equal(suite.T, 0, len(revokedRootCertificates.Certs))

	// Get all child certificates
	childCertificates, _ := GetAllChildX509Certs(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(suite.T, testconstants.RootSubject, childCertificates.Issuer)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, childCertificates.AuthorityKeyId)
	require.Equal(suite.T, testconstants.IntermediateSubject, childCertificates.CertIds[0].Subject)
	require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, childCertificates.CertIds[0].SubjectKeyId)

	childCertificates, _ = GetAllChildX509Certs(suite, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Equal(suite.T, testconstants.IntermediateSubject, childCertificates.Issuer)
	require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, childCertificates.AuthorityKeyId)
	require.Equal(suite.T, 1, len(childCertificates.CertIds))
	require.Equal(suite.T, testconstants.LeafSubject, childCertificates.CertIds[0].Subject)
	require.Equal(suite.T, testconstants.LeafSubjectKeyID, childCertificates.CertIds[0].SubjectKeyId)

	_, err = GetAllChildX509Certs(suite, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	suite.AssertNotFound(err)

	// User (Not Trustee) revokes Intermediate certificate. This must also revoke its child - Leaf certificate.
	msgRevokeX509Cert := pkitypes.MsgRevokeX509Cert{
		Subject:      testconstants.IntermediateSubject,
		SubjectKeyId: testconstants.IntermediateSubjectKeyID,
		Signer:       vendorAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRevokeX509Cert}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Request all Root certificates proposed to revoke
	proposedRevocationCertificates, _ = GetAllProposedRevocationX509Certs(suite)
	require.Equal(suite.T, 0, len(proposedRevocationCertificates))

	// Request all revoked certificates
	revokedCertificates, _ = GetAllRevokedX509Certs(suite)
	require.Equal(suite.T, 2, len(revokedCertificates))
	require.Equal(suite.T, testconstants.LeafSubject, revokedCertificates[0].Subject)
	require.Equal(suite.T, testconstants.LeafSubjectKeyID, revokedCertificates[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.LeafSubjectAsText, revokedCertificates[0].Certs[0].SubjectAsText)
	require.Equal(suite.T, testconstants.IntermediateSubject, revokedCertificates[1].Subject)
	require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, revokedCertificates[1].SubjectKeyId)
	require.Equal(suite.T, testconstants.IntermediateSubjectAsText, revokedCertificates[1].Certs[0].SubjectAsText)

	// Request all revoked Root certificates
	revokedRootCertificates, _ = GetAllRevokedRootX509Certs(suite)
	require.Equal(suite.T, 0, len(revokedRootCertificates.Certs))

	// Request revoked Intermediate certificate
	revokedCertificate, _ := GetRevokedX509Cert(suite, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Equal(suite.T, 1, len(revokedCertificate.Certs))
	require.Equal(suite.T, testconstants.IntermediateSubject, revokedCertificate.Subject)
	require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, revokedCertificate.SubjectKeyId)
	require.Equal(suite.T, testconstants.IntermediateSubject, revokedCertificate.Certs[0].Subject)
	require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, revokedCertificate.Certs[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.IntermediateCertPem, revokedCertificate.Certs[0].PemCert)
	require.Equal(suite.T, vendorAccount.Address, revokedCertificate.Certs[0].Owner)
	require.Equal(suite.T, testconstants.IntermediateSubjectAsText, revokedCertificate.Certs[0].SubjectAsText)
	require.False(suite.T, revokedCertificate.Certs[0].IsRoot)

	// Request revoked Leaf certificate
	revokedCertificate, _ = GetRevokedX509Cert(suite, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Equal(suite.T, 1, len(revokedCertificate.Certs))
	require.Equal(suite.T, testconstants.LeafSubject, revokedCertificate.Subject)
	require.Equal(suite.T, testconstants.LeafSubjectKeyID, revokedCertificate.SubjectKeyId)
	require.Equal(suite.T, testconstants.LeafSubject, revokedCertificate.Certs[0].Subject)
	require.Equal(suite.T, testconstants.LeafSubjectKeyID, revokedCertificate.Certs[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.LeafCertPem, revokedCertificate.Certs[0].PemCert)
	require.Equal(suite.T, aliceAccount.Address, revokedCertificate.Certs[0].Owner)
	require.Equal(suite.T, testconstants.LeafSubjectAsText, revokedCertificate.Certs[0].SubjectAsText)
	require.False(suite.T, revokedCertificate.Certs[0].IsRoot)

	// Request all approved certificates
	certificates, _ = GetAllX509Certs(suite)
	require.Equal(suite.T, 1, len(certificates))
	require.Equal(suite.T, testconstants.RootSubject, certificates[0].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, certificates[0].SubjectKeyId)

	// Jack (Trustee) proposes to revoke Root certificate
	msgProposeRevokeX509RootCert := pkitypes.MsgProposeRevokeX509RootCert{
		Subject:      testconstants.RootSubject,
		SubjectKeyId: testconstants.RootSubjectKeyID,
		Signer:       jackAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgProposeRevokeX509RootCert}, jackName, jackAccount)
	require.NoError(suite.T, err)

	// Request all Root certificates proposed to revoke
	proposedRevocationCertificates, _ = GetAllProposedRevocationX509Certs(suite)
	require.Equal(suite.T, 1, len(proposedRevocationCertificates))
	require.Equal(suite.T, testconstants.RootSubject, proposedRevocationCertificates[0].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, proposedRevocationCertificates[0].SubjectKeyId)

	// Request all revoked certificates
	revokedCertificates, _ = GetAllRevokedX509Certs(suite)
	require.Equal(suite.T, 2, len(revokedCertificates))
	require.Equal(suite.T, testconstants.LeafSubject, revokedCertificates[0].Subject)
	require.Equal(suite.T, testconstants.LeafSubjectKeyID, revokedCertificates[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.IntermediateSubject, revokedCertificates[1].Subject)
	require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, revokedCertificates[1].SubjectKeyId)

	// Request all revoked Root certificates
	revokedRootCertificates, _ = GetAllRevokedRootX509Certs(suite)
	require.Equal(suite.T, 0, len(revokedRootCertificates.Certs))

	// Request Root certificate proposed to revoke
	proposedCertificateRevocation, _ := GetProposedRevocationX509Cert(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(suite.T, testconstants.RootSubject, proposedCertificateRevocation.Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, proposedCertificateRevocation.SubjectKeyId)
	require.True(suite.T, proposedCertificateRevocation.HasRevocationFrom(jackAccount.Address))

	// Request all approved certificates
	certificates, _ = GetAllX509Certs(suite)
	require.Equal(suite.T, 1, len(certificates))
	require.Equal(suite.T, testconstants.RootSubject, certificates[0].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, certificates[0].SubjectKeyId)

	// Alice (Trustee) approves to revoke Root certificate
	msgApproveRevokeX509RootCert := pkitypes.MsgApproveRevokeX509RootCert{
		Subject:      proposedCertificate.Subject,
		SubjectKeyId: proposedCertificate.SubjectKeyId,
		Signer:       aliceAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgApproveRevokeX509RootCert}, aliceName, aliceAccount)
	require.NoError(suite.T, err)

	// Request all Root certificates proposed to revoke
	proposedRevocationCertificates, _ = GetAllProposedRevocationX509Certs(suite)
	require.Equal(suite.T, 0, len(proposedRevocationCertificates))

	// Request all revoked certificates
	revokedCertificates, _ = GetAllRevokedX509Certs(suite)
	require.Equal(suite.T, 3, len(revokedCertificates))
	require.Equal(suite.T, testconstants.LeafSubject, revokedCertificates[0].Subject)
	require.Equal(suite.T, testconstants.LeafSubjectKeyID, revokedCertificates[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.RootSubject, revokedCertificates[1].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, revokedCertificates[1].SubjectKeyId)
	require.Equal(suite.T, testconstants.IntermediateSubject, revokedCertificates[2].Subject)
	require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, revokedCertificates[2].SubjectKeyId)

	// Request all revoked Root certificates
	revokedRootCertificates, _ = GetAllRevokedRootX509Certs(suite)
	require.Equal(suite.T, 1, len(revokedRootCertificates.Certs))
	require.Equal(suite.T, testconstants.RootSubject, revokedRootCertificates.Certs[0].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, revokedRootCertificates.Certs[0].SubjectKeyId)

	// Request revoked Root certificate
	revokedCertificate, _ = GetRevokedX509Cert(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(suite.T, 1, len(revokedCertificate.Certs))
	require.Equal(suite.T, testconstants.RootSubject, revokedCertificate.Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, revokedCertificate.SubjectKeyId)
	require.Equal(suite.T, testconstants.RootSubject, revokedCertificate.Certs[0].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, revokedCertificate.Certs[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.RootCertPem, revokedCertificate.Certs[0].PemCert)
	require.Equal(suite.T, jackAccount.Address, revokedCertificate.Certs[0].Owner)
	require.Equal(suite.T, testconstants.RootSubjectAsText, revokedCertificate.Certs[0].SubjectAsText)
	require.True(suite.T, revokedCertificate.Certs[0].IsRoot)

	certificates, _ = GetAllX509Certs(suite)
	require.Equal(suite.T, 0, len(certificates))

	_, err = GetProposedX509RootCert(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	suite.AssertNotFound(err)

	_, err = GetX509Cert(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	suite.AssertNotFound(err)

	_, err = GetX509Cert(suite, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	suite.AssertNotFound(err)

	_, err = GetX509Cert(suite, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	suite.AssertNotFound(err)

	proposedRevocationCertificates, _ = GetAllProposedRevocationX509Certs(suite)
	require.Equal(suite.T, 0, len(proposedRevocationCertificates))

	_, err = GetProposedRevocationX509Cert(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	suite.AssertNotFound(err)

	rootCertificates, _ = GetAllRootX509Certs(suite)
	require.Equal(suite.T, 0, len(rootCertificates.Certs))

	_, err = GetAllChildX509Certs(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	suite.AssertNotFound(err)

	_, err = GetAllChildX509Certs(suite, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	suite.AssertNotFound(err)

	_, err = GetAllChildX509Certs(suite, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	suite.AssertNotFound(err)

	_, err = GetAllX509CertsBySubject(suite, testconstants.RootSubject)
	suite.AssertNotFound(err)

	_, err = GetAllX509CertsBySubject(suite, testconstants.IntermediateSubject)
	suite.AssertNotFound(err)

	_, err = GetAllX509CertsBySubject(suite, testconstants.LeafSubject)
	suite.AssertNotFound(err)

	// CHECK GOOGLE ROOT CERTIFICATE WHICH INCLUDES VID
	_, err = GetRevokedX509Cert(suite, testconstants.GoogleSubject, testconstants.GoogleSubjectKeyID)
	suite.AssertNotFound(err)

	_, err = GetProposedRevocationX509Cert(suite, testconstants.GoogleSubject, testconstants.GoogleSubjectKeyID)
	suite.AssertNotFound(err)

	// Alice (Trustee) propose Google Root certificate
	msgProposeAddX509GoogleRootcert := pkitypes.MsgProposeAddX509RootCert{
		Cert:   testconstants.GoogleCertPem,
		Signer: aliceAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgProposeAddX509GoogleRootcert}, aliceName, aliceAccount)
	require.NoError(suite.T, err)

	// Request all proposed Root certificate
	proposedCertificates, _ = GetAllProposedX509RootCerts(suite)
	require.Equal(suite.T, 1, len(proposedCertificates))
	require.Equal(suite.T, testconstants.GoogleSubject, proposedCertificates[0].Subject)
	require.Equal(suite.T, testconstants.GoogleSubjectKeyID, proposedCertificates[0].SubjectKeyId)

	// Request all approved certificates
	certificates, _ = GetAllX509Certs(suite)
	require.Equal(suite.T, 0, len(certificates))

	// Request proposed Google Root certificate
	proposedCertificate, _ = GetProposedX509RootCert(suite, testconstants.GoogleSubject, testconstants.GoogleSubjectKeyID)
	require.Equal(suite.T, testconstants.GoogleCertPem, proposedCertificate.PemCert)
	require.Equal(suite.T, aliceAccount.Address, proposedCertificate.Owner)
	require.Equal(suite.T, 1, len(proposedCertificate.Approvals))

	// Jack (Trustee) rejects Root certificate after approval
	msgRejectAddX509RootCert := pkitypes.MsgRejectAddX509RootCert{
		Subject:      proposedCertificate.Subject,
		SubjectKeyId: proposedCertificate.SubjectKeyId,
		Signer:       jackAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRejectAddX509RootCert}, jackName, jackAccount)
	require.NoError(suite.T, err)

	// Request all proposed Root certificates
	proposedCertificates, _ = GetAllProposedX509RootCerts(suite)
	require.Equal(suite.T, 1, len(proposedCertificates))
	require.Equal(suite.T, testconstants.GoogleSubject, proposedCertificates[0].Subject)
	require.Equal(suite.T, testconstants.GoogleSubjectKeyID, proposedCertificates[0].SubjectKeyId)

	// Request all approved certificates
	certificates, _ = GetAllX509Certs(suite)
	require.Equal(suite.T, 0, len(certificates))

	// Request proposed Root certificate
	proposedCertificate, _ = GetProposedX509RootCert(suite, testconstants.GoogleSubject, testconstants.GoogleSubjectKeyID)
	require.Equal(suite.T, testconstants.GoogleCertPem, proposedCertificate.PemCert)
	require.Equal(suite.T, aliceAccount.Address, proposedCertificate.Owner)
	require.True(suite.T, proposedCertificate.HasApprovalFrom(aliceAccount.Address))

	// Jack (Trustee) approves Root certificate
	msgApproveAddX509RootCert := pkitypes.MsgApproveAddX509RootCert{
		Subject:      proposedCertificate.Subject,
		SubjectKeyId: proposedCertificate.SubjectKeyId,
		Signer:       jackAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgApproveAddX509RootCert}, jackName, jackAccount)
	require.NoError(suite.T, err)

	// Request all proposed Root certificates
	proposedCertificates, _ = GetAllProposedX509RootCerts(suite)
	require.Equal(suite.T, 0, len(proposedCertificates))

	// Request all approved certificates
	certificates, _ = GetAllX509Certs(suite)
	require.Equal(suite.T, 1, len(certificates))
	require.Equal(suite.T, testconstants.GoogleSubject, certificates[0].Subject)
	require.Equal(suite.T, testconstants.GoogleSubjectKeyID, certificates[0].SubjectKeyId)

	// Request all approved Root certificates
	rootCertificates, _ = GetAllRootX509Certs(suite)
	require.Equal(suite.T, 1, len(rootCertificates.Certs))
	require.Equal(suite.T, testconstants.GoogleSubject, rootCertificates.Certs[0].Subject)
	require.Equal(suite.T, testconstants.GoogleSubjectKeyID, rootCertificates.Certs[0].SubjectKeyId)

	// Request approved Google Root certificate
	certificate, _ = GetX509Cert(suite, testconstants.GoogleSubject, testconstants.GoogleSubjectKeyID)
	require.Equal(suite.T, testconstants.GoogleSubject, certificate.Subject)
	require.Equal(suite.T, testconstants.GoogleSubjectKeyID, certificate.SubjectKeyId)
	require.Equal(suite.T, 1, len(certificate.Certs))
	require.Equal(suite.T, testconstants.GoogleCertPem, certificate.Certs[0].PemCert)
	require.Equal(suite.T, aliceAccount.Address, certificate.Certs[0].Owner)
	require.Equal(suite.T, testconstants.GoogleSubjectAsText, certificate.Certs[0].SubjectAsText)
	require.True(suite.T, certificate.Certs[0].IsRoot)

	// Jack (Trustee) proposes to revoke Google Root certificate
	msgProposeRevokeX509RootCert = pkitypes.MsgProposeRevokeX509RootCert{
		Subject:      testconstants.GoogleSubject,
		SubjectKeyId: testconstants.GoogleSubjectKeyID,
		Signer:       jackAccount.Address,
	}

	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgProposeRevokeX509RootCert}, jackName, jackAccount)
	require.NoError(suite.T, err)

	// Request all Root certificates proposed to revoke
	proposedRevocationCertificates, _ = GetAllProposedRevocationX509Certs(suite)
	require.Equal(suite.T, 1, len(proposedRevocationCertificates))
	require.Equal(suite.T, testconstants.GoogleSubject, proposedRevocationCertificates[0].Subject)
	require.Equal(suite.T, testconstants.GoogleSubjectKeyID, proposedRevocationCertificates[0].SubjectKeyId)

	// Request all revoked certificates
	revokedCertificates, _ = GetAllRevokedX509Certs(suite)
	require.Equal(suite.T, 3, len(revokedCertificates))
	require.Equal(suite.T, testconstants.LeafSubject, revokedCertificates[0].Subject)
	require.Equal(suite.T, testconstants.LeafSubjectKeyID, revokedCertificates[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.LeafSubjectAsText, revokedCertificates[0].Certs[0].SubjectAsText)
	require.Equal(suite.T, testconstants.RootSubject, revokedCertificates[1].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, revokedCertificates[1].SubjectKeyId)
	require.Equal(suite.T, testconstants.RootSubjectAsText, revokedCertificates[1].Certs[0].SubjectAsText)
	require.Equal(suite.T, testconstants.IntermediateSubject, revokedCertificates[2].Subject)
	require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, revokedCertificates[2].SubjectKeyId)
	require.Equal(suite.T, testconstants.IntermediateSubjectAsText, revokedCertificates[2].Certs[0].SubjectAsText)

	// Request all revoked Root certificates
	revokedRootCertificates, _ = GetAllRevokedRootX509Certs(suite)
	require.Equal(suite.T, 1, len(revokedRootCertificates.Certs))
	require.Equal(suite.T, testconstants.RootSubject, revokedRootCertificates.Certs[0].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, revokedRootCertificates.Certs[0].SubjectKeyId)

	// Request Google Root certificate proposed to revoke
	proposedCertificateRevocation, _ = GetProposedRevocationX509Cert(suite, testconstants.GoogleSubject, testconstants.GoogleSubjectKeyID)
	require.Equal(suite.T, testconstants.GoogleSubject, proposedCertificateRevocation.Subject)
	require.Equal(suite.T, testconstants.GoogleSubjectKeyID, proposedCertificateRevocation.SubjectKeyId)
	require.True(suite.T, proposedCertificateRevocation.HasRevocationFrom(jackAccount.Address))

	// Request all approved certificates
	certificates, _ = GetAllX509Certs(suite)
	require.Equal(suite.T, 1, len(certificates))
	require.Equal(suite.T, testconstants.GoogleSubject, certificates[0].Subject)
	require.Equal(suite.T, testconstants.GoogleSubjectKeyID, certificates[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.GoogleSubjectAsText, certificates[0].Certs[0].SubjectAsText)

	// Alice (Trustee) approves to revoke Google Root certificate
	msgApproveRevokeX509RootCert = pkitypes.MsgApproveRevokeX509RootCert{
		Subject:      proposedCertificate.Subject,
		SubjectKeyId: proposedCertificate.SubjectKeyId,
		Signer:       aliceAccount.Address,
	}

	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgApproveRevokeX509RootCert}, aliceName, aliceAccount)
	require.NoError(suite.T, err)

	// Request all Root certificates proposed to revoke
	proposedRevocationCertificates, _ = GetAllProposedRevocationX509Certs(suite)
	require.Equal(suite.T, 0, len(proposedRevocationCertificates))

	// Request all revoked certificates
	revokedCertificates, _ = GetAllRevokedX509Certs(suite)
	require.Equal(suite.T, 4, len(revokedCertificates))
	require.Equal(suite.T, testconstants.LeafSubject, revokedCertificates[0].Subject)
	require.Equal(suite.T, testconstants.LeafSubjectKeyID, revokedCertificates[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.LeafSubjectAsText, revokedCertificates[0].Certs[0].SubjectAsText)
	require.Equal(suite.T, testconstants.RootSubject, revokedCertificates[1].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, revokedCertificates[1].SubjectKeyId)
	require.Equal(suite.T, testconstants.RootSubjectAsText, revokedCertificates[1].Certs[0].SubjectAsText)
	require.Equal(suite.T, testconstants.IntermediateSubject, revokedCertificates[2].Subject)
	require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, revokedCertificates[2].SubjectKeyId)
	require.Equal(suite.T, testconstants.IntermediateSubjectAsText, revokedCertificates[2].Certs[0].SubjectAsText)
	require.Equal(suite.T, testconstants.GoogleSubject, revokedCertificates[3].Subject)
	require.Equal(suite.T, testconstants.GoogleSubjectKeyID, revokedCertificates[3].SubjectKeyId)
	require.Equal(suite.T, testconstants.GoogleSubjectAsText, revokedCertificates[3].Certs[0].SubjectAsText)

	// Request all revoked Root certificates
	revokedRootCertificates, _ = GetAllRevokedRootX509Certs(suite)
	require.Equal(suite.T, 2, len(revokedRootCertificates.Certs))
	require.Equal(suite.T, testconstants.RootSubject, revokedRootCertificates.Certs[0].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, revokedRootCertificates.Certs[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.GoogleSubject, revokedRootCertificates.Certs[1].Subject)
	require.Equal(suite.T, testconstants.GoogleSubjectKeyID, revokedRootCertificates.Certs[1].SubjectKeyId)

	// Request revoked Google Root certificate
	revokedCertificate, _ = GetRevokedX509Cert(suite, testconstants.GoogleSubject, testconstants.GoogleSubjectKeyID)
	require.Equal(suite.T, 1, len(revokedCertificate.Certs))
	require.Equal(suite.T, testconstants.GoogleSubject, revokedCertificate.Subject)
	require.Equal(suite.T, testconstants.GoogleSubjectKeyID, revokedCertificate.SubjectKeyId)
	require.Equal(suite.T, testconstants.GoogleSubject, revokedCertificate.Certs[0].Subject)
	require.Equal(suite.T, testconstants.GoogleSubjectKeyID, revokedCertificate.Certs[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.GoogleCertPem, revokedCertificate.Certs[0].PemCert)
	require.Equal(suite.T, aliceAccount.Address, revokedCertificate.Certs[0].Owner)
	require.Equal(suite.T, testconstants.GoogleSubjectAsText, revokedCertificate.Certs[0].SubjectAsText)
	require.True(suite.T, revokedCertificate.Certs[0].IsRoot)

	certificates, _ = GetAllX509Certs(suite)
	require.Equal(suite.T, 0, len(certificates))

	_, err = GetProposedX509RootCert(suite, testconstants.GoogleSubject, testconstants.GoogleSubjectKeyID)
	suite.AssertNotFound(err)

	_, err = GetX509Cert(suite, testconstants.GoogleSubject, testconstants.GoogleSubjectKeyID)
	suite.AssertNotFound(err)

	proposedRevocationCertificates, _ = GetAllProposedRevocationX509Certs(suite)
	require.Equal(suite.T, 0, len(proposedRevocationCertificates))

	_, err = GetProposedRevocationX509Cert(suite, testconstants.GoogleSubject, testconstants.GoogleSubjectKeyID)
	suite.AssertNotFound(err)

	rootCertificates, _ = GetAllRootX509Certs(suite)
	require.Equal(suite.T, 0, len(rootCertificates.Certs))

	_, err = GetAllChildX509Certs(suite, testconstants.GoogleSubject, testconstants.GoogleSubjectKeyID)
	suite.AssertNotFound(err)

	_, err = GetAllX509CertsBySubject(suite, testconstants.GoogleSubject)
	suite.AssertNotFound(err)

	// CHECK TEST ROOT CERTIFICATE FOR REJECTING
	_, err = GetRejectedX509RootCert(suite, testconstants.TestSubject, testconstants.TestSubjectKeyID)
	suite.AssertNotFound(err)

	_, err = GetProposedX509RootCert(suite, testconstants.TestSubject, testconstants.TestSubjectKeyID)
	suite.AssertNotFound(err)

	// Alice (Trustee) propose Test Root certificate
	msgProposeAddX509TestRootCert := pkitypes.MsgProposeAddX509RootCert{
		Cert:   testconstants.TestCertPem,
		Signer: aliceAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgProposeAddX509TestRootCert}, aliceName, aliceAccount)
	require.NoError(suite.T, err)

	// Request proposed Test Root certificate
	proposedCertificate, _ = GetProposedX509RootCert(suite, testconstants.TestSubject, testconstants.TestSubjectKeyID)
	require.Equal(suite.T, testconstants.TestCertPem, proposedCertificate.PemCert)
	require.Equal(suite.T, aliceAccount.Address, proposedCertificate.Owner)

	// Alice (Trustee) rejects Test Root certificate
	msgRejectX509TestRootCert := pkitypes.MsgRejectAddX509RootCert{
		Subject:      proposedCertificate.Subject,
		SubjectKeyId: proposedCertificate.SubjectKeyId,
		Signer:       aliceAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRejectX509TestRootCert}, aliceName, aliceAccount)
	require.NoError(suite.T, err)

	_, err = GetProposedX509RootCert(suite, testconstants.TestSubject, testconstants.TestSubjectKeyID)
	suite.AssertNotFound(err)

	_, err = GetX509Cert(suite, testconstants.TestSubject, testconstants.TestSubjectKeyID)
	suite.AssertNotFound(err)

	_, err = GetRevokedX509Cert(suite, testconstants.TestSubject, testconstants.TestSubjectKeyID)
	suite.AssertNotFound(err)

	// Alice (Trustee) propose Test Root certificate
	msgProposeAddX509TestRootCert = pkitypes.MsgProposeAddX509RootCert{
		Cert:   testconstants.TestCertPem,
		Signer: aliceAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgProposeAddX509TestRootCert}, aliceName, aliceAccount)
	require.NoError(suite.T, err)

	// Request all proposed Root certificates
	proposedCertificates, _ = GetAllProposedX509RootCerts(suite)
	require.Equal(suite.T, 1, len(proposedCertificates))
	require.Equal(suite.T, testconstants.TestSubject, proposedCertificates[0].Subject)
	require.Equal(suite.T, testconstants.TestSubjectKeyID, proposedCertificates[0].SubjectKeyId)

	// Request all rejected Root certificates
	rejectedCertificates, _ := GetAllRejectedX509RootCerts(suite)
	require.Equal(suite.T, 0, len(rejectedCertificates))

	// Request all approved Root certificates
	certificates, _ = GetAllX509Certs(suite)
	require.Equal(suite.T, 0, len(certificates))

	// Request proposed Test Root certificate
	proposedCertificate, _ = GetProposedX509RootCert(suite, testconstants.TestSubject, testconstants.TestSubjectKeyID)
	require.Equal(suite.T, testconstants.TestCertPem, proposedCertificate.PemCert)
	require.Equal(suite.T, aliceAccount.Address, proposedCertificate.Owner)
	require.Equal(suite.T, 1, len(proposedCertificate.Approvals))

	// Jack (Trustee) rejects Root certificate
	msgRejectAddX509RootCert = pkitypes.MsgRejectAddX509RootCert{
		Subject:      proposedCertificate.Subject,
		SubjectKeyId: proposedCertificate.SubjectKeyId,
		Signer:       jackAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRejectAddX509RootCert}, jackName, jackAccount)
	require.NoError(suite.T, err)

	// Jack (Trustee) rejects Root certificate for the second time
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRejectAddX509RootCert}, jackName, jackAccount)
	require.Error(suite.T, err)

	// Request all proposed Root certificates
	proposedCertificates, _ = GetAllProposedX509RootCerts(suite)
	require.Equal(suite.T, 1, len(proposedCertificates))
	require.Equal(suite.T, testconstants.TestSubject, proposedCertificates[0].Subject)
	require.Equal(suite.T, testconstants.TestSubjectKeyID, proposedCertificates[0].SubjectKeyId)

	// Request all rejected Root certificates
	rejectedCertificates, _ = GetAllRejectedX509RootCerts(suite)
	require.Equal(suite.T, 0, len(rejectedCertificates))

	// Request all approved Root certificates
	certificates, _ = GetAllX509Certs(suite)
	require.Equal(suite.T, 0, len(certificates))

	// Request proposed Test Root certificate
	proposedCertificate, _ = GetProposedX509RootCert(suite, testconstants.TestSubject, testconstants.TestSubjectKeyID)
	require.Equal(suite.T, testconstants.TestCertPem, proposedCertificate.PemCert)
	require.Equal(suite.T, aliceAccount.Address, proposedCertificate.Owner)
	require.Equal(suite.T, 1, len(proposedCertificate.Approvals))
	require.True(suite.T, proposedCertificate.HasRejectFrom(jackAccount.Address))

	// Alice (Trustee) rejects Root certificate
	secondMsgRejectAddX509RootCert := pkitypes.MsgRejectAddX509RootCert{
		Subject:      proposedCertificate.Subject,
		SubjectKeyId: proposedCertificate.SubjectKeyId,
		Signer:       aliceAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&secondMsgRejectAddX509RootCert}, aliceName, aliceAccount)
	require.NoError(suite.T, err)

	// Request all proposed Root certificates
	proposedCertificates, _ = GetAllProposedX509RootCerts(suite)
	require.Equal(suite.T, 0, len(proposedCertificates))

	// Request all approved Root certificates
	certificates, _ = GetAllX509Certs(suite)
	require.Equal(suite.T, 0, len(certificates))

	// Request all rejected Root certificates
	rejectedCertificates, _ = GetAllRejectedX509RootCerts(suite)
	require.Equal(suite.T, 1, len(rejectedCertificates))
	require.Equal(suite.T, testconstants.TestSubject, rejectedCertificates[0].Subject)
	require.Equal(suite.T, testconstants.TestSubjectKeyID, rejectedCertificates[0].SubjectKeyId)

	// Request rejected Test Root certificate
	rejectedCertificate, _ := GetRejectedX509RootCert(suite, testconstants.TestSubject, testconstants.TestSubjectKeyID)
	require.Equal(suite.T, testconstants.TestSubject, rejectedCertificate.Subject)
	require.Equal(suite.T, testconstants.TestSubjectKeyID, rejectedCertificate.SubjectKeyId)
	require.Equal(suite.T, 1, len(rejectedCertificate.Certs))
	require.Equal(suite.T, testconstants.TestCertPem, rejectedCertificate.Certs[0].PemCert)
	require.Equal(suite.T, testconstants.TestSubjectAsText, rejectedCertificate.Certs[0].SubjectAsText)
	require.Equal(suite.T, aliceAccount.Address, rejectedCertificate.Certs[0].Owner)
	require.Equal(suite.T, jackAccount.Address, rejectedCertificate.Certs[0].Rejects[0].Address)
	require.Equal(suite.T, aliceAccount.Address, rejectedCertificate.Certs[0].Rejects[1].Address)
}
