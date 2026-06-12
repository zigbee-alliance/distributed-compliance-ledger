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

//nolint:testpackage
package x509

import (
	x509std "crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"testing"
	"time"

	tmrand "github.com/cometbft/cometbft/libs/rand"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
)

func Test_DecodeCertificates(t *testing.T) {
	// decode leaf certificate
	certificate, err := DecodeX509Certificate(testconstants.LeafCertPem)
	require.Nil(t, err)
	require.False(t, certificate.IsSelfSigned())
	require.Equal(t, testconstants.IntermediateSubject, certificate.Issuer)
	require.Equal(t, testconstants.LeafSubjectAsText, certificate.SubjectAsText)
	require.Equal(t, testconstants.IntermediateSubjectKeyID, certificate.AuthorityKeyID)
	require.Equal(t, testconstants.LeafSerialNumber, certificate.SerialNumber)
	require.Equal(t, testconstants.LeafSubject, certificate.Subject)
	require.Equal(t, testconstants.LeafSubjectKeyID, certificate.SubjectKeyID)

	// decode intermediate certificate
	certificate, err = DecodeX509Certificate(testconstants.IntermediateCertPem)
	require.Nil(t, err)
	require.False(t, certificate.IsSelfSigned())
	require.Equal(t, testconstants.RootSubject, certificate.Issuer)
	require.Equal(t, testconstants.IntermediateSubjectAsText, certificate.SubjectAsText)
	require.Equal(t, testconstants.RootSubjectKeyID, certificate.AuthorityKeyID)
	require.Equal(t, testconstants.IntermediateSerialNumber, certificate.SerialNumber)
	require.Equal(t, testconstants.IntermediateSubject, certificate.Subject)
	require.Equal(t, testconstants.IntermediateSubjectKeyID, certificate.SubjectKeyID)

	// decode root certificate
	certificate, err = DecodeX509Certificate(testconstants.RootCertPem)
	require.Nil(t, err)
	require.True(t, certificate.IsSelfSigned())
	require.Equal(t, testconstants.RootSubject, certificate.Issuer)
	require.Equal(t, testconstants.RootSubjectAsText, certificate.SubjectAsText)
	require.Equal(t, testconstants.RootSerialNumber, certificate.SerialNumber)
	require.Equal(t, testconstants.RootSubject, certificate.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, certificate.SubjectKeyID)
}

func Test_DecodeCertificatesWithVID(t *testing.T) {
	// decode root google certificate with vid
	certificate, err := DecodeX509Certificate(testconstants.GoogleCertPem)
	require.Nil(t, err)
	require.True(t, certificate.IsSelfSigned())
	require.Equal(t, testconstants.GoogleSubject, certificate.Issuer)
	require.Equal(t, testconstants.GoogleSubjectAsText, certificate.SubjectAsText)
	require.Equal(t, testconstants.GoogleSerialNumber, certificate.SerialNumber)
	require.Equal(t, testconstants.GoogleSubject, certificate.Subject)
	require.Equal(t, testconstants.GoogleSubjectKeyID, certificate.SubjectKeyID)

	// decode root test certificate with vid
	certificate, err = DecodeX509Certificate(testconstants.TestCertPem)
	require.Nil(t, err)
	require.True(t, certificate.IsSelfSigned())
	require.Equal(t, testconstants.TestSubject, certificate.Issuer)
	require.Equal(t, testconstants.TestSubjectAsText, certificate.SubjectAsText)
	require.Equal(t, testconstants.TestSerialNumber, certificate.SerialNumber)
	require.Equal(t, testconstants.TestSubject, certificate.Subject)
	require.Equal(t, testconstants.TestSubjectKeyID, certificate.SubjectKeyID)
	require.Equal(t, testconstants.TestAuthorityKeyID, certificate.AuthorityKeyID)
}

func Test_VerifyLeafCertificate(t *testing.T) {
	certificate, _ := DecodeX509Certificate(testconstants.LeafCertPem)
	parentCertificate, _ := DecodeX509Certificate(testconstants.IntermediateCertPem)
	blockTime := time.Date(2022, 12, 22, 22, 22, 22, 22, time.UTC)

	err := certificate.Verify(parentCertificate, blockTime)
	require.Nil(t, err)
}

func Test_VerifyRootCertificate(t *testing.T) {
	certificate, _ := DecodeX509Certificate(testconstants.RootCertPem)
	blockTime := time.Date(2022, 12, 22, 22, 22, 22, 22, time.UTC)

	err := certificate.Verify(certificate, blockTime)
	require.Nil(t, err)
}

func Test_FastSync_VerifyExpiredRootCertificateWhenBlockTimeInPast(t *testing.T) {
	certificate, _ := DecodeX509Certificate(testconstants.PAACertExpired)
	blockTime := time.Date(2022, 5, 4, 22, 22, 22, 22, time.UTC)

	err := certificate.Verify(certificate, blockTime)
	require.Nil(t, err)
}

