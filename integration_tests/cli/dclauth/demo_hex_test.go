package dclauth

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// It creates a Vendor account using a hex-format VID, adds a model with the hex VID,
// verifies that using a different VID is rejected, updates the model, and queries it.
func TestAuthDemoHex(t *testing.T) {
	jack := testconstants.JackAccount

	// Use random VID/PID expressed as hex to avoid collisions across test runs.
	vid := rand.Intn(65534) + 1
	pid := rand.Intn(65534) + 1
	vidHex := fmt.Sprintf("0x%X", vid)
	pidHex := fmt.Sprintf("0x%X", pid)

	// Generate and add key
	name := "hexvendor" + utils.RandString()
	err := AddKey(name)
	require.NoError(t, err)

	userAddr, err := GetAddress(name)
	require.NoError(t, err)
	userPubkey, err := GetPubkey(name)
	require.NoError(t, err)

	jackAddr, err := GetAddress(jack)
	require.NoError(t, err)

	t.Run("ProposeVendorAccountWithHexVID", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "auth", "propose-add-account",
			"--info", "Jack is proposing this account",
			"--address", userAddr,
			"--pubkey", userPubkey,
			"--roles", "Vendor",
			"--vid", vidHex,
			"--from", jack,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Vendor accounts need only 1/3 approvals — should be active immediately
		out, err := QueryAllAccountsRaw()
		require.NoError(t, err)
		require.Contains(t, string(out), userAddr)

		out, err = QueryAccountRaw(userAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), userAddr)
		require.Contains(t, string(out), jackAddr)
		require.Contains(t, string(out), "Jack is proposing this account")

		out, err = QueryProposedAccount(userAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryAllProposedAccounts()
		require.NoError(t, err)
		require.Contains(t, string(out), "[]")
	})

	t.Run("AddModelWithHexVID", func(t *testing.T) {
		productName := "Device #1"
		txResult, err := utils.ExecuteTx("tx", "model", "add-model",
			"--vid", vidHex,
			"--pid", pidHex,
			"--productName", productName,
			"--productLabel", "Device Description",
			"--commissioningCustomFlow", "0",
			"--deviceTypeID", "12",
			"--partNumber", "12",
			"--enhancedSetupFlowOptions", "0",
			"--from", userAddr,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("AddModelWithWrongVID_Fails", func(t *testing.T) {
		vidPlusOneHex := "0xA14"
		productName := "Device #1"

		txResult, err := utils.ExecuteTx("tx", "model", "add-model",
			"--vid", vidPlusOneHex,
			"--pid", pidHex,
			"--productName", productName,
			"--productLabel", "Device Description",
			"--commissioningCustomFlow", "0",
			"--deviceTypeID", "12",
			"--partNumber", "12",
			"--enhancedSetupFlowOptions", "0",
			"--from", userAddr,
		)
		// With broadcast-mode sync, rejection might come at CLI level or on-chain.
		if err != nil {
			require.Contains(t, err.Error(), "vendorID")

			return
		}
		// Await on-chain result to drain tx from mempool before next subtest.
		txData, awaitErr := utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, awaitErr)
		require.Contains(t, string(txData), "vendorID")
	})

	t.Run("UpdateModel", func(t *testing.T) {
		productName := "Device #1"
		txResult, err := utils.ExecuteTx("tx", "model", "update-model",
			"--vid", vidHex,
			"--pid", pidHex,
			"--productName", productName,
			"--productLabel", "Device Description",
			"--partNumber", "12",
			"--enhancedSetupFlowOptions", "2",
			"--from", userAddr,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("QueryModel", func(t *testing.T) {
		out, err := utils.ExecuteCLI("query", "model", "get-model",
			"--vid", vidHex,
			"--pid", pidHex,
			"-o", "json",
		)
		require.NoError(t, err)
		require.Contains(t, string(out), `"vid":`+formatInt(vid))
		require.Contains(t, string(out), `"pid":`+formatInt(pid))
		require.Contains(t, string(out), "Device #1")
	})
}

func formatInt(n int) string {
	return itoa(n)
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	neg := false
	if n < 0 {
		neg = true
		n = -n
	}
	var buf [20]byte
	pos := len(buf)
	for n > 0 {
		pos--
		buf[pos] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		pos--
		buf[pos] = '-'
	}

	return string(buf[pos:])
}
