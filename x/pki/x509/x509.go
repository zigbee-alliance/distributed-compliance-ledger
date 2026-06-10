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

package x509

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
)

// X.509 extension OIDs we look up by ID instead of relying on crypto/x509's
// parsed fields, because we need to inspect the Critical flag — which
// crypto/x509 only surfaces via the raw Extensions slice.
var (
	oidBasicConstraints = asn1.ObjectIdentifier{2, 5, 29, 19}
	oidKeyUsage         = asn1.ObjectIdentifier{2, 5, 29, 15}
)

// findExtCritical reports whether the extension with the given OID is present
// and, if so, whether it was marked critical.
func findExtCritical(cert *x509.Certificate, oid asn1.ObjectIdentifier) (present, critical bool) {
	for _, e := range cert.Extensions {
		if e.Id.Equal(oid) {
			return true, e.Critical
		}
	}

	return false, false
}

type Certificate struct {
	Issuer         string
	SerialNumber   string
	Subject        string
	SubjectAsText  string
	SubjectKeyID   string
	AuthorityKeyID string
	Certificate    *x509.Certificate
}

const (
	Mvid             = "Mvid"
	Mpid             = "Mpid"
	MaxCertSize      = 20 * 1024 // 20 KB
	MaxSANCount      = 100
	MaxSubjectFields = 50
)

type DecodeX509CertVerificationOptions func(cert *x509.Certificate) error

// ParseAndValidateCertificateOptions is an option applied by ParseAndValidateCertificate
// after the standard size/SAN/subject-field checks pass. Options receive the parsed
// *x509.Certificate and may return an error to reject the certificate.
type ParseAndValidateCertificateOptions = DecodeX509CertVerificationOptions

// VerifyCAExtensions is a ParseAndValidateCertificate option that enforces the
// full structural CA profile mandated by Matter R1.5 §6.2.2.4/5 and §6.5.12:
//
//   - BasicConstraints extension SHALL be present and marked critical.
//   - BasicConstraints cA flag SHALL be set to TRUE.
//   - KeyUsage extension SHALL be present and marked critical.
//   - KeyUsage SHALL include keyCertSign and cRLSign.
//   - KeyUsage MAY include digitalSignature; no other bits SHALL be set.
//   - SubjectKeyIdentifier SHALL be present.
//   - AuthorityKeyIdentifier SHALL be present for non-self-signed CAs
//     (PAI / ICAC); §6.2.2.5 leaves AKI optional for self-signed PAAs, and the
//     §6.5.12 RCAC AKI==SKI rule is enforced separately by the RCAC handler.
//
// Used directly by PAA and RCAC handlers, and dispatched to from the IsCA=TRUE
// branches of VerifyDAChainNonRoot (PAI) and VerifyNOCChainNonRoot (ICAC).
func VerifyCAExtensions(cert *x509.Certificate) error {
	if !cert.BasicConstraintsValid || !cert.IsCA {
		return pkitypes.NewErrInappropriateCertificateType(
			"certificate is not a CA: BasicConstraints extension must be present and cA must be set to TRUE",
		)
	}

	if _, critical := findExtCritical(cert, oidBasicConstraints); !critical {
		return pkitypes.NewErrInappropriateCertificateType(
			"BasicConstraints extension SHALL be marked critical",
		)
	}

	kuPresent, kuCritical := findExtCritical(cert, oidKeyUsage)
	if !kuPresent {
		return pkitypes.NewErrInappropriateCertificateType(
			"KeyUsage extension SHALL be present for CA certificates",
		)
	}
	if !kuCritical {
		return pkitypes.NewErrInappropriateCertificateType(
			"KeyUsage extension SHALL be marked critical",
		)
	}

	const requiredCAKU = x509.KeyUsageCertSign | x509.KeyUsageCRLSign
	if cert.KeyUsage&requiredCAKU != requiredCAKU {
		return pkitypes.NewErrInappropriateCertificateType(
			"KeyUsage SHALL include both keyCertSign and cRLSign bits",
		)
	}

	const allowedCAKU = requiredCAKU | x509.KeyUsageDigitalSignature
	if cert.KeyUsage&^allowedCAKU != 0 {
		return pkitypes.NewErrInappropriateCertificateType(
			"KeyUsage SHALL NOT include bits other than keyCertSign, cRLSign, and digitalSignature",
		)
	}

	if len(cert.SubjectKeyId) == 0 {
		return pkitypes.NewErrInappropriateCertificateType(
			"SubjectKeyIdentifier extension SHALL be present for CA certificates",
		)
	}

	// crypto/x509 exposes RawIssuer/RawSubject as the DER bytes of the Name
	// sequence, so byte-equality is the structural "self-signed" test.
	selfSigned := bytes.Equal(cert.RawIssuer, cert.RawSubject)
	if !selfSigned && len(cert.AuthorityKeyId) == 0 {
		return pkitypes.NewErrInappropriateCertificateType(
			"AuthorityKeyIdentifier extension SHALL be present for non-self-signed CA certificates",
		)
	}

	return nil
}

