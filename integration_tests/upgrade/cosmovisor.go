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

import "fmt"

// SeedCosmovisorGenesis ensures the cosmovisor `genesis/bin` directory exists
// inside the container and that the active dcld binary is copied in. This is
// the ritual the bash scripts repeat before any `cosmovisor add-upgrade`
// invocation against an ad-hoc container.
func SeedCosmovisorGenesis(container string) error {
	if _, err := DockerExec(container, "mkdir", "-p", CosmovisorGenesisBin); err != nil {
		return err
	}

	if _, err := DockerExec(container, "cp", "-f", "./dcld", CosmovisorGenesisBin+"/"); err != nil {
		return err
	}

	return nil
}

// CosmovisorAddUpgrade registers a new upgrade plan binary with cosmovisor
// inside `container`. `binaryInContainer` is the path to the new binary that
// has already been `docker cp`-ed into the container (typically
// "$DCLDir/dcld").
func CosmovisorAddUpgrade(container, planName, binaryInContainer string) error {
	// `cosmovisor add-upgrade` requires the plan name and the binary path.
	// Quote both to survive embedded spaces if any future plan names use them.
	cmd := fmt.Sprintf("cosmovisor add-upgrade %q %q", planName, binaryInContainer)
	_, err := DockerExecShell(container, cmd)

	return err
}
