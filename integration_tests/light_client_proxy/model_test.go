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

package lightclientproxy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// TestLightClientProxyModel is the Go translation of
// integration_tests/light_client_proxy/model.sh.
//
//nolint:funlen
func TestLightClientProxyModel(t *testing.T) {
	skipIfDisabled(t)

	// 1. Random vid/pid/sv → every single-record model query returns
	//    Not Found via the proxy. (model.sh lines 13-41)
	mustRun(t, "NotFound_BeforeAdd", func(t *testing.T) {
		t.Helper()
		vid, pid, sv := randomUint16(), randomUint16(), randomUint16()

		out, qerr := queryWithRetry(LightClientProxyAddr,
			"query", "model", "get-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
		)
		require.NoError(t, qerr)
		assertContains(t, out, "Not Found", "get-model")

		out, qerr = queryWithRetry(LightClientProxyAddr,
			"query", "model", "vendor-models",
			"--vid", fmt.Sprintf("%d", vid),
		)
		require.NoError(t, qerr)
		assertContains(t, out, "Not Found", "vendor-models")

		out, qerr = queryWithRetry(LightClientProxyAddr,
			"query", "model", "model-version",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--softwareVersion", fmt.Sprintf("%d", sv),
		)
		require.NoError(t, qerr)
		assertContains(t, out, "Not Found", "model-version")

		out, qerr = queryWithRetry(LightClientProxyAddr,
			"query", "model", "all-model-versions",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
		)
		require.NoError(t, qerr)
		assertContains(t, out, "Not Found", "all-model-versions")
	})

	// 2. The all-models list is rejected by the proxy.
	mustRun(t, "ListQuery_Rejected", func(t *testing.T) {
		t.Helper()
		out, qerr := queryWithRetry(LightClientProxyAddr, "query", "model", "all-models")
		assertRejectionContains(t, out, qerr, listQueryRejection, "all-models")
	})

	// 3. Seed vendor + model + model-version via the full node.
	//    Account name is suffixed with utils.RandString() because the five
	//    light_client_proxy tests share one init_pool (see run-all.sh).
	vid, pid, sv := randomUint16(), randomUint16(), randomUint16()
	vendorAccount := "model_vendor_" + utils.RandString()
	const productLabel = "Device #1"

	mustRun(t, "Seed_VendorAndModel", func(t *testing.T) {
		t.Helper()
		_ = proposeVendorAccount(t, vendorAccount, vid)

		tx, err := utils.ExecuteTx(
			"tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--deviceTypeID", "1",
			"--productName", "TestProduct",
			"--productLabel", productLabel,
			"--partNumber", "1",
			"--commissioningCustomFlow", "0",
			"--enhancedSetupFlowOptions", "0",
			"--from", vendorAccount,
			"--node", FullNodeAddr,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, "add-model: %s", tx.RawLog)

		tx, err = utils.ExecuteTx(
			"tx", "model", "add-model-version",
			"--cdVersionNumber", "1",
			"--maxApplicableSoftwareVersion", "10",
			"--minApplicableSoftwareVersion", "1",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--softwareVersionString", "1",
			"--from", vendorAccount,
			"--node", FullNodeAddr,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, "add-model-version: %s", tx.RawLog)
	})

	// 4. Proxy now serves the new records. (model.sh lines 104-140)
	//    First read polls through the proxy's post-write sync window
	//    (bash sleeps 5; we poll up to 30s); subsequent queries reuse the
	//    now-synced state.
	mustRun(t, "Found_AfterAdd", func(t *testing.T) {
		t.Helper()
		out, qerr := queryUntilContains(LightClientProxyAddr, productLabel,
			"query", "model", "get-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
		)
		require.NoError(t, qerr)
		assertContains(t, out, fmt.Sprintf("%d", vid), "get-model.vid")
		assertContains(t, out, fmt.Sprintf("%d", pid), "get-model.pid")
		assertContains(t, out, productLabel, "get-model.productLabel")

		out, qerr = queryWithRetry(LightClientProxyAddr,
			"query", "model", "model-version",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--softwareVersion", fmt.Sprintf("%d", sv),
		)
		require.NoError(t, qerr)
		assertContains(t, out, fmt.Sprintf("%d", vid), "model-version.vid")
		assertContains(t, out, fmt.Sprintf("%d", pid), "model-version.pid")
		assertContains(t, out, fmt.Sprintf("%d", sv), "model-version.softwareVersion")
		assertContains(t, out, `"softwareVersionString": "1"`, "model-version.softwareVersionString")
		require.True(t,
			containsAnyLocal(out, `"softwareVersionValid": true`, `"softwareVersionValid":true`),
			"expected softwareVersionValid=true, got: %s", string(out))

		out, qerr = queryWithRetry(LightClientProxyAddr,
			"query", "model", "vendor-models",
			"--vid", fmt.Sprintf("%d", vid),
		)
		require.NoError(t, qerr)
		assertContains(t, out, fmt.Sprintf("%d", pid), "vendor-models.pid")

		out, qerr = queryWithRetry(LightClientProxyAddr,
			"query", "model", "all-model-versions",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
		)
		require.NoError(t, qerr)
		assertContains(t, out, fmt.Sprintf("%d", vid), "all-model-versions.vid")
		assertContains(t, out, fmt.Sprintf("%d", sv), "all-model-versions.softwareVersion")
	})

	// 5. Unrelated vid/pid/sv still returns Not Found through the proxy.
	mustRun(t, "NotFound_OtherKeys", func(t *testing.T) {
		t.Helper()
		otherV, otherP, otherSv := randomUint16(), randomUint16(), randomUint16()
		out, qerr := queryWithRetry(LightClientProxyAddr,
			"query", "model", "get-model",
			"--vid", fmt.Sprintf("%d", otherV),
			"--pid", fmt.Sprintf("%d", otherP),
		)
		require.NoError(t, qerr)
		assertContains(t, out, "Not Found", "get-model (other)")

		out, qerr = queryWithRetry(LightClientProxyAddr,
			"query", "model", "model-version",
			"--vid", fmt.Sprintf("%d", otherV),
			"--pid", fmt.Sprintf("%d", otherP),
			"--softwareVersion", fmt.Sprintf("%d", otherSv),
		)
		require.NoError(t, qerr)
		assertContains(t, out, "Not Found", "model-version (other)")

		out, qerr = queryWithRetry(LightClientProxyAddr,
			"query", "model", "vendor-models",
			"--vid", fmt.Sprintf("%d", otherV),
		)
		require.NoError(t, qerr)
		assertContains(t, out, "Not Found", "vendor-models (other)")

		out, qerr = queryWithRetry(LightClientProxyAddr,
			"query", "model", "all-model-versions",
			"--vid", fmt.Sprintf("%d", otherV),
			"--pid", fmt.Sprintf("%d", otherP),
		)
		require.NoError(t, qerr)
		assertContains(t, out, "Not Found", "all-model-versions (other)")
	})

	// 6. Write attempts through the proxy are rejected. Both add-model and
	//    add-model-version. (model.sh lines 188-199)
	mustRun(t, "Write_Rejected", func(t *testing.T) {
		t.Helper()
		out, err := executeCLIWithNode(LightClientProxyAddr,
			"tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--deviceTypeID", "1",
			"--productName", "TestProduct",
			"--productLabel", productLabel,
			"--partNumber", "1",
			"--commissioningCustomFlow", "0",
			"--enhancedSetupFlowOptions", "0",
			"--from", vendorAccount,
			"--yes", "-o", "json", "--keyring-backend", "test",
		)
		assertRejectionContains(t, out, err, writeRejection, "add-model")

		out, err = executeCLIWithNode(LightClientProxyAddr,
			"tx", "model", "add-model-version",
			"--cdVersionNumber", "1",
			"--maxApplicableSoftwareVersion", "10",
			"--minApplicableSoftwareVersion", "1",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--softwareVersionString", "1",
			"--from", vendorAccount,
			"--yes", "-o", "json", "--keyring-backend", "test",
		)
		assertRejectionContains(t, out, err, writeRejection, "add-model-version")
	})
}
