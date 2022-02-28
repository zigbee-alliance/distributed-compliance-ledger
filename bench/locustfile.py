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

import yaml
import json
import time
import random
import logging
from pathlib import Path
from locust import HttpUser, task, events, LoadTestShape
# import locust_plugins

DEFAULT_TARGET_HOST = "http://localhost:26657"
DEFAULT_REST_HOST = "http://localhost:26640"

txns = []
dcl_hosts = []
dcl_rest_hosts = []
users_done = {}

logger = logging.getLogger('dclbench')


@events.init_command_line_parser.add_listener
def init_paraser(parser):
    parser.add_argument(
        "--dcl-users", type=int,
        include_in_web_ui=True,
        default=10, help="Peak number of concurrent Locust users")
    parser.add_argument(
        "--dcl-spawn-rate", type=int,
        include_in_web_ui=True,
        default=1, help="Rate to spawn users at (users per second)")
    parser.add_argument(
        "--dcl-txn-file", type=str, env_var="LOCUST_TXN_FILE",
        include_in_web_ui=True,
        default="./txns", help="Path to a file with write transactions")
    # Set `include_in_web_ui` to False if you want to hide from the web UI
    parser.add_argument(
        "--dcl-hosts", metavar="DCL_HOSTS",
        include_in_web_ui=True,
        default=DEFAULT_TARGET_HOST,
        help="Comma separated list of DCL hosts to target")


@events.test_start.add_listener
def _(environment, **kw):
    logger.info(f"dcl-users: {environment.parsed_options.dcl_users}")
    logger.info(f"dcl-spawn-rate: {environment.parsed_options.dcl_spawn_rate}")
    logger.info(f"dcl-txn-file: {environment.parsed_options.dcl_txn_file}")
    logger.info(f"dcl-hosts: {environment.parsed_options.dcl_hosts}")

    if environment.parsed_options.dcl_hosts:
        dcl_hosts.extend(environment.parsed_options.dcl_hosts.split(","))

    _txns = yaml.safe_load(
        Path(environment.parsed_options.dcl_txn_file).read_text())

    # user only necessary number of users
    txns.extend(list(_txns.items())[:environment.parsed_options.dcl_users])

    for user in txns:
        users_done[user[0]] = False

    logger.info(f"Users number: {len(txns)}, users: {list(users_done)}")


@events.test_stop.add_listener
def _(environment, **kw):
    # reset the stat as we considering next runs
    logger.info("Resetting users progress")
    users_done.clear()

# curl --header "Content-Type: application/json" --request POST --data  localhost:26657  # noqa


# the most reliable (explored so far) how to interrupt locust once
# all requests are done:
#   - it explicitely tells locust to stop keeping users count
#   - other ways (call locust env to quite or use `-i` iteration
#     max count param from locust_plugins) don't work well since
#     after stopping the users locust tries to re-spawn them again
#     to keep the initial number of them
class DCLTestShape(LoadTestShape):
    def tick(self):
        logger.debug(f"{users_done}, users {self.runner.user_count}")
        if users_done and all(users_done.values()):
            logger.info("All users are done")
            return None
        else:
            return (
                self.runner.environment.parsed_options.dcl_users,
                self.runner.environment.parsed_options.dcl_spawn_rate
            )


class DCLUser(HttpUser):
    username = None
    txns = None
    host = ""
    rest_host = ""
    # DEFAULT_TARGET_HOST

    #@task
    def add_model(self):
        logger.debug(f"{self.username}: {len(self.txns or [])} txns remain")
        if self.txns:
            txn = self.txns.pop(0)
            payload = {
                "method": "broadcast_tx_sync",
                "params": {"tx": txn},
                "id": 1
            }
            with self.client.post(
                f"{self.host}/", json.dumps(payload), name="write-txn",
                catch_response=True
            ) as response:
                # logger.debug(f"{self.username}: response {response.__dict__}")
                logger.debug(f"{self.username}: response {response.text}")
                payload = json.loads(response.text)
                if "error" in payload:
                    response.failure(json.dumps(payload["error"]))
                elif "result" in payload:
                    if payload["result"].get("code") != 0:
                        error = dict(payload["result"])
                        # to keep failure stat condensed
                        error.pop("hash", None)
                        response.failure(json.dumps(error))
                else:
                    response.failure("malformed txn: {response.text}")

            time.sleep(0.1)
        else:
            if self.username:
                users_done[self.username] = True
            time.sleep(1)

    def on_start(self):
        global txns
        if len(txns):
            self.username, self.txns = txns.pop(0)
            if dcl_hosts:
                self.host = random.choice(dcl_hosts)
            else:
                self.host = DEFAULT_TARGET_HOST
            logger.info(
                f"{self.username}: started, num txns {len(self.txns)},"
                f" target host {self.host}")
        else:
            logger.warning("unexpected user: no more data")
