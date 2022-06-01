# Running a Seed Node manually

## Prerequisites
Make sure you have all [prerequisites](./prerequisites.md) set up
## Deployment steps

#### 1. Put `cosmovisor` binary to `/usr/bin/`, set proper owner and execution permissions.

#### 2. Locate the genesis app version to genesis application version directory:
- Create `$HOME/.dcl/cosmovisor/genesis/bin` directory.
- Copy `dcld` binary to it, set proper owner and execution permissions.
    Please note that execution permissions on `dcld` should be granted to all (i.e. User, Group and Others classes)
    because cosmovisor requires execution permission on the application binary to be granted to Others class.

#### 3. Configure CLI:
- `./dcld config chain-id <chain-id>`
    - Use `testnet-2.0` for `<chain-id>` if you want to connect to the persistent Test Net
- `./dcld config output json` - Output format (text/json).

#### 4. Initilize the node:

```bash
./dcld init "<node-name>" --chain-id "<chain-id>"
```
- Use `testnet-2.0` for `<chain-id>` if you want to connect to the persistent Test Net

#### 5. Enable seed mode in config `$HOME/.dcl/config/config.toml`:
```toml
[p2p]
seed_mode=true
```

#### 6. (Optional) Consider enabling `state sync` in the configuration if you are joining long-running network
- For more information refer to [running-node-in-existing-network.md](../advanced/running-node-in-existing-network.md)

#### *** Step 7 can be automated using `run_dcl_node` script:
Run node:

```bash
./run_dcl_node -t observer -c "<chain-id>" "<node-name>"
```

> Notes:
>
> * the script assumes that:
>   * current user is going to be used for `cosmovisor` service to run as
>   * current user is in sudoers list

#### 7. Run node:
- Put `genesis.json` into dcld's config directory (usually `$HOME/.dcl/config/`).
    - Use `deployment/persistent_chains/testnet/genesis.json` if you want to connect to the persistent Test Net
- Open `$HOME/.dcl/config/config.toml` file in your favorite text editor:
    - Tell node how to connect to the network:
        - Set the value for the `persistent_peers` field as the content of `persistent_peers.txt` file.
        - For `testnet-2.0` get the latest `persistent_peers` string from the CSA slack channel
    - Make your node public:
        - Open `$HOME/.dcl/config/config.toml`
        - Find the line under `# TCP or UNIX socket address for the RPC server to listen on`
        - Change it to: `laddr = "tcp://0.0.0.0:26657"`
    - Optionally change other setting.
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
#### 8. Check the seed node is running and getting all the transactions:
- Get the node status: `dcld status --node tcp://localhost:26657`.
- Make sure that `result.sync_info.latest_block_height` is increasing over the time (once in about 10 mins). When you see the `catching_up` as `true` that signifies that the node is still downloading all the transactions. Once it has fully synced this will value will turn to `false`

#### 9. Congrats! You are now running an observer node.