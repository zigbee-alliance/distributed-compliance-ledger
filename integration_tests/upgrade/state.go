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

package upgrade

// UpgradeTestState carries the on-chain identifiers and per-version constants
// that earlier scripts (01-07) put into the chain. Phase 1 subtests depend on
// these; once phase 2/3/4 migrates the producing scripts to Go, the same
// struct will be populated by Go setup code.
//
// Values mirror the literals in the existing bash scripts. Names match the
// bash variable names to keep cross-referencing easy during migration.
type UpgradeTestState struct {
	// Trustees set up by 01-test-upgrade-initialize-0.12.sh.
	Trustee1 string // "jack"   — hardcoded in 01.sh
	Trustee2 string // "alice"  — hardcoded in 01.sh
	Trustee3 string // "bob"    — hardcoded in 01.sh
	Trustee4 string // randomized in 01.sh; phase 2 will lock to a stable name.

	// Vendor + model identifiers seeded by 01.
	VendorAccount   string
	VID             int
	PID2            int
	ProductName     string
	ProductLabel    string
	PartNumber      string
	SoftwareVersion int

	// Vendor account created by 03-test-upgrade-0.12-to-1.2.sh.
	VendorAccountFor1_2 string

	// Constants set by 05 (lines 22-29) — referenced by 09's "ISSUE #593"
	// block BEFORE 09 redefines them at line 174+.
	VIDFor1_6_0_FromScript5              int
	PID3For1_6_0_FromScript5             int
	SoftwareVersion1For1_6_0_FromScript5 int
	SoftwareVersion2For1_6_0_FromScript5 int

	// User accounts created in 01 — referenced by later scripts' pre-upgrade
	// verifications. Names are randomized; addresses + pubkeys are propagated
	// for later auth-flow steps that need them.
	User1Address, User1Pubkey string // revoked by 01 (revoked-accounts list)
	User2Address, User2Pubkey string // active CertCenter, then propose-revoke in 01
	User3Address, User3Pubkey string // remains proposed in 01

	// User accounts created in 03.
	User4Address, User4Pubkey string
	User5Address, User5Pubkey string
	User6Address, User6Pubkey string

	// User accounts created in 05.
	User7Address, User7Pubkey string
	User8Address, User8Pubkey string
	User9Address, User9Pubkey string

	// User accounts created in 06.
	User10Address, User10Pubkey string
	User11Address, User11Pubkey string
	User12Address, User12Pubkey string

	// User accounts created in 07.
	User13Address, User13Pubkey string
	User14Address, User14Pubkey string
	User15Address, User15Pubkey string

	// Master upgrade plan name = git short HEAD hash, computed when script 10
	// builds the dcld-build-master image. Used by script 11 to verify the new
	// observer eventually reports the same version after catch-up.
	MasterPlanName string

	// Validator-demo node bookkeeping. Set by AddValidatorNode (script 01)
	// and re-used by all per-script disable/enable flows.
	ValidatorAccountName string // random name on the validator-demo's keyring
	ValidatorAddress     string // cosmosvaloper... owner from `query validator node`

	// Constants set by 07-test-upgrade-1.4.4-to-1.5.1.sh post-upgrade block.
	VIDFor1_5_1                                 int
	PID1For1_5_1                                int
	PID2For1_5_1                                int
	ProductLabelFor1_5_1                        string
	CommissioningModeSecondaryStepsHintFor1_5_1 int
	PartNumberFor1_5_1                          string
	SoftwareVersionFor1_5_1                     int
	MinApplicableSoftwareVersionFor1_5_1        int
	MaxApplicableSoftwareVersionFor1_5_1        int
}

