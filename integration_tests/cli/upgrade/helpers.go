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
	"encoding/json"
	"fmt"

	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
	upgradetypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
)

// ProposeUpgradeOpts holds optional flags for propose-upgrade.
// UpgradeInfo emits --upgrade-info; Info emits --info (proposer note).
type ProposeUpgradeOpts struct {
	UpgradeInfo string
	Info        string
	Extra       []string
}

func (o ProposeUpgradeOpts) args() []string {
	var args []string
	if o.UpgradeInfo != "" {
		args = append(args, "--upgrade-info", o.UpgradeInfo)
	}
	if o.Info != "" {
		args = append(args, "--info", o.Info)
	}

	return append(args, o.Extra...)
}

// UpgradeActionOpts holds optional flags for approve-upgrade / reject-upgrade.
type UpgradeActionOpts struct {
	Info  string
	Extra []string
}

func (o UpgradeActionOpts) args() []string {
	var args []string
	if o.Info != "" {
		args = append(args, "--info", o.Info)
	}

	return append(args, o.Extra...)
}

// ProposeUpgrade proposes a software upgrade.
func ProposeUpgrade(name, height, from string, opts ...ProposeUpgradeOpts) (*utils.TxResult, error) {
	args := []string{
		"tx", "dclupgrade", "propose-upgrade",
		"--name", name,
		"--upgrade-height", height,
		"--from", from,
	}
	for _, o := range opts {
		args = append(args, o.args()...)
	}

	return utils.ExecuteTx(args...)
}

// ApproveUpgrade approves a proposed software upgrade.
func ApproveUpgrade(name, from string, opts ...UpgradeActionOpts) (*utils.TxResult, error) {
	args := []string{
		"tx", "dclupgrade", "approve-upgrade",
		"--name", name,
		"--from", from,
	}
	for _, o := range opts {
		args = append(args, o.args()...)
	}

	return utils.ExecuteTx(args...)
}

// RejectUpgrade rejects a proposed software upgrade.
func RejectUpgrade(name, from string, opts ...UpgradeActionOpts) (*utils.TxResult, error) {
	args := []string{
		"tx", "dclupgrade", "reject-upgrade",
		"--name", name,
		"--from", from,
	}
	for _, o := range opts {
		args = append(args, o.args()...)
	}

	return utils.ExecuteTx(args...)
}

// getSingle runs a single-item dcld query and unmarshals into v. Returns
// (false, nil) when the CLI emitted "Not Found".
func getSingle(v interface{}, args ...string) (found bool, err error) {
	out, err := utils.ExecuteCLI(args...)
	if err != nil {
		return false, err
	}
	if utils.IsNotFound(out) {
		return false, nil
	}
	if err := json.Unmarshal(out, v); err != nil {
		return false, fmt.Errorf("parse %T: %w, output: %s", v, err, string(out))
	}

	return true, nil
}

// GetProposedUpgrade queries a proposed upgrade plan by name. Returns nil when
// no proposal exists.
func GetProposedUpgrade(name string) (*upgradetypes.ProposedUpgrade, error) {
	var res upgradetypes.ProposedUpgrade
	found, err := getSingle(&res,
		"query", "dclupgrade", "proposed-upgrade",
		"--name", name,
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetApprovedUpgrade queries an approved upgrade plan by name. Returns nil
// when no approved record exists.
func GetApprovedUpgrade(name string) (*upgradetypes.ApprovedUpgrade, error) {
	var res upgradetypes.ApprovedUpgrade
	found, err := getSingle(&res,
		"query", "dclupgrade", "approved-upgrade",
		"--name", name,
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetRejectedUpgrade queries a rejected upgrade plan by name. Returns nil when
// no rejected record exists.
func GetRejectedUpgrade(name string) (*upgradetypes.RejectedUpgrade, error) {
	var res upgradetypes.RejectedUpgrade
	found, err := getSingle(&res,
		"query", "dclupgrade", "rejected-upgrade",
		"--name", name,
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetUpgradePlan queries the currently scheduled upgrade plan from the Cosmos
// SDK upgrade module. The CLI prints just the Plan body (not the
// QueryCurrentPlanResponse wrapper). Returns the underlying CLI error when no
// plan is scheduled — the SDK exits with "no upgrade scheduled".
func GetUpgradePlan() (*upgradetypes.Plan, error) {
	out, err := utils.ExecuteCLI("query", "upgrade", "plan", "-o", "json")
	if err != nil {
		return nil, err
	}
	var res upgradetypes.Plan
	if err := json.Unmarshal(out, &res); err != nil {
		return nil, fmt.Errorf("parse Plan: %w, output: %s", err, string(out))
	}

	return &res, nil
}
