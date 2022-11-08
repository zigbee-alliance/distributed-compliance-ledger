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

> **_Note:_** Steps [1-2] are done for every node while steps [3-6] are done only once
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
    sudo systemctl stop cosmvisor
    sudo systemctl stop dcld
    sudo rm -f "$(which cosmovisor)"
    sudo rm -f "$(which dcld)"
    rm -rf "$HOME/.dcl"
    ```

    1.5. Make sure that no processes running on ports `26656` and `26657`
    ```bash
    # install lsof
    sudo apt install lsof

    # kill the processes running on ports 26656 and 26657
    # command fails if no process is running on given port
    sudo kill -9 $(sudo lsof -t -i:26656)
    sudo kill -9 $(sudo lsof -t -i:26657)
    ```

    1.6. Get the release artifacts (DCL v0.12.0):

    ```bash
    curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v0.12.0/dcld
    curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v0.12.0/cosmovisor
    curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v0.12.0/cosmovisor.service
    curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v0.12.0/run_dcl_node
    ```

    1.7. Put `cosmovisor` binary in a folder listed in `$PATH` (e.g. `/usr/bin/`) and set a proper owner and executable permissions.

    ```bash
    sudo cp -f ./cosmovisor -t /usr/bin
    sudo chown "<dcl-user>" /usr/bin/cosmovisor
    sudo chmod u+x /usr/bin/cosmovisor
    ```

    1.8. Configure the firewall

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

    3.3. Share generated `address` and `pubkey` in #dcl-enrollment Slack Channel.

    `address` and `pubkey` can be found in the `dcld keys show --output text "<admin-account-name>"` output.

4. **[Optional] Generate Trustee keys**

    4.1. Choose a machine where Trustee keys will be hold (it can be either VN Node, or a separate machine with `dcld` binary)

    4.2. Generate keys

    ```bash
    ./dcld keys add "<trustee-account-name>" 2>&1 | tee "<trustee-account-name>.dclkey.data"
    ```

    **IMPORTANT** keep generated data (especially the mnemonic) securely.

    4.3. Share generated `address` and `pubkey` in #dcl-enrollment Slack Channel.

    `address` and `pubkey` can be found in the `dcld keys show --output text "<trustee-account-name>"` output.

5. **Share Node info with other Node Admins** in #dcl-enrollment Slack Channel.

    5.1. Share Sentry Node's public IP address (VN's if a sentry node is not used) 

    5.2. Share Sentry Node's `id` (VN's if a sentry node is not used)
    - Get node's `id` using the following command:
        ```bash
        ./dcld tendermint show-node-id
        ```
        
    5.3. CSA will create and share persistent_peers.txt based on the shared data.

6. **[VN only] Update Empty Block Interval to 600s**.
    
    6.1 Open dcl configuration file.
        ```bash
        ~/.dcl/config/config.toml
        ```
    
    6.2 Locate section `[consensus]`.
    
    6.3 Configure as follows:
    
    ```bash 
    -- FROM: --
    # EmptyBlocks mode and possible interval between empty blocks
    create_empty_blocks = true
    create_empty_blocks_interval = "0s"

    -- TO: --
    # EmptyBlocks mode and possible interval between empty blocks
    create_empty_blocks = false
    create_empty_blocks_interval = "600s" #10min 
    ```
## II. Ceremony
7. **Wait until the CSA Mainnet infrastructure is up and running**

    7.1. Access DCL [Web UI][1] and follow the ceremony procedure

8. **Add Trustee accounts to Mainnet**

    8.1. To be able to approve proposals using [Web UI][1], all Trustees (including CSA) should create a wallet in [Web UI][1] using mnemonics generated while creating a trustee key.

    8.2. CSA proposes all Trustee accounts using [Web UI][1]

    8.3. Approved Trustees approve other proposed Trustee accounts using [Web UI][1]

9. **Add NodeAdmin accounts to Mainnet**

    9.1. CSA proposes all NodeAdmin accounts using [Web UI][1]

    9.2. Trustees approve proposed NodeAdmin accounts

10. **Run Sentry Node (skip this step if you are not running a sentry node)**

    10.1. Download geneåsis

    ```bash
    curl -L -O https://raw.githubusercontent.com/zigbee-alliance/distributed-compliance-ledger/master/deployment/persistent_chains/main-net/genesis.json
    ```

    10.2. Prepare `persistent_peers.txt` file 
    - download or copy-paste up-to date `persistent_peers.txt` file to the same directory as `run_dcl_node`

    10.3. Run Sentry Node

    ```bash
    chmod u+x run_dcl_node
    ./run_dcl_node -t private-sentry -c main-net "<node-name>"
    ```
    * Naming convention for `<node-name>` is `[company name]-[node type]-[sequence number]` e.g. `<CSA-SN-01>`


    10.4. Wait until catchup is finished: `./dcld status` returns `"catching_up": false`

11. **Sentry Node Deployment Verification (skip this step if you are not running a sentry node)**

    11.1. Check the account presence on the ledger: `./dcld query auth account --address="<address>"`.

    11.2. Check the cosmovisor service is running: `systemctl status cosmovisor`

    11.3. Check the node gets new blocks: `./dcld status`. Make sure that `result.sync_info.latest_block_height` is increasing over the time (once in about 5 sec).

12. **Run Validator Node**

    12.1. Wait until your NodeAdmin account has been apporoved by CSA

    12.2. Download genesis

    ```bash
    curl -L -O https://raw.githubusercontent.com/zigbee-alliance/distributed-compliance-ledger/master/deployment/persistent_chains/main-net/genesis.json
    ```

    12.3. Prepare `persistent_peers.txt` file
    - If you are not running a sentry node - copy-paste the up-to-date `persistent_peers.txt` file to the same directory as `run_dcl_node`.
    - If you are running a sentry node - specify only your sentry node's address in `persistent_peers.txt` file in the following format: 
        - `<sentry node id>`@`<sentry node's private/public IP address>`
        - Use the following command to get node id of a node
            ```bash
            ./dcld tendermint show-node-id
            ```
        > _Note_: It is better to communicate with a sentry node using internal private ip address if both validator and sentry nodes are in the same (logical) network

    12.4. Run Validator Node (VN)

    ```bash
    chmod u+x run_dcl_node
    ./run_dcl_node -t validator -c main-net "<node-name>"
    ```
    * Naming convention for `<node-name>` is `[company name]-[node type]-[sequence number]` e.g. `<CSA-VN-01>`


    12.5. Wait until catchup is finished: `./dcld status` returns `"catching_up": false`

    12.6. Make the node a validator

    ```bash
    ./dcld tx validator add-node --pubkey="<protobuf JSON encoded validator-pubkey>" --moniker="<node-name>" --from="<admin-account-name>"
    ```

       * `[Note]` Run the following command to get `<protobuf JSON encoded validator-pubkey>`
       
         ```bash
         ./dcld tendermint show-validator
         ```

    ```bash
    ie.
    ./dcld tx validator add-node --pubkey='{"@type":"/cosmos.crypto.ed25519.PubKey","key":"MH2rju5vHc/nE6yH+SIsQZsXcsHhOVI9Zv8Pf+lm36o="}' --moniker='CSA-VN-01' --from='csa-account'
    ```

    (once transaction is successfully written you should see `"code": 0` in the JSON output.)


13. **VN Deployment Verification**

    13.1. Check the account presence on the ledger: `./dcld query auth account --address="<address>"`.

    13.2. Check the cosmovisor service is running: `systemctl status cosmovisor`

    13.3. Check the node gets new blocks: `./dcld status`. Make sure that `result.sync_info.latest_block_height` is increasing over the time (once in about 5 sec).

    13.4. Make sure the VN participates in consensus: `./dcld query tendermint-validator-set` must contain the VN's address.
    * `[Note]` Get VN's address using the following command

        ```bash
        dcld tendermint show-address
        ```


## III: Post-Ceremony: Validation

**Make sure that Sentry (VN in case Sentry is not used) nodes accept incoming connections from this node for the given persistent peers file**

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


[1]: https://webui.dcl.csa-iot.org