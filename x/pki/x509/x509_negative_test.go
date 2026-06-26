//nolint:testpackage
package x509

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
)

// decodeStd decodes a PEM fixture and returns the underlying crypto/x509 cert
// so individual parsed fields can be mutated to exercise one profile rule at a
// time (the validators read the parsed struct, not the signature).
func decodeStd(t *testing.T, pemCert string) *x509.Certificate {
	t.Helper()
	c, err := DecodeX509Certificate(pemCert)
	require.NoError(t, err)

	return c.Certificate
}

func removeExt(cert *x509.Certificate, oid asn1.ObjectIdentifier) {
	out := cert.Extensions[:0]
	for _, e := range cert.Extensions {
		if !e.Id.Equal(oid) {
			out = append(out, e)
		}
	}
	cert.Extensions = out
}

func setExtCriticalFlag(cert *x509.Certificate, oid asn1.ObjectIdentifier, critical bool) {
	for i := range cert.Extensions {
		if cert.Extensions[i].Id.Equal(oid) {
			cert.Extensions[i].Critical = critical

			return
		}
	}
}

func TestVerifyCAExtensions_NegativeBranches(t *testing.T) {
	t.Run("KeyUsage absent", func(t *testing.T) {
		cert := decodeStd(t, testconstants.PAACertWithNumericVid)
		removeExt(cert, OIDKeyUsage)
		require.Error(t, VerifyCAExtensions(cert))
	})
	t.Run("SubjectKeyId wrong length", func(t *testing.T) {
		cert := decodeStd(t, testconstants.PAACertWithNumericVid)
		cert.SubjectKeyId = []byte{0x01, 0x02, 0x03}
		require.Error(t, VerifyCAExtensions(cert))
	})
	t.Run("non-self-signed AuthorityKeyId wrong length", func(t *testing.T) {
		cert := decodeStd(t, testconstants.IntermediateCertPem) // non-self-signed CA
		cert.AuthorityKeyId = []byte{0x01, 0x02, 0x03}
		require.Error(t, VerifyCAExtensions(cert))
	})
}

func TestVerifyEndEntityExtensions_NegativeBranches(t *testing.T) {
	base := func(t *testing.T) *x509.Certificate {
		t.Helper()

		return decodeStd(t, testconstants.MatterDACShaped)
	}
	t.Run("BasicConstraints not critical", func(t *testing.T) {
		cert := base(t)
		setExtCriticalFlag(cert, OIDBasicConstraints, false)
		require.Error(t, verifyEndEntityExtensions(cert, "DAC"))
	})
	t.Run("KeyUsage absent", func(t *testing.T) {
		cert := base(t)
		removeExt(cert, OIDKeyUsage)
		require.Error(t, verifyEndEntityExtensions(cert, "DAC"))
	})
	t.Run("KeyUsage not critical", func(t *testing.T) {
		cert := base(t)
		setExtCriticalFlag(cert, OIDKeyUsage, false)
		require.Error(t, verifyEndEntityExtensions(cert, "DAC"))
	})
	t.Run("SubjectKeyId wrong length", func(t *testing.T) {
		cert := base(t)
		cert.SubjectKeyId = []byte{0x01}
		require.Error(t, verifyEndEntityExtensions(cert, "DAC"))
	})
	t.Run("AuthorityKeyId wrong length", func(t *testing.T) {
		cert := base(t)
		cert.AuthorityKeyId = []byte{0x01}
		require.Error(t, verifyEndEntityExtensions(cert, "DAC"))
	})
}

func TestVerifyDAChainNonRoot_CAExtensionsFail(t *testing.T) {
	cert := decodeStd(t, testconstants.PAICertWithNumericVid) // cA=TRUE
	removeExt(cert, OIDKeyUsage)
	require.Error(t, VerifyDAChainNonRoot(cert))
}

func TestVerifyPAAPathLen_PresentNotOne(t *testing.T) {
	cert := decodeStd(t, testconstants.PAACertWithNumericVid)
	cert.MaxPathLen = 2
	cert.MaxPathLenZero = false
	require.Error(t, VerifyPAAPathLen(cert))
}

func TestVerifyECDSAP256SHA256_WrongCurve(t *testing.T) {
	cert := decodeStd(t, testconstants.PAACertWithNumericVid)
	pub, ok := cert.PublicKey.(*ecdsa.PublicKey)
	require.True(t, ok)
	pub.Curve = elliptic.P224()
	cert.PublicKey = pub
	require.Error(t, VerifyECDSAP256SHA256(cert))
}

