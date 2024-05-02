## [v1.4.0-dev3](https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/tag/v1.4.0-dev3) - 2024-04-24

### Dependency Changes
* Upgrade cosmos-sdk to v0.47.8
  * Migrate to CometBFT. Follow the migration instructions in the [upgrade guide](https://github.com/cosmos/cosmos-sdk/blob/main/UPGRADING.md#migration-to-cometbft-part-1).
  * Upgrade all related dependencies too 
* Upgrade Golang version to 1.20

### Breaking Changes
* Transaction broadcasting `block` mode has been removed from the updated cosmos-sdk. 
Starting from this version, dcl has only two modes: `sync` and `async`, with the default being `sync`.
In this mode, to obtain the actual result of a transaction (`txn`), an additional `query` call with the `txHash` must be executed. For example:
    `dcld query tx txHash` - where `txHash` represents the hash of the previously executed transaction."
* `starport` cli tool is no longer supported. Please use `v0.27.2` version of [ignite](https://github.com/ignite/cli).
* Due to upgrading `cosmovisor` to v1.3.0 in Docker and shell files, the node starting command has changed from `cosmovisor start` to `cosmovisor run start`
