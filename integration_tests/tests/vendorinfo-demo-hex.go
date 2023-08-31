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

func Vendorinfo_demo_hex() {

	//Preparation of Actors
	vid_in_hex_format := 0xA13
	vid2_in_hex_format := 0xA14
	vid := 2579
	//vid2 := 2580
	vendor_account := "vendor_account_" + string(vid_in_hex_format)
	second_vendor_account := "vendor_account_" + string(vid2_in_hex_format)
	fmt.Printf("Create First Vendor Account - " + vendor_account)
	cli.Create_new_vendor_account(vendor_account, string(vid_in_hex_format))
	fmt.Printf("Create Second Vendor Account - " + second_vendor_account)
	cli.Create_new_vendor_account(second_vendor_account, string(vid2_in_hex_format))

	cli.Test_divider()

	// Create a vendor info record
	fmt.Printf("Create VendorInfo Record for VID: " + string(vid_in_hex_format))
	companyLegalName := "XYZ IOT Devices Inc"
	vendorName := "XYZ Devices"
	cmd, _ := cli.AddVendor(string(vid_in_hex_format), companyLegalName, vendorName, vendor_account)
	result, _ := cli.Execute("bash", "-c", "echo test1234 | ", cmd)
	cli.Check_response(result, "\"code\": 0")
	fmt.Printf(result)

	cli.Test_divider()

	// Query vendor info record
	fmt.Printf("Verify if VendorInfo Record for VID: " + string(vid_in_hex_format) + "is present or not")
	result, _ = cli.Vendor(string(vid_in_hex_format))
	cli.Check_response(result, "\"vendorID\": ", string(vid))
	cli.Check_response(result, "\"companyLegalName\": \"", companyLegalName, "\"")
	cli.Check_response(result, "\"vendorName\": \"", vendorName, "\"")
	fmt.Printf(result)

	cli.Test_divider()

	// Update vendor info record
	fmt.Printf("Update vendor info record for VID: " + string(vid_in_hex_format))
	companyLegalName = "ABC Subsidiary Corporation"
	vendorLandingPageURL := "https://www.w3.org/"
	cmd, _ = cli.UpdateVendor(string(vid_in_hex_format), companyLegalName, vendorLandingPageURL, vendorName, vendor_account)
	result, _ = cli.Execute("bash", "-c", "echo test1234 | ", cmd)
	cli.Check_response(result, "\"code\": 0")
	fmt.Printf(result)

	cli.Test_divider()

	// Query updated vendor info record
	fmt.Printf("Verify if VendorInfo Record for VID: " + string(vid_in_hex_format) + "is updated or not")
	result, _ = cli.Vendor(string(vid_in_hex_format))
	cli.Check_response(result, "\"vendorID\": ", string(vid))
	cli.Check_response(result, "\"companyLegalName\": \"", companyLegalName, "\"")
	cli.Check_response(result, "\"vendorName\": \"", vendorName, "\"")
	cli.Check_response(result, "\"vendorLandingPageURL\": \"", vendorLandingPageURL, "\"")
	fmt.Printf(result)

	cli.Test_divider()

	// Create a vendor info record from a vendor account belonging to another vendor_account
	vid1_in_hex_format := 0xA15
	vid1 := 2581
	cmd, _ = cli.AddVendor(string(vid1_in_hex_format), companyLegalName, vendorName, vendor_account)
	result, _ = cli.Execute("bash", "-c", "echo test1234 | ", cmd)
	result, _ = cli.Execute("bash", "-c", result, "|| true")
	fmt.Printf(result)
	cli.Check_response_and_report(result, "transaction should be signed by a vendor account associated with the vendorID ", string(vid1))

	cli.Test_divider()

	// Update a vendor info record from a vendor account belonging to another vendor_account
	cmd, _ = cli.UpdateVendor(string(vid_in_hex_format), companyLegalName, vendorName, "", string(second_vendor_account))
	result, _ = cli.Execute("bash", "-c", "echo test1234 | ", cmd)
	result, _ = cli.Execute("bash", "-c", result, "|| true")
	fmt.Printf(result)
	cli.Check_response_and_report(result, "transaction should be signed by a vendor account associated with the vendorID ", string(vid))

	cli.Test_divider()
}
