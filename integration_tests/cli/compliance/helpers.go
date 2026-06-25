package compliance

import (
	"strconv"

	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
	compliancetypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
)

// OptionalFields holds the optional Matter compliance fields shared by
// certify-model and provision-model. Each non-empty field emits its flag.
type OptionalFields struct {
	ProgramType                        string
	ProgramTypeVersion                 string
	FamilyID                           string
	SupportedClusters                  string
	CompliantPlatformUsed              string
	CompliantPlatformVersion           string
	OSVersion                          string
	CertificationRoute                 string
	Transport                          string
	ParentChild                        string
	CertificationIDOfSoftwareComponent string
}

func (o OptionalFields) args() []string {
	pairs := []struct{ flag, val string }{
		{"--programType", o.ProgramType},
		{"--programTypeVersion", o.ProgramTypeVersion},
		{"--familyId", o.FamilyID},
		{"--supportedClusters", o.SupportedClusters},
		{"--compliantPlatformUsed", o.CompliantPlatformUsed},
		{"--compliantPlatformVersion", o.CompliantPlatformVersion},
		{"--OSVersion", o.OSVersion},
		{"--certificationRoute", o.CertificationRoute},
		{"--transport", o.Transport},
		{"--parentChild", o.ParentChild},
		{"--certificationIDOfSoftwareComponent", o.CertificationIDOfSoftwareComponent},
	}
	var args []string
	for _, p := range pairs {
		if p.val != "" {
			args = append(args, p.flag, p.val)
		}
	}

	return args
}

// CertifyModelOpts holds parameters for the certify-model transaction.
// Use VIDHex / PIDHex to pass hex-formatted identifiers (e.g. "0xA13"); when
// those are empty, the numeric VID / PID fields are formatted as decimal.
// Zero-valued SpecificationVersion and CDVersionNumber default to 1 (pass a
// non-zero value to exercise a mismatch).
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
	Reason                string
	Optional              OptionalFields
	From                  string
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
		"--vid", cliputils.FlagOrHex(opts.VID, opts.VIDHex),
		"--pid", cliputils.FlagOrHex(opts.PID, opts.PIDHex),
		"--softwareVersion", strconv.Itoa(opts.SoftwareVersion),
		"--softwareVersionString", opts.SoftwareVersionString,
		"--certificationType", opts.CertificationType,
		"--specificationVersion", strconv.Itoa(specVersion),
		"--certificationDate", opts.CertificationDate,
		"--cdCertificateId", opts.CDCertificateID,
		"--cdVersionNumber", strconv.Itoa(cdVersion),
		"--from", opts.From,
	}
	if opts.Reason != "" {
		args = append(args, "--reason", opts.Reason)
	}
	args = append(args, opts.Optional.args()...)

	return utils.ExecuteTx(args...)
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
}

// RevokeModel executes the revoke-model transaction.
func RevokeModel(opts RevokeModelOpts) (*utils.TxResult, error) {
	cdVersion := opts.CDVersionNumber
	if cdVersion == 0 {
		cdVersion = 1
	}

	args := []string{
		"tx", "compliance", "revoke-model",
		"--vid", cliputils.FlagOrHex(opts.VID, opts.VIDHex),
		"--pid", cliputils.FlagOrHex(opts.PID, opts.PIDHex),
		"--softwareVersion", strconv.Itoa(opts.SoftwareVersion),
		"--softwareVersionString", opts.SoftwareVersionString,
		"--certificationType", opts.CertificationType,
		"--revocationDate", opts.RevocationDate,
		"--cdVersionNumber", strconv.Itoa(cdVersion),
		"--from", opts.From,
	}
	if opts.Reason != "" {
		args = append(args, "--reason", opts.Reason)
	}

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
	Optional              OptionalFields
	From                  string
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
		"--vid", cliputils.FlagOrHex(opts.VID, opts.VIDHex),
		"--pid", cliputils.FlagOrHex(opts.PID, opts.PIDHex),
		"--softwareVersion", strconv.Itoa(opts.SoftwareVersion),
		"--softwareVersionString", opts.SoftwareVersionString,
		"--certificationType", opts.CertificationType,
		"--specificationVersion", strconv.Itoa(specVersion),
		"--provisionalDate", opts.ProvisionalDate,
		"--cdCertificateId", opts.CDCertificateID,
		"--cdVersionNumber", strconv.Itoa(cdVersion),
		"--from", opts.From,
	}
	if opts.Reason != "" {
		args = append(args, "--reason", opts.Reason)
	}
	args = append(args, opts.Optional.args()...)

	return utils.ExecuteTx(args...)
}

// UpdateComplianceInfoOpts holds parameters for update-compliance-info. The
// base fields (VID/PID/SoftwareVersion/CertificationType/From) are always sent;
// the rest emit only when non-zero/non-empty.
type UpdateComplianceInfoOpts struct {
	VID               int
	PID               int
	SoftwareVersion   int
	CertificationType string
	CDCertificateID   string
	CDVersionNumber   int
	CertificationDate string
	Reason            string
	Optional          OptionalFields
	From              string
}

