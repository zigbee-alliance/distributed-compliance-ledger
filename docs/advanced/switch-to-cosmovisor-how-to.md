# Switch to Cosmovisor: How To

This document describes the procedure how to switch a node from direct use of `dcld` binary to use of `cosmovisor` process manager which controls `dcld` process and supports DCL application upgrades that includes `dcld` binary updates and store migrations.

Switching to use of `cosmovisor` is performed by `switch_to_cosmovisor` script. This procedure does not include any store migrations. So it can be applied only if the diff between the target and source versions of `dcld` does not include any breaking changes of the store.

Assumptions:
* `switch_to_cosmovisor` script assumes that old `dcld` binary is located in `/usr/bin` directory. The script is operable in this case only.

Pre-requisites:
* `dcld` is launched as `dcld` systemd service.
* `dcld` systemd service is currently running, i.e. is in active state.
* The following files, taken from a DCL release, have been put into the current working directory from where the script is executed:
    * new `dcld` binary (that will be controlled by `cosmovisor` and so should include the upgrade feature)
    * `cosmovisor` binary
    * `cosmovisor.service`
