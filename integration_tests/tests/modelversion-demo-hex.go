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

func ModelVersion_demo_hex() {

	// Preparation of Actors

	vid_in_hex_format := 0xA13
	pid_in_hex_format := 0xA11
	vid := 2579
	pid := 2577

	vendor_account := "vendor_account_" + string(vid_in_hex_format)
	fmt.Printf("Create Vendor account - " + vendor_account)
	cli.Create_new_vendor_account(vendor_account, string(vid_in_hex_format))

	// Create a new model version

	fmt.Printf("Add Model with VID: " + string(vid_in_hex_format) + "PID: " + string(pid_in_hex_format))
	cmd, _ := cli.AddModel(string(vid_in_hex_format), string(pid_in_hex_format), "1", "TestProduct", " Test Product", "1", "0", vendor_account)
	result, _ := cli.Execute("bash", "-c", "echo test1234 | ", cmd)
	cli.Check_response(result, "\"code\": 0")

	cli.Test_divider()

	sv := cli.Random()
	fmt.Printf("Create a Device Model Version with VID: " + string(vid) + "PID: " + string(pid) + "SV: " + string(sv))
	cmd, _ = cli.AddModelVersion(1, 10, 1, string(vid), string(pid), string(sv), 1, vendor_account)
	result, _ = cli.Execute("bash", "-s", "echo test1234 | ", cmd)
	fmt.Printf(result)
	cli.Check_response(result, "\"code\": 0")

	cli.Test_divider()

	// Query the model version
	fmt.Printf("Query Device Model Version with VID: " + string(vid_in_hex_format) + "PID: " + string(pid_in_hex_format) + "SV: " + string(sv))
	result, _ = cli.ModelVersion(string(vid_in_hex_format), string(pid_in_hex_format), string(sv))
	fmt.Printf(result)
	cli.Check_response(result, "\"vid\": ", string(vid))
	cli.Check_response(result, "\"pid\": ", string(pid))
	cli.Check_response(result, "\"softwareVersion\": ", string(sv))
	cli.Check_response(result, "\"softwareVersionString\": \"1\"")
	cli.Check_response(result, "\"cdVersionNumber\": 1")
	cli.Check_response(result, "\"softwareVersionValid\": ", "true")
	cli.Check_response(result, "\"minApplicableSoftwareVersion\": 1")
	cli.Check_response(result, "\"maxApplicableSoftwareVersion\": 10")

	cli.Test_divider()

	// Query all models versions
	fmt.Printf("Query all model versions with VID: " + string(vid_in_hex_format) + "PID: " + string(pid_in_hex_format))
	result, _ = cli.AllModelVersion(string(vid_in_hex_format), string(pid_in_hex_format))
	fmt.Printf(result)
	cli.Check_response(result, "\"vid\": ", string(vid))
	cli.Check_response(result, "\"pid\": ", string(pid))
	cli.Check_response(result, "\"softwareVersions\"")
	cli.Check_response(result, string(sv))

	cli.Test_divider()

	// Query non existent model version
	fmt.Printf("Query Device Model Version with VID: " + string(vid_in_hex_format) + "PID: " + string(pid_in_hex_format) + "SV: 123456")
	result, _ = cli.ModelVersion(string(vid_in_hex_format), string(pid_in_hex_format), "123456")
	cli.Check_response(result, "Not Found")

	cli.Test_divider()

	// Query non existent model versions
	vid1_in_hex_format := 0xA14
	pid1_in_hex_format := 0xA15
	fmt.Printf("Query all Device Model Versions with VID: " + string(vid1_in_hex_format) + "PID: " + string(pid1_in_hex_format))
	result, _ = cli.AllModelVersion(string(vid1_in_hex_format), string(pid1_in_hex_format))
	cli.Check_response(result, "Not Found")

	cli.Test_divider()

	// Update the existing model version
	fmt.Printf("Update Device Model Version with VID: " + string(vid_in_hex_format) + "PID: " + string(pid_in_hex_format) + "SV: " + string(sv))
	cmd, _ = cli.UpdateModelVersion(string(vid_in_hex_format), string(pid_in_hex_format), 2, 10, string(sv), false, vendor_account)
	result, _ = cli.Execute("bash", "-c", "echo test1234 | ", cmd)
	cli.Check_response(result, "\"code\": 0")

	cli.Test_divider()

	// Query Updated model version
	fmt.Printf("Query updated Device Model Version with VID: " + string(vid_in_hex_format) + "PID: " + string(pid_in_hex_format) + "SV: " + string(sv))
	result, _ = cli.ModelVersion(string(vid_in_hex_format), string(pid_in_hex_format), string(sv))
	fmt.Printf(result)
	cli.Check_response(result, "\"vid\": ", string(vid))
	cli.Check_response(result, "\"pid\": ", string(pid))
	cli.Check_response(result, "\"softwareVersion\": ", string(sv))
	cli.Check_response(result, "\"softwareVersionString\": \"1\"")
	cli.Check_response(result, "\"cdVersionNumber\": 1")
	cli.Check_response(result, "\"softwareVersionValid\": ", "false")
	cli.Check_response(result, "\"minApplicableSoftwareVersion\": 2")
	cli.Check_response(result, "\"maxApplicableSoftwareVersion\": 10")

	cli.Test_divider()

	// Add second model version
	sv2 := cli.Random()
	fmt.Printf("Create a Second Device Model Version with VID: " + string(vid_in_hex_format) + "PID: " + string(pid_in_hex_format) + "SV: " + string(sv2))
	cmd, _ = cli.AddModelVersion(1, 10, 1, string(vid), string(pid), string(sv2), 1, vendor_account)
	result, _ = cli.Execute("bash", "-c", "echo test1234 | ", cmd)
	fmt.Printf(result)
	cli.Check_response(result, "\"code\": 0")

	cli.Test_divider()

	// Query all model versions
	fmt.Printf("Query all model versions with VID: " + string(vid_in_hex_format) + "PID: " + string(pid_in_hex_format))
	result, _ = cli.AllModelVersion(string(vid_in_hex_format), string(pid_in_hex_format))
	fmt.Printf(result)
	cli.Check_response(result, "\"vid\": ", string(vid))
	cli.Check_response(result, "\"pid\": ", string(pid))
	cli.Check_response(result, "\"softwareVersions\"")
	cli.Check_response(result, string(sv))
	cli.Check_response(result, string(sv2))

	cli.Test_divider()

	// Create model version with vid belonging to another vendor
	fmt.Printf("Create a Device Model Version with VID: " + string(vid_in_hex_format) + "PID: " + string(pid_in_hex_format) + "SV: " + string(sv) + "from a different vendor account")
	new_vid_in_hex_format := "0xB17"
	different_vendor_account := "vendor_account_" + string(new_vid_in_hex_format)
	cli.Create_new_vendor_account(different_vendor_account, new_vid_in_hex_format)
	cmd, _ = cli.AddModelVersion(1, 10, 1, string(vid), string(pid), string(sv), 1, different_vendor_account)
	result, _ = cli.Execute("bash", "-c", "echo test1234 | ", cmd)
	cli.Check_response_and_report(result, "transaction should be signed by a vendor account containing the vendorID ", string(vid))

	cli.Test_divider()

	// Update model version with vid belonging to another vendor
	fmt.Printf("Update a Device Model Version with VID: " + string(vid_in_hex_format) + "PID: " + string(pid_in_hex_format) + "SV: " + string(sv) + "from a different vendor account")
	cmd, _ = cli.UpdateModelVersion(string(vid_in_hex_format), string(pid_in_hex_format), sv, 0, "", false, string(different_vendor_account))
	result, _ = cli.Execute("bash", "-c", "echo test1234 | ", cmd)
	cli.Check_response_and_report(result, "transaction should be signed by a vendor account containing the vendorID ", string(vid))

}