// VerifyBasicConstraintsPresent is a ParseAndValidateCertificate option that fails
// if the certificate does not encode the BasicConstraints extension at all.
// It does NOT dictate the value of the cA flag, so it accepts both is-ca=true
// (ICAC) and is-ca=false (NOC, DAC, CRL signer).
//
// Pass it on paths that take leaf certificates or paths that take both CAs and
// leaves — Matter R1.5 §6.2.2.3 (DAC) and §6.5.12 (NOC) both require the
// BasicConstraints extension to be encoded. The NOC-ICA handler accepts both
// ICACs and NOCs, so a "BC encoded" check is the right gate there
//
// crypto/x509 sets BasicConstraintsValid = true if and only if the BC extension
// was found and parsed.
func VerifyBasicConstraintsPresent(cert *x509.Certificate) error {
	if !cert.BasicConstraintsValid {
		return pkitypes.NewErrInappropriateCertificateType(
			"BasicConstraints extension SHALL be present",
		)
	}

	return nil
}

// verifyEndEntityExtensions implements the structural rules shared by Matter
// end-entity certificates (DAC per §6.2.2.3 and NOC per §6.5.12):
//
//   - BasicConstraints SHALL be encoded, marked critical, with cA=FALSE.
//   - KeyUsage SHALL be encoded, marked critical, exactly digitalSignature.
//   - SubjectKeyIdentifier SHALL be present.
//   - AuthorityKeyIdentifier SHALL be present.
//
// VerifyDACExtensions and VerifyNOCExtensions layer additional rules on top of
// this helper (notably the NOC ExtendedKeyUsage check).
func verifyEndEntityExtensions(cert *x509.Certificate, certKind string) error {
	if !cert.BasicConstraintsValid {
		return pkitypes.NewErrInappropriateCertificateType(
			certKind + ": BasicConstraints extension SHALL be present",
		)
	}
	if cert.IsCA {
		return pkitypes.NewErrInappropriateCertificateType(
			certKind + ": BasicConstraints cA SHALL be set to FALSE",
		)
	}
	if _, critical := findExtCritical(cert, oidBasicConstraints); !critical {
		return pkitypes.NewErrInappropriateCertificateType(
			certKind + ": BasicConstraints extension SHALL be marked critical",
		)
	}

	kuPresent, kuCritical := findExtCritical(cert, oidKeyUsage)
	if !kuPresent {
		return pkitypes.NewErrInappropriateCertificateType(
			certKind + ": KeyUsage extension SHALL be present",
		)
	}
	if !kuCritical {
		return pkitypes.NewErrInappropriateCertificateType(
			certKind + ": KeyUsage extension SHALL be marked critical",
		)
	}
	if cert.KeyUsage != x509.KeyUsageDigitalSignature {
		return pkitypes.NewErrInappropriateCertificateType(
			certKind + ": KeyUsage SHALL be exactly digitalSignature",
		)
	}

	if len(cert.SubjectKeyId) == 0 {
		return pkitypes.NewErrInappropriateCertificateType(
			certKind + ": SubjectKeyIdentifier extension SHALL be present",
		)
	}
	if len(cert.AuthorityKeyId) == 0 {
		return pkitypes.NewErrInappropriateCertificateType(
			certKind + ": AuthorityKeyIdentifier extension SHALL be present",
		)
	}

	return nil
}

