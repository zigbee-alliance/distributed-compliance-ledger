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

import "strconv"

// Typed builders for the read-only queries the light_client_proxy suite runs
// against the proxy. Each Build returns the assembled `dcld query ...` args so
// the caller can dispatch via queryWithRetry / queryUntilContains as needed —
// the proxy's cold-start retry policy lives in queryWithRetry, so the builders
// stay transport-agnostic.
//
// Single-record queries return Not Found before the corresponding write lands
// on chain; list (`all-*`) queries are refused by the proxy. The Module field
// is included where helpful so a builder generalizes across modules with
// identical flag shapes.

// ─── auth ──────────────────────────────────────────────────────────────────

// AuthByAddress → `query auth <cmd> --address <addr>`. Used for account /
// proposed-account / proposed-account-to-revoke.
func AuthByAddress(cmd, address string) []string {
	return []string{"query", "auth", cmd, "--address", address}
}

// ─── compliance ────────────────────────────────────────────────────────────

// ComplianceByKeys → `query compliance <cmd> --vid V --pid P --softwareVersion SV
// --certificationType CT`. Used for compliance-info / certified-model /
// revoked-model / provisional-model.
func ComplianceByKeys(cmd string, vid, pid, sv int, certType string) []string {
	return []string{
		"query", "compliance", cmd,
		"--vid", strconv.Itoa(vid),
		"--pid", strconv.Itoa(pid),
		"--softwareVersion", strconv.Itoa(sv),
		"--certificationType", certType,
	}
}

// ─── model ─────────────────────────────────────────────────────────────────

// GetModel → `query model get-model --vid V --pid P`.
func GetModel(vid, pid int) []string {
	return []string{
		"query", "model", "get-model",
		"--vid", strconv.Itoa(vid),
		"--pid", strconv.Itoa(pid),
	}
}

// VendorModels → `query model vendor-models --vid V`.
func VendorModels(vid int) []string {
	return []string{"query", "model", "vendor-models", "--vid", strconv.Itoa(vid)}
}

// ModelVersion → `query model model-version --vid V --pid P --softwareVersion SV`.
func ModelVersion(vid, pid, sv int) []string {
	return []string{
		"query", "model", "model-version",
		"--vid", strconv.Itoa(vid),
		"--pid", strconv.Itoa(pid),
		"--softwareVersion", strconv.Itoa(sv),
	}
}

// AllModelVersions → `query model all-model-versions --vid V --pid P`. The
// proxy serves this when the (vid, pid) is known — returns Not Found
// otherwise. Despite the `all-` prefix this is not a list query, so the
// proxy does NOT refuse it.
func AllModelVersions(vid, pid int) []string {
	return []string{
		"query", "model", "all-model-versions",
		"--vid", strconv.Itoa(vid),
		"--pid", strconv.Itoa(pid),
	}
}

// ─── vendorinfo ────────────────────────────────────────────────────────────

// Vendor → `query vendorinfo vendor --vid V`.
func Vendor(vid int) []string {
	return []string{"query", "vendorinfo", "vendor", "--vid", strconv.Itoa(vid)}
}

// ─── pki ───────────────────────────────────────────────────────────────────

// PkiBySubjectAndSKID → `query pki <cmd> --subject S --subject-key-id SKID`.
// Covers x509-cert, revoked-x509-cert, proposed-x509-root-cert,
// proposed-x509-root-cert-to-revoke, and all-child-x509-certs.
func PkiBySubjectAndSKID(cmd, subject, skid string) []string {
	return []string{
		"query", "pki", cmd,
		"--subject", subject,
		"--subject-key-id", skid,
	}
}

// PkiBySubject → `query pki <cmd> --subject S`. Used for all-subject-x509-certs.
func PkiBySubject(cmd, subject string) []string {
	return []string{"query", "pki", cmd, "--subject", subject}
}

// PkiNoArgs → `query pki <cmd>`. Used for all-x509-root-certs /
// all-revoked-x509-root-certs which take no flags.
func PkiNoArgs(cmd string) []string {
	return []string{"query", "pki", cmd}
}
