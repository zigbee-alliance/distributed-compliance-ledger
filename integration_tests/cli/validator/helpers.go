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

package validator

import (
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// AddNode adds a validator node.
func AddNode(pubkey, moniker, from string, extra ...string) (*utils.TxResult, error) {
	args := []string{
		"tx", "validator", "add-node",
		"--pubkey", pubkey,
		"--moniker", moniker,
		"--from", from,
	}
	args = append(args, extra...)

	return utils.ExecuteTx(args...)
}

// DisableNode disables a validator node (node admin self-disable).
func DisableNode(from string) (*utils.TxResult, error) {
	return utils.ExecuteTx("tx", "validator", "disable-node", "--from", from)
}

// EnableNode enables a previously disabled validator node.
func EnableNode(from string) (*utils.TxResult, error) {
	return utils.ExecuteTx("tx", "validator", "enable-node", "--from", from)
}

// ProposeDisableNode proposes disabling a validator node.
func ProposeDisableNode(address, from string) (*utils.TxResult, error) {
	return utils.ExecuteTx("tx", "validator", "propose-disable-node",
		"--address", address,
		"--from", from,
	)
}

// ApproveDisableNode approves disabling a validator node.
func ApproveDisableNode(address, from string) (*utils.TxResult, error) {
	return utils.ExecuteTx("tx", "validator", "approve-disable-node",
		"--address", address,
		"--from", from,
	)
}

// RejectDisableNode rejects a proposal to disable a validator node.
func RejectDisableNode(address, from string) (*utils.TxResult, error) {
	return utils.ExecuteTx("tx", "validator", "reject-disable-node",
		"--address", address,
		"--from", from,
	)
}

// QueryNode queries a validator node by owner address.
func QueryNode(address string) ([]byte, error) {
	return utils.ExecuteCLI("query", "validator", "node",
		"--address", address,
		"-o", "json",
	)
}

// QueryAllNodes queries all validator nodes.
func QueryAllNodes() ([]byte, error) {
	return utils.ExecuteCLI("query", "validator", "all-nodes", "-o", "json")
}

// QueryLastPower queries the last power of a validator by owner address.
func QueryLastPower(address string) ([]byte, error) {
	return utils.ExecuteCLI("query", "validator", "last-power",
		"--address", address,
		"-o", "json",
	)
}

// QueryDisabledNode queries a disabled validator node by address.
func QueryDisabledNode(address string) ([]byte, error) {
	return utils.ExecuteCLI("query", "validator", "disabled-node",
		"--address", address,
		"-o", "json",
	)
}

// QueryProposedDisableNode queries a proposed-to-disable validator node.
func QueryProposedDisableNode(address string) ([]byte, error) {
	return utils.ExecuteCLI("query", "validator", "proposed-disable-node",
		"--address", address,
		"-o", "json",
	)
}

// QueryAllProposedDisableNodes queries all proposed-to-disable validator nodes.
func QueryAllProposedDisableNodes() ([]byte, error) {
	return utils.ExecuteCLI("query", "validator", "all-proposed-disable-nodes", "-o", "json")
}

// QueryRejectedDisableNode queries a rejected disable-node proposal by address.
func QueryRejectedDisableNode(address string) ([]byte, error) {
	return utils.ExecuteCLI("query", "validator", "rejected-disable-node",
		"--address", address,
		"-o", "json",
	)
}

// QueryAllRejectedDisableNodes queries all rejected disable-node proposals.
func QueryAllRejectedDisableNodes() ([]byte, error) {
	return utils.ExecuteCLI("query", "validator", "all-rejected-disable-nodes", "-o", "json")
}
