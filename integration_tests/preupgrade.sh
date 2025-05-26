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

info_file="$DAEMON_HOME/cosmovisor/upgrade_info.csv"

if [ -f $info_file ]; then
    exit 1
fi

echo "Writing preupgrade info to $info_file" >&2

current_dcld_path=$(readlink -f $DAEMON_HOME/cosmovisor/current)
upgrade_name=$1
upgrade_height=$2

echo "$upgrade_name,$current_dcld_path,$upgrade_height" > "$info_file"

exit 0