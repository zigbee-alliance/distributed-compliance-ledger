################################################################################
# Image of a machine with installed swagger-combine and starport.
################################################################################

FROM node

# Install golang
ARG GO_VERSION
ENV GO_VERSION=${GO_VERSION:-1.20}
ENV BASH_ENV=/etc/bashrc
ENV PATH="${PATH}:/usr/local/go/bin"
RUN curl -L https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz -o /tmp/go.tar.gz \
    && rm -rf /usr/local/go \
    && tar -C /usr/local -xzf /tmp/go.tar.gz \ 
    && go version \
    && rm -f /tmp/go.tar.gz

# Install tools
ENV BUF_BUILD=v1.29.0
ENV PROTOC_GEN_GRPC_GATEWAY_VERSION=v2.19.0
ENV PROTOC_GEN_OPENAPIV2_VERSION=v2.19.0
ENV PROTOC_GEN_SWAGGER_VERSION=v1.16.0
ENV PROTOC_GEN_GO_COSMOS=v1.4.10
ENV PROTOC_GEN_GO_PULSAR=v1.0.0-beta.2
ENV PROTOC_GEN_GO=v1.32.0
ENV PROTOC_GEN_GO_GRPC=v1.1.0

RUN go install \
        github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@${PROTOC_GEN_GRPC_GATEWAY_VERSION} \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@${PROTOC_GEN_OPENAPIV2_VERSION} \
        github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@${PROTOC_GEN_SWAGGER_VERSION} \

RUN go install github.com/bufbuild/buf/cmd/buf@${BUF_BUILD}
RUN go install github.com/cosmos/gogoproto/protoc-gen-gocosmos@${PROTOC_GEN_GO_COSMOS}
RUN go install github.com/cosmos/cosmos-proto/cmd/protoc-gen-go-pulsar@${PROTOC_GEN_GO_PULSAR}
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@${PROTOC_GEN_GO_GO}
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@${PROTOC_GEN_GO_GRPC}

RUN npm install -g swagger-combine


ENV PATH="/root/go/bin:${PATH}"
ENV PATH="${PATH}:/usr/local/bin"

WORKDIR /dcl