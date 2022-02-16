
# TestNet 2.0 Launch Guide

This document is a step-by-step guide for DCL TestNet 2.0 launch ceremony and describes in details the following
stages:

* Pre-Ceremony 
* Ceremony
* Post-Ceremony

Please see [Running a DCLedger Node](../../../docs/running-node.md) to get the general understanding
of a node setup logic along with requirements for the hardware and operating system.

## I. Pre-Ceremony
The following steps are expected to be done **before** the ceremony.

1. Configure VN Node

   1.1. Ensure a DCL user is in sudoers list (required only for the ceremony):
    *   Note. by default `ubuntu` user is expected as a running user for the DCL service.
        You can use/create another one if it doesn't work for you. In any case you will need to ensure
        that the user can do `sudo`.

   1.2. Login as a DCL user

   1.3. (Optional) Clean up the system

    ```bash
    $ sudo systemctl stop dcld
    $ sudo rm -f "$(which dcld)"
    $ rm -rf "$HOME/.dcl"
    ```

   1.4. Get the release artifacts:

    ```bash
    $ curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v0.6.1/dcld
    $ curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v0.6.1/dcld.service
    $ curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v0.6.1/run_dcl_node
    $ curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v0.6.1/update_peers
    ```

    1.5. Put `dlcd` binary in a folder listed in `$PATH` (e.g. `/usr/bin/`) and set a proper owner and executable permissions. 
    
    ```bash
    $ sudo cp -f ./dcld -t /usr/bin
    $ sudo chown <dcl-user> /usr/bin/dcld
    $ sudo chmod u+x /usr/bin/dcld
    # verification
    $ dcld version
    ```

    1.6. Configure the firewall
     * p2p and RPC (by default: `26656` and `26657` respectively) should be available for TCP connections.
       For Ubuntu:

        ```bash
        $ sudo ufw allow 26656/tcp
        $ sudo ufw allow 26657/tcp
        ```
     *  In case of IP filtering rules ensure they allow incoming and outcoming connections from/to other peers.

     * (Optional) you may consider to allow the following ports:
        *   gRPC (default: `9090`)
        *   REST (default: `1317`)

    1.7. Share VN's IP address (in Slack or in a special doc) 

 2. Generate NodeAdmin keys

    2.1. Go to VN Node

    2.2. Generate keys

    ```bash
    dcld keys add <key-name> 2>&1 | tee <key-name>.dclkey.data
    ```

    **IMPORTANT** keep generated data (especially the mnemonic) securely
    
    2.3. Share generated `address` and `pubkey` (in Slack or in a special doc) 

3. [Optional] Generate Trustee keys

    3.1. Choose a machine where Trustee keys will be hold (it can be either VN Node, or a separate machine with `dcld` binary)

    3.2. Generate keys

    ```bash
    dcld keys add <key-name> 2>&1 | tee <key-name>.dclkey.data
    ```

    **IMPORTANT** keep generated data (especially the mnemonic) securely

    3.3. Share generated `address` and `pubkey` (in Slack or in a special doc) 

4. [Optional] Configure ON Node

    Do steps 1.1 - 1.6 for all ON Nodes.
    
5. [CSA Only] Create `persistent_peers` file containing, share in Slack/doc

## II. Ceremony: Genesis Node
The following steps are expected to be done **during** the ceremony.

   6. [CSA Only] Run genesis node

        6.1. Make sure that all VNs accept incoming connections from this node for the given persistent peers file

        ```bash
        TBD
        ```

        6.2. Update persistent peers
        
        ```bash
        # by default path to a file is './persistent_peers'
        ./update_peers [PATH-TO-PEERS-FILE]
        ```
        6.3. Init genesis VN

        ```bash
        $ ./run_dcl_node -t genesis -c testnet-2.0 --gen-key-name <node-admin-key> [--gen-key-name-trustee <trustee-key>] node0
        ```

        6.4. Share genesis file (in Slack, doc, or GitHub)

## III. Ceremony: For Every Validator Node
The following steps are expected to be done **during** the ceremony.

   7. Add `NodeAdmin` account

        7.1. A Trustee proposes a NodeAdmin account

        ```bash
        dcld tx auth propose-add-account --address=<bench32 encoded string> --pubkey=<protobuf JSON encoded> --roles=NodeAdmin --from=<account-name>
        ```

        7.2. Trustees approve the NodeAdmin account

        ```bash
        dcld tx auth approve-add-account --address=<bench32 encoded string> --from=<account-name>
        ```


   8. Add VN node

        8.1. Download genesis 

        ```bash
        $ curl -L -O https://raw.githubusercontent.com/zigbee-alliance/distributed-compliance-ledger/master/deployment/persistent_chains/testnet-2.0/genesis.json
        ```


        8.2. Update persistent peers

        ```bash
        # by default path to a file is './persistent_peers'
        ./update_peers [PATH-TO-PEERS-FILE]
        ```

        8.3. Make sure that all VNs accept incoming connections from this node for the given persistent peers file

        ```bash
        TBD
        ```

        8.4. Init VN

        ```bash
        $ ./run_dcl_node -c testnet-2.0 <node-name>
        ```

        8.5 Wait until catchup is finished

        `dcld status` returns `"catching_up": false`

        8.6. Make a node a validator

        ```bash
        $ dcld tx validator add-node --pubkey=<validator-pubkey> \
            --moniker=<node-name> --from=<node-admin-key-name>
        ```

        (once transaction is successfully written you should see `"code": 0` in the JSON output.)

   9. [Optional] Add Trustee account  

        9.1. A Trustee proposes Trustee account

        ```bash
        dcld tx auth propose-add-account --address=<bench32 encoded string> --pubkey=<protobuf JSON encoded> --roles=Trustee --from=<account-name>
        ```

        9.2. Trustees approve Trustee account

        ```bash
        dcld tx auth approve-add-account --address=<bench32 encoded string> --from=<account-name>
        ```


## IV. Post-Ceremony: For every Observer Node

The following steps are expected to be done **after** the ceremony.

   10. Add ON Node

        10.1 Download genesis 

        ```bash
        $ curl -L -O https://raw.githubusercontent.com/zigbee-alliance/distributed-compliance-ledger/master/deployment/persistent_chains/testnet-2.0/genesis.json
        ```

        10.2. Set persistent peers. The list of persistent peers for an observer is not required to match the one used by the validators. As a general guidance you may consider to use only the peers you own and/or trust.

        10.3. Init ON

        ```bash
        $ ./run_dcl_node -t observer -c testnet-2.0 <node-name>
        ```

## V. Post-Ceremony: Validator Node Maintenance
*   on any changes in persistent peers list 
    * update `persistent_peers` field in `$HOME/.dcl/config/config.toml`cript

        ```bash
        # by default path to a file is './persistent_peers'
        ./update_peers [PATH-TO-PEERS-FILE]
        ```

    * (Optional) Update IP filtering firewall rules
    *  Restart `dcld` service 

        ```bash
        TBD
        ```