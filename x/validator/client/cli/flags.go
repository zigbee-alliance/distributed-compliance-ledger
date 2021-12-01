package cli

import (
	"flag"
)

const (
	FlagAddressValidator = "validator"
	FlagPubKey           = "pubkey"

	FlagName     = "name"
	FlagIdentity = "identity"
	FlagWebsite  = "website"
	FlagDetails  = "details"

	FlagNodeID = "node-id"
	FlagIP     = "ip"
	FlagState  = "state"

	FlagGenesisFormat = "genesis-format"
)

// common flagsets to add to various functions
var (
	fsValidator = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	fsValidator.String(FlagAddressValidator, "", "The Bech32 address of the validator")
}

// FlagSetPublicKey Returns the flagset for Public Key related operations.
func FlagSetPublicKey() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.String(FlagPubKey, "", "The validator's Protobuf JSON encoded public key")
	return fs
}

func flagSetDescriptionCreate() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagMoniker, "", "The validator's name")
	fs.String(FlagIdentity, "", "The optional identity signature (ex. UPort or Keybase)")
	fs.String(FlagWebsite, "", "The validator's (optional) website")
	fs.String(FlagSecurityContact, "", "The validator's (optional) security contact email")
	fs.String(FlagDetails, "", "The validator's (optional) details")

	return fs
}