// DefaultBashState returns the contract that Phase 1 subtests assume. Values
// match the literals in the corresponding bash scripts. When phases 2-3
// migrate those scripts to Go, the producing code must keep these literal
// values in lockstep (or this struct must be updated together with the
// migration).
//
// The empty Trustee4 captures that the bash version randomizes that name; the
// migrated phase-2 code will fill it in.
func DefaultBashState() *UpgradeTestState {
	return &UpgradeTestState{
		// 01-test-upgrade-initialize-0.12.sh
		Trustee1:        "jack",
		Trustee2:        "alice",
		Trustee3:        "bob",
		Trustee4:        "", // randomized — phase 2 will pin this
		VendorAccount:   "vendor_account",
		VID:             1,
		PID2:            2,
		ProductName:     "ProductName",
		ProductLabel:    "ProductLabel",
		PartNumber:      "RCU2205A",
		SoftwareVersion: 1,

		// 03-test-upgrade-0.12-to-1.2.sh
		VendorAccountFor1_2: "vendor_account_4701",

		// 05-test-upgrade-1.2-to-1.4.3.sh lines 22-29 (used early in 09)
		VIDFor1_6_0_FromScript5:              4701, // = $vid_for_1_2
		PID3For1_6_0_FromScript5:             160,
		SoftwareVersion1For1_6_0_FromScript5: 100001,
		SoftwareVersion2For1_6_0_FromScript5: 200002,

		// 07-test-upgrade-1.4.4-to-1.5.1.sh lines 489-507
		VIDFor1_5_1:          65529,
		PID1For1_5_1:         79,
		PID2For1_5_1:         89,
		ProductLabelFor1_5_1: "ProductLabel_1_5_1",
		CommissioningModeSecondaryStepsHintFor1_5_1: 8,
		PartNumberFor1_5_1:                          "RCU2245M",
		SoftwareVersionFor1_5_1:                     4,
		MinApplicableSoftwareVersionFor1_5_1:        8,
		MaxApplicableSoftwareVersionFor1_5_1:        8000,
	}
}

// Constants used by the rollback portion of script 02
// (02-test-upgrade-0.12-rollback.sh).
const (
	VIDForRollback                          = 4705
	PID1ForRollback                         = 11
	PID2ForRollback                         = 22
	PID3ForRollback                         = 33
	DeviceTypeIDForRollback                 = 1234
	ProductNameForRollback                  = "ProductName_0.12_r"
	ProductLabelForRollback                 = "ProductLabe_0.12_r"
	PartNumberForRollback                   = "RCU2205B"
	SoftwareVersionForRollback              = 2
	SoftwareVersionStringForRollback        = "2.0"
	CDVersionNumberForRollback              = 313
	MinApplicableSoftwareVersionForRollback = 2
	MaxApplicableSoftwareVersionForRollback = 2000

	CertificationTypeForRollback = "matter"
	CertificationDateForRollback = "2021-02-01T00:00:00Z"
	ProvisionalDateForRollback   = "2010-11-12T00:00:00Z"
	CDCertificateIDForRollback   = "12345678910abcdefgh"

	VendorNameForRollback           = "VendorName_r"
	CompanyLegalNameForRollback     = "LegalCompanyName_r"
	CompanyPreferredNameForRollback = "CompanyPreferredName_r"
	VendorLandingPageURLForRollback = "https://www.newexample_rollback.com"

	VendorAccountForRollback              = "vendor_account_r"
	CertificationCenterAccountForRollback = "certification_center_account_r"

	WrongPlanName        = "wrong_plan_name"
	WrongPlanChecksumV12 = "sha256:3f2b2a98b7572c6598383f7798c6bc16b4e432ae5cfd9dc8e84105c3d53b5026"
)

// Constants used by the rollback portion of script 04
// (04-test-upgrade-1.2-rollback.sh).
const (
	VIDFor1_2_R2                          = 4703
	PID1For1_2_R2                         = 16
	PID2For1_2_R2                         = 27
	PID3For1_2_R2                         = 38
	DeviceTypeIDFor1_2_R2                 = 1239
	ProductNameFor1_2_R2                  = "ProductName1.2_r2"
	ProductLabelFor1_2_R2                 = "ProductLabe1.2_r2"
	PartNumberFor1_2_R2                   = "RCU2205F"
	SoftwareVersionFor1_2_R2              = 2
	SoftwareVersionStringFor1_2_R2        = "2.0"
	CDVersionNumberFor1_2_R2              = 313
	MinApplicableSoftwareVersionFor1_2_R2 = 2
	MaxApplicableSoftwareVersionFor1_2_R2 = 2000

	CertificationTypeFor1_2_R2 = "matter"
	CertificationDateFor1_2_R2 = "2021-01-03T00:00:00Z"
	ProvisionalDateFor1_2_R2   = "2010-12-11T00:00:00Z"
	CDCertificateIDFor1_2_R2   = "12345678910abcdefgh"

	VendorNameFor1_2_R2           = "VendorName4705"
	CompanyLegalNameFor1_2_R2     = "LegalCompanyName4705"
	CompanyPreferredNameFor1_2_R2 = "CompanyPreferredName4705"
	VendorLandingPageURLFor1_2_R2 = "https://www.newexample_R2.com"

	VendorAccountFor1_2_R2 = "vendor_account_4705"

	WrongPlanName2        = "wrong_plan_name_2"
	WrongPlanChecksumV143 = "sha256:a007f58d61632af107a09c89b7392eedd05d8127d0df67ace50f318948c62001"
)

