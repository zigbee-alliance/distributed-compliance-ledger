package compliance

import (
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// CertifyModel executes the certify-model transaction.
func CertifyModel(vid, pid, sv int, svs, certType, certDate, cdCertID, from string, extra ...string) (*utils.TxResult, error) {
	args := []string{
		"tx", "compliance", "certify-model",
		"--vid", itoa(vid),
		"--pid", itoa(pid),
		"--softwareVersion", itoa(sv),
		"--softwareVersionString", svs,
		"--certificationType", certType,
		"--certificationDate", certDate,
		"--cdCertificateId", cdCertID,
		"--from", from,
	}
	args = append(args, extra...)
	return utils.ExecuteTx(args...)
}

// RevokeModel executes the revoke-model transaction.
func RevokeModel(vid, pid, sv int, svs, certType, revocationDate, from string, extra ...string) (*utils.TxResult, error) {
	args := []string{
		"tx", "compliance", "revoke-model",
		"--vid", itoa(vid),
		"--pid", itoa(pid),
		"--softwareVersion", itoa(sv),
		"--softwareVersionString", svs,
		"--certificationType", certType,
		"--revocationDate", revocationDate,
		"--from", from,
	}
	args = append(args, extra...)
	return utils.ExecuteTx(args...)
}

// ProvisionModel executes the provision-model transaction.
func ProvisionModel(vid, pid, sv int, svs, certType, provisionalDate, cdCertID, from string, extra ...string) (*utils.TxResult, error) {
	args := []string{
		"tx", "compliance", "provision-model",
		"--vid", itoa(vid),
		"--pid", itoa(pid),
		"--softwareVersion", itoa(sv),
		"--softwareVersionString", svs,
		"--certificationType", certType,
		"--provisionalDate", provisionalDate,
		"--cdCertificateId", cdCertID,
		"--from", from,
	}
	args = append(args, extra...)
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

// QueryComplianceInfo queries compliance-info for a given vid/pid/sv/certType.
func QueryComplianceInfo(vid, pid, sv int, certType string) ([]byte, error) {
	return utils.ExecuteCLI("query", "compliance", "compliance-info",
		"--vid", itoa(vid),
		"--pid", itoa(pid),
		"--softwareVersion", itoa(sv),
		"--certificationType", certType,
		"-o", "json",
	)
}

// QueryCertifiedModel queries the certified-model endpoint.
func QueryCertifiedModel(vid, pid, sv int, certType string) ([]byte, error) {
	return utils.ExecuteCLI("query", "compliance", "certified-model",
		"--vid", itoa(vid),
		"--pid", itoa(pid),
		"--softwareVersion", itoa(sv),
		"--certificationType", certType,
		"-o", "json",
	)
}

// QueryRevokedModel queries the revoked-model endpoint.
func QueryRevokedModel(vid, pid, sv int, certType string) ([]byte, error) {
	return utils.ExecuteCLI("query", "compliance", "revoked-model",
		"--vid", itoa(vid),
		"--pid", itoa(pid),
		"--softwareVersion", itoa(sv),
		"--certificationType", certType,
		"-o", "json",
	)
}

// QueryProvisionalModel queries the provisional-model endpoint.
func QueryProvisionalModel(vid, pid, sv int, certType string) ([]byte, error) {
	return utils.ExecuteCLI("query", "compliance", "provisional-model",
		"--vid", itoa(vid),
		"--pid", itoa(pid),
		"--softwareVersion", itoa(sv),
		"--certificationType", certType,
		"-o", "json",
	)
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
