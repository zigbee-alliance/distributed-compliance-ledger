# Running a Private Sentry Node manually

## Prerequisites
Make sure you have all [prerequisites](./prerequisites.md) set up
## Deployment steps
#### Steps are the same as in [full-node.md](./full-node.md) 
except:
- Set `<node-type>` to `"private-sentry"` when using [run_dcl_node](./full-node.md#step-8-can-be-automated-using-rundclnode-script)
- Configure node specific parameters before running the node:
    
    [`config.toml`]
    ```toml
    [p2p]
    pex = true
    persistent_peers = # `Validator` node with private IP + other orgs' validator/sentry nodes with public IPs
    private_peer_ids = # `Validator` node id
    unconditional_peers = # `Validator` node id
    addr_book_strict = false
    ```

    [`app.toml`]

    ```toml
    [state-sync]
    snapshot-interval = "snapshot-interval"
    snapshot-keep-recent = "snapshot-keep-recent"
    ```