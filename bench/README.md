# DCLedger Load Testing

DCLedger testing is implemented in python3 and based on [Locust](https://locust.io/) framework.

## Requirements

*   python >= 3.7

## Installation

Run (consider to use virtual environment):

```bash
pip3 install -r bench/requirements.txt
```

**Optional** If you need to monitor server side metrics please install [Prometheus](https://prometheus.io/docs/prometheus/latest/getting_started/).

## Preparation

Each write transactions is signed and thus requires:

*   an account with write permissions (e.g. Vendor account)
*   proper values for txn `sequence` which enforces txns ordering for an account

By that reason load test uses prepared load data which can be generated as follows:

*   Initialize the pool and test accounts (**Warning** applicable to local in-docker pool only for now):

    ```bash
    make localnet_clean

    # DCL_OBSERVERS=1 make localnet_init  # to initialize observers as well
    make localnet_init
    
    # ./gentestaccounts.sh [<NUM-USERS>]
    ./gentestaccounts.sh

    make localnet_start
    # Note: once started ledger may require some time to complete the initialization.
    ```
*   Generate test transactions:

    ```bash
    # DCLBENCH_WRITE_USERS_COUNT=<NUM-USERS> DCLBENCH_WRITE_USERS_Q_COUNT=<NUM-REQ-PER-USER> python bench/generate.py bench/test.spec.yaml bench/txns
    python bench/generate.py bench/test.spec.yaml bench/txns
    ```

Here the following (**optional**) inputs are considered:

*   `NUM-USERS`: number of client accounts with write access (created as Vendors). Default: 10
*   `NUM-REQ-PER-USER`: number of write txns to perform per a user. Default: 1000

## Run

### (Optional) Launch Prometheus

```bash
prometheus --config.file=bench/prometheus.yml
```

And open <http://localhost:9090/> to query and monitor the server side metrics.

### Headless

```bash
locust --headless
```

### Web UI

```bash
locust
```

Then you can open <http://localhost:8089/> and launch the tests from the browser.

### Configuration

Run options (DCLedger custom ones):

*   `--dcl-users`: number of users
*   `--dcl-spawn-rate` Rate to spawn users at (users per second)
*   `--dcl-hosts <comma-sepated-list>`: list of DCL nodes to target. Each user randomly picks one
    E.g. for local ledger `http://localhost:26657,http://localhost:26659,http://localhost:26661,http://localhost:26663` will specify all the nodes.
*   `--dcl-txn-file` path to a file with generated txns

Statistic options:

[Locust](https://locust.io/) provides the following options to present the results:

*   `--csv <prefix>`: generates a set of stat files (summary, failures, exceptions and stats history) with the provided `<prefix>`
*   `--csv-full-history`: populates the stats history with more entries (including each specific request type)
*   `--html <path>`: generates an html report
*   Web UI also includes `Download Data` tab where the reports can be found.

More details can be found in:

*   [locust.conf](../locust.conf): default values
*   `locust --help` (being in the project root)
*   [locust configuration](https://docs.locust.io/en/stable/configuration.html)
*   [locust stats](https://docs.locust.io/en/stable/retrieving-stats.html)

### Re-run

**Warning** applicable to local in-docker pool only for now

Next time when you run the test using the same data you will likely get many (all) failures since DCLedger
will complain about already written data or wrong sequence numbers.

For that case you may consider to reset the ledger as follows:

```bash
make localnet_reset localnet_start
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
*   consider different types of tx: async, sync (currently used), commit
