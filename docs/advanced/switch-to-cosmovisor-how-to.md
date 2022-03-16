# Switch to Cosmovisor: How To

This document describes the procedure of how to switch a node from direct use of
`dcld` binary to use of `cosmovisor` process manager which controls `dcld`
process and supports DCL application upgrades that include `dcld` binary updates
and store migrations.

Switching to use of `cosmovisor` is performed by `switch_to_cosmovisor` script.
This procedure does not include any store migrations. So it can be applied only
if the difference between the previously installed stand-alone `dcld` binary and
`dcld` binary to install with cosmovisor does not include any breaking changes
of the store.

**Pre-requisites:**

* `dcld` is launched as `dcld` systemd service.
* `dcld` service is currently in active state (i.e. running).
* The current user is the user on behalf of whom `dcld` service is launched.

**Steps:**

* Download new `dcld`, `cosmovisor` and `cosmovisor.service` from GitHub
  [release page](https://github.com/zigbee-alliance/distributed-compliance-ledger/releases)

    Example using curl:
    ```bash
    curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/<release>/dcld
    curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/<release>/cosmovisor
    curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/<release>/cosmovisor.service
    ```

* Download `switch_to_cosmovisor` script from [repository](../../deployment/scripts/)

    Example using curl:
    ```bash
    curl -L -O https://raw.githubusercontent.com/zigbee-alliance/distributed-compliance-ledger/master/deployment/scripts/switch_to_cosmovisor
    ```

* Run `switch_to_cosmovisor` script:

    ```bash
    ./switch_to_cosmovisor
    ```

    When it is done, it will print:
    ```
    Done
    ```
