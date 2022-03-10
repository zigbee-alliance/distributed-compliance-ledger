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
      This can be done using `sha256sum` or `sha512sum` tool. For example:
      ```
      sha256sum ./dcld
      ```

   2. Sends [`ProposeUpgrade`](./transactions.md#propose_upgrade) transaction
      with the name of the new upgrade handler, the chosen ledger height and the
      info containing URLs of the new application version binaries for supported
      platforms with the calculated checksums. For example:
      ```
      dcld tx dclupgrade propose-upgrade --name=v0.7.0 --upgrade-height=10000 --upgrade-info="{\"binaries\":{\"linux/amd64\":\"https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v0.7.0/dcld?checksum=sha256:50708d4f7e00da347d4e678bf26780cd424232461c4bb414f72391c75e39545a\"}}" --from=alice
      ```

4. **[Trustees] ApproveUpgrade**: Other trustees approve the proposed upgrade
   until it turns into the approved state and is scheduled (this happens when
   the count of approvals including the proposal reaches 2/3 of the total count
   of trustees). Each of them uses the following steps to accomplish this:

   1. Re-calculates checksums of the new application version binaries (for the
      supported platforms) taken from the project release. This can be done
      using `sha256sum` / `sha512sum` tool. For example:
      ```
      sha256sum ./dcld
      ```

   2. Ensures that the re-calculated values are equal to the checksums specified
      in the proposed upgrade plan `Info` field. Example how to view the
      proposed upgrade plan:
      ```
      dcld query dclupgrade proposed-upgrade --name=v0.7.0
      ```

   3. Verifies that the application binaries URLs provided in the proposed
      upgrade plan `Info` field are valid and that the files referenced by them
      match the provided checksums.

   4. Verifies that `Height` field of the proposed upgrade plan has the proper
      value.

   5. Sends [`ApproveUpgrade`](./transactions.md#approve_upgrade) transaction
      with the name of the proposed upgrade. For example:
      ```
      dcld tx dclupgrade approve-upgrade --name=v0.7.0 --from=bob
      ```

5. **Ensure That Upgrade Has Been Scheduled**: It makes sense to ensure that the
   upgrade has been approved and scheduled. Example how to view the approved
   upgrade plan:
   ```
   dcld query dclupgrade approved-upgrade --name=v0.7.0
   ```
   Command to view the current scheduled upgrade plan:
   ```
   dcld query upgrade plan
   ```

6. **[All Node Admins] Download New Binary**: Before the ledger reaches the
   height specified in the upgrade plan, each node admin does the following
   steps:

    1. Downloads the application binary from the URL specified in the upgrade
       plan `Info` field and corresponding to the node platform. Command to view
       the current scheduled upgrade plan:
       ```
       dcld query upgrade plan
       ```

    2. Verifies that the downloaded application binary matches the checksum
       specified in the URL. This can be done automatically together with the
       previous step by [`go-getter`](https://github.com/hashicorp/go-getter)
       download tool. For example:
       ```
       go-getter https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v0.7.0/dcld?checksum=sha256:50708d4f7e00da347d4e678bf26780cd424232461c4bb414f72391c75e39545a ~/Downloads
       ```
       `go-getter` verifies that the downloaded file matches the checksum when
       the URL is provided in the specified format. If the downloaded file
       checksum does not equal to the checksum provided in the URL, `go-getter`
       reports that checksums did not match.

    3. Creates a directory with the name of the upgrade within
       `/var/lib/<USERNAME>/.dcl/cosmovisor/` (where <USERNAME> is the name of
       the user on behalf of whom `cosmovisor` service is running), creates
       `bin` sub-directory within the created directory, and puts the new
       application binary within `bin` sub-directory. Example (for the user
       `ubuntu`):
       ```
       cd /var/lib/ubuntu/.dcl/cosmovisor
       mkdir v0.7.0
       cd v0.7.0
       mkdir bin
       cd ~/Downloads
       cp ./dcld /var/lib/ubuntu/.dcl/cosmovisor/v0.7.0/bin
       ```

7. **Upgrade Is Performed**: The upgrade is performed on all the nodes in the
   pool when the ledger reaches the height specified in the upgrade plan.