// UpdateComplianceInfo executes the update-compliance-info transaction.
func UpdateComplianceInfo(opts UpdateComplianceInfoOpts) (*utils.TxResult, error) {
	args := []string{
		"tx", "compliance", "update-compliance-info",
		"--vid", strconv.Itoa(opts.VID),
		"--pid", strconv.Itoa(opts.PID),
		"--softwareVersion", strconv.Itoa(opts.SoftwareVersion),
		"--certificationType", opts.CertificationType,
		"--from", opts.From,
	}
	if opts.CDCertificateID != "" {
		args = append(args, "--cdCertificateId", opts.CDCertificateID)
	}
	if opts.CDVersionNumber != 0 {
		args = append(args, "--cdVersionNumber", strconv.Itoa(opts.CDVersionNumber))
	}
	if opts.CertificationDate != "" {
		args = append(args, "--certificationDate", opts.CertificationDate)
	}
	if opts.Reason != "" {
		args = append(args, "--reason", opts.Reason)
	}
	args = append(args, opts.Optional.args()...)

	return utils.ExecuteTx(args...)
}

// DeleteComplianceInfo executes the delete-compliance-info transaction.
func DeleteComplianceInfo(vid, pid, sv int, certType, from string) (*utils.TxResult, error) {
	return utils.ExecuteTx("tx", "compliance", "delete-compliance-info",
		"--vid", strconv.Itoa(vid),
		"--pid", strconv.Itoa(pid),
		"--softwareVersion", strconv.Itoa(sv),
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
		"--vid", cliputils.FlagOrHex(o.VID, o.VIDHex),
		"--pid", cliputils.FlagOrHex(o.PID, o.PIDHex),
		"--softwareVersion", strconv.Itoa(o.SoftwareVersion),
		"--certificationType", o.CertificationType,
		"-o", "json",
	}
}

