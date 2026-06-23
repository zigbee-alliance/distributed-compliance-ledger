// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package upgrade

import (
	"strconv"

	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// Typed builders for the module txs the migration tests seed (model /
// vendorinfo / compliance). Each runs against an explicit binPath (a specific
// dcld release) via ExecuteTxWithBin and emits a flag ONLY when the
// corresponding field is set. The base fields (always sent) match the v0.12
// flag set common to every release; the optional fields cover the flags later
// releases added — so a call emits exactly the flags it populates, identical to
// the raw arg slice it replaces, on whichever binary version it targets.
//
// Numeric optionals emit when non-zero and string optionals when non-empty.
// A zero/empty value therefore omits the flag (equivalent to the CLI default);
// no migration test sends these optionals as an explicit zero.

// intPtr returns a pointer to n, for optional numeric flags whose explicit zero
// value is meaningful (e.g. --commissioningCustomFlow 0) and so must be emitted
// rather than treated as "unset".
func intPtr(n int) *int { return &n }

// AddModelArgs → `tx model add-model`.
type AddModelArgs struct {
	VID          int
	PID          int
	DeviceTypeID int
	ProductName  string
	ProductLabel string
	PartNumber   string

	// Optional flags added by later releases. CommissioningCustomFlow is a
	// pointer because v1.6+/master send it explicitly as 0, which must still be
	// emitted (a nil pointer means "older binary, omit the flag").
	CommissioningCustomFlow             *int
	DiscoveryCapabilitiesBitmask        int
	CommissioningModeSecondaryStepsHint int
	IcdUserActiveModeTriggerHint        int
	IcdUserActiveModeTriggerInstruction string
	FactoryResetStepsHint               int
	FactoryResetStepsInstruction        string

	From string
}

// AddModel executes `tx model add-model` with the binary at binPath.
func AddModel(binPath string, a AddModelArgs) (*utils.TxResult, error) {
	args := []string{
		"tx", "model", "add-model",
		"--vid", strconv.Itoa(a.VID),
		"--pid", strconv.Itoa(a.PID),
		"--deviceTypeID", strconv.Itoa(a.DeviceTypeID),
		"--productName", a.ProductName,
		"--productLabel", a.ProductLabel,
		"--partNumber", a.PartNumber,
	}
	if a.CommissioningCustomFlow != nil {
		args = append(args, "--commissioningCustomFlow", strconv.Itoa(*a.CommissioningCustomFlow))
	}
	if a.DiscoveryCapabilitiesBitmask != 0 {
		args = append(args, "--discoveryCapabilitiesBitmask", strconv.Itoa(a.DiscoveryCapabilitiesBitmask))
	}
	if a.CommissioningModeSecondaryStepsHint != 0 {
		args = append(args, "--commissioningModeSecondaryStepsHint", strconv.Itoa(a.CommissioningModeSecondaryStepsHint))
	}
	if a.IcdUserActiveModeTriggerHint != 0 {
		args = append(args, "--icdUserActiveModeTriggerHint", strconv.Itoa(a.IcdUserActiveModeTriggerHint))
	}
	if a.IcdUserActiveModeTriggerInstruction != "" {
		args = append(args, "--icdUserActiveModeTriggerInstruction", a.IcdUserActiveModeTriggerInstruction)
	}
	if a.FactoryResetStepsHint != 0 {
		args = append(args, "--factoryResetStepsHint", strconv.Itoa(a.FactoryResetStepsHint))
	}
	if a.FactoryResetStepsInstruction != "" {
		args = append(args, "--factoryResetStepsInstruction", a.FactoryResetStepsInstruction)
	}
	args = append(args, "--from", a.From)

	return ExecuteTxWithBin(binPath, args...)
}

// AddModelVersionArgs → `tx model add-model-version`.
type AddModelVersionArgs struct {
	VID                          int
	PID                          int
	SoftwareVersion              int
	SoftwareVersionString        string
	CDVersionNumber              int
	MinApplicableSoftwareVersion int
	MaxApplicableSoftwareVersion int
	SpecificationVersion         int // optional (added in v1.5)
	From                         string
}

// AddModelVersion executes `tx model add-model-version` with the binary at binPath.
func AddModelVersion(binPath string, a AddModelVersionArgs) (*utils.TxResult, error) {
	args := []string{
		"tx", "model", "add-model-version",
		"--vid", strconv.Itoa(a.VID),
		"--pid", strconv.Itoa(a.PID),
		"--softwareVersion", strconv.Itoa(a.SoftwareVersion),
		"--softwareVersionString", a.SoftwareVersionString,
		"--cdVersionNumber", strconv.Itoa(a.CDVersionNumber),
		"--minApplicableSoftwareVersion", strconv.Itoa(a.MinApplicableSoftwareVersion),
		"--maxApplicableSoftwareVersion", strconv.Itoa(a.MaxApplicableSoftwareVersion),
	}
	if a.SpecificationVersion != 0 {
		args = append(args, "--specificationVersion", strconv.Itoa(a.SpecificationVersion))
	}
	args = append(args, "--from", a.From)

	return ExecuteTxWithBin(binPath, args...)
}

// UpdateModelArgs → `tx model update-model`.
type UpdateModelArgs struct {
	VID          int
	PID          int
	ProductName  string
	ProductLabel string
	PartNumber   string
	From         string
}

// UpdateModel executes `tx model update-model` with the binary at binPath.
func UpdateModel(binPath string, a UpdateModelArgs) (*utils.TxResult, error) {
	return ExecuteTxWithBin(binPath,
		"tx", "model", "update-model",
		"--vid", strconv.Itoa(a.VID),
		"--pid", strconv.Itoa(a.PID),
		"--productName", a.ProductName,
		"--productLabel", a.ProductLabel,
		"--partNumber", a.PartNumber,
		"--from", a.From,
	)
}

// UpdateModelVersionArgs → `tx model update-model-version`.
type UpdateModelVersionArgs struct {
	VID                          int
	PID                          int
	SoftwareVersion              int
	MinApplicableSoftwareVersion int
	MaxApplicableSoftwareVersion int
	From                         string
}

// UpdateModelVersion executes `tx model update-model-version` with the binary at binPath.
func UpdateModelVersion(binPath string, a UpdateModelVersionArgs) (*utils.TxResult, error) {
	return ExecuteTxWithBin(binPath,
		"tx", "model", "update-model-version",
		"--vid", strconv.Itoa(a.VID),
		"--pid", strconv.Itoa(a.PID),
		"--softwareVersion", strconv.Itoa(a.SoftwareVersion),
		"--minApplicableSoftwareVersion", strconv.Itoa(a.MinApplicableSoftwareVersion),
		"--maxApplicableSoftwareVersion", strconv.Itoa(a.MaxApplicableSoftwareVersion),
		"--from", a.From,
	)
}

// DeleteModel executes `tx model delete-model` with the binary at binPath.
func DeleteModel(binPath string, vid, pid int, from string) (*utils.TxResult, error) {
	return ExecuteTxWithBin(binPath,
		"tx", "model", "delete-model",
		"--vid", strconv.Itoa(vid),
		"--pid", strconv.Itoa(pid),
		"--from", from,
	)
}

// DeleteModelVersion executes `tx model delete-model-version` with the binary at binPath.
func DeleteModelVersion(binPath string, vid, pid, softwareVersion int, from string) (*utils.TxResult, error) {
	return ExecuteTxWithBin(binPath,
		"tx", "model", "delete-model-version",
		"--vid", strconv.Itoa(vid),
		"--pid", strconv.Itoa(pid),
		"--softwareVersion", strconv.Itoa(softwareVersion),
		"--from", from,
	)
}

// CertifyModelArgs → `tx compliance certify-model`.
type CertifyModelArgs struct {
	VID                   int
	PID                   int
	SoftwareVersion       int
	SoftwareVersionString string
	CertificationType     string
	CertificationDate     string
	CDCertificateID       string
	CDVersionNumber       int    // optional
	SpecificationVersion  int    // optional (schema-v1)
	SchemaVersion         string // optional (schema-v1)
	From                  string
}

// CertifyModel executes `tx compliance certify-model` with the binary at binPath.
func CertifyModel(binPath string, a CertifyModelArgs) (*utils.TxResult, error) {
	args := []string{
		"tx", "compliance", "certify-model",
		"--vid", strconv.Itoa(a.VID),
		"--pid", strconv.Itoa(a.PID),
		"--softwareVersion", strconv.Itoa(a.SoftwareVersion),
		"--softwareVersionString", a.SoftwareVersionString,
		"--certificationType", a.CertificationType,
		"--certificationDate", a.CertificationDate,
		"--cdCertificateId", a.CDCertificateID,
	}
	if a.CDVersionNumber != 0 {
		args = append(args, "--cdVersionNumber", strconv.Itoa(a.CDVersionNumber))
	}
	if a.SpecificationVersion != 0 {
		args = append(args, "--specificationVersion", strconv.Itoa(a.SpecificationVersion))
	}
	if a.SchemaVersion != "" {
		args = append(args, "--schemaVersion", a.SchemaVersion)
	}
	args = append(args, "--from", a.From)

	return ExecuteTxWithBin(binPath, args...)
}

// ProvisionModelArgs → `tx compliance provision-model`.
type ProvisionModelArgs struct {
	VID                   int
	PID                   int
	SoftwareVersion       int
	SoftwareVersionString string
	CertificationType     string
	ProvisionalDate       string
	CDCertificateID       string
	CDVersionNumber       int    // optional
	SpecificationVersion  int    // optional (schema-v1)
	SchemaVersion         string // optional (schema-v1)
	From                  string
}

// ProvisionModel executes `tx compliance provision-model` with the binary at binPath.
func ProvisionModel(binPath string, a ProvisionModelArgs) (*utils.TxResult, error) {
	args := []string{
		"tx", "compliance", "provision-model",
		"--vid", strconv.Itoa(a.VID),
		"--pid", strconv.Itoa(a.PID),
		"--softwareVersion", strconv.Itoa(a.SoftwareVersion),
		"--softwareVersionString", a.SoftwareVersionString,
		"--certificationType", a.CertificationType,
		"--provisionalDate", a.ProvisionalDate,
		"--cdCertificateId", a.CDCertificateID,
	}
	if a.CDVersionNumber != 0 {
		args = append(args, "--cdVersionNumber", strconv.Itoa(a.CDVersionNumber))
	}
	if a.SpecificationVersion != 0 {
		args = append(args, "--specificationVersion", strconv.Itoa(a.SpecificationVersion))
	}
	if a.SchemaVersion != "" {
		args = append(args, "--schemaVersion", a.SchemaVersion)
	}
	args = append(args, "--from", a.From)

	return ExecuteTxWithBin(binPath, args...)
}

// RevokeModelArgs → `tx compliance revoke-model`.
type RevokeModelArgs struct {
	VID                   int
	PID                   int
	SoftwareVersion       int
	SoftwareVersionString string
	CertificationType     string
	RevocationDate        string
	CDVersionNumber       int // optional
	From                  string
}

// RevokeModel executes `tx compliance revoke-model` with the binary at binPath.
func RevokeModel(binPath string, a RevokeModelArgs) (*utils.TxResult, error) {
	args := []string{
		"tx", "compliance", "revoke-model",
		"--vid", strconv.Itoa(a.VID),
		"--pid", strconv.Itoa(a.PID),
		"--softwareVersion", strconv.Itoa(a.SoftwareVersion),
		"--softwareVersionString", a.SoftwareVersionString,
		"--certificationType", a.CertificationType,
		"--revocationDate", a.RevocationDate,
	}
	if a.CDVersionNumber != 0 {
		args = append(args, "--cdVersionNumber", strconv.Itoa(a.CDVersionNumber))
	}
	args = append(args, "--from", a.From)

	return ExecuteTxWithBin(binPath, args...)
}

// VendorArgs → `tx vendorinfo add-vendor` / `update-vendor` (same flag set).
type VendorArgs struct {
	VID                  int
	VendorName           string
	CompanyLegalName     string
	CompanyPreferredName string
	VendorLandingPageURL string
	From                 string
}

func (a VendorArgs) args(subcmd string) []string {
	return []string{
		"tx", "vendorinfo", subcmd,
		"--vid", strconv.Itoa(a.VID),
		"--vendorName", a.VendorName,
		"--companyLegalName", a.CompanyLegalName,
		"--companyPreferredName", a.CompanyPreferredName,
		"--vendorLandingPageURL", a.VendorLandingPageURL,
		"--from", a.From,
	}
}

// AddVendor executes `tx vendorinfo add-vendor` with the binary at binPath.
func AddVendor(binPath string, a VendorArgs) (*utils.TxResult, error) {
	return ExecuteTxWithBin(binPath, a.args("add-vendor")...)
}

// UpdateVendor executes `tx vendorinfo update-vendor` with the binary at binPath.
func UpdateVendor(binPath string, a VendorArgs) (*utils.TxResult, error) {
	return ExecuteTxWithBin(binPath, a.args("update-vendor")...)
}

// --- PKI: x509 root certificate lifecycle ------------------------------------

// ProposeAddX509RootCert executes `tx pki propose-add-x509-root-cert`. vid is
// the VID flag token (string and int constants are used across versions),
// emitted only when non-empty.
func ProposeAddX509RootCert(binPath, certPath, vid, from string) (*utils.TxResult, error) {
	args := []string{"tx", "pki", "propose-add-x509-root-cert", "--certificate", certPath}
	if vid != "" {
		args = append(args, "--vid", vid)
	}
	args = append(args, "--from", from)

	return ExecuteTxWithBin(binPath, args...)
}

// ApproveAddX509RootCert executes `tx pki approve-add-x509-root-cert`.
func ApproveAddX509RootCert(binPath, subject, subjectKeyID, from string) (*utils.TxResult, error) {
	return ExecuteTxWithBin(binPath, "tx", "pki", "approve-add-x509-root-cert",
		"--subject", subject, "--subject-key-id", subjectKeyID, "--from", from)
}

// RejectAddX509RootCert executes `tx pki reject-add-x509-root-cert`.
func RejectAddX509RootCert(binPath, subject, subjectKeyID, from string) (*utils.TxResult, error) {
	return ExecuteTxWithBin(binPath, "tx", "pki", "reject-add-x509-root-cert",
		"--subject", subject, "--subject-key-id", subjectKeyID, "--from", from)
}

// ProposeRevokeX509RootCert executes `tx pki propose-revoke-x509-root-cert`.
func ProposeRevokeX509RootCert(binPath, subject, subjectKeyID, from string) (*utils.TxResult, error) {
	return ExecuteTxWithBin(binPath, "tx", "pki", "propose-revoke-x509-root-cert",
		"--subject", subject, "--subject-key-id", subjectKeyID, "--from", from)
}

// ApproveRevokeX509RootCert executes `tx pki approve-revoke-x509-root-cert`.
func ApproveRevokeX509RootCert(binPath, subject, subjectKeyID, from string) (*utils.TxResult, error) {
	return ExecuteTxWithBin(binPath, "tx", "pki", "approve-revoke-x509-root-cert",
		"--subject", subject, "--subject-key-id", subjectKeyID, "--from", from)
}

// AddX509Cert executes `tx pki add-x509-cert`.
func AddX509Cert(binPath, certPath, from string) (*utils.TxResult, error) {
	return ExecuteTxWithBin(binPath, "tx", "pki", "add-x509-cert", "--certificate", certPath, "--from", from)
}

// RevokeX509Cert executes `tx pki revoke-x509-cert`. serialNumber is emitted
// only when non-empty (revoke a single cert by serial vs the whole subject).
func RevokeX509Cert(binPath, subject, subjectKeyID, serialNumber, from string) (*utils.TxResult, error) {
	args := []string{"tx", "pki", "revoke-x509-cert", "--subject", subject, "--subject-key-id", subjectKeyID}
	if serialNumber != "" {
		args = append(args, "--serial-number", serialNumber)
	}
	args = append(args, "--from", from)

	return ExecuteTxWithBin(binPath, args...)
}

// AssignVid executes `tx pki assign-vid`. vid is the VID flag token as-is.
func AssignVid(binPath, subject, subjectKeyID, vid, from string) (*utils.TxResult, error) {
	return ExecuteTxWithBin(binPath, "tx", "pki", "assign-vid",
		"--subject", subject, "--subject-key-id", subjectKeyID, "--vid", vid, "--from", from)
}

// --- PKI: NOC certificates ---------------------------------------------------

// AddNocX509RootCert executes `tx pki add-noc-x509-root-cert`.
func AddNocX509RootCert(binPath, certPath, from string) (*utils.TxResult, error) {
	return ExecuteTxWithBin(binPath, "tx", "pki", "add-noc-x509-root-cert", "--certificate", certPath, "--from", from)
}

// AddNocX509IcaCert executes `tx pki add-noc-x509-ica-cert`.
func AddNocX509IcaCert(binPath, certPath, from string) (*utils.TxResult, error) {
	return ExecuteTxWithBin(binPath, "tx", "pki", "add-noc-x509-ica-cert", "--certificate", certPath, "--from", from)
}

// RevokeNocX509RootCert executes `tx pki revoke-noc-x509-root-cert`.
func RevokeNocX509RootCert(binPath, subject, subjectKeyID, from string) (*utils.TxResult, error) {
	return ExecuteTxWithBin(binPath, "tx", "pki", "revoke-noc-x509-root-cert",
		"--subject", subject, "--subject-key-id", subjectKeyID, "--from", from)
}

// RevokeNocX509IcaCert executes `tx pki revoke-noc-x509-ica-cert`.
func RevokeNocX509IcaCert(binPath, subject, subjectKeyID, from string) (*utils.TxResult, error) {
	return ExecuteTxWithBin(binPath, "tx", "pki", "revoke-noc-x509-ica-cert",
		"--subject", subject, "--subject-key-id", subjectKeyID, "--from", from)
}

// RemoveNocX509RootCert executes `tx pki remove-noc-x509-root-cert`.
func RemoveNocX509RootCert(binPath, subject, subjectKeyID, from string) (*utils.TxResult, error) {
	return ExecuteTxWithBin(binPath, "tx", "pki", "remove-noc-x509-root-cert",
		"--subject", subject, "--subject-key-id", subjectKeyID, "--from", from)
}

// RemoveNocX509IcaCert executes `tx pki remove-noc-x509-ica-cert`.
func RemoveNocX509IcaCert(binPath, subject, subjectKeyID, from string) (*utils.TxResult, error) {
	return ExecuteTxWithBin(binPath, "tx", "pki", "remove-noc-x509-ica-cert",
		"--subject", subject, "--subject-key-id", subjectKeyID, "--from", from)
}

// --- PKI: revocation distribution points -------------------------------------

// AddRevocationPointArgs → `tx pki add-revocation-point`. CertificateDelegator
// is emitted only when non-empty (delegated-by-PAI points).
type AddRevocationPointArgs struct {
	VID                  int
	RevocationType       string
	IsPAA                bool
	Certificate          string
	CertificateDelegator string
	Label                string
	DataURL              string
	IssuerSubjectKeyID   string
	From                 string
}

// AddRevocationPoint executes `tx pki add-revocation-point` with the binary at binPath.
func AddRevocationPoint(binPath string, a AddRevocationPointArgs) (*utils.TxResult, error) {
	args := []string{
		"tx", "pki", "add-revocation-point",
		"--vid", strconv.Itoa(a.VID),
		"--revocation-type", a.RevocationType,
		"--is-paa=" + strconv.FormatBool(a.IsPAA),
		"--certificate", a.Certificate,
	}
	if a.CertificateDelegator != "" {
		args = append(args, "--certificate-delegator", a.CertificateDelegator)
	}
	args = append(args,
		"--label", a.Label,
		"--data-url", a.DataURL,
		"--issuer-subject-key-id", a.IssuerSubjectKeyID,
		"--from", a.From,
	)

	return ExecuteTxWithBin(binPath, args...)
}

// UpdateRevocationPointArgs → `tx pki update-revocation-point`.
type UpdateRevocationPointArgs struct {
	VID                  int
	Certificate          string
	CertificateDelegator string
	Label                string
	DataURL              string
	IssuerSubjectKeyID   string
	From                 string
}

// UpdateRevocationPoint executes `tx pki update-revocation-point` with the binary at binPath.
func UpdateRevocationPoint(binPath string, a UpdateRevocationPointArgs) (*utils.TxResult, error) {
	args := []string{
		"tx", "pki", "update-revocation-point",
		"--vid", strconv.Itoa(a.VID),
		"--certificate", a.Certificate,
	}
	if a.CertificateDelegator != "" {
		args = append(args, "--certificate-delegator", a.CertificateDelegator)
	}
	args = append(args,
		"--label", a.Label,
		"--data-url", a.DataURL,
		"--issuer-subject-key-id", a.IssuerSubjectKeyID,
		"--from", a.From,
	)

	return ExecuteTxWithBin(binPath, args...)
}

// DeleteRevocationPoint executes `tx pki delete-revocation-point`.
func DeleteRevocationPoint(binPath string, vid int, label, issuerSubjectKeyID, from string) (*utils.TxResult, error) {
	return ExecuteTxWithBin(binPath, "tx", "pki", "delete-revocation-point",
		"--vid", strconv.Itoa(vid), "--label", label, "--issuer-subject-key-id", issuerSubjectKeyID, "--from", from)
}
