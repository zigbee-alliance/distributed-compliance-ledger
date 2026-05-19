package model

import (
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// AddModel executes the add-model transaction.
func AddModel(vid, pid int, from string, extra ...string) (*utils.TxResult, error) {
	args := []string{
		"tx", "model", "add-model",
		"--vid", itoa(vid),
		"--pid", itoa(pid),
		"--deviceTypeID", "1",
		"--productName", "TestProduct",
		"--productLabel", "TestingProductLabel",
		"--partNumber", "1",
		"--commissioningCustomFlow", "0",
		"--enhancedSetupFlowOptions", "0",
		"--from", from,
	}
	args = append(args, extra...)

	return utils.ExecuteTx(args...)
}

// UpdateModel executes the update-model transaction.
func UpdateModel(vid, pid int, from string, extra ...string) (*utils.TxResult, error) {
	args := []string{
		"tx", "model", "update-model",
		"--vid", itoa(vid),
		"--pid", itoa(pid),
		"--from", from,
	}
	args = append(args, extra...)

	return utils.ExecuteTx(args...)
}

// DeleteModel executes the delete-model transaction.
func DeleteModel(vid, pid int, from string) (*utils.TxResult, error) {
	return utils.ExecuteTx("tx", "model", "delete-model",
		"--vid", itoa(vid),
		"--pid", itoa(pid),
		"--from", from,
	)
}

// AddModelVersion executes the add-model-version transaction.
func AddModelVersion(vid, pid, sv int, svs string, from string, extra ...string) (*utils.TxResult, error) {
	args := []string{
		"tx", "model", "add-model-version",
		"--vid", itoa(vid),
		"--pid", itoa(pid),
		"--softwareVersion", itoa(sv),
		"--softwareVersionString", svs,
		"--cdVersionNumber", "1",
		"--maxApplicableSoftwareVersion", "10",
		"--minApplicableSoftwareVersion", "1",
		"--from", from,
	}
	args = append(args, extra...)

	return utils.ExecuteTx(args...)
}

// UpdateModelVersion executes the update-model-version transaction.
func UpdateModelVersion(vid, pid, sv int, from string, extra ...string) (*utils.TxResult, error) {
	args := []string{
		"tx", "model", "update-model-version",
		"--vid", itoa(vid),
		"--pid", itoa(pid),
		"--softwareVersion", itoa(sv),
		"--from", from,
	}
	args = append(args, extra...)

	return utils.ExecuteTx(args...)
}

// DeleteModelVersion executes the delete-model-version transaction.
func DeleteModelVersion(vid, pid, sv int, from string) (*utils.TxResult, error) {
	return utils.ExecuteTx("tx", "model", "delete-model-version",
		"--vid", itoa(vid),
		"--pid", itoa(pid),
		"--softwareVersion", itoa(sv),
		"--from", from,
	)
}

// QueryModel queries a specific model by vid/pid.
func QueryModel(vid, pid int) ([]byte, error) {
	return utils.ExecuteCLI("query", "model", "get-model",
		"--vid", itoa(vid),
		"--pid", itoa(pid),
		"-o", "json",
	)
}

// QueryModelHex queries a model using hex-format vid/pid strings.
func QueryModelHex(vid, pid string) ([]byte, error) {
	return utils.ExecuteCLI("query", "model", "get-model",
		"--vid", vid,
		"--pid", pid,
		"-o", "json",
	)
}

// QueryAllModels queries all models.
func QueryAllModels() ([]byte, error) {
	return utils.ExecuteCLI("query", "model", "all-models", "-o", "json")
}

// QueryVendorModels queries all models for a given vendor.
func QueryVendorModels(vid int) ([]byte, error) {
	return utils.ExecuteCLI("query", "model", "vendor-models",
		"--vid", itoa(vid),
		"-o", "json",
	)
}

// QueryModelVersion queries a specific model version.
func QueryModelVersion(vid, pid, sv int) ([]byte, error) {
	return utils.ExecuteCLI("query", "model", "model-version",
		"--vid", itoa(vid),
		"--pid", itoa(pid),
		"--softwareVersion", itoa(sv),
		"-o", "json",
	)
}

// QueryAllModelVersions queries all model versions for a given vid/pid.
func QueryAllModelVersions(vid, pid int) ([]byte, error) {
	return utils.ExecuteCLI("query", "model", "all-model-versions",
		"--vid", itoa(vid),
		"--pid", itoa(pid),
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
