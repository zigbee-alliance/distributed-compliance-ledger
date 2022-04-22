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


def test_configuration(host):
    all_variables = host.ansible.get_variables()
    config = host.file(DCLD_HOME + "/config/")
    assert config.exists
    assert config.is_directory

    config_files = host.file(DCLD_HOME + "/config").listdir()
    for filename in ["app", "client", "config"]:
        assert (filename + ".toml") in config_files

    assert "chain_id" in all_variables
    assert (
        all_variables["chain_id"]
        in host.file(DCLD_HOME + "/config/client.toml").content_string
    )