// Constants used by the 1.2 portion of script 03.
const (
	VIDFor1_2                          = 4701
	PID1For1_2                         = 11
	PID2For1_2                         = 22
	PID3For1_2                         = 33
	DeviceTypeIDFor1_2                 = 1234
	ProductNameFor1_2                  = "ProductName1.2"
	ProductLabelFor1_2                 = "ProductLabe1.2"
	PartNumberFor1_2                   = "RCU2205B"
	SoftwareVersionFor1_2              = 2
	SoftwareVersionStringFor1_2        = "2.0"
	CDVersionNumberFor1_2              = 313
	MinApplicableSoftwareVersionFor1_2 = 2
	MaxApplicableSoftwareVersionFor1_2 = 2000

	CertificationTypeFor1_2 = "matter"
	CertificationDateFor1_2 = "2021-01-01T00:00:00Z"
	ProvisionalDateFor1_2   = "2010-12-12T00:00:00Z"
	CDCertificateIDFor1_2   = "12345678910abcdefgh"

	VendorNameFor1_2           = "VendorName4701"
	CompanyLegalNameFor1_2     = "LegalCompanyName4701"
	CompanyPreferredNameFor1_2 = "CompanyPreferredName4701"
	VendorLandingPageURLFor1_2 = "https://www.newexample.com"

	VendorAdminAccount               = "vendor_admin_account"
	CertificationCenterAccountFor1_2 = "certification_center_account"

	UpgradeChecksumV1_2 = "sha256:3f2b2a98b7572c6598383f7798c6bc16b4e432ae5cfd9dc8e84105c3d53b5026"
	PlanNameV1_2        = "v1.2"
	BinaryVersionV1_2   = "1.2.2"

	TestDataURL        = "https://url.data.dclmodel"
	IssuerSubjectKeyID = "5A880E6C3653D07FB08971A3F473790930E62BDB"
)

// PKI certs introduced by script 03 (1.2 era).
const (
	RootCertPathFor1_2         = "integration_tests/constants/google_root_cert_gsr4"
	RootCertSubjectFor1_2      = "MFAxJDAiBgNVBAsTG0dsb2JhbFNpZ24gRUNDIFJvb3QgQ0EgLSBSNDETMBEGA1UEChMKR2xvYmFsU2lnbjETMBEGA1UEAxMKR2xvYmFsU2lnbg=="
	RootCertSubjectKeyIDFor1_2 = "54:B0:7B:AD:45:B8:E2:40:7F:FB:0A:6E:FB:BE:33:C9:3C:A3:84:D5"
	RootCertRandomVIDFor1_2    = "1234"

	TestRootCertPathFor1_2         = "integration_tests/constants/paa_cert_numeric_vid"
	TestRootCertSubjectFor1_2      = "MDAxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBQTEUMBIGCisGAQQBgqJ8AgEMBEZGRjE="
	TestRootCertSubjectKeyIDFor1_2 = "6A:FD:22:77:1F:51:1F:EC:BF:16:41:97:67:10:DC:DC:31:A1:71:7E"
	TestRootCertVIDFor1_2          = "65521"
	TestRootCertVIDForAssign       = "65521" // = test_root_cert_vid in 01.sh's --vid flag

	GoogleRootCertPathFor1_2         = "integration_tests/constants/google_root_cert_r2"
	GoogleRootCertSubjectFor1_2      = "MEcxCzAJBgNVBAYTAlVTMSIwIAYDVQQKExlHb29nbGUgVHJ1c3QgU2VydmljZXMgTExDMRQwEgYDVQQDEwtHVFMgUm9vdCBSMg=="
	GoogleRootCertSubjectKeyIDFor1_2 = "BB:FF:CA:8E:23:9F:4F:99:CA:DB:E2:68:A6:A5:15:27:17:1E:D9:0E"
	GoogleRootCertRandomVIDFor1_2    = "1234"

	IntermediateCertPathFor1_2         = "integration_tests/constants/intermediate_cert_gsr4"
	IntermediateCertSubjectFor1_2      = "MEYxCzAJBgNVBAYTAlVTMSIwIAYDVQQKExlHb29nbGUgVHJ1c3QgU2VydmljZXMgTExDMRMwEQYDVQQDEwpHVFMgQ0EgMkQ0"
	IntermediateCertSubjectKeyIDFor1_2 = "A8:88:D9:8A:39:AC:65:D5:82:4B:37:A8:95:6C:65:43:CD:44:01:E0"
)

