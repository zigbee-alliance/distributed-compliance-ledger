package pki

import (
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// ProposeAddX509RootCert proposes adding an x509 root certificate.
func ProposeAddX509RootCert(certFile, from string, extra ...string) (*utils.TxResult, error) {
	args := []string{
		"tx", "pki", "propose-add-x509-root-cert",
		"--certificate", certFile,
		"--from", from,
	}
	args = append(args, extra...)
	return utils.ExecuteTx(args...)
}

// ApproveAddX509RootCert approves adding an x509 root certificate.
func ApproveAddX509RootCert(subject, subjectKeyID, from string, extra ...string) (*utils.TxResult, error) {
	args := []string{
		"tx", "pki", "approve-add-x509-root-cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"--from", from,
	}
	args = append(args, extra...)
	return utils.ExecuteTx(args...)
}

// RejectAddX509RootCert rejects adding an x509 root certificate.
func RejectAddX509RootCert(subject, subjectKeyID, from string, extra ...string) (*utils.TxResult, error) {
	args := []string{
		"tx", "pki", "reject-add-x509-root-cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"--from", from,
	}
	args = append(args, extra...)
	return utils.ExecuteTx(args...)
}

// AddX509Cert adds a non-root x509 certificate.
func AddX509Cert(certFile, from string, extra ...string) (*utils.TxResult, error) {
	args := []string{
		"tx", "pki", "add-x509-cert",
		"--certificate", certFile,
		"--from", from,
	}
	args = append(args, extra...)
	return utils.ExecuteTx(args...)
}

// ProposeRevokeX509RootCert proposes revoking an x509 root certificate.
func ProposeRevokeX509RootCert(subject, subjectKeyID, from string, extra ...string) (*utils.TxResult, error) {
	args := []string{
		"tx", "pki", "propose-revoke-x509-root-cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"--from", from,
	}
	args = append(args, extra...)
	return utils.ExecuteTx(args...)
}

// ApproveRevokeX509RootCert approves revoking an x509 root certificate.
func ApproveRevokeX509RootCert(subject, subjectKeyID, from string, extra ...string) (*utils.TxResult, error) {
	args := []string{
		"tx", "pki", "approve-revoke-x509-root-cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"--from", from,
	}
	args = append(args, extra...)
	return utils.ExecuteTx(args...)
}

// RevokeX509Cert revokes an x509 certificate.
func RevokeX509Cert(subject, subjectKeyID, from string, extra ...string) (*utils.TxResult, error) {
	args := []string{
		"tx", "pki", "revoke-x509-cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"--from", from,
	}
	args = append(args, extra...)
	return utils.ExecuteTx(args...)
}

// AddNocRootCert adds a NOC root certificate.
func AddNocRootCert(certFile, from string, extra ...string) (*utils.TxResult, error) {
	args := []string{
		"tx", "pki", "add-noc-x509-root-cert",
		"--certificate", certFile,
		"--from", from,
	}
	args = append(args, extra...)
	return utils.ExecuteTx(args...)
}

// AddNocX509IcaCert adds a NOC ICA certificate.
func AddNocX509IcaCert(certFile, from string, extra ...string) (*utils.TxResult, error) {
	args := []string{
		"tx", "pki", "add-noc-x509-ica-cert",
		"--certificate", certFile,
		"--from", from,
	}
	args = append(args, extra...)
	return utils.ExecuteTx(args...)
}

// RevokeNocRootCert revokes a NOC root certificate.
func RevokeNocRootCert(subject, subjectKeyID, from string, extra ...string) (*utils.TxResult, error) {
	args := []string{
		"tx", "pki", "revoke-noc-x509-root-cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"--from", from,
	}
	args = append(args, extra...)
	return utils.ExecuteTx(args...)
}

// RevokeNocX509IcaCert revokes a NOC ICA certificate.
func RevokeNocX509IcaCert(subject, subjectKeyID, from string, extra ...string) (*utils.TxResult, error) {
	args := []string{
		"tx", "pki", "revoke-noc-x509-ica-cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"--from", from,
	}
	args = append(args, extra...)
	return utils.ExecuteTx(args...)
}

// RemoveX509Cert removes an x509 certificate.
func RemoveX509Cert(subject, subjectKeyID, from string, extra ...string) (*utils.TxResult, error) {
	args := []string{
		"tx", "pki", "remove-x509-cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"--from", from,
	}
	args = append(args, extra...)
	return utils.ExecuteTx(args...)
}

// RemoveNocCert removes a NOC ICA certificate.
func RemoveNocCert(subject, subjectKeyID, from string, extra ...string) (*utils.TxResult, error) {
	args := []string{
		"tx", "pki", "remove-noc-x509-ica-cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"--from", from,
	}
	args = append(args, extra...)
	return utils.ExecuteTx(args...)
}

// RemoveNocRootCert removes a NOC root certificate.
func RemoveNocRootCert(subject, subjectKeyID, from string, extra ...string) (*utils.TxResult, error) {
	args := []string{
		"tx", "pki", "remove-noc-x509-root-cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"--from", from,
	}
	args = append(args, extra...)
	return utils.ExecuteTx(args...)
}

// AddRevocationPoint adds a PKI revocation distribution point.
func AddRevocationPoint(from string, extra ...string) (*utils.TxResult, error) {
	args := []string{"tx", "pki", "add-revocation-point", "--from", from}
	args = append(args, extra...)
	return utils.ExecuteTx(args...)
}

// UpdateRevocationPoint updates a PKI revocation distribution point.
func UpdateRevocationPoint(from string, extra ...string) (*utils.TxResult, error) {
	args := []string{"tx", "pki", "update-revocation-point", "--from", from}
	args = append(args, extra...)
	return utils.ExecuteTx(args...)
}

// DeleteRevocationPoint deletes a PKI revocation distribution point.
func DeleteRevocationPoint(from string, extra ...string) (*utils.TxResult, error) {
	args := []string{"tx", "pki", "delete-revocation-point", "--from", from}
	args = append(args, extra...)
	return utils.ExecuteTx(args...)
}

// AssignVid assigns a VID to a root certificate.
func AssignVid(subject, subjectKeyID string, vid int, from string) (*utils.TxResult, error) {
	return utils.ExecuteTx("tx", "pki", "assign-vid",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"--vid", itoa(vid),
		"--from", from,
	)
}

// AddPkiRevocationDistributionPoint adds a PKI revocation distribution point.
func AddPkiRevocationDistributionPoint(from string, extra ...string) (*utils.TxResult, error) {
	args := []string{"tx", "pki", "add-pki-revocation-distribution-point", "--from", from}
	args = append(args, extra...)
	return utils.ExecuteTx(args...)
}

// UpdatePkiRevocationDistributionPoint updates a PKI revocation distribution point.
func UpdatePkiRevocationDistributionPoint(from string, extra ...string) (*utils.TxResult, error) {
	args := []string{"tx", "pki", "update-pki-revocation-distribution-point", "--from", from}
	args = append(args, extra...)
	return utils.ExecuteTx(args...)
}

// DeletePkiRevocationDistributionPoint deletes a PKI revocation distribution point.
func DeletePkiRevocationDistributionPoint(from string, extra ...string) (*utils.TxResult, error) {
	args := []string{"tx", "pki", "delete-pki-revocation-distribution-point", "--from", from}
	args = append(args, extra...)
	return utils.ExecuteTx(args...)
}

// QueryX509Cert queries an approved x509 certificate.
func QueryX509Cert(subject, subjectKeyID string) ([]byte, error) {
	return utils.ExecuteCLI("query", "pki", "x509-cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"-o", "json",
	)
}

// QueryCert queries any certificate type (DA or NOC) by subject and subjectKeyID.
func QueryCert(subject, subjectKeyID string) ([]byte, error) {
	return utils.ExecuteCLI("query", "pki", "cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"-o", "json",
	)
}

// QueryAllX509Certs queries all approved x509 certificates.
func QueryAllX509Certs(extra ...string) ([]byte, error) {
	args := []string{"query", "pki", "all-x509-certs", "-o", "json"}
	args = append(args, extra...)
	return utils.ExecuteCLI(args...)
}

// QueryProposedX509RootCert queries a proposed root certificate.
func QueryProposedX509RootCert(subject, subjectKeyID string) ([]byte, error) {
	return utils.ExecuteCLI("query", "pki", "proposed-x509-root-cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"-o", "json",
	)
}

// QueryAllProposedX509RootCerts queries all proposed root certificates.
func QueryAllProposedX509RootCerts() ([]byte, error) {
	return utils.ExecuteCLI("query", "pki", "all-proposed-x509-root-certs", "-o", "json")
}

// QueryRevokedX509Cert queries a revoked x509 certificate.
func QueryRevokedX509Cert(subject, subjectKeyID string) ([]byte, error) {
	return utils.ExecuteCLI("query", "pki", "revoked-x509-cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"-o", "json",
	)
}

// QueryAllRevokedX509Certs queries all revoked x509 certificates.
func QueryAllRevokedX509Certs() ([]byte, error) {
	return utils.ExecuteCLI("query", "pki", "all-revoked-x509-certs", "-o", "json")
}

// QueryProposedRevokedX509RootCert queries a proposed revocation for a root cert.
func QueryProposedRevokedX509RootCert(subject, subjectKeyID string) ([]byte, error) {
	return utils.ExecuteCLI("query", "pki", "proposed-x509-root-cert-to-revoke",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"-o", "json",
	)
}

// QueryAllProposedRevokedX509RootCerts queries all proposed revocations for root certs.
func QueryAllProposedRevokedX509RootCerts() ([]byte, error) {
	return utils.ExecuteCLI("query", "pki", "all-proposed-x509-root-certs-to-revoke", "-o", "json")
}

// QueryRejectedX509RootCert queries a rejected root certificate.
func QueryRejectedX509RootCert(subject, subjectKeyID string) ([]byte, error) {
	return utils.ExecuteCLI("query", "pki", "rejected-x509-root-cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"-o", "json",
	)
}

// QueryAllRejectedX509RootCerts queries all rejected root certificates.
func QueryAllRejectedX509RootCerts() ([]byte, error) {
	return utils.ExecuteCLI("query", "pki", "all-rejected-x509-root-certs", "-o", "json")
}

// QueryNocRootCerts queries NOC root certificates for a vid.
func QueryNocRootCerts(vid int) ([]byte, error) {
	return utils.ExecuteCLI("query", "pki", "noc-x509-root-certs",
		"--vid", itoa(vid),
		"-o", "json",
	)
}

// QueryAllNocRootCerts queries all NOC root certificates.
func QueryAllNocRootCerts() ([]byte, error) {
	return utils.ExecuteCLI("query", "pki", "all-noc-x509-root-certs", "-o", "json")
}

// QueryNocX509IcaCerts queries NOC ICA certificates for a vid.
func QueryNocX509IcaCerts(vid int) ([]byte, error) {
	return utils.ExecuteCLI("query", "pki", "noc-x509-ica-certs",
		"--vid", itoa(vid),
		"-o", "json",
	)
}

// QueryAllNocX509IcaCerts queries all NOC ICA certificates.
func QueryAllNocX509IcaCerts() ([]byte, error) {
	return utils.ExecuteCLI("query", "pki", "all-noc-x509-ica-certs", "-o", "json")
}

// QueryAllNocX509Certs queries all NOC certificates (root + ICA).
func QueryAllNocX509Certs() ([]byte, error) {
	return utils.ExecuteCLI("query", "pki", "all-noc-x509-certs", "-o", "json")
}

// QueryNocCert queries a NOC certificate by various flags (pass as extra: --subject, --subject-key-id, --vid, etc.)
func QueryNocCert(extra ...string) ([]byte, error) {
	args := []string{"query", "pki", "noc-x509-cert", "-o", "json"}
	args = append(args, extra...)
	return utils.ExecuteCLI(args...)
}

// QueryNocSubjectCerts queries all NOC certificates by subject.
func QueryNocSubjectCerts(subject string) ([]byte, error) {
	return utils.ExecuteCLI("query", "pki", "all-noc-subject-x509-certs",
		"--subject", subject,
		"-o", "json",
	)
}

// QueryAllRevokedNocRootCerts queries all revoked NOC root certificates.
func QueryAllRevokedNocRootCerts() ([]byte, error) {
	return utils.ExecuteCLI("query", "pki", "all-revoked-noc-x509-root-certs", "-o", "json")
}

// QueryRevokedNocRootCert queries a revoked NOC root certificate by subject and subjectKeyID.
func QueryRevokedNocRootCert(subject, subjectKeyID string) ([]byte, error) {
	return utils.ExecuteCLI("query", "pki", "revoked-noc-x509-root-cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"-o", "json",
	)
}

// QueryAllRevokedNocX509IcaCerts queries all revoked NOC ICA certificates.
func QueryAllRevokedNocX509IcaCerts() ([]byte, error) {
	return utils.ExecuteCLI("query", "pki", "all-revoked-noc-x509-ica-certs", "-o", "json")
}

// QueryAllX509RootCerts queries all approved root certificates.
func QueryAllX509RootCerts() ([]byte, error) {
	return utils.ExecuteCLI("query", "pki", "all-x509-root-certs", "-o", "json")
}

// QueryAllRevokedX509RootCerts queries all revoked root certificates.
func QueryAllRevokedX509RootCerts() ([]byte, error) {
	return utils.ExecuteCLI("query", "pki", "all-revoked-x509-root-certs", "-o", "json")
}

// QueryX509CertBySKID queries an x509 certificate by subject-key-id only (no subject required).
func QueryX509CertBySKID(skid string) ([]byte, error) {
	return utils.ExecuteCLI("query", "pki", "x509-cert",
		"--subject-key-id", skid,
		"-o", "json",
	)
}

// QueryChildX509Certs queries all child certificates by subject and subjectKeyID.
func QueryChildX509Certs(subject, subjectKeyID string) ([]byte, error) {
	return utils.ExecuteCLI("query", "pki", "all-child-x509-certs",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"-o", "json",
	)
}

// QueryPkiRevocationDistributionPoint queries a revocation distribution point.
func QueryPkiRevocationDistributionPoint(vid int, label, issuerSubjectKeyID string) ([]byte, error) {
	return utils.ExecuteCLI("query", "pki", "revocation-point",
		"--vid", itoa(vid),
		"--label", label,
		"--issuer-subject-key-id", issuerSubjectKeyID,
		"-o", "json",
	)
}

// QueryAllPkiRevocationDistributionPoints queries all revocation distribution points.
func QueryAllPkiRevocationDistributionPoints() ([]byte, error) {
	return utils.ExecuteCLI("query", "pki", "all-revocation-points", "-o", "json")
}

// QueryPkiRevocationDistributionPointsByIssuer queries revocation distribution points by issuer subject key ID.
func QueryPkiRevocationDistributionPointsByIssuer(issuerSubjectKeyID string) ([]byte, error) {
	return utils.ExecuteCLI("query", "pki", "revocation-points",
		"--issuer-subject-key-id", issuerSubjectKeyID,
		"-o", "json",
	)
}

// QueryX509CertsBySubject queries all certs by subject.
func QueryX509CertsBySubject(subject string) ([]byte, error) {
	return utils.ExecuteCLI("query", "pki", "all-subject-x509-certs",
		"--subject", subject,
		"-o", "json",
	)
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	neg := false
	if n < 0 {
		neg = true
		n = -n
	}
	var buf [20]byte
	pos := len(buf)
	for n > 0 {
		pos--
		buf[pos] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		pos--
		buf[pos] = '-'
	}
	return string(buf[pos:])
}
