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
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
	validatortypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
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

// GetNode queries a validator node by owner address. Returns nil when the
// validator does not exist.
func GetNode(address string) (*validatortypes.Validator, error) {
	var res validatortypes.Validator
	found, err := cliputils.GetSingle(&res,
		"query", "validator", "node",
		"--address", address,
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetAllNodes queries all validator nodes.
func GetAllNodes() ([]validatortypes.Validator, error) {
	var res validatortypes.QueryAllValidatorResponse
	if err := cliputils.GetList(&res, "query", "validator", "all-nodes", "-o", "json"); err != nil {
		return nil, err
	}

	return res.Validator, nil
}

// GetLastPower queries the last power record of a validator by owner address.
// Returns nil when no record exists for the address.
func GetLastPower(address string) (*validatortypes.LastValidatorPower, error) {
	var res validatortypes.LastValidatorPower
	found, err := cliputils.GetSingle(&res,
		"query", "validator", "last-power",
		"--address", address,
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetDisabledNode queries a disabled validator node by address. Returns nil
// when no disabled record exists.
func GetDisabledNode(address string) (*validatortypes.DisabledValidator, error) {
	var res validatortypes.DisabledValidator
	found, err := cliputils.GetSingle(&res,
		"query", "validator", "disabled-node",
		"--address", address,
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetProposedDisableNode queries a proposed-to-disable validator node by
// address. Returns nil when no proposal exists.
func GetProposedDisableNode(address string) (*validatortypes.ProposedDisableValidator, error) {
	var res validatortypes.ProposedDisableValidator
	found, err := cliputils.GetSingle(&res,
		"query", "validator", "proposed-disable-node",
		"--address", address,
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetAllProposedDisableNodes queries all proposed-to-disable validator nodes.
func GetAllProposedDisableNodes() ([]validatortypes.ProposedDisableValidator, error) {
	var res validatortypes.QueryAllProposedDisableValidatorResponse
	if err := cliputils.GetList(&res, "query", "validator", "all-proposed-disable-nodes", "-o", "json"); err != nil {
		return nil, err
	}

	return res.ProposedDisableValidator, nil
}

// GetRejectedDisableNode queries a rejected disable-node proposal by address.
// Returns nil when no rejection record exists.
func GetRejectedDisableNode(address string) (*validatortypes.RejectedDisableValidator, error) {
	var res validatortypes.RejectedDisableValidator
	found, err := cliputils.GetSingle(&res,
		"query", "validator", "rejected-disable-node",
		"--address", address,
		"-o", "json",
	)
	if err != nil || !found {
		return nil, err
	}

	return &res, nil
}

// GetAllRejectedDisableNodes queries all rejected disable-node proposals.
func GetAllRejectedDisableNodes() ([]validatortypes.RejectedDisableValidator, error) {
	var res validatortypes.QueryAllRejectedDisableValidatorResponse
	if err := cliputils.GetList(&res, "query", "validator", "all-rejected-disable-nodes", "-o", "json"); err != nil {
		return nil, err
	}

	return res.RejectedValidator, nil
}

// containsValidatorByOwner reports whether list has a Validator with the given owner address.
func containsValidatorByOwner(list []validatortypes.Validator, owner string) bool {
	for i := range list {
		if list[i].Owner == owner {
			return true
		}
	}

	return false
}

// containsProposedByAddress reports whether list has a proposed-disable entry for address.
func containsProposedByAddress(list []validatortypes.ProposedDisableValidator, address string) bool {
	for i := range list {
		if list[i].Address == address {
			return true
		}
	}

	return false
}

// containsRejectedByAddress reports whether list has a rejected-disable entry for address.
func containsRejectedByAddress(list []validatortypes.RejectedDisableValidator, address string) bool {
	for i := range list {
		if list[i].Address == address {
			return true
		}
	}

	return false
}
