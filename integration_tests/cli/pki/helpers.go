package pki

import (
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// X509ProposeOpts holds optional flags for propose-add-x509-root-cert /
// add-x509-cert. VID > 0 emits --vid; SchemaVersion / Info / Time are sent
// only when non-empty.
type X509ProposeOpts struct {
	VID           int
	VIDHex        string
	SchemaVersion string
	Info          string
	Time          string
	Extra         []string
}

func (o X509ProposeOpts) args() []string {
	var args []string
	if o.VID != 0 || o.VIDHex != "" {
		args = append(args, "--vid", flagOrHex(o.VID, o.VIDHex))
	}
	if o.SchemaVersion != "" {
		args = append(args, "--schemaVersion", o.SchemaVersion)
	}
	if o.Info != "" {
		args = append(args, "--info", o.Info)
	}
	if o.Time != "" {
		args = append(args, "--time", o.Time)
	}

	return append(args, o.Extra...)
}

// X509ActionOpts holds optional flags for approve / reject /
// propose-revoke / approve-revoke X509 root cert. Reason emits --reason;
// Info emits --info; RevokeChild emits --revoke-child=true.
type X509ActionOpts struct {
	Reason       string
	Info         string
	RevokeChild  bool
	SerialNumber string
	Extra        []string
}

func (o X509ActionOpts) args() []string {
	var args []string
	if o.Info != "" {
		args = append(args, "--info", o.Info)
	}
	if o.Reason != "" {
		args = append(args, "--reason", o.Reason)
	}
	if o.SerialNumber != "" {
		args = append(args, "--serial-number", o.SerialNumber)
	}
	if o.RevokeChild {
		args = append(args, "--revoke-child=true")
	}

	return append(args, o.Extra...)
}

// ProposeAddX509RootCert proposes adding an x509 root certificate.
func ProposeAddX509RootCert(certFile, from string, opts ...X509ProposeOpts) (*utils.TxResult, error) {
	args := []string{
		"tx", "pki", "propose-add-x509-root-cert",
		"--certificate", certFile,
		"--from", from,
	}
	for _, o := range opts {
		args = append(args, o.args()...)
	}

	return utils.ExecuteTx(args...)
}

// ApproveAddX509RootCert approves adding an x509 root certificate.
func ApproveAddX509RootCert(subject, subjectKeyID, from string, opts ...X509ActionOpts) (*utils.TxResult, error) {
	args := []string{
		"tx", "pki", "approve-add-x509-root-cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"--from", from,
	}
	for _, o := range opts {
		args = append(args, o.args()...)
	}

	return utils.ExecuteTx(args...)
}

// RejectAddX509RootCert rejects adding an x509 root certificate.
func RejectAddX509RootCert(subject, subjectKeyID, from string, opts ...X509ActionOpts) (*utils.TxResult, error) {
	args := []string{
		"tx", "pki", "reject-add-x509-root-cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"--from", from,
	}
	for _, o := range opts {
		args = append(args, o.args()...)
	}

	return utils.ExecuteTx(args...)
}

// AddX509Cert adds a non-root x509 certificate.
func AddX509Cert(certFile, from string, opts ...X509ProposeOpts) (*utils.TxResult, error) {
	args := []string{
		"tx", "pki", "add-x509-cert",
		"--certificate", certFile,
		"--from", from,
	}
	for _, o := range opts {
		args = append(args, o.args()...)
	}

	return utils.ExecuteTx(args...)
}

// ProposeRevokeX509RootCert proposes revoking an x509 root certificate.
func ProposeRevokeX509RootCert(subject, subjectKeyID, from string, opts ...X509ActionOpts) (*utils.TxResult, error) {
	args := []string{
		"tx", "pki", "propose-revoke-x509-root-cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"--from", from,
	}
	for _, o := range opts {
		args = append(args, o.args()...)
	}

	return utils.ExecuteTx(args...)
}

// ApproveRevokeX509RootCert approves revoking an x509 root certificate.
func ApproveRevokeX509RootCert(subject, subjectKeyID, from string, opts ...X509ActionOpts) (*utils.TxResult, error) {
	args := []string{
		"tx", "pki", "approve-revoke-x509-root-cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"--from", from,
	}
	for _, o := range opts {
		args = append(args, o.args()...)
	}

	return utils.ExecuteTx(args...)
}

// RevokeX509Cert revokes an x509 certificate.
func RevokeX509Cert(subject, subjectKeyID, from string, opts ...RevokeNocCertOpts) (*utils.TxResult, error) {
	args := []string{
		"tx", "pki", "revoke-x509-cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"--from", from,
	}
	for _, o := range opts {
		args = append(args, o.args()...)
	}

	return utils.ExecuteTx(args...)
}

// AddNocCertOpts is shared by AddNocRootCert / AddNocX509IcaCert.
// IsVidVerificationSigner emits --is-vid-verification-signer=true (Matter §6.5.12
// VVSC flow). SchemaVersion is sent only when non-empty.
type AddNocCertOpts struct {
	IsVidVerificationSigner bool
	SchemaVersion           string
	Extra                   []string
}

func (o AddNocCertOpts) args() []string {
	var extra []string
	if o.IsVidVerificationSigner {
		extra = append(extra, "--is-vid-verification-signer=true")
	}
	if o.SchemaVersion != "" {
		extra = append(extra, "--schemaVersion", o.SchemaVersion)
	}

	return append(extra, o.Extra...)
}

// AddNocRootCert adds a NOC root certificate.
func AddNocRootCert(certFile, from string, opts ...AddNocCertOpts) (*utils.TxResult, error) {
	args := []string{
		"tx", "pki", "add-noc-x509-root-cert",
		"--certificate", certFile,
		"--from", from,
	}
	for _, o := range opts {
		args = append(args, o.args()...)
	}

	return utils.ExecuteTx(args...)
}

// AddNocX509IcaCert adds a NOC ICA certificate.
func AddNocX509IcaCert(certFile, from string, opts ...AddNocCertOpts) (*utils.TxResult, error) {
	args := []string{
		"tx", "pki", "add-noc-x509-ica-cert",
		"--certificate", certFile,
		"--from", from,
	}
	for _, o := range opts {
		args = append(args, o.args()...)
	}

	return utils.ExecuteTx(args...)
}

// RevokeNocCertOpts is shared by RevokeNocRootCert / RevokeNocX509IcaCert.
// SerialNumber narrows the revocation to a single serial (default: all).
// RevokeChild cascades revocation to child certificates.
type RevokeNocCertOpts struct {
	SerialNumber string
	RevokeChild  bool
	Extra        []string
}

func (o RevokeNocCertOpts) args() []string {
	var extra []string
	if o.SerialNumber != "" {
		extra = append(extra, "--serial-number", o.SerialNumber)
	}
	if o.RevokeChild {
		extra = append(extra, "--revoke-child=true")
	}

	return append(extra, o.Extra...)
}

// RevokeNocRootCert revokes a NOC root certificate.
func RevokeNocRootCert(subject, subjectKeyID, from string, opts ...RevokeNocCertOpts) (*utils.TxResult, error) {
	args := []string{
		"tx", "pki", "revoke-noc-x509-root-cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"--from", from,
	}
	for _, o := range opts {
		args = append(args, o.args()...)
	}

	return utils.ExecuteTx(args...)
}

// RevokeNocX509IcaCert revokes a NOC ICA certificate.
func RevokeNocX509IcaCert(subject, subjectKeyID, from string, opts ...RevokeNocCertOpts) (*utils.TxResult, error) {
	args := []string{
		"tx", "pki", "revoke-noc-x509-ica-cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"--from", from,
	}
	for _, o := range opts {
		args = append(args, o.args()...)
	}

	return utils.ExecuteTx(args...)
}

// RemoveX509Cert removes an x509 certificate.
func RemoveX509Cert(subject, subjectKeyID, from string, opts ...RevokeNocCertOpts) (*utils.TxResult, error) {
	args := []string{
		"tx", "pki", "remove-x509-cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"--from", from,
	}
	for _, o := range opts {
		args = append(args, o.args()...)
	}

	return utils.ExecuteTx(args...)
}

// RemoveNocCert removes a NOC ICA certificate.
func RemoveNocCert(subject, subjectKeyID, from string, opts ...RevokeNocCertOpts) (*utils.TxResult, error) {
	args := []string{
		"tx", "pki", "remove-noc-x509-ica-cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"--from", from,
	}
	for _, o := range opts {
		args = append(args, o.args()...)
	}

	return utils.ExecuteTx(args...)
}

// RemoveNocRootCert removes a NOC root certificate.
func RemoveNocRootCert(subject, subjectKeyID, from string, opts ...RevokeNocCertOpts) (*utils.TxResult, error) {
	args := []string{
		"tx", "pki", "remove-noc-x509-root-cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"--from", from,
	}
	for _, o := range opts {
		args = append(args, o.args()...)
	}

	return utils.ExecuteTx(args...)
}

// RevocationPointOpts holds parameters for add/update/delete-revocation-point.
// Each field is sent only when non-zero/non-empty; the boolean IsPAA emits
// --is-paa=true when set. Use Extra for unmodeled flags.
type RevocationPointOpts struct {
	VID                  int
	VIDHex               string
	PID                  int
	PIDHex               string
	IsPAA                bool
	Certificate          string
	CertificateDelegator string
	Label                string
	DataURL              string
	IssuerSubjectKeyID   string
	RevocationType       string
	CRLSignerCertificate string
	SchemaVersion        string
	Extra                []string
}

func (o RevocationPointOpts) args() []string {
	var args []string
	if o.VID != 0 || o.VIDHex != "" {
		args = append(args, "--vid", flagOrHex(o.VID, o.VIDHex))
	}
	if o.PID != 0 || o.PIDHex != "" {
		args = append(args, "--pid", flagOrHex(o.PID, o.PIDHex))
	}
	if o.IsPAA {
		args = append(args, "--is-paa=true")
	}
	if o.Certificate != "" {
		args = append(args, "--certificate", o.Certificate)
	}
	if o.CertificateDelegator != "" {
		args = append(args, "--certificate-delegator", o.CertificateDelegator)
	}
	if o.Label != "" {
		args = append(args, "--label", o.Label)
	}
	if o.DataURL != "" {
		args = append(args, "--data-url", o.DataURL)
	}
	if o.IssuerSubjectKeyID != "" {
		args = append(args, "--issuer-subject-key-id", o.IssuerSubjectKeyID)
	}
	if o.RevocationType != "" {
		args = append(args, "--revocation-type", o.RevocationType)
	}
	if o.CRLSignerCertificate != "" {
		args = append(args, "--crl-signer-certificate", o.CRLSignerCertificate)
	}
	if o.SchemaVersion != "" {
		args = append(args, "--schemaVersion", o.SchemaVersion)
	}

	return append(args, o.Extra...)
}

// AddRevocationPoint adds a PKI revocation distribution point.
func AddRevocationPoint(from string, opts RevocationPointOpts) (*utils.TxResult, error) {
	args := []string{"tx", "pki", "add-revocation-point", "--from", from}
	args = append(args, opts.args()...)

	return utils.ExecuteTx(args...)
}

// UpdateRevocationPoint updates a PKI revocation distribution point.
func UpdateRevocationPoint(from string, opts RevocationPointOpts) (*utils.TxResult, error) {
	args := []string{"tx", "pki", "update-revocation-point", "--from", from}
	args = append(args, opts.args()...)

	return utils.ExecuteTx(args...)
}

// DeleteRevocationPoint deletes a PKI revocation distribution point.
func DeleteRevocationPoint(from string, opts RevocationPointOpts) (*utils.TxResult, error) {
	args := []string{"tx", "pki", "delete-revocation-point", "--from", from}
	args = append(args, opts.args()...)

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

// flagOrHex returns hex if non-empty, otherwise the decimal-formatted n.
func flagOrHex(n int, hex string) string {
	if hex != "" {
		return hex
	}

	return itoa(n)
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