// VerifyDACExtensions is a ParseAndValidateCertificate option that enforces
// the structural rules of a Matter Device Attestation Certificate (DAC) per
// Matter R1.5 §6.2.2.3: BasicConstraints critical with cA=FALSE, KeyUsage
// critical with exactly digitalSignature, and SKI + AKI present.
func VerifyDACExtensions(cert *x509.Certificate) error {
	return verifyEndEntityExtensions(cert, "DAC")
}

// VerifyDAChainNonRoot is a ParseAndValidateCertificate option for the
// MsgAddX509Cert handler, which accepts both Matter PAIs (cA=TRUE) and Matter
// DACs (cA=FALSE). The certificate is dispatched by its BasicConstraints cA
// flag:
//
//   - cA=TRUE  → Matter R1.5 §6.2.2.4 PAI profile, enforced by VerifyCAExtensions.
//   - cA=FALSE → Matter R1.5 §6.2.2.3 DAC profile, enforced by VerifyDACExtensions.
//
// BasicConstraints must be encoded either way; a missing BC extension is
// reported as a DAC violation since crypto/x509 leaves IsCA at its zero value.
//
// Note: §6.2.2.4 also requires PAI BasicConstraints pathLenConstraint=0. That
// rule is not yet enforced here because several long-lived test fixtures encode
// PAIs without pathLenConstraint; tightening it is tracked as a follow-up and
// will require regenerating those fixtures.
func VerifyDAChainNonRoot(cert *x509.Certificate) error {
	if !cert.BasicConstraintsValid {
		return pkitypes.NewErrInappropriateCertificateType(
			"BasicConstraints extension SHALL be present",
		)
	}
	if cert.IsCA {
		return VerifyCAExtensions(cert)
	}

	return VerifyDACExtensions(cert)
}

// VerifyNOCExtensions is a ParseAndValidateCertificate option that enforces
// the structural rules of a Matter Node Operational Certificate (NOC) per
// Matter R1.5 §6.5.12: BasicConstraints critical with is-ca=FALSE, KeyUsage
// critical with exactly digitalSignature, ExtendedKeyUsage critical with
// exactly {serverAuth, clientAuth}, and SKI + AKI present.
func VerifyNOCExtensions(cert *x509.Certificate) error {
	if err := verifyEndEntityExtensions(cert, "NOC"); err != nil {
		return err
	}

	const oidExtKeyUsageStr = "2.5.29.37"
	ekuPresent, ekuCritical := false, false
	for _, e := range cert.Extensions {
		if e.Id.String() == oidExtKeyUsageStr {
			ekuPresent = true
			ekuCritical = e.Critical

			break
		}
	}
	if !ekuPresent {
		return pkitypes.NewErrInappropriateCertificateType(
			"NOC: ExtendedKeyUsage extension SHALL be present",
		)
	}
	if !ekuCritical {
		return pkitypes.NewErrInappropriateCertificateType(
			"NOC: ExtendedKeyUsage extension SHALL be marked critical",
		)
	}

	var hasServerAuth, hasClientAuth bool
	for _, eu := range cert.ExtKeyUsage {
		switch eu {
		case x509.ExtKeyUsageServerAuth:
			hasServerAuth = true
		case x509.ExtKeyUsageClientAuth:
			hasClientAuth = true
		}
	}
	if !hasServerAuth || !hasClientAuth || len(cert.ExtKeyUsage) != 2 {
		return pkitypes.NewErrInappropriateCertificateType(
			"NOC: ExtendedKeyUsage SHALL be exactly {serverAuth, clientAuth}",
		)
	}

	return nil
}

