package pki

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
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

	nocLeafCert1Path         = "../../constants/noc_leaf_cert_1"
	nocLeafCert1Subject      = "MIGBMQswCQYDVQQGEwJVWjETMBEGA1UECBMKU29tZSBTdGF0ZTETMBEGA1UEBxMKU29tZSBTdGF0ZTEYMBYGA1UEChMPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMRMwEQYDVQQDEwpOT0MtbGVhZi0x"
	nocLeafCert1SubjectKeyID = "F0:3A:A5:96:8F:60:63:F0:15:21:24:0C:DA:0A:E6:2B:CC:A0:58:F9"
	nocLeafCert1SerialNumber = "201294310322324358101935163941973786245732555938"

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
		// Query by VID — Not Found
		out, err := QueryNocRootCerts(nocVid)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert1Subject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert1SubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1SerialNumber))

		// Query by VID + SKID for cert1 — Not Found
		out, err = QueryNocCert("--vid", fmt.Sprintf("%d", nocVid), "--subject-key-id", nocRootCert1SubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert1Subject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert1SubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1SerialNumber))

		// Query by VID + SKID for cert2 — Not Found
		out, err = QueryNocCert("--vid", fmt.Sprintf("%d", nocVid), "--subject-key-id", nocRootCert2SubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert2Subject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert2SubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert2SerialNumber))

		// Query by VID + SKID for cert3 — Not Found
		out, err = QueryNocCert("--vid", fmt.Sprintf("%d", nocVid), "--subject-key-id", nocRootCert3SubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert3Subject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert3SubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert3SerialNumber))

		// Query all — empty
		out, err = QueryAllNocRootCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), "[]")
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert1Subject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert1SubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1SerialNumber))

		// Query by subject + SKID — Not Found
		out, err = QueryNocCert("--subject", nocRootCert1Subject, "--subject-key-id", nocRootCert1SubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert1Subject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert1SubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1SerialNumber))

		// Query by subject alone — Not Found
		out, err = QueryNocSubjectCerts(nocRootCert1Subject)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert1Subject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert1SubjectKeyID))

		// Query by SKID alone — Not Found
		out, err = QueryNocCert("--subject-key-id", nocRootCert1SubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert1Subject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert1SubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1SerialNumber))
	})

	t.Run("AddNocRootCerts", func(t *testing.T) {
		// Try to add intermediate cert using add-noc-x509-root-cert — should fail
		txResult, err := AddNocRootCert(intermediateCertPath, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(414), txResult.Code)

		// Add first NOC root certificate
		txResult, err = AddNocRootCert(nocRootCert1Path, vendorAccount, "--schemaVersion", "0")
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Add second NOC root certificate
		txResult, err = AddNocRootCert(nocRootCert2Path, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Add third NOC root certificate (different VID vendor)
		txResult, err = AddNocRootCert(nocRootCert3Path, vendorAccount2)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("QueryNocRootCertsByVid", func(t *testing.T) {
		// Query by VID — both cert1 and cert2 present with all fields
		out, err := QueryNocRootCerts(nocVid)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert1SubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectAsText":"%s"`, nocRootCert1SubjectAsText))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert2Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert2SubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert2SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectAsText":"%s"`, nocRootCert2SubjectAsText))
		require.Contains(t, string(out), `"schemaVersion":0`)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, nocVid))

		// Query by VID + SKID for cert1
		out, err = QueryNocCert("--vid", fmt.Sprintf("%d", nocVid), "--subject-key-id", nocRootCert1SubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert1SubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectAsText":"%s"`, nocRootCert1SubjectAsText))
		require.Contains(t, string(out), `"schemaVersion":0`)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, nocVid))
		require.Contains(t, string(out), `"tq":1`)

		// Query by VID + SKID for cert2
		out, err = QueryNocCert("--vid", fmt.Sprintf("%d", nocVid), "--subject-key-id", nocRootCert2SubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert2Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert2SubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert2SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectAsText":"%s"`, nocRootCert2SubjectAsText))
		require.Contains(t, string(out), `"schemaVersion":0`)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, nocVid))
		require.Contains(t, string(out), `"tq":1`)

		// Query all NOC root certs — all three certs from both VIDs
		out, err = QueryAllNocRootCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert1SubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectAsText":"%s"`, nocRootCert1SubjectAsText))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert2Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert2SubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert2SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectAsText":"%s"`, nocRootCert2SubjectAsText))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert3Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert3SubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert3SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectAsText":"%s"`, nocRootCert3SubjectAsText))
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, nocVid))
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, nocVid2))

		// Query by subject + SKID using noc-x509-cert
		out, err = QueryNocCert("--subject", nocRootCert1Subject, "--subject-key-id", nocRootCert1SubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert1SubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectAsText":"%s"`, nocRootCert1SubjectAsText))
		require.Contains(t, string(out), `"approvals":[]`)

		// Query by subject + SKID using generic cert command
		out, err = QueryCert(nocRootCert1Subject, nocRootCert1SubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert1SubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectAsText":"%s"`, nocRootCert1SubjectAsText))
		require.Contains(t, string(out), `"approvals":[]`)

		// Query by subject alone
		out, err = QueryNocSubjectCerts(nocRootCert1Subject)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"%s"`, nocRootCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"%s"`, nocRootCert1SubjectKeyID))

		// Query by SKID alone
		out, err = QueryNocCert("--subject-key-id", nocRootCert1SubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert1SubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectAsText":"%s"`, nocRootCert1SubjectAsText))
	})

	t.Run("AddNocIcaCerts", func(t *testing.T) {
		// Add first ICA cert
		txResult, err := AddNocX509IcaCert(nocCert1Path, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// ICA certs by VID — cert1 present with all fields
		out, err := QueryNocX509IcaCerts(nocVid)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocCert1SubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocCert1SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, nocVid))

		// Child certs of root1 — cert1 present
		out, err = QueryChildX509Certs(nocRootCert1Subject, nocRootCert1SubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocCert1SubjectKeyID))

		// Try to add ICA with different VID — should fail
		txResult, err = AddNocX509IcaCert(nocCert2Path, vendorAccount2)
		require.NoError(t, err)
		require.Equal(t, uint32(439), txResult.Code)

		// Add second ICA cert
		txResult, err = AddNocX509IcaCert(nocCert2Path, vendorAccount, "--schemaVersion", "0")
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Add cert copy
		txResult, err = AddNocX509IcaCert(nocCert1CopyPath, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// All ICA certs — cert1 (both serials), cert2, vid, schemaVersion
		out, err = QueryAllNocX509IcaCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocCert1SubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocCert1SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocCert1CopySerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocCert2Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocCert2SubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocCert2SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, nocVid))
		require.Contains(t, string(out), `"schemaVersion":0`)

		// NOC certs must NOT appear in the DA approved cert list
		out, err = QueryAllX509Certs()
		require.NoError(t, err)
		require.NotContains(t, string(out), nocRootCert1Subject)
		require.NotContains(t, string(out), nocRootCert1SubjectKeyID)
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1SerialNumber))
		require.NotContains(t, string(out), nocCert1Subject)
		require.NotContains(t, string(out), nocCert1SubjectKeyID)
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocCert1SerialNumber))

		// All NOC certs — root and ICA both present
		out, err = QueryAllNocX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert1SubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocCert1SubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocCert1SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocCert1CopySerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocCert2Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocCert2SubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocCert2SerialNumber))
	})

	t.Run("AddAndRevokeNocRootCert", func(t *testing.T) {
		// Add root cert copy
		txResult, err := AddNocRootCert(nocRootCert1CopyPath, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Add a Matter §6.4.5.4 VVSC chain (self-issued VVSC root, VVSC intermediate,
		// VVSC leaf) so the leaf-level operations have a §6.5.12-compliant chain to
		// exercise. NocLeafCert1 is a NOC end-entity (cA=FALSE / NOC profile) and is
		// no longer accepted by the stricter add-noc-x509-ica-cert handler.
		txResult, err = AddNocRootCert(vvscRootCert1Path, vendorAccount, "--is-vid-verification-signer=true")
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = AddNocX509IcaCert(vvscIcaCert1Path, vendorAccount, "--is-vid-verification-signer=true")
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = AddNocX509IcaCert(vvscLeafCert1Path, vendorAccount, "--is-vid-verification-signer=true")
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Verify root state before revocation
		out, err := QueryAllNocRootCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1CopySerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert2SerialNumber))

		// Verify ICA state before revocation
		out, err = QueryAllNocX509IcaCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocCert1SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocCert1CopySerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocCert2SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, vvscLeafCert1SerialNumber))

		// Try to revoke with different VID — should fail
		txResult, err = RevokeNocRootCert(nocRootCert1Subject, nocRootCert1SubjectKeyID, vendorAccount2)
		require.NoError(t, err)
		require.Equal(t, uint32(439), txResult.Code)

		// Revoke root cert without child flag — ICA must survive
		txResult, err = RevokeNocRootCert(nocRootCert1Subject, nocRootCert1SubjectKeyID, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// All revoked NOC root certs — both serials, ICA/leaf absent
		out, err = QueryAllRevokedNocRootCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s`, nocRootCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert1SubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1CopySerialNumber))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocCert1Subject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, vvscLeafCert1Subject))

		// Revoked NOC root cert by subject + SKID
		out, err = QueryRevokedNocRootCert(nocRootCert1Subject, nocRootCert1SubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s`, nocRootCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert1SubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1CopySerialNumber))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert2Subject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert3Subject))

		// DA revoked certs must NOT contain revoked NOC root certs
		out, err = QueryAllRevokedX509Certs()
		require.NoError(t, err)
		require.NotContains(t, string(out), nocRootCert1Subject)
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1SerialNumber))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1CopySerialNumber))

		// Active NOC root certs by VID — cert2 present, cert1 absent
		out, err = QueryNocRootCerts(nocVid)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert2Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert2SubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert2SerialNumber))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert1Subject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert1SubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1SerialNumber))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1CopySerialNumber))

		// Query by VID + SKID for cert1 — Not Found
		out, err = QueryNocCert("--vid", fmt.Sprintf("%d", nocVid), "--subject-key-id", nocRootCert1SubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert1Subject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert1SubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1SerialNumber))

		// Query by VID + SKID for cert2 — present with all fields
		out, err = QueryNocCert("--vid", fmt.Sprintf("%d", nocVid), "--subject-key-id", nocRootCert2SubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert2Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert2SubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert2SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectAsText":"%s"`, nocRootCert2SubjectAsText))
		require.Contains(t, string(out), `"schemaVersion":0`)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, nocVid))
		require.Contains(t, string(out), `"tq":1`)

		// Query by subject for cert1 — gone
		out, err = QueryNocSubjectCerts(nocRootCert1Subject)
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert1Subject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert1SubjectKeyID))

		// Query by SKID alone for cert1 — Not Found
		out, err = QueryNocCert("--subject-key-id", nocRootCert1SubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert1Subject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert1SubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1SerialNumber))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1CopySerialNumber))

		// ICA certs by VID — ICA and leaf still active
		out, err = QueryNocX509IcaCerts(nocVid)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, vvscLeafCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocCert1SubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, vvscLeafCert1SubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocCert1SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocCert1CopySerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, vvscLeafCert1SerialNumber))

		// All NOC certs — ICA/leaf present, revoked root1 absent
		out, err = QueryAllNocX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocCert1SubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocCert1SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocCert1CopySerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, vvscLeafCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, vvscLeafCert1SubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, vvscLeafCert1SerialNumber))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert1Subject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert1SubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1SerialNumber))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1CopySerialNumber))
	})

	t.Run("RevokeNocIcaCert", func(t *testing.T) {
		// Try to revoke with different VID — should fail
		txResult, err := RevokeNocX509IcaCert(nocCert1Subject, nocCert1SubjectKeyID, vendorAccount2)
		require.NoError(t, err)
		require.Equal(t, uint32(439), txResult.Code)

		// Revoke ICA cert without child flag — leaf must survive
		txResult, err = RevokeNocX509IcaCert(nocCert1Subject, nocCert1SubjectKeyID, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Revoked ICA list — cert1 present (with schemaVersion), leaf absent
		out, err := QueryAllRevokedNocX509IcaCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocCert1SubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocCert1SerialNumber))
		require.Contains(t, string(out), `"schemaVersion":0`)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, vvscLeafCert1Subject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, vvscLeafCert1SubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, vvscLeafCert1SerialNumber))

		// Revoked root list must not contain ICA or leaf
		out, err = QueryAllRevokedNocRootCerts()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocCert1SubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, vvscLeafCert1SubjectKeyID))

		// Query by subject for cert1 — gone
		out, err = QueryNocSubjectCerts(nocCert1Subject)
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocCert1Subject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocCert1SubjectKeyID))

		// Query by SKID alone for cert1 — Not Found
		out, err = QueryNocCert("--subject-key-id", nocCert1SubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocCert1Subject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocCert1SubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocCert1SerialNumber))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocCert1CopySerialNumber))

		// Active ICA certs by VID — only leaf remains
		out, err = QueryNocX509IcaCerts(nocVid)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, vvscLeafCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, vvscLeafCert1SubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocCert1Subject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocCert1SubjectKeyID))

		// Query by VID + SKID for cert1 — Not Found
		out, err = QueryNocCert("--vid", fmt.Sprintf("%d", nocVid), "--subject-key-id", nocCert1SubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		// Query by VID + SKID for leaf — present
		out, err = QueryNocCert("--vid", fmt.Sprintf("%d", nocVid), "--subject-key-id", vvscLeafCert1SubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, vvscLeafCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, vvscLeafCert1SubjectKeyID))

		// All NOC certs — leaf present, cert1 and root1 absent
		out, err = QueryAllNocX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, vvscLeafCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, vvscLeafCert1SubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, vvscLeafCert1SerialNumber))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocCert1Subject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocCert1SubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocCert1SerialNumber))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert1Subject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert1SubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1SerialNumber))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1CopySerialNumber))
	})
}
