package pki

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
)

const (
	nocRootCert1Path          = "../../constants/noc_root_cert_1"
	nocRootCert1Subject       = "MHoxCzAJBgNVBAYTAlVaMRMwEQYDVQQIEwpTb21lIFN0YXRlMREwDwYDVQQHEwhUYXNoa2VudDEYMBYGA1UEChMPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMQ4wDAYDVQQDEwVOT0MtMQ=="
	nocRootCert1SubjectKeyID  = "0E:10:B8:5D:96:7A:08:33:C7:C5:44:49:0E:28:0F:C1:6E:D5:D4:7C"
	nocRootCert1SerialNumber  = "313831573505791137291636389937677533381171619492"
	nocRootCert1SubjectAsText = "CN=NOC-1,OU=Testing Division,O=Example Company,L=Tashkent,ST=Some State,C=UZ"

	nocRootCert1CopyPath         = "../../constants/noc_root_cert_1_copy"
	nocRootCert1CopySerialNumber = "12722088350714347345576486793058060481880825999"

	nocRootCert2Path          = "../../constants/noc_root_cert_2"
	nocRootCert2Subject       = "MHoxCzAJBgNVBAYTAlVaMRMwEQYDVQQIEwpTb21lIFN0YXRlMREwDwYDVQQHEwhUYXNoa2VudDEYMBYGA1UEChMPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMQ4wDAYDVQQDEwVOT0MtMg=="
	nocRootCert2SubjectKeyID  = "46:C0:B0:74:0C:63:C8:9E:E0:5C:14:C2:71:62:F8:67:24:5C:8E:29"
	nocRootCert2SerialNumber  = "727423814323052015089749828769570958840545369270"
	nocRootCert2SubjectAsText = "CN=NOC-2,OU=Testing Division,O=Example Company,L=Tashkent,ST=Some State,C=UZ"

	nocRootCert3Path          = "../../constants/noc_root_cert_3"
	nocRootCert3Subject       = "MFUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpTb21lLVN0YXRlMSEwHwYDVQQKExhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQxDjAMBgNVBAMTBU5PQy0z"
	nocRootCert3SubjectKeyID  = "0F:D2:F8:12:06:F1:38:2D:D2:19:2F:29:52:42:AA:FB:E7:2F:7B:A3"
	nocRootCert3SerialNumber  = "620481712672111766723531823383547399894194653186"
	nocRootCert3SubjectAsText = "CN=NOC-3,O=Internet Widgits Pty Ltd,ST=Some-State,C=AU"

	nocCert1Path         = "../../constants/noc_cert_1"
	nocCert1Subject      = "MIGCMQswCQYDVQQGEwJVWjETMBEGA1UECBMKU29tZSBTdGF0ZTETMBEGA1UEBxMKU29tZSBTdGF0ZTEYMBYGA1UEChMPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMRQwEgYDVQQDEwtOT0MtY2hpbGQtMQ=="
	nocCert1SubjectKeyID = "06:9F:5A:E0:1F:23:3E:9F:C7:4F:B6:F9:A2:33:47:33:62:7A:07:C5"
	nocCert1SerialNumber = "577430346509479530103103319788179390906984119670"

	nocCert1CopyPath         = "../../constants/noc_cert_1_copy"
	nocCert1CopySerialNumber = "617357865778805507017637943649984133152592305888"

	nocCert2Path         = "../../constants/noc_cert_2"
	nocCert2Subject      = "MIGCMQswCQYDVQQGEwJVWjETMBEGA1UECBMKU29tZSBTdGF0ZTETMBEGA1UEBxMKU29tZSBTdGF0ZTEYMBYGA1UEChMPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMRQwEgYDVQQDEwtOT0MtY2hpbGQtMg=="
	nocCert2SubjectKeyID = "17:E2:72:19:E1:7F:19:D7:0D:02:1A:B0:40:7B:04:26:CC:D4:2B:F5"
	nocCert2SerialNumber = "634591262660314610068979921875981241084684028375"

	// Matter R1.6 §6.5.12 VVSC fixtures (CertificateType_VIDSignerPKI). All
	// subjects encode matter-vid=0001 in addition to the OperationalPKI-style
	// DN. The chain is vvscRoot1 → vvscIca1 → vvscLeaf1 (path length 3).
	vvscRootCert1Path         = "../../constants/vvsc_root_cert_1"
	vvscRootCert1Subject      = "MIGWMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTERMA8GA1UEBwwIVGFzaGtlbnQxGDAWBgNVBAoMD0V4YW1wbGUgQ29tcGFueTEZMBcGA1UECwwQVGVzdGluZyBEaXZpc2lvbjEUMBIGA1UEAwwLVlZTQy1Sb290LTExFDASBgorBgEEAYKifAIBDAQwMDAx"
	vvscRootCert1SubjectKeyID = "21:B9:21:60:2D:53:8B:86:DA:A4:16:5C:AA:40:90:25:EB:FE:7E:28"
	vvscRootCert1SerialNumber = "5068329979261235249"

	vvscIcaCert1Path         = "../../constants/vvsc_ica_cert_1"
	vvscIcaCert1Subject      = "MIGXMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRMwEQYDVQQDDApWVlNDLUlDQS0xMRQwEgYKKwYBBAGConwCAQwEMDAwMQ=="
	vvscIcaCert1SubjectKeyID = "98:4B:EE:D7:40:A2:FE:29:CB:AF:C0:0A:67:B7:AE:FF:12:A5:DA:DD"
	vvscIcaCert1SerialNumber = "5068329979109130545"

	vvscLeafCert1Path         = "../../constants/vvsc_leaf_cert_1"
	vvscLeafCert1Subject      = "MIGYMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRQwEgYDVQQDDAtWVlNDLUxlYWYtMTEUMBIGCisGAQQBgqJ8AgEMBDAwMDE="
	vvscLeafCert1SubjectKeyID = "42:24:A6:34:C8:C1:2F:88:9D:9C:7F:BE:8A:7A:6E:40:DB:C8:2B:F1"
	vvscLeafCert1SerialNumber = "5068329979159654449"

	nocVid  = 24582
	nocVid2 = 4701
)