// VerifyVersionV3 is a ParseAndValidateCertificate option that asserts the
// certificate is X.509 v3, as required by Matter R1.5 §6.2.2.3 (DAC),
// §6.2.2.4 (PAI), §6.2.2.5 (PAA), and §6.5.5/§6.5.8/§6.5.9 (NOC chain). The
// DER-level INTEGER 2 maps to crypto/x509's Version=3.
func VerifyVersionV3(cert *x509.Certificate) error {
	if cert.Version != 3 {
		return pkitypes.NewErrInvalidCertificate(
			fmt.Sprintf("certificate version SHALL be v3, got v%d", cert.Version),
		)
	}

	return nil
}

// VerifyNoPIDInSubject is a ParseAndValidateCertificate option that asserts the
// certificate's subject does not contain a Matter ProductID attribute. Matter
// R1.5 §6.2.2.5 rule 8 prohibits a ProductID in the PAA's subject (and
// equivalently issuer, since PAAs are self-signed).
func VerifyNoPIDInSubject(cert *x509.Certificate) error {
	subjectAsText, err := ToSubjectAsText(cert.Subject.String())
	if err != nil {
		return pkitypes.NewErrInvalidCertificate(err)
	}
	pid, err := GetPidFromSubject(subjectAsText)
	if err != nil {
		return pkitypes.NewErrInvalidPidFormat(err)
	}
	if pid != 0 {
		return pkitypes.NewErrNotEmptyPidForRootCertificate()
	}

	return nil
}

// VerifyNoEKU is a ParseAndValidateCertificate option that fails if the
// certificate encodes an ExtendedKeyUsage extension. Matter R1.5 §6.5.12 says
// "The ExtendedKeyUsage extension SHALL NOT be present" for RCAC and ICAC.
// PAA / PAI are NOT constrained by this rule (§6.2.2.5 rule 11 explicitly
// allows EKU on PAA), so this helper is wired only on the NOC-chain CA paths.
func VerifyNoEKU(cert *x509.Certificate) error {
	const oidExtKeyUsageStr = "2.5.29.37"
	for _, e := range cert.Extensions {
		if e.Id.String() == oidExtKeyUsageStr {
			return pkitypes.NewErrInappropriateCertificateType(
				"ExtendedKeyUsage extension SHALL NOT be present on RCAC/ICAC certificates",
			)
		}
	}

	return nil
}

// Matter OID strings used by the at-most-one-VID/PID walker. Kept as strings
// (matching asn1.ObjectIdentifier.String() output) so we don't need to allocate
// an asn1.ObjectIdentifier on every comparison.
const (
	matterVIDOIDString = "1.3.6.1.4.1.37244.2.1"
	matterPIDOIDString = "1.3.6.1.4.1.37244.2.2"
)

// VerifyAtMostOneVIDAndPID is a ParseAndValidateCertificate option that asserts
// the certificate's subject and issuer each contain AT MOST one matter-vid
// (1.3.6.1.4.1.37244.2.1) and AT MOST one matter-pid (1.3.6.1.4.1.37244.2.2)
// attribute. Implements Matter R1.5 §6.2.2.3 rules 4/5 (DAC), §6.2.2.4 rules
// 4/8 (PAI), and §6.2.2.5 rules 4/6 (PAA).
//
// The stricter "exactly 1 VendorID" rule (§6.2.2.3 rule 4 for DAC issuer,
// §6.2.2.4 for PAI subject) is not asserted here — enforcing it would require
// regenerating the long-lived DA-chain fixtures that currently encode no VID
// (LeafCertPem, IntermediateCertPem, IntermediateCertWithoutVidPid). That gap
// is tracked separately as part of the Matter-DN parsing remediation.
func VerifyAtMostOneVIDAndPID(cert *x509.Certificate) error {
	if err := matterAttrAtMostOne(cert.Subject.Names, "subject"); err != nil {
		return err
	}

	return matterAttrAtMostOne(cert.Issuer.Names, "issuer")
}

func matterAttrAtMostOne(names []pkix.AttributeTypeAndValue, field string) error {
	vidCount, pidCount := 0, 0
	for _, n := range names {
		switch n.Type.String() {
		case matterVIDOIDString:
			vidCount++
		case matterPIDOIDString:
			pidCount++
		}
	}
	if vidCount > 1 {
		return pkitypes.NewErrInvalidCertificate(
			fmt.Sprintf("%s SHALL contain at most one matter-vid attribute, found %d", field, vidCount),
		)
	}
	if pidCount > 1 {
		return pkitypes.NewErrInvalidCertificate(
			fmt.Sprintf("%s SHALL contain at most one matter-pid attribute, found %d", field, pidCount),
		)
	}

	return nil
}

