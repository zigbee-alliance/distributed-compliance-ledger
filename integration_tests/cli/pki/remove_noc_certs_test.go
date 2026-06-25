package pki

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
)

const (
	removeNocRootCertVid = 65521
	removeNocOtherVid    = 65522

	removeNocRootCert1SerialNumber = "313831573505791137291636389937677533381171619492"
	removeNocRootCert1CopySerial   = "12722088350714347345576486793058060481880825999"
	removeNocIntermCert1Serial     = "577430346509479530103103319788179390906984119670"
	removeNocIntermCert2Serial     = "617357865778805507017637943649984133152592305888"
	// removeNocLeafCertSerial holds the VVSC leaf 1 serial number — NocLeafCert1
	// (NOC end-entity) is no longer accepted by the strict §6.5.12 ICA handler
	// so we use the Matter VVSC leaf instead. Variable names keep the "leaf"
	// terminology so downstream assertions read naturally.
	removeNocLeafCertSerial = "5068329979159654449"

	removeNocRootCertSubject      = "MHoxCzAJBgNVBAYTAlVaMRMwEQYDVQQIEwpTb21lIFN0YXRlMREwDwYDVQQHEwhUYXNoa2VudDEYMBYGA1UEChMPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMQ4wDAYDVQQDEwVOT0MtMQ=="
	removeNocRootCertSubjectKeyID = "0E:10:B8:5D:96:7A:08:33:C7:C5:44:49:0E:28:0F:C1:6E:D5:D4:7C"

	removeNocIntermCertSubject      = "MIGCMQswCQYDVQQGEwJVWjETMBEGA1UECBMKU29tZSBTdGF0ZTETMBEGA1UEBxMKU29tZSBTdGF0ZTEYMBYGA1UEChMPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMRQwEgYDVQQDEwtOT0MtY2hpbGQtMQ=="
	removeNocIntermCertSubjectKeyID = "06:9F:5A:E0:1F:23:3E:9F:C7:4F:B6:F9:A2:33:47:33:62:7A:07:C5"

	// removeNocLeafCert* — same VVSC leaf 1 fixture as the rest of the package.
	removeNocLeafCertSubject      = "MIGYMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRQwEgYDVQQDDAtWVlNDLUxlYWYtMTEUMBIGCisGAQQBgqJ8AgEMBDAwMDE="
	removeNocLeafCertSubjectKeyID = "42:24:A6:34:C8:C1:2F:88:9D:9C:7F:BE:8A:7A:6E:40:DB:C8:2B:F1"
)