func Test_BytesToHex(t *testing.T) {
	tests := []struct {
		subjectKeyID []byte
		result       string
	}{
		{
			subjectKeyID: []byte("\xb0\x00V\x81\xb8\x88b\x89b\x80\xe1!\x18\xa1\xa8\xbe\tޓ!"),
			result:       "B0:00:56:81:B8:88:62:89:62:80:E1:21:18:A1:A8:BE:09:DE:93:21",
		},
		{
			subjectKeyID: []byte("␍6\x9c<\xa3\xc1\x13\xbb\t\xe2M\xc1\xccŦf\x91\xd4"),
			result:       "E2:90:8D:36:9C:3C:A3:C1:13:BB:09:E2:4D:C1:CC:C5:A6:66:91:D4",
		},
	}

	for _, tt := range tests {
		result := BytesToHex(tt.subjectKeyID)
		require.Equal(t, result, tt.result)
	}
}

func Test_FormatVID(t *testing.T) {
	positiveTests := []struct {
		header string
		oldKey string
		newKey string
		result string
	}{
		{
			header: "CN=Matter PAA 1,O=Google,C=US,1.3.6.1.4.1.37244.2.1=#130436303036",
			oldKey: "1.3.6.1.4.1.37244.2.1",
			newKey: "vid",
			result: "CN=Matter PAA 1,O=Google,C=US,vid=0x6006",
		},
		{
			header: "CN=Matter Test PAA,1.3.6.1.4.1.37244.2.1=#130431323544",
			oldKey: "1.3.6.1.4.1.37244.2.1",
			newKey: "vid",
			result: "CN=Matter Test PAA,vid=0x125D",
		},
	}

	negativeTests := []struct {
		header string
		oldKey string
		newKey string
		result string
	}{
		// set an incorrect header
		{
			header: "CN=Matter PAA 1,O=Google,C=US,1.3.6=#130436303036",
			oldKey: "1.3.6.1.4.1.37244.2.1",
			newKey: "vid",
			result: "CN=Matter PAA 1,O=Google,C=US,1.3.6=#130436303036",
		},
	}

	for _, tt := range positiveTests {
		result, err := FormatOID(tt.header, tt.oldKey, tt.newKey)
		require.NoError(t, err)
		require.Equal(t, tt.result, result)
	}

	for _, tt := range negativeTests {
		result, err := FormatOID(tt.header, tt.oldKey, tt.newKey)
		require.NoError(t, err)
		require.Equal(t, tt.result, result)
	}
}

func Test_FormatPID(t *testing.T) {
	positiveTests := []struct {
		header string
		oldKey string
		newKey string
		result string
	}{
		{
			header: "CN=Matter PAA 1,O=Google,C=US,1.3.6.1.4.1.37244.2.2=#130436303036",
			oldKey: "1.3.6.1.4.1.37244.2.2",
			newKey: "pid",
			result: "CN=Matter PAA 1,O=Google,C=US,pid=0x6006",
		},
		{
			header: "CN=Matter Test PAA,1.3.6.1.4.1.37244.2.2=#130431323544",
			oldKey: "1.3.6.1.4.1.37244.2.2",
			newKey: "pid",
			result: "CN=Matter Test PAA,pid=0x125D",
		},
		{
			header: "CN=Matter Test PAA,1.3.6.1.4.1.37244.2.2=#130431323544,SomeWord",
			oldKey: "1.3.6.1.4.1.37244.2.2",
			newKey: "pid",
			result: "CN=Matter Test PAA,pid=0x125D,SomeWord",
		},
	}

	// "Unchanged" cases — oldKey does not match any entry, so FormatOID
	// returns the header verbatim with no error. These cover legitimate input
	// where the rewrite is simply not applicable.
	unchangedTests := []struct {
		header string
		oldKey string
		newKey string
		result string
	}{
		// Different OID under the same prefix tree.
		{
			header: "CN=Matter PAA 1,O=Google,C=US,1.3.6.1.4.1.37244.2.1=#130436303036",
			oldKey: "1.3.6.1.4.1.37244.2.2",
			newKey: "vid",
			result: "CN=Matter PAA 1,O=Google,C=US,1.3.6.1.4.1.37244.2.1=#130436303036",
		},
		// Prefix-only match must NOT trigger rewriting: the OID
		// 1.3.6.1.4.1.37244.2.10 (hypothetical) shares a textual prefix with
		// 1.3.6.1.4.1.37244.2.1 but is a different attribute.
		{
			header: "CN=X,1.3.6.1.4.1.37244.2.10=#130436303036",
			oldKey: "1.3.6.1.4.1.37244.2.1",
			newKey: "vid",
			result: "CN=X,1.3.6.1.4.1.37244.2.10=#130436303036",
		},
	}

	// "Decoded" cases — oldKey matches and the value is well-formed but uses
	// encodings the original FormatOID was brittle about.
	decodedTests := []struct {
		header string
		oldKey string
		newKey string
		result string
	}{
		// UTF8String-tagged values (tag 0x0c) must be decoded too.
		{
			header: "CN=X,1.3.6.1.4.1.37244.2.2=#0c0438303030",
			oldKey: "1.3.6.1.4.1.37244.2.2",
			newKey: "pid",
			result: "CN=X,pid=0x8000",
		},
		// 6-byte value — the hardened parser uses the DER length byte, not a
		// hardcoded 8-hex-char slice.
		{
			header: "CN=X,1.3.6.1.4.1.37244.2.2=#1306414243444546",
			oldKey: "1.3.6.1.4.1.37244.2.2",
			newKey: "pid",
			result: "CN=X,pid=0xABCDEF",
		},
	}

	// "Error" cases — oldKey matches BUT the value is malformed. FormatOID
	// must refuse to project these into the readable form: silently keeping
	// the raw OID entry would make GetVidFromSubject / GetPidFromSubject
	// return 0 and bypass every VID/PID-based check.
	errorTests := []struct {
		header       string
		oldKey       string
		newKey       string
		expectSubstr string
	}{
		// Unknown tag (e.g. OCTET STRING 0x04).
		{
			header:       "CN=X,1.3.6.1.4.1.37244.2.2=#040436303036",
			oldKey:       "1.3.6.1.4.1.37244.2.2",
			newKey:       "pid",
			expectSubstr: "PrintableString or UTF8String",
		},
		// Truncated DER: length byte says 4 but only 1 byte follows.
		{
			header:       "CN=X,1.3.6.1.4.1.37244.2.2=#130436",
			oldKey:       "1.3.6.1.4.1.37244.2.2",
			newKey:       "pid",
			expectSubstr: "DER length",
		},
		// Non-hex content after `#`.
		{
			header:       "CN=X,1.3.6.1.4.1.37244.2.2=#13ZZZZZZ",
			oldKey:       "1.3.6.1.4.1.37244.2.2",
			newKey:       "pid",
			expectSubstr: "not valid hex",
		},
	}

	for _, tt := range positiveTests {
		result, err := FormatOID(tt.header, tt.oldKey, tt.newKey)
		require.NoError(t, err)
		require.Equal(t, tt.result, result)
	}

	for _, tt := range unchangedTests {
		result, err := FormatOID(tt.header, tt.oldKey, tt.newKey)
		require.NoError(t, err)
		require.Equal(t, tt.result, result)
	}

	for _, tt := range decodedTests {
		result, err := FormatOID(tt.header, tt.oldKey, tt.newKey)
		require.NoError(t, err)
		require.Equal(t, tt.result, result)
	}

	for _, tt := range errorTests {
		result, err := FormatOID(tt.header, tt.oldKey, tt.newKey)
		require.Error(t, err)
		require.Empty(t, result)
		require.Contains(t, err.Error(), tt.expectSubstr)
	}
}

