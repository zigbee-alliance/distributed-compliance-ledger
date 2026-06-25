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

// Package upgrade is the end-to-end Go test suite for cosmovisor chain
// upgrades across historical dcld binary versions.
package upgrade

const (
	// DCLUserHome is the home directory of the `dcl` user inside the dcledger
	// container image (matches Dockerfile DAEMON_HOME parent).
	DCLUserHome = "/var/lib/dcl"

	// DCLDir is the dcld application directory inside the container, derived
	// from DCLUserHome. Matches DAEMON_HOME in the Dockerfile.
	DCLDir = DCLUserHome + "/.dcl"

	// CosmovisorGenesisBin is the cosmovisor "genesis" binary directory inside
	// the container. dcld is copied here before any add-upgrade calls.
	CosmovisorGenesisBin = DCLDir + "/cosmovisor/genesis/bin"

	// DockerNetwork is the network name docker-compose creates for the
	// localnet (project name + "_localnet").
	DockerNetwork = "distributed-compliance-ledger_localnet"

	// ChainID matches the chain id baked into the localnet by genlocalnetconfig.sh.
	ChainID = "dclchain"

	// Node0Conn is the gRPC endpoint of node0 inside the localnet docker network.
	Node0Conn = "tcp://192.167.10.2:26657"

	// Passphrase is the keyring passphrase used by every account across the
	// localnet (only suitable for tests).
	Passphrase = "test1234" //nolint:gosec

	// LocalnetDir is the host-side directory holding per-node config / data,
	// produced by `make localnet_init`.
	LocalnetDir = ".localnet"
)

// Container names — kept stable so cleanup can target them by name.
const (
	MasterUpgradeContainerName = "dcld-build-master-inst"
	NewObserverContainerName   = "new-observer"
	ValidatorDemoContainerName = "validator-demo"
)

// Default IPs / ports for ad-hoc containers attached to the localnet network.
const (
	ValidatorDemoIP         = "192.167.10.6"
	ValidatorDemoP2PPort    = 26670
	ValidatorDemoClientPort = 26671
)

// BinariesDir is the host directory holding downloaded historical dcld binaries.
const BinariesDir = "/tmp/dcld_bins"

// HistoricalVersions is the set of dcld releases the upgrade tests exercise.
// Kept in sync with prepare-dcld-versions.sh's download loop.
var HistoricalVersions = []string{
	"0.12.0", "0.12.1", "1.2.2", "1.4.3", "1.4.4", "1.5.1", "1.5.2",
	"1.6.0",
}
