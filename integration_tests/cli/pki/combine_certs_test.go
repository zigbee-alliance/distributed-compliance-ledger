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
	rootWithSameSubjectAndSkidSubject      = "MIGCMQswCQYDVQQGEwJVUzERMA8GA1UECBMITmV3IFlvcmsxETAPBgNVBAcTCE5ldyBZb3JrMRgwFgYDVQQKEw9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsTEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMTD3d3dy5leGFtcGxlLmNvbQ=="
	rootWithSameSubjectAndSkidSubjectKeyID = "C1:48:66:ED:6F:23:D8:28:1A:D9:37:7C:58:AC:3F:DA:04:C1:41:E8"
)

// Tests that multiple certificates with the same subject and SKID can coexist.
func TestPKICombineCerts(t *testing.T) {
	jack := testconstants.JackAccount
	alice := testconstants.AliceAccount

	vid := 65521

	vendorAccount := fmt.Sprintf("vendor_account_%d", vid)
	cliputils.CreateVendorAccount(t, vendorAccount, vid)

	t.Run("ProposeAndApproveFirstRootCert", func(t *testing.T) {
		txResult, err := ProposeAddX509RootCert(rootWithSameSubjectAndSkid1Path, jack, X509ProposeOpts{VID: vid})
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

		cert, err := GetX509Cert(rootWithSameSubjectAndSkidSubject, rootWithSameSubjectAndSkidSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, cert)
		require.Equal(t, rootWithSameSubjectAndSkidSubject, cert.Subject)
	})

	t.Run("ProposeAndApproveSecondRootCert_SameSubjectSkid", func(t *testing.T) {
		txResult, err := ProposeAddX509RootCert(rootWithSameSubjectAndSkid2Path, jack, X509ProposeOpts{VID: vid})
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

		// Now both certs with same subject+skid should coexist (2 entries in Certs).
		cert, err := GetX509Cert(rootWithSameSubjectAndSkidSubject, rootWithSameSubjectAndSkidSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, cert)
		require.Equal(t, rootWithSameSubjectAndSkidSubject, cert.Subject)
		require.Len(t, cert.Certs, 2)
	})
}
