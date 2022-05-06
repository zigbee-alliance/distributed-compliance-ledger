#!/bin/bash
# Copyright 2022 Samsung Corporation
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

set -o nounset
set -o errexit
set -o pipefail
if [[ "${DEBUG:-false}" == "true" ]]; then
    set -o xtrace
fi

function info {
    _print_msg "INFO" "$1"
}

function warn {
    _print_msg "WARN" "$1"
}

function error {
    _print_msg "ERROR" "$1"
    exit 1
}

function _print_msg {
    echo "$(date +%H:%M:%S) - $1: $2"
}

function get_github_latest_release {
    version=""
    attempt_counter=0
    max_attempts=5

    until [ "$version" ]; do
        url_effective=$(curl -sL -o /dev/null -w '%{url_effective}' "https://github.com/$1/releases/latest")
        if [ "$url_effective" ]; then
            version="${url_effective##*/}"
            break
        elif [ ${attempt_counter} -eq ${max_attempts} ];then
            echo "Max attempts reached"
            exit 1
        fi
        attempt_counter=$((attempt_counter+1))
        sleep $((attempt_counter*2))
    done
    echo "${version#v}"
}

function main {
    readonly dcl_home=${DCL_HOME:-"$HOME/.dcl"}
    local version=${VERSION:-$(get_github_latest_release zigbee-alliance/distributed-compliance-ledger)}
    local dest=${DEST:-"$dcl_home/cosmovisor/upgrades/v$version/bin"}

    if [ ! -f "$dest/dcld" ] || [[ "$("$dest/dcld" version)" != "$version" ]]; then
        info "Installing DCL client $version version..."

        distro="macos"
        if [ "$(uname)" == "Linux" ]; then
            # shellcheck disable=SC1091
            source /etc/os-release || source /usr/lib/os-release
            distro="${ID,,}"
        fi

        info "Downloading the DCL binary for $distro..."
        mkdir -p "$dest"
        url="https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v$version/dcld.$distro.tar.gz"
        if command -v wget > /dev/null; then
            wget -q -c "$url" -O - | sudo tar -xz -C "$dest"
        elif command -v curl > /dev/null; then
            curl -s "$url" | sudo tar -xz -C "$dest"
        fi

        sudo chown -R "$(whoami)" "$dcl_home"
        sudo chmod a+x "$dest/dcld"
    fi

    if [ -n "${SHA256SUM:-}" ] && command -v sha256sum > /dev/null; then
        info "Validating $SHA256SUM checksum..."
        sha256sum --check --quiet <(echo "${SHA256SUM} $dest/dcld")
    else
        warn "Checksum wasn't performed"
    fi
}

# Validators
info "Validating $USER permissions"
if ! sudo -n "true"; then
    error "passwordless sudo is needed for '$(whoami)' user."
fi

info "Validating cosmovisor service"
if ! systemctl is-active --quiet cosmovisor; then
    error "cosmovisor service is not active"
fi

info "Validating cosmovisor user"
# shellcheck disable=SC2009
if [[ "$(whoami)" != "$(ps -o comm,user | grep cosmovisor | awk '{print $2}')" ]]; then
    error "This script needs to be executed with comovisor user"
fi

main
