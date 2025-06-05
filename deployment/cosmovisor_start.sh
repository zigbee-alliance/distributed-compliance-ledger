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

# this script runs dcld or restores the previous version and runs, deletes the upgrade information file if the upgrade was successful

set -euo pipefail

trap 'kill $(jobs -p)' EXIT

info_file="$DAEMON_HOME/cosmovisor/upgrade.info"
upgrade_is_running=0

clear_upgrade_flag() {
    sed -i -e "s/UPGRADE_IS_RUNNING=1/UPGRADE_IS_RUNNING=0/" "$info_file"
}

if [ -f "$info_file" ]; then
    prev_dcld_path="$(grep "^PREV_DCLD_PATH=" "$info_file" | cut -d'=' -f2-)"
    upgrade_height="$(grep "^UPGRADE_HEIGHT=" "$info_file" | cut -d'=' -f2-)"
    upgrade_is_running="$(grep "^UPGRADE_IS_RUNNING=" "$info_file" | cut -d'=' -f2-)"

    # better to clean flag before the cycle below
    if [ "$upgrade_is_running" = 1 ]; then
        clear_upgrade_flag
    fi
else
    # create an empty file
    echo -e "UPGRADE_PLAN=\nPREV_DCLD_PATH=\nUPGRADE_HEIGHT=\nUPGRADE_IS_RUNNING=\n" > "$info_file"
fi

while true; do
    sleep 1

    #check if upgrade is running
    if [ "$(grep "^UPGRADE_IS_RUNNING=" "$info_file" | cut -d'=' -f2-)" = 1 ]; then
        upgrade_plan="$(grep "^UPGRADE_PLAN=" "$info_file" | cut -d'=' -f2-)"

        if dcld query upgrade applied "$upgrade_plan" &>/dev/null; then
            clear_upgrade_flag
        fi
    fi
done &

exit_code=0

#if upgrade is currently in progress, need to roll back
if [ "$upgrade_is_running" = 1 ]; then
    echo "Restoring old dcld version"

    rm -f "$DAEMON_HOME/cosmovisor/current"
    rm -f "$DAEMON_HOME/data/upgrade-info.json"
    ln -s "$prev_dcld_path" "$DAEMON_HOME"/cosmovisor/current

    echo "Execution 'cosmovisor run start --unsafe-skip-upgrades $upgrade_height'"
    cosmovisor run start --unsafe-skip-upgrades "$upgrade_height" || exit_code="$?"
else
    echo "Execution 'cosmovisor run start'"
    cosmovisor run start || exit_code="$?"
fi

# normal exit if there is an error during upgrade (rollback will be performed after rerun this script)
if [[ "$exit_code" -ne 0 && "$(grep "^UPGRADE_IS_RUNNING=" "$info_file" | cut -d'=' -f2-)" = 1 ]]; then
    exit 0
fi

exit 1
