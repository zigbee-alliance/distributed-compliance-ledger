# Copyright 2022 DSR Corporation
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

import os

import testinfra.utils.ansible_runner

testinfra_hosts = testinfra.utils.ansible_runner.AnsibleRunner(
    os.environ["MOLECULE_INVENTORY_FILE"]
).get_hosts("all")


def test_run(host):
    assert host.run_test("/usr/bin/dcld version")


def test_configuration(host):
    config = host.file("/home/dcld/.dcl/config/")
    assert config.exists
    assert config.is_directory
    assert config.user == "dcld"

    assert host.file("/home/dcld/.dcl/jack.info").exists
    config_files = host.file("/home/dcld/.dcl/config").listdir()
    for filename in ["app", "client", "config"]:
        assert (filename + ".toml") in config_files

    assert "test-net" in host.file("/home/dcld/.dcl/config/client.toml").content_string
