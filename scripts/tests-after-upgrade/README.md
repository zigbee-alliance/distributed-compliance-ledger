# Tests after upgrade

 Tests after upgrade are used to verify the functionality of the pool after the upgrade.  
 The test suite contains Smoke test cases that cover basic read and write (excluding the MainNet environment) commands.

## Environment Setup

The script supports three environments:
- `TestNet`
- `MainNet` 
- `local`

Each environment has a corresponding `.env` file located in the `scripts/tests-after-upgrade/` directory. 

Run the script with the desired environment as an argument:
```
./scripts/tests-after-upgrade/tests-after-upgrade.sh testnet  # Use TestNet environment
./scripts/tests-after-upgrade/tests-after-upgrade.sh mainnet  # Use MainNet environment
./scripts/tests-after-upgrade/tests-after-upgrade.sh local    # Use local environment
```

If no environment is specified, the script defaults to `TestNet`.

## Passphrase Requirement

For the `local` and `TestNet` environments, the corresponding `.env` file **must** include a non-empty `passphrase` variable:

```
passphrase="your wallet passphrase here"
```
