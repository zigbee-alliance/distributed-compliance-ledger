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

func model_demo() {

	// Preparation of Actors

	vid := cli.Random()
	pid := cli.Random()
	vendor_account := "vendor_account_" + string(vid)
	fmt.Printf("Create Vendor account - ", vendor_account)
	cli.Create_new_vendor_account(vendor_account, string(vid))

	cli.Test_divider()

	// Body

	fmt.Printf("Query non existent model")
	result, _ := cli.GetModel(string(vid), string(pid))
	cli.Check_response(result, "Not Found")
	fmt.Printf(result)

	cli.Test_divider()

	fmt.Printf("Query non existent Vendor Models")
	result, _ = cli.VendorModel(string(vid))
	cli.Check_response(result, "Not Found")
	fmt.Printf(result)

	cli.Test_divider()

	fmt.Printf("Request all models must be empty")
	result, _ = cli.AllProposedModels()
	cli.Check_response(result, "\"[\"]")
	fmt.Printf(result)

	cli.Test_divider()

	productLabel := "Device #1"
	fmt.Printf("Add Model with VID: " + string(vid) + "PID: " + string(pid))
	cmd, _ := cli.AddModel(string(vid), string(pid), "1", "TestProduct", productLabel, "1", "0", vendor_account)
	result, _ = cli.Execute("bash", "-c", "echo test1234 | ", cmd)
	cli.Check_response(result, "\"code\": 0")
	fmt.Printf(result)

	cli.Test_divider()

	fmt.Printf("Get Model with VID: " + string(vid) + "PID: " + string(pid))
	result, _ = cli.GetModel(string(vid), string(pid))
	cli.Check_response(result, "\"vid\": ", string(vid))
	cli.Check_response(result, "\"pid\": ", string(pid))
	cli.Check_response(result, "\"productLabel\": \"", productLabel, "\"")
	fmt.Printf(result)

	cli.Test_divider()

	sv := 1
	cd_version_num := 10
	fmt.Printf("Create Model Versions with VID: " + string(vid) + "PID: " + string(pid) + "SoftwareVersion: " + string(sv))
	cmd, _ = cli.AddModelVersion(vid, pid, sv, "1", "15", "sv", cd_version_num, vendor_account)
	result, _ = cli.Execute("bash", "-c", "echo test1234 | ", cmd)
	cli.Check_response(result, "\"code\": 0")
	fmt.Printf(result)

	cli.Test_divider()

	fmt.Printf("Get all models")
	result, _ = cli.AllProposedModels()
	cli.Check_response(result, "\"vid\": ", string(vid))
	cli.Check_response(result, "\"pid\": ", string(pid))
	fmt.Printf(result)

	cli.Test_divider()

	fmt.Printf("Get Vendor Models with VID: " + string(vid))
	result, _ = cli.VendorModel(string(vid))
	cli.Check_response(result, "\"pid\": ", string(pid))
	fmt.Printf(result)

	cli.Test_divider()

	fmt.Printf("Update Model with VID: " + string(vid) + "PID: " + string(pid) + "with new description")
	description := "New Device Description"
	cmd, _ = cli.UpdateModel(string(vid), string(pid), vendor_account, "", "", string(description))
	result, _ = cli.Execute("bash", "-c", "echo test1234 | ", cmd)
	cli.Check_response(result, "\"code\": 0")
	fmt.Printf(result)

	cli.Test_divider()

	fmt.Printf("Get Model with VID: " + string(vid) + "PID: " + string(pid))
	result, _ = cli.GetModel(string(vid), string(pid))
	cli.Check_response(result, "\"vid\": ", string(vid))
	cli.Check_response(result, "\"pid\": ", string(pid))
	cli.Check_response(result, "\"productLabel\": \"", description, "\"")
	fmt.Printf(result)

	cli.Test_divider()

	fmt.Printf("Update Model with VID: " + string(vid) + "PID: " + string(pid) + "modifying supportURL")
	supportURL := "https://newsupporturl.test"
	cmd, _ = cli.UpdateModel(string(vid), string(pid), vendor_account, "", "", string(supportURL))
	result, _ = cli.Execute("bash", "-c", "echo test1234 | ", cmd)
	cli.Check_response(result, "\"code\": 0")
	fmt.Printf(result)

	cli.Test_divider()

	fmt.Printf("Get Model with VID: " + string(vid) + "PID: " + string(pid))
	result, _ = cli.GetModel(string(vid), string(pid))
	cli.Check_response(result, "\"vid\": ", string(vid))
	cli.Check_response(result, "\"pid\": ", string(pid))
	cli.Check_response(result, "\"supportUrl\": \"", supportURL, "\"")
	fmt.Printf(result)

	cli.Test_divider()

	fmt.Printf("Delete Model with VID: " + string(vid) + "PID: " + string(pid))
	result, _ = cli.DeleteModel(string(vid), string(pid), vendor_account)
	fmt.Printf(result)

	cli.Test_divider()

	fmt.Printf("Query non existent model")
	result, _ = cli.GetModel(string(vid), string(pid))
	cli.Check_response(result, "Not Found")
	fmt.Printf(result)

	cli.Test_divider()

	fmt.Printf("Query model versions for deleted model")
	result, _ = cli.ModelVersion(string(vid), string(pid), string(sv))
	cli.Check_response(result, "Not Found")
	fmt.Printf(result)
}
