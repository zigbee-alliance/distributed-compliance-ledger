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

// TestLightClientProxyCompliance is the Go translation of
// integration_tests/light_client_proxy/compliance.sh.
//
//nolint:funlen
func TestLightClientProxyCompliance(t *testing.T) {
	skipIfDisabled(t)

	const (
		certType          = "zigbee"
		certificationDate = "2020-01-01T00:00:01Z"
		cdCertificateID   = "12345678910abcdefgh"
	)

	// Helper closures keep the per-step blocks readable. complianceQuery
	// wraps queryWithRetry with the four positional flags every compliance
	// single-record query needs.
	complianceQuery := func(t *testing.T, cmd string, vid, pid, sv int) []byte {
		t.Helper()
		out, qerr := queryWithRetry(LightClientProxyAddr,
			"query", "compliance", cmd,
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--certificationType", certType,
		)
		require.NoError(t, qerr, "%s query", cmd)

		return out
	}

	// 1. Random vid/pid/sv → every single-record compliance query returns
	//    Not Found via the proxy.
	mustRun(t, "NotFound_BeforeAdd", func(t *testing.T) {
		t.Helper()
		vid, pid, sv := randomUint16(), randomUint16(), randomUint16()
		for _, cmd := range []string{
			"compliance-info", "certified-model", "revoked-model", "provisional-model",
		} {
			out := complianceQuery(t, cmd, vid, pid, sv)
			assertContains(t, out, "Not Found", cmd)
		}
	})

	// 2. List queries rejected by the proxy.
	mustRun(t, "ListQueries_Rejected", func(t *testing.T) {
		t.Helper()
		for _, q := range []string{
			"all-compliance-info", "all-certified-models",
			"all-revoked-models", "all-provisional-models",
		} {
			out, qerr := queryWithRetry(LightClientProxyAddr, "query", "compliance", q)
			assertRejectionContains(t, out, qerr, listQueryRejection, q)
		}
	})

	// 3. Seed: vendor account → certification center → model → certify.
	//    Account names are suffixed with utils.RandString() — the five
	//    light_client_proxy tests share one init_pool so the keyring is
	//    shared (see integration_tests/run-all.sh).
	vid, pid, sv := randomUint16(), randomUint16(), randomUint16()
	svs := fmt.Sprintf("%d", randomUint16())
	vendorAccount := "comp_vendor_" + utils.RandString()
	certCenter := "comp_certctr_" + utils.RandString()

	mustRun(t, "Seed_Vendor_CertCenter_Model", func(t *testing.T) {
		t.Helper()
		_ = proposeVendorAccount(t, vendorAccount, vid)
		_ = proposeAndApproveAccount(t, certCenter, "CertificationCenter")
		addModelAndVersion(t, vid, pid, sv, svs, vendorAccount)

		// Certify the model with zigbee certification.
		tx, err := utils.ExecuteTx(
			"tx", "compliance", "certify-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--softwareVersionString", svs,
			"--certificationType", certType,
			"--certificationDate", certificationDate,
			"--cdCertificateId", cdCertificateID,
			"--cdVersionNumber", "1",
			"--from", certCenter,
			"--node", FullNodeAddr,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, "certify-model: %s", tx.RawLog)
	})

	// 4. Proxy now serves the new record. certified=true, revoked=false,
	//    provisional=false. compliance-info reports softwareVersionCertificationStatus=2.
	//    First do a poll-until-contains read to absorb the proxy's post-write
	//    sync window (bash sleeps 5; we poll up to 30s); subsequent queries
	//    in this block reuse the now-synced state.
	mustRun(t, "Found_AfterCertify", func(t *testing.T) {
		t.Helper()
		_, qerr := queryUntilContains(LightClientProxyAddr, fmt.Sprintf("%d", vid),
			"query", "compliance", "compliance-info",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--certificationType", certType,
		)
		require.NoError(t, qerr)

		out := complianceQuery(t, "compliance-info", vid, pid, sv)
		assertContains(t, out, fmt.Sprintf("%d", vid), "compliance-info.vid")
		assertContains(t, out, fmt.Sprintf("%d", pid), "compliance-info.pid")
		// JSON spacing varies between dcld versions; both forms appear in
		// the wild. Accept either.
		require.True(t,
			containsAnyLocal(out, `"softwareVersionCertificationStatus": 2`,
				`"softwareVersionCertificationStatus":2`),
			"expected softwareVersionCertificationStatus=2, got: %s", string(out),
		)
		assertContains(t, out, certificationDate, "compliance-info.date")
		assertContains(t, out, certType, "compliance-info.certificationType")

		out = complianceQuery(t, "certified-model", vid, pid, sv)
		require.True(t,
			containsAnyLocal(out, `"value": true`, `"value":true`),
			"expected certified value=true, got: %s", string(out))

		out = complianceQuery(t, "revoked-model", vid, pid, sv)
		require.True(t,
			containsAnyLocal(out, `"value": false`, `"value":false`),
			"expected revoked value=false, got: %s", string(out))

		out = complianceQuery(t, "provisional-model", vid, pid, sv)
		require.True(t,
			containsAnyLocal(out, `"value": false`, `"value":false`),
			"expected provisional value=false, got: %s", string(out))
	})

	// 5. An unrelated vid/pid/sv still returns Not Found.
	mustRun(t, "NotFound_OtherKeys", func(t *testing.T) {
		t.Helper()
		v, p, s := randomUint16(), randomUint16(), randomUint16()
		for _, cmd := range []string{
			"compliance-info", "certified-model", "revoked-model", "provisional-model",
		} {
			out := complianceQuery(t, cmd, v, p, s)
			assertContains(t, out, "Not Found", cmd)
		}
	})

	// 6. Write through the proxy is rejected.
	mustRun(t, "Write_Rejected", func(t *testing.T) {
		t.Helper()
		out, err := executeCLIWithNode(LightClientProxyAddr,
			"tx", "compliance", "certify-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--softwareVersionString", svs,
			"--certificationType", certType,
			"--certificationDate", certificationDate,
			"--cdCertificateId", cdCertificateID,
			"--cdVersionNumber", "1",
			"--from", certCenter,
			"--yes", "-o", "json", "--keyring-backend", "test",
		)
		assertRejectionContains(t, out, err, writeRejection, "certify-model")
	})
}
