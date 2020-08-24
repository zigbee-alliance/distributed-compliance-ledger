package rest_test

import (
	"testing"

	testconstants "git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/utils"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

//nolint:godox
/*
	To Run test you need:
		* Run LocalNet with: `make install && make localnet_init && make localnet_start`
		* run RPC service with `zblcli rest-server --chain-id zblchain`

	TODO: provide tests for error cases
*/

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
	userKeyInfo := utils.CreateNewAccount(auth.AccountRoles{})

	// Query all proposed certificates
	proposedCertificates, _ := utils.GetAllProposedX509RootCerts()
	require.Equal(t, 0, len(proposedCertificates.Items))

	// Query all approved certificates
	certificates, _ := utils.GetAllX509Certs()
	require.Equal(t, 0, len(certificates.Items))

	// User (Not Trustee) propose Root certificate
	msgProposeAddX509RootCert := pki.MsgProposeAddX509RootCert{
		Cert:   testconstants.RootCertPem,
		Signer: userKeyInfo.Address,
	}
	_, _ = utils.ProposeAddX509RootCert(msgProposeAddX509RootCert,
		userKeyInfo.Name, testconstants.Passphrase)

	//Request all proposed Root certificates
	proposedCertificates, _ = utils.GetAllProposedX509RootCerts()
	require.Equal(t, 1, len(proposedCertificates.Items))
	require.Equal(t, testconstants.RootSubject, proposedCertificates.Items[0].Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, proposedCertificates.Items[0].SubjectKeyID)

	// Request proposed Root certificate
	proposedCertificate, _ :=
		utils.GetProposedX509RootCert(testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, msgProposeAddX509RootCert.Cert, proposedCertificate.PemCert)
	require.Equal(t, msgProposeAddX509RootCert.Signer, proposedCertificate.Owner)
	require.Nil(t, proposedCertificate.Approvals)

	// Query all approved certificates
	certificates, _ = utils.GetAllX509Certs()
	require.Equal(t, 0, len(certificates.Items))

	// Request all active Root certificates
	certificates, _ = utils.GetAllX509RootCerts()
	require.Equal(t, 0, len(certificates.Items))

	// Jack (Trustee) approve Root certificate
	msgApproveAddX509RootCert := pki.MsgApproveAddX509RootCert{
		Subject:      proposedCertificate.Subject,
		SubjectKeyID: proposedCertificate.SubjectKeyID,
		Signer:       jackKeyInfo.Address,
	}
	_, _ = utils.ApproveAddX509RootCert(msgApproveAddX509RootCert,
		jackKeyInfo.Name, testconstants.Passphrase)

	// Certificate must be still in Proposed state. Request proposed Root certificate. 1 Approval received
	proposedCertificate, _ =
		utils.GetProposedX509RootCert(testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, msgProposeAddX509RootCert.Cert, proposedCertificate.PemCert)
	require.Equal(t, msgProposeAddX509RootCert.Signer, proposedCertificate.Owner)
	require.Equal(t, []sdk.AccAddress{jackKeyInfo.Address}, proposedCertificate.Approvals)

	// Request all proposed Root certificates. Still contains 1 certificate
	proposedCertificates, _ = utils.GetAllProposedX509RootCerts()
	require.Equal(t, 1, len(proposedCertificates.Items))
	require.Equal(t, testconstants.RootSubject, proposedCertificates.Items[0].Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, proposedCertificates.Items[0].SubjectKeyID)

	// Query all approved certificates
	certificates, _ = utils.GetAllX509Certs()
	require.Equal(t, 0, len(certificates.Items))

	// Request all active Root certificates
	certificates, _ = utils.GetAllX509RootCerts()
	require.Equal(t, 0, len(certificates.Items))

	// Alice (Trustee) approve Root certificate
	secondMsgApproveAddX509RootCert := pki.MsgApproveAddX509RootCert{
		Subject:      proposedCertificate.Subject,
		SubjectKeyID: proposedCertificate.SubjectKeyID,
		Signer:       aliceKeyInfo.Address,
	}
	_, _ = utils.ApproveAddX509RootCert(secondMsgApproveAddX509RootCert,
		aliceKeyInfo.Name, testconstants.Passphrase)

	// Certificate must be Approved. Request Root certificate
	certificate, _ := utils.GetX509Cert(testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, msgProposeAddX509RootCert.Cert, certificate.PemCert)
	require.Equal(t, msgProposeAddX509RootCert.Signer, certificate.Owner)
	require.True(t, certificate.IsRoot)

	// Request certificate chain for Root certificate
	certificateChain, _ := utils.GetX509CertChain(testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Len(t, certificateChain.Items, 1)
	require.Equal(t, msgProposeAddX509RootCert.Cert, certificateChain.Items[0].PemCert)
	require.Equal(t, msgProposeAddX509RootCert.Signer, certificateChain.Items[0].Owner)
	require.True(t, certificateChain.Items[0].IsRoot)

	// Request all proposed Root certificates must be empty
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

	// User (Not Trustee) add Intermediate certificate
	msgAddX509Cert := pki.MsgAddX509Cert{
		Cert:   testconstants.IntermediateCertPem,
		Signer: userKeyInfo.Address,
	}
	_, _ = utils.AddX509Cert(msgAddX509Cert, userKeyInfo.Name, testconstants.Passphrase)

	// Request all proposed Root certificates must be empty
	proposedCertificates, _ = utils.GetAllProposedX509RootCerts()
	require.Equal(t, 0, len(proposedCertificates.Items))

	// Request all approved certificates
	certificates, _ = utils.GetAllX509Certs()
	require.Equal(t, 2, len(certificates.Items))
	require.Equal(t, testconstants.RootSubject, certificates.Items[0].Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, certificates.Items[0].SubjectKeyID)
	require.Equal(t, testconstants.IntermediateSubject, certificates.Items[1].Subject)
	require.Equal(t, testconstants.IntermediateSubjectKeyID, certificates.Items[1].SubjectKeyID)

	// Request all approved Root certificates
	certificates, _ = utils.GetAllX509RootCerts()
	require.Equal(t, 1, len(certificates.Items))
	require.Equal(t, testconstants.RootSubjectKeyID, certificates.Items[0].SubjectKeyID)

	// Request Intermediate certificate
	certificate, _ = utils.GetX509Cert(testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Equal(t, msgAddX509Cert.Cert, certificate.PemCert)
	require.Equal(t, msgAddX509Cert.Signer, certificate.Owner)
	require.False(t, certificate.IsRoot)

	// Request certificate chain for Intermediate certificate
	certificateChain, _ = utils.GetX509CertChain(testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Len(t, certificateChain.Items, 2)
	require.Equal(t, msgAddX509Cert.Cert, certificateChain.Items[0].PemCert)
	require.Equal(t, msgAddX509Cert.Signer, certificateChain.Items[0].Owner)
	require.False(t, certificateChain.Items[0].IsRoot)
	require.Equal(t, msgProposeAddX509RootCert.Cert, certificateChain.Items[1].PemCert)
	require.Equal(t, msgProposeAddX509RootCert.Signer, certificateChain.Items[1].Owner)
	require.True(t, certificateChain.Items[1].IsRoot)

	// Alice (Trustee) add Leaf certificate
	secondMsgAddX509Cert := pki.MsgAddX509Cert{
		Cert:   testconstants.LeafCertPem,
		Signer: aliceKeyInfo.Address,
	}
	_, _ = utils.AddX509Cert(secondMsgAddX509Cert, aliceKeyInfo.Name, testconstants.Passphrase)

	// Request all proposed Root certificates must be empty
	proposedCertificates, _ = utils.GetAllProposedX509RootCerts()
	require.Equal(t, 0, len(proposedCertificates.Items))

	// Request all approved certificates
	certificates, _ = utils.GetAllX509Certs()
	require.Equal(t, 3, len(certificates.Items))
	require.Equal(t, testconstants.RootSubject, certificates.Items[0].Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, certificates.Items[0].SubjectKeyID)
	require.Equal(t, testconstants.IntermediateSubject, certificates.Items[1].Subject)
	require.Equal(t, testconstants.IntermediateSubjectKeyID, certificates.Items[1].SubjectKeyID)
	require.Equal(t, testconstants.LeafSubject, certificates.Items[2].Subject)
	require.Equal(t, testconstants.LeafSubjectKeyID, certificates.Items[2].SubjectKeyID)

	// Request all approved Root certificates
	certificates, _ = utils.GetAllX509RootCerts()
	require.Equal(t, 1, len(certificates.Items))
	require.Equal(t, testconstants.RootSubjectKeyID, certificates.Items[0].SubjectKeyID)

	// Request Leaf certificate
	certificate, _ = utils.GetX509Cert(testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Equal(t, secondMsgAddX509Cert.Cert, certificate.PemCert)
	require.Equal(t, secondMsgAddX509Cert.Signer, certificate.Owner)
	require.False(t, certificate.IsRoot)

	// Request certificate chain for Leaf certificate
	certificateChain, _ = utils.GetX509CertChain(testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Len(t, certificateChain.Items, 3)
	require.Equal(t, secondMsgAddX509Cert.Cert, certificateChain.Items[0].PemCert)
	require.Equal(t, secondMsgAddX509Cert.Signer, certificateChain.Items[0].Owner)
	require.False(t, certificateChain.Items[0].IsRoot)
	require.Equal(t, msgAddX509Cert.Cert, certificateChain.Items[1].PemCert)
	require.Equal(t, msgAddX509Cert.Signer, certificateChain.Items[1].Owner)
	require.False(t, certificateChain.Items[1].IsRoot)
	require.Equal(t, msgProposeAddX509RootCert.Cert, certificateChain.Items[2].PemCert)
	require.Equal(t, msgProposeAddX509RootCert.Signer, certificateChain.Items[2].Owner)
	require.True(t, certificateChain.Items[2].IsRoot)

	// Request all Subject certificates
	certificates, _ = utils.GetAllSubjectX509Certs(testconstants.LeafSubject)
	require.Equal(t, 1, len(certificates.Items))
	require.Equal(t, testconstants.LeafSubjectKeyID, certificates.Items[0].SubjectKeyID)
}
