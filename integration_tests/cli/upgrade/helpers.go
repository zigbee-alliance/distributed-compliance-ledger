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
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// ProposeUpgrade proposes a software upgrade.
func ProposeUpgrade(name, height, from string, extra ...string) (*utils.TxResult, error) {
	args := []string{
		"tx", "dclupgrade", "propose-upgrade",
		"--name", name,
		"--upgrade-height", height,
		"--from", from,
	}
	args = append(args, extra...)

	return utils.ExecuteTx(args...)
}

// ApproveUpgrade approves a proposed software upgrade.
func ApproveUpgrade(name, from string) (*utils.TxResult, error) {
	return utils.ExecuteTx("tx", "dclupgrade", "approve-upgrade",
		"--name", name,
		"--from", from,
	)
}

// RejectUpgrade rejects a proposed software upgrade.
func RejectUpgrade(name, from string) (*utils.TxResult, error) {
	return utils.ExecuteTx("tx", "dclupgrade", "reject-upgrade",
		"--name", name,
		"--from", from,
	)
}

// QueryProposedUpgrade queries a proposed upgrade plan by name.
func QueryProposedUpgrade(name string) ([]byte, error) {
	return utils.ExecuteCLI("query", "dclupgrade", "proposed-upgrade",
		"--name", name,
		"-o", "json",
	)
}

// QueryApprovedUpgrade queries an approved upgrade plan by name.
func QueryApprovedUpgrade(name string) ([]byte, error) {
	return utils.ExecuteCLI("query", "dclupgrade", "approved-upgrade",
		"--name", name,
		"-o", "json",
	)
}

// QueryRejectedUpgrade queries a rejected upgrade plan by name.
func QueryRejectedUpgrade(name string) ([]byte, error) {
	return utils.ExecuteCLI("query", "dclupgrade", "rejected-upgrade",
		"--name", name,
		"-o", "json",
	)
}

// QueryUpgradePlan queries the currently scheduled upgrade plan.
func QueryUpgradePlan() ([]byte, error) {
	return utils.ExecuteCLI("query", "upgrade", "plan", "-o", "json")
}
