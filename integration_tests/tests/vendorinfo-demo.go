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

func Vendorinfo_demo() {

	// Preparation of Actors
	vid := cli.Random()
	vid2 := cli.Random()
	vendor_account := "vendor_account_" + string(vid)
	second_vendor_account := "vendor_account_" + string(vid2)
	var vendor_admin_account string
	fmt.Printf("Create First Vendor Account - " + vendor_account)
	cli.Create_new_vendor_account(vendor_account, string(vid))
	fmt.Printf("Create Second Vendor Account - " + second_vendor_account)
	cli.Create_new_vendor_account(second_vendor_account, string(vid2))
	fmt.Printf("Create A VendorAdmin Account - " + vendor_admin_account)
	cli.Create_new_account(vendor_admin_account, "VendorAdmin")

	cli.Test_divider()

	// Query non existent
	fmt.Printf("Query non existent vendorinfo")
	result, _ := cli.Vendor(string(vid))
	cli.Check_response(result, "Not Found")
	fmt.Printf(result)

	cli.Test_divider()

	fmt.Printf("Request all vendor info must be empty")
	result, _ = cli.AllVendors()
	cli.Check_response(result, "\"[\"]")
	fmt.Printf(result)

	cli.Test_divider()

	// Create a vendor info record
	fmt.Printf("Create VendorInfo Record for VID: " + string(vid))
	companyLegalName := "XYZ IOT Devices Inc"
	vendorName := "XYZ Devices"
	cmd, _ := cli.AddVendor(string(vid), companyLegalName, vendorName, vendor_account)
	result, _ = cli.Execute("bash", "-c", "echo test1234 | ", cmd)
	cli.Check_response(result, "\"code\": 0")
	fmt.Printf(result)

	cli.Test_divider()

	// Query vendor info record
	fmt.Printf("Verify if VendorInfo Record for VID: " + string(vid) + "is present or not")
	result, _ = cli.Vendor(string(vid))
	cli.Check_response(result, "\"vendorID\": ", string(vid))
	cli.Check_response(result, "\"companyLegalName\": \"", companyLegalName, "\"")
	cli.Check_response(result, "\"vendorName\": \"", vendorName, "\"")
	fmt.Printf(result)

	cli.Test_divider()

	fmt.Printf("Request all vendor info")
	result, _ = cli.AllVendors()
	cli.Check_response(result, "\"vendorID\": ", string(vid))
	cli.Check_response(result, "\"companyLegalName\": \"", companyLegalName, "\"")
	cli.Check_response(result, "\"vendorName\": \"", vendorName, "\"")
	fmt.Printf(result)

	cli.Test_divider()

	// Update vendor info with empty optional fields
	fmt.Printf("Update vendor info record for VID: " + string(vid) + "(with required fields only)")
	result, _ = cli.UpdateVendor(string(vid), "", "", "", string(vendor_account))
	cli.Check_response(result, "\"code\": 0")
	fmt.Printf(result)

	cli.Test_divider()

	// Query updated vendor info record
	fmt.Printf("Verify that omitted fields in update object are not set to empty string")
	result, _ = cli.Vendor(string(vid))
	cli.Check_response(result, "\"vendorID\": ", string(vid))
	cli.Check_response(result, "\"companyLegalName\": \"", companyLegalName, "\"")
	cli.Check_response(result, "\"vendorName\": \"", vendorName, "\"")
	fmt.Printf(result)

	cli.Test_divider()

	// Update vendor info record
	fmt.Printf("Update vendor info record for VID: ", vid)
	companyLegalName = "ABC Subsidiary Corporation"
	vendorLandingPageURL := "https://www.w3.org/"
	cmd, _ = cli.UpdateVendor(string(vid), companyLegalName, vendorLandingPageURL, vendorName, vendor_account)
	result, _ = cli.Execute("bash", "-c", "echo test1234 | ", cmd)
	cli.Check_response(result, "\"code\": 0")
	fmt.Printf(result)

	cli.Test_divider()

	// Query updated vendor info record
	fmt.Printf("Verify if VendorInfo Record for VID: " + string(vid) + "is updated or not")
	result, _ = cli.Vendor(string(vid))
	cli.Check_response(result, "\"vendorID\": ", string(vid))
	cli.Check_response(result, "\"companyLegalName\": \"", companyLegalName, "\"")
	cli.Check_response(result, "\"vendorName\": \"", vendorName, "\"")
	cli.Check_response(result, "\"vendorLandingPageURL\": \"", vendorLandingPageURL, "\"")
	fmt.Printf(result)

	cli.Test_divider()

	// Create a vendor info record from a vendor account belonging to another vendor_account
	vid1 := cli.Random()
	cmd, _ = cli.AddVendor(string(vid1), companyLegalName, vendorName, vendor_account)
	result, _ = cli.Execute("bash", "-c", "echo test1234 | ")
	result, _ = cli.Execute("bash", "-c", result, "|| true")
	fmt.Printf(result)
	cli.Check_response_and_report(result, "transaction should be signed by a vendor account associated with the vendorID ", string(vid1))

	cli.Test_divider()

	// Update a vendor info record from a vendor account belonging to another vendor_account
	cmd, _ = cli.UpdateVendor(string(vid), companyLegalName, vendorName, "", string(second_vendor_account))
	result, _ = cli.Execute("bash", "-c", "echo test1234 | ", cmd)
	result, _ = cli.Execute("bash", "-c", result, "|| true")
	fmt.Printf(result)
	cli.Check_response_and_report(result, "transaction should be signed by a vendor account associated with the vendorID ", string(vid))

	cli.Test_divider()

	// Create a vendor info reacord from a vendor admin account
	fmt.Printf("Create a vendor info reacord from a vendor admin account")
	vid = cli.Random()
	cmd, _ = cli.AddVendor(string(vid), companyLegalName, vendorName, vendor_admin_account)
	result, _ = cli.Execute("bash", "-c", "echo test1234 | ", cmd)
	result, _ = cli.Execute("bash", "-c", result, "|| true")
	fmt.Printf(result)
	cli.Check_response(result, "\"code\": 0")
	fmt.Printf(result)

	cli.Test_divider()

	// Update the vendor info record by a vendor admin account
	fmt.Printf("Update the vendor info record by a vendor admin account")
	companyLegalName1 := "New Corp"
	vendorName1 := "New Vendor Name"
	cmd, _ = cli.UpdateVendor(string(vid), companyLegalName1, vendorName1, "", string(vendor_admin_account))
	result, _ = cli.Execute("bash", "-c", "echo test1234 | ", cmd)
	cli.Check_response(result, "\"code\": 0")
	fmt.Printf(result)

	cli.Test_divider()
}
