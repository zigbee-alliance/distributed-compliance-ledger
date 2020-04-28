package rest

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/utils"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
	To Run test you need:
		* Run LocalNet with: `make install && make localnet_init && make localnet_start`
		* run RPC service with `zblcli rest-server --chain-id zblchain`

	TODO: provide tests for error cases
*/

const (
	jackAccount  = "jack"
	aliceAccount = "alice"
	bobAccount   = "bob"
	annaAccount  = "anna"
)

func /*Test*/ PkiDemo(t *testing.T) {
	// Get key info for Jack
	jackKeyInfo, _ := utils.GetKeyInfo(jackAccount)

	// Get key info for Alice
	aliceKeyInfo, _ := utils.GetKeyInfo(aliceAccount)

	// Get key info for Anna
	annaKeyInfo, _ := utils.GetKeyInfo(annaAccount)

	// Get key info for Bob
	bobKeyInfo, _ := utils.GetKeyInfo(bobAccount)

	// Assign Trustee role to Jack
	utils.AssignRole(jackKeyInfo.Address, jackKeyInfo, authz.Trustee)

	// Assign Trustee role to Alice
	utils.AssignRole(aliceKeyInfo.Address, jackKeyInfo, authz.Trustee)

	// Query all proposed certificates
	proposedCertificates, _ := utils.GetAllProposedX509RootCerts()
	require.Equal(t, 0, len(proposedCertificates.Items))

	// Query all approved certificates
	certificates, _ := utils.GetAllX509Certs()
	require.Equal(t, 0, len(certificates.Items))

	// Anna (Not Trustee) propose Root certificate
	msgProposeAddX509RootCert := pki.MsgProposeAddX509RootCert{
		Cert:   test_constants.RootCertPem,
		Signer: annaKeyInfo.Address,
	}
	_, _ = utils.ProposeAddX509RootCert(msgProposeAddX509RootCert, annaAccount, test_constants.Passphrase)

	//Request all proposed Root certificate
	proposedCertificates, _ = utils.GetAllProposedX509RootCerts()
	require.Equal(t, 1, len(proposedCertificates.Items))
	require.Equal(t, test_constants.RootSubject, proposedCertificates.Items[0].Subject)
	require.Equal(t, test_constants.RootSubjectKeyId, proposedCertificates.Items[0].SubjectKeyId)

	// Request proposed Root certificate
	proposedCertificate, _ := utils.GetProposedX509RootCert(test_constants.RootSubject, test_constants.RootSubjectKeyId)
	require.Equal(t, msgProposeAddX509RootCert.Cert, proposedCertificate.PemCert)
	require.Equal(t, msgProposeAddX509RootCert.Signer, proposedCertificate.Owner)
	require.Nil(t, proposedCertificate.Approvals)

	// Query all approved certificates
	certificates, _ = utils.GetAllX509Certs()
	require.Equal(t, 0, len(certificates.Items))

	// Request all active root certificates
	certificates, _ = utils.GetAllX509RootCerts()
	require.Equal(t, 0, len(certificates.Items))

	// Jack (Trustee) approve Root certificate
	msgApproveAddX509RootCert := pki.MsgApproveAddX509RootCert{
		Subject:      proposedCertificate.Subject,
		SubjectKeyId: proposedCertificate.SubjectKeyId,
		Signer:       jackKeyInfo.Address,
	}
	_, _ = utils.ApproveAddX509RootCert(msgApproveAddX509RootCert, jackAccount, test_constants.Passphrase)

	// Certificate mut be still in Proposed state. Request proposed Root certificate. 1 Approval received
	proposedCertificate, _ = utils.GetProposedX509RootCert(test_constants.RootSubject, test_constants.RootSubjectKeyId)
	require.Equal(t, msgProposeAddX509RootCert.Cert, proposedCertificate.PemCert)
	require.Equal(t, msgProposeAddX509RootCert.Signer, proposedCertificate.Owner)
	require.Equal(t, []sdk.AccAddress{jackKeyInfo.Address}, proposedCertificate.Approvals)

	// Request all proposed Root certificates. Still contains 1 certificate
	proposedCertificates, _ = utils.GetAllProposedX509RootCerts()
	require.Equal(t, 1, len(proposedCertificates.Items))
	require.Equal(t, test_constants.RootSubject, proposedCertificates.Items[0].Subject)
	require.Equal(t, test_constants.RootSubjectKeyId, proposedCertificates.Items[0].SubjectKeyId)

	// Query all approved certificates
	certificates, _ = utils.GetAllX509Certs()
	require.Equal(t, 0, len(certificates.Items))

	// Request all active root certificates
	certificates, _ = utils.GetAllX509RootCerts()
	require.Equal(t, 0, len(certificates.Items))

	// Alice (Trustee) approve Root certificate
	secondMsgApproveAddX509RootCert := pki.MsgApproveAddX509RootCert{
		Subject:      proposedCertificate.Subject,
		SubjectKeyId: proposedCertificate.SubjectKeyId,
		Signer:       aliceKeyInfo.Address,
	}
	_, _ = utils.ApproveAddX509RootCert(secondMsgApproveAddX509RootCert, aliceAccount, test_constants.Passphrase)

	// Certificate mut be Approved. Request Root certificate
	certificate, _ := utils.GetX509Cert(test_constants.RootSubject, test_constants.RootSubjectKeyId)
	require.Equal(t, msgProposeAddX509RootCert.Cert, certificate.PemCert)
	require.Equal(t, msgProposeAddX509RootCert.Signer, certificate.Owner)
	require.Equal(t, pki.RootCertificate, certificate.Type)

	// Request all proposed Root certificates must be empty
	proposedCertificates, _ = utils.GetAllProposedX509RootCerts()
	require.Equal(t, 0, len(proposedCertificates.Items))

	// Request all approved certificates
	certificates, _ = utils.GetAllX509Certs()
	require.Equal(t, 1, len(certificates.Items))
	require.Equal(t, test_constants.RootSubject, certificates.Items[0].Subject)
	require.Equal(t, test_constants.RootSubjectKeyId, certificates.Items[0].SubjectKeyId)

	// Request all approved Root certificates
	certificates, _ = utils.GetAllX509RootCerts()
	require.Equal(t, 1, len(certificates.Items))
	require.Equal(t, test_constants.RootSubject, certificates.Items[0].Subject)
	require.Equal(t, test_constants.RootSubjectKeyId, certificates.Items[0].SubjectKeyId)

	// Bob (Not Trustee) add intermediate certificate
	msgAddX509Cert := pki.MsgAddX509Cert{
		Cert:   test_constants.IntermediateCertPem,
		Signer: bobKeyInfo.Address,
	}
	_, _ = utils.AddX509Cert(msgAddX509Cert, bobAccount, test_constants.Passphrase)

	// Request all proposed Root certificates must be empty
	proposedCertificates, _ = utils.GetAllProposedX509RootCerts()
	require.Equal(t, 0, len(proposedCertificates.Items))

	// Request all approved certificates
	certificates, _ = utils.GetAllX509Certs()
	require.Equal(t, 2, len(certificates.Items))
	require.Equal(t, test_constants.RootSubject, certificates.Items[0].Subject)
	require.Equal(t, test_constants.RootSubjectKeyId, certificates.Items[0].SubjectKeyId)
	require.Equal(t, test_constants.IntermediateSubject, certificates.Items[1].Subject)
	require.Equal(t, test_constants.IntermediateSubjectKeyId, certificates.Items[1].SubjectKeyId)

	// Request all approved Root certificates
	certificates, _ = utils.GetAllX509RootCerts()
	require.Equal(t, 1, len(certificates.Items))
	require.Equal(t, test_constants.RootSubjectKeyId, certificates.Items[0].SubjectKeyId)

	// Request intermediate certificate
	certificate, _ = utils.GetX509Cert(test_constants.IntermediateSubject, test_constants.IntermediateSubjectKeyId)
	require.Equal(t, msgAddX509Cert.Cert, certificate.PemCert)
	require.Equal(t, msgAddX509Cert.Signer, certificate.Owner)
	require.Equal(t, pki.IntermediateCertificate, certificate.Type)

	// Alice (Trustee) add leaf certificate
	secondMsgAddX509Cert := pki.MsgAddX509Cert{
		Cert:   test_constants.LeafCertPem,
		Signer: aliceKeyInfo.Address,
	}
	_, _ = utils.AddX509Cert(secondMsgAddX509Cert, aliceAccount, test_constants.Passphrase)

	// Request all proposed Root certificates must be empty
	proposedCertificates, _ = utils.GetAllProposedX509RootCerts()
	require.Equal(t, 0, len(proposedCertificates.Items))

	// Request all approved certificates
	certificates, _ = utils.GetAllX509Certs()
	require.Equal(t, 3, len(certificates.Items))
	require.Equal(t, test_constants.RootSubject, certificates.Items[0].Subject)
	require.Equal(t, test_constants.RootSubjectKeyId, certificates.Items[0].SubjectKeyId)
	require.Equal(t, test_constants.IntermediateSubject, certificates.Items[1].Subject)
	require.Equal(t, test_constants.IntermediateSubjectKeyId, certificates.Items[1].SubjectKeyId)
	require.Equal(t, test_constants.LeafSubject, certificates.Items[2].Subject)
	require.Equal(t, test_constants.LeafSubjectKeyId, certificates.Items[2].SubjectKeyId)

	// Request all approved Root certificates
	certificates, _ = utils.GetAllX509RootCerts()
	require.Equal(t, 1, len(certificates.Items))
	require.Equal(t, test_constants.RootSubjectKeyId, certificates.Items[0].SubjectKeyId)

	// Request intermediate certificate
	certificate, _ = utils.GetX509Cert(test_constants.LeafSubject, test_constants.LeafSubjectKeyId)
	require.Equal(t, secondMsgAddX509Cert.Cert, certificate.PemCert)
	require.Equal(t, secondMsgAddX509Cert.Signer, certificate.Owner)
	require.Equal(t, pki.IntermediateCertificate, certificate.Type)

	// Request all Subject certificates
	certificates, _ = utils.GetAllSubjectX509Certs(test_constants.LeafSubject)
	require.Equal(t, 1, len(certificates.Items))
	require.Equal(t, test_constants.LeafSubjectKeyId, certificates.Items[0].SubjectKeyId)
}
