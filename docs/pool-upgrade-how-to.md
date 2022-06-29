# Pool Upgrade How To

Below there are steps which need to be done to upgrade a pool to a new DCL
application version:

1. **[Dev Team] New Release and Upgrade Name**: A new release of
   [distributed-compliance-ledger](https://github.com/zigbee-alliance/distributed-compliance-ledger)
   project is issued. The code of the new released application version must be
   augmented by a new upgrade handler which will handle an upgrade to this new
   version. This upgrade handler must have a unique name which will serve as the
   name of the upgrade to this new version.

2. **[Trustees] Upgrade Height**: A ledger height not reached yet at which all
   the nodes in the pool must be upgraded to the new application version is
   chosen.

3. **[A Trustee] ProposeUpgrade**: One of trustees proposes the upgrade using
   the following steps:

   1. Calculates SHA-256 or SHA-512 checksums of the new application version
      binaries (for the supported platforms) taken from the project release.
      This can be done using `sha256sum` or `sha512sum` tool.

      For example:

      ```bash
      sha256sum ./dcld
      ```

   2. Sends [`ProposeUpgrade`](./transactions.md#propose_upgrade) transaction
      with the name of the new upgrade handler, the chosen ledger height and the
      info containing URLs of the new application version binaries for supported
      platforms with the calculated checksums.

      For example:

      ```bash
      dcld tx dclupgrade propose-upgrade --name=vX.X.X --upgrade-height=<int64> --upgrade-info="{\"binaries\":{\"linux/amd64\":\"https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/vX.X.X/dcld?checksum=sha256:50708d4f7e00da347d4e678bf26780cd424232461c4bb414f72391c75e39545a\"}}" --from=alice
      ```

4. **[Trustees] ApproveUpgrade**: Other trustees approve the proposed upgrade
   until it turns into the approved state and is scheduled (this happens when
   the count of approvals including the proposal reaches 2/3 of the total count
   of trustees). Each of them uses the following steps to accomplish this:

   1. Re-calculates checksums of the new application version binaries (for the
      supported platforms) taken from the project release. This can be done
      using `sha256sum` / `sha512sum` tool.

      For example:

      ```bash
      sha256sum ./dcld
      ```

   2. Ensures that the re-calculated values are equal to the checksums specified
      in the proposed upgrade plan `Info` field.

      Example how to view the proposed upgrade plan:

      ```bash
      dcld query dclupgrade proposed-upgrade --name=vX.X.X
      ```

   3. Verifies that the application binaries URLs provided in the proposed
      upgrade plan `Info` field are valid and that the files referenced by them
      match the provided checksums.

   4. Verifies that `Height` field of the proposed upgrade plan has the proper
      value.

   5. Sends [`ApproveUpgrade`](./transactions.md#approve_upgrade) transaction
      with the name of the proposed upgrade.

      For example:

      ```bash
      dcld tx dclupgrade approve-upgrade --name=vX.X.X --from=bob
      ```

5. **[Anyone] Ensure That Upgrade Has Been Scheduled**: It makes sense to ensure
   that the upgrade has been approved and scheduled.

   Example how to view the approved upgrade plan:

   ```bash
   dcld query dclupgrade approved-upgrade --name=vX.X.X
   ```

   Command to view the current scheduled upgrade plan:

   ```bash
   dcld query upgrade plan
   ```

6. **[All Node Admins]**: 
   
   - If `auto-download` is enabled , then no manual steps are required to be done by `Node Admins`. Auto-download verifies that the downloaded binary hash is equal to the one specified in the `propose-upgrade` request, so the downloaded binary can be trusted.
   Nevertheless, it's recommended for all `Node Admins` to manually verify and put the `binaries` as described below.<br>
      - Make sure that the `auto-download` parameter is present in `cosmovisor.service` (in the folder - `/etc/systemd/system/`) and if the `auto-download` parameter does not exist, then add it to the  `Service` section (see [cosmovisor.service](../deployment/cosmovisor.service)).
         
         For example:

         ```bash
         Environment="DAEMON_ALLOW_DOWNLOAD_BINARIES=true"
         ```

         > **_Note:_**  In this case, `true` means that the **`auto-download`** is enabled. If you want to disable **`auto-download`**, you can set the value to`false` (`DAEMON_ALLOW_DOWNLOAD_BINARIES=false`) or remove the above line from the `cosmovisor.service`.

      - If you change the configuration in the `cosmovisor.service`, you need to restart `cosmovisor.service`:
         ```bash
         sudo systemctl daemon-reload
         sudo systemctl restart cosmovisor
         ```

         > **_Note:_** Also make sure that propose plan should contain the link and the checksum to the binary according to the format specified in cosmovisor docs. 
         See the [Command Line Arguments And Environment Variables in Cosmovisor](https://github.com/cosmos/cosmos-sdk/tree/main/cosmovisor#command-line-arguments-and-environment-variables).

      - If `auto-download` is enabled and no `binary` is put manually, then the correct `binary` will be downloaded and the `checksum` will be verified automatically. 
   
   - **(Recommended)** If `auto-download` is enabled and the `binary` is put manually to the correct folder, then this `binary` will be used for upgrade and no `auto-download` will happen. If you follow this way, then you need to complete the `steps` described below.

      1. Switches current user to the user on behalf of whom `cosmovisor` service
         is running:

         ```bash
         su - <USERNAME>
         ```

         where `<USERNAME>` is the corresponding username

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

         #### *** Steps [3-5] can be automated using the following command which downloads and verifies that the downloaded application binary matches the checksum specified in the URL
            command:

            For example:

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

7. **Upgrade Is Applied**: The upgrade is applied on all the nodes in the pool
   when the ledger reaches the height specified in the upgrade plan.

   Command to view the current pool status:

   ```bash
   dcld status
   ```

   The current ledger height is reported in `SyncInfo` -> `latest_block_height`
   field.

   Example of command to check whether the upgrade was applied:

   ```bash
   dcld query upgrade applied vX.X.X
   ```

   If the upgrade with the passed name was applied, this command output will
   contain information about it.