func Test_ParseAndValidateCertificate(t *testing.T) {
	positiveTests := []struct {
		name            string
		certPem         string
		expectedSubject string
		expectedKeyID   string
		isSelfSigned    bool
	}{
		{
			name:            "valid leaf certificate",
			certPem:         testconstants.LeafCertPem,
			expectedSubject: testconstants.LeafSubject,
			expectedKeyID:   testconstants.LeafSubjectKeyID,
			isSelfSigned:    false,
		},
		{
			name:            "valid intermediate certificate",
			certPem:         testconstants.IntermediateCertPem,
			expectedSubject: testconstants.IntermediateSubject,
			expectedKeyID:   testconstants.IntermediateSubjectKeyID,
			isSelfSigned:    false,
		},
		{
			name:            "valid root certificate",
			certPem:         testconstants.RootCertPem,
			expectedSubject: testconstants.RootSubject,
			expectedKeyID:   testconstants.RootSubjectKeyID,
			isSelfSigned:    true,
		},
		{
			name:            "valid Google certificate with VID",
			certPem:         testconstants.GoogleCertPem,
			expectedSubject: testconstants.GoogleSubject,
			expectedKeyID:   testconstants.GoogleSubjectKeyID,
			isSelfSigned:    true,
		},
		{
			name:            "valid PAA certificate with VID",
			certPem:         testconstants.PAACertWithNumericVid,
			expectedSubject: testconstants.PAACertWithNumericVidSubject,
			expectedKeyID:   testconstants.PAACertWithNumericVidSubjectKeyID,
			isSelfSigned:    true,
		},
		{
			name:            "valid PAI certificate with VID",
			certPem:         testconstants.PAICertWithNumericVid,
			expectedSubject: testconstants.PAICertWithNumericVidSubject,
			expectedKeyID:   testconstants.PAICertWithNumericVidSubjectKeyID,
			isSelfSigned:    false,
		},
	}

	negativeTests := []struct {
		name              string
		certPem           string
		expectErrorSubstr string
	}{
		{
			name:              "empty certificate string",
			certPem:           "",
			expectErrorSubstr: "failed to parse certificate",
		},
		{
			name:              "malformed PEM certificate",
			certPem:           "-----BEGIN CERTIFICATE-----\nInvalidData\n-----END CERTIFICATE-----",
			expectErrorSubstr: "failed to parse certificate",
		},
		{
			name:              "certificate size exceeds 20 KB",
			certPem:           tmrand.Str(20481),
			expectErrorSubstr: "exceeds maximum limit",
		},
		{
			name:              "certificate serial number exceeds 20 octets",
			certPem:           testconstants.CertWithSerialNumber21octets,
			expectErrorSubstr: "serial number exceeds",
		},
		{
			name:              "certificate serial number is negative",
			certPem:           testconstants.CertWithInvalidSerialNumber,
			expectErrorSubstr: "serial number must be a positive",
		},
	}

	for _, tt := range positiveTests {
		certificate, err := ParseAndValidateCertificate(tt.certPem)
		require.NoError(t, err)
		require.NotNil(t, certificate)
		require.Equal(t, tt.expectedSubject, certificate.Subject)
		require.Equal(t, tt.expectedKeyID, certificate.SubjectKeyID)
		require.Equal(t, tt.isSelfSigned, certificate.IsSelfSigned())
	}

	for _, tt := range negativeTests {
		certificate, err := ParseAndValidateCertificate(tt.certPem)
		require.Error(t, err)
		require.Nil(t, certificate)
		if tt.expectErrorSubstr != "" {
			require.Contains(t, err.Error(), tt.expectErrorSubstr)
		}
	}
}

