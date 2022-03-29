# Running Nodes in an existing network
> **_NOTE:_** Extra node specific configuration (out of scope of this section) may apply depending on node type or infrastructure.<br>
> Refer to [running-node.md](./running-node.md) for more details

Possible options when adding Validator, Observer, Sentry or Seed nodes to existing DCL network:

### 1) State sync (recommended)
- Prerequisites:
  - Latest binary version running in existing network
  - State sync snapshots should be enabled in at least one node in the network.<br>
  Perform the following steps to enable state sync snapshots in running node 
    
    - Edit the following attributes in `$HOME/.dcl/config/app.toml`
      ```
      # snapshot-interval specifies the block interval at which local state sync snapshots are
      # taken (0 to disable). Must be a multiple of pruning-keep-every.
      snapshot-interval = <snapshot-interval>

      # snapshot-keep-recent specifies the number of recent snapshots to keep and serve (0 to keep all).
      snapshot-keep-recent = <snapshot-keep-recent>
      ```
      > **_NOTE:_** Snapshot interval must currently be a multiple of the pruning-keep-every(default 100), to prevent heights from being pruned while taking snapshots. 
      > It’s also usually a good idea to keep at least 2 recent snapshots, such that the previous snapshot isn’t removed while a node is attempting to state sync using it.
    - Restart the node
    > **_NOTE:_** Nodes with snapshots should allow `tendermint p2p` connections to be able to share snapshots with synching peers.<br>
    > - Do not forget to check firewall and security settings<br> 
    > - We do not recommend enabling state sync for `Validator Nodes` for security reasons
- Steps:
  - Init node:

    ```bash
    ./dcld init "<node-name>" --chain-id "<chain-id>"
    ```
  - Enable state sync in config `$HOME/.dcl/config/config.toml`:
    ```
    [statesync]
    enable = true

    # RPC servers (comma-separated) for light client verification of the synced state machine and
    # retrieval of state data for node bootstrapping. Also needs a trusted height and corresponding
    # header hash obtained from a trusted source, and a period during which validators can be trusted.
    #
    # For Cosmos SDK-based chains, trust_period should usually be about 2/3 of the unbonding time (~2
    # weeks) during which they can be financially punished (slashed) for misbehavior.

    rpc_servers = "http(s)://<host>:<port>,http(s)://<host>:<port>"
    trust_height = <trust-height>
    trust_hash = "<trust-hash>"
    trust_period = "168h0m0s"
    ```
    > **_NOTE:_**  You should provide at least 2 addresses for `rpc_servers`. It can be 2 identical addresses

    You can use the following command to obtain `<trust-height>` and `<trust-hash>` of your network
    ```
    curl -s http(s)://<host>:<port>/commit | jq "{height: .result.signed_header.header.height, hash: .result.signed_header.commit.block_id.hash}"
    ```
    - `<host>` - RPC endpoint host of the network being joined
    - `<port>` - RPC endpoint port of the network being joined

  - Run the new node (refer to [running-node.md](./running-node.md) for node specific instructions)

    > **_NOTE:_** State sync is not attempted if the node has any local state (LastBlockHeight > 0)

- Pros:
  - No manual activities are needed (except configuring `$HOME/.dcl/config/config.toml`)
  - Several order of magnitudes faster than catchup
  - Requires only latest binary version

- Cons:
  - The node will have a truncated block history, starting from the height of the snapshot.

* References:
  - https://blog.cosmos.network/cosmos-sdk-state-sync-guide-99e4cf43be2f

### 2) Take the data from another node
- Prerequisites:
  - Access to a machine running an up-to-date node to copy data from

- Steps:
  - Stop running an up-to-date node (to have consistent data), and copy the data folder
  - Put that data to the new node's data folder (can be automated)
  - Run the new node

- Pros:
  - No specific node configuration needed (just start the new node)

- Cons:
  - Needs an access to another machine (may not be possible to everyone)
  - Node that data is copied from should be stopped (downtime)
  - Probably error prone due to manual operations

### 3) Catchup from genesis:
- Prerequisites:
  - All binary versions used for upgrading (using `cosmovisor`) existing network up to current state must be available
- Steps:
  - Add a node with a binary version as was at the genesis state
  - Let the nodes catch-up and play all updates/migrations

- Pros:
  - The new node contains all the blocks from blockchain (full history)

- Cons:
  - Can take quite a lot of time (depends on blockchain size)
  - Probably error-prone (if at least one migration has a bug, catchup fails)

* References:
- https://docs.cosmos.network/master/core/upgrade.html#syncing-a-full-node-to-an-upgraded-blockchain