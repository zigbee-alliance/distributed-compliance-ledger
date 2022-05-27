# DCLedger Load Testing

DCLedger testing is implemented in python3 and based on [Locust](https://locust.io/) framework.

## Requirements

* python >= 3.7
* Install Docker as described at https://docs.docker.com/engine/install/ubuntu/
    - In `Installation methods` section follow `Install using the repository` method
        - Check whether your user of Ubuntu has been added to `docker` group using the following command:

            ```bash
            getent group docker | awk -F: '{print $4}'
            ```

            - If it has not been added, add it using `Manage Docker as a non-root user` section from <https://docs.docker.com/engine/install/linux-postinstall/>
* Install Docker Compose as described at <https://docs.docker.com/compose/install/>

## Installation

Run (consider to use virtual environment):

```bash
pip3 install -r bench/requirements.txt
```

**Optional** If you need to monitor server-side metrics please install [Prometheus](https://prometheus.io/docs/prometheus/latest/getting_started/).

## Preparation

Each write transactions is signed and thus requires:

* an account with write permissions (e.g. Vendor account)
* proper values for txn `sequence` which enforces txns ordering for an account
* Initialize the pool and test accounts (**Warning** applicable to local in-docker pool only for now):

    ```bash
    make localnet_clean

    make localnet_init

    make localnet_start
    # Note: once started ledger may require some time to complete the initialization.
    ```

* Copy local accounts keys to folder `~/.dcl/keyring-test`:
    ```bash
    cp ./.localnet/node0/keyring-test/* ~/.dcl/keyring-test
    ```
* Enter to `bench` folder
    ```bash
    cd bench
    ```
* Build `docker-compose` file
    ```bash
    docker-compose build
    ```

## Run

### (Optional) Launch Prometheus

```bash
prometheus --config.file=bench/prometheus.yml
```

And open `http://localhost:9090/` to query and monitor the server-side metrics.

## Headless

To run write load tests

* Open the docker-compose.yml and update field `command` for `master` and `worker`.  
    e.g.: for `master`
    ```yml
    command: -f /dcl/bench/locustfile.py --headless WriteModelLoadTest --master -H http://master:8089 
    ```
    e.g.: for `worker`:
    ```yml
    command: -f /dcl/bench/locustfile.py --headless WriteModelLoadTest --worker --master-host master
    ```

* Run docker-compose.yml
    ```bash
    docker-compose up --scale worker=<workers-count>
    # The workers run your Users and send back statistics to the master. The master instance doesn't run any Users itself. Both the master and worker machines must have a copy of the locustfile when running Locust distributed.
    ```
    `<workers-count>`- number of machine.

To run read load tests

* Open the docker-compose.yml and update field `command` for `master` and `worker`.  
    e.g.: for `master`
    ```yml
    command: -f /dcl/bench/locustfile.py --headless ReadModelLoadTest --master -H http://master:8089 
    ```
    e.g.: for `worker`:
    ```yml
    command: -f /dcl/bench/locustfile.py --headless ReadModelLoadTest --worker --master-host master
    ```

* Run docker-compose.yml
    ```bash
    docker-compose up --scale worker=<workers-count>
    # The workers run your Users and send back statistics to the master. The master instance doesn't run any Users itself. Both the master and worker machines must have a copy of the locustfile when running Locust distributed.
    ```
    `<workers-count>`- number of machine.


To run write/read load tests

* Open the docker-compose.yml and update field `command` for `master` and `worker`.  
    e.g.: for `master`
    ```yml
    command: -f /dcl/bench/locustfile.py --headless --master -H http://master:8089 
    ```
    e.g.: for `worker`:
    ```yml
    command: -f /dcl/bench/locustfile.py --headless --worker --master-host master
    ```

* Run docker-compose.yml
    ```bash
    docker-compose up --scale worker=<workers-count>
    # The workers run your Users and send back statistics to the master. The master instance doesn't run any Users itself. Both the master and worker machines must have a copy of the locustfile when running Locust distributed.
    ```
    `<workers-count>`- number of machine.


## Web UI

To run write load tests

* Open the docker-compose.yml and update field `command` for `master` and `worker`.  
    <br> e.g.: for `master`
    ```yml
    command: -f /dcl/bench/locustfile.py WriteModelLoadTest --master -H http://master:8089 
    ```
    e.g.: for `worker`:
    ```yml
    command: -f /dcl/bench/locustfile.py WriteModelLoadTest --worker --master-host master
    ```

* Run docker-compose.yml
    ```bash
    docker-compose up --scale worker=<workers-count>
    # The workers run your Users and send back statistics to the master. The master instance doesn't run any Users itself. Both the master and worker machines must have a copy of the locustfile when running Locust distributed.
    ```
    `<workers-count>`- number of machine.

To run read load tests

* Open the docker-compose.yml and update field `command` for `master` and `worker`.  
    <br> e.g.: for `master`
    ```yml
    command: -f /dcl/bench/locustfile.py ReadModelLoadTest --master -H http://master:8089 
    ```
    e.g.: for `worker`:
    ```yml
    command: -f /dcl/bench/locustfile.py ReadModelLoadTest --worker --master-host master
    ```

* Run docker-compose.yml
    ```bash
    docker-compose up --scale worker=<workers-count>
    # The workers run your Users and send back statistics to the master. The master instance doesn't run any Users itself. Both the master and worker machines must have a copy of the locustfile when running Locust distributed.
    ```
    `<workers-count>`- number of machine.

To run write/read load tests

* Open the docker-compose.yml and update field `command` for `master` and `worker`.  
    <br> e.g.: for `master`
    ```yml
    command: -f /dcl/bench/locustfile.py --master -H http://master:8089 
    ```
    e.g.: for `worker`:
    ```yml
    command: -f /dcl/bench/locustfile.py --worker --master-host master
    ```

* Run docker-compose.yml
    ```bash
    docker-compose up --scale worker=<workers-count>
    # The workers run your Users and send back statistics to the master. The master instance doesn't run any Users itself. Both the master and worker machines must have a copy of the locustfile when running Locust distributed.
    ```
    `<workers-count>`- number of machine.

Then you can open `http://localhost:8089/` and launch the tests from the browser.

### Configuration

Run options (DCLedger custom ones):

* `--dcl-users`: number of users
* `--dcl-spawn-rate` Rate to spawn users at (users per second)
* `--dcl-hosts <comma-sepated-list>`: list of DCL nodes to target. Each user randomly picks one
    E.g. for local ledger `http://localhost:26657,http://localhost:26659,http://localhost:26661,http://localhost:26663` will specify all the nodes.
* `--dcl-rest-hosts <comma-sepated-list>`: list of DCL nodes to target. Each user randomly picks one
    E.g. for local ledger `http://localhost:26640,http://localhost:26641,http://localhost:26642,http://localhost:26643` will specify all the nodes.
* `--dcl-trustee-account-name`: name of existing Trustee account.
    E.g. for local ledger `jack`. Jack is Trustee account name


Statistic options:

[Locust](https://locust.io/) provides the following options to present the results:

* `--csv <prefix>`: generates a set of stat files (summary, failures, exceptions and stats history) with the provided `<prefix>`
* `--csv-full-history`: populates the stats history with more entries (including each specific request type)
* `--html <path>`: generates an html report
* Web UI also includes `Download Data` tab where the reports can be found.

More details can be found in:

* [locust.conf](../locust.conf): default values
* `locust --help` (being in the project root)
* [locust configuration](https://docs.locust.io/en/stable/configuration.html)
* [locust stats](https://docs.locust.io/en/stable/retrieving-stats.html)

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

* `man limits.conf`
* [RedHat: How to set ulimit values](https://access.redhat.com/solutions/61334)

## ToDo

* explore the options to export test accounts to commit static data (accounts and test txns)
* read requests loads
* combined (and configured) loads: write + read
* stat gathering and interpretation
* non-local setups automation and targeting (e.g. AWS)
* harden data generation scripts
* consider different types of tx: async, sync (currently used), commit
