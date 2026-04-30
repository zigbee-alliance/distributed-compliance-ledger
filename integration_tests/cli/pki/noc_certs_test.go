package pki

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

const (
	nocRootCert1Path       = "../../constants/noc_root_cert_1"
	nocRootCert1Subject    = "MHoxCzAJBgNVBAYTAlVaMRMwEQYDVQQIDApTb21lIFN0YXRlMREwDwYDVQQHDAhUYXNoa2VudDEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMQ4wDAYDVQQDDAVOT0MtMQ=="
	nocRootCert1SubjectKeyID = "44:EB:4C:62:6B:25:48:CD:A2:B3:1C:87:41:5A:08:E7:2B:B9:83:26"
	nocRootCert1SerialNumber = "47211865327720222621302679792296833381734533449"
	nocRootCert1SubjectAsText = "CN=NOC-1,OU=Testing Division,O=Example Company,L=Tashkent,ST=Some State,C=UZ"

	nocRootCert1CopyPath         = "../../constants/noc_root_cert_1_copy"
	nocRootCert1CopySerialNumber = "460647353168152946606945669687905527879095841977"

	nocRootCert2Path         = "../../constants/noc_root_cert_2"
	nocRootCert2Subject      = "MHoxCzAJBgNVBAYTAlVaMRMwEQYDVQQIDApTb21lIFN0YXRlMREwDwYDVQQHDAhUYXNoa2VudDEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMQ4wDAYDVQQDDAVOT0MtMg=="
	nocRootCert2SubjectKeyID = "CF:E6:DD:37:2B:4C:B2:B9:A9:F2:75:30:1C:AA:B1:37:1B:11:7F:1B"
	nocRootCert2SerialNumber = "332802481233145945539125204504842614737181725760"
	nocRootCert2SubjectAsText = "CN=NOC-2,OU=Testing Division,O=Example Company,L=Tashkent,ST=Some State,C=UZ"

	nocRootCert3Path         = "../../constants/noc_root_cert_3"
	nocRootCert3Subject      = "MFUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQxDjAMBgNVBAMMBU5PQy0z"
	nocRootCert3SubjectKeyID = "88:0D:06:D9:64:22:29:34:78:7F:8C:3B:AE:F5:08:93:86:8F:0D:20"
	nocRootCert3SerialNumber = "38457288443253426021793906708335409501754677187"
	nocRootCert3SubjectAsText = "CN=NOC-3,O=Internet Widgits Pty Ltd,ST=Some-State,C=AU"

	nocCert1Path         = "../../constants/noc_cert_1"
	nocCert1Subject      = "MIGCMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRQwEgYDVQQDDAtOT0MtY2hpbGQtMQ=="
	nocCert1SubjectKeyID = "02:72:6E:BC:BB:EF:D6:BD:8D:9B:42:AE:D4:3C:C0:55:5F:66:3A:B3"
	nocCert1SerialNumber = "631388393741945881054190991612463928825155142122"

	nocCert1CopyPath         = "../../constants/noc_cert_1_copy"
	nocCert1CopySerialNumber = "169445068204646961882009388640343665944683778293"

	nocCert2Path         = "../../constants/noc_cert_2"
	nocCert2Subject      = "MIGCMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRQwEgYDVQQDDAtOT0MtY2hpbGQtMg=="
	nocCert2SubjectKeyID = "87:48:A2:33:12:1F:51:5C:93:E6:90:40:4A:2C:AB:9E:D6:19:E5:AD"
	nocCert2SerialNumber = "361372967010167010646904372658654439710639340814"

	nocLeafCert1Path         = "../../constants/noc_leaf_cert_1"
	nocLeafCert1Subject      = "MIGBMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRMwEQYDVQQDDApOT0MtbGVhZi0x"
	nocLeafCert1SubjectKeyID = "77:1F:DB:C4:4C:B1:29:7E:3C:EB:3E:D8:2A:38:0B:63:06:07:00:01"
	nocLeafCert1SerialNumber = "281347277961838999749763518155363401757954575313"

	nocVid  = 24582
	nocVid2 = 4701
)

