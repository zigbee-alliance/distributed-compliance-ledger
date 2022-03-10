# Pool Upgrade How To

Below there are steps which need to be done to upgrade a pool to a new DCL application version:

1. A new release of [distributed-compliance-ledger](https://github.com/zigbee-alliance/distributed-compliance-ledger) project is issued. The code of the new released application version must be augmented by a new upgrade handler with a unique name which will serve as the name of the upgrade to this version.
1. A ledger height not reached yet at which all the nodes in the pool must be upgraded to the new application version is chosen.
1. Checksums of the application binaries for all the supported platforms are calculated.
1. One of trustees sends `ProposeUpgrade` transaction (see [Pool Upgrade documentation](./pool-upgrade.md) for fields specification). So the upgrade plan is proposed.
1. Other trustees send `ApproveUpgrade` transactions with the name of the proposed upgrade until the count of approvals (including the proposal) reaches 2/3 of the total count of trustees. So the upgrade plan is approved and scheduled.
1. Before the ledger reaches the height specified in the upgrade plan, each node admin does the following steps:
    1. Downloads the application binary for the node platform from the new release.
    1. Verifies that the downloaded application binary matches the checksum specified in the URL. This can be done automatically by [`go-getter`](https://github.com/hashicorp/go-getter) download tool.
    1. Creates a directory with the name of the upgrade within `/var/lib/<USERNAME>/.dcl/cosmovisor/` (where <USERNAME> is the name of the user on behalf of whom `cosmovisor` service is running), creates `bin` sub-directory within the created directory, and puts the new application binary within `bin` sub-directory.
1. The upgrade is performed on all the nodes in the pool when the ledger reaches the height specified in the upgrade plan.