func Test_ParseAndValidateCertificate_VerifyBasicConstraintsPresent(t *testing.T) {
	positiveTests := []struct {
		name    string
		certPem string
	}{
		{name: "root CA (BC encoded, cA=true)", certPem: testconstants.RootCertPem},
		{name: "intermediate CA (BC encoded, cA=true)", certPem: testconstants.IntermediateCertPem},
		{name: "leaf (BC encoded, cA=false)", certPem: testconstants.LeafCertPem},
		{name: "PAA with VID (BC encoded, cA=true)", certPem: testconstants.PAACertWithNumericVid},
	}
	for _, tt := range positiveTests {
		t.Run("ok/"+tt.name, func(t *testing.T) {
			cert, err := ParseAndValidateCertificate(tt.certPem, VerifyBasicConstraintsPresent)
			require.NoError(t, err)
			require.NotNil(t, cert)
		})
	}

	// A cert that has no BasicConstraints extension at all must be rejected.
	t.Run("reject/BC-extension-absent", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.LeafCertPem)
		require.NoError(t, err)
		cert.Certificate.BasicConstraintsValid = false

		err = VerifyBasicConstraintsPresent(cert.Certificate)
		require.Error(t, err)
		require.ErrorIs(t, err, pkitypes.ErrInappropriateCertificateType)
	})
}

func Test_ParseAndValidateCertificate_VerifyCAExtensions(t *testing.T) {
	// Positive fixtures: known CA certs that comply with the full Matter R1.6
	// CA profile (BC critical + cA=TRUE, KU critical with at least keyCertSign
	// and cRLSign, no disallowed bits).
	positiveTests := []struct {
		name    string
		certPem string
	}{
		{name: "PAA with VID", certPem: testconstants.PAACertWithNumericVid},
		{name: "PAA no VID", certPem: testconstants.PAACertNoVid},
		{name: "Google PAA", certPem: testconstants.GoogleCertPem},
		{name: "PAI", certPem: testconstants.PAICertWithNumericVid},
		{name: "regenerated RCAC NocRootCert2", certPem: testconstants.NocRootCert2},
	}
	for _, tt := range positiveTests {
		t.Run("ok/"+tt.name, func(t *testing.T) {
			cert, err := ParseAndValidateCertificate(tt.certPem, VerifyCAExtensions)
			require.NoError(t, err)
			require.NotNil(t, cert)
		})
	}

	// Negative fixtures: legacy roots that still violate the strict profile —
	// they're kept on disk because many tests depend on them. They exercise the
	// "is not a CA / missing critical / wrong KU" failure modes.
	type negCase struct {
		name         string
		certPem      string
		expectSubstr string
	}
	negativeTests := []negCase{
		// NOTE: RootCertPem and NocRootCert1 were previously non-compliant
		// fixtures (KU missing / BC not critical). They have since been
		// regenerated to pass the strict profile so they can flow through
		// MsgProposeAddX509RootCert and MsgAddNocX509RootCert under
		// VerifyCAExtensions. The negative coverage for "BC not critical",
		// "KU not critical", and "KU missing-bit" is now provided by the
		// synthetic mutation block below.
		{
			name:         "leaf (cA=FALSE)",
			certPem:      testconstants.LeafCertPem,
			expectSubstr: "BasicConstraints extension must be present and cA must be set to TRUE",
		},
	}
	for _, tt := range negativeTests {
		t.Run("reject/"+tt.name, func(t *testing.T) {
			cert, err := ParseAndValidateCertificate(tt.certPem, VerifyCAExtensions)
			require.Error(t, err)
			require.Nil(t, cert)
			require.ErrorIs(t, err, pkitypes.ErrInappropriateCertificateType)
			require.Contains(t, err.Error(), tt.expectSubstr)
		})
	}

	// Synthetic edge cases — start from a compliant cert and flip just one
	// invariant at a time, so each branch of VerifyCAExtensions has dedicated
	// coverage independent of which fixtures happen to violate which rules.
	flipCases := []struct {
		name         string
		mutate       func(*x509std.Certificate)
		expectSubstr string
	}{
		{
			name: "KU disallowed bit set",
			mutate: func(c *x509std.Certificate) {
				c.KeyUsage |= x509std.KeyUsageDataEncipherment
			},
			expectSubstr: "SHALL NOT include bits other than",
		},
		{
			name: "KU keyCertSign missing",
			mutate: func(c *x509std.Certificate) {
				c.KeyUsage &^= x509std.KeyUsageCertSign
			},
			expectSubstr: "keyCertSign and cRLSign",
		},
		{
			name: "BC not critical",
			mutate: func(c *x509std.Certificate) {
				for i := range c.Extensions {
					if c.Extensions[i].Id.String() == "2.5.29.19" {
						c.Extensions[i].Critical = false

						return
					}
				}
				panic("test fixture has no BasicConstraints extension")
			},
			expectSubstr: "BasicConstraints extension SHALL be marked critical",
		},
		{
			name: "KU not critical",
			mutate: func(c *x509std.Certificate) {
				for i := range c.Extensions {
					if c.Extensions[i].Id.String() == "2.5.29.15" {
						c.Extensions[i].Critical = false

						return
					}
				}
				panic("test fixture has no KeyUsage extension")
			},
			expectSubstr: "KeyUsage extension SHALL be marked critical",
		},
	}
	for _, tt := range flipCases {
		t.Run("reject/synthetic/"+tt.name, func(t *testing.T) {
			cert, err := ParseAndValidateCertificate(testconstants.PAACertWithNumericVid)
			require.NoError(t, err)
			tt.mutate(cert.Certificate)

			err = VerifyCAExtensions(cert.Certificate)
			require.Error(t, err)
			require.ErrorIs(t, err, pkitypes.ErrInappropriateCertificateType)
			require.Contains(t, err.Error(), tt.expectSubstr)
		})
	}
}

