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
	rootCertWithVid = 65521

	// Use google_root_cert_r2 (no embedded Matter VID) to avoid conflicting with root_cert used by TestPKIDemo.
	assignVidTestRootCertPath         = "../../constants/google_root_cert_r2"
	assignVidTestRootCertSubject      = "MEcxCzAJBgNVBAYTAlVTMSIwIAYDVQQKExlHb29nbGUgVHJ1c3QgU2VydmljZXMgTExDMRQwEgYDVQQDEwtHVFMgUm9vdCBSMg=="
	assignVidTestRootCertSubjectKeyID = "BB:FF:CA:8E:23:9F:4F:99:CA:DB:E2:68:A6:A5:15:27:17:1E:D9:0E"
)

// TestPKIAssignVid translates pki-assign-vid.sh.
func TestPKIAssignVid(t *testing.T) {
	jack := testconstants.JackAccount
	alice := testconstants.AliceAccount

	vendorAdminAccount := cliputils.CreateAccount(t, "VendorAdmin")

	t.Run("AssignVidToRootCertThatAlreadyHasVid_Fails", func(t *testing.T) {
		// Propose and approve a root cert with an assigned VID.
		// Use google_root_cert_r2 (no embedded Matter VID) to avoid conflicting with root_cert
		// used by TestPKIDemo.
		txResult, err := ProposeAddX509RootCert(assignVidTestRootCertPath, jack,
			"--vid", fmt.Sprintf("%d", rootCertWithVid),
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = ApproveAddX509RootCert(assignVidTestRootCertSubject, assignVidTestRootCertSubjectKeyID, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Assign VID to a cert that already has a VID — should fail.
		txResult, err = AssignVid(assignVidTestRootCertSubject, assignVidTestRootCertSubjectKeyID, rootCertWithVid, vendorAdminAccount)
		// expect error or non-zero code
		if err != nil {
			require.Contains(t, err.Error(), "vid is not empty")
		} else if txResult != nil {
			require.Contains(t, txResult.RawLog, "vid is not empty")
		}
	})

	_ = jack
	_ = alice
}
