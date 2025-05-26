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

trap 'kill $(jobs -p)' EXIT

info_file="$DAEMON_HOME/cosmovisor/upgrade_info.csv"
do_restore=0

if [ -f $info_file ]; then
    current_dcld_path=$(cut -d, -f2 $info_file)
    upgrade_height=$(cut -d, -f3 $info_file)
    do_restore=1
    rm -f $info_file
fi

while true; do
    sleep 1
    if [ ! -f $info_file ]; then
        continue
    fi

    upgrade_name=$(cut -d, -f1 $info_file)

    while [ -f $info_file ]; do
        sleep 1
        if dcld query upgrade applied $upgrade_name &>/dev/null; then
            rm -f $info_file
        fi
    done
done &

if [ "$do_restore" = 1 ]; then
    echo "Restoring old dcld version" >&2

    rm -f $DAEMON_HOME/cosmovisor/current
    rm -f $DAEMON_HOME/data/upgrade-info.json
    ln -s $current_dcld_path $DAEMON_HOME/cosmovisor/current

    echo "Execution 'cosmovisor run start --unsafe-skip-upgrades $upgrade_height'" >&2
    cosmovisor run start --unsafe-skip-upgrades $upgrade_height
else
    echo "Execution 'cosmovisor run start'" >&2
    cosmovisor run start
fi

if [[ $? -ne 0 && -f $info_file ]]; then
    exit 0
fi

exit 1