// TestPKINocCerts translates pki-noc-certs.sh.
func TestPKINocCerts(t *testing.T) {
	vendorAccount := fmt.Sprintf("vendor_account_%d", nocVid)
	cliputils.CreateVendorAccount(t, vendorAccount, nocVid)

	vendorAccount2 := fmt.Sprintf("vendor_account_%d", nocVid2)
	cliputils.CreateVendorAccount(t, vendorAccount2, nocVid2)

	t.Run("QueryAllEmpty", func(t *testing.T) {
		out, err := QueryNocRootCerts(nocVid)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryNocCert("--vid", fmt.Sprintf("%d", nocVid), "--subject-key-id", nocRootCert1SubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryAllNocRootCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), "[]")

		out, err = QueryNocCert("--subject", nocRootCert1Subject, "--subject-key-id", nocRootCert1SubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
	})

	t.Run("AddNocRootCerts", func(t *testing.T) {
		// Try to add intermediate cert using add-noc-x509-root-cert command — should fail
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
		out, err := QueryNocRootCerts(nocVid)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert1SubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert2Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert2SubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, nocVid))

		out, err = QueryAllNocRootCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert2Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert3Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, nocVid))
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, nocVid2))

		out, err = QueryNocCert("--subject", nocRootCert1Subject, "--subject-key-id", nocRootCert1SubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1SerialNumber))
		require.Contains(t, string(out), `"approvals":[]`)
	})

	t.Run("AddNocIcaCerts", func(t *testing.T) {
		// Add first ICA cert
		txResult, err := AddNocX509IcaCert(nocCert1Path, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Check child certs
		out, err := QueryChildX509Certs(nocRootCert1Subject, nocRootCert1SubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocCert1Subject))

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

		out, err = QueryAllNocX509IcaCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocCert1SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocCert1CopySerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocCert2Subject))

		// Approved x509 certs should be empty
		out, err = QueryAllX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), "[]")
		require.NotContains(t, string(out), nocRootCert1Subject)

		// All NOC certs should include both root and ICA
		out, err = QueryAllNocX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocCert1Subject))
	})

	t.Run("AddAndRevokeNocRootCert", func(t *testing.T) {
		// Add root cert copy
		txResult, err := AddNocRootCert(nocRootCert1CopyPath, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Add leaf cert
		txResult, err = AddNocX509IcaCert(nocLeafCert1Path, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Try to revoke with different VID — should fail
		txResult, err = RevokeNocRootCert(nocRootCert1Subject, nocRootCert1SubjectKeyID, vendorAccount2)
		require.NoError(t, err)
		require.Equal(t, uint32(439), txResult.Code)

		// Revoke root cert (without child flag — should not revoke ICA)
		txResult, err = RevokeNocRootCert(nocRootCert1Subject, nocRootCert1SubjectKeyID, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryAllRevokedNocRootCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s`, nocRootCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocRootCert1SubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, nocRootCert1CopySerialNumber))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocCert1Subject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocLeafCert1Subject))

		// Revoked NOC root certs should NOT appear in x509 revoked root certs
		out, err = QueryAllRevokedX509Certs()
		require.NoError(t, err)
		require.NotContains(t, string(out), nocRootCert1Subject)
	})

	t.Run("RevokeNocIcaCert", func(t *testing.T) {
		// Try to revoke with different VID — should fail
		txResult, err := RevokeNocX509IcaCert(nocCert1Subject, nocCert1SubjectKeyID, vendorAccount2)
		require.NoError(t, err)
		require.Equal(t, uint32(439), txResult.Code)

		// Revoke ICA cert (without child flag — should not revoke leaf)
		txResult, err = RevokeNocX509IcaCert(nocCert1Subject, nocCert1SubjectKeyID, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryAllRevokedNocX509IcaCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, nocCert1SubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocLeafCert1Subject))

		// Active ICA certs query by VID should only have leaf
		out, err = QueryNocX509IcaCerts(nocVid)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocLeafCert1Subject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocCert1Subject))

		// All NOC certs should not contain revoked ICA
		out, err = QueryAllNocX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocLeafCert1Subject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocCert1Subject))
		require.NotContains(t, string(out), nocRootCert1SerialNumber)
	})
}
