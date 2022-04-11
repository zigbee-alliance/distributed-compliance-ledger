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


def test_accounts_creation(host):
    all_variables = host.ansible.get_variables()
    assert "accounts" in all_variables
    for account in all_variables["accounts"]:
        assert "passphrase" in account
        assert "name" in account
        cmd = host.run(
            f"echo {account['passphrase']}"
            f" | /var/lib/dcl/.dcl/cosmovisor/genesis/bin/dcld keys show {account['name']}"
            f" --home {DCLD_HOME} --output json"
        )
        assert cmd.succeeded
        assert len(cmd.stdout) > 0
        key_name = json.loads(cmd.stdout)
        for key in ["name", "type", "address", "pubkey"]:
            assert key in key_name
        assert key_name["name"] == account["name"]
        assert key_name["type"] == "local"
        assert host.file(f"{DCLD_HOME}{account['name']}.info").exists
