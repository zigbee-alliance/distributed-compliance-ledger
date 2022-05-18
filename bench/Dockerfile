FROM locustio/locust:2.4.1

USER root

RUN apt-get update && apt-get install -y curl && pip install pyyaml

ARG DCLD_VERSION=v0.9.0
ARG DCLD_NODE=https://on.matterprotocol.com:26657
ARG DCLD_CHAIN_ID=test-net2

RUN curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/${DCLD_VERSION}/dcld

RUN cp dcld /usr/local/bin
RUN chmod a+x /usr/local/bin/dcld


USER locust
WORKDIR /home/locust

RUN dcld config node ${DCLD_NODE}
RUN dcld config chain-id ${DCLD_CHAIN_ID}
RUN dcld config keyring-backend test
RUN dcld config broadcast-mode block
RUN dcld config output json

RUN dcld keys add test

COPY 6f0d716d1d5c0ef842c5e2e61c21fdc79d9a9079.address /home/locust/.dcl/keyring-test
COPY validator-key.info /home/locust/.dcl/keyring-test