# Copyright 2020 DSR Corporation
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

################################################################################
# Image of a node in DCL pool.
################################################################################

############################
# STEP 1 build cosmovisor
############################
FROM ubuntu:20.04 AS builder

ARG GO_VERSION
ENV GO_VERSION=1.20

RUN apt-get update --fix-missing
RUN apt-get install -y wget git gcc

RUN wget -P /tmp "https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz"

RUN tar -C /usr/local -xzf "/tmp/go${GO_VERSION}.linux-amd64.tar.gz"
RUN rm "/tmp/go${GO_VERSION}.linux-amd64.tar.gz"

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

RUN go install cosmossdk.io/tools/cosmovisor/cmd/cosmovisor@v1.5.0

############################
# STEP 2 build node image
############################
FROM ubuntu:20.04

COPY --from=builder /go/bin/cosmovisor /usr/bin/

# test user
ARG TEST_USER
ENV TEST_USER=${TEST_USER:-dcl}

#ARG TEST_USER_GROUP
#ENV TEST_USER_GROUP=${TEST_USER_GROUP:-dcl}

ARG TEST_UID
ENV TEST_UID=${TEST_UID:-1000}
#ARG gid=1000
RUN adduser --disabled-password --uid ${TEST_UID} --home /var/lib/${TEST_USER} --gecos 'DCLedger user' ${TEST_USER}
ENV DAEMON_HOME=/var/lib/${TEST_USER}/.dcl
ENV DAEMON_NAME=dcld
ENV DAEMON_ALLOW_DOWNLOAD_BINARIES=true
ENV COSMOVISOR_CUSTOM_PREUPGRADE=preupgrade.sh

RUN apt-get update
RUN apt-get install -y ca-certificates

RUN update-ca-certificates

VOLUME /var/lib/${TEST_USER}

EXPOSE 26656 26657 1317 26660 8888

STOPSIGNAL SIGTERM

USER ${TEST_USER}

COPY integration_tests/node_helper.sh /var/lib/${TEST_USER}/
COPY deployment/dcld_manager.sh /var/lib/${TEST_USER}/
COPY deployment/preupgrade.sh /var/lib/${TEST_USER}/

ENV PATH=$PATH:${DAEMON_HOME}/cosmovisor/current/bin

WORKDIR /var/lib/${TEST_USER}