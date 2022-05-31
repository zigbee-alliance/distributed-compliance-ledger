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

* If you run load tests on a local pool, follow the steps below (**this item is needed only for a local pool**):<br>
  * Each write transactions is signed and thus requires:

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

* Copy accounts keys to folder `~/.dcl/keyring-test`:
    * After connecting to the network, we must copy the `Trustee` account keys to `~/.dcl/keyring-test`. <br>
    `The number of accounts with role Trustee must not exceed 3`.
* Enter to `bench` folder
    ```bash
    cd bench
    ```

* Open the .env file:
    * update fields.<br>
        e.g.:
        ```yml
        DCLD_VERSION=v0.11.0
        DCLD_NODE=tcp://host.docker.internal:26657
        DCLD_CHAIN_ID=dclchain

        WRITE_HOSTS=http://host.docker.internal:26657
        READ_HOSTS=http://host.docker.internal:26640
        TRUSTEE_ACCOUNT_NAME=jack
        COUNT_USERS=4 
        ```
        `<DCLD_VERSION>` - dcld binary version.<br>
        `<DCLD_NODE>` - Address `<host>:<port>` of the node to connect. This node needs for adding account with role `Vendor` in write load tests.<br>
        `<DCLD_CHAIN_ID>` - unique chain ID of the network you are going to connect.<br>
        `<WRITE_HOSTS>` - hosts for writing load tests.<br>
        `<READ_HOSTS>` - hosts for reading load tests.<br>
        `<TRUSTEE_ACCOUNT_NAME>` - trustee account name. `Trustee` account, which will add account with `Vendor` role in write load tests.<br>
        `<COUNT_USERS>` - number of users in load tests. `Number of users should be equal to number of workers for write tests.`


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

### Headless

To run write load tests

* Open the docker-compose.yml and add a new `command` for `master` and `worker`.  
    e.g.: for `master`
    ```yml
    command:
      - "--headless"
      - "WriteModelLoadTest"
    ```
    e.g.: for `worker`:
    ```yml
    command: 
      - "--headless"
      - "WriteModelLoadTest"
    ```
* Open the docker-compose.yml and add a new field `extra_hosts` for `master` and `worker` (**this item need only for the local pool**):<br>
  e.g.: for `master`
  ```yml
  extra_hosts:
    - "host.docker.internal:host-gateway"
  ```
  e.g.: for `worker`
  ```yml
  extra_hosts:
    - "host.docker.internal:host-gateway"
  ```
  `<host.docker.internal:host-gateway>` - need to connecting to localhost from locust container.

* Run docker-compose.yml
    ```bash
    docker-compose up --scale worker=<workers-count>
    # The workers run your Users and send back statistics to the master. The master instance doesn't run any Users itself. Both the master and worker machines must have a copy of the locustfile when running Locust distributed.
    ```
    `<workers-count>`- number of machine. `Number of users should be equal to number of workers for write tests.`

To run read load tests

* Open the docker-compose.yml and add a new `command` for `master` and `worker`.  
    e.g.: for `master`
    ```yml
    command: 
      - "--headless"
      - "ReadModelLoadTest" 
    ```
    e.g.: for `worker`:
    ```yml
    command: 
      - "--headless"
      - "ReadModelLoadTest" 
    ```
* Open the docker-compose.yml and add a new field `extra_hosts` for `master` and `worker` (**this item need only for the local pool**):<br>
  e.g.: for `master`
  ```yml
  extra_hosts:
    - "host.docker.internal:host-gateway"
  ```
  e.g.: for `worker`
  ```yml
  extra_hosts:
    - "host.docker.internal:host-gateway"
  ```
  `<host.docker.internal:host-gateway>` - need to connecting to localhost from locust container.

* Run docker-compose.yml
    ```bash
    docker-compose up --scale worker=<workers-count>
    # The workers run your Users and send back statistics to the master. The master instance doesn't run any Users itself. Both the master and worker machines must have a copy of the locustfile when running Locust distributed.
    ```
    `<workers-count>`- number of machine.


To run write/read load tests

