package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagAddress = "address"
	FlagPubKey  = "pubkey"

	FlagMoniker  = "moniker"
	FlagIdentity = "identity"
	FlagWebsite  = "website"
	FlagDetails  = "details"

	FlagNodeID = "node-id"
	FlagIP     = "ip"
	FlagState  = "state"

	FlagGenesisFormat = "genesis-format"
	FlagInfo          = "info"
)

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
	fs.String(FlagDetails, "", "The validator's (optional) details")

	return fs
}