func TestPKIRemoveNocCertificates(t *testing.T) {
	jack := testconstants.JackAccount

	vendorAccount65521 := fmt.Sprintf("vendor_account_%d", removeNocRootCertVid)
	cliputils.CreateVendorAccount(t, vendorAccount65521, removeNocRootCertVid)

	vendorAccount65522 := fmt.Sprintf("vendor_account_%d", removeNocOtherVid)
	cliputils.CreateVendorAccount(t, vendorAccount65522, removeNocOtherVid)

	// Prior NOC tests add noc_root_cert_1, noc_cert_1, vvsc_root_cert_1, vvsc_ica_cert_1,
	// and vvsc_leaf_cert_1 under VID 24582 (nocVid) and leave them in the revoked
	// state. The unique-cert store retains their serial entries, preventing
	// re-addition by VID 65521. Remove them via the owning VID 24582 account first.
	vendorAccount24582 := fmt.Sprintf("vendor_account_%d", nocVid)
	cliputils.CreateVendorAccount(t, vendorAccount24582, nocVid)
	RemoveNocRootCert(nocRootCert1Subject, nocRootCert1SubjectKeyID, vendorAccount24582)   //nolint:errcheck
	RemoveNocCert(nocCert1Subject, nocCert1SubjectKeyID, vendorAccount24582)               //nolint:errcheck
	RemoveNocRootCert(vvscRootCert1Subject, vvscRootCert1SubjectKeyID, vendorAccount24582) //nolint:errcheck
	RemoveNocCert(vvscIcaCert1Subject, vvscIcaCert1SubjectKeyID, vendorAccount24582)       //nolint:errcheck
	RemoveNocCert(vvscLeafCert1Subject, vvscLeafCert1SubjectKeyID, vendorAccount24582)     //nolint:errcheck

	t.Run("SetupCerts", func(t *testing.T) {
		// Add root cert
		txResult, err := AddNocRootCert(nocRootCert1Path, vendorAccount65521)
		cliputils.RequireTxOK(t, txResult, err)

		// Add NOC ICA certs
		txResult, err = AddNocX509IcaCert(nocCert1Path, vendorAccount65521)
		cliputils.RequireTxOK(t, txResult, err)

		txResult, err = AddNocX509IcaCert(nocCert1CopyPath, vendorAccount65521)
		cliputils.RequireTxOK(t, txResult, err)

		// Pre-seed the VVSC chain (Matter §6.4.5.4) so the leaf below has a
		// §6.4.10 step 12.a.iii path-length-3 chain to validate against.
		txResult, err = AddNocRootCert(vvscRootCert1Path, vendorAccount65521, AddNocCertOpts{IsVidVerificationSigner: true})
		cliputils.RequireTxOK(t, txResult, err)

		txResult, err = AddNocX509IcaCert(vvscIcaCert1Path, vendorAccount65521, AddNocCertOpts{IsVidVerificationSigner: true})
		cliputils.RequireTxOK(t, txResult, err)

		// Add VVSC leaf certificate (replaces the legacy NocLeafCert1).
		txResult, err = AddNocX509IcaCert(vvscLeafCert1Path, vendorAccount65521, AddNocCertOpts{IsVidVerificationSigner: true})
		cliputils.RequireTxOK(t, txResult, err)

		all, err := GetAllNocX509Certs()
		require.NoError(t, err)
		require.NotNil(t, all)
		require.True(t, containsCertSerial(all.Certs, removeNocRootCert1SerialNumber))
		require.True(t, containsCertSerial(all.Certs, removeNocIntermCert1Serial))
		require.True(t, containsCertSerial(all.Certs, removeNocIntermCert2Serial))
		require.True(t, containsCertSerial(all.Certs, removeNocLeafCertSerial))
	})

	t.Run("RevokeAndRemoveIcaCert", func(t *testing.T) {
		// Revoke first ICA cert
		txResult, err := RevokeNocX509IcaCert(removeNocIntermCertSubject, removeNocIntermCertSubjectKeyID, vendorAccount65521, RevokeNocCertOpts{SerialNumber: removeNocIntermCert1Serial})
		cliputils.RequireTxOK(t, txResult, err)

		// Try to remove with invalid serial
		txResult, err = RemoveNocCert(removeNocIntermCertSubject, removeNocIntermCertSubjectKeyID, vendorAccount65521, RevokeNocCertOpts{SerialNumber: "invalid"})
		require.NoError(t, err)
		require.Equal(t, uint32(404), txResult.Code)

		// Try to remove when sender is not Vendor
		txResult, err = RemoveNocCert(removeNocIntermCertSubject, removeNocIntermCertSubjectKeyID, jack, RevokeNocCertOpts{SerialNumber: removeNocIntermCert1Serial})
		require.NoError(t, err)
		require.Equal(t, uint32(4), txResult.Code)

		// Try to remove with different VID vendor (returns 439 ErrCertVidNotEqualAccountVid, not 4)
		txResult, err = RemoveNocCert(removeNocIntermCertSubject, removeNocIntermCertSubjectKeyID, vendorAccount65522, RevokeNocCertOpts{SerialNumber: removeNocIntermCert1Serial})
		require.NoError(t, err)
		require.Equal(t, uint32(439), txResult.Code)

		// Remove revoked ICA cert by serial
		txResult, err = RemoveNocCert(removeNocIntermCertSubject, removeNocIntermCertSubjectKeyID, vendorAccount65521, RevokeNocCertOpts{SerialNumber: removeNocIntermCert1Serial})
		cliputils.RequireTxOK(t, txResult, err)

		// Only second ICA cert should remain.
		cert, err := GetNocCert("--subject", removeNocIntermCertSubject, "--subject-key-id", removeNocIntermCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, cert)
		require.True(t, containsCertSerial(cert.Certs, removeNocIntermCert2Serial))
		require.False(t, containsCertSerial(cert.Certs, removeNocIntermCert1Serial))

		// ICA certs by VID still carry the remaining ICA + the leaf; the removed serial is gone.
		icas, err := GetNocX509IcaCerts(removeNocRootCertVid)
		require.NoError(t, err)
		require.NotNil(t, icas)
		require.True(t, containsCertSubjectSerial(icas.Certs, removeNocIntermCertSubject, removeNocIntermCert2Serial))
		require.True(t, containsCertSerial(icas.Certs, removeNocLeafCertSerial))
		require.False(t, containsCertSerial(icas.Certs, removeNocIntermCert1Serial))

		// Remove remaining ICA cert by subject+subjectKeyID
		txResult, err = RemoveNocCert(removeNocIntermCertSubject, removeNocIntermCertSubjectKeyID, vendorAccount65521)
		cliputils.RequireTxOK(t, txResult, err)

		cert, err = GetNocCert("--subject", removeNocIntermCertSubject, "--subject-key-id", removeNocIntermCertSubjectKeyID)
		require.NoError(t, err)
		require.Nil(t, cert)

		// ICA certs by VID now carry only the leaf; the ICA subject is gone.
		icas, err = GetNocX509IcaCerts(removeNocRootCertVid)
		require.NoError(t, err)
		require.NotNil(t, icas)
		require.True(t, containsCertSerial(icas.Certs, removeNocLeafCertSerial))
		for _, c := range icas.Certs {
			require.NotEqual(t, removeNocIntermCertSubject, c.Subject)
		}

		// This test's ICA certs (subject = removeNocIntermCertSubject) should not be in the revoked list.
		// (Prior NOC tests may have left other subjects in the revoked list, so we cannot assert empty.)
		revoked, err := GetAllRevokedNocX509IcaCerts()
		require.NoError(t, err)
		require.False(t, containsRevokedNocIcaCertSubject(revoked, removeNocIntermCertSubject))
	})

	t.Run("RemoveLeafCert", func(t *testing.T) {
		txResult, err := RemoveNocCert(removeNocLeafCertSubject, removeNocLeafCertSubjectKeyID, vendorAccount65521)
		cliputils.RequireTxOK(t, txResult, err)

		cert, err := GetNocCert("--subject", removeNocLeafCertSubject, "--subject-key-id", removeNocLeafCertSubjectKeyID)
		require.NoError(t, err)
		require.Nil(t, cert)

		// Once this test's ICA and leaf are removed, none of its serials/subject
		// remain under the VID. (Other NOC tests share VID 65521 on the shared
		// ledger, so the by-VID list itself may still be non-empty.)
		icas, err := GetNocX509IcaCerts(removeNocRootCertVid)
		require.NoError(t, err)
		if icas != nil {
			require.False(t, containsCertSerial(icas.Certs, removeNocLeafCertSerial))
			require.False(t, containsCertSerial(icas.Certs, removeNocIntermCert2Serial))
			for _, c := range icas.Certs {
				require.NotEqual(t, removeNocIntermCertSubject, c.Subject)
			}
		}

		// All NOC certs still carry the root, but neither leaf nor ICA serials.
		all, err := GetAllNocX509Certs()
		require.NoError(t, err)
		require.NotNil(t, all)
		require.True(t, containsCertSerial(all.Certs, removeNocRootCert1SerialNumber))
		require.False(t, containsCertSerial(all.Certs, removeNocLeafCertSerial))
		require.False(t, containsCertSerial(all.Certs, removeNocIntermCert2Serial))
	})

	t.Run("RemoveNocRootCert", func(t *testing.T) {
		// Add root cert copy
		txResult, err := AddNocRootCert(nocRootCert1CopyPath, vendorAccount65521)
		cliputils.RequireTxOK(t, txResult, err)

		// Re-add ICA cert
		txResult, err = AddNocX509IcaCert(nocCert1Path, vendorAccount65521)
		cliputils.RequireTxOK(t, txResult, err)

		// Try to remove root with invalid serial
		txResult, err = RemoveNocRootCert(removeNocRootCertSubject, removeNocRootCertSubjectKeyID, vendorAccount65521, RevokeNocCertOpts{SerialNumber: "invalid"})
		require.NoError(t, err)
		require.Equal(t, uint32(404), txResult.Code)

		// Try to remove when sender is not Vendor
		txResult, err = RemoveNocRootCert(removeNocRootCertSubject, removeNocRootCertSubjectKeyID, jack, RevokeNocCertOpts{SerialNumber: removeNocRootCert1SerialNumber})
		require.NoError(t, err)
		require.Equal(t, uint32(4), txResult.Code)

		// Try to remove with different VID vendor (returns 439 ErrCertVidNotEqualAccountVid, not 4)
		txResult, err = RemoveNocRootCert(removeNocRootCertSubject, removeNocRootCertSubjectKeyID, vendorAccount65522, RevokeNocCertOpts{SerialNumber: removeNocRootCert1SerialNumber})
		require.NoError(t, err)
		require.Equal(t, uint32(439), txResult.Code)

		// Revoke root cert
		txResult, err = RevokeNocRootCert(removeNocRootCertSubject, removeNocRootCertSubjectKeyID, vendorAccount65521, RevokeNocCertOpts{SerialNumber: removeNocRootCert1SerialNumber})
		cliputils.RequireTxOK(t, txResult, err)

		// Revoked NOC root list now carries the revoked root serial.
		revokedRoots, err := GetAllRevokedNocRootCerts()
		require.NoError(t, err)
		require.True(t, containsRevokedNocRootCertSerial(revokedRoots, removeNocRootCert1SerialNumber))

		// Remove revoked root cert by serial
		txResult, err = RemoveNocRootCert(removeNocRootCertSubject, removeNocRootCertSubjectKeyID, vendorAccount65521, RevokeNocCertOpts{SerialNumber: removeNocRootCert1SerialNumber})
		cliputils.RequireTxOK(t, txResult, err)

		// Only copy root should remain.
		roots, err := GetNocRootCerts(removeNocRootCertVid)
		require.NoError(t, err)
		require.NotNil(t, roots)
		require.True(t, containsCertSerial(roots.Certs, removeNocRootCert1CopySerial))
		require.False(t, containsCertSerial(roots.Certs, removeNocRootCert1SerialNumber))

		// Re-add root cert and then remove both
		txResult, err = AddNocRootCert(nocRootCert1Path, vendorAccount65521)
		cliputils.RequireTxOK(t, txResult, err)

		// Remove all root certs by subject+subjectKeyID (no serial)
		txResult, err = RemoveNocRootCert(removeNocRootCertSubject, removeNocRootCertSubjectKeyID, vendorAccount65521)
		cliputils.RequireTxOK(t, txResult, err)

		cert, err := GetNocCert("--subject", removeNocRootCertSubject, "--subject-key-id", removeNocRootCertSubjectKeyID)
		require.NoError(t, err)
		require.Nil(t, cert)

		// NOC root certs by VID are empty once both roots are removed.
		roots, err = GetNocRootCerts(removeNocRootCertVid)
		require.NoError(t, err)
		if roots != nil {
			require.False(t, containsCertSerial(roots.Certs, removeNocRootCert1SerialNumber))
			require.False(t, containsCertSerial(roots.Certs, removeNocRootCert1CopySerial))
		}

		// noc-x509-cert by VID+SKID no longer carries this test's root serials.
		// (The root subject/SKID is shared with other NOC tests on the shared
		// ledger, so the VID+SKID entry itself may still exist for their certs.)
		cert, err = GetNocCert("--vid", fmt.Sprintf("%d", removeNocRootCertVid), "--subject-key-id", removeNocRootCertSubjectKeyID)
		require.NoError(t, err)
		if cert != nil {
			require.False(t, containsCertSerial(cert.Certs, removeNocRootCert1SerialNumber))
			require.False(t, containsCertSerial(cert.Certs, removeNocRootCert1CopySerial))
		}

		// ICA cert should still be present.
		all, err := GetAllNocX509Certs()
		require.NoError(t, err)
		require.NotNil(t, all)
		require.True(t, containsCertSerial(all.Certs, removeNocIntermCert1Serial))

		// noc-x509-cert by VID+SKID for the re-added ICA is present.
		cert, err = GetNocCert("--vid", fmt.Sprintf("%d", removeNocRootCertVid), "--subject-key-id", removeNocIntermCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, cert)
		require.True(t, containsCertSerial(cert.Certs, removeNocIntermCert1Serial))
	})
}
