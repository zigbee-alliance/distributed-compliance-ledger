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

source integration_tests/utils/utils.sh

set -eEuo pipefail

MASTER_UPGRADE_IMAGE="dcld-build-master"
MASTER_UPGRADE_CONTAINER_NAME="$MASTER_UPGRADE_IMAGE-inst"
NEW_OBSERVER_CONTAINER_NAME="new-observer"
VALIDATOR_DEMO_CONTAINER_NAME="validator-demo"

cleanup_containers() {
    echo "Cleanup containers"
    cleanup_container "$MASTER_UPGRADE_CONTAINER_NAME"
    cleanup_container "$NEW_OBSERVER_CONTAINER_NAME"
    cleanup_container "$VALIDATOR_DEMO_CONTAINER_NAME"
}

trap cleanup_containers EXIT

cleanup_containers

source integration_tests/upgrade/01-test-upgrade-initialize-0.12.sh
source integration_tests/upgrade/02-test-upgrade-0.12-rollback.sh
source integration_tests/upgrade/03-test-upgrade-0.12-to-1.2.sh
source integration_tests/upgrade/04-test-upgrade-1.2-rollback.sh
source integration_tests/upgrade/05-test-upgrade-1.2-to-1.4.3.sh
source integration_tests/upgrade/06-test-upgrade-1.4.3-to-1.4.4.sh
source integration_tests/upgrade/07-test-upgrade-1.4.4-to-1.5.0.sh
source integration_tests/upgrade/08-test-upgrade-1.5.0-to-master.sh
source integration_tests/upgrade/09-add-new-node-after-upgrade.sh
