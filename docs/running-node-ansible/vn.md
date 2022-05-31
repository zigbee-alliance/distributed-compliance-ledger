# Running Validator Node using Ansible
## Prerequisites
Make sure you have all [prerequisites](./prerequisites.md) set up

## Configure DCL network parameters (local machine)
#### 1. Set network chain id
[`deployment/ansible/inventory/hosts.yml`]
```yaml
all:
  vars:
    chain_id: <chain-id>
  ...
```
Every network must have a unique chain ID (e.g. `test-net`, `main-net` etc.)

<details>
<summary>Example for Testnet 2.0</summary>

```yaml
all:
  vars:
    chain_id: testnet-2.0
  ...
```
</details>

#### 2. Put `genesis.json` file under specific directory
Get or download `genesis.json` file of a network your node will be joining and put it under the following path:
```
deployment/persistent_chains/<chain-id>/genesis.json
```
where `<chain-id>` is the chain id of a network spefied in the previous step
## Configure node type specific parameters (local machine)
#### 1. Specify target instance address in the inventory file
[`deployment/ansible/inventory/hosts.yml`]
```yaml
all:
  ...
  children:
    ...
    validators:
      hosts:
        <validator node IP address or hostname>
    ...
```

#### 2. Set persistent peers string in validator configuration
[`deployment/ansible/roles/configure/vars/validator.yml`]
```yaml
config:
  p2p:
    persistent_peers: "<node1-ID>@<node1-IP>:26656,..."
  ...
```

<details>
<summary>Example for Testnet 2.0</summary>

```yaml
config:
  p2p:
    persistent_peers: "8091122d82075eff187434f22055b193aa797835@52.21.51.189:26656,f0a652e3f08f0a9ddead545aae5233fbbaf82bd0@44.239.5.82:26656,cdff3160145059f8835e9489a9ea944b775a375e@3.11.5.22:26656,803d6f7d3489b618bc736de4914998b55c0a8240@35.224.92.247:26656,841db5c52e432a89b614d3413a639889b81394d0@34.231.53.112:26656,3029ef695af6a147d0701dd5c21154e60e299801@13.229.73.209:26656,1e35166c26761555dc63c95cef64895eff52b899@51.136.18.38:26656,52dce03ebd9a28b51d496a75d1aee690b009c72a@54.156.145.239:26656,81c2e4201bf8579f1db216dbb230d09e2dffb6b9@3.89.241.25:26656,f07ad3735ea1c883ea4c1a8d9d19f6a02e241ca5@13.209.123.210:26656"
```
</details>

#### 3. (Optional) If you are joining a long-running network, enable `statesync` or use one of the options in [running-node-in-existing-network.md](../advanced/running-node-in-existing-network.md)
[`deployment/ansible/roles/configure/vars/validator.yml`]

```yaml
config:
...
  statesync:
    enable: true
    rpc_servers: "http(s):<node1-IP>:26657, ..."
    trust_height: <trust-height>
    trust_hash: "<trust-hash>"
...
```
</details>

> **_NOTE:_**  You should provide at least 2 addresses for `rpc_servers`. It can be 2 identical addresses

You can use the following command to obtain `<trust-height>` and `<trust-hash>` of your network

```bash
curl -s http(s)://<host>:<port>/commit | jq "{height: .result.signed_header.header.height, hash: .result.signed_header.commit.block_id.hash}"
```
<details>
<summary>Example for Testnet 2.0</summary>

```bash
curl -s https://on.test-net.dcl.csa-iot.org:26657/commit | jq "{height: .result.signed_header.header.height, hash: .result.signed_header.commit.block_id.hash}"
```
</details>

- `<host>` - RPC endpoint host of the network being joined
- `<port>` - RPC endpoint port of the network being joined

> **_NOTE:_** State sync is not attempted if the node has any local state (LastBlockHeight > 0)

## Run ansible (local machine)
#### 1. Verify that all the configuration parameters from the previous section are correct
#### 2. Run ansible
```bash
ansible-playbook -i ./deployment/ansible/inventory  -u <target-host-ssh-user> ./deployment/ansible/deploy.yml
```
- `<target-host-ssh-username>` - target host ssh user
- Ansible provisioning can take several minutes depending on number of nodes being provisioned

## Deployment Verification (target machine)
#### 1. Switch to cosmovisor user
```
sudo su -s /bin/bash cosmovisor
```

#### 2. Query status
```
dcld status
```

## Make your node a Validator (target machine)

#### 1. Switch to cosmovisor user
```
sudo su -s /bin/bash cosmovisor
```

#### 2. Generate NodeAdmin keys

```bash
dcld keys add "<admin-account-name>" 2>&1 | tee "<admin-account-name>.dclkey.data"
```

IMPORTANT keep generated data (especially the mnemonic) securely.

#### 3. Share generated address and pubkey with the community

address and pubkey can be found using

```bash
dcld keys show --output text "<admin-account-name>"
```

#### 4. Wait until your NodeAdmin key is proposed and approved by the quorum of trustees

- Make sure the Node Admin account is proposed by a Trustee (usually CSA). The account will appear in the "Accounts" / "All Proposed Accounts" tab in https://testnet.iotledger.io/accounts.

- Make sure that the proposed account is approved by at least 2/3 of Trustees. The account must disappear from the "Accounts" / "All Proposed Accounts" tab in https://testnet.iotledger.io/accounts, and appear in  "Accounts" / "All Active Accounts" tab.

#### 5. Check the account presence on the ledger
`dcld query auth account --address="<address>"`


#### 6. Make the node a validator

```bash
dcld tx validator add-node --pubkey="<protobuf JSON encoded validator-pubkey>" --moniker="<node-name>" --from="<admin-account-name>"
```

> **_Note:_** Get `<protobuf JSON encoded validator-pubkey>` using `dcld tendermint show-validator` command

(once transaction is successfully written you should see "code": 0 in the JSON output.)

#### 7. Make sure the VN participates in consensus
`dcld query tendermint-validator-set` must contain the VN's address

>**_Note:_** Get your VN's address using `dcld tendermint show-address` command.