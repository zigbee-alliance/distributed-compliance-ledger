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
    ```bash
    make test
    ```

3. Run integration tests.

    The integration tests are run against a local pool of nodes in Docker.
    REST integration tests need to have a backend running (CLI in REST mode).

    The following script will start all necessary things and run all the tests:
    ```bash
    ./integration_tests/run-all.sh
    ```
    If you want to run a particular group of tests (cli, light, rest), you can
    ```bash
    ./integration_tests/run-all.sh cli
    ./integration_tests/run-all.sh light
    ./integration_tests/run-all.sh rest
    ./integration_tests/run-all.sh cli,light
    ```

    If you want to run a particular test you may:

    ```bash
    ./integration_tests/start-pool.sh
   
    bash <path-to-shell-script>  # to run a cli test
    # OR
    go test <path-to-go-test-file> # to run REST or gRPC go test
    ```
   
4. Run deployment test
    
    The deployment test verifies deployment steps described in [docs/running-node.md](./docs/running-node.md).

    ```bash
    ./integration_tests/deploy/test_deploy.sh
    ```


## Run local pool
The easiest way to run a local pool is to start it in Docker.

Validator nodes only (no Observers):
```bash
make install
make localnet_init
make localnet_start
```

Validator and Observer nodes:
```bash
make install
DCL_OBSERVERS=1 make localnet_init
make localnet_start
```    

This will start a local pool of 4 validator nodes in Docker. The nodes will expose their RPC enpoints on ports `26657`, `26659`, `26661`, `26663` correspondingly.

 Stopping the network: 

    make localnet_stop

 Then you can start the network again with the existing data using `make localnet_start`

If you need to start a new clean network then run `make localnet_rebuild` prior to executing `make localnet_start`.
It will remove `.dcl` directories from your user home directory (`~`), remove `.localnet` directory from the root directory of the cloned project,
and initialize a new network data using `make localnet_init`.

## CLI
Start a local pool as described above, and then just execute
```
dcld
```
Have a look at [How To](docs/how-to.md) and [transactions](docs/transactions.md) for instructions how to configure and use the CLI.

## REST
Start a local pool as described above.

Every node exposes a REST API at `http://<node-host>:1317` (see https://docs.cosmos.network/master/core/grpc_rest.html).

Have a look at [transactions](docs/transactions.md) for a full list of REST endpoints.

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


## How To Add a new Module
- Use [starport](https://github.com/tendermint/starport) command to scaffold the module.
  Consider using the provided [Dockerfile](scripts/starportDockerfile) to use predictable version of starport. See [README.md](scripts/starport/README.md).
- Have a look at the scripts and commands used for generation of existing modules and do it in a similar way
  (for example [PKI module commands](scripts/starport/upgrade-0.44/07.pki_types.sh)).
- Adjust the generated code
  - correct REST endpoints: `/dcl` instead of `/zigbee-alliance/distributedcomplianceledger` in `proto/<module>/query.proto`
  - add message validation as annotations (`validate` tags) in `proto/<module>/tx.proto`
  - add `(cosmos_proto.scalar) = "cosmos.AddressString"` annotation for all fields with address/account type (such as `signer` or `owner`).
  - fix types if needed in `proto/<module>/<entity>.proto` files
    - Note1: `unit64` will be returned as string if the output is a JSON format. So, it's better to use `uint64` only when it's really `uint64`. 
    - Note2: for `uint16` type: use `int32` during starport scaffolding, and add custom validation (annotations above) to check the lower and upper bounds.
    - Note3: for `uint32` type: use `int32` during starport scaffolding, then replace it by `uint32` in .proto files, re-generate the code and fix compilation errors.
  - build proto (for example `starport chain build`). Fix compilation errors if any.
  - **Note1**: colons (`:`) are part of subject-id in PKI module, but colons are not allowed in gRPC REST URLs by default.
    `allow_colon_final_segments=true` should be used as a workaround.
    So, make sure that `runtime.AssumeColonVerbOpt(true)` in `/x/pki/types/query.pb.gw.go`. 
    It's usually sufficient to revert the generated changes in `/x/pki/types/query.pb.gw.go`.
  - **Note2**: starport will include all default cosmos modules (even if we don't use them from DCL) into `docs/static/openapi.yaml`. 
    Revert the default cosmos modules keeping only DCL ones.   
- Call `validator.Validate(msg)` in `ValidateBasic` methods for all generated messages
- Implement business logic in `msg_server_xxx.go`
- Improve `NotFound` error processing:
    - replace `status.Error(codes.InvalidArgument, "not found")` to `status.Error(codes.NotFound, "not found")` in every generated `grpc_query_xxx.go` to return 404 error in REST.
- Support state proof for single value queries in CLI:
    - use `cli.QueryWithProof` instead of cosmos queries that doesn't support state proofs
    - add proper handling for list queries and write requests when used with a light client proxy 
      (see `IsWriteInsteadReadRpcError` and `IsKeyNotFoundRpcError`)
- Add unit tests (see other modules for reference)
- Add CLI-based integration tests to `integration_tests/cli/<module>` (see other modules for reference)
- Add gRPC/REST-based integration tests to `integration_tests/grpc_rest/<module>` (see other modules for reference)

## How To Make Changes in Data Model for Existing Modules
- Use [starport](https://github.com/tendermint/starport) command to scaffold the module.
  Consider using the provided [Dockerfile](scripts/starportDockerfile) to use predictable version of starport. See [README.md](scripts/starport/README.md).
- **Never change `.pb` files manually**. Do the changes in `.proto` files.
- Every time `.proto` files change, re-generate the code (for example `starport chain build`) and fix compilation errors if any.
- **Note1**: colons (`:`) are part of subject-id in PKI module, but colons are not allowed in gRPC REST URLs by default.
  `allow_colon_final_segments=true` should be used as a workaround.
  So, make sure that `runtime.AssumeColonVerbOpt(true)` in `/x/pki/types/query.pb.gw.go`. 
  It's usually sufficient to revert the generated changes in `/x/pki/types/query.pb.gw.go`.
- **Note2**: starport will include all default cosmos modules (even if we don't use them from DCL) into `docs/static/openapi.yaml`. 
    Revert the default cosmos modules keeping only DCL ones.   

## Update Cosmos-sdk Version
Re-generate cosmos base openapi (service API from cosmos exposed in DCL) using [cosmos-base-swagger-gen.sh](scripts/cosmos-base-swagger-gen.sh).

## Update Tendermint Version
Please note, that we depend on the Tendermint fork https://github.com/zigbee-alliance/tendermint/releases/tag/v0.34.140 
due to hotfixes for https://github.com/tendermint/tendermint/issues/7640 and https://github.com/tendermint/tendermint/issues/7641
requitred for Light Client Proxy.
Now that fixes are merged to Tendermint master, so check if we still need to depend on the fork.

Also don't forget to update the link to the Tendermint RPC in [Swagger UI](docs/index.html).

## Other
For more details, please have a look at [Cosmos SDK tutorial](https://tutorials.cosmos.network/).

