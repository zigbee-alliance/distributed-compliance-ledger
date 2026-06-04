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

	revChildRootCertSubject      = "MIGCMQswCQYDVQQGEwJVUzERMA8GA1UECAwITmV3IFlvcmsxETAPBgNVBAcMCE5ldyBZb3JrMRgwFgYDVQQKDA9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsMEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMMD3d3dy5leGFtcGxlLmNvbQ=="
	revChildRootCertSubjectKeyID = "33:5E:0C:07:44:F8:B5:9C:CD:55:01:9B:6D:71:23:83:6F:D0:D4:BE"

	revChildIntermCertSubject      = "MEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQ="
	revChildIntermCertSubjectKeyID = "2E:13:3B:44:52:2C:30:E9:EC:FB:45:FA:5D:E5:04:0A:C1:C6:E6:B9"

	revChildLeafCertSubject      = "MEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQ="
	revChildLeafCertSubjectKeyID = "12:16:55:8E:5E:2A:DF:04:D7:E6:FE:D1:53:69:61:98:EF:17:2F:03"
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
		out, err := QueryX509Cert(revChildRootCertSubject, revChildRootCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, revChildRootCertSubject))

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
		out, err = QueryAllX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, revChildRootCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, revChildIntermCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, revChildRootCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, revChildIntermCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, revChildLeafCertSubjectKeyID))
	})

	t.Run("RevokeIntermediateCertWithChildFlag", func(t *testing.T) {
		// Revoke intermediate certs and their child certificates
		txResult, err := RevokeX509Cert(revChildIntermCertSubject, revChildIntermCertSubjectKeyID, vendorAccount,
			"--revoke-child=true",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Revoked certs should contain both intermediate and leaf
		out, err := QueryAllRevokedX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, revChildIntermCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, revChildIntermCertSubjectKeyID))
		require.Contains(t, string(out), revChildLeafCertSubjectKeyID)
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revChildIntermCert1SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revChildIntermCert2SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revChildLeafCertSerialNumber))

		// Approved certs should contain only two root certs
		out, err = QueryAllX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, revChildRootCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, revChildRootCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revChildRootCert1SerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revChildRootCert2SerialNumber))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, revChildIntermCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revChildIntermCert1SerialNumber))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revChildIntermCert2SerialNumber))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, revChildLeafCertSerialNumber))
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
