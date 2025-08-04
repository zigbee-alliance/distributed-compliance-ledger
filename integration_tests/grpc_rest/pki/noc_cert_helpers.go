package pki

import (
	"context"
	"fmt"
	"net/url"

	tmrand "github.com/cometbft/cometbft/libs/rand"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	test_dclauth "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/grpc_rest/dclauth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	crttypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

//nolint:godox
/*
	To Run test you need:
		* Run LocalNet with: `make install && make localnet_init && make localnet_start`

	TODO: provide tests for error cases
*/

func GetNocX509Cert(suite *utils.TestSuite, subject string, subjectKeyID string) (*pkitypes.NocCertificates, error) {
	var res pkitypes.NocCertificates
	if suite.Rest {
		var resp pkitypes.QueryGetNocCertificatesResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/pki/all-noc-certificates/%s/%s",
				url.QueryEscape(subject), url.QueryEscape(subjectKeyID),
			),
			&resp,
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetNocCertificates()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.NocCertificates(
			context.Background(),
			&pkitypes.QueryGetNocCertificatesRequest{
				Subject:      subject,
				SubjectKeyId: subjectKeyID,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetNocCertificates()
	}

	return &res, nil
}

func GetAllNocX509CertsBySubject(suite *utils.TestSuite, subject string) (*pkitypes.NocCertificatesBySubject, error) {
	var res pkitypes.NocCertificatesBySubject
	if suite.Rest {
		var resp pkitypes.QueryGetNocCertificatesBySubjectResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/pki/all-noc-certificates/%s",
				url.QueryEscape(subject),
			),
			&resp,
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetNocCertificatesBySubject()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.NocCertificatesBySubject(
			context.Background(),
			&pkitypes.QueryGetNocCertificatesBySubjectRequest{
				Subject: subject,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetNocCertificatesBySubject()
	}

	return &res, nil
}

func GetAllNocX509Certs(suite *utils.TestSuite) (res []pkitypes.NocCertificates, err error) {
	return getAllNocX509Certs(suite, "")
}

func GetAllNocX509certsBySubjectKeyID(suite *utils.TestSuite, subjectKeyID string) (res []pkitypes.NocCertificates, err error) {
	return getAllNocX509Certs(suite, subjectKeyID)
}

func getAllNocX509Certs(suite *utils.TestSuite, subjectKeyID string) (res []pkitypes.NocCertificates, err error) {
	if suite.Rest {
		var resp pkitypes.QueryNocCertificatesResponse
		err := suite.QueryREST(fmt.Sprintf("/dcl/pki/all-noc-certificates?subjectKeyId=%s", subjectKeyID), &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetNocCertificates()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.NocCertificatesAll(
			context.Background(),
			&pkitypes.QueryNocCertificatesRequest{SubjectKeyId: subjectKeyID},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetNocCertificates()
	}

	return res, nil
}

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

func GetAllNocX509IcaCerts(suite *utils.TestSuite) (res []pkitypes.NocIcaCertificates, err error) {
	if suite.Rest {
		var resp pkitypes.QueryAllNocIcaCertificatesResponse
		err := suite.QueryREST("/dcl/pki/noc-ica-certificates/", &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetNocIcaCertificates()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.NocIcaCertificatesAll(
			context.Background(),
			&pkitypes.QueryAllNocIcaCertificatesRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetNocIcaCertificates()
	}

	return res, nil
}

func GetNocX509RootCerts(suite *utils.TestSuite, vendorID int32) (*pkitypes.NocRootCertificates, error) {
	var res pkitypes.NocRootCertificates
	if suite.Rest {
		var resp pkitypes.QueryGetNocRootCertificatesResponse
		err := suite.QueryREST(fmt.Sprintf("/dcl/pki/noc-vid-root-certificates/%v", vendorID), &resp)
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

func GetNocX509CertsByVidAndSkid(suite *utils.TestSuite, vendorID int32, subjectKeyID string) (*pkitypes.NocCertificatesByVidAndSkid, error) {
	var res pkitypes.NocCertificatesByVidAndSkid
	if suite.Rest {
		var resp pkitypes.QueryGetNocCertificatesByVidAndSkidResponse
		err := suite.QueryREST(fmt.Sprintf("/dcl/pki/noc-vid-certificates/%v/%s", vendorID, url.QueryEscape(subjectKeyID)), &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetNocCertificatesByVidAndSkid()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.NocCertificatesByVidAndSkid(
			context.Background(),
			&pkitypes.QueryGetNocCertificatesByVidAndSkidRequest{Vid: vendorID, SubjectKeyId: subjectKeyID},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetNocCertificatesByVidAndSkid()
	}

	return &res, nil
}

func GetNocX509IcaCerts(suite *utils.TestSuite, vendorID int32) (*pkitypes.NocIcaCertificates, error) {
	var res pkitypes.NocIcaCertificates
	if suite.Rest {
		var resp pkitypes.QueryGetNocIcaCertificatesResponse
		err := suite.QueryREST(fmt.Sprintf("/dcl/pki/noc-vid-ica-certificates/%v", vendorID), &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetNocIcaCertificates()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.NocIcaCertificates(
			context.Background(),
			&pkitypes.QueryGetNocIcaCertificatesRequest{Vid: vendorID},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetNocIcaCertificates()
	}

	return &res, nil
}

func GetNocX509IcaCertsBySubjectAndSKID(suite *utils.TestSuite, vendorID int32, subject, subjectKeyID string) []*pkitypes.Certificate {
	certs, _ := GetNocX509IcaCerts(suite, vendorID)
	for i := 0; i < len(certs.Certs); {
		cert := certs.Certs[i]
		if cert.Subject != subject || cert.SubjectKeyId != subjectKeyID {
			certs.Certs = append(certs.Certs[:i], certs.Certs[i+1:]...)
		} else {
			i++
		}
	}

	return certs.Certs
}

func GetRevokedNocX509RootCert(suite *utils.TestSuite, subject string, subjectKeyID string) (*pkitypes.RevokedNocRootCertificates, error) {
	var res pkitypes.RevokedNocRootCertificates
	if suite.Rest {
		var resp pkitypes.QueryGetRevokedNocRootCertificatesResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/pki/revoked-noc-root-certificates/%s/%s",
				url.QueryEscape(subject), url.QueryEscape(subjectKeyID),
			),
			&resp,
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetRevokedNocRootCertificates()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.RevokedNocRootCertificates(
			context.Background(),
			&pkitypes.QueryGetRevokedNocRootCertificatesRequest{
				Subject:      subject,
				SubjectKeyId: subjectKeyID,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetRevokedNocRootCertificates()
	}

	return &res, nil
}

func GetRevokedNocX509IcaCert(suite *utils.TestSuite, subject string, subjectKeyID string) (*pkitypes.RevokedNocIcaCertificates, error) {
	var res pkitypes.RevokedNocIcaCertificates
	if suite.Rest {
		var resp pkitypes.QueryGetRevokedNocIcaCertificatesResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/pki/revoked-noc-ica-certificates/%s/%s",
				url.QueryEscape(subject), url.QueryEscape(subjectKeyID),
			),
			&resp,
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetRevokedNocIcaCertificates()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.RevokedNocIcaCertificates(
			context.Background(),
			&pkitypes.QueryGetRevokedNocIcaCertificatesRequest{
				Subject:      subject,
				SubjectKeyId: subjectKeyID,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetRevokedNocIcaCertificates()
	}

	return &res, nil
}

//nolint:funlen
func NocCertDemo(suite *utils.TestSuite) {
	// Generate VIDs
	vid1 := int32(tmrand.Uint16())
	vid2 := int32(tmrand.Uint16())

	// All requests return empty or 404 value
	allNocCertificates, _ := GetAllNocX509RootCerts(suite)
	require.Equal(suite.T, 0, len(allNocCertificates))

	_, err := GetNocX509RootCerts(suite, testconstants.Vid)
	suite.AssertNotFound(err)

	_, err = GetNocX509CertsByVidAndSkid(suite, vid1, testconstants.NocRootCert1SubjectKeyID)
	suite.AssertNotFound(err)
	_, err = GetNocX509CertsByVidAndSkid(suite, vid1, testconstants.NocRootCert2SubjectKeyID)
	suite.AssertNotFound(err)
	_, err = GetNocX509CertsByVidAndSkid(suite, vid2, testconstants.NocRootCert3SubjectKeyID)
	suite.AssertNotFound(err)

	_, err = GetCert(suite, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	suite.AssertNotFound(err)

	// Alice and Jack are predefined Trustees
	aliceName := testconstants.AliceAccount
	aliceKeyInfo, err := suite.Kr.Key(aliceName)
	require.NoError(suite.T, err)
	address, err := aliceKeyInfo.GetAddress()
	require.NoError(suite.T, err)
	aliceAccount, err := test_dclauth.GetAccount(suite, address)
	require.NoError(suite.T, err)

	jackName := testconstants.JackAccount
	jackKeyInfo, err := suite.Kr.Key(jackName)
	require.NoError(suite.T, err)
	address, err = jackKeyInfo.GetAddress()
	require.NoError(suite.T, err)
	jackAccount, err := test_dclauth.GetAccount(suite, address)
	require.NoError(suite.T, err)

	// Register first Vendor account
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

	// Try to add intermediate cert
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
		Signer:                  vendor1Account.Address,
		Cert:                    testconstants.NocRootCert2,
		IsVidVerificationSigner: true,
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
	require.Equal(suite.T, testconstants.SchemaVersion, nocCertificates.Certs[0].SchemaVersion)
	require.Equal(suite.T, crttypes.CertificateType_OperationalPKI, nocCertificates.Certs[0].CertificateType)
	require.Equal(suite.T, testconstants.NocRootCert2Subject, nocCertificates.Certs[1].Subject)
	require.Equal(suite.T, testconstants.NocRootCert2SubjectKeyID, nocCertificates.Certs[1].SubjectKeyId)
	require.Equal(suite.T, crttypes.CertificateType_VIDSignerPKI, nocCertificates.Certs[1].CertificateType)
	require.Equal(suite.T, testconstants.SchemaVersion, nocCertificates.SchemaVersion)

	// Request NOC root certificate by VID1 and SKID1
	nocCertificatesByVidAndSkid, _ := GetNocX509CertsByVidAndSkid(suite, vid1, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(suite.T, 1, len(nocCertificatesByVidAndSkid.Certs))
	require.Equal(suite.T, testconstants.NocRootCert1Subject, nocCertificatesByVidAndSkid.Certs[0].Subject)
	require.Equal(suite.T, testconstants.NocRootCert1SubjectKeyID, nocCertificatesByVidAndSkid.Certs[0].SubjectKeyId)
	require.Equal(suite.T, float32(1), nocCertificatesByVidAndSkid.Tq)

	// Request NOC root certificate by VID1 and SKID2
	nocCertificatesByVidAndSkid, _ = GetNocX509CertsByVidAndSkid(suite, vid1, testconstants.NocRootCert2SubjectKeyID)
	require.Equal(suite.T, 1, len(nocCertificatesByVidAndSkid.Certs))
	require.Equal(suite.T, testconstants.NocRootCert2Subject, nocCertificatesByVidAndSkid.Certs[0].Subject)
	require.Equal(suite.T, testconstants.NocRootCert2SubjectKeyID, nocCertificatesByVidAndSkid.Certs[0].SubjectKeyId)
	require.Equal(suite.T, float32(1), nocCertificatesByVidAndSkid.Tq)

	// Request NOC root certificate by VID2 and SKID3
	nocCertificatesByVidAndSkid, _ = GetNocX509CertsByVidAndSkid(suite, vid2, testconstants.NocRootCert3SubjectKeyID)
	require.Equal(suite.T, 1, len(nocCertificatesByVidAndSkid.Certs))
	require.Equal(suite.T, testconstants.NocRootCert3Subject, nocCertificatesByVidAndSkid.Certs[0].Subject)
	require.Equal(suite.T, testconstants.NocRootCert3SubjectKeyID, nocCertificatesByVidAndSkid.Certs[0].SubjectKeyId)
	require.Equal(suite.T, float32(1), nocCertificatesByVidAndSkid.Tq)

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
	require.Equal(suite.T, crttypes.CertificateType_OperationalPKI, certsWithVid1[0].CertificateType)
	require.Equal(suite.T, testconstants.NocRootCert2Subject, certsWithVid1[1].Subject)
	require.Equal(suite.T, testconstants.NocRootCert2SubjectKeyID, certsWithVid1[1].SubjectKeyId)
	require.Equal(suite.T, crttypes.CertificateType_VIDSignerPKI, certsWithVid1[1].CertificateType)
	require.Equal(suite.T, testconstants.NocRootCert3Subject, certsWithVid2[0].Subject)
	require.Equal(suite.T, testconstants.NocRootCert3SubjectKeyID, certsWithVid2[0].SubjectKeyId)
	require.Equal(suite.T, crttypes.CertificateType_OperationalPKI, certsWithVid2[0].CertificateType)

	// Request NOC root certificate by Subject and SubjectKeyID
	certificate, _ := GetNocX509Cert(suite, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(suite.T, testconstants.NocRootCert1Subject, certificate.Subject)
	require.Equal(suite.T, testconstants.NocRootCert1SubjectKeyID, certificate.SubjectKeyId)
	require.Equal(suite.T, testconstants.NocRootCert1SubjectAsText, certificate.Certs[0].SubjectAsText)
	require.Equal(suite.T, 1, len(certificate.Certs))
	require.Equal(suite.T, testconstants.NocRootCert1, certificate.Certs[0].PemCert)
	require.Equal(suite.T, vendor1Account.Address, certificate.Certs[0].Owner)
	require.True(suite.T, certificate.Certs[0].IsRoot)

	// Request NOC root certificate by Subject
	subjectCertificates, _ := GetAllNocX509CertsBySubject(suite, testconstants.NocRootCert1Subject)
	require.Equal(suite.T, testconstants.NocRootCert1Subject, subjectCertificates.Subject)
	require.Equal(suite.T, 1, len(subjectCertificates.SubjectKeyIds))
	require.Equal(suite.T, testconstants.NocRootCert1SubjectKeyID, subjectCertificates.SubjectKeyIds[0])

	// Request NOC root certificate by SubjectKeyID
	certsBySubjectKeyID, _ := GetAllNocX509certsBySubjectKeyID(suite, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(suite.T, 1, len(certsBySubjectKeyID))
	require.Equal(suite.T, 1, len(certsBySubjectKeyID[0].Certs))
	require.Equal(suite.T, testconstants.NocRootCert1SubjectKeyID, certsBySubjectKeyID[0].Certs[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.NocRootCert1Subject, certsBySubjectKeyID[0].Certs[0].Subject)
	require.Equal(suite.T, testconstants.NocRootCert1SubjectAsText, certsBySubjectKeyID[0].Certs[0].SubjectAsText)

	// Request NOC root certificate by Subject and SubjectKeyID (using global command)
	globalCertificate, _ := GetCert(suite, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(suite.T, testconstants.NocRootCert1Subject, globalCertificate.Subject)
	require.Equal(suite.T, testconstants.NocRootCert1SubjectKeyID, globalCertificate.SubjectKeyId)
	require.Equal(suite.T, testconstants.NocRootCert1SubjectAsText, globalCertificate.Certs[0].SubjectAsText)
	require.Equal(suite.T, 1, len(globalCertificate.Certs))
	require.Equal(suite.T, testconstants.NocRootCert1, globalCertificate.Certs[0].PemCert)
	require.Equal(suite.T, vendor1Account.Address, globalCertificate.Certs[0].Owner)
	require.True(suite.T, globalCertificate.Certs[0].IsRoot)

	// Request NOC root certificate by Subject (using global command)
	globalSubjectCertificates, _ := GetAllCertsBySubject(suite, testconstants.NocRootCert1Subject)
	require.Equal(suite.T, testconstants.NocRootCert1Subject, globalSubjectCertificates.Subject)
	require.Equal(suite.T, 1, len(globalSubjectCertificates.SubjectKeyIds))
	require.Equal(suite.T, testconstants.NocRootCert1SubjectKeyID, globalSubjectCertificates.SubjectKeyIds[0])

	// Add intermediate NOC certificate
	msgAddNocCertificate := pkitypes.MsgAddNocX509IcaCert{
		Signer: vendor1Account.Address,
		Cert:   testconstants.NocCert1,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddNocCertificate}, vendor1Name, vendor1Account)
	require.NoError(suite.T, err)

	// Request certificate by VID1
	nocCert, _ := GetNocX509IcaCerts(suite, vid1)
	require.Equal(suite.T, 1, len(nocCert.Certs))
	require.Equal(suite.T, testconstants.NocCert1Subject, nocCert.Certs[0].Subject)
	require.Equal(suite.T, testconstants.NocCert1SubjectKeyID, nocCert.Certs[0].SubjectKeyId)

	// Request NOC ICA certificate by VID1
	nocCertificatesByVidAndSkid, _ = GetNocX509CertsByVidAndSkid(suite, vid1, testconstants.NocCert1SubjectKeyID)
	require.Equal(suite.T, 1, len(nocCertificatesByVidAndSkid.Certs))
	require.Equal(suite.T, testconstants.NocCert1Subject, nocCertificatesByVidAndSkid.Certs[0].Subject)
	require.Equal(suite.T, testconstants.NocCert1SubjectKeyID, nocCertificatesByVidAndSkid.Certs[0].SubjectKeyId)

	// Request Child certificates list
	childCertificates, _ := GetAllChildX509Certs(suite, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(suite.T, testconstants.NocRootCert1Subject, childCertificates.Issuer)
	require.Equal(suite.T, testconstants.NocRootCert1SubjectKeyID, childCertificates.AuthorityKeyId)
	require.Equal(suite.T, testconstants.NocCert1Subject, childCertificates.CertIds[0].Subject)
	require.Equal(suite.T, testconstants.NocCert1SubjectKeyID, childCertificates.CertIds[0].SubjectKeyId)

	// Try to add third intermediate NOC certificate with different vid
	msgAddNocCertificate = pkitypes.MsgAddNocX509IcaCert{
		Signer: vendor1Account.Address,
		Cert:   testconstants.NocCert2,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddNocCertificate}, vendor2Name, vendor2Account)
	require.Error(suite.T, err)

	// Add second intermediate NOC certificate
	msgAddNocCertificate = pkitypes.MsgAddNocX509IcaCert{
		Signer:                  vendor1Account.Address,
		Cert:                    testconstants.NocCert2,
		IsVidVerificationSigner: true,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddNocCertificate}, vendor1Name, vendor1Account)
	require.NoError(suite.T, err)

	// Request certificate by VID1
	nocCerts, _ := GetAllNocX509IcaCerts(suite)
	require.Equal(suite.T, 1, len(nocCerts))
	require.Equal(suite.T, 2, len(nocCerts[0].Certs))

	// Request NOC certificate by Subject and SubjectKeyID
	certs, _ := GetNocX509Cert(suite, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(suite.T, testconstants.NocCert1Subject, certs.Subject)
	require.Equal(suite.T, testconstants.NocCert1SubjectKeyID, certs.SubjectKeyId)
	require.Equal(suite.T, 1, len(certs.Certs))
	require.Equal(suite.T, testconstants.NocCert1, certs.Certs[0].PemCert)
	require.Equal(suite.T, vendor1Account.Address, certs.Certs[0].Owner)
	require.False(suite.T, certs.Certs[0].IsRoot)
	require.Equal(suite.T, crttypes.CertificateType_OperationalPKI, certs.Certs[0].CertificateType)

	certs, _ = GetNocX509Cert(suite, testconstants.NocCert2Subject, testconstants.NocCert2SubjectKeyID)
	require.Equal(suite.T, testconstants.NocCert2Subject, certs.Subject)
	require.Equal(suite.T, testconstants.NocCert2SubjectKeyID, certs.SubjectKeyId)
	require.Equal(suite.T, 1, len(certs.Certs))
	require.Equal(suite.T, testconstants.NocCert2, certs.Certs[0].PemCert)
	require.Equal(suite.T, vendor1Account.Address, certs.Certs[0].Owner)
	require.False(suite.T, certs.Certs[0].IsRoot)
	require.Equal(suite.T, crttypes.CertificateType_VIDSignerPKI, certs.Certs[0].CertificateType)

	// Check Revocation
	// Add NOC root certificate with same subject and skid as testconstants.NocCert1 cert
	msgAddNocRootCert := pkitypes.MsgAddNocX509RootCert{
		Signer: vendor1Account.Address,
		Cert:   testconstants.NocRootCert1Copy,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddNocRootCert}, vendor1Name, vendor1Account)
	require.NoError(suite.T, err)

	// Request NOC root certificate by VID1
	nocCertificates, _ = GetNocX509RootCerts(suite, vid1)
	require.Equal(suite.T, 3, len(nocCertificates.Certs))

	// Request NOC root certificate by VID1 and SKID1
	nocCertificatesByVidAndSkid, _ = GetNocX509CertsByVidAndSkid(suite, vid1, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(suite.T, 2, len(nocCertificatesByVidAndSkid.Certs))
	require.Equal(suite.T, testconstants.NocRootCert1Subject, nocCertificatesByVidAndSkid.Certs[0].Subject)
	require.Equal(suite.T, testconstants.NocRootCert1SubjectKeyID, nocCertificatesByVidAndSkid.Certs[0].SubjectKeyId)
	require.Equal(suite.T, float32(1), nocCertificatesByVidAndSkid.Tq)

	// Add NOC leaf certificate
	msgAddNocCert := pkitypes.MsgAddNocX509IcaCert{
		Signer: vendor1Account.Address,
		Cert:   testconstants.NocLeafCert1,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddNocCert}, vendor1Name, vendor1Account)
	require.NoError(suite.T, err)

	nocCerts, _ = GetAllNocX509IcaCerts(suite)
	require.Equal(suite.T, 1, len(nocCerts))
	require.Equal(suite.T, 3, len(nocCerts[0].Certs))

	// Try to revoke NOC 1 root with different serial number
	msgRevokeNocRootCert := pkitypes.MsgRevokeNocX509RootCert{
		Signer:       vendor1Account.Address,
		Subject:      testconstants.NocRootCert1Subject,
		SubjectKeyId: testconstants.NocRootCert1SubjectKeyID,
		SerialNumber: "1234",
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRevokeNocRootCert}, vendor1Name, vendor1Account)
	require.Error(suite.T, err)

	// Try to revoke NOC 1 root with another Vendor Account
	msgRevokeNocRootCert = pkitypes.MsgRevokeNocX509RootCert{
		Signer:       vendor2Account.Address,
		Subject:      testconstants.NocRootCert1Subject,
		SubjectKeyId: testconstants.NocRootCert1SubjectKeyID,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRevokeNocRootCert}, vendor2Name, vendor2Account)
	require.Error(suite.T, err)

	// Add ICA certificate
	msgAddNocCert = pkitypes.MsgAddNocX509IcaCert{
		Signer: vendor1Account.Address,
		Cert:   testconstants.NocCert1Copy,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddNocCert}, vendor1Name, vendor1Account)
	require.NoError(suite.T, err)

	// Revoke intermediate certificate with invalid serialNumber
	msgRevokeCert := pkitypes.MsgRevokeNocX509IcaCert{
		Subject:      testconstants.NocCert1Subject,
		SubjectKeyId: testconstants.NocCert1SubjectKeyID,
		SerialNumber: "invalid",
		Signer:       vendor1Account.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRevokeCert}, vendor1Name, vendor1Account)
	require.Error(suite.T, err)

	// Revoke intermediate certificate with serialNumber 1 only(child certs should not be removed)
	msgRevokeCert.SerialNumber = testconstants.NocCert1SerialNumber
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRevokeCert}, vendor1Name, vendor1Account)
	require.NoError(suite.T, err)

	// Request revoked certificate
	revokedCerts, _ := GetRevokedNocX509IcaCert(suite, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(suite.T, 1, len(revokedCerts.Certs))
	require.Equal(suite.T, testconstants.NocCert1Subject, revokedCerts.Subject)
	require.Equal(suite.T, testconstants.NocCert1SubjectKeyID, revokedCerts.SubjectKeyId)
	require.Equal(suite.T, testconstants.NocCert1SerialNumber, revokedCerts.Certs[0].SerialNumber)

	// Check approved certificates
	certs, _ = GetNocX509Cert(suite, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(suite.T, 1, len(certs.Certs))
	require.Equal(suite.T, testconstants.NocCert1CopySerialNumber, certs.Certs[0].SerialNumber)
	certs, _ = GetNocX509Cert(suite, testconstants.NocLeafCert1Subject, testconstants.NocLeafCert1SubjectKeyID)
	require.Equal(suite.T, 1, len(certs.Certs))
	require.Equal(suite.T, testconstants.NocLeafCert1SerialNumber, certs.Certs[0].SerialNumber)

	icaCerts, _ := GetNocX509IcaCerts(suite, vid1)
	require.Equal(suite.T, 3, len(icaCerts.Certs))

	// Revoke Root certificate with serialNumber(child certs should not be removed)
	msgRevokeRootCert := pkitypes.MsgRevokeNocX509RootCert{
		Signer:       vendor1Account.Address,
		Subject:      testconstants.NocRootCert1Subject,
		SubjectKeyId: testconstants.NocRootCert1SubjectKeyID,
		SerialNumber: testconstants.NocRootCert1SerialNumber,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRevokeRootCert}, vendor1Name, vendor1Account)
	require.NoError(suite.T, err)

	// Request revoked root certificate
	revokedRootCerts, _ := GetRevokedNocX509RootCert(suite, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(suite.T, 1, len(revokedRootCerts.Certs))
	require.Equal(suite.T, testconstants.NocRootCert1Subject, revokedRootCerts.Subject)
	require.Equal(suite.T, testconstants.NocRootCert1SubjectKeyID, revokedRootCerts.SubjectKeyId)
	require.Equal(suite.T, testconstants.NocRootCert1SerialNumber, revokedRootCerts.Certs[0].SerialNumber)

	// Check approved certificate
	certs, _ = GetNocX509Cert(suite, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(suite.T, 1, len(certs.Certs))
	require.Equal(suite.T, testconstants.NocRootCert1CopySerialNumber, certs.Certs[0].SerialNumber)
	certs, _ = GetNocX509Cert(suite, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(suite.T, 1, len(certs.Certs))
	require.Equal(suite.T, testconstants.NocCert1CopySerialNumber, certs.Certs[0].SerialNumber)

	nocRootCerts, _ := GetNocX509CertsByVidAndSkid(suite, vid1, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(suite.T, 1, len(nocRootCerts.Certs))
	require.Equal(suite.T, testconstants.NocRootCert1CopySerialNumber, nocRootCerts.Certs[0].SerialNumber)

	// Remove ICA certificate with invalid serialNumber
	msgRemoveCert := pkitypes.MsgRemoveNocX509IcaCert{
		Signer:       vendor1Account.Address,
		Subject:      testconstants.NocCert1Subject,
		SubjectKeyId: testconstants.NocCert1SubjectKeyID,
		SerialNumber: "invalid",
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRemoveCert}, vendor1Name, vendor1Account)
	require.Error(suite.T, err)

	// Try to Remove ICA certificate by subject and subject key id when sender is not Vendor account
	msgRemoveCert = pkitypes.MsgRemoveNocX509IcaCert{
		Signer:       aliceAccount.Address,
		Subject:      testconstants.NocCert1Subject,
		SubjectKeyId: testconstants.NocCert1SubjectKeyID,
		SerialNumber: testconstants.NocCert1CopySerialNumber,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRemoveCert}, aliceName, aliceAccount)
	require.Error(suite.T, err)

	// Remove revoked ICA certificate by subject and subject key id
	msgRemoveCert = pkitypes.MsgRemoveNocX509IcaCert{
		Signer:       vendor1Account.Address,
		Subject:      testconstants.NocCert1Subject,
		SubjectKeyId: testconstants.NocCert1SubjectKeyID,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRemoveCert}, vendor1Name, vendor1Account)
	require.NoError(suite.T, err)

	// Request NOC ICA certificate by VID1
	_, err = GetNocX509CertsByVidAndSkid(suite, vid1, testconstants.NocCert1Subject)
	suite.AssertNotFound(err)

	// Check that two intermediate ICA certificates are removed
	_, err = GetRevokedNocX509IcaCert(suite, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	suite.AssertNotFound(err)
	_, err = GetNocX509Cert(suite, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	suite.AssertNotFound(err)

	certificates := GetNocX509IcaCertsBySubjectAndSKID(suite, vid1, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Empty(suite.T, certificates)

	// Remove leaf ICA certificate by subject and subject key id
	msgRemoveCert = pkitypes.MsgRemoveNocX509IcaCert{
		Signer:       vendor1Account.Address,
		Subject:      testconstants.NocLeafCert1Subject,
		SubjectKeyId: testconstants.NocLeafCert1SubjectKeyID,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRemoveCert}, vendor1Name, vendor1Account)
	require.NoError(suite.T, err)

	// Check that leaf ICA certificate is removed
	_, err = GetNocX509Cert(suite, testconstants.NocLeafCert1Subject, testconstants.NocLeafCert1SubjectKeyID)
	suite.AssertNotFound(err)

	// Request NOC ICA certificate by VID1
	_, err = GetNocX509CertsByVidAndSkid(suite, vid1, testconstants.NocLeafCert1SubjectKeyID)
	suite.AssertNotFound(err)

	certificates = GetNocX509IcaCertsBySubjectAndSKID(suite, vid1, testconstants.NocLeafCert1Subject, testconstants.NocLeafCert1SubjectKeyID)
	require.Empty(suite.T, certificates)

	// Remove ICA certificates by subject, subject key id and serial number
	// Add ICA certificates
	msgAddNocCert = pkitypes.MsgAddNocX509IcaCert{
		Signer: vendor1Account.Address,
		Cert:   testconstants.NocCert1,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddNocCert}, vendor1Name, vendor1Account)
	require.NoError(suite.T, err)

	msgAddNocCert = pkitypes.MsgAddNocX509IcaCert{
		Signer: vendor1Account.Address,
		Cert:   testconstants.NocCert1Copy,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddNocCert}, vendor1Name, vendor1Account)
	require.NoError(suite.T, err)

	msgAddNocCert = pkitypes.MsgAddNocX509IcaCert{
		Signer: vendor1Account.Address,
		Cert:   testconstants.NocLeafCert1,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddNocCert}, vendor1Name, vendor1Account)
	require.NoError(suite.T, err)

	// Remove ICA certificate by serial number
	msgRemoveCert = pkitypes.MsgRemoveNocX509IcaCert{
		Signer:       vendor1Account.Address,
		Subject:      testconstants.NocCert1Subject,
		SubjectKeyId: testconstants.NocCert1SubjectKeyID,
		SerialNumber: testconstants.NocCert1SerialNumber,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRemoveCert}, vendor1Name, vendor1Account)
	require.NoError(suite.T, err)

	// Check that leaf and ICA with different serial number is not removed
	certs, _ = GetNocX509Cert(suite, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(suite.T, 1, len(certs.Certs))
	require.Equal(suite.T, testconstants.NocCert1CopySerialNumber, certs.Certs[0].SerialNumber)
	certs, _ = GetNocX509Cert(suite, testconstants.NocLeafCert1Subject, testconstants.NocLeafCert1SubjectKeyID)
	require.Equal(suite.T, 1, len(certs.Certs))

	certificates = GetNocX509IcaCertsBySubjectAndSKID(suite, vid1, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(suite.T, 1, len(certificates))
	require.Equal(suite.T, testconstants.NocCert1CopySerialNumber, certificates[0].SerialNumber)
	certificates = GetNocX509IcaCertsBySubjectAndSKID(suite, vid1, testconstants.NocLeafCert1Subject, testconstants.NocLeafCert1SubjectKeyID)
	require.Equal(suite.T, 1, len(certificates))

	// Revoke root cert and its child
	msgRevokeRootCert = pkitypes.MsgRevokeNocX509RootCert{
		Signer:       vendor1Account.Address,
		Subject:      testconstants.NocRootCert1Subject,
		SubjectKeyId: testconstants.NocRootCert1SubjectKeyID,
		RevokeChild:  true,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRevokeRootCert}, vendor1Name, vendor1Account)
	require.NoError(suite.T, err)

	// Check that all 3 certificates are revoked
	revokedRootCerts, _ = GetRevokedNocX509RootCert(suite, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(suite.T, 2, len(revokedRootCerts.Certs))
	revokedCerts, _ = GetRevokedNocX509IcaCert(suite, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(suite.T, 1, len(revokedCerts.Certs))
	revokedCerts, _ = GetRevokedNocX509IcaCert(suite, testconstants.NocLeafCert1Subject, testconstants.NocLeafCert1SubjectKeyID)
	require.Equal(suite.T, 1, len(revokedCerts.Certs))

	_, err = GetNocX509Cert(suite, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	suite.AssertNotFound(err)
	_, err = GetNocX509Cert(suite, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	suite.AssertNotFound(err)
	_, err = GetNocX509Cert(suite, testconstants.NocLeafCert1Subject, testconstants.NocLeafCert1SubjectKeyID)
	suite.AssertNotFound(err)

	_, err = GetNocX509CertsByVidAndSkid(suite, vid1, testconstants.NocRootCert1SubjectKeyID)
	suite.AssertNotFound(err)
	_, err = GetNocX509CertsByVidAndSkid(suite, vid1, testconstants.NocCert1SubjectKeyID)
	suite.AssertNotFound(err)
	_, err = GetNocX509CertsByVidAndSkid(suite, vid1, testconstants.NocLeafCert1SubjectKeyID)
	suite.AssertNotFound(err)

	_, err = GetNocX509CertsByVidAndSkid(suite, vid1, testconstants.NocRootCert1SubjectKeyID)
	suite.AssertNotFound(err)
	certificates = GetNocX509IcaCertsBySubjectAndSKID(suite, vid1, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Empty(suite.T, certificates)
	certificates = GetNocX509IcaCertsBySubjectAndSKID(suite, vid1, testconstants.NocLeafCert1Subject, testconstants.NocLeafCert1SubjectKeyID)
	require.Empty(suite.T, certificates)

	// Remove revoked NOC root certificate by invalid serial number
	msgRemoveRootCert := pkitypes.MsgRemoveNocX509RootCert{
		Signer:       vendor1Account.Address,
		Subject:      testconstants.NocRootCert1Subject,
		SubjectKeyId: testconstants.NocRootCert1SubjectKeyID,
		SerialNumber: "invalid",
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRemoveRootCert}, vendor1Name, vendor1Account)
	require.Error(suite.T, err)

	// Remove revoked NOC root certificate by serial number
	msgRemoveRootCert = pkitypes.MsgRemoveNocX509RootCert{
		Signer:       vendor1Account.Address,
		Subject:      testconstants.NocRootCert1Subject,
		SubjectKeyId: testconstants.NocRootCert1SubjectKeyID,
		SerialNumber: testconstants.NocRootCert1SerialNumber,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRemoveRootCert}, vendor1Name, vendor1Account)
	require.NoError(suite.T, err)

	// Check that NOC root certificate is removed
	revokedRootCerts, _ = GetRevokedNocX509RootCert(suite, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(suite.T, 1, len(revokedRootCerts.Certs))
	require.Equal(suite.T, testconstants.NocRootCert1CopySerialNumber, revokedRootCerts.Certs[0].SerialNumber)
	_, err = GetNocX509Cert(suite, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	suite.AssertNotFound(err)

	// Add root NOC certificate
	msgAddNocRootCert = pkitypes.MsgAddNocX509RootCert{
		Signer: vendor1Account.Address,
		Cert:   testconstants.NocRootCert1,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddNocRootCert}, vendor1Name, vendor1Account)
	require.NoError(suite.T, err)

	// Remove revoked ICA certificate to re-add it again
	msgRemoveCert = pkitypes.MsgRemoveNocX509IcaCert{
		Signer:       vendor1Account.Address,
		Subject:      testconstants.NocCert1Subject,
		SubjectKeyId: testconstants.NocCert1SubjectKeyID,
		SerialNumber: "",
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRemoveCert}, vendor1Name, vendor1Account)
	require.NoError(suite.T, err)

	// Add ICA certificates
	msgAddNocCert = pkitypes.MsgAddNocX509IcaCert{
		Signer: vendor1Account.Address,
		Cert:   testconstants.NocCert1,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddNocCert}, vendor1Name, vendor1Account)
	require.NoError(suite.T, err)

	certs, _ = GetNocX509Cert(suite, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	require.Equal(suite.T, 1, len(certs.Certs))

	// Remove revoked NOC root certificates
	msgRemoveRootCert = pkitypes.MsgRemoveNocX509RootCert{
		Signer:       vendor1Account.Address,
		Subject:      testconstants.NocRootCert1Subject,
		SubjectKeyId: testconstants.NocRootCert1SubjectKeyID,
		SerialNumber: "",
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRemoveRootCert}, vendor1Name, vendor1Account)
	require.NoError(suite.T, err)

	// Check that certificates are removed
	_, err = GetNocX509Cert(suite, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	suite.AssertNotFound(err)
	_, err = GetRevokedNocX509RootCert(suite, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	suite.AssertNotFound(err)
	_, err = GetNocX509CertsByVidAndSkid(suite, vid1, testconstants.NocRootCert1SubjectKeyID)
	suite.AssertNotFound(err)
	_, err = GetCert(suite, testconstants.NocRootCert1Subject, testconstants.NocRootCert1SubjectKeyID)
	suite.AssertNotFound(err)

	// Check that child is not removed
	certs, _ = GetNocX509Cert(suite, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(suite.T, 1, len(certs.Certs))
	_, err = GetRevokedNocX509IcaCert(suite, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	suite.AssertNotFound(err)
	certificates = GetNocX509IcaCertsBySubjectAndSKID(suite, vid1, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(suite.T, 1, len(certificates))
	require.Equal(suite.T, testconstants.NocCert1SerialNumber, certificates[0].SerialNumber)
	globalCerts, _ := GetCert(suite, testconstants.NocCert1Subject, testconstants.NocCert1SubjectKeyID)
	require.Equal(suite.T, 1, len(globalCerts.Certs))
}
