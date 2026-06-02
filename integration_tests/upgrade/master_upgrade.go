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
	"os"
	"strings"
	"time"
)

// BuildMasterImage runs `docker build -f Dockerfile-build-master` to produce
// the dcld-build-master image. Equivalent to bash:
//
//	docker build -f "$MASTER_UPGRADE_DOCKERFILE" -t "$MASTER_UPGRADE_IMAGE" .
//
// The master build downloads a fresh Go module cache inside the container
// (~2 GB), which trips "no space left on device" on the GitHub runner unless
// we reclaim space first. Historical dcld binaries (already used by phases
// 01-09) and dangling Docker artifacts are removed before the build.
func BuildMasterImage() error {
	freeDiskBeforeMasterBuild()

	_, err := dockerCmd("build",
		"-f", MasterUpgradeDockerfile,
		"-t", MasterUpgradeImage,
		".",
	)

	return err
}

// freeDiskBeforeMasterBuild reclaims disk space before the master container is
// built. All steps are best-effort — failures are ignored because the build
// itself will surface a clearer error if anything is actually broken.
func freeDiskBeforeMasterBuild() {
	// Historical dcld binaries downloaded by EnsureAllBinaries are no longer
	// needed after we've reached the master upgrade phase (~80-100 MB each).
	// The directory itself must stay because ExtractMasterBinary's `docker cp`
	// uses BinariesDir/dcld_master as its destination and docker cp requires
	// the parent directory to already exist.
	_ = os.RemoveAll(BinariesDir)
	_ = os.MkdirAll(BinariesDir, 0o755)

	// Reclaim Docker layer / build / volume cache. -af keeps no questions; the
	// localnet containers are still running and pinned, so their images stay.
	_, _ = dockerCmd("system", "prune", "-af")
	_, _ = dockerCmd("builder", "prune", "-af")
}

// CreateMasterContainer creates (but does not start) a container from the
// dcld-build-master image so we can `docker cp` the binary out.
func CreateMasterContainer() error {
	_, err := dockerCmd("container", "create",
		"--name", MasterUpgradeContainerName,
		MasterUpgradeImage,
	)

	return err
}

// ExtractMasterBinary copies /go/bin/dcld out of the master container onto the
// host so subsequent steps can hand it off to localnet nodes.
func ExtractMasterBinary(hostPath string) error {
	_, err := DockerCp(MasterUpgradeContainerName+":/go/bin/dcld", hostPath)
	if err != nil {
		return err
	}

	return os.Chmod(hostPath, 0o755)
}

// GetMasterPlanName runs the master image and returns the short git HEAD hash
// of the bundled checkout — bash:
//
//	docker run "$MASTER_UPGRADE_IMAGE" /bin/sh -c "cd /go/src/distributed-compliance-ledger && git rev-parse --short HEAD"
func GetMasterPlanName() (string, error) {
	out, err := dockerCmd("run", "--rm", MasterUpgradeImage, "/bin/sh", "-c",
		"cd /go/src/distributed-compliance-ledger && git rev-parse --short HEAD",
	)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}

// PrepareCosmovisorUpgradeOnLocalnetNodes hands the master binary to every
// localnet node container that is present (node0..node3, observer0,
// lightclient0) and registers it with cosmovisor under `planName`. Mirrors
// the per-node loop in script 10.
func PrepareCosmovisorUpgradeOnLocalnetNodes(planName, hostBinaryPath string) error {
	for _, node := range []string{"node0", "node1", "node2", "node3", "observer0", "lightclient0"} {
		// Skip nodes that don't have a host-side .localnet directory.
		if _, err := os.Stat(LocalnetDir + "/" + node); err != nil {
			continue
		}

		// Drop the host master binary into the node's bind-mounted directory.
		dst := LocalnetDir + "/" + node + "/dcld"
		if err := copyFile(hostBinaryPath, dst); err != nil {
			return fmt.Errorf("copy master binary to %s: %w", dst, err)
		}

		// Register the upgrade with cosmovisor running inside the node.
		if err := CosmovisorAddUpgrade(node, planName, DCLDir+"/dcld"); err != nil {
			return fmt.Errorf("cosmovisor add-upgrade on %s: %w", node, err)
		}
	}

	return nil
}

// copyFile copies src to dst preserving mode 0755 (executable).
func copyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	return os.WriteFile(dst, data, 0o755)
}

// WaitForObserverVersion polls `dcld version` inside containerName until it
// equals expectedVersion or timeout elapses. Returns nil on match.
func WaitForObserverVersion(containerName, expectedVersion string, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		time.Sleep(time.Second)

		if !containerIsRunning(containerName) {
			continue
		}

		out, err := DockerExec(containerName, "dcld", "version")
		if err != nil {
			continue
		}
		if strings.TrimSpace(string(out)) == expectedVersion {
			return nil
		}
	}

	return fmt.Errorf("container %s did not report dcld version %q within %s",
		containerName, expectedVersion, timeout)
}

// WaitForCatchingUpStatus polls `dcld status` inside containerName until its
// catching_up field equals expected (true or false) or timeout elapses.
func WaitForCatchingUpStatus(containerName string, expected bool, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	want := fmt.Sprintf(`"catching_up":%t`, expected)

	for time.Now().Before(deadline) {
		time.Sleep(time.Second)

		if !containerIsRunning(containerName) {
			continue
		}

		// Run as root so we don't trip on home-dir permissions during catch-up.
		out, err := dockerCmd("exec", "--user", "root", containerName, "dcld", "status")
		if err != nil {
			continue
		}
		if strings.Contains(string(out), want) {
			return nil
		}
	}

	return fmt.Errorf("container %s did not reach catching_up=%v within %s",
		containerName, expected, timeout)
}

// containerIsRunning returns true if the named container exists and is in
// "running" state. Errors are treated as not-running.
func containerIsRunning(name string) bool {
	out, err := DockerInspect(name)
	if err != nil {
		return false
	}

	return strings.Contains(string(out), `"Status": "running"`)
}
