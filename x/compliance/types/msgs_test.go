package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
)

func TestValidateMsgCertifyModel(t *testing.T) {
	cases := []struct {
		valid bool
		msg   *MsgCertifyModel
	}{

		{true, newMsgCertifyModel(1, 1, 1, "1", testconstants.Signer)},
		{true, newMsgCertifyModel(65535, 65535, 1, "1", testconstants.Signer)},
		// zero PID/VID/SV - OK
		{true, newMsgCertifyModel(0, 0, 0, "1", testconstants.Signer)},

		// negative VID - not OK
		{false, newMsgCertifyModel(-1, 1, 1, "1", testconstants.Signer)},
		// negative PID - not OK
		{false, newMsgCertifyModel(1, -1, 1, "1", testconstants.Signer)},

		// too large VID - not OK
		{false, newMsgCertifyModel(65535+1, 1, 1, "1", testconstants.Signer)},
		// too large PID - not OK
		{false, newMsgCertifyModel(1, 65535+1, 1, "1", testconstants.Signer)},
	}

	for _, tc := range cases {
		err := tc.msg.ValidateBasic()

		if tc.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestValidateMsgRevokeModel(t *testing.T) {
	cases := []struct {
		valid bool
		msg   *MsgRevokeModel
	}{

		{true, newMsgRevokeModel(1, 1, 1, "1", testconstants.Signer)},
		{true, newMsgRevokeModel(65535, 65535, 1, "1", testconstants.Signer)},
		// zero PID/VID/SV - OK
		{true, newMsgRevokeModel(0, 0, 0, "1", testconstants.Signer)},

		// negative VID - not OK
		{false, newMsgRevokeModel(-1, 1, 1, "1", testconstants.Signer)},
		// negative PID - not OK
		{false, newMsgRevokeModel(1, -1, 1, "1", testconstants.Signer)},

		// too large VID - not OK
		{false, newMsgRevokeModel(65535+1, 1, 1, "1", testconstants.Signer)},
		// too large PID - not OK
		{false, newMsgRevokeModel(1, 65535+1, 1, "1", testconstants.Signer)},
	}

	for _, tc := range cases {
		err := tc.msg.ValidateBasic()

		if tc.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestValidateMsgProvisionModel(t *testing.T) {
	cases := []struct {
		valid bool
		msg   *MsgProvisionModel
	}{

		{true, newMsgProvisionModel(1, 1, 1, "1", testconstants.Signer)},
		{true, newMsgProvisionModel(65535, 65535, 1, "1", testconstants.Signer)},
		// zero PID/VID/SV - OK
		{true, newMsgProvisionModel(0, 0, 0, "1", testconstants.Signer)},

		// negative VID - not OK
		{false, newMsgProvisionModel(-1, 1, 1, "1", testconstants.Signer)},
		// negative PID - not OK
		{false, newMsgProvisionModel(1, -1, 1, "1", testconstants.Signer)},

		// too large VID - not OK
		{false, newMsgProvisionModel(65535+1, 1, 1, "1", testconstants.Signer)},
		// too large PID - not OK
		{false, newMsgProvisionModel(1, 65535+1, 1, "1", testconstants.Signer)},
	}

	for _, tc := range cases {
		err := tc.msg.ValidateBasic()

		if tc.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func newMsgCertifyModel(
	vid int32,
	pid int32,
	softwareVersion uint32,
	softwareVersionString string,
	signer sdk.AccAddress,
) *MsgCertifyModel {

	return &MsgCertifyModel{
		Signer:                signer.String(),
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       softwareVersion,
		SoftwareVersionString: softwareVersionString,
		CDVersionNumber:       uint32(testconstants.CdVersionNumber),
		CertificationDate:     testconstants.TestDate,
		CertificationType:     testconstants.CertificationType,
		Reason:                testconstants.Reason,
	}
}

func newMsgRevokeModel(
	vid int32,
	pid int32,
	softwareVersion uint32,
	softwareVersionString string,
	signer sdk.AccAddress,
) *MsgRevokeModel {

	return &MsgRevokeModel{
		Signer:                signer.String(),
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       softwareVersion,
		SoftwareVersionString: softwareVersionString,
		CDVersionNumber:       uint32(testconstants.CdVersionNumber),
		RevocationDate:        testconstants.TestDate,
		CertificationType:     testconstants.CertificationType,
		Reason:                testconstants.Reason,
	}
}

func newMsgProvisionModel(
	vid int32,
	pid int32,
	softwareVersion uint32,
	softwareVersionString string,
	signer sdk.AccAddress,
) *MsgProvisionModel {

	return &MsgProvisionModel{
		Signer:                signer.String(),
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       softwareVersion,
		SoftwareVersionString: softwareVersionString,
		CDVersionNumber:       uint32(testconstants.CdVersionNumber),
		ProvisionalDate:       testconstants.TestDate,
		CertificationType:     testconstants.CertificationType,
		Reason:                testconstants.Reason,
	}
}