func Test_ParseAndValidateCertificate_verifyDACExtensions(t *testing.T) {
	// Positive case: a DAC-shaped cert (BC critical + cA=FALSE, KU critical
	// with exactly digitalSignature, SKI + AKI present).
	t.Run("ok/DAC-shaped leaf", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.MatterDACShaped, verifyDACExtensions)
		require.NoError(t, err)
		require.NotNil(t, cert)
	})

	// LeafCertPem was regenerated to satisfy the DAC profile so the DA-chain
	// handler can dispatch on cA and apply verifyDACExtensions to leaves.
	t.Run("ok/LeafCertPem", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.LeafCertPem, verifyDACExtensions)
		require.NoError(t, err)
		require.NotNil(t, cert)
	})

	// Negative cases that exercise each branch of the helper.
	negativeTests := []struct {
		name         string
		certPem      string
		mutate       func(*x509std.Certificate) // optional; applied after parse
		expectSubstr string
	}{
		{
			name:         "CA cert (cA=TRUE)",
			certPem:      testconstants.PAACertWithNumericVid,
			expectSubstr: "DAC: BasicConstraints cA SHALL be set to FALSE",
		},
		{
			name:         "BC extension absent",
			certPem:      testconstants.LeafCertWithoutBasicConstraints,
			expectSubstr: "DAC: BasicConstraints extension SHALL be present",
		},
		{
			// Synthetic: take a DAC-shaped leaf and OR an extra KU bit (cRLSign)
			// — DAC profile (§6.2.2.3) bans any KU bit other than
			// digitalSignature, so this must trip the "exactly digitalSignature"
			// rule.
			name:    "synthetic: KU has cRLSign bit",
			certPem: testconstants.MatterDACShaped,
			mutate: func(c *x509std.Certificate) {
				c.KeyUsage |= x509std.KeyUsageCRLSign
			},
			expectSubstr: "DAC: KeyUsage SHALL be exactly digitalSignature",
		},
		{
			name:    "synthetic: SKI cleared",
			certPem: testconstants.MatterDACShaped,
			mutate: func(c *x509std.Certificate) {
				c.SubjectKeyId = nil
			},
			expectSubstr: "DAC: SubjectKeyIdentifier extension SHALL be present",
		},
		{
			name:    "synthetic: AKI cleared",
			certPem: testconstants.MatterDACShaped,
			mutate: func(c *x509std.Certificate) {
				c.AuthorityKeyId = nil
			},
			expectSubstr: "DAC: AuthorityKeyIdentifier extension SHALL be present",
		},
	}
	for _, tt := range negativeTests {
		t.Run("reject/"+tt.name, func(t *testing.T) {
			if tt.mutate == nil {
				cert, err := ParseAndValidateCertificate(tt.certPem, verifyDACExtensions)
				require.Error(t, err)
				require.Nil(t, cert)
				require.ErrorIs(t, err, pkitypes.ErrInappropriateCertificateType)
				require.Contains(t, err.Error(), tt.expectSubstr)

				return
			}
			cert, err := ParseAndValidateCertificate(tt.certPem)
			require.NoError(t, err)
			tt.mutate(cert.Certificate)
			err = verifyDACExtensions(cert.Certificate)
			require.Error(t, err)
			require.ErrorIs(t, err, pkitypes.ErrInappropriateCertificateType)
			require.Contains(t, err.Error(), tt.expectSubstr)
		})
	}
}

