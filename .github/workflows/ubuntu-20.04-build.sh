#!/bin/bash

set -eux

_HOME="$HOME"
BIN_NAME="${BIN_NAME:-dcld}"

GO_ROOT="$(realpath $(dirname $(which go))/..)"
_GO_ROOT="/usr/local/go"
_GO_BIN="$_GO_ROOT/bin"

! read -r -d '' script << EOF
    set -eux

    export PATH="$_GO_BIN:\$PATH"
    go version

    # install system deps
    apt-get update
    apt-get install -y software-properties-common
    DEBIAN_FRONTEND=noninteractive add-apt-repository -y ppa:git-core/ppa
    apt-get update
    apt-get install -y make g++ git
    git version

    git config --global --add safe.directory "\$PWD"
    make build

    chown "$(id -u):$(id -g)" "build/$BIN_NAME"
    "build/$BIN_NAME" version
EOF


docker run --rm -w "$PWD" \
    -e GOROOT="$_GO_ROOT" \
    -e HOME="$_HOME" \
    -v "$HOME":"$_HOME" \
    -v "$GO_ROOT:$_GO_ROOT" \
    ubuntu:20.04 bash -c "$script"
