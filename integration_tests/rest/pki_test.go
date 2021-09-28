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

package rest_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki"
)

//nolint:godox
/*
	To Run test you need:
		* Run LocalNet with: `make install && make localnet_init && make localnet_start`
		* run RPC service with `dclcli rest-server --chain-id dclchain`

	TODO: provide tests for error cases
*/

// nolint:godox,funlen
// FIXME: `GetX509CertChain` calls within this test may fail with EOF on an attempt to read the response
// from REST API server and so leave the resulting `Certificates.Items` slice empty.
// The issue seems to be caused by the connection being closed due to timeout while REST API server
// is gathering the reply which consists of multiple replies from the pool.
// However, `net/http.Client` does not report the request as timed out while actually it seems to be so.
//nolint:funlen
func TestPkiDemo(t *testing.T) {
	// Get key info for Jack
	jackKeyInfo, _ := utils.GetKeyInfo(testconstants.JackAccount)

	// Get key info for Alice
	aliceKeyInfo, _ := utils.GetKeyInfo(testconstants.AliceAccount)

	// Create account for Anna
	userKeyInfo := utils.CreateNewAccount(auth.AccountRoles{}, 0)

	// Request all proposed certificates
	proposedCertificates, _ := utils.GetAllProposedX509RootCerts()
	require.Equal(t, 0, len(proposedCertificates.Items))

	// Request all approved certificates
	certificates, _ := utils.GetAllX509Certs()
	require.Equal(t, 0, len(certificates.Items))

	// User (Not Trustee) propose Root certificate
	msgProposeAddX509RootCert := pki.MsgProposeAddX509RootCert{
		Cert:   testconstants.RootCertPem,
		Signer: userKeyInfo.Address,
	}
	_, _ = utils.ProposeAddX509RootCert(msgProposeAddX509RootCert, userKeyInfo.Name, testconstants.Passphrase)

	// Request all proposed Root certificates
	proposedCertificates, _ = utils.GetAllProposedX509RootCerts()
	require.Equal(t, 1, len(proposedCertificates.Items))
	require.Equal(t, testconstants.RootSubject, proposedCertificates.Items[0].Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, proposedCertificates.Items[0].SubjectKeyID)

	// Request all approved certificates
	certificates, _ = utils.GetAllX509Certs()
	require.Equal(t, 0, len(certificates.Items))

	// Request proposed Root certificate
	proposedCertificate, _ :=
		utils.GetProposedX509RootCert(testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, testconstants.RootCertPem, proposedCertificate.PemCert)
	require.Equal(t, userKeyInfo.Address, proposedCertificate.Owner)
	require.Nil(t, proposedCertificate.Approvals)

	// Jack (Trustee) approve Root certificate
	msgApproveAddX509RootCert := pki.MsgApproveAddX509RootCert{
		Subject:      proposedCertificate.Subject,
		SubjectKeyID: proposedCertificate.SubjectKeyID,
		Signer:       jackKeyInfo.Address,
	}
	_, _ = utils.ApproveAddX509RootCert(msgApproveAddX509RootCert, jackKeyInfo.Name, testconstants.Passphrase)

	// Request all proposed Root certificates
	proposedCertificates, _ = utils.GetAllProposedX509RootCerts()
	require.Equal(t, 1, len(proposedCertificates.Items))
	require.Equal(t, testconstants.RootSubject, proposedCertificates.Items[0].Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, proposedCertificates.Items[0].SubjectKeyID)

	// Request all approved certificates
	certificates, _ = utils.GetAllX509Certs()
	require.Equal(t, 0, len(certificates.Items))

	// Request proposed Root certificate
	proposedCertificate, _ =
		utils.GetProposedX509RootCert(testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, testconstants.RootCertPem, proposedCertificate.PemCert)
	require.Equal(t, userKeyInfo.Address, proposedCertificate.Owner)
	require.Equal(t, []sdk.AccAddress{jackKeyInfo.Address}, proposedCertificate.Approvals)

	// Alice (Trustee) approve Root certificate
	secondMsgApproveAddX509RootCert := pki.MsgApproveAddX509RootCert{
		Subject:      proposedCertificate.Subject,
		SubjectKeyID: proposedCertificate.SubjectKeyID,
		Signer:       aliceKeyInfo.Address,
	}
	_, _ = utils.ApproveAddX509RootCert(secondMsgApproveAddX509RootCert, aliceKeyInfo.Name, testconstants.Passphrase)

	// Request all proposed Root certificates
	proposedCertificates, _ = utils.GetAllProposedX509RootCerts()
	require.Equal(t, 0, len(proposedCertificates.Items))

	// Request all approved certificates
	certificates, _ = utils.GetAllX509Certs()
	require.Equal(t, 1, len(certificates.Items))
	require.Equal(t, testconstants.RootSubject, certificates.Items[0].Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, certificates.Items[0].SubjectKeyID)

	// Request all approved Root certificates
	certificates, _ = utils.GetAllX509RootCerts()
	require.Equal(t, 1, len(certificates.Items))
	require.Equal(t, testconstants.RootSubject, certificates.Items[0].Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, certificates.Items[0].SubjectKeyID)

	// Request approved Root certificate
	certificate, _ := utils.GetX509Cert(testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, testconstants.RootCertPem, certificate.PemCert)
	require.Equal(t, userKeyInfo.Address, certificate.Owner)
	require.True(t, certificate.IsRoot)

	// Request certificate chain for Root certificate
	certificateChain, _ := utils.GetX509CertChain(testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Len(t, certificateChain.Items, 1)
	require.Equal(t, testconstants.RootCertPem, certificateChain.Items[0].PemCert)
	require.Equal(t, userKeyInfo.Address, certificateChain.Items[0].Owner)
	require.True(t, certificateChain.Items[0].IsRoot)

	// User (Not Trustee) add Intermediate certificate
	msgAddX509Cert := pki.MsgAddX509Cert{
		Cert:   testconstants.IntermediateCertPem,
		Signer: userKeyInfo.Address,
	}
	_, _ = utils.AddX509Cert(msgAddX509Cert, userKeyInfo.Name, testconstants.Passphrase)

	// Request all proposed Root certificates
	proposedCertificates, _ = utils.GetAllProposedX509RootCerts()
	require.Equal(t, 0, len(proposedCertificates.Items))

	// Request all approved certificates
	certificates, _ = utils.GetAllX509Certs()
	require.Equal(t, 2, len(certificates.Items))
	require.Equal(t, testconstants.IntermediateSubject, certificates.Items[0].Subject)
	require.Equal(t, testconstants.IntermediateSubjectKeyID, certificates.Items[0].SubjectKeyID)
	require.Equal(t, testconstants.RootSubject, certificates.Items[1].Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, certificates.Items[1].SubjectKeyID)

	// Request all approved Root certificates
	certificates, _ = utils.GetAllX509RootCerts()
	require.Equal(t, 1, len(certificates.Items))
	require.Equal(t, testconstants.RootSubject, certificates.Items[0].Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, certificates.Items[0].SubjectKeyID)

	// Request Intermediate certificate
	certificate, _ = utils.GetX509Cert(testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Equal(t, testconstants.IntermediateCertPem, certificate.PemCert)
	require.Equal(t, userKeyInfo.Address, certificate.Owner)
	require.False(t, certificate.IsRoot)

	// Request certificate chain for Intermediate certificate
	certificateChain, _ = utils.GetX509CertChain(testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Len(t, certificateChain.Items, 2)
	require.Equal(t, testconstants.IntermediateCertPem, certificateChain.Items[0].PemCert)
	require.Equal(t, userKeyInfo.Address, certificateChain.Items[0].Owner)
	require.False(t, certificateChain.Items[0].IsRoot)
	require.Equal(t, testconstants.RootCertPem, certificateChain.Items[1].PemCert)
	require.Equal(t, userKeyInfo.Address, certificateChain.Items[1].Owner)
	require.True(t, certificateChain.Items[1].IsRoot)

	// Alice (Trustee) add Leaf certificate
	secondMsgAddX509Cert := pki.MsgAddX509Cert{
		Cert:   testconstants.LeafCertPem,
		Signer: aliceKeyInfo.Address,
	}
	_, _ = utils.AddX509Cert(secondMsgAddX509Cert, aliceKeyInfo.Name, testconstants.Passphrase)

	// Request all proposed Root certificates
	proposedCertificates, _ = utils.GetAllProposedX509RootCerts()
	require.Equal(t, 0, len(proposedCertificates.Items))

	// Request all approved certificates
	certificates, _ = utils.GetAllX509Certs()
	require.Equal(t, 3, len(certificates.Items))
	require.Equal(t, testconstants.IntermediateSubject, certificates.Items[0].Subject)
	require.Equal(t, testconstants.IntermediateSubjectKeyID, certificates.Items[0].SubjectKeyID)
	require.Equal(t, testconstants.LeafSubject, certificates.Items[1].Subject)
	require.Equal(t, testconstants.LeafSubjectKeyID, certificates.Items[1].SubjectKeyID)
	require.Equal(t, testconstants.RootSubject, certificates.Items[2].Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, certificates.Items[2].SubjectKeyID)

	// Request all approved Root certificates
	certificates, _ = utils.GetAllX509RootCerts()
	require.Equal(t, 1, len(certificates.Items))
	require.Equal(t, testconstants.RootSubject, certificates.Items[0].Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, certificates.Items[0].SubjectKeyID)

	// Request Leaf certificate
	certificate, _ = utils.GetX509Cert(testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Equal(t, testconstants.LeafCertPem, certificate.PemCert)
	require.Equal(t, aliceKeyInfo.Address, certificate.Owner)
	require.False(t, certificate.IsRoot)

	// Request certificate chain for Leaf certificate
	certificateChain, _ = utils.GetX509CertChain(testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Len(t, certificateChain.Items, 3)
	require.Equal(t, testconstants.LeafCertPem, certificateChain.Items[0].PemCert)
	require.Equal(t, aliceKeyInfo.Address, certificateChain.Items[0].Owner)
	require.False(t, certificateChain.Items[0].IsRoot)
	require.Equal(t, testconstants.IntermediateCertPem, certificateChain.Items[1].PemCert)
	require.Equal(t, userKeyInfo.Address, certificateChain.Items[1].Owner)
	require.False(t, certificateChain.Items[1].IsRoot)
	require.Equal(t, testconstants.RootCertPem, certificateChain.Items[2].PemCert)
	require.Equal(t, userKeyInfo.Address, certificateChain.Items[2].Owner)
	require.True(t, certificateChain.Items[2].IsRoot)

	// Request all Subject certificates
	certificates, _ = utils.GetAllSubjectX509Certs(testconstants.LeafSubject)
	require.Equal(t, 1, len(certificates.Items))
	require.Equal(t, testconstants.LeafSubjectKeyID, certificates.Items[0].SubjectKeyID)

	// Request all Root certificates proposed to revoke
	proposedCertificateRevocations, _ := utils.GetAllProposedX509RootCertsToRevoke()
	require.Equal(t, 0, len(proposedCertificateRevocations.Items))

	// Request all revoked certificates
	revokedCertificates, _ := utils.GetAllRevokedX509Certs()
	require.Equal(t, 0, len(revokedCertificates.Items))

	// User (Not Trustee) revokes Intermediate certificate. This must also revoke its child - Leaf certificate.
	msgRevokeX509Cert := pki.MsgRevokeX509Cert{
		Subject:      testconstants.IntermediateSubject,
		SubjectKeyID: testconstants.IntermediateSubjectKeyID,
		Signer:       userKeyInfo.Address,
	}
	_, _ = utils.RevokeX509Cert(msgRevokeX509Cert, userKeyInfo.Name, testconstants.Passphrase)

	// Request all Root certificates proposed to revoke
	proposedCertificateRevocations, _ = utils.GetAllProposedX509RootCertsToRevoke()
	require.Equal(t, 0, len(proposedCertificateRevocations.Items))

	// Request all revoked certificates
	revokedCertificates, _ = utils.GetAllRevokedX509Certs()
	require.Equal(t, 2, len(revokedCertificates.Items))
	require.Equal(t, testconstants.IntermediateSubject, revokedCertificates.Items[0].Subject)
	require.Equal(t, testconstants.IntermediateSubjectKeyID, revokedCertificates.Items[0].SubjectKeyID)
	require.Equal(t, testconstants.LeafSubject, revokedCertificates.Items[1].Subject)
	require.Equal(t, testconstants.LeafSubjectKeyID, revokedCertificates.Items[1].SubjectKeyID)

	// Request all revoked Root certificates
	revokedCertificates, _ = utils.GetAllRevokedX509RootCerts()
	require.Equal(t, 0, len(revokedCertificates.Items))

	// Request revoked Intermediate certificate
	revokedCertificate, _ := utils.GetRevokedX509Cert(
		testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Equal(t, testconstants.IntermediateCertPem, revokedCertificate.PemCert)
	require.Equal(t, userKeyInfo.Address, revokedCertificate.Owner)
	require.False(t, revokedCertificate.IsRoot)

	// Request revoked Leaf certificate
	revokedCertificate, _ = utils.GetRevokedX509Cert(testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Equal(t, testconstants.LeafCertPem, revokedCertificate.PemCert)
	require.Equal(t, aliceKeyInfo.Address, revokedCertificate.Owner)
	require.False(t, revokedCertificate.IsRoot)

	// Request all approved certificates
	certificates, _ = utils.GetAllX509Certs()
	require.Equal(t, 1, len(certificates.Items))
	require.Equal(t, testconstants.RootSubject, certificates.Items[0].Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, certificates.Items[0].SubjectKeyID)

	// Jack (Trustee) proposes to revoke Root certificate
	msgProposeRevokeX509RootCert := pki.MsgProposeRevokeX509RootCert{
		Subject:      testconstants.RootSubject,
		SubjectKeyID: testconstants.RootSubjectKeyID,
		Signer:       jackKeyInfo.Address,
	}
	_, _ = utils.ProposeRevokeX509RootCert(msgProposeRevokeX509RootCert, jackKeyInfo.Name, testconstants.Passphrase)

	// Request all Root certificates proposed to revoke
	proposedCertificateRevocations, _ = utils.GetAllProposedX509RootCertsToRevoke()
	require.Equal(t, 1, len(proposedCertificateRevocations.Items))
	require.Equal(t, testconstants.RootSubject, proposedCertificateRevocations.Items[0].Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, proposedCertificateRevocations.Items[0].SubjectKeyID)

	// Request all revoked certificates
	revokedCertificates, _ = utils.GetAllRevokedX509Certs()
	require.Equal(t, 2, len(revokedCertificates.Items))
	require.Equal(t, testconstants.IntermediateSubject, revokedCertificates.Items[0].Subject)
	require.Equal(t, testconstants.IntermediateSubjectKeyID, revokedCertificates.Items[0].SubjectKeyID)
	require.Equal(t, testconstants.LeafSubject, revokedCertificates.Items[1].Subject)
	require.Equal(t, testconstants.LeafSubjectKeyID, revokedCertificates.Items[1].SubjectKeyID)

	// Request all revoked Root certificates
	revokedCertificates, _ = utils.GetAllRevokedX509RootCerts()
	require.Equal(t, 0, len(revokedCertificates.Items))

	// Request Root certificate proposed to revoke
	proposedCertificateRevocation, _ :=
		utils.GetProposedX509RootCertToRevoke(testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, []sdk.AccAddress{jackKeyInfo.Address}, proposedCertificateRevocation.Approvals)

	// Request all approved certificates
	certificates, _ = utils.GetAllX509Certs()
	require.Equal(t, 1, len(certificates.Items))
	require.Equal(t, testconstants.RootSubject, certificates.Items[0].Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, certificates.Items[0].SubjectKeyID)

	// Alice (Trustee) approves to revoke Root certificate
	msgApproveRevokeX509RootCert := pki.MsgApproveRevokeX509RootCert{
		Subject:      proposedCertificate.Subject,
		SubjectKeyID: proposedCertificate.SubjectKeyID,
		Signer:       aliceKeyInfo.Address,
	}
	_, _ = utils.ApproveRevokeX509RootCert(msgApproveRevokeX509RootCert, aliceKeyInfo.Name, testconstants.Passphrase)

	// Request all Root certificates proposed to revoke
	proposedCertificateRevocations, _ = utils.GetAllProposedX509RootCertsToRevoke()
	require.Equal(t, 0, len(proposedCertificateRevocations.Items))

	// Request all revoked certificates
	revokedCertificates, _ = utils.GetAllRevokedX509Certs()
	require.Equal(t, 3, len(revokedCertificates.Items))
	require.Equal(t, testconstants.IntermediateSubject, revokedCertificates.Items[0].Subject)
	require.Equal(t, testconstants.IntermediateSubjectKeyID, revokedCertificates.Items[0].SubjectKeyID)
	require.Equal(t, testconstants.LeafSubject, revokedCertificates.Items[1].Subject)
	require.Equal(t, testconstants.LeafSubjectKeyID, revokedCertificates.Items[1].SubjectKeyID)
	require.Equal(t, testconstants.RootSubject, revokedCertificates.Items[2].Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, revokedCertificates.Items[2].SubjectKeyID)

	// Request all revoked Root certificates
	revokedCertificates, _ = utils.GetAllRevokedX509RootCerts()
	require.Equal(t, 1, len(revokedCertificates.Items))
	require.Equal(t, testconstants.RootSubject, revokedCertificates.Items[0].Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, revokedCertificates.Items[0].SubjectKeyID)

	// Request revoked Root certificate
	revokedCertificate, _ = utils.GetRevokedX509Cert(testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, testconstants.RootCertPem, revokedCertificate.PemCert)
	require.Equal(t, userKeyInfo.Address, revokedCertificate.Owner)
	require.True(t, revokedCertificate.IsRoot)

	// Request all approved certificates
	certificates, _ = utils.GetAllX509Certs()
	require.Equal(t, 0, len(certificates.Items))
}
