package model

import (
	"strconv"

	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
	modeltypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
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

	// Commissioning hints/instructions and product URLs. Hints have on-chain
	// defaults (1/4/1/1) so they are emitted only when non-zero; the rest emit
	// when non-empty.
	CommissioningCustomFlowURL                 string
	CommissioningModeInitialStepsHint          int
	CommissioningModeInitialStepsInstruction   string
	CommissioningModeSecondaryStepsHint        int
	CommissioningModeSecondaryStepsInstruction string
	IcdUserActiveModeTriggerHint               int
	IcdUserActiveModeTriggerInstruction        string
	FactoryResetStepsHint                      int
	FactoryResetStepsInstruction               string
	UserManualURL                              string
	ProductURL                                 string
	LsfURL                                     string
	SupportURL                                 string

	// SchemaVersion is sent only when non-empty (the on-chain default is 0).
	SchemaVersion string

	From string
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
		"--vid", cliputils.FlagOrHex(opts.VID, opts.VIDHex),
		"--pid", cliputils.FlagOrHex(opts.PID, opts.PIDHex),
		"--deviceTypeID", strconv.Itoa(deviceType),
		"--productName", productName,
		"--productLabel", productLabel,
		"--partNumber", partNumber,
		"--commissioningCustomFlow", strconv.Itoa(opts.CommissioningCustomFlow),
		"--enhancedSetupFlowOptions", strconv.Itoa(opts.EnhancedSetupFlowOptions),
		"--from", opts.From,
	}

	if opts.EnhancedSetupFlowTCUrl != "" {
		args = append(args, "--enhancedSetupFlowTCUrl", opts.EnhancedSetupFlowTCUrl)
	}
	if opts.EnhancedSetupFlowTCRevision != 0 {
		args = append(args, "--enhancedSetupFlowTCRevision", strconv.Itoa(opts.EnhancedSetupFlowTCRevision))
	}
	if opts.EnhancedSetupFlowTCDigest != "" {
		args = append(args, "--enhancedSetupFlowTCDigest", opts.EnhancedSetupFlowTCDigest)
	}
	if opts.EnhancedSetupFlowTCFileSize != 0 {
		args = append(args, "--enhancedSetupFlowTCFileSize", strconv.Itoa(opts.EnhancedSetupFlowTCFileSize))
	}
	if opts.MaintenanceURL != "" {
		args = append(args, "--maintenanceUrl", opts.MaintenanceURL)
	}
	if opts.CommissioningFallbackURL != "" {
		args = append(args, "--commissioningFallbackUrl", opts.CommissioningFallbackURL)
	}
	if opts.DiscoveryCapabilitiesBitmask != 0 {
		args = append(args, "--discoveryCapabilitiesBitmask", strconv.Itoa(opts.DiscoveryCapabilitiesBitmask))
	}
	if opts.CommissioningCustomFlowURL != "" {
		args = append(args, "--commissioningCustomFlowURL", opts.CommissioningCustomFlowURL)
	}
	if opts.CommissioningModeInitialStepsHint != 0 {
		args = append(args, "--commissioningModeInitialStepsHint", strconv.Itoa(opts.CommissioningModeInitialStepsHint))
	}
	if opts.CommissioningModeInitialStepsInstruction != "" {
		args = append(args, "--commissioningModeInitialStepsInstruction", opts.CommissioningModeInitialStepsInstruction)
	}
	if opts.CommissioningModeSecondaryStepsHint != 0 {
		args = append(args, "--commissioningModeSecondaryStepsHint", strconv.Itoa(opts.CommissioningModeSecondaryStepsHint))
	}
	if opts.CommissioningModeSecondaryStepsInstruction != "" {
		args = append(args, "--commissioningModeSecondaryStepsInstruction", opts.CommissioningModeSecondaryStepsInstruction)
	}
	if opts.IcdUserActiveModeTriggerHint != 0 {
		args = append(args, "--icdUserActiveModeTriggerHint", strconv.Itoa(opts.IcdUserActiveModeTriggerHint))
	}
	if opts.IcdUserActiveModeTriggerInstruction != "" {
		args = append(args, "--icdUserActiveModeTriggerInstruction", opts.IcdUserActiveModeTriggerInstruction)
	}
	if opts.FactoryResetStepsHint != 0 {
		args = append(args, "--factoryResetStepsHint", strconv.Itoa(opts.FactoryResetStepsHint))
	}
	if opts.FactoryResetStepsInstruction != "" {
		args = append(args, "--factoryResetStepsInstruction", opts.FactoryResetStepsInstruction)
	}
	if opts.UserManualURL != "" {
		args = append(args, "--userManualURL", opts.UserManualURL)
	}
	if opts.ProductURL != "" {
		args = append(args, "--productURL", opts.ProductURL)
	}
	if opts.LsfURL != "" {
		args = append(args, "--lsfURL", opts.LsfURL)
	}
	if opts.SupportURL != "" {
		args = append(args, "--supportURL", opts.SupportURL)
	}
	if opts.SchemaVersion != "" {
		args = append(args, "--schemaVersion", opts.SchemaVersion)
	}

	return utils.ExecuteTx(args...)
}