// Constants used by the 1.4.3 portion of script 05.
const (
	VIDFor1_4_3                          = 65521
	PID1For1_4_3                         = 44
	PID2For1_4_3                         = 55
	PID3For1_4_3                         = 66
	DeviceTypeIDFor1_4_3                 = 4321
	ProductNameFor1_4_3                  = "ProductName13"
	ProductLabelFor1_4_3                 = "ProductLabel13"
	PartNumberFor1_4_3                   = "RCU2225B"
	SoftwareVersionFor1_4_3              = 2
	SoftwareVersionStringFor1_4_3        = "3.0"
	CDVersionNumberFor1_4_3              = 413
	MinApplicableSoftwareVersionFor1_4_3 = 3
	MaxApplicableSoftwareVersionFor1_4_3 = 3000

	CertificationTypeFor1_4_3 = "matter"
	CertificationDateFor1_4_3 = "2022-01-01T00:00:00Z"
	ProvisionalDateFor1_4_3   = "2012-12-12T00:00:00Z"
	CDCertificateIDFor1_4_3   = "12345678910abcdefgh"

	VendorNameFor1_4_3           = "Vendor65521"
	CompanyLegalNameFor1_4_3     = "LegalCompanyName65521"
	CompanyPreferredNameFor1_4_3 = "CompanyPreferredName65521"
	VendorLandingPageURLFor1_4_3 = "https://www.new65521example.com"
	VendorAccountFor1_4_3        = "vendor_account_65521"

	TestDataURLFor1_4_3 = "https://url.data.dclmodel-1.4"

	UpgradeChecksumV1_4 = "sha256:a007f58d61632af107a09c89b7392eedd05d8127d0df67ace50f318948c62001"
	PlanNameV1_4        = "v1.4"
	BinaryVersionV1_4_3 = "1.4.3"
)

// ISSUE #593 pre-upgrade ghost-model constants written by script 05's prelude
// (these are referenced by script 09 later under the "FromScript5" naming).
const (
	DeviceTypeIDForIssue593       = 4321
	ProductNameForIssue593        = "ProductName13"
	ProductLabelForIssue593       = "ProductLabel13"
	PartNumberForIssue593         = "RCU2225B"
	SoftwareVersionStringIssue593 = "3.0"
	CDVersionNumberIssue593       = 413
	MinSWVerIssue593              = 3
	MaxSWVerIssue593              = 3000
)

