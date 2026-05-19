package pki

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

const (
	removeNocRootCertVid = 65521
	removeNocOtherVid    = 65522

	removeNocRootCert1SerialNumber = "47211865327720222621302679792296833381734533449"
	removeNocRootCert1CopySerial   = "460647353168152946606945669687905527879095841977"
	removeNocIntermCert1Serial     = "631388393741945881054190991612463928825155142122"
	removeNocIntermCert2Serial     = "169445068204646961882009388640343665944683778293"
	removeNocLeafCertSerial        = "281347277961838999749763518155363401757954575313"

	removeNocRootCertSubject      = "MHoxCzAJBgNVBAYTAlVaMRMwEQYDVQQIDApTb21lIFN0YXRlMREwDwYDVQQHDAhUYXNoa2VudDEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMQ4wDAYDVQQDDAVOT0MtMQ=="
	removeNocRootCertSubjectKeyID = "44:EB:4C:62:6B:25:48:CD:A2:B3:1C:87:41:5A:08:E7:2B:B9:83:26"

	removeNocIntermCertSubject      = "MIGCMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRQwEgYDVQQDDAtOT0MtY2hpbGQtMQ=="
	removeNocIntermCertSubjectKeyID = "02:72:6E:BC:BB:EF:D6:BD:8D:9B:42:AE:D4:3C:C0:55:5F:66:3A:B3"

	removeNocLeafCertSubject      = "MIGBMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRMwEQYDVQQDDApOT0MtbGVhZi0x"
	removeNocLeafCertSubjectKeyID = "77:1F:DB:C4:4C:B1:29:7E:3C:EB:3E:D8:2A:38:0B:63:06:07:00:01"
)

