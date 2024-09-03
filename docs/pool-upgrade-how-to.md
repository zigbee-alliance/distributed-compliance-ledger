# Pool Upgrade How To

1. [Upgrade Steps](#upgrade-steps)
   - [Before Upgrade Event](#before-upgrade-event)
   - [During Upgrade Event](#during-upgrade-event)
2. [Upgrade Troubleshooting Guide](#upgrade-troubleshooting-guide)
3. [(Optional for Node Admins) Manual Upgrade](#optional-for-node-admins-manually-put-new-versions-binaries)

## Upgrade Steps

Below there are steps which need to be done to upgrade a pool to a new DCL
application version:

### Before Upgrade Event

#### 1.1. **[All Node Admins] Check that VN is running and auto-download is enabled**

Make sure that the `auto-download` parameter is present in `cosmovisor.service` (in the folder - `/etc/systemd/system/`).

If the `auto-download` parameter does not exist, then add it to the  `Service` section (see [cosmovisor.service](../deployment/cosmovisor.service)):

     ```bash
     Environment="DAEMON_ALLOW_DOWNLOAD_BINARIES=true"
     ```
After this restart the `cosmovisor.service`
   ```bash
   sudo systemctl daemon-reload
   sudo systemctl restart cosmovisor
   ```

 **_Note:_**  In this case, `true` means that the **`auto-download`** is enabled. If you want to disable **`auto-download`**, you can set the value to`false` (`DAEMON_ALLOW_DOWNLOAD_BINARIES=false`) or remove the above line from the `cosmovisor.service`.


#### 1.2. **[Dev Team] New Release and Upgrade Name**

A new release of [distributed-compliance-ledger](https://github.com/zigbee-alliance/distributed-compliance-ledger)
   project is issued. The code of the new released application version must be
   augmented by a new upgrade handler which will handle an upgrade to this new
   version. This upgrade handler must have a unique name which will serve as the
   name of the upgrade to this new version.

#### 1.3. **[Trustees] Choose Upgrade Height**:
Choose a ledger height not reached yet at which all the nodes in the pool must be upgraded to the new application version is
   chosen.

### During Upgrade Event

#### 2.1. **[A Trustee] ProposeUpgrade Transaction**:

One of the trustees proposes the upgrade using the following steps:

   1. Calculates SHA-256 or SHA-512 checksums of the new application version
      binaries (for the supported platforms, usually Ubuntu) taken from the project release.
      This can be done using `sha256sum` or `sha512sum` tool.

      For example:

      ```bash
      sha256sum ./dcld
      ```
      Please note, that it must be called against the `dcld` binary, not the platform archive itself. So, for Ubuntu, either take a `dcld` binary from the root folder of the release, or extract it from ` dcld.ubuntu.tar.gz`.

   2. Sends [`ProposeUpgrade`](./transactions.md#propose_upgrade) transaction
      with the name of the new upgrade handler, the chosen ledger height and the
      info containing URLs of the new application version binaries for supported
      platforms with the calculated checksums.

      For example:

      ```bash
      dcld tx dclupgrade propose-upgrade --name=vX.X --upgrade-height=<int64> --upgrade-info="{\"binaries\":{\"linux/amd64\":\"https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/vX.X.X/dcld?checksum=sha256:50708d4f7e00da347d4e678bf26780cd424232461c4bb414f72391c75e39545a\"}}" --from=alice
      ```

####  2.2. **[Trustees] ApproveUpgrade Transactions**:

Other trustees approve the proposed upgrade until it turns into the approved state and is scheduled (this happens when the count of approvals including the proposal reaches 2/3 of the total number of trustees).

Each of them uses the following steps to accomplish this:

1. Verifies that the application binaries URLs provided in the proposed
   upgrade plan `Info` field are valid and that the files referenced by them match the provided checksums.

    1. View the proposed upgrade plan:

       ```bash
       dcld query dclupgrade proposed-upgrade --name=vX.X
       ```
    2. Re-calculates checksums of the new application version binaries (for the supported platforms) taken from the project release. This can be done using `sha256sum` tool.

         ```bash
         sha256sum ./dcld
         ```
       Please note, that it must be called against the `dcld` binary, not the platform archive itself. So, for Ubuntu, either take a `dcld` binary from the root folder of the release, or extract it from ` dcld.ubuntu.tar.gz`.

2. Sends [`ApproveUpgrade`](./transactions.md#approve_upgrade) transaction
    with the name of the proposed upgrade.

    For example:

    ```bash
    dcld tx dclupgrade approve-upgrade --name=vX.X --from=bob
    ```

#### 2.3. **[Trustees] Ensure That Upgrade Has Been Scheduled**:

It makes sense to ensure that the upgrade has been approved and scheduled.

   Example how to view the approved upgrade plan:

   ```bash
   dcld query dclupgrade approved-upgrade --name=vX.X
   ```

   Command to view the currently scheduled upgrade plan:

   ```bash
   dcld query upgrade plan
   ```

#### 2.4. **[Trusteees, Node Admins] Wait until the upgrade height is reached**

The upgrade is applied on all the nodes in the pool when the ledger reaches the height specified in the upgrade plan.

Command to view the current pool status:

   ```bash
   dcld status
   ```

The current ledger height is reported in `SyncInfo` -> `latest_block_height`
field.

Example of command to check whether the upgrade was applied:

   ```bash
   dcld query upgrade applied vX.X
   ```

If the upgrade with the passed name was applied, this command output will
contain information about it.

 
## Upgrade Troubleshooting Guide

### Useful Commands

- Start node: `sudo systemctl start cosmovisor`
- Stop node: `sudo systemctl stop cosmovisor`
- Restart node: `sudo systemctl restart cosmovisor`
- Node service status: `systemctl status cosmovisor`
- Node logs: `journalctl -u cosmovisor.service -f`
- Node status via CLI: `./dcld status`. The value of `latest_block_height` reflects the current node height.

### Check Transaction Result

There is CometBFT Endpoint to get a transaction by its hash:
```
curl -X GET "https://<node-ip>:<port>/tx?hash=0x74504B24ED59A424E436656E5E9A11034C7A7C7ED3BE7C3CDEA1ED387EF62967&prove=true" -H "accept: application/json"
```
So, for the Test Net and CSA ON it will look like
```
https://on.test-net.dcl.csa-iot.org:26657/tx?hash=0xADB4ED112D00BFDE33958BCD108865686ADF51210C32C20272A55C23E5C26C0D&prove=true
```

### Up-to-date Persistent Peers
Make sure that VN has a correct list of persistent peers set (`persistent_peers` field in `$HOME/.dcl/config/config.toml`).
See [VN Running Docs](running-node-manual/vn.md).

On any changes in persistent peers list, update `persistent_peers` field in `$HOME/.dcl/config/config.toml`

```bash
curl -L -O https://raw.githubusercontent.com/zigbee-alliance/distributed-compliance-ledger/master/deployment/scripts/update_peers

# by default path to a file is './persistent_peers.txt'
./update_peers [PATH-TO-PEERS-FILE]
```

### Firewall Rules
If adding a new VN, it must inherit current firewall rules for the existing VN.

### Config File
Ensure configuration file for a VN is correct.
See [VN](running-node-manual/vn.md) and [Full Node](running-node-manual/full-node.md).

Make sure that the configuration file is created for a correct Network (TestNet or MainNet).
In particular, that correct `persistent_peers` are set depending on the network.

### Deleting and restoring the DB
In some cases, DB needs to be deleted and restored from another node. For example, when 
- a node is not at the current version after upgrade (Observer Node may also experience this)
- Corrupted database as a result of breaking changes in code not applied correctly after upgrade

Steps to restore a DB from another machine:
- Delete data `$HOME/.dcl/data`  on Node1
- Stop a running up-to-date Node2 (to have consistent data) `sudo systemctl stop cosmovisor`
- Copy the data folder `$HOME/.dcl/data` from Node2 to Node1 
- Run both nodes `sudo systemctl start cosmovisor`

### Manual upgrade
If a node fails to download the required version binary, manual upgrade will be needed. See [Manual Upgrade](#optional-for-node-admins-manual-upgrade).

### Block height increase scenario even with wrong software version
It should not happen, VN should stop adding to the chain if it is not at the right software version.






## (Optional for Node Admins) Manual Upgrade

If `auto-download` is enabled (see Step 1.1), then no manual steps are required to be done by `Node Admins`.
The correct `binary` will be downloaded and the `checksum` will be verified automatically (downloaded binary hash is equal to the one specified in the `propose-upgrade` request), so the downloaded binary can be trusted.

Nevertheless, it's recommended for all `Node Admins` to manually verify and put the `binaries` as described below prior to reaching the upgrade height.

**_Note_**: If `auto-download` is enabled and the `binary` is put manually to the correct folder, then this `binary` will be used for upgrade and no `auto-download` will happen. If you follow this way, then you need to complete the `steps` described below.

1. Switches current user to the user on behalf of whom `cosmovisor` service
   is running:

   ```bash
   su - <USERNAME>
   ```

   where `<USERNAME>` is the corresponding username.

   The command will ask for the user's password. Enter it.

2. Downloads the application binary from the URL specified in the upgrade
   plan `Info` field and corresponding to the node platform.

   Command to view the current scheduled upgrade plan:

   ```bash
   dcld query upgrade plan
   ```

3. Verifies that the downloaded application binary matches the checksum
   specified in the URL. This can be done automatically together with the
   previous step by [`go-getter`](https://github.com/hashicorp/go-getter)
   download tool (its executable binaries for various platforms can be
   downloaded from <https://github.com/hashicorp/go-getter/releases>).

   For example:

   ```bash
   go-getter https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/vX.X.X/dcld?checksum=sha256:50708d4f7e00da347d4e678bf26780cd424232461c4bb414f72391c75e39545a $HOME/Downloads
   ```

   `go-getter` verifies that the downloaded file matches the checksum when
   the URL is provided in the specified format. If the downloaded file
   checksum does not equal to the checksum provided in the URL, `go-getter`
   reports that checksums did not match.

4. Creates a directory with the name of the upgrade within
   `$HOME/.dcl/cosmovisor/upgrades`, creates `bin` sub-directory within the
   created directory, and puts the new application binary into `bin`
   sub-directory.

   For example:

   ```bash
   cd $HOME/.dcl/cosmovisor
   mkdir -p upgrades
   cd upgrades
   mkdir vX.X.X
   cd vX.X.X
   mkdir bin
   cd $HOME/Downloads
   cp ./dcld $HOME/.dcl/cosmovisor/upgrades/vX.X.X/bin/
   ```

5. Sets proper owner and permissions for the new application binary.

   For example:

   ```bash
   sudo chown $(whoami) $HOME/.dcl/cosmovisor/upgrades/vX.X.X/bin/dcld
   sudo chmod a+x $HOME/.dcl/cosmovisor/upgrades/vX.X.X/bin/dcld
   ```

**Steps [3-5] can be automated using the following command which downloads and verifies that the downloaded application binary matches the checksum specified in the URL**

      ```bash
      curl -fsSL https://raw.githubusercontent.com/zigbee-alliance/distributed-compliance-ledger/master/deployment/scripts/install.sh | SHA256SUM=ea0e16eed3cc30b5a7f17299aca01b5d827b9a04576662d957af02608bca0fb6 bash
      ```

   | Variable   | Default                                     | Description                                  |
               |:-----------|:--------------------------------------------|----------------------------------------------|
   | DEBUG      | false                                       | Enables verbose mode during the execution    |
   | DCL_HOME   | $HOME/.dcl                                  | DCL home folder                              |
   | VERSION    |                                             | DCL binary version to be upgraded            |
   | DEST       | $DCL_HOME/cosmovisor/upgrades/v$VERSION/bin | Destination path for DCL binary              |
   | SHA256SUM  |                                             | SHA256 sum value for DCL binary verification |
