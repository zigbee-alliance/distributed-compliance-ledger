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

func Model_negative_cases() {

	// Preparation of Actors

	fmt.Printf("Create regular account")
	cli.Create_new_account("certification_house", "CertificationCenter")

	vid := cli.Random()
	pid := cli.Random()
	softwareVersionString := cli.Random()
	vendor_account := "vendor_account_" + string(vid)
	fmt.Printf("Vendor account - " + vendor_account)
	cli.Create_new_vendor_account(vendor_account, string(vid))

	cli.Test_divider()

	fmt.Printf("Create CertificationCenter account")
	cli.Create_new_account("zb_account", "CertificationCenter")

	// Body

	// Ledger side errors

	fmt.Printf("Add Model with VID: " + string(vid) + "PID: " + string(pid) + "  : Not Vendor")
	cmd, _ := cli.AddModel(string(vid), string(pid), "1", "TestProduct", "TestingProductLabel", "1", "0", "certification_house")
	result, _ := cli.Execute("bash", "-c", "echo test1234 | ", cmd)
	cli.Check_response_and_report(result, "\"code\":", "4")
	fmt.Printf(result)

	cli.Test_divider()

	vid1 := cli.Random()
	fmt.Printf("Add Model with VID: " + string(vid1) + "PID: " + string(pid) + ": Vendor ID does not belong to vendor")
	cmd, _ = cli.AddModel(string(vid1), string(pid), "1", "TestProduct", "TestingProductLabel", "1", "0", vendor_account)
	result, _ = cli.Execute("bash", "-c", "echo test1234 | ", cmd)
	cli.Check_response_and_report(result, "\"code\":", "4")
	fmt.Printf(result)

	cli.Test_divider()

	fmt.Printf("Add Model with VID: " + string(vid) + "PID: " + string(pid) + ": Twice")
	cmd, _ = cli.AddModel(string(vid), string(pid), "1", "TestProduct", "TestingProductLabel", "1", "0", vendor_account)
	result, _ = cli.Execute("bash", "-c", "echo test1234 | ", cmd)
	cli.Check_response_and_report(result, "\"code\": ", "501")
	fmt.Printf(result)

	cli.Test_divider()

	sv := cli.Random()
	svs := cli.Random()
	fmt.Printf("Create a Device Model Version with VID: " + string(vid) + "PID: " + string(pid) + "SV: " + string(sv))
	cmd, _ = cli.AddModelVersion(1, 10, 1, string(vid), string(pid), string(sv), softwareVersionString, vendor_account)
	result, _ = cli.Execute("bash", "-c", "echo test1234 | ", cmd)
	fmt.Printf(result)
	cli.Check_response(result, "\"code\": 0")

	certification_date := "2020-01-01T00:00:01Z"
	zigbee_certification_type := "zigbee"
	//matter_certification_typ := "matter"
	cd_certificate_id := "123"
	cmd, _ = cli.CertifyModel(string(vid), string(pid), string(sv), "1", zigbee_certification_type, certification_date, string(softwareVersionString), cd_certificate_id, "zb_account")
	result, _ = cli.Execute("bash", "-c", "echo test1234 | ", cmd)
	fmt.Printf(result)

	fmt.Printf("Delete Model with VID: " + string(vid) + "PID: " + string(pid))
	result, _ = cli.DeleteModel(string(vid), string(pid), vendor_account)
	cli.Check_response(result, "\"code\": 525") // code for model certified error

	// CLI side errors

	fmt.Printf("Add Model with VID: " + string(vid) + "PID: " + string(pid) + ": Unknown account")
	cmd, _ = cli.AddModel(string(vid), string(pid), "1", "TestProduct", "TestingProductLabel", "1", "0", vendor_account)
	result, _ = cli.Execute("bash", "-c", "echo test1234 | ", cmd)
	result, _ = cli.AddModel(string(vid), string(pid), "1", "TestProduct", "TestingProductLabel", "1", "0", "Unknown")
	result, _ = cli.Execute("bash", "-c", result, "|| true")
	cli.Check_response_and_report(result, "key not found", "raw")

	cli.Test_divider()

	fmt.Printf("Add model with invalid VID/PID")
	i := "-1"
	cmd, _ = cli.AddModel(i, string(pid), "1", "TestProduct", "TestingProductLabel", "1", "0", vendor_account)
	result, _ = cli.Execute("bash", "-c", "echo test1234 |", cmd)
	result, _ = cli.Execute("bash", "-c", result, "|| true")
	cli.Check_response_and_report(result, "Vid must not be less than 1", "raw")

	cmd, _ = cli.AddModel(string(vid), i, "1", "TestProduct", "TestingProductLabel", "1", "0", vendor_account)
	result, _ = cli.Execute("bash", "-c", "echo test1234 |", cmd)
	result, _ = cli.Execute("bash", "-c", result, "|| true")
	cli.Check_response_and_report(result, "Pid must not be less than 1", "raw")

	i = "0"
	cmd, _ = cli.AddModel(i, string(pid), "1", "TestProduct", "TestingProductLabel", "1", "0", vendor_account)
	result, _ = cli.Execute("bash", "-c", "echo test1234 |", cmd)
	result, _ = cli.Execute("bash", "-c", result, "|| true")
	cli.Check_response_and_report(result, "Vid must not be less than 1", "raw")

	cmd, _ = cli.AddModel(string(vid), i, "1", "TestProduct", "TestingProductLabel", "1", "0", vendor_account)
	result, _ = cli.Execute("bash", "-c", "echo test1234 |", cmd)
	result, _ = cli.Execute("bash", "-c", result, "|| true")
	cli.Check_response_and_report(result, "Pid must not be less than 1", "raw")

	i = "65536"
	cmd, _ = cli.AddModel(i, string(pid), "1", "TestProduct", "TestingProductLabel", "1", "0", vendor_account)
	result, _ = cli.Execute("bash", "-c", "echo test1234 |", cmd)
	result, _ = cli.Execute("bash", "-c", result, "|| true")
	cli.Check_response_and_report(result, "Vid must not be greater than 65535", "raw")

	cmd, _ = cli.AddModel(string(vid), i, "1", "TestProduct", "TestingProductLabel", "1", "0", vendor_account)
	result, _ = cli.Execute("bash", "-c", "echo test1234 |", cmd)
	result, _ = cli.Execute("bash", "-c", result, "|| true")
	cli.Check_response_and_report(result, "Pid must not be greater than 65535", "raw")

	i = "string"
	cmd, _ = cli.AddModel(i, string(pid), "1", "TestProduct", "TestingProductLabel", "1", "0", vendor_account)
	result, _ = cli.Execute("bash", "-c", "echo test1234 |", cmd)
	result, _ = cli.Execute("bash", "-c", result, "|| true")
	cli.Check_response_and_report(result, "invalid syntax", "raw")

	cmd, _ = cli.AddModel(string(vid), i, "1", "TestProduct", "TestingProductLabel", "1", "0", vendor_account)
	result, _ = cli.Execute("bash", "-c", "echo test1234 |", cmd)
	result, _ = cli.Execute("bash", "-c", result, "|| true")
	cli.Check_response_and_report(result, "invalid syntax", "raw")

	cli.Test_divider()

	fmt.Printf("Add model with empty name")
	cmd, _ = cli.AddModel(string(vid), string(pid), "1", "", "TestingProductLabel", "1", "0", vendor_account)
	result, _ = cli.Execute("bash", "-c", "echo test1234 | ", cmd)
	result, _ = cli.Execute("bash", "-c", result, "|| true")
	cli.Check_response_and_report(result, "ProductName is a required field", "raw")

	cli.Test_divider()

	fmt.Printf("Add model with empty --from flag")
	result, _ = cli.AddModel(string(vid), string(pid), "1", "TestProduct", "TestingProductLabel", "1", "0", "")
	result, _ = cli.Execute("bash", "-c", result, "|| true")
	cli.Check_response_and_report(result, "invalid creator address (empty address string is not allowed)", "raw")

	cli.Test_divider()

	fmt.Printf("Add model without --from flag")
	result, _ = cli.AddModel(string(vid), string(pid), "1", "TestProduct", "TestingProductLabel", "", "1", "0")
	result, _ = cli.Execute("bash", "-c", result, "|| true")
	cli.Check_response_and_report(result, "required flag(s) \"from\" not set", "raw")

	cli.Test_divider()

	fmt.Printf("Update model with Non Mutable fields")
	pid = cli.Random()
	sv = cli.Random()
	svs = cli.Random()
	cli.Create_model_and_version(string(vid), string(pid), string(sv), string(svs), vendor_account)
	result, _ = cli.GetModel(string(vid), string(pid))
	fmt.Printf(result)

}
