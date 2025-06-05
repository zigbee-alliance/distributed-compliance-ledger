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

# this script the script will write file with information before the upcoming upgrade

set -euo pipefail

info_file="$DAEMON_HOME/cosmovisor/upgrade.info"

#check if upgrade is running

if [ "$(grep "^UPGRADE_IS_RUNNING=" "$info_file" | cut -d'=' -f2-)" = 1 ]; then
    echo -n "The upgrade has already started, exit"
    exit 1
fi

echo -n "Writing preupgrade info to $info_file"

prev_dcld_path=$(readlink -f "$DAEMON_HOME/cosmovisor/current")
upgrade_plan="$1"
upgrade_height="$2"
upgrade_is_running=1

echo -e "UPGRADE_PLAN=$upgrade_plan\nPREV_DCLD_PATH=$prev_dcld_path\nUPGRADE_HEIGHT=$upgrade_height\nUPGRADE_IS_RUNNING=$upgrade_is_running\n" > "$info_file"

exit 0