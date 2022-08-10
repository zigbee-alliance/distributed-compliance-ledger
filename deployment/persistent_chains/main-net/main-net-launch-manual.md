# DCL Mainnet Launch Guide

<!-- markdownlint-disable MD029 -->
This document is a step-by-step guide for those who prefer manual configuration of DCL nodes.

It  describes the following stages in details:

* Pre-Ceremony
* Ceremony
* Post-Ceremony

Please see [Running a DCLedger Node](../../../docs/running-node.md) to get the general understanding
of a node setup logic along with requirements for the hardware and operating system.

## I. Pre-Ceremony

The following steps are expected to be done **before** the ceremony.

0. **Check connectivity of your instances with other companies' instances**
    
    The following steps should be performed on all public facing nodes (VN or Sentry depending on the configuration).

    You can also use the following instructions to check the connectivity between your own internal nodes (Sentry-VN, Sentry-Observer, etc.)

    0.1. Download the following scripts:
    ```bash
    # fetch the helper scripts
    curl -L -O https://raw.githubusercontent.com/zigbee-alliance/distributed-compliance-ledger/master/deployment/scripts/utils/run_stub_server
    curl -L -O https://raw.githubusercontent.com/zigbee-alliance/distributed-compliance-ledger/master/deployment/scripts/utils/test_endpoints

    # give execute permission
    sudo chmod u+x run_stub_server test_endpoints
    ```

    0.2. Run stub server on Sentry (VN if Sentry is not used) Node instance:
    ```bash
    ./run_stub_server <port-number>
    ```
    - Script requires sudo permission
    - Stub server helps to troubleshoot network (connection, firewall, etc) issues before actual mainnet launch
    - `<port-number>` - port number the stub server will listen on
    - Run the stub server on the following ports:
        - 26656 (tendermint p2p port)
            ```bash
            ./run_stub_server 26656
            ```
    
    0.3. Once the stub servers on all public facing nodes (VN or Sentry depending on configuration) are running. It is time to check connectivity other nodes with yours.
    - Create or copy `persistent_endpoints.txt` file into the the directory as `test_endpoints` script
    - `persistent_endpoints.txt` file contains comma-separated (without spaces) list of `[IP:PORT]` pairs of all public facing nodes
        - example:
            ```txt
            1.1.1.1:26656,2.2.2.2:26656
            ```
        - Cordinate with other NodeAdmins in `#dcl-node-admins` slack channel to get up-to date list of addresses
    - Run the `test_endpoints` script to check connectivity with nodes specified in the `persistent_endpoints.txt`
        ```bash
        ./test_endpoints
        ```
        - you should see connectivity status for all [IP:PORT] pairs in the output

> **_Note:_** Steps [1-2] are done for every validator node while steps [3-5] are done only once
1. **Configure Validator/Sentry Node**

    1.1. `Ubuntu 20.04 LTS` is recommended.

    1.2. Ensure a DCL user is in sudoers list (required only for the ceremony):

    * Note. by default `ubuntu` user is expected as a running user for the DCL service.
        You can use/create another one if it doesn't work for you. In any case you will need to ensure
        that the user can do `sudo`.

    1.3. Login as a DCL user

    1.4. (Optional) Clean up the system

    ```bash
    # clean earlier version of DCL installed on the same computer.
    sudo systemctl stop dcld
    sudo rm -f "$(which dcld)"
    rm -rf "$HOME/.dcl"

    # kill any processes running on ports 26656 and 26657
    sudo kill -9 $(sudo lsof -t -i:26656)
    sudo kill -9 $(sudo lsof -t -i:26657)
    ```

    1.5. Get the release artifacts (DCL v0.12.0):

    ```bash
    curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v0.12.0/dcld
    curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v0.12.0/cosmovisor
    curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v0.12.0/cosmovisor.service
    curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v0.12.0/run_dcl_node
    ```

    1.6. Put `cosmovisor` binary in a folder listed in `$PATH` (e.g. `/usr/bin/`) and set a proper owner and executable permissions.

    ```bash
    sudo cp -f ./cosmovisor -t /usr/bin
    sudo chown "<dcl-user>" /usr/bin/cosmovisor
    sudo chmod u+x /usr/bin/cosmovisor
    ```

    1.7. Configure the firewall

    * p2p and RPC (by default: `26656` and `26657` respectively) should be available for TCP connections.
        For Ubuntu:

        ```bash
        sudo ufw allow 26656/tcp
        sudo ufw allow 26657/tcp
        ```

    * In case of IP filtering rules ensure they allow incoming and outcoming connections from/to other peers.

    * (Optional) you may consider to allow the following ports:
        * gRPC (default: `9090`)
        * REST (default: `1317`)

