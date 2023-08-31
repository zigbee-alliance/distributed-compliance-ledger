package cli

const (
	FlagVersionID = "--version-id"
)

func Tx(module string, tx string, txArgs []string) (string, error) {
	args := []string{"tx", module, tx}

	// Other args
	args = append(args, txArgs...)

	resp, err := Execute(args...)

	return resp, err
}

func ProposeAddAccount(info string, user_address string, user_pubkey string, roles string, vid_in_hex_format string, from string) (string, error) {

	var txArgs []string

	txArgs = append(txArgs, "--info", info)
	txArgs = append(txArgs, "--address", user_address)
	txArgs = append(txArgs, "--pubkey", user_pubkey)
	txArgs = append(txArgs, "--roles", roles)
	txArgs = append(txArgs, "--vid", vid_in_hex_format)
	txArgs = append(txArgs, "--from", from)
	txArgs = append(txArgs, "--yes")

	return Tx("auth", "propose-add-account", txArgs)
}

func AddModel(vid_in_hex_format string, pid_in_hex_format string, productName string, productLabel string, CustomFlow string, typeId string, partNumber string, from string) (string, error) {

	var txArgs []string

	txArgs = append(txArgs, "--vid", vid_in_hex_format)
	txArgs = append(txArgs, "--pid", pid_in_hex_format)
	txArgs = append(txArgs, "--productName", productName)
	txArgs = append(txArgs, "--productLabel", productLabel)
	txArgs = append(txArgs, "--commissioningCustomFlow", CustomFlow)
	txArgs = append(txArgs, "--deviceTypeID", typeId)
	txArgs = append(txArgs, "--partNumber", partNumber)
	txArgs = append(txArgs, "--from", from)
	txArgs = append(txArgs, "--yes")

	return Tx("model", "add-model", txArgs)
}

func UpdateModel(vid_in_hex_format string, pid_in_hex_format string, productName string, productLabel string, partNumber string, from string) (string, error) {

	var txArgs []string

	txArgs = append(txArgs, "--vid", vid_in_hex_format)
	txArgs = append(txArgs, "--pid", pid_in_hex_format)
	txArgs = append(txArgs, "--productName", productName)
	txArgs = append(txArgs, "--productLabel", productLabel)
	txArgs = append(txArgs, "--partNumber", partNumber)
	txArgs = append(txArgs, "--from", from)
	txArgs = append(txArgs, "--yes")

	return Tx("model", "update-model", txArgs)
}

func DeleteModel(vid_in_hex_format string, pid_in_hex_format string, from string) (string, error) {

	var txArgs []string

	txArgs = append(txArgs, "--vid", vid_in_hex_format)
	txArgs = append(txArgs, "--pid", pid_in_hex_format)
	txArgs = append(txArgs, "--from", from)
	txArgs = append(txArgs, "--yes")

	return Tx("model", "delete-model", txArgs)
}

func AddVendor(vid_in_hex_format string, company_legal_name string, vendorName string, from string) (string, error) {

	var txArgs []string

	txArgs = append(txArgs, "--vid", vid_in_hex_format)
	txArgs = append(txArgs, "--companyLegalName", company_legal_name)
	txArgs = append(txArgs, "--vendorName", vendorName)
	txArgs = append(txArgs, "--from", from)
	txArgs = append(txArgs, "--yes")

	return Tx("vendorinfo", "add-vendor", txArgs)
}

func UpdateVendor(vid_in_hex_format string, company_legal_name string, vendorLandingPageURL string, vendorName string, from string) (string, error) {

	var txArgs []string

	txArgs = append(txArgs, "--vid", vid_in_hex_format)
	txArgs = append(txArgs, "--companyLegalName", company_legal_name)
	txArgs = append(txArgs, "--vendorLandingPageURL", vendorLandingPageURL)
	txArgs = append(txArgs, "--vendorName", vendorName)
	txArgs = append(txArgs, "--from", from)
	txArgs = append(txArgs, "--yes")

	return Tx("vendorinfo", "update-vendor", txArgs)
}

func ProposeAddx509(certificate string, vid string, from string) (string, error) {

	var txArgs []string

	txArgs = append(txArgs, "--certificate", certificate)
	txArgs = append(txArgs, "--vid", vid)
	txArgs = append(txArgs, "--from", from)
	txArgs = append(txArgs, "--yes")

	return Tx("pki", "propose-add-x509-root-cert", txArgs)
}

func ApproveAddx509(subject string, subject_key_id string, from string) (string, error) {

	var txArgs []string

	txArgs = append(txArgs, "--subject", subject)
	txArgs = append(txArgs, "-subject-key-id", subject_key_id)
	txArgs = append(txArgs, "--from", from)
	txArgs = append(txArgs, "--yes")

	return Tx("pki", "approve-add-x509-root-cert", txArgs)
}

