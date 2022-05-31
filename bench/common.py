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

import subprocess
import random
import string
import tempfile
import json
import time

from pathlib import Path


DCLCLI = "dcld"
DCL_CHAIN_ID = "dclchain"


def create_vendor_account(vendor_name, vid, trustee_account_name):
    try:
        keys_delete(vendor_name)
    except Exception:
        print("We don't remove that user, because that user does not exist in dcld")

    keys_add(vendor_name)

    # Get a Vendor address and pubkey
    vendor_address = keys_show_address(vendor_name)
    vendor_pubkey = keys_show_pubkey(vendor_name)

    # Send to request to another node to propose
    cmd = [
        DCLCLI,
        "tx",
        "auth",
        "propose-add-account",
        "--address=" + vendor_address,
        "--pubkey=" + vendor_pubkey,
        "--roles=Vendor",
        "--vid=" + str(vid),
        "--from=" + trustee_account_name,
        "--yes",
    ]
    result = run_shell_cmd(cmd)

    while "account sequence mismatch" in str(result):
        time.sleep(random.randint(1, 20))
        print(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
        print("[WARNING! READDING VENDOR ACCOUNT!]")
        result = run_shell_cmd(cmd)
        print(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

    print(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
    print("[SUCCESS! VENDOR ACCOUNT HAS ADDED!]")
    print(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")


def generate_txns(model_id, model_sequence, vendor_address, vendor_id, account_number):
    with tempfile.TemporaryDirectory() as tmpdirname:
        random_file_name = generate_random_name()
        tmp_file = (Path(tmpdirname) / random_file_name).resolve()
        
        tmp_file.write_text(create_model(vendor_address, model_id, vendor_id))
        tmp_file.write_text(txn_sign(vendor_address, account_number, model_sequence, str(tmp_file)))
        
        return txn_encode(str(tmp_file))



def keys_delete(key_name):
    cmd = [DCLCLI, "keys", "delete", key_name, "--yes"]
    return run_shell_cmd(cmd).stdout

def keys_add(key_name):
    cmd = [DCLCLI, "keys", "add", key_name]
    return run_shell_cmd(cmd).stdout

def keys_show_address(key_name):
    cmd = [DCLCLI, "keys", "show", key_name, "-a"]
    return run_shell_cmd(cmd).stdout.rstrip("\n")

def keys_show_pubkey(key_name):
    cmd = [DCLCLI, "keys", "show", key_name, "-p"]
    return run_shell_cmd(cmd).stdout.rstrip("\n")

def get_account_number(vendor_address):
    cmd = [DCLCLI, "query", "auth", "account", "--address", vendor_address]
    result = run_shell_cmd(cmd).stdout

    json_result = json.loads(result)

    return int(json_result["base_account"]["account_number"])



def create_model(vendor_address, current_model_id, vendor_id):
    cmd = [
        DCLCLI,
        "tx",
        "model",
        "add-model",
        "--vid=" + str(vendor_id),
        "--pid=" + str(current_model_id),
        "--deviceTypeID=" + str(current_model_id),
        "--productName=ProductName" + str(current_model_id),
        "--productLabel=ProductLabel" + str(current_model_id),
        "--partNumber=PartNumber" + str(current_model_id),
        "--from=" + vendor_address,
        "--yes",
        "--generate-only",
    ]
    return run_shell_cmd(cmd).stdout

def txn_sign(vendor_address, account, sequence_n, f_path):
    cmd = [DCLCLI, "tx", "sign", "--chain-id", DCL_CHAIN_ID]
    params = {"from": vendor_address}
    cmd += to_cli_args(account_number=account, sequence=sequence_n, gas="auto", **params)
    cmd.extend(["--offline", f_path])
    
    return run_shell_cmd(cmd).stdout

def txn_encode(f_path):
    cmd = [DCLCLI, "tx", "encode", f_path]
    return run_shell_cmd(cmd).stdout



def generate_random_name():
    letters = string.ascii_lowercase
    return ''.join(random.choice(letters) for i in range(10)) 

def generate_random_number():
    return random.randint(1000, 65000)



def run_shell_cmd(cmd, **kwargs):
    _kwargs = dict(
        check=True,
        universal_newlines=True,
        stdout=subprocess.PIPE,
        stderr=subprocess.PIPE,
    )

    _kwargs.update(kwargs)

    if not _kwargs.get("shell") and type(cmd) is str:
        cmd = cmd.split()

    try:
        return subprocess.run(cmd, **_kwargs)
    except (subprocess.CalledProcessError, FileNotFoundError) as exc:
        raise RuntimeError(f"command '{cmd}' failed: {exc.stderr}") from exc

def to_cli_args(**kwargs):
    res = []
    for k, v in kwargs.items():
        k = "--{}".format(k.replace("_", "-"))
        res.extend([k, str(v)])
    return res
