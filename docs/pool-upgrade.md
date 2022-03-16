# Pool Upgrade

## Overview

A pool upgrade is an automated procedure of updating DCL application on all the
nodes in the pool. Besides the application binary update, the upgrade procedure
can include store migrations for application modules.

An upgrade is scheduled in advance for a specific height of the ledger which is
not reached yet. When the ledger reaches the height of the scheduled upgrade,
the upgrade procedure is performed on all the nodes in the pool simultaneously.

Any upgrade plan has the following fields:

* `Name: string` - the upgrade name for which the new application version must
    contain an associated upgrade handler; this upgrade handler is responsible
    for running store migrations for application modules that change in the new
    application version (the upgrade name is unique; after an upgrade has been
    proposed, it is impossible to propose an upgrade with the same name any time
    in the future).
* `Height: int64` - the height of the ledger at which the upgrade must be
    applied on all the nodes in the pool.
* `Info: optional(string)` - a string containing any additional information
    about the upgrade, e.g. URLs for downloading the new application version
    binaries for supported platforms (see below).

## Workflow

Initially an upgrade plan is proposed by one of trustees using
[propose-upgrade](./transactions.md#propose_upgrade) command. Then the proposed
upgrade plan has to be approved by the majority of trustees (2/3 including the
trustee who has proposed it) using
[approve-upgrade](./transactions.md#approve_upgrade) command. When the necessary
count of approvals is gathered, the upgrade plan turns into the approved state
and is actually scheduled.

There can be multiple proposed upgrade plans at the same time but not more than
one scheduled upgrade plan at a time. If there is currently a scheduled upgrade
plan and another upgrade plan turns into the approved state, then the latter is
scheduled and the former is actually cancelled. When the upgrade procedure is
completed, the current scheduled upgrade plan is cleared. Please note, once an
upgrade is approved, the approved upgrade entity remains in the store forever
(no matter if the upgrade is later completed or cancelled).

For the upgrade procedure to be feasible in an automated mode, the application
process `dcld` is controlled as a sub-process by the parent process
`cosmovisor`. Cosmovisor is a standard process manager for Cosmos SDK
application binaries. `cosmovisor` uses a directory tree where for each next
scheduled upgrade the node admin must in advance create a new directory with the
name of this upgrade and put the new application version binary to its `bin`
sub-directory. `cosmovisor` also maintains `current` symbolic link which points
to the current application version directory. See [cosmovisor
documentation](https://github.com/cosmos/cosmos-sdk/tree/cosmovisor/v1.0.0/cosmovisor)
for details.

When the ledger reaches the height specified in the current scheduled upgrade
plan, `dcld` notifies `cosmovisor` that the upgrade must be applied and stops.
`cosmovisor`, having been notified, performs data back-up, switches `current`
symbolic link to the new application version directory and launches the new
`dcld` binary which performs necessary store migrations and clears the current
scheduled upgrade plan on start.

## Application Binary Download

Downloading the new application version binary and providing it to the proper
location for cosmovisor is a manual routine for the node admin. For all the node
admins to be aware of where to download the new application version binary from,
it is recommended to provide a value for `Info` field of an upgrade plan in the
format specified in **Item 1** at
<https://github.com/cosmos/cosmos-sdk/tree/cosmovisor/v1.0.0/cosmovisor#auto-download>.
DCL application does not support the audo-download feature for which this format
is primarily intended but it is quite convenient for providing URLs for
downloading the application version binaries for supported platforms. Please
note that the URLs should include checksums. This allows to verify that no false
binary is run. It is recommended to use
[`go-getter`](https://github.com/hashicorp/go-getter) tool for downloading the
application binaries because it verifies that the downloaded file matches the
checksum when the URL is provided in the specified format. If the downloaded
file checksum does not equal to the checksum provided in the URL, `go-getter`
reports that checksums did not match. To view `Info` field value of an upgrade
plan, just execute an appropriate query command from `dclupgrade` or `upgrade`
module. See [Upgrade CLI commands reference](./transactions.md#upgrade) for
details.
