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

package lightclientproxy

import (
	"strconv"

	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// Typed builders for the txs the light_client_proxy suite broadcasts via
// utils.ExecuteTx against the full node. Each Args struct carries only the
// fields the suite actually sets — fields beyond those are left at the dcld
// CLI's defaults. Send(from) appends --from and --node FullNodeAddr and
// invokes utils.ExecuteTx, returning the raw TxResult; callers pair with
// requireTxOK to assert success.
//
// Build returns the assembled tx args minus --from / --node so the
// Write_Rejected blocks can reuse the same struct to construct a tx that's
// then sent through the proxy via assertWriteRejected.

// ─── auth ──────────────────────────────────────────────────────────────────

// ProposeAddAccountArgs → `tx auth propose-add-account`.
type ProposeAddAccountArgs struct {
	Address string
	Pubkey  string
	Roles   string
	VID     int // 0 → omit --vid (non-Vendor roles)
}

func (a ProposeAddAccountArgs) Build() []string {
	args := []string{
		"tx", "auth", "propose-add-account",
		"--address", a.Address,
		"--pubkey", a.Pubkey,
		"--roles", a.Roles,
	}
	if a.VID > 0 {
		args = append(args, "--vid", strconv.Itoa(a.VID))
	}

	return args
}

func (a ProposeAddAccountArgs) Send(from string) (*utils.TxResult, error) {
	return runFullNodeTx(from, a.Build())
}

// ApproveAddAccountArgs → `tx auth approve-add-account`.
type ApproveAddAccountArgs struct {
	Address string
}

func (a ApproveAddAccountArgs) Build() []string {
	return []string{
		"tx", "auth", "approve-add-account",
		"--address", a.Address,
	}
}

func (a ApproveAddAccountArgs) Send(from string) (*utils.TxResult, error) {
	return runFullNodeTx(from, a.Build())
}

// ─── model ─────────────────────────────────────────────────────────────────

// AddModelArgs → `tx model add-model`. The suite always seeds the same
// fixture-style fields, so the unused-but-required CLI flags
// (--deviceTypeID, --productName, --partNumber, --commissioningCustomFlow,
// --enhancedSetupFlowOptions) are hard-coded in Build. ProductLabel is the
// only freely-set string; vid/pid are random.
type AddModelArgs struct {
	VID          int
	PID          int
	ProductLabel string
}

func (a AddModelArgs) Build() []string {
	return []string{
		"tx", "model", "add-model",
		"--vid", strconv.Itoa(a.VID),
		"--pid", strconv.Itoa(a.PID),
		"--deviceTypeID", "1",
		"--productName", "TestProduct",
		"--productLabel", a.ProductLabel,
		"--partNumber", "1",
		"--commissioningCustomFlow", "0",
		"--enhancedSetupFlowOptions", "0",
	}
}

func (a AddModelArgs) Send(from string) (*utils.TxResult, error) {
	return runFullNodeTx(from, a.Build())
}

// AddModelVersionArgs → `tx model add-model-version`. Same rationale as
// AddModelArgs — fixture-style fields hard-coded, only the keys vary.
type AddModelVersionArgs struct {
	VID                   int
	PID                   int
	SoftwareVersion       int
	SoftwareVersionString string
}

func (a AddModelVersionArgs) Build() []string {
	return []string{
		"tx", "model", "add-model-version",
		"--cdVersionNumber", "1",
		"--maxApplicableSoftwareVersion", "10",
		"--minApplicableSoftwareVersion", "1",
		"--vid", strconv.Itoa(a.VID),
		"--pid", strconv.Itoa(a.PID),
		"--softwareVersion", strconv.Itoa(a.SoftwareVersion),
		"--softwareVersionString", a.SoftwareVersionString,
	}
}

func (a AddModelVersionArgs) Send(from string) (*utils.TxResult, error) {
	return runFullNodeTx(from, a.Build())
}

// ─── compliance ────────────────────────────────────────────────────────────

// CertifyModelArgs → `tx compliance certify-model`.
type CertifyModelArgs struct {
	VID                   int
	PID                   int
	SoftwareVersion       int
	SoftwareVersionString string
	SpecificationVersion  int
	CertificationType     string
	CertificationDate     string
	CDCertificateID       string
}

func (a CertifyModelArgs) Build() []string {
	return []string{
		"tx", "compliance", "certify-model",
		"--vid", strconv.Itoa(a.VID),
		"--pid", strconv.Itoa(a.PID),
		"--softwareVersion", strconv.Itoa(a.SoftwareVersion),
		"--softwareVersionString", a.SoftwareVersionString,
		"--specificationVersion", strconv.Itoa(a.SpecificationVersion),
		"--certificationType", a.CertificationType,
		"--certificationDate", a.CertificationDate,
		"--cdCertificateId", a.CDCertificateID,
		"--cdVersionNumber", "1",
	}
}

func (a CertifyModelArgs) Send(from string) (*utils.TxResult, error) {
	return runFullNodeTx(from, a.Build())
}

// ─── vendorinfo ────────────────────────────────────────────────────────────

// AddVendorArgs → `tx vendorinfo add-vendor`.
type AddVendorArgs struct {
	VID              int
	CompanyLegalName string
	VendorName       string
}

func (a AddVendorArgs) Build() []string {
	return []string{
		"tx", "vendorinfo", "add-vendor",
		"--vid", strconv.Itoa(a.VID),
		"--companyLegalName", a.CompanyLegalName,
		"--vendorName", a.VendorName,
	}
}

func (a AddVendorArgs) Send(from string) (*utils.TxResult, error) {
	return runFullNodeTx(from, a.Build())
}

// ─── pki ───────────────────────────────────────────────────────────────────

// ProposeAddX509RootCertArgs → `tx pki propose-add-x509-root-cert`.
type ProposeAddX509RootCertArgs struct {
	Certificate string
	VID         int // 0 → omit --vid
}

func (a ProposeAddX509RootCertArgs) Build() []string {
	args := []string{
		"tx", "pki", "propose-add-x509-root-cert",
		"--certificate", a.Certificate,
	}
	if a.VID > 0 {
		args = append(args, "--vid", strconv.Itoa(a.VID))
	}

	return args
}

func (a ProposeAddX509RootCertArgs) Send(from string) (*utils.TxResult, error) {
	return runFullNodeTx(from, a.Build())
}

// CertRefArgs identifies a root cert by (subject, subject-key-id) — used by
// every approve / propose-revoke flow.
type CertRefArgs struct {
	Subject      string
	SubjectKeyID string
}

func (a CertRefArgs) flags() []string {
	return []string{"--subject", a.Subject, "--subject-key-id", a.SubjectKeyID}
}

// ApproveAddX509RootCertArgs → `tx pki approve-add-x509-root-cert`.
type ApproveAddX509RootCertArgs struct {
	CertRefArgs
}

func (a ApproveAddX509RootCertArgs) Build() []string {
	return append([]string{"tx", "pki", "approve-add-x509-root-cert"}, a.flags()...)
}

func (a ApproveAddX509RootCertArgs) Send(from string) (*utils.TxResult, error) {
	return runFullNodeTx(from, a.Build())
}

// AddX509CertArgs → `tx pki add-x509-cert`.
type AddX509CertArgs struct {
	Certificate string
}

func (a AddX509CertArgs) Build() []string {
	return []string{
		"tx", "pki", "add-x509-cert",
		"--certificate", a.Certificate,
	}
}

func (a AddX509CertArgs) Send(from string) (*utils.TxResult, error) {
	return runFullNodeTx(from, a.Build())
}

// RevokeX509CertArgs → `tx pki revoke-x509-cert`.
type RevokeX509CertArgs struct {
	CertRefArgs
}

func (a RevokeX509CertArgs) Build() []string {
	return append([]string{"tx", "pki", "revoke-x509-cert"}, a.flags()...)
}

func (a RevokeX509CertArgs) Send(from string) (*utils.TxResult, error) {
	return runFullNodeTx(from, a.Build())
}

// ProposeRevokeX509RootCertArgs → `tx pki propose-revoke-x509-root-cert`.
type ProposeRevokeX509RootCertArgs struct {
	CertRefArgs
}

func (a ProposeRevokeX509RootCertArgs) Build() []string {
	return append([]string{"tx", "pki", "propose-revoke-x509-root-cert"}, a.flags()...)
}

func (a ProposeRevokeX509RootCertArgs) Send(from string) (*utils.TxResult, error) {
	return runFullNodeTx(from, a.Build())
}

// ─── internal ──────────────────────────────────────────────────────────────

// runFullNodeTx appends --from + --node FullNodeAddr and runs the tx via
// utils.ExecuteTx. Centralized so every Send method routes through the same
// node + signer convention.
func runFullNodeTx(from string, args []string) (*utils.TxResult, error) {
	args = append(args, "--from", from, "--node", FullNodeAddr)

	return utils.ExecuteTx(args...)
}
