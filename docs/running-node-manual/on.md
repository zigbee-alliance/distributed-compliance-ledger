# Running an Observer Node manually

## Prerequisites

Make sure you have all [prerequisites](./prerequisites.md) set up

## Deployment steps

### Steps are the same as in [full-node.md](./full-node.md)

except:

- Set `<node-type>` to `"observer"` when using [run_dcl_node](./full-node.md#step-8-can-be-automated-using-rundclnode-script)
- Configure node specific parameters before running the node depending on the ON deployment option.

#### Option 1: an Observer node is connected to another organization's Public Sentries via Seed nodes

This is the main option if you want to connect an ON to CSA public nodes.

[`config.toml`]

```toml
[p2p]
pex = true
seeds = "<seed-node-ID>@<seed-node-public-IP>:26656"
addr_book_strict = false
```

[`app.toml`]

```toml
[api]
enable = true
```

<details>
<summary>CSA `seeds` Example for Testnet 2.0 (clickable) </summary>

```bash
seeds = "8190bf7a220892165727896ddac6e71e735babe5@100.25.175.140:26656"
```

</details>

<details>
<summary>CSA `seeds` Example for MainNet (clickable) </summary>

  ```bash
seeds = "ba1f547b83040904568f181a39ebe6d7e29dd438@54.183.6.67:26656"
```

</details>

#### Option 2: an Observer node is connected to another organization's public nodes

This option can be used if you have a trusted relationship with some organization and that organization
provided you with access to its nodes.   

[`config.toml`]

  ```toml
  [p2p]
  pex = true
  persistent_peers = "<node1-ID>@<node1-IP>:26656,..." # See the comment below on what values should be set here
  addr_book_strict = false
  ```

[`app.toml`]

  ```toml
  [api]
  enable = true
  ```

`persistent_peers` value:
  - another organization nodes with public IPs that this organization shared with you. 

    
#### Option 3: an Observer node is connected to my organization nodes

This is the option when you have a VN and want to create an ON connected to it.

  [`config.toml`]

  ```toml
  [p2p]
  pex = true
  persistent_peers = "<node1-ID>@<node1-IP>:26656,..." # See the comment below on what values should be set here
  addr_book_strict = false
  ```

  [`app.toml`]

  ```toml
  [api]
  enable = true
  ```
  
  `persistent_peers` value:
   - If your VN doesn't use any Private Sentry nodes, then it must point to the `Validator` node with private IP.
   - Otherwise, it must point to the Private Sentry nodes with private IPs.
   - Use the following command to get `node-ID` of a node: `./dcld tendermint show-node-id`.

