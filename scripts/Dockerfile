################################################################################
# Image of a machine with installed swagger-combine, buf and ignite.
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

# Install swagger-combine
RUN npm install -g swagger-combine

# Install buf v1.29.0
RUN cd tmp
RUN curl -L https://github.com/bufbuild/buf/releases/download/v1.29.0/buf-Linux-x86_64 -o buf && \
    chmod 777 buf && \
    cp buf /usr/local/bin
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.19.0

# Install ignite v0.27.2
RUN curl -L https://github.com/ignite/cli/releases/download/v0.27.2/ignite_0.27.2_linux_amd64.tar.gz -o ignite.targ.gz && \
    tar -xvzf ignite.targ.gz && \
    cp ignite /usr/local/bin

ENV PATH="/root/go/bin:${PATH}"
ENV PATH="${PATH}:/usr/local/bin"

WORKDIR /dcl