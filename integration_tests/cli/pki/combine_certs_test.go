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
	rootWithSameSubjectAndSkid1Path = "../../constants/root_with_same_subject_and_skid_1"
	rootWithSameSubjectAndSkid2Path = "../../constants/root_with_same_subject_and_skid_2"
	// Subject and SKID match the actual "Example Company" cert files (not Amazon Root CA).
	rootWithSameSubjectAndSkidSubject      = "MIGCMQswCQYDVQQGEwJVUzERMA8GA1UECAwITmV3IFlvcmsxETAPBgNVBAcMCE5ldyBZb3JrMRgwFgYDVQQKDA9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsMEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMMD3d3dy5leGFtcGxlLmNvbQ=="
	rootWithSameSubjectAndSkidSubjectKeyID = "33:5E:0C:07:44:F8:B5:9C:CD:55:01:9B:6D:71:23:83:6F:D0:D4:BE"
)

// Tests that multiple certificates with the same subject and SKID can coexist.
func TestPKICombineCerts(t *testing.T) {
	jack := testconstants.JackAccount
	alice := testconstants.AliceAccount

	vid := 65521

	vendorAccount := fmt.Sprintf("vendor_account_%d", vid)
	cliputils.CreateVendorAccount(t, vendorAccount, vid)

	t.Run("ProposeAndApproveFirstRootCert", func(t *testing.T) {
		txResult, err := ProposeAddX509RootCert(rootWithSameSubjectAndSkid1Path, jack,
			"--vid", fmt.Sprintf("%d", vid),
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// With 3 trustees, quorum=2: jack proposes + alice approves = cert is approved.
		txResult, err = ApproveAddX509RootCert(rootWithSameSubjectAndSkidSubject, rootWithSameSubjectAndSkidSubjectKeyID, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryX509Cert(rootWithSameSubjectAndSkidSubject, rootWithSameSubjectAndSkidSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootWithSameSubjectAndSkidSubject))
	})

	t.Run("ProposeAndApproveSecondRootCert_SameSubjectSkid", func(t *testing.T) {
		txResult, err := ProposeAddX509RootCert(rootWithSameSubjectAndSkid2Path, jack,
			"--vid", fmt.Sprintf("%d", vid),
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// With 3 trustees, quorum=2: jack proposes + alice approves = cert is approved.
		txResult, err = ApproveAddX509RootCert(rootWithSameSubjectAndSkidSubject, rootWithSameSubjectAndSkidSubjectKeyID, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Now both certs with same subject+skid should coexist
		out, err := QueryX509Cert(rootWithSameSubjectAndSkidSubject, rootWithSameSubjectAndSkidSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootWithSameSubjectAndSkidSubject))
	})
}
