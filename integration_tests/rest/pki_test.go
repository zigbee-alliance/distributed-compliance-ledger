package rest_test

//nolint:goimports
import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/utils"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

//nolint:godox
/*
	To Run test you need:
		* Run LocalNet with: `make install && make localnet_init && make localnet_start`
		* run RPC service with `zblcli rest-server --chain-id zblchain`

	TODO: provide tests for error cases
*/

//nolint:funlen
func TestPkiDemo(t *testing.T) {
	// Get key info for Jack
	jackKeyInfo, _ := utils.GetKeyInfo(testconstants.JackAccount)

	// Get key info for Alice
	aliceKeyInfo, _ := utils.GetKeyInfo(testconstants.AliceAccount)

	// Get key info for Anna
	annaKeyInfo, _ := utils.GetKeyInfo(testconstants.AnnaAccount)

	// Get key info for Bob
	bobKeyInfo, _ := utils.GetKeyInfo(testconstants.BobAccount)

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
		Cert:   testconstants.RootCertPem,
		Signer: annaKeyInfo.Address,
	}
	_, _ = utils.ProposeAddX509RootCert(msgProposeAddX509RootCert,
		testconstants.AnnaAccount, testconstants.Passphrase)

	//Request all proposed Root certificate
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

	// Request all active root certificates
	certificates, _ = utils.GetAllX509RootCerts()
	require.Equal(t, 0, len(certificates.Items))

	// Jack (Trustee) approve Root certificate
	msgApproveAddX509RootCert := pki.MsgApproveAddX509RootCert{
		Subject:      proposedCertificate.Subject,
		SubjectKeyID: proposedCertificate.SubjectKeyID,
		Signer:       jackKeyInfo.Address,
	}
	_, _ = utils.ApproveAddX509RootCert(msgApproveAddX509RootCert,
		testconstants.JackAccount, testconstants.Passphrase)

	// Certificate mut be still in Proposed state. Request proposed Root certificate. 1 Approval received
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

	// Request all active root certificates
	certificates, _ = utils.GetAllX509RootCerts()
	require.Equal(t, 0, len(certificates.Items))

	// Alice (Trustee) approve Root certificate
	secondMsgApproveAddX509RootCert := pki.MsgApproveAddX509RootCert{
		Subject:      proposedCertificate.Subject,
		SubjectKeyID: proposedCertificate.SubjectKeyID,
		Signer:       aliceKeyInfo.Address,
	}
	_, _ = utils.ApproveAddX509RootCert(secondMsgApproveAddX509RootCert,
		testconstants.AliceAccount, testconstants.Passphrase)

	// Certificate mut be Approved. Request Root certificate
	certificate, _ := utils.GetX509Cert(testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, msgProposeAddX509RootCert.Cert, certificate.PemCert)
	require.Equal(t, msgProposeAddX509RootCert.Signer, certificate.Owner)
	require.Equal(t, pki.RootCertificate, certificate.Type)

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

	// Bob (Not Trustee) add intermediate certificate
	msgAddX509Cert := pki.MsgAddX509Cert{
		Cert:   testconstants.IntermediateCertPem,
		Signer: bobKeyInfo.Address,
	}
	_, _ = utils.AddX509Cert(msgAddX509Cert, testconstants.BobAccount, testconstants.Passphrase)

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

	// Request intermediate certificate
	certificate, _ = utils.GetX509Cert(testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Equal(t, msgAddX509Cert.Cert, certificate.PemCert)
	require.Equal(t, msgAddX509Cert.Signer, certificate.Owner)
	require.Equal(t, pki.IntermediateCertificate, certificate.Type)

	// Alice (Trustee) add leaf certificate
	secondMsgAddX509Cert := pki.MsgAddX509Cert{
		Cert:   testconstants.LeafCertPem,
		Signer: aliceKeyInfo.Address,
	}
	_, _ = utils.AddX509Cert(secondMsgAddX509Cert, testconstants.AliceAccount, testconstants.Passphrase)

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

	// Request intermediate certificate
	certificate, _ = utils.GetX509Cert(testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Equal(t, secondMsgAddX509Cert.Cert, certificate.PemCert)
	require.Equal(t, secondMsgAddX509Cert.Signer, certificate.Owner)
	require.Equal(t, pki.IntermediateCertificate, certificate.Type)

	// Request all Subject certificates
	certificates, _ = utils.GetAllSubjectX509Certs(testconstants.LeafSubject)
	require.Equal(t, 1, len(certificates.Items))
	require.Equal(t, testconstants.LeafSubjectKeyID, certificates.Items[0].SubjectKeyID)
}
