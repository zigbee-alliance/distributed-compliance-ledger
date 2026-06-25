package pki

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

const (
	rootWithSameSubjectAndSkid1Path = "../../constants/root_with_same_subject_and_skid_1"
	rootWithSameSubjectAndSkid2Path = "../../constants/root_with_same_subject_and_skid_2"
	// Subject and SKID match the actual "Example Company" cert files (not Amazon Root CA).
	rootWithSameSubjectAndSkidSubject      = "MIGCMQswCQYDVQQGEwJVUzERMA8GA1UECBMITmV3IFlvcmsxETAPBgNVBAcTCE5ldyBZb3JrMRgwFgYDVQQKEw9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsTEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMTD3d3dy5leGFtcGxlLmNvbQ=="
	rootWithSameSubjectAndSkidSubjectKeyID = "C1:48:66:ED:6F:23:D8:28:1A:D9:37:7C:58:AC:3F:DA:04:C1:41:E8"
)

// Tests that multiple certificates with the same subject and SKID can coexist.
func TestPKICombineCerts(t *testing.T) {
	jack := testconstants.JackAccount
	alice := testconstants.AliceAccount

	vid := 65521

	vendorAccount := fmt.Sprintf("vendor_account_%d", vid)
	cliputils.CreateVendorAccount(t, vendorAccount, vid)

	subj := rootWithSameSubjectAndSkidSubject
	skid := rootWithSameSubjectAndSkidSubjectKeyID

	// certsContainID reports whether a list of CertificateIdentifiers contains
	// (subj, skid).
	certsContainID := func(ids []*pkitypes.CertificateIdentifier) bool {
		for _, id := range ids {
			if id != nil && id.Subject == subj && id.SubjectKeyId == skid {
				return true
			}
		}

		return false
	}
	// certsContainSKID reports whether a list of Certificates contains skid.
	certsContainSKID := func(certs []*pkitypes.Certificate) bool {
		for _, c := range certs {
			if c != nil && c.SubjectKeyId == skid {
				return true
			}
		}

		return false
	}
	// approvedContain reports whether a list of ApprovedCertificates contains
	// (subj, skid).
	approvedContain := func(list []pkitypes.ApprovedCertificates) bool {
		for i := range list {
			if list[i].Subject == subj && list[i].SubjectKeyId == skid {
				return true
			}
		}

		return false
	}

	// Scoped to this test's own subject (the global all-* lists already hold
	// certs from earlier tests on the shared ledger, so only per-subject
	// emptiness can be asserted here).
	t.Run("QueryBeforeAdd_NotFound", func(t *testing.T) {
		x509, err := GetX509Cert(subj, skid)
		require.NoError(t, err)
		require.Nil(t, x509)

		noc, err := GetNocCert("--subject", subj, "--subject-key-id", skid)
		require.NoError(t, err)
		require.Nil(t, noc)

		proposed, err := GetProposedX509RootCert(subj, skid)
		require.NoError(t, err)
		require.Nil(t, proposed)

		revoked, err := GetRevokedX509Cert(subj, skid)
		require.NoError(t, err)
		require.Nil(t, revoked)

		bySubject, err := GetX509CertsBySubject(subj)
		require.NoError(t, err)
		require.Nil(t, bySubject)

		children, err := GetChildX509Certs(subj, skid)
		require.NoError(t, err)
		require.Nil(t, children)
	})

	t.Run("ProposeAndApproveFirstRootCert", func(t *testing.T) {
		txResult, err := ProposeAddX509RootCert(rootWithSameSubjectAndSkid1Path, jack, X509ProposeOpts{VID: vid})
		cliputils.RequireTxOK(t, txResult, err)

		// With 3 trustees, quorum=2: jack proposes + alice approves = cert is approved.
		txResult, err = ApproveAddX509RootCert(rootWithSameSubjectAndSkidSubject, rootWithSameSubjectAndSkidSubjectKeyID, alice)
		cliputils.RequireTxOK(t, txResult, err)

		cert, err := GetX509Cert(rootWithSameSubjectAndSkidSubject, rootWithSameSubjectAndSkidSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, cert)
		require.Equal(t, rootWithSameSubjectAndSkidSubject, cert.Subject)
	})

	t.Run("ProposeAndApproveSecondRootCert_SameSubjectSkid", func(t *testing.T) {
		txResult, err := ProposeAddX509RootCert(rootWithSameSubjectAndSkid2Path, jack, X509ProposeOpts{VID: vid})
		cliputils.RequireTxOK(t, txResult, err)

		// With 3 trustees, quorum=2: jack proposes + alice approves = cert is approved.
		txResult, err = ApproveAddX509RootCert(rootWithSameSubjectAndSkidSubject, rootWithSameSubjectAndSkidSubjectKeyID, alice)
		cliputils.RequireTxOK(t, txResult, err)

		// Now both certs with same subject+skid should coexist (2 entries in Certs).
		cert, err := GetX509Cert(rootWithSameSubjectAndSkidSubject, rootWithSameSubjectAndSkidSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, cert)
		require.Equal(t, rootWithSameSubjectAndSkidSubject, cert.Subject)
		require.Len(t, cert.Certs, 2)
	})

	// The query commands dispatch correctly and the DA (x509) and NOC namespaces
	// are kept separate — a DA certificate is visible through every x509/global
	// query but invisible through every NOC query. (The full DA+NOC+VVSC chain
	// is owned by TestPKIDemo / TestPKINocCerts here, so the separation is
	// asserted from the DA side using this test's own root.)
	t.Run("QueryDispatchAndNamespaceSeparation", func(t *testing.T) {
		// Visible through the DA / global single-cert queries.
		x509, err := GetX509Cert(subj, skid)
		require.NoError(t, err)
		require.NotNil(t, x509)
		require.True(t, certsContainSKID(x509.Certs))

		global, err := GetCert(subj, skid)
		require.NoError(t, err)
		require.NotNil(t, global)
		require.True(t, certsContainSKID(global.Certs))

		bySkid, err := GetX509CertBySKID(skid)
		require.NoError(t, err)
		require.NotNil(t, bySkid)
		require.True(t, certsContainSKID(bySkid.Certs))

		// Invisible through the NOC single-cert query (namespace separation).
		nocCert, err := GetNocCert("--subject", subj, "--subject-key-id", skid)
		require.NoError(t, err)
		require.Nil(t, nocCert)

		// Visible through the DA by-subject query; invisible through NOC by-subject.
		bySubject, err := GetX509CertsBySubject(subj)
		require.NoError(t, err)
		require.NotNil(t, bySubject)
		require.Contains(t, bySubject.SubjectKeyIds, skid)

		nocBySubject, err := GetNocSubjectCerts(subj)
		require.NoError(t, err)
		require.Nil(t, nocBySubject)

		// Visible through all-x509-certs and all-x509-root-certs; absent from
		// all-noc-x509-certs.
		allX509, err := GetAllX509Certs()
		require.NoError(t, err)
		require.True(t, approvedContain(allX509))

		allRoots, err := GetAllX509RootCerts()
		require.NoError(t, err)
		require.NotNil(t, allRoots)
		require.True(t, certsContainID(allRoots.Certs))

		allNoc, err := GetAllNocX509Certs()
		require.NoError(t, err)
		require.NotNil(t, allNoc)
		require.False(t, certsContainSKID(allNoc.Certs))

		// No child certs exist under this root (its children are added later by
		// the revocation tests).
		children, err := GetChildX509Certs(subj, skid)
		require.NoError(t, err)
		require.Nil(t, children)
	})
}