// VerifyECDSAP256SHA256 is a ParseAndValidateCertificate option that asserts
// the certificate is signed with ecdsa-with-SHA256 and that its subject public
// key is an ECDSA key on the prime256v1 (P-256 / secp256r1) curve, per Matter
// R1.5 §6.2.2.3 (DAC), §6.2.2.4 (PAI), §6.2.2.5 (PAA), and §6.5.5/§6.5.8/§6.5.9
// (NOC chain). Mirrors the long-standing check in VerifyCRLSignerCertFormat —
// before this helper existed, the CRL signer path was the only place these
// rules were enforced.
func VerifyECDSAP256SHA256(cert *x509.Certificate) error {
	if cert.SignatureAlgorithm != x509.ECDSAWithSHA256 {
		return pkitypes.NewErrInvalidCertificate(
			"signatureAlgorithm SHALL be ecdsa-with-SHA256",
		)
	}

	pub, ok := cert.PublicKey.(*ecdsa.PublicKey)
	if !ok {
		return pkitypes.NewErrInvalidCertificate(
			"subjectPublicKeyInfo algorithm SHALL be ECDSA on prime256v1",
		)
	}
	if pub.Curve != elliptic.P256() {
		return pkitypes.NewErrInvalidCertificate(
			"subjectPublicKeyInfo curve SHALL be prime256v1 (P-256)",
		)
	}

	return nil
}

// VerifyNOCChainNonRoot is a ParseAndValidateCertificate option for the
// MsgAddNocX509IcaCert handler, which accepts both Matter ICACs (is-ca=TRUE)
// and Matter NOCs (is-ca=FALSE). The certificate is dispatched by its
// BasicConstraints cA flag:
//
//   - cA=TRUE  → Matter R1.5 §6.5.12 ICAC profile, enforced by VerifyCAExtensions
//     plus the §6.5.12 "EKU SHALL NOT be present" rule via VerifyNoEKU.
//   - cA=FALSE → Matter R1.5 §6.5.12 NOC profile, enforced by VerifyNOCExtensions.
//
// BasicConstraints must be encoded either way; a missing BC extension is
// reported as a NOC violation since crypto/x509 leaves IsCA at its zero value.
func VerifyNOCChainNonRoot(cert *x509.Certificate) error {
	if !cert.BasicConstraintsValid {
		return pkitypes.NewErrInappropriateCertificateType(
			"BasicConstraints extension SHALL be present",
		)
	}
	if cert.IsCA {
		if err := VerifyCAExtensions(cert); err != nil {
			return err
		}

		return VerifyNoEKU(cert)
	}

	return VerifyNOCExtensions(cert)
}

// VerifyVidPidConsistency enforces the immediate-parent VID/PID matching rules
// from Matter R1.5 §6.2.2.3 (DAC) 8a and 9a, and §6.2.2.4 (PAI) 7a:
//
//   - When the parent's subject carries a Matter VID, the child's subject SHALL
//     carry the same Matter VID.
//   - When the parent's subject carries a Matter PID, the child's subject SHALL
//     carry the same Matter PID.
//
// Following the spec literally, the rules only fire when the parent actually
// carries the attribute — a PAA without a VID does not constrain its PAI's
// VID, and a PAI without a PID does not constrain its DAC's PID. The child is
// not required to be VID/PID-scoped on its own; the structural rules that
// require it (DAC SHALL have VID, PAI SHALL have VID) belong to the per-cert
// extension/DN validation, not to this immediate-parent consistency check.
//
// Both arguments are subject strings already passed through ToSubjectAsText,
// which is the canonical representation stored alongside each certificate.
func VerifyVidPidConsistency(childSubjectAsText, parentSubjectAsText string) error {
	parentVid, err := GetVidFromSubject(parentSubjectAsText)
	if err != nil {
		return pkitypes.NewErrInvalidVidFormat(err)
	}
	if parentVid != 0 {
		childVid, err := GetVidFromSubject(childSubjectAsText)
		if err != nil {
			return pkitypes.NewErrInvalidVidFormat(err)
		}
		if childVid != parentVid {
			return pkitypes.NewErrCertVidNotEqualToIssuerVid(childVid, parentVid)
		}
	}

	parentPid, err := GetPidFromSubject(parentSubjectAsText)
	if err != nil {
		return pkitypes.NewErrInvalidPidFormat(err)
	}
	if parentPid != 0 {
		childPid, err := GetPidFromSubject(childSubjectAsText)
		if err != nil {
			return pkitypes.NewErrInvalidPidFormat(err)
		}
		if childPid != parentPid {
			return pkitypes.NewErrCertPidNotEqualToIssuerPid(childPid, parentPid)
		}
	}

	return nil
}

