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

	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
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

func GetProposedX509RootCert(suite *utils.TestSuite, subject string, subjectKeyId string) (res *pkitypes.ProposedCertificate, err error) {
	if suite.Rest {
		var resp pkitypes.QueryGetProposedCertificateResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/pki/proposed-certificates/%s/%s",
				url.QueryEscape(subject), url.QueryEscape(subjectKeyId),
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
				SubjectKeyId: subjectKeyId,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetProposedCertificate()
	}

	return res, nil
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

func GetX509Cert(suite *utils.TestSuite, subject string, subjectKeyId string) (res *pkitypes.ApprovedCertificates, err error) {
	if suite.Rest {
		var resp pkitypes.QueryGetApprovedCertificatesResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/pki/certificates/%s/%s",
				url.QueryEscape(subject), url.QueryEscape(subjectKeyId),
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
				SubjectKeyId: subjectKeyId,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetApprovedCertificates()
	}

	return res, nil
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

func GetRevokedX509Cert(suite *utils.TestSuite, subject string, subjectKeyId string) (res *pkitypes.RevokedCertificates, err error) {
	if suite.Rest {
		var resp pkitypes.QueryGetRevokedCertificatesResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/pki/revoked-certificates/%s/%s",
				url.QueryEscape(subject), url.QueryEscape(subjectKeyId),
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
				SubjectKeyId: subjectKeyId,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetRevokedCertificates()
	}

	return res, nil
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

func GetProposedRevocationX509Cert(suite *utils.TestSuite, subject string, subjectKeyId string) (res *pkitypes.ProposedCertificateRevocation, err error) {
	if suite.Rest {
		var resp pkitypes.QueryGetProposedCertificateRevocationResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/pki/proposed-revocation-certificates/%s/%s",
				url.QueryEscape(subject), url.QueryEscape(subjectKeyId),
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
				SubjectKeyId: subjectKeyId,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetProposedCertificateRevocation()
	}

	return res, nil
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

func GetAllX509CertsBySubject(suite *utils.TestSuite, subject string) (res *pkitypes.ApprovedCertificatesBySubject, err error) {
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

	return res, nil
}

func GetAllChildX509Certs(suite *utils.TestSuite, subject string, subjectKeyId string) (res *pkitypes.ChildCertificates, err error) {
	if suite.Rest {
		var resp pkitypes.QueryGetChildCertificatesResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/pki/child-certificates/%s/%s",
				url.QueryEscape(subject), url.QueryEscape(subjectKeyId),
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
				AuthorityKeyId: subjectKeyId,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetChildCertificates()
	}

	return res, nil
}

// Common Test Logic

//nolint:funlen
func PKIDemo(suite *utils.TestSuite) {
	// rand.Seed(time.Now().UnixNano())

	// jackName := testconstants.JackAccount
	// aliceName := testconstants.AliceAccount

	// jackKeyInfo, _ := suite.Kr.Key(jackName)
	// aliceKeyInfo, _ := suite.Kr.Key(aliceName)

	// // build map with an acc address as a key
	// inputAccounts, err := dclauthhelpers.GetAccounts(suite)
	// require.NoError(suite.T, err)
	// accDataInitial := make(map[string]dclauthtypes.Account)
	// for _, acc := range inputAccounts {
	// 	accDataInitial[acc.GetAddress().String()] = acc
	// }

	// // Create account for Anna
	// vendorAccountName := "newVendorAccount" + strconv.Itoa(rand.Intn(1000))
	// _ = dclauthhelpers.CreateAccount(
	// 	suite,
	// 	vendorAccountName, dclauthtypes.AccountRoles{dclauthtypes.Vendor}, uint16(testconstants.VID),
	// 	jackName, accDataInitial[jackKeyInfo.GetAddress().String()],
	// 	aliceName, accDataInitial[aliceKeyInfo.GetAddress().String()],
	// )

	// All requests return empty or 404 value
	proposedCertificates, _ := GetAllProposedX509RootCerts(suite)
	require.Equal(suite.T, 0, len(proposedCertificates))

	_, err := GetProposedX509RootCert(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Error(suite.T, err)
	require.Contains(suite.T, err.Error(), "rpc error: code = NotFound desc = not found")

	allCertificates, _ := GetAllX509Certs(suite)
	require.Equal(suite.T, 0, len(allCertificates))

	allRevokedCertificates, _ := GetAllRevokedX509Certs(suite)
	require.Equal(suite.T, 0, len(allRevokedCertificates))

	allProposedRevocationCertificates, _ := GetAllProposedRevocationX509Certs(suite)
	require.Equal(suite.T, 0, len(allProposedRevocationCertificates))

	allRootCertificates, _ := GetAllRootX509Certs(suite)
	require.Equal(suite.T, 0, len(allRootCertificates.Certs))

	allRevokedRootCertificates, _ := GetAllRevokedRootX509Certs(suite)
	require.Equal(suite.T, 0, len(allRevokedRootCertificates.Certs))

	_, err = GetAllChildX509Certs(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Error(suite.T, err)
	require.Contains(suite.T, err.Error(), "rpc error: code = NotFound desc = not found")

	_, err = GetAllX509CertsBySubject(suite, testconstants.RootSubject)
	require.Error(suite.T, err)
	require.Contains(suite.T, err.Error(), "rpc error: code = NotFound desc = not found")

	// // Request all approved certificates
	// certificates, _ := utils.GetAllX509Certs()
	// require.Equal(suite.T, 0, len(certificates.Items))

	// // User (Not Trustee) propose Root certificate
	// msgProposeAddX509RootCert := pki.MsgProposeAddX509RootCert{
	// 	Cert:   testconstants.RootCertPem,
	// 	Signer: userKeyInfo.Address,
	// }
	// _, _ = utils.ProposeAddX509RootCert(msgProposeAddX509RootCert, userKeyInfo.Name, testconstants.Passphrase)

	// // Request all proposed Root certificates
	// proposedCertificates, _ = utils.GetAllProposedX509RootCerts()
	// require.Equal(suite.T, 1, len(proposedCertificates.Items))
	// require.Equal(suite.T, testconstants.RootSubject, proposedCertificates.Items[0].Subject)
	// require.Equal(suite.T, testconstants.RootSubjectKeyID, proposedCertificates.Items[0].SubjectKeyID)

	// // Request all approved certificates
	// certificates, _ = utils.GetAllX509Certs()
	// require.Equal(suite.T, 0, len(certificates.Items))

	// // Request proposed Root certificate
	// proposedCertificate, _ :=
	// 	utils.GetProposedX509RootCert(testconstants.RootSubject, testconstants.RootSubjectKeyID)
	// require.Equal(suite.T, testconstants.RootCertPem, proposedCertificate.PemCert)
	// require.Equal(suite.T, userKeyInfo.Address, proposedCertificate.Owner)
	// require.Nil(suite.T, proposedCertificate.Approvals)

	// // Jack (Trustee) approve Root certificate
	// msgApproveAddX509RootCert := pki.MsgApproveAddX509RootCert{
	// 	Subject:      proposedCertificate.Subject,
	// 	SubjectKeyID: proposedCertificate.SubjectKeyID,
	// 	Signer:       jackKeyInfo.Address,
	// }
	// _, _ = utils.ApproveAddX509RootCert(msgApproveAddX509RootCert, jackKeyInfo.Name, testconstants.Passphrase)

	// // Request all proposed Root certificates
	// proposedCertificates, _ = utils.GetAllProposedX509RootCerts()
	// require.Equal(suite.T, 1, len(proposedCertificates.Items))
	// require.Equal(suite.T, testconstants.RootSubject, proposedCertificates.Items[0].Subject)
	// require.Equal(suite.T, testconstants.RootSubjectKeyID, proposedCertificates.Items[0].SubjectKeyID)

	// // Request all approved certificates
	// certificates, _ = utils.GetAllX509Certs()
	// require.Equal(suite.T, 0, len(certificates.Items))

	// // Request proposed Root certificate
	// proposedCertificate, _ =
	// 	utils.GetProposedX509RootCert(testconstants.RootSubject, testconstants.RootSubjectKeyID)
	// require.Equal(suite.T, testconstants.RootCertPem, proposedCertificate.PemCert)
	// require.Equal(suite.T, userKeyInfo.Address, proposedCertificate.Owner)
	// require.Equal(suite.T, []sdk.AccAddress{jackKeyInfo.Address}, proposedCertificate.Approvals)

	// // Alice (Trustee) approve Root certificate
	// secondMsgApproveAddX509RootCert := pki.MsgApproveAddX509RootCert{
	// 	Subject:      proposedCertificate.Subject,
	// 	SubjectKeyID: proposedCertificate.SubjectKeyID,
	// 	Signer:       aliceKeyInfo.Address,
	// }
	// _, _ = utils.ApproveAddX509RootCert(secondMsgApproveAddX509RootCert, aliceKeyInfo.Name, testconstants.Passphrase)

	// // Request all proposed Root certificates
	// proposedCertificates, _ = utils.GetAllProposedX509RootCerts()
	// require.Equal(suite.T, 0, len(proposedCertificates.Items))

	// // Request all approved certificates
	// certificates, _ = utils.GetAllX509Certs()
	// require.Equal(suite.T, 1, len(certificates.Items))
	// require.Equal(suite.T, testconstants.RootSubject, certificates.Items[0].Subject)
	// require.Equal(suite.T, testconstants.RootSubjectKeyID, certificates.Items[0].SubjectKeyID)

	// // Request all approved Root certificates
	// certificates, _ = utils.GetAllX509RootCerts()
	// require.Equal(suite.T, 1, len(certificates.Items))
	// require.Equal(suite.T, testconstants.RootSubject, certificates.Items[0].Subject)
	// require.Equal(suite.T, testconstants.RootSubjectKeyID, certificates.Items[0].SubjectKeyID)

	// // Request approved Root certificate
	// certificate, _ := utils.GetX509Cert(testconstants.RootSubject, testconstants.RootSubjectKeyID)
	// require.Equal(suite.T, testconstants.RootCertPem, certificate.PemCert)
	// require.Equal(suite.T, userKeyInfo.Address, certificate.Owner)
	// require.True(suite.T, certificate.IsRoot)

	// // Request certificate chain for Root certificate
	// certificateChain, _ := utils.GetX509CertChain(testconstants.RootSubject, testconstants.RootSubjectKeyID)
	// require.Len(suite.T, certificateChain.Items, 1)
	// require.Equal(suite.T, testconstants.RootCertPem, certificateChain.Items[0].PemCert)
	// require.Equal(suite.T, userKeyInfo.Address, certificateChain.Items[0].Owner)
	// require.True(suite.T, certificateChain.Items[0].IsRoot)

	// // User (Not Trustee) add Intermediate certificate
	// msgAddX509Cert := pki.MsgAddX509Cert{
	// 	Cert:   testconstants.IntermediateCertPem,
	// 	Signer: userKeyInfo.Address,
	// }
	// _, _ = utils.AddX509Cert(msgAddX509Cert, userKeyInfo.Name, testconstants.Passphrase)

	// // Request all proposed Root certificates
	// proposedCertificates, _ = utils.GetAllProposedX509RootCerts()
	// require.Equal(suite.T, 0, len(proposedCertificates.Items))

	// // Request all approved certificates
	// certificates, _ = utils.GetAllX509Certs()
	// require.Equal(suite.T, 2, len(certificates.Items))
	// require.Equal(suite.T, testconstants.IntermediateSubject, certificates.Items[0].Subject)
	// require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, certificates.Items[0].SubjectKeyID)
	// require.Equal(suite.T, testconstants.RootSubject, certificates.Items[1].Subject)
	// require.Equal(suite.T, testconstants.RootSubjectKeyID, certificates.Items[1].SubjectKeyID)

	// // Request all approved Root certificates
	// certificates, _ = utils.GetAllX509RootCerts()
	// require.Equal(suite.T, 1, len(certificates.Items))
	// require.Equal(suite.T, testconstants.RootSubject, certificates.Items[0].Subject)
	// require.Equal(suite.T, testconstants.RootSubjectKeyID, certificates.Items[0].SubjectKeyID)

	// // Request Intermediate certificate
	// certificate, _ = utils.GetX509Cert(testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	// require.Equal(suite.T, testconstants.IntermediateCertPem, certificate.PemCert)
	// require.Equal(suite.T, userKeyInfo.Address, certificate.Owner)
	// require.False(t, certificate.IsRoot)

	// // Request certificate chain for Intermediate certificate
	// certificateChain, _ = utils.GetX509CertChain(testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	// require.Len(suite.T, certificateChain.Items, 2)
	// require.Equal(suite.T, testconstants.IntermediateCertPem, certificateChain.Items[0].PemCert)
	// require.Equal(suite.T, userKeyInfo.Address, certificateChain.Items[0].Owner)
	// require.False(t, certificateChain.Items[0].IsRoot)
	// require.Equal(suite.T, testconstants.RootCertPem, certificateChain.Items[1].PemCert)
	// require.Equal(suite.T, userKeyInfo.Address, certificateChain.Items[1].Owner)
	// require.True(suite.T, certificateChain.Items[1].IsRoot)

	// // Alice (Trustee) add Leaf certificate
	// secondMsgAddX509Cert := pki.MsgAddX509Cert{
	// 	Cert:   testconstants.LeafCertPem,
	// 	Signer: aliceKeyInfo.Address,
	// }
	// _, _ = utils.AddX509Cert(secondMsgAddX509Cert, aliceKeyInfo.Name, testconstants.Passphrase)

	// // Request all proposed Root certificates
	// proposedCertificates, _ = utils.GetAllProposedX509RootCerts()
	// require.Equal(suite.T, 0, len(proposedCertificates.Items))

	// // Request all approved certificates
	// certificates, _ = utils.GetAllX509Certs()
	// require.Equal(suite.T, 3, len(certificates.Items))
	// require.Equal(suite.T, testconstants.IntermediateSubject, certificates.Items[0].Subject)
	// require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, certificates.Items[0].SubjectKeyID)
	// require.Equal(suite.T, testconstants.LeafSubject, certificates.Items[1].Subject)
	// require.Equal(suite.T, testconstants.LeafSubjectKeyID, certificates.Items[1].SubjectKeyID)
	// require.Equal(suite.T, testconstants.RootSubject, certificates.Items[2].Subject)
	// require.Equal(suite.T, testconstants.RootSubjectKeyID, certificates.Items[2].SubjectKeyID)

	// // Request all approved Root certificates
	// certificates, _ = utils.GetAllX509RootCerts()
	// require.Equal(suite.T, 1, len(certificates.Items))
	// require.Equal(suite.T, testconstants.RootSubject, certificates.Items[0].Subject)
	// require.Equal(suite.T, testconstants.RootSubjectKeyID, certificates.Items[0].SubjectKeyID)

	// // Request Leaf certificate
	// certificate, _ = utils.GetX509Cert(testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	// require.Equal(suite.T, testconstants.LeafCertPem, certificate.PemCert)
	// require.Equal(suite.T, aliceKeyInfo.Address, certificate.Owner)
	// require.False(t, certificate.IsRoot)

	// // Request certificate chain for Leaf certificate
	// certificateChain, _ = utils.GetX509CertChain(testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	// require.Len(suite.T, certificateChain.Items, 3)
	// require.Equal(suite.T, testconstants.LeafCertPem, certificateChain.Items[0].PemCert)
	// require.Equal(suite.T, aliceKeyInfo.Address, certificateChain.Items[0].Owner)
	// require.False(t, certificateChain.Items[0].IsRoot)
	// require.Equal(suite.T, testconstants.IntermediateCertPem, certificateChain.Items[1].PemCert)
	// require.Equal(suite.T, userKeyInfo.Address, certificateChain.Items[1].Owner)
	// require.False(t, certificateChain.Items[1].IsRoot)
	// require.Equal(suite.T, testconstants.RootCertPem, certificateChain.Items[2].PemCert)
	// require.Equal(suite.T, userKeyInfo.Address, certificateChain.Items[2].Owner)
	// require.True(suite.T, certificateChain.Items[2].IsRoot)

	// // Request all Subject certificates
	// certificates, _ = utils.GetAllSubjectX509Certs(testconstants.LeafSubject)
	// require.Equal(suite.T, 1, len(certificates.Items))
	// require.Equal(suite.T, testconstants.LeafSubjectKeyID, certificates.Items[0].SubjectKeyID)

	// // Request all Root certificates proposed to revoke
	// proposedCertificateRevocations, _ := utils.GetAllProposedX509RootCertsToRevoke()
	// require.Equal(suite.T, 0, len(proposedCertificateRevocations.Items))

	// // Request all revoked certificates
	// revokedCertificates, _ := utils.GetAllRevokedX509Certs()
	// require.Equal(suite.T, 0, len(revokedCertificates.Items))

	// // User (Not Trustee) revokes Intermediate certificate. This must also revoke its child - Leaf certificate.
	// msgRevokeX509Cert := pki.MsgRevokeX509Cert{
	// 	Subject:      testconstants.IntermediateSubject,
	// 	SubjectKeyID: testconstants.IntermediateSubjectKeyID,
	// 	Signer:       userKeyInfo.Address,
	// }
	// _, _ = utils.RevokeX509Cert(msgRevokeX509Cert, userKeyInfo.Name, testconstants.Passphrase)

	// // Request all Root certificates proposed to revoke
	// proposedCertificateRevocations, _ = utils.GetAllProposedX509RootCertsToRevoke()
	// require.Equal(suite.T, 0, len(proposedCertificateRevocations.Items))

	// // Request all revoked certificates
	// revokedCertificates, _ = utils.GetAllRevokedX509Certs()
	// require.Equal(suite.T, 2, len(revokedCertificates.Items))
	// require.Equal(suite.T, testconstants.IntermediateSubject, revokedCertificates.Items[0].Subject)
	// require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, revokedCertificates.Items[0].SubjectKeyID)
	// require.Equal(suite.T, testconstants.LeafSubject, revokedCertificates.Items[1].Subject)
	// require.Equal(suite.T, testconstants.LeafSubjectKeyID, revokedCertificates.Items[1].SubjectKeyID)

	// // Request all revoked Root certificates
	// revokedCertificates, _ = utils.GetAllRevokedX509RootCerts()
	// require.Equal(suite.T, 0, len(revokedCertificates.Items))

	// // Request revoked Intermediate certificate
	// revokedCertificate, _ := utils.GetRevokedX509Cert(
	// 	testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	// require.Equal(suite.T, testconstants.IntermediateCertPem, revokedCertificate.PemCert)
	// require.Equal(suite.T, userKeyInfo.Address, revokedCertificate.Owner)
	// require.False(t, revokedCertificate.IsRoot)

	// // Request revoked Leaf certificate
	// revokedCertificate, _ = utils.GetRevokedX509Cert(testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	// require.Equal(suite.T, testconstants.LeafCertPem, revokedCertificate.PemCert)
	// require.Equal(suite.T, aliceKeyInfo.Address, revokedCertificate.Owner)
	// require.False(t, revokedCertificate.IsRoot)

	// // Request all approved certificates
	// certificates, _ = utils.GetAllX509Certs()
	// require.Equal(suite.T, 1, len(certificates.Items))
	// require.Equal(suite.T, testconstants.RootSubject, certificates.Items[0].Subject)
	// require.Equal(suite.T, testconstants.RootSubjectKeyID, certificates.Items[0].SubjectKeyID)

	// // Jack (Trustee) proposes to revoke Root certificate
	// msgProposeRevokeX509RootCert := pki.MsgProposeRevokeX509RootCert{
	// 	Subject:      testconstants.RootSubject,
	// 	SubjectKeyID: testconstants.RootSubjectKeyID,
	// 	Signer:       jackKeyInfo.Address,
	// }
	// _, _ = utils.ProposeRevokeX509RootCert(msgProposeRevokeX509RootCert, jackKeyInfo.Name, testconstants.Passphrase)

	// // Request all Root certificates proposed to revoke
	// proposedCertificateRevocations, _ = utils.GetAllProposedX509RootCertsToRevoke()
	// require.Equal(suite.T, 1, len(proposedCertificateRevocations.Items))
	// require.Equal(suite.T, testconstants.RootSubject, proposedCertificateRevocations.Items[0].Subject)
	// require.Equal(suite.T, testconstants.RootSubjectKeyID, proposedCertificateRevocations.Items[0].SubjectKeyID)

	// // Request all revoked certificates
	// revokedCertificates, _ = utils.GetAllRevokedX509Certs()
	// require.Equal(suite.T, 2, len(revokedCertificates.Items))
	// require.Equal(suite.T, testconstants.IntermediateSubject, revokedCertificates.Items[0].Subject)
	// require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, revokedCertificates.Items[0].SubjectKeyID)
	// require.Equal(suite.T, testconstants.LeafSubject, revokedCertificates.Items[1].Subject)
	// require.Equal(suite.T, testconstants.LeafSubjectKeyID, revokedCertificates.Items[1].SubjectKeyID)

	// // Request all revoked Root certificates
	// revokedCertificates, _ = utils.GetAllRevokedX509RootCerts()
	// require.Equal(suite.T, 0, len(revokedCertificates.Items))

	// // Request Root certificate proposed to revoke
	// proposedCertificateRevocation, _ :=
	// 	utils.GetProposedX509RootCertToRevoke(testconstants.RootSubject, testconstants.RootSubjectKeyID)
	// require.Equal(suite.T, []sdk.AccAddress{jackKeyInfo.Address}, proposedCertificateRevocation.Approvals)

	// // Request all approved certificates
	// certificates, _ = utils.GetAllX509Certs()
	// require.Equal(suite.T, 1, len(certificates.Items))
	// require.Equal(suite.T, testconstants.RootSubject, certificates.Items[0].Subject)
	// require.Equal(suite.T, testconstants.RootSubjectKeyID, certificates.Items[0].SubjectKeyID)

	// // Alice (Trustee) approves to revoke Root certificate
	// msgApproveRevokeX509RootCert := pki.MsgApproveRevokeX509RootCert{
	// 	Subject:      proposedCertificate.Subject,
	// 	SubjectKeyID: proposedCertificate.SubjectKeyID,
	// 	Signer:       aliceKeyInfo.Address,
	// }
	// _, _ = utils.ApproveRevokeX509RootCert(msgApproveRevokeX509RootCert, aliceKeyInfo.Name, testconstants.Passphrase)

	// // Request all Root certificates proposed to revoke
	// proposedCertificateRevocations, _ = utils.GetAllProposedX509RootCertsToRevoke()
	// require.Equal(suite.T, 0, len(proposedCertificateRevocations.Items))

	// // Request all revoked certificates
	// revokedCertificates, _ = utils.GetAllRevokedX509Certs()
	// require.Equal(suite.T, 3, len(revokedCertificates.Items))
	// require.Equal(suite.T, testconstants.IntermediateSubject, revokedCertificates.Items[0].Subject)
	// require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, revokedCertificates.Items[0].SubjectKeyID)
	// require.Equal(suite.T, testconstants.LeafSubject, revokedCertificates.Items[1].Subject)
	// require.Equal(suite.T, testconstants.LeafSubjectKeyID, revokedCertificates.Items[1].SubjectKeyID)
	// require.Equal(suite.T, testconstants.RootSubject, revokedCertificates.Items[2].Subject)
	// require.Equal(suite.T, testconstants.RootSubjectKeyID, revokedCertificates.Items[2].SubjectKeyID)

	// // Request all revoked Root certificates
	// revokedCertificates, _ = utils.GetAllRevokedX509RootCerts()
	// require.Equal(suite.T, 1, len(revokedCertificates.Items))
	// require.Equal(suite.T, testconstants.RootSubject, revokedCertificates.Items[0].Subject)
	// require.Equal(suite.T, testconstants.RootSubjectKeyID, revokedCertificates.Items[0].SubjectKeyID)

	// // Request revoked Root certificate
	// revokedCertificate, _ = utils.GetRevokedX509Cert(testconstants.RootSubject, testconstants.RootSubjectKeyID)
	// require.Equal(suite.T, testconstants.RootCertPem, revokedCertificate.PemCert)
	// require.Equal(suite.T, userKeyInfo.Address, revokedCertificate.Owner)
	// require.True(suite.T, revokedCertificate.IsRoot)

	// // Request all approved certificates
	// certificates, _ = utils.GetAllX509Certs()
	// require.Equal(suite.T, 0, len(certificates.Items))
}
