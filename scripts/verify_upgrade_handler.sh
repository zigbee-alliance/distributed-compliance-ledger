#!/usr/bin/env bash

set -euo pipefail

VER="\"$1\""

if echo "$VER" | grep 'dev"' &>/dev/null; then
    echo "Develop release, skipping the upgrade handler verify"
elif echo "$VER" | grep 'pre"' &>/dev/null; then
    echo "Pre-release, skipping the upgrade handler verify"
elif grep 'SetUpgradeHandler(' -a1 ./app/app.go | grep "$VER" &>/dev/null; then
    echo "Upgrade handler $VER exists"
else
    echo "Upgrade handler $VER not exists"
    exit 1
fi

exit 0