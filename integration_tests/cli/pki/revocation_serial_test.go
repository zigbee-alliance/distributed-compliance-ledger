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
)

// Cert fixtures — same paths as revocation_child_test.go, aliased here to
// keep this file self-contained.
const (
	revSerialRootCert1SerialNumber = revChildRootCert1SerialNumber
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
		all, err := GetAllX509Certs()
		require.NoError(t, err)
		require.True(t, containsApprovedCertSerial(all, revSerialRootCert1SerialNumber))
		require.True(t, containsApprovedCertSerial(all, revSerialRootCert2SerialNumber))
		require.True(t, containsApprovedCertSerial(all, revSerialIntermCert1SerialNumber))
		require.True(t, containsApprovedCertSerial(all, revSerialIntermCert2SerialNumber))
		require.True(t, containsApprovedCertSerial(all, revSerialLeafCertSerialNumber))
	})

	t.Run("RevokeIntermWithInvalidSerialNumber", func(t *testing.T) {
		txResult, err := RevokeX509Cert(revSerialIntermCertSubject, revSerialIntermCertSubjectKeyID, vendorAccount, RevokeNocCertOpts{SerialNumber: "invalid"})
		cliputils.RequireTxFailCode(t, txResult, err, 404)
	})

	t.Run("RevokeIntermWithSerialNumber3Only", func(t *testing.T) {
		// Revoke with serial number 3 only — child certs should remain
		txResult, err := RevokeX509Cert(revSerialIntermCertSubject, revSerialIntermCertSubjectKeyID, vendorAccount, RevokeNocCertOpts{SerialNumber: revSerialIntermCert1SerialNumber})
		cliputils.RequireTxOK(t, txResult, err)

		// Revoked list should contain only intermediate cert with serial 3.
		revoked, err := GetAllRevokedX509Certs()
		require.NoError(t, err)
		require.True(t, containsRevokedCertSerial(revoked, revSerialIntermCert1SerialNumber))
		require.False(t, containsRevokedCertSerial(revoked, revSerialIntermCert2SerialNumber))
		require.False(t, containsRevokedCertSerial(revoked, revSerialLeafCertSerialNumber))

		// Approved intermediate certs should contain only cert with serial 4.
		intermediate, err := GetX509Cert(revSerialIntermCertSubject, revSerialIntermCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, intermediate)
		require.True(t, containsCertSerial(intermediate.Certs, revSerialIntermCert2SerialNumber))
		require.False(t, containsCertSerial(intermediate.Certs, revSerialIntermCert1SerialNumber))

		// Leaf cert should still be present.
		leaf, err := GetX509Cert(revSerialLeafCertSubject, revSerialLeafCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, leaf)
		require.True(t, containsCertSerial(leaf.Certs, revSerialLeafCertSerialNumber))
	})

	t.Run("RevokeIntermWithSerial4AndChildFlag", func(t *testing.T) {
		// Revoke intermediate with serial 4 and its children
		txResult, err := RevokeX509Cert(revSerialIntermCertSubject, revSerialIntermCertSubjectKeyID, vendorAccount, RevokeNocCertOpts{SerialNumber: revSerialIntermCert2SerialNumber, RevokeChild: true})
		cliputils.RequireTxOK(t, txResult, err)

		// Revoked list should contain two intermediates and leaf.
		revoked, err := GetAllRevokedX509Certs()
		require.NoError(t, err)
		require.True(t, containsRevokedCertSerial(revoked, revSerialIntermCert1SerialNumber))
		require.True(t, containsRevokedCertSerial(revoked, revSerialIntermCert2SerialNumber))
		require.True(t, containsRevokedCertSerial(revoked, revSerialLeafCertSerialNumber))

		// Approved certs should contain only the two root certs.
		all, err := GetAllX509Certs()
		require.NoError(t, err)
		require.True(t, containsApprovedCertSubjectSerial(all, revSerialRootCertSubject, revSerialRootCert1SerialNumber))
		require.True(t, containsApprovedCertSubjectSerial(all, revSerialRootCertSubject, revSerialRootCert2SerialNumber))
		require.False(t, containsApprovedCertSubjectSerial(all, revSerialIntermCertSubject, revSerialIntermCert1SerialNumber))
		require.False(t, containsApprovedCertSubjectSerial(all, revSerialIntermCertSubject, revSerialIntermCert2SerialNumber))
		require.False(t, containsApprovedCertSubjectSerial(all, revSerialLeafCertSubject, revSerialLeafCertSerialNumber))
	})

	t.Run("ReAddCertsForRootRevocationTest", func(t *testing.T) {
		// Remove revoked certs
		txResult, err := RemoveX509Cert(revSerialIntermCertSubject, revSerialIntermCertSubjectKeyID, vendorAccount)
		cliputils.RequireTxOK(t, txResult, err)

		txResult, err = RemoveX509Cert(revSerialLeafCertSubject, revSerialLeafCertSubjectKeyID, vendorAccount)
		cliputils.RequireTxOK(t, txResult, err)

		// Re-add intermediate and leaf certs
		txResult, err = AddX509Cert(revSerialIntermCert1Path, vendorAccount)
		cliputils.RequireTxOK(t, txResult, err)

		txResult, err = AddX509Cert(revSerialIntermCert2Path, vendorAccount)
		cliputils.RequireTxOK(t, txResult, err)

		txResult, err = AddX509Cert(revSerialLeafCertPath, vendorAccount)
		cliputils.RequireTxOK(t, txResult, err)

		// Verify all certs present.
		all, err := GetAllX509Certs()
		require.NoError(t, err)
		require.True(t, containsApprovedCertSerial(all, revSerialRootCert1SerialNumber))
		require.True(t, containsApprovedCertSerial(all, revSerialIntermCert1SerialNumber))
		require.True(t, containsApprovedCertSerial(all, revSerialIntermCert2SerialNumber))
		require.True(t, containsApprovedCertSerial(all, revSerialLeafCertSerialNumber))
	})

	t.Run("ProposeRevokeRootWithInvalidSerialNumber", func(t *testing.T) {
		txResult, err := ProposeRevokeX509RootCert(revSerialRootCertSubject, revSerialRootCertSubjectKeyID, jack, X509ActionOpts{SerialNumber: "invalid"})
		cliputils.RequireTxFailCode(t, txResult, err, 404)
	})

	t.Run("ProposeAndApproveRevokeRootSerial1Only", func(t *testing.T) {
		// Propose revoke root with serial 1 (child certs should remain).
		txResult, err := ProposeRevokeX509RootCert(revSerialRootCertSubject, revSerialRootCertSubjectKeyID, jack, X509ActionOpts{SerialNumber: revSerialRootCert1SerialNumber})
		cliputils.RequireTxOK(t, txResult, err)

		txResult, err = ApproveRevokeX509RootCert(revSerialRootCertSubject, revSerialRootCertSubjectKeyID, alice, X509ActionOpts{SerialNumber: revSerialRootCert1SerialNumber})
		cliputils.RequireTxOK(t, txResult, err)

		// Revoked list should contain only one root entry with serial 1.
		revoked, err := GetAllRevokedX509Certs()
		require.NoError(t, err)
		require.True(t, containsRevokedCertSerial(revoked, revSerialRootCert1SerialNumber))
		require.False(t, containsRevokedCertSerial(revoked, revSerialRootCert2SerialNumber))
		require.False(t, containsRevokedCertSerial(revoked, revSerialIntermCert1SerialNumber))
		require.False(t, containsRevokedCertSerial(revoked, revSerialIntermCert2SerialNumber))
		require.False(t, containsRevokedCertSerial(revoked, revSerialLeafCertSerialNumber))

		// Verify root cert 1 was revoked by querying the specific cert.
		root, err := GetX509Cert(revSerialRootCertSubject, revSerialRootCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, root)
		require.True(t, containsCertSerial(root.Certs, revSerialRootCert2SerialNumber))
		require.False(t, containsCertSerial(root.Certs, revSerialRootCert1SerialNumber))

		// Intermediates and leaf should still be approved.
		all, err := GetAllX509Certs()
		require.NoError(t, err)
		require.True(t, containsApprovedCertSerial(all, revSerialIntermCert1SerialNumber))
		require.True(t, containsApprovedCertSerial(all, revSerialIntermCert2SerialNumber))
		require.True(t, containsApprovedCertSerial(all, revSerialLeafCertSerialNumber))
	})

	t.Run("ProposeAndApproveRevokeRootSerial2WithChild", func(t *testing.T) {
		// Propose revoke root with serial 2 and its children.
		txResult, err := ProposeRevokeX509RootCert(revSerialRootCertSubject, revSerialRootCertSubjectKeyID, jack, X509ActionOpts{
			SerialNumber: revSerialRootCert2SerialNumber,
			RevokeChild:  true,
		})
		cliputils.RequireTxOK(t, txResult, err)

		txResult, err = ApproveRevokeX509RootCert(revSerialRootCertSubject, revSerialRootCertSubjectKeyID, alice, X509ActionOpts{SerialNumber: revSerialRootCert2SerialNumber})
		cliputils.RequireTxOK(t, txResult, err)

		// Revoked list should contain two root, two intermediate and leaf.
		revoked, err := GetAllRevokedX509Certs()
		require.NoError(t, err)
		require.True(t, containsRevokedCertSerial(revoked, revSerialRootCert1SerialNumber))
		require.True(t, containsRevokedCertSerial(revoked, revSerialRootCert2SerialNumber))
		require.True(t, containsRevokedCertSerial(revoked, revSerialIntermCert1SerialNumber))
		require.True(t, containsRevokedCertSerial(revoked, revSerialIntermCert2SerialNumber))
		require.True(t, containsRevokedCertSerial(revoked, revSerialLeafCertSerialNumber))

		// Approved root certs should not include this chain's root or intermediate (by Subject+SKID).
		approvedRoots, err := GetAllX509RootCerts()
		require.NoError(t, err)
		if approvedRoots != nil {
			for _, id := range approvedRoots.Certs {
				if id.Subject == revSerialRootCertSubject && id.SubjectKeyId == revSerialRootCertSubjectKeyID {
					t.Fatalf("approved root certs still contains revoked root: %+v", id)
				}
				if id.Subject == revSerialIntermCertSubject && id.SubjectKeyId == revSerialIntermCertSubjectKeyID {
					t.Fatalf("approved root certs still contains intermediate: %+v", id)
				}
			}
		}
	})
}