// TestPKIRemoveNocCertificates translates pki-remove-noc-certificates.sh.
func TestPKIRemoveNocCertificates(t *testing.T) {
	jack := testconstants.JackAccount

	vendorAccount65521 := fmt.Sprintf("vendor_account_%d", removeNocRootCertVid)
	cliputils.CreateVendorAccount(t, vendorAccount65521, removeNocRootCertVid)

	vendorAccount65522 := fmt.Sprintf("vendor_account_%d", removeNocOtherVid)
	cliputils.CreateVendorAccount(t, vendorAccount65522, removeNocOtherVid)

	// Prior NOC tests add noc_root_cert_1, noc_cert_1, and noc_leaf_cert_1 under VID 24582 (nocVid)
	// and leave them in the revoked state. The unique-cert store retains their serial entries,
	// preventing re-addition by VID 65521. Remove them via the owning VID 24582 account first.
	vendorAccount24582 := fmt.Sprintf("vendor_account_%d", nocVid)
	cliputils.CreateVendorAccount(t, vendorAccount24582, nocVid)
	RemoveNocRootCert(nocRootCert1Subject, nocRootCert1SubjectKeyID, vendorAccount24582) //nolint:errcheck
	RemoveNocCert(nocCert1Subject, nocCert1SubjectKeyID, vendorAccount24582)             //nolint:errcheck
	RemoveNocCert(nocLeafCert1Subject, nocLeafCert1SubjectKeyID, vendorAccount24582)     //nolint:errcheck

	t.Run("SetupCerts", func(t *testing.T) {
		// Add root cert
		txResult, err := AddNocRootCert(nocRootCert1Path, vendorAccount65521)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Add ICA certs
		txResult, err = AddNocX509IcaCert(nocCert1Path, vendorAccount65521)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = AddNocX509IcaCert(nocCert1CopyPath, vendorAccount65521)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Add leaf cert
		txResult, err = AddNocX509IcaCert(nocLeafCert1Path, vendorAccount65521)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryAllNocX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), removeNocRootCert1SerialNumber)
		require.Contains(t, string(out), removeNocIntermCert1Serial)
		require.Contains(t, string(out), removeNocIntermCert2Serial)
		require.Contains(t, string(out), removeNocLeafCertSerial)
	})

	t.Run("RevokeAndRemoveIcaCert", func(t *testing.T) {
		// Revoke first ICA cert
		txResult, err := RevokeNocX509IcaCert(removeNocIntermCertSubject, removeNocIntermCertSubjectKeyID, vendorAccount65521,
			"--serial-number", removeNocIntermCert1Serial,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Try to remove with invalid serial
		txResult, err = RemoveNocCert(removeNocIntermCertSubject, removeNocIntermCertSubjectKeyID, vendorAccount65521,
			"--serial-number", "invalid",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(404), txResult.Code)

		// Try to remove when sender is not Vendor
		txResult, err = RemoveNocCert(removeNocIntermCertSubject, removeNocIntermCertSubjectKeyID, jack,
			"--serial-number", removeNocIntermCert1Serial,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(4), txResult.Code)

		// Try to remove with different VID vendor (returns 439 ErrCertVidNotEqualAccountVid, not 4)
		txResult, err = RemoveNocCert(removeNocIntermCertSubject, removeNocIntermCertSubjectKeyID, vendorAccount65522,
			"--serial-number", removeNocIntermCert1Serial,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(439), txResult.Code)

		// Remove revoked ICA cert by serial
		txResult, err = RemoveNocCert(removeNocIntermCertSubject, removeNocIntermCertSubjectKeyID, vendorAccount65521,
			"--serial-number", removeNocIntermCert1Serial,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Only second ICA cert should remain
		out, err := QueryNocCert("--subject", removeNocIntermCertSubject, "--subject-key-id", removeNocIntermCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), removeNocIntermCert2Serial)
		require.NotContains(t, string(out), removeNocIntermCert1Serial)

		// Remove remaining ICA cert by subject+subjectKeyID
		txResult, err = RemoveNocCert(removeNocIntermCertSubject, removeNocIntermCertSubjectKeyID, vendorAccount65521)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err = QueryNocCert("--subject", removeNocIntermCertSubject, "--subject-key-id", removeNocIntermCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		// This test's ICA certs (subject = removeNocIntermCertSubject) should not be in the revoked list.
		// (Prior NOC tests may have left other subjects in the revoked list, so we cannot assert "[]".)
		out, err = QueryAllRevokedNocX509IcaCerts()
		require.NoError(t, err)
		require.NotContains(t, string(out), removeNocIntermCertSubject)
	})

	t.Run("RemoveLeafCert", func(t *testing.T) {
		txResult, err := RemoveNocCert(removeNocLeafCertSubject, removeNocLeafCertSubjectKeyID, vendorAccount65521)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryNocCert("--subject", removeNocLeafCertSubject, "--subject-key-id", removeNocLeafCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
	})

	t.Run("RemoveNocRootCert", func(t *testing.T) {
		// Add root cert copy
		txResult, err := AddNocRootCert(nocRootCert1CopyPath, vendorAccount65521)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Re-add ICA cert
		txResult, err = AddNocX509IcaCert(nocCert1Path, vendorAccount65521)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Try to remove root with invalid serial
		txResult, err = RemoveNocRootCert(removeNocRootCertSubject, removeNocRootCertSubjectKeyID, vendorAccount65521,
			"--serial-number", "invalid",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(404), txResult.Code)

		// Try to remove when sender is not Vendor
		txResult, err = RemoveNocRootCert(removeNocRootCertSubject, removeNocRootCertSubjectKeyID, jack,
			"--serial-number", removeNocRootCert1SerialNumber,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(4), txResult.Code)

		// Try to remove with different VID vendor (returns 439 ErrCertVidNotEqualAccountVid, not 4)
		txResult, err = RemoveNocRootCert(removeNocRootCertSubject, removeNocRootCertSubjectKeyID, vendorAccount65522,
			"--serial-number", removeNocRootCert1SerialNumber,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(439), txResult.Code)

		// Revoke root cert
		txResult, err = RevokeNocRootCert(removeNocRootCertSubject, removeNocRootCertSubjectKeyID, vendorAccount65521,
			"--serial-number", removeNocRootCert1SerialNumber,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Remove revoked root cert by serial
		txResult, err = RemoveNocRootCert(removeNocRootCertSubject, removeNocRootCertSubjectKeyID, vendorAccount65521,
			"--serial-number", removeNocRootCert1SerialNumber,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Only copy root should remain
		out, err := QueryNocRootCerts(removeNocRootCertVid)
		require.NoError(t, err)
		require.Contains(t, string(out), removeNocRootCert1CopySerial)
		require.NotContains(t, string(out), removeNocRootCert1SerialNumber)

		// Re-add root cert and then remove both
		txResult, err = AddNocRootCert(nocRootCert1Path, vendorAccount65521)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Remove all root certs by subject+subjectKeyID (no serial)
		txResult, err = RemoveNocRootCert(removeNocRootCertSubject, removeNocRootCertSubjectKeyID, vendorAccount65521)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err = QueryNocCert("--subject", removeNocRootCertSubject, "--subject-key-id", removeNocRootCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		// ICA cert should still be present
		out, err = QueryAllNocX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), removeNocIntermCert1Serial)
	})
}
