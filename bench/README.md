# DCLedger Load Testing

## Requirements

*   python >= 3.7

## Installation

Run (consider to use virtual environment):

```bash
pip3 install -r bench/requirements.txt
```

## Preparation

Each write transactions is signed and thus requires:

*   an account with write permissions (e.g. Vendor account)
*   proper values for txn `sequence` which enforces txns ordering for an account

By that reason load test uses prepared load data which can be generated as follows:

```bash
$ sudo make localnet_clean
$ make localnet_init
$ ./gentestaccounts.sh <NUM-USERS>
$ make localnet_start
# Note: once started ledger may require some time to complete the initialization.
$ DCLBENCH_WRITE_USERS_COUNT=<NUM-USERS> DCLBENCH_WRITE_USERS_Q_COUNT=<NUM-REQ-PER-USER> python bench/generate.py bench/test.spec.yaml ./txns
```

Here the following inputs are considered:

*   `NUM-USERS`: number of client accounts with write access (created as Vendors)
*   `NUM-REQ-PER-USER`: number of write txns to perform per a user

## Run

### Headless

```bash
locust -f bench/locustfile.py --headless --dcl-users <NUM-USERS> -s 10
```

### Web UI

```bash
locust -f bench/locustfile.py --dcl-users <NUM-USERS> -s 10
```

Then you can open <http://localhost:8089/> and launch the tests from the browser.

### Configuration

*   `--dcl-users`: number of users
*   `--dcl-spawn-rate` Rate to spawn users at (users per second)
*   `--dcl-hosts <comma-sepated-list>`: list of DCL nodes to target. Each user randomly picks one
    E.g. for local ledger `http://localhost:26657,http://localhost:26659,http://localhost:26661,http://localhost:26663` will specify all the nodes.
*   `--dcl-txn-file` path to a file with generated txns

Please check `locust -f bench/locustfile.py --help` for the more details.

### Re-run

Next time when you run the test using the same data you will likely get many (all) failures since DCLedger
will complain about already written data or wrong sequence numbers.

For that case you may consider to reset the ledger as follows:

```bash
$ make localnet_reset localnet_start
```

## FAQ

### locust may complain about ulimit values for open files

Please check the details [here](https://github.com/locustio/locust/wiki/Installation#increasing-maximum-number-of-open-files-limit).

Additional sources (linux):

*   `man limits.conf`
*   [RedHat: How to set ulimit values](https://access.redhat.com/solutions/61334)

## ToDo

*   explore the options to export test accounts to commit static data (accounts and test txns)
*   read requests loads
*   combined (and configured) loads: write + read
*   stat gathering and interpretation
*   non-local setups automation and targeting (e.g. AWS)
*   harden data generation scripts