// UpdateModelOpts holds parameters for update-model. VID/PID (or their hex
// forms) and From are always sent; every other field is emitted only when
// non-zero/non-empty. Because proto3 cannot distinguish an explicit zero from an
// omitted value, emitting numeric flags only when non-zero matches the update
// handler's "treat zero as leave-unchanged" semantics.
type UpdateModelOpts struct {
	VID    int
	VIDHex string
	PID    int
	PIDHex string

	ProductName  string
	ProductLabel string
	PartNumber   string

	CommissioningCustomFlowURL                 string
	CommissioningFallbackURL                   string
	CommissioningModeInitialStepsHint          int
	CommissioningModeInitialStepsInstruction   string
	CommissioningModeSecondaryStepsHint        int
	CommissioningModeSecondaryStepsInstruction string
	IcdUserActiveModeTriggerHint               int
	IcdUserActiveModeTriggerInstruction        string
	FactoryResetStepsHint                      int
	FactoryResetStepsInstruction               string

	UserManualURL  string
	ProductURL     string
	LsfURL         string
	LsfRevision    int
	SupportURL     string
	MaintenanceURL string

	EnhancedSetupFlowOptions    int
	EnhancedSetupFlowTCUrl      string
	EnhancedSetupFlowTCRevision int
	EnhancedSetupFlowTCDigest   string
	EnhancedSetupFlowTCFileSize int

	SchemaVersion string

	From string
}

