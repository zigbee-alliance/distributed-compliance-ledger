package compliance

import (
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// CertifyModelOpts holds parameters for the certify-model transaction.
// Use VIDHex / PIDHex to pass hex-formatted identifiers (e.g. "0xA13"); when
// those are empty, the numeric VID / PID fields are formatted as decimal.
// Zero-valued SpecificationVersion and CDVersionNumber default to 1.
// Extra appends arbitrary flag tokens for fields not modeled here.
type CertifyModelOpts struct {
	VID                   int
	VIDHex                string
	PID                   int
	PIDHex                string
	SoftwareVersion       int
	SoftwareVersionString string
	CertificationType     string
	SpecificationVersion  int
	CertificationDate     string
	CDCertificateID       string
	CDVersionNumber       int
	From                  string
	Extra                 []string
}

// CertifyModel executes the certify-model transaction.
func CertifyModel(opts CertifyModelOpts) (*utils.TxResult, error) {
	specVersion := opts.SpecificationVersion
	if specVersion == 0 {
		specVersion = 1
	}
	cdVersion := opts.CDVersionNumber
	if cdVersion == 0 {
		cdVersion = 1
	}

	args := []string{
		"tx", "compliance", "certify-model",
		"--vid", flagOrHex(opts.VID, opts.VIDHex),
		"--pid", flagOrHex(opts.PID, opts.PIDHex),
		"--softwareVersion", itoa(opts.SoftwareVersion),
		"--softwareVersionString", opts.SoftwareVersionString,
		"--certificationType", opts.CertificationType,
		"--specificationVersion", itoa(specVersion),
		"--certificationDate", opts.CertificationDate,
		"--cdCertificateId", opts.CDCertificateID,
		"--cdVersionNumber", itoa(cdVersion),
		"--from", opts.From,
	}
	args = append(args, opts.Extra...)

	return utils.ExecuteTx(args...)
}

// flagOrHex returns hex if non-empty, otherwise the decimal-formatted n.
func flagOrHex(n int, hex string) string {
	if hex != "" {
		return hex
	}

	return itoa(n)
}

// RevokeModelOpts holds parameters for the revoke-model transaction.
// Zero-valued CDVersionNumber defaults to 1.
type RevokeModelOpts struct {
	VID                   int
	VIDHex                string
	PID                   int
	PIDHex                string
	SoftwareVersion       int
	SoftwareVersionString string
	CertificationType     string
	RevocationDate        string
	Reason                string
	CDVersionNumber       int
	From                  string
	Extra                 []string
}

// RevokeModel executes the revoke-model transaction.
func RevokeModel(opts RevokeModelOpts) (*utils.TxResult, error) {
	cdVersion := opts.CDVersionNumber
	if cdVersion == 0 {
		cdVersion = 1
	}

	args := []string{
		"tx", "compliance", "revoke-model",
		"--vid", flagOrHex(opts.VID, opts.VIDHex),
		"--pid", flagOrHex(opts.PID, opts.PIDHex),
		"--softwareVersion", itoa(opts.SoftwareVersion),
		"--softwareVersionString", opts.SoftwareVersionString,
		"--certificationType", opts.CertificationType,
		"--revocationDate", opts.RevocationDate,
		"--cdVersionNumber", itoa(cdVersion),
		"--from", opts.From,
	}
	if opts.Reason != "" {
		args = append(args, "--reason", opts.Reason)
	}
	args = append(args, opts.Extra...)

	return utils.ExecuteTx(args...)
}

// ProvisionModelOpts holds parameters for the provision-model transaction.
// Zero-valued SpecificationVersion and CDVersionNumber default to 1.
type ProvisionModelOpts struct {
	VID                   int
	VIDHex                string
	PID                   int
	PIDHex                string
	SoftwareVersion       int
	SoftwareVersionString string
	CertificationType     string
	SpecificationVersion  int
	ProvisionalDate       string
	CDCertificateID       string
	CDVersionNumber       int
	Reason                string
	From                  string
	Extra                 []string
}

// ProvisionModel executes the provision-model transaction.
func ProvisionModel(opts ProvisionModelOpts) (*utils.TxResult, error) {
	specVersion := opts.SpecificationVersion
	if specVersion == 0 {
		specVersion = 1
	}
	cdVersion := opts.CDVersionNumber
	if cdVersion == 0 {
		cdVersion = 1
	}

	args := []string{
		"tx", "compliance", "provision-model",
		"--vid", flagOrHex(opts.VID, opts.VIDHex),
		"--pid", flagOrHex(opts.PID, opts.PIDHex),
		"--softwareVersion", itoa(opts.SoftwareVersion),
		"--softwareVersionString", opts.SoftwareVersionString,
		"--certificationType", opts.CertificationType,
		"--specificationVersion", itoa(specVersion),
		"--provisionalDate", opts.ProvisionalDate,
		"--cdCertificateId", opts.CDCertificateID,
		"--cdVersionNumber", itoa(cdVersion),
		"--from", opts.From,
	}
	if opts.Reason != "" {
		args = append(args, "--reason", opts.Reason)
	}
	args = append(args, opts.Extra...)

	return utils.ExecuteTx(args...)
}

// UpdateComplianceInfo executes the update-compliance-info transaction.
func UpdateComplianceInfo(vid, pid, sv int, certType, from string, extra ...string) (*utils.TxResult, error) {
	args := []string{
		"tx", "compliance", "update-compliance-info",
		"--vid", itoa(vid),
		"--pid", itoa(pid),
		"--softwareVersion", itoa(sv),
		"--certificationType", certType,
		"--from", from,
	}
	args = append(args, extra...)

	return utils.ExecuteTx(args...)
}

// DeleteComplianceInfo executes the delete-compliance-info transaction.
func DeleteComplianceInfo(vid, pid, sv int, certType, from string) (*utils.TxResult, error) {
	return utils.ExecuteTx("tx", "compliance", "delete-compliance-info",
		"--vid", itoa(vid),
		"--pid", itoa(pid),
		"--softwareVersion", itoa(sv),
		"--certificationType", certType,
		"--from", from,
	)
}

// ComplianceQueryOpts holds parameters shared by the per-model compliance
// queries (compliance-info, certified-model, revoked-model, provisional-model).
// Use VIDHex / PIDHex for hex-formatted identifiers; otherwise the numeric
// VID / PID are formatted as decimal.
type ComplianceQueryOpts struct {
	VID               int
	VIDHex            string
	PID               int
	PIDHex            string
	SoftwareVersion   int
	CertificationType string
}

func (o ComplianceQueryOpts) args() []string {
	return []string{
		"--vid", flagOrHex(o.VID, o.VIDHex),
		"--pid", flagOrHex(o.PID, o.PIDHex),
		"--softwareVersion", itoa(o.SoftwareVersion),
		"--certificationType", o.CertificationType,
		"-o", "json",
	}
}

// QueryComplianceInfo queries compliance-info for a given vid/pid/sv/certType.
func QueryComplianceInfo(opts ComplianceQueryOpts) ([]byte, error) {
	return utils.ExecuteCLI(append([]string{"query", "compliance", "compliance-info"}, opts.args()...)...)
}

// QueryCertifiedModel queries the certified-model endpoint.
func QueryCertifiedModel(opts ComplianceQueryOpts) ([]byte, error) {
	return utils.ExecuteCLI(append([]string{"query", "compliance", "certified-model"}, opts.args()...)...)
}

// QueryRevokedModel queries the revoked-model endpoint.
func QueryRevokedModel(opts ComplianceQueryOpts) ([]byte, error) {
	return utils.ExecuteCLI(append([]string{"query", "compliance", "revoked-model"}, opts.args()...)...)
}

// QueryProvisionalModel queries the provisional-model endpoint.
func QueryProvisionalModel(opts ComplianceQueryOpts) ([]byte, error) {
	return utils.ExecuteCLI(append([]string{"query", "compliance", "provisional-model"}, opts.args()...)...)
}

// QueryDeviceSoftwareCompliance queries device-software-compliance by CDCertificateID.
func QueryDeviceSoftwareCompliance(cdCertificateID string) ([]byte, error) {
	return utils.ExecuteCLI("query", "compliance", "device-software-compliance",
		"--cdCertificateId", cdCertificateID,
		"-o", "json",
	)
}

// QueryAllComplianceInfo queries all compliance info records.
func QueryAllComplianceInfo() ([]byte, error) {
	return utils.ExecuteCLI("query", "compliance", "all-compliance-info", "-o", "json")
}

// QueryAllCertifiedModels queries all certified models.
func QueryAllCertifiedModels() ([]byte, error) {
	return utils.ExecuteCLI("query", "compliance", "all-certified-models", "-o", "json")
}

// QueryAllRevokedModels queries all revoked models.
func QueryAllRevokedModels() ([]byte, error) {
	return utils.ExecuteCLI("query", "compliance", "all-revoked-models", "-o", "json")
}

// QueryAllProvisionalModels queries all provisional models.
func QueryAllProvisionalModels() ([]byte, error) {
	return utils.ExecuteCLI("query", "compliance", "all-provisional-models", "-o", "json")
}

// QueryAllDeviceSoftwareCompliance queries all device software compliance records.
func QueryAllDeviceSoftwareCompliance() ([]byte, error) {
	return utils.ExecuteCLI("query", "compliance", "all-device-software-compliance", "-o", "json")
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
