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
	"testing"
	"time"

	tmrand "github.com/cometbft/cometbft/libs/rand"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
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
		result := FormatOID(tt.header, tt.oldKey, tt.newKey)
		require.Equal(t, result, tt.result)
	}

	for _, tt := range negativeTests {
		result := FormatOID(tt.header, tt.oldKey, tt.newKey)
		require.Equal(t, result, tt.result)
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

	negativeTests := []struct {
		header string
		oldKey string
		newKey string
		result string
	}{
		// set incorrect oldKey
		{
			header: "CN=Matter PAA 1,O=Google,C=US,1.3.6.1.4.1.37244.2.1=#130436303036",
			oldKey: "1.3.6.1.4.1.37244.2.2",
			newKey: "vid",
			result: "CN=Matter PAA 1,O=Google,C=US,1.3.6.1.4.1.37244.2.1=#130436303036",
		},
	}

	for _, tt := range positiveTests {
		result := FormatOID(tt.header, tt.oldKey, tt.newKey)
		require.Equal(t, result, tt.result)
	}

	for _, tt := range negativeTests {
		result := FormatOID(tt.header, tt.oldKey, tt.newKey)
		require.Equal(t, result, tt.result)
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

// Test_FormatOID_HexDecodeError reproduces the bug fixed in this PR: previously
// FormatOID() discarded the error from hex.DecodeString() with `decoded, _ := ...`,
// causing malformed hex values in certificate subject fields to silently produce
// corrupt decoded output (e.g. "vid=0x" with empty bytes). In a compliance ledger
// this is a security concern because an attacker-controlled or accidentally
// truncated certificate could yield bogus normalized OIDs.
//
// After the fix, malformed entries are skipped: the original value remains in
// the output unchanged.
func Test_FormatOID_HexDecodeError(t *testing.T) {
	tests := []struct {
		name   string
		header string
		oldKey string
		newKey string
		// Expected output: original entry must be preserved unchanged when
		// hex.DecodeString cannot parse the trailing 8 chars.
		expected string
	}{
		{
			name:     "non-hex characters in last 8 chars are skipped",
			header:   "CN=Matter PAA 1,O=Google,C=US,1.3.6.1.4.1.37244.2.1=#1304ZZNOTHEX",
			oldKey:   "1.3.6.1.4.1.37244.2.1",
			newKey:   "vid",
			expected: "CN=Matter PAA 1,O=Google,C=US,1.3.6.1.4.1.37244.2.1=#1304ZZNOTHEX",
		},
		{
			name:     "odd-length hex (invalid) is skipped",
			header:   "CN=Matter PAA 1,1.3.6.1.4.1.37244.2.1=#13043600060",
			oldKey:   "1.3.6.1.4.1.37244.2.1",
			newKey:   "vid",
			expected: "CN=Matter PAA 1,1.3.6.1.4.1.37244.2.1=#13043600060",
		},
		{
			name:     "value containing whitespace yields decode error",
			header:   "CN=Matter PAA 1,1.3.6.1.4.1.37244.2.1=# 1304600",
			oldKey:   "1.3.6.1.4.1.37244.2.1",
			newKey:   "vid",
			expected: "CN=Matter PAA 1,1.3.6.1.4.1.37244.2.1=# 1304600",
		},
		{
			name:     "valid hex still works (regression guard for positive path)",
			header:   "CN=Matter PAA 1,O=Google,C=US,1.3.6.1.4.1.37244.2.1=#130436303036",
			oldKey:   "1.3.6.1.4.1.37244.2.1",
			newKey:   "vid",
			expected: "CN=Matter PAA 1,O=Google,C=US,vid=0x6006",
		},
		{
			name:     "multiple OIDs - one valid, one invalid",
			header:   "CN=Test,1.3.6.1.4.1.37244.2.1=#130436303036,1.3.6.1.4.1.37244.2.2=#1304ZZZZZZ",
			oldKey:   "1.3.6.1.4.1.37244.2.2",
			newKey:   "pid",
			expected: "CN=Test,1.3.6.1.4.1.37244.2.1=#130436303036,1.3.6.1.4.1.37244.2.2=#1304ZZZZZZ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatOID(tt.header, tt.oldKey, tt.newKey)
			require.Equal(t, tt.expected, result,
				"FormatOID must NOT silently emit corrupted output when hex.DecodeString fails")
		})
	}
}

// Test_FormatOID_NoCorruptOutput is an explicit assertion that the post-fix
// behaviour cannot produce the buggy output pattern "newKey=0x" (empty hex)
// which the pre-fix code could emit when hex.DecodeString returned an error.
func Test_FormatOID_NoCorruptOutput(t *testing.T) {
	// Crafted input where the last 8 chars contain a non-hex sentinel.
	header := "CN=Test,1.3.6.1.4.1.37244.2.1=#1304NOTHEXX"
	result := FormatOID(header, "1.3.6.1.4.1.37244.2.1", "vid")

	require.NotContains(t, result, "vid=0x ",
		"output must not contain 'vid=0x' followed by empty/whitespace bytes")
	require.NotContains(t, result, "vid=0x\x00",
		"output must not contain a null-byte after the hex prefix")
	// Original entry must survive untouched.
	require.Contains(t, result, "1.3.6.1.4.1.37244.2.1=#1304NOTHEXX",
		"malformed entry must be preserved verbatim, not silently rewritten")
}
