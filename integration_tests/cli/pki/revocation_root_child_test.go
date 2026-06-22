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
	"testing"

	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// TestPKIRevokeRootCertWholeSubjectWithChild ports the whole-subject (no serial
// number) root revocation flow from pki-revocation-with-revoking-child.sh:140-167:
//
//	dcld tx pki propose-revoke-x509-root-cert --subject ... --subject-key-id ... --revoke-child=true
//	dcld tx pki approve-revoke-x509-root-cert --subject ... --subject-key-id ...
//
// i.e. a trustee revokes a root by Subject/SubjectKeyID with no --serial-number,
// which revokes ALL serials of that root at once, and --revoke-child cascades the
// revocation to every certificate issued under it. This is distinct from the
// serial-number root revocation in revocation_serial_test.go, which targets a
// single serial.
//
// Chain selection: this revokes the google_root_cert_gsr4 → intermediate_cert_gsr4
// chain left APPROVED on-chain by TestPKIAddVendorX509Certificates (and used
// read-only by TestPKIAssignVid). Both run before this test. No later test depends
// on that chain (TestPKIRevocationWithSerialNumber uses the
// root_with_same_subject_and_skid chain instead), so revoking it here is safe.
//
// A dedicated chain is required because a revoked root can never be re-added — the
// revoke handlers do not release the unique Issuer/SerialNumber registration
// (no RemoveUniqueCertificate on the revoke path), so ProposeAddX509RootCert would
// fail with "certificate already exists". The shared revocation chain is therefore
// consumed permanently by the serial-number test and cannot be reused here.
func TestPKIRevokeRootCertWholeSubjectWithChild(t *testing.T) {
	jack := testconstants.JackAccount
	alice := testconstants.AliceAccount

	t.Run("VerifyChainOnChain", func(t *testing.T) {
		// Root (google_root_cert_gsr4) and its intermediate (intermediate_cert_gsr4)
		// must be approved before we can revoke them.
		root, err := GetX509Cert(addVendorRootCertSubject, addVendorRootCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, root, "google_root_cert_gsr4 must be approved on-chain (added by TestPKIAddVendorX509Certificates)")

		interm, err := GetX509Cert(addVendorIntermCertSubject, addVendorIntermCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, interm, "intermediate_cert_gsr4 must be approved on-chain (added by TestPKIAddVendorX509Certificates)")
	})

	t.Run("ProposeRevokeRootWholeSubjectWithChild", func(t *testing.T) {
		// Trustee proposes to revoke the root cert (whole subject, no serial) and
		// its children. With 3 trustees this single proposal is not yet enough
		// (revocation needs 2/3), so the root stays approved until the approval.
		txResult, err := ProposeRevokeX509RootCert(addVendorRootCertSubject, addVendorRootCertSubjectKeyID, jack, X509ActionOpts{RevokeChild: true})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Pending revocation is recorded.
		proposedRev, err := GetProposedRevokedX509RootCert(addVendorRootCertSubject, addVendorRootCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, proposedRev)

		// Root and intermediate are still approved (quorum not reached yet).
		root, err := GetX509Cert(addVendorRootCertSubject, addVendorRootCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, root)
		interm, err := GetX509Cert(addVendorIntermCertSubject, addVendorIntermCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, interm)
	})

	t.Run("ApproveRevokeRootWholeSubjectWithChild", func(t *testing.T) {
		// Second trustee approves → 2/3 quorum reached → root revoked.
		txResult, err := ApproveRevokeX509RootCert(addVendorRootCertSubject, addVendorRootCertSubjectKeyID, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Root is now revoked and no longer approved.
		revokedRoot, err := GetRevokedX509Cert(addVendorRootCertSubject, addVendorRootCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, revokedRoot)
		require.Equal(t, addVendorRootCertSubject, revokedRoot.Subject)
		require.Equal(t, addVendorRootCertSubjectKeyID, revokedRoot.SubjectKeyId)

		approvedRoot, err := GetX509Cert(addVendorRootCertSubject, addVendorRootCertSubjectKeyID)
		require.NoError(t, err)
		require.Nil(t, approvedRoot)

		// The child intermediate was cascaded into the revoked store too.
		revokedInterm, err := GetRevokedX509Cert(addVendorIntermCertSubject, addVendorIntermCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, revokedInterm)
		require.Equal(t, addVendorIntermCertSubject, revokedInterm.Subject)

		approvedInterm, err := GetX509Cert(addVendorIntermCertSubject, addVendorIntermCertSubjectKeyID)
		require.NoError(t, err)
		require.Nil(t, approvedInterm)

		// The approved-root list must no longer contain this root.
		approvedRoots, err := GetAllX509RootCerts()
		require.NoError(t, err)
		if approvedRoots != nil {
			for _, id := range approvedRoots.Certs {
				require.False(t,
					id.Subject == addVendorRootCertSubject && id.SubjectKeyId == addVendorRootCertSubjectKeyID,
					"approved root list still contains the revoked root: %+v", id)
			}
		}
	})
}
