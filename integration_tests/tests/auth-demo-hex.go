//!/bin/bash
// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"

	cli "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli"
)

func auth_demo_hex() {
	var user string
	user = cli.Random_string(user)
	fmt.Printf(user + " generates keys")
	aux, _ := cli.AddKeys(user)
	cmd := "(echo " + cli.Passphrase + "; echo " + cli.Passphrase + ") | " + aux
	result, _ := cli.Execute("bash", "-c", cmd)

	cli.Test_divider()

	fmt.Printf("Get key info for " + user)
	cmd = "echo " + cli.Passphrase + " | " + aux
	result, _ = cli.Execute("bash", "-c", cmd)
	cli.Check_response(result, "\"name\": \"", user, "\"")

	cli.Test_divider()

	user_address, _ := cli.GetUserAddress(user, cli.Passphrase)
	user_pubkey, _ := cli.GetUserPubKey(user, cli.Passphrase)

	jack_address, _ := cli.GetUserAddress("jack", cli.Passphrase)
	//alice_address, _ := cli.GetUserAddress("alice", cli.Passphrase)
	//bob_address, _ := cli.GetUserAddress("bob", cli.Passphrase)
	//anna_address, _ := cli.GetUserAddress("anna", cli.Passphrase)

	vid_in_hex_format := 0xA13
	pid_in_hex_format := 0xA11
	vid := 2579
	pid := 2577

	fmt.Printf("Jack proposes account for " + user)
	cmd, _ = cli.ProposeAddAccount("Jack is proposing this account", user_address, user_pubkey, "Vendor", string(vid_in_hex_format), "jack")
	result, _ = cli.Execute("bash", "-c", "echo ", cli.Passphrase, " | ", cmd)
	cli.Check_response(result, "\"code\": 0")

	cli.Test_divider()

	fmt.Printf("Get all active accounts. $user account in the list because has enough approvals")
	result, _ = cli.AllProposedAccounts()
	cli.Check_response(result, "\"address\": \"", user_address, "\"")

	cli.Test_divider()

	fmt.Printf("Get an account for " + user)
	result, _ = cli.Account(user_address)
	cli.Check_response(result, "\"address\": \"", user_address, "\"")
	cli.Check_response_and_report(result, jack_address, "json")
	cli.Check_response_and_report(result, "info: Jack is proposing this account", "json")

	cli.Test_divider()

	fmt.Printf("Get an proposed account for $user is not found")
	result, _ = cli.ProposeAccount(user_address)
	cli.Check_response(result, "Not Found")

	cli.Test_divider()

	fmt.Printf("Get all proposed accounts. " + user + "account is not in the list")
	result, _ = cli.AllProposedAccounts()
	cli.Check_response(result, "\"[\"]")

	cli.Test_divider()

	productName := "Device #1"
	fmt.Printf(string(user)+" adds Model with VID: "+string(vid_in_hex_format)+"PID: ", pid)
	cmd, _ = cli.AddModel(string(vid_in_hex_format), string(pid_in_hex_format), productName, "Device Description", string(0), string(12), string(12), user_address)
	result, _ = cli.Execute("bach", "-c", "echo test1234 | ", cmd)
	cli.Check_response_and_report(result, "\"code\":", "0")

	cli.Test_divider()

	vid_plus_one_in_hex_format := 0xA14
	vidPlusOne := vid_in_hex_format + 1
	fmt.Printf(string(user) + "adds Model with a VID: " + string(vid_plus_one_in_hex_format) + "PID: " + string(pid_in_hex_format) + " This fails with Permission denied as the VID is not associated with this vendor account.")
	cmd, _ = cli.AddModel(string(vid_in_hex_format), string(pid_in_hex_format), productName, "Device Description", string(0), string(12), string(12), user_address)
	result, _ = cli.Execute("bash", "-c", "echo test1234 | ", cmd, "|| true")
	cli.Check_response_and_report(result, "transaction should be signed by a vendor account containing the vendorID ", string(vidPlusOne))

	cli.Test_divider()

	fmt.Printf(string(user) + " updates Model with VID: " + string(vid_in_hex_format) + "PID: " + string(pid_in_hex_format))
	cmd, _ = cli.UpdateModel(string(vid_in_hex_format), string(pid_in_hex_format), productName, "Device Description", string(12), user_address)
	result, _ = cli.Execute("bash", "-c", "echo test1234 | ", cmd)

	cli.Test_divider()

	fmt.Printf("Get Model with VID: "+string(vid_in_hex_format)+" PID: ", pid_in_hex_format)
	result, _ = cli.GetModel(string(vid_in_hex_format), string(pid_in_hex_format))
	cli.Check_response(result, "\"vid\": ", string(vid))
	cli.Check_response(result, "\"pid\": ", string(pid))
	cli.Check_response(result, "\"productName\": \"", productName, "\"")

	fmt.Printf("PASSED")
}
