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
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

const (
	revChildRootCert1SerialNumber = "1"
	revChildRootCert2SerialNumber = "2"
	revChildRootCertVid           = 65521

	revChildIntermCert1Path         = "../../constants/intermediate_with_same_subject_and_skid_1"
	revChildIntermCert1SerialNumber = "3"
	revChildIntermCert2Path         = "../../constants/intermediate_with_same_subject_and_skid_2"
	revChildIntermCert2SerialNumber = "4"

	revChildLeafCertPath         = "../../constants/leaf_with_same_subject_and_skid"
	revChildLeafCertSerialNumber = "5"

	revChildRootCertSubject      = "MIGCMQswCQYDVQQGEwJVUzERMA8GA1UECBMITmV3IFlvcmsxETAPBgNVBAcTCE5ldyBZb3JrMRgwFgYDVQQKEw9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsTEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMTD3d3dy5leGFtcGxlLmNvbQ=="
	revChildRootCertSubjectKeyID = "C1:48:66:ED:6F:23:D8:28:1A:D9:37:7C:58:AC:3F:DA:04:C1:41:E8"

	revChildIntermCertSubject      = "MEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpTb21lLVN0YXRlMSEwHwYDVQQKExhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQ="
	revChildIntermCertSubjectKeyID = "A1:E0:92:89:FA:18:82:12:14:9D:B8:AE:19:43:BE:44:31:6B:F1:F5"

	revChildLeafCertSubject      = "MEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpTb21lLVN0YXRlMSEwHwYDVQQKExhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQ="
	revChildLeafCertSubjectKeyID = "90:81:84:C7:EC:B8:81:14:66:61:2F:82:BB:E9:51:67:F2:4D:99:A3"
)

// Root certs (root_with_same_subject_and_skid_1/2) are already on chain from TestPKICombineCerts,
// so SetupCerts only adds intermediate and leaf certs.
// Root revocation tests run in TestPKIRevocationWithSerialNumber to avoid state conflicts.
func TestPKIRevocationWithRevokingChild(t *testing.T) {
	vendorAccount := fmt.Sprintf("vendor_account_%d", revChildRootCertVid)
	cliputils.CreateVendorAccount(t, vendorAccount, revChildRootCertVid)

	t.Run("SetupCerts", func(t *testing.T) {
		// Root certs 1 and 2 are already approved on-chain from TestPKICombineCerts.
		// Verify they exist before proceeding.
		rootCert, err := GetX509Cert(revChildRootCertSubject, revChildRootCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, rootCert)
		require.Equal(t, revChildRootCertSubject, rootCert.Subject)

		// Add intermediate cert 1
		txResult, err := AddX509Cert(revChildIntermCert1Path, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Add intermediate cert 2
		txResult, err = AddX509Cert(revChildIntermCert2Path, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Add leaf cert
		txResult, err = AddX509Cert(revChildLeafCertPath, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Verify all certs exist
		all, err := GetAllX509Certs()
		require.NoError(t, err)
		require.True(t, containsApprovedCertSerial(all, revChildRootCert1SerialNumber))
		require.True(t, containsApprovedCertSerial(all, revChildIntermCert1SerialNumber))
		require.True(t, containsApprovedCertSerial(all, revChildIntermCert2SerialNumber))
		require.True(t, containsApprovedCertSerial(all, revChildLeafCertSerialNumber))
	})

	t.Run("RevokeIntermediateCertWithChildFlag", func(t *testing.T) {
		// Revoke intermediate certs and their child certificates
		txResult, err := RevokeX509Cert(revChildIntermCertSubject, revChildIntermCertSubjectKeyID, vendorAccount, RevokeNocCertOpts{RevokeChild: true})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Revoked certs should contain both intermediate and leaf
		revoked, err := GetAllRevokedX509Certs()
		require.NoError(t, err)
		require.True(t, containsRevokedCertSerial(revoked, revChildIntermCert1SerialNumber))
		require.True(t, containsRevokedCertSerial(revoked, revChildIntermCert2SerialNumber))
		require.True(t, containsRevokedCertSerial(revoked, revChildLeafCertSerialNumber))

		// Approved certs should contain only two root certs (no intermediate/leaf).
		all, err := GetAllX509Certs()
		require.NoError(t, err)
		require.True(t, containsApprovedCertSerial(all, revChildRootCert1SerialNumber))
		require.True(t, containsApprovedCertSerial(all, revChildRootCert2SerialNumber))
		require.False(t, containsApprovedCertSerial(all, revChildIntermCert1SerialNumber))
		require.False(t, containsApprovedCertSerial(all, revChildIntermCert2SerialNumber))
		require.False(t, containsApprovedCertSerial(all, revChildLeafCertSerialNumber))
	})

	t.Run("ReAddCertsAfterRevocation", func(t *testing.T) {
		// Remove revoked certs from revoked list
		txResult, err := RemoveX509Cert(revChildIntermCertSubject, revChildIntermCertSubjectKeyID, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = RemoveX509Cert(revChildLeafCertSubject, revChildLeafCertSubjectKeyID, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Re-add intermediate cert 1
		txResult, err = AddX509Cert(revChildIntermCert1Path, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Re-add intermediate cert 2
		txResult, err = AddX509Cert(revChildIntermCert2Path, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Re-add leaf cert
		txResult, err = AddX509Cert(revChildLeafCertPath, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})
}
