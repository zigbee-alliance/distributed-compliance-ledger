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
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// AddValidatorNode is the Go translation of script 01's `add_validator_node`
// function. It spins up a fresh validator-demo container at 192.167.10.6,
// configures dcld, creates a NodeAdmin account, gets it approved by trustees
// jack/alice/bob on the main chain, submits `tx validator add-node`, seeds
// cosmovisor, starts the node helper, and records the resulting validator
// owner address + account name on `state` for downstream disable/enable flows.
//
//nolint:funlen
func AddValidatorNode(t *testing.T, state *UpgradeTestState, dcldHost string) {
	t.Helper()

	DockerCleanup(ValidatorDemoContainerName)

	accountName := RandomString()

	// 1. Start container on the localnet network.
	portMap := fmt.Sprintf("%d-%d:26656-26657",
		ValidatorDemoP2PPort, ValidatorDemoClientPort)
	_, err := DockerRun(
		"-d",
		"--name", ValidatorDemoContainerName,
		"--ip", ValidatorDemoIP,
		"-p", portMap,
		"--network", DockerNetwork,
		"-i", "dcledger",
	)
	require.NoError(t, err, "docker run validator-demo")

	// 2. Drop the dcld binary into the container.
	_, err = DockerCp(dcldHost, ValidatorDemoContainerName+":"+DCLUserHome+"/dcld")
	require.NoError(t, err, "docker cp dcld into validator-demo")

	// 3. Configure dcld client inside the container.
	for _, args := range [][]string{
		{"./dcld", "config", "chain-id", ChainID},
		{"./dcld", "config", "output", "json"},
		{"./dcld", "config", "node", Node0Conn},
		{"./dcld", "config", "keyring-backend", "test"},
		{"./dcld", "config", "broadcast-mode", "block"},
	} {
		_, err = DockerExec(ValidatorDemoContainerName, args...)
		require.NoError(t, err, "dcld config %v", args)
	}

	// 4. Initialize the new node.
	_, err = DockerExec(ValidatorDemoContainerName,
		"./dcld", "init", "node-demo", "--chain-id", ChainID)
	require.NoError(t, err, "dcld init")

	// 5. Copy genesis.json from node0.
	_, err = DockerCp(
		LocalnetDir+"/node0/config/genesis.json",
		ValidatorDemoContainerName+":"+DCLDir+"/config",
	)
	require.NoError(t, err, "docker cp genesis.json")

	// 6. Inject persistent_peers + listen address into config.toml.
	peers, err := readPersistentPeersFromHost(LocalnetDir + "/node0/config/config.toml")
	require.NoError(t, err)

	_, err = DockerExecShell(ValidatorDemoContainerName, fmt.Sprintf(
		`sed -i 's|persistent_peers = ""|persistent_peers = "%s"|g' %s/config/config.toml`,
		peers, DCLDir,
	))
	require.NoError(t, err, "sed persistent_peers")

	_, err = DockerExecShell(ValidatorDemoContainerName, fmt.Sprintf(
		`sed -i 's|laddr = "tcp://127.0.0.1:26657"|laddr = "tcp://0.0.0.0:26657"|g' %s/config/config.toml`,
		DCLDir,
	))
	require.NoError(t, err, "sed laddr")

	// 7. Generate a fresh key inside the container's keyring.
	_, err = DockerExec(ValidatorDemoContainerName,
		"./dcld", "keys", "add", accountName, "--keyring-backend", "test")
	require.NoError(t, err, "keys add")

	addrOut, err := DockerExec(ValidatorDemoContainerName,
		"./dcld", "keys", "show", accountName, "-a", "--keyring-backend", "test")
	require.NoError(t, err)
	address := strings.TrimSpace(string(addrOut))
	require.NotEmpty(t, address)

	pubOut, err := DockerExec(ValidatorDemoContainerName,
		"./dcld", "keys", "show", accountName, "-p", "--keyring-backend", "test")
	require.NoError(t, err)
	pubkey := strings.TrimSpace(string(pubOut))
	require.NotEmpty(t, pubkey)

	// 8. Propose + 2 approvals for the NodeAdmin account on the main chain.
	tx, err := ProposeAddAccount(dcldHost, address, pubkey, state.Trustee1,
		ProposeAddAccountArgs{VID: -1, Roles: "NodeAdmin"})
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, tx.RawLog)

	for _, who := range []string{state.Trustee2, state.Trustee3} {
		tx, err = ApproveAddAccount(dcldHost, address, who)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, "approve %s: %s", who, tx.RawLog)
	}

	// 9. Pre-add sanity check: pool should not yet know about this address.
	out, _ := ExecuteCLIWithBin(dcldHost,
		"query", "validator", "node", "--address", address)
	require.True(t, strings.Contains(string(out), "Not Found"),
		"validator node should not exist pre-add, got: %s", string(out))

	// 10. Get the tendermint validator pubkey from inside the container.
	vpubOut, err := DockerExec(ValidatorDemoContainerName,
		"./dcld", "tendermint", "show-validator")
	require.NoError(t, err)
	vpubkey := strings.TrimSpace(string(vpubOut))

	// 11. Submit `tx validator add-node` from inside the container.
	_, err = DockerExecShell(ValidatorDemoContainerName, fmt.Sprintf(
		`echo test1234 | ./dcld tx validator add-node --pubkey='%s' --moniker=node-demo --from=%s --yes`,
		vpubkey, accountName,
	))
	require.NoError(t, err, "tx validator add-node")

	// 12. Seed cosmovisor/genesis/bin so the helper has a binary to launch.
	require.NoError(t, SeedCosmovisorGenesis(ValidatorDemoContainerName))

	// 13. Start node_helper in the background and give it a moment to bind.
	_, err = dockerCmd("exec", "-d", ValidatorDemoContainerName,
		DCLUserHome+"/node_helper.sh")
	require.NoError(t, err, "start node_helper.sh")
	time.Sleep(10 * time.Second)

	// 14. Capture the owner (cosmosvaloper...) address from the new pool entry.
	state.ValidatorAccountName = accountName
	state.ValidatorAddress = mustExtractOwner(t, dcldHost, address)
	t.Logf("validator-demo address=%s, owner=%s, account=%s",
		address, state.ValidatorAddress, state.ValidatorAccountName)
}

