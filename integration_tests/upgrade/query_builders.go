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

import "strconv"

// Typed builders for the read-only queries the migration tests run against a
// specific dcld release. Each wraps ExecuteCLIWithBin and returns the raw CLI
// output for the caller to assert on. Query keys (vid/pid/subject/...) are
// stable across releases, so these are version-agnostic.

// QueryAllAccounts runs `query auth all-accounts`.
func QueryAllAccounts(binPath string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "auth", "all-accounts")
}

// QueryAllProposedAccounts runs `query auth all-proposed-accounts`.
func QueryAllProposedAccounts(binPath string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "auth", "all-proposed-accounts")
}

// QueryAllProposedAccountsToRevoke runs `query auth all-proposed-accounts-to-revoke`.
func QueryAllProposedAccountsToRevoke(binPath string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "auth", "all-proposed-accounts-to-revoke")
}

// QueryAllRevokedAccounts runs `query auth all-revoked-accounts`.
func QueryAllRevokedAccounts(binPath string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "auth", "all-revoked-accounts")
}

// QueryAllCertifiedModels runs `query compliance all-certified-models`.
func QueryAllCertifiedModels(binPath string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "compliance", "all-certified-models")
}

// QueryAllComplianceInfo runs `query compliance all-compliance-info`.
func QueryAllComplianceInfo(binPath string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "compliance", "all-compliance-info")
}

// QueryAllDeviceSoftwareCompliance runs `query compliance all-device-software-compliance`.
func QueryAllDeviceSoftwareCompliance(binPath string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "compliance", "all-device-software-compliance")
}

// QueryAllProvisionalModels runs `query compliance all-provisional-models`.
func QueryAllProvisionalModels(binPath string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "compliance", "all-provisional-models")
}

// QueryAllRevokedModels runs `query compliance all-revoked-models`.
func QueryAllRevokedModels(binPath string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "compliance", "all-revoked-models")
}

// QueryAllModels runs `query model all-models`.
func QueryAllModels(binPath string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "model", "all-models")
}

// QueryAllCerts runs `query pki all-certs`.
func QueryAllCerts(binPath string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "pki", "all-certs")
}

// QueryAllNocX509Certs runs `query pki all-noc-x509-certs`.
func QueryAllNocX509Certs(binPath string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "pki", "all-noc-x509-certs")
}

// QueryAllProposedX509RootCerts runs `query pki all-proposed-x509-root-certs`.
func QueryAllProposedX509RootCerts(binPath string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "pki", "all-proposed-x509-root-certs")
}

// QueryAllProposedX509RootCertsToRevoke runs `query pki all-proposed-x509-root-certs-to-revoke`.
func QueryAllProposedX509RootCertsToRevoke(binPath string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "pki", "all-proposed-x509-root-certs-to-revoke")
}

// QueryAllRevocationPoints runs `query pki all-revocation-points`.
func QueryAllRevocationPoints(binPath string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "pki", "all-revocation-points")
}

// QueryAllRevokedNocX509IcaCerts runs `query pki all-revoked-noc-x509-ica-certs`.
func QueryAllRevokedNocX509IcaCerts(binPath string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "pki", "all-revoked-noc-x509-ica-certs")
}

// QueryAllRevokedNocX509RootCerts runs `query pki all-revoked-noc-x509-root-certs`.
func QueryAllRevokedNocX509RootCerts(binPath string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "pki", "all-revoked-noc-x509-root-certs")
}

// QueryAllRevokedX509Certs runs `query pki all-revoked-x509-certs`.
func QueryAllRevokedX509Certs(binPath string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "pki", "all-revoked-x509-certs")
}

// QueryAllRevokedX509RootCerts runs `query pki all-revoked-x509-root-certs`.
func QueryAllRevokedX509RootCerts(binPath string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "pki", "all-revoked-x509-root-certs")
}

// QueryAllX509Certs runs `query pki all-x509-certs`.
func QueryAllX509Certs(binPath string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "pki", "all-x509-certs")
}

// QueryAllX509RootCerts runs `query pki all-x509-root-certs`.
func QueryAllX509RootCerts(binPath string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "pki", "all-x509-root-certs")
}