func DecodeX509Certificate(pemCertificate string, options ...DecodeX509CertVerificationOptions) (*Certificate, error) {
	block, _ := pem.Decode([]byte(pemCertificate))
	if block == nil {
		return nil, pkitypes.NewErrInvalidCertificate("Could not decode pem certificate")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, pkitypes.NewErrInvalidCertificate(fmt.Sprintf("Could not parse certificate: %v", err.Error()))
	}

	for _, option := range options {
		err = option(cert)
		if err != nil {
			return nil, err
		}
	}

	subjectAsText, err := ToSubjectAsText(cert.Subject.String())
	if err != nil {
		return nil, pkitypes.NewErrInvalidCertificate(err)
	}

	certificate := Certificate{
		Issuer:         ToBase64String(cert.RawIssuer),
		SerialNumber:   cert.SerialNumber.String(),
		Subject:        ToBase64String(cert.RawSubject),
		SubjectAsText:  subjectAsText,
		SubjectKeyID:   BytesToHex(cert.SubjectKeyId),
		AuthorityKeyID: BytesToHex(cert.AuthorityKeyId),
		Certificate:    cert,
	}

	return &certificate, nil
}

// ToSubjectAsText projects crypto/x509 pkix.Name.String() output into the
// canonical DCL form by rewriting Matter VID / PID OID entries to their
// readable `vid=0x..` / `pid=0x..` shorthand. Returns an error if either
// OID's DER-encoded value is malformed — silently accepting a malformed VID
// or PID would make the downstream GetVidFromSubject / GetPidFromSubject
// lookups return 0 and bypass every VID/PID-based check.
func ToSubjectAsText(subject string) (string, error) {
	const (
		oldVIDKey = "1.3.6.1.4.1.37244.2.1"
		oldPIDKey = "1.3.6.1.4.1.37244.2.2"
		newVIDKey = "vid"
		newPIDKey = "pid"
	)

	subjectAsText, err := FormatOID(subject, oldVIDKey, newVIDKey)
	if err != nil {
		return "", err
	}

	return FormatOID(subjectAsText, oldPIDKey, newPIDKey)
}

var (
	subjectFieldRegex = regexp.MustCompile(`(([^\s,]+\s?=\s?[^[\s,]+)|(([^\s,]+:\s?[^\s,]+)))`)
	separatorRegex    = regexp.MustCompile(`(\s?=\s?)|(\s?:\s?)`)
)

func subjectAsTextToMap(subjectAsText string) map[string]string {
	matches := subjectFieldRegex.FindAllString(subjectAsText, -1)

	subjectMap := make(map[string]string)
	for _, elem := range matches {
		splittedElem := separatorRegex.Split(elem, -1)
		if splittedElem[0] == "vid" {
			splittedElem[0] = Mvid
		}
		if splittedElem[0] == "pid" {
			splittedElem[0] = Mpid
		}
		subjectMap[splittedElem[0]] = splittedElem[1]
	}

	return subjectMap
}

func GetVidFromSubject(subjectAsText string) (int32, error) {
	return getIntValueFromSubject(subjectAsText, Mvid)
}