// mustExtractOwner queries `validator node --address` and returns the `owner`
// field. The bash version pipes through `jq -r '.owner'` after a fixed sleep.
func mustExtractOwner(t *testing.T, dcldHost, address string) string {
	t.Helper()

	var lastErr error
	for i := 0; i < 30; i++ {
		out, err := ExecuteCLIWithBin(dcldHost,
			"query", "validator", "node", "--address", address,
		)
		if err == nil && !strings.Contains(string(out), "Not Found") {
			var parsed struct {
				Owner string `json:"owner"`
			}
			if jerr := json.Unmarshal(out, &parsed); jerr == nil && parsed.Owner != "" {
				return parsed.Owner
			}
			lastErr = fmt.Errorf("unable to parse owner from: %s", string(out))
		} else if err != nil {
			lastErr = err
		}
		time.Sleep(time.Second)
	}

	t.Fatalf("could not resolve validator owner address: %v", lastErr)

	return ""
}

// DisableValidatorNode submits `tx validator disable-node --from=<account>`
// from inside the validator-demo container. Mirrors the per-script disable.
func DisableValidatorNode(account string) error {
	_, err := DockerExecShell(ValidatorDemoContainerName, fmt.Sprintf(
		`echo test1234 | dcld tx validator disable-node --from=%s --yes`, account,
	))

	return err
}

// EnableValidatorNode submits `tx validator enable-node --from=<account>`
// from inside the validator-demo container.
func EnableValidatorNode(account string) error {
	_, err := DockerExecShell(ValidatorDemoContainerName, fmt.Sprintf(
		`echo test1234 | dcld tx validator enable-node --from=%s --yes`, account,
	))

	return err
}

// ProposeDisableValidatorNode runs propose-disable-node from the host chain.
func ProposeDisableValidatorNode(dcldBin, validatorAddress, from string) (*utils.TxResult, error) {
	return ExecuteTxWithBin(dcldBin,
		"tx", "validator", "propose-disable-node",
		"--address", validatorAddress,
		"--from", from,
	)
}