// QueryAllNodes runs `query validator all-nodes`.
func QueryAllNodes(binPath string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "validator", "all-nodes")
}

// QueryAllVendors runs `query vendorinfo all-vendors`.
func QueryAllVendors(binPath string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "vendorinfo", "all-vendors")
}

// QueryAccount runs `query auth account`.
func QueryAccount(binPath string, address string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "auth", "account", "--address", address)
}

// QueryProposedAccount runs `query auth proposed-account`.
func QueryProposedAccount(binPath string, address string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "auth", "proposed-account", "--address", address)
}

// QueryProposedAccountToRevoke runs `query auth proposed-account-to-revoke`.
func QueryProposedAccountToRevoke(binPath string, address string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "auth", "proposed-account-to-revoke", "--address", address)
}

// QueryRevokedAccount runs `query auth revoked-account`.
func QueryRevokedAccount(binPath string, address string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "auth", "revoked-account", "--address", address)
}

// QueryLastPower runs `query validator last-power`.
func QueryLastPower(binPath string, address string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "validator", "last-power", "--address", address)
}

// QueryProposedDisableNode runs `query validator proposed-disable-node`.
func QueryProposedDisableNode(binPath string, address string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "validator", "proposed-disable-node", "--address", address)
}

// QueryVendorModels runs `query model vendor-models`.
func QueryVendorModels(binPath string, vid int) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "model", "vendor-models", "--vid", strconv.Itoa(vid))
}

// QueryVendor runs `query vendorinfo vendor`.
func QueryVendor(binPath string, vid int) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "vendorinfo", "vendor", "--vid", strconv.Itoa(vid))
}

// QueryNocX509RootCerts runs `query pki noc-x509-root-certs`.
func QueryNocX509RootCerts(binPath string, vid int) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "pki", "noc-x509-root-certs", "--vid", strconv.Itoa(vid))
}

// QueryGetModel runs `query model get-model`.
func QueryGetModel(binPath string, vid int, pid int) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "model", "get-model", "--vid", strconv.Itoa(vid), "--pid", strconv.Itoa(pid))
}

// QueryAllModelVersions runs `query model all-model-versions`.
func QueryAllModelVersions(binPath string, vid int, pid int) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "model", "all-model-versions", "--vid", strconv.Itoa(vid), "--pid", strconv.Itoa(pid))
}

// QueryModelVersion runs `query model model-version`.
func QueryModelVersion(binPath string, vid int, pid int, sv int) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "model", "model-version", "--vid", strconv.Itoa(vid), "--pid", strconv.Itoa(pid), "--softwareVersion", strconv.Itoa(sv))
}

// QueryCertifiedModel runs `query compliance certified-model`.
func QueryCertifiedModel(binPath string, vid int, pid int, sv int, certType string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "compliance", "certified-model", "--vid", strconv.Itoa(vid), "--pid", strconv.Itoa(pid), "--softwareVersion", strconv.Itoa(sv), "--certificationType", certType)
}

// QueryComplianceInfo runs `query compliance compliance-info`.
func QueryComplianceInfo(binPath string, vid int, pid int, sv int, certType string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "compliance", "compliance-info", "--vid", strconv.Itoa(vid), "--pid", strconv.Itoa(pid), "--softwareVersion", strconv.Itoa(sv), "--certificationType", certType)
}

// QueryProvisionalModel runs `query compliance provisional-model`.
func QueryProvisionalModel(binPath string, vid int, pid int, sv int, certType string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "compliance", "provisional-model", "--vid", strconv.Itoa(vid), "--pid", strconv.Itoa(pid), "--softwareVersion", strconv.Itoa(sv), "--certificationType", certType)
}

// QueryRevokedModel runs `query compliance revoked-model`.
func QueryRevokedModel(binPath string, vid int, pid int, sv int, certType string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "compliance", "revoked-model", "--vid", strconv.Itoa(vid), "--pid", strconv.Itoa(pid), "--softwareVersion", strconv.Itoa(sv), "--certificationType", certType)
}