func TestVerifyVidPidConsistency_Branches(t *testing.T) {
	// parent VID malformed
	require.Error(t, VerifyVidPidConsistency("vid=0x1", "vid=0xZZ"))
	// child VID differs from parent VID
	require.Error(t, VerifyVidPidConsistency("vid=0x2", "vid=0x1"))
	// parent PID malformed
	require.Error(t, VerifyVidPidConsistency("pid=0x1", "pid=0xZZ"))
	// child PID differs from parent PID
	require.Error(t, VerifyVidPidConsistency("pid=0x2", "pid=0x1"))
	// happy path: parent carries neither -> no constraint
	require.NoError(t, VerifyVidPidConsistency("vid=0x2,pid=0x5", "CN=parent"))
}

func TestFormatOID_Branches(t *testing.T) {
	const vidOID = "1.3.6.1.4.1.37244.2.1"

	// value already in readable form (no '#') -> returned untouched
	out, err := FormatOID(vidOID+"=0x6006", vidOID, "vid")
	require.NoError(t, err)
	require.Equal(t, vidOID+"=0x6006", out) // no '#': left as-is

	// non-hex DER
	_, err = FormatOID(vidOID+"=#ZZ", vidOID, "vid")
	require.Error(t, err)

	// DER too short (<2 bytes)
	_, err = FormatOID(vidOID+"=#13", vidOID, "vid")
	require.Error(t, err)

	// wrong DER tag (0x05 != PrintableString/UTF8String)
	_, err = FormatOID(vidOID+"=#0500", vidOID, "vid")
	require.Error(t, err)

	// DER length exceeds remaining bytes
	_, err = FormatOID(vidOID+"=#1305", vidOID, "vid")
	require.Error(t, err)
}

func TestGetIntValueFromSubject_BadHex(t *testing.T) {
	_, err := GetVidFromSubject("vid=0xZZ")
	require.Error(t, err)
}

func TestToSubjectAsText_MalformedVID(t *testing.T) {
	_, err := ToSubjectAsText("1.3.6.1.4.1.37244.2.1=#ZZ")
	require.Error(t, err)
}

func TestCertificatePEMsEqual_InvalidPEM(t *testing.T) {
	require.False(t, CertificatePEMsEqual("not a pem", testconstants.RootCertPem))
	require.True(t, CertificatePEMsEqual(testconstants.RootCertPem, testconstants.RootCertPem))
}

func TestDecodeX509Certificate_BadDER(t *testing.T) {
	const badDER = "-----BEGIN CERTIFICATE-----\nAQID\n-----END CERTIFICATE-----\n"
	_, err := DecodeX509Certificate(badDER)
	require.Error(t, err)
}

func TestVerifyVVSCSignature_Branches(t *testing.T) {
	leaf, err := DecodeX509Certificate(testconstants.LeafCertPem)
	require.NoError(t, err)
	intermediate, err := DecodeX509Certificate(testconstants.IntermediateCertPem)
	require.NoError(t, err)

	// not yet valid
	future := leaf.Certificate.NotBefore.Add(-1 * time.Hour)
	require.Error(t, leaf.VerifyVVSCSignature(intermediate, future))

	// expired
	expired := leaf.Certificate.NotAfter.Add(1 * time.Hour)
	require.Error(t, leaf.VerifyVVSCSignature(intermediate, expired))

	// signature mismatch: wrong parent (RootCertPem did not sign the leaf)
	root, err := DecodeX509Certificate(testconstants.RootCertPem)
	require.NoError(t, err)
	within := leaf.Certificate.NotBefore.Add(time.Hour)
	require.Error(t, leaf.VerifyVVSCSignature(root, within))
}

func TestVerifyVVSCExtensions_Wrapper(t *testing.T) {
	// Valid DAC-shaped cert satisfies the shared end-entity profile.
	require.NoError(t, VerifyVVSCExtensions(decodeStd(t, testconstants.MatterDACShaped)))
}

func TestVerifyPAIPathLen_Absent(t *testing.T) {
	// A PAA has no pathLenConstraint=0, so the PAI rule rejects it.
	cert := decodeStd(t, testconstants.PAACertWithNumericVid)
	cert.MaxPathLenZero = false
	require.Error(t, VerifyPAIPathLen(cert))
}

func TestVerifyVidPidConsistency_ChildVidMalformed(t *testing.T) {
	// parent carries a valid VID, child VID is malformed
	require.Error(t, VerifyVidPidConsistency("vid=0xZZ", "vid=0x1"))
	// parent carries a valid PID, child PID is malformed
	require.Error(t, VerifyVidPidConsistency("pid=0xZZ", "pid=0x1"))
}

func TestVerifyVVSCSignature_Success(t *testing.T) {
	leaf, err := DecodeX509Certificate(testconstants.LeafCertPem)
	require.NoError(t, err)
	intermediate, err := DecodeX509Certificate(testconstants.IntermediateCertPem)
	require.NoError(t, err)

	within := leaf.Certificate.NotBefore.Add(time.Hour)
	require.NoError(t, leaf.VerifyVVSCSignature(intermediate, within))
}

// ensure pkix import is used to keep the import meaningful if other refs change.
var _ = pkix.Name{}
