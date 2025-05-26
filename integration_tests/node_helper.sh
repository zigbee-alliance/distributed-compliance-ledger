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

# this script is the entry point of docker contain

timestamp=0
cooldown_time=10

RUN_CMD=/var/lib/dcl/./dcld_manager.sh
ln -s /var/lib/dcl/preupgrade.sh /var/lib/dcl/.dcl/cosmovisor/preupgrade.sh

while $RUN_CMD; do
    now=$(date +%s)

    if (( now < timestamp + cooldown_time )); then
        echo "Unacceptable to restart '$RUN_CMD' more than every $cooldown_time seconds, exit" >&2
        exit 1
    fi

    timestamp=$now
    sleep 1
    echo "'$RUN_CMD' stopped, restarting" >&2
done

echo "'$RUN_CMD' crashed with an error code" >&2

