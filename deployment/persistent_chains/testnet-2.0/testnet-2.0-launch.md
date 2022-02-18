# TestNet 2.0 Launch Guide

This document is a step-by-step guide for DCL TestNet 2.0 launch ceremony. It  describes the following stages in details:

*   Pre-Ceremony
*   Ceremony
*   Post-Ceremony

Please see [Running a DCLedger Node](../../../docs/running-node.md) to get the general understanding
of a node setup logic along with requirements for the hardware and operating system.

## I. Pre-Ceremony

The following steps are expected to be done **before** the ceremony.

1.  **Configure VN Node**

    1.1. `Ubuntu 20.04 LTS` is recommended.

    1.2. Ensure a DCL user is in sudoers list (required only for the ceremony):

    *   Note. by default `ubuntu` user is expected as a running user for the DCL service.
        You can use/create another one if it doesn't work for you. In any case you will need to ensure
        that the user can do `sudo`.

    1.3. Login as a DCL user

    1.4. (Optional) Clean up the system

    ```bash
    $ sudo systemctl stop dcld
    $ sudo rm -f "$(which dcld)"
    $ rm -rf "$HOME/.dcl"
    ```

    1.5. Get the release artifacts (DCL 0.7.0):

    ```bash
    $ curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v0.7.0/dcld
    $ curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v0.7.0/dcld.service
    $ curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v0.7.0/run_dcl_node
    ```

    1.6. Put `dlcd` binary in a folder listed in `$PATH` (e.g. `/usr/bin/`) and set a proper owner and executable permissions.

    ```bash
    $ sudo cp -f ./dcld -t /usr/bin
    $ sudo chown "<dcl-user>" /usr/bin/dcld
    $ sudo chmod u+x /usr/bin/dcld

    # verification
    $ dcld version
    ```

    1.7. Configure the firewall

    *   p2p and RPC (by default: `26656` and `26657` respectively) should be available for TCP connections.
        For Ubuntu:

        ```bash
        $ sudo ufw allow 26656/tcp
        $ sudo ufw allow 26657/tcp
        ```

    *   In case of IP filtering rules ensure they allow incoming and outcoming connections from/to other peers.

    *   (Optional) you may consider to allow the following ports:
        *   gRPC (default: `9090`)
        *   REST (default: `1317`)

2.  **Init VN Node**

    ```bash
    dcld init "<node-name>" --chain-id "testnet-2.0"
    ```

3.  **Share VN info with other Node Admins** (in Slack or in a special doc)

    3.1. Share VN's IP address

    3.2. Share VN's `id` (`node_id` field in `dcld init` command output
         or `id` field in `dcld status` command output in case node is running)

4.  **Generate NodeAdmin keys**

    4.1. Go to VN Node

    4.2. Generate keys

    ```bash
    dcld keys add "<key-name>" 2>&1 | tee "<key-name>.dclkey.data"
    ```

    **IMPORTANT** keep generated data (especially the mnemonic) securely.

    4.3. Share generated `address` and `pubkey` (in Slack or in a special doc).

    `address` and `pubkey` can be found in the `dcld keys show --output text "<key-name>"` output.

5.  **[Optional] Generate Trustee keys**

    5.1. Choose a machine where Trustee keys will be hold (it can be either VN Node, or a separate machine with `dcld` binary)

    5.2. Generate keys

    ```bash
    dcld keys add "<key-name>" 2>&1 | tee "<key-name>.dclkey.data"
    ```

    **IMPORTANT** keep generated data (especially the mnemonic) securely.

    5.3. Share generated `address` and `pubkey` (in Slack or in a special doc).

    `address` and `pubkey` can be found in the `dcld keys show --output text "<key-name>"` output.

6.  **[Optional] Configure ON Nodes**

    Do steps 1.1 - 1.7 for all ON Nodes.

7.  [CSA Only] Create `persistent_peers.txt` file containing `<node1-ID>@<node1-IP>:26656,...`  for all VNs. Share in Slack/doc.

## II. Ceremony: Genesis Node (CSA Only)

The following steps are expected to be done **during** the ceremony.

8.  **Run genesis node**

    8.1. Prepare `persistent_peers.txt` file (download or copy-paste into the file
    in the same directory as `run_dcl_node`).

    8.2. Make sure that all VNs accept incoming connections from this node for the given persistent peers file

    ```bash
    # fetch the helper script
    curl -L -O https://raw.githubusercontent.com/zigbee-alliance/distributed-compliance-ledger/master/deployment/scripts/test_peers_conn

    # run, by default it expects persistent_peers.txt in the current directory
    ./test_peers_conn
    ```

    8.3. Run genesis VN

    ```bash
    ./run_dcl_node -t genesis -c testnet-2.0 --gen-key-name "<node-admin-key>" [--gen-key-name-trustee "<trustee-key>"] "<node-name>"
    ```

    8.4. Put genesis file to GitHub (`zigbee-alliance/distributed-compliance-ledger/master/deployment/persistent_chains/testnet-2.0/genesis.json`)

