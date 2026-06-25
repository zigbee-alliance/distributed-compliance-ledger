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
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// TestLightClientProxyModel exercises the dcld model module against the
// light client proxy.
func TestLightClientProxyModel(t *testing.T) {
	skipIfDisabled(t)

	// 1. Random vid/pid/sv → every single-record model query returns
	//    Not Found via the proxy.
	mustRun(t, "NotFound_BeforeAdd", func(t *testing.T) {
		t.Helper()
		vid, pid, sv := randomUint16(), randomUint16(), randomUint16()
		assertNotFoundOnProxy(t, "get-model", GetModel(vid, pid)...)
		assertNotFoundOnProxy(t, "vendor-models", VendorModels(vid)...)
		assertNotFoundOnProxy(t, "model-version", ModelVersion(vid, pid, sv)...)
		assertNotFoundOnProxy(t, "all-model-versions", AllModelVersions(vid, pid)...)
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
	//    Same isolation contract every test in this package uses.
	vid, pid, sv := randomUint16(), randomUint16(), randomUint16()
	vendorAccount := "model_vendor_" + utils.RandString()
	const productLabel = "Device #1"

	addModelArgs := AddModelArgs{VID: vid, PID: pid, ProductLabel: productLabel}
	addModelVersionArgs := AddModelVersionArgs{
		VID: vid, PID: pid, SoftwareVersion: sv, SoftwareVersionString: "1",
	}

	mustRun(t, "Seed_VendorAndModel", func(t *testing.T) {
		t.Helper()
		_ = proposeVendorAccount(t, vendorAccount, vid)

		tx, err := addModelArgs.Send(vendorAccount)
		requireTxOK(t, tx, err, "add-model")

		tx, err = addModelVersionArgs.Send(vendorAccount)
		requireTxOK(t, tx, err, "add-model-version")
	})

	// 4. Proxy now serves the new records.
	//    The proxy syncs headers monotonically, so we warm up by polling the
	//    *latest* write (add-model-version) until visible — once that's
	//    visible, every earlier write (add-model) is guaranteed visible too.
	//    Polling on the first write (get-model) here would race
	//    add-model-version which lands in a later block. Poll up to 30s.
	mustRun(t, "Found_AfterAdd", func(t *testing.T) {
		t.Helper()
		out, qerr := queryUntilContains(LightClientProxyAddr, strconv.Itoa(sv),
			ModelVersion(vid, pid, sv)...)
		require.NoError(t, qerr)
		assertContains(t, out, strconv.Itoa(vid), "model-version.vid")
		assertContains(t, out, strconv.Itoa(pid), "model-version.pid")
		assertContains(t, out, strconv.Itoa(sv), "model-version.softwareVersion")
		// Accept both legacy spaced and master-binary compact JSON forms.
		// containsAnyLocal handles the same legacy-vs-compact split we see
		// across the upgrade suite's mixed-binary queries.
		require.True(t,
			containsAnyLocal(out, `"softwareVersionString": "1"`, `"softwareVersionString":"1"`),
			"expected softwareVersionString=%q, got: %s", "1", string(out))
		require.True(t,
			containsAnyLocal(out, `"softwareVersionValid": true`, `"softwareVersionValid":true`),
			"expected softwareVersionValid=true, got: %s", string(out))
		require.True(t,
			containsAnyLocal(out, `"cdVersionNumber": 1`, `"cdVersionNumber":1`),
			"expected cdVersionNumber=1, got: %s", string(out))
		require.True(t,
			containsAnyLocal(out, `"minApplicableSoftwareVersion": 1`, `"minApplicableSoftwareVersion":1`),
			"expected minApplicableSoftwareVersion=1, got: %s", string(out))
		require.True(t,
			containsAnyLocal(out, `"maxApplicableSoftwareVersion": 10`, `"maxApplicableSoftwareVersion":10`),
			"expected maxApplicableSoftwareVersion=10, got: %s", string(out))

		out, qerr = queryWithRetry(LightClientProxyAddr, GetModel(vid, pid)...)
		require.NoError(t, qerr)
		assertContains(t, out, strconv.Itoa(vid), "get-model.vid")
		assertContains(t, out, strconv.Itoa(pid), "get-model.pid")
		assertContains(t, out, productLabel, "get-model.productLabel")

		out, qerr = queryWithRetry(LightClientProxyAddr, VendorModels(vid)...)
		require.NoError(t, qerr)
		assertContains(t, out, strconv.Itoa(pid), "vendor-models.pid")

		out, qerr = queryWithRetry(LightClientProxyAddr, AllModelVersions(vid, pid)...)
		require.NoError(t, qerr)
		assertContains(t, out, strconv.Itoa(vid), "all-model-versions.vid")
		assertContains(t, out, strconv.Itoa(sv), "all-model-versions.softwareVersion")
	})

	// 5. Unrelated vid/pid/sv still returns Not Found through the proxy.
	mustRun(t, "NotFound_OtherKeys", func(t *testing.T) {
		t.Helper()
		otherV, otherP, otherSv := randomUint16(), randomUint16(), randomUint16()
		assertNotFoundOnProxy(t, "get-model (other)", GetModel(otherV, otherP)...)
		assertNotFoundOnProxy(t, "model-version (other)", ModelVersion(otherV, otherP, otherSv)...)
		assertNotFoundOnProxy(t, "vendor-models (other)", VendorModels(otherV)...)
		assertNotFoundOnProxy(t, "all-model-versions (other)", AllModelVersions(otherV, otherP)...)
	})

	// 6. Write attempts through the proxy are rejected. Both add-model and
	//    add-model-version.
	mustRun(t, "Write_Rejected", func(t *testing.T) {
		t.Helper()
		args := append(addModelArgs.Build(), "--from", vendorAccount)
		assertWriteRejected(t, "add-model", args...)

		args = append(addModelVersionArgs.Build(), "--from", vendorAccount)
		assertWriteRejected(t, "add-model-version", args...)
	})
}