func Test_ParseAndValidateCertificate_verifyNOCExtensions(t *testing.T) {
	// Positive: a NOC-shaped cert with EKU critical and {serverAuth, clientAuth}.
	t.Run("ok/NOC-shaped leaf", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.MatterNOCShaped, verifyNOCExtensions)
		require.NoError(t, err)
		require.NotNil(t, cert)
	})

	// Negative: DAC-shaped leaf has no EKU → fails at the EKU presence check.
	t.Run("reject/DAC-shaped (no EKU)", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.MatterDACShaped, verifyNOCExtensions)
		require.Error(t, err)
		require.Nil(t, cert)
		require.ErrorIs(t, err, pkitypes.ErrInappropriateCertificateType)
		require.Contains(t, err.Error(), "NOC: ExtendedKeyUsage extension SHALL be present")
	})

	// Negative synthetic: NOC with extra EKU entry.
	t.Run("reject/synthetic NOC with 3 EKUs", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.MatterNOCShaped)
		require.NoError(t, err)
		cert.Certificate.ExtKeyUsage = append(cert.Certificate.ExtKeyUsage, x509std.ExtKeyUsageCodeSigning)
		err = verifyNOCExtensions(cert.Certificate)
		require.Error(t, err)
		require.ErrorIs(t, err, pkitypes.ErrInappropriateCertificateType)
		require.Contains(t, err.Error(), "NOC: ExtendedKeyUsage SHALL be exactly {serverAuth, clientAuth}")
	})

	// Negative: CA cert with cA=TRUE.
	t.Run("reject/CA cert", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.PAACertWithNumericVid, verifyNOCExtensions)
		require.Error(t, err)
		require.Nil(t, cert)
		require.ErrorIs(t, err, pkitypes.ErrInappropriateCertificateType)
		require.Contains(t, err.Error(), "NOC: BasicConstraints cA SHALL be set to FALSE")
	})
}

// VerifyDAChainNonRoot dispatches on cA: PAIs go to VerifyCAExtensions + the
// §6.2.2.4 pathLen=0 rule, DACs go to verifyDACExtensions. Both branches must
// accept their respective fixtures and reject the opposite shape.
func Test_ParseAndValidateCertificate_VerifyDAChainNonRoot(t *testing.T) {
	t.Run("ok/PAI (cA=TRUE, pathLen=0)", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.PAICertWithNumericVid, VerifyDAChainNonRoot)
		require.NoError(t, err)
		require.NotNil(t, cert)
	})

	t.Run("ok/IntermediateCertPem (DA chain ICA)", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.IntermediateCertPem, VerifyDAChainNonRoot)
		require.NoError(t, err)
		require.NotNil(t, cert)
	})

	t.Run("ok/LeafCertPem (DAC-shaped)", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.LeafCertPem, VerifyDAChainNonRoot)
		require.NoError(t, err)
		require.NotNil(t, cert)
	})

	t.Run("ok/MatterDACShaped", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.MatterDACShaped, VerifyDAChainNonRoot)
		require.NoError(t, err)
		require.NotNil(t, cert)
	})

	t.Run("reject/BC-extension-absent", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.LeafCertWithoutBasicConstraints, VerifyDAChainNonRoot)
		require.Error(t, err)
		require.Nil(t, cert)
		require.ErrorIs(t, err, pkitypes.ErrInappropriateCertificateType)
		require.Contains(t, err.Error(), "BasicConstraints extension SHALL be present")
	})
}

// VerifyECDSAP256SHA256 must accept every fixture currently in the codebase
// (all are ECDSA-with-SHA256 on P-256) and reject any mutation that changes
// either the signature algorithm or the public-key curve.
func Test_ParseAndValidateCertificate_VerifyECDSAP256SHA256(t *testing.T) {
	positiveTests := []struct {
		name    string
		certPem string
	}{
		{name: "PAA with VID", certPem: testconstants.PAACertWithNumericVid},
		{name: "PAI", certPem: testconstants.PAICertWithNumericVid},
		{name: "DAC-shaped leaf", certPem: testconstants.MatterDACShaped},
		{name: "NOC-shaped leaf", certPem: testconstants.MatterNOCShaped},
		{name: "RCAC NocRootCert1", certPem: testconstants.NocRootCert1},
		{name: "ICAC NocCert1", certPem: testconstants.NocCert1},
		{name: "NOC leaf NocLeafCert1", certPem: testconstants.NocLeafCert1},
		{name: "LeafCertPem", certPem: testconstants.LeafCertPem},
	}
	for _, tt := range positiveTests {
		t.Run("ok/"+tt.name, func(t *testing.T) {
			cert, err := ParseAndValidateCertificate(tt.certPem, VerifyECDSAP256SHA256)
			require.NoError(t, err)
			require.NotNil(t, cert)
		})
	}

	t.Run("reject/non-ECDSA-SHA256 signature", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.PAACertWithNumericVid)
		require.NoError(t, err)
		cert.Certificate.SignatureAlgorithm = x509std.ECDSAWithSHA384
		err = VerifyECDSAP256SHA256(cert.Certificate)
		require.Error(t, err)
		require.ErrorIs(t, err, pkitypes.ErrInvalidCertificate)
		require.Contains(t, err.Error(), "ecdsa-with-SHA256")
	})

	t.Run("reject/non-ECDSA public key", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.PAACertWithNumericVid)
		require.NoError(t, err)
		cert.Certificate.PublicKey = struct{}{}
		err = VerifyECDSAP256SHA256(cert.Certificate)
		require.Error(t, err)
		require.ErrorIs(t, err, pkitypes.ErrInvalidCertificate)
		require.Contains(t, err.Error(), "ECDSA on prime256v1")
	})
}

