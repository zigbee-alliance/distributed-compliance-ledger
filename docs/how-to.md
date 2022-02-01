# How To

This document contains tutorials demonstrating how to accomplish common tasks using CLI.

- Please configure the CLI before using (see [CLI Configuration](#cli-configuration) section).
- If write requests to the Ledger needs to be sent, please make sure that you have
an Account created on the Ledger with an appropriate role (see [Getting Account](#getting-account) section in [how-to.md](how-to.md)).
- Sending read requests to the Ledger doesn't require an Account (Ledger is public for reads).
- A full list of all CLI commands with all parameters can be found there: [transactions](transactions.md).
- After the CLI is configured and Account with an appropriate role is created,
the following instructions can be used for every role (see [Use Case Diagrams](use_cases)):
    - [Trustee](#trustee-instructions) 
        - propose new accounts
        - approve new accounts
        - propose revocation of accounts
        - approve revocation of accounts
        - propose X509 root certificates
        - approve X509 root certificates
        - propose revocation of X509 root certificates
        - approve revocation of X509 root certificates
        - publish X509 certificates
        - revoke X509 certificates        
    - [CA](#ca-instructions)
        - propose X509 root certificates
        - publish X509 certificates     
        - revoke X509 certificates           
    - [Vendor](#vendor-instructions) 
        - publish device model info
        - publish X509 certificates
        - revoke X509 certificates        
    - [Test House](#test-house-instructions) 
        - publish compliance test results
        - publish X509 certificates
        - revoke X509 certificates        
    - [Certification Center](#certification-center-instructions)
        - certify or revoke certification of device models
        - publish X509 certificates
        - revoke X509 certificates        
    - [Node Admin](#node-admin-instructions-setting-up-a-new-validator-node) 
        - add a new Validator node
        - publish X509 certificates
        - revoke X509 certificates

## CLI Configuration

CLI configuration file can be created or updated by executing of the command: `dcld config <key> [value]`.
Here is the list of supported settings:
* chain-id <chain id> - unique chain ID of the network you are going to connect to
* output <type> - Output format (text/json)
* node <node-ip> - Address `<host>:<port>` of the node to connect. 
* broadcast-mode <mode> - Write transaction broadcast mode to use (one of: `sync`, `async`, `block`. `block` is default).

In order to connect the CLI to a DC Ledger Network (Chain), the following parameters should be used:

* `dcld config chain-id <chain-id>` - `<chain-id>` defines the Network you want to connect to
    * Use `testnet` if you want to connect to persistent Test Net
    * A full list of available persistent chains can be found in [Persistent Chains](../deployment/persistent_chains)
    where every sub-folder matches the corresponding chain-id.
* `dcld config output json` - Output format (text/json).
* `dcld config node tcp://<host>:<port>` - Address of a node to connect.
    * Example: `tcp://18.157.114.34:26657` or `on1.testnet.dcl.dev.dsr-corporation.com:26657`.
    * The address there is the address (IP or domain name) of one of the nodes (Validators or Observers) from the Network you are going to connect to.
    * A list of Observer Node addresses for persistent networks (such as the Test Net)
    can be found in the corresponding subfolders within [Persistent Chains](../deployment/persistent_chains).
    * Please look at [Observer Nodes](../deployment/persistent_chains/testnet/nodes.md) for a list of Observer nodes for the current `testnet`.
    * If you don't have any trusted nodes (Observers or Validators) for connection,
      consider running a Light Client Proxy Server and connecting a CLI to it. 
      See [Run Light Client Proxy](running-light-client-proxy.md).

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
* Generate keys and local account: `dcld keys add <account name>`.
    * You can remember and securely save the mnemonic phrase shown after the key is created
        to be able to recover the key later.
* Share generated `address` and `pubkey` with a number of `Trustee`s sufficient for account addition operation.
* One of `Trustee`s proposes to add the account to the ledger: `dcld tx auth propose-add-account --address=<account address> --pubkey=<account pubkey> --roles=<role1,role2,...> --vid=<vendorID for vendor role> --from=<trustee name>` (p.s. A vendor role is tied to a Vendor ID, hence while proposing a Vendor Role vid is a required field.)
* Sufficient number of other `Trustee`s approve the proposed account: `dcld tx auth approve-add-account --address=<account address> --from=<trustee name>`
* Check that the active account exists: `dcld query auth account --address=<account address>`

Example:
* `dcld keys add steve`
* `dcld keys show steve -a`   # assume it returns cosmos1sug8cquqnn5jddkqt4ud6hcr290sn4wh96x5tv
* `dcld keys show steve -p`   # assume it returns {"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A9v90lbd1tCtvXTKH3Fmir9wIg/cLlWU+/HSDnDYfaMm"}
* `dcld tx auth propose-add-account --address=cosmos1sug8cquqnn5jddkqt4ud6hcr290sn4wh96x5tv --pubkey={"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A9v90lbd1tCtvXTKH3Fmir9wIg/cLlWU+/HSDnDYfaMm"} --roles=Vendor,NodeAdmin --vid=4563 --from jack`
* `dcld tx auth approve-add-account --address=cosmos1sug8cquqnn5jddkqt4ud6hcr290sn4wh96x5tv --from alice`
* `dcld query auth account --address=cosmos1sug8cquqnn5jddkqt4ud6hcr290sn4wh96x5tv`

## Exporting and Importing Accounts

It's possible to export and import accounts by exporting and importing the corresponding keys.

There are two options on how it can be done:
1. Recover existing keys from the mnemonic passphrase saved when key was created.
2. Export a private key to ASCII-armored encrypted format, and then import it.

Example:
* Recovering an exiting key:

    `dcld keys add jack --recover`

* Exporting and importing a key:

    * `dcld keys export jack`
    * store the exported ASCII-armored encrypted key to a file `jack_exported_priv_key_file.txt`
    * `dcld keys import jack jack_exported_priv_key_file.txt`


## Trustee Instructions

Account creation consists of two parts. One of the trustees should propose an account by posting `propose-add-account` transaction.
After this account goes to the proposed accounts set. Then this account must be approved by a majority of trustees (2/3+ including the trustee who has proposed the account).
Once approved the account can be used to send transactions. See [use_case_txn_auth](use_cases/use_cases_txn_auth.png).

##### 1. Create an Account proposal for the user
```
dcld tx auth propose-add-account --address=<bench32 encoded string> --pubkey=<protobuf JSON encoded> --roles=<role1,role2,...> --vid=<uint16> --from=<account>
```

##### 2. Approve proposed Account
```
dcld tx auth approve-add-account --address=<bench32 encoded string> --from=<account>
```

##### 3. Propose revocation of an Account
```
dcld tx auth propose-revoke-account --address=<bench32 encoded string> --from=<account>
```

##### 4. Approve revocation of an Account
```
dcld tx auth approve-revoke-account --address=<bench32 encoded string> --from=<account>
```

##### 5. Propose a new self-signed root certificate
```
dcld tx pki propose-add-x509-root-cert --certificate=<string-or-path> --from=<account>
```
Example: `dcld tx pki propose-add-x509-root-cert --certificate="/path/to/certificate/file" --from=jack`
  
Example: `dcld tx pki propose-add-x509-root-cert --certificate="----BEGIN CERTIFICATE----- ......" --from=jack`

##### 6. Approve proposed X509 root certificate  
```
dcld tx pki approve-add-x509-root-cert --subject=<string> --subject-key-id=<hex string> --from=<account>
```
 
##### 7. Propose revocation of an X509 root certificate
```
dcld tx pki propose-revoke-x509-root-cert --subject=<string> --subject-key-id=<hex string> --from=<account>
```
  
 
##### 8. Approve revocation of an X509 root certificate  
```
dcld tx pki approve-revoke-x509-root-cert --subject=<string> --subject-key-id=<hex string> --from=<account>
```
 
  
## CA instructions
Currently any role can propose an X509 root certificate, or publish 
(intermediate or leaf) X509 certificates. 

##### 1. Propose a new self-signed root certificate
```
dcld tx pki propose-add-x509-root-cert --certificate=<string-or-path> --from=<account>
```
Example: `dcld tx pki propose-add-x509-root-cert --certificate="/path/to/certificate/file" --from=jack`
  
Example: `dcld tx pki propose-add-x509-root-cert --certificate="----BEGIN CERTIFICATE----- ......" --from=jack`

##### 2. Publish an intermediate or leaf X509 certificate
```
dcld tx pki add-x509-cert --certificate=<string-or-path> --from=<account>
```
The certificate must be signed by a chain of certificates which must be already present on the ledger.
 
  Example: `dcld tx pki add-x509-cert --certificate="/path/to/certificate/file" --from=jack`
  
  Example: `dcld tx pki add-x509-cert --certificate="----BEGIN CERTIFICATE----- ......" --from=jack`  
    
##### 3. Revoke an intermediate or leaf X509 certificate
```
dcld tx pki revoke-x509-cert --subject=<string> --subject-key-id=<hex string> --from=<account>
```
Can be done by the certificate's issuer only.
 
   
## Vendor Instructions

##### 1. Publish an intermediate or leaf X509 certificate(s) to be used for signing X509 Certificates for every Device
```
dcld tx pki add-x509-cert --certificate=<string-or-path> --from=<account>
```
The certificate must be signed by a chain of certificates which must be already present on the ledger.

##### 2. Add Vendor Info
```
dcld tx vendorinfo add-vendor --vid=<uint16> --vendorName=<string> --companyLegalName=<string> --companyPreferredName=<string> --vendorLandingPageURL=<string> --from=<account>
```

##### 3. Add a new model info with the given VID/PID
Minimal command:
```
 dcld tx model add-model --vid=<uint16> --pid=<uint16> --deviceTypeID=<uint16> --productName=<string> --productLabel=<string or path> --partNumber=<string> 
 --from=<account>
```
Full command:
```
dcld tx model add-model --vid=<uint16> --pid=<uint16> --deviceTypeID=<uint16> --productName=<string> --productLabel=<string or path> --partNumber=<string> 
    --commissioningCustomFlow=<uint8> --commissioningCustomFlowUrl=<string> --commissioningModeInitialStepsHint=<uint32> --commissioningModeInitialStepsInstruction=<string>
    --commissioningModeSecondaryStepsHint=<uint32> --commissioningModeSecondaryStepsInstruction=<string> --userManualURL=<string> --supportURL=<string> --productURL=<string> --lsfURL=<string> --lsfRevision=<uint32>
    --from=<account>
```

##### 4. Add a new model version for the given VID/PID and Software Version
Minimal command:
```
dcld tx model add-model-version --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --softwareVersionString=<string> --cdVersionNumber=<uint32>
--minApplicableSoftwareVersion=<uint32> --maxApplicableSoftwareVersion=<uint32> --from=<account>
```
Full command:
```
dcld tx model add-model-version --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --softwareVersionString=<string> --cdVersionNumber=<uint32>
--minApplicableSoftwareVersion=<uint32> --maxApplicableSoftwareVersion=<uint32>
--firmwareDigests=<string> --softwareVersionValid=<bool> --otaURL=<string> --otaFileSize=<string> --otaChecksum=<string> --otaChecksumType=<string> --releaseNotesURL=<string> 
--from=<account>
```

##### 5. Edit Vendor Info
```
dcld tx vendorinfo update-vendor --vid=<uint16> ... --from=<account>
```

##### 6. Edit Model Info
```
dcld tx model update-model --vid=<uint16> --pid=<uint16> ... --from=<account>
```

##### 7. Edit Model Version
```
dcld tx model update-model-version --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> ... --from=<account>
```

## Test House Instructions
##### 1A. Publish an intermediate or leaf X509 certificate(s) to be used for signing the Certification blob
This step is needed for off-ledger certification use case only, see [use_cases_device_off_ledger_certification](use_cases/use_cases_device_off_ledger_certification.png).

The certificate must be signed by a chain of certificates which must be already present on the ledger.
 
```
dcld tx pki add-x509-cert --certificate=<string-or-path> --from=<account>
```
    
##### 1B. Add a new testing result for the device model with the given VID/PID
This step is needed for on-ledger certification use case only, see [use_cases_device_on_ledger_certification](use_cases/use_cases_device_on_ledger_certification.png).

The corresponding model must present on the ledger.

```
dcld tx compliancetest add-test-result --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --softwareVersionString=<string> --test-result=<string-or-path> --test-date=<rfc3339 encoded date> --from=<account>
```
 
   
## Certification Center Instructions
##### 1A. Publish an intermediate or leaf X509 certificate(s) to be used for signing the Certification blob
This step is needed for off-ledger certification use case only, see [use_cases_device_off_ledger_certification](use_cases/use_cases_device_off_ledger_certification.png).

The certificate must be signed by a chain of certificates which must be already present on the ledger.
 
```
dcld tx pki add-x509-cert --certificate=<string-or-path> --from=<account>
```
    
##### 1B. Certify the device model with the given VID/PID
This step is needed for on-ledger certification use case only, see [use_cases_device_on_ledger_certification](use_cases/use_cases_device_on_ledger_certification.png).

The corresponding model and the test results must be present on the ledger.

```
dcld tx compliance certify-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --softwareVersionString=<string>  --certificationType=<matter|zigbee> --certificationDate=<rfc3339 encoded date> --reason=<string> --from=<account>
```

##### 2. Revoke certification for the device model with the given VID/PID
This step can be used in either on-ledger certification use case
 ([use_cases_device_on_ledger_certification](use_cases/use_cases_device_on_ledger_certification.png))
  or off-ledger certification use case ([use_cases_device_off_ledger_certification](use_cases/use_cases_device_off_ledger_certification.png)).
 
The corresponding Model Info and test results are not required to be on the ledger.

```
dcld tx compliance revoke-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<matter|zigbee> --revocationDate=<rfc3339 encoded date> --reason=<string> --from=<account>
```
  
 
## Node Admin Instructions (Setting up a new Validator Node)

Validators are responsible for committing of new blocks to the ledger.

A more detailed instruction on how to add a validator node to an existing DC Ledger network
can be found here: [running-node.md](running-node.md).

* `dcld init node-name --chain-id dclchain`
* `cp /source/genesis.json $HOME/.dcl/config/`
* `sed -i "s/persistent_peers = \"\"/<node id>@<node ip>,<node2 id>@<node2 ip>/g" $HOME/.dcl/config/config.toml`
* `sudo ufw allow 26656/tcp`
* `sudo ufw allow 26657/tcp`
* `dcld start`
* `dcld status`
* `dcld tx validator add-node --pubkey=(dcld tendermint show-validator) --moniker=<string> --from=<account>`
* `dcld query validator all-nodes`

##### Policy

**Please note, that Jailing is currently disabled!**

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