// 1.4.3-era PKI cert constants.
const (
	RootCertWithVIDPathFor1_4_3         = "integration_tests/constants/root_cert_with_vid"
	RootCertWithVIDSubjectFor1_4_3      = "MIGYMQswCQYDVQQGEwJVUzERMA8GA1UECAwITmV3IFlvcmsxETAPBgNVBAcMCE5ldyBZb3JrMRgwFgYDVQQKDA9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsMEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMMD3d3dy5leGFtcGxlLmNvbTEUMBIGCisGAQQBgqJ8AgEMBEZGRjE="
	RootCertWithVIDSubjectKeyIDFor1_4_3 = "CE:A8:92:66:EA:E0:80:BD:2B:B5:68:E4:0B:07:C4:FA:2C:34:6D:31"
	RootCertVIDFor1_4_3                 = 65521

	PaaCertNoVIDPathFor1_4_3         = "integration_tests/constants/paa_cert_no_vid"
	PaaCertNoVIDSubjectFor1_4_3      = "MBoxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBQQ=="
	PaaCertNoVIDSubjectKeyIDFor1_4_3 = "78:5C:E7:05:B8:6B:8F:4E:6F:C7:93:AA:60:CB:43:EA:69:68:82:D5"

	RootCertSubjectFor1_4_3      = "MIGCMQswCQYDVQQGEwJVUzERMA8GA1UECAwITmV3IFlvcmsxETAPBgNVBAcMCE5ldyBZb3JrMRgwFgYDVQQKDA9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsMEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMMD3d3dy5leGFtcGxlLmNvbQ=="
	RootCertSubjectKeyIDFor1_4_3 = "33:5E:0C:07:44:F8:B5:9C:CD:55:01:9B:6D:71:23:83:6F:D0:D4:BE"
	RootCertPathFor1_4_3         = "integration_tests/constants/root_with_same_subject_and_skid_1"

	IntermediateCertWithVIDPathFor1_4_3         = "integration_tests/constants/intermediate_cert_with_vid_1"
	IntermediateCertWithVIDSubjectFor1_4_3      = "MIGuMQswCQYDVQQGEwJVUzERMA8GA1UECAwITmV3IFlvcmsxETAPBgNVBAcMCE5ldyBZb3JrMRgwFgYDVQQKDA9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsMEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMMD3d3dy5leGFtcGxlLmNvbTEUMBIGCisGAQQBgqJ8AgEMBEZGRjExFDASBgorBgEEAYKifAICDARGRkYx"
	IntermediateCertWithVIDSubjectKeyIDFor1_4_3 = "0E:8C:E8:C8:B8:AA:50:BC:25:85:56:B9:B1:9C:C2:C7:D9:C5:2F:17"
	IntermediateCertWithVIDSerialNumberFor1_4_3 = "3"

	NOCRootCert1PathFor1_4_3         = "integration_tests/constants/noc_root_cert_1"
	NOCRootCert1SubjectFor1_4_3      = "MHoxCzAJBgNVBAYTAlVaMRMwEQYDVQQIDApTb21lIFN0YXRlMREwDwYDVQQHDAhUYXNoa2VudDEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMQ4wDAYDVQQDDAVOT0MtMQ=="
	NOCRootCert1SubjectKeyIDFor1_4_3 = "44:EB:4C:62:6B:25:48:CD:A2:B3:1C:87:41:5A:08:E7:2B:B9:83:26"

	NOCICACert1PathFor1_4_3         = "integration_tests/constants/noc_cert_1"
	NOCICACert1SubjectFor1_4_3      = "MIGCMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRQwEgYDVQQDDAtOT0MtY2hpbGQtMQ=="
	NOCICACert1SubjectKeyIDFor1_4_3 = "02:72:6E:BC:BB:EF:D6:BD:8D:9B:42:AE:D4:3C:C0:55:5F:66:3A:B3"

	CRLSignerDelegatedByPAI1         = "integration_tests/constants/leaf_cert_with_vid_65521"
	DelegatorCertWithVID65521Path    = "integration_tests/constants/intermediate_cert_with_vid_1"
	DelegatorCertWithVIDSubjectKeyID = "0E8CE8C8B8AA50BC258556B9B19CC2C7D9C52F17"
)

// Constants used by the 1.4.4 portion of script 06.
const (
	VIDFor1_4_4                          = 65522
	PID1For1_4_4                         = 77
	PID2For1_4_4                         = 88
	PID3For1_4_4                         = 99
	DeviceTypeIDFor1_4_4                 = 4433
	ProductNameFor1_4_4                  = "ProductName1.4.4"
	ProductLabelFor1_4_4                 = "ProductLabel1.4.4"
	PartNumberFor1_4_4                   = "RCU2245B"
	SoftwareVersionFor1_4_4              = 2
	SoftwareVersionStringFor1_4_4        = "4.0"
	CDVersionNumberFor1_4_4              = 513
	MinApplicableSoftwareVersionFor1_4_4 = 4
	MaxApplicableSoftwareVersionFor1_4_4 = 4000

	CertificationTypeFor1_4_4 = "matter"
	CertificationDateFor1_4_4 = "2023-01-01T00:00:00Z"
	ProvisionalDateFor1_4_4   = "2014-12-12T00:00:00Z"
	CDCertificateIDFor1_4_4   = "12345678910abcdefgh"

	VendorNameFor1_4_4           = "Vendor65522"
	CompanyLegalNameFor1_4_4     = "LegalCompanyName65522"
	CompanyPreferredNameFor1_4_4 = "CompanyPreferredName65522"
	VendorLandingPageURLFor1_4_4 = "https://www.new65522example.com"
	VendorAccountFor1_4_4        = "vendor_account_65522"

	TestDataURLFor1_4_4 = "https://url.data.dclmodel-1.4.4"

	UpgradeChecksumV1_4_4 = "sha256:e4031c6a77aa8e58add391be671a334613271bcf6e7f11d23b04a0881ece6958"
	PlanNameV1_4_4        = "v1.4.4"
	BinaryVersionV1_4_4   = "1.4.4"
)