func Test_VerifyVidPidConsistency(t *testing.T) {
	t.Run("ok/parent has no VID — rule does not fire", func(t *testing.T) {
		// Spec only constrains the child when the parent carries the attribute.
		require.NoError(t, VerifyVidPidConsistency(
			"O=child,vid=0xFFF1",
			"O=parent",
		))
	})

	t.Run("ok/matching VID", func(t *testing.T) {
		require.NoError(t, VerifyVidPidConsistency(
			"CN=Matter Test DAC,vid=0xFFF1",
			"CN=Matter Test PAI,vid=0xFFF1",
		))
	})

	t.Run("ok/matching VID + matching PID", func(t *testing.T) {
		require.NoError(t, VerifyVidPidConsistency(
			"CN=Matter Test DAC,pid=0x8000,vid=0xFFF1",
			"CN=Matter Test PAI,pid=0x8000,vid=0xFFF1",
		))
	})

	t.Run("ok/parent has no PID — PID rule does not fire", func(t *testing.T) {
		require.NoError(t, VerifyVidPidConsistency(
			"CN=DAC,pid=0x1234,vid=0xFFF1",
			"CN=PAI,vid=0xFFF1",
		))
	})

	t.Run("reject/child VID mismatches parent VID", func(t *testing.T) {
		err := VerifyVidPidConsistency(
			"CN=DAC,vid=0xFFF2",
			"CN=PAI,vid=0xFFF1",
		)
		require.Error(t, err)
		require.ErrorIs(t, err, pkitypes.ErrCertVidNotEqualToIssuerVid)
	})

	t.Run("reject/parent has VID, child missing VID", func(t *testing.T) {
		// Child VID parsed as 0; parent's non-zero VID means mismatch.
		err := VerifyVidPidConsistency(
			"O=DAC",
			"CN=PAI,vid=0xFFF1",
		)
		require.Error(t, err)
		require.ErrorIs(t, err, pkitypes.ErrCertVidNotEqualToIssuerVid)
	})

	t.Run("reject/PID mismatch", func(t *testing.T) {
		err := VerifyVidPidConsistency(
			"CN=DAC,pid=0x8001,vid=0xFFF1",
			"CN=PAI,pid=0x8000,vid=0xFFF1",
		)
		require.Error(t, err)
		require.ErrorIs(t, err, pkitypes.ErrCertPidNotEqualToIssuerPid)
	})
}

// VerifyCAExtensions now also asserts SKI presence (always) and AKI presence
// (only for non-self-signed CAs). Synthetic mutations exercise each branch.
func Test_VerifyCAExtensions_SKIAndAKI(t *testing.T) {
	t.Run("reject/SKI absent on self-signed CA", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.PAACertWithNumericVid)
		require.NoError(t, err)
		cert.Certificate.SubjectKeyId = nil
		err = VerifyCAExtensions(cert.Certificate)
		require.Error(t, err)
		require.ErrorIs(t, err, pkitypes.ErrInappropriateCertificateType)
		require.Contains(t, err.Error(), "SubjectKeyIdentifier extension SHALL be present")
	})

	t.Run("reject/AKI absent on non-self-signed CA", func(t *testing.T) {
		// IntermediateCertPem is a non-self-signed PAI; clearing its AKI must
		// trip the new presence rule.
		cert, err := ParseAndValidateCertificate(testconstants.IntermediateCertPem)
		require.NoError(t, err)
		cert.Certificate.AuthorityKeyId = nil
		err = VerifyCAExtensions(cert.Certificate)
		require.Error(t, err)
		require.ErrorIs(t, err, pkitypes.ErrInappropriateCertificateType)
		require.Contains(t, err.Error(), "AuthorityKeyIdentifier extension SHALL be present")
	})

	t.Run("ok/AKI absent on self-signed CA (PAA)", func(t *testing.T) {
		// PAAs MAY omit AKI per §6.2.2.5; the rule only fires for non-self-signed.
		cert, err := ParseAndValidateCertificate(testconstants.RootCertWithVid, VerifyCAExtensions)
		require.NoError(t, err)
		require.NotNil(t, cert)
	})
}

