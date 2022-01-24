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

############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

# Git is required for fetching the dependencies,
# make is required for building.
# build-base is needed because gcc is required by delve
RUN apk update && apk add --no-cache git make build-base

# Build Delve - This is helpful if you want to do remote debugging by attaching to one of the docker containers remotely
RUN go get github.com/go-delve/delve/cmd/dlv

WORKDIR /go/src/dc-ledger
COPY app ./app/
COPY cmd ./cmd/
COPY testutil ./testutil/
COPY utils ./utils/
COPY x ./x/
COPY go.mod go.sum Makefile ./
COPY docs/docs.go ./docs/
COPY docs/static ./docs/static/

ARG DCL_VERSION
ARG DCL_COMMIT
RUN make

############################
# STEP 2 build an image
############################
FROM alpine:latest

COPY --from=builder /go/bin/dcld /usr/bin/dcld
COPY --from=builder /go/bin/dlv /usr/bin/dlv

# test user
ARG TEST_USER
ENV TEST_USER=${TEST_USER:-dcl}

#ARG TEST_USER_GROUP
#ENV TEST_USER_GROUP=${TEST_USER_GROUP:-dcl}

ARG TEST_UID
ENV TEST_UID=${TEST_UID:-1000}
#ARG gid=1000
RUN adduser -D -u ${TEST_UID} -h /var/lib/${TEST_USER} -g 'DCLedger user' ${TEST_USER}

VOLUME /var/lib/${TEST_USER}

EXPOSE 26656 26657 1317 26660 8888

STOPSIGNAL SIGTERM

USER ${TEST_USER}
WORKDIR /var/lib/${TEST_USER}