// UpdateModel executes the update-model transaction.
func UpdateModel(opts UpdateModelOpts) (*utils.TxResult, error) {
	args := []string{
		"tx", "model", "update-model",
		"--vid", cliputils.FlagOrHex(opts.VID, opts.VIDHex),
		"--pid", cliputils.FlagOrHex(opts.PID, opts.PIDHex),
		"--from", opts.From,
	}
	if opts.ProductName != "" {
		args = append(args, "--productName", opts.ProductName)
	}
	if opts.ProductLabel != "" {
		args = append(args, "--productLabel", opts.ProductLabel)
	}
	if opts.PartNumber != "" {
		args = append(args, "--partNumber", opts.PartNumber)
	}
	if opts.CommissioningCustomFlowURL != "" {
		args = append(args, "--commissioningCustomFlowURL", opts.CommissioningCustomFlowURL)
	}
	if opts.CommissioningFallbackURL != "" {
		args = append(args, "--commissioningFallbackUrl", opts.CommissioningFallbackURL)
	}
	if opts.CommissioningModeInitialStepsHint != 0 {
		args = append(args, "--commissioningModeInitialStepsHint", strconv.Itoa(opts.CommissioningModeInitialStepsHint))
	}
	if opts.CommissioningModeInitialStepsInstruction != "" {
		args = append(args, "--commissioningModeInitialStepsInstruction", opts.CommissioningModeInitialStepsInstruction)
	}
	if opts.CommissioningModeSecondaryStepsHint != 0 {
		args = append(args, "--commissioningModeSecondaryStepsHint", strconv.Itoa(opts.CommissioningModeSecondaryStepsHint))
	}
	if opts.CommissioningModeSecondaryStepsInstruction != "" {
		args = append(args, "--commissioningModeSecondaryStepsInstruction", opts.CommissioningModeSecondaryStepsInstruction)
	}
	if opts.IcdUserActiveModeTriggerHint != 0 {
		args = append(args, "--icdUserActiveModeTriggerHint", strconv.Itoa(opts.IcdUserActiveModeTriggerHint))
	}
	if opts.IcdUserActiveModeTriggerInstruction != "" {
		args = append(args, "--icdUserActiveModeTriggerInstruction", opts.IcdUserActiveModeTriggerInstruction)
	}
	if opts.FactoryResetStepsHint != 0 {
		args = append(args, "--factoryResetStepsHint", strconv.Itoa(opts.FactoryResetStepsHint))
	}
	if opts.FactoryResetStepsInstruction != "" {
		args = append(args, "--factoryResetStepsInstruction", opts.FactoryResetStepsInstruction)
	}
	if opts.UserManualURL != "" {
		args = append(args, "--userManualURL", opts.UserManualURL)
	}
	if opts.ProductURL != "" {
		args = append(args, "--productURL", opts.ProductURL)
	}
	if opts.LsfURL != "" {
		args = append(args, "--lsfURL", opts.LsfURL)
	}
	if opts.LsfRevision != 0 {
		args = append(args, "--lsfRevision", strconv.Itoa(opts.LsfRevision))
	}
	if opts.SupportURL != "" {
		args = append(args, "--supportURL", opts.SupportURL)
	}
	if opts.MaintenanceURL != "" {
		args = append(args, "--maintenanceUrl", opts.MaintenanceURL)
	}
	if opts.EnhancedSetupFlowOptions != 0 {
		args = append(args, "--enhancedSetupFlowOptions", strconv.Itoa(opts.EnhancedSetupFlowOptions))
	}
	if opts.EnhancedSetupFlowTCUrl != "" {
		args = append(args, "--enhancedSetupFlowTCUrl", opts.EnhancedSetupFlowTCUrl)
	}
	if opts.EnhancedSetupFlowTCRevision != 0 {
		args = append(args, "--enhancedSetupFlowTCRevision", strconv.Itoa(opts.EnhancedSetupFlowTCRevision))
	}
	if opts.EnhancedSetupFlowTCDigest != "" {
		args = append(args, "--enhancedSetupFlowTCDigest", opts.EnhancedSetupFlowTCDigest)
	}
	if opts.EnhancedSetupFlowTCFileSize != 0 {
		args = append(args, "--enhancedSetupFlowTCFileSize", strconv.Itoa(opts.EnhancedSetupFlowTCFileSize))
	}
	if opts.SchemaVersion != "" {
		args = append(args, "--schemaVersion", opts.SchemaVersion)
	}

	return utils.ExecuteTx(args...)
}

