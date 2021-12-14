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
# FIXME issue 99: that is dirty and invalidates the cache almost always during the dev
COPY . .

RUN make

############################
# STEP 2 build an image
############################
FROM alpine:latest

COPY --from=builder /go/bin/dcld /usr/bin/dcld
COPY --from=builder /go/bin/dlv /usr/bin/dlv

VOLUME /root/.dcl

EXPOSE 26656 26657

STOPSIGNAL SIGTERM