// ApproveDisableValidatorNode runs approve-disable-node from the host chain.
func ApproveDisableValidatorNode(dcldBin, validatorAddress, from string) (*utils.TxResult, error) {
	return ExecuteTxWithBin(dcldBin,
		"tx", "validator", "approve-disable-node",
		"--address", validatorAddress,
		"--from", from,
	)
}

// HasProposedDisable reports whether a pending propose-disable-node proposal
// exists for the validator address. Used to decide whether the per-script
// flow should propose first or jump straight to approvals.
func HasProposedDisable(dcldBin, validatorAddress string) bool {
	out, err := ExecuteCLIWithBin(dcldBin,
		"query", "validator", "proposed-disable-node",
		"--address", validatorAddress,
	)
	if err != nil {
		return false
	}

	return !strings.Contains(string(out), "Not Found")
}

// RunValidatorDisableEnableFlow is the per-script docker exec disable/enable
// sequence common to scripts 01/03/05/06/07/10. The exact step ordering
// matches the bash:
//
//	docker exec disable-node           (from validator-demo's account)
//	docker exec enable-node
//	host propose-disable-node          (from trustee_1) — only if no pending
//	                                    proposal exists; script 02+ inherits
//	                                    the previous script's tail-propose.
//	host approve-disable-node × N      (from trustee_2 … trustee_(approvers+1))
//	docker exec enable-node
//	host propose-disable-node          (from trustee_1) — leaves it proposed
//
// `approvers` controls how many trustees approve before the final
// re-enable. Scripts 01/02/04 use 2; scripts 03+ use 3 (also Trustee4).
func RunValidatorDisableEnableFlow(t *testing.T, state *UpgradeTestState, dcldBin string, approvers []string) {
	t.Helper()

	if state.ValidatorAddress == "" || state.ValidatorAccountName == "" {
		t.Log("RunValidatorDisableEnableFlow: validator-demo not initialized; skipping")

		return
	}

	require.NoError(t, DisableValidatorNode(state.ValidatorAccountName), "disable-node")
	require.NoError(t, EnableValidatorNode(state.ValidatorAccountName), "enable-node")

	// Propose only if no pending proposal carried over from the previous
	// script's tail-propose. Bash scripts 02+ skip this initial propose for
	// the same reason — the proposal is already on-chain.
	if !HasProposedDisable(dcldBin, state.ValidatorAddress) {
		tx, err := ProposeDisableValidatorNode(dcldBin, state.ValidatorAddress, state.Trustee1)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, "propose-disable-node: %s", tx.RawLog)
	}

	// Approvals.
	for _, who := range approvers {
		tx, err := ApproveDisableValidatorNode(dcldBin, state.ValidatorAddress, who)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, "approve-disable-node %s: %s", who, tx.RawLog)
	}

	// Final re-enable inside the container.
	require.NoError(t, EnableValidatorNode(state.ValidatorAccountName), "final enable-node")

	// Closing propose-disable-node (leaves the node in proposed-disable state
	// for the next script to inherit). Skip when a proposal is still open —
	// the approval count above may not have reached the disable threshold
	// (e.g. 5-trustee genesis where ceil(2/3*5)=4 approvals are required but
	// the per-script flow only contributes 3 incl. the implicit proposer
	// vote), or the inherited proposal carried over the v0.12→v1.2 binary
	// boundary with a different approvals shape. In either case there's
	// already an open proposal for the next script to inherit, so a redundant
	// propose would fail with "Disable proposal already exists".
	if !HasProposedDisable(dcldBin, state.ValidatorAddress) {
		tx, err := ProposeDisableValidatorNode(dcldBin, state.ValidatorAddress, state.Trustee1)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, "trailing propose-disable-node: %s", tx.RawLog)
	}
}

// QueryAllValidatorNodes runs `dcld query validator all-nodes` from inside the
// validator-demo container — mirroring the bash "Get node" verification.
func QueryAllValidatorNodes() ([]byte, error) {
	return DockerExecShell(ValidatorDemoContainerName,
		`echo test1234 | dcld query validator all-nodes`,
	)
}
