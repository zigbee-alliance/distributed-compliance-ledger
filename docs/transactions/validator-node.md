# Validator Node Module

<!-- markdownlint-disable MD036 -->

## VALIDATOR_NODE

### ADD_VALIDATOR_NODE

**Status: Implemented**

Adds a new Validator node.

- Parameters:
  - pubkey: `string` - The validator's Protobuf JSON encoded public key
  - moniker: `string` - The validator's human-readable name
  - identity: `optional(string)` - identity signature (ex. UPort or Keybase)
  - website: `optional(string)` - The validator's site link
  - details: `optional(string)` - The validator's details
  - ip: `optional(string)` - The node's public IP
  - node-id: `optional(string)` - The node's ID
- In State: `validator/Validator/value/<owner-address>`
- Who can send:
  - NodeAdmin
- CLI command:
  - `dcld tx validator add-node --pubkey='<protobuf JSON encoded>' --moniker=<string> --from=<account>`

### DISABLE_VALIDATOR_NODE

**Status: Implemented**

Disables the Validator node (removes from the validator set) by the owner.

- Who can send:
  - NodeAdmin; owner
- Parameters: No
- CLI command:
  - `dcld tx validator disable-node --from=<account>`

### PROPOSE_DISABLE_VALIDATOR_NODE

**Status: Implemented**

Proposes disabling of the Validator node from the validator set by a Trustee.

If more than 1 Trustee signature is required to disable a node, the disable
will be in a pending state until sufficient number of approvals is received.

- Parameters:
  - address: `string` - Bech32 encoded validator address or owner account
  - info: `optional(string)` - information/notes for the proposal
- Who can send:
  - Trustee
- CLI command:
  - `dcld tx validator propose-disable-node --address=<validator address> --from=<account>`
   e.g.:

    ```bash
    dcld query validator propose-disable-node --address=cosmos1nlt926tzc280ntkdmqvqumgrnvym8xc5wqwg3q --from alice
    ```

> **_Note:_** You can get Validator's address or owner address using query [GET_VALIDATOR](#get_validator)

### APPROVE_DISABLE_VALIDATOR_NODE

**Status: Implemented**

Approves disabling of the Validator node by a Trustee. It also can be used for revote (i.e. change vote from reject to approve)

The validator node is not disabled until sufficient number of Trustees approve it.

- Parameters:
  - address: `string` - Bech32 encoded validator address or owner account
  - info: `optional(string)` - information/notes for the approval
- Who can send:
  - Trustee
- Number of required approvals:
  - greater than 2/3 of Trustees (proposal by a Trustee is also counted as an approval)
- CLI command:
  - `dcld tx validator approve-disable-node --address=<validator address> --from=<account>`
   e.g.:

    ```bash
    dcld tx validator approve-disable-node --address=cosmos1nlt926tzc280ntkdmqvqumgrnvym8xc5wqwg3q from alice
    ```

> **_Note:_** You can get Validator's address or owner address using query [GET_VALIDATOR](#get_validator)

### REJECT_DISABLE_VALIDATOR_NODE

**Status: Implemented**

Rejects disabling of the Validator node by a Trustee. It also can be used for revote (i.e. change vote from approve to reject)

If disable validator proposal has only proposer's approval and no rejects then proposer can send this transaction to remove the proposal

The validator node is not reject until sufficient number of Trustees rejects it.

- Parameters:
  - address: `string` - Bech32 encoded validator address
  - info: `optional(string)` - information/notes for the reject
- Who can send:
  - Trustee
- Number of required rejects:
  - more than 1/3 of Trustees
- CLI command:
  - `dcld tx validator reject-disable-node --address=<validator address> --from=<account>`
   e.g.:

  ```bash
  dcld tx validator reject-disable-node --address=cosmos1nlt926tzc280ntkdmqvqumgrnvym8xc5wqwg3q --from alice
  ```

> **_Note:_** You can get Validator's address or owner address using query [GET_VALIDATOR](#get_validator)

### ENABLE_VALIDATOR_NODE

**Status: Implemented**

Enables the Validator node (returns to the validator set) by the owner.

the node will be enabled and returned to the active validator set.

- Who can send:
  - NodeAdmin; owner
- Parameters: No
- CLI command:
  - `dcld tx validator enable-node --from=<account>`

### GET_VALIDATOR

**Status: Implemented**

Gets a validator node.

- Parameters:
  - address: `string` - Bech32 encoded validator address or owner account
- CLI command:
  - `dcld query validator node --address=<validator address|account>`  e.g.:

    ```bash
    dcld query validator node --address=cosmosvaloper1qse069r3w0d82dul4xluqapxfg62qlndsdw9ms
    ```

    or

    ```bash
    dcld query validator node --address=cosmos1nlt926tzc280ntkdmqvqumgrnvym8xc5wqwg3q
    ```

- REST API:
  - GET `/dcl/validator/nodes/{owner}`

