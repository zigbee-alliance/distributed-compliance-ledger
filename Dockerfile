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
FROM golang:alpine AS builder

# git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

RUN go install github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor@v1.0.0

############################
# STEP 2 build node image
############################
FROM alpine:latest

COPY --from=builder /go/bin/cosmovisor /usr/bin/

# test user
ARG TEST_USER
ENV TEST_USER=${TEST_USER:-dcl}

#ARG TEST_USER_GROUP
#ENV TEST_USER_GROUP=${TEST_USER_GROUP:-dcl}

ARG TEST_UID
ENV TEST_UID=${TEST_UID:-1000}
#ARG gid=1000
RUN adduser -D -u ${TEST_UID} -h /var/lib/${TEST_USER} -g 'DCLedger user' ${TEST_USER}

ENV DAEMON_HOME=/var/lib/${TEST_USER}/.dcl
ENV DAEMON_NAME=dcld

VOLUME /var/lib/${TEST_USER}

EXPOSE 26656 26657 1317 26660 8888

STOPSIGNAL SIGTERM

USER ${TEST_USER}

ENV PATH=$PATH:${DAEMON_HOME}/cosmovisor/current/bin

WORKDIR /var/lib/${TEST_USER}
