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
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
)

// dockerCmd shells out to `docker` with the given args and returns combined
// stdout/stderr. Errors include the full output to make CI logs actionable.
func dockerCmd(args ...string) ([]byte, error) {
	cmd := exec.CommandContext(context.Background(), "docker", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return out, fmt.Errorf("docker %s: %w, output: %s",
			strings.Join(args, " "), err, string(out))
	}

	return out, nil
}

// DockerRun starts a container. Pass `docker run` flags — typically
// -d --name --ip -p --network -i <image>.
func DockerRun(args ...string) ([]byte, error) {
	return dockerCmd(append([]string{"run"}, args...)...)
}

// DockerExec runs a single command inside a running container (argv form,
// not shell-evaluated).
func DockerExec(container string, args ...string) ([]byte, error) {
	return dockerCmd(append([]string{"exec", container}, args...)...)
}

// DockerExecShell runs a shell command inside a container via `sh -c`.
func DockerExecShell(container, command string) ([]byte, error) {
	return DockerExec(container, "/bin/sh", "-c", command)
}

// DockerCp copies a file/directory between host and container or vice versa.
// Use `<container>:<path>` to denote the container side.
func DockerCp(src, dst string) ([]byte, error) {
	return dockerCmd("cp", src, dst)
}

// DockerKillPID1 sends SIGTERM to PID 1 inside the container, which causes
// the container's main process (cosmovisor / dcld) to exit and the container
// to stop. Used to harvest coverage from a node before pool teardown.
func DockerKillPID1(container string) ([]byte, error) {
	return DockerExec(container, "kill", "1")
}

// DockerWait blocks until the named container exits, returning its exit code.
func DockerWait(container string) ([]byte, error) {
	return dockerCmd("wait", container)
}

// DockerInspect returns the raw JSON from `docker container inspect`.
func DockerInspect(container string) ([]byte, error) {
	return dockerCmd("container", "inspect", container)
}

// DockerCleanup stops and removes a container if it exists. Errors are
// silently swallowed — best-effort cleanup.
func DockerCleanup(container string) {
	inspectOut, err := exec.CommandContext(context.Background(), "docker", "container", "inspect", container).CombinedOutput()
	if err != nil {
		// Container does not exist.
		return
	}

	if bytes.Contains(inspectOut, []byte(`"Status": "running"`)) {
		_, _ = dockerCmd("container", "kill", container)
	}

	_, _ = dockerCmd("container", "rm", "-f", container)
}
