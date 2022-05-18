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

import json
import logging
import string
import random
import common

from locust import HttpUser, events, task


DEFAULT_TARGET_HOST = "http://localhost:26657"
DEFAULT_REST_HOST = "http://localhost:26640"



dcl_hosts = []
dcl_rest_hosts = []

logger = logging.getLogger("dclbench")


@events.init_command_line_parser.add_listener
def init_paraser(parser):
    # Set `include_in_web_ui` to False if you want to hide from the web UI
    parser.add_argument(
        "--dcl-hosts",
        metavar="DCL_HOSTS",
        include_in_web_ui=True,
        default=DEFAULT_TARGET_HOST,
        help="Comma separated list of DCL hosts to target",
    )
    parser.add_argument(
        "--dcl-rest-hosts",
        metavar="DCL_REST_HOSTS",
        include_in_web_ui=True,
        default=DEFAULT_REST_HOST,
        help="Comma separated list of DCL REST hosts to target",
    )


@events.test_start.add_listener
def _(environment, **kw):
    logger.info(f"dcl-hosts: {environment.parsed_options.dcl_hosts}")

    if environment.parsed_options.dcl_hosts:
        dcl_hosts.extend(environment.parsed_options.dcl_hosts.split(","))

    if environment.parsed_options.dcl_rest_hosts:
        dcl_rest_hosts.extend(environment.parsed_options.dcl_rest_hosts.split(","))

class DCLWriteUser(HttpUser):
    host = ""
    weight = 5
    
    vendor_account_name = common.generate_random_name()
    vendor_id = common.generate_random_number()
    vendor_account_number = int
    vendor_account_address = string

    model_id = 1  
    model_sequence = 0
    

    @task
    def add_model(self):
        payload = {
            "method": "broadcast_tx_sync", 
            "params": 
                {
                    "tx": 
                        common.generate_txns(
                            self.model_id, 
                            self.model_sequence,  
                            self.vendor_account_address, 
                            self.vendor_id,
                            self.vendor_account_number, 
                        )
                }, 
            "id": 1
            }

        with self.client.post(
                f"{self.host}/",
                json.dumps(payload),
                name="write transactions",
                catch_response=True,
            ) as response:
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
                
                self.model_sequence += 1
                self.model_id += 1

    
    def on_start(self):
        # Get RPC endpoint
        if dcl_hosts:
            self.host = random.choice(dcl_hosts)
        else:
            self.host = DEFAULT_TARGET_HOST
        
        common.create_vendor_account(self.vendor_account_name, self.vendor_id)

        self.vendor_account_address = common.keys_show_address(self.vendor_account_name)
        self.vendor_account_number = common.get_account_number(self.vendor_account_address)
