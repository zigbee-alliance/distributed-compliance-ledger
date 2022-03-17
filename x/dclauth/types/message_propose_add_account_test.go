package types

// TODO issue 99
import (
	fmt "fmt"
	"testing"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
)

func NewMsgProposeAddAccountWrapper(
	t *testing.T,
	signer sdk.AccAddress,
	address sdk.AccAddress,
	pubKey cryptotypes.PubKey,
	roles AccountRoles,
	vendorID int32,
) *MsgProposeAddAccount {
	msg, err := NewMsgProposeAddAccount(signer, address, pubKey, roles, vendorID, testconstants.Info)
	require.NoError(t, err)
	return msg
}

func TestNewMsgProposeAddAccount(t *testing.T) {
	msg := NewMsgProposeAddAccountWrapper(
		t,
		testconstants.Signer,
		testconstants.Address1, testconstants.PubKey1,
		AccountRoles{}, testconstants.VendorID1,
	)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "propose_add_account")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestValidateMsgProposeAddAccount(t *testing.T) {
	tests := []struct {
		valid bool
		msg   *MsgProposeAddAccount
	}{
		{
			valid: false,
			msg: NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
				AccountRoles{}, 1), // no roles provided
		},
		{
			valid: true,
			msg: NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
				AccountRoles{NodeAdmin}, 1),
		},
		{
			valid: true,
			msg: NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
				AccountRoles{Vendor, NodeAdmin}, testconstants.VendorID1),
		},
		// zero VID with Vendor role - error - can not create Vendor with vid=0 (reserved)
		{
			valid: false,
			msg: NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
				AccountRoles{Vendor, NodeAdmin}, 0),
		},
		// zero VID without Vendor role - no error
		{
			valid: true,
			msg: NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
				AccountRoles{NodeAdmin}, 0),
		},
		// negative VID - error
		{
			valid: false,
			msg: NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
				AccountRoles{Vendor, NodeAdmin}, -1),
		},
		// too large VID - error
		{
			valid: false,
			msg: NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
				AccountRoles{Vendor, NodeAdmin}, 65535+1),
		},
		{
			valid: true,
			msg: NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
				AccountRoles{Vendor, NodeAdmin}, testconstants.VendorID1),
		},
		{
			valid: false,
			msg: NewMsgProposeAddAccountWrapper(t, testconstants.Signer, nil, testconstants.PubKey1,
				AccountRoles{NodeAdmin}, 1),
		},
		//{
		//valid: false,
		//msg: NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, "",
		//AccountRoles{}, 1),
		//},
		{
			valid: false,
			msg: NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
				AccountRoles{"Wrong Role"}, 1),
		},
		{
			valid: false,
			msg: NewMsgProposeAddAccountWrapper(t, nil, testconstants.Address1, testconstants.PubKey1,
				AccountRoles{NodeAdmin}, 1),
		},
	}

	for _, tt := range tests {
		err := tt.msg.ValidateBasic()

		if tt.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestMsgProposeAddAccountGetSignBytes(t *testing.T) {
	msg := NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address2, testconstants.PubKey2,
		AccountRoles{}, testconstants.VendorID1)
	transcationTime := msg.Time
	expected := fmt.Sprintf(`{"address":"cosmos1nl4uaesk9gtu7su3n89lne6xpa6lq8gljn79rq","info":"Information for Proposal/Approval/Revoke","pubKey":{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A2wJ7uOEE5Zm04K52czFTXfDj1qF2mholzi1zOJVlKlr"},"roles":[],"signer":"cosmos1s5xf3aanx7w84hgplk9z3l90qfpantg6nsmhpf","time":"%v","vendorID":1000}`,
		transcationTime)

	require.Equal(t, expected, string(msg.GetSignBytes()))
}
