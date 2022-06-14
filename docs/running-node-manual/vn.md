# Running a Validator Node manually

## Prerequisites

Make sure you have all [prerequisites](./prerequisites.md) set up

## Deployment steps

### Steps [1-9] are the same as in [full-node.md](./full-node.md)

except:

- Set `<node-type>` to `"validator"` when using [run_dcl_node](./full-node.md#step-8-can-be-automated-using-rundclnode-script)

- Configure node specific parameters before running the node:
  
    [`config.toml`]

    ```toml
    [p2p]
    pex = false
    persistent_peers = # `Private Sentry` nodes with private IPs
    addr_book_strict = false

    [consensus]
    create_empty_blocks = false
    create_empty_blocks_interval = "600s" # 10 mins
    ```

    [`app.toml`]

    ```toml
    [state-sync]
    snapshot-interval = "snapshot-interval"
    snapshot-keep-recent = "snapshot-keep-recent"
    ```

### 10. Create keys for a node admin and a trustee (optional) accounts

```bash
./dcld keys add "<key-name>" 2>&1 | tee "<key-name>.dclkey.data"
```

- Remember generated `address` and `pubkey` they will be used later.
You can retrieve `address` and `pubkey` values anytime using `./dcld keys show <name>`.
Of course, only on the machine where the keypair was generated.

> Notes: It's important to keep the generated data (especially a mnemonic that allows to recover a key) in a safe place

### 11. Add validator node to the network

- Get this node's tendermint validator address: `./dcld tendermint show-address`.
    Expected output format:

    ```text
    cosmosvalcons1yrs697lxpwugy7h465wskwu2a5w9dgklx608f0
    ```

- Get this node's tendermint validator pubkey: `./dcld tendermint show-validator`.
    Expected output format:

    ```text
    cosmosvalconspub1zcjduepqcwg4eenpcxgs0269xuup5jlzj3pdquxlvj494cjxtqtcathsq7esfrsapa
    ```

- Note that *validator address* and *validator pubkey* are not the same as `address` and `pubkey` were used for account creation.

- Wait until the value of `catching_up` field gets to `false` value.

- Add validator node: `dcld tx validator add-node --pubkey=<validator pubkey> --moniker=<node name> --from=<key name>`.
If the transaction has been successfully written you would find `"code": 0` in the output JSON.

### 12. Check the node is running and participates in consensus

- Get the list of all nodes: `dcld query validator all-nodes`.
The node must present in the list and has the following params: `power:10` and `jailed:false`.

- Get the node status: `dcld status --node tcp://<node_ip>:26657`.
The value of `node ip` matches to `[rpc] laddr` field in `$HOME/.dcl/config/config.toml`
(TCP or UNIX socket address for the RPC server to listen on).  
Make sure that `result.sync_info.latest_block_height` is increasing over the time (once in about 10 mins).

- Get the list of nodes participating in the consensus for the last block: `dcld query tendermint-validator-set`.
  - You can pass the additional value to get the result for a specific height: `dcld query tendermint-validator-set 100`.
