package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

func TestMsgCreateVendorInfo_ValidateBasic(t *testing.T) {
	negative_tests := []struct {
		name string
		msg  MsgCreateVendorInfo
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateVendorInfo{
				Creator:          "invalid_address",
				VendorID:         testconstants.VendorID1,
				VendorName:       testconstants.VendorName,
				CompanyLegalName: testconstants.CompanyLegalName,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "vid less than 0",
			msg: MsgCreateVendorInfo{
				Creator:          sample.AccAddress(),
				VendorID:         -1,
				VendorName:       testconstants.VendorName,
				CompanyLegalName: testconstants.CompanyLegalName,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid is 0",
			msg: MsgCreateVendorInfo{
				Creator:          sample.AccAddress(),
				VendorID:         0,
				VendorName:       testconstants.VendorName,
				CompanyLegalName: testconstants.CompanyLegalName,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid is bigger than 65535",
			msg: MsgCreateVendorInfo{
				Creator:          sample.AccAddress(),
				VendorID:         65536,
				VendorName:       testconstants.VendorName,
				CompanyLegalName: testconstants.CompanyLegalName,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "vendor name len < 2",
			msg: MsgCreateVendorInfo{
				Creator:          sample.AccAddress(),
				VendorID:         testconstants.VendorID1,
				VendorName:       "a",
				CompanyLegalName: testconstants.CompanyLegalName,
			},
			err: validator.ErrFieldMinLengthNotReached,
		},
		{
			name: "vendor name len > 32",
			msg: MsgCreateVendorInfo{
				Creator:          sample.AccAddress(),
				VendorID:         testconstants.VendorID1,
				VendorName:       "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",
				CompanyLegalName: testconstants.CompanyLegalName,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "company legal name len < 2",
			msg: MsgCreateVendorInfo{
				Creator:          sample.AccAddress(),
				VendorID:         testconstants.VendorID1,
				VendorName:       testconstants.VendorName,
				CompanyLegalName: "a",
			},
			err: validator.ErrFieldMinLengthNotReached,
		},
		{
			name: "company legal name len > 64",
			msg: MsgCreateVendorInfo{
				Creator:          sample.AccAddress(),
				VendorID:         testconstants.VendorID1,
				VendorName:       testconstants.VendorName,
				CompanyLegalName: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789123",
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "company preffered name len > 64",
			msg: MsgCreateVendorInfo{
				Creator:              sample.AccAddress(),
				VendorID:             testconstants.VendorID1,
				VendorName:           testconstants.VendorName,
				CompanyLegalName:     testconstants.CompanyPreferredName,
				CompanyPrefferedName: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789123",
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "vendor landing page URL is not URL",
			msg: MsgCreateVendorInfo{
				Creator:              sample.AccAddress(),
				VendorID:             testconstants.VendorID1,
				VendorName:           testconstants.VendorName,
				CompanyLegalName:     testconstants.CompanyPreferredName,
				VendorLandingPageURL: "ABC",
			},
			err: validator.ErrFieldNotValid,
		},
		{
			name: "vendor landing page URL len > 256",
			msg: MsgCreateVendorInfo{
				Creator:              sample.AccAddress(),
				VendorID:             testconstants.VendorID1,
				VendorName:           testconstants.VendorName,
				CompanyLegalName:     testconstants.CompanyPreferredName,
				VendorLandingPageURL: "https://www.example.com/ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
	}

	positive_tests := []struct {
		name string
		msg  MsgCreateVendorInfo
	}{
		{
			name: "valid address",
			msg: MsgCreateVendorInfo{
				Creator:          sample.AccAddress(),
				VendorID:         testconstants.VendorID1,
				VendorName:       testconstants.VendorName,
				CompanyLegalName: testconstants.CompanyLegalName,
			},
		},
	}
	for _, tt := range positive_tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.NoError(t, err)
		})
	}

	for _, tt := range negative_tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.Error(t, err)
			require.ErrorIs(t, err, tt.err)
		})
	}
}

func TestMsgUpdateVendorInfo_ValidateBasic(t *testing.T) {
	negative_tests := []struct {
		name string
		msg  MsgUpdateVendorInfo
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateVendorInfo{
				Creator:          "invalid_address",
				VendorID:         testconstants.VendorID1,
				VendorName:       testconstants.VendorName,
				CompanyLegalName: testconstants.CompanyLegalName,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "vid less than 0",
			msg: MsgUpdateVendorInfo{
				Creator:          sample.AccAddress(),
				VendorID:         -1,
				VendorName:       testconstants.VendorName,
				CompanyLegalName: testconstants.CompanyLegalName,
			},
		},
		{
			name: "vid is 0",
			msg: MsgUpdateVendorInfo{
				Creator:          sample.AccAddress(),
				VendorID:         0,
				VendorName:       testconstants.VendorName,
				CompanyLegalName: testconstants.CompanyLegalName,
			},
		},
		{
			name: "vid is bigger than 65535",
			msg: MsgUpdateVendorInfo{
				Creator:          sample.AccAddress(),
				VendorID:         65536,
				VendorName:       testconstants.VendorName,
				CompanyLegalName: testconstants.CompanyLegalName,
			},
		},
		{
			name: "vendor name len < 2",
			msg: MsgUpdateVendorInfo{
				Creator:          sample.AccAddress(),
				VendorID:         testconstants.VendorID1,
				VendorName:       "a",
				CompanyLegalName: testconstants.CompanyLegalName,
			},
		},
		{
			name: "vendor name len > 32",
			msg: MsgUpdateVendorInfo{
				Creator:          sample.AccAddress(),
				VendorID:         testconstants.VendorID1,
				VendorName:       "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",
				CompanyLegalName: testconstants.CompanyLegalName,
			},
		},
		{
			name: "company legal name len < 2",
			msg: MsgUpdateVendorInfo{
				Creator:          sample.AccAddress(),
				VendorID:         testconstants.VendorID1,
				VendorName:       testconstants.VendorName,
				CompanyLegalName: "a",
			},
		},
		{
			name: "company legal name len > 64",
			msg: MsgUpdateVendorInfo{
				Creator:          sample.AccAddress(),
				VendorID:         testconstants.VendorID1,
				VendorName:       testconstants.VendorName,
				CompanyLegalName: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789123",
			},
		},
		{
			name: "company preffered name len > 64",
			msg: MsgUpdateVendorInfo{
				Creator:              sample.AccAddress(),
				VendorID:             testconstants.VendorID1,
				VendorName:           testconstants.VendorName,
				CompanyLegalName:     testconstants.CompanyPreferredName,
				CompanyPrefferedName: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789123",
			},
		},
		{
			name: "vendor landing page URL is not URL",
			msg: MsgUpdateVendorInfo{
				Creator:              sample.AccAddress(),
				VendorID:             testconstants.VendorID1,
				VendorName:           testconstants.VendorName,
				CompanyLegalName:     testconstants.CompanyPreferredName,
				VendorLandingPageURL: "ABC",
			},
		},
		{
			name: "vendor landing page URL len > 256",
			msg: MsgUpdateVendorInfo{
				Creator:              sample.AccAddress(),
				VendorID:             testconstants.VendorID1,
				VendorName:           testconstants.VendorName,
				CompanyLegalName:     testconstants.CompanyPreferredName,
				VendorLandingPageURL: "https://www.example.com/ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",
			},
		},
	}

	positive_tests := []struct {
		name string
		msg  MsgCreateVendorInfo
		err  error
	}{
		{
			name: "valid address",
			msg: MsgCreateVendorInfo{
				Creator:          sample.AccAddress(),
				VendorID:         testconstants.VendorID1,
				VendorName:       testconstants.VendorName,
				CompanyLegalName: testconstants.CompanyLegalName,
			},
		},
	}
	for _, tt := range positive_tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.NoError(t, err)
		})
	}

	for _, tt := range negative_tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.Error(t, err)
		})
	}
}
