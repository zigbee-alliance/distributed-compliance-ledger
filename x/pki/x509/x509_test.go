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

func Test_ParseAndValidateCertificate_VerifyIsCACertificate(t *testing.T) {
	positiveTests := []struct {
		name    string
		certPem string
	}{
		{name: "self-signed root CA", certPem: testconstants.RootCertPem},
		{name: "non-self-signed intermediate CA", certPem: testconstants.IntermediateCertPem},
		{name: "PAA with VID (CA)", certPem: testconstants.PAACertWithNumericVid},
	}

	for _, tt := range positiveTests {
		t.Run("ok/"+tt.name, func(t *testing.T) {
			cert, err := ParseAndValidateCertificate(tt.certPem, VerifyIsCACertificate)
			require.NoError(t, err)
			require.NotNil(t, cert)
		})
	}

	negativeTests := []struct {
		name    string
		certPem string
	}{
		// BasicConstraintsValid=true, IsCA=false — explicit non-CA end-entity
		{name: "leaf certificate (IsCA=false)", certPem: testconstants.LeafCertPem},
	}

	for _, tt := range negativeTests {
		t.Run("reject/"+tt.name, func(t *testing.T) {
			cert, err := ParseAndValidateCertificate(tt.certPem, VerifyIsCACertificate)
			require.Error(t, err)
			require.Nil(t, cert)
			require.ErrorIs(t, err, pkitypes.ErrInappropriateCertificateType)
		})
	}

	// Without the option, the same leaf certificate must parse successfully —
	// confirms the rejection is driven by VerifyIsCACertificate, not by some
	// other validation that already existed.
	t.Run("leaf passes without VerifyIsCACertificate", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.LeafCertPem)
		require.NoError(t, err)
		require.NotNil(t, cert)
	})
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
	// Positive fixtures: known CA certs that comply with the full Matter R1.5
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
		massage      func(*x509std.Certificate)
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

func Test_ParseAndValidateCertificate_VerifyDACExtensions(t *testing.T) {
	// Positive case: a DAC-shaped cert (BC critical + cA=FALSE, KU critical
	// with exactly digitalSignature, SKI + AKI present).
	t.Run("ok/DAC-shaped leaf", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.MatterDACShaped, VerifyDACExtensions)
		require.NoError(t, err)
		require.NotNil(t, cert)
	})

	// LeafCertPem was regenerated to satisfy the DAC profile so the DA-chain
	// handler can dispatch on cA and apply VerifyDACExtensions to leaves.
	t.Run("ok/LeafCertPem", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.LeafCertPem, VerifyDACExtensions)
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
				cert, err := ParseAndValidateCertificate(tt.certPem, VerifyDACExtensions)
				require.Error(t, err)
				require.Nil(t, cert)
				require.ErrorIs(t, err, pkitypes.ErrInappropriateCertificateType)
				require.Contains(t, err.Error(), tt.expectSubstr)
				return
			}
			cert, err := ParseAndValidateCertificate(tt.certPem)
			require.NoError(t, err)
			tt.mutate(cert.Certificate)
			err = VerifyDACExtensions(cert.Certificate)
			require.Error(t, err)
			require.ErrorIs(t, err, pkitypes.ErrInappropriateCertificateType)
			require.Contains(t, err.Error(), tt.expectSubstr)
		})
	}
}

func Test_ParseAndValidateCertificate_VerifyNOCExtensions(t *testing.T) {
	// Positive: a NOC-shaped cert with EKU critical and {serverAuth, clientAuth}.
	t.Run("ok/NOC-shaped leaf", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.MatterNOCShaped, VerifyNOCExtensions)
		require.NoError(t, err)
		require.NotNil(t, cert)
	})

	// Negative: DAC-shaped leaf has no EKU → fails at the EKU presence check.
	t.Run("reject/DAC-shaped (no EKU)", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.MatterDACShaped, VerifyNOCExtensions)
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
		err = VerifyNOCExtensions(cert.Certificate)
		require.Error(t, err)
		require.ErrorIs(t, err, pkitypes.ErrInappropriateCertificateType)
		require.Contains(t, err.Error(), "NOC: ExtendedKeyUsage SHALL be exactly {serverAuth, clientAuth}")
	})

	// Negative: CA cert with cA=TRUE.
	t.Run("reject/CA cert", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.PAACertWithNumericVid, VerifyNOCExtensions)
		require.Error(t, err)
		require.Nil(t, cert)
		require.ErrorIs(t, err, pkitypes.ErrInappropriateCertificateType)
		require.Contains(t, err.Error(), "NOC: BasicConstraints cA SHALL be set to FALSE")
	})
}

// VerifyDAChainNonRoot dispatches on cA: PAIs go to VerifyCAExtensions + the
// §6.2.2.4 pathLen=0 rule, DACs go to VerifyDACExtensions. Both branches must
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

// VerifyNOCChainNonRoot dispatches on cA: ICACs go to VerifyCAExtensions, NOCs
// go to VerifyNOCExtensions. Both branches must accept their respective fixtures.
func Test_ParseAndValidateCertificate_VerifyNOCChainNonRoot(t *testing.T) {
	t.Run("ok/NocCert1 (ICAC)", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.NocCert1, VerifyNOCChainNonRoot)
		require.NoError(t, err)
		require.NotNil(t, cert)
	})

	t.Run("ok/NocLeafCert1 (NOC end-entity)", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.NocLeafCert1, VerifyNOCChainNonRoot)
		require.NoError(t, err)
		require.NotNil(t, cert)
	})

	t.Run("ok/MatterNOCShaped", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.MatterNOCShaped, VerifyNOCChainNonRoot)
		require.NoError(t, err)
		require.NotNil(t, cert)
	})

	t.Run("reject/BC-extension-absent", func(t *testing.T) {
		cert, err := ParseAndValidateCertificate(testconstants.LeafCertWithoutBasicConstraints, VerifyNOCChainNonRoot)
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
