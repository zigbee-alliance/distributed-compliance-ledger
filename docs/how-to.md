# How To

This document contains tutorials demonstrating how to accomplish common tasks using CLI.

- Please configure the CLI before using (see [CLI Configuration](#cli-configuration) section).
- If write requests to the Ledger needs to be sent, please make sure that you have
an Account created on the Ledger with an appropriate role (see [Getting Account](#getting-account) section in [how-to.md](how-to.md)).
- Sending read requests to the Ledger doesn't require an Account (Ledger is public for reads).
- A full list of all CLI commands can be found there: [cli-help](cli-help.md).
- After the CLI is configured and Account with an appropriate role is created,
the following instructions can be used for every role (see [Use Case Diagrams](use_cases)):
    - [Trustee](#trustee-instructions) 
        - create new accounts
        - assign roles to the account
        - revoke roles from the account
        - approve X509 root certificates
        - publish X509 certificates
    - [CA](#ca-instructions)
        - propose X509 root certificates
        - publish X509 certificates        
    - [Vendor](#vendor-instructions) 
        - publish device model info
        - publish X509 certificates
    - [Test House](#test-house-instructions) 
        - publish compliance test results
        - publish X509 certificates
    - [ZB Certification Center](#certification-center-instructions)
        - certify or revoke certification of device models
        - publish X509 certificates
    - [Node Admin](#node-admin-instructions-setting-up-a-new-validator-node) 
        - add a new Validator node
        - publish X509 certificates

## CLI Configuration

CLI configuration file can be created or updated by executing of the command: `zblcli config <key> [value]`.
Here is the list of supported settings:
* chain-id <chain id> - Chain ID of pool node
* output <type> - Output format (text/json)
* indent <bool> - Add indent to JSON response
* trust-node <bool> - Trust connected full node (don't verify proofs for responses). The `false` value is recommended.
* node <node-ip> - Address `<host>:<port>` of the node to connect. 
* trace <bool> - Print out full stack trace on errors.
* broadcast-mode <mode> - Write transaction broadcast mode to use (one of: `sync`, `async`, `block`. `block` is default).

In order to connect the CLI to the ZB Ledger Demo Pool, the following parameters should be used:

* `zblcli config chain-id zblchain`
* `zblcli config output json` - Output format (text/json).
* `zblcli config indent true` - Add indent to JSON response.
* `zblcli config trust-node false` - Verify proofs for node responses.
* `zblcli config node <ip address>` - Address of a node to connect. 
Choose one of the listed in `persistent_peers.txt` file. 
Example: `tcp://18.157.114.34:26657`.


## Getting Account
Ledger is public for read which means that anyone can read from the Ledger without a need to have 
an Account but it is private for write. 
In order to send write transactions to the ledger you need: 

  - Have a private/public key pair.
  - Have an Account created on the ledger via `ACCOUNT` transaction (see [Use Case Txn Auth](use_cases/use_cases_txn_auth.puml)).
      - The Account stores the public part of the key
      - The Account has an associated role. The role is used for authorization policies.
  - Sign every transaction by the private key.

Here is steps for getting an account:
* Generate keys and local account: `zblcli keys add <name>`.
* Share generated `address` and `pubkey` with a number of `Trustee`s sufficient for account addition operation.
* One of `Trustee`s proposes to add the account to the ledger: `zblcli tx auth propose-add-account --address=<account address> --pubkey=<account pubkey> --roles=<role1,role2,...> --from=<account>`
* Sufficient number of other `Trustee`s approve the proposed account: `zblcli tx auth approve-add-account --address=<account address> --from=<account>`
* Check that the active account exists: `zblcli query auth account --address=<account address>`

Example:
* `zblcli keys add steve`
* `zblcli tx auth propose-add-account --address=cosmos1sug8cquqnn5jddkqt4ud6hcr290sn4wh96x5tv --pubkey=cosmospub1addwnpepqvnfd2f99vew4t7phe3mqprmceq3jgavm0rguef3gkv8z8jd6lg25egq6d5 --roles=Vendor,NodeAdmin --from jack`
* `zblcli tx auth approve-add-account --address=cosmos1sug8cquqnn5jddkqt4ud6hcr290sn4wh96x5tv --from alice`
* `zblcli query auth account --address=cosmos1sug8cquqnn5jddkqt4ud6hcr290sn4wh96x5tv`

## Trustee Instructions

Account creation consists of two parts. One of the trustees should propose an account by posting `propose-add-account` transaction.
After this account goes to the proposed accounts set. Then this account must be approved by a majority of trustees (2/3+ including the trustee who has proposed the account).
Once approved the account can be used to send transactions. See [use_case_txn_auth](use_cases/use_cases_txn_auth.png).

##### 1. Create an Account proposal for the user
  Command: `zblcli tx auth propose-add-account --address=<string> --pubkey=<string> --roles=<roles> --from=<account>`

  Flags:
  - address: `string` - bench32 encoded account address
  - pubkey: `string` - bench32 encoded public key
  - roles: `optional(string)` - comma-separated list of roles (supported roles: Vendor, TestHouse, ZBCertificationCenter, Trustee, NodeAdmin)
  - from: `string` - name or address of private key with which to sign

  Example: `zblcli tx auth propose-add-account --address=cosmos15ljvz60tfekhstz8lcyy0c9l8dys5qa2nnx4d7 --pubkey=cosmospub1addwnpepqtrnrp93hswlsrzvltc3n8z7hjg9dxuh3n4rkp2w2verwfr8yg27c95l4k3 --roles=Vendor,NodeAdmin --from=jack`

##### 2. Approve proposed Account
  Command: `zblcli tx auth approve-add-account --address=<string> --from=<account>`

  Flags:
  - address: `string` - bench32 encoded account address to approve
  - from: `string` - name or address of private key with which to sign

  Example: `zblcli tx auth approve-add-account --address=cosmos15ljvz60tfekhstz8lcyy0c9l8dys5qa2nnx4d7 --from=jack`

##### 3. Approve proposed X509 root certificate  
  Command: `zblcli tx pki approve-add-x509-root-cert --subject=<string> --subject-key-id=<hex string> --from=<account>`

  Flags:
   - subject: `string` - certificates's `Subject`.
   - subject-key-id: `string` - certificates's `Subject Key ID` (hex-encoded uppercase string).
   - from: `string` - Name or address of private key with which to sign.

  Example: `zblcli tx pki approve-add-x509-root-cert --subject="CN=dsr-corporation.com" --subject-key-id="8A:E9:AC:D4:16:81:2F:87:66:8E:61:BE:A9:C5:1C:0:1B:F7:BB:AE" --from=jack`
 
## CA instructions
Currently any role can propose an X509 root certificate, or publish 
(intermediate or leaf) X509 certificates. 

##### 1. Propose a new self-signed root certificate
  Command: `zblcli tx pki propose-add-x509-root-cert --certificate=<string-or-path> --from=<account>`

  Flags:
   - certificate: `string` - PEM encoded certificate (string or path to file containing data).
   - from: `string` - Name or address of private key with which to sign.

  Example: `zblcli tx pki propose-add-x509-root-cert --certificate="/path/to/certificate/file" --from=jack`
  
  Example: `zblcli tx pki propose-add-x509-root-cert --certificate="----BEGIN CERTIFICATE----- ......" --from=jack`

##### 2. Publish an intermediate or leaf X509 certificate
 The certificate must be signed by a chain of certificates which must be already present on the ledger.
 
 Command: `zblcli tx pki add-x509-cert --certificate=<string-or-path> --from=<account>`

  Flags:
   - certificate: `string` - PEM encoded certificate (string or path to file containing data).
   - from: `string` - Name or address of private key with which to sign.

  Example: `zblcli tx pki add-x509-cert --certificate="/path/to/certificate/file" --from=jack`
  
  Example: `zblcli tx pki add-x509-cert --certificate="----BEGIN CERTIFICATE----- ......" --from=jack`  
    

## Vendor Instructions

##### 1. Publish an intermediate or leaf X509 certificate(s) to be used for signing X509 Certificates for every Device
The certificate must be signed by a chain of certificates which must be already present on the ledger.
 
 Command: `zblcli tx pki add-x509-cert --certificate=<string-or-path> --from=<account>`

  Flags:
   - certificate: `string` - PEM encoded certificate (string or path to file containing data).
   - from: `string` - Name or address of private key with which to sign.

  Example: `zblcli tx pki add-x509-cert --certificate="/path/to/certificate/file" --from=jack`
  
  Example: `zblcli tx pki add-x509-cert --certificate="----BEGIN CERTIFICATE----- ......" --from=jack`  
    
##### 2. Add a new model info with the given VID/PID

  Command: `zblcli tx modelinfo add-model --vid=<uint16> --pid=<uint16> --name=<string> --description=<string or path> --sku=<string> 
--firmware-version=<string> --hardware-version=<string> --tis-or-trp-testing-completed=<bool> --from=<account>`

  Flags:
  - vid: `uint16` -  model vendor ID
  - pid: `uint16` -  model product ID
  - name: `string` -  model name
  - description: `string` -  model description (string or path to file containing data)
  - sku: `string` -  stock keeping unit
  - firmware-version: `string` -  version of model firmware
  - hardware-version: `string` -  version of model hardware
  - hardware-version: `string` -  version of model hardware
  - tis-or-trp-testing-completed: `bool` -  whether model has successfully completed TIS/TRP testing
  - from: `string` - Name or address of private key with which to sign
  - cid: `optional(uint16)` - model category ID
  - custom: `optional(string)` - custom information (string or path to file containing data)

  Example: `zblcli tx modelinfo add-model --vid=1 --pid=1 --name="Device #1" --description="Device Description" --sku="SKU12FS" --firmware-version="1.0" --hardware-version="2.0" --tis-or-trp-testing-completed=true --from=jack`
  
  Example: `zblcli tx modelinfo add-model --vid=1 --pid=1 --name="Device #1" --description="Device Description" --sku="SKU12FS" --firmware-version="1.0" --hardware-version="2.0" --tis-or-trp-testing-completed=true --from=jack --cid=1 --custom="Some Custom information"`


## Test House Instructions
##### 1A. Publish an intermediate or leaf X509 certificate(s) to be used for signing the Certification blob
This step is needed for off-ledger certification use case only, see [use_cases_device_off_ledger_certification](use_cases/use_cases_device_off_ledger_certification.png).

The certificate must be signed by a chain of certificates which must be already present on the ledger.
 
 Command: `zblcli tx pki add-x509-cert --certificate=<string-or-path> --from=<account>`

  Flags:
   - certificate: `string` - PEM encoded certificate (string or path to file containing data).
   - from: `string` - Name or address of private key with which to sign.

  Example: `zblcli tx pki add-x509-cert --certificate="/path/to/certificate/file" --from=jack`
  
  Example: `zblcli tx pki add-x509-cert --certificate="----BEGIN CERTIFICATE----- ......" --from=jack`  
    
##### 1B. Add a new testing result for the device model with the given VID/PID
This step is needed for on-ledger certification use case only, see [use_cases_device_on_ledger_certification](use_cases/use_cases_device_on_ledger_certification.png).

 The corresponding model must present on the ledger.
 
 Command: ` zblcli tx compliancetest add-test-result --vid=<uint16> --pid=<uint16> --test-result=<string> --test-date=<rfc3339 encoded date> --from=<account>`

  Flags:
  - vid: `uint16` -  model vendor ID
  - pid: `uint16` -  model product ID
  - test-result: `string` -  test result (string or path to file containing data)
  - test-date: `string` -  date of test result (rfc3339 encoded)
  - from: `string` - Name or address of private key with which to sign

  Example: `zblcli tx compliancetest add-test-result --vid=1 --pid=1 --test-result="Test Document" --test-date="2020-04-16T06:04:57.05Z" --from=jack`
  
  Example: `zblcli tx compliancetest add-test-result --vid=1 --pid=1 --test-result="path/to/document" --test-date="2020-04-16T06:04:57.05Z" --from=jack`
  
   
## Certification Center Instructions
##### 1A. Publish an intermediate or leaf X509 certificate(s) to be used for signing the Certification blob
This step is needed for off-ledger certification use case only, see [use_cases_device_off_ledger_certification](use_cases/use_cases_device_off_ledger_certification.png).

The certificate must be signed by a chain of certificates which must be already present on the ledger.
 
 Command: `zblcli tx pki add-x509-cert --certificate=<string-or-path> --from=<account>`

  Flags:
   - certificate: `string` - PEM encoded certificate (string or path to file containing data).
   - from: `string` - Name or address of private key with which to sign.

  Example: `zblcli tx pki add-x509-cert --certificate="/path/to/certificate/file" --from=jack`
  
  Example: `zblcli tx pki add-x509-cert --certificate="----BEGIN CERTIFICATE----- ......" --from=jack`  
    
##### 1B. Certify the device model with the given VID/PID
This step is needed for on-ledger certification use case only, see [use_cases_device_on_ledger_certification](use_cases/use_cases_device_on_ledger_certification.png).

The corresponding model and the test results must present on the ledger.

 Command: `zblcli tx compliance certify-model --vid=<uint16> --pid=<uint16> --certification-type=<zb> --certification-date=<rfc3339 encoded date> --from=<account>`
  
  Flags:
  - vid: `uint16` -  model vendor ID
  - pid: `uint16` -  model product ID
  - certification-type: `string` -  certification type (zb` is the only supported value now)
  - certification-date: `string` -  the date of model certification (rfc3339 encoded)
  - from: `string` - name or address of private key with which to sign
  - reason: `optional(string)` -  an optional comment describing the reason of certification

  Example: `zblcli tx compliance certify-model --vid=1 --pid=1 --certification-type="zb" --certification-date="2020-04-16T06:04:57.05Z" --from=jack`

##### 2. Revoke certification for the device model with the given VID/PID
This step can be used in either on-ledger certification use case
 ([use_cases_device_on_ledger_certification](use_cases/use_cases_device_on_ledger_certification.png))
  or off-ledger certification use case ([use_cases_device_off_ledger_certification](use_cases/use_cases_device_off_ledger_certification.png)).
 
  Command: ` zblcli tx compliance revoke-model --vid=<uint16> --pid=<uint16> --certification-type=<zb> --revocation-date=<rfc3339 encoded date> --from=<account>`

  Flags:
  - vid: `uint16` -  model vendor ID
  - pid: `uint16` -  model product ID
  - certification-type: `string` -  certification type (zb` is the only supported value now)
  - revocation-date: `string` -  the date of model revocation (rfc3339 encoded)
  - from: `string` - name or address of private key with which to sign
  - reason: `optional(string)` -  an optional comment describing the reason of revocation

  Example: `zblcli tx compliance revoke-model --vid=1 --pid=1 --certification-type="zb" --revocation-date="2020-04-16T06:04:57.05Z" --from=jack`
  
  Example: `zblcli tx compliance revoke-model --vid=1 --pid=1 --certification-type="zb" --revocation-date="2020-04-16T06:04:57.05Z" --reason "Some Reason" --from=jack`
  
 
## Node Admin Instructions (Setting up a new Validator Node)

Validators are responsible for committing of new blocks to the ledger.
Here are steps for setting up a new validator node.

A more detailed instruction on how to add a validator node to the current ZB Ledger pool
can be found here: [running-node.md](running-node.md).

* Firstly you have to posses an account with `NodeAdmin` role on the ledger. See [Getting Account](#getting-account):
    
* Initialize the node and create the necessary config files:
    * Init Node: `zbld init <node name> --chain-id <chain id>`.
    * Fetch the network `genesis.json` file and put it into zbld's config directory (usually `$HOME/.zbld/config/`).
    * In order to join network your node needs to know how to find alive peers. 
    Update `persistent_peers` field of `$HOME/.zbld/config/config.toml` file to contain peers info in the format:
    `<node1 id>@<node1 listen_addr>,<node2 id>@<node2 listen_addr>,.....`
    * Open `26656` (p2p) and `26657` (RPC) ports.

* Add validator node to the network:
    * Get this node's tendermint validator *consensus address*: `zbld tendermint show-address`
    * Get this node's tendermint validator *consensus pubkey*: `zbld tendermint show-validator`
    * Note that *consensus address* and *consensus pubkey* are not the same as `address` and `pubkey` were used for account creation.
    * Start node: `zbld start`
    * In the output, you can notice that `height` increases quickly over time. 
    This means that the node is in updating to the latest network state (it takes some time).
    
        You can also check node status by connecting CLI to your local node `zblcli config node tcp://0.0.0.0:26657`
        and executing the command `zblcli status` to get the current status.
        The `true` value for `catching_up` field means that the node is in the updating process.
        The value of `latest_block_height` reflects the current node height.

    * Wait until the value of `catching_up` field gets to `false` value.
    * Add validator node: `zblcli tx validator add-node --validator-address=<validator address> --validator-pubkey=<validator pubkey> --name=<node name> --from=<name>`

* Congrats! You are an owner of the validator node.

* Check node is alive and participate in consensus:
    * Get the list of all nodes: `zblcli query validator all-nodes`. 
    The node must present in the list and has the following params: `power:10` and `jailed:false`.
    * Get the list of nodes participating in the consensus for the last block: `zblcli tendermint-validator-set`
        * You can pass the additional value to get the result for a specific height: `zblcli tendermint-validator-set 100`  
    * Get the node status: `zblcli status --node <node ip>`

Example:
* `zbld init node-name --chain-id zblchain`
* `cp /source/genesis.json $HOME/.zbld/config/`
* `sed -i "s/persistent_peers = \"\"/<node id>@<node ip>,<node2 id>@<node2 ip>/g" $HOME/.zbld/config/config.toml`
* `sudo ufw allow 26656/tcp`
* `sudo ufw allow 26657/tcp`
* `zbld start`
* `zbld status`
* `zblcli tx validator add-node --validator-address=$(zbld tendermint show-address) --validator-pubkey=$(zbld tendermint show-validator) --name=node-name --from=node_admin`
* `zblcli query validator all-nodes`

##### Policy
- Maximum number of nodes (`MaxNodes`): 100 
- Size (number of blocks) of the sliding window used to track validator liveness (`SignedBlocksWindow`): 100
- Minimal number of blocks must have been signed per window (`MinSignedPerWindow`): 50

Node will be jailed and removed from the active validator set in the following conditions are met:
- Node passed minimum height: `node start height + SignedBlocksWindow`
- Node exceeded the maximum number of unsigned blocks withing the window: `SignedBlocksWindow - MinSignedPerWindow`

Note that jailed node will not be removed from the main index to track validator nodes. 
So it is not possible to create a new node with the same `validator address`.
In order to unjail the node and return it to the active tendermint validator set the sufficient number of Trustee's approvals is needed 
(see authorization rules).
