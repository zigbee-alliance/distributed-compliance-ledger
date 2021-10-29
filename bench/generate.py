#!/usr/bin/env python
# -*- coding: utf-8 -*-
#
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

import sys
import os
import yaml
import json
import subprocess
import tempfile
from struct import pack
from pathlib import Path

from render import render


DCLCLI = "dclcli"

DEF_ACCOUNT_N_START = 4
DEF_SEQUENCE_START = 0

ACCOUNT_N_START_F = "account-number-start"
SEQUENCE_START_F = "sequence-number-start"
QUERIES_F = "q"

TEST_PASSWORD = "test1234"

MODEL_INFO_PREFIX = 1
VENDOR_PRODUCTS_PREFIX = 2


def pack_model_info_key(vid, pid):
    return pack('<bhh', MODEL_INFO_PREFIX, vid, pid)


def run_shell_cmd(cmd, **kwargs):
    _kwargs = dict(
        check=True,
        universal_newlines=True,
        stdout=subprocess.PIPE,
        stderr=subprocess.PIPE
    )

    _kwargs.update(kwargs)

    if not _kwargs.get("shell") and type(cmd) is str:
        cmd = cmd.split()

    try:
        return subprocess.run(cmd, **_kwargs)
    except (subprocess.CalledProcessError, FileNotFoundError) as exc:
        raise RuntimeError(f"command '{cmd}' failed: {exc.stderr}") from exc


def run_for_json_res(cmd, **kwargs):
    return json.loads(run_shell_cmd(cmd, **kwargs).stdout)


def to_cli_args(**kwargs):
    res = []
    for k, v in kwargs.items():
        k = "--{}".format(k.replace("_", "-"))
        res.extend([k, str(v)])
    return res


def yaml_dump(
    data,
    stream=None,
    width=1,
    indent=4,
    default_flow_style=False,
    canonical=False,
    **kwargs
):
    return yaml.safe_dump(
        data,
        stream,
        default_flow_style=default_flow_style,
        canonical=canonical,
        width=width,
        indent=indent,
        **kwargs
    )


def resolve_users():
    return {u["name"]: u for u in run_for_json_res([DCLCLI, "keys", "list"])}


def txn_generate(u_address, txn_t_cls, txn_t_cmd, **params):
    cmd = [DCLCLI, "tx", txn_t_cls, txn_t_cmd]
    params["from"] = u_address
    cmd += to_cli_args(**params)
    cmd.append("--generate-only")
    return run_shell_cmd(cmd).stdout


def txn_sign(u_address, account, sequence, f_path):
    cmd = [DCLCLI, "tx", "sign"]
    params = {"from": u_address}
    cmd += to_cli_args(
        account_number=account, sequence=sequence, gas="auto", **params
    )
    cmd.extend(["--offline", f_path])
    cmd = f"echo '{TEST_PASSWORD}' | {' '.join(cmd)}"
    return run_shell_cmd(cmd, shell=True).stdout


def txn_encode(f_path):
    cmd = [DCLCLI, "tx", "encode", f_path]
    return run_shell_cmd(cmd).stdout


ENV_PREFIX = "DCLBENCH_"


def main():
    render_ctx = {
        k.split(ENV_PREFIX)[1].lower(): v
        for k, v in os.environ.items()
        if k.startswith(ENV_PREFIX)
    }

    # TODO argument parsing using argparse
    spec_yaml = render(sys.argv[1], ctx=render_ctx)
    spec = yaml.safe_load(spec_yaml)

    try:
        out_file = Path(sys.argv[2]).resolve()
    except IndexError:
        out_file = None

    account_n_start = spec["defaults"].get(
        ACCOUNT_N_START_F, DEF_ACCOUNT_N_START)
    sequence_start = spec["defaults"].get(
        SEQUENCE_START_F, DEF_SEQUENCE_START)

    users = resolve_users()

    res = {}

    account_n = account_n_start
    with tempfile.TemporaryDirectory() as tmpdirname:

        for user, u_data in spec["users"].items():
            res[user] = []
            tmp_file = (Path(tmpdirname) / user).resolve()

            u_address = users[user]["address"]

            sequence = sequence_start
            for q in u_data[QUERIES_F]:
                q_id, q_data = next(iter(q.items()))
                q_cls, q_t, q_cmd = q_id.split("/")

                if q_cls == "tx":
                    tmp_file.write_text(
                        txn_generate(u_address, q_t, q_cmd, **q_data)
                    )
                    # XXX by some reason pipe to encode doesn't work
                    tmp_file.write_text(
                        txn_sign(u_address, account_n, sequence, str(tmp_file))
                    )

                    txn_encoded = txn_encode(str(tmp_file))
                    res[user].append(txn_encoded.strip().strip('"'))
                    sequence += 1
                else:
                    raise ValueError("Unexpected query class: {q_cls}")

            if out_file:
                print(f"User {user}: done")

            account_n += 1

    # TODO optimize for big data
    if out_file is None:
        print(yaml_dump(res))
    else:
        with out_file.open('w') as fd:
            yaml_dump(res, fd)


if __name__ == "__main__":
    main()