2. **Initialize Node**

    ```bash
    ./dcld init "<node-name>" --chain-id "main-net"
    ```

    * Naming convention for `<node-name>` is `[company name]-[node type]-[sequence number]` e.g. `<CSA-VN-01>`

3. **Generate Validator NodeAdmin keys (VN only)**

    3.1. Go to VN Node

    3.2. Generate keys

    ```bash
    ./dcld keys add "<admin-account-name>" 2>&1 | tee "<admin-account-name>.dclkey.data"
    ```

    **IMPORTANT** keep generated data (especially the mnemonic) securely.

    3.3. Share generated `address` and `pubkey` (in Slack or in a special doc).

    `address` and `pubkey` can be found in the `dcld keys show --output text "<admin-account-name>"` output.

4. **[Optional] Generate Trustee keys (VN only)**

    4.1. Choose a machine where Trustee keys will be hold (it can be either VN Node, or a separate machine with `dcld` binary)

    4.2. Generate keys

    ```bash
    ./dcld keys add "<trustee-account-name>" 2>&1 | tee "<trustee-account-name>.dclkey.data"
    ```

    **IMPORTANT** keep generated data (especially the mnemonic) securely.

    4.3. Share generated `address` and `pubkey` (in Slack or in a special doc).

    `address` and `pubkey` can be found in the `dcld keys show --output text "<trustee-account-name>"` output.

5. **Share Node info with other Node Admins** (in Slack or in a special doc)

    5.1. Share Sentry Node's (VN's if a sentry node is not used) public IP address

    5.2. Share Sentry Node's (VN's if a sentry node is not used) `id`
    - Get node's `id` using the following command:
        ```bash
        ./dcld tendermint show-node-id
        ```

## II. Ceremony
6. **Wait until the CSA Mainnet infrastructure is up and running**

    6.1. Access DCL Web UI link provided by CSA and follow the ceremony procedure

7. **Run Sentry Node (skip this section if you are not running a sentry node)**

    7.1. Download genesis

    ```bash
    curl -L -O https://raw.githubusercontent.com/zigbee-alliance/distributed-compliance-ledger/master/deployment/persistent_chains/main-net/genesis.json
    ```

    7.2. Prepare `persistent_peers.txt` file 
    - download or copy-paste up-to date `persistent_peers.txt` file to the same directory as `run_dcl_node`

    7.3. Run Sentry Node

    ```bash
    chmod u+x run_dcl_node
    ./run_dcl_node -t private-sentry -c main-net "<node-name>"
    ```
    * Naming convention for `<node-name>` is `[company name]-[node type]-[sequence number]` e.g. `<CSA-SN-01>`


    7.4. Wait until catchup is finished: `./dcld status` returns `"catching_up": false`

8. **Sentry Node Deployment Verification (skip this section if you are not running a sentry node)**

    8.1. Check the account presence on the ledger: `./dcld query auth account --address="<address>"`.

    8.2. Check the cosmovisor service is running: `systemctl status cosmovisor`

    8.3. Check the node gets new blocks: `./dcld status`. Make sure that `result.sync_info.latest_block_height` is increasing over the time (once in about 5 sec).