func AssignVid(subject string, subject_key_id string, vid string, from string) (string, error) {

	var txArgs []string

	txArgs = append(txArgs, "--subject", subject)
	txArgs = append(txArgs, "-subject-key-id", subject_key_id)
	txArgs = append(txArgs, "-vid", vid)
	txArgs = append(txArgs, "--from", from)
	txArgs = append(txArgs, "--yes")

	return Tx("pki", "assign-vid", txArgs)
}

func AddModelVersion(cdVersionNumber int, maxApplicableSoftwareVersion int, minApplicableSoftwareVersion int, vid string, pid string, softwareVersion string, softwareVersionString int, from string) (string, error) {

	var txArgs []string

	txArgs = append(txArgs, "--cdVersionNumber", string(rune(cdVersionNumber)))
	txArgs = append(txArgs, "--maxApplicableSoftwareVersion", string(rune(maxApplicableSoftwareVersion)))
	txArgs = append(txArgs, "--minApplicableSoftwareVersion", string(rune(minApplicableSoftwareVersion)))
	txArgs = append(txArgs, "--vid", vid)
	txArgs = append(txArgs, "--pid", pid)
	txArgs = append(txArgs, "--softwareVersion", softwareVersion)
	txArgs = append(txArgs, "--softwareVersionString", string(rune(softwareVersionString)))
	txArgs = append(txArgs, "--from", from)
	txArgs = append(txArgs, "--yes")

	return Tx("model", "add-model-version", txArgs)
}

func UpdateModelVersion(vid string, pid string, minApplicableSoftwareVersion int, maxApplicableSoftwareVersion int, sv string, softwareVersionValid bool, from string) (string, error) {

	var txArgs []string

	txArgs = append(txArgs, "--vid", vid)
	txArgs = append(txArgs, "--pid", pid)
	txArgs = append(txArgs, "--minApplicableSoftwareVersion", string(rune(minApplicableSoftwareVersion)))
	txArgs = append(txArgs, "--maxApplicableSoftwareVersion", string(rune(maxApplicableSoftwareVersion)))
	txArgs = append(txArgs, "--softwareVersion", sv)
	txArgs = append(txArgs, "--from", from)
	txArgs = append(txArgs, "--yes")

	return Tx("model", "update-model-version", txArgs)
}

func DeleteModelVersion(vid string, pid string, sv string, from string) (string, error) {

	var txArgs []string

	txArgs = append(txArgs, "--vid", vid)
	txArgs = append(txArgs, "--pid", pid)
	txArgs = append(txArgs, "--softwareVersion", sv)
	txArgs = append(txArgs, "--from", from)
	txArgs = append(txArgs, "--yes")

	return Tx("model", "delete-model-version", txArgs)
}

func CertifyModel(vid string, pid string, sv string, cdVersionNumber string, certificationType string, certificationDate string, softwareVersionString string, cdCertificateId string, from string) (string, error) {

	var txArgs []string

	txArgs = append(txArgs, "--vid", vid)
	txArgs = append(txArgs, "--pid", pid)
	txArgs = append(txArgs, "--softwareVersion", sv)
	txArgs = append(txArgs, "--cdVersionNumber", cdVersionNumber)
	txArgs = append(txArgs, "--certificationType", certificationType)
	txArgs = append(txArgs, "--certificationDate", certificationDate)
	txArgs = append(txArgs, "--softwareVersionString", softwareVersionString)
	txArgs = append(txArgs, "--cdCertificateId", cdCertificateId)
	txArgs = append(txArgs, "--from", from)
	txArgs = append(txArgs, "--yes")

	return Tx("complaince", "certify-model", txArgs)
}

func Approve_add_account(address string, from string) (string, error) {

	var txArgs []string

	txArgs = append(txArgs, "--address", address)
	txArgs = append(txArgs, "--from", from)
	txArgs = append(txArgs, "--yes")

	return Tx("auth", "approve-add-account", txArgs)
}

func ProposeRevokeX509(subject, id, from string) (string, error) {

	var txArgs []string

	txArgs = append(txArgs, "--subject", subject)
	txArgs = append(txArgs, "--subject-key-id", id)
	txArgs = append(txArgs, "--from", from)
	txArgs = append(txArgs, "--yes")

	return Tx("pki", "propose-revoke-x509-root-cert", txArgs)
}

func ProposedX509CertToRevoke(subject, id string) (string, error) {

	var txArgs []string

	txArgs = append(txArgs, "--subject", subject)
	txArgs = append(txArgs, "--subject-key-id", id)
	txArgs = append(txArgs, "--yes")

	return Tx("pki", "proposed-x509-root-cert-to-revoke", txArgs)
}

func ApproveRevokeX509(subject, id, from string) (string, error) {

	var txArgs []string

	txArgs = append(txArgs, "--subject", subject)
	txArgs = append(txArgs, "--subject-key-id", id)
	txArgs = append(txArgs, "--from", from)
	txArgs = append(txArgs, "--yes")

	return Tx("pki", "approve-revoke-x509-root-cert", txArgs)
}
