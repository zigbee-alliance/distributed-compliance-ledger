# Running a Public Sentry Node manually

## Prerequisites

Make sure you have all [prerequisites](./prerequisites.md) set up.

## Deployment steps

### Steps are the same as in [full-node.md](./full-node.md)

except:

- Set `<node-type>` to `"public-sentry"` when using [run_dcl_node](./full-node.md#step-8-can-be-automated-using-rundclnode-script)
- Configure node specific parameters before running the node:

    [`config.toml`]

    ```toml
    [p2p]
    pex = true
    persistent_peers = "<node1-ID>@<node1-IP>:26656,..." # See the comment below on what values should be set here 
    ```

    [`app.toml`]

    ```toml
    [state-sync]
    snapshot-interval = "snapshot-interval"
    snapshot-keep-recent = "snapshot-keep-recent"
    ```
- `persistent_peers` value:
  - If your VN doesn't use any Private Sentry nodes, then it must point to the `Validator` node with private IP.
  - Otherwise, it must point to the Private Sentry nodes with private IPs.
  - Use the following command to get `node-ID` of a node: `./dcld tendermint show-node-id`.