9. **Run Validator Node**

    9.1. Wait until your NodeAdmin account has been apporoved by CSA

    9.2. Download genesis

    ```bash
    curl -L -O https://raw.githubusercontent.com/zigbee-alliance/distributed-compliance-ledger/master/deployment/persistent_chains/main-net/genesis.json
    ```

    9.3. Prepare `persistent_peers.txt` file
    - If you are not running a sentry node - download or copy-paste up-to date `persistent_peers.txt` file to the same directory as `run_dcl_node`.
    - If you are running a sentry node - specify only your sentry node's address in `persistent_peers.txt` file in the following format: 
        - `<sentry node id>`@`<sentry node's private/public IP address>`
        - Use the following command to get node id of a node
            ```bash
            ./dcld tendermint show-validator
            ```
        > _Note_: It is better to communicate with a senrty node using internal private ip address if both validator and sentry nodes are in the same (logical) network

    9.4. Run VN

    ```bash
    chmod u+x run_dcl_node
    ./run_dcl_node -t validator -c main-net "<node-name>"
    ```
    * Naming convention for `<node-name>` is `[company name]-[node type]-[sequence number]` e.g. `<CSA-VN-01>`


    9.5. Wait until catchup is finished: `./dcld status` returns `"catching_up": false`

    9.6. Make the node a validator

    ```bash
    ./dcld tx validator add-node --pubkey="<protobuf JSON encoded validator-pubkey>" --moniker="<node-name>" --from="<admin-account-name>"
    ```

    * `[Note]` Run the following command to get `<protobuf JSON encoded validator-pubkey>`

        ```bash
        ./dcld tendermint show-validator
        ```

    (once transaction is successfully written you should see `"code": 0` in the JSON output.)


10. **VN Deployment Verification**

    10.1. Check the account presence on the ledger: `./dcld query auth account --address="<address>"`.

    10.2. Check the cosmovisor service is running: `systemctl status cosmovisor`

    10.3. Check the node gets new blocks: `./dcld status`. Make sure that `result.sync_info.latest_block_height` is increasing over the time (once in about 5 sec).

    10.4. Make sure the VN participates in consensus: `./dcld query tendermint-validator-set` must contain the VN's address.
    * `[Note]` Get VN's address using the following command

        ```bash
        dcld tendermint show-address
        ```


## III: Post-Ceremony: Validation

13. **Make sure that Sentry (VN in case Sentry is not used) nodes accept incoming connections from this node for the given persistent peers file**

    ```bash
    # fetch the helper script
    curl -L -O https://raw.githubusercontent.com/zigbee-alliance/distributed-compliance-ledger/master/deployment/scripts/test_peers_conn

    # give execute permission
    sudo chmod u+x test_peers_conn

    # run, by default it expects persistent_peers.txt in the current directory
    ./test_peers_conn
    ```

## IV. Post-Ceremony: Node Maintenance

* On any changes in persistent peers list
  * update `persistent_peers` field in `$HOME/.dcl/config/config.toml`

    ```bash
    curl -L -O https://raw.githubusercontent.com/zigbee-alliance/distributed-compliance-ledger/master/deployment/scripts/update_peers

    # by default path to a file is './persistent_peers.txt'
    ./update_peers [PATH-TO-PEERS-FILE]
    ```

* (Optional) Update IP filtering firewall rules

* Restart `cosmovisor` service

    ```bash
    systemctl restart cosmovisor
    ```

* Useful commands
  * keys:
    * `dcld keys show --output text "<name>"`: to get address and pubkey for a keyname
  * node status:
    * `systemctl status cosmovisor`: to get the node service status.
    * `journalctl -u cosmovisor.service -f`: to see node logs.
    * `dcld status [--node "tcp://<node host>:<node port>"]`: to get the current status.
    * `dcld query tendermint-validator-set [height]`: list of nodes participating in consensus
  * account status:
    * `dcld query auth account --address="<address>"`: to ensure that account is created and has assigned role
<!-- markdownlint-enable MD029 -->

## V. Post-Ceremony: Adding new nodes to mainnet

When adding new nodes to mainnet after a while, you might consider one of the options described in [running-node.md](../../../docs/running-node.md).