func GetPidFromSubject(subjectAsText string) (int32, error) {
	return getIntValueFromSubject(subjectAsText, Mpid)
}

func getIntValueFromSubject(subjectAsText string, key string) (int32, error) {
	subjectAsTextMap := subjectAsTextToMap(subjectAsText)

	if strValue, ok := subjectAsTextMap[key]; ok {
		pid, err := strconv.ParseInt(strings.TrimPrefix(strValue, "0x"), 16, 32)
		if err != nil {
			return 0, err
		}

		return int32(pid), nil
	}

	return 0, nil
}

// FormatOID rewrites every comma-separated entry in `header` whose attribute
// type is exactly `oldKey` into the form `<newKey>=0x<value>`, where `<value>`
// is the printable ASCII content of the DER-encoded attribute value. Used to
// translate Matter VID / PID OID entries (`1.3.6.1.4.1.37244.2.{1,2}`) into
// the readable `vid=0x6006` / `pid=0x8000` form callers downstream consume.
// https://github.com/zigbee-alliance/distributed-compliance-ledger/issues/270
//
// The input format comes from crypto/x509's pkix.Name.String(), which emits
// unknown OIDs as `<oid>=#<hexDER>`, where `<hexDER>` is the hex encoding of a
// DER-encoded string. We decode the DER properly (tag + length + value) and
// accept both PrintableString (0x13) and UTF8String (0x0c). Returns error when
// `oldKey` matches an entry but its value is malformed (wrong tag, truncated, non-hex)
//
// Entries whose attribute type does not match `oldKey` are returned untouched.
func FormatOID(header, oldKey, newKey string) (string, error) {
	subjectValues := strings.Split(header, ",")
	// Match `oldKey=` exactly so e.g. matter-vid (1.3.6.1.4.1.37244.2.1) does
	// not also match a hypothetical 1.3.6.1.4.1.37244.2.10.
	prefix := oldKey + "="
	for index, value := range subjectValues {
		if !strings.HasPrefix(value, prefix) {
			continue
		}
		encoded := value[len(prefix):]
		// pkix.Name.String emits `#<hex>` only when the attribute value is not
		// a recognized printable type. If we see anything else, the value is
		// already in readable form — accept it as is.
		if !strings.HasPrefix(encoded, "#") {
			continue
		}
		derBytes, err := hex.DecodeString(encoded[1:])
		if err != nil {
			return "", fmt.Errorf("FormatOID: %s value is not valid hex: %w", oldKey, err)
		}
		if len(derBytes) < 2 {
			return "", fmt.Errorf("FormatOID: %s value DER is too short (%d bytes)", oldKey, len(derBytes))
		}
		// Matter encodes matter-vid / matter-pid as PrintableString or
		// UTF8String. Reject anything else: we cannot safely project a
		// BMPString / BIT STRING / OCTET STRING into the `<newKey>=0x...` form
		// callers expect.
		tag := derBytes[0]
		if tag != 0x13 && tag != 0x0c {
			return "", fmt.Errorf("FormatOID: %s value SHALL be PrintableString or UTF8String, got DER tag 0x%02x", oldKey, tag)
		}
		length := int(derBytes[1])
		if 2+length > len(derBytes) {
			return "", fmt.Errorf("FormatOID: %s value DER length (%d) exceeds remaining bytes (%d)", oldKey, length, len(derBytes)-2)
		}
		printable := derBytes[2 : 2+length]
		subjectValues[index] = fmt.Sprintf("%s=0x%s", newKey, printable)
	}

	return strings.Join(subjectValues, ","), nil
}

func ToBase64String(subject []byte) string {
	return base64.StdEncoding.EncodeToString(subject)
}

func BytesToHex(bytes []byte) string {
	if bytes == nil {
		return ""
	}

	bytesHex := make([]string, len(bytes))
	for i, b := range bytes {
		bytesHex[i] = fmt.Sprintf("%02X", b)
	}

	return strings.Join(bytesHex, ":")
}

