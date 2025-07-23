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

# this script is the entry point of docker contain (used only in tests)

set -euo pipefail

timestamp=0
max_restart_cnt=10
restart_cnt=0

RUN_CMD="$HOME"/./cosmovisor_start.sh
ln -s "$HOME"/cosmovisor_preupgrade.sh "$HOME"/.dcl/cosmovisor/cosmovisor_preupgrade.sh

cosmovisor_stop() {
    pkill -f cosmovisor_start
    sleep 1
    exit 0
}

if env | grep GOCOVERDIR; then
    trap cosmovisor_stop EXIT
fi

while $RUN_CMD; do
    # in tests, the number of restarts is very limited, mainly when an upgrade occurs
    # exit with error if there are any unexpected unnecessary restarts
    if (( restart_cnt > max_restart_cnt )) ; then
        echo "The maximum number ($max_restart_cnt) of restart attempts has been reached, exit"
        exit 1
    fi

    sleep 1
    echo "'$RUN_CMD' stopped, restarting"

   restart_cnt=$((restart_cnt+1))
done

echo "'$RUN_CMD' crashed with an error code"

