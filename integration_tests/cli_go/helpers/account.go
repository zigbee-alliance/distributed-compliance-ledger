package helpers

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/go-bip39"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

type AccountInfo struct {
	Name    string
	Address string
	Key     string
}

func CreateAccountInfo(suite *utils.TestSuite) AccountInfo {
	name := RandomString()
	entropySeed, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropySeed)
	account, _ := suite.Kr.NewAccount(name, mnemonic, testconstants.Passphrase, sdk.FullFundraiserPath, hd.Secp256k1)

	address, _ := account.GetAddress()
	pubKey, _ := account.GetPubKey()

	return AccountInfo{
		Name:    name,
		Address: address.String(),
		Key:     FormatKey(pubKey),
	}
}

func FormatKey(pk cryptotypes.PubKey) string {
	apk, _ := codectypes.NewAnyWithValue(pk)
	bz, _ := codec.ProtoMarshalJSON(apk, nil)

	return string(bz)
}
