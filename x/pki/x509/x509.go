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
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
)

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

	certificate := Certificate{
		Issuer:         ToBase64String(cert.RawIssuer),
		SerialNumber:   cert.SerialNumber.String(),
		Subject:        ToBase64String(cert.RawSubject),
		SubjectAsText:  ToSubjectAsText(cert.Subject.String()),
		SubjectKeyID:   BytesToHex(cert.SubjectKeyId),
		AuthorityKeyID: BytesToHex(cert.AuthorityKeyId),
		Certificate:    cert,
	}

	return &certificate, nil
}

func ToSubjectAsText(subject string) string {
	oldVIDKey := "1.3.6.1.4.1.37244.2.1"
	oldPIDKey := "1.3.6.1.4.1.37244.2.2"

	newVIDKey := "vid"
	newPIDKey := "pid"

	subjectAsText := subject
	subjectAsText = FormatOID(subjectAsText, oldVIDKey, newVIDKey)
	subjectAsText = FormatOID(subjectAsText, oldPIDKey, newPIDKey)

	return subjectAsText
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

// This function is needed to patch the Issuer/Subject(vid/pid) field of certificate to hex format.
// https://github.com/zigbee-alliance/distributed-compliance-ledger/issues/270
func FormatOID(header, oldKey, newKey string) string {
	subjectValues := strings.Split(header, ",")

	// When translating a string number into a hexadecimal number,
	// we must take 8 numbers of this string number from the end so that it needs to fit into an integer number.
	hexStringIntegerLength := 8
	for index, value := range subjectValues {
		if strings.HasPrefix(value, oldKey) {
			// get value from header
			value = value[len(value)-hexStringIntegerLength:]

			decoded, err := hex.DecodeString(value)
			if err != nil {
				continue
			}

			subjectValues[index] = fmt.Sprintf("%s=0x%s", newKey, decoded)
		}
	}

	return strings.Join(subjectValues, ",")
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

func RemoveWhitespaces(pem string) string {
	var builder strings.Builder

	for _, r := range pem {
		if !unicode.IsSpace(r) {
			builder.WriteRune(r)
		}
	}

	return builder.String()
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
//
// Returns:
//   - *Certificate: Parsed certificate structure if validation succeeds
//   - error: Error if validation fails or certificate cannot be parsed
func ParseAndValidateCertificate(pemCertificate string) (*Certificate, error) {
	// 1. Check Certificate Size
	if len(pemCertificate) > MaxCertSize {
		return nil, pkitypes.NewErrInvalidCertificate(fmt.Sprintf("certificate size (%d bytes) exceeds maximum limit of %d bytes", len(pemCertificate), MaxCertSize))
	}

	serialNumberVerification := func(cert *x509.Certificate) error {
		serial := cert.SerialNumber
		// RFC 5280 requires serial numbers to be positive
		if serial.Sign() <= 0 {
			return pkitypes.NewErrInvalidCertificate("serial number must be positive")
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

	return cert, nil
}