// GetComplianceInfo queries compliance-info for a given vid/pid/sv/certType.
// Returns nil when the record does not exist.
func GetComplianceInfo(opts ComplianceQueryOpts) (*compliancetypes.ComplianceInfo, error) {
	var res compliancetypes.ComplianceInfo
	found, err := cliputils.GetSingle(&res, append([]string{"query", "compliance", "compliance-info"}, opts.args()...)...)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetCertifiedModel queries the certified-model endpoint. Returns nil when the
// record does not exist.
func GetCertifiedModel(opts ComplianceQueryOpts) (*compliancetypes.CertifiedModel, error) {
	var res compliancetypes.CertifiedModel
	found, err := cliputils.GetSingle(&res, append([]string{"query", "compliance", "certified-model"}, opts.args()...)...)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetRevokedModel queries the revoked-model endpoint. Returns nil when the
// record does not exist.
func GetRevokedModel(opts ComplianceQueryOpts) (*compliancetypes.RevokedModel, error) {
	var res compliancetypes.RevokedModel
	found, err := cliputils.GetSingle(&res, append([]string{"query", "compliance", "revoked-model"}, opts.args()...)...)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetProvisionalModel queries the provisional-model endpoint. Returns nil when
// the record does not exist.
func GetProvisionalModel(opts ComplianceQueryOpts) (*compliancetypes.ProvisionalModel, error) {
	var res compliancetypes.ProvisionalModel
	found, err := cliputils.GetSingle(&res, append([]string{"query", "compliance", "provisional-model"}, opts.args()...)...)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetDeviceSoftwareCompliance queries device-software-compliance by
// CDCertificateID. Returns nil when the record does not exist.
func GetDeviceSoftwareCompliance(cdCertificateID string) (*compliancetypes.DeviceSoftwareCompliance, error) {
	var res compliancetypes.DeviceSoftwareCompliance
	found, err := cliputils.GetSingle(&res,
		"query", "compliance", "device-software-compliance",
		"--cdCertificateId", cdCertificateID,
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetAllComplianceInfo queries all compliance info records.
func GetAllComplianceInfo() ([]compliancetypes.ComplianceInfo, error) {
	var res compliancetypes.QueryAllComplianceInfoResponse
	if err := cliputils.GetList(&res, "query", "compliance", "all-compliance-info", "-o", "json"); err != nil {
		return nil, err
	}

	return res.ComplianceInfo, nil
}

// GetAllCertifiedModels queries all certified models.
func GetAllCertifiedModels() ([]compliancetypes.CertifiedModel, error) {
	var res compliancetypes.QueryAllCertifiedModelResponse
	if err := cliputils.GetList(&res, "query", "compliance", "all-certified-models", "-o", "json"); err != nil {
		return nil, err
	}

	return res.CertifiedModel, nil
}

// GetAllRevokedModels queries all revoked models.
func GetAllRevokedModels() ([]compliancetypes.RevokedModel, error) {
	var res compliancetypes.QueryAllRevokedModelResponse
	if err := cliputils.GetList(&res, "query", "compliance", "all-revoked-models", "-o", "json"); err != nil {
		return nil, err
	}

	return res.RevokedModel, nil
}

// GetAllProvisionalModels queries all provisional models.
func GetAllProvisionalModels() ([]compliancetypes.ProvisionalModel, error) {
	var res compliancetypes.QueryAllProvisionalModelResponse
	if err := cliputils.GetList(&res, "query", "compliance", "all-provisional-models", "-o", "json"); err != nil {
		return nil, err
	}

	return res.ProvisionalModel, nil
}

// GetAllDeviceSoftwareCompliance queries all device software compliance records.
func GetAllDeviceSoftwareCompliance() ([]compliancetypes.DeviceSoftwareCompliance, error) {
	var res compliancetypes.QueryAllDeviceSoftwareComplianceResponse
	if err := cliputils.GetList(&res, "query", "compliance", "all-device-software-compliance", "-o", "json"); err != nil {
		return nil, err
	}

	return res.DeviceSoftwareCompliance, nil
}

// containsComplianceInfo reports whether list has an entry matching (vid, pid).
func containsComplianceInfo(list []compliancetypes.ComplianceInfo, vid, pid int32) bool {
	for i := range list {
		if list[i].Vid == vid && list[i].Pid == pid {
			return true
		}
	}

	return false
}

// containsCertifiedModel reports whether list has an entry matching (vid, pid).
func containsCertifiedModel(list []compliancetypes.CertifiedModel, vid, pid int32) bool {
	for i := range list {
		if list[i].Vid == vid && list[i].Pid == pid {
			return true
		}
	}

	return false
}

// containsRevokedModel reports whether list has an entry matching (vid, pid).
func containsRevokedModel(list []compliancetypes.RevokedModel, vid, pid int32) bool {
	for i := range list {
		if list[i].Vid == vid && list[i].Pid == pid {
			return true
		}
	}

	return false
}

// containsProvisionalModel reports whether list has an entry matching (vid, pid).
func containsProvisionalModel(list []compliancetypes.ProvisionalModel, vid, pid int32) bool {
	for i := range list {
		if list[i].Vid == vid && list[i].Pid == pid {
			return true
		}
	}

	return false
}

// containsDeviceSoftwareCompliance reports whether list has an entry whose
// CDCertificateID matches cdCertID.
func containsDeviceSoftwareCompliance(list []compliancetypes.DeviceSoftwareCompliance, cdCertID string) bool {
	for i := range list {
		if list[i].CDCertificateId == cdCertID {
			return true
		}
	}

	return false
}

// containsComplianceInfoCertType reports whether list contains a record with
// the given (vid, pid, certType).
func containsComplianceInfoCertType(list []*compliancetypes.ComplianceInfo, vid, pid int32, certType string) bool {
	for _, ci := range list {
		if ci != nil && ci.Vid == vid && ci.Pid == pid && ci.CertificationType == certType {
			return true
		}
	}

	return false
}

// hasComplianceInfoCertType reports whether list contains a record matching
// (vid, pid, certType, date).
func hasComplianceInfoCertType(list []compliancetypes.ComplianceInfo, vid, pid int32, certType, date string) bool {
	for i := range list {
		m := &list[i]
		if m.Vid == vid && m.Pid == pid && m.CertificationType == certType && m.Date == date {
			return true
		}
	}

	return false
}

// hasCertifiedModelCertType reports whether list contains a record matching
// (vid, pid, certType).
func hasCertifiedModelCertType(list []compliancetypes.CertifiedModel, vid, pid int32, certType string) bool {
	for i := range list {
		m := &list[i]
		if m.Vid == vid && m.Pid == pid && m.CertificationType == certType {
			return true
		}
	}

	return false
}

// hasRevokedWithValue reports whether list contains a record matching
// (vid, pid) with the given Value.
func hasRevokedWithValue(list []compliancetypes.RevokedModel, vid, pid int32, value bool) bool {
	for i := range list {
		m := &list[i]
		if m.Vid == vid && m.Pid == pid && m.Value == value {
			return true
		}
	}

	return false
}

// hasProvisionalWithValue reports whether list contains a record matching
// (vid, pid) with the given Value.
func hasProvisionalWithValue(list []compliancetypes.ProvisionalModel, vid, pid int32, value bool) bool {
	for i := range list {
		m := &list[i]
		if m.Vid == vid && m.Pid == pid && m.Value == value {
			return true
		}
	}

	return false
}

// hasCertifiedTrueAt reports whether list contains a CertifiedModel matching
// (vid, pid, sv, certType) with Value=true.
func hasCertifiedTrueAt(list []compliancetypes.CertifiedModel, vid, pid int32, sv uint32, certType string) bool {
	for i := range list {
		m := &list[i]
		if m.Vid == vid && m.Pid == pid && m.SoftwareVersion == sv && m.CertificationType == certType && m.Value {
			return true
		}
	}

	return false
}

// hasComplianceInfoStatus reports whether list contains a record matching
// (vid, pid) with the given SoftwareVersionCertificationStatus.
func hasComplianceInfoStatus(list []compliancetypes.ComplianceInfo, vid, pid int32, status uint32) bool {
	for i := range list {
		m := &list[i]
		if m.Vid == vid && m.Pid == pid && m.SoftwareVersionCertificationStatus == status {
			return true
		}
	}

	return false
}