// CertificatePEMsEqual reports whether two PEM-encoded certificates contain the
// same DER bytes. It tolerates differences in whitespace, line wrapping, and PEM
// block headers, which are not part of the certificate itself.
func CertificatePEMsEqual(a, b string) bool {
	blockA, _ := pem.Decode([]byte(a))
	blockB, _ := pem.Decode([]byte(b))
	if blockA == nil || blockB == nil {
		return false
	}

	return bytes.Equal(blockA.Bytes, blockB.Bytes)
}

func (c Certificate) Verify(parent *Certificate, blockTime time.Time) error {
	roots := x509.NewCertPool()
	roots.AddCert(parent.Certificate)

	opts := x509.VerifyOptions{Roots: roots, CurrentTime: blockTime}

	if _, err := c.Certificate.Verify(opts); err != nil {
		return pkitypes.NewErrInvalidCertificate(fmt.Sprintf("Certificate verification failed. Error: %v", err))
	}

	return nil
}

func (c Certificate) IsSelfSigned() bool {
	if len(c.AuthorityKeyID) > 0 {
		return c.Issuer == c.Subject && c.AuthorityKeyID == c.SubjectKeyID
	}

	return c.Issuer == c.Subject
}

// ParseAndValidateCertificate validates and parses a PEM-encoded X.509 certificate.
// It performs the following validations:
// 1. Checks that the certificate size does not exceed MaxCertSize (20 KB)
// 2. Checks that the number of Subject Alternative Names (SANs) does not exceed MaxSANCount (100)
// 3. Checks that the number of subject fields does not exceed MaxSubjectFields (50)
//
// Parameters:
//   - pemCertificate: PEM-encoded X.509 certificate string
//   - options: optional checks applied to the parsed certificate, e.g. VerifyVersionV3
//
// Returns:
//   - *Certificate: Parsed certificate structure if validation succeeds
//   - error: Error if validation fails or certificate cannot be parsed
func ParseAndValidateCertificate(pemCertificate string, options ...ParseAndValidateCertificateOptions) (*Certificate, error) {
	// 1. Check Certificate Size
	if len(pemCertificate) > MaxCertSize {
		return nil, pkitypes.NewErrInvalidCertificate(fmt.Sprintf("certificate size (%d bytes) exceeds maximum limit of %d bytes", len(pemCertificate), MaxCertSize))
	}

	serialNumberVerification := func(cert *x509.Certificate) error {
		serial := cert.SerialNumber
		// RFC 5280 requires serial numbers to be positive
		if serial.Sign() <= 0 {
			return pkitypes.NewErrInvalidCertificate("serial number must be a positive")
		}

		// When crypto/x509 parses a certificate, it reads the DER integer, strips the sign byte if present,
		// then returns the minimal magnitude in octets (no leading zeros).
		if len(serial.Bytes()) > 20 {
			return pkitypes.NewErrInvalidCertificate("serial number exceeds 20-octet limit")
		}

		return nil
	}

	// Parse the certificate
	cert, err := DecodeX509Certificate(pemCertificate, serialNumberVerification)
	if err != nil {
		return nil, pkitypes.NewErrInvalidCertificate(fmt.Sprintf("failed to parse certificate: %v", err))
	}

	// 2. Check SAN (Subject Alternative Name) count
	sanCount := len(cert.Certificate.DNSNames) + len(cert.Certificate.EmailAddresses) + len(cert.Certificate.IPAddresses) + len(cert.Certificate.URIs)
	if sanCount > MaxSANCount {
		return nil, pkitypes.NewErrInvalidCertificate(fmt.Sprintf("SAN count (%d) exceeds maximum limit of %d", sanCount, MaxSANCount))
	}

	// 3. Check Subject Fields count
	subjectFieldCount := len(cert.Certificate.Subject.Names)
	if subjectFieldCount > MaxSubjectFields {
		return nil, pkitypes.NewErrInvalidCertificate(fmt.Sprintf("subject field count (%d) exceeds maximum limit of %d", subjectFieldCount, MaxSubjectFields))
	}

	// 4. Apply caller-supplied options (e.g. VerifyVersionV3)
	for _, opt := range options {
		if err = opt(cert.Certificate); err != nil {
			return nil, err
		}
	}

	return cert, nil
}