## III. Ceremony: For Every Validator Node

The following steps are expected to be done **during** the ceremony.

9.  **Add `NodeAdmin` account**

    9.1. A Trustee proposes a NodeAdmin account

    ```bash
    dcld tx auth propose-add-account --address='<bench32 encoded string>' --pubkey='<protobuf JSON encoded>' --roles=NodeAdmin --from='<account-name>'
    ```

    9.2. Trustees approve the NodeAdmin account

    ```bash
    dcld tx auth approve-add-account --address='<bench32 encoded string>' --from='<account-name>'
    ```

10. **[Optional] Add Trustee account**

    12.1. A Trustee proposes Trustee account

    ```bash
    dcld tx auth propose-add-account --address='<bench32 encoded string>' --pubkey='<protobuf JSON encoded>' --roles=Trustee --from='<account-name>'
    ```

    12.2. Trustees approve Trustee account

    ```bash
    dcld tx auth approve-add-account --address='<bench32 encoded string>' --from='<account-name>'
    ```

11. **Run VN node**

    10.1. Download genesis

    ```bash
    $ curl -L -O https://raw.githubusercontent.com/zigbee-alliance/distributed-compliance-ledger/master/deployment/persistent_chains/testnet-2.0/genesis.json
    ```

    10.2. Prepare `persistent_peers.txt` file (download or copy-paste into the file
    in the same directory as `run_dcl_node`).

    10.3. Make sure that all VNs accept incoming connections from this node for the given persistent peers file

    ```bash
    # fetch the helper script
    curl -L -O https://raw.githubusercontent.com/zigbee-alliance/distributed-compliance-ledger/master/deployment/scripts/test_peers_conn

    # run, by default it expects persistent_peers.txt in the current directory
    ./test_peers_conn
    ```

    10.4. Run VN

    ```bash
    ./run_dcl_node -c testnet-2.0 "<node-name>"
    ```

    10.5 Wait until catchup is finished: `dcld status` returns `"catching_up": false`

    10.6. Make the node a validator

    ```bash
    $ dcld tx validator add-node --pubkey="<validator-pubkey>" --moniker="<node-name>" --from="<node-admin-key-name>"
    ```

    (once transaction is successfully written you should see `"code": 0` in the JSON output.)

12. **VN Deployment Verification**

    11.1. Check the account presence on the ledger: `dcld query auth account --address="<address>"`.

    11.2. Check the node service is running: `systemctl status dcld`

    11.3. Check the node gets new blocks: `dcld status`. Make sure that `result.sync_info.latest_block_height` is increasing over the time (once in about 5 sec).

    11.4. Make sure the VN participates in consensus: `dcld query tendermint-validator-set` must contain the VN's address.

## IV. Post-Ceremony: For every Observer Node

The following steps can be done **after** the ceremony.

13. **Add ON Node**

    13.1 Download genesis

    ```bash
    $ curl -L -O https://raw.githubusercontent.com/zigbee-alliance/distributed-compliance-ledger/master/deployment/persistent_chains/testnet-2.0/genesis.json
    ```

    13.2. Set persistent peers by updating `persistent_peers` field in `$HOME/.dcl/config/config.toml`.
    The list of persistent peers for an observer is not required to match the one used by the validators.
    As a general guidance you may consider to use only the peers you own and/or trust.

    13.3. Init ON

    ```bash
    dcld init "<node-name>" --chain-id "testnet-2.0"
    ```

    13.4. Run ON

    ```bash
    ./run_dcl_node -t observer -c testnet-2.0 "<node-name>"
    ```

14. **ON Deployment Verification**

    14.1. Check the node service is running: `systemctl status dcld`

    14.2. Check the node gets new blocks: `dcld status`. Make sure that `result.sync_info.latest_block_height` is increasing over the time (once in about 5 sec).

## V. Post-Ceremony: Node Maintenance

*   On any changes in persistent peers list
    *   update `persistent_peers` field in `$HOME/.dcl/config/config.toml`

        ```bash
        curl -L -O https://raw.githubusercontent.com/zigbee-alliance/distributed-compliance-ledger/master/deployment/scripts/update_peers

        # by default path to a file is './persistent_peers.txt'
        ./update_peers [PATH-TO-PEERS-FILE]
        ```

    *   (Optional) Update IP filtering firewall rules

    *   Restart `dcld` service

        ```bash
        systemctl restart dcld
        ```
*   Useful commands
    *   keys:
        *   `dcld keys show --output text "<name>"`: to get address and pubkey for a keyname
    *   node status:
        *   `systemctl status dcld`: to get the node service status.
        *   `journalctl -u dcld.service -f`: to see node logs.
        *   `dcld status [--node "tcp://<node host>:<node port>"]`: to get the current status.
        *   `dcld query tendermint-validator-set [height]`: list of nodes participating in consensus
    *   account status:
        *   `dcld query auth account --address="<address>"`: to ensure that account is created and has assigned role
