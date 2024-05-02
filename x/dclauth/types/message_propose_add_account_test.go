package types

// TODO issue 99.
import (
	fmt "fmt"
	"testing"

	"cosmossdk.io/errors"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	validator "github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/common/types"
)

func NewMsgProposeAddAccountWrapper(
	t *testing.T,
	signer sdk.AccAddress,
	address sdk.AccAddress,
	pubKey cryptotypes.PubKey,
	roles AccountRoles,
	vendorID int32,
	productIDs []*types.Uint16Range,
) *MsgProposeAddAccount {
	t.Helper()
	msg, err := NewMsgProposeAddAccount(signer, address, pubKey, roles, vendorID, productIDs, testconstants.Info)
	require.NoError(t, err)

	return msg
}

func TestNewMsgProposeAddAccount(t *testing.T) {
	msg := NewMsgProposeAddAccountWrapper(
		t,
		testconstants.Signer,
		testconstants.Address1, testconstants.PubKey1,
		AccountRoles{}, testconstants.VendorID1, testconstants.ProductIDsEmpty,
	)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "propose_add_account")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestValidateMsgProposeAddAccount(t *testing.T) {
	positiveTests := []struct {
		valid bool
		msg   *MsgProposeAddAccount
	}{
		{
			valid: true,
			msg: NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
				AccountRoles{NodeAdmin}, 1, []*types.Uint16Range{{Min: 1, Max: 1000}}),
		},
		{
			valid: true,
			msg: NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
				AccountRoles{Vendor, NodeAdmin}, testconstants.VendorID1, []*types.Uint16Range{{Min: 1, Max: 65535}}),
		},
		{
			valid: true,
			msg: NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
				AccountRoles{Vendor, NodeAdmin}, testconstants.VendorID1, testconstants.ProductIDsEmpty),
		},
		// zero VID without Vendor role - no error
		{
			valid: true,
			msg: NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
				AccountRoles{NodeAdmin}, 0, []*types.Uint16Range{{Min: 1, Max: 2}}),
		},
		{
			valid: true,
			msg: NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
				AccountRoles{Vendor, NodeAdmin}, testconstants.VendorID1, []*types.Uint16Range{{Min: 1, Max: 1}}),
		},
		{
			valid: true,
			msg: NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
				AccountRoles{Vendor, NodeAdmin}, testconstants.VendorID1, []*types.Uint16Range{{Min: 1, Max: 1}, {Min: 2, Max: 2}}),
		},
	}

	negativeTests := []struct {
		valid bool
		msg   *MsgProposeAddAccount
		err   error
	}{
		{
			valid: false,
			msg: NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
				AccountRoles{}, 1, testconstants.ProductIDsEmpty), // no roles provided
			err: errors.Wrapf(MissingRoles,
				"No roles provided"),
		},
		// zero VID with Vendor role - error - can not create Vendor with vid=0 (reserved)
		{
			valid: false,
			msg: NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
				AccountRoles{Vendor, NodeAdmin}, 0, testconstants.ProductIDsEmpty),
			err: errors.Wrapf(MissingVendorIDForVendorAccount,
				"No Vendor ID is provided in the Vendor Role for the new account"),
		},
		// negative VID - error
		{
			valid: false,
			msg: NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
				AccountRoles{Vendor, NodeAdmin}, -1, testconstants.ProductIDsEmpty),
			err: errors.Wrapf(MissingVendorIDForVendorAccount,
				"No Vendor ID is provided in the Vendor Role for the new account"),
		},
		// too large VID - error
		{
			valid: false,
			msg: NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
				AccountRoles{Vendor, NodeAdmin}, 65535+1, testconstants.ProductIDsEmpty),
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			valid: false,
			msg: NewMsgProposeAddAccountWrapper(t, testconstants.Signer, nil, testconstants.PubKey1,
				AccountRoles{NodeAdmin}, 1, testconstants.ProductIDsEmpty),
			err: sdkerrors.ErrInvalidAddress,
		},
		// {
		// valid: false,
		// msg: NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, "",
		// AccountRoles{}, 1),
		// err: sdkerrors.ErrInvalidType,
		// },
		{
			valid: false,
			msg: NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
				AccountRoles{"Wrong Role"}, 1, testconstants.ProductIDsEmpty),
			err: sdkerrors.ErrUnknownRequest,
		},
		{
			valid: false,
			msg: NewMsgProposeAddAccountWrapper(t, nil, testconstants.Address1, testconstants.PubKey1,
				AccountRoles{NodeAdmin}, 1, testconstants.ProductIDsEmpty),
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			valid: false,
			msg: NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
				AccountRoles{Vendor}, 1, []*types.Uint16Range{{Min: 1, Max: 200}, {Min: 101, Max: 200}}),
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			valid: false,
			msg: NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
				AccountRoles{Vendor}, 1, []*types.Uint16Range{{Min: 10, Max: 100}, {Min: 100, Max: 200}}),
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			valid: false,
			msg: NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
				AccountRoles{Vendor}, 1, []*types.Uint16Range{{Min: 100, Max: 1}}),
			err: sdkerrors.ErrInvalidRequest,
		},
	}

	for _, tt := range positiveTests {
		err := tt.msg.ValidateBasic()

		if tt.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}

	for _, tt := range negativeTests {
		err := tt.msg.ValidateBasic()

		if tt.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
			require.ErrorIs(t, err, tt.err)
		}
	}
}

func TestMsgProposeAddAccountGetSignBytes(t *testing.T) {
	msg := NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address2, testconstants.PubKey2,
		AccountRoles{}, testconstants.VendorID1, testconstants.ProductIDsEmpty)
	transactionTime := msg.Time
	expected := fmt.Sprintf(`{"address":"cosmos1nl4uaesk9gtu7su3n89lne6xpa6lq8gljn79rq","info":"Information for Proposal/Approval/Revoke","productIDs":[],"pubKey":{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A2wJ7uOEE5Zm04K52czFTXfDj1qF2mholzi1zOJVlKlr"},"roles":[],"signer":"cosmos1s5xf3aanx7w84hgplk9z3l90qfpantg6nsmhpf","time":"%v","vendorID":1000}`,
		transactionTime)

	require.Equal(t, expected, string(msg.GetSignBytes()))
}
