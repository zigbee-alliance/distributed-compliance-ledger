package compliance

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/test_constants"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHandler_HandleAddDeleteModelInfo(t *testing.T) {
	setup := Setup()
	owner := setup.Manufacturer()

	// add new model
	modelInfo := TestMsgAddModelInfo(owner)
	result := setup.Handler(setup.Ctx, modelInfo)
	require.Equal(t, sdk.CodeOK, result.Code)

	// delete model
	result = setup.Handler(setup.Ctx, MsgDeleteModelInfo{
		ID:     modelInfo.ID,
		Signer: modelInfo.Owner,
	})
	require.Equal(t, sdk.CodeOK, result.Code)
}

func TestHandler_HandleUpdateModel(t *testing.T) {
	setup := Setup()
	owner := setup.Manufacturer()

	// try update not present model
	result := setup.Handler(setup.Ctx, MsgUpdateModelInfo{
		ID:                          test_constants.Id,
		NewName:                     test_constants.Name,
		NewOwner:                    owner,
		NewDescription:              test_constants.Description,
		NewSKU:                      test_constants.Sku,
		NewFirmwareVersion:          test_constants.FirmwareVersion,
		NewHardwareVersion:          test_constants.HardwareVersion,
		NewCertificateID:            test_constants.CertificateID,
		NewCertifiedDate:            test_constants.CertifiedDate,
		NewTisOrTrpTestingCompleted: false,
		Signer:                      owner,
	})
	require.Equal(t, types.CodeModelInfoDoesNotExist, result.Code)

	// add new model
	modelInfo := TestMsgAddModelInfo(owner)
	result = setup.Handler(setup.Ctx, modelInfo)
	require.Equal(t, sdk.CodeOK, result.Code)

	// update existing model
	result = setup.Handler(setup.Ctx, MsgUpdateModelInfo{
		ID:                          test_constants.Id,
		NewName:                     "New Name",
		NewOwner:                    owner,
		NewDescription:              "New Description",
		NewSKU:                      test_constants.Sku,
		NewFirmwareVersion:          test_constants.FirmwareVersion,
		NewHardwareVersion:          test_constants.HardwareVersion,
		NewCertificateID:            test_constants.CertificateID,
		NewCertifiedDate:            test_constants.CertifiedDate,
		NewTisOrTrpTestingCompleted: false,
		Signer:                      owner,
	})
	require.Equal(t, sdk.CodeOK, result.Code)
}

func TestHandler_HandleDeleteModel(t *testing.T) {
	setup := Setup()

	// try to delete not present model
	result := setup.Handler(setup.Ctx, MsgDeleteModelInfo{
		ID:     test_constants.Id,
		Signer: test_constants.Owner,
	})
	require.Equal(t, types.CodeModelInfoDoesNotExist, result.Code)
}