// VerifyNoEKU rejects certs that carry an ExtendedKeyUsage extension. All
// current RCAC/ICAC fixtures comply (none encode EKU). Synthetic mutations
// exercise the rejection.
func Test_VerifyNoEKU(t *testing.T) {
	t.Run("ok/RCAC NocRootCert1 has no EKU", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.NocRootCert1, VerifyNoEKU)
		require.NoError(t, err)
		require.NotNil(t, cert)
	})

	t.Run("ok/ICAC NocCert1 has no EKU", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.NocCert1, VerifyNoEKU)
		require.NoError(t, err)
		require.NotNil(t, cert)
	})

	t.Run("reject/synthetic EKU added", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.NocRootCert1)
		require.NoError(t, err)
		cert.Certificate.Extensions = append(cert.Certificate.Extensions, pkix.Extension{
			Id:       asn1.ObjectIdentifier{2, 5, 29, 37},
			Critical: false,
			Value:    []byte{0x30, 0x00},
		})
		err = VerifyNoEKU(cert.Certificate)
		require.Error(t, err)
		require.ErrorIs(t, err, pkitypes.ErrInappropriateCertificateType)
		require.Contains(t, err.Error(), "ExtendedKeyUsage extension SHALL NOT be present")
	})
}

func Test_VerifyVersionV3(t *testing.T) {
	t.Run("ok/RootCertPem", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.RootCertPem, VerifyVersionV3)
		require.NoError(t, err)
		require.NotNil(t, cert)
	})

	t.Run("reject/synthetic v1", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.PAACertWithNumericVid)
		require.NoError(t, err)
		cert.Certificate.Version = 1
		err = VerifyVersionV3(cert.Certificate)
		require.Error(t, err)
		require.ErrorIs(t, err, pkitypes.ErrInvalidCertificate)
		require.Contains(t, err.Error(), "v3")
	})
}

func Test_VerifyAtMostOneVIDAndPID(t *testing.T) {
	t.Run("ok/cert with one VID + one PID", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.PAICertWithNumericPidVid, VerifyAtMostOneVIDAndPID)
		require.NoError(t, err)
		require.NotNil(t, cert)
	})

	t.Run("ok/cert with no VID and no PID", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.RootCertPem, VerifyAtMostOneVIDAndPID)
		require.NoError(t, err)
		require.NotNil(t, cert)
	})

	t.Run("reject/duplicate VID in subject", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.PAACertWithNumericVid)
		require.NoError(t, err)
		// Duplicate the matter-vid attribute that's already in the subject.
		var vid pkix.AttributeTypeAndValue
		for _, n := range cert.Certificate.Subject.Names {
			if n.Type.String() == "1.3.6.1.4.1.37244.2.1" {
				vid = n

				break
			}
		}
		require.NotZero(t, vid.Type, "fixture must already carry a matter-vid")
		cert.Certificate.Subject.Names = append(cert.Certificate.Subject.Names, vid)
		err = VerifyAtMostOneVIDAndPID(cert.Certificate)
		require.Error(t, err)
		require.ErrorIs(t, err, pkitypes.ErrInvalidCertificate)
		require.Contains(t, err.Error(), "subject SHALL contain at most one matter-vid")
	})

	t.Run("reject/duplicate PID in issuer", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.PAICertWithNumericPidVid)
		require.NoError(t, err)
		cert.Certificate.Issuer.Names = append(cert.Certificate.Issuer.Names,
			pkix.AttributeTypeAndValue{
				Type:  asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 37244, 2, 2},
				Value: "8000",
			},
			pkix.AttributeTypeAndValue{
				Type:  asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 37244, 2, 2},
				Value: "8001",
			},
		)
		err = VerifyAtMostOneVIDAndPID(cert.Certificate)
		require.Error(t, err)
		require.ErrorIs(t, err, pkitypes.ErrInvalidCertificate)
		require.Contains(t, err.Error(), "issuer SHALL contain at most one matter-pid")
	})
}

func Test_VerifyNoPIDInSubject(t *testing.T) {
	t.Run("ok/PAACertWithNumericVid (no PID)", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.PAACertWithNumericVid, VerifyNoPIDInSubject)
		require.NoError(t, err)
		require.NotNil(t, cert)
	})

	t.Run("ok/RootCertPem (no PID)", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.RootCertPem, VerifyNoPIDInSubject)
		require.NoError(t, err)
		require.NotNil(t, cert)
	})

	t.Run("reject/PAI cert with PID is rejected if used as PAA", func(t *testing.T) {
		// PAICertWithNumericPidVid carries pid=0x8000. The rule applies only to
		// the PAA add path; this confirms the helper trips on a PID-bearing DN.
		cert, err := ParseAndValidateCertificate(testconstants.PAICertWithNumericPidVid, VerifyNoPIDInSubject)
		require.Error(t, err)
		require.Nil(t, cert)
		require.ErrorIs(t, err, pkitypes.ErrNotEmptyPid)
	})
}
