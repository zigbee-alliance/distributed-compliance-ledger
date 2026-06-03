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

// TestLightClientProxyVendorInfo is the Go translation of
// integration_tests/light_client_proxy/vendorinfo.sh.
func TestLightClientProxyVendorInfo(t *testing.T) {
	skipIfDisabled(t)

	// 1. Random VID — no vendorinfo record exists yet. Proxy returns Not Found.
	mustRun(t, "NotFound_BeforeAdd", func(t *testing.T) {
		t.Helper()
		vid := randomUint16()
		out, qerr := queryWithRetry(LightClientProxyAddr,
			"query", "vendorinfo", "vendor", "--vid", fmt.Sprintf("%d", vid),
		)
		require.NoError(t, qerr)
		assertContains(t, out, "Not Found", "vendor query")
	})

	// 2. Listing all vendors via the proxy is rejected.
	mustRun(t, "ListQuery_Rejected", func(t *testing.T) {
		t.Helper()
		out, qerr := queryWithRetry(LightClientProxyAddr,
			"query", "vendorinfo", "all-vendors",
		)
		assertRejectionContains(t, out, qerr, listQueryRejection, "all-vendors")
	})

	// 3. Propose Vendor account against the full node, then add the vendor
	//    info record. Mirrors bash lines 42-58.
	//
	//    Account name is suffixed with utils.RandString() so all five tests
	//    in this package can share one init_pool without colliding on the
	//    shared keyring (see integration_tests/run-all.sh).
	vid := randomUint16()
	vendorAccount := "vinfo_vendor_" + utils.RandString()
	const (
		companyLegalName = "XYZ IOT Devices Inc"
		vendorName       = "XYZ Devices"
	)

	mustRun(t, "AddVendorInfo", func(t *testing.T) {
		t.Helper()
		_ = proposeVendorAccount(t, vendorAccount, vid)

		tx, err := utils.ExecuteTx(
			"tx", "vendorinfo", "add-vendor",
			"--vid", fmt.Sprintf("%d", vid),
			"--companyLegalName", companyLegalName,
			"--vendorName", vendorName,
			"--from", vendorAccount,
			"--node", FullNodeAddr,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, "add-vendor: %s", tx.RawLog)
	})

	// 4. Now the proxy serves the new record. Poll through the proxy's
	//    post-write sync window (bash sleeps 5; we poll up to 30s).
	mustRun(t, "Found_AfterAdd", func(t *testing.T) {
		t.Helper()
		out, qerr := queryUntilContains(LightClientProxyAddr, companyLegalName,
			"query", "vendorinfo", "vendor", "--vid", fmt.Sprintf("%d", vid),
		)
		require.NoError(t, qerr)
		assertContains(t, out, fmt.Sprintf("%d", vid), "vendorID")
		assertContains(t, out, companyLegalName, "companyLegalName")
		assertContains(t, out, vendorName, "vendorName")
	})

	// 5. An unrelated VID still returns Not Found through the proxy.
	mustRun(t, "NotFound_OtherVID", func(t *testing.T) {
		t.Helper()
		otherVID := randomUint16()
		out, qerr := queryWithRetry(LightClientProxyAddr,
			"query", "vendorinfo", "vendor", "--vid", fmt.Sprintf("%d", otherVID),
		)
		require.NoError(t, qerr)
		assertContains(t, out, "Not Found", "vendor query")
	})

	// 6. Write attempt through the proxy is rejected.
	mustRun(t, "Write_Rejected", func(t *testing.T) {
		t.Helper()
		out, err := executeCLIWithNode(LightClientProxyAddr,
			"tx", "vendorinfo", "add-vendor",
			"--vid", fmt.Sprintf("%d", vid),
			"--companyLegalName", companyLegalName,
			"--vendorName", vendorName,
			"--from", vendorAccount,
			"--yes", "-o", "json", "--keyring-backend", "test",
		)
		assertRejectionContains(t, out, err, writeRejection, "add-vendor")
	})
}