// QueryDeviceSoftwareCompliance runs `query compliance device-software-compliance`.
func QueryDeviceSoftwareCompliance(binPath string, cdCertID string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "compliance", "device-software-compliance", "--cdCertificateId", cdCertID)
}

// QueryAllSubjectCerts runs `query pki all-subject-certs`.
func QueryAllSubjectCerts(binPath string, subject string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "pki", "all-subject-certs", "--subject", subject)
}

// QueryAllSubjectX509Certs runs `query pki all-subject-x509-certs`.
func QueryAllSubjectX509Certs(binPath string, subject string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "pki", "all-subject-x509-certs", "--subject", subject)
}

// QueryAllNocSubjectX509Certs runs `query pki all-noc-subject-x509-certs`.
func QueryAllNocSubjectX509Certs(binPath string, subject string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "pki", "all-noc-subject-x509-certs", "--subject", subject)
}

// QueryCert runs `query pki cert`.
func QueryCert(binPath string, subject string, subjectKeyID string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "pki", "cert", "--subject", subject, "--subject-key-id", subjectKeyID)
}

// QueryNocX509Cert runs `query pki noc-x509-cert`.
func QueryNocX509Cert(binPath string, subject string, subjectKeyID string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "pki", "noc-x509-cert", "--subject", subject, "--subject-key-id", subjectKeyID)
}

// QueryProposedX509RootCert runs `query pki proposed-x509-root-cert`.
func QueryProposedX509RootCert(binPath string, subject string, subjectKeyID string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "pki", "proposed-x509-root-cert", "--subject", subject, "--subject-key-id", subjectKeyID)
}

// QueryProposedX509RootCertToRevoke runs `query pki proposed-x509-root-cert-to-revoke`.
func QueryProposedX509RootCertToRevoke(binPath string, subject string, subjectKeyID string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "pki", "proposed-x509-root-cert-to-revoke", "--subject", subject, "--subject-key-id", subjectKeyID)
}

// QueryRevokedNocX509RootCert runs `query pki revoked-noc-x509-root-cert`.
func QueryRevokedNocX509RootCert(binPath string, subject string, subjectKeyID string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "pki", "revoked-noc-x509-root-cert", "--subject", subject, "--subject-key-id", subjectKeyID)
}

// QueryRevokedX509Cert runs `query pki revoked-x509-cert`.
func QueryRevokedX509Cert(binPath string, subject string, subjectKeyID string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "pki", "revoked-x509-cert", "--subject", subject, "--subject-key-id", subjectKeyID)
}

// QueryX509Cert runs `query pki x509-cert`.
func QueryX509Cert(binPath string, subject string, subjectKeyID string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "pki", "x509-cert", "--subject", subject, "--subject-key-id", subjectKeyID)
}

// QueryNocX509Certs runs `query pki noc-x509-certs` (the plural command, as it
// existed up to v1.4.3). v1.4.4 renamed it to the singular `noc-x509-cert`
// (see QueryNocX509CertByVid).
func QueryNocX509Certs(binPath string, subjectKeyID string, vid int) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "pki", "noc-x509-certs", "--subject-key-id", subjectKeyID, "--vid", strconv.Itoa(vid))
}

// QueryNocX509CertByVid runs `query pki noc-x509-cert --vid --subject-key-id`,
// the v1.4.4+ form of the VID+SKID NOC lookup (v1.4.3 and earlier used the
// plural noc-x509-certs — see QueryNocX509Certs).
func QueryNocX509CertByVid(binPath string, vid int, subjectKeyID string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "pki", "noc-x509-cert", "--vid", strconv.Itoa(vid), "--subject-key-id", subjectKeyID)
}

// QueryRevocationPoint runs `query pki revocation-point`.
func QueryRevocationPoint(binPath string, vid int, label string, issuerSubjectKeyID string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "pki", "revocation-point", "--vid", strconv.Itoa(vid), "--label", label, "--issuer-subject-key-id", issuerSubjectKeyID)
}

// QueryRevocationPoints runs `query pki revocation-points`.
func QueryRevocationPoints(binPath string, issuerSubjectKeyID string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "pki", "revocation-points", "--issuer-subject-key-id", issuerSubjectKeyID)
}