### GET_ALL_VALIDATORS

**Status: Implemented**

Gets the list of all validator nodes from the store.

> **_Note:_**  All stored validator nodes (`active` and `jailed`) will be returned by default.
In order to get an active validator set use specific command [validator set](#validator-set).

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query validator all-nodes`
- REST API:
  - GET `/dcl/validator/nodes`

### GET_PROPOSED_DISABLE_VALIDATOR

**Status: Implemented**

Gets a proposed validator node.

- Parameters:
  - address: `string` - Bech32 encoded validator address or owner account
- CLI command:
  - `dcld query validator proposed-disable-node --address=<validator address|account>`  e.g.:

    ```bash
    dcld query validator proposed-disable-node --address=cosmosvaloper1qse069r3w0d82dul4xluqapxfg62qlndsdw9ms
    ```

    or

    ```bash
    dcld query validator proposed-disable-node --address=cosmos1nlt926tzc280ntkdmqvqumgrnvym8xc5wqwg3q
    ```

- REST API:
  - GET `/dcl/validator/proposed-disable-nodes/{address}`

### GET_ALL_PROPOSED_DISABLE_VALIDATORS

**Status: Implemented**

Gets the list of all proposed disable validator nodes from the store.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query validator all-proposed-disable-nodes`
- REST API:
  - GET `/dcl/validator/proposed-disable-nodes`

### GET_REJECTED_DISABLE_VALIDATOR

**Status: Implemented**

Gets a rejected validator node.

- Parameters:
  - address: `string` - Bech32 encoded validator address or owner account
- CLI command:
  - `dcld query validator rejected-disable-node --address=<validator address|account>`  e.g.:

    ```bash
    dcld query validator rejected-disable-node --address=cosmosvaloper1qse069r3w0d82dul4xluqapxfg62qlndsdw9ms
    ```

    or

    ```bash
    dcld query validator rejected-disable-node --address=cosmos1nlt926tzc280ntkdmqvqumgrnvym8xc5wqwg3q
    ```

- REST API:
  - GET `/dcl/validator/rejected-disable-nodes/{address}`

### GET_ALL_REJECTED_DISABLE_VALIDATORS

**Status: Implemented**

Gets the list of all rejected disable validator nodes from the store.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query validator all-rejected-disable-nodes`
- REST API:
  - GET `/dcl/validator/rejected-disable-nodes`

### GET_DISABLED_VALIDATOR

**Status: Implemented**

Gets a disabled validator node.

- Parameters:
  - address: `string` - Bech32 encoded validator address or owner account
- CLI command:
  - `dcld query validator disabled-node --address=<validator address|account>`
   e.g.:

    ```bash
    dcld query validator disabled-node --address=cosmosvaloper1qse069r3w0d82dul4xluqapxfg62qlndsdw9ms
    ```

    or

    ```bash
    dcld query validator disabled-node --address=cosmos1nlt926tzc280ntkdmqvqumgrnvym8xc5wqwg3q
    ```

- REST API:
  - GET `/dcl/validator/disabled-nodes/{address}`

### GET_ALL_DISABLED_VALIDATORS

**Status: Implemented**

Gets the list of all disabled validator nodes from the store.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query validator all-disabled-nodes`
- REST API:
  - GET `/dcl/validator/disabled-nodes`

### GET_LAST_VALIDATOR_POWER

**Status: Implemented**

Gets a last validator node power.

- Parameters:
  - address: `string` - Bech32 encoded validator address or owner account
- CLI command:
  - `dcld query validator last-power --address=<validator address|account>`
   e.g.:

    ```bash
    dcld query validator last-power --address=cosmosvaloper1qse069r3w0d82dul4xluqapxfg62qlndsdw9ms
    ```

    or

    ```bash
    dcld query validator last-power --address=cosmos1nlt926tzc280ntkdmqvqumgrnvym8xc5wqwg3q
    ```

- REST API:
  - GET `/dcl/validator/last-powers/{owner}`

### GET_ALL_LAST_VALIDATORS_POWER

**Status: Implemented**

Gets the list of all last validator nodes power from the store.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query validator all-last-powers`
- REST API:
  - GET `/dcl/validator/last-powers`

### UPDATE_VALIDATOR_NODE

**Status: Not Implemented**

Updates the Validator node by the owner.  
`address` is used to reference the node, but can not be changed.

- Parameters:
  - address: `string` - Bech32 encoded validator address or owner account
  - moniker: `string` - The validator's human-readable name
  - identity: `optional(string)` - identity signature (ex. UPort or Keybase)
  - website: `optional(string)` - The validator's site link
  - details: `optional(string)` - The validator's details
  - ip: `optional(string)` - The node's public IP
  - node-id: `optional(string)` - The node's ID
- Who can send:
  - NodeAdmin; owner

<!-- markdownlint-enable MD036 -->
