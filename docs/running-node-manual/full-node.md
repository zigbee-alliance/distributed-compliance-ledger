# Running a Full Node manually
<!-- markdownlint-disable MD033 -->

## Prerequisites

Make sure you have all [prerequisites](./prerequisites.md) set up.

## Deployment steps

### 1. Put `cosmovisor` binary to `/usr/bin/`, set proper owner and execution permissions

### 2. Locate the genesis app version to genesis application version directory

- Create `$HOME/.dcl/cosmovisor/genesis/bin` directory.
- Copy `dcld` binary to it, set proper owner and execution permissions.
    Please note that execution permissions on `dcld` should be granted to all (i.e. User, Group and Others classes)
    because cosmovisor requires execution permission on the application binary to be granted to Others class.

### 3. Configure CLI

- `./dcld config chain-id <chain-id>`
  - Use `testnet-2.0` for `<chain-id>` if you want to connect to the persistent Test Net
  - Use `main-net` for `<chain-id>` if you want to connect to the persistent Main Net
- `./dcld config output json` - Output format (text/json).

### 4. Initilize the node

```bash
./dcld init "<node-name>" --chain-id "<chain-id>"
```

- Use `testnet-2.0` for `<chain-id>` if you want to connect to the persistent Test Net
- Use `main-net` for `<chain-id>` if you want to connect to the persistent Main Net

### 5. (Optional) Enable `state sync` in the configuration if you are joining long-running network

[`$HOME/.dcl/config/config.toml`]

```toml
[statesync]
enable = true

rpc_servers = "http(s)://<host>:<port>,http(s)://<host>:<port>"
trust_height = <trust-height>
trust_hash = "<trust-hash>"
trust_period = "168h0m0s"
```

<details>
<summary>Example for Testnet 2.0 (clickable) </summary>

```toml
[statesync]
enable = true
rpc_servers = "https://on.test-net.dcl.csa-iot.org:26657,https://on.test-net.dcl.csa-iot.org:26657"
```

</details>

<details>
<summary>Example for Mainnet (clickable) </summary>

```toml
[statesync]
enable = true
rpc_servers = "https://on.dcl.csa-iot.org:26657,https://on.dcl.csa-iot.org:26657"
```

</details>

> **_NOTE:_**  You should provide at least 2 addresses for `rpc_servers`. It can be 2 identical addresses

You can use the following command to obtain `<trust-height>` and `<trust-hash>` of your network

```bash
curl -s http(s)://<host>:<port>/commit | jq "{height: .result.signed_header.header.height, hash: .result.signed_header.commit.block_id.hash}"
```

<details>
<summary>Example for Testnet 2.0 (clickable) </summary>

```bash
curl -s https://on.test-net.dcl.csa-iot.org:26657/commit | jq "{height: .result.signed_header.header.height, hash: .result.signed_header.commit.block_id.hash}"
```

</details>

<details>
<summary>Example for Mainnet (clickable) </summary>

```bash
curl -s https://on.dcl.csa-iot.org:26657/commit | jq "{height: .result.signed_header.header.height, hash: .result.signed_header.commit.block_id.hash}"
```

</details>

- `<host>` - RPC endpoint host of the network being joined
- `<port>` - RPC endpoint port of the network being joined

> **_NOTE:_** State sync is not attempted if the node has any local state (LastBlockHeight > 0)

### 6. (Optional) Enable `state sync` snapshots in `[~/.dcl/config/app.toml]` file

```toml
[state-sync]
snapshot-interval = "snapshot-interval"
snapshot-keep-recent = "snapshot-keep-recent"
```

### 7. Get genesis and persistent peers files

- Put `genesis.json` into dcld's config directory (usually `$HOME/.dcl/config/`).
  - Use `deployment/persistent_chains/testnet-2.0/genesis.json` if you want to connect to the persistent Testnet 2.0
  - Use `deployment/persistent_chains/main-net/genesis.json` if you want to connect to the persistent Mainnet
- Create an empty `persistent_peers.txt` in the current path because this file is required by `run_dcl_node` script
    ```bash
    touch persistent_peers.txt
    ```

### *** Step 8 can be automated using `run_dcl_node` script

Run node:

```bash
./run_dcl_node -t "<node-type>" -c "<chain-id>" "<node-name>"
```

- `<node-type>` - one of the following types depending on which type of node is being run:
  - genesis
  - validator
  - observer
  - private-sentry
  - public-sentry
  - seed

> Notes:
>
> - the script assumes that:
>   - current user is going to be used for `cosmovisor` service to run as
>   - current user is in sudoers list

### 8. Run node

- Open `$HOME/.dcl/config/config.toml` file in your favorite text editor:
  - Make your node public:
    - Open `$HOME/.dcl/config/config.toml`
    - Find the line under `# TCP or UNIX socket address for the RPC server to listen on`
    - Change it to: `laddr = "tcp://0.0.0.0:26657"`
  - Change other settings (see specific instructions for every Node type)
- Open `26656` (p2p) and `26657` (RPC) ports.
  - `sudo ufw allow 26656/tcp`
  - `sudo ufw allow 26657/tcp`
- Edit `cosmovisor.service`
  - Replace `ubuntu` with a username you want to start service on behalf
- Copy service configuration.
  - `cp cosmovisor.service /etc/systemd/system/`
- Enable the service: `sudo systemctl enable cosmovisor`
- Start node: `sudo systemctl start cosmovisor`
- For testing purpose the node process can be started directly: `./dcld start` (instead of two previous `systemctl` commands using `cosmovisor` service).
Service mode is recommended for demo and production environment.

- Use `systemctl status cosmovisor` to get the node service status.
- Use `journalctl -u cosmovisor.service -f` to see node logs.
- You can also check node status by executing the command `./dcld status` to get the current status.
    The value of `latest_block_height` reflects the current node height.

- Add the following line to the end of `$HOME/.profile` file:
  - `export PATH=$PATH:$HOME/.dcl/cosmovisor/current/bin`
- Execute the following command to apply the updated `$PATH` immediately:
  - `source $HOME/.profile`

### 9. Check the node is running and getting all the transactions

- Get the node status: `dcld status --node tcp://localhost:26657`.
- Make sure that `result.sync_info.latest_block_height` is increasing over the time (once in about 10 mins). When you see the `catching_up` as `true` that signifies that the node is still downloading all the transactions. Once it has fully synced this will value will turn to `false`
