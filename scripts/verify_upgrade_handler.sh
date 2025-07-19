#!/usr/bin/env bash

set -euo pipefail

VER="\"$1\""

if echo "$VER" | grep 'dev"' &>/dev/null; then
    echo "Develop release, skip upgrade handle verify"
elif echo "$VER" | grep 'pre"' &>/dev/null; then
    echo "Pre-release, skip upgrade handle verify"
elif grep 'SetUpgradeHandler(' -a1 ../app/app.go | grep "$VER" &>/dev/null; then
    echo "Upgrade handle $VER exists"
else
    echo "Upgrade handle not $VER exists"
    exit 1
fi

exit 0