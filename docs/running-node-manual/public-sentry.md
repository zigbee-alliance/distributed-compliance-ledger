# Running a Public Sentry Node manually

## Prerequisites

Make sure you have all [prerequisites](./prerequisites.md) set up

## Deployment steps

### Steps are the same as in [full-node.md](./full-node.md)

except:

- Set `<node-type>` to `"public-sentry"` when using [run_dcl_node](./full-node.md#step-8-can-be-automated-using-rundclnode-script)
- Configure node specific parameters before running the node:

    [`config.toml`]

    ```toml
    [p2p]
    pex = true
    persistent_peers = # `Private Sentry` nodes with private IPs
    ```

    [`app.toml`]

    ```toml
    [state-sync]
    snapshot-interval = "snapshot-interval"
    snapshot-keep-recent = "snapshot-keep-recent"
    ```
