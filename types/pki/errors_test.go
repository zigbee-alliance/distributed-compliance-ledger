package types_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
)

// TestErrorConstructors exercises the error constructors that build the module's
// typed error surface so each wrapped sentinel is produced at least once.
func TestErrorConstructors(t *testing.T) {
	errs := []error{
		pkitypes.NewErrRevokedCertificateDoesNotExist("subj", "skid"),
		pkitypes.NewErrProvidedNotNocCertButExistingNoc("subj", "skid"),
		pkitypes.NewErrProvidedNotNocCertButRootIsNoc(),
		pkitypes.NewErrCRLSignerCertificatePidNotEqualRevocationPointPid(1, 2),
		pkitypes.NewErrInvalidAuthorityKeyIDFormat(),
		pkitypes.NewErrVidNotFound("vid"),
		pkitypes.NewErrPemValuesNotEqual("subj", "skid"),
		pkitypes.NewErrUnsupportedOperation("op"),
		pkitypes.NewErrInvalidVidFormat("bad"),
		pkitypes.NewErrInvalidPidFormat("bad"),
	}
	for _, err := range errs {
		require.Error(t, err)
		require.NotEmpty(t, err.Error())
		require.NotNil(t, errors.Unwrap(err))
	}
}
