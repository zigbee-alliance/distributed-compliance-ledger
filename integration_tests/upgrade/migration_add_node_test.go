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
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// runAddNewNodeAfterUpgrade spins up a fresh observer container at
// IP 192.167.10.28 installed with dcld
// v0.12.0. Seeds cosmovisor with the master binary under the master plan
// name. Starts the node and polls until:
//
//  1. dcld version reports v0.12.0 (old binary running)
//  2. catching_up=true (catch-up procedure started)
//  3. catching_up=false (catch-up procedure finished)
//  4. dcld version reports MasterPlanName (cosmovisor swapped the binary)
//
//nolint:funlen
func runAddNewNodeAfterUpgrade(t *testing.T, state *UpgradeTestState) {
	t.Helper()

	require.NotEmpty(t, state.MasterPlanName,
		"MasterPlanName must be populated by the master upgrade step first")

	dcldOld, err := EnsureBinary(DCLDVersionV012)
	require.NoError(t, err)

	t.Cleanup(func() { DockerCleanup(NewObserverContainerName) })
	DockerCleanup(NewObserverContainerName)

	// Step 1: run the new observer container on the localnet network.
	MustRun(t, "StartNewObserverContainer", func(t *testing.T) {
		t.Helper()
		portMap := fmt.Sprintf("%d-%d:26656-26657",
			NewObserverNodeP2PPort, NewObserverNodeClientPort)
		_, runErr := DockerRun(
			"-d",
			"--name", NewObserverContainerName,
			"--ip", NewObserverIP,
			"-p", portMap,
			"--network", DockerNetwork,
			"-i", "dcledger",
		)
		require.NoError(t, runErr, "docker run new observer")
	})

	// Step 2: install dcld v0.12.0 into the container.
	MustRun(t, "InstallOldDcld", func(t *testing.T) {
		t.Helper()
		_, err := DockerCp(dcldOld, NewObserverContainerName+":"+DCLUserHome+"/dcld")
		require.NoError(t, err, "docker cp v0.12.0 into observer")
	})

	// Step 3: configure node identity + peers + listen address.
	MustRun(t, "ConfigureNodeFiles", func(t *testing.T) {
		t.Helper()
		_, err := DockerExec(NewObserverContainerName,
			"./dcld", "init", NewObserverContainerName, "--chain-id", ChainID)
		require.NoError(t, err)

		_, err = DockerCp(
			LocalnetDir+"/node0/config/genesis.json",
			NewObserverContainerName+":"+DCLDir+"/config",
		)
		require.NoError(t, err)

		peers, err := readPersistentPeersFromHost(LocalnetDir + "/node0/config/config.toml")
		require.NoError(t, err)

		// Replace persistent_peers and laddr inside the container's config.toml.
		_, err = DockerExecShell(NewObserverContainerName,
			fmt.Sprintf(`sed -i 's|persistent_peers = ""|persistent_peers = "%s"|g' %s/config/config.toml`,
				peers, DCLDir))
		require.NoError(t, err)

		_, err = DockerExecShell(NewObserverContainerName,
			fmt.Sprintf(`sed -i 's|laddr = "tcp://127.0.0.1:26657"|laddr = "tcp://0.0.0.0:26657"|g' %s/config/config.toml`, DCLDir))
		require.NoError(t, err)
	})

	// Step 4: seed cosmovisor/genesis/bin with the old binary.
	MustRun(t, "SeedCosmovisorGenesis", func(t *testing.T) {
		t.Helper()
		require.NoError(t, SeedCosmovisorGenesis(NewObserverContainerName))
	})

	// Step 5: register the master upgrade in cosmovisor.
	MustRun(t, "RegisterMasterUpgradeOnObserver", func(t *testing.T) {
		t.Helper()
		_, err := DockerCp(DcldMasterBinaryPath,
			NewObserverContainerName+":"+DCLDir+"/dcld")
		require.NoError(t, err, "docker cp dcld_master into observer")

		require.NoError(t,
			CosmovisorAddUpgrade(NewObserverContainerName, state.MasterPlanName, DCLDir+"/dcld"),
		)
	})

	// Step 6: start the cosmovisor-managed node helper.
	MustRun(t, "StartNodeHelper", func(t *testing.T) {
		t.Helper()
		_, err := dockerCmd("exec", "-d", NewObserverContainerName,
			"sh", "-c",
			"/var/lib/dcl/./node_helper.sh | tee /proc/1/fd/1",
		)
		require.NoError(t, err)
	})

	// Step 7: verify pre-upgrade dcld reports v0.12.0 inside the container.
	MustRun(t, "ObserverReportsOldVersion", func(t *testing.T) {
		t.Helper()
		require.NoError(t, WaitForObserverVersion(
			NewObserverContainerName, DCLDVersionV012, 10*time.Second,
		))
	})

	// Step 8 + 9: catch-up starts and then finishes. 15min polling window
	// per status.
	MustRun(t, "ObserverCatchUpLifecycle", func(t *testing.T) {
		t.Helper()
		require.NoError(t, WaitForCatchingUpStatus(
			NewObserverContainerName, true, 15*time.Minute,
		), "expected catching_up=true (start of catch-up)")

		require.NoError(t, WaitForCatchingUpStatus(
			NewObserverContainerName, false, 15*time.Minute,
		), "expected catching_up=false (catch-up complete)")
	})

	// Step 10: post-catch-up the binary should be the master plan name.
	MustRun(t, "ObserverReportsMasterVersion", func(t *testing.T) {
		t.Helper()
		require.NoError(t, WaitForObserverVersion(
			NewObserverContainerName, state.MasterPlanName, 30*time.Second,
		), "observer should have been upgraded to master plan binary")
	})
}
