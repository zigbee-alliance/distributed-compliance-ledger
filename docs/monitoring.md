# DCLedger Monitoring

## Overview

As long as DCLedger is built on top of Cosmos SDK and Tendermint their monitoring abilities can be considered for DCLedger as well.

*   Starting from `v0.21.0` Tendermint provides Prometheus compatible [metrics](https://docs.tendermint.com/v0.34/tendermint-core/metrics.html).
*   Starting from `v0.40.0` Cosmos SDK provides [telemetry](https://docs.cosmos.network/master/core/telemetry.html) package as a server-side support for application performance and behavior explorations.

## Installation & Configuration

Server:

*   set server settings as described in [Tendermint Metrics](https://docs.tendermint.com/v0.34/tendermint-core/metrics.html)
*   restart the `cosmovisor` service
*   configure a firewall (if any) so incoming HTTP connections to prometheus port would be allowed
*   verify (e.g. `curl IP:PORT`): you should see the metrics along with the values

Client:

*   install [Prometheus](https://prometheus.io/docs/prometheus/latest/getting_started/).
*   prepare prometheus configuration file
    *   please check Prometheus docs for more details
    *   example: [bench/prometheus.yml](../bench/prometheus.yml)

## Prometheus Run

Run

```bash
prometheus --config.file=prometheus.yml
```

And open `http://localhost:9090/` to query and monitor the server-side metrics.
