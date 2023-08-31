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

// Check add model with fieds VID/PID in hex format and
// get model with fields VID/PID in hex format

// Preperation of Actors

package main

import (
	"fmt"

	cli "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli"
)

func model_demo_hex() {

	vid_in_hex_format := 0xA13
	pid_in_hex_format := 0xA11
	vid := 2579
	pid := 2577

	vendor_account := "vendor_account_" + string(vid_in_hex_format)
	fmt.Printf("Create Vendor account - " + vendor_account)
	cli.Create_new_vendor_account(vendor_account, string(vid_in_hex_format))

	cli.Test_divider()

	// Body

	fmt.Printf("Query non existent model")
	result, _ := cli.GetModel(string(vid_in_hex_format), string(pid_in_hex_format))
	cli.Check_response(result, "Not Found")
	fmt.Printf(result)

	cli.Test_divider()

	fmt.Printf("Query non existent Vendor Models")
	result, _ = cli.VendorModel(string(vid_in_hex_format))
	cli.Check_response(result, "Not Found")
	fmt.Printf(result)

	cli.Test_divider()

	fmt.Printf("Request all models must be empty")
	result, _ = cli.AllProposedModels()
	cli.Check_response(result, "\"[\"]")
	fmt.Printf(result)

	cli.Test_divider()

	productLabel := "Device #1"
	fmt.Printf("Add Model with VID: " + string(vid_in_hex_format) + " PID: " + string(pid_in_hex_format))
	cmd, _ := cli.AddModel(string(vid_in_hex_format), string(pid_in_hex_format), string(1), "TestProduct", productLabel, string(1), string(0), vendor_account)
	result, _ = cli.Execute("bash", "-c", "echo test1234 | ", cmd)
	cli.Check_response(result, "\"code\": 0")
	fmt.Printf(result)

	cli.Test_divider()

	fmt.Printf("Get Model with VID: " + string(vid_in_hex_format) + " PID: " + string(pid_in_hex_format))
	result, _ = cli.GetModel(string(vid_in_hex_format), string(pid_in_hex_format))
	cli.Check_response(result, "\"vid\": ", string(vid))
	cli.Check_response(result, "\"pid\": ", string(pid))
	cli.Check_response(result, "\"productLabel\": \""+productLabel+"\"")
	fmt.Printf(result)

	cli.Test_divider()

	fmt.Printf("Get all models")
	result, _ = cli.AllProposedModels()
	cli.Check_response(result, "\"vid\": ", string(vid))
	cli.Check_response(result, "\"pid\": ", string(pid))
	fmt.Printf(result)

	cli.Test_divider()

	fmt.Printf("Get Vendor Models with VID: " + string(vid_in_hex_format))
	result, _ = cli.VendorModel(string(vid_in_hex_format))
	cli.Check_response(result, "\"pid\": ", string(pid))
	fmt.Printf(result)

	cli.Test_divider()

	fmt.Printf("Update Model with VID: " + string(vid_in_hex_format) + " PID: " + string(pid_in_hex_format) + " with new description")
	description := "New Device Description"
	cmd, _ = cli.UpdateModel(string(vid_in_hex_format), string(pid_in_hex_format), description, "", "", string(vendor_account))
	result, _ = cli.Execute("bash", "-c", "echo test1234 | ", cmd)
	cli.Check_response(result, "\"code\": 0")
	fmt.Printf(result)

	cli.Test_divider()

	fmt.Printf("Get Model with VID: " + string(vid_in_hex_format) + "PID: " + string(pid_in_hex_format))
	result, _ = cli.GetModel(string(vid_in_hex_format), string(pid_in_hex_format))
	cli.Check_response(result, "\"vid\": ", string(vid))
	cli.Check_response(result, "\"pid\": ", string(pid))
	cli.Check_response(result, "\"productLabel\": \"", description, "\"")
	fmt.Printf(result)

	cli.Test_divider()

	fmt.Printf("Update Model with VID: " + string(vid_in_hex_format) + "PID: " + string(pid_in_hex_format) + "modifying supportURL")
	supportURL := "https://newsupporturl.test"
	cmd, _ = cli.UpdateModel(string(vid_in_hex_format), string(pid_in_hex_format), vendor_account, "", "", string(supportURL))
	result, _ = cli.Execute("bash", "-c", "echo test1234 | ", cmd)
	cli.Check_response(result, "\"code\": 0")
	fmt.Printf(result)

	cli.Test_divider()
	fmt.Printf("Get Model with VID: " + string(vid_in_hex_format) + " PID: " + string(pid_in_hex_format))
	result, _ = cli.GetModel(string(vid_in_hex_format), string(pid_in_hex_format))
	cli.Check_response(result, "\"vid\": ", string(vid))
	cli.Check_response(result, "\"pid\": ", string(pid))
	cli.Check_response(result, "\"supportUrl\": \"", supportURL, "\"")
	fmt.Printf(result)

	cli.Test_divider()

	fmt.Printf("Delete Model with VID: " + string(vid_in_hex_format) + " PID: " + string(pid_in_hex_format))
	result, _ = cli.DeleteModel(string(vid_in_hex_format), string(pid_in_hex_format), string(vendor_account))
	fmt.Printf(result)

	cli.Test_divider()

	fmt.Printf("Query non existent model")
	result, _ = cli.GetModel(string(vid_in_hex_format), string(pid_in_hex_format))
	cli.Check_response(result, "Not Found")
	fmt.Printf(result)

	cli.Test_divider()
}
