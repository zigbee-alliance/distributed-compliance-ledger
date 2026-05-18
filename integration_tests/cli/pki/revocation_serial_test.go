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

package pki

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// Constants for pki-revocation-with-serial-number.sh — reuses same cert paths as revocation_child_test.go.
// We alias them here to keep this file self-contained.
const (
	revSerialRootCert1Path         = revChildRootCert1Path
	revSerialRootCert1SerialNumber = revChildRootCert1SerialNumber
	revSerialRootCert2Path         = revChildRootCert2Path
	revSerialRootCert2SerialNumber = revChildRootCert2SerialNumber
	revSerialRootCertVid           = revChildRootCertVid

	revSerialIntermCert1Path         = revChildIntermCert1Path
	revSerialIntermCert1SerialNumber = revChildIntermCert1SerialNumber
	revSerialIntermCert2Path         = revChildIntermCert2Path
	revSerialIntermCert2SerialNumber = revChildIntermCert2SerialNumber

	revSerialLeafCertPath         = revChildLeafCertPath
	revSerialLeafCertSerialNumber = revChildLeafCertSerialNumber

	revSerialRootCertSubject      = revChildRootCertSubject
	revSerialRootCertSubjectKeyID = revChildRootCertSubjectKeyID

	revSerialIntermCertSubject      = revChildIntermCertSubject
	revSerialIntermCertSubjectKeyID = revChildIntermCertSubjectKeyID

	revSerialLeafCertSubject      = revChildLeafCertSubject
	revSerialLeafCertSubjectKeyID = revChildLeafCertSubjectKeyID
)

