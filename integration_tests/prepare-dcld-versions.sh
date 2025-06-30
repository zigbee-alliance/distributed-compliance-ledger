#!/bin/bash
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

set -euo pipefail

LOG_PREFIX="[prepare dcld versions] "
TMP_DCLD_BINS_DIR=/tmp/dcld_bins

log() {
  echo "${LOG_PREFIX}$1"
}

mkdir -p $TMP_DCLD_BINS_DIR

for version in 0.12.0 0.12.1 1.2.2 1.4.3 1.4.4; do
    DCLD_BIN_FILE="$TMP_DCLD_BINS_DIR/dcld_v$version"

    if [[ ! -f $DCLD_BIN_FILE || "$(stat -c %a $DCLD_BIN_FILE)" != "775" ]]; then
        log "$DCLD_BIN_FILE not found! The download is running"

        wget -q --show-progress -O $DCLD_BIN_FILE "https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v$version/dcld"
        chmod ugo+x $DCLD_BIN_FILE
    fi
done