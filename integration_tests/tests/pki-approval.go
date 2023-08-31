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

func main() {

	root_cert_subject := "MDQxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRAwDgYDVQQKDAdyb290LWNh"
	root_cert_subject_key_id := "5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB"
	//root_cert_serial_number := "442314047376310867378175982234956458728610743315"
	vid := 1

	//intermediate_cert_subject := "MDwxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRgwFgYDVQQKDA9pbnRlcm1lZGlhdGUtY2E="
	//intermediate_cert_subject_key_id := "4E:3B:73:F4:70:4D:C2:98:0D:DB:C8:5A:5F:02:3B:BF:86:25:56:2B"
	//intermediate_cert_serial_number := "169917617234879872371588777545667947720450185023"

	//leaf_cert_subject := "MDExCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMQ0wCwYDVQQKDARsZWFm"
	//leaf_cert_subject_key_id := "30:F4:65:75:14:20:B2:AF:3D:14:71:17:AC:49:90:93:3E:24:A0:1F"
	//leaf_cert_serial_number := "143290473708569835418599774898811724528308722063"

	// Preparation of Actors

	first_trustee_account := "jack"
	second_trustee_account := "alice"
	third_trustee_account := "bob"

	cmd, _ := cli.ShowKeys("jack", "-a")
	first_trustee_address, _ := cli.Execute("bash", "-c", "echo", cli.Passphrase, "| ", cmd)
	cmd, _ = cli.ShowKeys("alice", "-a")
	second_trustee_address, _ := cli.Execute("bash", "-c", "echo", cli.Passphrase, "| ", cmd)
	cmd, _ = cli.ShowKeys("bob", "-a")
	third_trustee_address, _ := cli.Execute("bash", "-c", "echo", cli.Passphrase, "| ", cmd)

	fmt.Printf("Create regular account")
	var user_account string
	user_account = cli.Create_new_account("user_account", "CertificationCenter")

	cli.Test_divider()

	fmt.Printf("Add 3 new Trustee accounts, this will result in a total of 6 Trustees and 4 approvals needed for 2/3 quorum")

	cli.Test_divider()

	fmt.Printf("Add Fourth Trustee Account")
	var fourth_trustee_account string
	fourth_trustee_account = cli.Random_string(fourth_trustee_account)
	fmt.Printf(fourth_trustee_account + "generates keys")
	result, _ := cli.AddKeys(fourth_trustee_account)
	cmd, _ = cli.ShowKeys(fourth_trustee_account, "-a")
	var fourth_trustee_address string
	fourth_trustee_address, _ = cli.Execute("bash", "-c", "echo ", cli.Passphrase, "|", cmd)
	cmd, _ = cli.ShowKeys(fourth_trustee_account, "-p")
	//var fourth_trustee_pubkey string
	//fourth_trustee_pubkey, _ = cli.Execute("bash", "-c", "echo ", cli.Passphrase, "|", cmd)

	fmt.Printf(first_trustee_account + "proposes account for " + fourth_trustee_account)
	cmd, _ = cli.ProposeAddAccount("Jack is proposing this account", fourth_trustee_address, "Trustee", "", "", first_trustee_account)
	result, _ = cli.Execute("bash", "-c", "echo ", cli.Passphrase, cmd)
	cli.Check_response(result, "\"code\": 0")

	fmt.Printf(second_trustee_account + "approves account for " + fourth_trustee_account)
	cmd, _ = cli.Approve_add_account(fourth_trustee_address, second_trustee_address)
	result, _ = cli.Execute("bash", "-c", "echo ", cmd)
	cli.Check_response(result, "\"code\": 0")

	fmt.Printf("Verify that the account is now present")
	cmd, _ = cli.Account(fourth_trustee_address)
	result, _ = cli.Execute("bash", "-c", "echo ", cli.Passphrase, "|", cmd)
	cli.Check_response(result, "\"address\": \"", fourth_trustee_address, "\"")

	cli.Test_divider()

	fmt.Printf("Add Fifth Trustee Account")
	var fifth_trustee_account string
	fifth_trustee_account = cli.Random_string(fifth_trustee_account)
	fmt.Printf(fifth_trustee_account + "generates keys")
	result, _ = cli.AddKeys(fifth_trustee_account)

	var fifth_trustee_address string
	cmd, _ = cli.ShowKeys(fifth_trustee_account, "-a")
	fifth_trustee_address, _ = cli.Execute("bash", "-c", "echo ", cli.Passphrase, "|", cmd)
	var fifth_trustee_pubkey string
	cmd, _ = cli.ShowKeys(fifth_trustee_account, "-p")
	fifth_trustee_pubkey, _ = cli.Execute("bash", "-c", "echo ", cli.Passphrase, "|", cmd)

	fmt.Printf(first_trustee_account + "proposes account for " + fifth_trustee_account)
	cmd, _ = cli.ProposeAddAccount("Jack is proposing this account", fifth_trustee_address, fifth_trustee_pubkey, "Trustee", "", first_trustee_account)
	result, _ = cli.Execute("bash", "-c", "echo ", cli.Passphrase, "|", cmd)
	cli.Check_response(result, "\"code\": 0")

	fmt.Printf(second_trustee_account + "approves account for " + fifth_trustee_account)
	cmd, _ = cli.Approve_add_account(fifth_trustee_account, second_trustee_account)
	result, _ = cli.Execute("bash", "-c", "echo ", cli.Passphrase, "|", cmd)
	cli.Check_response(result, "\"code\": 0")

	fmt.Printf(third_trustee_account + "approves account for " + fifth_trustee_account)
	cmd, _ = cli.Approve_add_account(fifth_trustee_account, third_trustee_account)
	result, _ = cli.Execute("bash", "-c", "echo ", cli.Passphrase, "|", cmd)
	cli.Check_response(result, "\"code\": 0")

	fmt.Printf("Verify that fifth account is now present")
	cmd, _ = cli.Account(fifth_trustee_account)
	result, _ = cli.Execute("bash", "-c", "echo ", cli.Passphrase, "|", cmd)
	cli.Check_response(result, "\"address\": \"", fifth_trustee_address, "\"")

	cli.Test_divider()

	fmt.Printf("Add sixth Trustee Account")
	var sixth_trustee_account string
	sixth_trustee_account = cli.Random_string(sixth_trustee_account)
	fmt.Printf(sixth_trustee_account + "generates keys")
	result, _ = cli.AddKeys(sixth_trustee_account)
	cmd, _ = cli.ShowKeys(sixth_trustee_account, "-a")
	var sixth_trustee_address string
	sixth_trustee_address, _ = cli.Execute("bash", "-c", "echo ", cli.Passphrase, "|", cmd)
	cmd, _ = cli.ShowKeys(sixth_trustee_account, "-p")
	var sixth_trustee_pubkey string
	sixth_trustee_pubkey, _ = cli.Execute("bash", "-c", "echo ", cli.Passphrase, "|", cmd)

	fmt.Printf(first_trustee_account + "proposes account for " + sixth_trustee_account)
	cmd, _ = cli.ProposeAddAccount("Jack is proposing this account", sixth_trustee_address, sixth_trustee_pubkey, "Trustee", "", first_trustee_account)
	result, _ = cli.Execute("bash", "-c", "echo", cli.Passphrase, "|", cmd)
	cli.Check_response(result, "\"code\": 0")

	fmt.Printf(second_trustee_account + "approves account for " + sixth_trustee_account)
	cmd, _ = cli.Approve_add_account(sixth_trustee_account, second_trustee_account)
	result, _ = cli.Execute("bash", "-c", "echo ", cli.Passphrase, "| ", cmd)
	cli.Check_response(result, "\"code\": 0")

	fmt.Printf(third_trustee_account + "approves account for " + sixth_trustee_account)
	cmd, _ = cli.Approve_add_account(sixth_trustee_address, third_trustee_account)
	result, _ = cli.Execute("bash", "-c", "echo ", cli.Passphrase, "| ", cmd)
	cli.Check_response(result, "\"code\": 0")

	fmt.Printf(fourth_trustee_account + "approves account for " + sixth_trustee_account)
	cmd, _ = cli.Approve_add_account(sixth_trustee_address, fourth_trustee_account)
	result, _ = cli.Execute("bash", "-c", "echo ", cli.Passphrase, "|", cmd)
	cli.Check_response(result, "\"code\": 0")

	fmt.Printf("Verify that sixth account is now present")
	cmd, _ = cli.Account(sixth_trustee_address)
	result, _ = cli.Execute("bash", "-c", "echo ", cli.Passphrase, "|", cmd)
	cli.Check_response(result, "\"address\": \"", sixth_trustee_address, "\"")

	cli.Test_divider()

	fmt.Printf("PROPOSE ROOT CERT")
	fmt.Printf(user_account + "(Not Trustee) propose Root certificate")
	root_path := "integration_tests/constants/root_cert"
	cmd, _ = cli.ProposeAddx509(root_path, string(vid), user_account)
	result, _ = cli.Execute("bash", "-c", "echo", cli.Passphrase, " |", cmd)
	cli.Response_does_not_contain(result, "\"code\":", "0")

	fmt.Printf(fourth_trustee_account + "(Trustee) propose Root certificate")
	cmd, _ = cli.ProposeAddx509(root_path, string(vid), fourth_trustee_account)
	result, _ = cli.Execute("bash", "-c", "echo ", cli.Passphrase, " | ", cmd)
	cli.Check_response(result, "\"code\": 0")

	cli.Test_divider()

	fmt.Printf("Approve Root certificate now 4 Approvals are needed as we have 6 trustees")

	fmt.Printf(first_trustee_account + "(Trustee) approve Root certificate")
	cmd, _ = cli.ApproveAddx509(root_cert_subject, root_cert_subject_key_id, first_trustee_account)
	result, _ = cli.Execute("bash", "-c", "echo", cli.Passphrase, " |", cmd)
	cli.Check_response(result, "\"code\": 0")

	cli.Test_divider()

	fmt.Printf("Verify Root certificate is not present in approved state")
	cmd, _ = cli.X509Cert(root_cert_subject, root_cert_subject_key_id)
	result, _ = cli.Execute("bash", "-c", "echo ", cli.Passphrase, " | ", cmd)
	cli.Check_response(result, "Not Found")

	cli.Test_divider()

	fmt.Printf(second_trustee_account + "(Trustee) approve Root certificate")
	cmd, _ = cli.ApproveAddx509(root_cert_subject, root_cert_subject_key_id, second_trustee_account)
	result, _ = cli.Execute("bash", "-c", "echo ", cli.Passphrase, " | ", cmd)
	cli.Check_response(result, "\"code\": 0")

	cli.Test_divider()

	fmt.Printf("Verify Root certificate is not present in approved state")
	cmd, _ = cli.X509Cert(root_cert_subject, root_cert_subject_key_id)
	result, _ = cli.Execute("bash", "-c", "echo ", cli.Passphrase, " | ", cmd)
	cli.Check_response(result, "Not Found")

	cli.Test_divider()

	fmt.Printf(third_trustee_account + "(Trustee) approve Root certificate")
	cmd, _ = cli.ApproveAddx509(root_cert_subject, root_cert_subject_key_id, third_trustee_account)
	result, _ = cli.Execute("bash", "-c", "echo ", cli.Passphrase, " | ", cmd)
	cli.Check_response(result, "\"code\": 0")

	cli.Test_divider()

	fmt.Printf("Verify Root certificate is in approved state")
	cmd, _ = cli.X509Cert(root_cert_subject, root_cert_subject_key_id)
	result, _ = cli.Execute("bash", "-c", "echo ", cli.Passphrase, " |", cmd)
	cli.Execute("bash", "-c", "echo", result, "| jq")
	cli.Check_response(result, "\"subject\": \"", root_cert_subject, "\"")
	cli.Check_response(result, "\"address\": \"", first_trustee_address, "\"")
	cli.Check_response(result, "\"address\": \"", second_trustee_address, "\"")
	cli.Check_response(result, "\"address\": \"", third_trustee_address, "\"")

	cli.Test_divider()

	fmt.Printf(sixth_trustee_account + "proposes revoke Root certificate")
	cmd, _ = cli.ProposeRevokeX509(root_cert_subject, root_cert_subject_key_id, sixth_trustee_account)
	result, _ = cli.Execute("bash", "_c", "echo ", cli.Passphrase, " |", cmd)
	cli.Check_response(result, "\"code\": 0")

	fmt.Printf("Request root certificate proposed to revoke and verify that it contains approval from " + sixth_trustee_account)
	cmd, _ = cli.ProposedX509CertToRevoke(root_cert_subject, root_cert_subject_key_id)
	result = cmd
	cli.Execute("bash", "-c", "echo", result, "| jq")
	cli.Check_response(result, "\"", root_cert_subject, "\"")
	cli.Check_response(result, "\"", root_cert_subject_key_id, "\"")
	cli.Check_response(result, "\"address\": \"", sixth_trustee_address, "\"")

	cli.Test_divider()

	fmt.Printf(fifth_trustee_account + "revokes Root certificate")
	cmd, _ = cli.ApproveRevokeX509(root_cert_subject, root_cert_subject_key_id, fifth_trustee_account)
	result, _ = cli.Execute("bash", "-c", "echo "+cli.Passphrase, " | ", cmd)
	cli.Check_response(result, "\"code\": 0")

	fmt.Printf("Request root certificate proposed to revoke and verify that it contains approval from " + fifth_trustee_account)
	cmd, _ = cli.ProposedX509CertToRevoke(root_cert_subject, root_cert_subject_key_id)
	result = cmd
	cli.Execute("bash", "-c", "echo", result, "| jq")
	cli.Check_response(result, "\"", root_cert_subject, "\"")
	cli.Check_response(result, "\"", root_cert_subject_key_id, "\"")
	cli.Check_response(result, "\"address\": \"", sixth_trustee_address, "\"")
	cli.Check_response(result, "\"address\": \"", fifth_trustee_address, "\"")

	fmt.Printf(fourth_trustee_account + "revokes Root certificate")
	cmd, _ = cli.ApproveRevokeX509(root_cert_subject, root_cert_subject_key_id, fourth_trustee_account)
	result, _ = cli.Execute("bash", "-c", "echo ", cli.Passphrase, " | ", cmd)
	cli.Check_response(result, "\"code\": 0")

	fmt.Printf("Request root certificate proposed to revoke and verify that it contains approval from " + fourth_trustee_account)
	cmd, _ = cli.ProposedX509CertToRevoke(root_cert_subject, root_cert_subject_key_id)
	result = cmd
	cli.Execute("bash", "-c", "echo", result, "| jq")
	cli.Check_response(result, "\"", root_cert_subject, "\"")
	cli.Check_response(result, "\"", root_cert_subject_key_id, "\"")
	cli.Check_response(result, "\"address\": \"", sixth_trustee_address, "\"")
	cli.Check_response(result, "\"address\": \"", fifth_trustee_address, "\"")
	cli.Check_response(result, "\"address\": \"", fourth_trustee_address, "\"")

	fmt.Printf(third_trustee_account + "revokes Root certificate")
	cmd, _ = cli.ApproveRevokeX509(root_cert_subject, root_cert_subject_key_id, third_trustee_account)
	result, _ = cli.Execute("bash", "-c", "echo ", cli.Passphrase, " | ", cmd)
	cli.Check_response(result, "\"code\": 0")

	fmt.Printf("Verify Root certificate is now revoked")
	cmd, _ = cli.RevokedX509(root_cert_subject, root_cert_subject_key_id)
	result, _ = cli.Execute("bash", "-c", "echo ", cli.Passphrase, " |", cmd)
	cli.Execute("bash", "-c", "echo", result, "| jq")
	cli.Check_response(result, "\"", root_cert_subject, "\"")
	cli.Check_response(result, "\"", root_cert_subject_key_id, "\"")
	cli.Check_response(result, "\"address\": \"", sixth_trustee_address, "\"")
	cli.Check_response(result, "\"address\": \"", fifth_trustee_address, "\"")
	cli.Check_response(result, "\"address\": \"", fourth_trustee_address, "\"")
	cli.Check_response(result, "\"address\": \"", third_trustee_address, "\"")

}
