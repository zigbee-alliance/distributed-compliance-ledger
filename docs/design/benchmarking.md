# Benchmarking

## Client Side Metrics

*   `response time` (percentiles): the time between client's initial request and the last byte of a validator response
*   `requests per second (RPS)`: number of requests per second
    *   `transactions per second (TPS)`: number of write requests (txns) per second
*   `number of clients`: number of concurrent clients that ledger serves
*   (optional) `throughtput` (in/out): number of KB per second. Marked as optional since we don't consider  much in/out data due to relatively small txns payloads.

## Server Side Metrics

Starting from `v0.40.0` Cosmos SDK provides [telemetry](https://docs.cosmos.network/master/core/telemetry.html) package as a server side support for application performance and behavior explorations.

The following [metrics](https://docs.cosmos.network/master/core/telemetry.html#supported-metrics) make sense to consider:

*   `tx_count`: Total number of txs processed via DeliverTx (tx)
*   `tx_successful`: Total number of successful txs processed via DeliverTx  (tx)
*   `tx_failed`: Total number of failed txs processed via DeliverTx
*   `abci_deliver_tx`: Duration of ABCI DeliverTx  (ms)
*   `abci_commit`: Duration of ABCI Commit (ms)
*   `abci_query`: Duration of ABCI Query  (ms)
*   `abci_begin_block`: Duration of ABCI BeginBlock (ms)
*   `abci_end_block`: Duration of ABCI EndBlock   (ms)
*   `begin_blocker`: Duration of BeginBlock for a given module   (ms)
*   `end_blocker`: Duration of EndBlock for a given module (ms)
*   `store_iavl_get`: Duration of an IAVL Store#Get call  (ms)
*   `store_iavl_set`: Duration of an IAVL Store#Set call  (ms)
*   `store_iavl_has`: Duration of an IAVL Store#Has call  (ms)
*   `store_iavl_delete`: Duration of an IAVL Store#Delete call   (ms)
*   `store_iavl_commit`: Duration of an IAVL Store#Commit call   (ms)
*   `store_iavl_query`: Duration of an IAVL Store#Query  (ms)

## Environment

**Note**. For the moment it's not clear enough what production setup will look like, in particular:

*   number of vendor companies (number of validators)
*   type of external endpoints, options are [Cosmos SDK / Tendermint endpoints](https://docs.cosmos.network/master/core/grpc_rest.html)
*   type and number of proxies for validator-validator and client-validator connections

Current assumptions:

*   multiple companies (vendors) will manage one/multiple validators
*   while some common requirements and recommendations would be provided each vendor will deploy the infrastructure independently with some freedom regarding internal architecture
*   there would be a set of external (for clients) and internal (for validators to support txn flows) endpoints
    *   most likely observer nodes along with REST http servers with clients authentication would be in front of the client endpoints

## Workloads

### Transactions Types

*   write txns:
    *   `tx/modelinfo/add-model`
        *   with `vid` constant for a particular client
        *   variable (incremented) `pid`
    *   **ToDo** consider other request types (e.g. `update-model`)
*   read txns:
    *   `query/modelinfo/model`
    *   **ToDo** consider other request types (e.g. `all-models`)

### Clients

**ToDo** define which client endpoints are considered in production

As long as CosmosSDK (Tendermint) provides multiple client [endpoints](https://docs.cosmos.network/master/core/grpc_rest.html) makes sense to benchmark all of them (separately and in a combination), in particular:

*   http RPC
*   websocket RPC
*   http REST

### Load Types

*   per txns types:
    *   only write txns: to measure server-side (consensus related) bottlenecks and limitations
    *   only read txns: to measure client-side (setup related) limitations
    *   combined loads with read/write ratio as a parameter
        *   **ToDo** define anticipated real loads
*   per scenario:
    *   stepping load: to identify the point where performance degrades significantly

    *   waves: to emulate peaks and troughs in clients behavior

## Load Generation Framework

As long as DCledger based on Cosmos SDK and Tendermint which provide standard HTTP/websocket RPC and REST  [endpoints](https://docs.cosmos.network/master/core/grpc_rest.html) to perform both read & write txns generic production ready tools like [jMeter](https://jmeter.apache.org/), [Locust](https://locust.io/), [K6](https://k6.io/) may be considered.

[Locust](https://locust.io/) looks like the most easy-to-go option:

*   tests can be configured using simple python scripts (version control, CI/CD), in comparison:
    *   JS based configuration for [K6](https://k6.io/) will likely require more efforts
    *   [jMeter](https://jmeter.apache.org/) configuration is mostly about UI but not coding
*   [distributed testing](https://docs.locust.io/en/stable/running-locust-distributed.html) with results aggregation is supported (if we decide to use it)
*   there are some [concerns](https://k6.io/blog/comparing-best-open-source-load-testing-tools/) regarding its performance and accuracy but the current vision is that it should be acceptable for our case

## Testing Environment Provisioning Automation

General considerations:

*   as long as target production deploy architecture is not yet defined automation for testing environment provisioning would simplify comparison of the options
*   single cloud provider as a starting point
*   multiple cloud provides as a very possible production case
*   tools: [terraform](https://www.terraform.io/) and [pulumi](https://www.pulumi.com/) as the preferred options

## ToDo

*   define acceptance criteria (target metrics values)
*   define target environment
