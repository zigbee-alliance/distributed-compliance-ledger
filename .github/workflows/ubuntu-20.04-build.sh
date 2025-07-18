#!/bin/bash

set -eux

_HOME="$HOME"
BIN_NAME="${BIN_NAME:-dcld}"
COSMOVISOR_BIN_NAME="${COSMOVISOR_BIN_NAME:-cosmovisor}"
BUILD_DIR=build

if [[ -z "${COSMOVISOR_VERSION:-}" ]]; then
    >&2 echo "COSMOVISOR_VERSION is not defined"
    exit 1
fi

GO_ROOT="$(realpath $(dirname $(which go))/..)"
_GO_ROOT="/usr/local/go"
_GO_ROOT_BIN="$_GO_ROOT/bin"

! read -r -d '' script << EOF
    set -eux

    export PATH="$_GO_ROOT_BIN:\$PATH"
    GOPATH="\${GOPATH:-\${HOME}/go}"
    GOBIN="\${GOBIN:-\${GOPATH}/bin}"

    go version
    go env GOBIN
    go env GOPATH

    # install system deps
    apt-get update
    apt-get install -y software-properties-common
    DEBIAN_FRONTEND=noninteractive add-apt-repository -y ppa:git-core/ppa
    apt-get update
    apt-get install -y make g++ git ca-certificates
    git version

    git config --global --add safe.directory "\$PWD"
    make build
    "$BUILD_DIR/$BIN_NAME" version

    go install "cosmossdk.io/tools/cosmovisor/cmd/cosmovisor@v$COSMOVISOR_VERSION"
    cp "\$GOPATH/bin/$COSMOVISOR_BIN_NAME" "$BUILD_DIR"
    chown -R "$(id -u):$(id -g)" "$BUILD_DIR"
EOF


docker run --rm -w "$PWD" \
    -e GOROOT="$_GO_ROOT" \
    -e HOME="$_HOME" \
    -v "$HOME":"$_HOME" \
    -v "$GO_ROOT:$_GO_ROOT" \
    ubuntu:20.04 bash -c "$script"
