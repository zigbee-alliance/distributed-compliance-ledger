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

import testinfra.utils.ansible_runner

testinfra_hosts = testinfra.utils.ansible_runner.AnsibleRunner(
    os.environ["MOLECULE_INVENTORY_FILE"]
).get_hosts("all")
DCLD_HOME = "/var/lib/dcl/.dcl/"


def test_binary_version(host):
    assert host.run_test(DCLD_HOME + "cosmovisor/genesis/bin/dcld version").succeeded


def test_service(host):
    svc = host.file("/etc/systemd/system/cosmovisor.service")
    assert svc.exists
    for prop in [
        "User=cosmovisor",
        "Group=dcl",
        'Environment="DAEMON_HOME=/var/lib/dcl/.dcl" "DAEMON_NAME=dcld"',
        "ExecStart=/usr/bin/cosmovisor start",
    ]:
        assert prop in svc.content_string
