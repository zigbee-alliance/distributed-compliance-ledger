package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

func TestMsgCreateVendorInfo_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
		name string
		msg  MsgCreateVendorInfo
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateVendorInfo{
				Creator:               "invalid_address",
				VendorID:              testconstants.VendorID1,
				VendorName:            testconstants.VendorName,
				CompanyLegalName:      testconstants.CompanyLegalName,
				CompanyPrefferredName: testconstants.CompanyPreferredName,
				VendorLandingPageURL:  testconstants.VendorLandingPageURL,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "vendor name is not set",
			msg: MsgCreateVendorInfo{
				Creator:               sample.AccAddress(),
				VendorID:              testconstants.VendorID1,
				VendorName:            "",
				CompanyLegalName:      testconstants.CompanyLegalName,
				CompanyPrefferredName: testconstants.CompanyPreferredName,
				VendorLandingPageURL:  testconstants.VendorLandingPageURL,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "company legal name is not set",
			msg: MsgCreateVendorInfo{
				Creator:               sample.AccAddress(),
				VendorID:              testconstants.VendorID1,
				VendorName:            testconstants.VendorName,
				CompanyLegalName:      "",
				CompanyPrefferredName: testconstants.CompanyPreferredName,
				VendorLandingPageURL:  testconstants.VendorLandingPageURL,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "vid less than 0",
			msg: MsgCreateVendorInfo{
				Creator:               sample.AccAddress(),
				VendorID:              -1,
				VendorName:            testconstants.VendorName,
				CompanyLegalName:      testconstants.CompanyLegalName,
				CompanyPrefferredName: testconstants.CompanyPreferredName,
				VendorLandingPageURL:  testconstants.VendorLandingPageURL,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid is 0",
			msg: MsgCreateVendorInfo{
				Creator:               sample.AccAddress(),
				VendorID:              0,
				VendorName:            testconstants.VendorName,
				CompanyLegalName:      testconstants.CompanyLegalName,
				CompanyPrefferredName: testconstants.CompanyPreferredName,
				VendorLandingPageURL:  testconstants.VendorLandingPageURL,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid is bigger than 65535",
			msg: MsgCreateVendorInfo{
				Creator:               sample.AccAddress(),
				VendorID:              65536,
				VendorName:            testconstants.VendorName,
				CompanyLegalName:      testconstants.CompanyLegalName,
				CompanyPrefferredName: testconstants.CompanyPreferredName,
				VendorLandingPageURL:  testconstants.VendorLandingPageURL,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "vendor name len > 128",
			msg: MsgCreateVendorInfo{
				Creator:               sample.AccAddress(),
				VendorID:              testconstants.VendorID1,
				VendorName:            tmrand.Str(129),
				CompanyLegalName:      testconstants.CompanyLegalName,
				CompanyPrefferredName: testconstants.CompanyPreferredName,
				VendorLandingPageURL:  testconstants.VendorLandingPageURL,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "company legal name len > 256",
			msg: MsgCreateVendorInfo{
				Creator:               sample.AccAddress(),
				VendorID:              testconstants.VendorID1,
				VendorName:            testconstants.VendorName,
				CompanyLegalName:      tmrand.Str(257),
				CompanyPrefferredName: testconstants.CompanyPreferredName,
				VendorLandingPageURL:  testconstants.VendorLandingPageURL,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "company preffered name len > 256",
			msg: MsgCreateVendorInfo{
				Creator:               sample.AccAddress(),
				VendorID:              testconstants.VendorID1,
				VendorName:            testconstants.VendorName,
				CompanyLegalName:      testconstants.CompanyPreferredName,
				CompanyPrefferredName: tmrand.Str(257),
				VendorLandingPageURL:  testconstants.VendorLandingPageURL,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "vendor landing page URL is not URL",
			msg: MsgCreateVendorInfo{
				Creator:               sample.AccAddress(),
				VendorID:              testconstants.VendorID1,
				VendorName:            testconstants.VendorName,
				CompanyLegalName:      testconstants.CompanyPreferredName,
				VendorLandingPageURL:  "ABC",
				CompanyPrefferredName: testconstants.CompanyPreferredName,
			},
			err: validator.ErrFieldNotValid,
		},
		{
			name: "vendor landing page URL len > 256",
			msg: MsgCreateVendorInfo{
				Creator:               sample.AccAddress(),
				VendorID:              testconstants.VendorID1,
				VendorName:            testconstants.VendorName,
				CompanyLegalName:      testconstants.CompanyPreferredName,
				VendorLandingPageURL:  "https://www.example.com/ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",
				CompanyPrefferredName: testconstants.CompanyPreferredName,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
	}

	positiveTests := []struct {
		name string
		msg  MsgCreateVendorInfo
	}{
		{
			name: "valid create vendorinfo message",
			msg: MsgCreateVendorInfo{
				Creator:               sample.AccAddress(),
				VendorID:              testconstants.VendorID1,
				VendorName:            testconstants.VendorName,
				CompanyLegalName:      testconstants.CompanyLegalName,
				CompanyPrefferredName: testconstants.CompanyPreferredName,
				VendorLandingPageURL:  testconstants.VendorLandingPageURL,
			},
		},
	}
	for _, tt := range positiveTests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.NoError(t, err)
		})
	}

	for _, tt := range negativeTests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.Error(t, err)
			require.ErrorIs(t, err, tt.err)
		})
	}
}

func TestMsgUpdateVendorInfo_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
		name string
		msg  MsgUpdateVendorInfo
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateVendorInfo{
				Creator:               "invalid_address",
				VendorID:              testconstants.VendorID1,
				VendorName:            testconstants.VendorName,
				CompanyLegalName:      testconstants.CompanyLegalName,
				CompanyPrefferredName: testconstants.CompanyPreferredName,
				VendorLandingPageURL:  testconstants.VendorLandingPageURL,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "vendor name is not set",
			msg: MsgUpdateVendorInfo{
				Creator:               sample.AccAddress(),
				VendorID:              testconstants.VendorID1,
				VendorName:            "",
				CompanyLegalName:      testconstants.CompanyLegalName,
				CompanyPrefferredName: testconstants.CompanyPreferredName,
				VendorLandingPageURL:  testconstants.VendorLandingPageURL,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "company legal name is not set",
			msg: MsgUpdateVendorInfo{
				Creator:               sample.AccAddress(),
				VendorID:              testconstants.VendorID1,
				VendorName:            testconstants.VendorName,
				CompanyLegalName:      "",
				CompanyPrefferredName: testconstants.CompanyPreferredName,
				VendorLandingPageURL:  testconstants.VendorLandingPageURL,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "vid less than 0",
			msg: MsgUpdateVendorInfo{
				Creator:               sample.AccAddress(),
				VendorID:              -1,
				VendorName:            testconstants.VendorName,
				CompanyLegalName:      testconstants.CompanyLegalName,
				CompanyPrefferredName: testconstants.CompanyPreferredName,
				VendorLandingPageURL:  testconstants.VendorLandingPageURL,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid is 0",
			msg: MsgUpdateVendorInfo{
				Creator:               sample.AccAddress(),
				VendorID:              0,
				VendorName:            testconstants.VendorName,
				CompanyLegalName:      testconstants.CompanyLegalName,
				CompanyPrefferredName: testconstants.CompanyPreferredName,
				VendorLandingPageURL:  testconstants.VendorLandingPageURL,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid is bigger than 65535",
			msg: MsgUpdateVendorInfo{
				Creator:               sample.AccAddress(),
				VendorID:              65536,
				VendorName:            testconstants.VendorName,
				CompanyLegalName:      testconstants.CompanyLegalName,
				CompanyPrefferredName: testconstants.CompanyPreferredName,
				VendorLandingPageURL:  testconstants.VendorLandingPageURL,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "vendor name len > 128",
			msg: MsgUpdateVendorInfo{
				Creator:               sample.AccAddress(),
				VendorID:              testconstants.VendorID1,
				VendorName:            tmrand.Str(129),
				CompanyLegalName:      testconstants.CompanyLegalName,
				CompanyPrefferredName: testconstants.CompanyPreferredName,
				VendorLandingPageURL:  testconstants.VendorLandingPageURL,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "company legal name len > 256",
			msg: MsgUpdateVendorInfo{
				Creator:               sample.AccAddress(),
				VendorID:              testconstants.VendorID1,
				VendorName:            testconstants.VendorName,
				CompanyLegalName:      tmrand.Str(257),
				CompanyPrefferredName: testconstants.CompanyPreferredName,
				VendorLandingPageURL:  testconstants.VendorLandingPageURL,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "company preffered name len > 256",
			msg: MsgUpdateVendorInfo{
				Creator:               sample.AccAddress(),
				VendorID:              testconstants.VendorID1,
				VendorName:            testconstants.VendorName,
				CompanyLegalName:      testconstants.CompanyPreferredName,
				CompanyPrefferredName: tmrand.Str(257),
				VendorLandingPageURL:  testconstants.VendorLandingPageURL,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "vendor landing page URL is not URL",
			msg: MsgUpdateVendorInfo{
				Creator:               sample.AccAddress(),
				VendorID:              testconstants.VendorID1,
				VendorName:            testconstants.VendorName,
				CompanyLegalName:      testconstants.CompanyPreferredName,
				VendorLandingPageURL:  "ABC",
				CompanyPrefferredName: testconstants.CompanyPreferredName,
			},
			err: validator.ErrFieldNotValid,
		},
		{
			name: "vendor landing page URL len > 256",
			msg: MsgUpdateVendorInfo{
				Creator:               sample.AccAddress(),
				VendorID:              testconstants.VendorID1,
				VendorName:            testconstants.VendorName,
				CompanyLegalName:      testconstants.CompanyPreferredName,
				VendorLandingPageURL:  "https://www.example.com/ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",
				CompanyPrefferredName: testconstants.CompanyPreferredName,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
	}

	positiveTests := []struct {
		name string
		msg  MsgUpdateVendorInfo
	}{
		{
			name: "valid create vendorinfo message",
			msg: MsgUpdateVendorInfo{
				Creator:               sample.AccAddress(),
				VendorID:              testconstants.VendorID1,
				VendorName:            testconstants.VendorName,
				CompanyLegalName:      testconstants.CompanyLegalName,
				CompanyPrefferredName: testconstants.CompanyPreferredName,
				VendorLandingPageURL:  testconstants.VendorLandingPageURL,
			},
		},
		{
			name: "valid create vendorinfo message",
			msg: MsgUpdateVendorInfo{
				Creator:              sample.AccAddress(),
				VendorID:             testconstants.VendorID1,
				VendorName:           testconstants.VendorName,
				CompanyLegalName:     testconstants.CompanyLegalName,
				VendorLandingPageURL: testconstants.VendorLandingPageURL,
			},
		},
	}
	for _, tt := range positiveTests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.NoError(t, err)
		})
	}

	for _, tt := range negativeTests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.Error(t, err)
			require.ErrorIs(t, err, tt.err)
		})
	}
}
