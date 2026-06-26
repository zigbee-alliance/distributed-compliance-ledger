package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

func TestErrorConstructors(t *testing.T) {
	errs := []error{
		types.NewErrVendorProductsDoNotExist(1),
		types.NewErrSoftwareVersionStringInvalid("v"),
		types.NewErrFirmwareInformationInvalid("fw"),
		types.NewErrCDVersionNumberInvalid(1),
		types.NewErrOtaURLInvalid("url"),
		types.NewErrReleaseNotesURLInvalid("url"),
		types.NewErrNoModelVersionsExist(1, 2),
		types.NewErrModelVersionAlreadyExists(1, 2, 3),
		types.NewErrorOtaURLNotProvidedButOtherOtaFieldsProvided(),
	}
	for _, err := range errs {
		require.Error(t, err)
		require.NotEmpty(t, err.Error())
	}
}