// DeleteModel executes the delete-model transaction.
func DeleteModel(vid, pid int, from string) (*utils.TxResult, error) {
	return DeleteModelHex(strconv.Itoa(vid), strconv.Itoa(pid), from)
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

	OtaURL          string
	OtaFileSize     int
	OtaChecksum     string
	OtaChecksumType int

	FirmwareInformation  string
	SpecificationVersion int
	ReleaseNotesURL      string

	SchemaVersion string

	From string
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
		"--vid", cliputils.FlagOrHex(opts.VID, opts.VIDHex),
		"--pid", cliputils.FlagOrHex(opts.PID, opts.PIDHex),
		"--softwareVersion", strconv.Itoa(opts.SoftwareVersion),
		"--softwareVersionString", opts.SoftwareVersionString,
		"--cdVersionNumber", strconv.Itoa(cdVersion),
		"--maxApplicableSoftwareVersion", strconv.Itoa(maxSV),
		"--minApplicableSoftwareVersion", strconv.Itoa(minSV),
		"--from", opts.From,
	}

	if opts.OtaURL != "" {
		args = append(args, "--otaURL", opts.OtaURL)
	}
	if opts.OtaFileSize != 0 {
		args = append(args, "--otaFileSize", strconv.Itoa(opts.OtaFileSize))
	}
	if opts.OtaChecksum != "" {
		args = append(args, "--otaChecksum", opts.OtaChecksum)
	}
	if opts.OtaChecksumType != 0 {
		args = append(args, "--otaChecksumType", strconv.Itoa(opts.OtaChecksumType))
	}
	if opts.FirmwareInformation != "" {
		args = append(args, "--firmwareInformation", opts.FirmwareInformation)
	}
	if opts.SpecificationVersion != 0 {
		args = append(args, "--specificationVersion", strconv.Itoa(opts.SpecificationVersion))
	}
	if opts.ReleaseNotesURL != "" {
		args = append(args, "--releaseNotesURL", opts.ReleaseNotesURL)
	}
	if opts.SchemaVersion != "" {
		args = append(args, "--schemaVersion", opts.SchemaVersion)
	}

	return utils.ExecuteTx(args...)
}

// UpdateModelVersionOpts holds parameters for update-model-version. The base
// fields (VID/PID/SoftwareVersion/From) are always sent; the rest emit only when
// non-zero/non-empty. SoftwareVersionValid is a *bool so that both true and
// false can be sent explicitly (and omitted when nil). SpecificationVersion is
// retained only to drive the negative "unknown flag" CLI test — update-model-
// version does not accept it.
type UpdateModelVersionOpts struct {
	VID                          int
	PID                          int
	SoftwareVersion              int
	MinApplicableSoftwareVersion int
	MaxApplicableSoftwareVersion int
	SoftwareVersionValid         *bool
	OtaURL                       string
	OtaFileSize                  int
	OtaChecksum                  string
	OtaChecksumType              int
	ReleaseNotesURL              string
	SpecificationVersion         int
	SchemaVersion                string
	From                         string
}

// boolPtr returns a pointer to b, for setting optional *bool flags inline.
func boolPtr(b bool) *bool { return &b }

// UpdateModelVersion executes the update-model-version transaction.
func UpdateModelVersion(opts UpdateModelVersionOpts) (*utils.TxResult, error) {
	args := []string{
		"tx", "model", "update-model-version",
		"--vid", strconv.Itoa(opts.VID),
		"--pid", strconv.Itoa(opts.PID),
		"--softwareVersion", strconv.Itoa(opts.SoftwareVersion),
		"--from", opts.From,
	}
	if opts.MinApplicableSoftwareVersion != 0 {
		args = append(args, "--minApplicableSoftwareVersion", strconv.Itoa(opts.MinApplicableSoftwareVersion))
	}
	if opts.MaxApplicableSoftwareVersion != 0 {
		args = append(args, "--maxApplicableSoftwareVersion", strconv.Itoa(opts.MaxApplicableSoftwareVersion))
	}
	if opts.SoftwareVersionValid != nil {
		if *opts.SoftwareVersionValid {
			args = append(args, "--softwareVersionValid=true")
		} else {
			args = append(args, "--softwareVersionValid=false")
		}
	}
	if opts.OtaURL != "" {
		args = append(args, "--otaURL", opts.OtaURL)
	}
	if opts.OtaFileSize != 0 {
		args = append(args, "--otaFileSize", strconv.Itoa(opts.OtaFileSize))
	}
	if opts.OtaChecksum != "" {
		args = append(args, "--otaChecksum", opts.OtaChecksum)
	}
	if opts.OtaChecksumType != 0 {
		args = append(args, "--otaChecksumType", strconv.Itoa(opts.OtaChecksumType))
	}
	if opts.ReleaseNotesURL != "" {
		args = append(args, "--releaseNotesURL", opts.ReleaseNotesURL)
	}
	if opts.SpecificationVersion != 0 {
		args = append(args, "--specificationVersion", strconv.Itoa(opts.SpecificationVersion))
	}
	if opts.SchemaVersion != "" {
		args = append(args, "--schemaVersion", opts.SchemaVersion)
	}

	return utils.ExecuteTx(args...)
}

