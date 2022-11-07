# Running Node using Ansible (common part)

## Prerequisites

Make sure you have all [prerequisites](./prerequisites.md) set up

## Configure DCL network parameters (local machine)

### 1. Set network chain ID

[`deployment/ansible/inventory/group_vars/all.yml`]

```yaml
  chain_id: <chain-id>
  ...
```

Every network must have a unique chain ID (e.g. `test-net`, `main-net` etc.)

<details>
<summary>Example for Testnet 2.0 (clickable) </summary>

```yaml
  chain_id: testnet-2.0
```

</details>

<details>
<summary>Example for Mainnet (clickable) </summary>

```yaml
  chain_id: main-net
```

</details>

### 2. Set other parameters

[`deployment/ansible/inventory/group_vars/all.yml`]

```yaml
  company_name: <your-company-name>
  dcl_version: <latest-DCL-version>
  ...
```

### 3. Put `genesis.json` file under specific directory

- Get or download `genesis.json` file of a network your node will be joining and put it under the following path:

  ```text
  deployment/persistent_chains/<chain-id>/genesis.json
  ```

  where `<chain-id>` is the chain ID of a network spefied in the previous step.

  <details>
  <summary>Example for Testnet 2.0 (clickable) </summary>

  For `testnet-2.0` the genesis file is already in place. So you don't need to do anything!

  ```text
  deployment/persistent_chains/testnet-2.0/genesis.json
  ```

  </details>

  <details>
  <summary>Example for Mainnet (clickable) </summary>

  For `main-net` the genesis file is already in place. So you don't need to do anything!

  ```text
  deployment/persistent_chains/main-net/genesis.json
  ```

  </details>
