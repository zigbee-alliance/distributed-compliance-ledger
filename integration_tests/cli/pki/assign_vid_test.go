package pki

import (
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
)

const (
	rootCertWithVid = 65521
)

func TestPKIAssignVid(t *testing.T) {
	vendorAdminAccount := cliputils.CreateAccount(t, "VendorAdmin")

	t.Run("AssignVidToRootCertThatAlreadyHasVid_Fails", func(t *testing.T) {
		// Reuse google_root_cert_gsr4 already on the ledger from
		// TestPKIAddVendorX509Certificates (added with VID=pkiDemoVid=1, so its
		// stored Vid is non-zero). It's ECDSA P-256 + SHA-256 — Matter R1.6
		// regenerated/RSA roots (e.g. google_root_cert_r2) fail the §6.2.2.5
		// VerifyECDSAP256SHA256 check at propose time, so we cannot freshly
		// propose them here. The handler's assignVid loop skips certs whose
		// stored Vid is already non-zero and returns ErrNotEmptyVid when
		// nothing was modified, which is exactly the path under test.
		txResult, err := AssignVid(addVendorRootCertSubject, addVendorRootCertSubjectKeyID, rootCertWithVid, vendorAdminAccount)
		// expect error or non-zero code
		if err != nil {
			require.Contains(t, err.Error(), "vid is not empty")
		} else if txResult != nil {
			require.Contains(t, txResult.RawLog, "vid is not empty")
		}
	})
}
