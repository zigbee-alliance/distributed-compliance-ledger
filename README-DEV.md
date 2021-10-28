# Development Guideline

Please note, that the only officially supported platform now is Linux.
It's recommended to develop and deploy the App on Ubuntu 18.04 or Ubuntu 20.04.

## Pre-requisites
1. Install Go as described at https://golang.org/doc/install
    - Ensure that the following line has been appended to `/etc/profile`:
        ```
        export PATH=$PATH:/usr/local/go/bin
        ```
    - Ensure that the following line has been appended to `~/.profile`:
        ```
        export PATH=$PATH:~/go/bin
        ```
2. Install Docker as described at https://docs.docker.com/engine/install/ubuntu/
    - In `Installation methods` section follow `Install using the repository` method
    - Check whether your user of Ubuntu has been added to `docker` group using the following command:
        ```
        getent group docker | awk -F: '{print $4}'
        ```
        - If it has not been added, add it using `Manage Docker as a non-root user` section from https://docs.docker.com/engine/install/linux-postinstall/
 3. Install Docker Compose as described at https://docs.docker.com/compose/install/

## Build and test

1. Building
    ```
    make build
    make install
    ```


2. Run unit tests
    ```
    make test
    ```

3. Run integration tests.

    The integration tests are run against a local pool of nodes in Docker.
    REST integration tests need to have a backend running (CLI in REST mode).

    The following script will start all necessary things and run the tests:
    ```
    ./integration_tests/ci/run-all.sh
    ```


## Run local pool
The easiest way to run a local pool is to start it in Docker:

    make install
    make localnet_init
    make localnet_start

This will start a local pool of 4 validator nodes in Docker. The nodes will expose their RPC enpoints on ports `26657`, `26659`, `26661`, `26663` correspondingly.

 Stopping the network: 

    make localnet_stop

 Then you can start the network again with the existing data using `make localnet_start`

If you need to start a new clean network then do the following steps prior to executing `make localnet_start`:
  - Remove `.dclcli` and `.dcld` directories from your user home directory (`~`)
  - Remove `localnet` directory from the root directory of the cloned project
  - Initialize the new network data using `make localnet_init` 
## Run CLI
Start a local pool as described above, and then just execute
```
dclcli
```
Have a look at [How To](docs/how-to.md) and [CLI Help](docs/cli-help.md) for instructions how to configure and use the CLI.

## Remote Debugging local pool 
If you want to remotely debug the node running in docker in order to inspect and step thru the code, modify the following lines in the `docker-compose.yml`. Comment out the default `dcld start` command with the `dlv` as shown below (delve is the go remote debugger) 
```
    # uncomment following line if starting in debug mode 
    - "2345:2345"

    # command: dcld start    
    # Please use the following as the entry point if you want to start this node in debug mode for easy debugging
    command: dlv --listen=:2345 --headless=true --log=true --log-output=debugger,debuglineerr,gdbwire,lldbout,rpc --accept-multiclient --api-version=2 exec /usr/bin/dcld start
```
After making the changes re-initialize and start the localnet
  - `make localnet_stop` -If your localnet is still up
  - `make localnet_init` 
  - `make localnet_start`

Once all the four nodes are running, the last node i.e. node03 will be listing on port `2345` to which you can attach a debug process from any IDE (e.g. Visual Studio) and step thru the code. (p.s. node03 will only start working as validator node once the debugger is attached.). More details about IDE configuration can be found at https://github.com/go-delve/delve/blob/master/Documentation/EditorIntegration.md

## Contributing
Please take into account the following when sending a PR:
1) Make sure there is a license header added:
    - Have a look at `make license` and `make license-check` command in [Makefile](Makefile).

2) Make sure the new functionality has unit tests added

3) Make sure the new functionality has integration tests added
    - [CLI-based tests](integration_tests/cli)
    - [REST-based tests](integration_tests/rest)

4) There is CI based on GitHub Actions that will do the following for every Pull Request:
    - make sure the app can be built
    - run go linter
    - run unit tests
    - run integratioins test
    - make sure there is a license header in all the files


## Other
For more details, please have a look at [Cosmos SDK tutorial](https://github.com/cosmos/sdk-tutorials/blob/master/nameservice/tutorial).
Use __dcld__, __dclcli__ instead of __nameserviced__, __nameservicecli__.
