# Copyright 2022 Samsung Corporation
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

import json
import os
import time

import testinfra.utils.ansible_runner

testinfra_hosts = testinfra.utils.ansible_runner.AnsibleRunner(
    os.environ["MOLECULE_INVENTORY_FILE"]
).get_hosts("all")
dcld_cmd = "/var/lib/dcl/.dcl/cosmovisor/current/bin/dcld {} --home=/var/lib/dcl/.dcl/"


def test_service(host):
    svc = host.service("cosmovisor")
    assert svc.is_enabled
    assert svc.is_running


def test_node_not_jailed(host):
    # Get the list of all nodes
    cmd = host.run(dcld_cmd.format("query validator all-nodes"))
    assert len(cmd.stdout) > 0
    query = json.loads(cmd.stdout)
    assert "validator" in query
    for validator in query["validator"]:
        assert validator["power"] == 10
        assert validator["jailed"] is False


def test_block_height_changes(host):
    latest_block_height = 0
    for x in range(3):
        # Get the node status
        cmd = host.run(dcld_cmd.format("status --node tcp://127.0.0.1:26657"))
        assert len(cmd.stdout) > 0
        status = json.loads(cmd.stdout)
        assert "SyncInfo" in status
        assert "latest_block_height" in status["SyncInfo"]
        assert latest_block_height != status["SyncInfo"]["latest_block_height"]
        latest_block_height = status["SyncInfo"]["latest_block_height"]
        time.sleep(5)


def test_total_participating_nodes(host):
    # Get the list of nodes participating in the consensus for the last block
    cmd = host.run(dcld_cmd.format("query tendermint-validator-set"))
    assert len(cmd.stdout) > 0
    query = json.loads(cmd.stdout)
    assert "total" in query
    assert query["total"] == "1"
