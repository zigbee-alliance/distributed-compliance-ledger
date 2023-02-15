# Copyright 2022
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
# Image of a machine with installed swagger-combine and starport.
################################################################################

FROM alpine:3.15 as builder

ARG STARPORT_VERSION

ENV STARPORT_VERSION=${STARPORT_VERSION:-v0.19.3}
#ENV STARPORT_VERSION=${STARPORT_VERSION:-dev}

RUN apk --no-cache add \
  make=4.3-r0 \
  go=1.17.9-r0 \
  git=2.34.2-r0 \
  bash=5.1.16-r0

# Install starport
WORKDIR /tmp
RUN git clone https://github.com/tendermint/starport.git \
  && if [ "$STARPORT_VERSION" != "dev"  ]; then \
    git -C /tmp/starport checkout -b ${STARPORT_VERSION:-v0.19.3} tags/${STARPORT_VERSION:-v0.19.3}; \
  fi \
  && make -C /tmp/starport build

FROM alpine:3.15

ENV PROTOC_GEN_GRPC_GATEWAY_VERSION=v2.8.0
ENV PROTOC_GEN_OPENAPIV2_VERSION=v2.8.0
ENV PROTOC_GEN_SWAGGER_VERSION=v1.16.0
ENV PATH="/root/go/bin:${PATH}"

COPY --from=builder /tmp/starport/dist/starport /usr/local/bin/starport

RUN apk --no-cache add \
  bash=5.1.16-r0 \
  git=2.34.2-r0 \
  go=1.17.9-r0 \
  protobuf=3.18.1-r1 \
  protobuf-dev=3.18.1-r1 \
  npm=8.1.3-r0 \
  && npm install -g swagger-combine@1.4.0

# Install grpc-gateway tools

RUN go install \
  github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@${PROTOC_GEN_GRPC_GATEWAY_VERSION} \
  github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@${PROTOC_GEN_OPENAPIV2_VERSION} \
  && go install \
  github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@${PROTOC_GEN_SWAGGER_VERSION}

WORKDIR /dcl