// 1.4.4-era DA + NOC PKI cert constants.
const (
	DARootCert1PathFor1_4_4         = "integration_tests/constants/upgrade_1_4_4_da_root_cert"
	DARootCert1SubjectFor1_4_4      = "MEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQ="
	DARootCert1SubjectKeyIDFor1_4_4 = "A8:A0:95:18:9B:9F:81:4D:C7:9F:5E:B5:82:09:27:95:13:0C:9F:87"

	DAIntermediateCert1PathFor1_4_4         = "integration_tests/constants/upgrade_1_4_4_da_intermediate_cert"
	DAIntermediateCert1SubjectFor1_4_4      = "MIGCMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRQwEgYDVQQDDAtOT0MtY2hpbGQtMw=="
	DAIntermediateCert1SubjectKeyIDFor1_4_4 = "A8:A0:95:18:9B:9F:81:4D:C7:9F:5E:B5:82:09:27:95:13:0C:9F:87"
	DAIntermediateCert1SerialNumberFor1_4_4 = "3"

	DARootCert2PathFor1_4_4         = "integration_tests/constants/upgrade_1_4_4_da_root_cert_2"
	DARootCert2SubjectFor1_4_4      = "MDsxCzAJBgNVBAYTAlRFMRMwEQYDVQQIDApTb21lLVN0YXRlMRcwFQYDVQQKDA5VcGdyYWRlMS40LjRfMQ=="
	DARootCert2SubjectKeyIDFor1_4_4 = "A8:A0:95:18:9B:9F:81:4D:C7:9F:5E:B5:82:09:27:95:13:0C:9F:87"

	DAIntermediateCert2PathFor1_4_4         = "integration_tests/constants/upgrade_1_4_4_da_intermediate_cert_2"
	DAIntermediateCert2SubjectFor1_4_4      = "MIGBMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRMwEQYDVQQDDApEQS1jaGlsZC0z"
	DAIntermediateCert2SubjectKeyIDFor1_4_4 = "A8:A0:95:18:9B:9F:81:4D:C7:9F:5E:B5:82:09:27:95:13:0C:9F:87"

	NOCRootCert1V144PathFor1_4_4         = "integration_tests/constants/noc_root_cert_2"
	NOCRootCert1V144SubjectFor1_4_4      = "MHoxCzAJBgNVBAYTAlVaMRMwEQYDVQQIDApTb21lIFN0YXRlMREwDwYDVQQHDAhUYXNoa2VudDEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMQ4wDAYDVQQDDAVOT0MtMg=="
	NOCRootCert1V144SubjectKeyIDFor1_4_4 = "CF:E6:DD:37:2B:4C:B2:B9:A9:F2:75:30:1C:AA:B1:37:1B:11:7F:1B"

	NOCICACert1V144PathFor1_4_4         = "integration_tests/constants/noc_cert_2"
	NOCICACert1V144SubjectFor1_4_4      = "MIGCMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRQwEgYDVQQDDAtOT0MtY2hpbGQtMg=="
	NOCICACert1V144SubjectKeyIDFor1_4_4 = "87:48:A2:33:12:1F:51:5C:93:E6:90:40:4A:2C:AB:9E:D6:19:E5:AD"

	NOCRootCert2V144PathFor1_4_4         = "integration_tests/constants/noc_root_cert_3"
	NOCRootCert2V144SubjectFor1_4_4      = "MFUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQxDjAMBgNVBAMMBU5PQy0z"
	NOCRootCert2V144SubjectKeyIDFor1_4_4 = "88:0D:06:D9:64:22:29:34:78:7F:8C:3B:AE:F5:08:93:86:8F:0D:20"

	NOCICACert2V144PathFor1_4_4         = "integration_tests/constants/noc_cert_3"
	NOCICACert2V144SubjectFor1_4_4      = "MIGCMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRQwEgYDVQQDDAtOT0MtY2hpbGQtMw=="
	NOCICACert2V144SubjectKeyIDFor1_4_4 = "DE:F9:1D:90:D5:A1:0F:23:59:5C:3F:5C:C7:2D:31:58:2F:A8:79:33"
)

