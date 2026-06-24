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

// TestLightClientProxyCompliance exercises the dcld compliance module
// against the light client proxy.
func TestLightClientProxyCompliance(t *testing.T) {
	skipIfDisabled(t)

	const (
		certType          = "zigbee"
		certificationDate = "2020-01-01T00:00:01Z"
		cdCertificateID   = "12345678910abcdefgh"
	)

	complianceSingleRecordCmds := []string{
		"compliance-info", "certified-model", "revoked-model", "provisional-model",
	}

	// queryCompliance wraps queryWithRetry with the four positional flags
	// every compliance single-record query needs.
	queryCompliance := func(t *testing.T, cmd string, vid, pid, sv int) []byte {
		t.Helper()
		out, qerr := queryWithRetry(LightClientProxyAddr,
			ComplianceByKeys(cmd, vid, pid, sv, certType)...)
		require.NoError(t, qerr, "%s query", cmd)

		return out
	}

	// 1. Random vid/pid/sv → every single-record compliance query returns
	//    Not Found via the proxy.
	mustRun(t, "NotFound_BeforeAdd", func(t *testing.T) {
		t.Helper()
		vid, pid, sv := randomUint16(), randomUint16(), randomUint16()
		for _, cmd := range complianceSingleRecordCmds {
			assertNotFoundOnProxy(t, cmd, ComplianceByKeys(cmd, vid, pid, sv, certType)...)
		}
	})

	// 2. List queries rejected by the proxy.
	mustRun(t, "ListQueries_Rejected", func(t *testing.T) {
		t.Helper()
		assertListQueriesRejected(t, "compliance",
			"all-compliance-info", "all-certified-models",
			"all-revoked-models", "all-provisional-models")
	})

	// 3. Seed: vendor account → certification center → model → certify.
	//    Account names are suffixed with utils.RandString() — the five
	//    light_client_proxy tests share one init_pool so the keyring is
	//    shared (see run-all.sh).
	vid, pid, sv := randomUint16(), randomUint16(), randomUint16()
	svs := strconv.Itoa(randomUint16())
	vendorAccount := "comp_vendor_" + utils.RandString()
	certCenter := "comp_certctr_" + utils.RandString()

	certifyArgs := CertifyModelArgs{
		VID:                   vid,
		PID:                   pid,
		SoftwareVersion:       sv,
		SoftwareVersionString: svs,
		CertificationType:     certType,
		CertificationDate:     certificationDate,
		CDCertificateID:       cdCertificateID,
	}

	mustRun(t, "Seed_Vendor_CertCenter_Model", func(t *testing.T) {
		t.Helper()
		_ = proposeVendorAccount(t, vendorAccount, vid)
		_ = proposeAndApproveAccount(t, certCenter, "CertificationCenter")
		addModelAndVersion(t, vid, pid, sv, svs, vendorAccount)

		tx, err := certifyArgs.Send(certCenter)
		requireTxOK(t, tx, err, "certify-model")
	})

	// 4. Proxy now serves the new record. certified=true, revoked=false,
	//    provisional=false. compliance-info reports softwareVersionCertificationStatus=2.
	//    First do a poll-until-contains read to absorb the proxy's post-write
	//    sync window (up to 30s); subsequent queries in this block reuse the
	//    now-synced state.
	mustRun(t, "Found_AfterCertify", func(t *testing.T) {
		t.Helper()
		_, qerr := queryUntilContains(LightClientProxyAddr, strconv.Itoa(vid),
			ComplianceByKeys("compliance-info", vid, pid, sv, certType)...)
		require.NoError(t, qerr)

		out := queryCompliance(t, "compliance-info", vid, pid, sv)
		assertContains(t, out, strconv.Itoa(vid), "compliance-info.vid")
		assertContains(t, out, strconv.Itoa(pid), "compliance-info.pid")
		// JSON spacing varies between dcld versions; both forms appear in
		// the wild. Accept either.
		require.True(t,
			containsAnyLocal(out, `"softwareVersionCertificationStatus": 2`,
				`"softwareVersionCertificationStatus":2`),
			"expected softwareVersionCertificationStatus=2, got: %s", string(out),
		)
		assertContains(t, out, certificationDate, "compliance-info.date")
		assertContains(t, out, certType, "compliance-info.certificationType")

		out = queryCompliance(t, "certified-model", vid, pid, sv)
		require.True(t,
			containsAnyLocal(out, `"value": true`, `"value":true`),
			"expected certified value=true, got: %s", string(out))

		out = queryCompliance(t, "revoked-model", vid, pid, sv)
		require.True(t,
			containsAnyLocal(out, `"value": false`, `"value":false`),
			"expected revoked value=false, got: %s", string(out))

		out = queryCompliance(t, "provisional-model", vid, pid, sv)
		require.True(t,
			containsAnyLocal(out, `"value": false`, `"value":false`),
			"expected provisional value=false, got: %s", string(out))
	})

	// 5. An unrelated vid/pid/sv still returns Not Found.
	mustRun(t, "NotFound_OtherKeys", func(t *testing.T) {
		t.Helper()
		v, p, s := randomUint16(), randomUint16(), randomUint16()
		for _, cmd := range complianceSingleRecordCmds {
			assertNotFoundOnProxy(t, cmd, ComplianceByKeys(cmd, v, p, s, certType)...)
		}
	})

	// 6. Write through the proxy is rejected.
	mustRun(t, "Write_Rejected", func(t *testing.T) {
		t.Helper()
		args := append(certifyArgs.Build(), "--from", certCenter)
		assertWriteRejected(t, "certify-model", args...)
	})
}
