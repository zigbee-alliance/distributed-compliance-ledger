package cli

func Query(module string, query string, queryArgs []string) (string, error) {
	args := []string{"query", module, query}

	args = append(args, queryArgs...)

	resp, err := Execute(args...)

	return resp, err
}

func ProposedAccount(info string, user_address string, user_pubkey string, roles string, vid_in_hex_format string, from string) (string, error) {

	var queryArgs []string

	queryArgs = append(queryArgs, "--info", info)
	queryArgs = append(queryArgs, "--address", user_address)
	queryArgs = append(queryArgs, "--pubkey", user_pubkey)
	queryArgs = append(queryArgs, "--roles", roles)
	queryArgs = append(queryArgs, "--vid", vid_in_hex_format)
	queryArgs = append(queryArgs, "--from", from)
	queryArgs = append(queryArgs, "--yes")

	return Query("auth", "propose-add-account", queryArgs)
}

func AllProposedAccounts() (string, error) {

	resp, err := Execute("query", "auth", "all-accounts")

	return resp, err
}

func Account(user_address string) (string, error) {

	var queryArgs []string

	queryArgs = append(queryArgs, "--address=", user_address)

	return Query("auth", "account", queryArgs)
}

func GetModel(vid_in_hex_format string, pid_in_hex_format string) (string, error) {

	var queryArgs []string

	queryArgs = append(queryArgs, "--vid", vid_in_hex_format)
	queryArgs = append(queryArgs, "--pid", pid_in_hex_format)

	return Query("model", "get-model", queryArgs)
}

func VendorModel(vid_in_hex_format string) (string, error) {

	var queryArgs []string

	queryArgs = append(queryArgs, "--vid", vid_in_hex_format)

	return Query("model", "vendor-models", queryArgs)

}

func AllProposedModels() (string, error) {

	resp, err := Execute("query", "model", "all-models")

	return resp, err
}

func Vendor(vid_in_hex_format string) (string, error) {

	var queryArgs []string

	queryArgs = append(queryArgs, "--vid", vid_in_hex_format)
	queryArgs = append(queryArgs, "--yes")

	return Query("vendorinfo", "vendor", queryArgs)
}

func AllVendors() (string, error) {
	resp, err := Execute("query", "vendorinfo", "all-vendors")

	return resp, err
}

func ModelVersion(vid, pid, sv string) (string, error) {

	var queryArgs []string

	queryArgs = append(queryArgs, "--vid", vid)
	queryArgs = append(queryArgs, "--pid", pid)
	queryArgs = append(queryArgs, "--softwareVersion", sv)

	return Query("model", "model-version", queryArgs)
}

func AllModelVersion(vid, pid string) (string, error) {

	var queryArgs []string

	queryArgs = append(queryArgs, "--vid", vid)
	queryArgs = append(queryArgs, "--pid", pid)

	return Query("model", "all-model-versions", queryArgs)
}

func ProposeAccount(address string) (string, error) {

	var queryArgs []string

	queryArgs = append(queryArgs, "--address", address)

	return Query("auth", "proposed-account", queryArgs)
}

func X509Cert(subject, id string) (string, error) {

	var queryArgs []string

	queryArgs = append(queryArgs, "--subject", subject)
	queryArgs = append(queryArgs, "--subject-key-id", id)
	queryArgs = append(queryArgs, "--yes")

	return Query("pki", "x509-cert", queryArgs)
}

func RevokedX509(subject, id string) (string, error) {

	var queryArgs []string

	queryArgs = append(queryArgs, "--subject", subject)
	queryArgs = append(queryArgs, "--subject-key-id", id)
	queryArgs = append(queryArgs, "--yes")

	return Query("pki", "revoked-x509-cert", queryArgs)
}
