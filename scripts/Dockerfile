################################################################################
# Image of a machine with installed swagger-combine and starport.
################################################################################

FROM node

# Install golang
ARG GO_VERSION
ENV GO_VERSION=${GO_VERSION:-1.17.2}
ENV BASH_ENV=/etc/bashrc
ENV PATH="${PATH}:/usr/local/go/bin"
RUN curl -L https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz -o /tmp/go.tar.gz \
    && rm -rf /usr/local/go \
    && tar -C /usr/local -xzf /tmp/go.tar.gz \ 
    && go version \
    && rm -f /tmp/go.tar.gz

# Install protoc
ARG PROTOC_VERSION
ENV PROTOC_VERSION=${PROTOC_VERSION:-3.19.4}
RUN PROTOC_ZIP=protoc-${PROTOC_VERSION}-linux-x86_64.zip \
    && curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/${PROTOC_ZIP} \
    && unzip -o $PROTOC_ZIP -d /usr/local bin/protoc \
    && unzip -o $PROTOC_ZIP -d /usr/local 'include/*' \
    && rm -f $PROTOC_ZIP

# Install grpc-gateway tools
ENV PROTOC_GEN_GRPC_GATEWAY_VERSION=v2.8.0
ENV PROTOC_GEN_OPENAPIV2_VERSION=v2.8.0
ENV PROTOC_GEN_SWAGGER_VERSION=v1.16.0

RUN go install \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@${PROTOC_GEN_GRPC_GATEWAY_VERSION} \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@${PROTOC_GEN_OPENAPIV2_VERSION}

RUN go install \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@${PROTOC_GEN_SWAGGER_VERSION}

RUN npm install -g swagger-combine


# Install starport
ARG STARPORT_VERSION
ENV STARPORT_VERSION=${STARPORT_VERSION:-0.19.3}
#ENV STARPORT_VERSION=${STARPORT_VERSION:-dev}

RUN if [ "$STARPORT_VERSION" = "dev" ]; then \
      curl -L https://github.com/tendermint/starport/archive/refs/heads/develop.zip -o /tmp/starport.zip && \
      cd /tmp \
        && unzip starport.zip \
        && cd starport-develop \
        && make build \
        && cp ./dist/starport /usr/local/bin; \
    else \
      curl https://get.starport.network/starport@v${STARPORT_VERSION}! -o /tmp/startport \
      && bash /tmp/startport \
      && rm /tmp/startport; \
    fi  

ENV PATH="/root/go/bin:${PATH}"
ENV PATH="${PATH}:/usr/local/bin"

WORKDIR /dcl