# Switch to Cosmovisor: How To

This document describes the procedure of how to switch a node from direct use of
`dcld` binary to use of `cosmovisor` process manager which controls `dcld`
process and supports DCL application upgrades that include `dcld` binary updates
and store migrations. This procedure must be performed one by one on all the
nodes in the pool (validators, observers, seed nodes, sentry nodes).

Switching to use of `cosmovisor` is performed by `switch_to_cosmovisor` script.
This procedure does not include any store migrations. So it can be applied only
if the difference between the previously installed stand-alone `dcld` binary and
`dcld` binary to install with cosmovisor does not include any breaking changes
of the store.

**Pre-requisites:**

* `dcld` is launched as `dcld` systemd service.
* `dcld` service is currently in active state (i.e. running).

**Steps:**

1. Switch current user to the user on behalf of whom `dcld` service is running:

    ```bash
    su - <USERNAME>
    ```
    where `<USERNAME>` is the corresponding user name

    The command will ask for the user's password. Enter it.

2. Download new `dcld`, `cosmovisor` and `cosmovisor.service` from GitHub
  [release page](https://github.com/zigbee-alliance/distributed-compliance-ledger/releases). *(To ensure that no old `dcld` binary remains in the current direcory, try to remove it at first.)*

    Example using curl:
    ```bash
    sudo rm -f ./dcld
    curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/<release>/dcld
    curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/<release>/cosmovisor
    curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/<release>/cosmovisor.service
    ```

3. Download `switch_to_cosmovisor` script from [repository](../../deployment/scripts/)

    Example using curl:
    ```bash
    curl -L -O https://raw.githubusercontent.com/zigbee-alliance/distributed-compliance-ledger/master/deployment/scripts/switch_to_cosmovisor
    ```

    > Note:
    >
    > * `switch_to_cosmovisor` script adds the cosmovisor-controlled directory
    containing the current version of `dcld` binary to `$PATH` of current user.
    To do this the script adds a line doing the corresponding `$PATH` assignment
    to `$HOME/.profile` file. If for some reason it is not effective for your
    environment, please modify the corresponding line in the script in the way
    you need or just comment out the corresponding line and manually add
    `$HOME/.dcl/cosmovisor/current/bin` to `$PATH` of current user after
    `switch_to_cosmovisor` script is executed (see below).

4. Grant execution permission on `switch_to_cosmovisor` script:

    ```bash
    chmod u+x ./switch_to_cosmovisor
    ```

5. Run `switch_to_cosmovisor` script:

    ```bash
    ./switch_to_cosmovisor
    ```

    When it is done, it will print:
    ```
    Done
    ```
