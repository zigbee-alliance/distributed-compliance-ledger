package model

import (
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// AddModelOpts holds parameters for the add-model transaction.
// Zero / empty values for the "primary" fields fall back to test-friendly
// defaults (DeviceTypeID=1, ProductName="TestProduct", ProductLabel=
// "TestingProductLabel", PartNumber="1"). CommissioningCustomFlow and
// EnhancedSetupFlowOptions are always sent — pass 0 for the typical case.
// Fields with no default (URLs, TC fields, schema version, etc.) are only
// emitted as flags when non-zero / non-empty.
type AddModelOpts struct {
	VID    int
	VIDHex string
	PID    int
	PIDHex string

	DeviceTypeID             int
	ProductName              string
	ProductLabel             string
	PartNumber               string
	CommissioningCustomFlow  int
	EnhancedSetupFlowOptions int

	EnhancedSetupFlowTCUrl       string
	EnhancedSetupFlowTCRevision  int
	EnhancedSetupFlowTCDigest    string
	EnhancedSetupFlowTCFileSize  int
	MaintenanceURL               string
	CommissioningFallbackURL     string
	DiscoveryCapabilitiesBitmask int

	// SchemaVersion is sent only when non-empty (the on-chain default is 0).
	SchemaVersion string

	From  string
	Extra []string
}

// AddModel executes the add-model transaction.
func AddModel(opts AddModelOpts) (*utils.TxResult, error) {
	deviceType := opts.DeviceTypeID
	if deviceType == 0 {
		deviceType = 1
	}
	productName := opts.ProductName
	if productName == "" {
		productName = "TestProduct"
	}
	productLabel := opts.ProductLabel
	if productLabel == "" {
		productLabel = "TestingProductLabel"
	}
	partNumber := opts.PartNumber
	if partNumber == "" {
		partNumber = "1"
	}

	args := []string{
		"tx", "model", "add-model",
		"--vid", flagOrHex(opts.VID, opts.VIDHex),
		"--pid", flagOrHex(opts.PID, opts.PIDHex),
		"--deviceTypeID", itoa(deviceType),
		"--productName", productName,
		"--productLabel", productLabel,
		"--partNumber", partNumber,
		"--commissioningCustomFlow", itoa(opts.CommissioningCustomFlow),
		"--enhancedSetupFlowOptions", itoa(opts.EnhancedSetupFlowOptions),
		"--from", opts.From,
	}

	if opts.EnhancedSetupFlowTCUrl != "" {
		args = append(args, "--enhancedSetupFlowTCUrl", opts.EnhancedSetupFlowTCUrl)
	}
	if opts.EnhancedSetupFlowTCRevision != 0 {
		args = append(args, "--enhancedSetupFlowTCRevision", itoa(opts.EnhancedSetupFlowTCRevision))
	}
	if opts.EnhancedSetupFlowTCDigest != "" {
		args = append(args, "--enhancedSetupFlowTCDigest", opts.EnhancedSetupFlowTCDigest)
	}
	if opts.EnhancedSetupFlowTCFileSize != 0 {
		args = append(args, "--enhancedSetupFlowTCFileSize", itoa(opts.EnhancedSetupFlowTCFileSize))
	}
	if opts.MaintenanceURL != "" {
		args = append(args, "--maintenanceUrl", opts.MaintenanceURL)
	}
	if opts.CommissioningFallbackURL != "" {
		args = append(args, "--commissioningFallbackUrl", opts.CommissioningFallbackURL)
	}
	if opts.DiscoveryCapabilitiesBitmask != 0 {
		args = append(args, "--discoveryCapabilitiesBitmask", itoa(opts.DiscoveryCapabilitiesBitmask))
	}
	if opts.SchemaVersion != "" {
		args = append(args, "--schemaVersion", opts.SchemaVersion)
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

// UpdateModel executes the update-model transaction.
func UpdateModel(vid, pid int, from string, extra ...string) (*utils.TxResult, error) {
	return UpdateModelHex(itoa(vid), itoa(pid), from, extra...)
}

// UpdateModelHex executes update-model using hex vid/pid strings.
func UpdateModelHex(vid, pid, from string, extra ...string) (*utils.TxResult, error) {
	args := []string{
		"tx", "model", "update-model",
		"--vid", vid,
		"--pid", pid,
		"--from", from,
	}
	args = append(args, extra...)

	return utils.ExecuteTx(args...)
}

// DeleteModel executes the delete-model transaction.
func DeleteModel(vid, pid int, from string) (*utils.TxResult, error) {
	return DeleteModelHex(itoa(vid), itoa(pid), from)
}

// DeleteModelHex executes delete-model using hex vid/pid strings.
func DeleteModelHex(vid, pid, from string) (*utils.TxResult, error) {
	return utils.ExecuteTx("tx", "model", "delete-model",
		"--vid", vid,
		"--pid", pid,
		"--from", from,
	)
}

// AddModelVersionOpts holds parameters for the add-model-version transaction.
// CDVersionNumber, MinApplicableSoftwareVersion default to 1;
// MaxApplicableSoftwareVersion defaults to 10.
type AddModelVersionOpts struct {
	VID                          int
	VIDHex                       string
	PID                          int
	PIDHex                       string
	SoftwareVersion              int
	SoftwareVersionString        string
	CDVersionNumber              int
	MinApplicableSoftwareVersion int
	MaxApplicableSoftwareVersion int

	OtaURL      string
	OtaFileSize int
	OtaChecksum string

	SchemaVersion string

	From  string
	Extra []string
}

// AddModelVersion executes the add-model-version transaction.
func AddModelVersion(opts AddModelVersionOpts) (*utils.TxResult, error) {
	cdVersion := opts.CDVersionNumber
	if cdVersion == 0 {
		cdVersion = 1
	}
	minSV := opts.MinApplicableSoftwareVersion
	if minSV == 0 {
		minSV = 1
	}
	maxSV := opts.MaxApplicableSoftwareVersion
	if maxSV == 0 {
		maxSV = 10
	}

	args := []string{
		"tx", "model", "add-model-version",
		"--vid", flagOrHex(opts.VID, opts.VIDHex),
		"--pid", flagOrHex(opts.PID, opts.PIDHex),
		"--softwareVersion", itoa(opts.SoftwareVersion),
		"--softwareVersionString", opts.SoftwareVersionString,
		"--cdVersionNumber", itoa(cdVersion),
		"--maxApplicableSoftwareVersion", itoa(maxSV),
		"--minApplicableSoftwareVersion", itoa(minSV),
		"--from", opts.From,
	}

	if opts.OtaURL != "" {
		args = append(args, "--otaURL", opts.OtaURL)
	}
	if opts.OtaFileSize != 0 {
		args = append(args, "--otaFileSize", itoa(opts.OtaFileSize))
	}
	if opts.OtaChecksum != "" {
		args = append(args, "--otaChecksum", opts.OtaChecksum)
	}
	if opts.SchemaVersion != "" {
		args = append(args, "--schemaVersion", opts.SchemaVersion)
	}
	args = append(args, opts.Extra...)

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
	return QueryVendorModelsHex(itoa(vid))
}

// QueryVendorModelsHex queries all models for a given vendor using a hex vid.
func QueryVendorModelsHex(vid string) ([]byte, error) {
	return utils.ExecuteCLI("query", "model", "vendor-models",
		"--vid", vid,
		"-o", "json",
	)
}

// QueryModelVersion queries a specific model version.
func QueryModelVersion(vid, pid, sv int) ([]byte, error) {
	return QueryModelVersionHex(itoa(vid), itoa(pid), sv)
}

// QueryModelVersionHex queries a specific model version using hex vid/pid.
func QueryModelVersionHex(vid, pid string, sv int) ([]byte, error) {
	return utils.ExecuteCLI("query", "model", "model-version",
		"--vid", vid,
		"--pid", pid,
		"--softwareVersion", itoa(sv),
		"-o", "json",
	)
}

// QueryAllModelVersions queries all model versions for a given vid/pid.
func QueryAllModelVersions(vid, pid int) ([]byte, error) {
	return QueryAllModelVersionsHex(itoa(vid), itoa(pid))
}

// QueryAllModelVersionsHex queries all model versions using hex vid/pid.
func QueryAllModelVersionsHex(vid, pid string) ([]byte, error) {
	return utils.ExecuteCLI("query", "model", "all-model-versions",
		"--vid", vid,
		"--pid", pid,
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