// Constants used by the 1.5.1 portion of script 07 (the rest are already on
// UpgradeTestState since Phase 1 anticipated them).
const (
	PID3For1_5_1                                = 97
	DeviceTypeIDFor1_5_1                        = 4433
	ProductNameFor1_5_1                         = "ProductName_1_5_1"
	ICDUserActiveModeTriggerHintFor1_5_1        = 4
	ICDUserActiveModeTriggerInstructionFor1_5_1 = "icd_user_active_mode_trigger_hint_for_1_5_1"
	FactoryResetStepsHintFor1_5_1               = 3
	FactoryResetStepsInstructionFor1_5_1        = "factory_reset_steps_instruction_for_1_5_1"
	SpecificationVersionFor1_5_1                = 2
	SoftwareVersionStringFor1_5_1               = "4.3"
	CDVersionNumberFor1_5_1                     = 513

	CertificationTypeFor1_5_1 = "matter"
	CertificationDateFor1_5_1 = "2024-01-01T00:00:00Z"
	ProvisionalDateFor1_5_1   = "2016-12-12T00:00:00Z"
	CDCertificateIDFor1_5_1   = "12345678910abcdefgh"

	VendorNameFor1_5_1           = "Vendor_1_5_1"
	CompanyLegalNameFor1_5_1     = "LegalCompanyName_1_5_1"
	CompanyPreferredNameFor1_5_1 = "CompanyPreferredName_1_5_1"
	VendorLandingPageURLFor1_5_1 = "https://www.new_1_5_1_example.com"
	VendorAccountFor1_5_1        = "vendor_account_1_5_1"

	TestDataURLFor1_5_1 = "https://url.data.dclmodel-1.5"

	UpgradeChecksumV1_5_1 = "sha256:21550db9f1018b7d464b0bca7440dc4aee4ee13932ff4f9e2b405b342e2e0a75"
	PlanNameV1_5_1        = "v1.5.1"
	BinaryVersionV1_5_1   = "1.5.1"
)

// Constants used by the 1.5.2 portion of script 08 (defined in 08 itself, not
// inherited from earlier scripts).
const (
	VIDFor1_5_2                                 = 65519
	PID1For1_5_2                                = 59
	PID2For1_5_2                                = 69
	PID3For1_5_2                                = 57
	DeviceTypeIDFor1_5_2                        = 4433
	ProductNameFor1_5_2                         = "ProductName_1_5_2"
	ProductLabelFor1_5_2                        = "ProductLabel_1_5_2"
	ICDUserActiveModeTriggerHintFor1_5_2        = 4
	ICDUserActiveModeTriggerInstructionFor1_5_2 = "icd_user_active_mode_trigger_hint_for_1_5_2"
	FactoryResetStepsHintFor1_5_2               = 3
	FactoryResetStepsInstructionFor1_5_2        = "factory_reset_steps_instruction_for_1_5_2"
	CommissioningModeSecondaryStepsHintFor1_5_2 = 7
	SpecificationVersionFor1_5_2                = 2
	PartNumberFor1_5_2                          = "RCU2245M"
	SoftwareVersionFor1_5_2                     = 4
	SoftwareVersionStringFor1_5_2               = "4.3"
	CDVersionNumberFor1_5_2                     = 513
	MinApplicableSoftwareVersionFor1_5_2        = 8
	MaxApplicableSoftwareVersionFor1_5_2        = 8000

	VendorAccountFor1_5_2 = "vendor_account_1_5_2"
)

