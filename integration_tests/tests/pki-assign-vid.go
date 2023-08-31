package main

import (
	"fmt"

	cli "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli"
)

func Pki_addign_vid() {

	root_cert_subject_path := "integration_tests/constants/root_cert"
	root_cert_subject := "MDQxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRAwDgYDVQQKDAdyb290LWNh"
	root_cert_subject_key_id := "5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB"
	root_cert_vid := 65521

	trustee_account := "jack"
	second_trustee_account := "alice"

	fmt.Printf("Create a VendorAdmin Account")
	cli.Create_new_account("vendor_admin_account", "VendorAdmin")

	cli.Test_divider()

	// ASSIGN VID TO ROOT CERTIFICATE THAT ALREADY HAS VID
	fmt.Printf("ASSIGN VID TO ROOT CERTIFICATE THAT ALREADY HAS VID")

	fmt.Printf("Propose and approve root certificate")
	cmd, _ := cli.ProposeAddx509(root_cert_subject_path, string(root_cert_vid), trustee_account)
	result, _ := cli.Execute("bash", "-c", "echo ", cli.Passphrase, "| ", cmd)
	cli.Check_response(result, "\"code\": 0")
	cmd, _ = cli.ApproveAddx509(root_cert_subject, root_cert_subject_key_id, second_trustee_account)
	result, _ = cli.Execute("bash", "-c", "echo ", cli.Passphrase, " | ", cmd)
	cli.Check_response(result, "\"code\": 0")

	fmt.Printf("Assing VID")
	result, _ = cli.AssignVid(root_cert_subject, root_cert_subject_key_id, string(root_cert_vid), "vendor_admin_account")
	cli.Check_response(result, "vid is not empty")

	cli.Test_divider()
}
