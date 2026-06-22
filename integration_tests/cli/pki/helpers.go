package pki

import (
	"encoding/json"
	"fmt"

	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
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

// getSingle runs a single-item dcld query and unmarshals the JSON output into
// v. Returns (false, nil) when the CLI returned "Not Found" so callers can
// branch on a nil result; otherwise (true, nil) and v is populated.
func getSingle(v interface{}, args ...string) (found bool, err error) {
	out, err := utils.ExecuteCLI(args...)
	if err != nil {
		return false, err
	}
	if utils.IsNotFound(out) {
		return false, nil
	}
	out = utils.NormalizeProtoJSON(out)
	if err := json.Unmarshal(out, v); err != nil {
		return false, fmt.Errorf("parse %T: %w, output: %s", v, err, string(out))
	}

	return true, nil
}

// getList runs an all-* dcld query and unmarshals the wrapper response into v.
func getList(v interface{}, args ...string) error {
	out, err := utils.ExecuteCLI(args...)
	if err != nil {
		return err
	}
	out = utils.NormalizeProtoJSON(utils.StripPagination(out))
	if err := json.Unmarshal(out, v); err != nil {
		return fmt.Errorf("parse %T: %w, output: %s", v, err, string(out))
	}

	return nil
}

// GetX509Cert queries an approved x509 certificate by (subject, SKID). Returns
// nil when no record matches.
func GetX509Cert(subject, subjectKeyID string) (*pkitypes.ApprovedCertificates, error) {
	var res pkitypes.ApprovedCertificates
	found, err := getSingle(&res,
		"query", "pki", "x509-cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetX509CertBySKID queries an x509 certificate by SKID only. Returns nil when
// no record matches.
func GetX509CertBySKID(skid string) (*pkitypes.ApprovedCertificatesBySubjectKeyId, error) {
	var res pkitypes.ApprovedCertificatesBySubjectKeyId
	found, err := getSingle(&res,
		"query", "pki", "x509-cert",
		"--subject-key-id", skid,
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetCert queries any certificate (DA or NOC) by (subject, SKID).
func GetCert(subject, subjectKeyID string) (*pkitypes.AllCertificates, error) {
	var res pkitypes.AllCertificates
	found, err := getSingle(&res,
		"query", "pki", "cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetAllX509Certs queries all approved x509 certificates.
func GetAllX509Certs(extra ...string) ([]pkitypes.ApprovedCertificates, error) {
	var res pkitypes.QueryAllApprovedCertificatesResponse
	args := []string{"query", "pki", "all-x509-certs", "-o", "json"}
	args = append(args, extra...)
	if err := getList(&res, args...); err != nil {
		return nil, err
	}

	return res.ApprovedCertificates, nil
}

// GetAllX509RootCerts queries all approved root certificates.
func GetAllX509RootCerts() (*pkitypes.ApprovedRootCertificates, error) {
	var res pkitypes.ApprovedRootCertificates
	found, err := getSingle(&res, "query", "pki", "all-x509-root-certs", "-o", "json")
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetX509CertsBySubject queries all approved certs by subject. Returns nil
// when no record matches.
func GetX509CertsBySubject(subject string) (*pkitypes.ApprovedCertificatesBySubject, error) {
	var res pkitypes.ApprovedCertificatesBySubject
	found, err := getSingle(&res,
		"query", "pki", "all-subject-x509-certs",
		"--subject", subject,
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetProposedX509RootCert queries a proposed root certificate. Returns nil
// when no record matches.
func GetProposedX509RootCert(subject, subjectKeyID string) (*pkitypes.ProposedCertificate, error) {
	var res pkitypes.ProposedCertificate
	found, err := getSingle(&res,
		"query", "pki", "proposed-x509-root-cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetAllProposedX509RootCerts queries all proposed root certificates.
func GetAllProposedX509RootCerts() ([]pkitypes.ProposedCertificate, error) {
	var res pkitypes.QueryAllProposedCertificateResponse
	if err := getList(&res, "query", "pki", "all-proposed-x509-root-certs", "-o", "json"); err != nil {
		return nil, err
	}

	return res.ProposedCertificate, nil
}

// GetRevokedX509Cert queries a revoked x509 certificate. Returns nil when no
// record matches.
func GetRevokedX509Cert(subject, subjectKeyID string) (*pkitypes.RevokedCertificates, error) {
	var res pkitypes.RevokedCertificates
	found, err := getSingle(&res,
		"query", "pki", "revoked-x509-cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetAllRevokedX509Certs queries all revoked x509 certificates.
func GetAllRevokedX509Certs() ([]pkitypes.RevokedCertificates, error) {
	var res pkitypes.QueryAllRevokedCertificatesResponse
	if err := getList(&res, "query", "pki", "all-revoked-x509-certs", "-o", "json"); err != nil {
		return nil, err
	}

	return res.RevokedCertificates, nil
}

// GetAllRevokedX509RootCerts queries all revoked root certificates.
func GetAllRevokedX509RootCerts() (*pkitypes.RevokedRootCertificates, error) {
	var res pkitypes.RevokedRootCertificates
	found, err := getSingle(&res, "query", "pki", "all-revoked-x509-root-certs", "-o", "json")
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetProposedRevokedX509RootCert queries a proposed revocation for a root
// cert. Returns nil when no record matches.
func GetProposedRevokedX509RootCert(subject, subjectKeyID string) (*pkitypes.ProposedCertificateRevocation, error) {
	var res pkitypes.ProposedCertificateRevocation
	found, err := getSingle(&res,
		"query", "pki", "proposed-x509-root-cert-to-revoke",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetAllProposedRevokedX509RootCerts queries all proposed revocations.
func GetAllProposedRevokedX509RootCerts() ([]pkitypes.ProposedCertificateRevocation, error) {
	var res pkitypes.QueryAllProposedCertificateRevocationResponse
	if err := getList(&res, "query", "pki", "all-proposed-x509-root-certs-to-revoke", "-o", "json"); err != nil {
		return nil, err
	}

	return res.ProposedCertificateRevocation, nil
}

// GetRejectedX509RootCert queries a rejected root certificate. Returns nil
// when no record matches.
func GetRejectedX509RootCert(subject, subjectKeyID string) (*pkitypes.RejectedCertificate, error) {
	var res pkitypes.RejectedCertificate
	found, err := getSingle(&res,
		"query", "pki", "rejected-x509-root-cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetAllRejectedX509RootCerts queries all rejected root certificates.
func GetAllRejectedX509RootCerts() ([]pkitypes.RejectedCertificate, error) {
	var res pkitypes.QueryAllRejectedCertificatesResponse
	if err := getList(&res, "query", "pki", "all-rejected-x509-root-certs", "-o", "json"); err != nil {
		return nil, err
	}

	return res.RejectedCertificate, nil
}

// GetChildX509Certs queries all child certificates by (subject, SKID). Returns
// nil when no record matches.
func GetChildX509Certs(subject, subjectKeyID string) (*pkitypes.ChildCertificates, error) {
	var res pkitypes.ChildCertificates
	found, err := getSingle(&res,
		"query", "pki", "all-child-x509-certs",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetNocRootCerts queries NOC root certificates for a vid. Returns nil when
// no record matches.
func GetNocRootCerts(vid int) (*pkitypes.NocRootCertificates, error) {
	var res pkitypes.NocRootCertificates
	found, err := getSingle(&res,
		"query", "pki", "noc-x509-root-certs",
		"--vid", itoa(vid),
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetAllNocRootCerts queries all NOC root certificates.
func GetAllNocRootCerts() ([]pkitypes.NocRootCertificates, error) {
	var res pkitypes.QueryAllNocRootCertificatesResponse
	if err := getList(&res, "query", "pki", "all-noc-x509-root-certs", "-o", "json"); err != nil {
		return nil, err
	}

	return res.NocRootCertificates, nil
}

// GetNocX509IcaCerts queries NOC ICA certificates for a vid. Returns nil when
// no record matches.
func GetNocX509IcaCerts(vid int) (*pkitypes.NocIcaCertificates, error) {
	var res pkitypes.NocIcaCertificates
	found, err := getSingle(&res,
		"query", "pki", "noc-x509-ica-certs",
		"--vid", itoa(vid),
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetAllNocX509IcaCerts queries all NOC ICA certificates.
func GetAllNocX509IcaCerts() ([]pkitypes.NocIcaCertificates, error) {
	var res pkitypes.QueryAllNocIcaCertificatesResponse
	if err := getList(&res, "query", "pki", "all-noc-x509-ica-certs", "-o", "json"); err != nil {
		return nil, err
	}

	return res.NocIcaCertificates, nil
}

// GetAllNocX509Certs queries all NOC certificates (root + ICA). The CLI emits
// a QueryNocCertificatesResponse wrapping a list of NocCertificates, each
// carrying its own Certs slice. We flatten the slices into a single
// AllCertificates wrapper so callers can iterate over one Certs list.
func GetAllNocX509Certs() (*pkitypes.AllCertificates, error) {
	var res pkitypes.QueryNocCertificatesResponse
	if err := getList(&res, "query", "pki", "all-noc-x509-certs", "-o", "json"); err != nil {
		return nil, err
	}
	all := &pkitypes.AllCertificates{}
	for i := range res.NocCertificates {
		all.Certs = append(all.Certs, res.NocCertificates[i].Certs...)
	}

	return all, nil
}

// GetNocCert queries a NOC certificate by various flags. Pass --subject /
// --subject-key-id / --vid as extra args. Returns nil when no record matches.
// The CLI dispatches to NocCertificates or NocCertificatesByVidAndSkid depending
// on the flags, but both share the cert layout; the wider NocCertificates
// shape is used here.
func GetNocCert(extra ...string) (*pkitypes.NocCertificates, error) {
	var res pkitypes.NocCertificates
	args := []string{"query", "pki", "noc-x509-cert", "-o", "json"}
	args = append(args, extra...)
	found, err := getSingle(&res, args...)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetNocSubjectCerts queries all NOC certificates by subject. Returns nil when
// no record matches.
func GetNocSubjectCerts(subject string) (*pkitypes.NocCertificatesBySubject, error) {
	var res pkitypes.NocCertificatesBySubject
	found, err := getSingle(&res,
		"query", "pki", "all-noc-subject-x509-certs",
		"--subject", subject,
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetAllRevokedNocRootCerts queries all revoked NOC root certificates.
func GetAllRevokedNocRootCerts() ([]pkitypes.RevokedNocRootCertificates, error) {
	var res pkitypes.QueryAllRevokedNocRootCertificatesResponse
	if err := getList(&res, "query", "pki", "all-revoked-noc-x509-root-certs", "-o", "json"); err != nil {
		return nil, err
	}

	return res.RevokedNocRootCertificates, nil
}

// GetRevokedNocRootCert queries a revoked NOC root certificate. Returns nil
// when no record matches.
func GetRevokedNocRootCert(subject, subjectKeyID string) (*pkitypes.RevokedNocRootCertificates, error) {
	var res pkitypes.RevokedNocRootCertificates
	found, err := getSingle(&res,
		"query", "pki", "revoked-noc-x509-root-cert",
		"--subject", subject,
		"--subject-key-id", subjectKeyID,
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetAllRevokedNocX509IcaCerts queries all revoked NOC ICA certificates.
func GetAllRevokedNocX509IcaCerts() ([]pkitypes.RevokedNocIcaCertificates, error) {
	var res pkitypes.QueryAllRevokedNocIcaCertificatesResponse
	if err := getList(&res, "query", "pki", "all-revoked-noc-x509-ica-certs", "-o", "json"); err != nil {
		return nil, err
	}

	return res.RevokedNocIcaCertificates, nil
}

// GetPkiRevocationDistributionPoint queries a revocation distribution point.
// Returns nil when no record matches.
func GetPkiRevocationDistributionPoint(vid int, label, issuerSubjectKeyID string) (*pkitypes.PkiRevocationDistributionPoint, error) {
	var res pkitypes.PkiRevocationDistributionPoint
	found, err := getSingle(&res,
		"query", "pki", "revocation-point",
		"--vid", itoa(vid),
		"--label", label,
		"--issuer-subject-key-id", issuerSubjectKeyID,
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetAllPkiRevocationDistributionPoints queries all revocation distribution points.
func GetAllPkiRevocationDistributionPoints() ([]pkitypes.PkiRevocationDistributionPoint, error) {
	var res pkitypes.QueryAllPkiRevocationDistributionPointResponse
	if err := getList(&res, "query", "pki", "all-revocation-points", "-o", "json"); err != nil {
		return nil, err
	}

	return res.PkiRevocationDistributionPoint, nil
}

// GetPkiRevocationDistributionPointsByIssuer queries revocation distribution
// points by issuer SKID. Returns nil when no record matches.
func GetPkiRevocationDistributionPointsByIssuer(issuerSubjectKeyID string) (*pkitypes.PkiRevocationDistributionPointsByIssuerSubjectKeyID, error) {
	var res pkitypes.PkiRevocationDistributionPointsByIssuerSubjectKeyID
	found, err := getSingle(&res,
		"query", "pki", "revocation-points",
		"--issuer-subject-key-id", issuerSubjectKeyID,
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// findCertBySerial returns the *Certificate in certs whose SerialNumber matches.
func findCertBySerial(certs []*pkitypes.Certificate, serial string) *pkitypes.Certificate {
	for _, c := range certs {
		if c != nil && c.SerialNumber == serial {
			return c
		}
	}

	return nil
}

// containsCertSerial reports whether certs has an entry with the given serial.
func containsCertSerial(certs []*pkitypes.Certificate, serial string) bool {
	return findCertBySerial(certs, serial) != nil
}

// containsCertSubjectSerial reports whether certs has a (subject, serial) match.
func containsCertSubjectSerial(certs []*pkitypes.Certificate, subject, serial string) bool {
	for _, c := range certs {
		if c != nil && c.Subject == subject && c.SerialNumber == serial {
			return true
		}
	}

	return false
}

// grantsContain reports whether a list of approval/reject grants includes the
// given address. Used to assert which trustees voted on a proposed/approved/
// revoked/rejected certificate.
func grantsContain(grants []*pkitypes.Grant, addr string) bool {
	for _, g := range grants {
		if g != nil && g.Address == addr {
			return true
		}
	}

	return false
}

// containsApprovedCertSerial reports whether any of the ApprovedCertificates'
// inner Certs slices contains the given serial.
func containsApprovedCertSerial(list []pkitypes.ApprovedCertificates, serial string) bool {
	for i := range list {
		if findCertBySerial(list[i].Certs, serial) != nil {
			return true
		}
	}

	return false
}

// containsApprovedCertSubjectSerial reports whether any ApprovedCertificates
// entry has a cert matching both the subject and the serial — used when
// multiple unrelated certs in the ledger share a serial number.
func containsApprovedCertSubjectSerial(list []pkitypes.ApprovedCertificates, subject, serial string) bool {
	for i := range list {
		if containsCertSubjectSerial(list[i].Certs, subject, serial) {
			return true
		}
	}

	return false
}

// containsRevokedCertSerial does the same for RevokedCertificates.
func containsRevokedCertSerial(list []pkitypes.RevokedCertificates, serial string) bool {
	for i := range list {
		if findCertBySerial(list[i].Certs, serial) != nil {
			return true
		}
	}

	return false
}

// containsRevokedCertSubjectSerial reports whether any RevokedCertificates entry
// has a cert matching both the subject and the serial — used when multiple
// unrelated certs in the ledger share a serial number.
func containsRevokedCertSubjectSerial(list []pkitypes.RevokedCertificates, subject, serial string) bool {
	for i := range list {
		if containsCertSubjectSerial(list[i].Certs, subject, serial) {
			return true
		}
	}

	return false
}

// containsRevokedNocRootCertSerial does the same for RevokedNocRootCertificates.
func containsRevokedNocRootCertSerial(list []pkitypes.RevokedNocRootCertificates, serial string) bool {
	for i := range list {
		if findCertBySerial(list[i].Certs, serial) != nil {
			return true
		}
	}

	return false
}

// containsRevokedNocIcaCertSerial does the same for RevokedNocIcaCertificates.
func containsRevokedNocIcaCertSerial(list []pkitypes.RevokedNocIcaCertificates, serial string) bool {
	for i := range list {
		if findCertBySerial(list[i].Certs, serial) != nil {
			return true
		}
	}

	return false
}

// containsRevokedNocIcaCertSubject reports whether any RevokedNocIcaCertificates
// entry has the given subject.
func containsRevokedNocIcaCertSubject(list []pkitypes.RevokedNocIcaCertificates, subject string) bool {
	for i := range list {
		if list[i].Subject == subject {
			return true
		}
	}

	return false
}

// containsNocRootCertSerial does the same for NocRootCertificates.
func containsNocRootCertSerial(list []pkitypes.NocRootCertificates, serial string) bool {
	for i := range list {
		if findCertBySerial(list[i].Certs, serial) != nil {
			return true
		}
	}

	return false
}

// containsNocIcaCertSerial does the same for NocIcaCertificates.
func containsNocIcaCertSerial(list []pkitypes.NocIcaCertificates, serial string) bool {
	for i := range list {
		if findCertBySerial(list[i].Certs, serial) != nil {
			return true
		}
	}

	return false
}

// containsRevocationPointByLabel reports whether the list has an entry with
// matching vid + label.
func containsRevocationPointByLabel(list []pkitypes.PkiRevocationDistributionPoint, vid int32, label string) bool {
	for i := range list {
		if list[i].Vid == vid && list[i].Label == label {
			return true
		}
	}

	return false
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
