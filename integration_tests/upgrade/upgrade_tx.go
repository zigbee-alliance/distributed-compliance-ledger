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

package upgrade

import (
	"fmt"

	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// UpgradeInfoForVersion returns the standard `--upgrade-info` JSON payload
// pointing at the released linux/amd64 dcld binary for `version`. `checksum`
// is the sha256 verifying the binary. Pass an empty checksum to omit it.
func UpgradeInfoForVersion(version, checksum string) string {
	url := fmt.Sprintf(BinaryURLTemplate, version)
	if checksum != "" {
		url += "?checksum=" + checksum
	}

	return fmt.Sprintf(`{"binaries":{"linux/amd64":"%s"}}`, url)
}

// ProposeUpgrade submits a `dclupgrade propose-upgrade` tx using the binary at
// binPath. The active dcld may be running an older release than the host
// build, hence the explicit binary path. Result is the confirmed TxResult.
func ProposeUpgrade(binPath, planName string, planHeight int64, upgradeInfo, from string) (*utils.TxResult, error) {
	return ExecuteTxWithBin(binPath,
		"tx", "dclupgrade", "propose-upgrade",
		"--name", planName,
		"--upgrade-height", fmt.Sprintf("%d", planHeight),
		"--upgrade-info", upgradeInfo,
		"--from", from,
	)
}

// ApproveUpgrade submits a `dclupgrade approve-upgrade` tx using the binary at
// binPath.
func ApproveUpgrade(binPath, planName, from string) (*utils.TxResult, error) {
	return ExecuteTxWithBin(binPath,
		"tx", "dclupgrade", "approve-upgrade",
		"--name", planName,
		"--from", from,
	)
}

// RejectUpgrade submits a `dclupgrade reject-upgrade` tx using the binary at
// binPath.
func RejectUpgrade(binPath, planName, from string) (*utils.TxResult, error) {
	return ExecuteTxWithBin(binPath,
		"tx", "dclupgrade", "reject-upgrade",
		"--name", planName,
		"--from", from,
	)
}

// QueryProposedUpgrade queries a proposed upgrade plan via the binary at binPath.
func QueryProposedUpgrade(binPath, planName string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath,
		"query", "dclupgrade", "proposed-upgrade",
		"--name", planName,
		"-o", "json",
	)
}

// QueryApprovedUpgrade queries an approved upgrade plan via the binary at binPath.
func QueryApprovedUpgrade(binPath, planName string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath,
		"query", "dclupgrade", "approved-upgrade",
		"--name", planName,
		"-o", "json",
	)
}

// QueryUpgradePlan queries the currently scheduled cosmos-sdk upgrade plan.
func QueryUpgradePlan(binPath string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "upgrade", "plan", "-o", "json")
}

// QueryAppliedPlan queries an applied upgrade plan via the binary at binPath.
func QueryAppliedPlan(binPath, planName string) ([]byte, error) {
	return ExecuteCLIWithBin(binPath, "query", "upgrade", "applied", planName, "-o", "json")
}