// Constants used by the 1.6.0 portion of script 09.
const (
	VIDFor1_6_0                                 = 65520
	PID1For1_6_0                                = 60
	PID2For1_6_0                                = 70
	PID3For1_6_0                                = 58
	DeviceTypeIDFor1_6_0                        = 4434
	ProductNameFor1_6_0                         = "ProductName_1_6_0"
	ProductLabelFor1_6_0                        = "ProductLabel_1_6_0"
	ICDUserActiveModeTriggerHintFor1_6_0        = 5
	ICDUserActiveModeTriggerInstructionFor1_6_0 = "icd_user_active_mode_trigger_hint_for_1_6_0"
	FactoryResetStepsHintFor1_6_0               = 4
	FactoryResetStepsInstructionFor1_6_0        = "factory_reset_steps_instruction_for_1_6_0"
	CommissioningModeSecondaryStepsHintFor1_6_0 = 8
	SpecificationVersionFor1_6_0                = 3
	PartNumberFor1_6_0                          = "RCU2246M"
	SoftwareVersionFor1_6_0                     = 5
	SoftwareVersionStringFor1_6_0               = "5.0"
	CDVersionNumberFor1_6_0                     = 514
	MinApplicableSoftwareVersionFor1_6_0        = 9
	MaxApplicableSoftwareVersionFor1_6_0        = 9000

	VendorAccountFor1_6_0 = "vendor_account_1_6_0"
)

// Constants used by the master portion of script 10
// (10-test-upgrade-1.6.0-to-master.sh). The plan name itself is the master
// branch's git short hash, computed at runtime.
const (
	VIDForMaster                          = 62529
	PID1ForMaster                         = 89
	PID2ForMaster                         = 99
	PID3ForMaster                         = 77
	DeviceTypeIDForMaster                 = 3433
	ProductNameForMaster                  = "ProductName_master"
	ProductLabelForMaster                 = "ProductLabel_master"
	PartNumberForMaster                   = "ZCU2245M"
	SoftwareVersionForMaster              = 4
	SoftwareVersionStringForMaster        = "5.3"
	CDVersionNumberForMaster              = 743
	MinApplicableSoftwareVersionForMaster = 4
	MaxApplicableSoftwareVersionForMaster = 4000

	CertificationTypeForMaster = "matter"
	CertificationDateForMaster = "2024-02-01T00:00:00Z"
	ProvisionalDateForMaster   = "2016-10-12T00:00:00Z"
	CDCertificateIDForMaster   = "12345678910masterAB"

	VendorNameForMaster           = "Vendor_master"
	CompanyLegalNameForMaster     = "LegalCompanyName_master"
	CompanyPreferredNameForMaster = "CompanyPreferredName_master"
	VendorLandingPageURLForMaster = "https://www.new_master_example.com"
	VendorAccountForMaster        = "vendor_account_master"

	TestDataURLForMaster = "https://url.data.dclmodel-master"

	// MasterUpgradeImage is the locally-built image tag (matches the bash
	// MASTER_UPGRADE_IMAGE variable). Defined in test-upgrade.sh as
	// "dcld-build-master".
	MasterUpgradeImage      = "dcld-build-master"
	MasterUpgradeDockerfile = "./integration_tests/upgrade/Dockerfile-build-master"

	// DcldMasterBinaryPath is the host path where the freshly-built master
	// binary is copied for use by the test sequence.
	DcldMasterBinaryPath = BinariesDir + "/dcld_master"
)

// Constants used by script 11 (add new node after upgrade).
const (
	NewObserverNodeP2PPort    = 26570
	NewObserverNodeClientPort = 26571
	NewObserverIP             = "192.167.10.28"
	DCLDVersionV012           = "0.12.0"
)

// Upgrade plan checksums copied from the bash scripts.
const (
	UpgradeChecksumV1_5_2 = "sha256:746e4d24969f45f55b7d4a2f143ffe9609cf4f7a60c1472e38ecfe781b2327dc"
	UpgradeChecksumV1_6_0 = "sha256:63a27d50985947dc79f03421a2f7c6475adb84932b431d26d0b6d0a2e2d539ad"
)

// Plan names and binary version identifiers.
const (
	PlanNameV1_5_2 = "v1.5.2"
	PlanNameV1_6_0 = "v1.6.0"

	// BinaryVersionV1_6_0 is the release tag used by the upgrade plan info.
	// The bash carries a TODO to switch this to "1.6.0" once that release
	// ships; until then it points at the dev tag.
	BinaryVersionV1_5_2 = "1.5.2"
	BinaryVersionV1_6_0 = "1.6.0-0.dev.2"
)
