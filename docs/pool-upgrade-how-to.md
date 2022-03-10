# Pool Upgrade How To

Below there are steps which need to be done to upgrade a pool to a new DCL
application version:

1. A new release of
   [distributed-compliance-ledger](https://github.com/zigbee-alliance/distributed-compliance-ledger)
   project is issued. The code of the new released application version must be
   augmented by a new upgrade handler which will handle an upgrade to this new
   version. This upgrade handler must have a unique name which will serve as the
   name of the upgrade to this new version.
2. A ledger height not reached yet at which all the nodes in the pool must be
   upgraded to the new application version is chosen.
3. SHA-256 or SHA-512 checksums of the new application version binaries for
   supported platforms are calculated.
4. One of trustees sends [`ProposeUpgrade`](./transactions.md#propose_upgrade)
   transaction with the name of the new upgrade handler, the chosen ledger
   height and the info containing URLs of the application binaries for supported
   platforms with the calculated checksums. So the upgrade plan is proposed.
5. Other trustees send [`ApproveUpgrade`](./transactions.md#approve_upgrade)
   transactions with the name of the proposed upgrade until the count of
   approvals (including the proposal) reaches 2/3 of the total count of
   trustees. So the upgrade plan is approved and scheduled.
6. Before the ledger reaches the height specified in the upgrade plan, each node
   admin does the following steps:
    1. Downloads the application binary at the URL specified in the upgrade plan
       `Info` field and corresponding to the node platform.
    2. Verifies that the downloaded application binary matches the checksum
       specified in the URL. This can be done automatically by
       [`go-getter`](https://github.com/hashicorp/go-getter) download tool.
    3. Creates a directory with the name of the upgrade within
       `/var/lib/<USERNAME>/.dcl/cosmovisor/` (where <USERNAME> is the name of
       the user on behalf of whom `cosmovisor` service is running), creates
       `bin` sub-directory within the created directory, and puts the new
       application binary within `bin` sub-directory.
7. The upgrade is performed on all the nodes in the pool when the ledger reaches
   the height specified in the upgrade plan.