func TestPKINocCerts(t *testing.T) {
	vendorAccount := fmt.Sprintf("vendor_account_%d", nocVid)
	cliputils.CreateVendorAccount(t, vendorAccount, nocVid)

	vendorAccount2 := fmt.Sprintf("vendor_account_%d", nocVid2)
	cliputils.CreateVendorAccount(t, vendorAccount2, nocVid2)

	t.Run("QueryAllEmpty", func(t *testing.T) {
		// Query by VID — not present.
		roots, err := GetNocRootCerts(nocVid)
		require.NoError(t, err)
		require.Nil(t, roots)

		// Query by VID + SKID for each known cert — not present.
		for _, skid := range []string{nocRootCert1SubjectKeyID, nocRootCert2SubjectKeyID, nocRootCert3SubjectKeyID} {
			cert, err := GetNocCert("--vid", fmt.Sprintf("%d", nocVid), "--subject-key-id", skid)
			require.NoError(t, err)
			require.Nil(t, cert)
		}

		// Query all NOC root certs — empty.
		allRoots, err := GetAllNocRootCerts()
		require.NoError(t, err)
		require.Empty(t, allRoots)

		// Query by subject + SKID — not present.
		cert, err := GetNocCert("--subject", nocRootCert1Subject, "--subject-key-id", nocRootCert1SubjectKeyID)
		require.NoError(t, err)
		require.Nil(t, cert)

		// Query by subject alone — not present.
		subjCerts, err := GetNocSubjectCerts(nocRootCert1Subject)
		require.NoError(t, err)
		require.Nil(t, subjCerts)

		// Query by SKID alone — not present.
		cert, err = GetNocCert("--subject-key-id", nocRootCert1SubjectKeyID)
		require.NoError(t, err)
		require.Nil(t, cert)
	})

	t.Run("AddNocRootCerts", func(t *testing.T) {
		// Try to add intermediate cert using add-noc-x509-root-cert — should fail
		txResult, err := AddNocRootCert(intermediateCertPath, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(414), txResult.Code)

		// Add first NOC root certificate
		txResult, err = AddNocRootCert(nocRootCert1Path, vendorAccount, AddNocCertOpts{SchemaVersion: "0"})
		cliputils.RequireTxOK(t, txResult, err)

		// Add second NOC root certificate
		txResult, err = AddNocRootCert(nocRootCert2Path, vendorAccount)
		cliputils.RequireTxOK(t, txResult, err)

		// Add third NOC root certificate (different VID vendor)
		txResult, err = AddNocRootCert(nocRootCert3Path, vendorAccount2)
		cliputils.RequireTxOK(t, txResult, err)
	})

	t.Run("QueryNocRootCertsByVid", func(t *testing.T) {
		// Query by VID — both cert1 and cert2 are present.
		roots, err := GetNocRootCerts(nocVid)
		require.NoError(t, err)
		require.NotNil(t, roots)
		require.True(t, containsCertSubjectSerial(roots.Certs, nocRootCert1Subject, nocRootCert1SerialNumber))
		require.True(t, containsCertSubjectSerial(roots.Certs, nocRootCert2Subject, nocRootCert2SerialNumber))
		c1 := findCertBySerial(roots.Certs, nocRootCert1SerialNumber)
		require.NotNil(t, c1)
		require.Equal(t, nocRootCert1SubjectKeyID, c1.SubjectKeyId)
		require.Equal(t, nocRootCert1SubjectAsText, c1.SubjectAsText)
		require.Equal(t, int32(nocVid), c1.Vid)
		require.Equal(t, uint32(0), c1.SchemaVersion)

		// Query by VID + SKID for cert1
		cert, err := GetNocCert("--vid", fmt.Sprintf("%d", nocVid), "--subject-key-id", nocRootCert1SubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, cert)
		require.True(t, containsCertSubjectSerial(cert.Certs, nocRootCert1Subject, nocRootCert1SerialNumber))
		require.Equal(t, float32(1), cert.Tq)

		// Query by VID + SKID for cert2
		cert, err = GetNocCert("--vid", fmt.Sprintf("%d", nocVid), "--subject-key-id", nocRootCert2SubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, cert)
		require.True(t, containsCertSubjectSerial(cert.Certs, nocRootCert2Subject, nocRootCert2SerialNumber))

		// Query all NOC root certs — all three certs from both VIDs.
		allRoots, err := GetAllNocRootCerts()
		require.NoError(t, err)
		require.True(t, containsNocRootCertSerial(allRoots, nocRootCert1SerialNumber))
		require.True(t, containsNocRootCertSerial(allRoots, nocRootCert2SerialNumber))
		require.True(t, containsNocRootCertSerial(allRoots, nocRootCert3SerialNumber))

		// Query by subject + SKID using noc-x509-cert.
		cert, err = GetNocCert("--subject", nocRootCert1Subject, "--subject-key-id", nocRootCert1SubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, cert)
		require.True(t, containsCertSubjectSerial(cert.Certs, nocRootCert1Subject, nocRootCert1SerialNumber))

		// Query by subject + SKID using generic cert command.
		gen, err := GetCert(nocRootCert1Subject, nocRootCert1SubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, gen)
		require.True(t, containsCertSubjectSerial(gen.Certs, nocRootCert1Subject, nocRootCert1SerialNumber))

		// Query by subject alone.
		subjCerts, err := GetNocSubjectCerts(nocRootCert1Subject)
		require.NoError(t, err)
		require.NotNil(t, subjCerts)
		require.Equal(t, nocRootCert1Subject, subjCerts.Subject)
		require.Contains(t, subjCerts.SubjectKeyIds, nocRootCert1SubjectKeyID)

		// Query by SKID alone.
		cert, err = GetNocCert("--subject-key-id", nocRootCert1SubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, cert)
		require.True(t, containsCertSubjectSerial(cert.Certs, nocRootCert1Subject, nocRootCert1SerialNumber))
	})

	t.Run("AddNocIcaCerts", func(t *testing.T) {
		// Add first ICA cert
		txResult, err := AddNocX509IcaCert(nocCert1Path, vendorAccount)
		cliputils.RequireTxOK(t, txResult, err)

		// ICA certs by VID — cert1 present.
		icas, err := GetNocX509IcaCerts(nocVid)
		require.NoError(t, err)
		require.NotNil(t, icas)
		require.True(t, containsCertSubjectSerial(icas.Certs, nocCert1Subject, nocCert1SerialNumber))

		// Child certs of root1 — cert1 present.
		children, err := GetChildX509Certs(nocRootCert1Subject, nocRootCert1SubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, children)
		foundChild := false
		for _, id := range children.CertIds {
			if id.Subject == nocCert1Subject && id.SubjectKeyId == nocCert1SubjectKeyID {
				foundChild = true

				break
			}
		}
		require.True(t, foundChild)

		// Try to add ICA with different VID — should fail
		txResult, err = AddNocX509IcaCert(nocCert2Path, vendorAccount2)
		require.NoError(t, err)
		require.Equal(t, uint32(439), txResult.Code)

		// Add second ICA cert
		txResult, err = AddNocX509IcaCert(nocCert2Path, vendorAccount, AddNocCertOpts{SchemaVersion: "0"})
		cliputils.RequireTxOK(t, txResult, err)

		// Add cert copy
		txResult, err = AddNocX509IcaCert(nocCert1CopyPath, vendorAccount)
		cliputils.RequireTxOK(t, txResult, err)

		// All ICA certs include cert1 (both serials), cert2.
		allIcas, err := GetAllNocX509IcaCerts()
		require.NoError(t, err)
		require.True(t, containsNocIcaCertSerial(allIcas, nocCert1SerialNumber))
		require.True(t, containsNocIcaCertSerial(allIcas, nocCert1CopySerialNumber))
		require.True(t, containsNocIcaCertSerial(allIcas, nocCert2SerialNumber))

		// NOC certs must NOT appear in the DA approved cert list.
		da, err := GetAllX509Certs()
		require.NoError(t, err)
		require.False(t, containsApprovedCertSerial(da, nocRootCert1SerialNumber))
		require.False(t, containsApprovedCertSerial(da, nocCert1SerialNumber))

		// All NOC certs — root and ICA both present.
		all, err := GetAllNocX509Certs()
		require.NoError(t, err)
		require.NotNil(t, all)
		require.True(t, containsCertSerial(all.Certs, nocRootCert1SerialNumber))
		require.True(t, containsCertSerial(all.Certs, nocCert1SerialNumber))
		require.True(t, containsCertSerial(all.Certs, nocCert1CopySerialNumber))
		require.True(t, containsCertSerial(all.Certs, nocCert2SerialNumber))
	})

	t.Run("AddAndRevokeNocRootCert", func(t *testing.T) {
		// Add root cert copy
		txResult, err := AddNocRootCert(nocRootCert1CopyPath, vendorAccount)
		cliputils.RequireTxOK(t, txResult, err)

		// Add a Matter §6.4.5.4 VVSC chain (self-issued VVSC root, VVSC intermediate,
		// VVSC leaf) so the leaf-level operations have a §6.5.12-compliant chain to
		// exercise. NocLeafCert1 is a NOC end-entity (cA=FALSE / NOC profile) and is
		// no longer accepted by the stricter add-noc-x509-ica-cert handler.
		txResult, err = AddNocRootCert(vvscRootCert1Path, vendorAccount, AddNocCertOpts{IsVidVerificationSigner: true})
		cliputils.RequireTxOK(t, txResult, err)

		txResult, err = AddNocX509IcaCert(vvscIcaCert1Path, vendorAccount, AddNocCertOpts{IsVidVerificationSigner: true})
		cliputils.RequireTxOK(t, txResult, err)

		txResult, err = AddNocX509IcaCert(vvscLeafCert1Path, vendorAccount, AddNocCertOpts{IsVidVerificationSigner: true})
		cliputils.RequireTxOK(t, txResult, err)

		// Verify root state before revocation.
		allRoots, err := GetAllNocRootCerts()
		require.NoError(t, err)
		require.True(t, containsNocRootCertSerial(allRoots, nocRootCert1SerialNumber))
		require.True(t, containsNocRootCertSerial(allRoots, nocRootCert1CopySerialNumber))
		require.True(t, containsNocRootCertSerial(allRoots, nocRootCert2SerialNumber))

		// Verify ICA state before revocation.
		allIcas, err := GetAllNocX509IcaCerts()
		require.NoError(t, err)
		require.True(t, containsNocIcaCertSerial(allIcas, nocCert1SerialNumber))
		require.True(t, containsNocIcaCertSerial(allIcas, nocCert1CopySerialNumber))
		require.True(t, containsNocIcaCertSerial(allIcas, nocCert2SerialNumber))
		require.True(t, containsNocIcaCertSerial(allIcas, vvscLeafCert1SerialNumber))

		// Try to revoke with different VID — should fail
		txResult, err = RevokeNocRootCert(nocRootCert1Subject, nocRootCert1SubjectKeyID, vendorAccount2)
		require.NoError(t, err)
		require.Equal(t, uint32(439), txResult.Code)

		// Revoke root cert without child flag — ICA must survive
		txResult, err = RevokeNocRootCert(nocRootCert1Subject, nocRootCert1SubjectKeyID, vendorAccount)
		cliputils.RequireTxOK(t, txResult, err)

		// All revoked NOC root certs — both serials, ICA/leaf absent.
		revokedRoots, err := GetAllRevokedNocRootCerts()
		require.NoError(t, err)
		require.True(t, containsRevokedNocRootCertSerial(revokedRoots, nocRootCert1SerialNumber))
		require.True(t, containsRevokedNocRootCertSerial(revokedRoots, nocRootCert1CopySerialNumber))

		// Revoked NOC root cert by subject + SKID.
		revokedRoot, err := GetRevokedNocRootCert(nocRootCert1Subject, nocRootCert1SubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, revokedRoot)
		require.True(t, containsCertSerial(revokedRoot.Certs, nocRootCert1SerialNumber))
		require.True(t, containsCertSerial(revokedRoot.Certs, nocRootCert1CopySerialNumber))

		// DA revoked certs must NOT contain revoked NOC root certs.
		daRevoked, err := GetAllRevokedX509Certs()
		require.NoError(t, err)
		require.False(t, containsRevokedCertSerial(daRevoked, nocRootCert1SerialNumber))
		require.False(t, containsRevokedCertSerial(daRevoked, nocRootCert1CopySerialNumber))

		// Active NOC root certs by VID — cert2 present, cert1 absent.
		roots, err := GetNocRootCerts(nocVid)
		require.NoError(t, err)
		require.NotNil(t, roots)
		require.True(t, containsCertSerial(roots.Certs, nocRootCert2SerialNumber))
		require.False(t, containsCertSerial(roots.Certs, nocRootCert1SerialNumber))
		require.False(t, containsCertSerial(roots.Certs, nocRootCert1CopySerialNumber))

		// Query by VID + SKID for cert1 — gone.
		cert, err := GetNocCert("--vid", fmt.Sprintf("%d", nocVid), "--subject-key-id", nocRootCert1SubjectKeyID)
		require.NoError(t, err)
		require.Nil(t, cert)

		// Query by VID + SKID for cert2 — present.
		cert, err = GetNocCert("--vid", fmt.Sprintf("%d", nocVid), "--subject-key-id", nocRootCert2SubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, cert)
		require.True(t, containsCertSubjectSerial(cert.Certs, nocRootCert2Subject, nocRootCert2SerialNumber))

		// Query by subject for cert1 — gone (either nil or doesn't list the SKID).
		subjCerts, err := GetNocSubjectCerts(nocRootCert1Subject)
		require.NoError(t, err)
		if subjCerts != nil {
			require.NotContains(t, subjCerts.SubjectKeyIds, nocRootCert1SubjectKeyID)
		}

		// Query by SKID alone for cert1 — gone.
		cert, err = GetNocCert("--subject-key-id", nocRootCert1SubjectKeyID)
		require.NoError(t, err)
		require.Nil(t, cert)

		// ICA certs by VID — ICA and leaf still active.
		icas, err := GetNocX509IcaCerts(nocVid)
		require.NoError(t, err)
		require.NotNil(t, icas)
		require.True(t, containsCertSubjectSerial(icas.Certs, nocCert1Subject, nocCert1SerialNumber))
		require.True(t, containsCertSubjectSerial(icas.Certs, nocCert1Subject, nocCert1CopySerialNumber))
		require.True(t, containsCertSubjectSerial(icas.Certs, vvscLeafCert1Subject, vvscLeafCert1SerialNumber))

		// All NOC certs — ICA/leaf present, revoked root1 absent.
		all, err := GetAllNocX509Certs()
		require.NoError(t, err)
		require.NotNil(t, all)
		require.True(t, containsCertSerial(all.Certs, nocCert1SerialNumber))
		require.True(t, containsCertSerial(all.Certs, nocCert1CopySerialNumber))
		require.True(t, containsCertSerial(all.Certs, vvscLeafCert1SerialNumber))
		require.False(t, containsCertSerial(all.Certs, nocRootCert1SerialNumber))
		require.False(t, containsCertSerial(all.Certs, nocRootCert1CopySerialNumber))
	})

	t.Run("RevokeNocIcaCert", func(t *testing.T) {
		// Try to revoke with different VID — should fail
		txResult, err := RevokeNocX509IcaCert(nocCert1Subject, nocCert1SubjectKeyID, vendorAccount2)
		require.NoError(t, err)
		require.Equal(t, uint32(439), txResult.Code)

		// Revoke ICA cert without child flag — leaf must survive
		txResult, err = RevokeNocX509IcaCert(nocCert1Subject, nocCert1SubjectKeyID, vendorAccount)
		cliputils.RequireTxOK(t, txResult, err)

		// Revoked ICA list — cert1 present, leaf absent.
		revokedIcas, err := GetAllRevokedNocX509IcaCerts()
		require.NoError(t, err)
		require.True(t, containsRevokedNocIcaCertSerial(revokedIcas, nocCert1SerialNumber))
		require.False(t, containsRevokedNocIcaCertSerial(revokedIcas, vvscLeafCert1SerialNumber))
		require.False(t, containsRevokedNocIcaCertSubject(revokedIcas, vvscLeafCert1Subject))

		// Revoked root list must not contain ICA or leaf.
		revokedRoots, err := GetAllRevokedNocRootCerts()
		require.NoError(t, err)
		for _, r := range revokedRoots {
			require.NotEqual(t, nocCert1SubjectKeyID, r.SubjectKeyId)
			require.NotEqual(t, vvscLeafCert1SubjectKeyID, r.SubjectKeyId)
		}

		// Query by subject for cert1 — gone (either nil or doesn't list the SKID).
		subjCerts, err := GetNocSubjectCerts(nocCert1Subject)
		require.NoError(t, err)
		if subjCerts != nil {
			require.NotContains(t, subjCerts.SubjectKeyIds, nocCert1SubjectKeyID)
		}

		// Query by SKID alone for cert1 — gone.
		cert, err := GetNocCert("--subject-key-id", nocCert1SubjectKeyID)
		require.NoError(t, err)
		require.Nil(t, cert)

		// Active ICA certs by VID — only leaf remains.
		icas, err := GetNocX509IcaCerts(nocVid)
		require.NoError(t, err)
		require.NotNil(t, icas)
		require.True(t, containsCertSubjectSerial(icas.Certs, vvscLeafCert1Subject, vvscLeafCert1SerialNumber))
		for _, c := range icas.Certs {
			require.NotEqual(t, nocCert1Subject, c.Subject)
		}

		// Query by VID + SKID for cert1 — gone.
		cert, err = GetNocCert("--vid", fmt.Sprintf("%d", nocVid), "--subject-key-id", nocCert1SubjectKeyID)
		require.NoError(t, err)
		require.Nil(t, cert)

		// Query by VID + SKID for leaf — present.
		cert, err = GetNocCert("--vid", fmt.Sprintf("%d", nocVid), "--subject-key-id", vvscLeafCert1SubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, cert)
		require.True(t, containsCertSubjectSerial(cert.Certs, vvscLeafCert1Subject, vvscLeafCert1SerialNumber))

		// All NOC certs — leaf present, cert1 and root1 absent.
		all, err := GetAllNocX509Certs()
		require.NoError(t, err)
		require.NotNil(t, all)
		require.True(t, containsCertSerial(all.Certs, vvscLeafCert1SerialNumber))
		require.False(t, containsCertSerial(all.Certs, nocCert1SerialNumber))
		require.False(t, containsCertSerial(all.Certs, nocRootCert1SerialNumber))
		require.False(t, containsCertSerial(all.Certs, nocRootCert1CopySerialNumber))
	})
}