// DeleteModelVersion executes the delete-model-version transaction.
func DeleteModelVersion(vid, pid, sv int, from string) (*utils.TxResult, error) {
	return utils.ExecuteTx("tx", "model", "delete-model-version",
		"--vid", strconv.Itoa(vid),
		"--pid", strconv.Itoa(pid),
		"--softwareVersion", strconv.Itoa(sv),
		"--from", from,
	)
}

// GetModel queries a specific model by vid/pid. Returns nil when the model
// does not exist.
func GetModel(vid, pid int) (*modeltypes.Model, error) {
	return GetModelHex(strconv.Itoa(vid), strconv.Itoa(pid))
}

// GetModelHex queries a model using hex-format vid/pid strings.
func GetModelHex(vid, pid string) (*modeltypes.Model, error) {
	var res modeltypes.Model
	found, err := cliputils.GetSingle(&res,
		"query", "model", "get-model",
		"--vid", vid,
		"--pid", pid,
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetAllModels queries all models.
func GetAllModels() ([]modeltypes.Model, error) {
	var res modeltypes.QueryAllModelResponse
	if err := cliputils.GetList(&res, "query", "model", "all-models", "-o", "json"); err != nil {
		return nil, err
	}

	return res.Model, nil
}

// GetVendorModels queries all models for a given vendor. Returns nil when the
// vendor has no products on chain.
func GetVendorModels(vid int) (*modeltypes.VendorProducts, error) {
	return GetVendorModelsHex(strconv.Itoa(vid))
}

// GetVendorModelsHex queries vendor models using a hex vid string.
func GetVendorModelsHex(vid string) (*modeltypes.VendorProducts, error) {
	var res modeltypes.VendorProducts
	found, err := cliputils.GetSingle(&res,
		"query", "model", "vendor-models",
		"--vid", vid,
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetModelVersion queries a specific model version. Returns nil when the
// record does not exist.
func GetModelVersion(vid, pid, sv int) (*modeltypes.ModelVersion, error) {
	return GetModelVersionHex(strconv.Itoa(vid), strconv.Itoa(pid), sv)
}

// GetModelVersionHex queries a model version using hex vid/pid.
func GetModelVersionHex(vid, pid string, sv int) (*modeltypes.ModelVersion, error) {
	var res modeltypes.ModelVersion
	found, err := cliputils.GetSingle(&res,
		"query", "model", "model-version",
		"--vid", vid,
		"--pid", pid,
		"--softwareVersion", strconv.Itoa(sv),
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetAllModelVersions queries all model versions for a given vid/pid. Returns
// nil when the model has no versions on chain.
func GetAllModelVersions(vid, pid int) (*modeltypes.ModelVersions, error) {
	return GetAllModelVersionsHex(strconv.Itoa(vid), strconv.Itoa(pid))
}

// GetAllModelVersionsHex queries all model versions using hex vid/pid.
func GetAllModelVersionsHex(vid, pid string) (*modeltypes.ModelVersions, error) {
	var res modeltypes.ModelVersions
	found, err := cliputils.GetSingle(&res,
		"query", "model", "all-model-versions",
		"--vid", vid,
		"--pid", pid,
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// containsModelByPid reports whether list has a Model with the given (vid, pid).
func containsModelByPid(list []modeltypes.Model, vid, pid int32) bool {
	for i := range list {
		if list[i].Vid == vid && list[i].Pid == pid {
			return true
		}
	}

	return false
}

// containsProductByPid reports whether products has an entry with the given pid.
func containsProductByPid(products []*modeltypes.Product, pid int32) bool {
	for _, p := range products {
		if p != nil && p.Pid == pid {
			return true
		}
	}

	return false
}