* Open the docker-compose.yml and add a new `command` for `master` and `worker`.  
    e.g.: for `master`
    ```yml
    command: 
      - "--headless"
    ```
    e.g.: for `worker`:
    ```yml
    command: 
      - "--headless" 
    ```
* Open the docker-compose.yml and add a new field `extra_hosts` for `master` and `worker` (**this item need only for the local pool**):<br>
  e.g.: for `master`
  ```yml
  extra_hosts:
    - "host.docker.internal:host-gateway"
  ```
  e.g.: for `worker`
  ```yml
  extra_hosts:
    - "host.docker.internal:host-gateway"
  ```
  `<host.docker.internal:host-gateway>` - need to connecting to localhost from locust container.

* Run docker-compose.yml
    ```bash
    docker-compose up --scale worker=<workers-count>
    # The workers run your Users and send back statistics to the master. The master instance doesn't run any Users itself. Both the master and worker machines must have a copy of the locustfile when running Locust distributed.
    ```
    `<workers-count>`- number of machine. `Number of users should be equal to number of workers for write tests.`


## Web UI

To run write load tests

* Open the docker-compose.yml and add a new `command` for `master` and `worker`.  
    <br> e.g.: for `master`
    ```yml
    command: 
      - "WriteModelLoadTest" 
    ```
    e.g.: for `worker`:
    ```yml
    command: 
      - "WriteModelLoadTest" 
    ```

* Open the docker-compose.yml and add a new field `extra_hosts` for `master` and `worker` (**this item need only for the local pool**):<br>
  e.g.: for `master`
  ```yml
  extra_hosts:
    - "host.docker.internal:host-gateway"
  ```
  e.g.: for `worker`
  ```yml
  extra_hosts:
    - "host.docker.internal:host-gateway"
  ```
  `<host.docker.internal:host-gateway>` - need to connecting to localhost from locust container.

* Run docker-compose.yml
    ```bash
    docker-compose up --scale worker=<workers-count>
    # The workers run your Users and send back statistics to the master. The master instance doesn't run any Users itself. Both the master and worker machines must have a copy of the locustfile when running Locust distributed.
    ```
    `<workers-count>`- number of machine. `Number of users should be equal to number of workers for write tests.`

To run read load tests

* Open the docker-compose.yml and add a new `command` for `master` and `worker`.  
    <br> e.g.: for `master`
    ```yml
    command:
      - "ReadModelLoadTest" 
    ```
    e.g.: for `worker`: 
    ```yml
    command: 
      - "ReadModelLoadTest" 
    ```

* Open the docker-compose.yml and add a new field `extra_hosts` for `master` and `worker` (**this item need only for the local pool**):<br>
  e.g.: for `master`
  ```yml
  extra_hosts:
    - "host.docker.internal:host-gateway"
  ```
  e.g.: for `worker`
  ```yml
  extra_hosts:
    - "host.docker.internal:host-gateway"
  ```
  `<host.docker.internal:host-gateway>` - need to connecting to localhost from locust container.

* Run docker-compose.yml
    ```bash
    docker-compose up --scale worker=<workers-count>
    # The workers run your Users and send back statistics to the master. The master instance doesn't run any Users itself. Both the master and worker machines must have a copy of the locustfile when running Locust distributed.
    ```
    `<workers-count>`- number of machine.

To run write/read load tests

* Open the docker-compose.yml and add a new field `extra_hosts` for `master` and `worker` (**this item need only for the local pool**):<br>
  e.g.: for `master`
  ```yml
  extra_hosts:
    - "host.docker.internal:host-gateway"
  ```
  e.g.: for `worker`
  ```yml
  extra_hosts:
    - "host.docker.internal:host-gateway"
  ```
  `<host.docker.internal:host-gateway>` - need to connecting to localhost from locust container.

* Run docker-compose.yml
    ```bash
    docker-compose up --scale worker=<workers-count>
    # The workers run your Users and send back statistics to the master. The master instance doesn't run any Users itself. Both the master and worker machines must have a copy of the locustfile when running Locust distributed.
    ```
    `<workers-count>`- number of machine. `Number of users should be equal to number of workers for write tests.`

Then you can open `http://localhost:8089/` and launch the tests from the browser.

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
* consider different types of tx: async, sync (currently used), commit