// TestPKIRevocationWithSerialNumber translates pki-revocation-with-serial-number.sh.
// Root certs are already on-chain from TestPKICombineCerts.
// Intermediate/leaf certs are re-added by TestPKIRevocationWithRevokingChild.ReAddCertsAfterRevocation.
// This test also covers RevokeRootCertWithChildFlag (moved here from revocation_child_test.go
// to avoid permanently revoking root certs before this test runs).
func TestPKIRevocationWithSerialNumber(t *testing.T) {
	jack := testconstants.JackAccount
	alice := testconstants.AliceAccount

	// Use same vendor account as revocation_child_test.go (same VID).
	vendorAccount := fmt.Sprintf("vendor_account_%d", revSerialRootCertVid)
	cliputils.CreateVendorAccount(t, vendorAccount, revSerialRootCertVid)

	t.Run("VerifyCertsOnChain", func(t *testing.T) {
		// Root certs 1 and 2 are already approved (from TestPKICombineCerts).
		// Intermediate and leaf certs are already on-chain (re-added by TestPKIRevocationWithRevokingChild).
		out, err := QueryAllX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, revSerialRootCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, revSerialIntermCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialRootCert1SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialIntermCert1SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialRootCert2SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialIntermCert2SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialLeafCertSerialNumber))
	})

	t.Run("RevokeIntermWithInvalidSerialNumber", func(t *testing.T) {
		txResult, err := RevokeX509Cert(revSerialIntermCertSubject, revSerialIntermCertSubjectKeyID, vendorAccount,
			"--serial-number", "invalid",
		)
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), txResult.Code)
	})

	t.Run("RevokeIntermWithSerialNumber3Only", func(t *testing.T) {
		// Revoke with serial number 3 only — child certs should remain
		txResult, err := RevokeX509Cert(revSerialIntermCertSubject, revSerialIntermCertSubjectKeyID, vendorAccount,
			"--serial-number", revSerialIntermCert1SerialNumber,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Revoked list should contain only intermediate cert with serial 3
		out, err := QueryAllRevokedX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, revSerialIntermCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, revSerialIntermCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialIntermCert1SerialNumber))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialIntermCert2SerialNumber))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialLeafCertSerialNumber))

		// Approved intermediate certs should contain only cert with serial 4
		out, err = QueryX509Cert(revSerialIntermCertSubject, revSerialIntermCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, revSerialIntermCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, revSerialIntermCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialIntermCert2SerialNumber))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialIntermCert1SerialNumber))

		// Leaf cert should still be present
		out, err = QueryX509Cert(revSerialLeafCertSubject, revSerialLeafCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, revSerialLeafCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, revSerialLeafCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialLeafCertSerialNumber))
	})

	t.Run("RevokeIntermWithSerial4AndChildFlag", func(t *testing.T) {
		// Revoke intermediate with serial 4 and its children
		txResult, err := RevokeX509Cert(revSerialIntermCertSubject, revSerialIntermCertSubjectKeyID, vendorAccount,
			"--serial-number", revSerialIntermCert2SerialNumber,
			"--revoke-child=true",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Revoked list should contain two intermediate and leaf
		out, err := QueryAllRevokedX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, revSerialIntermCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, revSerialIntermCertSubjectKeyID))
		require.Contains(t, string(out), revSerialLeafCertSubjectKeyID)
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialIntermCert1SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialIntermCert2SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialLeafCertSerialNumber))

		// Approved certs should contain only two root certs.
		// Use subjectKeyId to distinguish intermediate/leaf from root (they share the same subject).
		out, err = QueryAllX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, revSerialRootCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, revSerialRootCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialRootCert1SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialRootCert2SerialNumber))
		require.NotContains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, revSerialIntermCertSubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, revSerialLeafCertSubjectKeyID))
	})

	t.Run("ReAddCertsForRootRevocationTest", func(t *testing.T) {
		// Remove revoked certs
		txResult, err := RemoveX509Cert(revSerialIntermCertSubject, revSerialIntermCertSubjectKeyID, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = RemoveX509Cert(revSerialLeafCertSubject, revSerialLeafCertSubjectKeyID, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Re-add intermediate and leaf certs
		txResult, err = AddX509Cert(revSerialIntermCert1Path, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = AddX509Cert(revSerialIntermCert2Path, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = AddX509Cert(revSerialLeafCertPath, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Verify all certs present
		out, err := QueryAllX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, revSerialRootCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, revSerialIntermCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, revSerialRootCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, revSerialIntermCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, revSerialLeafCertSubjectKeyID))
	})

	t.Run("ProposeRevokeRootWithInvalidSerialNumber", func(t *testing.T) {
		txResult, err := ProposeRevokeX509RootCert(revSerialRootCertSubject, revSerialRootCertSubjectKeyID, jack,
			"--serial-number", "invalid",
		)
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), txResult.Code)
	})

	t.Run("ProposeAndApproveRevokeRootSerial1Only", func(t *testing.T) {
		// Propose revoke root with serial 1 (child certs should remain)
		txResult, err := ProposeRevokeX509RootCert(revSerialRootCertSubject, revSerialRootCertSubjectKeyID, jack,
			"--serial-number", revSerialRootCert1SerialNumber,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = ApproveRevokeX509RootCert(revSerialRootCertSubject, revSerialRootCertSubjectKeyID, alice,
			"--serial-number", revSerialRootCert1SerialNumber,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Revoked list should contain one root with serial 1
		out, err := QueryAllRevokedX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, revSerialRootCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, revSerialRootCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialRootCert1SerialNumber))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialRootCert2SerialNumber))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialIntermCert1SerialNumber))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialIntermCert2SerialNumber))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialLeafCertSerialNumber))

		// Approved certs should still contain root 2, intermediates and leaf
		out, err = QueryAllX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, revSerialRootCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, revSerialIntermCertSubject))
		require.Contains(t, string(out), revSerialRootCertSubjectKeyID)
		require.Contains(t, string(out), revSerialIntermCertSubjectKeyID)
		require.Contains(t, string(out), revSerialLeafCertSubjectKeyID)
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialRootCert2SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialIntermCert1SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialIntermCert2SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialLeafCertSerialNumber))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialRootCert1SerialNumber))
	})

	t.Run("ProposeAndApproveRevokeRootSerial2WithChild", func(t *testing.T) {
		// Propose revoke root with serial 2 and its children
		txResult, err := ProposeRevokeX509RootCert(revSerialRootCertSubject, revSerialRootCertSubjectKeyID, jack,
			"--serial-number", revSerialRootCert2SerialNumber,
			"--revoke-child=true",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = ApproveRevokeX509RootCert(revSerialRootCertSubject, revSerialRootCertSubjectKeyID, alice,
			"--serial-number", revSerialRootCert2SerialNumber,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Revoked list should contain two root, two intermediate and leaf
		out, err := QueryAllRevokedX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, revSerialRootCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, revSerialIntermCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, revSerialRootCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, revSerialIntermCertSubjectKeyID))
		require.Contains(t, string(out), revSerialLeafCertSubjectKeyID)
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialRootCert1SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialIntermCert1SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialRootCert2SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialIntermCert2SerialNumber))
		require.Contains(t, string(out), revSerialLeafCertSerialNumber)

		// Approved root certs should be empty
		out, err = QueryAllX509RootCerts()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, revSerialRootCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, revSerialIntermCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, revSerialRootCertSubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, revSerialIntermCertSubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialRootCert1SerialNumber))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialIntermCert1SerialNumber))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialRootCert2SerialNumber))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revSerialIntermCert2SerialNumber))
	})
}
